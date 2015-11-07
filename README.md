## `subbs`
**subbs** is a simple program written in Go which calculates a file hash and use it to search for corresponding subtitle(s). The differential is that it's entirely based on the hash instead of metadata.

### Features
- Scans directory matching `.mp4`, `.mkv` and `.avi` files;
- Allows to specify preferred languages.

### Limitations
- Only `.srt` supported;
- Daily limited by third-party API(s);
- No metadata based search (for the accuracy sake).

**Assumption**: as many players picks the first `.srt` file matching the same name, **subbs** will also do that when downloading subtitles.
