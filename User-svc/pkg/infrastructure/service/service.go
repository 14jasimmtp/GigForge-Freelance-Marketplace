package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/pb/auth"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/pkg/infrastructure/repository"
	jwtoken "github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/utils/jwt"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/utils/otp"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo repository.RepoIfc
	auth.UnimplementedAuthServiceServer
}

func NewService(repo repository.RepoIfc) Service {
	return Service{repo: repo}
}

func (s *Service) Login(ctx context.Context, req *auth.UserLoginReq) (*auth.UserLoginRes, error) {
	// err := s.repo.CheckUserExist(req.Email,"")
	// if err != nil {
	// 	return &auth.UserLoginRes{
	// 		Status: http.StatusNotFound,
	// 		Error: err.Error(),
	// 	}, nil
	// }

	// _ = s.repo.CheckIsUserActive(req.Email)
	// if err != nil {
	// 	return &auth.UserLoginRes{
	// 		Status: http.StatusUnauthorized,
	// 		Error:  err.Error(),
	// 	}, nil
	// }

	user, err := s.repo.GetUser(req.Email)
	if err != nil {
		return &auth.UserLoginRes{
			Status:  http.StatusNotFound,
			Message: "user doesn't exist with this email",
		}, nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return &auth.UserLoginRes{
			Status:  http.StatusUnauthorized,
			Message: "password doesn't match",
			Error:   err.Error(),
		}, nil
	}
	fmt.Println(viper.GetString("ATokenSecret"))
	accessToken, err := jwtoken.GenerateAccessToken(viper.GetString("ATokenSecret"), user)
	if err != nil {
		return &auth.UserLoginRes{
			Status:  http.StatusInternalServerError,
			Message: "Err in token generation",
			Error:   err.Error(),
		}, nil
	}

	return &auth.UserLoginRes{
		Status:  200,
		Message: "user logged in successfully",
		Token:   accessToken,
	}, nil
}

func (s *Service) Signup(ctx context.Context, user *auth.UserSignupReq) (*auth.UserSignupRes, error) {
	log.Print("signup service")
	err := s.repo.CheckUserExist(user.Email, user.Phone)
	if err != nil {
		return &auth.UserSignupRes{
			Status:  http.StatusConflict,
			Message: "user already exists",
			Error:   err.Error(),
		}, nil
	}

	fmt.Println("otp sented")
	otpp, err := otp.SendVerificationOtp(user.Email)
	if err != nil {
		return &auth.UserSignupRes{
			Status:  http.StatusForbidden,
			Message: "error while sending otp",
			Error:   err.Error(),
		}, nil
	}
	exp := time.Now().Add(2 * time.Minute)
	err = s.repo.SaveOTP(otpp, user.Email, exp)
	if err != nil {
		return &auth.UserSignupRes{
			Status:  http.StatusInternalServerError,
			Message: "error while saving otp",
			Error:   err.Error(),
		}, nil
	}
	fmt.Println("token generated")
	token, err := jwtoken.GenerateTemporaryTokenToVerify(viper.GetString("ATokenSecret"), user)
	if err != nil {
		return &auth.UserSignupRes{Status: 400, Message: "Error while generating token to verify", Error: err.Error()}, nil
	}

	return &auth.UserSignupRes{
		Status:  201,
		Message: "OTP sent to email.Verify to get access",
		Token:   token,
	}, nil
}

func (s *Service) Verify(ctx context.Context, req *auth.VerifyReq) (*auth.VerifyRes, error) {
	user, err := jwtoken.FetchUserVerifyDetailsFromToken(req.Token)
	if err != nil {
		return &auth.VerifyRes{
			Status: 401,
			Message: "token expired",
			Error: err.Error(),
		}, nil
	}

	userOTP, exp, err := s.repo.CheckOTP(user.Email)
	if err != nil {
		return &auth.VerifyRes{Status: 400, Message: "otp doesn't sent.gignup again", Error: err.Error()}, nil
	}
	fmt.Println(userOTP, " ", req.OTP)
	if userOTP != req.OTP {
		return &auth.VerifyRes{Status: 400, Message: "otp doesn't match"}, nil
	}
	if time.Now().After(exp) {
		return &auth.VerifyRes{Status: 400, Message: "otp expired"}, nil
	}

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return &auth.VerifyRes{
			Status:  http.StatusInternalServerError,
			Message: "Error while hashing password",
			Error:   err.Error(),
		}, nil
	}
	user.Password = string(hashedPwd)

	_, err = s.repo.CreateUser(user)
	if err != nil {
		return &auth.VerifyRes{
			Status:  http.StatusNotImplemented,
			Message: "error while creating user",
			Error:   err.Error(),
		}, nil
	}

	return &auth.VerifyRes{Status: 200, Message: "user verified successfully.Login to get access."}, nil

}

// func (s *Service) ForgotPassword(ctx context.Context, req *auth.FPreq) (*auth.FPres, error) {

// }

// func (s *Service) ResetPassword(ctx context.Context, req *auth.RPreq) (*auth.RPres, error) {

// }


func (s *Service) AddEducation(ctx context.Context,req *auth.AddEducationReq) (*auth.AddEducationRes,error){
	res,err:=s.repo.AddEducation(req)
	if err != nil{
		return &auth.AddEducationRes{
			Status: 400,
			Message: "something went wrong",
			Error: err.Error(),
		},nil
	}
	return &auth.AddEducationRes{
		Status: 200,
		Message: "Education added successfully",
		Body: &auth.Education{
			EducationId: int64(res.ID),
			School: res.School,
			Course: res.Course,
			Date_Started: res.Year_Started,
			Date_Ended: res.Year_Ended,
			AreaOfStudy: res.Area_Of_Study,
			Description: res.Description,
		},
	},nil
}

func (s *Service) UpdateEducation(ctx context.Context,req *auth.UpdateEducationReq) (*auth.UpdateEducationRes,error){
	res,err:=s.repo.UpdateEducation(req)
	if err != nil{
		return &auth.UpdateEducationRes{
			Status: 400,
			Message: "something went wrong",
			Error: err.Error(),
		},nil
	}
	return &auth.UpdateEducationRes{
		Status: 200,
		Message: "Education added successfully",
		Body: &auth.Education{
			EducationId: int64(res.ID),
			School: res.School,
			Course: res.Course,
			Date_Started: res.Year_Started,
			Date_Ended: res.Year_Ended,
			AreaOfStudy: res.Area_Of_Study,
			Description: res.Description,
		},
	},nil
}

func (s *Service) DeleteEducation(ctx context.Context,req *auth.DeleteEducationReq) (*auth.DeleteEducationRes,error){
	err:=s.repo.DeleteEducation(req)
	if err != nil{
		return &auth.DeleteEducationRes{
			Status: 400,
			Message: 	
			"something went wrong",
			Error: err.Error(),
		},nil
	}
	return &auth.DeleteEducationRes{
		Status: 200,
		Message: "Education deleted successfully",
	},nil
}

func (s *Service) AddProfileDescription(ctx context.Context, req *auth.APDReq) (*auth.APDRes,error){
	_,err:=s.repo.AddProfileDescription(req)
	if err != nil{
		return &auth.APDRes{
			Status: 400,
			Message: 	
			"something went wrong",
			Error: err.Error(),
		},nil
	}
	return &auth.APDRes{
		Status: 200,
		Message: "Education deleted successfully",
	},nil
}

func (s *Service) UpdateProfileDescription(ctx context.Context, req *auth.UPDReq) (*auth.UPDRes,error){
	_,err:=s.repo.UpdateProfileDescription(req)
	if err != nil{
		return &auth.UPDRes{
			Status: 400,
			Message: 	
			"something went wrong",
			Error: err.Error(),
		},nil
	}
	return &auth.UPDRes{
		Status: 200,
		Message: "description updated successfully",
	},nil

}

// func (s *Service) GetProfile(ctx context.Context, req *auth.GetProfileReq) (*auth.GetProfileRes,error){

// }