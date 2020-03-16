## Test 
```bash
curl -H "Content-type: application/json" -X POST \
  -d '{"receiver": "team-sms", "status": "firing", "alerts": [{"status": "firing", "labels": {"alertname": "test-123"} }], "commonLabels": {"key": "value"}}' \
  http://localhost:8080/alert/teh-1/critical
```

## ENV 
```bash
KAVE_API_KEY=
RECIPIENT=
```

## AlertManager Configuration
Check [Documents](https://prometheus.io/docs/alerting/configuration/#webhook_config)
```yaml
# Whether or not to notify about resolved alerts.
[ send_resolved: <boolean> | default = true ]

# The endpoint to send HTTP POST requests to.
url: URL/alert/:region/:severity

# The HTTP client's configuration.
[ http_config: <http_config> | default = global.http_config ]
```