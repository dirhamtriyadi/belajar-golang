package main

import (
	"belajar-golang/internal/delivery/http"
	"belajar-golang/internal/infra/mysql"
	"belajar-golang/internal/repository"
	"belajar-golang/internal/usecase"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize database connection sqlite
	// db, err := sqlite.InitDB()
	// if err != nil {
	// 	log.Fatalf("failed to initialize db: %v", err)
	// }

	// Initialize database connection mysql
	db, err := mysql.InitDB()
	if err != nil {
		log.Fatalf("failed to initialize db: %v", err)
	}

	// Initialize repository
	userRepository := repository.NewUserRepository(db)

	// Initialize usecase
	userUsecase := usecase.NewUserUsecase(userRepository)

	// Initialize fiber app
	app := fiber.New()

	// Initialize user handler
	http.NewUserHandler(app, userUsecase)

	// Run the server
	log.Fatal(app.Listen(":8080"))
}
