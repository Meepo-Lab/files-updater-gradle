package updater

import (
	"fmt"
	"strings"

	"github.com/apex/log"
)

var NAME = "Updater Gradle"
var FUVERSION = "dev"

type Updater struct {
}

func (u *Updater) Init(m map[string]string) error {
	return nil
}

func (u *Updater) Name() string {
	return NAME
}

func (u *Updater) Version() string {
	return FUVERSION
}

func (u *Updater) ForFiles() string {
	return "gradle\\.properties"
}

func (u *Updater) Apply(file, newVersion string) error {
	log.Infof(fmt.Sprintf("file=%s, newVersion=%s", file, newVersion))
	config, err := ReadPropertiesFile(file, true)
	if err != nil {
		return err
	}

	for k, v := range config {
		if equal := strings.Index(k, "version"); equal >= 0 {
			if v != newVersion {
				config[k] = newVersion
			}
			break
		}
	}

	if err := WritePropertiesFile(file, config); err != nil {
		return err
	}
	return nil
}
