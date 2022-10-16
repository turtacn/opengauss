# 验证GORM PostgreSQL Driver适配OpenGauss

## 准备环境

### 启一个opengauss db实例
```shell script
docker run --name opengauss --privileged=true -d -e GS_PASSWORD=Enmo@123 -p 15432:5432 enmotech/opengauss:3.0.0 
```

## 测试代码
```shell script
go mod tidy
go run gstest.go
```
