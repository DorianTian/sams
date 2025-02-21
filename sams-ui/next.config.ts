import type { NextConfig } from 'next'

const nextConfig: NextConfig = {
  /* config options here */
  rewrites: async () => {
    return [
      {
        source: '/:path*',
        destination: 'http://192.168.3.32:8080/:path*', // 目标服务器地址
      },
    ]
  },
}

export default nextConfig
