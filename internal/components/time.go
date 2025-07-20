package components

import (
	"time"

	"github.com/xeyossr/pulsarship/internal/models"
	"github.com/xeyossr/pulsarship/internal/utils"
)

type TimeComponent struct {
	Config  models.TimeConfig
	Palette models.PaletteConfig
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
	utils.SetDefault(&t.Config.Format, "{time}")
	var format string = *t.Config.Format

	val, err := t.Val()
	if err != nil {
		return models.Result{Skip: true}, err
	}

	rendered, err := utils.RenderFormat(format, map[string]string{
		"time": val,
	}, (*map[string]string)(&t.Palette))

	if err != nil {
		return models.Result{Skip: true}, err
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
