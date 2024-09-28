package routing

import (
	"database/sql"

	"github.com/labstack/echo"
	service "github.com/mBuergi86/devseconnect/internal/application"
	"github.com/mBuergi86/devseconnect/internal/domain/handler"
	"github.com/mBuergi86/devseconnect/internal/domain/repository"
)

func SetupRouting(e *echo.Echo, db *sql.DB) {
	userRepo := repository.NewPostgresUsersRepository(db)

	userService := service.NewUserService(userRepo)

	userHandler := &handler.UserHandler{Service: userService}

	e.GET("/users/:id", userHandler.GetUsersByID)
	e.GET("/users/email/:email", userHandler.FindUserByEmail)
	e.POST("/users", userHandler.CreateUser)
	e.PUT("users/:id", userHandler.UpdateUser)
	e.DELETE("users/:id", userHandler.DeleteUser)
}
