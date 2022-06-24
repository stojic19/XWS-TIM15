FROM node:13.12.0-alpine

WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH

COPY ./dislinkt-frontend/package.json ./
COPY ./dislinkt-frontend/package-lock.json ./
RUN npm install --silent
RUN npm install react-scripts@3.4.1 -g --silent

COPY ./dislinkt-frontend ./

CMD ["npm", "run", "start"]
EXPOSE 3000