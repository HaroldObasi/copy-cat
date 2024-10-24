package sourcecontrol

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"net/http"

	"github.com/HaroldObasi/copy-cat/utils"
)

type CreateRepoRequest struct{
	Name string `json:"name"`
	Description string `json:"description"`
	Homepage string `json:"homepage"`
	Private bool `json:"private"`
	IsTemplate bool `json:"is_template"`
}

type CreateRepoResponse struct{
	HtmlUrl string `json:"html_url"`
	Url string `json:"url"`
}

var client = &http.Client{}


func CreateRepo(repoName, token string) string {

	data := CreateRepoRequest{
		Name: repoName,
		Description: "This is a test repo",
		Homepage: "https://github.com",
		Private: true,
		IsTemplate: false,
	}

	jsonData, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	url := "https://api.github.com/user/repos"

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))

	request.Header.Set("Authorization", "Bearer " + token)
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

	return response.Url
}

func UploadDir(dir, apiUrl, token string) []string {
	content, err := utils.GetFilesInDirectory("./" ,dir)
	if err != nil {
		panic(err)
	}

	fmt.Println(content)
	return []string{}
}