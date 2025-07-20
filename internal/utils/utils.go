package utils

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/xeyossr/pulsarship/internal/models"
)

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
	if err != nil {
		return text
	}

	r, g, b, err := HexToRGB(hex)
	if err != nil {
		return text
	}
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s\x1b[0m", r, g, b, text)
}

var (
	bracedVarRegex  = regexp.MustCompile(`\{([a-zA-Z0-9_]+)\}`)
	coloredBlockReg = regexp.MustCompile(`\^\((#[a-fA-F0-9]{6}|[a-zA-Z_]+)\)(.*?)\^`)
)

func RenderFormat(format string, components map[string]models.Component, palette *map[string]string) (string, error) {
	var err error

	result := coloredBlockReg.ReplaceAllStringFunc(format, func(match string) string {
		if err != nil {
			return match
		}

		matches := coloredBlockReg.FindStringSubmatch(match)
		if len(matches) != 3 {
			return match
		}
		colorName := matches[1]
		content := matches[2]

		resolved := bracedVarRegex.ReplaceAllStringFunc(content, func(m string) string {
			if err != nil {
				return m
			}
			key := bracedVarRegex.FindStringSubmatch(m)[1]
			if comp, ok := components[key]; ok {
				var val string
				val, err = comp.Val()
				if err != nil {
					return m
				}
				return val
			}
			return m
		})

		return Print(resolved, colorName, palette)
	})

	if err != nil {
		return "", err
	}

	result = bracedVarRegex.ReplaceAllStringFunc(result, func(m string) string {
		if err != nil {
			return m
		}
		key := bracedVarRegex.FindStringSubmatch(m)[1]
		if comp, ok := components[key]; ok {
			var val string
			val, err = comp.Val()
			if err != nil {
				return m
			}
			return val
		}
		return m
	})

	if err != nil {
		return "", err
	}

	return result, nil
}
