gcloud functions deploy ${YOUR_APP_NAME} \
--runtime=go119 \
--entry-point CloneX \
--trigger-http \
--region asia-southeast1 \
--env-vars-file ./config/.env.yaml \
--allow-unauthenticated \
--source=.