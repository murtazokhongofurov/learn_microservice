FROM golang:1.19.3-alpine
RUN mkdir store
COPY . /store
WORKDIR /store
RUN go mod tidy
RUN go build -o main cmd/main.go
CMD ./main
EXPOSE 3000