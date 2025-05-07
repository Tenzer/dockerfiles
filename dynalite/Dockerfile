# syntax=docker/dockerfile:1
FROM node:24-alpine

RUN <<EOF
    apk add --no-cache tini
    mkdir /var/lib/dynalite /dynalite
    chown node:node /var/lib/dynalite /dynalite
EOF

USER node
WORKDIR /dynalite/

COPY --chown=node:node package.json .
RUN --mount=type=cache,target=/home/node/.npm,uid=1000 \
    npm install

ENV PATH=$PATH:/dynalite/node_modules/.bin

EXPOSE 4567
VOLUME /var/lib/dynalite

RUN dynalite --help

ENTRYPOINT ["/sbin/tini", "--", "dynalite", "--path=/var/lib/dynalite"]
