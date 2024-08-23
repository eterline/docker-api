package cli

import (
	"encoding/json"

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

func JsonPs() (PsList, error) {
	var list PsList
	err := json.Unmarshal([]byte(DockerPs()), &list)
	if err != nil {
		return list, nil
	}
	return list, nil
}

func JsonCtId(id string) (Container, error) {
	var res Container
	list, err := JsonPs()
	if err != nil {
		return res, err
	}
	for _, i := range list {
		if i.ID == id {
			res := i
			return res, err
		}
	}
	return res, nil
}
