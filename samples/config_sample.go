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
		"another-setting": "And another setting.",
	}
	// Can also be passed in / overridden by a "-config-path" CLI flag
        filePath := "./samples/sample.json"

	// 2. Create a new config instance.
	conf := simpleconfig.NewConfig(opts, filePath)
	// 3. Parse the options from both the JSON settings file and CLI flags.
	err := conf.Parse()

	if err != nil {
		panic(err.Error())
	}

	// 4. Config values are accessible via a normal map.
	fmt.Println(conf.Map)
}
