package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"
	"todo-api/config"
	modelstoken "todo-api/models/models-token"
	modelsusr "todo-api/models/models-usr"
	repocache "todo-api/repository/repo-cache"
	repousr "todo-api/repository/repo-usr"
	"todo-api/server"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

const (
	HASH_CONST = 8
)

func msgError(msg interface{}) gin.H {
	return gin.H{
		"message": msg,
	}
}

func HandlerWs(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		s.Hub().HandlerWs(ctx.Writer, ctx.Request)
	}
}

func HandlerSingUp(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := modelsusr.SingUp_Request{}
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, msgError(err))
			return
		}

		hashPass, err := bcrypt.GenerateFromPassword([]byte(req.Pass), HASH_CONST)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, msgError(err))
			return
		}

		var usr = modelsusr.SingUp_Request{
			Name:     req.Name,
			Lastname: req.Lastname,
			Email:    req.Email,
			Pass:     string(hashPass),
			Status:   true,
			Position: req.Position,
		}

		id, err := repousr.SingUp(&usr)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, msgError(err))
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"message": "User Created",
			"id_user": id,
		})
	}
}

func HandlerLogin(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		req := modelsusr.Login_Request{}

		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, msgError(err))
			return
		}

		usr, err := repousr.GetUsrByEmail(req.Email)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, msgError(err))
			return
		}

		if usr == nil {
			ctx.JSON(http.StatusUnauthorized, msgError(err))
			return
		}

		if err = bcrypt.CompareHashAndPassword([]byte(usr.Pass), []byte(req.Pass)); err != nil {
			ctx.JSON(http.StatusUnauthorized, msgError(err))
			return
		}

		claims := modelstoken.AppClaims{
			IdUser:   usr.ID,
			Position: usr.Position,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(2 * time.Hour * 24).Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenString, err := token.SignedString([]byte(s.Config().JWT_SECRET))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, msgError(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Token": tokenString,
		})

	}
}

func HandlerMe(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token, err := config.Token(s, ctx.Writer, ctx.Request)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, msgError(err))
			return
		}

		if claims, ok := token.Claims.(*modelstoken.AppClaims); ok && token.Valid {
			usr, hit, err := repocache.GetUser_ID(context.Background(), strconv.Itoa(int(claims.IdUser)))
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, msgError(err))
				return
			}

			ctx.JSON(http.StatusInternalServerError, gin.H{
				"cache": hit, "data": usr,
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, msgError(err))
	}
}
