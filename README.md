# 记录swagger使用示例
官方地址：https://swagger.io/

go版github地址：https://github.com/swaggo/swag

> Swagger 是一个规范和完整的框架，用于生成、描述、调用和可视化 RESTful 风格的 Web 服务。总体目标是使客户端和文件系统作为服务器以同样的速度来更新。文件的方法，参数和模型紧密集成到服务器端的代码，允许API来始终保持同步。

**作用：**
1. 接口的文档在线自动生成。
2. 功能测试。

### 项目运行
1. `go get all` 下载依赖包
2. 进入`swagger/app`目录，执行`swag init  -g cmd/main.go  -o ./docs`生成docs文件。
3. `routers.go` 引入docs目录 
```go
 import (
     _ "swagger/app/docs"
 )
 ```
 4. cd `cmd`目录，`go run main.go` 运行
 5. 浏览器输入地址：`http://127.0.0.1:8080/swagger/index.html` 访问
_____

### 需要注意的几处坑
1. **生成docs文件命令**
   如果解析的`api`目录和`main.go`文件是同级，使用以下命令即可：<br>
   <br>
   `swag init`
   <br>
   但是如果 main.go文件和本文示例一样放在cmd目录下，使用上面的命令是无法生成docs文件的。<br>
   <br>
   **解决方案：** 例如我的项目目录（进入到app目录） `C:\code\swagger\app`
    执行 `swag init  -g cmd/main.go  -o ./docs` 就可以成功生成了。
    <br>
2. **路由文件需要引入docs目录**
    &emsp;例如项目中的`routers.go`文件
    ```go
    import (
        _ "swagger/app/docs"
    )
    ```

3. 如果`main.go`文件中定义的host是`127.0.0.1:8080`，在浏览器使用`localhost:8080`进行访问是可以正常访问的，但是在发起请求测试时，会出现跨域的问题。需要把修改浏览器地址改为`127.0.0.1:8080`