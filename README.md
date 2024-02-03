## Install

1. Install `go`:

```bash
https://golang.org/dl/
```

2. Clone this repository:

```bash
git clone https://github.com/TimeReapz/national-itmx.git
```

3. Run script for create database:

```bash
sh script.sh
```

4. Run `main.go`:

```bash
go run main.go
```

## Example curl command

### Create a new customer

```bash
curl -X POST -H "Content-Type: application/json" -d '{"name": "Robert Downey Jr", "age": 40}' http://localhost:8080/customers
```

### Get a customer

```bash
curl -X GET http://localhost:8080/customers/1
```

### Update a customer

```bash
curl -X PUT -H "Content-Type: application/json" -d '{"name": "Robert Downey Jr", "age": 40}' http://localhost:8080/customers/1
```

### Delete a customer

```bash
curl -X DELETE http://localhost:8080/customers/1
```
