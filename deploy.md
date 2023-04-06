## 线上
```shell
ssh starwiz@82.1557.180.248
cd /home/starwiz/app/pacific-insurance-api
docker-compose restar app
```

## 42
```shell
ssh starwiz@192.168.3.42
cd /home/starwiz/app/pacific-insurance-api
docker-compose restar app
```

遥感影像处理目前数据量少，目前手动处理，步骤是用gdal_translate转成mbtiles格式
```shell
gdal_translate demo.tif demo.mbtiles -of MBTILES
gdaladdo -r average demo.mbtiles 2 4 8 16 32 64 128 256
mv demo.mbtiles /home/starwiz/app/pacific-insurance-api/tilesets
# 手动把链接写入数据库，样式参照数据库原有数据 /demo字段加入raster表的mbtiles字段
# 项目tif上传通过上传tif接口上传，把项目tif转成cog格式（转方法见 https://gdal.org/drivers/raster/cog.html）在上传，
```