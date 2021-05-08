FROM node:14

ENV BOT_TOKEN=default \
    K8_URI=http://localhost:8080

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm install

RUN npm ci --only=production

COPY . .

CMD [ "node", "app.js" ]
