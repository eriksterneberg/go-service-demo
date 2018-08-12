## go-user-service
A simple service for storing usernames, email addresses and preferences. No security review has been done, so it is not
production ready.

Properties:
 - ("Data on the outside")[https://www.confluent.io/blog/data-dichotomy-rethinking-the-way-we-treat-data-and-services/], enabling us to treat a stateful database as cattle and not as a pet

## Quickstart
Build and run service locally:
`$ docker-compose up`

In another terminal, run:
```
$ curl localhost:8080/health
OK
```

## Development
### Scaling
Run `docker-compose up --scale user-service=3` to run more than one Go service.
You can test that it works by entering one of the containers:
```
$ docker container exec -it go-user-service_user-service_1 /bin/sh
/ # wget user-service/health
Connecting to user-service (172.25.0.5:80)
health               100% |**********|    11   0:00:00 ETA
/ # cat health
OK
```

## Deployment
Todo


## Messages
Todo
 - Need to create topic inside the container for docker-compose.yml (dev-only)

## Security
- User credentials is protected using SSL encryption
- Only store hashes of user credentials (password) in database


## Tech Stack
### Database
The optimal setup would be to use a Postgres database cluster due to its ACID compliance.
Setting this up takes a lot of manual work however, and so for this simple service I will choose MongoDB instead.

Todo:
 - Set password using Secrets

### Web Server
Golang's built in webserver. It doesn't service static files well, but it is fast and good enough for a pure API.

### Swagger
I choose the tool (Swagger)[https://swagger.io/] to generate a server stub to make setting up the endpoints faster. Swagger will provide also validate parameters and body in a request automatically.

Installation:
On MacOS:
`$ brew install swagger-codegen`

Generate or update server stub from swagger.yaml:
`$ swagger-codegen generate -i swagger.yaml -l go-server -o src/`

Todo:
- Add `swagger-codegen` to the Makefile.


## Docker
I choose to containerize the microservice using Docker since it will make it easier to develop, test and deploy. I choose Docker Swarm over Kubernetes for deployment since it is much easier to setup.


## Todo
 -