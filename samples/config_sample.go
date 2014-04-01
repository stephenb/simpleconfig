package main

import (
	"fmt"
	"github.com/stephenb/simpleconfig"
)

func main() {
	// 1. Setup a map of supported options and their descriptions.
	opts := map[string]string {
		"setting1": "Setting one description.",
		"setting2": "Setting two description.",
	}

	// 2. Create a new config instance.
	config := simpleconfig.NewConfig(opts)
	config.JsonPath = "./samples/sample.json" // Could also use "config-path" CLI flag
	// 3. Parse the options from both the JSON settings file and CLI flags.
	err := config.Parse()

	if err != nil {
		panic(err.Error())
	}

	// 4. Config values are accessible via a normal map.
	fmt.Println("Config Map: ", config.Map)
}
