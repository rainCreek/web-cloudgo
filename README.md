# web-cloudgo
开发简单 web 服务程序 cloudgo，了解 web 服务器工作原理。

环境：Window10/
curl-7.56.1/
Apache 2.4.29 x64

工具：VS Code



## 框架选择

此次选用的是  **martini**

理由：

* 简单，上手快。
* 可与其他 Go 的包配合工作；可轻松添加工具

最主要的是GitHub上有中文版文档，非常有助于理解。

-----------


## 测试

主要是使用curl和ab测试

### curl测试

**1.缺省测试**

用VS code终端输入：
```
go run main.go
```

打开浏览器输入：
```
http://localhost:8080
```

打开另一cmd，输入：
```
curl -v http://localhost:8080
```

![default](http://img.blog.csdn.net/20171114194156548)


**2.9090端口**

用VS code终端输入：
```
go run main.go -p9090
```


打开浏览器输入：
```
http://localhost:9090/hello/15331262raoyuxi
```

打开另一cmd，输入：
```
curl -v http://localhost:9090/hello/15331262raoyuxi_test_for_curl
```

![port9090](http://img.blog.csdn.net/20171114194216309)

-----------

### ab测试

用VS code终端输入：
```
go run main.go -p9090
```

打开另一cmd，输入：
```
ab -n 1000 -c 100 http://localhost:9090/hello/15331262raoyuxi_abtest
```

![ab](http://img.blog.csdn.net/20171114194246493)

-----------



## 杂记：遇到的一些小问题

问题其实不是很多，都是一些小问题。总结一下，为下次节省时间。

问题一：
![problem1](http://img.blog.csdn.net/20171114194303055)


解决方案：
这是个老问题了，用cmd在$GOPATH下get到martini即可。

```
go get -u github.com/go-martini/martini
```



问题二：编写server.go中的NewServer函数时，不能确定port number的类型。

解决方案：
用reflect输出port的类型看看。

![problem2](http://img.blog.csdn.net/20171114194315766)

