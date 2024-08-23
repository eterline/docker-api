package cli

import (
	"encoding/json"
	"log"

	utilla "github.com/eterline/utills"
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

var DOCKER_PS = `docker ps --all --no-trunc --format="{{json . }}" | jq -s --tab .`

func DockerPs() []byte {
	return utilla.ExecCmd(DOCKER_PS)
}

func JsonPs() PsList {
	var list PsList
	err := json.Unmarshal([]byte(DockerPs()), &list)
	if err != nil {
		log.Fatal(err.Error())
	}
	return list
}

func JsonCtId(id string) Container {
	var res Container
	for _, i := range JsonPs() {
		if i.ID == id {
			res := i
			return res
		}
	}
	return res
}

func JsonCtNames(names string) Container {
	var res Container
	for _, i := range JsonPs() {
		if i.Names == names {
			res := i
			return res
		}
	}
	return res
}
