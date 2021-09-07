package config

import (
	"fmt"

	"github.com/amido/stacks-cli/internal/constants"
)

// import (
// 	"github.com/dnitsch/scaffold/internal/util"
// 	"github.com/dnitsch/scaffold/pkg/scaffold"
// )

// func ReadSelfConfigFile(input scaffold.InputConfig) (*scaffold.SelfConfig, error) {
// 	return readSelfConfigFile(input)
// }

// // readSelfConfigFile constructs self config for CLI based on bundle resources
// func readSelfConfigFile(input scaffold.InputConfig) (*scaffold.SelfConfig, error) {

// 	sharedT, err := ParseShared()
// 	util.CheckErrors(err)

// 	specificT, err := ParseSpecific(input.Platform, input.Deployment, input.ProjectType)

// 	s := scaffold.SelfConfig{
// 		Shared:   &sharedT,
// 		Specific: &specificT,
// 	}
// 	// TODO: feat request allow overwrite of self config from outside (as long as it can be parsed back to a SelfConfig)
// 	return &s, err
// }

type TypeDetail struct {
	Gitrepo                  string    `yaml:"git_repo"`
	Gitref                   string    `yaml:"git_ref"`
	Localpath                string    `yaml:"local_path"`
	FilenameReplacementPaths []string  `yaml:"filename_replacement_paths,omitempty"`
	Searchvalue              string    `yaml:"search_value,omitempty"`
	Foldermap                Foldermap `mapstructure:"folder_map"`
}

type Foldermap struct {
	Src  string `mapstructure:"src"`
	Dest string `mapstructure:"dest"`
}

type SelfConfig struct {
	Shared   *TypeDetail
	Specific *TypeDetail
}

type OutputConfig struct {
	TmpPath   string
	ZipPath   string
	UnzipPath string
	NewPath   string
}

type ReplaceConfig struct {
	Files  []string          `yaml:"files"`
	Values map[string]string `yaml:"values"`
}

type Config struct {
	Input   InputConfig
	Self    SelfConfig
	Output  OutputConfig
	Replace []ReplaceConfig
}

// GetVersion returns the current version of the application
// It will check to see uif the Version is empty, if it is, it will
// set and identifiable local build version
func (config *Config) GetVersion() string {
	var version string

	version = config.Input.Version

	fmt.Println(version)

	if version == "" {
		version = constants.DefaultVersion
	}

	return version
}

/*
// Create creates a config object based on parsed input config
func New(data InputConfig) (*Config, error) {
	tmpPath, err := os.MkdirTemp("", "source")
	if err != nil {
		return nil, err
	}

	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	selfConf, err := readSelfConfigFile(data)
	if err != nil {
		helper.TraceInfo("Failed to read self config")
		return nil, err
	}

	conf := Config{
		Output: &OutputConfig{
			NewPath:   fmt.Sprintf("%s/%s", pwd, data.ProjectName),
			TmpPath:   tmpPath,
			ZipPath:   fmt.Sprintf("%s/source.zip", tmpPath),
			UnzipPath: path.Join(tmpPath, "wip", selfConf.Specific.Localpath),
		},
		Input: &data,
		Self:  selfConf,
	}

	helper.TraceInfo(fmt.Sprintf("New Project Dir: %s\n", conf.Output.NewPath))

	helper.TraceInfo(fmt.Sprintf("Temp Path: %s\n", conf.Output.TmpPath))

	return &conf, err
}

// Create creates a config object based on bytes stream read from a config file
func NewBytes(data []byte) (*Config, error) {

	t := InputConfig{}

	if err := yaml.Unmarshal(data, &t); err != nil {
		return nil, err
	}

	conf, err := New(t)
	return conf, err
}

// readSelfConfigFile constructs self config for CLI based on bundle resources
func readSelfConfigFile(input InputConfig) (*SelfConfig, error) {
	// var err error

	sharedT, err := ParseLocalShared()
	// if errShared != nil {
	if err != nil {
		return nil, err
	}

	specificT, err := ParseLocalSpecific(input.Platform, input.Deployment, input.ProjectType)

	if err != nil {
		return nil, err
	}

	s := SelfConfig{
		Shared:   &sharedT,
		Specific: &specificT,
	}

	// TODO: feat request allow overwrite of self config from outside (as long as it can be parsed back to a SelfConfig)
	return &s, err
}
*/
