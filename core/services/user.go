package services

import (
	"context"
	"demo_time_sheet_server/apis/entities"
	"demo_time_sheet_server/common/logger"
	"demo_time_sheet_server/common/utils"
	"demo_time_sheet_server/core/adapters/cache"
	customerrors "demo_time_sheet_server/core/custom_errors"
	"demo_time_sheet_server/core/domain"
	"fmt"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	user  domain.RepositoryUser
	log   *logger.Logger
	jwt   *JwtService
	cache cache.CacheOperations
	trans domain.RepositoryTransactionHelper
}

func NewUserService(user domain.RepositoryUser,
	log *logger.Logger,
	jwt *JwtService,
	cache cache.CacheOperations,
	trans domain.RepositoryTransactionHelper,
) *UserService {
	return &UserService{
		user:  user,
		log:   log,
		jwt:   jwt,
		cache: cache,
		trans: trans,
	}
}

func (u *UserService) Register(ctx context.Context, req *entities.CreateUserRequest) *customerrors.CustomError {
	userNameTrSp := strings.TrimSpace(req.UserName)
	userID := utils.GenUUID()
	user, err := u.user.FindByUsername(ctx, userNameTrSp)
	if err != nil {
		u.log.Error("database error during user lookup", err)
		return customerrors.ErrDB
	}
	if user != nil {
		u.log.Warn("User already exists")
		return customerrors.ErrUserExists
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(strings.TrimSpace(req.Password)), bcrypt.DefaultCost)
	if err != nil {
		u.log.Error("error hash password", err)
		return customerrors.ErrHashPassword
	}

	model := &domain.Users{
		Id:          userID,
		UserName:    userNameTrSp,
		Password:    string(passwordHash),
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Avatar:      req.Avatar,
		CreatedAt:   utils.GenTime(),
		UpdatedAt:   utils.GenTime(),
	}

	if err := u.trans.Transaction(ctx, func(ctx context.Context, db *gorm.DB) error {
		err = u.user.Create(ctx, db, model)
		if err != nil {
			u.log.Error("Failed to create user", err)
			return customerrors.ErrDB
		}

		key := fmt.Sprintf("user:%v", userID)
		err = u.cache.Set(ctx, key, model, 0)
		if err != nil {
			u.log.Error("Failed add info after to create user", err)
			return customerrors.ErrDB
		}

		return nil
	}); err != nil {
		return customerrors.ErrDB
	}

	return nil
}

func (u *UserService) Login(ctx context.Context, user_name, password string) (*entities.LoginResponse, *customerrors.CustomError) {
	user, err := u.user.FindByUsername(ctx, user_name)
	if err != nil {
		return nil, customerrors.ErrDB
	}
	if user == nil {
		return nil, customerrors.ErrNotFound
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, customerrors.ErrNotFound
	}
	genToken, err := u.jwt.GenToken(ctx, user.UserName, user.Id, user.UpdatedAt)
	if err != nil {
		return nil, customerrors.ErrAuth
	}

	return &entities.LoginResponse{
		Token:     genToken.Token,
		CreatedAt: utils.GenTime().UTC(),
	}, nil
}

func (u *UserService) Profile(ctx context.Context, userID int64) (*entities.GetProfile, *customerrors.CustomError) {
	key := fmt.Sprintf("user:%v", userID)
	var user domain.Users
	err := u.cache.Get(ctx, key, &user)
	if err != nil {
		u.log.Error("error get data from cache", err)
		return nil, customerrors.ErrDB
	}
	if user.UserName == "" {
		return nil, customerrors.ErrNotFound
	}

	return &entities.GetProfile{
		Id:          userID,
		UserName:    user.UserName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}, nil
}

func (u *UserService) LoginWithGG(ctx context.Context, tokenFromGG string) (*entities.LoginResponse, *customerrors.CustomError) {
	var userID int64
	var token string
	var updateTime time.Time
	var passWord string
	inforUser, err := utils.VerifyGoogleToken(tokenFromGG)
	if err != nil {
		return nil, customerrors.ErrVerifyTokenEmail
	}

	genPassWord := utils.GenPasswordString(8)
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(genPassWord), bcrypt.DefaultCost)
	if err != nil {
		u.log.Error("error hash password", err)
		return nil, customerrors.ErrHashPassword
	}

	if err := u.trans.Transaction(ctx, func(ctx context.Context, db *gorm.DB) error {
		user, err := u.user.GetUserByGoogleUserID(ctx, inforUser.Sub)
		if err != nil {
			u.log.Error("Failed to fetch user", err)
			return customerrors.ErrDB
		}
		if user == nil {
			userID = utils.GenUUID()
			updateTime = utils.GenTime()
			passWord = string(passwordHash)
		} else {
			userID = user.Id
			updateTime = user.UpdatedAt
			passWord = user.Password

			subject := "Gửi bạn tài khoản và mật khẩu đăng nhập"
			body := fmt.Sprintf(
				"Chào %v,\n\nChúng tôi đã tạo tài khoản cho bạn. Bạn có thể đăng nhập với tài khoản và mật khẩu sau:\n\nTài khoản: %v\nMật khẩu: %v\n\nChúc bạn sử dụng dịch vụ vui vẻ!",
				inforUser.Name,
				inforUser.Name,
				genPassWord,
			)

			err = utils.SendEmail(inforUser.Email, subject, body)
			if err != nil {
				u.log.Error("Failed to send email", err)
				return customerrors.ErrorSendEmail
			}
		}

		model := &domain.Users{
			Id:           userID,
			UserName:     inforUser.Name,
			Password:     passWord,
			GoogleUserId: inforUser.Sub,
			Email:        inforUser.Email,
			Avatar:       inforUser.Picture,
			CreatedAt:    utils.GenTime(),
			UpdatedAt:    updateTime,
		}

		err = u.user.Create(ctx, db, model)
		if err != nil {
			u.log.Error("Failed to create user", err)
			return customerrors.ErrDB
		}

		key := fmt.Sprintf("user:%v", userID)
		err = u.cache.Set(ctx, key, model, 0)
		if err != nil {
			u.log.Error("Failed add info after to create user", err)
			return customerrors.ErrDB
		}

		genToken, err := u.jwt.GenToken(ctx, inforUser.Name, userID, updateTime)
		if err != nil {
			return customerrors.ErrAuth
		}
		token = genToken.Token

		return nil

	}); err != nil {
		return nil, customerrors.ErrDB
	}

	return &entities.LoginResponse{
		Token:     token,
		CreatedAt: utils.GenTime(),
	}, nil
}
