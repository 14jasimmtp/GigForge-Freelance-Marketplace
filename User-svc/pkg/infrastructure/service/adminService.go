package service

import (
	"context"

	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/pb/auth"
)

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