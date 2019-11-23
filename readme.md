# Introduction

This is my solution for the assignment for Salary.

# Design

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
