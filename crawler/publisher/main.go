package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

const basePath = "https://musicbrainz.org/ws/2/"

type Track struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Position int    `json:"position"`
}

type Media struct {
	Tracks []Track `json:"tracks"`
}
type Release struct {
	Id     string  `json:"id"`
	Title  string  `json:"title"`
	Date   string  `json:"date"`
	Medias []Media `json:"media"`
}
type ReleaseResponse struct {
	Releases []Release `json:"releases"`
}

type Artist struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
type ArtistSearchResponse struct {
	Artists []Artist `json:"artists"`
}

type Message struct {
	Artist   Artist    `json:"artist"`
	Releases []Release `json:"releases"`
}

func searchArtist(name string) []Artist {
	logText := fmt.Sprintf("[searching] %s", name)
	fmt.Println(logText)
	req, err := http.NewRequest("GET", basePath+"artist", nil)
	if err != nil {
		panic(err)
	}

	q := req.URL.Query()
	q.Add("query", name)
	q.Add("fmt", "json")
	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	var data ArtistSearchResponse
	json.Unmarshal(resBody, &data)
	logText = fmt.Sprintf("[searching] Response %v", data)
	fmt.Println(logText)
	return data.Artists
}

func getRelease(artist string) []Release {
	req, err := http.NewRequest("GET", basePath+"release", nil)
	if err != nil {
		panic(err)
	}

	q := req.URL.Query()
	q.Add("artist", artist)
	q.Add("inc", "recordings")
	q.Add("fmt", "json")
	req.URL.RawQuery = q.Encode()
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	var data ReleaseResponse
	json.Unmarshal(resBody, &data)
	return data.Releases
}

func main() {
	godotenv.Load("../../.env")
	rHost := os.Getenv("RABBITMQ_HOST")
	rVHost := os.Getenv("RABBITMQ_VHOST")
	rUser := os.Getenv("RABBITMQ_USER")
	rPwd := os.Getenv("RABBITMQ_PWD")

	rURI := fmt.Sprintf("amqp://%s:%s@%s%s", rUser, rPwd, rHost, rVHost)
	conn, err := amqp.Dial(rURI)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"music-creation", // name
		true,             // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		panic(err)
	}

	fileContent, err := ioutil.ReadFile("./bands.txt")
	if err != nil {
		log.Fatal(err)
	}

	bands := strings.Split(string(fileContent), "\n")
	for _, band := range bands {
		artists := searchArtist(band)
		for _, a := range artists {
			releases := getRelease(a.Id)
			message := Message{
				Artist:   a,
				Releases: releases,
			}
			err = publish(ch, &q, message)
			if err != nil {
				panic(err)
			}
		}

		time.Sleep(5 * time.Second)
	}

}

func publish(ch *amqp.Channel, q *amqp.Queue, message Message) error {
	msg, err := json.Marshal(&message)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err != nil {
		return err
	}

	return ch.PublishWithContext(
		ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msg,
		},
	)
}
