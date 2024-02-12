# PROBLEM :-
    - Read data from different json file saved in json-data folder
    - Create a new json ouput file with contain same data as final-output.json
    - Refactor the format to reflect the output same as final-output.json

# APPROACH :-
    - Created a schema as struct data structure in go lang and mapped the data with the json file
    - used json library and marshaling functionality to achieve the same
    - than using os library we can perform io operation in clients memory

# STEPS :-
    1. create the formated struct for each json file wich will tackle the input
    2. create the formated struct for ouput json file with will be ouput for json file
    3. map the input with output using loops and transform the output structure 
    4. perform io operation and write the output.json file

# NOTES :-
    > The progrma may perform well if any one of data or any one of field are not available or 
    > not in valid format so we have to make sure the data from api are in valid formate and
    > in proper struct schema defined