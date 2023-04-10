FROM debian:11

WORKDIR /mariomang

RUN apt-get update && apt-get install -y ca-certificates

COPY ./build/bin/bio /mariomang/bio

ENV GITHUB_TOKEN=your_token

CMD ["/mariomang/bio"]