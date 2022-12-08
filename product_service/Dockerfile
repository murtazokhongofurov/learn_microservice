FROM golang:1.19.3-alpine
RUN mkdir product
COPY . /product
WORKDIR /product
RUN go mod tidy
RUN go build -o main cmd/main.go
CMD ./main
EXPOSE 5000