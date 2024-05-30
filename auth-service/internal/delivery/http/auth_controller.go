package deliver_http

import (
	"errors"
	"github.com/MuhamadAndre/auth-service/internal/models"
	"github.com/MuhamadAndre/auth-service/internal/usecase"
	"github.com/MuhamadAndre/auth-service/internal/utils"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type AuthController struct {
	log         *zap.Logger
	authUseCase *usecase.AuthUseCase
}

func NewAuthController(log *zap.Logger, authUseCase *usecase.AuthUseCase) *AuthController {
	return &AuthController{
		log:         log,
		authUseCase: authUseCase,
	}
}

func (c *AuthController) SignIn(w http.ResponseWriter, r *http.Request) {

	var payload models.LoginRequest

	err := utils.DecodeJson(w, r, &payload)
	if err != nil {
		c.log.Error("Error decoding json body", zap.Error(err))
		utils.WriteErrorJson(w, err, http.StatusBadRequest)
		return
	}

	userResp, err := c.authUseCase.SignIn(r.Context(), &payload)
	if err != nil {
		c.log.Error("Error sign in", zap.Error(err))
		utils.WriteErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	token, err := utils.CreateToken(time.Duration(300)*time.Second, userResp.ID)
	if err != nil {
		c.log.Error("Error create token", zap.Error(err))
		utils.WriteErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJsonBody(w, http.StatusOK, utils.WebResponse[*models.UserResponse]{
		Message: "login success",
		Token:   token,
		Data:    userResp,
	}, http.Header{"Authorization": {token}})

}

func (c *AuthController) SignUp(w http.ResponseWriter, r *http.Request) {

	var payload models.RegisterRequest

	err := utils.DecodeJson(w, r, &payload)
	if err != nil {
		c.log.Error("Error decoding json body", zap.Error(err))
		utils.WriteErrorJson(w, err, http.StatusBadRequest)
		return
	}

	response, err := c.authUseCase.SignUp(r.Context(), &payload)
	if err != nil {
		c.log.Error("Error sign up", zap.Error(err))
		utils.WriteErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	token, err := utils.CreateToken(time.Duration(180)*time.Second, response.Email)
	if err != nil {
		c.log.Error("Error create token", zap.Error(err))
		utils.WriteErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJsonBody(w, http.StatusOK, utils.WebResponse[*models.UserResponse]{
		Message: "Register success",
		Token:   token,
		Data:    response,
	}, http.Header{"Authorization": {token}})

}

func (c *AuthController) VerifyUser(w http.ResponseWriter, r *http.Request) {
	// get token from params
	tknFromParam := chi.URLParam(r, "token")

	// validate token
	payload, err := utils.VerifyToken(tknFromParam)
	if err != nil {
		if errors.Is(err, errors.New("error parse token: token has invalid claims: token is expired")) {
			c.authUseCase.DeleteUser(r.Context(), payload.(string))
			return
		} else {
			c.log.Info("token not valid", zap.Error(err))
			utils.WriteErrorJson(w, err, http.StatusUnauthorized)
			return
		}
	}

	response, err := c.authUseCase.VerifyUser(r.Context(), payload.(string))
	if err != nil {
		c.log.Error("failed to verify user", zap.Error(err))
		utils.WriteErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	// create a new token
	token, err := utils.CreateToken(time.Duration(30)*time.Minute, response.ID)
	if err != nil {
		c.log.Error("Error create token", zap.Error(err))
		utils.WriteErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJsonBody(w, http.StatusOK, utils.WebResponse[*models.UserResponse]{
		Message: "User verified",
		Data:    response,
		Token:   token,
	})

}
