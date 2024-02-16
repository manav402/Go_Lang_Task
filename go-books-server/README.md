## Problem :-
- For given struct create a data table and perform CRUD operations over it.
- Use mux and gorm for routing and data operations.
- Create moduler code with each service implemented.

## Solution :-
1. First define each router using mux.Router which has params for handlerfunction.
2. Than connect to database with gorm driver and create struct for each table models.
3. Perform auto migrate so gorm can create required tables.
4. Create services with db pointer and write api functionalities.

## Available Apis :-
#### GET
```
    http://localhost:8000/Books         -- returns all books
    localhost:8000/Book/{id}            -- return a single book matching id
```
#### POST
```
    localhost:8000/Book                 -- creates a book
```
#### PATCH
```
    localhost:8000/Book/{id}             -- Edits a book matching id
```

# Table Schemas :-

This document describes the schema of the `Book` table.

| Column    | Type     | Constraints   |
| --------- | -------- | ------------- |
| title     | VARCHAR  | None          |
| author    | VARCHAR  | None          |
| isbn      | VARCHAR  | PRIMARY KEY   |
| publisher | VARCHAR  | None          |
| year      | INTEGER  | None          |
| genre     | VARCHAR  | None          |

# Running The Server :-
first create a env file with following details name must and only be 
> ".env"

write the following contant with your details.
```
HOST = "YOUR HOST NAME"
PORT = "YOUR PORT NUMBER"
USER = "YOUR POSTGRAS USER NAME"
PASSWORD = "YOUR PASSWORD"
DBNAME = "YOUR DATABASE NAME"
```

use the following command to start the server
> go run main.go


### NOTE :-

- grom is and object relational model for go lang.
- which can efficiently write query with minimal overhead for developer 
- grom provide methods to developer so they don't have to write sql queries.
- mux is a powerfull library for go which can override default router of http library.
- providing developer a more control over api parsing like query and parameter finding
- filtering based on state and host name etc.


