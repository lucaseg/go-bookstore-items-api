# go-bookstore-items-api

This project use a Elastic search so, you have to run that using the docker compose file:

```
docker-compose up -d
```

Create the Index to store items

`PUT: 127.0.0.1:9200/items`

Body
```
{
    "settings": {
        "index": {
            "number_of_shards": 4,
            "number_of_replicas": 2
        }
    }
}
```