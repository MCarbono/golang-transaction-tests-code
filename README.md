<h1 align="center"> Financial transaction registration </h1>

---

## About

This is a simple programm to study about tests, folder struct and GOLANG sintax. All the code is based on a live of a youtube channel called "Full Cycle".

---

## Libs

- SQLite3
- gomock
- testify
- migrate

--- 

## Mocks Setup

The mocks are created based on an interface. If there's by any chance an update on it, runs the command below to adds new mocks 
for the methods.

```bash
    $ mockgen -source=entity/repository.go -destination=entity/mock/mock.go
```

## 🗄 Database Setup

```bash
    #executes database sqlite3
    $ sqlite3 test.db

    #Copy and paste the code below
    $ CREATE TABLE transactions(
        id TEXT NOT NULL,
        account_id TEXT NOT NULL,
        amount REAL NOT NULL,
        status TEXT NOT NULL,
        error_message TEXT NOT NULL,
        created_at TEXT NOT NULL,
        updated_at TEXT NOT NULL 
    );
```

## Run

```bash
    #Run with Makefile
    $ make

   #Run main.go file
   $ go run cmd/main.go
```

## 🧪 Run Tests

```bash
    #run tests with Makefile
    $ make tests

    #run all tests
    $ go test ./... -v
```

