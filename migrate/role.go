package migrate

import (
	"encoding/json"
	"fireboom-migrate/consts"
	"fireboom-migrate/types/cloud"
	"fireboom-migrate/types/origin"
	"fireboom-migrate/utils"
	"fmt"
)

// 迁移角色
// store/list/FbRole --> store-cloud/role/xx.json
func migrateRole() {
	path := consts.StoreListRolePath
	content, err := utils.ReadFile(path)
	if err != nil {
		fmt.Errorf("read file failed, err: %v\n", err.Error())
		return
	}

	var roles = make([]origin.FbRole, 0)
	err = json.Unmarshal(content, &roles)
	if err != nil {
		fmt.Errorf("unmarshal failed, err: %v\n", err.Error())
		return
	}

	for _, role := range roles {
		var newRole = &cloud.Role{
			Code:       role.Code,
			Remark:     role.Remark,
			CreateTime: role.CreateTime,
			UpdateTime: role.UpdateTime,
			DeleteTime: role.DeleteTime,
		}

		// 序列化datasource，写入文件
		bytes, err := json.MarshalIndent(newRole, "", "    ")
		if err != nil {
			fmt.Errorf("marshal newRole failed, err: %v\n", err.Error())
			return
		}

		dsPath := utils.GetStoreCloudModelPath(consts.StoreCloudRolePath, role.Code)
		if utils.NotExistFile(dsPath) {
			utils.CreateFile(dsPath)
		}

		err = utils.WriteFile(dsPath, bytes)
		fmt.Printf("write file %s\n", dsPath)
		if err != nil {
			fmt.Errorf(err.Error())
		}
	}
}
