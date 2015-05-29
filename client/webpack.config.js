var path = require('path');

module.exports = {

  entry: {
    'bundle': './src/main.js',
  },

  output: {
    path: './',
    filename: 'dist/[name].js',
  },

  module: {
    loaders: [
      { test: /\.js$/, exclude: /node_modules/, loader: 'babel-loader' },
      // { test: /\.react.js/, loader: 'jsx-loader' },
      { test: /\.scss$/, loader: 'style-loader!css-loader!sass-loader' },
    ]
  }

}
