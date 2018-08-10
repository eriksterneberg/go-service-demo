## go-user-service
Todo

## Quickstart
Build and run service locally:
`$ docker-compose up`

In another terminal, run:
```
$ curl localhost:8080/health
OK - pid 1
```

## Development
Todo

## Deployment
Todo

## Tech Stack
### Database
PostgreSQL
 - ACID

### Web Server
Golang's built in webserver. It doesn't service static files well, but it is fast and good enough for a pure API.

### Swagger
I chose the tool (Swagger)[https://swagger.io/] to generate a server stub to make setting up the endpoints faster. Swagger will provide also validate parameters and body in a request automatically.

Installation:
On MacOS:
`$ brew install swagger-codegen`

Generate or update server stub from swagger.yaml:
`$ swagger-codegen generate -i swagger.yaml -l go-server -o src/`


## Docker
Todo


## Todo
 -