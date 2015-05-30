'use strict';

var $ = require('jquery');
var flux = require('../../flux');
var actionTypes = require('./action-types');
var getters = require('./getters');

const baseUrl = 'http://localhost:9000/api/';

exports.searchSeries = function (query) {
  $.ajax({
    type: 'get',
    url: baseUrl + '/series',
    data: {
      filter: query,
      limit: 300,
    },
  }).then(result => {
    flux.dispatch(actionTypes.SET_SERIES_LIST, result);
  });
};

exports.viewSeries = function (id) {
  $.ajax({
    type: 'get',
    url: baseUrl + '/series/' + id,
  }).then(result => {
    console.log(result);
  });
};
