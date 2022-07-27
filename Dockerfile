FROM golang:latest 

WORKDIR /
# COPY go.mod ./
# COPY go.sum ./
COPY  . .
EXPOSE 1313

# RUN go get -d -v
RUN go build cmd/main.go
CMD ["./main"]
