#!/bin/bash

RUNTIME="-go119"
PROJECT_ID=$1

gcloud functions deploy TransactionCreated \
  --entry-point "OnTransactionCreated" \
  --runtime "$RUNTIME" \
  --trigger-event "providers/cloud.firestore/eventTypes/document.create" \
  --trigger-resource "projects/$PROJECT_ID/databases/(default)/documents/user-transactions/{userId}/transactions/{transactionId}" \
  --retry

gcloud functions deploy TransactionUpdated \
  --entry-point "OnTransactionUpdated" \
  --runtime "$RUNTIME" \
  --trigger-event "providers/cloud.firestore/eventTypes/document.update" \
  --trigger-resource "projects/$PROJECT_ID/databases/(default)/documents/user-transactions/{userId}/transactions/{transactionId}" \
  --retry

gcloud functions deploy TransactionDeleted \
  --entry-point "OnTransactionUpdated" \
  --runtime "$RUNTIME" \
  --trigger-event "providers/cloud.firestore/eventTypes/document.delete" \
  --trigger-resource "projects/$PROJECT_ID/databases/(default)/documents/user-transactions/{userId}/transactions/{transactionId}" \
  --retry