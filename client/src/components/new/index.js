'use strict';

var React = require('react');

var flux = require('../../flux');
var Lightbox = require('../../modules/lightbox');
var ShowList = require('../common/showlist');

var NewReleases = React.createClass({
  mixins: [flux.ReactMixin],

  getDataBindings() {
    return {
      shows: Lightbox.getters.shows,
    };
  },

	render() {
    var shows = this.state.shows.sortBy(show => {
      return 0 - show.get('Released').getTime();
    }).slice(0, 30);

    return (
      <div className='route-new'>
        <h1>New Releases</h1>
        <ShowList shows={shows} />
      </div>
    );
	},

});

module.exports = NewReleases;
