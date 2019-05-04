[![pipeline status](https://gitlab.com/michaelmworthington/sonatype-iq-slackbot/badges/master/pipeline.svg)](https://gitlab.com/michaelmworthington/sonatype-iq-slackbot/commits/master) [![coverage report](https://gitlab.com/michaelmworthington/sonatype-iq-slackbot/badges/master/coverage.svg)](https://gitlab.com/michaelmworthington/sonatype-iq-slackbot/commits/master)

# sonatype-iq-slackbot

written in go

```
GOOS=linux GOARCH=amd64 go build
docker build --rm -f "Dockerfile" -t sonatype-iq-slackbot:latest .
docker run -it -p 9000:9000 sonatype-iq-slackbot:latest
docker run --name sonatype-iq-slackbot --rm -d -p 9000:9000 sonatype-iq-slackbot:latest
docker tag sonatype-iq-slackbot:latest local-mike:19447/michaelmworthington/sonatype-iq-slackbot:1.0.0
docker push local-mike:19447/michaelmworthington/sonatype-iq-slackbot:1.0.0
docker images | grep slack
```

build outside = 14 MB
build inside = 805 MB

docker compose:
```
docker-compose build
docker-compose up -d
docker-compose down
```