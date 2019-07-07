package githubdisplay

import (
	"fmt"
	"github.com/n7down/pitftdisplays/internal/githubapi"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type GithubReleasesDisplay struct {
	config *viper.Viper
}

func NewGithubReleasesDisplay(config *viper.Viper) *GithubReleasesDisplay {
	return &GithubReleasesDisplay{
		config: config,
	}
}

func (g GithubReleasesDisplay) Refresh() bool {
	return false
}

func (g GithubReleasesDisplay) Render() {
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

	fmt.Printf("betaflight: %s\n", betaflightReleases[0].Name)
	fmt.Printf("betaflight configurator: %s\n", betaflightConfiguratorReleases[0].Name)
	//fmt.Printf("%s\n", goReleases[0].Name)
	fmt.Printf("godot: %s\n", godotReleases[0].Name)
	//fmt.Printf("i3: %s\n", i3Releases[0].Name)
	//fmt.Printf("%s\n", linuxReleases)
	fmt.Printf("httpie: %s\n", httpieReleases[0].Name)
	fmt.Printf("neovim: %s\n", neovimReleases[0].Name)
	//fmt.Printf("%s\n", fzfReleases[0].Name)
}
