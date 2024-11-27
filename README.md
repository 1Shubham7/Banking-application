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

https://codecov.io/gh/1Shubham7/Banking-application/graphs/sunburst.svg?token=X5WO4RO683



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
