FROM node:8.1 

WORKDIR /app
COPY app . 

RUN npm install

EXPOSE 9200 

CMD ["npm", "start"]
