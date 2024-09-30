FROM node:18-alpine
WORKDIR /cutshort-url/
COPY /public /cutshort-url/public
COPY /src /cutshort-url/src
COPY package.json package-lock.json .env ./
RUN npm install
EXPOSE 3000
CMD [ "npm", "start" ]
