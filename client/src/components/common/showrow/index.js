'use strict';

var React = require('react');

var Show = require('../show/');

const ITEM_WIDTH = 187;
const ITEM_MARGIN = 20;

var ShowRow = React.createClass({

  propTypes: {
    shows: React.PropTypes.object,
  },

  getInitialState() {
    return {
      offset: 0,
    };
  },

  render() {
    var shows = this.props.shows;

    var showElements = shows.map(show => {
      return (
        <Show key={show.get('ID')} show={show} />
      );
    });

    var transform = 'translate(' + (this.state.offset * (ITEM_WIDTH + ITEM_MARGIN) * -1) + 'px, 0px)';

    return (
      <div className='show-row'> 
        <div className='buttons'>
          <button onClick={this.goLeft}><span className='icon icon-left-open' /></button>
          <button onClick={this.goRight}><span className='icon icon-right-open' /></button>
        </div>
        <div className='shows' style={{
          width: (ITEM_WIDTH + ITEM_MARGIN) * this.props.shows.size,
          transform: transform,
          WebkitTransform: transform,
        }}>
          {showElements}
        </div>
      </div>
    );
  },

  goLeft() {
    var offset = this.state.offset - this.getViewableItems();
    if (offset < 0) {
      offset = 0;
    }
    this.setState({ offset });
  },

  goRight() {
    var viewableItems = this.getViewableItems();
    var max = this.props.shows.size;

    if (this.state.offset + viewableItems >= max) {
      return;
    }

    var offset = this.state.offset + viewableItems;
    if (offset > max) {
      offset = max;
    }
    this.setState({ offset });
  },

  getViewableItems() {
    var containerWidth = this.getDOMNode().offsetWidth;
    return Math.floor(containerWidth / (ITEM_WIDTH + ITEM_MARGIN));
  },

});

module.exports = ShowRow;
