module.exports = {
  productionSourceMap: false,
  configureWebpack: {
    devServer: {
      'headers': {
        "Access-Control-Allow-Origin": "*",
        "Access-Control-Allow-Methods": "GET, POST, PUT, DELETE",
        "Access-Control-Allow-Headers": "content-type, Authorization"
      }
    }
  },
}
