package migrate

func Start() {
	migrateDatasource()
	migrateRole()
	migrateSdk()
	migrateGlobalSettings()
	migrateOperation()
	migrateGlobalOperation()
	migrateStorage()
	migrateEnv()
}
