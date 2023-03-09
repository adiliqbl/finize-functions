#!/bin/bash

RUNTIME="go119"
PROJECT_ID=$1
SERVICE_ACCOUNT=$2

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

gcloud functions deploy OnBudgetDeleted \
  --entry-point "OnBudgetDeleted" \
  --runtime $RUNTIME \
  --trigger-event "providers/cloud.firestore/eventTypes/document.delete" \
  --trigger-resource "projects/$PROJECT_ID/databases/(default)/documents/user-budgets/{userId}/budgets/{budgetId}" \
  --retry

gcloud functions deploy OnAccountUpdated \
  --entry-point "OnAccountUpdated" \
  --runtime $RUNTIME \
  --trigger-event "providers/cloud.firestore/eventTypes/document.update" \
  --trigger-resource "projects/$PROJECT_ID/databases/(default)/documents/user-accounts/{userId}/accounts/{accountId}" \
  --retry

gcloud functions deploy OnUserUpdated \
  --entry-point "OnUserUpdated" \
  --runtime $RUNTIME \
  --trigger-event "providers/cloud.firestore/eventTypes/document.update" \
  --trigger-resource "projects/$PROJECT_ID/databases/(default)/documents/users/{userId}" \
  --retry

gcloud functions deploy ProcessRecurringTasks \
  --entry-point "ProcessRecurringTasks" \
  --runtime $RUNTIME \
  --region europe-west1 \
  --trigger-http \
  --timeout 600s

gcloud functions deploy GetExchangeRate \
  --entry-point "GetExchangeRate" \
  --set-secrets 'EXCHANGE_RATES_API=EXCHANGE_RATES_API:1' \
  --runtime $RUNTIME \
  --region europe-west1 \
  --trigger-http \
  --timeout 180s

if [ -n "$SERVICE_ACCOUNT" ]; then
  gcloud scheduler jobs create http ProcessRecurringTasks \
    --schedule="0 1 * * *" \
    --uri="https://europe-west1-${PROJECT_ID}.cloudfunctions.net/ProcessRecurringTasks/" \
    --http-method=POST \
    --oidc-service-account-email="$SERVICE_ACCOUNT" \
    --oidc-token-audience="https://europe-west1-${PROJECT_ID}.cloudfunctions.net/ProcessRecurringTasks" \
    --project="${PROJECT_ID}" \
    --time-zone="Etc/UTC" \
    --attempt-deadline=1800s
fi
