.PHONY: amq-build
amq-build:
	-eval $(minishift docker-env)
	docker build -t docker.io/surajd/activemq-artemis:2.6.0 amq

.PHONY: amq-deploy
amq-deploy:
	-oc new-project amq
	-oc delete configmap apache-artemis-cm
	-oc create configmap apache-artemis-cm --from-file=amq/config
	oc apply -f amq/activemq-artemis-deploy.yaml

.PHONY: test1-electron
test1-electron:
	cd test1-electron && ./build-deploy.sh

.PHONY: test2-electron
test2-electron:
	cd test2-electron && ./build-deploy.sh

.PHONY: test3-electron
test3-electron:
	cd test3-electron && ./build-deploy.sh

.PHONY: test4-electron
test4-electron:
	cd test4-electron && ./build-deploy.sh

.PHONY: test5-electron
test5-electron:
	cd test5-electron && ./build-deploy.sh

.PHONY: test6-electron
test6-electron:
	cd test6-electron && ./build-deploy.sh
