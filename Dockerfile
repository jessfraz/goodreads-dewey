FROM golang:alpine as builder
MAINTAINER Jessica Frazelle <jess@linux.com>

ENV PATH /go/bin:/usr/local/go/bin:$PATH
ENV GOPATH /go

RUN	apk add --no-cache \
	bash \
	ca-certificates

COPY . /go/src/github.com/jessfraz/goodreads-dewey

RUN set -x \
	&& apk add --no-cache --virtual .build-deps \
		git \
		gcc \
		libc-dev \
		libgcc \
		make \
	&& cd /go/src/github.com/jessfraz/goodreads-dewey \
	&& make static \
	&& mv goodreads-dewey /usr/bin/goodreads-dewey \
	&& apk del .build-deps \
	&& rm -rf /go \
	&& echo "Build complete."

FROM alpine:latest

COPY --from=builder /usr/bin/goodreads-dewey /usr/bin/goodreads-dewey
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs

ENTRYPOINT [ "goodreads-dewey" ]
CMD [ "--help" ]
