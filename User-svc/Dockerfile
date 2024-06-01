FROM golang:1-alpine3.19 AS BaseImage

WORKDIR /User-svc

COPY . ./

RUN go mod download

RUN GOOS=linux go build -o ./user-svc ./cmd

RUN apk add -U --no-cache ca-certificates


FROM scratch AS build-stage 

WORKDIR /

COPY --from=BaseImage ./User-svc/user-svc /
COPY --from=BaseImage ./User-svc/config.env /
COPY --from=BaseImage /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/


CMD [ "/user-svc"]