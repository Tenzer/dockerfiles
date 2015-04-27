FROM ruby:latest
MAINTAINER Jeppe Toustrup <docker@tenzer.dk>

RUN apt-get update \
	&& apt-get install -y npm php-pear python-setuptools rpm \
	&& rm -rf /var/lib/apt/lists/*

RUN gem install fpm

CMD [ "fpm" ]
