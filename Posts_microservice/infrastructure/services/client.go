package services

import (
	"github.com/stojic19/XWS-TIM15/common/proto/followers"
	"github.com/stojic19/XWS-TIM15/common/proto/notifications"
	"github.com/stojic19/XWS-TIM15/common/proto/users"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewFollowersClient(address string) followers.FollowersServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Followers service: %v", err)
	}
	return followers.NewFollowersServiceClient(conn)
}

func NewUsersClient(address string) users.UsersServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Users service: %v", err)
	}
	return users.NewUsersServiceClient(conn)
}

func NewNotificationsClient(address string) notifications.NotificationsServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Users service: %v", err)
	}
	return notifications.NewNotificationsServiceClient(conn)
}

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
