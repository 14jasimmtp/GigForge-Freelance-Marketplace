FROM golang:1-alpine3.19 AS BaseImage

WORKDIR /Project-svc

COPY . ./

RUN go mod download

RUN GOOS=linux go build -o ./project-svc ./cmd

FROM scratch AS build-stage 

WORKDIR /

COPY --from=BaseImage ./Project-svc /
COPY --from=BaseImage ./Project-svc/config.env /

ENTRYPOINT [ "/project-svc"]