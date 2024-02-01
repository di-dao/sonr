# sonrhq/chain: Earthfile
# ---------------------------------------------------------------------
VERSION 0.7
PROJECT sonrhq/testnet-1
FROM golang:1.21.5-alpine

# ---------------------------------------------------------------------

# deps - downloads dependencies
deps:
    RUN apk add --update --no-cache \
    bash \
    binutils \
    ca-certificates \
    coreutils \
    curl \
    findutils \
    g++ \
    git \
    grep \
    make \
    nodejs \
    npm \
    openssl \
    util-linux
    COPY go.mod go.sum ./
    RUN go mod download
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum

# -------------------
# [PROTOBUF Commands]
# -------------------


# init - generates private key for matrix homeserver
cosmoscan:
    FROM matrixdotorg/dendrite-monolith:latest
    ARG serverName=matrix.sonr.run
    ARG psqlConnURI=postgres://postgres:pwd@postgres:5432/postgres?sslmode=disable
    RUN /usr/bin/generate-keys -private-key matrix_key.pem
    SAVE ARTIFACT matrix_key.pem AS LOCAL matrix_key.pem
    RUN /usr/bin/generate-config -server $serverName -db $psqlConnURI > dendrite.yaml
    SAVE ARTIFACT dendrite.yaml AS LOCAL dendrite.yaml


# init - generates private key for matrix homeserver
faucet:
    FROM matrixdotorg/dendrite-monolith:latest
    ARG serverName=matrix.sonr.run
    ARG psqlConnURI=postgres://postgres:pwd@postgres:5432/postgres?sslmode=disable
    RUN /usr/bin/generate-keys -private-key matrix_key.pem
    SAVE ARTIFACT matrix_key.pem AS LOCAL matrix_key.pem
    RUN /usr/bin/generate-config -server $serverName -db $psqlConnURI > dendrite.yaml
    SAVE ARTIFACT dendrite.yaml AS LOCAL dendrite.yaml


# init - generates private key for matrix homeserver
ipfs:
    FROM matrixdotorg/dendrite-monolith:latest
    ARG serverName=matrix.sonr.run
    ARG psqlConnURI=postgres://postgres:pwd@postgres:5432/postgres?sslmode=disable
    RUN /usr/bin/generate-keys -private-key matrix_key.pem
    SAVE ARTIFACT matrix_key.pem AS LOCAL matrix_key.pem
    RUN /usr/bin/generate-config -server $serverName -db $psqlConnURI > dendrite.yaml
    SAVE ARTIFACT dendrite.yaml AS LOCAL dendrite.yaml


# build - builds and configures monolithic dendrite matrix homeserver
matrix:
    FROM matrixdotorg/dendrite-monolith:latest
    ARG serverName=matrix.sonr.run
    SAVE IMAGE --push sonrhq/dendrite:latest
