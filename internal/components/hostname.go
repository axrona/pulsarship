package components

import (
	"os"

	"github.com/xeyossr/pulsarship/internal/models"
	"github.com/xeyossr/pulsarship/internal/utils"
)

type HostnameComponent struct {
	Config  models.HostnameConfig
	Palette models.PaletteConfig
}

func (h *HostnameComponent) Val() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}
	if hostname == "" {
		return "", nil
	}

	return hostname, nil
}

func (h *HostnameComponent) Render() (string, error) {
	utils.SetDefault(&h.Config.Format, "{hostname}")
	var format string = *h.Config.Format

	rendered, err := utils.RenderFormat(format, map[string]models.Component{
		"hostname": h,
	}, (*map[string]string)(&h.Palette))

	if err != nil {
		return "", err
	}

	return rendered, nil
}

func (c *HostnameComponent) RenderAsync() <-chan models.Result {
	ch := make(chan models.Result, 1)
	go func() {
		val, err := c.Render()
		ch <- models.Result{Value: val, Error: err}
	}()
	return ch
}

func (h HostnameComponent) Name() string {
	return "hostname"
}
