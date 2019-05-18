/* diabled-eslint */
module.exports = {
  // vue 项目静态文件的根路径
  publicPath: process.env.NODE_ENV === 'production'
    ? '/admin'
    : '/',
  // server 配置
  devServer: {
    host: '0.0.0.0',
    port: 8081,
    // open: true
    // proxy: {}
  },
  chainWebpack: config => {
    // 定义全局变量，项目中可通过 process.env 获取到
    config.plugin('define').tap(definitions => {
      Object.assign(definitions[0]['process.env'], {
        build_timestamp: (new Date()).getTime(),
        version: 1,
      });
      return definitions;
    });
  }
}
