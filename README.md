# Producer-concumer Prblm2

go-concurrency Producer-concumer prblm2 // make piza app

The **Producer-Consumer problem** is a classic synchronization problem in computer science where two types of processes, producers and consumers, share a common resource, typically a buffer.

**Producer-Consumer problem** involves two types of goroutines:

- **Producer**: Generates data and sends it to a shared buffer (channel).
- **Consumer**: Receives and processes the data from the buffer.

Go's channels act as the shared buffer, synchronizing data flow between the producer and consumer, avoiding race conditions.

run
```sh
go mod init xprblm2
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
