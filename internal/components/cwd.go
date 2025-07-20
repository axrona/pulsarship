package components

import (
	"os"
	"strings"

	"github.com/xeyossr/pulsarship/internal/models"
	"github.com/xeyossr/pulsarship/internal/utils"
)

type CwdComponent struct {
	Config  models.CwdConfig
	Palette models.PaletteConfig
}

func (c *CwdComponent) Val() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
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

func (c *CwdComponent) Render() (string, error) {
	utils.SetDefault(&c.Config.MaxLength, 3)
	utils.SetDefault(&c.Config.Format, "{cwd}")
	var format string = *c.Config.Format

	val, err := c.Val()
	if err != nil {
		return "", err
	}

	rendered, err := utils.RenderFormat(format, map[string]string{
		"cwd": val,
	}, (*map[string]string)(&c.Palette))

	if err != nil {
		return "", err
	}

	return rendered, nil
}

func (c *CwdComponent) RenderAsync() <-chan models.Result {
	ch := make(chan models.Result, 1)
	go func() {
		val, err := c.Render()
		ch <- models.Result{Value: val, Error: err}
	}()
	return ch
}

func (c CwdComponent) Name() string {
	return "cwd"
}
