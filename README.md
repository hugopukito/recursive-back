# Recursive structure

## Go & SQL

### Database

CREATE DATABASE IF NOT EXISTS recursive_db;

mysql -u root -p recursive_db < migration.sql

### Run

go run main.go

#### Html representation

-> http://localhost:8080

#### Json representation

-> http://localhost:8080/comments