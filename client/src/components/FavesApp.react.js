var React = require('react');
var User = require('./User.react');

var FavesApp = React.createClass({
  render() {
    return (
      <div className="favesapp">
        <header>
          <User />
        </header>
      </div>
    )
  }
});

module.exports = FavesApp;
