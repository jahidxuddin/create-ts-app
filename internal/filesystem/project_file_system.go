package filesystem

import (
	"path/filepath"

	"github.com/jahidxuddin/create-ts-app/internal/utils"
)

type PackageJSON struct {
	Name            string            `json:"name"`
	Version         string            `json:"version"`
	Description     string            `json:"description"`
	Main            string            `json:"main"`
	Scripts         map[string]string `json:"scripts"`
	Keywords        []string          `json:"keywords"`
	Author          string            `json:"author"`
	License         string            `json:"license"`
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
}

type TsconfigJSON struct {
	CompilerOptions struct {
		Module           string `json:"module"`
		ModuleResolution string `json:"moduleResolution"`
		Target           string `json:"target"`
		SourceMap        bool   `json:"sourceMap"`
		OutDir           string `json:"outDir"`
	} `json:"compilerOptions"`
	Include []string `json:"include"`
}

func CreateConfigFiles(projectName string) error {
	packageJSON := PackageJSON{
		Name:        "my-app",
		Version:     "1.0.0",
		Description: "",
		Main:        "index.js",
		Scripts: map[string]string{
			"build": "npx tsc",
			"dev":   "concurrently \"npx tsc -w\" \"nodemon dist/index.js\"",
			"start": "node dist/index.js",
		},
		Keywords: []string{},
		Author:   "",
		License:  "MIT",
		Dependencies: map[string]string{
			"express": "^4.19.2",
		},
		DevDependencies: map[string]string{
			"@types/express": "^4.17.21",
			"@types/node":    "^20.12.8",
			"concurrently":   "^8.2.2",
			"nodemon":        "^3.1.0",
			"ts-node":        "^10.9.2",
			"typescript":     "^5.4.5",
		},
	}

	tsconfigJSON := TsconfigJSON{
		CompilerOptions: struct {
			Module           string `json:"module"`
			ModuleResolution string `json:"moduleResolution"`
			Target           string `json:"target"`
			SourceMap        bool   `json:"sourceMap"`
			OutDir           string `json:"outDir"`
		}{
			Module:           "NodeNext",
			ModuleResolution: "NodeNext",
			Target:           "ES2022",
			SourceMap:        true,
			OutDir:           "dist",
		},
		Include: []string{"src/**/*"},
	}

	isPackageJSONCreated := utils.CreateJSONFile(filepath.Join(projectName, "package.json"), packageJSON)
	if isPackageJSONCreated != nil {
		return utils.NewError("Error while creating package.json: " + isPackageJSONCreated.Error())
	}

	isTsConfigCreated := utils.CreateJSONFile(filepath.Join(projectName, "tsconfig.json"), tsconfigJSON)
	if isTsConfigCreated != nil {
		return utils.NewError("Error while creating tsconfig.json: " + isTsConfigCreated.Error())
	}

	return nil
}
