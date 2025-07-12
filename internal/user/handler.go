package user

import (
	"net/http"
	"os"
	"strings"

	"splitbill-back/internal/auth"
	"splitbill-back/internal/db"

	"github.com/gin-gonic/gin"
)

func SignupHandler(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := User{
		Email:    input.Email,
		Login:    input.Login,
		Password: input.Password,  
	}

	if err := db.DB.Create(&user).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "User with such data already exists"})
				return 
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created",
		"user":    user,
	})
}

func LoginHandler(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user User
	if err := db.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if user.Password != input.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.SetCookie("token", token, 60*60*24*3, "/", os.Getenv("API_BASE_HOST"), false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Logged in successfully", "user": user})
}

func LogoutHandler (c *gin.Context) {
	c.SetCookie("token", "", -1, "/", os.Getenv("API_BASE_HOST"), false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
}

func GetUsersHandler(c *gin.Context) {
	var users []User
	if err := db.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func MeHandler(c *gin.Context) {
	userID := c.GetUint("user_id")

	var user User
	if err := db.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}


