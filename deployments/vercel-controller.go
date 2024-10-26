package deployments

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var client = &http.Client{}

func CreateProject(projectName, repo, token string) CreateProjectResponse {
	reqData := CreateProjectRequest{
		Name:                              projectName,
		BuildCommand:                      "npm run build",
		EnableAffectedProjectsDeployments: false,
		Framework:                         Vite,
		GitRepository: GitRepository{
			Repo: repo,
			Type: Github,
		},
		InstallCommand:           "npm install",
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

	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Set("Content-Type", "application/json")

	clientResp, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	defer clientResp.Body.Close()

	body, err := io.ReadAll(clientResp.Body)

	if err != nil {
		panic(err)
	}

	responseData := CreateProjectResponse{}

	err = json.Unmarshal(body, &responseData)

	if err != nil {
		panic(err)
	}

	return responseData
}

func CreateDeployment(projectName, token string, repoId int) {
	reqData := DeployProjectRequest{
		Name: projectName,
		GitSource: GitSource{
			Ref:    "main",
			RepoId: repoId,
			Type:   Github,
		},
		ProjectSettings: projectSettings{
			Framework: Vite,
		},
	}

	jsonData, err := json.Marshal(reqData)

	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest("POST", "https://api.vercel.com/v13/deployments", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}

	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Set("Content-Type", "application/json")

	clientResp, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	defer clientResp.Body.Close()

	response := DeployProjectResponse{}

	bytes, err := io.ReadAll(clientResp.Body)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bytes, &response)

	if err != nil {
		panic(err)
	}

	fmt.Print("deployment url: ", response.Alias[0])
}

func DeployToVercel(repoUrl, token, appName string) {

	parts := strings.Split(repoUrl, "/")
	repo := parts[len(parts)-2] + "/" + parts[len(parts)-1]

	response := CreateProject(appName, repo, token)

	linkRepo := response.Link.RepoId

	CreateDeployment(appName, token, linkRepo)
}
