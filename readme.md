# Template
1.定义模版引擎文件，解析模版引擎文件，执行模版引擎文件
2.语法 
{{ . }} .表示传入的数据
{{ if }} {{ else }} {{ end }} 条件判断
{{ range }} {{else}} {{ end }} 循环便利
{{ selfFunc }} 通过在template定义自定义函数完成功能
{{ define template }} 模版嵌套
{{ block “content” .}} {{ end }} 模版继承 多个不同页面共用一套主样式 数据不同
## Gin_Template

不同目录下的同名文件加载
静态文件的加载 r.Static(relativePath,fileSystemPath)

## Gin_function

返回数据 
获取参数 get query
        post form
        uri /path/:username/:id
参数绑定 绑定到结构体
文件上传 
请求重定向
路由和路由组 r.any/r.norouter