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
	Repo string `json:"repo"`
	Type GitRepositoryType `json:"type"`
}

type CreateProjectRequest struct {
	Name string `json:"name"`
	BuildCommand string `json:"buildCommand"`
	DevCommand string `json:"devCommand"`
	EnableAffectedProjectsDeployments bool `json:"enableAffectedProjectsDeployments"`
	Framework Framework `json:"framework"`
	GitRepository GitRepository `json:"gitRepository"`
	InstallCommand string `json:"installCommand"`
	SkipGitConnectDuringLink bool `json:"skipGitConnectDuringLink"`
}
