package main

import (
	"os"

	"github.com/HaroldObasi/copy-cat/deployments"
)

func main() {
	if len(os.Args) < 4 {
		panic("APP_NAME, Github token, and Vercel token are required")
	}
	APP_NAME := os.Args[1]
	// GITHUB_TOKEN := os.Args[2]
	VERCEL_TOKEN := os.Args[3]

	// bootstrap.BootStrapApp(APP_NAME)
	// template.FormatTemplate(APP_NAME)

	// apiUrl, htmlUrl := sourcecontrol.CreateRepo(APP_NAME, GITHUB_TOKEN)
	//
	// url := "https://api.github.com/repos/HaroldObasi/test2"
	// sourcecontrol.UploadDir(APP_NAME, url, GITHUB_TOKEN)

	htmlUrl := "https://github.com/HaroldObasi/test2"
	deployments.DeployToVercel(htmlUrl, VERCEL_TOKEN, APP_NAME)
}
