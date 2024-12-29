package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sk8sta13/API-Service/internal/dto"
	"github.com/sk8sta13/API-Service/internal/entity"
	"github.com/sk8sta13/API-Service/internal/infra/database"

	"github.com/go-chi/jwtauth"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserDB database.UserInterface
	JWT    *jwtauth.JWTAuth
	JWTTTL int
}

func NewUserHandler(db database.UserInterface, jwt *jwtauth.JWTAuth, ttl int) *UserHandler {
	return &UserHandler{
		UserDB: db,
		JWT:    jwt,
		JWTTTL: ttl,
	}
}

// Create user godoc
// @Summary Create user
// @Description Create user
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.CreateUser true "user request"
// @Success 201
// @Failure 500 {object} Error
// @Router /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userDTO dto.CreateUser
	err := json.NewDecoder(r.Body).Decode(&userDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := entity.NewUser(userDTO.Name, userDTO.Email, userDTO.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.UserDB.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}

// @GetJWT godoc
// @Summary Get a user JWT
// @Description Get a user JWT
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.GetToken true "user credentials"
// @Success 200 {object} dto.GetJWT
// @Failure 404
// @Failure 500 {object} Error
// @Router /users/generate_token [post]
func (h *UserHandler) GetToken(w http.ResponseWriter, r *http.Request) {
	var userDTO dto.GetToken
	err := json.NewDecoder(r.Body).Decode(&userDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.UserDB.FindByEmail(userDTO.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !user.CheckPassword(userDTO.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, tokenstr, _ := h.JWT.Encode(map[string]interface{}{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JWTTTL)).Unix(),
	})

	accessToken := dto.GetJWT{AccessToken: tokenstr}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}
