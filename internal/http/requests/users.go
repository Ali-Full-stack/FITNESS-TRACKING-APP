package requests

type UserRegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

 type UserRegisterResponse struct{
	ID int `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Profile  map[string]any `json:"profile,omitempty"`
 }

 type UserLoginRequest struct{
	ID int32  `json:"id"`
	Password string `json:"password"`
 }

 type UserLoginResponse struct{
	Token string `json:"token"`
 }
