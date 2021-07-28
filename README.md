# Echo tcp server
## How to run?
The first and the simplest solution is:
```
cd server
go run server.go
```

or you can run the server using docker-compose(remove the -d flag to see logs)
```
docker-compose up server
```

to run the client(remember to specify the port flag):
```
cd client
go run client.go --port=7777
```