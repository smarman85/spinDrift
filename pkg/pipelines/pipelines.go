package pipelines

import (
	"encoding/json"
	f "fmt"
	"github.com/smarman85/spinDrift/pkg/armory"
)

type Pipelines []struct {
	Name string `json:"name"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getPipelineNames(payload []byte) {
	var pipes Pipelines
	err := json.Unmarshal([]byte(payload), &pipes)
	check(err)
	for i := 0; i < len(pipes); i++ {
		f.Println(pipes[i].Name)
	}
}

func ListPipelines(app string) {
	pipelines := armory.ArmoryAPI("applications/" + app + "/pipelines")
	getPipelineNames(pipelines)
}
