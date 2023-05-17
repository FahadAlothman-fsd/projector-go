package projector

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

type Data struct {
	Projector map[string]map[string]string `json:"projector"`
}

type Projector struct {
	config Config
	data   Data
}

func CreateProjector(config *Config, data *Data) *Projector {
	return &Projector{
		config: *config,
		data:   *data,
	}
}

func (proj *Projector) GetValue(key string) (string, bool) {

	curr := proj.config.Pwd
	// fmt.Printf("curr: %v\n", curr)
	prev := ""

	out := ""
	found := false
	for curr != prev {
		// fmt.Printf("curr: %v\n", curr)
		if dir, ok := proj.data.Projector[curr]; ok {
			// fmt.Printf("dir: %v\n", dir)
			if value, ok := dir[key]; ok {
				out = value
				found = true
				break
			}
		}
		prev = curr
		curr = path.Dir(curr)

	}
	// fmt.Printf("out: %v\n", out)

	return out, found

}

func (proj *Projector) GetValueAll() map[string]string {

	out := map[string]string{}
	paths := []string{}

	curr := proj.config.Pwd
	prev := ""

	for curr != prev {
		paths = append(paths, curr)
		prev = curr
		curr = path.Dir(curr)
	}
	for i := len(paths) - 1; i >= 0; i-- {

		if dir, ok := proj.data.Projector[paths[i]]; ok {

			for key, value := range dir {
				out[key] = value
			}
		}
	}

	return out
}

func (proj *Projector) SetValue(key, value string) {
	pwd := proj.config.Pwd
	if _, ok := proj.data.Projector[pwd]; !ok {
		proj.data.Projector[pwd] = map[string]string{}

	}
	proj.data.Projector[pwd][key] = value

}

func (proj *Projector) RemoveValue(key string) (string, bool) {
	pwd := proj.config.Pwd
	value := ""
	if dir, ok := proj.data.Projector[pwd]; ok {
		fmt.Printf("dir: %v\n", dir)
		value = dir[key]
		delete(dir, key)
		fmt.Printf("dir: %v\n", dir)
	}
	fmt.Printf("proj: %v\n", proj.data.Projector[pwd])

	if value == "" {
		return value, false
	}
	return value, true
}

func defaultProjector(config *Config) *Projector {
	return &Projector{
		config: *config,
		data: Data{
			Projector: map[string]map[string]string{},
		},
	}
}

func NewProjector(config *Config) *Projector {

	if _, err := os.Stat(config.Config); !os.IsNotExist(err) {

		contents, err := os.ReadFile(config.Config)

		if err != nil {
			defaultProjector(config)
		}
		var data Data
		err = json.Unmarshal(contents, &data)
		if err != nil {
			return defaultProjector(config)
		}
		return &Projector{
			config: *config,
			data:   data,
		}
	}
	return defaultProjector(config)

}

func (proj *Projector) Save() error {
	dir := path.Dir(proj.config.Config)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}

	}

	jsonString, err := json.Marshal(proj.data)
	if err != nil {
		return err
	}
	ioutil.WriteFile(proj.config.Config, jsonString, 0644)
	return nil

}
