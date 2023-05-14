package projector

import (
	"encoding/json"
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

func (proj *Projector) GetValue(key string) (string, bool) {

	curr := proj.config.Pwd
	prev := ""

	out := ""
	found := false
	for curr != prev {
		if dir, ok := proj.data.Projector[curr]; ok {
			if value, ok := dir[key]; ok {
				out = value
				found = true
				break
			}
		}
		prev = curr
		curr = path.Dir(curr)

	}

	return out, found

}

func (proj *Projector) GetValueAll(key string, value string) map[string]string {

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

func (proj *Projector) setValue(key, value string) {
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
		value = dir[key]
		delete(dir, key)
	}
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

	if _, err := os.Stat(config.Config); err != nil {

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
