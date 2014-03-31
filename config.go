package simpleconfig

import(
	"flag"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"io/ioutil"
)

/*
All values can be passed via a JSON file or via command-line args.
If both are present, the command-line arg will be used.
*/
type Config struct {
	JsonPath string
	Map      map[string]interface{} // map of {title: parsed value}
	Options  map[string]string      // map of {title: description}
	flagSet  *flag.FlagSet
}

func NewConfig(opts map[string]string, jsonPath string) *Config {
	conf := Config{Options: opts, JsonPath: jsonPath}
	fs := flag.NewFlagSet("simpleconfig", flag.ExitOnError)
	fs.Usage = conf.Usage
	conf.flagSet = fs
	return &conf
}

func (conf *Config) Parse() error {
	// Always support config-path flag
	configPathFlag := flag.String("config-path", "", "Path to JSON-formatted config file.")

	// first, setup the flag options
	flagMap := make(map[string]*string)
	for opt, desc := range conf.Options {
		flagMap[opt] = conf.flagSet.String(opt, "", desc)
	}
	// parse them
	conf.flagSet.Parse(os.Args[1:])

	if *configPathFlag != "" {
		conf.JsonPath = *configPathFlag
	}

	if conf.JsonPath != "" {
		bytes, err := ioutil.ReadFile(conf.JsonPath)
		if err != nil {
			return errors.New("Unable to read JSON file: " + err.Error())
		}

		err = json.Unmarshal(bytes, &conf.Map)
		if err != nil {
			return errors.New("Unable to unmarshal JSON file: " + err.Error())
		}
	}

	// override/set any flags that are in the conf map
	for key, val := range flagMap {
		if *val != "" {
			conf.Map[key] = *val
		}
	}

	return nil
}

func (conf *Config) Usage() {
	fmt.Println("USAGE")
}
