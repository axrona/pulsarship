package config

import (
	"github.com/xeyossr/pulsarship/internal/models"
	"github.com/xeyossr/pulsarship/internal/utils"
)

var DefaultConfig = models.PromptConfig{
	Prompt:     "{cwd} {git_branch} {git_status}\n{character} ",
	AddNewLine: true,

	Hostname: models.HostnameConfig{
		Format: nil,
	},
	Username: models.UsernameConfig{
		Format: nil,
	},
	Cwd: models.CwdConfig{
		Format:    utils.Ptr("^(lavender){cwd}^"),
		MaxLength: utils.Ptr(1),
	},
	Time: models.TimeConfig{
		Format:     nil,
		TimeFormat: nil,
	},
	Character: models.CharacterConfig{
		Icon:   utils.Ptr("❯"),
		Format: utils.Ptr("^(peach){character}^"),
	},
	GitBranch: models.GitBranchConfig{
		Format: utils.Ptr("^(yellow)on^ ^(mauve) {branch}^"),
	},

	GitStatus: models.GitStatusConfig{
		Format:      utils.Ptr("^(red)[^{status}^(red)]^"),
		CleanSuffix: utils.Ptr(""),
		Conflicted:  utils.Ptr("^(red)✖^"),
		Ahead:       utils.Ptr("^(red)⇡^"),
		Behind:      utils.Ptr("^(red)⇣^"),
		Diverged:    utils.Ptr("^(red)⇕^"),
		UpToDate:    utils.Ptr("^(red)✓^"),
		Untracked:   utils.Ptr("^(red)?^"),
		Stashed:     utils.Ptr("^(red)S^"),
		Modified:    utils.Ptr("^(red)~^"),
		Staged:      utils.Ptr("^(red)+^"),
		Renamed:     utils.Ptr("^(red)»^"),
		Deleted:     utils.Ptr("^(red)✘^"),
	},

	Palette: models.PaletteConfig{
		"rosewater": "#f5e0dc",
		"flamingo":  "#f2cdcd",
		"pink":      "#f5c2e7",
		"mauve":     "#cba6f7",
		"red":       "#f38ba8",
		"maroon":    "#eba0ac",
		"peach":     "#fab387",
		"yellow":    "#f9e2af",
		"green":     "#a6e3a1",
		"teal":      "#94e2d5",
		"sky":       "#89dceb",
		"sapphire":  "#74c7ec",
		"blue":      "#89b4fa",
		"lavender":  "#b4befe",
		"text":      "#cdd6f4",
		"subtext1":  "#bac2de",
		"subtext0":  "#a6adc8",
		"overlay2":  "#9399b2",
		"overlay1":  "#7f849c",
		"overlay0":  "#6c7086",
		"surface2":  "#585b70",
		"surface1":  "#45475a",
		"surface0":  "#313244",
		"base":      "#1e1e2e",
		"mantle":    "#181825",
		"crust":     "#11111b",
	},
}
