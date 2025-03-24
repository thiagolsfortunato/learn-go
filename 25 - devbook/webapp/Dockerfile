FROM golang:1.23 as builder 

ARG WEBAPP_PORT=${WEBAPP_PORT}

# Add Maintainer Info
LABEL maintainer="Thiago Fortunato <thiagolsfortunato@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /webapp

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# Build the Go app
RUN go build -o webapp .

FROM alpine:latest

WORKDIR /webapp

COPY --from=builder /webapp /webapp

# Expose port 8080 to the outside world
EXPOSE ${WEBAPP_PORT}

# Run the executable
CMD ["./main"]
