package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gaba-bouliva/buyit/services/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)


type APIServer struct {
	Addr			string
	DB				*sql.DB
	InfoLog	 	*log.Logger
	ErrorLog  *log.Logger
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	infoLog 	:= log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog 	:= log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return &APIServer{
		Addr: addr,
		DB: db,
		InfoLog: infoLog,
		ErrorLog: errorLog,
	}
}

func (srv *APIServer)Run() error {
	srv.InfoLog.Println("API listening on port", srv.Addr)

	s := &http.Server{
		Addr: fmt.Sprint(srv.Addr),
		Handler: srv.routes(),
	}

	return s.ListenAndServe()
}

func (srv *APIServer)routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:				[]string{"https://*", "http://*"},
		AllowedMethods:				[]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:				[]string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:				[]string{"Link"},
		AllowCredentials: 		true,
		MaxAge: 							300,
	}))


	apiRouter := chi.NewRouter()

	mux.Mount("/api/v1", apiRouter)

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(apiRouter)


	return mux
}