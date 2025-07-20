package models

type Component interface {
	Render() (string, error)
	RenderAsync() <-chan Result
	Name() string
	Val() (string, error)
}

type Result struct {
	Value string
	Error error
}

type PromptConfig struct {
	Prompt string `toml:"prompt"`

	Hostname  HostnameConfig  `toml:"hostname"`
	Username  UsernameConfig  `toml:"username"`
	Cwd       CwdConfig       `toml:"cwd"`
	Time      TimeConfig      `toml:"time"`
	Character CharacterConfig `toml:"character"`
	Palette   PaletteConfig   `toml:"palette"`
}

type HostnameConfig struct {
	Format *string `toml:"format"`
}

type UsernameConfig struct {
	Format *string `toml:"format"`
}

type CwdConfig struct {
	Format    *string `toml:"format"`
	MaxLength *int    `toml:"max_length"`
}

type TimeConfig struct {
	Format     *string `toml:"format"`
	TimeFormat *string `toml:"time_format"`
}

type CharacterConfig struct {
	Icon   *string `toml:"icon"`
	Format *string `toml:"format"`
}

type PaletteConfig map[string]string
