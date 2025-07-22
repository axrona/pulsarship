package components

import (
	"github.com/xeyossr/pulsarship/internal/models"
	"github.com/xeyossr/pulsarship/internal/utils"
)

type CharacterComponent struct {
	Config  models.CharacterConfig
	Palette models.PaletteConfig
}

func init() {
	Registry["character"] = func(config models.PromptConfig) models.Component {
		return &CharacterComponent{
			Config:  config.Character,
			Palette: config.Palette,
		}
	}
}

func (c *CharacterComponent) Val() (string, error) {
	return *c.Config.Icon, nil
}

func (c *CharacterComponent) Render() (models.Result, error) {
	utils.SetDefault(&c.Config.Icon, ">")
	utils.SetDefault(&c.Config.Format, ">")
	var format string = *c.Config.Format

	val, err := c.Val()
	if err != nil {
		return models.Result{Skip: true}, err
	}

	rendered, err := utils.RenderFormat(format, map[string]string{
		"character": val,
	}, (*map[string]string)(&c.Palette))

	if err != nil {
		return models.Result{Skip: true}, err
	}
	return models.Result{Value: rendered}, nil
}

func (c *CharacterComponent) RenderAsync() <-chan models.Result {
	ch := make(chan models.Result, 1)
	go func() {
		val, err := c.Render()
		ch <- models.Result{Value: val.Value, Error: err}
	}()
	return ch
}

func (c CharacterComponent) Name() string {
	return "character"
}
