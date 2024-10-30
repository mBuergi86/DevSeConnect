# Stage 1: Build Stage
FROM node:lts-alpine AS build

# Set working directory inside the container
WORKDIR /app

# Install pnpm globally
RUN npm install -g pnpm

# Copy package.json and pnpm-lock.yaml
COPY package.json pnpm-lock.yaml ./

# Install dependencies using pnpm
RUN pnpm install --frozen-lockfile --dev

# Copy the rest of the application files
COPY . .

# Build the application
RUN pnpm run build

# Stage 2: Production Stage
FROM node:lts-alpine AS production

# Set working directory for the production container
WORKDIR /app

# Install pnpm globally
RUN npm install -g pnpm

# Copy only the built application from the build stage
COPY --from=build /app/build /app

# Install only production dependencies
COPY package.json pnpm-lock.yaml ./
RUN pnpm install --prod

# Expose the port the app will run on
EXPOSE 3000

# Command to run the application
CMD ["node", "build"]
