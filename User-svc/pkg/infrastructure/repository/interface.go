package repository

import (
	"time"

	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/pb/auth"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/pkg/domain"
)

type RepoIfc interface {
	CreateUser(user *auth.UserSignupReq) (*domain.User, error)
	CheckIsUserActive(email string) domain.User
	CheckUserExist(email, phone string) error
	CheckOTP(email string) (int64,time.Time,error)
	GetUser(email string) (*domain.UserModel, error)
	SaveOTP(otp int, email string, exp time.Time) error
	DeleteOTP() error

	AddEducation(edu *auth.AddEducationReq) (*domain.Freelancer_Education,error)
	UpdateEducation(edu *auth.UpdateEducationReq) (*domain.Freelancer_Education,error)
	DeleteEducation(edu *auth.DeleteEducationReq) (error)
	AddProfileDescription(req *auth.APDReq) (*domain.Freelancer_Description,error)
	UpdateProfileDescription(req *auth.UPDReq) (*domain.Freelancer_Description,error)
}
