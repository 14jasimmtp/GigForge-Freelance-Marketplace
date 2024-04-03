package service

import (
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/pb/auth"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/pkg/infrastructure/repository"
)

type ProfileService struct {
	Repo repository.RepoIfc
	auth.UnimplementedAuthServiceServer
}

func NewProfileService(r repository.RepoIfc) ProfileService {
	return ProfileService{Repo: r}
}

// func (s *ProfileService) AddEducation(ctx context.Context,req *auth.AddEducationReq) (*auth.AddEducationRes,error){
// 	res,err:=s.Repo.AddEducation(req)
// 	if err != nil{
// 		return &auth.AddEducationRes{
// 			Status: 400,
// 			Message: "something went wrong",
// 			Error: err.Error(),
// 		},nil
// 	}
// 	return &auth.AddEducationRes{
// 		Status: 200,
// 		Message: "Education added successfully",
// 		Body: &auth.Education{
// 			EducationId: int64(res.ID),
// 			School: res.School,
// 			Course: res.Course,
// 			Date_Started: res.Date_Started,
// 			Date_Ended: res.Date_Ended,
// 			AreaOfStudy: res.Area_Of_Study,
// 			Description: res.Description,
// 		},
// 	},nil
// }

// func (s *ProfileService) UpdateEducation(ctx context.Context,req *auth.UpdateEducationReq) (*auth.UpdateEducationRes,error){
// 	res,err:=s.Repo.UpdateEducation(req)
// 	if err != nil{
// 		return &auth.UpdateEducationRes{
// 			Status: 400,
// 			Message: "something went wrong",
// 			Error: err.Error(),
// 		},nil
// 	}
// 	return &auth.UpdateEducationRes{
// 		Status: 200,
// 		Message: "Education added successfully",
// 		Body: &auth.Education{
// 			EducationId: int64(res.ID),
// 			School: res.School,
// 			Course: res.Course,
// 			Date_Started: res.Date_Started,
// 			Date_Ended: res.Date_Ended,
// 			AreaOfStudy: res.Area_Of_Study,
// 			Description: res.Description,
// 		},
// 	},nil
// }

// func (s *ProfileService) DeleteEducation(ctx context.Context,req *auth.DeleteEducationReq) (*auth.DeleteEducationRes,error){
// 	err:=s.Repo.DeleteEducation(req)
// 	if err != nil{
// 		return &auth.DeleteEducationRes{
// 			Status: 400,
// 			Message:
// 			"something went wrong",
// 			Error: err.Error(),
// 		},nil
// 	}
// 	return &auth.DeleteEducationRes{
// 		Status: 200,
// 		Message: "Education deleted successfully",
// 	},nil
// }

// func (s *ProfileService) AddProfileDescription(ctx context.Context, req *auth.APDReq) (*auth.APDRes,error){
// 	_,err:=s.Repo.AddProfileDescription(req)
// 	if err != nil{
// 		return &auth.APDRes{
// 			Status: 400,
// 			Message:
// 			"something went wrong",
// 			Error: err.Error(),
// 		},nil
// 	}
// 	return &auth.APDRes{
// 		Status: 200,
// 		Message: "Education deleted successfully",
// 	},nil
// }

// func (s *ProfileService) UpdateProfileDescription(ctx context.Context, req *auth.UPDReq) (*auth.UPDRes,error){
// 	_,err:=s.Repo.UpdateProfileDescription(req)
// 	if err != nil{
// 		return &auth.UPDRes{
// 			Status: 400,
// 			Message:
// 			"something went wrong",
// 			Error: err.Error(),
// 		},nil
// 	}
// 	return &auth.UPDRes{
// 		Status: 200,
// 		Message: "description updated successfully",
// 	},nil

// }

// // func (s *ProfileService) GetProfile(ctx context.Context, req *auth.GetProfileReq) (*auth.GetProfileRes,error){

// // }
