package main

import (
	"fireboom-migrate/migrate"
	"fireboom-migrate/utils"
	"flag"
	"fmt"
)

var rollBackOption = flag.Bool("rollback", false, "是否回滚")

func main() {
	flag.Parse()
	if *rollBackOption {
		fmt.Println("确定要回滚？y/n")
		var option string
		fmt.Scan(&option)
		if option == "y" {
			fmt.Println("回滚中...")
			utils.RollBack()
		}
		return
	}
	migrate.Start()
}
