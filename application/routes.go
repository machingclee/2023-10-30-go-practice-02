package application

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"orders-api/handler"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	router.Route("/orders", loadOrderRoutes)
	return router
}

func loadOrderRoutes(router chi.Router) {
	orderHandeler := &handler.Order{}
	router.Post("/", orderHandeler.Create)
	router.Get("/", orderHandeler.List)
	router.Get("/{id}", orderHandeler.GetByID)
	router.Put("/{id}", orderHandeler.UpdateByID)
	router.Delete("/{id}", orderHandeler.DeleteByID)
}
