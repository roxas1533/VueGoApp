module.exports = {
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