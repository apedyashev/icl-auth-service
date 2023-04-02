package grpc

import (
	"context"
	"errors"
	"icl-auth/pkg/adapter/grpc/pb"
	"icl-auth/pkg/adapter/repository"
	"icl-auth/pkg/domain/service"
	"icl-auth/pkg/model"
	"icl-auth/pkg/usecase"
	"log"

	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type userServiceServer struct {
	pb.UnimplementedUserServiceServer
	authUsecase usecase.Auth
}

func NewUserServiceServer(db *gorm.DB) *userServiceServer {
	userService := service.NewUserService(
		repository.NewUserRepository(db),
	)
	return &userServiceServer{
		authUsecase: usecase.NewAuth(userService),
	}
}

func (s *userServiceServer) GetByCredentials(ctx context.Context, cred *pb.GetByCredentialsRequest) (*pb.User, error) {
	u, err := s.authUsecase.Login(&usecase.LogingDTO{
		Email:    cred.Email,
		Password: cred.Password,
	})
	if err != nil {
		log.Println("Unable to login", err)
		return nil, errors.New("Invalid credentials")
	}
	return &pb.User{
		ID:        uint64(u.ID),
		Name:      u.Name,
		Username:  u.Username,
		Email:     u.Email,
		CreatedAt: timestamppb.New(u.CreatedAt),
		UpdatedAt: timestamppb.New(u.UpdatedAt),
	}, nil
}
func (s *userServiceServer) Create(ctx context.Context, userReq *pb.CreateRequest) (*pb.User, error) {

	user := &model.User{
		Name:     userReq.Name,
		Username: userReq.Username,
		Email:    userReq.Email,
		Password: userReq.Password,
	}
	user, err := s.authUsecase.Register(user)
	if err != nil {
		log.Println("failed to register user", err)
		return nil, errors.New("failed to register user")
	}

	return &pb.User{
		ID:        uint64(user.ID),
		Name:      user.Name,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}, nil
}

func (s *userServiceServer) GetById(ctx context.Context, req *pb.GetByIdRequest) (*pb.User, error) {
	u, err := s.authUsecase.UserById(uint(req.ID))
	if err != nil {
		log.Printf("Unable to find user %d %+v\n", req.ID, err)
		return nil, errors.New("Invalid credentials")
	}
	return &pb.User{
		ID:        uint64(u.ID),
		Name:      u.Name,
		Username:  u.Username,
		Email:     u.Email,
		CreatedAt: timestamppb.New(u.CreatedAt),
		UpdatedAt: timestamppb.New(u.UpdatedAt),
	}, nil
}
