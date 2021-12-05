package controller

import (
	message "A/protobuffer/proto"
	"context"
	"fmt"
)

type MessageController struct {
	message.UnimplementedGetMessageServer
}

func (m MessageController) Get(ctx context.Context, input *message.ARequest) (*message.AResponse, error) {
	return &message.AResponse{
		Result: "Succeed for input: " + input.Name,
	}, nil
}

func (m MessageController) GetStreaming(input *message.ARequests, stream message.GetMessage_GetStreamingServer) error {
	fmt.Println("Receive requests: ", input.Names)
	for _, name := range input.Names {
		if err := stream.Send(&message.AResponse{
			Result: "Received name: " + name,
		}); err != nil {
			return err
		}
	}

	return nil
}
