# AMQP tests

This repository has the tests done using electron library for ActiveMQ Artemis.

# Try it out yourself

```bash
make test2
```

# Test cases being run

* Test Case 1 - one service send one msg sent to multiple service, one instance of each service recv msg.
    - T1 - srv1 send msg
    - T2 - srv2 (i1) get msg
    - T3 - srv3 (i1) get msg

* Test Case 2 - when one of the recv instance is down other instance should get msg (assuming srv3(i1) is down at T3)
    - T1 - srv1 send msg
    - T2 - srv2 (i1) get msg
    - T3 - srv3 (i2) get msg

* Test Case 3 - when many msg sent, on recv side the msg should be load balanced across multiple recv instances.
    - T1 - srv1 send multiple msg
    - T2 - srv2 (i1) get all msg
    - T3 - srv3 (i1) get few msg
    - T4 - srv3 (i2) get rest msg

Here, all msg should be load-balanced between srv3(i1) and srv3(i2). 

* Test Case 4 - when recv down initially and comes up later then it should get msg. (assuming srv2 (i1) is down initially)

    - T1 - srv1 send msg
    - T2 - srv3 (i1) get msg
    - T3 - nothing happened
    - T4 - srv2 comes up
    - T5 - srv2 get msg

* Test Case 5 - msg_sys should be durable.

    - T1 - srv1 send msg
    - T2 - nats has got msg
    - T3 - nats got down
    - T4 - nats comes up
    - T5 - srv2 get msg
    - T6 - srv3 get msg

* Test Case 6 - when recv failed to process msg correctly, msg system should re-deliver same msg.

    - T1 - srv1 send msg
    - T2 - srv2 (i1) get msg
    - T3 - srv3 (i1) get msg but failed to deliver to actual service
    - T4 - nats re-deliver same msg to srv3
    - T5 - srv3 (i1) get same msg

Re-delivery support either provided by msg system or done in the integration service itself
