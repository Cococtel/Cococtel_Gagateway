package authrepository

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Cococtel/Cococtel_Gagateway/internal/domain/dtos"
	"github.com/Cococtel/Cococtel_Gagateway/internal/domain/entities"
	"io/ioutil"
	"net/http"
	"os"
)

type (
	IAuth interface {
		Verify(token string) error
		Register(user dtos.Register) (*entities.User, error)
		Login(credentails dtos.Login) (*entities.SuccessfulLogin, error)
	}
	authRepository struct{}
)

var ms_authorization_endpoint string

func NewAuthRepository() IAuth {
	ms_authorization_endpoint = os.Getenv("MS_AUTH_DOMAIN")
	return &authRepository{}
}

func (r *authRepository) Verify(token string) error {
	url := fmt.Sprintf("%s/v1/verify", ms_authorization_endpoint)
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("x-auth-token", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("unauthorized")
	}

	return nil
}

func (r *authRepository) Register(user dtos.Register) (*entities.User, error) {
	url := fmt.Sprintf("%s/register", ms_authorization_endpoint)
	body, _ := json.Marshal(user)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var userResp entities.UserResponse
	if err := json.Unmarshal(responseBody, &userResp); err != nil {
		return nil, err
	}

	return &userResp.Data, nil
}

func (r *authRepository) Login(credentails dtos.Login) (*entities.SuccessfulLogin, error) {
	url := fmt.Sprintf("%s/login", ms_authorization_endpoint)
	body, _ := json.Marshal(credentails)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var loginResp entities.LoginResponse
	if err := json.Unmarshal(responseBody, &loginResp); err != nil {
		return nil, err
	}

	return &loginResp.Data, nil
}
