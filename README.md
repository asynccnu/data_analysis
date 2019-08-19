## 木犀团队 Go Web 工程模板

![](https://travis-ci.org/asynccnu/data_analysis_service_v1.svg?branch=master)

### 简介

Go HTTP 服务工程模板。参考自掘金小册[基于 Go 语言构建企业级的 RESTful API 服务](https://juejin.im/book/5b0778756fb9a07aa632301e)

主要依赖：gin + gorm + viper + lexkong/log

### Build and run

```
mkdir $GOPATH/src/github.com/muxih4ck && cd $GOPATH/src/github.com/muxih4ck
git clone https://github.com/asynccnu/data_analysis_service_v1.git
cd Go-Web-Application-Template
make
./main
```

### Testing

```
make test
```

### change

此服务使用了go mod
- 修改了config里conf的路径
- 使用时修改config.yaml内的influxDB.addr, 并且按需修改config.go里的path
