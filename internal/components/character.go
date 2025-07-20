package components

import (
	"github.com/xeyossr/pulsarship/internal/models"
	"github.com/xeyossr/pulsarship/internal/utils"
)

type CharacterComponent struct {
	Config  models.CharacterConfig
	Palette models.PaletteConfig
}

func (c *CharacterComponent) Val() (string, error) {
	return *c.Config.Icon, nil
}

func (c *CharacterComponent) Render() (string, error) {
	utils.SetDefault(&c.Config.Icon, ">")
	utils.SetDefault(&c.Config.Format, ">")
	var format string = *c.Config.Format

	rendered, err := utils.RenderFormat(format, map[string]models.Component{
		"character": c,
	}, (*map[string]string)(&c.Palette))

	if err != nil {
		return "", err
	}
	return rendered, nil
}

func (c *CharacterComponent) RenderAsync() <-chan models.Result {
	ch := make(chan models.Result, 1)
	go func() {
		val, err := c.Render()
		ch <- models.Result{Value: val, Error: err}
	}()
	return ch
}

func (c CharacterComponent) Name() string {
	return "character"
}
