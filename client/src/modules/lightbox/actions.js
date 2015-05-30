'use strict';

var $ = require('jquery');
var flux = require('../../flux');
var actionTypes = require('./action-types');
var getters = require('./getters');

const baseUrl = 'http://localhost:9000/api';

exports.searchShows = function (query) {
  $.ajax({
    type: 'get',
    url: baseUrl + '/shows',
    data: {
      filter: query,
      limit: 300,
    },
  }).then(result => {
    flux.dispatch(actionTypes.SetShows, result);
  });
};

exports.viewShow = function (id) {
  $.ajax({
    type: 'get',
    url: baseUrl + '/shows/' + id,
  }).then(result => {
    flux.dispatch(actionTypes.SetShow, result);
  });
};

exports.viewSeason = function (id) {
  $.ajax({
    type: 'get',
    url: baseUrl + '/seasons/' + id,
  }).then(result => {
    flux.dispatch(actionTypes.SetSeason, result);
  });
};
