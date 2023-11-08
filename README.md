# ods-pipeline-npm

[![Tests](https://github.com/opendevstack/ods-pipeline-npm/actions/workflows/main.yaml/badge.svg)](https://github.com/opendevstack/ods-pipeline-npm/actions/workflows/main.yaml)

Tekton task for use with [ODS Pipeline](https://github.com/opendevstack/ods-pipeline) to build applications with NPM.

## Usage

```yaml
tasks:
- name: build
  taskRef:
    resolver: git
    params:
    - { name: url, value: https://github.com/opendevstack/ods-pipeline-npm.git }
    - { name: revision, value: v0.2.0 }
    - { name: pathInRepo, value: tasks/build.yaml }
    workspaces:
    - { name: source, workspace: shared-workspace }
```

See the [documentation](https://github.com/opendevstack/ods-pipeline-npm/blob/main/docs/build.adoc) for details and available parameters.

## About this repository

`docs` and `tasks` are generated directories from recipes located in `build`. See the `Makefile` target for how everything fits together.
