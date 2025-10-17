# What
Confluent kafka setup using docker-compose with PLAINTEXT, SASL_PLAINTEXT, SSL and SASL_SSL listeners.

# How to run
`sudo docker-compose up -d`

# How to test
## client-sasl-ssl.properties
```
sasl.mechanism=PLAIN
security.protocol=SASL_SSL
group.id=console-consumer-group
sasl.jaas.config=org.apache.kafka.common.security.plain.PlainLoginModule required \
      username="alice" \
      password="alice-secret";

ssl.truststore.location=./kafka.broker.truststore.jks
ssl.truststore.password=broker-password
ssl.endpoint.identification.algorithm=
```

`bash kafka-topics.sh --list --bootstrap-server ec2.us-west-2.compute.amazonaws.com:9095 --command-config client-sasl-ssl.properties`
