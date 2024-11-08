FROM alpine:3.20.0

ENV PERCONA_VERSION=3.0.11
ADD https://www.percona.com/downloads/percona-toolkit/${PERCONA_VERSION}/binary/tarball/percona-toolkit-${PERCONA_VERSION}_x86_64.tar.gz /
RUN apk --update upgrade && \
    apk add curl ca-certificates tzdata perl perl-dbi perl-dbd-mysql && \
    apk add --virtual=build make && \
    tar zxf /percona-toolkit-${PERCONA_VERSION}_x86_64.tar.gz && \
    ( \
      cd percona-toolkit-* && \
      perl Makefile.PL && \
      make && \
      make install \
    ) && \
    update-ca-certificates && \
    find /usr/local/bin/ -name 'pt-*'  ! -name 'pt-online-schema-change' -delete &&\
    apk del make &&\
    rm -rf percona-toolkit* && \
    rm -rf /var/cache/apk/*

# copy binary into image
COPY sql-migrate /bin/sql-migrate