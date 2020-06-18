FROM golang:1.13.0

RUN apt-get update && apt-get install -y

ENV PKG_NAME=bukuduit-go/
ENV PKG_PATH=$GOPATH/src/$PKG_NAME
WORKDIR $PKG_PATH/


COPY . $PKG_PATH/

RUN echo $PWD
RUN go mod vendor

WORKDIR $PKG_PATH/server/
RUN echo $PWD

RUN go build main.go
EXPOSE 3000
CMD ["./main"]
