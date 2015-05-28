# TVDB Notes

API Key: `71D4D0F8D1336E0D`

## Initial Setup

### 1. Get Mirror List

Deprecated. Just use http://thetvdb.com.

### 2. Get Language List

`http://thetvdb.com/api/{apikey}/languages.xml`

The `abbreviation` is the important bit.

To make development easier, the interface is going to be English only, but it
would be a good idea to allow the user to select which language they want.

    <Languages>
        ...
        <Language>
            <name>English</name>
            <abbreviation>en</abbreviation>
            <id>7</id>
        </Language>
        ...
    </Languages>

### 3. Get Server Time

`http://thetvdb.com/api/Updates.php?type=none`

This is so you can get database updates in the future.

    <Items>
        <Time>1432765997</Time>
    </Items>

### 4. Find Series ID

This needs to be done for all the series available on Lightbox.

http://thetvdb.com/api/GetSeries.php?seriesname={seriesname}

Save `Data.Series.seriesid` to the local database.

    <Data>
        <Series>
            <seriesid>75932</seriesid>
            <language>en</language>
            <SeriesName>Fawlty Towers</SeriesName>
            <AliasNames>Faulty Towers</AliasNames>
            <banner>graphical/75932-g2.jpg</banner>
            <Overview>Fawlty Towers is a British series featuring John Cleese as Basil Fawlty, a hotel owner whose incompetence, short fuse and arrogance form a combination that ensures accidents and trouble are never far away. Although only twelve episodes were produced, the show left a lasting and powerful legacy and is considered by many to be the greatest of all TV comedies.</Overview>
            <FirstAired>1975-09-19</FirstAired>
            <Network>BBC Two</Network>
            <IMDB_ID>tt0072500</IMDB_ID>
            <zap2it_id>SH001588</zap2it_id>
            <id>75932</id>
        </Series>
    </Data>

### 4. Get Series Info

`http://thetvdb.com/api/{apikey}/series/{seriesid}/all/{language}.zip`

Extract the zip to get three files:

- `actors.xml` - list of actors
- `banners.xml` - series images and fan art
- `en.xml` - series description, metadata, seasons and episode info

Get banner image: `http://thetvdb.com/banners/{BannerPath}`

### 5. Get Episode Info

Get episode image: `http://thetvdb.com/banners/{episode.filepath}`

## Cheking for Updates

### 1. Check What Has Been Updated

`http://thetvdb.com/api/Updates.php?type=all&time={previoustime}`

Get `previoustime` from step 3 above.

- `Items.Time`
- `Items.Series`
- `Items.Episode`

    <Items>
        <Time>1432767639</Time>
        <Series>293517</Series>
        <Series>79490</Series>
        <Series>296233</Series>
        <Series>296237</Series>
        <Series>81798</Series>
        <Series>71595</Series>
        <Series>80379</Series>
        <Series>291215</Series>
        <Series>72775</Series>
        <Series>296245</Series>
        <Episode>5166697</Episode>
        <Episode>1494941</Episode>
        <Episode>1494951</Episode>
        <Episode>1494961</Episode>
        <Episode>5155186</Episode>
        <Episode>1494971</Episode>
        <Episode>5236913</Episode>
        <Episode>1494981</Episode>
        <Episode>1494991</Episode>
        <Episode>5236912</Episode>
        <Episode>5165642</Episode>
    </Items>

### 2. Update Series Info

`http://thetvdb.com/api/{apikey}/series/{seriesid}/en.xml`

### 3. Update Episode Info

`http://thetvdb.com/api/{apikey}/episode/{episodeid}/en.xml`
