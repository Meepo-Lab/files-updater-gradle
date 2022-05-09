package main

import (
	"github.com/apex/log"
	gradleUpdater "github.com/ted-vo/file-updater-gradle/pkg/updater"
	"github.com/ted-vo/semantic-release/v3/pkg/plugin"
	"github.com/ted-vo/semantic-release/v3/pkg/updater"
)

func main() {
	log.SetHandler(gradleUpdater.NewLogHandler())
	plugin.Serve(&plugin.ServeOpts{
		FilesUpdater: func() updater.FilesUpdater {
			return &gradleUpdater.Updater{}
		},
	})
}
