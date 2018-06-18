#!/bin/bash

eval $(minishift docker-env)

go build -o sender/sender sender/sender.go
docker build -t docker.io/surajd/activemq-sender:fed-1.0.1 ./sender

go build -o receiver/receiver receiver/receiver.go
docker build -t docker.io/surajd/activemq-receiver:fed-1.0.1 ./receiver

oc new-project test5
kedge apply -f kedge.yml

oc get pods


giventime=30
while :
do
    echo "Sleeping for $giventime seconds ..."
    sleep $giventime
    oc delete pods -n amq --all
    giventime=$(echo "$(($giventime*2))")
done
