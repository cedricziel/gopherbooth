package handlers
import (
	"net/http"
	"appengine"
	"appengine/file"
	"log"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/storage/v1"
)

// HandleBucket routes requests to the /bucket space.
func HandleBuckets(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleListBucket(w, r)
		break;
	}
}

// handleListBucket lists the contents of the default bucket
func handleListBucket(w http.ResponseWriter, r *http.Request) {

	ctx := appengine.NewContext(r)

	client, err := google.DefaultClient(context.Background(), storage.DevstorageFullControlScope)
	if err != nil {
		log.Println("Error initializing client")
	}

	service, err := storage.New(client)
	if err != nil {
		log.Println("No connection to cloud storage")
	}

	defaultBucketName, _ := file.DefaultBucketName(ctx)
	log.Print(defaultBucketName)

	_, err = service.Buckets.Get(defaultBucketName).Do()
	if err != nil {
		log.Println("Error retrieving Bucket")
	}

	service.Objects.List(defaultBucketName)


	//var client *storage.Client // See Example (Auth)
	//var query *storage.Query




	// objects, _ := client.Bucket(defaultBucketName).List(ctx, query)

	// for _, obj := range objects.Results {
	//	log.Printf("Got file %s", obj.Name)
	//}
}
