FROM golang:1-alpine3.19 AS BaseImage

WORKDIR /API-Gateway

COPY . ./

RUN go mod download

RUN GOOS=linux go build -o ./api-gateway ./cmd


FROM scratch AS build-stage 

WORKDIR /

COPY --from=BaseImage ./API-Gateway/api-gateway /
COPY --from=BaseImage ./API-Gateway/config.env /
COPY --from=BaseImage ./API-Gateway/template /template

CMD [ "/api-gateway"]   