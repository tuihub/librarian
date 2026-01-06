# tuihub-bangumi

A Bangumi plugin for tuihub that provides anime/manga database integration.

## Features

- App info source for anime, manga, games, books, and music
- Search functionality for Bangumi subjects
- Subject detail retrieval
- Required token authentication

## Configuration

The plugin requires a Bangumi API token to be configured in the PorterContext:

```json
{
  "token": "your-bangumi-api-token"
}
```

## API Support

This plugin uses the Bangumi API v0 to fetch:
- Subject information (anime/manga/game/book/music)
- Search results
- Metadata including descriptions, release dates, developers, and images

## App Types Supported

- Games (`SubjectTypeGame`)
- Anime (`SubjectTypeAnime`)  
- Books (`SubjectTypeBook`)
- Music (`SubjectTypeMusic`)
- Real (live action) (`SubjectTypeReal`)