package repository

import (
	"time"

	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/pb/auth"
	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/pkg/domain"
)

type RepoIfc interface {
	CreateUser(user *auth.UserSignupReq) (*domain.User, error)
	CheckIsUserActive(email string) error
	CheckUserExist(email, phone string) error
	CheckOTP(email string) (int64, time.Time, error)
	GetUser(email string) (*domain.UserModel, error)
	SaveOTP(otp int, email string, exp time.Time) error
	DeleteOTP() error
	AddEducation(edu *auth.AddEducationReq) (*domain.Freelancer_Education, error)
	UpdateEducation(edu *auth.UpdateEducationReq) (*domain.Freelancer_Education, error)
	DeleteEducation(edu *auth.DeleteEducationReq) error
	AddProfileDescription(req *auth.APDReq) (*domain.Freelancer_Description, error)
	UpdateProfileDescription(req *auth.UPDReq) (*domain.Freelancer_Description, error)
	GetExperience(id string) ([]*auth.ExpReq, error)
	GetEducations(id string) ([]*auth.Education, error)
	GetProfileDescription(id string) (*auth.UPDReq, error)
	GetUserWithId(id string) (*auth.User, error)
	AddExperience(edu *auth.ExpReq) error
	UpdateExperience(edu *auth.ExpReq) error
	DeleteExperience(edu *auth.DltExpReq) error
	UpdatePassword(password, email string) error
	AddSkill(req *auth.AddSkillReq) (int, error)
	BlockUser(userID string) (int, error)
	UnBlockUser(userID string) (int, error)
	UpdateProfilePhoto(userID,url string) error
	CheckUserOnboardStatus(user_id string) error
	UpdateSkillUserProfile(user_id string, skill []int64) ([]string,error)
	CheckSkillsExist(skills []int64) error
}
