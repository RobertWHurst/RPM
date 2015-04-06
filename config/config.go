package config

import "os"
import "strings"
import "github.com/mitchellh/go-homedir"
import "github.com/vaughan0/go-ini"

type Config struct {
	path string
	file ini.File
}

func New() Config {
	workingPath, _ := os.Getwd()
	configPath := findConfigPath(workingPath)
	configFile, _ := ini.LoadFile(configPath)
	return Config{configPath, configFile}
}

func (c *Config) Get(section string, key string) string {
	val, _ := c.file.Get(section, key)
	return val
}

func findConfigPath(workingPath string) string {
	paths := getConfigPaths(workingPath)
	for _, path := range paths {
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			return path
		}
	}
	return ""
}

func getHomePath(pathSep string) string {
	path, _ := homedir.Dir()
	if path != "" {
		path += pathSep + ".rpmrc"
	}
	return path
}

func getConfigPaths(workingPath string) []string {
	pathSep := string(os.PathSeparator)
	pathChunks := strings.Split(workingPath, pathSep)

	homePath := getHomePath(pathSep)
	isInHomePath := false
	if homePath == "" {
		isInHomePath = true
	}

	paths := make([]string, 24)
	for len(pathChunks) > 0 {
		path := strings.Join(append(pathChunks, ".rpmrc"), pathSep)
		if path == homePath {
			isInHomePath = true
		}
		paths = append(paths, path)
		pathChunks = pathChunks[:len(pathChunks)-1]
	}

	if !isInHomePath {
		paths = append(paths, homePath)
	}

	return paths
}
