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

    console.log(shows.toJS());

    return (
      <div className='route-watchlist'>
        <ShowList shows={shows} />
      </div>
    );
	},

});

module.exports = WatchList;
