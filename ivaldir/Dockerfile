FROM node:16.3

WORKDIR /usr/app

COPY ./package*.json ./

RUN yarn

COPY ./ ./

RUN yarn build

EXPOSE 3000

USER node

CMD ["yarn", "start"]