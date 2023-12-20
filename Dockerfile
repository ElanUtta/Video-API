FROM golang:1.21.4


WORKDIR /app



COPY . ./


#download dependencies
RUN go mod download



RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

CMD ["/docker-gs-ping"]