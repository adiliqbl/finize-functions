#!/bin/bash

RUNTIME="go119"
PROJECT_ID=$1

gcloud functions deploy TransactionCreated \
  --entry-point "OnTransactionCreated" \
  --runtime $RUNTIME \
  --trigger-event "providers/cloud.firestore/eventTypes/document.create" \
  --trigger-resource "projects/$PROJECT_ID/databases/(default)/documents/user-transactions/{userId}/transactions/{transactionId}" \
  --retry

gcloud functions deploy TransactionUpdated \
  --entry-point "OnTransactionUpdated" \
  --runtime $RUNTIME \
  --trigger-event "providers/cloud.firestore/eventTypes/document.update" \
  --trigger-resource "projects/$PROJECT_ID/databases/(default)/documents/user-transactions/{userId}/transactions/{transactionId}" \
  --retry

gcloud functions deploy TransactionDeleted \
  --entry-point "OnTransactionDeleted" \
  --runtime $RUNTIME \
  --trigger-event "providers/cloud.firestore/eventTypes/document.delete" \
  --trigger-resource "projects/$PROJECT_ID/databases/(default)/documents/user-transactions/{userId}/transactions/{transactionId}" \
  --retry

gcloud functions deploy ProcessRecurringTasks \
  --entry-point "ProcessRecurringTasks" \
  --runtime $RUNTIME \
  --trigger-http \
  --timeout 540s

gcloud functions deploy GetExchangeRate \
  --entry-point "GetExchangeRate" \
  --set-secrets 'EXCHANGE_RATES_API=EXCHANGE_RATES_API:1' \
  --runtime $RUNTIME \
  --trigger-http \
  --timeout 180s
