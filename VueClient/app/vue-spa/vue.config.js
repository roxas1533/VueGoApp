module.exports = {
  devServer: {
    host: '0.0.0.0',
    disableHostCheck: true
  },
  configureWebpack: {
    module: {
      rules: [
        {
          enforce: 'pre',
          test: /\.(jsx?|tsx?|vue)$/,
          exclude: /node_modules/,
          loader: 'eslint-loader',
          options: {
            fix: true
          }
        }
      ]
    }
  }
}