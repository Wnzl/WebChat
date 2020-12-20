FROM node:15.4.0-alpine

WORKDIR /frontend

COPY package.json package-lock.json ./

RUN npm install
RUN npm install react-scripts@3.4.1 -g

COPY . ./

CMD ["npm", "start"]

EXPOSE 3000