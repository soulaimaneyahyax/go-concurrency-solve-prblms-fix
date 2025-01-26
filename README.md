# The Dining Philosophers

<img src="./imgs/1.png" alt="x" />

- A classic computer science problem
- Five philosophers live in a house together, and they always dine together at the same table, sitting in the same place.
- They always eat a special kind of spaghetti which requires two forks
- There are two forks next to each plate, which means that no two neighbours be eating at the same time

run
```sh
go mod init xprblm1
go run .
```

build
```sh
go build main.go && ./main
```

test
```sh
go test .
```

test race condition
```sh
go run -race .
go test -race .
```
