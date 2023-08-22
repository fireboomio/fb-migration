package migrate

import (
	"encoding/json"
	"fireboom-migrate/consts"
	"fireboom-migrate/types/cloud"
	"fireboom-migrate/types/origin"
	"fireboom-migrate/types/wgpb"
	"fireboom-migrate/utils"
	"fmt"
	"path/filepath"
)

func migrateStorage() {
	path := consts.StoreListStoragePath
	content, err := utils.ReadFile(path)
	if err != nil || len(content) == 0 {
		return
	}

	storages := make([]origin.FbStorageBucket, 0)
	err = json.Unmarshal(content, &storages)
	if err != nil {
		fmt.Errorf(err.Error())
		return
	}

	for _, storage := range storages {
		newStorage := &cloud.Storage{
			Enabled:    storage.Enabled,
			CreateTime: storage.CreateTime,
			UpdateTime: storage.UpdateTime,
			DeleteTime: storage.DeleteTime,
			S3UploadConfiguration: wgpb.S3UploadConfiguration{
				Name: storage.Config.Name,
				Endpoint: &wgpb.ConfigurationVariable{
					Kind:                  0,
					StaticVariableContent: storage.Config.EndPoint,
				},
				AccessKeyID:     utils.GetConfigurationVariable(storage.Config.AccessKeyID),
				SecretAccessKey: utils.GetConfigurationVariable(storage.Config.SecretAccessKey),
				BucketName: &wgpb.ConfigurationVariable{
					Kind:                  0,
					StaticVariableContent: storage.Config.BucketName,
				},
				BucketLocation: &wgpb.ConfigurationVariable{
					Kind:                  0,
					StaticVariableContent: storage.Config.BucketLocation,
				},
				UseSSL:         storage.Config.UseSSL,
				UploadProfiles: storage.Config.UploadProfiles,
			},
		}

		bytes, _ := json.MarshalIndent(newStorage, "", "    ")
		targetStoragePath := filepath.Join(consts.StoreCloudStoragePath, storage.Config.Name) + consts.JsonExt
		err = utils.WriteFile(targetStoragePath, bytes)
		if err != nil {
			fmt.Errorf(err.Error())
		}
		fmt.Printf("write file %s\n", targetStoragePath)
	}

}
