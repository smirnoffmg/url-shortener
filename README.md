Url shortener. Pretty fast, as I can see
===

Benchmark with `ab`:

```bash
Document Path:          /r/f50fe5d6
Document Length:        41 bytes

Concurrency Level:      10
Time taken for tests:   0.777 seconds
Complete requests:      10000
Failed requests:        0
Non-2xx responses:      10000
Total transferred:      2310000 bytes
HTML transferred:       410000 bytes
Requests per second:    12877.04 [#/sec] (mean)
Time per request:       0.777 [ms] (mean)
Time per request:       0.078 [ms] (mean, across all concurrent requests)
Transfer rate:          2904.88 [Kbytes/sec] received
```

API
---

- `GET /r/{:alias}` - get 302 or 404

- `POST /api/v1/{:original_url}` - create short url

Example response:

```json
{
    "original": "https://google.com",
    "alias": "45a65bc8",
    "visits": 0
}
```
