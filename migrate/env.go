package migrate

import (
	"fireboom-migrate/consts"
	"fmt"
	"os"
)

// 增加ENV配置属性
func migrateEnv() {
	file, _ := os.OpenFile(".env", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	defer file.Close()

	// 追加写入环境变量
	envVariables := map[string]string{
		consts.PUBLIC_URL:   "http://localhost:9991",
		consts.INTERNAL_URL: "http://localhost:9991",
		consts.LISTEN_HOST:  "localhost",
		consts.LISTEN_PORT:  "9991",
		consts.LOG_LEVEL:    "debug",
	}

	file.Write([]byte("\n"))
	for key, value := range envVariables {
		os.Setenv(key, value)
		_, err := file.WriteString(fmt.Sprintf("%s=\"%s\"\n", key, value))
		if err != nil {
			fmt.Println("写入文件失败:", err)
			return
		}
	}
}
