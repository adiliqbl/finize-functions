#!/bin/bash

while getopts ":p:r:" opt; do
  case $opt in
    p) pid="$OPTARG"
    ;;
    r) runtime="$OPTARG"
    ;;
    \?) echo "Invalid option -$OPTARG" >&2
    exit 1
    ;;
  esac

  case $OPTARG in
    -*) echo "Option $opt needs a valid argument"
    exit 1
    ;;
  esac
done

gcloud functions deploy OnUserCreated \
  --entry-point ENTRY_POINT \
  --runtime "${runtime:-go119}" \
  --trigger-event "providers/cloud.firestore/eventTypes/document.write" \
  --trigger-resource "projects/$pid/databases/(default)/documents/users/{userId}" \
  --retry