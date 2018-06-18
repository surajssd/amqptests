#!/bin/bash

eval $(minishift docker-env)

go build -o sender/sender sender/sender.go
docker build -t docker.io/surajd/activemq-sender:fed-1.0.1 ./sender

go build -o receiver/receiver receiver/receiver.go
docker build -t docker.io/surajd/activemq-receiver:fed-1.0.1 ./receiver

oc new-project test4
kedge apply -f kedge.yml

oc get pods

# we should wait to see if msgs reach receiver2
# even if it came up late to the party
sleep 100
kedge apply -f kedge2.yml
