'use strict';

var React = require('react');
var classNames = require('classnames');

var DropdownButton = React.createClass({

  propTypes: {
    items: React.PropTypes.array.isRequired,
  },

  getInitialState() {
    return {
      open: false,
    };
  },

  render() {

    var items = this.props.items.map((item, i) => {
      return (
        <li key={i}><a>{item.label}</a></li>
      );
    });

    var classes = classNames({
      'dropdown-button': true,
      'open': this.state.open,
    });

    return (
      <span className={classes}>
        <a className='button' onClick={this.onButtonClick}>
          {this.props.children}
        </a>
        <ul>
          {items}
        </ul>
      </span>
    );
  },

  openMenu() {
    this.setState({ open: true, });
  },

  closeMenu() {
    this.setState({ open: false, });
  },

  onButtonClick(e) {
    e.stopPropagation();
    this.addEventListener();
    this.openMenu();
  },

  onDocumentClick() {
    this.removeEventListener();
    this.closeMenu();
  },

  addEventListener() {
    document.addEventListener('click', this.onDocumentClick);
  },

  removeEventListener() {
    document.removeEventListener('click', this.onDocumentClick);
  },

  shouldComponentUpdate(nextProps, nextState) {
    return nextState.open !== this.state.open;
  },

  componentWillUnmount() {
    this.removeEventListener();
  }

});

module.exports = DropdownButton;
