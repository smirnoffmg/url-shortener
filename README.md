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

```json
{
    "alias": "45a65bc8",
}
```

- `GET /api/v1/stats/{:alias}` - get short url stats

Example response:

```json
{
  "visits": [
    {
      "ID": 2,
      "CreatedAt": "2024-05-15T17:08:40.735617Z",
      "UpdatedAt": "2024-05-15T17:08:40.735617Z",
      "DeletedAt": null,
      "alias": "0f31ea3a",
      "ip_addr": "192.168.65.1",
      "user_agent": "ApacheBench/2.3"
    }
  ]
}

```

ERD
---

![ERD](docs/images/docs/ERD.png)

Techical stack
---

- Go
- Gin
- PostgreSQL
- Docker
- Docker-compose

We use two databases: one for the url records and the other for the visit records. This is done to optimize the performance of the application.

How to run
---

```bash
docker-compose up --build
```

Environment variables
---

- `URLS_DB_DSN` - DSN for the urls database
- `VISITS_DB_DSN` - DSN for the visits database

Example could be found in the `.env` file.
