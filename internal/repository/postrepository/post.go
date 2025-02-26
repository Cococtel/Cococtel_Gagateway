package postrepository

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Cococtel/Cococtel_Gagateway/internal/domain/dtos"
	"github.com/Cococtel/Cococtel_Gagateway/internal/domain/entities"
	"net/http"
	"os"
)

type IPost interface {
	FetchPosts() ([]entities.Post, error)
	FetchPostByID(id string) (*entities.Post, error)
	CreatePost(post dtos.Post) (*entities.Post, error)
	UpdatePost(id string, updates map[string]interface{}) (*entities.Post, error)
	DeletePost(id string) error
}

type postRepository struct{}

var ms_posts_endpoint string

func NewCatalogRepository() IPost {
	ms_posts_endpoint = os.Getenv("MS_CATALOG_DOMAIN")
	return &postRepository{}
}

// FetchPosts obtiene todos los posts.
func (r *postRepository) FetchPosts() ([]entities.Post, error) {
	resp, err := http.Get(ms_posts_endpoint + "/posts")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var posts []entities.Post
	if err := json.NewDecoder(resp.Body).Decode(&posts); err != nil {
		return nil, err
	}
	return posts, nil
}

// FetchPostByID obtiene un post por ID.
func (r *postRepository) FetchPostByID(id string) (*entities.Post, error) {
	url := fmt.Sprintf("%s/posts/%s", ms_posts_endpoint, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var post entities.Post
	if err := json.NewDecoder(resp.Body).Decode(&post); err != nil {
		return nil, err
	}
	if post.ID == "" {
		return nil, errors.New("post not found")
	}
	return &post, nil
}

// CreatePost crea un nuevo post.
func (r *postRepository) CreatePost(post dtos.Post) (*entities.Post, error) {
	body, err := json.Marshal(post)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(ms_posts_endpoint+"/posts", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var newPost entities.Post
	if err := json.NewDecoder(resp.Body).Decode(&newPost); err != nil {
		return nil, err
	}
	return &newPost, nil
}

// UpdatePost actualiza un post existente.
func (r *postRepository) UpdatePost(id string, updates map[string]interface{}) (*entities.Post, error) {
	url := fmt.Sprintf("%s/posts/%s", ms_posts_endpoint, id)
	body, err := json.Marshal(updates)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var updatedPost entities.Post
	if err := json.NewDecoder(resp.Body).Decode(&updatedPost); err != nil {
		return nil, err
	}
	return &updatedPost, nil
}

// DeletePost elimina un post por ID.
func (r *postRepository) DeletePost(id string) error {
	url := fmt.Sprintf("%s/posts/%s", ms_posts_endpoint, id)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
