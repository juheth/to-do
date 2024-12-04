package task

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type (
	Controller func(c *gin.Context)
	EndPoints  struct {
		CreateTask     Controller
		GetAllTaskById Controller
		GetAllTask     Controller
		UpDateTask     Controller
	}
	CreateReq struct {
		Name        string `form:"name"`
		Description string `form:"description"`
		DueDate     string `form:"due_date"`
		UserID      string `form:"user_id"`
	}
	UpDateReq struct {
		Name        string `form:"name"`
		Description string `form:"description"`
		DueDate     string `form:"due_date"`
		UserID      string `form:"id_user"`
		Status      bool   `form:"status"`
		Create      string `form:"create_at"`
		UpDate      string
	}
)

func MakeEnponints(s Service) EndPoints {
	return EndPoints{
		CreateTask:     makeCreateTask(s),
		GetAllTaskById: makeGetAllTaskById(s),
		UpDateTask:     makeUpdateTask(s),
		GetAllTask:     makeGetAllTask(s),
	}
}

func makeCreateTask(s Service) Controller {
	return func(c *gin.Context) {
		var req CreateReq
		c.ShouldBind(&req)

		if req.Name == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "name is required"})
			return
		}
		if req.DueDate == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "due_date is required"})
			return
		}
		if req.UserID == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "user_id is required"})
			return
		}

		err := s.GetUserById(req.UserID)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "user not found", "err": err})
			return
		}
		dueDate, err := time.Parse("2006-01-02 15:04:05", req.DueDate)

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": 500, "err": err})
			return
		}

		err = s.Create(req.Name, req.Description, req.UserID, dueDate)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": 500, "err": err})
			return
		}
		c.IndentedJSON(http.StatusCreated, gin.H{"status": 201, "data": req})
	}
}

func makeGetAllTaskById(s Service) Controller {
	return func(c *gin.Context) {
		userId := c.Param("id")
		if userId == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "user_id is required"})
			return
		}
		tasks, err := s.GetAllTaskById(userId)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": 500, "err": err})
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{"status": 200, "data": tasks})

	}
}

func makeGetAllTask(s Service) Controller {
	return func(c *gin.Context) {
		tasks, err := s.GetAllTask()
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": 500, "err": err})
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{"status": 200, "data": tasks})

	}
}

func makeUpdateTask(s Service) Controller {
	return func(c *gin.Context) {
		taskId := c.Param("id")
		if taskId == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "task_id is required"})
			return
		}
		task, err := s.GetTaskById(taskId)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "user not found", "err": err})
			return
		}

		var req UpDateReq
		c.ShouldBind(&req)

		if req.Name == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "name is required"})
			return
		}
		if req.Description == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "description is required"})
			return
		}
		if req.DueDate == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "due_date is required"})
			return
		}
		if req.UserID == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "user_id is required"})
			return
		}
		err = s.GetUserById(req.UserID)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "user not found", "err": err})
			return
		}
		dueDate, err := time.Parse("2006-01-02 15:04:05Z", req.DueDate)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": 500, "err": err})
			fmt.Println(err)
			return
		}
		req.Create = task.Create_at.Format("2006-01-02 15:04:05")

		update, err := s.UpDateTask(taskId, req.Name, req.Description, req.UserID, dueDate, task.Create_at, req.Status)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": 500, "err": err})
			return
		}
		req.UpDate = update
		c.IndentedJSON(http.StatusOK, gin.H{"status": 200, "data": req})
	}
}
