## Data analysis

![](https://travis-ci.org/asynccnu/data_analysis_service_v1.svg?branch=master)

### 简介

参考自掘金小册[基于 Go 语言构建企业级的 RESTful API 服务](https://juejin.im/book/5b0778756fb9a07aa632301e)

主要依赖：gin + gorm + viper + lexkong/log

### Build and run

```
mkdir $HOME/asynccnu && cd $HOME/asynccnu
git clone https://github.com/asynccnu/data_analysis_service_v1.git data_analysis
cd data_analysis
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
