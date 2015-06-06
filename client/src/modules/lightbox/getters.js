'use strict';

exports.sort       = ['sort'];
exports.show       = ['show'];
exports.shows      = ['shows', 'Shows'];
exports.showsSearch  = ['shows', 'Search'];
exports.season     = ['season'];
exports.episode    = ['episode'];
exports.categories = ['categories'];
exports.watchlist  = ['watchlist'];
exports.lists      = ['lists'];

exports.fanart = ['show', 'Fanart'];

exports.watchlistShows = [
  exports.watchlist,
  exports.shows,
  (watchlist, shows) => {
    return watchlist.map(item => {
      var show = shows.find(show => {
        return show.get('ID') === item.get('ShowID');
      });
      return show ? show.set('DateBookmarked', item.get('Date')) : null;
    }).filter(item => {
      return item != null;
    });
  },
];
