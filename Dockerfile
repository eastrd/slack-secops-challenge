FROM golang:1.13

WORKDIR /app
COPY . .

# Run unit tests with fake envionrment variables
ENV CERT=somecert/cert.crt KEY=somecert/key.key BASIC_USER=testuser BASIC_PASS=testpass
RUN go test 2>&1 | grep "PASS" || exit 2

# Compile to binary exec
RUN go build -o server api.go main.go structs.go

EXPOSE 443/tcp

CMD ["./server"]