package sourcecontrol

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

type Committer struct{
	Name string `json:"name"`
	Email string `json:"email"`
}

type UploadFileRequest struct{
	Message string `json:"message"`
	Content string `json:"content"`
	Committer Committer `json:"committer"`
}