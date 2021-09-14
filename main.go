package main

import (
	"net/http"
	"time"

	"github.com/ihulsbus/cookbook/internal/config"
	h "github.com/ihulsbus/cookbook/internal/handlers"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var (
	cnf = config.Configuration
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{"remote_addr": r.RemoteAddr, "method": r.Method, "uri": r.RequestURI}).Info()

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func main() {
	router := mux.NewRouter()
	router.Use(loggingMiddleware)

	// API versioning setup
	v1 := router.PathPrefix("/v1").Subrouter()

	// All GET routes
	v1get := v1.Methods("GET").Subrouter()
	v1get.Path("/recipes").HandlerFunc(h.RecipeGetAll)
	v1get.Path("/ingredients").HandlerFunc(h.IngredientGetAll)

	// All POST routes
	v1post := v1.Methods("POST").Subrouter()
	v1post.Path("/recipes").HandlerFunc(h.RecipeCreate)
	v1post.Path("/ingredients").HandlerFunc(h.IngredientCreate)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 300 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Info("server available on port 8080")
	log.Info(cnf)
	log.Fatal(srv.ListenAndServe())

}
