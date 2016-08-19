FROM resin/raspberrypi3-golang:1.6.2-slim-20160712
ENV VERSION=0.0.1
ENV GOPATH=/usr/src/app
WORKDIR $GOPATH
COPY . ./
RUN go build -o goapp.o github.com/shaunmulligan/goapp
ENV INITSYSTEM=on
WORKDIR /
CMD ["modprobe", "i2c-dev && $GOPATH/goapp.o"]
