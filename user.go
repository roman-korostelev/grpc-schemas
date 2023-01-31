package grpc_schemas

import (
	"context"
	taxi_order_service "github.com/roman-korostelev/grpc-schemas/grpc"
	"log"
	"time"
)

type User interface {
	Register(context.Context, string, string, string, string) (string, error)
	Login(context.Context, string, string) (string, error)
	DeleteProfile(context.Context, string) error
	DisableToken(context.Context, string) error
	VerifyToken(context.Context, string) (bool, error)
}

func FindUserOrders(client taxi_order_service.RouteGuideClient, in *taxi_order_service.FindOrdersReq) (*taxi_order_service.Orders, error) {
	log.Printf("Getting orders of driver with id : %v", in.UserId)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	orders, err := client.FindOrders(ctx, in)
	if err != nil {
		log.Println("Getting orders failed")
	}
	return orders, nil
}
