package bootstrap

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

type PackageJSON struct {
	Name        string            `json:"name"`
	Version     string            `json:"version"`
	Description string            `json:"description"`
	Main        string            `json:"main"`
	Scripts     map[string]string `json:"scripts"`
}

func (pj PackageJSON) String() string {
	return fmt.Sprintf("Name: %v", pj.Name)
}

func BootStrapApp(appName string) {


	// create folder for app
	appScaffold := exec.Command(
		"mkdir",
		appName,
	)

	_, err := appScaffold.Output()
	if err != nil {
		fmt.Println(err)
	}

	// initialize npm and install vite in app folder
	fmt.Println("Installing Vite...")
	npmInit := exec.Command(
		"npm",
		"init",
		"-y",
	)

	npmInit.Dir = appName

	_, err = npmInit.Output()
	if err != nil {
		fmt.Println(err)
	}

	npmInstallVite := exec.Command(
		"npm",
		"install",
		"-D",
		"vite",
	)

	npmInstallVite.Dir = appName

	_, err = npmInstallVite.Output()
	if err != nil {
		fmt.Println(err)
	}

	// add dev script to package.json
	var data PackageJSON

	jsonFile, err := os.ReadFile("./app/package.json")

	if err != nil{
		panic(err)
	}

	err = json.Unmarshal(jsonFile, &data)

	if err != nil {
		panic(err)
	}

	fmt.Print(data)
}