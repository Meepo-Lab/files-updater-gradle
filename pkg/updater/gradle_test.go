package updater

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGradleUpdater(t *testing.T) {
	require := require.New(t)
	updater := &Updater{
		VersionKey: "version",
	}

	defaultVer := "1.0.0-SNAPSHOT.297"
	nVer := "1.0.0-SNAPSHOT.298"

	gradlePropertiesPath := "../../test/gradle.properties"

	err := updater.Apply(gradlePropertiesPath, nVer)
	require.NoError(err)

	config, err := ReadPropertiesFile(gradlePropertiesPath, false)
	require.NoError(err)
	require.Equal(nVer, config["version"], nVer)

	err2 := updater.Apply(gradlePropertiesPath, defaultVer)
	require.NoError(err2)
	require.Equal(nVer, config["version"], defaultVer)
}

func TestWithTrimTagGradleUpdater(t *testing.T) {
	require := require.New(t)
	updater := &Updater{
		VersionKey: "version",
		TrimTag:    "-SNAPSHOT",
	}

	defaultVer := "1.0.0-SNAPSHOT.297"
	nVer := "1.0.0.298"

	gradlePropertiesPath := "../../test/gradle.properties"

	err := updater.Apply(gradlePropertiesPath, nVer)
	require.NoError(err)

	config, err := ReadPropertiesFile(gradlePropertiesPath, false)
	require.NoError(err)
	require.Equal(nVer, config["version"], nVer)

	updater.TrimTag = ""
	err2 := updater.Apply(gradlePropertiesPath, defaultVer)
	require.NoError(err2)
	require.Equal(nVer, config["version"], defaultVer)
}
