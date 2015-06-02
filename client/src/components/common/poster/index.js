'use strict';

var React = require('react');

var Poster = React.createClass({

  propTypes: {
    url: React.PropTypes.string,
    onPlay: React.PropTypes.func,
    onAdd: React.PropTypes.func,
  },

  render() {

    var url = "/default.jpg";

    if (this.props.url != null) {
      url = this.props.url;
    }

    return (
      <div className='poster' style={{
        backgroundImage: 'url(' + url + ')'
      }}>
        <div className='overlay'>
          <button className='icon icon-play' onClick={this.props.onPlay} />
          <button className='icon icon-plus' onClick={this.props.onAdd} />
        </div>
      </div>
    );
  },

});

module.exports = Poster;
