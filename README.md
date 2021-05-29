# gcs-proxy

Simple reverse proxy that modifies Access-Control-Allow-Origin due to GCS bug.


## Building

Build your container image using Cloud Build, by running the following command from the directory containing the Dockerfile:

```
gcloud builds submit --tag gcr.io/PROJECT-ID/gcs-proxy
```
where **PROJECT-ID** is your GCP project ID. You can get it by running `gcloud config get-value project`.

Upon success, you will see a SUCCESS message containing the image name (gcr.io/**PROJECT-ID**/gcs-proxy). The image is stored in Container Registry and can be re-used if desired.

## Deploying to Cloud Run

To deploy the container image:

Deploy using the following command:

```
gcloud run deploy --image gcr.io/PROJECT-ID/gcs-proxy
```

If prompted to enable the API, Reply **y** to enable.

Replace **PROJECT-ID** with your GCP project ID.

1. You will be prompted for the service name: press Enter to accept the default name, helloworld.
2. You will be prompted for region: select the region of your choice, for example us-central1.
3. You will be prompted to allow unauthenticated invocations: respond y .

Then wait a few moments until the deployment is complete. On success, the command line displays the service URL.

Visit your deployed container by opening the service URL in a web browser.
