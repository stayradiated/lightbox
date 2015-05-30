'use strict';

var React = require('react');

var SeriesItem = React.createClass({

  propTypes: {
    series: React.PropTypes.object,
  },

  render() {
    var series = this.props.series;

    var date = series.get('FirstAired').split('-')[0];

    return (
      <a href={'#/series/' + series.get('ID')} className='series-item'>
        <div className='poster' style={{
          backgroundImage: 'url(http://thetvdb.com/banners/' + series.get('Poster') + ')'
        }}>
          <div className='overlay' />
        </div>
        <h3>{series.get('Name')}</h3>
        <p>{date}</p>
      </a>
    );
  },

});

module.exports = SeriesItem;
