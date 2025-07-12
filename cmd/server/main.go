package main

import (
	"fmt"
	"os"
	"splitbill-back/internal/auth"
	"splitbill-back/internal/db"
	"splitbill-back/internal/team"
	"splitbill-back/internal/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err:= godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db.InitDatabase(dsn)

	// 🟢 Move AutoMigrate here to avoid import cycle
	db.DB.AutoMigrate(&user.User{}, &team.Team{})

	r := gin.Default()
	r.Use(cors.New(cors.Config{
			AllowOrigins:     []string{os.Getenv("API_FRONT_URL") }, // frontend origin
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type"},
			AllowCredentials: true,
	}))

	// Public routes
	r.POST("/signup", user.SignupHandler)
	r.POST("/login", user.LoginHandler)


	// Protected routes
	authRoutes := r.Group("/")
	authRoutes.Use(auth.JWTAuthMiddleware())
	authRoutes.GET("/me", user.MeHandler)
	authRoutes.GET("/users", user.GetUsersHandler)
	authRoutes.POST("/logout", user.LogoutHandler)

	authRoutes.POST("/teams", team.CreateTeamHandler)
	authRoutes.GET("/teams/:id", team.GetTeamHandler)

	r.Run(":8080")
}
