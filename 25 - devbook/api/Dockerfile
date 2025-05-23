FROM golang:1.12-alpine

ARG API_PORT=${API_PORT}

# Add Maintainer Info
LABEL maintainer="Thiago Fortunato <thiagolsfortunato@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# Build the Go app
RUN go build -o main .

# Expose port 3000 to the outside world
EXPOSE ${API_PORT}

# Run the executable
CMD ["./main"]
