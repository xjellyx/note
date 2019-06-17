#!/bin/bash
docker run -d \
  --restart=always \
  --name = mysql \
  -p 33306:3306 \
  -e MYSQL_ROOT_PASSWORD=mysql \
  -e MYSQL_DATABASE=business \
  -e MYSQL_USER=business \
  -e MYSQL_PASSWORD=business \
  -e TZ=Asia/Shanghai \
  -d mysql:8.0 --default-authentication-plugin=mysql_native_password
# waiting for init
sleep 5
