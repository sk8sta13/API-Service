package main

import (
	"log"
	"net/http"

	"github.com/sk8sta13/API-Service/configs"
	"github.com/sk8sta13/API-Service/internal/entity"
	"github.com/sk8sta13/API-Service/internal/infra/database"
	"github.com/sk8sta13/API-Service/internal/infra/web/handlers"

	_ "github.com/sk8sta13/API-Service/docs"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title          Simple API in GO
// @version        1.0
// @description    Simple API for learning using GO language.
// @termsOfService http://swagger.io/terms/

// @contact.name  Marcelo Soto Campos
// @contact.email sk8sta13@gmail.com
// @contact.url   https://github.com/sk8sta13

// @license.name sk8sta13
// @license.url  https://github.com/sk8sta13

// @host                       localhost:8000
// @BasePath                   /
// @securityDefinitions.apikey ApiKeyAuth
// @in                         header
// @name                       Authorization

func main() {
	config, err := configs.LoadConfigs(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("service.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.User{}, &entity.Product{})

	ProductDB := database.NewProduct(db)
	ProductHandler := handlers.NewProductHandler(ProductDB)

	UserDB := database.NewUser(db)
	UserHandler := handlers.NewUserHandler(UserDB, config.TokenAuth, config.JWTttl)

	r := chi.NewRouter()
	//r.Use(middleware.Logger)
	r.Use(LogRequest)

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/", ProductHandler.CreateProduct)
		r.Get("/", ProductHandler.GetProducts)
		r.Get("/{id}", ProductHandler.GetProduct)
		r.Put("/{id}", ProductHandler.UpdateProduct)
		r.Delete("/{id}", ProductHandler.DeleteProduct)
	})

	r.Post("/users", UserHandler.CreateUser)
	r.Post("/users/generate_token", UserHandler.GetToken)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":8000", r)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
