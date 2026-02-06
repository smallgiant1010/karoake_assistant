FROM node:25-alpine3.22 AS build

WORKDIR /app

COPY package.json package-lock.json ./

RUN npm install

COPY . .

RUN npm run build

FROM nginx:alpine

COPY --from=build /app/prod /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
