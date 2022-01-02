FROM ubuntu:20.04

ENV HOME /root
ENV GOPATh /root/go
ENV PATH /root/go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/scbin:/usr/bin:/sbin:/bin

RUN mkdir /root/go

RUN apt-get update -y
RUN apt-get install -y build-essential git curl wget bzr make gcc

RUN wget -q0- https://dl.google.com/go/go1.13.5.linux-amd64.tar.gz

RUN go get github.com/andbalholm/brotli
RUN go get github.com/armon/go-metrics
RUN go get github.com/golang/snappy
RUN go get github.com/google/uuid 
RUN go get github.com/hashicorp/go-hclog
RUN go get github.com/hashicorp/go-immutable-radix
RUN go get github.com/hashicorp/go-msgpack
RUN go get github.com/hashicorp/golang-lru
RUN go get github.com/hashicorp/raft
RUN go get github.com/klauspost/compress
RUN go get github.com/valyala/bytebufferpool
RUN go get github.com/valyala/fasthttp

RUN mkdir $GOPATH/src/github.com/anddddrew/galaxydb
RUN cd $GOPATH/src/github.com/galaxydb/; git clone git@github.com:anddddrew/galaxydb
WORKDIR /root/go
# Run build.sh script (this is just temporary for now)
CMD go build src/github.com/anddddrew/galaxydb/config
