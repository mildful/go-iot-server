# Use binaries.
# docker run --rm -it -v "$GOPATH":/gopath -v "$(pwd)":/app -e "GOPATH=/gopath" -w /app golang:1.4.2 sh -c 'CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o hello'
FROM iron/base
MAINTAINER Clement Flodrops<clement.flodrops@gmail.com>
WORKDIR /app
# copy binary into image
COPY hello /app/
ENTRYPOINT ["./hello"]
