'use strict';

var path = require('path');
var ExtractTextPlugin = require('extract-text-webpack-plugin');

function getEntrySources(sources) {
  if (process.env.NODE_ENV !== 'production') {
    sources.push('webpack-dev-server/client?http://localhost:8080');
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
      loader: ExtractTextPlugin.extract('css!sass'),
    });

  } else {

    config.module.loaders.push({
      test: /\.scss$/,
      loaders: ['style', 'css', 'sass']
    });

  }

  return config;
}

module.exports = handleHotSass({

  entry: {
    bundle: getEntrySources([
      './src/main.js',
    ]),
  },

  output: {
    path: './',
    filename: 'dist/[name].js',
    publicPath: 'http://localhost:8080/',
  },

  module: {
    loaders: [
      {
        test: /\.js$/,
        exclude: /node_modules/,
        loader: 'react-hot!jsx!babel',
      },
    ]
  },

});
