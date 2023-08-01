cat << EOF > /etc/logrotate.d/app_copytruncate
/var/log/app/*.log {
    rotate 7
    daily
    compress
    copytruncate
}
EOF