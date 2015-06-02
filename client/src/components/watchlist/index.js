'use strict';

var React = require('react');

var flux = require('../../flux');
var Lightbox = require('../../modules/lightbox');
var ShowList = require('../common/showlist');

var WatchList = React.createClass({
  mixins: [flux.ReactMixin],

  getDataBindings() {
    return {
      shows: Lightbox.getters.watchlistShows,
    };
  },

	render() {
    var shows = this.state.shows;

    return (
      <div className='route-watchlist'>
        <h1>Watch List</h1>
        <ShowList shows={shows} />
      </div>
    );
	},

});

module.exports = WatchList;
