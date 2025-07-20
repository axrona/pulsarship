package config

import (
	"github.com/xeyossr/pulsarship/internal/models"
	"github.com/xeyossr/pulsarship/internal/utils"
)

var DefaultConfig = models.PromptConfig{
	Prompt: "{cwd}  {git}\n{character} ",
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
	Git: models.GitConfig{
		Format:      utils.Ptr("^(mauve) {branch} {status}^ "),
		CleanSuffix: utils.Ptr("^(overlay1)✔^ "),
		Conflicted:  utils.Ptr("^(red)Conflicted {count}^ "),
		Ahead:       utils.Ptr("^(green)↑{count}^ "),
		Behind:      utils.Ptr("^(peach)↓{count}^ "),
		Diverged:    utils.Ptr("^(yellow)↕{count}^ "),
		UpToDate:    utils.Ptr("^(green)✓^ "),
		Untracked:   utils.Ptr("^(overlay1)?{count}^ "),
		Stashed:     utils.Ptr("^(sky)★{count}^ "),
		Modified:    utils.Ptr("^(flamingo)!{count}^ "),
		Staged:      utils.Ptr("^(lavender)+{count}^ "),
		Renamed:     utils.Ptr("^(sapphire)»{count}^ "),
		Deleted:     utils.Ptr("^(red)-{count} ^"),
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
