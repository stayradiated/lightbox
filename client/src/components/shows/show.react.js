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

    return (
      <Link to='show' params={{showID: show.get('ID')}} className='show'>
        <div className='poster' style={{
          backgroundImage: 'url(' +  show.get('Poster') + ')'
        }}>
          <div className='overlay' />
        </div>
        <h3>{show.get('Title')}</h3>
        <p>{show.get('Year')}</p>
      </Link>
    );
  },

});

module.exports = Show;
