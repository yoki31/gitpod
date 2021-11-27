// Copyright (c) 2020 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package ide

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"sigs.k8s.io/e2e-framework/pkg/envconf"
	"sigs.k8s.io/e2e-framework/pkg/features"

	agent "github.com/gitpod-io/gitpod/test/pkg/agent/workspace/api"
	"github.com/gitpod-io/gitpod/test/pkg/integration"
)

func poolTask(task func() (bool, error)) (bool, error) {
	timeout := time.After(5 * time.Minute)
	ticker := time.Tick(20 * time.Second)
	for {
		select {
		case <-timeout:
			return false, errors.New("timed out")
		case <-ticker:
			ok, err := task()
			if err != nil {
				return false, err
			} else if ok {
				return true, nil
			}
		}
	}
}

func TestPythonExtWorkspace(t *testing.T) {
	f := features.New("PythonExtensionWorkspace").
		WithLabel("component", "server").
		Assess("it can run python extension in a workspace", func(_ context.Context, t *testing.T, cfg *envconf.Config) context.Context {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
			defer cancel()

			api := integration.NewComponentAPI(ctx, cfg.Namespace(), cfg.Client())
			t.Cleanup(func() {
				api.Done(t)
			})

			nfo, stopWs, err := integration.LaunchWorkspaceFromContextURL(ctx, "github.com/jeanp413/python-test-workspace", username, api)
			if err != nil {
				t.Fatal(err)
			}
			defer stopWs(true)
			t.Log(">>>>>>>>>>>>>>>>>> before wait for workspace")
			_, err = integration.WaitForWorkspaceStart(ctx, nfo.LatestInstance.ID, api)
			if err != nil {
				t.Fatal(err)
			}

			// _, err = api.GitpodSession(nfo.LatestInstance.ID, integration.WithGitpodUser(username))
			// if err != nil {
			// 	t.Fatal(err)
			// }

			t.Log(">>>>>>>>>>>>>>>>>> before rpc into workspace")
			rsa, closer, err := integration.Instrument(integration.ComponentWorkspace, "workspace", cfg.Namespace(), cfg.Client(), integration.WithInstanceID(nfo.LatestInstance.ID), integration.WithWorkspacekitLift(true))
			if err != nil {
				t.Fatal(err)
			}
			defer rsa.Close()
			integration.DeferCloser(t, closer)
			t.Log(">>>>>>>>>>>>>>>>>> before exec")

			_, err = poolTask(func() (bool, error) {
				var resp agent.ExecResponse
				err = rsa.Call("WorkspaceAgent.Exec", &agent.ExecRequest{
					Dir:     "/workspace/python-test-workspace",
					Command: "test",
					Args: []string{
						"-f",
						"__init_task_done__",
					},
				}, &resp)

				return resp.ExitCode == 0, nil
			})
			if err != nil {
				t.Fatal(err)
			}

			// db, err := api.DB(integration.DBName("gitpod-sessions"))
			// if err != nil {
			// 	t.Fatal(err)
			// }

			// var rawCookieData string
			// err = db.QueryRow("SELECT data FROM sessions LIMIT 1").Scan(&rawCookieData)
			// if err != nil {
			// 	t.Fatal(err)
			// }

			// var cookieData integration.CookieData
			// err = json.Unmarshal([]byte(rawCookieData), &cookieData)
			// if err != nil {
			// 	t.Fatal(err)
			// }


			var resp agent.ExecResponse
			err = rsa.Call("WorkspaceAgent.Exec", &agent.ExecRequest{
				Dir:     "/workspace/python-test-workspace",
				Command: "yarn",
				Args: []string{
					"openvscode-server-test",
					fmt.Sprintf("--endpoint=%s", nfo.LatestInstance.IdeURL),
					"--workspacePath=./src/testWorkspace",
					"--extensionDevelopmentPath=./out",
					"--extensionTestsPath=./out/test/suite",
				},
			}, &resp)
			t.Log(">>>>>>>>>>>>>>>>>> after exec")
			if err != nil {
				t.Fatal(err)
			}
			t.Log(">>>>>>>>>>>>>>>>>> no exec errors")
			t.Log(">>>>>>>>>>>>>>>>>> stdout", resp.Stdout)
			t.Log(">>>>>>>>>>>>>>>>>> stderr", resp.Stderr)
			// if err != nil {
			// 	t.Fatal(err)
			// }
			// for _, f := range ls.Files {
			// 	t.Log(f)
			// }

			return ctx
		}).
		Feature()

	testEnv.Test(t, f)
}
