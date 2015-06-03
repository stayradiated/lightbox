'use strict';

var React = require('react');
var classNames = require('classnames');

var Poster = React.createClass({

  propTypes: {
    type: React.PropTypes.string,
    id: React.PropTypes.number,
    size: React.PropTypes.string,
    onPlay: React.PropTypes.func,
    onAdd: React.PropTypes.func,
    playOnly: React.PropTypes.bool,
  },

  getDefaultProps() {
    return {
      size: 'small',
    };
  },

  render() {

    var url = "./default.jpg";

    if (true) {
      url = [
        './images',
        this.props.type,
        this.props.size,
        this.props.id + '.jpg',
      ].join('/');
    }

    var classes = classNames({
      'poster': true,
      'large': this.props.size === 'large',
      'small': this.props.size === 'small',
      'play-only': this.props.playOnly,
    });

    return (
      <div className={classes} style={{
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
