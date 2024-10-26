package sourcecontrol

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"net/http"

	"github.com/HaroldObasi/copy-cat/utils"
)

var client = &http.Client{}

func GetUserInfo(token string) {
	url := "https://api.github.com/user"

	request, err := http.NewRequest("GET", url, nil)

	request.Header.Set("Authorization", "Bearer "+token)

	if err != nil {
		panic(err)
	}

	resp, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	githubUser := GithubUser{}

	err = json.Unmarshal(body, &githubUser)

	if err != nil {
		panic(err)
	}
}

func CreateRepo(repoName, token string) (string, string) {

	data := CreateRepoRequest{
		Name:        repoName,
		Description: "This is a test repo",
		Homepage:    "https://github.com",
		Private:     true,
		IsTemplate:  false,
	}

	jsonData, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	url := "https://api.github.com/user/repos"

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))

	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Set("Content-Type", "application/json")

	if err != nil {
		panic(err)
	}

	resp, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

	response := CreateRepoResponse{}

	err = json.Unmarshal(body, &response)

	if err != nil {
		panic(err)
	}

	fmt.Println("Html url: ", response.HtmlUrl)
	fmt.Println("Url: ", response.Url)

	return response.Url, response.HtmlUrl
}

func UploadFile(fileName, apiUrl, base64Encoded, token string) error {
	url := apiUrl + "/contents/" + fileName

	data := UploadFileRequest{
		Message: "Initial commit",
		Content: base64Encoded,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func UploadDir(dir, apiUrl, token string) {
	content, err := utils.GetFilesInDirectory("./", dir)
	if err != nil {
		panic(err)
	}

	// iterate over the files and upload them to the repo
	for _, file := range content {
		fileContent, err := os.ReadFile(dir + "/" + file)

		if err != nil {
			panic(err)
		}

		encoded := base64.StdEncoding.EncodeToString(fileContent)

		err = UploadFile(file, apiUrl, encoded, token)

		if err != nil {
			fmt.Println("Error uploading file: ", file)
		}
	}
}
