package server

import (
	"github.com/assaidy/bookstore/internals/handlers"
	// jwtware "github.com/gofiber/contrib/jwt"
)

func (s *FiberServer) RegisterRoutes() {
	var (
		userH = handlers.NewUserHandler(s.db)
	)

	s.Post("/users/register", userH.HandleRegisterUser)
	s.Post("/users/login", userH.HandleLoginUser)
	s.Get("/users", userH.HandleGetAllUsers)
	s.Get("/users/:id<int>", userH.HandleGetUserById)
	s.Put("/users/:id<int>", userH.HandleUpdateUserById)
	s.Delete("/users/:id<int>", userH.HandleDeleteUserById)

	// NOTE: this validates before our logging handler ie. it will not log errors
	// it sends 401: Invalid or expired JWT
	// s.Use(jwtware.New(jwtware.Config{
	// 	SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
	// }))

	// restricted here...
}
