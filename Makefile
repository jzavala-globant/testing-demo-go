run:
	@clear
	@echo "Running..."
	@echo
	@go run cmd/testing-demo/main.go
	@echo

test:
	@clear
	@echo "Testing..."
	@echo
	@go test -coverprofile output/cover.out ./internal/...;go-cover-treemap -coverprofile output/cover.out > output/heatmap.svg
	@echo

mock:
	@clear
	@echo "Creating mocks..."
	@echo
	mockery --all --dir=internal/services --output=internal/services/mocks
	@echo

show-dependencies:
	@clear
	@echo "Generating output/graph.svg..."
	@echo
	@goda graph ./...| dot -Tsvg -o output/graph.svg
	@echo

coverage: test
	@clear
	@echo "Getting coverage..."
	@echo
	@go-coverage -f output/cover.out --line-filter 0
	@echo