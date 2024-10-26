package deployments

type Framework string
type GitRepositoryType string

const (
	Vite Framework = "vite"
)

const (
	Github GitRepositoryType = "github"
)

type GitRepository struct {
	Repo string            `json:"repo"`
	Type GitRepositoryType `json:"type"`
}

type Link struct {
	Type   GitRepositoryType `json:"type"`
	Repo   string            `json:"repo"`
	RepoId int               `json:"repoId"`
}

type CreateProjectRequest struct {
	Name                              string        `json:"name"`
	BuildCommand                      string        `json:"buildCommand"`
	EnableAffectedProjectsDeployments bool          `json:"enableAffectedProjectsDeployments"`
	Framework                         Framework     `json:"framework"`
	GitRepository                     GitRepository `json:"gitRepository"`
	InstallCommand                    string        `json:"installCommand"`
	SkipGitConnectDuringLink          bool          `json:"skipGitConnectDuringLink"`
}

type CreateProjectResponse struct {
	CreatedAt int  `json:"createdAt"`
	Link      Link `json:"link"`
}

type GitSource struct {
	Ref    string            `json:"ref"`
	RepoId int               `json:"repoId"`
	Type   GitRepositoryType `json:"type"`
}

type projectSettings struct {
	Framework Framework `json:"framework"`
}

type DeployProjectRequest struct {
	Name            string          `json:"name"`
	GitSource       GitSource       `json:"gitSource"`
	ProjectSettings projectSettings `json:"projectSettings"`
}

type DeployProjectResponse struct {
	Alias     []string `json:"alias"`
	CreatedAt int      `json:"createdAt"`
}
