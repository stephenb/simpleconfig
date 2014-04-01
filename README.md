# simpleconfig

Simple configuration reader for Go that handles a JSON settings file,
along with command line flags and env vars, at the same time.


## Install

	go get github.com/stephenb/simpleconfig


## Example Usage

Given the following Go program,

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
		config.JsonPath = "sample.json" // Could also use "config-path" CLI flag

		// 3. Parse the options from both the JSON settings file and CLI flags.
		err := config.Parse()
		if err != nil {
			panic(err.Error())
		}

		// 4. Config values are accessible via a normal map.
		fmt.Println("Config Map: ", config.Map)
	}


and sample.json containing,

	{
		"setting1": "Hello",
		"setting2": "World"
	}

running the program will yield the following output:

With no CLI fags:

	Config Map: map[setting1:Hello setting2:World]

With "-setting2=Josh" specified in the CLI flags:

	Config Map: map[setting1:Hello setting2:Josh]

Once the Config is parsed, config.Map can be used just like a normal Go map.
Note that the Config.Map has type "map[string]interface{}", which means you will
need to use type assertion when using setting values.

Also, a basic usage explanation will be shown if you pass the "-h" CLI flag. This
can be customized by re-implementing the Config's Usage function, similar to how
you would customize the same Usage output in Go's own flag package.


## NEW

Now ENV vars are also read into the configuration. The order of precedence is:
1. CLI Flags, 2. ENV Vars, 3. JSON Settings.


## Notes

- Path to JSON config file can also be specified via "config-path" command-line flag
- CLI args will only be parsed into strings
- Default values aren't supported
