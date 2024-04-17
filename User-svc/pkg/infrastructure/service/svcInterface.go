package service

import (
	"context"

	"github.com/14jasimmtp/GigForge-Freelancer-Marketplace/User-Auth/pb/auth"
)

type ServiceIfc interface {
	Login(ctx context.Context, req *auth.UserLoginReq) (*auth.UserLoginRes, error)
	Signup(ctx context.Context, user *auth.UserSignupReq) (*auth.UserSignupRes, error)
	Verify(ctx context.Context, req *auth.VerifyReq) (*auth.VerifyRes, error)
	ForgotPassword(ctx context.Context, req *auth.FPreq) (*auth.FPres, error)
	ResetPassword(ctx context.Context, req *auth.RPreq) (*auth.RPres, error)
	AddEducation(ctx context.Context, req *auth.AddEducationReq) (*auth.AddEducationRes, error)
	UpdateEducation(ctx context.Context, req *auth.UpdateEducationReq) (*auth.UpdateEducationRes, error)
	DeleteEducation(ctx context.Context, req *auth.DeleteEducationReq) (*auth.DeleteEducationRes, error)
	AddProfileDescription(ctx context.Context, req *auth.APDReq) (*auth.APDRes, error)
	UpdateProfileDescription(ctx context.Context, req *auth.UPDReq) (*auth.UPDRes, error)
	AddExperience(ctx context.Context, req *auth.ExpReq) (*auth.ExpRes, error)
	UpdateExperience(ctx context.Context, req *auth.ExpReq) (*auth.ExpRes, error)
	DeleteExperience(ctx context.Context, req *auth.DltExpReq) (*auth.DltExpRes, error)
	GetProfile(ctx context.Context, req *auth.GetProfileReq) (*auth.GetProfileRes, error)
	UpdateProfilePhoto(ctx context.Context, req *auth.PhotoReq) (*auth.PhotoRes, error)
	AddSkill(ctx context.Context, req *auth.AddSkillReq) (*auth.AddSkillRes, error)
	UnBlockUser(ctx context.Context, req *auth.BlockReq) (*auth.BlockRes, error)
	BlockUser(ctx context.Context, req *auth.BlockReq) (*auth.BlockRes, error)
	EditSkill(ctx context.Context, req *auth.EditSkillReq) (*auth.EditSkillRes, error)
	GetProfileClient(ctx context.Context, req *auth.ClientProfileReq) (*auth.ClientProfileRes, error)
}
