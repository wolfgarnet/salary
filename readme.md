# Introduction

This is my solution for the assignment for Salary.

It is a solution that uses an Sqlite database. It could easily have been upgraded to use a MySQL or PostgreSQL database.
Furthermore, docker could have been an excellent choice to containerize the solution, 
which would have had a docker image for the solution and a docker image for the database.
Everything could have been orchestrated by docker-compose and made this as a neatly packaged solution.

# Building and running

In order to build the solution, go to ```cmd/app``` and run the command ```go buid```.
Run the application by executing ```./app```, this will start the server on port 6543.

# Testing

#### Running the tests

From the root folder run the command ```go test``` to run all the tests.

#### About the tests

Since this is a very simple application that relies heavily on database operations and the http server,
only integration tests have been implemented.
All the endpoints are tested, though even more special cases could have been tested.

Unit tests and component tests are a vital part of a test suite, but in this given solution,
there aren't much business logic to test and therefore, not feasible to test.

# Usage

#### Creating an account

Using the POST request http://localhost:6543/accounts

```bash
curl -XPOST -H 'Content-type: application/json' http://localhost:6543/accounts -d '{"name": "My account"}'
```

#### Listing transaction for specific account

Using the GET request http://localhost:6543/accounts/{id}/transactions

```bash
curl http://localhost:6543/accounts/2/transactions
```

#### Listing all accounts

Using the GET request http://localhost:6543/accounts

```bash
curl http://localhost:6543/accounts
```

#### Creating a transaction for an account

Using the POST request http://localhost:6543/accounts/{id}/transactions

```bash
curl -XPOST -H 'Content-type: application/json' http://localhost:6543/accounts/2/transactions -d '{"description": "Description", "amounet": 100}'
```

#### Deleting a transaction

Using the DELETE request http://localhost:6543/transactions/{id}

```bash
url -XDELETE http://localhost:6543/transactions/1
```
