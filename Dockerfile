FROM golang:1.17.3-bullseye as base
WORKDIR /app
COPY . .
RUN go mod download

FROM base as build
RUN go build -o ./bookstore-catalog-v1 .

#Target for integration testing
FROM base as integ
CMD ["go", "test", "-v", "."]

#Target for release
FROM build as release
CMD ["./bookstore-catalog-v1"]