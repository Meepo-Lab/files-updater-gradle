package updater

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
)

type AppConfigProperties map[string]string

func ReadPropertiesFile(fName string, order bool) (AppConfigProperties, error) {
	config := AppConfigProperties{}

	filePath := path.Join(path.Dir(fName), "gradle.properties")
	if _, err := os.Stat(filePath); err != nil {
		return config, fmt.Errorf(fmt.Sprintf("'%s' file does not exist", fName))
	}

	file, err := os.OpenFile(filePath, os.O_RDWR, 0)
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
				keyWithOrder := fmt.Sprintf("%d;%s", index, key)
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
	filePath := path.Join(path.Dir(fName), "gradle.properties")
	if _, err := os.Stat(filePath); err != nil {
		return fmt.Errorf(fmt.Sprintf("'%s' file does not exist", fName))
	}

	mapSortKey := map[int]string{}
	for k := range config {
		index, err := strconv.Atoi(k[0:strings.Index(k, ";")])
		if err != nil {
			log.Fatalln("Error when indexing properties")
		}
		value := k[strings.Index(k, ";")+1:]
		mapSortKey[index] = value
	}

	keys := make([]int, 0)

	for k := range mapSortKey {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	var content string
	for k := range keys {
		key := mapSortKey[k]
		content += fmt.Sprintf("%s=%s\n", key, config[fmt.Sprintf("%d;%s", k, key)])
	}
	if err := os.WriteFile(filePath, []byte(content), 0666); err != nil {
		return err
	}
	return nil
}

func propertyOf(str string) (key, val string, err error) {
	split := strings.Split(strings.TrimSpace(strings.TrimRight(str, "\n")), "=")
	if len(split) == 2 {
		return strings.TrimSpace(split[0]), strings.TrimSpace(split[1]), nil
	}
	return "", "", fmt.Errorf(fmt.Sprintf("string '%s' can not be read as property.", str))
}
