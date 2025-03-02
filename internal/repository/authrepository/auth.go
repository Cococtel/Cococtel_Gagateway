package authrepository

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Cococtel/Cococtel_Gagateway/internal/domain/dtos"
	"github.com/Cococtel/Cococtel_Gagateway/internal/domain/entities"
	"github.com/Cococtel/Cococtel_Gagateway/internal/utils"
)

type (
	IAuth interface {
		Verify(token string) utils.ApiError
		Register(user dtos.Register) (*entities.User, utils.ApiError)
		Login(credentails dtos.Login) (*entities.SuccessfulLogin, utils.ApiError)
		GetUser(id string, token string) (*entities.User, utils.ApiError)
		EditUser(user dtos.User, token string) utils.ApiError
	}
	authRepository struct{}
)

var ms_authorization_endpoint string

func NewAuthRepository() IAuth {
	ms_authorization_endpoint = os.Getenv("MS_AUTH_DOMAIN")
	return &authRepository{}
}

func (r *authRepository) Verify(token string) utils.ApiError {
	start := time.Now()
	url := fmt.Sprintf("%s/v1/verify", ms_authorization_endpoint)
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		log.Println(err)
		utils.MeasureRequest("auth", "Verify", start, http.StatusBadRequest, err)
		return utils.NewApiError(err, http.StatusBadRequest)
	}
	req.Header.Set("x-auth-token", token)
	client := &http.Client{}
	resp, err := client.Do(req)

	if resp.StatusCode != http.StatusOK {
		err = errors.New("unauthorized")
		utils.MeasureRequest("auth", "Verify", start, resp.StatusCode, err)
		return utils.NewApiError(err, resp.StatusCode)
	}

	utils.MeasureRequest("auth", "Verify", start, resp.StatusCode, nil)
	return nil
}

func (r *authRepository) Register(user dtos.Register) (*entities.User, utils.ApiError) {
	start := time.Now()
	url := fmt.Sprintf("%s/register", ms_authorization_endpoint)
	body, _ := json.Marshal(user)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Println(err)
		utils.MeasureRequest("auth", "Register", start, 0, err)
		return nil, utils.NewApiError(err, http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode > http.StatusMultipleChoices {
		log.Println(err)
		utils.MeasureRequest("auth", "Register", start, resp.StatusCode, errors.New(string(responseBody)))
		return nil, utils.NewApiError(errors.New(string(responseBody)), resp.StatusCode)
	}

	var userResp entities.UserResponse
	if err := json.Unmarshal(responseBody, &userResp); err != nil {
		utils.MeasureRequest("auth", "Register", start, resp.StatusCode, err)
		return nil, utils.NewApiError(err, http.StatusInternalServerError)
	}

	utils.MeasureRequest("auth", "Register", start, resp.StatusCode, nil)
	return &userResp.Data, nil
}

func (r *authRepository) Login(credentails dtos.Login) (*entities.SuccessfulLogin, utils.ApiError) {
	start := time.Now()
	url := fmt.Sprintf("%s/login", ms_authorization_endpoint)
	body, _ := json.Marshal(credentails)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Println(err)
		utils.MeasureRequest("auth", "Login", start, 0, err)
		return nil, utils.NewApiError(err, http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode > http.StatusMultipleChoices {
		log.Println(responseBody)
		utils.MeasureRequest("auth", "Login", start, resp.StatusCode, errors.New(string(responseBody)))
		return nil, utils.NewApiError(errors.New(string(responseBody)), resp.StatusCode)
	}

	var loginResp entities.LoginResponse
	if err := json.Unmarshal(responseBody, &loginResp); err != nil {
		utils.MeasureRequest("auth", "Login", start, resp.StatusCode, err)
		return nil, utils.NewApiError(err, http.StatusInternalServerError)
	}

	utils.MeasureRequest("auth", "Login", start, resp.StatusCode, nil)
	return &loginResp.Data, nil
}

func (r *authRepository) GetUser(id string, token string) (*entities.User, utils.ApiError) {
	start := time.Now()
	url := fmt.Sprintf("%s/v1/profile/%s", ms_authorization_endpoint, id)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		utils.MeasureRequest("auth", "GetUser", start, http.StatusBadRequest, err)
		return nil, utils.NewApiError(err, http.StatusBadRequest)
	}
	req.Header.Set("x-auth-token", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		utils.MeasureRequest("auth", "GetUser", start, 0, err)
		return nil, utils.NewApiError(err, resp.StatusCode)
	}
	defer resp.Body.Close()

	var response struct {
		Data entities.User `json:"data"`
	}

	if resp.StatusCode > http.StatusMultipleChoices {
		log.Println("error getting user")
		utils.MeasureRequest("auth", "GetUser", start, resp.StatusCode, errors.New("error getting user"))
		return nil, utils.NewApiError(errors.New("error getting user"), resp.StatusCode)
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		utils.MeasureRequest("auth", "GetUser", start, resp.StatusCode, err)
		return nil, utils.NewApiError(err, http.StatusInternalServerError)
	}

	if response.Data.Name == "" {
		err := errors.New("user not found")
		utils.MeasureRequest("auth", "GetUser", start, resp.StatusCode, err)
		return nil, utils.NewApiError(err, http.StatusNotFound)
	}

	utils.MeasureRequest("auth", "GetUser", start, resp.StatusCode, nil)
	return &response.Data, nil
}

func (r *authRepository) EditUser(user dtos.User, token string) utils.ApiError {
	start := time.Now()
	url := fmt.Sprintf("%s/v1/profile", ms_authorization_endpoint)
	body, _ := json.Marshal(user)
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
	if err != nil {
		log.Println(err)
		utils.MeasureRequest("auth", "EditUser", start, http.StatusBadRequest, err)
		return utils.NewApiError(err, http.StatusBadRequest)
	}
	req.Header.Set("x-auth-token", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		utils.MeasureRequest("auth", "EditUser", start, 0, err)
		return utils.NewApiError(err, http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	if resp.StatusCode > http.StatusMultipleChoices {
		err := errors.New(resp.Status)
		utils.MeasureRequest("auth", "EditUser", start, resp.StatusCode, errors.New(string(body)))
		return utils.NewApiError(err, resp.StatusCode)
	}

	utils.MeasureRequest("auth", "EditUser", start, resp.StatusCode, nil)
	return nil
}
