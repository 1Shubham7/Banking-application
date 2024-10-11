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

