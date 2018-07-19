# Lapin Cretin

## Run locally
In order to run this the app locally, make sure you have a RabbitMQ running somewhere (e.g. localhost:5672).

```
$ export VCAP_SERVICES="{ "p-rabbitmq": [ { "credentials": { "protocols": { "amqp": { "uri": "amqp://guest:guest@localhost:5672" } } } } ] }"
$ go run main.go
```

## Test
Requires a local RabbitMQ.

```
$ ginkgo -r
```

## Deploy
Using CF, create and bind the RabbitMQ service to the app, and then:
```
$ cf push -i 4
```
