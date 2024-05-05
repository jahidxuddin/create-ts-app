package cli

import (
	"regexp"

	"github.com/jahidxuddin/create-ts-app/internal/utils"
)

func validateProjectName(projectName string) error {
	if len(projectName) == 0 {
		return utils.NewError("Please enter a project name with at least one character.")
	}

	folderNamingPattern := `^[^/\\:*?"<>|\r\n]+$`

	regex := regexp.MustCompile(folderNamingPattern)

	if !regex.MatchString(projectName) {
		return utils.NewError("Please enter a valid project name.")
	}

	return nil
}

func ParseProjectNameFromArgs(args []string) (string, error) {
	if len(args) > 1 {
		return "", utils.NewError("Too many arguments.")
	}

	var projectName string
	if len(args) == 0 {
		name, err := PromptProjectName()
		if err != nil {
			return "", err
		}
		projectName = name
	} else {
		projectName = args[0]
	}

	if err := validateProjectName(projectName); err != nil {
		return "", err
	}

	return projectName, nil
}
