# Usage

```go
import (
    "github.com/Inasayang/logger"
)
```

## Copy and Truncate

```go
logger.Init("/var/log/{app}", "{app}", "info",nil)
```

```bash
cat << EOF > /etc/logrotate.d/{app}
/var/log/{app}/*.log {
    rotate 7
    daily
    compress
    copytruncate
}
EOF
```

## Rename and Create

```go
reloadCh := make(chan struct{}, 1)
logger.Init("/var/log/{app}", "{app}", "info",reloadCh)
...
capture user signal
...
```

```bash
cat << EOF > /etc/logrotate.d/{app}
/var/log/{app}/*.log {
    rotate 7
    daily
    compress
    create
    postrotate
        /usr/bin/kill -USR1 `pid of app`
    endscript
}
EOF
```

```bash
logrotate -f /etc/logrotate.d/{app}
```

- <https://grafana.com/docs/loki/latest/clients/promtail/logrotation/>
