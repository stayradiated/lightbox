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
      return null;
    }
    
    var year = show.get('Released').getFullYear();
    if (year === 1) {
      year = "";
    }

    return (
      <Link to='show' params={{showID: show.get('ID')}} className='show-item'>
        <Poster id={show.get('ID')} type='shows' onAdd={this.onAdd} />
        <h3>{show.get('Title')}</h3>
        <p>{year}</p>
      </Link>
    );
  },

  onAdd(e) {
    e.preventDefault();
    Lightbox.actions.bookmarkShow(this.props.show.get('ID'));
  },

});

module.exports = Show;
