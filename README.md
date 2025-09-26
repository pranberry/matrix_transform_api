# League Backend Challenge

- API works as intended
- All tests pass

## How to run:
Run web server
```
go run .
```
Run tests:
```
go test ./...
```
Send request
- use postman or curl
```
curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"
```

## Solution Notes

### Assumptions

- stdlib only
- A valid matrix is `[][]string`
- treating empty values (eg: `1,,2`) as legitimate cell values
- All valid matrices can be transposed and flattened. But only int-value matrices can be added or multiplied
- Desired response content-type not specified. Sending back txt/csv, not JSON

### What missing

- any need for concurrency features
- not enforcing access by HTTP Request types. GET, POST, etc. all have the same behaviour
- logging middleware
  - current logging clutters the code, could easily be done via middleware
  - would be nice to just define and wrap
  - central place to define log format
- Add support for generics
- Using contexts to enforce request deadlines

### Challenges

- determining the right level of abstraction is very important
  - this will determine the shape of validation logic
- every request requires a file be opened, matrix loaded, and matrix validate
  - if any of these operations fail, bail on the request

- I could easily operate on `[][]string`, converting values to int for the arithmetic funcs(), but...
  - creating a `type Matrix` would enable more complex behaviour
  - provide type safety
  - better validation of a given matrix
  - better testing, all on int type...
  - HOWEVER, will not be able to use strings.Join(row, ",") for doing the printing....big hit imo

- watchout for scope creep:
  - the requirement is pretty straight-forward

## Task Description

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

## What we're looking for

- The solution runs
- The solution performs all cases correctly
- The code is easy to read
- The code is reasonably documented
- The code is tested
- The code is robust and handles invalid input and provides helpful error messages

