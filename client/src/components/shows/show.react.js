'use strict';

var React = require('react');
var Router = require('react-router');
var { Link } = Router;

var Show = React.createClass({

  propTypes: {
    show: React.PropTypes.object,
  },

  render() {
    var show = this.props.show;

    var date = show.get('FirstAired').split('-')[0];

    return (
      <Link to='show' params={{showID: show.get('ID')}} className='show'>
        <div className='poster' style={{
          backgroundImage: 'url(http://thetvdb.com/banners/' + show.get('Poster') + ')'
        }}>
          <div className='overlay' />
        </div>
        <h3>{show.get('Name')}</h3>
        <p>{date}</p>
      </Link>
    );
  },

});

module.exports = Show;
