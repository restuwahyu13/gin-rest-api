package resendAuth

type InputResend struct {
	Email string `json:"email"  binding:"required"`
}
