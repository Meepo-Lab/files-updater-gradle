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
	log.Infof(fmt.Sprintf("file=%s, newVersion=%s, trimTag=%s", file, newVersion, u.TrimTag))

	u.NewVersion = newVersion
	if len(u.TrimTag) != 0 {
		u.NewVersion = strings.ReplaceAll(newVersion, u.TrimTag, "")
	}

	p, err := ReadPropertiesFile(file)
	if err != nil {
		return err
	}

	p.SetValue(u.VersionKey, u.NewVersion)
	if err := WritePropertiesFile(file, p); err != nil {
		return err
	}
	return nil
}
