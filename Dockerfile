# syntax=docker/dockerfile:1
# API
FROM alpine:latest

WORKDIR /

COPY scripts/setup.sh /
COPY pages/ pages/
COPY config/default.yaml config/
COPY exodialib-core /

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.3/wait /wait
RUN chmod +x /wait
RUN chmod +x /setup.sh
RUN cd pages && ls -al

EXPOSE 9001

ENTRYPOINT [ "/setup.sh" ]