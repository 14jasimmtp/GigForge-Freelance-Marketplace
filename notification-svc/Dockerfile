FROM golang:1-alpine3.19 AS BaseImage

WORKDIR /Notification-svc

COPY . ./

RUN go mod download

RUN GOOS=linux go build -o ./notification-svc ./cmd

RUN apk add -U --no-cache ca-certificates


FROM scratch AS build-stage 

WORKDIR /

COPY --from=BaseImage ./Notification-svc/notification-svc /
COPY --from=BaseImage ./Notification-svc/config.env /
COPY --from=BaseImage /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/


CMD [ "/notification-svc"]