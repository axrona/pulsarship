package config

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/pelletier/go-toml/v2"
	env "github.com/xeyossr/pulsarship/internal"
	"github.com/xeyossr/pulsarship/internal/models"
	"github.com/xeyossr/pulsarship/internal/utils"
)

func DeepMerge(dst, src interface{}) {
	dv := reflect.ValueOf(dst).Elem()
	sv := reflect.ValueOf(src).Elem()

	for i := 0; i < dv.NumField(); i++ {
		df := dv.Field(i)
		sf := sv.Field(i)

		// Skip unexported fields
		if !df.CanSet() {
			continue
		}

		switch df.Kind() {
		case reflect.Ptr:
			if !sf.IsNil() {
				df.Set(sf)
			}

		case reflect.Struct:
			// recurse into struct
			dfPtr := df.Addr().Interface()
			sfPtr := sf.Addr().Interface()
			DeepMerge(dfPtr, sfPtr)

		case reflect.Map:
			if !sf.IsNil() {
				if df.IsNil() {
					df.Set(reflect.MakeMap(df.Type()))
				}
				for _, key := range sf.MapKeys() {
					df.SetMapIndex(key, sf.MapIndex(key))
				}
			}

		default:
			zero := reflect.Zero(df.Type())
			if !reflect.DeepEqual(sf.Interface(), zero.Interface()) {
				df.Set(sf)
			}
		}
	}
}

func ExpandPath(path string) string {
	if strings.HasPrefix(path, "~/") {
		path = filepath.Join(env.HOME_ENV, path[2:])
	}
	return path
}

func GetConfigPath(configFlag string) string {
	if configFlag != "" {
		return ExpandPath(configFlag)
	}

	if env.PULSARSHIP_CONFIG != "" {
		return ExpandPath(env.PULSARSHIP_CONFIG)
	}

	return filepath.Join(env.HOME_ENV, ".config", "pulsarship", "pulsarship.toml")
}

// Read and parse the configuration file at the given path.
func ParseConfig(file string) (models.PromptConfig, error) {
	var config models.PromptConfig

	data, err := os.ReadFile(file)
	if err != nil {
		return models.PromptConfig{}, fmt.Errorf("could not read config file: %w", err)
	}

	if err := toml.Unmarshal(data, &config); err != nil {
		return models.PromptConfig{}, fmt.Errorf("could not parse config file: %w", err)
	}

	if config.Import != nil && *config.Import != "" {
		importPath := filepath.Join(filepath.Dir(file), ExpandPath(*config.Import))
		importData, err := ParseConfig(importPath)
		utils.IfNotDebug(func() {
			if err != nil {
				return
			}
			DeepMerge(&config, &importData)
		}, func() {
			if err != nil {
				panic(fmt.Errorf("could not import %s: %w", importPath, err))
			}
		})
	}

	return config, nil
}

// Write the default config to the given path.
func WriteDefaultConfig(path string) error {
	data, err := toml.Marshal(DefaultConfig)
	if err != nil {
		return fmt.Errorf("could not encode default config: %w", err)
	}

	err = os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return fmt.Errorf("could not create config directory: %w", err)
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return fmt.Errorf("could not write config file: %w", err)
	}

	return nil
}
