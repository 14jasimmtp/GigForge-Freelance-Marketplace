package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/pb/auth"
	jwtoken "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/utils/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) AdminLogin(ctx context.Context, req *auth.LoginReq) (*auth.LoginRes, error) {
	AdminDetails, err := s.repo.AdminLogin(req.Email)
	if err != nil {
		fmt.Println("Admin doesn't exist")
		return &auth.LoginRes{
			Status:   http.StatusUnauthorized,
			Error:    err.Error(),
			Response: "Email doesn't exist",
		}, nil
	}

	if bcrypt.CompareHashAndPassword([]byte(AdminDetails.Password), []byte(req.Password)) != nil {
		fmt.Println("wrong password")
		return &auth.LoginRes{
			Status:   http.StatusUnauthorized,
			Error:    errors.New("wrong password").Error(),
			Response: "Enter password correctly",
		}, nil
	}

	tokenString, err := jwtoken.AdminTokenGenerate(AdminDetails)
	if err != nil {
		fmt.Println("error generating token")
		return &auth.LoginRes{
			Status:   http.StatusInternalServerError,
			Error:    err.Error(),
			Response: "error while generating token for admin",
		}, nil
	}

	return &auth.LoginRes{
		Status:   http.StatusOK,
		Token:    tokenString,
		Response: "admin logged in successfully",
	}, nil
}

func (s *Service) BlockUser(ctx context.Context, req *auth.BlockReq) (*auth.BlockRes, error) {
	status, err := s.repo.BlockUser(req.UserId)
	if err != nil {
		return &auth.BlockRes{
			Status: int64(status),
			Error:  err.Error(),
		}, nil
	}
	return &auth.BlockRes{
		Status:   int64(status),
		Response: "user successfully blocked",
	}, nil
}

func (s *Service) UnBlockUser(ctx context.Context, req *auth.BlockReq) (*auth.BlockRes, error) {
	status, err := s.repo.UnBlockUser(req.UserId)
	if err != nil {
		return &auth.BlockRes{
			Status: int64(status),
			Error:  err.Error(),
		}, nil
	}
	return &auth.BlockRes{
		Status:   int64(status),
		Response: "user Unblocked successfully",
	}, nil
}

func (s *Service) AddSkill(ctx context.Context, req *auth.AddSkillReq) (*auth.AddSkillRes, error) {
	status, err := s.repo.AddSkill(req)
	if err != nil {
		return &auth.AddSkillRes{
			Status: int64(status),
			Error:  err.Error(),
		}, nil
	}
	return &auth.AddSkillRes{
		Status:   int64(status),
		Response: "skill added to database successfully",
	}, nil
}
