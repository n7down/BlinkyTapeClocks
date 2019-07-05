package githubdisplay

import (
	"fmt"
	"github.com/n7down/pitftdisplays/internal/githubapi"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type GithubDisplay struct {
	config *viper.Viper
}

func NewGithubDisplay(config *viper.Viper) *GithubDisplay {
	return &GithubDisplay{
		config: config,
	}
}

func (g GithubDisplay) Refresh() bool {
	return false
}

func (g GithubDisplay) Render() {
	owner := "betaflight"
	repo := "betaflight"
	githubToken := g.config.GetString("github")
	log.Info(githubToken)
	releases, err := githubapi.GetReleases(owner, repo, githubToken)
	if err != nil {
		log.Error(err)
	}
	fmt.Printf("%s", releases[0].Name)
}
