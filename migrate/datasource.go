package migrate

import (
	"encoding/json"
	"fireboom-migrate/consts"
	"fireboom-migrate/types/cloud"
	"fireboom-migrate/types/origin"
	"fireboom-migrate/types/wgpb"
	"fireboom-migrate/utils"
	"fmt"
	"github.com/tidwall/gjson"
	"path/filepath"
	"strings"
)

var dbMap map[string]wgpb.DataSourceKind

func init() {
	dbMap = make(map[string]wgpb.DataSourceKind, 9)
	dbMap["static"] = wgpb.DataSourceKind_STATIC
	dbMap["rest"] = wgpb.DataSourceKind_REST
	dbMap["graphql"] = wgpb.DataSourceKind_GRAPHQL
	dbMap["postgresql"] = wgpb.DataSourceKind_POSTGRESQL
	dbMap["mysql"] = wgpb.DataSourceKind_MYSQL
	dbMap["sqlserver"] = wgpb.DataSourceKind_SQLSERVER
	dbMap["mongodb"] = wgpb.DataSourceKind_MONGODB
	dbMap["sqlite"] = wgpb.DataSourceKind_SQLITE
	dbMap["prisma"] = wgpb.DataSourceKind_PRISMA

	//扫描 upload目录，进行迁移--> upload-cloud
	if !utils.NotExistFile(consts.UploadDbpath) {
		utils.CopyDir(consts.UploadDbpath, consts.UploadCloudDbpath)
	}

	if !utils.NotExistFile(consts.UploadGraphqlPath) {
		utils.CopyDir(consts.UploadGraphqlPath, consts.UploadCloudGraphqlPath)
	}
}

// 迁移数据源
// store/list/FbDatasource --> store-cloud/datasource/xx.json
func migrateDatasource() {
	path := consts.StoreListDatasourcePath
	content, err := utils.ReadFile(path)
	if err != nil {
		fmt.Errorf("read file failed, err: %v\n", err.Error())
		return
	}

	var fds = make([]origin.FbDataSource, 0)
	_ = json.Unmarshal(content, &fds)

	for _, source := range fds {
		if source.Name == "system" {
			// system数据源无需处理
			continue
		}
		sourceBytes, _ := json.Marshal(source)
		var ds = &cloud.Datasource{
			Name:       source.Name,
			CreateTime: source.CreateTime,
			UpdateTime: source.UpdateTime,
			DeleteTime: source.DeleteTime,
			Enabled:    source.Enabled,
		}

		sourceStr := string(sourceBytes)

		// 旧版数据源类型: 1-db 2-rest 3-graphql 4-自定义
		// 新版数据源：0-static 1-rest 2-graphql 3-postgresql 4-mysql 5-sqlserver 6-mongodb 7-sqlite 8-prisma
		switch source.SourceType {
		case 1:
			dbType := strings.ToLower(gjson.Get(sourceStr, "config.dbType").String())
			ds.Kind = dbMap[dbType]
			ds.CustomDatabase = getCustomBase(sourceStr)
			break
		case 2:
			ds.Kind = wgpb.DataSourceKind_REST
			ds.CustomRest = &wgpb.CustomRest{
				OasFilepath: gjson.Get(sourceStr, "config.filePath").String(),
				BaseUrl:     gjson.Get(sourceStr, "config.baseURL").String(),
				Headers:     getHeaders(sourceStr),
			}
			_ = utils.CopyFile(filepath.Join(consts.UploadPath, "oas", ds.CustomRest.OasFilepath),
				filepath.Join(consts.UploadCloudPath, "oas", ds.CustomRest.OasFilepath))
			break
		case 3:
			ds.Kind = wgpb.DataSourceKind_GRAPHQL
			ds.CustomGraphql = &wgpb.CustomGraphql{
				Url:          filepath.Join("graphql", gjson.Get(sourceStr, "config.url").String()),
				Headers:      getHeaders(sourceStr),
				Customized:   false,
				SchemaString: "",
			}
			break
		case 4:
			ds.Kind = wgpb.DataSourceKind_GRAPHQL
			ds.CustomGraphql = &wgpb.CustomGraphql{
				Headers:      getHeaders(sourceStr),
				Customized:   true,
				SchemaString: gjson.Get(sourceStr, "config.loadSchemaFromString").String(),
			}
			break
		}

		// 序列化datasource，写入文件
		bytes, _ := json.MarshalIndent(ds, "", "    ")
		dsPath := utils.GetStoreCloudModelPath(consts.StoreCloudDatasourcePath, source.Name)
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

func getHeaders(source string) map[string]*wgpb.HTTPHeader {
	var headersMap = make(map[string]*wgpb.HTTPHeader)
	headers := gjson.Get(source, "config.headers").Array()
	if len(headers) == 0 {
		return nil
	}

	for _, header := range headers {
		headerKey := header.Get("key").String()
		val, ok := headersMap[headerKey]
		if ok {
			val.Values = append(val.Values, getHeader(header))
			continue
		}

		headersMap[headerKey] = &wgpb.HTTPHeader{Values: []*wgpb.ConfigurationVariable{
			getHeader(header),
		}}
	}
	return headersMap
}

func getHeader(header gjson.Result) *wgpb.ConfigurationVariable {
	res := &wgpb.ConfigurationVariable{}
	kind := header.Get("kind").Int()
	res.Kind = int32(kind)
	headerVal := header.Get("val").String()
	switch kind {
	case 0:
		// 值
		res.StaticVariableContent = headerVal
		break
	case 1:
		// 环境变量
		res.EnvironmentVariableName = headerVal
		break
	case 2:
		// 转发自客户端
		res.PlaceholderVariableName = headerVal
		break
	}

	return res
}

func getCustomBase(source string) *wgpb.CustomDatabase {
	customDs := &wgpb.CustomDatabase{}
	appendType := gjson.Get(source, "config.appendType").Int()
	customDs.Kind = appendType
	if appendType == 1 {
		// 连接参数
		customDs.DatabaseAlone = &wgpb.CustomDatabaseAlone{
			Host:     gjson.Get(source, "config.host").String(),
			Port:     gjson.Get(source, "config.port").Int(),
			Database: gjson.Get(source, "config.dbName").String(),
			Username: gjson.Get(source, "config.userName.val").String(),
			Password: gjson.Get(source, "config.password.val").String(),
		}
	} else {
		dbUrlKind := gjson.Get(source, "config.databaseUrl.kind").Int()
		customDs.DatabaseUrl = &wgpb.ConfigurationVariable{
			Kind: int32(dbUrlKind),
		}
		if dbUrlKind == 0 {
			// 静态url
			customDs.DatabaseUrl.StaticVariableContent = gjson.Get(source, "config.databaseUrl.val").String()
		} else {
			// 环境变量
			customDs.DatabaseUrl.EnvironmentVariableName = gjson.Get(source, "config.databaseUrl.key").String()
		}
	}

	return customDs
}
