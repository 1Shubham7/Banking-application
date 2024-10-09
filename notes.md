```
docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e  POSTGRES_PASSWORD=secret -d postgres:12-alpine
```

We will need to download `golang-migrate/migrate` for migrations:

Step 1. Ensure $GOPATH/bin is in your PATH: The installed binaries are placed in $GOPATH/bin by default, so make sure it's in your PATH:
```
export PATH=$PATH:$(go env GOPATH)/bin
```

Then install it with root previliges:
```
sudo go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.1

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