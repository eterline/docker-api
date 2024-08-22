package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/eterline/docker-api/internal/cli"
	"github.com/gorilla/mux"
)

type Container struct {
	Command      string `json:'Command'`
	CreatedAt    string `json:'CreatedAt'`
	ID           string `json:'ID'`
	Image        string `json:'Image'`
	Labels       string `json:'Labels'`
	LocalVolumes string `json:'LocalVolumes'`
	Mounts       string `json:'Mounts'`
	Names        string `json:'Names'`
	Networks     string `json:'Networks'`
	Port         string `json:'Port'`
	Ports        string `json:'Ports'`
	RunningFor   string `json:'RunningFor'`
	Size         string `json:'Size'`
	State        string `json:'State'`
	Status       string `json:'Status'`
}

type PsList []Container

func StartServe(port uint, ip string) {
	router := mux.NewRouter()

	router.HandleFunc("/api/containers", getPosts).Methods("GET")
	router.HandleFunc("/api/containers/{id}", idContainer).Methods("GET")

	log.Print("API server has been started.")
	err := http.ListenAndServe(fmt.Sprintf("%s:%v", ip, port), router)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	list := cli.JsonPs()
	json.NewEncoder(w).Encode(list)
}

func idContainer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	ct := cli.JsonCtId(vars[`id`])
	json.NewEncoder(w).Encode(ct)
}
