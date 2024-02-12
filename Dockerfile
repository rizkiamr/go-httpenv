FROM golang:1.18.1 AS build

RUN mkdir -p /opt/build

WORKDIR /opt/build

# Copy only necessary files
COPY go.mod ./
RUN go mod download

# Copy the rest of the files
COPY . .

# Do the build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

FROM gcr.io/distroless/static
USER nobody:nobody
WORKDIR /
COPY --from=build /opt/build/main .
ENTRYPOINT []
