package updater

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"strings"

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
	p.Write(buf, properties.UTF8)
	if err := os.WriteFile(filePath, []byte(buf.String()), 0666); err != nil {
		return err
	}
	return nil
}

func propertyOf(str string) (key, val string, err error) {
	fmt.Println(str)
	split := strings.Split(str, "=")
	if len(split) != 2 {
		return "", "", fmt.Errorf(fmt.Sprintf("string '%s' can not be read as property.", str))
	}
	fmt.Println(fmt.Sprintf("k=%s,v=%s", split[0], split[1]))
	fmt.Println("")
	return split[0], split[1], nil
}
