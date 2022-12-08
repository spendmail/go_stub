FROM golang:1.17

LABEL ORGANIZATION="OTUS"
LABEL SERVICE="stub"
LABEL MAINTAINERS="admin@stub.ru"

ARG LDFLAGS
ARG CONFIG_FILE_NAME

ENV BIN_FILE /opt/stub/app
ENV CODE_DIR /go/src/
ENV CONFIG_FILE /etc/stub/${CONFIG_FILE_NAME}.toml

WORKDIR ${CODE_DIR}

#RUN apt-get update && apt-get install -y \
#    libmagickwand-dev \
#    libgraphicsmagick-dev \
#    imagemagick \
#    graphicsmagick \
#    && apt-get clean \
#    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* \
#    && pkg-config --cflags --libs MagickWand \
#    && export CGO_CFLAGS_ALLOW='-Xpreprocessor'

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . ${CODE_DIR}

RUN go build \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE} ${CODE_DIR}/cmd/stub/

COPY ./configs/${CONFIG_FILE_NAME}.toml ${CONFIG_FILE}

CMD ${BIN_FILE} -config ${CONFIG_FILE}
