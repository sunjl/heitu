package rpc

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/golang/protobuf/ptypes/wrappers"
	pb "github.com/jishufan/heitu/common/protobuf"
	"github.com/jishufan/heitu/server/repository"
	"github.com/jishufan/heitu/server/util"
	"golang.org/x/net/context"
)

type UserServiceServer struct{}

func convertToModelPageRequest(pageRequest *pb.PageRequest) repository.PageRequest {
	return repository.PageRequest{
		Page:      int(pageRequest.Page),
		Size:      int(pageRequest.Size),
		Order:     pageRequest.Order,
		Direction: pageRequest.Direction,
	}
}

func convertToModelUser(user *pb.User) repository.User {
	return repository.User{
		Id:          int(user.Id),
		Username:    user.Username,
		Password:    user.Password,
		DisplayName: user.DisplayName,
		Phone:       user.Phone,
		UserType:    user.UserType,
		Token:       user.Token,
	}
}

func convertToPBUser(user repository.User) *pb.User {
	return &pb.User{
		Id:          int32(user.Id),
		Username:    user.Username,
		Password:    user.Password,
		DisplayName: user.DisplayName,
		Phone:       user.Phone,
		UserType:    user.UserType,
		Token:       user.Token,
		CreateTime:  &timestamp.Timestamp{Seconds: user.CreateTime.Unix()},
		UpdateTime:  &timestamp.Timestamp{Seconds: user.UpdateTime.Unix()},
	}
}

func convertToPBUsers(users []repository.User) []*pb.User {
	items := make([]*pb.User, len(users))
	for i, user := range users {
		items[i] = convertToPBUser(user)
	}
	return items
}

func (server *UserServiceServer) Create(ctx context.Context, in *pb.User) (*pb.User, error) {
	user := convertToModelUser(in)
	createdUser, err := repository.CreateUser(user)
	return convertToPBUser(createdUser), err
}

func (server *UserServiceServer) SignIn(ctx context.Context, in *pb.UserCredentials) (*pb.User, error) {
	username := in.Username
	password := in.Password
	existingUser, err := repository.FindUserFullByUsername(username)
	if err != nil {
		return nil, err
	}
	if err := util.ComparePassword(existingUser.Password, password); err != nil {
		return nil, err
	}
	return convertToPBUser(existingUser), err
}

func (server *UserServiceServer) Get(ctx context.Context, in *pb.User) (*pb.User, error) {
	user, err := repository.FindUserFullById(int(in.Id))
	return convertToPBUser(user), err
}

func (server *UserServiceServer) Count(ctx context.Context, in *pb.User) (*wrappers.Int32Value, error) {
	query := make(map[string]interface{})
	if in.UserType != "" {
		query["user_type"] = in.UserType
	}
	count, err := repository.CountUser(query)
	return &wrappers.Int32Value{Value: int32(count)}, err
}

func (server *UserServiceServer) ListAll(ctx context.Context, in *pb.UserPageRequest) (*pb.Users, error) {
	query := make(map[string]interface{})
	if in.User.UserType != "" {
		query["user_type"] = in.User.UserType
	}
	pageRequest := convertToModelPageRequest(in.PageRequest)
	users, err := repository.FindAllUserFullList(query, pageRequest)
	items := convertToPBUsers(users)
	return &pb.Users{Items: items}, err
}

func (server *UserServiceServer) List(ctx context.Context, in *pb.UserPageRequest) (*pb.UserPageResponse, error) {
	query := make(map[string]interface{})
	if in.User.UserType != "" {
		query["user_type"] = in.User.UserType
	}
	pageRequest := convertToModelPageRequest(in.PageRequest)
	users, err := repository.FindAllUserFullList(query, pageRequest)
	items := convertToPBUsers(users)
	count, err := repository.CountUser(query)
	return &pb.UserPageResponse{
		Page:  int32(pageRequest.Page),
		Size:  int32(pageRequest.Size),
		Count: int32(count),
		Items: items,
	}, err
}

func (server *UserServiceServer) Update(ctx context.Context, in *pb.User) (*pb.User, error) {
	user := convertToModelUser(in)
	updatedUser, err := repository.UpdateUser(user)
	return convertToPBUser(updatedUser), err
}

func (server *UserServiceServer) Delete(ctx context.Context, in *pb.User) (*empty.Empty, error) {
	err := repository.DeleteUserById(int(in.Id))
	return &empty.Empty{}, err
}
