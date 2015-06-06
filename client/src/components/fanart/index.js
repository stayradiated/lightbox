'use strict';

var React  = require('react');
var Router = require('react-router');
var ImmutableRenderMixin = require('react-immutable-render-mixin');
var { RouteHandler} = Router;

var flux     = require('../../flux');
var Lightbox = require('../../modules/lightbox');

var Fanart = React.createClass({
	mixins: [flux.ReactMixin, ImmutableRenderMixin],

	getDataBindings() {
		return {
			show: Lightbox.getters.show,
		};
	},

	render() {
    var show = this.state.show;

    var style = {};
    if (show.has('ID')) {
      style.backgroundImage = 'url(./images/fanart/large/' + show.get('ID') + '.jpg)';
    }

		return (
			<div className='route-fanart'>
        <RouteHandler />
        <div className='background' style={style} />
			</div>
		);
	}

});

module.exports = Fanart;
