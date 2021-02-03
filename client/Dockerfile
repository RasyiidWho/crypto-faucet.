FROM node:14
RUN mkdir /app
ADD . /app
WORKDIR /app
COPY /package*.json /app/
RUN npm install
COPY . /app
EXPOSE 3000
CMD ["npm","start"]