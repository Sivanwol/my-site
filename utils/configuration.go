package utils

import "github.com/kataras/iris/v12"

func GetConfiguration() map[string]iris.Configuration {
	var cnf = make(map[string]iris.Configuration)
	var files []ConfigFileName
	files = FetchAllConfigFiles("./configs")
	for i := 0; i < len(files); i++ {
		var fileObj ConfigFileName
		fileObj = files[i]
		var cnfObj = iris.YAML("./configs/" + fileObj.FileNameWithExt)
		cnf[fileObj.FileName] = cnfObj
	}
	return cnf
}
