# Usage:
```go
import (
    "github.com/Inasayang/logger"
)
```
```go
logger.Init("/var/log/{serviceName}", "{serviceName}", "info")
```

```
cat << EOF > /etc/logrotate.d/{serviceName}
/var/log/{serviceName}/*.log {
    rotate 7
    daily
    compress
    copytruncate
}
EOF
```

```
logrotate -f /etc/logrotate.d/{serviceName}
```