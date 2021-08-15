FROM node:14-alpine

RUN apk add --no-cache gcc g++ make python3

WORKDIR /app

COPY package.json yarn.lock ./

RUN yarn --frozen-lockfile

# Use npm modules with binaries
ENV PATH /opt/mypc/node_modules/.bin:$PATH

COPY . .

EXPOSE 3001
