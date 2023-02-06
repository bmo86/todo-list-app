package handlers

import (
	"encoding/base64"
	"net/http"
	"time"
	"todo-api/config"
	modelsapp "todo-api/models/models-app"
	modelstoken "todo-api/models/models-token"
	repoapp "todo-api/repository/repo-app"
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
				Check:       req.Check,
				Status:      req.Status,
				DateFinish:  req.DateFinish,
				Image:       string(image),
			}

			idRes, err := repoapp.CreateTask(&taskSend)
			if err != nil {
				c.JSON(http.StatusInternalServerError, msgError(err))
				return
			}

			s.Hub().BroadCast(taskSend, nil)

			c.JSON(http.StatusCreated, gin.H{
				"id_create_task": idRes,
				"message":        "create Task",
			})
		} else {
			c.JSON(http.StatusInternalServerError, msgError(err))
		}

	}
}
