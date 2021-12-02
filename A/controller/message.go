package controller

import (
	"context"

	message "A/protobuffer/proto"
)

type MessageController struct {
	message.UnimplementedGetMessageServer
}

func (m MessageController) Get(ctx context.Context, input *message.ARequest) (*message.AResponse, error) {
	return &message.AResponse{
		Result: "Succeed for input: " + input.Name,
	}, nil
}
