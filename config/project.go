// Package config implements configuration variables
//
// project configuration
package config

import (
	"path/filepath"

	"supermercato24.it/configuration"
)

const (
	Name    = "colander"
	Version = "0.0.0"
)

func init() {
	configuration.DirProject = filepath.Join(configuration.DirProject, Name)
	configuration.BuildProject()
	configuration.DirBinStorage = filepath.Join(configuration.DirProject, configuration.PathStorage)
}