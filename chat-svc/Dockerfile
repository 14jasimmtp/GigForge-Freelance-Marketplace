FROM golang:1-alpine3.19 AS BaseImage

WORKDIR /Chat-svc

COPY . ./

RUN go mod download

RUN GOOS=linux go build -o ./chat-svc ./cmd

FROM scratch AS build-stage 

WORKDIR /

COPY --from=BaseImage ./Chat-svc/chat-svc /
COPY --from=BaseImage ./Chat-svc/config.env /

CMD [ "/chat-svc"]