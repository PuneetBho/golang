package services

import (
	sv "go-mongo/proto"
	"go-mongo/repository"
)

type Service struct {
	sv.PlayerServiceServer
	repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}
