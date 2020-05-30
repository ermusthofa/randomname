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
COPY ./env.sh .
COPY ./.env .
RUN apk add --no-cache bash
RUN chmod +x env.sh
EXPOSE 80
CMD ["/bin/bash", "-c", "/usr/share/nginx/html/env.sh && nginx -g \"daemon off;\""]