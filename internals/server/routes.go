package server

import (
	"github.com/assaidy/bookstore/internals/handlers"
	// jwtware "github.com/gofiber/contrib/jwt"
)

func (s *FiberServer) RegisterRoutes() {
	var (
		userH     = handlers.NewUserHandler(s.db)
		categoryH = handlers.NewCategoryHandler(s.db)
		coverH    = handlers.NewCoverHandler(s.db)
		bookH     = handlers.NewBookHandler(s.db)
	)

	s.Post("/user/register", userH.HandleRegisterUser)
	s.Post("/user/login", userH.HandleLoginUser)
	s.Get("/user", userH.HandleGetAllUsers)
	s.Get("/user/:id<int>", userH.HandleGetUserById)
	s.Put("/user/:id<int>", userH.HandleUpdateUserById)
	s.Delete("/user/:id<int>", userH.HandleDeleteUserById)

	s.Post("/category", categoryH.HandleCreateCategory)
	s.Get("/category", categoryH.HandleGetAllCategories)
	s.Put("/category/:id<int>", categoryH.HandleUpdateCategoryById)
	s.Delete("/category/:id<int>", categoryH.HandleDeleteCategoryById)

	s.Post("/cover", coverH.HandleCreateCover) // FIX: delete this routes
	s.Get("/cover/:id<int>", coverH.HandleGetCoverById)
	s.Put("/cover/:id<int>", coverH.HandleUpdateCoverById)
	s.Delete("/cover/:id<int>", coverH.HandleDeleteCoverById) // FIX: delete this routes

	s.Post("/book", bookH.HandleCreateBook)
	s.Get("/book", bookH.HandleGetAllBooks)
	s.Get("/book/:id<int>", bookH.HnadleGetBookById)
	s.Put("/book/:id<int>", bookH.HnadleUpdateBookById)
	s.Delete("/book/:id<int>", bookH.HnadleDeleteBookById)

	// NOTE: this validates before our logging handler ie. it will not log errors
	// it sends 401: Invalid or expired JWT
	// s.Use(jwtware.New(jwtware.Config{
	// 	SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
	// }))

	// restricted here...
}
