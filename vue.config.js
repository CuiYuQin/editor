module.exports = {
    publicPath : process.env.NODE_ENV === 'production' ? './' : '/',
    devServer: {
      proxy: {
        "/api": {
          target: 'http://localhost:7000', // 前端域名
          changeOrigin: true, // 如果接口跨域 ， 需要进行这个参数配置
          pathRewrite: {
            "^/api": ""
          },
        },
      },
    },
  };