# Fetching the latest node image on alpine linux
FROM node:slim AS development

# Declaring env
ENV NODE_ENV development

# Setting up the work directory
WORKDIR /ui

# Installing dependencies
COPY ./package.json /ui
RUN yarn install

# Copying all the files in our project
COPY . .

# Starting our application
CMD yarn start