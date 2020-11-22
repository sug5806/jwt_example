package model

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	JwtToken string `json:"jwt_token"`
}
