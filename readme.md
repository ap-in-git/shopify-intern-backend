#About

Project made for shopify technical challenge.

#Running project locally

* Clone the repository and open terminal in project folder
* run `go get ./`
* run cp `.env.example .env`
* put `SESSION_COOKIE_KEY` value in `.env` file
* run `go build -o gin main.go` to build the application
* run `./gin` in terminal and open localhost:8080 to see the application

#Used external packages
1. [https://github.com/go-gorm/gorm](https://github.com/go-gorm/gorm)
2. [https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
3. [https://github.com/joho/godotenv](https://github.com/joho/godotenv)
