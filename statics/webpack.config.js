var path = require('path')
var webpack = require('webpack')
const HtmlWebpackPlugin = require('html-webpack-plugin')

module.exports = {
  entry: {
    main: ['./src/main.js']
  },
  output: {
    path: path.resolve(__dirname, './dist'),
    publicPath: '/dist/',
    filename: '[name].js'
  },
  module: {
    rules: [
      {
        test: /\.vue$/,
        use: {
          loader: 'vue-loader',
          options: {
            postcss() {
              return [require('autoprefixer')]
            }
          }
        }
      },
      {
        test: /\.js$/,
        use: 'babel-loader',
        exclude: [/node_modules/, /src\/lib/]
      },
      {
        test: /\.css$/,
        loader: 'style-loader!css-loader!postcss-loader'
      },
      {
        test: /\.(png|jpg|gif|svg)$/,
        use: 'url-loader',
      }
    ]
  },
  plugins: [
    new HtmlWebpackPlugin({
      chunks: ['main'],
      hash: true,
      template: path.resolve('./src/index.html'),
      filename: path.resolve('./dist/index.html'),
    }),
    new webpack.LoaderOptionsPlugin({
      minimize: true,
      options: {
        test: /\.(css|less|scss)$/,
        context: __dirname,
        postcss: [require('autoprefixer')]
      }
    })
  ],
  devtool: '#eval-source-map'
}

if (process.env.NODE_ENV === 'production') {
  module.exports.devtool = '#source-map'
  module.exports.plugins = (module.exports.plugins || []).concat([
    new webpack.DefinePlugin({
      'process.env': {
        NODE_ENV: '"production"'
      }
    }),
    new webpack.optimize.UglifyJsPlugin({
      sourceMap: true,
      compress: {
        warnings: false
      }
    })
  ])
}
