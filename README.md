## Crud API written in Go and Gin
![example workflow](https://github.com/almcr/crud-go/actions/workflows/main.yml/badge.svg)
---
This is my take on using Go environment to create a simple crud application. Go has really low entry barrier, and has a huge community which makes it simple to find learning resources. 

### Technologies
* Go 1.17
* Gin web framwork
* go-jwt
* go driver for mongodb
* dotenv

### Main challenges
Depsite the ease of use of the language, some time have been taken to comprehend the package management system (pkg vs modules) and some language constructs (concurrency, data structures, value semantics...). 
It's also my first exposition to mongoDB, its such a great technology with great library support.
I also dabbled with jwt, the topic of security is so fascinating 

### Todo
- [ ]  tests
- [ ]  better error handling and recovery
- [ ]  setup CI
- [ ]  use change stream for update route
### Usage
Run the script `init_db.sh`  to launch a mongo container and `go run` to launch the app


[![Run in Postman](https://run.pstmn.io/button.svg)](https://god.postman.co/run-collection/ceee14bc8aeb05f44dfc?action=collection%2Fimport)
