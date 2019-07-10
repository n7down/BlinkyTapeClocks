package githubdisplay

import (
	"bytes"
	"fmt"
	"github.com/n7down/pitftdisplays/internal/githubapi"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type GithubReleasesDisplay struct {
	config *viper.Viper
}

func NewGithubReleasesDisplay(config *viper.Viper) (*GithubReleasesDisplay, error) {
	return &GithubReleasesDisplay{
		config: config,
	}, nil
}

func (g GithubReleasesDisplay) Refresh() bool {
	return false
}

func (g GithubReleasesDisplay) Render() string {
	var buffer bytes.Buffer
	githubToken := g.config.GetString("github")
	betaflightReleases, err := githubapi.GetReleases("betaflight", "betaflight", githubToken)
	if err != nil {
		log.Error(err)
	}

	betaflightConfiguratorReleases, err := githubapi.GetReleases("betaflight", "betaflight-configurator", githubToken)
	if err != nil {
		log.Error(err)
	}

	//goReleases, err := githubapi.GetReleases("golang", "go", githubToken)
	//if err != nil {
	//log.Error(err)
	//}

	godotReleases, err := githubapi.GetReleases("godotengine", "godot", githubToken)
	if err != nil {
		log.Error(err)
	}

	//i3Releases, err := githubapi.GetReleases("Airblader", "i3", githubToken)
	//if err != nil {
	//log.Error(err)
	//}

	//linuxReleases, err := githubapi.GetReleases("torvalds", "linux", githubToken)
	//if err != nil {
	//log.Error(err)
	//}

	httpieReleases, err := githubapi.GetReleases("jakubroztocil", "httpie", githubToken)
	if err != nil {
		log.Error(err)
	}

	neovimReleases, err := githubapi.GetReleases("neovim", "neovim", githubToken)
	if err != nil {
		log.Error(err)
	}

	//fzfReleases, err := githubapi.GetReleases("junegunn", "fzf", githubToken)
	//if err != nil {
	//log.Error(err)
	//}

	buffer.WriteString(fmt.Sprintf("betaflight: %s\n", betaflightReleases[0].Name))
	buffer.WriteString(fmt.Sprintf("betflight configurator: %s\n", betaflightConfiguratorReleases[0].Name))
	buffer.WriteString(fmt.Sprintf("godot: %s\n", godotReleases[0].Name))
	buffer.WriteString(fmt.Sprintf("httpie: %s\n", httpieReleases[0].Name))
	buffer.WriteString(fmt.Sprintf("neovim: %s\n", neovimReleases[0].Name))
	return buffer.String()
}
