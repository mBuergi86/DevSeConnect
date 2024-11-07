# Stage 1: Build Stage
FROM node:23.1.0-alpine AS build

RUN npm install -g pnpm@latest

WORKDIR /app

COPY pnpm-lock.yaml ./
RUN pnpm fetch --dev

COPY . ./
RUN pnpm install --offline --dev

COPY .env .env

RUN pnpm run build

EXPOSE 4000

CMD ["sh", "-c", "PORT=4000 node build"]
