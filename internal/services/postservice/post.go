package postservice

import (
	"errors"
	"github.com/Cococtel/Cococtel_Gagateway/internal/domain/dtos"
	"github.com/Cococtel/Cococtel_Gagateway/internal/domain/entities"
	"github.com/Cococtel/Cococtel_Gagateway/internal/repository/postrepository"
	"github.com/Cococtel/Cococtel_Gagateway/internal/utils"
	"log"
)

type PostsService interface {
	GetPosts() ([]entities.Post, utils.ApiError)
	GetPostByID(id string) (*entities.Post, utils.ApiError)
	CreatePost(post dtos.Post) (*entities.Post, utils.ApiError)
	UpdatePost(id string, updates map[string]interface{}) (*entities.Post, utils.ApiError)
	DeletePost(id string) utils.ApiError
}

type postsService struct {
	repo postrepository.IPost
}

func NewPostsService(repo postrepository.IPost) PostsService {
	return &postsService{repo: repo}
}

func (s *postsService) GetPosts() ([]entities.Post, utils.ApiError) {
	posts, err := s.repo.FetchPosts()
	if err != nil {
		log.Println(err)
		return nil, utils.NewApiError(errors.New("error fetching posts"), err.Status())
	}
	return posts, nil
}

func (s *postsService) GetPostByID(id string) (*entities.Post, utils.ApiError) {
	post, err := s.repo.FetchPostByID(id)
	if err != nil {
		log.Println(err)
		return nil, utils.NewApiError(errors.New("post not found"), err.Status())
	}
	return post, nil
}

func (s *postsService) CreatePost(post dtos.Post) (*entities.Post, utils.ApiError) {
	newPost, err := s.repo.CreatePost(post)
	if err != nil {
		log.Println(err)
		return nil, utils.NewApiError(errors.New("error creating post"), err.Status())
	}
	return newPost, nil
}

func (s *postsService) UpdatePost(id string, updates map[string]interface{}) (*entities.Post, utils.ApiError) {
	updatedPost, err := s.repo.UpdatePost(id, updates)
	if err != nil {
		log.Println(err)
		return nil, utils.NewApiError(errors.New("error updating post"), err.Status())
	}
	return updatedPost, nil
}

func (s *postsService) DeletePost(id string) utils.ApiError {
	err := s.repo.DeletePost(id)
	if err != nil {
		log.Println(err)
		return utils.NewApiError(errors.New("error deleting post"), err.Status())
	}
	return nil
}
