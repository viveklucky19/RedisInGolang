# RedisInGolang
Redis functionalities implemented in golang


In this Project, the following Redis operations are implemented in Golang using slice of integers:

1)Lpush
2)Rpush
3)Lpop
4)Rpop

Project implements following :

1) Apis inplementation to perform operations
2) Modularity between the components (handler,controller,redisFile, etc...) and interface design
3) Error Handling
4) Unit Testing


Steps to execute:

1) Clone the repository from github link: https://github.com/viveklucky19/RedisInGolang
2) Go to the folder path and  Run command "go run main.go" in command line or terminal
3) Server will start running at port 8080 
4) Hit Get method with url "localhost:8080/lpush?data=32" in postman or browser and see the results. Similarly you can see perform other operations as well.