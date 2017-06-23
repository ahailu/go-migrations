FROM golang:1.8.0-alpine

# HAVE TO RUN THE BUILD FIRST
RUN apk update
RUN apk add --no-cache bash ca-certificates
RUN update-ca-certificates \
  && apk add --update make \
  && apk add git \
  && git clone https://github.com/Masterminds/glide.git  /go/src/github.com/Masterminds/glide \
  && cd /go/src/github.com/Masterminds/glide \
  && make install \ 
  && mkdir -p /go/src/scheduler

# COPY BUILD
COPY . /go/src/scheduler
WORKDIR /go/src/scheduler

# RUN
RUN echo $GOPATH 
RUN glide install \
  && CGO_ENABLED=0 GOOS=$(uname -s | tr A-Z a-z) go build -ldflags "-s" -a -installsuffix cgo  -o /go/bin/scheduler scheduler

EXPOSE 8080
CMD ["/go/bin/scheduler"]
