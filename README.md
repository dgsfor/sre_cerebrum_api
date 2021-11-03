<h1 align="center">SRE稳定性运营平台后端</h1>
<div align="center">
一个致力于SRE稳定性运营工作的平台
</div>

<div align="center">

[![License](https://img.shields.io/npm/l/package.json.svg?style=flat)](https://github.com/gsgs-libin/sre_cerebrum/blob/main/LICENSE)

</div>

项目
----
整体项目分为前端和后端：(后续前端和后端会同步更新)

> 前端：[sre_cerebrum](https://github.com/gsgs-libin/sre_cerebrum)

> 后端：[sre_cerebrum_api](https://github.com/gsgs-libin/sre_cerebrum_api)

环境和依赖
----

- go
- mysql

项目下载和运行
----

- 拉取项目代码
```bash
git clone https://github.com/gsgs-libin/sre_cerebrum_api
cd sre_cerebrum_api
```

- 开发模式运行
```
go run main.go --env test
```

- 编译项目
```
sh build.sh
```



其他说明
----

跑项目需要先有mysql，配置可以在这里修改: 

> 测试环境：[点击我](/conf/config/test.ini)

> 生产环境：[点击我](/conf/config/release.ini)
