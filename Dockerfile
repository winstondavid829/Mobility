FROM golang:alpine

# Setting up necessary variable for the environment

ENV GO1111MODULE = on\
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH = amd64

# Moving files to docker working directory
WORKDIR /build

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container
COPY . .


# Build the application
RUN go build -v -o main .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

# Export necessary port
EXPOSE 8080

# Command to run when starting the container
CMD ["/dist/main"]