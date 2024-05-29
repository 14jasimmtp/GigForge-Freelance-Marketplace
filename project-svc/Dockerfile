FROM golang:1-alpine3.19 AS BaseImage

WORKDIR /Project-svc

COPY . ./

RUN go mod download

RUN GOOS=linux go build -o ./project-svc ./cmd

RUN apk add -U --no-cache ca-certificates


FROM scratch AS build-stage 

WORKDIR /

COPY --from=BaseImage ./Project-svc/project-svc /
# COPY --from=BaseImage ./Project-svc/config.env /
COPY --from=BaseImage /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD [ "/project-svc"]