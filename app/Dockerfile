FROM node:14.3.0-alpine as build-stage

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY ./ .

RUN npm run local-build

FROM nginx:1.19.0-alpine

COPY ./nginx/nginx.conf /etc/nginx/nginx.conf
COPY ./nginx/conf.d/gobel-admin-client.conf /etc/nginx/conf.d/goble-admin-client-example.conf
COPY --from=build-stage /app/dist /var/www/html