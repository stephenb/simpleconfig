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
	Usage    func()                 // Can be re-implemented if needed
	flagSet  *flag.FlagSet
}

func NewConfig(opts map[string]string, jsonPath string) *Config {
	conf := Config{Options: opts, JsonPath: jsonPath, Map: make(map[string]interface{})}
	fs := flag.NewFlagSet("simpleconfig", flag.ExitOnError)
	conf.Usage = conf.usage
	fs.Usage = conf.usage
	conf.flagSet = fs
	return &conf
}

func (conf *Config) Parse() error {
	// Always support config-path flag
	configPathFlag := conf.flagSet.String("config-path", "", "Path to JSON-formatted config file.")

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

	// override/set any vals that are set in env vars
	for key, _ := range conf.Options {
		if os.Getenv(key) != "" {
			conf.Map[key] = os.Getenv(key)
		}
	}

	// override/set any vals that are in the conf map
	for key, val := range flagMap {
		if *val != "" {
			conf.Map[key] = *val
		}
	}

	return nil
}

func (conf *Config) usage() {
	fmt.Println("USAGE:")
	fmt.Println("\t", os.Args[0], "[-config-path=\"path/to/json/config/file\"] [setting flags]")
	fmt.Println("")
	fmt.Println("\tSettings will be read from the specified JSON file, ENV vars set matching any of ")
	fmt.Println("\tthe specified settings, and matching setting flags used when the command is run.")
	fmt.Println("\tOrder of precendence is CLI flags, then ENV vars, then JSON file settings.")
	fmt.Println("")
	fmt.Println("\t", os.Args[0], "supports the following settings:")
	conf.flagSet.VisitAll(func(f *flag.Flag) {
		fmt.Println("\t\t", f.Name, ":", f.Usage)
	})
}
