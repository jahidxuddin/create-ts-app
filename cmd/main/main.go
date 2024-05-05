package main

import (
	"os"
	"path/filepath"

	"github.com/jahidxuddin/create-ts-app/internal/cli"
	"github.com/jahidxuddin/create-ts-app/internal/filesystem"
	"github.com/jahidxuddin/create-ts-app/internal/utils"
)

func main() {
	args := os.Args[1:]

	projectName, err := cli.ParseProjectNameFromArgs(args)
	if err != nil {
		println(err.Error())
		return
	}

	isProjectFolderCreated := os.MkdirAll(filepath.Join(projectName, "src"), 0755)
	if isProjectFolderCreated != nil {
		println("Error creating directory: ", isProjectFolderCreated.Error())
		return
	}

	areConfigFilesCreated := filesystem.CreateConfigFiles(projectName)
	if areConfigFilesCreated != nil {
		println("Error creating config files: ", areConfigFilesCreated.Error())
		return
	}

	_, isIndexTSFileCreated := utils.CreateTextFile(filepath.Join(projectName, "src", "index.ts"))
	if isIndexTSFileCreated != nil {
		println("Error creating index.ts: ", isIndexTSFileCreated.Error())
		return
	}

	println("Project created successfully.")
}
