FROM node:18-alpine

ENV NPM_CONFIG_PREFIX=/home/node/.npm-global \
    PATH=$PATH:/home/node/.npm-global/bin

RUN apk add --no-cache tini \
    && su node -c "npm install --global dynalite@3.2.2" \
    && su node -c "npm cache clean --force" \
    && mkdir /var/lib/dynalite \
    && chown node:node /var/lib/dynalite

USER node
EXPOSE 4567
VOLUME /var/lib/dynalite

RUN dynalite --help

ENTRYPOINT ["/sbin/tini", "--", "dynalite", "--path=/var/lib/dynalite"]
