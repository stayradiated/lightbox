'use strict';

var React = require('react');

var Fake = React.createClass({

  render() {
    var show = this.props.show;

    return (
      <div className='show fake'>
        <div className='poster'  />
        <h3>Title</h3>
        <p>2000</p>
      </div>
    );
  },

});

module.exports = Fake;
