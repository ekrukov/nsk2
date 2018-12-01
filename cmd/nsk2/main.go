package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	log.Print("Start application")

	port := os.Getenv("PORT")

	s := http.Server{
		Addr:    ":" + port,
		Handler: nil,
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/healthz", helthz)
	http.HandleFunc("/readyz", ready)

	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("Server is stopped with error: %v", err)
	}

	log.Print("Stop application")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, http.StatusText(http.StatusOK))
}

func helthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, http.StatusText(http.StatusOK))
}

func ready(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, http.StatusText(http.StatusOK))
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Print("The hello handler was called")

	token := os.Getenv("K8S_TOKEN")

	config := &rest.Config{
		Host:            "https://master.k8s.community:443",
		BearerToken:     token,
		TLSClientConfig: rest.TLSClientConfig{Insecure: true},
	}

	c, err := kubernetes.NewForConfig(config)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Couldn't create a k8s client: %v", err)
	}

	podlist, err := c.Core().Pods("ekrukov").List(meta.ListOptions{})
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Couldn't get the list of pods: %v", err)
	}

	var podnames []string
	for _, pod := range podlist.Items {
		podnames = append(podnames, pod.GetName())
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "The list of pods: [%v].", podnames)
}
