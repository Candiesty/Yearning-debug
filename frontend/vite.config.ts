import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import * as path from 'path';

// https://vitejs.dev/config/
export default defineConfig({
  base: '/front/',
  build: {
    minify: 'esbuild',
  },
  server: {
    proxy: {
      '/register': {
        target: 'http://127.0.0.1:8000',
      },
      '/fetch': {
        target: 'http://127.0.0.1:8000',
      },
      '/lang': {
        target: 'http://127.0.0.1:8000',
      },
      '/login': {
        target: 'http://127.0.0.1:8000',
      },
      '/api': {
        target: 'http://127.0.0.1:8000',
      },
      '/ldap': {
        target: 'http://127.0.0.1:8000',
      },
      '/downlaod/*': {
        target: 'http://127.0.0.1:8000',
      },
      '/oidc/state': {
        target: 'http://127.0.0.1:8000',
      },
    },
  },
  plugins: [vue()],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src/'),
    },
  },
  css: {
    preprocessorOptions: {
      less: {
        javascriptEnabled: true,
      },
    },
  },
});
