FROM node:14-alpine

WORKDIR /app

COPY package.json yarn.lock ./

RUN yarn --frozen-lockfile

# Use npm modules with binaries
ENV PATH /opt/mypc/node_modules/.bin:$PATH

COPY . .

EXPOSE 3001
