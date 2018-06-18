#!/bin/bash

eval $(minishift docker-env)

go build -o sender/sender sender/sender.go
docker build -t docker.io/surajd/activemq-sender:fed-1.0.1 ./sender

go build -o receiver/receiver receiver/receiver.go
docker build -t docker.io/surajd/activemq-receiver:fed-1.0.1 ./receiver

oc new-project test2
kedge apply -f kedge.yml

oc get pods

while true
do
    oc delete pod $(oc get pods | grep receiver2 | awk '{print $1}' | head -1)
    sleep 10
done
