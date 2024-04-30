package service

import (
	"context"
	"net/http"

	"github.com/14jasimmtp/GigForge-Freelance-Marketplace/Job-svc/pb/job"
)

// func (s *Service) GetCategory(ctx context.Context, req *job.GetCategoryReq) (*job.GetCategoryRes, error) {
// 	category, err := s.repo.GetCategory()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return category, nil

// }
func (s *Service) AddCategory(ctx context.Context, req *job.AddCategoryReq) (*job.AddCategoryRes, error) {
	categories, err := s.repo.AddCategory(req)
	if err != nil {
		return &job.AddCategoryRes{Status: http.StatusInternalServerError,Error: "something went wrong"}, nil
	}
	return categories, nil
}
