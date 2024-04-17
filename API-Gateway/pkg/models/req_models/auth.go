package req

type LoginRequest struct{
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SignupRequest struct{
	FirstName string `json:"firstname" validate:"required"`
	LastName string `json:"lastname" validate:"required"`
	Password string `json:"password" validate:"required"`
	Phone string `json:"phone" validate:"required,number,len=10"`
	Email string `json:"email" validate:"required,email"`
	Country string `json:"country" validate:"required"`
}

type ForgotPassword struct{
	Email string `json:"email" validate:"required,email"`
}

type ResetPassword struct{
	NewPassword string `json:"password"`
	OTP int64 `json:"otp"`
}

type Verify struct{
	OTP int64 `json:"otp" validate:"required"`
}

