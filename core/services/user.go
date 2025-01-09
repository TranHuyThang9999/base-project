package services

import (
	"context"
	"rices/apis/entities"
	"rices/common/logger"
	"rices/common/utils"
	customerrors "rices/core/custom_errors"
	"rices/core/domain"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	user domain.RepositoryUser
	log  *logger.Logger
	jwt  *JwtService
}

func NewUserService(user domain.RepositoryUser, log *logger.Logger, jwt *JwtService) *UserService {
	return &UserService{
		user: user,
		log:  log,
		jwt:  jwt,
	}
}

func (u *UserService) Register(ctx context.Context, req *entities.CreateUserRequest) *customerrors.CustomError {
	userNameTrSp := strings.TrimSpace(req.UserName)
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

	err = u.user.Create(ctx, &domain.Users{
		Id:          utils.NewUUID().GenUUID(),
		UserName:    userNameTrSp,
		Password:    string(passwordHash),
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		CreatedAt:   utils.NewUUID().GenTime(),
		UpdatedAt:   utils.NewUUID().GenTime(),
	})
	if err != nil {
		u.log.Error("Failed to create user", err)
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
		CreatedAt: utils.NewUUID().GenTime().UTC(),
	}, nil
}

func (u *UserService) Profile(ctx context.Context, userID int64) (*entities.GetProfile, *customerrors.CustomError) {
	user, err := u.user.FindByID(ctx, userID)
	if err != nil {
		return nil, customerrors.ErrDB
	}
	if user == nil {
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
