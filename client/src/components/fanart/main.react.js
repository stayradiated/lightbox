'use strict';

var React  = require('react');
var Router = require('react-router');
var { RouteHandler} = Router;

var flux     = require('../../flux');
var Lightbox = require('../../modules/lightbox');

var Fanart = React.createClass({
	mixins: [flux.ReactMixin],

	getDataBindings() {
		return {
			fanart: Lightbox.getters.fanart,
		};
	},

	render() {
    var url = '';

    if (this.state.fanart != null) {
      url = 'url(' + this.state.fanart + ')';
    }

		return (
			<div className='route-fanart'>
        <RouteHandler />
        <div className='background' style={{ backgroundImage: url }} />
			</div>
		);
	}

});

module.exports = Fanart;
