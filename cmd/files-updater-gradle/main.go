package main

import (
	"github.com/apex/log"
	"github.com/go-semantic-release/semantic-release/v2/pkg/plugin"
	"github.com/go-semantic-release/semantic-release/v2/pkg/updater"
	gradleUpdater "github.com/ted-vo/file-updater-gradle/pkg/updater"
)

func main() {
	log.SetHandler(gradleUpdater.NewLogHandler())
	plugin.Serve(&plugin.ServeOpts{
		FilesUpdater: func() updater.FilesUpdater {
			return &gradleUpdater.Updater{}
		},
	})
}
