# Stage 1: Build the application
FROM node:18 AS builder

# Set the working directory for the build stage
WORKDIR /usr/src/app

# Copy package.json and package-lock.json to install dependencies
COPY package*.json ./

# Install project dependencies
RUN npm install

# Copy the entire project source
COPY . .

# Build the application
RUN npm run build

# Stage 2: Create a production-ready image
FROM node:18 AS production

# Set the working directory for the production stage
WORKDIR /usr/src/app

# Copy the built application from the builder stage
COPY --from=builder /usr/src/app/dist ./dist

# Copy package.json and package-lock.json for production install
COPY package*.json ./

# Install only production dependencies
RUN npm install --only=production

# Remove unnecessary files
RUN rm package*.json

# Expose the port your application will run on
EXPOSE 3030

# Start your application
CMD ["node", "dist/main.js"]
