'use strict';

var React = require('react');
var { Link } = require('react-router');
var classNames = require('classnames');

var Dropdown = React.createClass({

  propTypes: {
    active: React.PropTypes.object.isRequired,
    list: React.PropTypes.object.isRequired,
    linkTo: React.PropTypes.string.isRequired,
    itemParams: React.PropTypes.func.isRequired,
    itemName: React.PropTypes.func.isRequired,
  },

  render() {

    var items = this.props.list.map(item => {

      var classes = classNames({
        active: item.get('ID') === this.props.active.get('ID'),
      });

      return (
        <li key={item.get('ID')} className={classes}>
          <Link to={this.props.linkTo} params={this.props.itemParams(item)}>
            {this.props.itemName(item)}
          </Link>
        </li>
      );

    });

    return (
      <div className='dropdown'>
        <Link to={this.props.linkTo} params={this.props.itemParams(this.props.active)}>
          {this.props.itemName(this.props.active)}
        </Link>
        <span className='icon icon-down-dir' />
        <ul>{items}</ul>
      </div>
    );
  },

});

module.exports = Dropdown;
