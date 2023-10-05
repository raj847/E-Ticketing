package handler

import (
	"errors"
	"eticketing/entity"
	"eticketing/handler/request"
	"eticketing/handler/response"
	"eticketing/service"
	"eticketing/validate"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(
	userService *service.UserService,
) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (u *UserHandler) Create(c echo.Context) error {

	var userReq request.User
	err := c.Bind(&userReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: "failed to read json request",
					Code:    "BAD_REQUEST",
				},
			},
		})
	}
	err = validate.Validate(userReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: err.Error(),
					Code:    "USER_INVALID",
				},
			},
		})
	}
	user := entity.User{
		Username: userReq.Username,
		Password: userReq.Password,
	}
	err = u.userService.Register(c.Request().Context(), &user)
	if err != nil {
		if errors.Is(err, service.ErrUserAlreadyExists) {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Errors: []response.ErrorDetail{
					{
						Message: "failed to create user",
						Code:    "USER-ALREADY-EXIST_CREATE-ERROR",
					},
				},
			})
		}
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: "failed to create user",
					Code:    "USER_CREATE-ERROR",
				},
			},
		})
	}
	res := response.BuildUser(user)
	return c.JSON(http.StatusCreated, res)
}

func (u *UserHandler) Login(c echo.Context) error {

	var userReq request.Login
	err := c.Bind(&userReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: "failed to read json request",
					Code:    "BAD_REQUEST",
				},
			},
		})
	}
	User := entity.User{
		Username: userReq.Username,
		Password: userReq.Password,
	}
	res, err := u.userService.LoginUser(c.Request().Context(), &User)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Errors: []response.ErrorDetail{
					{
						Message: "failed to login user",
						Code:    "USER-NOT-FOUND_LOGIN-ERROR",
					},
				},
			})
		}
		if errors.Is(err, service.ErrUserPasswordDontMatch) {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Errors: []response.ErrorDetail{
					{
						Message: "failed to login user",
						Code:    "PASSWORD-NOT-MATCH_LOGIN-ERROR",
					},
				},
			})
		}
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Errors: []response.ErrorDetail{
				{
					Message: "failed to login user",
					Code:    "USER_LOGIN-ERROR",
				},
			},
		})
	}
	ress := response.BuildUser(res)

	expiresAt := time.Now().Add(5 * time.Hour)
	claims := entity.Claims{
		UserID:   res.ID,
		Username: res.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	tokenString, _ := token.SignedString([]byte(os.Getenv("SIGNED_TOKEN")))

	resp := map[string]any{
		"data":  ress,
		"token": tokenString,
	}
	return c.JSON(http.StatusCreated, resp)
}
