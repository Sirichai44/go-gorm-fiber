CRUD with gorm and fiber
-
start with
  - 
  ```
    docker-compose up -d
    go run .
  ```   
postgres
  - 
  ```
    http://localhost:5050
  ```
fiber --> auth --> middleware
  -
  - register
```
  http://localhost:8080/register
  //method POST
    //example 
    {
      "email": "test@mail.com"
      "password": "1234"
    }
```
  - login
  ```
   //method POST
   http://localhost:8080/login
     //example 
    {
      "email": "test@mail.com"
      "password": "1234"
    }
  ```
  - create book
  ```
   //method POST
   http://localhost:8080/books
     //example 
    {
    "name": "Harry Potter",
    "author": "J.K. Rowling",
    "description": "The best book ever",
    "price": 100
   }
  ```
- update book
 ```
   //method PUT
   http://localhost:8080/books:id
     //example 
    {
    "name": "Harry Potter5",
    "author": "J.K. Rowling",
    "description": "The best book ever",
    "price": 100
   }
  ```
- delete book
```
   //method DELETE
   http://localhost:8080/books:id
```

