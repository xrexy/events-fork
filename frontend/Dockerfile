# BUILD STAGE
FROM node:18.12.1-alpine3.17 AS build
WORKDIR /app
COPY package.json yarn.lock ./
RUN yarn install
COPY . .
RUN yarn build

ARG API_HOST=http://localhost:8090
ENV VITE_API_HOST $API_HOST

# RUN STAGE
FROM nginx:1.23.3-alpine as prod-stage
COPY --from=build /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
