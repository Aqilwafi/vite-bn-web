import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

export default defineConfig(({ mode }) => ({
  plugins: [react()],
  server: {
    watch: {
      usePolling: true, // tetap untuk Docker
    },
    port: 5173,
    strictPort: true,
    host: true,
  },
}));
