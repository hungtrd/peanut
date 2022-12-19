package config

import "os"

var Locale string

func getConfig() {
	Locale = os.Getenv("PJ_LOCALE")
}
