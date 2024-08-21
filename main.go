package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
)

type ProjectFolder struct {
	Key  string `json:"key"`
	Path string `json:"path"`
}

type Config struct {
	ProjectFolders []ProjectFolder `json:"project_folders"`
}

const configFileName = "config.json"

func main() {
	var config Config

	fileContent, err := os.ReadFile(configFileName)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		panic(err)
	}

	directories := make([]string, len(config.ProjectFolders))
	for i, folder := range config.ProjectFolders {
		directories[i] = folder.Path
	}

	var chosenFile string
	filePaths := []string{}
	for _, directory := range directories {
		// Step 2: List files in the chosen directory
		f, err := ioutil.ReadDir(directory)
		if err != nil {
			panic(err)
		}

		for _, file := range f {
			if file.IsDir() {
				filePaths = append(filePaths, filepath.Join(directory, file.Name()))
			}
		}
	}

	// Step 3: Choose a file or folder
	prompt := &survey.Select{
		Message: "Choose a file or folder to open:",
		Options: filePaths,
	}
	survey.AskOne(prompt, &chosenFile)

	// Step 4: Open the chosen file or folder in VSCode
	err = exec.Command("code", chosenFile).Start()
	if err != nil {
		fmt.Printf("Error opening VSCode: %v\n", err)
	}

}
