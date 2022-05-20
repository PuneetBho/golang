package services

import (
	"context"
	"fmt"
	"go-mongo/model"
	sv "go-mongo/proto"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) AddPlayer(ctx context.Context, in *sv.AddPlayerRequest) (*sv.AddPlayerResponse, error) {
	log.Println("Invoking AddPlayer")

	player := model.Player{
		PlayerName: in.PlayerName,
		Position:   in.Position,
		Salary:     in.Salary,
		IsEnabled:  in.Enabled,
	}

	err := s.repo.AddPlayer(ctx, player)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error %v\n", err),
		)
	}

	return &sv.AddPlayerResponse{Message: "Player Added"}, nil
}

func (s *Service) GetPlayers(ctx context.Context, in *sv.GetPlayersRequest) (*sv.GetPlayersResponse, error) {
	log.Println("Invoking GetPlayers")
	log.Println(in.Enabled)
	players, err := s.repo.GetPlayers(ctx, in.Position, in.Salary, in.Enabled, in.Sortby)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error %v\n", err),
		)
	}

	playerList := make([]*sv.GetPlayersResponse_Player, 0)
	for _, player := range *players {
		playerList = append(playerList, &sv.GetPlayersResponse_Player{
			Id:         player.ID.Hex(),
			PlayerName: player.PlayerName,
			Position:   player.Position,
			Salary:     player.Salary,
			Enabled:    player.IsEnabled,
		})
	}

	return &sv.GetPlayersResponse{Players: playerList}, nil
}
