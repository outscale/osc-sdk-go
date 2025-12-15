package examples_test

import (
	"fmt"
	"testing"
	"time"

	"dario.cat/mergo"
	"github.com/outscale/osc-sdk-go/v3/pkg/oks"
	"github.com/outscale/osc-sdk-go/v3/pkg/profile"
	"github.com/outscale/osc-sdk-go/v3/pkg/utils"
)

func TestProject(t *testing.T) {
	userProfile, err := profile.NewProfileFromStandardConfiguration("", "")
	if err != nil {
		panic(err)
	}

	client, err := oks.NewClient(userProfile, utils.WithLogging(&testingLogger{t}))
	if err != nil {
		panic(err)
	}

	ctx := t.Context()

	_, err = client.GetProject(ctx, "a-non-existing-project")
	if err == nil {
		panic("a-non-existing-project should not exist")
	}
	if !oks.IsNotFound(err) {
		panic(fmt.Errorf("a-non-existing-project should be not found: %w", err))
	}

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

	var projectID string
	createProject, err := client.CreateProject(ctx, project)
	if err != nil {
		if oks.IsConflict(err) {
			readtProject, errList := client.ListProjects(ctx, &oks.ListProjectsParams{Name: &name})
			if errList != nil {
				panic(err)
			}
			if len(readtProject.Projects) != 1 {
				panic(fmt.Errorf("expected 1 project, got %d", len(readtProject.Projects)))
			}
			projectID = readtProject.Projects[0].Id
		} else {
			panic(err)
		}
	} else {
		projectID = createProject.Project.Id
	}

	for {
		readProject, err := client.GetProject(ctx, projectID)
		if err != nil {
			panic(err)
		}

		if readProject.Project.Status == oks.ProjectStatusReady {
			break
		}

		time.Sleep(10 * time.Second)
	}

	_, err = client.DeleteProject(ctx, projectID)
	if err != nil {
		panic(err)
	}
}
