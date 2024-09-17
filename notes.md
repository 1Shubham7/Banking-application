```
docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e  POSTGRES_PASSWORD=secret -d postgres:12-alpine
```

```
docker exec -it postgres12 /bin/sh
```

We used this to generate mocks of Store interface

```
mockgen --destination db/mock/store.go github.com/1shubham7/bank/db/sqlc Store
```

you can also choose which package you want for your generated file:

```
mockgen --destination db/mock/store.go --package mockdb  github.com/1shubham7/bank/db/sqlc Store
```