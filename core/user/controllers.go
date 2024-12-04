package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juheth/to-do/core/jwt"
	"github.com/juheth/to-do/core/models"
)

type (
	Controller func(c *gin.Context)
	EndPoints  struct {
		RegisterUser Controller
		LoginUser    Controller
		GetUser      Controller
	}
	LoginReq struct {
		Email    string `form:"email"`
		Password string `form:"password"`
	}
	UserRes struct {
		Id    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	RegisterReq struct {
		Name     string `form:"name"`
		Email    string `form:"email"`
		Password string `form:"password"`
	}
)

func MakeEnponints(s Service) EndPoints {
	return EndPoints{
		RegisterUser: makeRegisterUser(s),
		LoginUser:    makeLoginUser(s),
		GetUser:      makeGetUser(s),
	}
}

func makeRegisterUser(s Service) Controller {
	return func(c *gin.Context) {
		var req RegisterReq
		// Validar el payload
		if err := c.ShouldBind(&req); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": err.Error()})
			return
		}

		// Validaciones adicionales
		if req.Name == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Name is required"})
			return
		}
		if req.Email == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Email is required"})
			return
		}
		if !Service.IsValidMail(s, req.Email) {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid email format"})
			return
		}
		if len(req.Password) < 8 {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Password must be at least 8 characters"})
			return
		}

		// Verificar si el email ya existe
		existingUser, _ := s.GetUserByMail(req.Email)
		if existingUser.Email != "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Email already exists"})
			return
		}

		// Crear usuario
		err := s.RegisterUser(req.Name, req.Email, req.Password)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"status":  500,
				"message": "Error registering user",
				"error":   err.Error(),
			})
			return
		}

		// Generar el token JWT
		token, err := jwt.JWT(&models.User{Name: req.Name, Email: req.Email}) // Ajusta según la implementación de JWT
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"status":  500,
				"message": "Error generating token",
				"error":   err.Error(),
			})
			return
		}

		// Respuesta exitosa
		c.IndentedJSON(http.StatusCreated, gin.H{
			"status": 201,
			"token":  token,
		})
	}
}

func makeLoginUser(s Service) Controller {
	return func(c *gin.Context) {
		var req LoginReq
		err := c.ShouldBind(&req)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 500, "message": err})
			return
		}
		if req.Email == "" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "email is required"})
			return
		}
		if !Service.IsValidMail(s, req.Email) {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "email is not valid "})
			return
		}
		user, err := s.GetUserByMail(req.Email)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "the email no exists"})
			return
		}
		if len(req.Password) < 8 {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": ""})
			return
		}

		valid, err := s.ValidPassword(req.Email, req.Password)
		if !valid {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "invalid password"})
			return
		}

		data := UserRes{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}
		c.IndentedJSON(http.StatusAccepted, gin.H{"status": 202, "data": data})

	}
}
func makeGetUser(s Service) Controller {
	return func(c *gin.Context) {
		users, err := s.GetUser()
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": 500, "message": err})
			return
		}
		c.IndentedJSON(http.StatusAccepted, gin.H{"status": 202, "users": users})
	}
}
