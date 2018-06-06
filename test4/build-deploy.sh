#!/bin/bash

eval $(minishift docker-env)

go build -o sender/sender sender/sender.go
docker build -t docker.io/surajd/activemq-sender:1.0.1 ./sender

go build -o receiver1/receiver receiver1/receiver.go
docker build -t docker.io/surajd/activemq-receiver1:1.0.1 ./receiver1

go build -o receiver2/receiver receiver2/receiver.go
docker build -t docker.io/surajd/activemq-receiver2:1.0.1 ./receiver2

oc new-project amqptest4
kedge apply -f kedge.yml

sleep 20

kedge apply -f kedge-receiver2.yml

oc get pods
