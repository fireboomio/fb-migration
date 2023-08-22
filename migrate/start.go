package migrate

func Start() {
	migrateRole()
	migrateDatasource()
	migrateSdk()
	migrateGlobalSettings()
	migrateOperation()
	migrateGlobalOperation()
	migrateStorage()
	migrateEnv()
}
