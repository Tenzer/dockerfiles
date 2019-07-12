FROM node:10-alpine

ENV NPM_CONFIG_PREFIX=/home/node/.npm-global \
    PATH=$PATH:/home/node/.npm-global/bin

RUN apk add --no-cache tini \
    && apk add --no-cache --virtual .install-deps \
        g++ \
        make \
        python \
    && su node -c "npm install --global dynalite@2.3.2" \
    && su node -c "npm cache clean --force" \
    && apk del .install-deps \
    && mkdir /var/lib/dynalite \
    && chown node:node /var/lib/dynalite

USER node
EXPOSE 4567
VOLUME /var/lib/dynalite

ENTRYPOINT ["/sbin/tini", "--", "dynalite", "--path=/var/lib/dynalite"]
