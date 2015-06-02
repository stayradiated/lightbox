'use strict';

var $ = require('jquery');
var flux = require('../../flux');
var actionTypes = require('./action-types');
var getters = require('./getters');

const baseUrl = 'http://192.168.1.100:9000/api';

exports.fetchShows = function () {
  // flux.dispatch(actionTypes.SetShows, {});
  $.ajax({
    url: baseUrl + '/shows',
  }).then(result => {
    flux.dispatch(actionTypes.SetShows, result);
  });
};

exports.fetchShow = function (id) {
  // flux.dispatch(actionTypes.SetShow, {});
  $.ajax({
    url: baseUrl + '/shows/' + id,
  }).then(result => {
    flux.dispatch(actionTypes.SetShow, result);
  });
};

exports.fetchSeason = function (id) {
  // flux.dispatch(actionTypes.SetSeason, {});
  $.ajax({
    url: baseUrl + '/seasons/' + id,
  }).then(result => {
    flux.dispatch(actionTypes.SetSeason, result);
  });
};

exports.fetchCategory = function (id) {
  // flux.dispatch(actionTypes.SetShows, []);
  $.ajax({
    url: baseUrl + '/categories/' + id,
  }).then(result => {
    flux.dispatch(actionTypes.SetShows, result);
  });
};

exports.fetchEpisode = function (id) {
  // flux.dispatch(actionTypes.SetEpisode, {});
  $.ajax({
    url: baseUrl + '/episodes/' + id,
  }).then(result => {
    flux.dispatch(actionTypes.SetEpisode, result);
  });
};

exports.fetchCategories = function () {
  $.ajax({
    url: baseUrl + '/categories',
  }).then(result => {
    flux.dispatch(actionTypes.SetCategories, result);
  });
};

exports.bookmarkShow = function (showID) {
  flux.dispatch(actionTypes.BookmarkShow, showID);
};
