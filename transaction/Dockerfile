FROM golang:1.9

ADD . /go/src/simple-microservices/transaction
RUN go get gopkg.in/mgo.v2
RUN go get gopkg.in/mgo.v2/bson
RUN go get github.com/twitchtv/twirp
RUN go get github.com/twitchtv/twirp/ctxsetters
RUN go get github.com/golang/protobuf/proto
RUN go get github.com/golang/protobuf/jsonpb

WORKDIR /go/src/simple-microservices/transaction/cmd/server/
RUN go build main.go
