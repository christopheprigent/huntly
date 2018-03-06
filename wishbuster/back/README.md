Huntly Server
===

# Setup

First setup and start docker locally, then execute the `dbstart.sh` script.

# Usage

```
./back start | init | clean
```

For custom database deployment (outside the scope of the `dbstart.sh` script):
```
./back start --cockroach-user=root --cockroach-db=huntly --cockroach-host=localhost --cockroach-port=26257
```

# Testing

### Create Member and get token

```
curl -H "Content-Type: application/json" -X GET -d '{"id":"42"}' http://localhost:5050/token
```
