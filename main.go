package main

import "C"

import (
	"encoding/json"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	// "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {}

//export GetEvents
// GetEvents will return the events
func GetEvents() *C.char {
	// Use clientcmd.BuildConfigFromFlags to use config from a directory.
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)

	// config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	events, err := clientSet.CoreV1().Events("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	jsonbytes, err := json.Marshal(events.Items)
	if err != nil {
		panic(err)
	}

	return C.CString(string(jsonbytes))
}
