# Base image untuk build
FROM node:22-alpine 

WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

# Expose port 3000
EXPOSE 5173

# Jalankan nginx
CMD ["npm", "run", "dev"]
