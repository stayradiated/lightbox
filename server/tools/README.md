# Tools

Tools are listed in alphabetical order.

It's possible some tools no longer work as they rely on types specified in the
`server/db` package which when through quite a few changes recently.

## Aspect

Used to find the optimal sizes for displaying poster images in the browser.

Given two dimensions (X & Y) it will calculate possible sizes that have the
same aspect ratio and are still whole numbers.

I'm sure there is name for this, and probably an easier way to do it, but I was
in a hurry.

    > ./aspect -a 30 -b 20
    3 2
    6 4
    9 6
    12 8
    15 10
    18 12

## Consolidate Categories

Used to copy categories from Lightbox over to a new table.

Needed because when I created the new categories table all the IDs were
different from the original Lightbox category ID.

So categories are compared by name, and any categories that don't exist in the
new table are logged.

Probably could be replaced with a pure SQL script.

## Fetch

Interface for the Lightbox xstream API.

## Fetch Images

Downloads all images referred to in the database.

Runs five concurrent downloads at a time. Images are stored in their respective
folders and named as '{id}.jpg' where `id` is the id of whatever the image is
associated with.

Can handle:

- show posters
- season posters
- episode images
- show fanart

Doesn't take any options, just uncomment what needs to be downloaded.

## Fetch TVDB

Downloads a series zip file from [thetvdb.com](http://thetvdb.com).

Follows the [TVDB Programmers
API](http://thetvdb.com/wiki/index.php?title=Programmers_API).

There are two ways to use it.

Either by searching for a show by name:
    
    > ./fetch-tvdb -id 355 -name "The Lizzie Borden Chronicles

Or specifying a TVDB ID

    > ./fetch-tvdb -id 355 -tvdb 293505

Searching for shows by name is useful when needing to get data for hundreds of
shows. And then finding which shows have been matched incorrectly and manually
specifying their TVDB ID.

The `id` is only used to create the folder to store the data in. Running the
command:

    > ./fetch-tvdb -id 353 -name "X Company"

Creates `./series/353 - X Company/package.zip` and then extracts three files:

- `actors.xml` - Not used
- `banners.xml` - Contains links to images
- `en.xml` - Contains show metadata

## Import IMDB

Used to fetch IMDB show information using the [OMDb API](http://omdbapi.com).
    
    > ./import-imdb -id tt3032476 -lightbox 337

- `id` - IMDB ID
- `lightbox` - Lightbox ID

Inserts show information into `lightbox.imdb_shows`.

Also contains code to fetch episode information from IMDB, but it ended up not
getting used as the OMDB api wasn't always returning episode info.

    > ./import-imdb -show 'Better Call Saul' -season 1 -episode 1

## Import Lightbox

Parses JSON output of `fetch` and inserts information into the database.

    > ./import-lightbox -f ../fetch/output.json

Handles:

- Shows
- Seasons
- Episodes

## Import Lightbox Images

Parses JSON output of `fetch` in a similar way as `import-lightbox` but just
focuses on the `images` information.

    > ./import-lightbox-images -f ../fetch/output.json

Handles:

- Show Images
- Season Images
- Episode Images

## Import Lightbox List

Parses JSON input on `stdin` from `fetch` and inserts it into the database.

    > for i in {0..30} do ../fetch/fetch -list $i | ./import-lightbox-list
    > done

## Import TVDB

Bulk reads all `en.xml` and `banners.xml` files that are downloaded using
`fetch-tvdb`.

    > ./import-tvdb -f ../fetch-tvdb/series

Imports:

- Show info
- Season info
- Episode info
- Banner info

## List

Prints out all show information stored in the database encoded as JSON.

Useful for piping through `jq` to format and then `xargs` to run other tools.

    > ./list

## List IMDB Episodes

Prints out all episode information associated with shows in the IMDB shows
table.

Designed to be used with `import-imdb`.

    > ./list-imdb-episodes -show 'Better Call Saul' -season 1 -episode 1
## Make Categories

## Make Seasons

## Make Satic

## Season Banner

## Smoosh

## Trim Seasons

