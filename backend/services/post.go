package services

import (
	"github.com/nambuitechx/go-social/models"
	"github.com/nambuitechx/go-social/repositories"
)

type PostService struct {
	PostRepository *repositories.PostRepository
}

func NewPostService(postRepository *repositories.PostRepository) *PostService {
	return &PostService{ PostRepository: postRepository }
}

func (s *PostService) Health() string {
	return "Post service is available"
}

func (s *PostService) GetAllPosts(limit *int, offset *int) []models.PostModel {
	posts := s.PostRepository.SelectPosts(limit, offset)
	return posts
}

func (s *PostService) GetPostById(id *string) (*models.PostModel, error) {
	post, err := s.PostRepository.SelectPostById(id)
	return post, err
}

func (s *PostService) CreatePost(payload *models.CreatePostPayload, userId string) (*models.PostModel, error) {
	post, err := s.PostRepository.InsertPost(payload, userId)
	return post, err
}

func (s *PostService) DeletePostById(id *string) error {
	err := s.PostRepository.DeletePostById(id)
	return err
}
