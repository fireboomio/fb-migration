package utils

import (
	"fireboom-migrate/consts"
	"path/filepath"
)

func JoinPathAndToSlash(path ...string) string {
	return filepath.ToSlash(filepath.Join(path...))
}

func GetExportedOperationsGraphqlPathWithBool(path string, enabled bool) string {
	path = path + consts.OperationApiSuffix + consts.OperationAPISwitchMap[enabled]
	return JoinPathAndToSlash(consts.ExportedOperationsPath, path)
}

func GetStoreCloudModelPath(dirname, pathName string) string {
	return JoinPathAndToSlash(dirname, pathName) + consts.JsonExt
}
