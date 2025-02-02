package authservice

import (
	"errors"
	"github.com/Cococtel/Cococtel_Gagateway/internal/domain/dtos"
	"github.com/Cococtel/Cococtel_Gagateway/internal/domain/entities"
	"github.com/Cococtel/Cococtel_Gagateway/internal/repository/authrepository"
	"github.com/Cococtel/Cococtel_Gagateway/internal/utils"
	"net/http"
)

type (
	IAuth interface {
		Verify(token string) utils.ApiError
		Register(user dtos.Register) (*entities.User, utils.ApiError)
		Login(credentials dtos.Login) (*entities.SuccessfulLogin, utils.ApiError)
	}
	authService struct {
		authRepo authrepository.IAuth
	}
)

func NewAuthService(repo authrepository.IAuth) IAuth {
	return &authService{authRepo: repo}
}

func (s *authService) Verify(token string) utils.ApiError {
	err := s.authRepo.Verify(token)
	if err != nil {
		return utils.NewApiError(errors.New("unauthorized"), http.StatusUnauthorized)
	}
	return nil
}

func (s *authService) Register(user dtos.Register) (*entities.User, utils.ApiError) {
	newUser, err := s.authRepo.Register(user)
	if err != nil {
		return nil, utils.NewApiError(errors.New("register error"), http.StatusInternalServerError)
	}
	return newUser, nil
}

func (s *authService) Login(credentials dtos.Login) (*entities.SuccessfulLogin, utils.ApiError) {
	loginResponse, err := s.authRepo.Login(credentials)
	if err != nil {
		return nil, utils.NewApiError(errors.New("login error"), http.StatusInternalServerError)
	}
	return loginResponse, nil
}
