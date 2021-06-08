FROM node:14.3.0-alpine as build-stage

WORKDIR /go/gobel-admin-client-example/app

COPY . .

ARG VUE_APP_API_ENDPOINT
ENV VUE_APP_API_ENDPOINT=${VUE_APP_API_ENDPOINT}

RUN npm install && \
	npm run local-build

FROM nginx:1.19.0-alpine

COPY ./nginx/nginx.conf /etc/nginx/nginx.conf
COPY ./nginx/conf.d/gobel-admin-client-example.conf /etc/nginx/conf.d/gobel-admin-client-example.conf
COPY --from=build-stage /go/gobel-admin-client-example/app/dist /var/www/html