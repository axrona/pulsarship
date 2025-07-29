package components

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/xeyossr/pulsarship/internal/models"
	"github.com/xeyossr/pulsarship/internal/utils"
)

type GitStatusComponent struct {
	Config  models.GitStatusConfig
	Palette models.PaletteConfig
}

func init() {
	Registry["git_status"] = func(config models.PromptConfig) models.Component {
		return &GitStatusComponent{
			Config:  config.GitStatus,
			Palette: config.Palette,
		}
	}
}

func (g *GitStatusComponent) Val() (string, error) {
	utils.SetDefault(&g.Config.Format, "^(#f28fad)[^{status}^(#f28fad)]^")
	utils.SetDefault(&g.Config.CleanSuffix, "")
	utils.SetDefault(&g.Config.Conflicted, "^(#f28fad)✖^")
	utils.SetDefault(&g.Config.Ahead, "^(#f28fad)⇡^")
	utils.SetDefault(&g.Config.Behind, "^(#f28fad)⇣^")
	utils.SetDefault(&g.Config.Diverged, "^(#f28fad)⇕^")
	utils.SetDefault(&g.Config.UpToDate, "^(#f28fad)✓^")
	utils.SetDefault(&g.Config.Untracked, "^(#f28fad)?^")
	utils.SetDefault(&g.Config.Stashed, "^(#f28fad)S^")
	utils.SetDefault(&g.Config.Modified, "^(#f28fad)~^")
	utils.SetDefault(&g.Config.Staged, "^(#f28fad)+^")
	utils.SetDefault(&g.Config.Renamed, "^(#f28fad)»^")
	utils.SetDefault(&g.Config.Deleted, "^(#f28fad)✘^")

	cwd, _ := os.Getwd()
	_, err := findGitRoot(cwd)
	if err != nil {
		return "", nil
	}

	counts := map[string]int{
		"conflicted": 0,
		"untracked":  0,
		"modified":   0,
		"staged":     0,
		"renamed":    0,
		"deleted":    0,
		"stashed":    0,
		"ahead":      0,
		"behind":     0,
		"diverged":   0,
	}

	statusCmd := exec.Command("git", "status", "--porcelain", "--branch")
	statusOut, err := statusCmd.Output()
	if err != nil {
		return "", err
	}
	lines := strings.Split(string(statusOut), "\n")

	var statusSymbols []string

	for _, line := range lines {
		if strings.HasPrefix(line, "##") {
			if strings.Contains(line, "ahead") && strings.Contains(line, "behind") {
				counts["diverged"]++
			} else if strings.Contains(line, "ahead") {
				counts["ahead"]++
			} else if strings.Contains(line, "behind") {
				counts["behind"]++
			}
			continue
		}

		code := strings.TrimSpace(line)
		if code == "" {
			continue
		}

		switch {
		case strings.HasPrefix(code, "UU"):
			counts["conflicted"]++
		case strings.HasPrefix(code, "??"):
			counts["untracked"]++
		case strings.HasPrefix(code, "M "), strings.HasPrefix(code, " M"):
			counts["modified"]++
		case strings.HasPrefix(code, "A "), strings.HasPrefix(code, "AM"):
			counts["staged"]++
		case strings.HasPrefix(code, "R "):
			counts["renamed"]++
		case strings.HasPrefix(code, "D "), strings.HasPrefix(code, " D"):
			counts["deleted"]++
		}
	}

	stashCmd := exec.Command("git", "stash", "list")
	stashOut, err := stashCmd.Output()
	if err == nil && strings.TrimSpace(string(stashOut)) != "" {
		counts["stashed"] = len(strings.Split(strings.TrimSpace(string(stashOut)), "\n"))
	}

	types := map[string]*string{
		"conflicted": g.Config.Conflicted,
		"untracked":  g.Config.Untracked,
		"modified":   g.Config.Modified,
		"staged":     g.Config.Staged,
		"renamed":    g.Config.Renamed,
		"deleted":    g.Config.Deleted,
		"stashed":    g.Config.Stashed,
		"ahead":      g.Config.Ahead,
		"behind":     g.Config.Behind,
		"diverged":   g.Config.Diverged,
	}

	for key, val := range types {
		if counts[key] > 0 {
			formatted, err := utils.RenderFormat(*val, map[string]string{
				"count": fmt.Sprintf("%d", counts[key]),
			}, (*map[string]string)(&g.Palette))
			if err != nil {
				return "", err
			}
			statusSymbols = append(statusSymbols, formatted)
		}
	}

	if len(statusSymbols) == 0 {
		suffix, err := utils.RenderFormat(*g.Config.CleanSuffix, map[string]string{}, (*map[string]string)(&g.Palette))
		if err != nil {
			return "", err
		}
		return suffix, nil
	}

	suffix := strings.Join(statusSymbols, "")
	return suffix, nil
}

func (g *GitStatusComponent) Render() (models.Result, error) {
	utils.SetDefault(&g.Config.Format, "{status}")
	val, err := g.Val()
	if err != nil {
		return models.Result{Skip: true}, err
	}

	rendered, err := utils.RenderFormat(*g.Config.Format, map[string]string{
		"git_status": val,
		"status":     val,
	}, (*map[string]string)(&g.Palette))

	return models.Result{Value: rendered}, err
}

func (g *GitStatusComponent) RenderAsync() <-chan models.Result {
	ch := make(chan models.Result, 1)
	go func() {
		val, err := g.Render()
		ch <- models.Result{Value: val.Value, Error: err}
	}()
	return ch
}

func (g GitStatusComponent) Name() string {
	return "git_status"
}
