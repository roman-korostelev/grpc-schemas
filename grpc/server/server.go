package server

import (
	"context"
	grpc "github.com/roman-korostelev/grpc-schemas/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Repository interface {
	CreateOrder(context.Context, string, string, string, string, grpc.TaxiType) (string, error)
	FindOrderByFields(context.Context, string, string, string, string) ([]*grpc.Order, error)
}

type ImplementedRouteGuideServer struct {
	grpc.UnimplementedRouteGuideServer
	repo Repository
}

func NewRouteGuideServer(r Repository) *ImplementedRouteGuideServer {
	return &ImplementedRouteGuideServer{
		UnimplementedRouteGuideServer: grpc.UnimplementedRouteGuideServer{},
		repo:                          r,
	}
}

func (ImplementedRouteGuideServer) RateDriver(context.Context, *grpc.DriverRatingReq) (*grpc.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RateDriver not implemented")
}
func (ImplementedRouteGuideServer) RateUser(context.Context, *grpc.UserRatingReq) (*grpc.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RateUser not implemented")
}
func (ImplementedRouteGuideServer) CreateOrder(context.Context, *grpc.OrderReq) (*grpc.OrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (ImplementedRouteGuideServer) FindDriverForOrder(context.Context, *grpc.FindDriverReq) (*grpc.FindDriverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindDriverForOrder not implemented")
}

func (s ImplementedRouteGuideServer) FindOrders(ctx context.Context, req *grpc.FindOrdersReq) (*grpc.Orders, error) {
	orders, err := s.repo.FindOrderByFields(ctx, *req.UserId, *req.DriverId, *req.From, *req.To)
	if err != nil {
		return nil, err
	}
	ans := grpc.Orders{Order: orders}
	return &ans, err
}
