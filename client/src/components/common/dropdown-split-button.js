'use strict';

var React = require('react');
var classNames = require('classnames');
var { Link } = require('react-router');

var DropdownSplitButton = React.createClass({

  propTypes: {
    active:      React.PropTypes.object.isRequired,
    items:       React.PropTypes.object.isRequired,
    linkTo:      React.PropTypes.string.isRequired,
    itemParams:  React.PropTypes.func.isRequired,
    itemName:    React.PropTypes.func.isRequired,
  },

  getInitialState() {
    return {
      open: false,
    };
  },

  render() {

    var items = this.props.items.map(item => {

      var isActive = item.get('ID') === this.props.active.get('ID');

      return (
        <li key={item.get('ID')} className={classNames({active: isActive})}>
          <Link to={this.props.linkTo} params={this.props.itemParams(item)}>
            {this.props.itemName(item)}
            {isActive ? <span className='icon icon-ok' /> : null}
          </Link>
        </li>
      );

    });

    var classes = classNames({
      'dropdown-split-button': true,
      'open': this.state.open,
    });

    return (
      <span className={classes}>
        <Link to={this.props.linkTo} params={this.props.itemParams(this.props.active)}>
          {this.props.itemName(this.props.active)}
        </Link>
        <a className='button' onClick={this.onButtonClick}>
          <span className='icon icon-down-dir' />
        </a>
        <ul>{items}</ul>
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

module.exports = DropdownSplitButton;
