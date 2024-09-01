FROM docker.io/ppc64le/golang:latest
RUN pwd
WORKDIR /app
COPY . /app 
WORKDIR /app/metrics_server
RUN go get github.com/prometheus/client_golang/prometheus 
RUN go build -o metrics metrics.go
EXPOSE 9080
#ENTRYPOINT ["go","run","./metrics_server/metrics.go"]
CMD ["/app/metrics_server/metrics"]
