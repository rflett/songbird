# SongBird

*chirp chirp*

A Golang script that streams [@triplejplays](https://twitter.com/triplejplays) Twitter feed and parses the song names out of the tweets.

## Development

### Requirements

- Go 1.13+
- Twitter Development Account

Install Go dependencies:

```bash
go get github.com/dghubble/go-twitter/twitter
go get github.com/dghubble/oauth1
go get github.com/joho/godotenv
```

### Running locally

Update the `.env` file to include your Twitter API credentials and the API to send the song name to. 

The following request is made when a song is played:

```bash
curl -XPOST  \
     -H 'Content-Type: application/json' \
     --data '{"name": "SONG_NAME"}' \
     $API_URL
```

Then run `go run chirp.go`.

```bash
2020/01/04 16:10:26 Starting Stream...
2020/01/04 16:13:18 Black, White And Blue
```
