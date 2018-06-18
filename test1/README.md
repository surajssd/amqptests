# Test1

## What this test is?

Test Case 1 - one service send one msg sent to multiple service, one instance of each service recv msg.
  * T1 - srv1 send msg
  * T2 - srv2 (i1) get msg
  * T3 - srv3 (i1) get msg

## Seeing it in action

Running the make, which builds and deploys the images.

```bash
$ make test1
cd test1 && ./build-deploy.sh
/home/hummer/go/src/github.com/surajssd/amqptests/test1
./build-deploy.sh: line 3: Running: command not found
Sending build context to Docker daemon 4.291 MB
Step 1/4 : FROM fedora:28
 ---> cc510acfcd70
Step 2/4 : RUN dnf -y install qpid-proton-c-devel && dnf clean all
 ---> Using cache
 ---> 05c3cc498d9a
Step 3/4 : COPY sender /sender
 ---> Using cache
 ---> 1d9125e2d6db
Step 4/4 : CMD /sender
 ---> Using cache
 ---> 8183a7bf4413
Successfully built 8183a7bf4413
Sending build context to Docker daemon  4.29 MB
Step 1/4 : FROM fedora:28
 ---> cc510acfcd70
Step 2/4 : RUN dnf -y install qpid-proton-c-devel && dnf clean all
 ---> Using cache
 ---> 05c3cc498d9a
Step 3/4 : COPY receiver /receiver
 ---> Using cache
 ---> b854d72e2909
Step 4/4 : CMD /receiver
 ---> Using cache
 ---> 712a08f0c4bd
Successfully built 712a08f0c4bd
Now using project "test1" on server "https://192.168.122.1:8443".

You can add applications to this project with the 'new-app' command. For example, try:

    oc new-app centos/ruby-22-centos7~https://github.com/openshift/ruby-ex.git

to build a new example application in Ruby.
deployment.extensions "sender" created
deployment.extensions "receiver1" created
deployment.extensions "receiver2" created

NAME                         READY     STATUS              RESTARTS   AGE
receiver1-56568b75b-sq9xr    0/1       ContainerCreating   0          1s
receiver2-5695dc86fd-4nc5c   0/1       ContainerCreating   0          1s
sender-7ff856c596-9kzq8      0/1       ContainerCreating   0          1s
```

Pods that have come up:

```bash
$ oc get pods
NAME                         READY     STATUS    RESTARTS   AGE
receiver1-56568b75b-sq9xr    1/1       Running   0          5s
receiver2-5695dc86fd-4nc5c   1/1       Running   0          5s
sender-7ff856c596-9kzq8      1/1       Running   0          5s
```

Sender is sending:

```bash
$ oc logs sender-7ff856c596-9kzq8
2018/06/18 09:50:38 [*] sent message: hello from "sender" on "sender-7ff856c596-9kzq8"! 0
2018/06/18 09:50:40 [*] sent message: hello from "sender" on "sender-7ff856c596-9kzq8"! 1
2018/06/18 09:50:42 [*] sent message: hello from "sender" on "sender-7ff856c596-9kzq8"! 2
2018/06/18 09:50:44 [*] sent message: hello from "sender" on "sender-7ff856c596-9kzq8"! 3
2018/06/18 09:50:46 [*] sent message: hello from "sender" on "sender-7ff856c596-9kzq8"! 4
2018/06/18 09:50:48 [*] sent message: hello from "sender" on "sender-7ff856c596-9kzq8"! 5
```

Receiver1 gets all of them:

```bash
$ oc logs receiver1-56568b75b-sq9xr
2018/06/18 09:50:38 [*] message received on "receiver1" on "receiver1-56568b75b-sq9xr": hello from "sender" on "sender-7ff856c596-9kzq8"! 0
2018/06/18 09:50:40 [*] message received on "receiver1" on "receiver1-56568b75b-sq9xr": hello from "sender" on "sender-7ff856c596-9kzq8"! 1
2018/06/18 09:50:42 [*] message received on "receiver1" on "receiver1-56568b75b-sq9xr": hello from "sender" on "sender-7ff856c596-9kzq8"! 2
2018/06/18 09:50:44 [*] message received on "receiver1" on "receiver1-56568b75b-sq9xr": hello from "sender" on "sender-7ff856c596-9kzq8"! 3
2018/06/18 09:50:46 [*] message received on "receiver1" on "receiver1-56568b75b-sq9xr": hello from "sender" on "sender-7ff856c596-9kzq8"! 4
2018/06/18 09:50:48 [*] message received on "receiver1" on "receiver1-56568b75b-sq9xr": hello from "sender" on "sender-7ff856c596-9kzq8"! 5
2018/06/18 09:50:50 [*] message received on "receiver1" on "receiver1-56568b75b-sq9xr": hello from "sender" on "sender-7ff856c596-9kzq8"! 6
2018/06/18 09:50:52 [*] message received on "receiver1" on "receiver1-56568b75b-sq9xr": hello from "sender" on "sender-7ff856c596-9kzq8"! 7
```

Receiver2 also gets them all:

```bash
$ oc logs receiver2-5695dc86fd-4nc5c
2018/06/18 09:50:38 [*] message received on "receiver2" on "receiver2-5695dc86fd-4nc5c": hello from "sender" on "sender-7ff856c596-9kzq8"! 0
2018/06/18 09:50:40 [*] message received on "receiver2" on "receiver2-5695dc86fd-4nc5c": hello from "sender" on "sender-7ff856c596-9kzq8"! 1
2018/06/18 09:50:42 [*] message received on "receiver2" on "receiver2-5695dc86fd-4nc5c": hello from "sender" on "sender-7ff856c596-9kzq8"! 2
2018/06/18 09:50:44 [*] message received on "receiver2" on "receiver2-5695dc86fd-4nc5c": hello from "sender" on "sender-7ff856c596-9kzq8"! 3
2018/06/18 09:50:46 [*] message received on "receiver2" on "receiver2-5695dc86fd-4nc5c": hello from "sender" on "sender-7ff856c596-9kzq8"! 4
2018/06/18 09:50:48 [*] message received on "receiver2" on "receiver2-5695dc86fd-4nc5c": hello from "sender" on "sender-7ff856c596-9kzq8"! 5
2018/06/18 09:50:50 [*] message received on "receiver2" on "receiver2-5695dc86fd-4nc5c": hello from "sender" on "sender-7ff856c596-9kzq8"! 6
2018/06/18 09:50:52 [*] message received on "receiver2" on "receiver2-5695dc86fd-4nc5c": hello from "sender" on "sender-7ff856c596-9kzq8"! 7
2018/06/18 09:50:54 [*] message received on "receiver2" on "receiver2-5695dc86fd-4nc5c": hello from "sender" on "sender-7ff856c596-9kzq8"! 8
```

## Conclusion

What sender has sent is recieved by the receivers so the test passed!
