package updater

import (
	"fmt"
	"strings"

	"github.com/apex/log"
)

var NAME = "Updater Gradle"
var FUVERSION = "dev"

type Updater struct {
	VersionKey string
	TrimTag    string
	NewVersion string
}

func (u *Updater) Init(m map[string]string) error {
	log.Infof("Init %v", m)
	vKey := m["version-key"]
	if len(vKey) == 0 {
		vKey = "version"
	}
	u.VersionKey = vKey
	u.TrimTag = m["trim-tag"]
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

	u.NewVersion = newVersion
	if len(u.TrimTag) != 0 {
		u.NewVersion = strings.Trim(newVersion, u.TrimTag)
	}

	config, err := ReadPropertiesFile(file, true)
	if err != nil {
		return err
	}

	for k, v := range config {
		originalKey := k[strings.Index(k, ";")+1:]
		if originalKey == u.VersionKey {
			if v != u.NewVersion {
				config[k] = u.NewVersion
			}
			break
		}
	}

	if err := WritePropertiesFile(file, config); err != nil {
		return err
	}
	return nil
}
