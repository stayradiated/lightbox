'use strict';

var React = require('react');
var Router = require('react-router');
var { Link } = Router;

var Lightbox = require('../../../modules/lightbox');

var Poster = require('../poster/');

var Show = React.createClass({

  propTypes: {
    show: React.PropTypes.object,
  },

  render() {
    var show = this.props.show;

    if (! show.has('ID')) {
      console.log('FOUND NULL SHOW');
      return null;
    }

    return (
      <Link to='show' params={{showID: show.get('ID')}} className='show'>
        <Poster id={show.get('ID')} type='shows' onAdd={this.onAdd} />
        <h3>{show.get('Title')}</h3>
        <p>{show.get('Year')}</p>
      </Link>
    );
  },

  onAdd(e) {
    e.preventDefault();
    console.log('bookmarking show...');
    Lightbox.actions.bookmarkShow(this.props.show.get('ID'));
  },

});

module.exports = Show;
