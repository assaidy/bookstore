package server

import (
	"github.com/assaidy/bookstore/internals/handlers"
	jwtware "github.com/gofiber/contrib/jwt"
	"os"
)

func (s *FiberServer) RegisterRoutes() {
	var (
		userH     = handlers.NewUserHandler(s.db)
		categoryH = handlers.NewCategoryHandler(s.db)
		coverH    = handlers.NewCoverHandler(s.db)
		bookH     = handlers.NewBookHandler(s.db)
		favH      = handlers.NewFavouritesHandler(s.db)
		cartH     = handlers.NewCartHandler(s.db)
	)

	s.Post("/user/register", userH.HandleRegisterUser)
	s.Post("/user/login", userH.HandleLoginUser)

	s.Get("/category", categoryH.HandleGetAllCategories)

	// s.Post("/cover", coverH.HandleCreateCover) // FIX: delete this routes
	s.Get("/cover/:id<int>", coverH.HandleGetCoverById)
	// s.Delete("/cover/:id<int>", coverH.HandleDeleteCoverById) // FIX: delete this routes

	s.Get("/book", bookH.HandleGetAllBooks)
	s.Get("/book/:id<int>", bookH.HnadleGetBookById)

	s.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
	}))

    // TODO: handle admin in jwt token creation
    // TODO: create authenticate func: if user is not admin, check if id param maches token id (from context)
    s.Get("/user", userH.HandleGetAllUsers)
    s.Get("/user/:id<int>", userH.HandleGetUserById)
    s.Put("/user/:id<int>", userH.HandleUpdateUserById)
    s.Delete("/user/:id<int>", userH.HandleDeleteUserById)

    s.Post("/category", categoryH.HandleCreateCategory)
    s.Put("/category/:id<int>", categoryH.HandleUpdateCategoryById)
    s.Delete("/category/:id<int>", categoryH.HandleDeleteCategoryById)

    s.Put("/cover/:id<int>", coverH.HandleUpdateCoverById)

    s.Post("/book", bookH.HandleCreateBook)
    s.Put("/book/:id<int>", bookH.HnadleUpdateBookById)
    s.Delete("/book/:id<int>", bookH.HnadleDeleteBookById)

    s.Post("/user/:uid<int>/favourite/:bid<int>", favH.HandleAddBookToFavourites)
    s.Get("/user/:uid<int>/favourite", favH.HandleGetAllUserFavourites)
    s.Delete("/user/:uid<int>/favourite/:bid<int>", favH.HandleDeleteBookFromFavourites)

    s.Post("/user/:uid<int>/cart", cartH.HandleAddToCart)
    s.Get("/user/:uid<int>/cart", cartH.HandleGetBooksInCart)
    s.Delete("/user/:uid<int>/cart/:bid<int>", cartH.HandleDeleteBookFromCart)
}
