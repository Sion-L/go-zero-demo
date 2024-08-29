#!/bin/bash
mkdir /root/redis/data -pv
mkdir /usr/local/mysql/data -pv
mkdir /usr/local/mysql/logs -pv

docker run -d -p 6379:6379 -v /root/redis/data:/data --name redis redis redis-server --appendonly yes --requirepass "123456"
docker run -d --name etcd-server \
    --publish 2379:2379 \
    --publish 2380:2380 \
    --env ALLOW_NONE_AUTHENTICATION=yes \
    --env ETCD_ADVERTISE_CLIENT_URLS=http://etcd-server:2379 \
    bitnami/etcd:latest
docker run --name mysql8 -e MYSQL_ROOT_PASSWORD=123456 -v /usr/local/mysql/logs:/logs -v /usr/local/mysql/data:/var/lib/mysql -p 3306:3306 -d  mysql:8