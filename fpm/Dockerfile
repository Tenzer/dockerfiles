FROM ruby:2.3
MAINTAINER Jeppe Toustrup <docker@tenzer.dk>

RUN apt-get update \
	&& apt-get install -y npm php-pear python-setuptools rpm cpio \
	&& rm -rf /var/lib/apt/lists/*

RUN gem install fpm

ENTRYPOINT [ "fpm" ]
CMD [ "--help" ]
