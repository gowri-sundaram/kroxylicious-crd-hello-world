# What ?
### Confluent kafka setup using docker-compose with SASL_PLAINTEXT and PLAINTEXT listeners.

# How to run ?
`sudo docker-compose up -d`

# How to test ?
## client.properties
```
sasl.mechanism=PLAIN
security.protocol=SASL_PLAINTEXT
group.id=console-consumer-group
sasl.jaas.config=org.apache.kafka.common.security.plain.PlainLoginModule required \
      username="alice" \
      password="alice-secret";
```
` bash kafka-topics.sh --list --bootstrap-server ec2-35-90-106-177.us-west-2.compute.amazonaws.com:9093 --command-config client.properties`
