package components

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/axrona/pulsarship/internal/models"
	"github.com/axrona/pulsarship/internal/utils"
)

type GitBranchComponent struct {
	Config  models.GitBranchConfig
	Palette models.PaletteConfig
}

func init() {
	Registry["git_branch"] = func(config models.PromptConfig) models.Component {
		return &GitBranchComponent{
			Config:  config.GitBranch,
			Palette: config.Palette,
		}
	}
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

func (g *GitBranchComponent) Val() (string, error) {
	branchCmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	branchOut, err := branchCmd.Output()
	if def := utils.Must(err, ""); def != nil {
		return *def, err
	}
	branch := strings.TrimSpace(string(branchOut))

	return branch, nil
}

func (g *GitBranchComponent) Render() (models.Result, error) {
	utils.SetDefault(&g.Config.Format, "^(#e6e7ae)on^ ^(#b8a9f9) {branch}^")
	val, err := g.Val()
	if def := utils.Must(err, SkipComponent); def != nil {
		return *def, nil
	}

	rendered, err := utils.RenderFormat(*g.Config.Format, map[string]string{
		"git_branch": val,
		"branch":     val,
	}, (*map[string]string)(&g.Palette))

	return models.Result{Value: rendered}, err
}

func (g *GitBranchComponent) RenderAsync() <-chan models.Result {
	ch := make(chan models.Result, 1)
	go func() {
		val, err := g.Render()
		ch <- models.Result{Value: val.Value, Error: err}
	}()
	return ch
}

func (g GitBranchComponent) Name() string {
	return "git_branch"
}
