'use strict';

var React = require('react');
var classNames = require('classnames');
var { Link } = require('react-router');

var Dropdown = require('./dropdown');
var Categories = require('./categories');

var Header = React.createClass({

  propTypes: {
    show: React.PropTypes.object,
    season: React.PropTypes.object,
    episode: React.PropTypes.object,
  },
  
  render() {

    var activeShow = this.props.show;
    var activeSeason = this.props.season;
    var activeEpisode = this.props.episode;

    var showLink = null;
    if (activeShow != null) {
      showLink = (
        <span className='active-item'>
          <Link to='show' params={{showID: activeShow.get('ID')}}>
            {activeShow.get('Title')}
          </Link>
        </span>
      );
    }

    var seasonDropdown = null;
    if (activeShow != null && activeSeason != null) {
      seasonDropdown = <Dropdown
        active={activeSeason}
        list={activeShow.get('Seasons')}
        linkTo='season'
        itemParams={season => {
          return {
            showID: activeShow.get('ID'),
            seasonID: season.get('ID'),
          };
        }}
        itemName={season => {
          return 'Season ' + season.get('Number');
        }}
      />;
    }

    var episodeDropdown = null;
    if (activeShow != null && activeSeason != null && activeEpisode != null) {
      episodeDropdown = <Dropdown
        active={activeEpisode}
        list={activeSeason.get('Episodes')}
        linkTo='episode'
        itemParams={episode => {
          return {
            showID: activeShow.get('ID'),
            seasonID: activeSeason.get('ID'),
            episodeID: episode.get('ID'),
          };
        }}
        itemName={episode => {
          return episode.get('Title');
        }}
      />;
    }

    return (
      <header className='show-header'>
        <Categories />
        {showLink}
        {seasonDropdown}
        {episodeDropdown}
      </header>
    );
  },

});

module.exports = Header;
