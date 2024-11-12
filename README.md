<div align="center">
  <h1>Smyik - Banking application</h1>

  [![codecov](https://codecov.io/gh/1Shubham7/Banking-application/graph/badge.svg?token=X5WO4RO683)](https://codecov.io/gh/1Shubham7/Banking-application)
  [![GitHub last commit](https://img.shields.io/github/last-commit/1shubham7/banking-application)](#)
  ![GitHub language count](https://img.shields.io/github/languages/count/1shubham7/banking-application)
  ![GitHub top language](https://img.shields.io/github/languages/top/1shubham7/banking-application)

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

# Ensure Git is installed
# Visit https://git-scm.com to download and install console Git if not already installed

# Clone the repository
git clone https://github.com/1Shubham7/Banking-application.git

# Navigate to the project directory
cd Banking-application

make postgres

make createdb

# Install `go-migrate` with root previliges
sudo go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.1

make migrateup1

make migrateup

make server
```