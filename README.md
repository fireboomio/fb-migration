# 飞布迁移程序说明

新版飞布相对于老版做了很多改进和调整，现已不再兼容旧版，为了得到更好的开发与使用体验，请尽快升级至新版飞布，并参照本文档进行迁移老项目。已升级的可忽略。

## 使用

可直接下载本仓库提供的可执行文件，置于项目的根目录下，运行该文件，可看到项目下生成的文件，新旧版本对比如下：

- `store` 目录已迁移至 `store-cloud`目录【目录结构也有所调整】

```
store
├── hooks
│   ├── auth
│   ├── customize
│   ├── global
│   ├── hooks
│   └── uploads
├── list
│   ├── FbAuthentication
│   ├── FbDataSource
│   ├── FbOperation
│   ├── FbRole
│   ├── FbSDK
│   └── FbStorageBucket
└── object
    ├── global_config.json
    ├── global_operation_config.json
    ├── global_system_config.json
    └── operations
```

```
store-cloud
├── config
│   ├── global.operation.json
│   └── global.setting.json
├── datasource
│   ├── main.json
│   ├── system.json
├── operation
│   ├── xxx.graphql
│   ├── xxx.json
├── role
│   ├── admin.json
│   └── user.json
├── sdk
│   └── golang-server.json
└── storage
    └── aliyun.json
```

🎉重大更新🎉：

🌟`store/list/`目录下的配置文件不再以伪json格式存储，全部拆分迁移至`store-cloud`目录下了

🌟`store/object`目录下的全局配置文件和系统配置文件在新版已合并成`store-cloud/global.setting.json`和`global.operation.json`

🌟`graphql`文件和`operation`的配置文件不再分家， 全部存放至`store-cloud/operation`目录下

🌟`upload`目录下的文件已迁移至`upload-cloud`目录

## 钩子模板的更新

本次更新对钩子模板也进行了更新，适配了新版的飞布，如果在项目中使用了钩子模板，比如`Golang-server`或者`node-server`，在升级模板后需要做一点手动的适配，大部分代码我们都是兼容的，只有几处小的地方需要开发者手动处理，以`Golang-server`为例：

### 1. cutomize目录下创建的自定义数据源

示例旧版代码：

```go
var StatisticsSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: graphql.NewObject(graphql.ObjectConfig{
		Name: "query",
		Fields: graphql.Fields{
			"GetMonthlySales": &graphql.Field{
				Type: graphql.NewList(graphql.NewObject(graphql.ObjectConfig{
					Name: "GetMonthlySales",
					Fields: graphql.Fields{
						"months":     &graphql.Field{Type: graphql.String},
						"totalSales": &graphql.Field{Type: graphql.Float},
					},
				})),
				Resolve: func(p graphql.ResolveParams) (res interface{}, err error) {
					_, _, err = plugins.ResolveArgs[any](p)
					if err != nil {
						return
					}
					return "ok", nil
				},
			},
		},
	}),
})
```

新版需要在`init`函数中添加一行代码：

```go
func init() {
	plugins.RegisterGraphql(&StatisticsSchema)
}
```

【注】：要记得在custom-go/main.go 中自行引入，否则自定义的数据源不会生效

```go
import (
	_ "custom-go/customize"
	"custom-go/server"
)
```

### 2. proxys目录下注册的自定义路由函数

### 3. function目录下注册的函数















注：如果迁移成功，旧的文件可以选择删除，包括`store`、`exported`和`upload`目录下的文件

