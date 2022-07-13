package updater

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadProperties(t *testing.T) {
	require := require.New(t)
	gradlepropertiespath := "../../test/gradle.properties"
	p, err := ReadPropertiesFile(gradlepropertiespath)
	require.NotNil(p)
	require.NoError(err)
}

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

	config, err := ReadPropertiesFile(gradlePropertiesPath)
	require.NoError(err)
	require.Equal(nVer, config.MustGetString("version"), nVer)

	err2 := updater.Apply(gradlePropertiesPath, defaultVer)
	require.NoError(err2)
	require.Equal(nVer, config.MustGetString("version"), defaultVer)
}

func TestWithTrimTagGradleUpdater(t *testing.T) {
	require := require.New(t)
	updater := &Updater{
		VersionKey: "version",
		TrimTag:    "-COMMONS",
	}

	defaultVer := "1.0.0-COMMONS-SNAPSHOT.41"
	nVer := "1.0.0-SNAPSHOT.42"
	gradlePropertiesPath := "../../test/gradle.properties"

	err := updater.Apply(gradlePropertiesPath, nVer)
	require.NoError(err)

	config, err := ReadPropertiesFile(gradlePropertiesPath)
	require.NoError(err)
	require.Equal(nVer, config.MustGetString("version"), nVer)

	updater.TrimTag = ""
	err2 := updater.Apply(gradlePropertiesPath, defaultVer)
	require.NoError(err2)
	require.Equal(nVer, config.MustGetString("version"), defaultVer)
}
