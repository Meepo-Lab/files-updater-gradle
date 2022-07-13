package updater

import (
	"bytes"
	"fmt"
	"os"
	"path"

	"github.com/magiconair/properties"
)

func ReadPropertiesFile(fName string) (*properties.Properties, error) {
	filePath := path.Join(path.Dir(fName), "gradle.properties")
	if _, err := os.Stat(filePath); err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("'%s' file does not exist", fName))
	}

	return properties.MustLoadFile(fName, properties.UTF8), nil
}

func WritePropertiesFile(fName string, p *properties.Properties) error {
	filePath := path.Join(path.Dir(fName), "gradle.properties")
	if _, err := os.Stat(filePath); err != nil {
		return fmt.Errorf(fmt.Sprintf("'%s' file does not exist", fName))
	}

	buf := new(bytes.Buffer)
	if _, err := p.Write(buf, properties.UTF8); err != nil {
		return err
	}
	if err := os.WriteFile(filePath, buf.Bytes(), 0666); err != nil {
		return err
	}
	return nil
}
