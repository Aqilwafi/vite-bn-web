import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [react()],
  server: {
    watch: {
      usePolling: true,  // <-- ini kunci agar hot reload jalan di Docker
    },
    port: 5173,
    strictPort: true,
    host: true, // ini penting untuk Docker
    proxy: {
      '/api': {
        target: 'http://localhost:4000', 
        changeOrigin: true,
        secure: false,
      },
    },
  },
})
