<div align="center">
  <h1>Smyik - Banking application</h1>

  [![codecov](https://codecov.io/gh/1Shubham7/Banking-application/graph/badge.svg?token=X5WO4RO683)](https://codecov.io/gh/1Shubham7/Banking-application)
  [![GitHub last commit](https://img.shields.io/github/last-commit/1shubham7/banking-application)](#)
  ![GitHub language count](https://img.shields.io/github/languages/count/1shubham7/banking-application)
  ![GitHub top language](https://img.shields.io/github/languages/top/1shubham7/banking-application)
  ![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/1shubham7/banking-application/Test-and-Coverage)

</div>

This project is a banking application built using Go + PostgreSQL + Docker. I have used Go as the backend language, PostgreSQL as the database, SQLC for generating type-safe SQL queries, and `go-playground/validator` for input validation. The project also includes comprehensive unit tests written with `testify` and `GoMock` for mocking dependencies. The project also includes CI tests written using GitHub workflows and code coverage is also tracked and uploaded to Codecov ([find it here](https://app.codecov.io/gh/1shubham7/banking-application)). Build automation and commands are managed using Makefile.

⭐ Star us on GitHub — it motivates me a lot!

[![Share](https://img.shields.io/badge/share-000000?logo=x&logoColor=white)](https://x.com/intent/tweet?text=Check%20out%20this%20project%20on%20GitHub:%20https://github.com/1Shubham7/Banking-application%20%23OpenIDConnect%20%23Security%20%23Authentication)
[![Share](https://img.shields.io/badge/share-1877F2?logo=facebook&logoColor=white)](https://www.facebook.com/sharer/sharer.php?u=https://github.com/1Shubham7/Banking-application)
[![Share](https://img.shields.io/badge/share-0A66C2?logo=linkedin&logoColor=white)](https://www.linkedin.com/sharing/share-offsite/?url=https://github.com/1Shubham7/Banking-application)
[![Share](https://img.shields.io/badge/share-FF4500?logo=reddit&logoColor=white)](https://www.reddit.com/submit?title=Check%20out%20this%20project%20on%20GitHub:%20https://github.com/1Shubham7/Banking-application)
[![Share](https://img.shields.io/badge/share-0088CC?logo=telegram&logoColor=white)](https://t.me/share/url?url=https://github.com/1Shubham7/Banking-application&text=Check%20out%20this%20project%20on%20GitHub)

## 📚Tech Stack

- **Backend Language:** Go (Golang)
- **Database:** PostgreSQL
- **ORM/SQL Generator:** [SQLC](https://sqlc.dev/)
- **Validation:** go-playground/validator
- **Testing:** Comprehensive unit tests with GoMock for mocking
- **CI Pipeline:** GitHub Actions for CI tests and code coverage uploading, 
- **Code Coverage:** Tracking test coverage of the project using Codecov. ([find it here](https://app.codecov.io/gh/1shubham7/banking-application))
- **Build and Commands:** Managed using Makefile

![CodeCov](https://codecov.io/gh/1Shubham7/Banking-application/graphs/sunburst.svg?token=X5WO4RO683)

## 📝 How to Build

To build the packages, follow these steps:

```shell
# Open a terminal (Command Prompt or PowerShell for Windows, Terminal for macOS or Linux)

# Ensure Git and make is installed
# Visit https://git-scm.com to download and install console Git if not already installed

# Clone the repository
git clone https://github.com/1Shubham7/Banking-application.git

# Navigate to the project directory
cd Banking-application

make postgres

make createdb

make start_prod

# If you want to run it through local image:
# Install `go-migrate` with root previliges
sudo go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.1

make migrateup1

make migrateup

make server
```

## 🙋‍♂️ User Journey

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

## AWS Secrets

Command to generate the .env file from secrets stored in secrets manager and accessed using IAM (which has the required permissions).

```
aws secretsmanager get-secret-value --secret-id smyik-secret  --query SecretString --output text | jq -r 'to_entries|map("\(.key)=\(.value)")|.[]' > app.env
```

## 🤝 Contributing
  ![GitHub Issues or Pull Requests](https://img.shields.io/github/issues/1shubham7/banking-application)
  ![GitHub Issues or Pull Requests](https://img.shields.io/github/issues-closed/1shubham7/banking-application)
  ![GitHub Issues or Pull Requests](https://img.shields.io/github/issues-pr/1shubham7/banking-application) 
  ![GitHub Issues or Pull Requests](https://img.shields.io/github/issues-pr-closed/1shubham7/banking-application) 

Whether you have feedback on features, have encountered any bugs, or have suggestions for enhancements, we're eager to hear from you. Your insights help us make our application more robust and user-friendly. Please feel free to contribute by [submitting an issue](https://github.com/1Shubham7/Banking-application/issues). Each contribution helps us grow and improve. If you would like to contribute to this project, please read the [Contribution guidelines](CONTRIBUTING.md) and make sure to follow the [Code Of Conduct](CODE_OF_CONDUCT.md).

We appreciate your support and look forward to making our product even better with your help!

## 📃 License
![GitHub License](https://img.shields.io/github/license/1shubham7/banking-application)

This product is distributed under a Apache License Version 2.0, January 2004. You can review the full license agreement at [LICENSE](LICENSE)
