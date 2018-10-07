# Lapin Cretin
This app allows you to create loads of ghost connections and channels in [RabbitMQ](http://www.rabbitmq.com/). This is very useful when performing load or chaos testing, or even preparing firedrills.

## Run locally
In order to run this the app locally, make sure you have a RabbitMQ running somewhere (e.g. localhost:5672). The app will get the RabbitMQ connection URI from teh [`VCAP_SERVICES` environment variable](https://docs.cloudfoundry.org/devguide/deploy-apps/environment-variable.html#VCAP-SERVICES).

```
$ export VCAP_SERVICES='{ "p-rabbitmq": [ { "credentials": { "protocols": { "amqp": { "uri": "amqp://guest:guest@localhost:5672" } } } } ] }'
$ go install
$ lapin-cretin <# of connections> <# of channels per connection>
```

## Test
Requires a local RabbitMQ.

```
$ ginkgo -r
```

## Deploy
Using CF, create and bind the RabbitMQ service to the app, and then:
```
$ cf push
```
