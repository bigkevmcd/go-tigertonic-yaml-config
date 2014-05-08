package tigertonic_yaml

import (
	"github.com/bigkevmcd/go-tigertonic"
	"gopkg.in/v1/yaml"
	"io/ioutil"
	"os"
)

func init() {
	tigertonic.RegisterConfigExt(".yaml", ConfigureYAML)
}

// ConfigureYAML reads the given configuration file and unmarshals the YAML
// found into the given configuration structure.  For convenient use with
// the flags package, an empty pathname is not considered an error.
func ConfigureYAML(pathname string, i interface{}) error {
	f, err := os.Open(pathname)
	if err != nil {
		return err
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, i)
}
