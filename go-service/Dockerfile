FROM debian:wheezy

LABEL authors="Usama"

# Install Thrift.

ENV THRIFT_VERSION 0.11.0

RUN buildDeps=" \
		automake \
		bison \
		curl \
		flex \
		g++ \
		libboost-dev \
		libboost-filesystem-dev \
		libboost-program-options-dev \
		libboost-system-dev \
		libboost-test-dev \
		libevent-dev \
		libssl-dev \
		libtool \
		make \
		pkg-config \
	"; \
	apt-get update && apt-get install -y --no-install-recommends $buildDeps && rm -rf /var/lib/apt/lists/* \
	&& curl -sSL "http://apache.mirrors.spacedump.net/thrift/$THRIFT_VERSION/thrift-$THRIFT_VERSION.tar.gz" -o thrift.tar.gz \
	&& mkdir -p /usr/src/thrift \
	&& tar zxf thrift.tar.gz -C /usr/src/thrift --strip-components=1 \
	&& rm thrift.tar.gz \
	&& cd /usr/src/thrift \
	&& ./configure  --without-python --without-cpp \
	&& make \
	&& make install \
	&& cd / \
	&& rm -rf /usr/src/thrift \
	&& curl -k -sSL "https://storage.googleapis.com/golang/go1.4.linux-amd64.tar.gz" -o go.tar.gz \
	&& tar xzf go.tar.gz \
	&& rm go.tar.gz \
	&& cp go/bin/gofmt /usr/bin/gofmt \
	&& rm -rf go \
	&& apt-get purge -y --auto-remove $buildDeps

# Install Go.

RUN apt-get update
RUN apt-get -f install
RUN apt-get install -y ca-certificates
RUN apt-get install -y curl

RUN curl -O https://storage.googleapis.com/golang/go1.10.1.linux-amd64.tar.gz
#RUN sha256sum go1.10*.tar.gz
RUN tar xvf go1.10.1.linux-amd64.tar.gz

RUN chown -R root:root ./go
RUN mv go /usr/local
RUN export PATH="/usr/local/go/bin:$PATH"; \
	go version
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

# This will be required later for pulling in Go's Thrift library.
RUN apt-get install -y git

# Create and define the container's working directory.
RUN mkdir /theserver
WORKDIR /theserver
