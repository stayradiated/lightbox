'use strict';

var path = require('path');
var ExtractTextPlugin = require('extract-text-webpack-plugin');

function getEntrySources(sources) {
  if (process.env.NODE_ENV !== 'production') {
    sources.push('webpack-dev-server/client?http://192.168.1.100:8080');
    sources.push('webpack/hot/only-dev-server');
  }

  return sources;
}

function handleHotSass(config) {
  if (process.env.NODE_ENV === 'production') {

    config.plugins = [
      new ExtractTextPlugin('dist/style.css', {
        allChunks: true,
      }),
    ];

    config.module.loaders.push({
      test: /\.scss$/,
      loader: ExtractTextPlugin.extract('css!autoprefixer!sass'),
    });

  } else {

    config.module.loaders.push({
      test: /\.scss$/,
      loaders: ['style', 'css', 'autoprefixer', 'sass']
    });
    config.module.loaders.push({
      text: /\.(png|woff|woff2|eot|ttf|svg)$/,
      loader: 'url',
    });

  }

  return config;
}

module.exports = {

  entry: {
    bundle: getEntrySources([
      './src/main.js',
    ]),
  },

  output: {
    path: '.dist/',
    filename: '[name].js',
    publicPath: 'http://192.168.1.100:8080/',
  },

  module: {
    loaders: [
      {
        test: /\.js$/,
        exclude: /node_modules/,
        loaders: ['react-hot', 'jsx', 'babel'],
      },
      {
        test: /\.scss$/,
        loaders: ['style', 'css', 'autoprefixer', 'sass'],
      },
      {
        test: /\.(woff|eot|svg|ttf)$/,
        loaders: ['url?limit=8192'],
      },
    ]
  },

};
