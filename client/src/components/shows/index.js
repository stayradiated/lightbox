'use strict';

var React = require('react');

var flux     = require('../../flux');
var Lightbox = require('../../modules/lightbox');
var Sort     = require('./sort');
var ShowList = require('../common/showlist');

var Shows = React.createClass({
  mixins: [flux.ReactMixin],

  contextTypes: {
    router: React.PropTypes.func,
  },

  getDataBindings() {
    return {
      shows: Lightbox.getters.shows,
    };
  },

  filterShows(query) {
    var shows = this.state.shows;

    if (shows == null) {
      shows = [];
    }

    if (query == null) {
      return shows;
    }

    query = query.toLowerCase();
    return shows.filter(show => {
      return show.get('Title').toLowerCase().indexOf(query) >= 0;
    });
  },

  render() {
    var query = this.context.router.getCurrentParams().query;
    // var shows = this.filterShows(query).sortBy(show => 0 - show.get('Released').getTime());
    var shows = this.filterShows(query).sort((a, b) => {
      return a.get('Title').localeCompare(b.get('Title'));
    }).slice(0, 30);

    return (
      <div className='route-shows'>
        <ShowList shows={shows} />
      </div>
    );
  },

});

module.exports = Shows;
