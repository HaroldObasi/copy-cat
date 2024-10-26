package main

import (
	"os"

	"github.com/HaroldObasi/copy-cat/bootstrap"
	"github.com/HaroldObasi/copy-cat/deployments"
	sourcecontrol "github.com/HaroldObasi/copy-cat/source-control"
	"github.com/HaroldObasi/copy-cat/template"
)

func main() {
	if len(os.Args) < 4 {
		panic("APP_NAME, Github token, and Vercel token are required")
	}
	APP_NAME := os.Args[1]
	GITHUB_TOKEN := os.Args[2]
	VERCEL_TOKEN := os.Args[3]

	bootstrap.BootStrapApp(APP_NAME)
	template.FormatTemplate(APP_NAME)

	apiUrl, htmlUrl := sourcecontrol.CreateRepo(APP_NAME, GITHUB_TOKEN)
	sourcecontrol.UploadDir(APP_NAME, apiUrl, GITHUB_TOKEN)
	deployments.DeployToVercel(htmlUrl, VERCEL_TOKEN, APP_NAME)
}
