# grpc math
- server\
`url` : localhost:4040

- client\
`url` : localhost:8080

## Add
```sh
curl -X GET http://localhost:8080/add/1/1
```

`2xx`
```json
{
  "success": true,
  "data": {
    "result": 2
  }
}
```
`4xx`
```json
{
  "success": false,
  "data": {
    "error_msg":{}
  }
}
```

## Multiply
```sh
curl -X GET http://localhost:8080/mul/2/2
```

`2xx`
```json
{
  "success": true,
  "data": {
    "result": 4
  }
}
```
`4xx`
```json
{
  "success": false,
  "data": {
    "error_msg":{}
  }
}
```
### Divide
```sh
curl -X GET http://localhost:8080/div/20/4
```

`2xx`
```json
{
  "success": true,
  "data": {
    "result": 5
  }
}
```
`4xx`
```json
{
  "success": false,
  "data": {
    "error_msg":{}
  }
}
```
### Reduce
```sh
curl -X GET http://localhost:8080/red/100/1
```

`2xx`
```json
{
  "success": true,
  "data": {
    "result": 99
  }
}
```
`4xx`
```json
{
  "success": false,
  "data": {
    "error_msg":{}
  }
}
```

# Running End To End test
- turn on grpc server with command :
```sh
make server.on
```
- don't forget also turn on client :
```sh
make client.on
```
- running test
```sh
go test -v -run=. github.com/fanchann/grpc_math/tests
```