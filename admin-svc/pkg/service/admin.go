package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/admin-svc/pb"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/admin-svc/pkg/repository"
	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/admin-svc/utils/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AdminService struct {
	pb.UnimplementedAdminServiceServer
	repo repository.RepoIfc
}

type ServiceIfc interface {
}

func NewAdminService(repo repository.RepoIfc) AdminService {
	return AdminService{repo: repo}
}

func (s *AdminService) AdminLogin(ctx context.Context, admin *pb.LoginReq) (*pb.LoginRes, error) {
	AdminDetails, err := s.repo.AdminLogin(admin.Email)
	if err != nil {
		fmt.Println("Admin doesn't exist")
		return &pb.LoginRes{
			Status: http.StatusUnauthorized,
			Error: err.Error(),
			Response: "Email doesn't exist",
		}, nil
	}

	if bcrypt.CompareHashAndPassword([]byte(AdminDetails.Password), []byte(admin.Password)) != nil {
		fmt.Println("wrong password")
		return &pb.LoginRes{
			Status: http.StatusUnauthorized,
			Error: errors.New("wrong password").Error(),
			Response: "Enter password correctly",
		},nil
	}

	tokenString, err := jwt.AdminTokenGenerate(AdminDetails)
	if err != nil {
		fmt.Println("error generating token")
		return &pb.LoginRes{
			Status: http.StatusInternalServerError,
			Error: err.Error(),
			Response: "error while generating token for admin",
		}, nil
	}

	return &pb.LoginRes{
		Status:   http.StatusOK,
		Token: tokenString,
		Response: "admin logged in successfully",
	}, nil
}

