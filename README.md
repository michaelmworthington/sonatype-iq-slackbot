[![pipeline status](https://gitlab.com/michaelmworthington/sonatype-iq-slackbot/badges/master/pipeline.svg)](https://gitlab.com/michaelmworthington/sonatype-iq-slackbot/commits/master) [![coverage report](https://gitlab.com/michaelmworthington/sonatype-iq-slackbot/badges/master/coverage.svg)](https://gitlab.com/michaelmworthington/sonatype-iq-slackbot/commits/master)

# sonatype-iq-slackbot

written in go

## after building and running, make available via ngrok and update slack
```
./ngrok http 9000
update ngrok url in the slack bot slash command
```

## build the app:
```
go build
```

## run the app:
```
./sonatype-iq-slackbot
```

## build and publish with docker
build outside w/ alpine image = 14 MB

build inside w/ golang image = 805 MB

```
GOOS=linux GOARCH=amd64 go build

docker build --rm -f "Dockerfile" -t sonatype-iq-slackbot:latest .
docker tag sonatype-iq-slackbot:latest local-mike:19447/michaelmworthington/sonatype-iq-slackbot:1.0.0
docker push local-mike:19447/michaelmworthington/sonatype-iq-slackbot:1.0.0
docker images | grep slack
```

## running with docker:
```
docker run -it -p 9000:9000 sonatype-iq-slackbot:latest
docker run --name sonatype-iq-slackbot --rm -d -p 9000:9000 sonatype-iq-slackbot:latest
```

## running with docker compose:
```
docker-compose build
docker-compose up -d
docker-compose down
```

## running with kubernetes (simple test):
```
kubectl create -f 01-k8s-pod.yml (pull from local registry - see evernotes)
```

## running with kubernetes (more advanced test):
```
kubectl create -f 02-k8s-deployment.yml
kubectl create -f 03-k8s-service.yml
```

## kubernetes (scale, update & rollback):
```
kubectl scale deployment <name> —replicas=5
kubectl set image deployment sonatype-iq-slackbot-deployment sonatype-iq-slackbot=local-mike:19447/michaelmworthington/sonatype-iq-slackbot:1.0.3
kubectl rollout undo deployment sonatype-iq-slackbot-deployment
```