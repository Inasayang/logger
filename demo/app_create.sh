cat << EOF > /etc/logrotate.d/app_create
/var/log/app/*.log {
    rotate 7
    daily
    compress
    create
    postrotate
        /usr/bin/kill -USR1 `pid of app`
    endscript
}
EOF