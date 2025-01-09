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
}

func NewUserService(user domain.RepositoryUser, log *logger.Logger) *UserService {
	return &UserService{
		user: user,
		log:  log,
	}
}

func (u *UserService) Register(ctx context.Context, req *entities.CreateUserRequest) *customerrors.CustomError {
	user, err := u.user.FindByUsername(ctx, req.UserName)
	if err != nil {
		u.log.Error("database error during user lookup", err)
		return customerrors.ErrDB
	}
	if user != nil {
		u.log.Warn("User already exists")
		return customerrors.ErrUserExists
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		u.log.Error("error hash password", err)
		return customerrors.ErrHashPassword
	}
	err = u.user.Create(ctx, &domain.Users{
		Id:          utils.NewUUID().GenUUID(),
		UserName:    strings.TrimSpace(req.UserName),
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

func (u *UserService) Login(ctx context.Context, user_name, password string) *entities.LoginResponse {
	return &entities.LoginResponse{}
}
