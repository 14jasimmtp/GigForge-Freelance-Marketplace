package service

import (
	"context"

	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/pb/auth"
	s3 "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/utils/s3bucket"
)

func (s *Service) AddEducation(ctx context.Context, req *auth.AddEducationReq) (*auth.AddEducationRes, error) {
	res, err := s.repo.AddEducation(req)
	if err != nil {
		return &auth.AddEducationRes{
			Status:  400,
			Message: "something went wrong",
			Error:   err.Error(),
		}, nil
	}
	return &auth.AddEducationRes{
		Status:  200,
		Message: "Education added successfully",
		Body: &auth.Education{
			EducationId:  int64(res.ID),
			School:       res.School,
			Course:       res.Course,
			Date_Started: res.Year_Started,
			Date_Ended:   res.Year_Ended,
			AreaOfStudy:  res.Area_Of_Study,
			Description:  res.Description,
		},
	}, nil
}


func (s *Service) UpdateEducation(ctx context.Context, req *auth.UpdateEducationReq) (*auth.UpdateEducationRes, error) {
	res, err := s.repo.UpdateEducation(req)
	if err != nil {
		return &auth.UpdateEducationRes{
			Status:  400,
			Message: "something went wrong",
			Error:   err.Error(),
		}, nil
	}
	return &auth.UpdateEducationRes{
		Status:  200,
		Message: "Education added successfully",
		Body: &auth.Education{
			EducationId:  int64(res.ID),
			School:       res.School,
			Course:       res.Course,
			Date_Started: res.Year_Started,
			Date_Ended:   res.Year_Ended,
			AreaOfStudy:  res.Area_Of_Study,
			Description:  res.Description,
		},
	}, nil
}

func (s *Service) DeleteEducation(ctx context.Context, req *auth.DeleteEducationReq) (*auth.DeleteEducationRes, error) {
	err := s.repo.DeleteEducation(req)
	if err != nil {
		return &auth.DeleteEducationRes{
			Status:  400,
			Message: "something went wrong",
			Error:   err.Error(),
		}, nil
	}
	return &auth.DeleteEducationRes{
		Status:  200,
		Message: "Education deleted successfully",
	}, nil
}

func (s *Service) AddProfileDescription(ctx context.Context, req *auth.APDReq) (*auth.APDRes, error) {
	_, err := s.repo.AddProfileDescription(req)
	if err != nil {
		return &auth.APDRes{
			Status:  400,
			Message: "something went wrong",
			Error:   err.Error(),
		}, nil
	}
	return &auth.APDRes{
		Status:  200,
		Message: "added profile description successfully",
	}, nil
}

func (s *Service) UpdateProfileDescription(ctx context.Context, req *auth.UPDReq) (*auth.UPDRes, error) {
	_, err := s.repo.UpdateProfileDescription(req)
	if err != nil {
		return &auth.UPDRes{
			Status:  400,
			Message: "something went wrong",
			Error:   err.Error(),
		}, nil
	}
	return &auth.UPDRes{
		Status:  200,
		Message: "description updated successfully",
	}, nil

}

func (s *Service) AddExperience(ctx context.Context, req *auth.ExpReq) (*auth.ExpRes, error) {
	err := s.repo.AddExperience(req)
	if err != nil {
		return &auth.ExpRes{
			Status: 400,
			Error:  err.Error(),
		}, nil
	}
	return &auth.ExpRes{
		Status:   200,
		Response: "experience added to profile successfully",
	}, nil
}

func (s *Service) UpdateExperience(ctx context.Context, req *auth.ExpReq) (*auth.ExpRes, error) {
	err := s.repo.UpdateExperience(req)
	if err != nil {
		return &auth.ExpRes{
			Status: 400,
			Error:  err.Error(),
		}, nil
	}
	return &auth.ExpRes{
		Status:   200,
		Response: "experience updated successfully",
	}, nil
}

func (s *Service) DeleteExperience(ctx context.Context, req *auth.DltExpReq) (*auth.DltExpRes, error) {
	err := s.repo.DeleteExperience(req)
	if err != nil {
		return &auth.DltExpRes{
			Status: 400,
			Error:  err.Error(),
		}, nil
	}
	return &auth.DltExpRes{
		Status:   200,
		Response: "experience deleted successfully",
	}, nil
}

func (s *Service) GetProfile(ctx context.Context, req *auth.GetProfileReq) (*auth.GetProfileRes, error) {
	user, err := s.repo.GetUserWithId(req.UserId)
	if err != nil {
		return &auth.GetProfileRes{
			Status: 400,
			Error:  err.Error(),
		}, nil
	}
	description, err := s.repo.GetProfileDescription(req.UserId)
	if err != nil {
		return &auth.GetProfileRes{
			Status: 400,
			Error:  err.Error(),
		}, nil
	}
	education, err := s.repo.GetEducations(req.UserId)
	if err != nil {
		return &auth.GetProfileRes{
			Status: 400,
			Error:  err.Error(),
		}, nil
	}
	experience, err := s.repo.GetExperience(req.UserId)
	if err != nil {
		return &auth.GetProfileRes{
			Status: 400,
			Error:  err.Error(),
		}, nil
	}

	return &auth.GetProfileRes{
		User:        user,
		Description: description,
		Education:   education,
		Experience:  experience,
		Status:      200,
	}, nil
}

func (s *Service) UpdateProfilePhoto(ctx context.Context, req *auth.PhotoReq) (*auth.PhotoRes, error) {
	sess := s3.CreateSession()
	url, err := s3.UploadImageToS3(req.Image, sess)
	if err != nil {
		return &auth.PhotoRes{
			Status:   500,
			Error:    err.Error(),
			Response: "error while uploading profile image",
		}, nil
	}
	err = s.repo.UpdateProfilePhoto(req.UserId, url)
	if err != nil {
		return &auth.PhotoRes{
			Status:   400,
			Error:    err.Error(),
			Response: "something went wrong",
		}, nil
	}
	return &auth.PhotoRes{
		Status:   200,
		Response: "profile photo updated successfully",
	}, nil
}

// func (s *Service) UpdateCompanyDetails(ctx context.Context, req *auth.UPDCompanyDetailsReq) (*auth.UPDCompanyDetailsRes, error){
// 	err:=s.repo.UpdateCmpDtails(req)
// 	if err != nil{
// 		return &auth.UPDCompanyDetailsRes{}
// 	}
// }

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
