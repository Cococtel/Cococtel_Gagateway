package postrepository

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Cococtel/Cococtel_Gagateway/internal/domain/dtos"
	"github.com/Cococtel/Cococtel_Gagateway/internal/domain/entities"
	"github.com/Cococtel/Cococtel_Gagateway/internal/utils"
)

type IPost interface {
	FetchPosts() ([]entities.Post, utils.ApiError)
	FetchPostByID(id string) (*entities.Post, utils.ApiError)
	CreatePost(post dtos.Post) (*entities.Post, utils.ApiError)
	UpdatePost(id string, updates map[string]interface{}) (*entities.Post, utils.ApiError)
	DeletePost(id string) utils.ApiError
}

type postRepository struct{}

var ms_posts_endpoint string

func NewPostRepository() IPost {
	ms_posts_endpoint = os.Getenv("MS_CATALOG_DOMAIN")
	return &postRepository{}
}

// FetchPosts obtiene todos los posts.
func (r *postRepository) FetchPosts() ([]entities.Post, utils.ApiError) {
	start := time.Now()
	resp, err := http.Get(ms_posts_endpoint + "/posts")
	if err != nil {
		log.Println(err)
		utils.MeasureRequest("posts", "FetchPosts", start, resp.StatusCode, err)
		return nil, utils.NewApiError(err, resp.StatusCode)
	}
	defer resp.Body.Close()

	var posts []entities.Post
	if err := json.NewDecoder(resp.Body).Decode(&posts); err != nil {
		utils.MeasureRequest("posts", "FetchPosts", start, resp.StatusCode, err)
		return nil, utils.NewApiError(err, resp.StatusCode)
	}

	if resp.StatusCode > http.StatusMultipleChoices {
		log.Println(resp.Body)
		utils.MeasureRequest("posts", "FetchPosts", start, resp.StatusCode, errors.New("status code "+strconv.Itoa(resp.StatusCode)))
		return nil, utils.NewApiError(errors.New("error getting posts"), resp.StatusCode)
	}

	utils.MeasureRequest("posts", "FetchPosts", start, resp.StatusCode, nil)
	return posts, nil
}

// FetchPostByID obtiene un post por ID.
func (r *postRepository) FetchPostByID(id string) (*entities.Post, utils.ApiError) {
	start := time.Now()
	url := fmt.Sprintf("%s/posts/%s", ms_posts_endpoint, id)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		utils.MeasureRequest("posts", "FetchPostByID", start, resp.StatusCode, err)
		return nil, utils.NewApiError(err, resp.StatusCode)
	}
	defer resp.Body.Close()

	var post entities.Post
	if err := json.NewDecoder(resp.Body).Decode(&post); err != nil {
		utils.MeasureRequest("posts", "FetchPostByID", start, resp.StatusCode, err)
		return nil, utils.NewApiError(err, resp.StatusCode)
	}
	if resp.StatusCode > http.StatusMultipleChoices {
		log.Println(resp.Body)
		utils.MeasureRequest("posts", "FetchPostByID", start, resp.StatusCode, errors.New("status code "+strconv.Itoa(resp.StatusCode)))
		return nil, utils.NewApiError(errors.New("post not found"), resp.StatusCode)
	}
	if post.ID == "" {
		err := errors.New("post not found")
		utils.MeasureRequest("posts", "FetchPostByID", start, http.StatusNotFound, err)
		return nil, utils.NewApiError(errors.New("post not found"), http.StatusNotFound)
	}

	utils.MeasureRequest("posts", "FetchPostByID", start, resp.StatusCode, nil)
	return &post, nil
}

// CreatePost crea un nuevo post.
func (r *postRepository) CreatePost(post dtos.Post) (*entities.Post, utils.ApiError) {
	start := time.Now()
	body, err := json.Marshal(post)
	if err != nil {
		log.Println(err)
		utils.MeasureRequest("posts", "CreatePost", start, http.StatusBadRequest, err)
		return nil, utils.NewApiError(err, http.StatusBadRequest)
	}
	resp, err := http.Post(ms_posts_endpoint+"/posts", "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Println(err)
		utils.MeasureRequest("posts", "CreatePost", start, resp.StatusCode, err)
		return nil, utils.NewApiError(err, resp.StatusCode)
	}
	defer resp.Body.Close()

	var newPost entities.Post
	if err := json.NewDecoder(resp.Body).Decode(&newPost); err != nil {
		utils.MeasureRequest("posts", "CreatePost", start, resp.StatusCode, err)
		return nil, utils.NewApiError(err, http.StatusInternalServerError)
	}

	if resp.StatusCode > http.StatusMultipleChoices {
		log.Println(resp.Body)
		utils.MeasureRequest("posts", "CreatePost", start, resp.StatusCode, errors.New("status code "+strconv.Itoa(resp.StatusCode)))
		return nil, utils.NewApiError(errors.New("post not created"), resp.StatusCode)
	}

	utils.MeasureRequest("posts", "CreatePost", start, resp.StatusCode, nil)
	return &newPost, nil
}

// UpdatePost actualiza un post existente.
func (r *postRepository) UpdatePost(id string, updates map[string]interface{}) (*entities.Post, utils.ApiError) {
	start := time.Now()
	url := fmt.Sprintf("%s/posts/%s", ms_posts_endpoint, id)
	body, err := json.Marshal(updates)
	if err != nil {
		log.Println(err)
		utils.MeasureRequest("posts", "UpdatePost", start, http.StatusBadRequest, err)
		return nil, utils.NewApiError(err, http.StatusBadRequest)
	}
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
	if err != nil {
		log.Println(err)
		utils.MeasureRequest("posts", "UpdatePost", start, http.StatusBadRequest, err)
		return nil, utils.NewApiError(err, http.StatusBadRequest)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		utils.MeasureRequest("posts", "UpdatePost", start, resp.StatusCode, err)
		return nil, utils.NewApiError(err, resp.StatusCode)
	}
	defer resp.Body.Close()

	var updatedPost entities.Post
	if err := json.NewDecoder(resp.Body).Decode(&updatedPost); err != nil {
		utils.MeasureRequest("posts", "UpdatePost", start, resp.StatusCode, err)
		return nil, utils.NewApiError(err, resp.StatusCode)
	}

	if resp.StatusCode > http.StatusMultipleChoices {
		log.Println(resp.Body)
		utils.MeasureRequest("posts", "UpdatePost", start, resp.StatusCode, errors.New("status code "+strconv.Itoa(resp.StatusCode)))
		return nil, utils.NewApiError(errors.New("post not updated"), resp.StatusCode)
	}

	utils.MeasureRequest("posts", "UpdatePost", start, resp.StatusCode, nil)
	return &updatedPost, nil
}

// DeletePost elimina un post por ID.
func (r *postRepository) DeletePost(id string) utils.ApiError {
	start := time.Now()
	url := fmt.Sprintf("%s/posts/%s", ms_posts_endpoint, id)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		log.Println(err)
		utils.MeasureRequest("posts", "DeletePost", start, http.StatusBadRequest, err)
		return utils.NewApiError(err, http.StatusBadRequest)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		utils.MeasureRequest("posts", "DeletePost", start, resp.StatusCode, err)
		return utils.NewApiError(err, resp.StatusCode)
	}
	defer resp.Body.Close()

	if resp.StatusCode > http.StatusMultipleChoices {
		log.Println(resp.Body)
		utils.MeasureRequest("posts", "DeletePost", start, resp.StatusCode, errors.New("status code "+strconv.Itoa(resp.StatusCode)))
		return utils.NewApiError(errors.New("post not deleted"), resp.StatusCode)
	}

	utils.MeasureRequest("posts", "DeletePost", start, resp.StatusCode, nil)
	return nil
}
