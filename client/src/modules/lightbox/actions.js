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
      limit: 40,
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

exports.viewCategory = function (id) {
  $.ajax({
    type: 'get',
    url: baseUrl + '/categories/' + id,
  }).then(result => {
    flux.dispatch(actionTypes.SetShows, result);
  });
};

exports.viewEpisode = function (id) {
  $.ajax({
    type: 'get',
    url: baseUrl + '/episodes/' + id,
  }).then(result => {
    flux.dispatch(actionTypes.SetEpisode, result);
  });
};
