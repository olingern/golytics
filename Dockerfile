# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/olingern/golytics

WORKDIR /go/src/github.com/olingern/golytics

RUN go install

ENTRYPOINT /go/bin/golytics

# Document that the service listens on port 8080.
EXPOSE 8080