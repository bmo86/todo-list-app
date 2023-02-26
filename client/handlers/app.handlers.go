package handlers

import (
	"encoding/base64"
	"net/http"
	"strconv"
	"time"
	"todo-api/config"
	modelsapp "todo-api/models/models-app"
	modelscache "todo-api/models/models-cache"
	modelstoken "todo-api/models/models-token"
	repoapp "todo-api/repository/repo-app"
	repocache "todo-api/repository/repo-cache"
	"todo-api/server"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandlerCreateTask(s server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {

		token, err := config.Token(s, c.Writer, c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, msgError(err))
			return
		}

		if claims, ok := token.Claims.(*modelstoken.AppClaims); ok && token.Valid {

			req := modelsapp.Request_Task{}

			if err := c.ShouldBind(&req); err != nil {
				c.JSON(http.StatusBadRequest, msgError(err))
				return
			}

			image, err := base64.StdEncoding.DecodeString(string(req.Image))
			if err != nil {
				c.JSON(http.StatusInternalServerError, msgError(err))
				return
			}

			taskSend := modelsapp.Task{
				Model: gorm.Model{
					CreatedAt: time.Now().UTC(),
					UpdatedAt: time.Now().UTC(),
				},
				User_id:     claims.IdUser,
				Title:       req.Title,
				Description: req.Description,
				Status:      req.Status,
				DateFinish:  req.DateFinish,
				Image:       image,
			}

			idRes, err := repoapp.CreateTask(&taskSend)
			if err != nil {
				c.JSON(http.StatusInternalServerError, msgError(err))
				return
			}

			s.Hub().BroadCast(taskSend, nil)

			c.JSON(http.StatusCreated, gin.H{
				"id":      idRes,
				"message": "create Task",
			})
		} else {
			c.JSON(http.StatusInternalServerError, msgError(err))
		}

	}
}

func HandlerDeleteTask(s server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := config.Token(s, c.Writer, c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, msgError(err))
			return
		}

		if _, ok := token.Claims.(*modelstoken.AppClaims); ok && token.Valid {
			idReq := c.Param("id")
			if idReq == "" {
				c.JSON(http.StatusBadRequest, msgError("Not Found id"))
				return
			}

			id, err := strconv.ParseInt(idReq, 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, msgError(err))
				return
			}

			if err = repoapp.DeleteTask(uint(id)); err != nil {
				c.JSON(http.StatusInternalServerError, msgError(err))
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "Delete Taks",
			})

		} else {
			c.JSON(http.StatusInternalServerError, msgError(err))
			return
		}

	}
}

func HandlerUpdateTask(s server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := config.Token(s, c.Writer, c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, msgError(err))
			return
		}

		if _, ok := token.Claims.(*modelstoken.AppClaims); ok && token.Valid {
			idReq := c.Param("id")

			if idReq == "" {
				c.JSON(http.StatusBadRequest, msgError("Not Found id"))
				return
			}

			id, err := strconv.ParseInt(idReq, 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, msgError(err))
				return
			}

			var reqData modelsapp.Task

			if err = c.ShouldBind(&reqData); err != nil {
				c.JSON(http.StatusBadRequest, msgError(err))
				return
			}

			taskRes := modelsapp.Task{
				Title:       reqData.Title,
				Description: reqData.Description,
				Image:       reqData.Image,
				Model:       gorm.Model{UpdatedAt: reqData.UpdatedAt},
				Status:      reqData.Status,
			}

			idRes, err := repoapp.UpdateTask(uint(id), &taskRes)
			if err != nil {
				c.JSON(http.StatusInternalServerError, msgError(err))
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "update task",
				"id":      idRes,
			})

		}

	}
}

func HandlerGetTasks(s server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := config.Token(s, c.Writer, c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, msgError(err))
			return
		}

		if _, ok := token.Claims.(*modelstoken.AppClaims); ok && token.Valid {
			pageRes := c.Param("page")
			if pageRes == "" {
				pageRes = "1"
			}

			res, cacheHit, err := repocache.GetDataTasks(c.Request.Context(), pageRes)
			if err != nil {
				c.JSON(http.StatusInternalServerError, msgError(err))
				return
			}
			c.JSON(http.StatusOK, modelscache.CacheResponse{Cache: cacheHit, Data: res})

		} else {
			c.JSON(http.StatusInternalServerError, msgError(err))
		}
	}
}

func HandlerGetTask(s server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := config.Token(s, c.Writer, c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, msgError(err))
			return
		}

		if _, ok := token.Claims.(*modelstoken.AppClaims); ok && token.Valid {
			idReq := c.Param("id")
			if idReq == "" {
				c.JSON(http.StatusBadRequest, msgError("Not Found id"))
				return
			}

			res, cacheHit, err := repocache.GetDataTask(idReq)
			if err != nil {
				c.JSON(http.StatusInternalServerError, msgError(err))
				return
			}
			c.JSON(http.StatusOK, modelscache.CacheResponseOnlyOne{Cache: cacheHit, Data: res})
		} else {
			c.JSON(http.StatusInternalServerError, msgError(err))
			return
		}
	}
}
