# Template
1.定义模版引擎文件，解析模版引擎文件，执行模版引擎文件
2.语法 
{{ . }} .表示传入的数据
{{ if }} {{ else }} {{ end }} 条件判断
{{ range }} {{else}} {{ end }} 循环便利
{{ selfFunc }} 通过在template定义自定义函数完成功能
{{ define template }} 模版嵌套
{{ block “content” .}} {{ end }} 模版继承 多个不同页面共用一套主样式 数据不同