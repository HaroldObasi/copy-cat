package deployments

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var client = &http.Client{}

func CreateProject(projectName, repoUrl, token string){
	reqData := CreateProjectRequest{
		Name: projectName,
		BuildCommand: "npm run build",
		DevCommand: "npm run dev",
		EnableAffectedProjectsDeployments: false,
		Framework: Vite,
		GitRepository: GitRepository{
			Repo: repoUrl,
			Type: Github,
		},
		InstallCommand: "npm install",
		SkipGitConnectDuringLink: false,
	}

	jsonData, err := json.Marshal(reqData)

	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest("POST", "https://api.vercel.com/v12/projects", bytes.NewBuffer(jsonData))

	if err != nil {
		panic(err)
	}

	request.Header.Set("Authorization", "Bearer " + token)
	request.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	fmt.Println(resp)

}

func CreateDeployment(projectName, repoUrl string){

}

func DeployToVercel(repoUrl, token, appName string) string{
	CreateProject(appName, repoUrl, token)

	return ""
}