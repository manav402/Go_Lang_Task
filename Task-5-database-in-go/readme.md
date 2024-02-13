# PROBLEM :-
- Use a registration form to gather user data and populate your postgrace database from that data
- redirect user to a different page and show each stored user from your database

# STEPS :-
 1. setup each end point and datbase schema with proper validation
 2. create handler function for each route in go lang
 3. use struct, http library and template lib for playing with data 
 4. using postgrace database for performing above action

# How To Run :-
- change variable value stored in server.go file with your own postgrace  database value
- create a table named profile and create column mentioned in DB/query.go file with proper name and constraints.
- use belowe command to run the file
``` go run server.go ```
- after that open browser and navigate to this url
``` http://localhost:8000/ ```
- supported urls
``` 
GET
http://localhost:800/ 
http://localhost:800/allResult 

POST
http://localhost:800/register

```
