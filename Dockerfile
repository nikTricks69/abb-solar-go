FROM golang:1.16-alpine

ARG abb_user=user21
ARG abb_password=password
ARG abb_ip=192.168.80.172

ENV abb_user=$abb_user
ENV abb_password=$abb_password
ENV abb_ip=$abb_ip
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /docker-abb-prom-metrics

EXPOSE 8869

CMD [ "/docker-abb-prom-metrics" ]
