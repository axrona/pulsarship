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

func (u *UsernameComponent) Render() (string, error) {
	utils.SetDefault(&u.Config.Format, "{username}")
	var format string = *u.Config.Format

	rendered, err := utils.RenderFormat(format, map[string]models.Component{
		"username": u,
	}, (*map[string]string)(&u.Palette))

	if err != nil {
		return "", err
	}

	return rendered, nil
}

func (u *UsernameComponent) RenderAsync() <-chan models.Result {
	ch := make(chan models.Result, 1)
	go func() {
		val, err := u.Render()
		ch <- models.Result{Value: val, Error: err}
	}()
	return ch
}

func (u UsernameComponent) Name() string {
	return "username"
}
