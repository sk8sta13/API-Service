package dto

type CreateProduct struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CreateUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetToken struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetJWT struct {
	AccessToken string `json:"access_token"`
}
