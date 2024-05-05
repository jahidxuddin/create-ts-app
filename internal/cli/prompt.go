package cli

import (
	"bufio"
	"os"
	"strings"
)

func PromptProjectName() (string, error) {
	print("Enter project name: ")

	reader := bufio.NewReader(os.Stdin)
	projectName, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	projectName = strings.TrimSpace(projectName)

	return projectName, nil
}
