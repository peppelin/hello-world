FROM alpine

MAINTAINER Joan Gargallo <peppelin@gmail.com>

COPY out/bin/hello-world /usr/bin

ENTRYPOINT ["/usr/bin/hello-world"]