package email

import (
	"context"

	"github.com/Priyokumar/atekpa-email-grpc/api/pb"
)

type Service struct {
	pb.UnimplementedEmailServiceServer
}

func New() *Service {
	return &Service{}
}

func (s *Service) SendOrderConfiramtionEmail(ctx context.Context, req *pb.OrderConfirmationRequest) (*pb.Response, error) {
	return &pb.Response{Message: "success", Error: false}, nil
}
