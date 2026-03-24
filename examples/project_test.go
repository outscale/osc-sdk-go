package examples_test

import (
	"testing"
	"time"

	"dario.cat/mergo"
	"github.com/outscale/osc-sdk-go/v3/pkg/oks"
	"github.com/outscale/osc-sdk-go/v3/pkg/options"
	"github.com/outscale/osc-sdk-go/v3/pkg/profile"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Steps done in this test:
// 1. Read a non-existing project and validate the not-found error.
// 2. List projects by name.
// 3. Read the project template.
// 4. Create a project from the template.
// 5. Wait for the project to be ready.
// 6. Delete the project.
func TestProject(t *testing.T) {
	userProfile, err := profile.New()
	require.NoError(t, err)

	client, err := oks.NewClient(userProfile, options.WithLogging(&testingLogger{t}))
	require.NoError(t, err)

	ctx := t.Context()

	_, err = client.GetProject(ctx, "a-non-existing-project")
	require.Error(t, err, "a-non-existing-project should not exist")
	assert.True(t, oks.IsNotFound(err), "a-non-existing-project should be not found: %v", err)

	name := "osc-sdk-go-test-" + RandomString(10)
	_, err = client.ListProjects(ctx, &oks.ListProjectsParams{Name: &name})
	require.NoError(t, err)

	projectTemplate, err := client.GetProjectTemplate(ctx)
	require.NoError(t, err)
	require.NotNil(t, projectTemplate.Template)

	project := oks.CreateProjectJSONRequestBody{Name: name}
	err = mergo.Merge(&project, projectTemplate.Template)
	require.NoError(t, err)

	var projectID string
	createProject, err := client.CreateProject(ctx, project)
	if err != nil {
		if oks.IsConflict(err) {
			readtProject, errList := client.ListProjects(ctx, &oks.ListProjectsParams{Name: &name})
			require.NoError(t, errList)
			require.Len(t, readtProject.Projects, 1)
			projectID = readtProject.Projects[0].Id
		} else {
			require.NoError(t, err)
		}
	} else {
		projectID = createProject.Project.Id
	}
	assert.NotEmpty(t, projectID)

	deleted := false
	defer func() {
		if deleted || projectID == "" {
			return
		}

		_, _ = client.DeleteProject(ctx, projectID)
	}()

	for {
		readProject, err := client.GetProject(ctx, projectID)
		require.NoError(t, err)
		assert.Equal(t, name, readProject.Project.Name)

		if readProject.Project.Status == oks.ProjectStatusReady {
			break
		}

		time.Sleep(10 * time.Second)
	}

	_, err = client.DeleteProject(ctx, projectID)
	require.NoError(t, err)
	deleted = true
}
