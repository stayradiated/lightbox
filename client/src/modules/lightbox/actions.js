'use strict';

var $ = require('jquery');
var flux = require('../../flux');
var actionTypes = require('./action-types');
var getters = require('./getters');

const baseUrl = './api';
const ext = '.json';

exports.fetchShows = function () {
  $.ajax({
    url: baseUrl + '/shows' + ext,
  }).then(result => {
    flux.dispatch(actionTypes.SetShows, result);
  });
};

exports.fetchShow = function (id) {
  if (flux.evaluate(getters.show).get('ID') == id) {
    return;
  }

  flux.dispatch(actionTypes.ResetShow);

  $.ajax({
    url: baseUrl + '/shows/' + id + ext,
  }).then(result => {
    flux.dispatch(actionTypes.SetShow, result);
  });
};

exports.fetchSeason = function (id) {
  if (flux.evaluate(getters.season).get('ID') == id) {
    return;
  }

  flux.dispatch(actionTypes.ResetSeason);

  $.ajax({
    url: baseUrl + '/seasons/' + id + ext,
  }).then(result => {
    flux.dispatch(actionTypes.SetSeason, result);
  });
};

exports.fetchCategory = function (id) {
  $.ajax({
    url: baseUrl + '/categories/' + id + ext,
  }).then(result => {
    flux.dispatch(actionTypes.SetShows, result);
  });
};

exports.fetchEpisode = function (id) {
  if (flux.evaluate(getters.episode).get('ID') == id) {
    return;
  }

  flux.dispatch(actionTypes.ResetEpisode);

  $.ajax({
    url: baseUrl + '/episodes/' + id + ext,
  }).then(result => {
    flux.dispatch(actionTypes.SetEpisode, result);
  });
};

exports.fetchCategories = function () {
  $.ajax({
    url: baseUrl + '/categories' + ext,
  }).then(result => {
    flux.dispatch(actionTypes.SetCategories, result);
  });
};

exports.bookmarkShow = function (showID) {
  flux.dispatch(actionTypes.BookmarkShow, showID);
};

exports.unbookmarkShow = function (showID) {
  flux.dispatch(actionTypes.UnbookmarkShow, showID);
};

exports.setWatchlist = function (data) {
  flux.dispatch(actionTypes.SetWatchlist, data);
};
