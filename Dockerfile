# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/datdao0712/BUNTEAM 
RUN go get -u github.com/beego/bee 
RUN go get -u github.com/astaxie/beego
# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/datdao0712/BUNTEAM
RUN mv /go/bin/BUNTEAM /go/src/github.com/datdao0712/BUNTEAM
# Run the outyet command by default when the container starts.
WORKDIR /go/src/github.com/datdao0712/BUNTEAM
ENTRYPOINT ./BUNTEAM

# Document that the service listens on port 8080.
EXPOSE 8080
