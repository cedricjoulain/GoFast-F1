# GoFast-F1

***Early dev status***

F1 Telemetry tools for both online and offline (recorded) messages.
Based on the great work from:

https://github.com/theOehrly/Fast-F1.git

## F1Viewer

Very interesting tool created with Go:

https://github.com/SoMuchForSubtlety/f1viewer.git

Needs a F1 TV pro account.

### Trick to download Video from stream

From F1Viewer Copy URL to clipboard (here named M3U8_URL)

```bash
youtube-dl --list-formats M3U8_URL
youtube-dl -f FORMAT_CODE --hls-prefer-native M3U8_URL
```
