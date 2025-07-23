package env

import "os"

var (
	HOME_ENV          = os.Getenv("HOME")
	PULSARSHIP_CONFIG = os.Getenv("PULSARSHIP_CONFIG")
	PULSARSHIP_SHELL  = os.Getenv("PULSARSHIP_SHELL")
)
