package utils

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	env "github.com/xeyossr/pulsarship/internal"
	flagvars "github.com/xeyossr/pulsarship/internal/cli/flag_variables"
)

// check if the PULSARSHIP_DEBUG environment variable is set to "1"
func IsDebug() bool {
	return flagvars.DebugMode || env.PULSARSHIP_DEBUG == "1"
}

// if not in debug mode, execute the function
func IfNotDebug(fnNotDebug interface{}, fnDebug interface{}, params ...interface{}) {
	var fn interface{}
	if IsDebug() {
		fn = fnDebug
	} else {
		fn = fnNotDebug
	}

	fnValue := reflect.ValueOf(fn)
	if fnValue.Kind() != reflect.Func {
		if err, ok := fn.(error); ok {
			panic(err)
		}
		return
	}

	in := make([]reflect.Value, len(params))
	for i, param := range params {
		in[i] = reflect.ValueOf(param)
	}

	fnValue.Call(in)
}

// Create a pointer to a value of any type
func Ptr[T any](v T) *T {
	return &v
}

// set the pointer to a default value if it is nil
func SetDefault[T any](p **T, def T) {
	if *p == nil {
		*p = &def
	}
}

// Must returns a pointer to the default value if err is non-nil, otherwise nil.
func Must[T any](err error, def T) *T {
	if err != nil {
		return &def
	}

	return nil
}

// Resolves a color from a palette or returns the color itself if it's a valid hex code
func ResolveColor(color *string, palette map[string]string) (string, error) {
	if color == nil || *color == "" {
		return "", fmt.Errorf("color is nil or empty")
	}

	key := strings.TrimSpace(*color)
	if palette == nil {
		return key, nil
	}

	if val, ok := palette[key]; ok {
		return val, nil
	}

	if strings.HasPrefix(key, "#") {
		return key, nil
	}

	return "", fmt.Errorf("invalid color: %s", key)
}

// HexToRGB parses a hex color like "#FFAA33" and returns r,g,b (0-255)
func HexToRGB(hex string) (int, int, int, error) {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) != 6 {
		return 0, 0, 0, fmt.Errorf("invalid hex color: %s", hex)
	}
	var r, g, b int
	_, err := fmt.Sscanf(hex, "%02x%02x%02x", &r, &g, &b)
	return r, g, b, err
}

// Format the text with the specified color using ANSI escape codes
func Print(text, hex string, palette *map[string]string) string {
	hex, err := ResolveColor(&hex, *palette)
	if def := Must(err, text); def != nil {
		return *def
	}

	r, g, b, err := HexToRGB(hex)
	if def := Must(err, text); def != nil {
		return *def
	}

	colorSeq := fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
	resetSeq := "\x1b[0m"

	if strings.ToLower(env.PULSARSHIP_SHELL) == "zsh" {
		colorSeq = "%{" + colorSeq + "%}"
		resetSeq = "%{" + resetSeq + "%}"
	}

	return fmt.Sprintf("%s%s%s", colorSeq, text, resetSeq)
	//return fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s\x1b[0m", r, g, b, text)
}

var (
	bracedVarRegex  = regexp.MustCompile(`\{([a-zA-Z0-9_]+)\}`)
	coloredBlockReg = regexp.MustCompile(`\^\((#[a-fA-F0-9]{6}|[a-zA-Z0-9_]+)\)(.*?)\^`)
)

func RenderFormat(format string, vars map[string]string, palette *map[string]string) (string, error) {
	var err error

	result := coloredBlockReg.ReplaceAllStringFunc(format, func(match string) string {
		if def := Must(err, match); def != nil {
			return *def
		}

		matches := coloredBlockReg.FindStringSubmatch(match)
		if len(matches) != 3 {
			return match
		}
		colorName := matches[1]
		content := matches[2]

		resolved := bracedVarRegex.ReplaceAllStringFunc(content, func(m string) string {
			if def := Must(err, m); def != nil {
				return *def
			}
			key := bracedVarRegex.FindStringSubmatch(m)[1]
			if val, ok := vars[key]; ok {
				return val
			}
			return m
		})

		return Print(resolved, colorName, palette)
	})

	if def := Must(err, ""); def != nil {
		return *def, err
	}

	result = bracedVarRegex.ReplaceAllStringFunc(result, func(m string) string {
		if def := Must(err, m); def != nil {
			return *def
		}
		key := bracedVarRegex.FindStringSubmatch(m)[1]
		if val, ok := vars[key]; ok {
			return val
		}
		return m
	})

	if def := Must(err, ""); def != nil {
		return *def, err
	}

	return result, nil
}
