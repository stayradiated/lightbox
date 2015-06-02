'use strict';

var React = require('react');

var flux     = require('../../flux');
var Lightbox = require('../../modules/lightbox');
var Sort     = require('./sort');
var Header   = require('../common/header/');
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
    var shows = this.filterShows(this.context.router.getCurrentParams().query);

    return (
      <div className='route-shows'>
        <Header />
        <ShowList shows={shows} />
      </div>
    );
  },

});

module.exports = Shows;
