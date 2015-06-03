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
      category: Lightbox.getters.category,
    };
  },

  filterShows(query) {
    var shows = this.state.shows;
    var category = this.state.category;

    console.log('QUERY', query);
    console.log('SHOWS', shows.toJS());
    console.log('CATEGORY', category.toJS());

    if (shows == null) {
      shows = [];
    }

    if (query == null) {
      return shows;
    }

    query = query.toLowerCase();

    return shows.filter(show => {
      return
        show.get('Title').toLowerCase().indexOf(query) >= 0
      &&
        show.get('Categories').contains(category.get('ID'));
    });
  },

  render() {
    var query = this.context.router.getCurrentParams().query;

    // var shows = this.filterShows(query);
    // .sortBy(show => 0 - show.get('Rating'));

    var shows = this.filterShows(query);
    // .sort((a, b) => {
    //   return a.get('Title').localeCompare(b.get('Title'));
    // }).slice(0, 30);

    return (
      <div className='route-shows'>
        <ShowList shows={shows} />
      </div>
    );
  },

});

module.exports = Shows;
