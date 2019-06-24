package spacexapi

type ApiInfo struct {
	ProjectName      string `json:"project_name"`
	Version          string `json:"version"`
	ProjectLink      string `json:"project_link"`
	Organization     string `json:"organization"`
	OrganizationLink string `json:"organization_link"`
	Description      string `json:"description"`
}
