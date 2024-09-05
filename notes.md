```
docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e  POSTGRES_PASSWORD=secret -d postgres:12-alpine
```

```
docker exec -it postgres12 /bin/sh
```