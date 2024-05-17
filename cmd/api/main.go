package main 

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/Pranaenae/goauthapi/internal/handlers"
	log "github.com/sirupsen/logrus" //aliased as log
)

func main(){
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting Go API Service...")

	err := http.ListenandServer("localhost:8080",r)

	if err!=nil {
		log.Error(err)
	}
}

