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
		Login(credentails dtos.Login) (*entities.SuccessfulLogin, utils.ApiError)
		GetUser(id string, token string) (*entities.User, utils.ApiError)
		EditProfile(user dtos.User, token string) utils.ApiError
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

func (s *authService) Login(credentails dtos.Login) (*entities.SuccessfulLogin, utils.ApiError) {
	loginResponse, err := s.authRepo.Login(credentails)
	if err != nil {
		return nil, utils.NewApiError(errors.New("login error"), http.StatusInternalServerError)
	}
	return loginResponse, nil
}

func (s *authService) GetUser(id string, token string) (*entities.User, utils.ApiError) {
	user, err := s.authRepo.GetUser(id, token)
	user.UserID = id
	if err != nil {
		return nil, utils.NewApiError(errors.New("getting user error"), http.StatusInternalServerError)
	}
	return user, nil
}

func (s *authService) EditProfile(user dtos.User, token string) utils.ApiError {
	err := s.authRepo.EditUser(user, token)
	if err != nil {
		return utils.NewApiError(errors.New("editing user error"), http.StatusInternalServerError)
	}
	return nil
}
