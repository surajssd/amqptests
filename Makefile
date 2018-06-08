.PHONY: amq-build
amq-build:
	eval $(minishift docker-env)
	docker build -t docker.io/surajd/activemq-artemis:2.6.0 amq

.PHONY: amq-deploy
amq-deploy:
	-oc new-project amq
	oc apply -f amq/activemq-artemis-deploy.yaml

.PHONY: test1
test1:
	cd test1 && ./build-deploy.sh

.PHONY: test2
test2:
	cd test2 && ./build-deploy.sh

.PHONY: test3
test3:
	cd test3 && ./build-deploy.sh

.PHONY: test4
test4:
	cd test4 && ./build-deploy.sh

.PHONY: test6
test6:
	cd test6 && ./build-deploy.sh

.PHONY: test7
test7:
	cd test7 && ./build-deploy.sh
