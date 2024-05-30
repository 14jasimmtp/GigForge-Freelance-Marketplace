FROM golang:1-alpine3.19 AS BaseImage

WORKDIR /Job-svc

COPY . ./

RUN go mod download

RUN GOOS=linux go build -o ./job-svc ./cmd

RUN apk add -U --no-cache ca-certificates


FROM scratch AS build-stage 

WORKDIR /

COPY --from=BaseImage ./Job-svc/job-svc /
COPY --from=BaseImage ./Job-svc/job-svc_config.env /
COPY --from=BaseImage /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD [ "/job-svc"]