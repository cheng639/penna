
<div align=center>

</div>
<div align=center>
<img src="https://img.shields.io/badge/golang-1.24-blue"/>
<img src="https://img.shields.io/badge/gin-1.12.0-lightBlue"/>
<img src="https://img.shields.io/badge/gorm-1.31.1-red"/>
</div>



## 1. 基本介绍

### 1.1 项目介绍


> penna是一个基于gorm和gin开发的后台web框架，适合中小型项目开发RESTful API。对于常规的CURD接口，则只需定义Model文件，
> 和简单的控制器类(在继承base.Controller的前提下)，再定义好路由，即开发完成。无需再编写对应的控制器方法。当然在基类的方法
> 不适合具体情况时，在对应的Controller下添加同名方法即可覆盖，十分方便灵活。



### 1.2 版本列表
- master: 2.0.1

## 2. 使用说明

```
- 使用git克隆本项目
    - ```git
        git clone xxxx.
        go mod tidy
      ```
- golang版本 >= v1.24
- IDE推荐：Goland
- 初始化项目： 配置文件中添加数据库， Redis连接参数,否则项目无法启动
-编译运行，请将编译后的可执行文件输出到cmd目录，或将config.yaml移动到可执行文件所在目录，如果是Goland编辑器，可以在运行-编辑配置修改
输出目录和工作目录
--依赖管理，本项目使用go mod vendor管理依赖项，运行go mod vendor命令可将依赖文件添加/更新到/vendor目录，这样可以在内网/无网络环境
下编译代码，如果您用不到，可以忽略此项，按默认的go mod方式使用即可，Goland会自动处理依赖项
```
### 2.1 使用示例
```
-创建表，将/model/sql目录下的category.sql文件的表创建语句添加到你的数据库。本项目不采用表自动迁移机制，建议通过执行sql脚本实施表结构变更
。中大型项目的或生产环境中，grom表自动迁移应当关闭。否则你的服务可能因为表结构迁移执行时间太长而无法启动运行。

-编译运行后，通过curl或postman访问/router/api.vi.go下定义的categories路由。这些路由对应的控制器方法在base.Controller中实现。所以CategoryController
无需编写控制器方法即可实现CURD的功能。你可以在CategoryController或api.Controller重写同名方法以覆盖，或增加其他的控制器方法。

- 例如 GET请求 localhost:8090/api/v1/categories, GET请求localhost:8090/api/v1/categories/1, 其他接口请求参数请参考
/model/category.go文件下结构体binding注释
```



    


## 3. 技术选型


- 路由：用`Gin`快速搭建基础restful风格API，`Gin`是一个go语言编写的Web框架。
- 中间件：支持`Gin`中间件
- 验证器：Gin 可以解析并验证请求的输入参数，包括XML，Json, Form等，Gin使用 go-playground/validator/v10 进行验证

- 数据库：使用`gorm`实现对数据库的基本操作,已添加对sqlite数据库的支持。
- 缓存：使用`Redis` 缓存数据。
- 配置文件：Viper支持在运行时让应用程序实时读取配置文件。
- 日志：使用`zerolog`实现日志记录。


## 4. 项目架构

### 4.1 目录结构

```
    ├─gin-g  	        （文件夹）
    │  ├─cmd            （程序入口及配置文件目录）
    │  ├─config         （配置包）
    │  ├─common  	    （通用方法及常量）
    │  ├─middleware     （中间件）
    │  ├─model          （结构体层）
    │  ├─router         （路由）
    │  └─service	    （业务层）
    |  |-vendor	        （依赖项）

```


## 5.. 商用注意事项

如果您将此项目用于商业用途，请遵守Apache2.0协议并保留作者技术支持声明。
