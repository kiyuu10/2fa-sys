package handlers

type (
	OTPEmailReq struct {
		Email string `json:"email"`
	}

	OTPEmailVerifyReq struct {
		Email string `json:"email"`
		OTP   string `json:"otp"`
	}
)
