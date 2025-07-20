package components

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/xeyossr/pulsarship/internal/models"
	"github.com/xeyossr/pulsarship/internal/utils"
)

type GitComponent struct {
	Config  models.GitConfig
	Palette models.PaletteConfig
}

func findGitRoot(start string) (string, error) {
	dir := start
	for {
		if _, err := os.Stat(filepath.Join(dir, ".git")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf(".git directory not found")
		}
		dir = parent
	}
}

func (g *GitComponent) Val() (string, error) {
	utils.SetDefault(&g.Config.CleanSuffix, "")
	utils.SetDefault(&g.Config.DirtySuffix, " ^(#F38BA8)[*]^")
	cwd, _ := os.Getwd()
	_, err := findGitRoot(cwd)
	if err != nil {
		return "", nil
	}

	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	out, err := cmd.Output()
	if err != nil {
		return "", nil
	}
	branch := strings.TrimSpace(string(out))

	statusCmd := exec.Command("git", "status", "--porcelain")
	statusOut, err := statusCmd.Output()
	if err != nil {
		return branch, nil
	}
	clean := strings.TrimSpace(string(statusOut)) == ""
	var suffix string
	if clean {
		suffix, err = utils.RenderFormat(*g.Config.CleanSuffix, map[string]models.Component{}, (*map[string]string)(&g.Palette))
	} else {
		suffix, err = utils.RenderFormat(*g.Config.DirtySuffix, map[string]models.Component{}, (*map[string]string)(&g.Palette))
	}

	if err != nil {
		return "", err
	}

	return branch + suffix, nil
}

func (g *GitComponent) Render() (string, error) {
	utils.SetDefault(&g.Config.Format, "^(#f2a971) {git}^")
	val, err := g.Val()
	if err != nil || val == "" {
		return "", err
	}

	rendered, err := utils.RenderFormat(*g.Config.Format, map[string]models.Component{
		"git": g,
	}, (*map[string]string)(&g.Palette))

	return rendered, err
}

func (g *GitComponent) RenderAsync() <-chan models.Result {
	ch := make(chan models.Result, 1)
	go func() {
		val, err := g.Render()
		ch <- models.Result{Value: val, Error: err}
	}()
	return ch
}

func (g GitComponent) Name() string {
	return "git"
}
