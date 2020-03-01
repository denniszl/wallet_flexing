# Flexing on nerds

This is an api that has no purpose except to flex on people.

# Building
Run `make build` to build. Requires go 1.11 or later to use `mod` to install dependencies.

# Running
Run using `./bin/api`

The default port the http server runs on is `6969`. To change the port, set a `PORT` environment variable (e.g. `PORT=7000 ./bin/api`)

# Tests
Run tests using `make build`

# API
## curl to do a GET for some time period:

`curl '0.0.0.0:6969/transactions?start_date_time=2020-03-01T22%3A43%3A26%2B00%3A00&end_date_time=2020-03-01T22%3A44%3A26%2B00%3A00'`

## curl to do a GET for all transactions:
`curl 0.0.0.0:6969/transactions`

## curl to POST a transaction:
`curl -v -XPOST '0.0.0.0:6969/transactions' -d '{"datetime": "2019-10-05T13:00:00+00:00", "amount": 1000}'`

# Why are amounts sometimes higher in the past?
Since transactions are listed a descending order by timestamp and not by amounts, if an amount is reported late, the total will be adjusted for when it actually gets reported.

If it's a requested feature, perhaps it'd make sense to add things like a posted date or sorting by amount.

# Why is the GET time period range for query parameters and not in the body?
[GET is semantically supposed to ignore the body.](https://tools.ietf.org/html/rfc2616#section-4.3) Perhaps this could have been implemented as a POST, but that wouldn't have been RESTful.