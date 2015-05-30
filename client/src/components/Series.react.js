'use strict';

var React = require('react');

var Series = React.createClass({

  propTypes: {
    series: React.PropTypes.object,
  },

  render() {
    var series = this.props.series;

    if (series == null) {
      return null;
    }

    console.log('Rendering series');

    return (

      <div className='series-view'>

        <div className='poster' style={{
          backgroundImage: 'url(http://thetvdb.com/banners/' + series.get('Poster') + ')'
        }}>
          <div className='overlay' />
        </div>

        <h3>Parks and Recreation</h3>

        <p>22 min</p>

        <p>Overview</p>

        <p>Rated M</p>

      </div>

    );

  },

});

module.exports = Series;
