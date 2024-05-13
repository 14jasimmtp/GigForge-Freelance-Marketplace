FROM golang:1-alpine3.19 AS BaseImage

WORKDIR /User-svc

COPY . ./

RUN go mod download

RUN GOOS=linux go build -o ./user-svc ./cmd

FROM scratch AS build-stage 

WORKDIR /

COPY --from=BaseImage ./User-svc /
COPY --from=BaseImage ./User-svc/config.env /

ENTRYPOINT [ "/user-svc"]