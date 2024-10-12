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





Create User First:

```
curl -i -X POST http://localhost:8080/users -H "Content-Type: application/json" -d '{"username": "shiv", "password": "password", "full_name": "Shiv Pratap Singh Waghel", "email": "shivp007@gmail.com"}'
```

Then Login the User:

```
curl -i -X POST http://localhost:8080/users/login -H "Content-Type application/json" -d '{"username": "shiv", "password": "password"}'
```

Then create an account for the user:

```
curl -X POST http://localhost:8080/accounts \
-H "Authorization: Bearer <PAESTO_TOKEN>" \
-H "Content-Type: application/json" \
-d '{"balance": 10000, "currency": "USD"}'
```

Then you can get that account:

```
curl -X GET http://localhost:8080/accounts/1 -H "Authorization: Bearer <TOKEN>"
```







Building Docker image:

```
docker build -t smyik:latest .
```

```
docker run --name smyik -p 8080:8080 -e GIN_MODE=release smyik:latest
```

we are creating a docker network so that the smyik and postgres container can communicated within the same network, this could also be resolved my specifying the same ip address for smyik as it is for Postgres but when we restart a container the IP may change, so it is not a  best practice:

```
docker run --name smyik -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://roo
t:secret@172.17.0.2:5432/simple_bank?sslmode=disable" smyik:latest
```

So we will create our own network and put smyik and postgres containers in this network:

```
docker network create smyik-network
```

```
docker network connect smyik-network postgres12
```

(later we will make this change in Makefile itself so that it happens automatically)

```
docker run --name smyik --network smyik-network -p  8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secret@postgres12:5432/simple_bank?sslmode=disable" smyik:latest
```

