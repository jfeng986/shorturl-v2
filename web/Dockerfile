
FROM node:18 as build-stage

WORKDIR /app
COPY ./nginx.conf /nginx.conf
COPY package*.json /app/
RUN npm install

COPY ./ /app/
RUN npm run build


FROM nginx:alpine
EXPOSE 80
COPY --from=build-stage /app/dist/ /usr/share/nginx/html
COPY --from=build-stage /nginx.conf /etc/nginx/conf.d/default.conf
