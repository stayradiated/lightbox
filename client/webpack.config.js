'use strict';

var webpack = require('webpack');
var path = require('path');
var ExtractTextPlugin = require('extract-text-webpack-plugin');

module.exports = {

  entry: {
    bundle: [
      './src/main.js',
    ],
  },

  output: {
    path: '../gh-pages/',
    filename: '[name].js',
    publicPath: 'http://localhost:8080/',
  },

  module: {
    loaders: [
      {
        test: /\.js$/,
        exclude: /node_modules/,
        loaders: ['react-hot', 'jsx', 'babel'],
      },
      {
        test: /\.(woff|eot|svg|ttf|png)$/,
        loaders: ['url?limit=8192'],
      },
    ]
  },

};

if (process.env.NODE_ENV === 'production') {

  module.exports.output.publicPath = 'http://lightbox.mintco.de/';

  module.exports.plugins = [
    new ExtractTextPlugin('style.css', {
      allChunks: true,
    }),
    new webpack.optimize.UglifyJsPlugin({
      output: {comments: false},
    })
  ];

  module.exports.module.loaders.push({
    test: /\.scss$/,
    loader: ExtractTextPlugin.extract('css!autoprefixer!sass'),
  });

} else {

  var entryBundle = module.exports.entry.bundle;
  entryBundle.push('webpack-dev-server/client?http://localhost:8080');
  entryBundle.push('webpack/hot/only-dev-server');

  module.exports.module.loaders.push({
    test: /\.scss$/,
    loaders: ['style', 'css', 'autoprefixer', 'sass'],
  });

}
