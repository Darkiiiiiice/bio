FROM ubuntu:latest

WORKDIR /mariomang

RUN apt-get update && apt-get install -y ca-certificates libc6


COPY ./build/bin/bio /mariomang/bio

ENV GITHUB_TOKEN=your_token

CMD ["/mariomang/bio"]