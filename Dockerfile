# Use an official Golang runtime as a parent image
FROM golang:1.16-alpine

# Set the working directory in the container
WORKDIR /app

# Copy the local package files to the container's workspace
ADD . /app

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Run the executable
CMD ["app"]
