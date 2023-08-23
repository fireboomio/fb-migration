package migrate

import (
	"fireboom-migrate/consts"
	"fireboom-migrate/utils"
	"path/filepath"
)

func Start() {

	utils.CopyFile(".env", filepath.Join(consts.BackendPath, ".env"))

	migrateRole()
	migrateDatasource()
	migrateSdk()
	migrateGlobalSettings()
	migrateOperation()
	migrateGlobalOperation()
	migrateStorage()
	migrateEnv()

	utils.MoveDir()
	utils.RenameCustom()
}
