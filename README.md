## xmutex

<img src="./imgs/1.png" alt="x" />

run
```sh
go mod init xmutex
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

