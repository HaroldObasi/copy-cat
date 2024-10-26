package template

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
)

type Project struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	TechStack   []string `json:"techStack"`
	Github      string   `json:"github"`
}

type Social struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type UserInfo struct {
	Name     string    `json:"name"`
	JobTitle string    `json:"jobTitle"`
	Bio      string    `json:"bio"`
	Socials  []Social  `json:"socials"`
	Projects []Project `json:"projects"`
}

func (u UserInfo) String() string {
	return fmt.Sprintf("Name: %s\nJobTitle: %s\nBio: %s\nSocials: %v\nProjects: %v\n", u.Name, u.JobTitle, u.Bio, u.Socials, u.Projects)
}

func FormatTemplate(appName string) {
	fmt.Println("Setting up template")
	htmlTemplate, err := os.ReadFile("./template/template.html")

	if err != nil {
		fmt.Println(err)
	}

	t, err := template.New("webpage").Parse(string(htmlTemplate))
	if err != nil {
		fmt.Println(err)
	}

	jsonFile, err := os.ReadFile("./template/userInfo.json")

	if err != nil {
		fmt.Println(err)
	}

	var data UserInfo

	err = json.Unmarshal(jsonFile, &data)

	if err != nil {
		fmt.Println(err)
	}

	// f, err := os.Create("output.html")
	f, err := os.Create(fmt.Sprintf("./%s/index.html", appName))

	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Execute the template
	err = t.Execute(f, data)
	if err != nil {
		panic(err)
	}

	fmt.Println("Template created")
}
