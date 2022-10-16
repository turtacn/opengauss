# GORM PostgreSQL Driver适配OpenGauss

## 准备环境

### 启一个opengauss db实例
```shell script
# 起实例
docker run --name opengauss --privileged=true -d -e GS_PASSWORD=Enmo@123 -p 15432:5432 enmotech/opengauss:3.0.0 

# 创数据库
docker ps | grep opengauss # get container id
docker cp school.sql $container_id:/

docker exec -ti $container_id bash 
su - omm
/usr/local/opengauss/bin/gsql -f /school.sql  # create school db 

```

## 测试代码
```shell script
go mod tidy
go run gstest.go

```
