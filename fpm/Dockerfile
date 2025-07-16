# syntax=docker/dockerfile:1
FROM ruby:3.4.5-slim AS no-entrypoint

SHELL ["/bin/bash", "-o", "pipefail", "-c"]

RUN <<EOF
    apt-get update
    apt-get install -y --no-install-recommends \
        binutils \
        cpanminus \
        cpio \
        cpp \
        curl \
        gcc \
        libarchive-tools \
        make \
        php-pear \
        python3-pip \
        python3-virtualenv \
        python-is-python3 \
        rpm \
        squashfs-tools \
        unzip \
        xz-utils \
        zip \
        zstd
    apt-get clean
    rm --recursive /var/lib/apt/lists/*
EOF

RUN --mount=type=cache,target=/root/.cache/pip \
    pip install --break-system-packages \
        virtualenv-tools3

ARG NODEJS_VERSION=20.12.2
ARG TARGETPLATFORM
RUN <<EOF
    arch=x64
    [ "$TARGETPLATFORM" == "linux/arm64" ] && arch=arm64
    mkdir --parents /usr/local/node
    curl --silent "https://nodejs.org/dist/v${NODEJS_VERSION}/node-v${NODEJS_VERSION}-linux-${arch}.tar.xz" | \
    tar --extract --xz --directory /usr/local/node --strip-components 1 --file -
EOF

ENV PATH="$PATH:/usr/local/node/bin"

COPY Gemfile /tmp/Gemfile
RUN --mount=type=cache,target=/root/.bundle/cache \
    bundler install --gemfile=/tmp/Gemfile

RUN --mount=type=cache,target=/root/.cache/pip \
    --mount=type=cache,target=/root/.cpanm \
    <<EOF
    mkdir /tmp/tests
    pushd /tmp/tests || exit
    echo 'This is a test' > package-data

    # Dir to * tests
    # dir -> dir
    fpm --input-type=dir --output-type=dir --name=test package-data
    # dir -> deb
    fpm --input-type=dir --output-type=deb --name=test --architecture=all package-data
    # dir -> rpm
    fpm --input-type=dir --output-type=rpm --name=test --architecture=all package-data
    # dir -> tar
    fpm --input-type=dir --output-type=tar --name=test package-data
    # dir -> empty
    fpm --input-type=dir --output-type=empty --name=test package-data
    # dir -> freebsd
    fpm --input-type=dir --output-type=freebsd --name=test package-data
    # dir -> apk
    fpm --input-type=dir --output-type=apk --name=test --architecture=all package-data
    # dir -> snap
    fpm --input-type=dir --output-type=snap --name=test --architecture=all package-data
    # dir -> pacman
    fpm --input-type=dir --output-type=pacman --name=test --architecture=all package-data
    # dir -> sh
    fpm --input-type=dir --output-type=sh --name=test package-data
    # dir -> zip
    fpm --input-type=dir --output-type=zip --name=test package-data

    # * to dir tests
    # gem -> dir
    fpm --input-type=gem --output-type=dir test
    # deb -> dir
    fpm --input-type=deb --output-type=dir --package=deb test_1.0_all.deb
    # npm -> dir
    fpm --input-type=npm --output-type=dir is-number
    # rpm -> dir
    fpm --input-type=rpm --output-type=dir --package=rpm test-1.0-1.noarch.rpm
    # tar -> dir
    fpm --input-type=tar --output-type=dir --package=tar test.tar
    # cpan -> dir
    fpm --input-type=cpan --output-type=dir Text::Lorem
    # pear -> dir
    fpm --input-type=pear --output-type=dir Text_Statistics
    # empty -> dir
    fpm --input-type=empty --output-type=dir --name=empty
    # python -> dir
    fpm --input-type=python --output-type=dir cowsay
    # snap -> dir
    fpm --input-type=snap --output-type=dir --package=snap test_1.0_all.snap
    # pacman -> dir
    fpm --input-type=pacman --output-type=dir --package=pacman test-1.0-1-any.pkg.tar.zst
    # pleaserun -> dir
    fpm --input-type=pleaserun --output-type=dir --name=pleaserun-test /usr/bin/echo
    # virtualenv -> dir
    fpm --input-type=virtualenv --output-type=dir cowsay
    # zip -> dir
    fpm --input-type=zip --output-type=dir --package=zip test.zip

    # Cleanup
    popd || exit
    rm -r /tmp/tests
EOF

CMD ["fpm", "--help"]


FROM no-entrypoint AS with-entrypoint

ENTRYPOINT ["fpm"]
CMD ["--help"]
