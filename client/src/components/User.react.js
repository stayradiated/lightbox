var React = require('react');
var flux = require('../flux');
var Faves = require('../modules/faves');

var User = React.createClass({
  mixins: [flux.ReactMixin],

  getDataBindings() {
    return {
      username: Faves.getters.username,
    };
  },

  render() {
    var username = this.state.username;

    return (
      <form className='user' onSubmit={this.onSubmit}>
        <input ref='username' placeholder='Username' />
        <button type='submit'>Get Faves</button>
      </form>
    );
  },

  onSubmit(event) {
    console.log(event);
  },

});

module.exports = User;
