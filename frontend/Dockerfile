# syntax=docker/dockerfile:1.3-labs

# build stage
FROM node:lts-alpine as build-stage

ARG VUE_APP_COOKBOOK_AUTHORITY
ARG VUE_APP_COOKBOOK_CLIENTID
ARG VUE_APP_COOKBOOK_BACKEND_URL

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

# production stage
FROM nginx:stable-alpine as production-stage

COPY --from=build-stage /app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
