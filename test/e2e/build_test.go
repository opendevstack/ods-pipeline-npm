package e2e

import (
	"bytes"
	"path/filepath"
	"testing"

	ott "github.com/opendevstack/ods-pipeline/pkg/odstasktest"
	"github.com/opendevstack/ods-pipeline/pkg/pipelinectxt"
	ttr "github.com/opendevstack/ods-pipeline/pkg/tektontaskrun"
	tekton "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
)

func TestNPMBuildTask(t *testing.T) {
	if err := runTask(
		ttr.WithStringParams(map[string]string{
			"cache-build": "false",
		}),
		ott.WithGitSourceWorkspace(t, "../testdata/workspaces/typescript-sample-app", namespaceConfig.Name),
		ttr.AfterRun(func(config *ttr.TaskRunConfig, run *tekton.TaskRun, logs bytes.Buffer) {
			wd := config.WorkspaceConfigs["source"].Dir
			ott.AssertFilesExist(t, wd,
				filepath.Join(pipelinectxt.XUnitReportsPath, "report.xml"),
				filepath.Join(pipelinectxt.CodeCoveragesPath, "clover.xml"),
				filepath.Join(pipelinectxt.CodeCoveragesPath, "coverage-final.json"),
				filepath.Join(pipelinectxt.CodeCoveragesPath, "lcov.info"),
				filepath.Join(pipelinectxt.LintReportsPath, "report.txt"),
				"dist/src/index.js",
				"node_modules",
				"package.json",
				"package-lock.json",
			)
		}),
	); err != nil {
		t.Fatal(err)
	}
}
