package models

type Component interface {
	Render() (Result, error)
	RenderAsync() <-chan Result
	Name() string
	Val() (string, error)
}

type Result struct {
	Value string
	Error error
	Skip  bool
}

type PromptConfig struct {
	Prompt      string  `toml:"prompt"`
	RightPrompt *string `toml:"prompt_right"`
	AddNewLine  bool    `toml:"add_newline"`

	Custom    map[string]CustomComponentConfig `toml:"custom"`
	Hostname  HostnameConfig                   `toml:"hostname"`
	Username  UsernameConfig                   `toml:"username"`
	Cwd       CwdConfig                        `toml:"cwd"`
	Time      TimeConfig                       `toml:"time"`
	Character CharacterConfig                  `toml:"character"`
	Palette   PaletteConfig                    `toml:"palette"`
	GitBranch GitBranchConfig                  `toml:"git_branch"`
	GitStatus GitStatusConfig                  `toml:"git_status"`
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

type CustomComponentConfig struct {
	Format *string `toml:"format"`
	Run    *string `toml:"run"`
}

type GitBranchConfig struct {
	Format *string `toml:"format"`
}

type GitStatusConfig struct {
	Format      *string `toml:"format"`
	CleanSuffix *string `toml:"clean_suffix"`
	Conflicted  *string `toml:"conflicted"`
	Ahead       *string `toml:"ahead"`
	Behind      *string `toml:"behind"`
	Diverged    *string `toml:"diverged"`
	UpToDate    *string `toml:"up_to_date"`
	Untracked   *string `toml:"untracked"`
	Stashed     *string `toml:"stashed"`
	Modified    *string `toml:"modified"`
	Staged      *string `toml:"staged"`
	Renamed     *string `toml:"renamed"`
	Deleted     *string `toml:"deleted"`
}

type PaletteConfig map[string]string
