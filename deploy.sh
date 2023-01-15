#!/bin/bash

RUNTIME="-go119"
PROJECT_ID=$1

gcloud functions deploy OnUserCreated \
  --entry-point "OnUserCreated" \
  --runtime "$RUNTIME" \
  --trigger-event "providers/cloud.firestore/eventTypes/document.write" \
  --trigger-resource "projects/$PROJECT_ID/databases/(default)/documents/users/{userId}" \
  --retry