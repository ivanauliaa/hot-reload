FROM node:alpine3.14

RUN mkdir server-backend

WORKDIR /server-backend

RUN npm install -g localtunnel

CMD lt --port 9000 --local-host backend | tee output.txt