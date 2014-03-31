# config

Simple configuration reader for Go that handles a JSON settings file along with command line flags at the same time.


## Install

	go get github.com/stephenb/config


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
		conf := simpleconfig.NewConfig(opts, "sample.json")
		// 3. Parse the options from both the JSON settings file and CLI flags.
		err := conf.Parse()

		if err != nil {
			panic(err.Error())
		}

		// 4. Config values are accessible via a normal map.
		fmt.Println(conf.Map)
	}

and sample.json containing,

	{
		"setting1": "Hello",
		"setting2": "World"
	}

running the program will yield the following output:

With no CLI fags:

	map[setting1:Hello setting2:World]

With "-another-setting=Josh" specified in the CLI flags:

	map[setting1:Hello setting2:Josh]


## Notes

- Path to JSON config file can also be specified via "config-path" command-line flag
- CLI args will only be parsed into strings
- Default values aren't supported
