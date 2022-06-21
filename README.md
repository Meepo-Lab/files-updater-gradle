# :open_file_folder: file-updater-gradle

[![CI](https://github.com/meepo-lab/files-updater-gradle/workflows/CI/badge.svg?branch=main)](https://github.com/meepo-lab/files-updater-gradle/actions?query=workflow%3ACI+branch%main)
[![Go Report Card](https://goreportcard.com/badge/github.com/meepo-lab/files-updater-gradle)](https://goreportcard.com/report/github.com/meepo-lab/files-updater-gradle)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/meepo-lab/files-updater-gradle)](https://pkg.go.dev/github.com/meepo-lab/files-updater-gradle)

The gradle files updater for [semantic-release](https://github.com/meepo-lab/semantic-release).

Plugin update version in `grade.properties`

## Usage

Enable `--update` in your command line with value is path to your version files.
> Recommand at root of your project dir

e.g:

``` bash
./semantic-release \
      ...
      --update gradle.properties
```

And with config

In `.semrelrc`

``` json
{
  "plugins": {
    "files-updater": {
      "names": ["gradle"]
      "options": {
        "message": "ci(release):"
        "version-key": "customize-version-key-in-your-files"
      }
    }
  }
}
```

OR

In `command line`

``` bash
./semantic-release \
      --files-updater gradle \
      --files-updater-opt message="ci(release):" \
      --files-updater-opt version-key="customize-version-key-in-your-files" \
      --files-updater-opt trim-tag="text-will-be-trim-before-replace" \
      --update gradle.properties
```

| options	|  e.g	| description |
|---	    |---	  |---	        |
| `message`	| "ci(release): v1.0.0-SNAPSHOT.1" | prefix message commit with new version release  	|
| `version-key`	| customize-version-key-in-your-files | customize your key version in your gradle.properties files |
| `trim-tag`	| text will be trimed before replace to `gradle.properties` files | option to trim some prefix |

e.g:
``` gradle
org.gradle.parallel=true
org.gradle.jvmargs=-Xmx3000m
gRPCspringBootVersion=2.13.0.RELEASE
javaxValidationVersion=2.0.1.Final
version=1.0.0-SNAPSHOT.1
customize-version-key-in-your-files=1.0.0-CUSTOMIZE.1
```

Use this plugin by enabling it via `--update gradle.properties` for the version update file.

## Licence

The [MIT License (MIT)](http://opensource.org/licenses/MIT)

Copyright © 2020 [Ted Vo](https://tedvo.dev)
