package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	api    *API
	router *mux.Router
}

func StartServer(port uint, ip string) {
	srv := new(server)
	srv.router = mux.NewRouter()
	srv.api = &API{router: srv.router}
	srv.api.Endpoints()

	log.Printf("API server has been started on: %s:%v", ip, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%v", ip, port), srv.router))
}
