FROM debian:11

WORKDIR /mariomang
COPY ./build/bin/* /mariomang

ENV GITHUB_TOKEN=your_token

RUN apt-get update && apt-get install -y ca-certificates

CMD ["./bio"]