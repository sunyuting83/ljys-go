# 影视后端API
---
#### gin + gorm + levledb 练手作品
示例： [m.zhenguohe.com](http://m.zhenguohe.com "m.zhenguohe.com") （[React前端开源地址](https://github.com/sunyuting83/dbys-react "React前端开源地址")）

项目地址：[github.com/sunyuting83/ljys-go](https://github.com/sunyuting83/ljys-go "github.com/sunyuting83/ljys-go") 

## 依赖
- go 版本是 go1.14.3 linux/amd64
- [gin](https://github.com/gin-gonic/gin "gin")
- [gorm](https://github.com/go-gorm/gorm "gorm")
- [goleveldb](https://github.com/syndtr/goleveldb "goleveldb")
- [redix](https://github.com/alash3al/redix "redix")  github.com/alash3al/redix/kvstore/leveldb

## 环境
- 开发环境 Manjaro
- 服务器环境 Ubuntu 16
- 打包环境 Centos 7.5

其他环境没有测试过，不确定是否兼容（如 Windows、MAC）

## 快速开始
不需要自己编译，从releases下载最新版本的movie.zip 已经编译好版本，然后执行：
```bash
unzip movie.zip
cd movie
./movie -p 8123
```
参数说明：-p 监听端口号

浏览器访问 [http://localhost:8123/api/index](http://localhost:8123/api/index "http://localhost:8123/api/index")

### 接口说明

|  接口地址  | 说明  | 参数说明 |
| ------------ | ------------ | ------------ |
| /api/index  |  首页 | null |
| /api/getclass  |  大类 | ?id= 大类id 1 - 4 |
| /api/list  |  小类 | ?cid= 小类id 5以上  &page= 页码 |
| /api/getmovie  |  播放页 | ?id= 电影id |
| /api/getkey  |  搜索联想关键词 | ?key= 关键词 |
| /api/search  |  搜索结果 | ?word= 关键词 &page= 页码 |
| /api/gethot  |  搜索页 | null |
| /api/area  |  地区的影片列表 | ?id= 地区id &page= 页码 |
| /api/director  |  导演的影片列表 | ?id= 导演id &page= 页码 |
| /api/performer  |  演员的影片列表 | ?id= 演员id &page= 页码 |
