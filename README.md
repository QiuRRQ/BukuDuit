### Bukuduit backend

### Building

Building requires a
[working Go 1.13+ installation](http://golang.org/doc/install).

Main package
```
$ git clone https://github.com/kudajawa/bukuduit-go.git
$ go mod download
$ go mod vendor
$ cd server
$ go run main.go
```

AMQP Listener
```

$ cd amqp_listener_otp
$ go run main.go

```

Docker Building
```
$ docker build -t [image_name] . --no-cache
$ docker-compose up -d --build
```

# Repository structure
```
amqp_listener_otp = listener amqp for sending otp
db = contains files about database
├──migrations = Contains migration script for database
├──models = Contains struct files that represent entities from a table
├──repositories = repositories package
├──── actions = Contains implementation query function of interface function from contracts directory
├──── contracts = Contains interface query function
helpers = Helper function and usage of pakage that usually called in usecase
key = Credential file e.g. azure key, google key (note : for security reason this directory excluded from git hub repository)
server = Main package
├── bootstrap = Init cron job and routes
├── handler = Handler function to validate parameter inputed and handle response body
├── middleware = Route middleware
├── request = Request body struct
usecase = API logic flow
├── viewmodel = Struct of usecase response body
dbconfig.yml = configuration for migration function
go.mod = list of package dependencies
```