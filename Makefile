setup:
	firebase --project=demo-project setup:emulators:firestore

emulator:
	firebase emulators:start --project=demo-project --only=firestore

test:
	export FIRESTORE_EMULATOR_HOST="localhost:8080"
	go test ./tests/...

test-report:
	export FIRESTORE_EMULATOR_HOST="localhost:8080"
	go test ./tests/... -json