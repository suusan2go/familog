FROM golang:alpine AS build-env
WORKDIR /usr/local/go/src/github.com/suusan2go/familog
COPY . /usr/local/go/src/github.com/suusan2go/familog
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN go get ./...
RUN go build -o build/familog ./familog


FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=build-env /usr/local/go/src/github.com/suusan2go/familog/build/familog /bin/familog
CMD ["familog", "up"]
