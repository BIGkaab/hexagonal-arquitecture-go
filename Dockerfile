# Start from golang base image
FROM golang:1.19.0-alpine3.16

# Add Maintainer info
LABEL maintainer="BIGkaab"

# Setup folders
RUN mkdir /app
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container
COPY . .
COPY .env .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Build the Go app
RUN go build -o /build

# Expose port 3005 to the outside world
EXPOSE 3005

# Run the executable
CMD [ "/build" ]