## Installation

To use this project, you need to follow these steps:

1. Clone the repository: `git clone git@github.com:LeandroSAlmeida/golang-api.git`
2. create a database in mysql with the name: `golang_api`
3. Configure the MySQL connection in the `config/mysql.go` file, modifying the dsn variable as needed. The first "root" is the username and the second "root" is the password
4. Install the dependencies: `go mod download`
6. Run the application: `go run main.go`

## Used Tools

This project uses the following tools:

- [Golang](https://golang.org/) for backend development
- [Go-Gin](https://github.com/gin-gonic/gin) for route management
- [GoORM](https://gorm.io/) for database communication
- [MySQL](https://www.mysql.com/) as the database
- [Swagger](https://swagger.io/) for API documentation and testing
