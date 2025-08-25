package components

import "github.com/axrona/pulsarship/internal/models"

type ComponentFactory func(config any, palette models.PaletteConfig) models.Component

var Registry = make(map[string]func(config models.PromptConfig) models.Component)
var SkipComponent = models.Result{Skip: true}

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
