FROM local-mike:19443/alpine:latest
LABEL MAINTAINER="MW"

COPY ./sonatype-iq-slackbot /app/sonatype-iq-slackbot

ENV PORT 9000
EXPOSE 9000

ENTRYPOINT ["/app/sonatype-iq-slackbot"]