# pagerduty-api docker sample

## Test

``` shell
$ cd path/to/sample
$ docker-compose up -d

# test
$ curl -d '[{"labels": {"Alertname": "PagerDuty Test"}}]' http://localhost:9093/api/v1/alerts

# pagerduty-api logs

[GIN-debug] POST   /v2/enqueue               --> main.main.func3 (3 handlers)
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default

...

{"routing_key":"dummykey_length_is_32_characters","dedup_key":"<dedup_key>","event_action":"trigger","payload":{"summary":"[FIRING:1] PagerDuty Test ","source":"AlertManager","severity":"error","custom_details":{"firing":"Labels:\n - alertname = PagerDuty Test\nAnnotations:\nSource: \n","num_firing":"1","num_resolved":"0","resolved":""}},"client":"AlertManager","client_url":"http://alertmanager:9093/#/alerts?receiver=team-X-pager"}

[GIN] 2019/03/22 - 11:00:49 | 202 |      1.3033ms |      172.25.0.4 | POST     /v2/enqueue

...

{"routing_key":"dummykey_length_is_32_characters","dedup_key":"<dedup_key>","event_action":"resolve","payload":{"summary":"[RESOLVED] PagerDuty Test ","source":"AlertManager","severity":"error","custom_details":{"firing":"","num_firing":"0","num_resolved":"1","resolved":"Labels:\n - alertname = PagerDuty Test\nAnnotations:\nSource: \n"}},"client":"AlertManager","client_url":"http://alertmanager:9093/#/alerts?receiver=team-X-pager"}

[GIN] 2019/03/22 - 11:15:31 | 202 |       207.8µs |      172.25.0.4 | POST     /v2/enqueue
```
