FROM mysql:latest

ADD ./my.cnf /etc/mysql/conf.d/my.cnf
ADD ./hosts /etc/hosts

RUN ln -sf  /usr/share/zoneinfo/Asia/Tokyo /etc/localtime &&\
    chmod 644 /etc/mysql/conf.d/my.cnf
EXPOSE 3306
CMD [ "mysqld"]