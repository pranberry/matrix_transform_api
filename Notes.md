# Solution Notes (during and post)

## Challenges

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

## Assumptions

- stdlib only
- treating empty values (eg: `1,,2`) as legitimate
- Desired response content-type not specified. Sending back txt/csv, not JSON

## What missing

- any need for concurrency features
- not enforcing access by HTTP Request types. GET, POST, etc. all have the same behaviour
- logging middleware
  - current logging clutters the code, could easily be done via middleware
  - would be nice to just define and wrap
  - central place to define log format
