# 飞布迁移程序说明

新版飞布相对于老版做了很多改进和调整，现已不再兼容旧版，为了得到更好的开发与使用体验，请尽快升级至新版飞布，并参照本文档进行迁移老项目。已升级的可忽略。

## 使用

可直接下载本仓库提供的可执行文件：https://github.com/fireboomio/fb-migration/releases/tag/0.0.1

置于项目的根目录下，运行该文件，可看到项目下生成的文件，旧版的文件放置于`old`目录下，如有需要可对照查看

回滚命令：`./fb-migration-amd64 -rollback` 如果因为某些意外迁移失败，可以恢复成旧版

新旧版本对比如下：

- 旧版`store` 目录

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

- 新版`store`目录

  ```
  store
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

- 新版`upload`目录

  ```
  upload
  ├── graphql
  ├── oas
  │   └── casdoor.json
  └── sqlite
  ```
  



🎉重大更新🎉：

🌟`store/list/`目录下的配置文件不再以伪json格式存储，全部拆分迁移至新版`store`目录下了

🌟`store/object`目录下的全局配置文件和系统配置文件在新版已合并成`store/global.setting.json`和`global.operation.json`

🌟`graphql`文件和`operation`的配置文件不再分家， 全部存放至`store/operation`目录下

🌟`upload`目录下的db文件夹更名为`sqlite`

## 钩子模板的更新

本次更新对钩子模板也进行了更新，适配了新版的飞布，如果在项目中使用了钩子模板，比如`Golang-server`或者`node-server`，在升级模板后需要做一点手动的适配，大部分代码我们都是兼容的，只有几处小的地方需要开发者手动处理，以`Golang-server`为例：

### 1. 修改模板分支来升级为新的模板

手动在控制台升级钩子模板，新版模板的内容会输出到`template-cloud`目录下，模板的配置项可以在`store-cloud/sdk`目录下查看，以`golang-server`为例，它的配置文件如下：

```json
{
	"name": "golang-server",
	"enabled": true,
	"type": "server",
	"language": "go",
	"extension": ".go",
	"gitUrl": "https://code.100ai.com.cn/fireboomio/sdk-template_go-server.git",
	"gitBranch": "V2.0",
	"outputPath": "./custom-go",
	"createTime": "2023-08-21T18:51:18+08:00",
	"updateTime": "2023-08-21T18:51:28+08:00",
	"deleteTime": "",
	"icon": "...",
	"title": "Golang server",
	"author": "fireboom",
	"version": "latest",
	"description": "Golang hook server SDK template for fireboom"
}
```

新版增加了`gitBranch`选项，可以从新版的分支拉取最新的模板，目前新版在`V2.0`分支



### 2. cutomize目录下创建的自定义数据源

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

### 3. proxys目录下注册的自定义路由函数

- `proxys` **目录名称变更为** `proxy`，避免歧义。同时可以传入**RBAC权限认证的配置项**

  示例代码：

  ```go
  package proxy
  
  import (
  	"custom-go/pkg/base"
  	"custom-go/pkg/plugins"
  	"custom-go/pkg/wgpb"
  	"net/http"
  )
  
  func init() {
  	plugins.RegisterProxyHook(ping, conf)
  }
  
  var conf = &plugins.HookConfig{
  	AuthRequired: true,
  	AuthorizationConfig: &wgpb.OperationAuthorizationConfig{
  		RoleConfig: &wgpb.OperationRoleConfig{
  			RequireMatchAny: []string{"admin", "user"},
  		},
  	},
  	EnableLiveQuery: false,
  }
  
  func ping(hook *base.HttpTransportHookRequest, body *plugins.HttpTransportBody) (*base.ClientResponse, error) {
  	// do something here ...
  	body.Response = &base.ClientResponse{
  		StatusCode: http.StatusOK,
  	}
  	body.Response.OriginBody = []byte("ok")
  	return body.Response, nil
  }
  ```

- ⚠️要在`main.go`中引入

  ```go
  import (
  	_ "custom-go/proxy"
  	"custom-go/server"
  )
  ```

## Rest数据源适配

新版`Rest`数据源支持`swagger3`的格式，旧版数据源是`swagger2`格式的，需要做适配性改造，才能使飞布内省出数据源的超图。

此处改造是针对之前使用过Rest数据源并上传过`swagger2`格式的json文件。如果没有更好的方式来升级成swagger3格式的文件，可以参考以下示例

```shell
## 安装npm包
npm install -g swagger2openapi
## 转换json文件
swagger2openapi swagger2.json -o swagger3.json
```

💡很重要的一点：**为每个api加上 `operationId`**，新版的飞布根据这个自定义的属性来生成超图，可以自定义一个相对简短的名称

```json
"/api/login": {
      "post": {
        "operationId": "login"
        ......
      }
}
```

## 目录名称变更

- `auth`目录变更为`authentication`
- `proxys`目录变更为`proxy`
- API钩子目录名称变更：`hooks/A/postResolve.go`---->`operation/A/postResovle.go`
