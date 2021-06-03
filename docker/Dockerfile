#### Use latest Ubuntu
FROM osgeo/gdal:latest

# Update base container install
RUN apt-get update
#RUN apt-get upgrade -y


# Install GDAL dependencies
RUN apt-get install -y  python3-pip
# This will install latest version of GDAL
RUN pip3 config set global.index-url https://mirrors.aliyun.com/pypi/simple/
COPY requirements.txt .
RUN  python3 --version
RUN pip3 install -r requirements.txt