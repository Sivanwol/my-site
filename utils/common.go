package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

type ConfigFileName struct {
	FileName        string
	FileNameWithExt string
}

func FetchAllConfigFiles(path string) []ConfigFileName {
	var returnValue []ConfigFileName
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		f, err := os.Open(path)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		files, err := f.Readdir(0)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		for _, v := range files {
			if !v.IsDir() {
				if filepath.Ext(v.Name()) == ".yml" || filepath.Ext(v.Name()) == ".yaml" {
					var obj ConfigFileName
					obj.FileName = fileNameWithoutExtSliceNotation(v.Name())
					obj.FileNameWithExt = v.Name()
					returnValue = append(returnValue, obj)
				}
			}
		}
	}
	return returnValue
}

func fileNameWithoutExtSliceNotation(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}
