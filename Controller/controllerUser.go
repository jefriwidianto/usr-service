package Controller

import (
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"net/http"
	"time"
	"usr-service/Controller/Dto/Response"
	"usr-service/Repository"
	"usr-service/Services/Jwt"
	"usr-service/Utils"
)

type UserInterface interface {
	Login(ctx echo.Context) (err error)
}

func (c Controller) Login(ctx echo.Context) (err error) {
	username, password, ok := ctx.Request().BasicAuth()
	if !ok || (username == "" || password == "") {
		c.Log.Error(http.StatusText(http.StatusBadRequest), zap.String("_method", ctx.Request().Method), zap.String("_url_request", ctx.Request().RequestURI))
		return ctx.JSON(http.StatusBadRequest, &Response.Responses{
			Message: http.StatusText(http.StatusBadRequest),
		})
	}

	resp, err := Repository.ApplicationRepository.User.LoginUser(ctx.Request().Context(), username, Utils.GenerateHashPass(password))
	if err != nil {
		if err == sql.ErrNoRows {
			c.Log.Error("invalid username or password", zap.String("_method", ctx.Request().Method), zap.String("_url_request", ctx.Request().RequestURI))
			return ctx.JSON(http.StatusUnauthorized, &Response.Responses{
				Message: "invalid username or password",
			})
		}

		c.Log.Error(http.StatusText(http.StatusInternalServerError), zap.String("_method", ctx.Request().Method), zap.String("_url_request", ctx.Request().RequestURI))
		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
			Message: http.StatusText(http.StatusInternalServerError),
		})
	}

	location, _ := time.LoadLocation("Asia/Jakarta")
	param := jwt.MapClaims{
		"userId":      resp.Id,
		"name":        resp.Username,
		"email":       resp.Email,
		"phoneNumber": resp.Email,
		"expired":     time.Now().In(location).Add(time.Hour * 12).Format("2006-01-02 15:04:05"),
	}

	token, err := Jwt.AuthKey().Encode(param)
	if err != nil {
		c.Log.Error(http.StatusText(http.StatusInternalServerError), zap.String("_method", ctx.Request().Method), zap.String("_url_request", ctx.Request().RequestURI))
		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
			Message: http.StatusText(http.StatusInternalServerError),
		})
	}

	return ctx.JSON(http.StatusOK, &Response.Responses{
		Data:    token,
		Message: http.StatusText(http.StatusOK),
	})
}
