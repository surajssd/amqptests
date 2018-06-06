#!/bin/bash

eval $(minishift docker-env)

go build -o sender/sender sender/sender.go
docker build -t docker.io/surajd/activemq-sender:1.0.1 ./sender

go build -o receiver/receiver receiver/receiver.go
docker build -t docker.io/surajd/activemq-receiver:1.0.1 ./receiver

oc new-project amqptest4
kedge apply -f kedge.yml

sleep 20

kedge apply -f kedge-receiver2.yml

oc get pods
