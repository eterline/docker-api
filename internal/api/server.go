package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eterline/docker-api/internal/settings"
	"github.com/gorilla/mux"
)

type server struct {
	api    *API
	router *mux.Router
}

var Sets settings.Settings

func StartServer(port uint, ip string) {
	srv := new(server)
	srv.router = mux.NewRouter()
	srv.api = &API{router: srv.router}
	srv.api.router.Use(authMiddleware)
	srv.api.Endpoints()

	log.Printf("API server has been started on: %s:%v", ip, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%v", ip, port), srv.router))
}

func verifyPass(pass string) bool {
	return pass == Sets.Pass
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RemoteAddr, r.RequestURI)
		if !verifyPass(r.Header.Get("Password")) {
			http.Error(w, "Invalid password", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)

	})
}
