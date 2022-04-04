package updater

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type AppConfigProperties map[string]string

func ReadPropertiesFile(fName string, order bool) (AppConfigProperties, error) {
	config := AppConfigProperties{}

	if _, err := os.Stat(fName); err != nil {
		return config, fmt.Errorf(fmt.Sprintf("'%s' file does not exist", fName))
	}

	file, err := os.OpenFile(fName, os.O_RDWR, 0)
	if err != nil {
		return config, err
	}
	defer file.Close()

	var index uint32 = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if equal := strings.Index(line, "="); equal >= 0 {
			key, val, err := propertyOf(line)
			if err != nil {
				continue
			}

			if order {
				keyWithOrder := fmt.Sprintf("%d-%s", index, key)
				config[keyWithOrder] = val
			} else {
				config[key] = val
			}

			index++
		}
	}

	return config, nil
}

func WritePropertiesFile(fName string, config map[string]string) error {
	if _, err := os.Stat(fName); err != nil {
		return fmt.Errorf(fmt.Sprintf("'%s' file does not exist", fName))
	}

	keys := make([]string, 0, len(config))
	for k := range config {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var content string
	for _, k := range keys {
		content += fmt.Sprintf("%s=%s\n", k[2:], config[k])
	}
	os.WriteFile(fName, []byte(content), 0666)
	return nil
}

func propertyOf(str string) (key, val string, err error) {
	split := strings.Split(strings.TrimSpace(strings.TrimRight(str, "\n")), "=")
	if len(split) == 2 {
		return strings.TrimSpace(split[0]), strings.TrimSpace(split[1]), nil
	}
	return "", "", fmt.Errorf(fmt.Sprintf("string '%s' can not be read as property.", str))
}
