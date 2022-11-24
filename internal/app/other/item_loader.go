package other

import (
	"fmt"
	"io/ioutil"
	"os"
)

func LoadItems(dirPath, tabs string) (m map[string]Recipe) {
	if files, err := ioutil.ReadDir(dirPath); err == nil {

		m = make(map[string]Recipe)
		for _, file := range files {
			if file.IsDir() {
				for name, r := range LoadItems(dirPath+string(os.PathSeparator)+file.Name(), tabs+"\t") {
					m[name] = r
				}
			} else {
				content, err := ioutil.ReadFile(dirPath + string(os.PathSeparator) + file.Name())
				if err != nil {
					continue
				}

				recipes, err := NewTokenParser().Parse(string(content))
				if err != nil {
					continue
				}

				for _, r := range recipes {
					m[r.Name] = r
				}
			}
		}
	} else {
		fmt.Printf("Cannot read %s", dirPath)
	}

	return m
}
