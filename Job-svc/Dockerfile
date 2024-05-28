FROM golang:1-alpine3.19 AS BaseImage

WORKDIR /Job-svc

COPY . ./

RUN go mod download

RUN GOOS=linux go build -o ./job-svc ./cmd

FROM scratch AS build-stage 

WORKDIR /

COPY --from=BaseImage ./Job-svc/job-svc /
COPY --from=BaseImage ./Job-svc/config.env /

CMD [ "/job-svc"]