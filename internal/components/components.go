package components

import "github.com/xeyossr/pulsarship/internal/models"

// Create a map of components
func BuildComponentMap(config models.PromptConfig) map[string]models.Component {
	return map[string]models.Component{
		"username":  &UsernameComponent{Config: config.Username, Palette: config.Palette},
		"hostname":  &HostnameComponent{Config: config.Hostname, Palette: config.Palette},
		"cwd":       &CwdComponent{Config: config.Cwd, Palette: config.Palette},
		"time":      &TimeComponent{Config: config.Time, Palette: config.Palette},
		"character": &CharacterComponent{Config: config.Character, Palette: config.Palette},
	}
}
