FROM debian:stable

WORKDIR /mariomang
COPY ./build/* /mariomang

ENV GITHUB_TOKEN=your_token

RUN apt-get update && apt-get install -y ca-certificates

CMD ["./bio"]