package components

import (
	"strings"

	"github.com/xeyossr/pulsarship/internal/models"
)

type PromptPart struct {
	Value       string
	IsComponent bool
}

func SplitPrompt(prompt string) []PromptPart {
	var parts []PromptPart
	var builder strings.Builder
	inside := false

	for i := 0; i < len(prompt); i++ {
		c := prompt[i]

		if c == '{' {
			if builder.Len() > 0 && !inside {
				parts = append(parts, PromptPart{
					Value:       builder.String(),
					IsComponent: false,
				})
				builder.Reset()
			}
			inside = true
			continue
		} else if c == '}' {
			if inside {
				parts = append(parts, PromptPart{
					Value:       builder.String(),
					IsComponent: true,
				})
				builder.Reset()
				inside = false
				continue
			}
		}

		builder.WriteByte(c)
	}

	if builder.Len() > 0 {
		parts = append(parts, PromptPart{
			Value:       builder.String(),
			IsComponent: inside,
		})
	}

	return parts
}

func GenPrompt(config models.PromptConfig) (string, error) {
	if config.Prompt == "" {
		config.Prompt = "{cwd} > "
	}

	parts := SplitPrompt(config.Prompt)
	components := BuildComponentMap(config)

	resultChans := make(map[string]<-chan models.Result)

	for _, part := range parts {
		if part.IsComponent {
			if comp, ok := components[part.Value]; ok {
				resultChans[part.Value] = comp.RenderAsync()
			}
		}
	}

	results := make(map[string]models.Result)
	for key, ch := range resultChans {
		results[key] = <-ch
	}

	var builder strings.Builder

	for _, part := range parts {
		if part.IsComponent {
			_, ok := components[part.Value]
			if !ok {
				builder.WriteString("{" + part.Value + "}")
				continue
			}
			result := results[part.Value]
			if result.Skip {
				continue
			}
			builder.WriteString(result.Value)
		} else {
			builder.WriteString(part.Value)
		}
	}

	return builder.String(), nil
}
