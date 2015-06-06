'use strict';

var React = require('react');

var flux     = require('../../flux');
var Lightbox = require('../../modules/lightbox');
var ShowRow  = require('../common/showrow/');

var List = React.createClass({
  mixins: [flux.ReactMixin],

  propTypes: {
    listID: React.PropTypes.number.isRequired,
    showTitle: React.PropTypes.bool,
  },

  getDataBindings() {
    return {
      lists: Lightbox.getters.lists,
      shows: Lightbox.getters.shows,
    };
  },

	render() {
    var listID = this.props.listID;
    var list = this.state.lists.find(list => {
      return list.get('ID') === listID;
    });

    if (list == null) {
      return null;
    }

    var shows = this.state.shows;
    var listShows = shows.filter(show => {
      return list.get('Shows').contains(show.get('ID'));
    }).sortBy(show => {
      return list.get('Shows').indexOf(show.get('ID'));
    });

    return (
      <section>
        { this.props.showTitle ? (
          <h2>{list.get('Title')}</h2>
        ) : null }
        <ShowRow shows={listShows} />
      </section>
    );
	},

});

module.exports = List;
