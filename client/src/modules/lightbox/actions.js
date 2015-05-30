'use strict';

var $ = require('jquery');
var flux = require('../../flux');
var actionTypes = require('./action-types');
var getters = require('./getters');

exports.searchSeries = function (query) {
  $.ajax({
    type: 'get',
    url: 'http://localhost:9000/api/series',
    data: {
      filter: query,
      limit: 300,
    },
  }).then(result => {
    flux.dispatch(actionTypes.SET_SERIES_LIST, result);
  });
};
