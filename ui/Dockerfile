FROM node:16.13.1
WORKDIR /app
COPY package.json . 
RUN yarn install 
COPY . . 
RUN yarn global add vite 
EXPOSE 3000
CMD ["yarn", "run", "dev"]