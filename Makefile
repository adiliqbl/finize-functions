setup:
	firebase --project=demo-project setup:emulators:firestore

emulator:
	firebase emulators:start --project=demo-project --only=firestore

test:
	export FIRESTORE_EMULATOR_HOST="localhost:8080"
	firebase emulators:exec --project=demo-project --only=firestore "go test ./tests/... -p 1"

test-report:
	export FIRESTORE_EMULATOR_HOST="localhost:8080"
	firebase emulators:exec --project=demo-project --only=firestore "go test ./tests/... -json -p 1"
