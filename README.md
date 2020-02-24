# Hello Go

### Run locally
```shell script
$ go run main.go

  Print from the Go program
  Hello World
```

##### Run with environment
```shell script
$ export TEST_ENV=Phi
$ go run main.go

  Print from the Go program
  Hello Phi
```

### Build container image
```shell script
docker build . -t hello-go
```

### Run docker container
```shell script
docker run -it hello-go

  Print from the Go program
  Hello World
```

```shell script
docker run -it -e TEST_ENV=Phi hello-go

  Print from the Go program
  Hello Phi
```