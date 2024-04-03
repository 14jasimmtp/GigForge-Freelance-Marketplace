package req

type LoginRequest struct{
	Email string `json:"email" `
	Password string `json:"password"`
}

type SignupRequest struct{
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Password string `json:"password"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Country string `json:"country"`
}

type ForgotPassword struct{
	Email string `json:"email"`
}

type ResetPassword struct{
	NewPassword string `json:"password"`
	OTP int64 `json:"otp"`
}

type Verify struct{
	OTP int64 `json:"otp"`
}

