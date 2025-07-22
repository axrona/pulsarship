package components

import (
	"os/user"

	"github.com/xeyossr/pulsarship/internal/models"
	"github.com/xeyossr/pulsarship/internal/utils"
)

type UsernameComponent struct {
	Config  models.UsernameConfig
	Palette models.PaletteConfig
}

func init() {
	Registry["username"] = func(config models.PromptConfig) models.Component {
		return &UsernameComponent{
			Config:  config.Username,
			Palette: config.Palette,
		}
	}
}

func (u *UsernameComponent) Val() (string, error) {
	username, err := user.Current()
	if err != nil {
		return "", err
	}
	if username.Name == "" {
		return "", nil
	}

	return username.Name, nil
}

func (u *UsernameComponent) Render() (models.Result, error) {
	utils.SetDefault(&u.Config.Format, "{username}")
	var format string = *u.Config.Format

	val, err := u.Val()
	if err != nil {
		return models.Result{Skip: true}, err
	}

	rendered, err := utils.RenderFormat(format, map[string]string{
		"username": val,
	}, (*map[string]string)(&u.Palette))

	if err != nil {
		return models.Result{Skip: true}, err
	}

	return models.Result{Value: rendered}, nil
}

func (u *UsernameComponent) RenderAsync() <-chan models.Result {
	ch := make(chan models.Result, 1)
	go func() {
		val, err := u.Render()
		ch <- models.Result{Value: val.Value, Error: err}
	}()
	return ch
}

func (u UsernameComponent) Name() string {
	return "username"
}
