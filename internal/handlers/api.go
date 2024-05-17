package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/Pranaenae/goauthapi/internal/middleware"
)
func Handler(r *chi.Mux){
	//Global middleware
	r.Use(chimiddle.StripSlashes) //Ignore trailing slashes

	r.Route("/account", func(router chi.Router)){

		//Authorization middleware
		router.Use(middleware.Authorization){
			
		router.Get("/coins",GetCoinBalance)
		}
	}
}