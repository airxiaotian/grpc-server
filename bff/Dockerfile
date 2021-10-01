FROM node:14.3.0-alpine3.11 AS builder
WORKDIR /tmp
RUN apk add --no-cache python3 make g++
RUN wget -qO node-prune.tgz https://github.com/tj/node-prune/releases/download/v1.0.1/node-prune_1.0.1_linux_amd64.tar.gz && \
    tar xzfv node-prune.tgz && \
    mv node-prune /usr/local/bin/node-prune
WORKDIR /build
COPY package.json .
COPY yarn.lock .
RUN yarn install
COPY . .
RUN yarn build
RUN rm -rf node_modules
RUN yarn install --prod
RUN yarn cache clean
RUN node-prune

FROM node:14.3.0-alpine3.11
RUN apk add --no-cache tzdata ca-certificates
ENV NODE_ENV=production
ENV TZ=Asia/Tokyo
WORKDIR /usr/local/app
COPY --from=builder /build/node_modules ./node_modules
COPY --from=builder /build/package.json ./package.json
COPY --from=builder /build/.dist ./.dist
EXPOSE 3000
CMD ["yarn", "start"]
