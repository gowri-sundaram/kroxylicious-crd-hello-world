# Kroxylicious starter operator
## Prerequisites
### Setup remote kafka cluster
Guide: https://developer.confluent.io/confluent-tutorials/kafka-on-docker/#the-docker-compose-file

Change `KAFKA_ADVERTISER_LISTENERS` in the above guide to use the public DNS as required.

```KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://broker:29092,PLAINTEXT_HOST://ec2-abcd.us-west-2.compute.amazonaws.com:9092```

### Start local minikube cluster
```minikube start --driver=docker --ports=8443:8443```

## Running kroxy operator
### Build local docker image and add to minikube
```
docker build -t kroxylicious .
minikube image load kroxylicious:latest
```
### Add kroxy config file to minikube
```
minikube ssh
mkdir -p /tmp
vi /tmp/example-proxy-config.yaml  # Add the configs here
```
### Build and run the kroxy operator
```
go mod tidy
go build -o ./kroxy-operator main.go
./kroxy-operator
```
### Apply the CRD and custom resource
```
kubectl apply -f config/crd/bases/kroxy-crd.yaml
kubectl apply -f kroxy-container.yaml
```
### Verify that pods have been created
```
kubectl get kroxies
kubectl describe pod my-kroxy
```
### Enable port forwarding
```
kubectl port-forward pod/my-kroxy 9192:9192 9193:9193 9194:9194
```
## Testing kroxy operator
### Create topics
```
bash kafka-topics.sh --create --topic sushi --bootstrap-server localhost:9192
```
### Console producer
```
bash kafka-console-producer.sh  --topic sushi --bootstrap-server localhost:9192
```
### Console consumer
```
bash kafka-console-consumer.sh --topic sushi --from-beginning --bootstrap-server localhost:9192
```
## Cleanup
```
kubectl delete -f kroxy-container.yaml
kubectl delete -f config/crd/bases/kroxy-crd.yaml
minikube stop
```


