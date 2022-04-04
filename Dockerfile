FROM alpine

MAINTAINER Joan Gargallo <peppelin@gmail.com>

COPY out/bin/hello-world /usr/bin

WORKDIR /opt/project
ENTRYPOINT ["/usr/bin/hello-world"]