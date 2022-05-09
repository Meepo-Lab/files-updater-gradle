# :open_file_folder: file-updater-gradle

[![CI](https://github.com/ted-vo/files-updater-gradle/workflows/CI/badge.svg?branch=main)](https://github.com/ted-vo/files-updater-gradle/actions?query=workflow%3ACI+branch%main)
[![Go Report Card](https://goreportcard.com/badge/github.com/ted-vo/files-updater-gradle)](https://goreportcard.com/report/github.com/ted-vo/files-updater-gradle)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/ted-vo/files-updater-gradle)](https://pkg.go.dev/github.com/ted-vo/files-updater-gradle)

The gradle files updater for [go-semantic-release](https://github.com/go-semantic-release/semantic-release).

Plugin update version in `grade.properties`

## Usage

``` json
{
  "plugins": {
    "files-updater": {
      "names": ["gradle"]
    }
  }
}
```

Use this plugin by enabling it via `--update gradle.properties` for the version update file.

## Licence

The [MIT License (MIT)](http://opensource.org/licenses/MIT)

Copyright © 2020 [Ted Vo](https://tedvo.dev)
