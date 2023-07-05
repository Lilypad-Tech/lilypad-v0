FROM --platform=linux/x86-64 node:18

WORKDIR /usr/src/app

COPY package*.json ./
RUN npm install
COPY . .

CMD [ "node", "./index.js" ]
