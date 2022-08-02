FROM golang:latest 

WORKDIR /
# COPY go.mod ./
# COPY go.sum ./
COPY  ./ ./
EXPOSE 1313
EXPOSE 5555

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# # make wait-for-postgres.sh executable
# RUN chmod +x wait-for-postgres.sh

# build go app
RUN go mod download
RUN go build cmd/main.go
CMD ["./main"]
