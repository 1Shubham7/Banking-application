# Local User Journey

1. **Create a User/SignIn:**

```sh
curl -i -X POST http://localhost:8080/users -H "Content-Type: application/json" \
-d '{"username": "Shubham", "password": "secret", "full_name": "Shubham Singh", "email": "shubhammahar1306@gmail.com"}'
```
If running locally, you can checkout the db, we use postgres:

<img width="2784" height="1648" alt="image" src="https://github.com/user-attachments/assets/ac11f3c7-46f4-4d75-98cd-55202b3b097f" />


2. **Login User:**
```sh
curl -i -X POST http://localhost:8080/users/login -H "Content-Type application/json" \
-d '{"username": "Shubham", "password": "secret"}'
```

This will return you a JWT token:

```sh
{
  "session_id": "932680d9-a170-4e3d-8f1a-ed9aa090d21b",
  "access_token": "v2.local.nqUejIG_-A0-YKfedgDIjfFH9AyLr-idGE5e3QvduogV3hRsO0d1HTu3MoaXjHTfixDPjy7ZImbJZU_6S-LNZG3OqWtNZlQA0ta-UBg5vnnaGjzvahOT4CxwaBXOeZdxMjiZ0eRMB3yf4Kb8X0dYdjk0830h7b6b95gnmWfs85DfUUvh5JxuHaUwdT1SSiAf2r5-hh2d3qX1nFRnyeT0T0P3jOVrFd-Citqgz-ASCGMdtbJ1Dne6xjZ1SXsMlNt5zswiEkfTx5wOBshWV6oj.bnVsbA",
  "access_token_expires_at": "2024-12-13T12:15:43.050587204+05:30",
  "user": {
    "username": "Shubham",
    "full_name": "Shubham Singh",
    "email": "shubhammahar1306@gmail.com",
    "password_changed_at": "0001-01-01T05:53:28+05:53",
    "created_at": "2024-12-13T11:36:20.819497+05:30"
  },
  "refresh_token": "v2.local.aWZ1xqaQ6KZfgwmN-uPrwKgfu2uqzXILAA006JbGvIHdfmpAgWkj_vJ5bjPhUi8KDT1qGyzSYMTG1PD8tOTisGWVz__VYLKzjr7aPBkIB_Nj0yHBInYPXF27G1MK5E_DPi_GOYKiGjITW76CYN-ob7cM8p5ruCwxAB_Mz8Sr4XLlzGGSuRnuEaBVZkBd2rIQil2m5FjidkvyXtwztonC5wT93CxEeVCH7Nnx05zp_tKzUtAZV0vQU4xf_Z0S5ynPvrf0azYrHsyG_Jh2_t9s.bnVsbA",
  "refresh_token_expires_at": "2024-12-15T11:45:43.056058261+05:30"
}
```


3. **Create an account for the user:**

```sh
curl -X POST http://localhost:8080/accounts \
-H "Authorization: Bearer <Token> \
-H "Content-Type: application/json" \
-d '{"balance": 10000, "currency": "USD"}'
```

4. Now you can perform actions on the account, for example, get a user account (under authentication):

```sh
curl -X GET http://localhost:8080/accounts -H "Authorization: Bearer <Token>"
```

This will return a response:

```sh
{"id":1,"owner":"Shubham","balance":10000,"currency":"USD","created_at":"2024-12-13T12:01:09.869101+05:30"}
```

5. **Perform Money Tansfer:** 

For that we will have to create another user:

```sh
curl -i -X POST http://localhost:8080/users \
-H "Content-Type: application/json" \
-d '{
  "username": "Elon", 
  "password": "secret", 
  "full_name": "Elon Musk", 
  "email": "musk@Tesla.com"
}'
```

And then login the new user:

```sh
curl -i -X POST http://localhost:8080/users/login \
-H "Content-Type: application/json" \
-d '{
  "username": "Elon", 
  "password": "secret"
}'
```

This will return this result including the JWT/Paesto Token:

```sh
{
  "session_id": "0f48c550-16fc-418c-9e99-d5b378fcd53d",
  "access_token": "<Token>",
  "access_token_expires_at": "2024-12-13T13:23:37.434517608+05:30",
  "user": {
    "username": "Elon",
    "full_name": "Elon Musk",
    "email": "musk@Tesla.com",
    "password_changed_at": "0001-01-01T05:53:28+05:53",
    "created_at": "2024-12-13T12:52:38.271222+05:30"
  },
  "refresh_token": "<Token>",
  "refresh_token_expires_at": "2024-12-15T12:53:37.435101986+05:30"
}
```

Then create new account for this new user:

```sh
curl -X POST http://localhost:8080/accounts \
-H "Authorization: Bearer <Token>" \
-H "Content-Type: application/json" \
-d '{
  "balance": 100, 
  "currency": "USD"
}'
```

And then perform transfer:

```sh
curl -X POST http://localhost:8080/transfers \
-H "Authorization: Bearer <Token>" \
-d '{
  "from_account_id": 1, 
  "to_account_id": 2, 
  "amount": 100, 
  "currency": "USD"
}'
```

6. **For Refreshing Token:**

```sh
curl -X POST http://localhost:8080/tokens/renew_access \
-H "Authorization: Bearer <Token>" \
-d '{
  "refresh_token": "<Token>"
}'
```

This will return a new access token:

```sh
{
  "access_token": "<Token>",
  "access_token_expires_at": "2024-12-13T13:41:10.668632532+05:30"
}
```