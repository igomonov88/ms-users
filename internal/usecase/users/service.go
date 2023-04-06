package users

import (
	"context"

	userspb "github.com/igomonov88/crea-genproto-go/creaapis/ms-users/v1"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/igomonov88/ms-users/internal/operation"
)

type Service struct {
	op     *operation.Service
	logger *zap.SugaredLogger
}

func NewService(op *operation.Service, logger *zap.SugaredLogger) *Service {
	return &Service{
		op:     op,
		logger: logger,
	}
}

func (s *Service) CreateUser(ctx context.Context, req *userspb.User, resp *userspb.User) error {
	resp.UserId = 123
	resp.Name = req.Name
	s.logger.Info("CreateUser method called")
	return nil
}

func (s *Service) GetUser(ctx context.Context, req *userspb.GetUserRequest, resp *userspb.User) error {
	s.logger.Info("GetUser method called")
	return nil
}

func (s *Service) UpdateUser(ctx context.Context, req *userspb.User, resp *userspb.User) error {
	s.logger.Info("UpdateUser method called")
	return nil
}

func (s *Service) DeleteUser(ctx context.Context, req *userspb.DeleteUserRequest, resp *emptypb.Empty) error {
	s.logger.Info("DeleteUser method called")
	return nil
}
