package main

import (
	f "fmt"
	"github.com/smarman85/spinDrift/pkg/listApps"
	"github.com/smarman85/spinDrift/pkg/pipeConfigs"
	"github.com/smarman85/spinDrift/pkg/pipelines"
	"os"
)

func main() {
	if len(os.Args[1:]) > 0 {
		switch os.Args[1] {
		case "list-applications":
			listApps.GetApps()
		case "list-pipelines":
			pipelines.ListPipelines(os.Args[2])
		case "get-config":
			pipeConfigs.PipelineConfig(os.Args[2], os.Args[3])
		default:
			listApps.GetApps()
		}
	} else {
		f.Println("Nope")
		os.Exit(1)
	}
}
