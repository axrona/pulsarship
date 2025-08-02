package components

import (
	"bytes"
	"os/exec"

	"github.com/xeyossr/pulsarship/internal/models"
	"github.com/xeyossr/pulsarship/internal/utils"
)

type CustomComponent struct {
	ComponentName string
	Config        models.CustomComponentConfig
	Palette       models.PaletteConfig
}

func (c *CustomComponent) Val() (string, error) {
	if c.Config.Run == nil {
		return "", nil
	}

	runCmd := exec.Command("bash", "-c", *c.Config.Run)

	var stdoutBuf bytes.Buffer
	runCmd.Stdout = &stdoutBuf

	err := runCmd.Run()
	if def := utils.Must(err, ""); def != nil {
		return *def, err
	}

	stdoutStr := stdoutBuf.String()

	return stdoutStr, nil
}

func (c *CustomComponent) Render() (models.Result, error) {
	utils.SetDefault(&c.Config.Format, "{output}")
	var format string = *c.Config.Format

	val, err := c.Val()
	if def := utils.Must(err, SkipComponent); def != nil {
		return *def, err
	}

	rendered, err := utils.RenderFormat(format, map[string]string{
		"output": val,
	}, (*map[string]string)(&c.Palette))

	if def := utils.Must(err, SkipComponent); def != nil {
		return *def, err
	}

	return models.Result{Value: rendered}, nil
}

func (c *CustomComponent) RenderAsync() <-chan models.Result {
	ch := make(chan models.Result, 1)
	go func() {
		val, err := c.Render()
		ch <- models.Result{Value: val.Value, Error: err}
	}()
	return ch
}

func (c CustomComponent) Name() string {
	return c.ComponentName
}
