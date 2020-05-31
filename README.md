# RedisInGolang
Redis functionalities implemented in golang


In this Project, the following Redis operations are implemented in Golang using slice of integers:

A) List 
    1)Lpush
    2)Rpush
    3)Lpop
    4)Rpop
    5)Lrange


B) String
    1)Append
    2)SetNx
    3)INCR
    4)GetRange

C) Set
    1)SADD
    2)SCARD
    3)SREM
    4)SUNION    

Project implements the following :

1) Apis inplementation to perform operations
2) Modularity between the components (handler,controller,redisFile, etc...) and interface design
3) Error Handling
4) Unit Testing


Steps to execute:

1) Clone the repository from github link: https://github.com/viveklucky19/RedisInGolang
2) Go to the folder path and  Run command "go run main.go" in command line or terminal
3) Server will start running at port 8080 
4) Hit Get method with url "localhost:8080/lpush?data=32" in postman or browser and see the results. Similarly you can see perform other operations as well.


Postman Collect Link : https://www.getpostman.com/collections/e4cec9911c07fc2418a3