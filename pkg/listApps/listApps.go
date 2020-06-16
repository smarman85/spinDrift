package listApps

import (
	"encoding/json"
	f "fmt"
	"github.com/smarman85/spinDrift/pkg/armory"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Applications []struct {
	Name string `json:"name"`
}

func getAppNames(payload []byte) {

	var apps Applications
	err := json.Unmarshal([]byte(payload), &apps)
	check(err)
	for i := 0; i < len(apps); i++ {
		f.Println(apps[i].Name)
	}

}

func GetApps() {
	apps := armory.ArmoryAPI("applications")
	getAppNames(apps)
}
