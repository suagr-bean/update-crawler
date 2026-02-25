
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { resolve } from 'path'

import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { VantResolver } from '@vant/auto-import-resolver'
// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue(),
    AutoImport({
      resolvers:[VantResolver()],
    }),
   Components({
     resolvers:[VantResolver()],
   }),
  ],
  resolve: {
    alias: {
      '@': resolve(__dirname, 'src')
    }
  },
  server:{
    proxy:{
      '/api':{
        target:'http://localhost:8080',
        changeOrigin:true,
        rewrite: (path) => path.replace(/^\/api/, '')
      }
    }
  }
})
