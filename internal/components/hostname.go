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

func init() {
	Registry["hostname"] = func(config models.PromptConfig) models.Component {
		return &HostnameComponent{
			Config:  config.Hostname,
			Palette: config.Palette,
		}
	}
}

func (h *HostnameComponent) Val() (string, error) {
	hostname, err := os.Hostname()
	if def := utils.Must(err, ""); def != nil {
		return *def, err
	}
	if hostname == "" {
		return "", nil
	}

	return hostname, nil
}

func (h *HostnameComponent) Render() (models.Result, error) {
	utils.SetDefault(&h.Config.Format, "{hostname}")
	var format string = *h.Config.Format

	val, err := h.Val()
	if def := utils.Must(err, SkipComponent); def != nil {
		return *def, err
	}

	rendered, err := utils.RenderFormat(format, map[string]string{
		"hostname": val,
	}, (*map[string]string)(&h.Palette))

	if def := utils.Must(err, SkipComponent); def != nil {
		return *def, err
	}

	return models.Result{Value: rendered}, nil
}

func (c *HostnameComponent) RenderAsync() <-chan models.Result {
	ch := make(chan models.Result, 1)
	go func() {
		val, err := c.Render()
		ch <- models.Result{Value: val.Value, Error: err}
	}()
	return ch
}

func (h HostnameComponent) Name() string {
	return "hostname"
}
