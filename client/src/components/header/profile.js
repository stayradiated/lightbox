'use strict';

var React = require('react');
var DropdownButton = require('../common/dropdown-button');

var profileItems = [
  { label: 'Add Profile', },
  { label: 'Settings', },
  { label: 'Help', },
  { label: 'Logout', },
];

var Profile = React.createClass({

  render() {
    return (
      <div className='profile'>
        <DropdownButton items={profileItems}>
          George Czabania 
          <span className='icon icon-down-open' />
        </DropdownButton>
      </div>
    );
  },

});

module.exports = Profile;
