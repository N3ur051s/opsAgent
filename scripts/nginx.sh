#!/bin/sh

# Source function library.
. /etc/rc.d/init.d/functions

# Source networking configuration.
. /etc/sysconfig/network

# Check env variable
. /etc/profile

# Check that networking is up.
[ "$NETWORKING" = "no" ] && exit 0

nginx="/usr/local/openresty/nginx/sbin/nginx"
prog=$(basename $nginx)
counter=$(sudo ps -C nginx --no-heading|wc -l)

NGINX_CONF_FILE="/usr/local/openresty/nginx/conf/nginx.conf"

start() {
    [ -x $nginx ] || exit 5
    [ -f $NGINX_CONF_FILE ] || exit 6
    echo -n $"Starting $prog: "
    daemon $nginx -c $NGINX_CONF_FILE
    [ $? -ne 0 ] && echo "负载均衡启动失败"
}

stop() {
    echo -n $"Stopping $prog: "
    ps -A | grep nginx | grep -v grep | awk '{ print $1}' | xargs -L 1 kill -QUIT
    retval=$?
    return $retval
}

status() {
    sudo ps -ef|grep nginx | grep -v 'grep' | grep -v 'nginx.sh'
    retval=$?
    [ $retval -eq 0 ] && echo "当前负载均衡正常运行" && exit 0
}

restart() {
    stop
    start
}

reload() {
    sudo $nginx -s reload
    retval=$?
    [ $retval -eq 0 ] && echo "配置重载成功\n"
}

configtest() {
    sudo $nginx -t -c $NGINX_CONF_FILE
    if [ $? -ne 0 ]; then
        echo "负载均衡配置检测失败" && exit 6
    fi 
}


configtest
if [ $counter -eq 0 ]; then
    start
    status
fi
if [ $counter -gt 0 ]; then
    reload 
    status
fi


# case $1 in
# start)
# start
#     ;;
# status)
# status
#     ;;
# stop)
# stop
#     ;;
# restart)
# restart
#     ;;
# reload)
# reload
#     ;;
# test)
# configtest
#     ;;
# *)
# echo    "使用格式:$0 start|stop|status|restart|reload|test"
# esac