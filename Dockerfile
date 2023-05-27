FROM golang:1.20.4-bullseye AS build-env

WORKDIR /go/src/github.com/fury-labs/blackfury

COPY . .

RUN make build

FROM golang:1.20.4-bullseye

RUN apt-get update  \ 
&& apt-get install ca-certificates jq=1.6-2.1 -y --no-install-recommends

WORKDIR /root

COPY --from=build-env /go/src/github.com/fury-labs/blackfury/build/blackfuryd /usr/bin/blackfuryd

EXPOSE 26656 26657 1317 9090 8545 8546

CMD ["blackfuryd"]
