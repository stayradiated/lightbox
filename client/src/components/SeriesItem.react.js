'use strict';

var React = require('react');

var SeriesItem = React.createClass({

  propTypes: {
    series: React.PropTypes.object,
  },

  render() {
    var series = this.props.series;

    return (
      <div className='series-item'>
        <div className='overlay' />
        <img src={'http://thetvdb.com/banners/' + series.get('Poster')} />
        <h3>{series.get('Name')}</h3>
        <p>{series.get('FirstAired')}</p>
      </div>
    );
  },

});

module.exports = SeriesItem;
