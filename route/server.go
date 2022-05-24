package route

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var (
	store    *session.Store
	AUTH_KEY string = "authenticated"
	USER_ID  string = "user_id"
)

func GetServer() *fiber.App {
	server := fiber.New()

	store = session.New(session.Config{
		CookieHTTPOnly: true,
		Expiration:     time.Hour * 4,
	})

	server.Use(NewMiddleware(), cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
		AllowHeaders:     "Access-Control-Allow-Origin, Content-Type, Origin, Accept",
	}))

	path, err := os.Getwd()
	if err == nil {
		server.Static("/", path+"/../public")
	}

	server.Post("/register", Register)
	server.Post("/login", Login)
	server.Post("/logout", Logout)
	server.Get("/ping", Ping)
	server.Get("/user", GetUser)

	return server
}
