FROM node:10-alpine

ENV NPM_CONFIG_PREFIX=/home/node/.npm-global \
    PATH=$PATH:/home/node/.npm-global/bin

RUN apk add --no-cache tini \
    && apk add --no-cache --virtual .install-deps \
        g++ \
        make \
        python \
    && su node -c "npm install --global kinesalite@3.1.0" \
    && su node -c "npm cache clean --force" \
    && apk del .install-deps \
    && mkdir /var/lib/kinesalite \
    && chown node:node /var/lib/kinesalite

USER node
EXPOSE 4567
VOLUME /var/lib/kinesalite

ENTRYPOINT ["/sbin/tini", "--", "kinesalite", "--path=/var/lib/kinesalite"]
