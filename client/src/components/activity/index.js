'use strict';

var React = require('react');

var flux     = require('../../flux');
var Lightbox = require('../../modules/lightbox');
var List     = require('../common/list');

var Recommended = React.createClass({
  mixins: [flux.ReactMixin],

  getDataBindings() {
    return {
      lists: Lightbox.getters.lists,
    };
  },

	render() {
    var lists = this.state.lists;

    return (
      <div className='route-lists'>
        <h1>Recommended</h1>
        {
          lists.map(list => {
            return (
              <List key={list.get('ID')} listID={list.get('ID')} showTitle={true}/>
            );
          })
        }
      </div>
    );
	},

});

module.exports = Recommended;
