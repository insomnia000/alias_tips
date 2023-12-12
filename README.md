# AliasTips

 AliasTips是一个管理Linux命令别名alias的小工具，基于Golang实现。



### 一、前言

服务端同学可能会遇到一个问题是，随着学习到的Linux命令越来越多，有一些命令还特别长，很难记得住这么多命令。这时我们可以给这些命令起一个别名，以后只要输入别名就可以了

```
#启动 ElasticSearch
alias es_start='docker run -d --name elasticsearch -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" -e ES_JAVA_OPTS="-Xms64m -Xmx512m" elasticsearch:7.14.0 '
```



但是随着别名越来越多，我们甚至别名也会忘记了，要去shell配置文件翻查；这时候，我们可以写一个小工具，来统一维护我们的别名，aliasTips就诞生了。



aliasTips期望实现的效果是：

1、在命令行中输入 `tips init` 时，把当前目录下的 config 目录中所有 json文件的alias载入shell配置文件 `.bashrc` 或 `.zshrc` 中。

2、在命令行中输入 `tips` 时，列出所有管理的命令别名，此时等待用户输入一个命令编号，用户输入后，直接执行该编号对应的命令。

```
# tang @ macdeMacBook-Pro in ~/www_go/alias_tips  [18:03:04]
$ tips
+-----+--------+----------+--------------------------+------------------------------------+
| NUM | ALIAS  |   TAGS   |       DESCRIPTION        |              COMMAND               |
+-----+--------+----------+--------------------------+------------------------------------+
|   1 | cdgo   | comm     | 进入Go项目根目录         | cd /Users/shaochao/www_go && ls -l |
|   2 | tree2  | comm     | 显示2层目录结构          | tree -L 2                          |
|   3 | tree3  | comm     | 显示3层目录结构          | tree -L 3                          |
|   4 | gorun  | comm,go  | 编译运行main.go          | go run main.go                     |
|   5 | cdwww  | comm     | 进入网站根目录           | cd /Users/shaochao/www && ls -l    |
|   6 | sshdev | comm,ssh | 登录测试服务器           | ssh  root@120.79.5.254             |
|   7 | dkimg  | docker   | 查看docker镜像           | docker images                      |
|   8 | dkps   | docker   | 查看docker正在运行的容器 | docker ps                          |
+-----+--------+----------+--------------------------+------------------------------------+
请输入想要执行的命令编号，按Enter确定 (按ESC键或q键退出) ...
2
即将执行命令： tree -L 2
---------------------------------------------------------

.
├── README.md
├── config
│   ├── common.json
│   └── docker.json
├── go.mod
├── go.sum
├── logic
│   ├── alias.go
├── main
├── main.go
├── output_alias.sh
└── run.sh

3 directories, 15 files
```





3、命令`tips` 还接受一个tags参数，以供用户进行命令筛选，比如 	`tips docker` 只会列出docker相关的命令以供用户选择。用户可以在config文件中，给各个命令打 tag ，进行筛选。

```
# tang @ macdeMacBook-Pro in ~/www_go/alias_tips [18:03:36]
$ tips docker
+-----+-------+--------+--------------------------+---------------+
| NUM | ALIAS |  TAGS  |       DESCRIPTION        |    COMMAND    |
+-----+-------+--------+--------------------------+---------------+
|   1 | dkimg | docker | 查看docker镜像           | docker images |
|   2 | dkps  | docker | 查看docker正在运行的容器 | docker ps     |
+-----+-------+--------+--------------------------+---------------+
请输入想要执行的命令编号，按Enter确定 (按ESC键或q键退出) ...

```



这样我们只要把config文件夹中的alias维护好就好了，如果忘记了命令或者忘记了别名，直接 `tips` 看一下。

 

### 二、使用

1、下载