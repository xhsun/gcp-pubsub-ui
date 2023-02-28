ARG NODE_VERSION=18.14-alpine
FROM node:$NODE_VERSION as build

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm ci

COPY . .

ARG PUBSUB_UI_BACKEND_URL=http://localhost:9090
RUN sed -i "s|PUBSUB_UI_BACKEND_URL|$PUBSUB_UI_BACKEND_URL|" src/environments/environment.ts

RUN npm run build:prod

FROM nginx:stable-alpine

COPY --from=build /app/dist/pubsub-ui-client /usr/share/nginx/html
