package grpc_schemas

import (
	"context"
	"fmt"
	taxi_order_service "github.com/roman-korostelev/grpc-schemas/grpc"
	"github.com/sirupsen/logrus"
	"time"
)

type Driver interface {
	Register(context.Context, *logrus.Logger, string, string, string, string, taxi_order_service.TaxiType) (string, error)
	Login(context.Context, *logrus.Logger, string, string) (string, error)
	GetRating(context.Context, *logrus.Logger, string) (float32, error)
	ChangeStatus(context.Context, *logrus.Logger, string) error
}

func FindOrders(logger *logrus.Logger, client taxi_order_service.RouteGuideClient, in *taxi_order_service.FindOrdersReq) (*taxi_order_service.Orders, error) {
	logger.Infof("Getting orders of driver with id : %v", in.DriverId)
	fmt.Println("dsa")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	orders, err := client.FindOrders(ctx, in)
	if err != nil {
		logger.Infoln("Getting orders failed")
	}
	return orders, nil
}
