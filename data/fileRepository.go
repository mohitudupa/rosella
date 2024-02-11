package data

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"

	"githib.com/mohitudupa/rosella/utils"
)

const (
	FormatYaml = "YAML"
	FormatJson = "JSON"
)

type Group struct {
	Flags   map[string]bool    `yaml:"flags" json:"flags"`
	Limits  map[string]float64 `yaml:"limits" json:"limits"`
	Values  map[string]string  `yaml:"values" json:"values"`
	Configs map[string]Config  `yaml:"configs" json:"configs"`
}

func NewGroup() *Group {
	return &Group{
		Flags:   make(map[string]bool),
		Limits:  make(map[string]float64),
		Values:  make(map[string]string),
		Configs: make(map[string]Config),
	}
}

type FileRepository struct {
	groups map[string]Group
	path   string
}

func NewFileRepository(path string) *FileRepository {
	return &FileRepository{groups: make(map[string]Group), path: path}
}

func (fr *FileRepository) loadYaml() error {
	yamlFile, err := os.ReadFile(fr.path)
	if err != nil {
		return errors.New("failed to locate yaml file at " + fr.path)
	}

	err = yaml.Unmarshal(yamlFile, fr.groups)
	if err != nil {
		return errors.New("failed to load yaml file at " + fr.path)
	}

	log.Printf("INFO: finished loading yaml sources for fileRepository")
	return nil
}

func (fr *FileRepository) loadJson() error {
	jsonFile, err := os.ReadFile(fr.path)
	if err != nil {
		return errors.New("failed to locate json file at " + fr.path)
	}

	err = json.Unmarshal(jsonFile, &fr.groups)
	if err != nil {
		return errors.New("failed to load json file at " + fr.path)
	}

	log.Printf("INFO: finished loading json sources for fileRepository")
	return nil
}

func (fr *FileRepository) Connect() error {
	splitPath := strings.Split(fr.path, ".")
	switch strings.ToUpper(splitPath[len(splitPath)-1]) {
	case FormatYaml:
		return fr.loadYaml()
	case FormatJson:
		return fr.loadJson()
	default:
		return errors.New("file format should be either one of " + FormatYaml + " or " + FormatJson)
	}
}

func (fr *FileRepository) Close() error {
	// Do nothing
	return nil
}

// Flag CRUD
func (fr *FileRepository) ListFlags(group string) ([]string, error) {
	data, ok := fr.groups[group]
	if !ok {
		return nil, &utils.GroupNotFound{}
	}
	return utils.Keys(data.Flags), nil
}

func (fr *FileRepository) GetFlag(group string, key string) (bool, error) {
	data, ok := fr.groups[group]
	if !ok {
		return false, &utils.GroupNotFound{}
	}

	result, ok := data.Flags[key]
	if !ok {
		return false, &utils.FlagNotFound{}
	}

	return result, nil
}

func (fr *FileRepository) SetFlag(group string, key string, value bool) error {
	data, ok := fr.groups[group]
	if !ok {
		fr.groups[group] = Group{}
		data = fr.groups[group]
	}

	data.Flags[key] = value
	return nil
}

func (fr *FileRepository) DeleteFlag(group string, key string) error {
	data, ok := fr.groups[group]
	if !ok {
		return nil
	}

	delete(data.Flags, key)
	return nil
}

// Limit CRUD
func (fr *FileRepository) ListLimits(group string) ([]string, error) {
	data, ok := fr.groups[group]
	if !ok {
		return nil, &utils.GroupNotFound{}
	}
	return utils.Keys(data.Limits), nil
}

func (fr *FileRepository) GetLimit(group string, key string) (float64, error) {
	data, ok := fr.groups[group]
	if !ok {
		return 0, &utils.GroupNotFound{}
	}

	result, ok := data.Limits[key]
	if !ok {
		return 0, &utils.LimitNotFound{}
	}

	return result, nil
}

func (fr *FileRepository) SetLimit(group string, key string, value float64) error {
	data, ok := fr.groups[group]
	if !ok {
		fr.groups[group] = Group{}
		data = fr.groups[group]
	}

	data.Limits[key] = value
	return nil
}

func (fr *FileRepository) DeleteLimit(group string, key string) error {
	data, ok := fr.groups[group]
	if !ok {
		return nil
	}

	delete(data.Limits, key)
	return nil
}

// Value CRUD
func (fr *FileRepository) ListValues(group string) ([]string, error) {
	data, ok := fr.groups[group]
	if !ok {
		return nil, &utils.GroupNotFound{}
	}
	return utils.Keys(data.Values), nil
}

func (fr *FileRepository) GetValue(group string, key string) (string, error) {
	data, ok := fr.groups[group]
	if !ok {
		return "", &utils.GroupNotFound{}
	}

	result, ok := data.Values[key]
	if !ok {
		return "", &utils.ValueNotFound{}
	}

	return result, nil
}

func (fr *FileRepository) SetValue(group string, key string, value string) error {
	data, ok := fr.groups[group]
	if !ok {
		fr.groups[group] = Group{}
		data = fr.groups[group]
	}

	data.Values[key] = value
	return nil
}

func (fr *FileRepository) DeleteValue(group string, key string) error {
	data, ok := fr.groups[group]
	if !ok {
		return nil
	}

	delete(data.Values, key)
	return nil
}

// Config CRUD
func (fr *FileRepository) ListConfigs(group string) ([]string, error) {
	data, ok := fr.groups[group]
	if !ok {
		return nil, &utils.GroupNotFound{}
	}
	return utils.Keys(data.Configs), nil
}

func (fr *FileRepository) GetConfig(group string, key string) (Config, error) {
	data, ok := fr.groups[group]
	if !ok {
		return nil, &utils.GroupNotFound{}
	}

	result, ok := data.Configs[key]
	if !ok {
		return nil, &utils.ConfigNotFound{}
	}

	return result, nil
}

func (fr *FileRepository) SetConfig(group string, key string, value Config) error {
	data, ok := fr.groups[group]
	if !ok {
		fr.groups[group] = Group{}
		data = fr.groups[group]
	}

	data.Configs[key] = value
	return nil
}

func (fr *FileRepository) DeleteConfig(group string, key string) error {
	data, ok := fr.groups[group]
	if !ok {
		return nil
	}

	delete(data.Configs, key)
	return nil
}
