### How to build
1. `docker compose up -d`
2. `docker compose exec go sh`
3. `go run api/api.go`

if it doesn't work...
set go.mod file and `go mod tidy` on shell

### How to use api
Push Send Request button in `api/requests.http`

### How to connect Sequel Ace
Name: {any}  
Host: localhost  
Username: root  
Password: ginapi  
Database: ginapi  
Port: 3306  