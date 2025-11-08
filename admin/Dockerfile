# Stage: development & build
FROM node:22-alpine

WORKDIR /app

# Salin package info dulu untuk caching layer
COPY package*.json ./

# Jika file kosong, skip install dulu (nanti setelah create)
RUN npm install 

COPY . .

# Command default: development server
CMD ["npm", "run", "dev"]
