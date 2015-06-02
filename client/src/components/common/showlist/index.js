'use strict';

var React = require('react');

var Show = require('./show');

var ShowList = React.createClass({

  propTypes: {
    shows: React.PropTypes.object,
  },

  render() {
    var shows = this.props.shows;

    var showElements = shows.map(show => {
      return (
        <Show key={show.get('ID')} show={show} />
      );
    });

    return (
      <div className='show-list'>
        {showElements}
      </div>
    );
  },

});

module.exports = ShowList;
