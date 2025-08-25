package components

import (
	"time"

	"github.com/axrona/pulsarship/internal/models"
	"github.com/axrona/pulsarship/internal/utils"
)

type TimeComponent struct {
	Config  models.TimeConfig
	Palette models.PaletteConfig
}

func init() {
	Registry["time"] = func(config models.PromptConfig) models.Component {
		return &TimeComponent{
			Config:  config.Time,
			Palette: config.Palette,
		}
	}
}

func (t *TimeComponent) Val() (string, error) {
	currentTime := time.Now().Format(*t.Config.TimeFormat)
	if currentTime == "" {
		return "", nil
	}

	return currentTime, nil
}

func (t *TimeComponent) Render() (models.Result, error) {
	utils.SetDefault(&t.Config.TimeFormat, "15:04:05")
	utils.SetDefault(&t.Config.Format, "^(#97a1c3){time}^")
	var format string = *t.Config.Format

	val, err := t.Val()
	if def := utils.Must(err, SkipComponent); def != nil {
		return *def, err
	}

	rendered, err := utils.RenderFormat(format, map[string]string{
		"time": val,
	}, (*map[string]string)(&t.Palette))

	if def := utils.Must(err, SkipComponent); def != nil {
		return *def, err
	}

	return models.Result{Value: rendered}, nil
}

func (t *TimeComponent) RenderAsync() <-chan models.Result {
	ch := make(chan models.Result, 1)
	go func() {
		val, err := t.Render()
		ch <- models.Result{Value: val.Value, Error: err}
	}()
	return ch
}

func (t TimeComponent) Name() string {
	return "time"
}
