package api

import (
	"encoding/json"
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

type API struct {
	router *mux.Router
}

func (api *API) Endpoints() {
	api.router.HandleFunc("/api/v1/containers", api.containers).Methods(http.MethodGet)
	api.router.HandleFunc("/api/v1/containers/{id}", api.idContainer).Methods(http.MethodGet)
}

func (api *API) containers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	list, err := cli.JsonPs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(list)
}

func (api *API) idContainer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	ct, err := cli.JsonCtId(vars[`id`])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(ct)
}
