FROM ubuntu:latest

WORKDIR /mariomang
COPY ./build/bin/bio /mariomang/bio

RUN apt-get update && apt-get install -y ca-certificates

ENV GITHUB_TOKEN=your_token

CMD ["/mariomang/bio"]