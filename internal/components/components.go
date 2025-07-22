package components

import "github.com/xeyossr/pulsarship/internal/models"

type ComponentFactory func(config any, palette models.PaletteConfig) models.Component

var Registry = make(map[string]func(config models.PromptConfig) models.Component)

// Create a map of components
func BuildComponentMap(config models.PromptConfig) map[string]models.Component {
	components := make(map[string]models.Component)
	for name, factory := range Registry {
		components[name] = factory(config)
	}

	for name, customCfg := range config.Custom {
		components[name] = &CustomComponent{
			ComponentName: name,
			Config:        customCfg,
			Palette:       config.Palette,
		}
	}

	return components
}
