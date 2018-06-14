#!/bin/bash

set -e
eval $(minishift docker-env)
set +e

go build -o sender/sender sender/sender.go
docker build -t docker.io/surajd/activemq-sender:fed-1.0.1 ./sender

go build -o receiver/receiver receiver/receiver.go
docker build -t docker.io/surajd/activemq-receiver:fed-1.0.1 ./receiver

oc new-project test3-electron
kedge apply -f kedge.yml

oc get pods
