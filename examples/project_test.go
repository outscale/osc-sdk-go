package examples_test

import (
	"context"
	"testing"
	"time"

	"dario.cat/mergo"
	"github.com/outscale/osc-sdk-go/v3/internal/oks"
	"github.com/outscale/osc-sdk-go/v3/pkg/client"
)

func TestProject(t *testing.T) {
	client, err := client.NewOKSClient()
	if err != nil {
		panic(err)
	}

	ctx := context.TODO()

	name := "test"
	_, err = client.ListProjects(ctx, &oks.ListProjectsParams{Name: &name})
	if err != nil {
		panic(err)
	}

	projectTemplate, err := client.GetProjectTemplate(ctx)
	if err != nil {
		panic(err)
	}

	project := oks.CreateProjectJSONRequestBody{Name: name}
	err = mergo.Merge(&project, projectTemplate.Template)
	if err != nil {
		panic(err)
	}

	createProject, err := client.CreateProject(ctx, project)
	if err != nil {
		panic(err)
	}

	for {
		readProject, err := client.GetProject(ctx, createProject.Project.Id)
		if err != nil {
			panic(err)
		}

		if readProject.Project.Status == "ready" {
			break
		}

		time.Sleep(10 * time.Second)
	}

	_, err = client.DeleteProject(ctx, createProject.Project.Id)
	if err != nil {
		panic(err)
	}
}
