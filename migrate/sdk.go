package migrate

import (
	"encoding/json"
	"fireboom-migrate/consts"
	"fireboom-migrate/types/cloud"
	"fireboom-migrate/types/origin"
	"fireboom-migrate/utils"
	"fmt"
)

var langExtMap map[string]string

func init() {
	langExtMap = make(map[string]string)
	langExtMap["go"] = ".go"
	langExtMap["typescript"] = ".ts"
	langExtMap["dart"] = ".dart"
}

// 迁移sdk
// store/list/FbSDK --> store-cloud/sdk/xx.json
func migrateSdk() {
	path := consts.StoreListSDKPath
	content, err := utils.ReadFile(path)
	if err != nil {
		fmt.Errorf("read file failed, err: %v\n", err.Error())
		return
	}

	var sdks = make([]origin.FbSDK, 0)
	err = json.Unmarshal(content, &sdks)
	if err != nil {
		fmt.Errorf("unmarshal failed, err: %v\n", err.Error())
		return
	}

	for _, sdk := range sdks {
		newSdk := &cloud.Sdk{
			Name:        sdk.Name,
			Enabled:     sdk.Enabled,
			Type:        sdk.Type,
			Language:    sdk.Language,
			Extension:   langExtMap[sdk.Language],
			GitUrl:      sdk.Url,
			GitBranch:   "V2.0",
			OutputPath:  sdk.OutputPath,
			CreateTime:  sdk.CreateTime,
			UpdateTime:  sdk.UpdateTime,
			DeleteTime:  sdk.DeleteTime,
			Icon:        sdk.Icon,
			Title:       sdk.Title,
			Author:      sdk.Author,
			Version:     sdk.Version,
			Description: sdk.Description,
		}

		// 序列化sdk, 写入文件
		bytes, err := json.MarshalIndent(newSdk, "", "    ")
		if err != nil {
			fmt.Errorf("marshal newSdk failed, err: %v\n", err.Error())
			return
		}

		sdkPath := utils.GetStoreCloudModelPath(consts.StoreCloudSDKPath, sdk.Name)
		if utils.NotExistFile(sdkPath) {
			utils.CreateFile(sdkPath)
		}

		err = utils.WriteFile(sdkPath, bytes)
		fmt.Printf("write file %s\n", sdkPath)
		if err != nil {
			fmt.Errorf(err.Error())
		}
	}
}
