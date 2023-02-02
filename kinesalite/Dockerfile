FROM node:19-alpine

RUN apk add --no-cache tini && \
    mkdir /var/lib/kinesalite /kinesalite && \
    chown node:node /var/lib/kinesalite /kinesalite

USER node
WORKDIR /kinesalite/

COPY --chown=node:node package.json .
RUN --mount=type=cache,target=/home/node/.npm,uid=1000 \
    npm install

ENV PATH=$PATH:/kinesalite/node_modules/.bin

EXPOSE 4567
VOLUME /var/lib/kinesalite

RUN kinesalite --help

ENTRYPOINT ["/sbin/tini", "--", "kinesalite", "--path=/var/lib/kinesalite"]
