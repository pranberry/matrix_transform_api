# League Backend Challenge

In main.go you will find a basic web server written in GoLang. It accepts a single request _/echo_. Extend the webservice with the ability to perform the following operations

Given an uploaded csv file
```
1,2,3
4,5,6
7,8,9
```

1. Echo (given)
    - Return the matrix as a string in matrix format.
    
    ```
    // Expected output
    1,2,3
    4,5,6
    7,8,9
    ``` 
2. Invert
    - Return the matrix as a string in matrix format where the columns and rows are inverted
    ```
    // Expected output
    1,4,7
    2,5,8
    3,6,9
    ``` 
3. Flatten
    - Return the matrix as a 1 line string, with values separated by commas.
    ```
    // Expected output
    1,2,3,4,5,6,7,8,9
    ``` 
4. Sum
    - Return the sum of the integers in the matrix
    ```
    // Expected output
    45
    ``` 
5. Multiply
    - Return the product of the integers in the matrix
    ```
    // Expected output
    362880
    ``` 

The input file to these functions is a matrix, of any dimension where the number of rows are equal to the number of columns (square). Each value is an integer, and there is no header row. matrix.csv is example valid input.  

Run web server
```
go run .
```

Send request
```
curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"
```

## What we're looking for

- The solution runs
- The solution performs all cases correctly
- The code is easy to read
- The code is reasonably documented
- The code is tested
- The code is robust and handles invalid input and provides helpful error messages

## Initial thinking

- validate input file:
  - contains a matrix on upload. that is, number of rows == num columns
  - the matrix is a SQUARE matrix, note that for the arithemetic operations
  - all values are ints
- seems i must provide file with all REST calls.
- their `/echo` is using stdlib, will continue to use stdlib, no framework.
- Re: Testing:
  - the tests ~seem~ straight forward...
  - tricky parts seem to be the transformations
    - should test data be randomly generated or fixed?
    - how do you test the "flatness" of a flattened matrix? just that its a single slice?
    - transposing a matrix (switching rows and columns)
    - comparisons could be tricky
  - sure, start with the tests
- matrix transformation (rotation, etc) seems to be the most challenging part. matrix math should be pretty straightforward
- no requests types specified, use intelligent defaults
- add improved logging on connect
- move endpoint funcs out of main()
- seems that by default csv reader returns rows as strings. so conversion will be required
- every endpoint will need to open and read the file, move that into its own function
- while none of these requests will/should take very long (especially given the small test file), contexts should be used to enforce timeouts and cancellations