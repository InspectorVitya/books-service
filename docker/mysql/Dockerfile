FROM mysql:8.0.27

MAINTAINER me

ENV MYSQL_DATABASE=books \
    MYSQL_ROOT_PASSWORD=root

ADD schema.sql /docker-entrypoint-initdb.d

EXPOSE 3306