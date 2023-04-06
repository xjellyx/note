
- server 服务基本配置

    - staticurl: http://127.0.0.1:8101/static 修改成当前系统的ip地址或者域名
    - http
      - host: 0.0.0.0 无需修改
      - port: 8100 无需修改
      - addr: 0.0.0.0:8101 无需修改d
      - timeout: 1s 无需修改

- database 数据库连接配置
  - driver: postgres 默认是postgres，若改为其他数据库，需要修改
  - automigrate：false 设置成true会根据定义的结构体自动生成或同步数据表结构信息
  - source：postgres://postgres:starwiz123@127.0.0.1:54399/gis?sslmode=disable 数据库连接链接

- ml 内部机器学习服务
    - addr: http://127.0.0.1:8301 如使用docker脚本部署，无需修改
    - ship: /ship 船舶识别路由，无需修改,内部请求是会是http://ml:8300/ship,下同
    - airplane： /airplane 飞机识别
    - many: /many 飞机，船舶，汽车等多个识别的模型
    - wrap: /target-extract 无需修改，内部逻辑使用,目标提取
    - tiff2png: /tiff2png 内部遥感数据转换使用
    - tiffinfo: /tiffinfo 内部遥感数据读取信息使用
    - changedetect: /change-detect 变化监测模型识别
    - features: /features 地物分类模型识别
    - gdal2tiles: /gdal2tiles 内部处理遥感影像切片
- file 静态文件保存路径，务必有/home/starwiz/data目录
  - basedir: /home/starwiz/data/local-data-2t 项目文件保存地址,local-data-2t是目前系统挂载的2t固态硬盘
  - thumb: /home/starwiz/data/local-data-2t/analysis_information_data/thumb 拇指图输入地址，前缀必须和basedir一致
  - airplane: /home/starwiz/data/local-data-2t/analysis_information_data/thumb  飞机识别输出地址，前缀必须和basedir一致
  - tiles: /home/starwiz/ghk-data-api/remote_sensing_information_analysis/static/tiles 切片保存的路径
  - tiff: /home/starwiz/data/local-data-2t/analysis_information_data/tiff tif保存路径，前缀必须和basedir一致

- environment: test 日志数据控制

- serviceid: 888 服务id
- servicename: "information_dig" 服务名称
- nacos 网关服务
  - register: true true时才会注册网关
  - clientip: 127.0.0.1 注册客户端地址,本项目服务ip地址
  - clientport: 8101 注册客户端端口,本项目的端口号
  - host: 127.0.0.1 nacos网关ip地址
  - port: 8848 nacos网关端口号
- log
  - filename: ./log/logs.log 日志保存路径
  - maxsize: 20 日志文件最大值，单位M
  - maxbackups: 30 备份日志，天
  - maxages: 30 最长保存时间，天
  - compress: false true压缩

