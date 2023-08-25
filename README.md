# 飞布迁移程序说明

新版飞布相对于老版做了很多改进和调整，现已不再兼容旧版，为了得到更好的开发与使用体验，请尽快升级至新版飞布，并参照本文档进行迁移老项目。已升级的可忽略。

## 使用

可直接下载本仓库提供的可执行文件：https://github.com/fireboomio/fb-migration/releases/tag/0.0.1

置于项目的根目录下，运行该文件，可看到项目下生成的文件，旧版的文件放置于`old`目录下，如有需要可对照查看

## 注意事项

⚠️**不要多次执行迁移程序**，因为这涉及到文件目录的操作，如果你想恢复到以前的版本，可以执行回滚命令

**回滚命令**：`./fb-migration-amd64 -rollback` 如果因为某些意外迁移失败，可以恢复成旧版

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

手动在控制台升级钩子模板，新版模板的内容会输出到`template`目录下，模板的配置项可以在`store/sdk`目录下查看，以`golang-server`为例，它的配置文件如下：

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

当运行迁移程序时，会要求`输入需要修改的目录名称(例如：custom-go 回车换行可输入多个，两次回车结束输入；若无需修改，直接回车)`

其实就是将钩子的目录名称进行了变更，无需再手动操作，假如你不需要变更模板（如custom-go/custom-ts），直接回车即可

如果需要变更，可以输出钩子模板的路径，并且可以输出多个，**注意这里仅仅是针对服务端模板，不是客户端模板**



## 其他升级内容

1. 新增fromClaim参数customConfig，入参jsonPathComponents，填写参数路径，从user的customClaims中获取参数，类型，必填，名称根据schema自动解析
2. **新增@transaction参数，允许mutation执行事务**，要求同一个数据源且为数据库类型
3. 新增编译orderby，or数组类型参数判断，若参数为object且子参数数量大于1时提示非法
4. 底层存储升级，目录结构更加清晰，优化底层逻辑，采用单个json文件存储各个模块配置
5. 优化编译生成的启动配置，复用缓存配置，减少graphql编译次数，**大幅缩小配置文件大小和编译时间**
6. **钩子customize，proxy、function**由生成-启动-编译方式改为由钩子服务通过health上报，自动更新对应的存储
7. **模版下载提供指定分支和commitHash方式，**重新部署后自动下载对应commitHash版本的模板
8. 新增prisma数据源，上传prisma文件来实现数据源的查询操作
9. 新增导入导出功能，各个模块可以根据名称筛选导出压缩包，并导入到其他项目，包括api，数据源，s3，sdk，role等
10. 飞布web控制台对应的swagger升级，采用通用路由配合占位符，范型动态解析真实路由
11. **生成`graphql api`时`Decimal` `bytes`去掉前缀，json变全大写**
