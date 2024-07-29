test-integration:
	TEST_INTEGRATION=true go test ./... -v

test-unit:
	TEST_INTEGRATION=false go test ./... -v

serve-dev:
	go run queensbattle serve