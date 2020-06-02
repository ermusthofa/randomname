############################
# build executable binary
############################
FROM node:10 as build
WORKDIR /usr/src/app
COPY package.json yarn.lock ./
RUN yarn
COPY . ./
RUN yarn build

############################
# build service image
############################
FROM nginx:1.17.10-alpine
WORKDIR /usr/share/nginx/html
COPY --from=build /usr/src/app/build .
COPY ./nginx/default.conf .
RUN apk add --no-cache bash
EXPOSE 80
CMD ["/bin/bash", "-c", "/bin/cat default.conf | envsubst > /etc/nginx/conf.d/default.conf; nginx -g \"daemon off;\""]