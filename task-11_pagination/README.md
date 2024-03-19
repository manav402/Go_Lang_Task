# TASK :
1. Implement pagination and searching using mongodb database and sql dataset.
2. Create a command flag on flag present read the data from csv file and populate the mongo database

# APIS :-
| API     | METHOD |
|---------|---------|
| /getall | GET    |
| /get/{v}| GET    |
| /search | POST   |

## Feature:
- implemented lazy searching so after user completed typing only a single request will be send
- if user remove the search box text the previous data will be re rendered

