package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

func sendToAPI(song string) {
	apiURL := os.Getenv("API_URL")

	// build the request body
	requestBody, err := json.Marshal(map[string]string{
		"name": song,
	})

	if err != nil {
		log.Fatalln(err)
	}

	// send song to api
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		log.Fatalln(err)
	}

	// close response
	defer resp.Body.Close()
}

func handleTweet(t string) {
	/*
	 Tweets from triplejplays are usually in the format '.@AtheAstronaut - I Like To Dance [10:12]', so
	 we get the song from after the dash, and then remove the timestamp (7 chars).

	 The artist name is sometimes a twitter handle and other times just a string so we will ignore it
	*/
	dashIndex := strings.Index(t, "-")
	songName := strings.TrimSpace(t[dashIndex+2 : len(t)-7])

	/* Sometimes the song name contains braces because it features another artist:
	.@GlassAnimals - Tokyo Drifting {ft. Denzel Curry} [17:16]

	so we should remove from the braces
	*/
	bracesIndex := strings.Index(songName, "{")

	if bracesIndex >= 0 {
		songName = strings.TrimSpace(songName[:bracesIndex])
	}

	log.Println(songName)
	sendToAPI(songName)
}

func streamFeed(client *twitter.Client) {
	// setup the demux with a tweet handler
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		handleTweet(tweet.Text)
	}

	log.Println("Starting Stream...")

	// filter the stream to follow triplejplays
	filterParams := &twitter.StreamFilterParams{
		Follow:        []string{"86848460"}, // @triplejplays user ID
		StallWarnings: twitter.Bool(true),
	}
	stream, err := client.Streams.Filter(filterParams)
	if err != nil {
		log.Fatal(err)
	}

	// receive messages until stopped or stream quits
	go demux.HandleChan(stream.Messages)

	// wait for SIGINT and SIGTERM
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)

	log.Println("Stopping Stream...")
	stream.Stop()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env found, using environment only")
	}

	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessSecret := os.Getenv("ACCESS_SECRET")

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	streamFeed(client)
}
