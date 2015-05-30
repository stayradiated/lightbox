'use strict';

var React = require('react');

var flux     = require('../../flux');
var Lightbox = require('../../modules/lightbox');
var Show     = require('./show.react.js');

var Shows = React.createClass({
  mixins: [flux.ReactMixin],

  getDataBindings() {
    return {
      shows: Lightbox.getters.shows,
    };
  },

  render() {
    var shows = this.state.shows;

    if (shows == null) {
      shows = [];
    }

    var showElements = shows.map(show => {
      return (
        <Show
          key={show.get('ID')}
          show={show}
        />
      );
    });

    return (
      <div className='route-shows'>
        <header>
          All TV
          Most Popular
        </header>
        <div className='list'>
          {showElements}
        </div>
      </div>
    );
  },

});

module.exports = Shows;
