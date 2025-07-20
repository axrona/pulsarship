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

func Status(g *GitComponent) (string, error) {
	utils.SetDefault(&g.Config.CleanSuffix, "")
	utils.SetDefault(&g.Config.UpToDate, " ^(#ff0000)[✓]^")
	utils.SetDefault(&g.Config.Conflicted, " ^(#ff0000)[!?]^")
	utils.SetDefault(&g.Config.Ahead, " ^(#ff0000)[↑+]^")
	utils.SetDefault(&g.Config.Behind, " ^(#ff0000)[↓-]^")
	utils.SetDefault(&g.Config.Diverged, " ^(#ff0000)[⇅!?]^")
	utils.SetDefault(&g.Config.Untracked, " ^(#ff0000)[?{count}]^")
	utils.SetDefault(&g.Config.Stashed, " ^(#ff0000)[S{count}]^")
	utils.SetDefault(&g.Config.Modified, " ^(#ff0000)[M{count}]^")
	utils.SetDefault(&g.Config.Staged, " ^(#ff0000)[+{count}]^")
	utils.SetDefault(&g.Config.Renamed, " ^(#ff0000)[R{count}]^")
	utils.SetDefault(&g.Config.Deleted, " ^(#ff0000)[X{count}]^")

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

func (g *GitComponent) Val() (string, error) {
	branchCmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	branchOut, err := branchCmd.Output()
	if err != nil {
		return "", err
	}
	branch := strings.TrimSpace(string(branchOut))

	return branch, nil
}

func (g *GitComponent) Render() (models.Result, error) {
	utils.SetDefault(&g.Config.Format, "^(#f2a971) {git}^")
	val, err := g.Val()
	if err != nil {
		return models.Result{Skip: true}, err
	}

	status, err := Status(g)
	if err != nil {
		return models.Result{Skip: true}, err
	}

	rendered, err := utils.RenderFormat(*g.Config.Format, map[string]string{
		"git":    val,
		"branch": val,
		"status": status,
	}, (*map[string]string)(&g.Palette))

	return models.Result{Value: rendered}, err
}

func (g *GitComponent) RenderAsync() <-chan models.Result {
	ch := make(chan models.Result, 1)
	go func() {
		val, err := g.Render()
		ch <- models.Result{Value: val.Value, Error: err}
	}()
	return ch
}

func (g GitComponent) Name() string {
	return "git"
}
