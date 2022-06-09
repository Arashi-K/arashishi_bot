package handler

import (
	"arashishi_bot/config"
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/slack-go/slack"
)

func RunHandlers() {
	handleEvent()
	handleAction()

	log.Println("[INFO] Server listening")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}

func slackVerificationMiddleware(path string, next http.HandlerFunc) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		log.Println(path)
		verifier, err := slack.NewSecretsVerifier(r.Header, config.SlackSecretKey)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		bodyReader := io.TeeReader(r.Body, &verifier)
		body, err := ioutil.ReadAll(bodyReader)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := verifier.Ensure(); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		next.ServeHTTP(w, r)

		log.Println(3)
	})
	log.Println(4)
}
