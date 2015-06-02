'use strict';

exports.show       = ['show'];
exports.shows      = ['shows'];
exports.season     = ['season'];
exports.episode    = ['episode'];
exports.categories = ['categories'];
exports.watchlist  = ['watchlist'];

exports.fanart = ['show', 'Fanart'];

exports.watchlistShows = [
  exports.watchlist,
  exports.shows,
  (watchlist, shows) => {
    console.log(watchlist.toJS());
    return watchlist.map(item => {
      console.log('item', item);
      var show = shows.find(show => {
        return show.get('ID') === item.get('ShowID');
      });
      return show ? show.set('DateBookmarked', item.get('Date')) : null;
    }).filter(item => {
      return item != null;
    });
  },
];
