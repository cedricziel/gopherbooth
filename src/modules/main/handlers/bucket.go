package handlers

import (
	"fmt"
	"net/http"

	"io/ioutil"

	"google.golang.org/appengine"
	"google.golang.org/appengine/file"
	"google.golang.org/appengine/log"
	"google.golang.org/cloud/storage"
)

// bucket is a local cache of the app's default bucket name.
var bucket string // or: var bucket = "<your-app-id>.appspot.com"

// HandleBucket routes requests to the /bucket space.
func HandleBuckets(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleListBucket(w, r)
		break
	}
}

// handleListBucket lists the contents of the default bucket
func handleListBucket(w http.ResponseWriter, r *http.Request) {

	ctx := appengine.NewContext(r)
	if bucket == "" {
		var err error
		if bucket, err = file.DefaultBucketName(ctx); err != nil {
			log.Errorf(ctx, "failed to get default GCS bucket name: %v", err)
			return
		}
	}

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Errorf(ctx, "failed to get default GCS bucket name: %v", err)
		return
	}
	defer client.Close()

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "Demo GCS Application running from Version: %v\n", appengine.VersionID(ctx))
	fmt.Fprintf(w, "Using bucket name: %v\n\n", bucket)

	writer := client.Bucket(bucket).Object("foo").NewWriter(ctx)

	_, err = writer.Write([]byte("Yoho!"))
	if err != nil {
		log.Errorf(ctx, "Error while creating file")
	}
	log.Infof(ctx, "Created a new file. So yay!")

	reader, err := client.Bucket(bucket).Object("foo").NewReader(ctx)
	if err != nil {
		log.Errorf(ctx, "Error creating reader for object from Bucket. %s", err)
	}

	content, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Errorf(ctx, "Error reading object")
	}

	fmt.Fprintf(w, "Got: %s", &content)
}
