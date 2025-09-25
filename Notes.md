# Solution Notes (during and post):

- determining the right level of abstraction is very important
  - this will determine the shape of validation logic
    - every request requires a file be opened, matrix loaded, and matrix validate
    - if any of these operations fail, bail on the request
- I could easily operate on `[][]string`, converting values to int for the arithematic funcs(), but...
  - creating a `type Matrix` would enable more complex behaviour
  - provide type safety
  - better validation of a given matrix
  - better testing, all on int type...
  - HOWEVER, will not be able to use strings.Join(row, ",") for doing the printing....big hit imo
- watchout for scope creep:
  - the requirement is pretty straight-forward
- no concurrency here
- not enforcing HTTP Request types