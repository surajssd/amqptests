#!/bin/bash

go build -o sender/sender sender/sender.go
docker build -t docker.io/surajd/rabbitmq-sender:1.0.1 ./sender

go build -o receiver/receiver receiver/receiver.go
docker build -t docker.io/surajd/rabbitmq-receiver:1.0.1 ./receiver

oc new-project amqptest1
kedge apply -f kedge.yml

oc get pods