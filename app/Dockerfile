FROM --platform=$BUILDPLATFORM node:14.3.0-alpine as build-stage

WORKDIR /go/gobel-admin-client-example/app

COPY . .

RUN npm install

