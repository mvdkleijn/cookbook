# syntax=docker/dockerfile:1.3-labs

# build stage
FROM node:lts-alpine as build-stage

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

# production stage
FROM nginx:stable-alpine as production-stage

COPY --from=build-stage /app/dist /usr/share/nginx/html

RUN <<EOR
  /bin/cat <<'EOF' > /etc/nginx/conf.d/cookbook.conf
  server {
    listen       80;
    server_name  localhost;
    location / {
      root   /usr/share/nginx/html;
      index  index.html;
      try_files $uri $uri/ /index.html;
    }
    error_page   500 502 503 504  /50x.html;
    location = /50x.html {
      root   /usr/share/nginx/html;
    }
  }
EOR

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
