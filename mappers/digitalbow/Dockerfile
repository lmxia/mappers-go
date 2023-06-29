
FROM ubuntu:16.04

RUN mkdir -p kubeedge

COPY ./bin/digitalbow kubeedge/
COPY ./config.yaml kubeedge/

WORKDIR kubeedge

ENTRYPOINT ["/kubeedge/digitalbow", "--v", "5"]
