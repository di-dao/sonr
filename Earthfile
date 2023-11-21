VERSION 0.7
PROJECT sonrhq/sonr-testnet-0

faucet:
    FROM node:18.7-alpine
    ARG cosmosjsVersion=0.28.11
    RUN npm install @cosmjs/faucet@$cosmosjsVersion --global --production
    COPY ./scripts/start-faucet.sh /usr/local/bin/faucet
    RUN chmod +x /usr/local/bin/faucet
    EXPOSE 4500
    ENTRYPOINT ["faucet"]
    SAVE IMAGE sonrhq/faucet:latest
