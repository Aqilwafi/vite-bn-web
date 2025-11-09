# =========================
# Stage 1: Build frontend
# =========================
FROM node:22-alpine AS build

WORKDIR /app

# Copy package dan install dependencies
COPY package*.json ./
RUN npm install

# Copy semua source code
COPY . .

# Build project (Vite akan menghasilkan folder 'dist')
RUN npm run build

# =========================
# Stage 2: Serve dengan Nginx
# =========================
FROM nginx:alpine

# Hapus default nginx content
RUN rm -rf /usr/share/nginx/html/*

# Tambahkan custom nginx.conf
COPY ./nginx.conf /etc/nginx/conf.d/default.conf


# Copy hasil build dari stage sebelumnya
COPY --from=build /app/dist /usr/share/nginx/html

# Copy konfigurasi nginx custom (opsional)
# COPY ./nginx.conf /etc/nginx/conf.d/default.conf

# Expose port 80
EXPOSE 80

# Jalankan Nginx
CMD ["nginx", "-g", "daemon off;"]
