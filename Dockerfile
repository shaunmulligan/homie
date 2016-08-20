FROM resin/raspberrypi3-golang:1.6.2-slim-20160712
ENV VERSION=0.0.1

RUN apt-get update && apt-get install -yq \
    apt-transport-https && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

RUN curl -sL https://repos.influxdata.com/influxdb.key | sudo apt-key add -
RUN echo "deb https://repos.influxdata.com/debian jessie stable" | sudo tee /etc/apt/sources.list.d/influxdb.list

RUN apt-get update && apt-get install -yq \
    influxdb && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

ENV GOPATH=/usr/src/app
WORKDIR $GOPATH
COPY . ./
RUN go build -o goapp.o github.com/shaunmulligan/goapp
ENV INITSYSTEM=on
CMD ["./start.sh"]
