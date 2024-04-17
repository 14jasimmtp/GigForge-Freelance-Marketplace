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

func NewService(repo repository.RepoIfc) ServiceIfc {
	return &Service{repo: repo}
}

func (s *Service) Login(ctx context.Context, req *auth.UserLoginReq) (*auth.UserLoginRes, error) {
	
	user, err := s.repo.GetUser(req.Email)
	if err != nil {
		return &auth.UserLoginRes{
			Status:  http.StatusNotFound,
			Message: "user doesn't exist with this email",
		}, nil
	}

	err = s.repo.CheckIsUserActive(req.Email)
	if err != nil {
		return &auth.UserLoginRes{
			Status: http.StatusUnauthorized,
			Error:  err.Error(),
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
			Status:  401,
			Message: "token expired",
			Error:   err.Error(),
		}, nil
	}

	userOTP, exp, err := s.repo.CheckOTP(user.Email)
	if err != nil {
		return &auth.VerifyRes{Status: 400, Message: "otp doesn't sent.signup again", Error: err.Error()}, nil
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

func (s *Service) ForgotPassword(ctx context.Context, req *auth.FPreq) (*auth.FPres, error) {
	otpp, err := otp.SendVerificationOtp(req.Email)
	if err != nil {
		return &auth.FPres{
			Status:   http.StatusInternalServerError,
			Response: "error while sending otp",
			Error:    err.Error(),
		}, nil
	}
	exp := time.Now().Add(2 * time.Minute)
	err = s.repo.SaveOTP(otpp, req.Email, exp)
	if err != nil {
		return &auth.FPres{
			Status:   http.StatusInternalServerError,
			Response: "error while saving otp",
			Error:    err.Error(),
		}, nil
	}
	token, err := jwtoken.GenerateTemporaryTokenToResetPwd(viper.GetString("ATokenSecret"), req.Email)
	if err != nil {
		return &auth.FPres{
			Error:  err.Error(),
			Status: 400,
		}, nil
	}
	return &auth.FPres{Status: 200, Response: token}, nil
}

func (s *Service) ResetPassword(ctx context.Context, req *auth.RPreq) (*auth.RPres, error) {
	email, err := jwtoken.FetchEmailFromToken(req.Token)
	fmt.Println(email)
	if err != nil {
		return &auth.RPres{
			Status:   401,
			Response: "token problem",
			Error:    err.Error(),
		}, nil
	}

	userOTP, exp, err := s.repo.CheckOTP(email)
	if err != nil {
		return &auth.RPres{Status: 400, Response: "otp doesn't sent.signup again", Error: err.Error()}, nil
	}
	fmt.Println(userOTP, " ", req.OTP)
	if userOTP != req.OTP {
		return &auth.RPres{Status: 400, Response: "otp doesn't match"}, nil
	}
	if time.Now().After(exp) {
		return &auth.RPres{Status: 400, Response: "otp expired"}, nil
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	pq := string(hashed)

	err = s.repo.UpdatePassword(pq, email)
	if err != nil {
		return &auth.RPres{Status: 400, Response: "something went wrong", Error: err.Error()}, nil
	}
	return &auth.RPres{Status: 200, Response: "password changed successfully"}, nil
}

