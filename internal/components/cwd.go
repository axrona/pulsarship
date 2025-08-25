package components

import (
	"os"
	"strings"

	"github.com/axrona/pulsarship/internal/models"
	"github.com/axrona/pulsarship/internal/utils"
)

type CwdComponent struct {
	Config  models.CwdConfig
	Palette models.PaletteConfig
}

func init() {
	Registry["cwd"] = func(config models.PromptConfig) models.Component {
		return &CwdComponent{
			Config:  config.Cwd,
			Palette: config.Palette,
		}
	}
}

func (c *CwdComponent) Val() (string, error) {
	cwd, err := os.Getwd()
	if def := utils.Must(err, ""); def != nil {
		return *def, err
	}

	cwd = strings.ReplaceAll(cwd, os.Getenv("HOME"), "~")

	if *c.Config.MaxLength > 0 {
		cwdParts := strings.Split(cwd, string(os.PathSeparator))
		if len(cwdParts) > *c.Config.MaxLength {
			cwd = strings.Join(cwdParts[len(cwdParts)-*c.Config.MaxLength:], string(os.PathSeparator))
		}
	}

	return cwd, nil
}

func (c *CwdComponent) Render() (models.Result, error) {
	utils.SetDefault(&c.Config.MaxLength, 3)
	utils.SetDefault(&c.Config.Format, "^(#8e9ae6){cwd}^")
	var format string = *c.Config.Format

	val, err := c.Val()
	if def := utils.Must(err, SkipComponent); def != nil {
		return *def, err
	}

	rendered, err := utils.RenderFormat(format, map[string]string{
		"cwd": val,
	}, (*map[string]string)(&c.Palette))

	if def := utils.Must(err, SkipComponent); def != nil {
		return *def, err
	}

	return models.Result{Value: rendered}, nil
}

func (c *CwdComponent) RenderAsync() <-chan models.Result {
	ch := make(chan models.Result, 1)
	go func() {
		val, err := c.Render()
		ch <- models.Result{Value: val.Value, Error: err}
	}()
	return ch
}

func (c CwdComponent) Name() string {
	return "cwd"
}
