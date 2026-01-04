# -------------------------------------------------
# Reports
# -------------------------------------------------

ROOT_DIR   			:= $(CURDIR)
BUILD_DIR			:= $(ROOT_DIR)
REPORT_DIR 			:= $(BUILD_DIR)/build/reports
TEST_RESULTS_DIR    := $(REPORT_DIR)/test-results
COVERAGE   			:= $(REPORT_DIR)/coverage.out
JUNIT     			:= $(TEST_RESULTS_DIR)/go-test.xml
GOVULNCHECK_VERSION := v1.1.4

# NOTE: Intentionally using "." to also format generated files
GO_PKG_DIRS = .

define go_test
	go test ./... \
		-coverprofile=$(COVERAGE) \
		-v -json | go-junit-report > $(JUNIT)
endef

.PHONY: mod-tidy
mod-tidy:
	go mod tidy

.PHONY: clean
clean:
	@echo "🧹 Cleaning build artifacts..."
	rm -rf build

.PHONY: tools
tools:
	@echo "🔧 Installing tools..."
	go install mvdan.cc/gofumpt@v0.6.0
	go install golang.org/x/tools/cmd/goimports@v0.16.0
	go install github.com/jstemmer/go-junit-report/v2@v2.1.0
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.0.0

.PHONY: generate
generate:
	@echo ">> Generating contracts"
	go generate ./...
	@echo ">> Staging generated files"
	git status --porcelain | grep '\.gen\.go$$' | awk '{print $$2}' | xargs -r git add

.PHONY: generate-check
generate-check:
	go generate ./...
	git diff --exit-code

# -------------------------------------------------
# Format
# -------------------------------------------------
.PHONY: fmt
fmt:
	@echo "🎨 Formatting..."
	@if [ -n "$(GO_PKG_DIRS)" ]; then \
		echo "→ gofumpt"; \
		gofumpt -w $(GO_PKG_DIRS); \
		echo "→ goimports"; \
		goimports -w $(GO_PKG_DIRS); \
	else \
		echo "ℹ️ No Go packages found, skipping fmt"; \
	fi

# -------------------------------------------------
# Vet
# -------------------------------------------------
.PHONY: vet
vet:
	@echo "🔍 Vetting..."
	go vet ./...

# -------------------------------------------------
# Prepare
# -------------------------------------------------
.PHONY: prepare
prepare:
	@mkdir -p $(REPORT_DIR) $(TEST_RESULTS_DIR)

# -------------------------------------------------
# Test
# -------------------------------------------------
.PHONY: test
test: prepare
	@echo "🧪 Testing..."
	$(call go_test)

.PHONY: test-ci
test-ci: prepare
	@echo "🧪 Testing (CI)..."
	$(call go_test)

# -------------------------------------------------
# Lint
# -------------------------------------------------
.PHONY: lint
lint: prepare
	@echo "🧹 Linting..."
	golangci-lint run ./...

# -------------------------------------------------
# Vulnerability scan
# -------------------------------------------------
.PHONY: vuln
vuln:
	@echo "🔒 Running govulncheck..."
	go run golang.org/x/vuln/cmd/govulncheck@$(GOVULNCHECK_VERSION) ./...

# -------------------------------------------------
# Full checks
# -------------------------------------------------
.PHONY: check
check: mod-tidy fmt vet lint test
	@echo "✅ All local checks passed"

.PHONY: check-ci
check-ci: lint test-ci vuln
	@echo "✅ All CI checks passed"

.PHONY: ensure-clean
ensure-clean:
	@if [ -n "$$(git status --porcelain)" ]; then \
		echo "❌ Working tree is dirty. Commit or stash your changes before releasing."; \
		git status --short; \
		exit 1; \
	fi

# -------------------------------------------------
# Sonar
# -------------------------------------------------
sonar: check-ci
	@echo "📊 Running Sonar analysis..."
	sonar-scanner

sonar-ci:
	@echo "📊 Running Sonar analysis..."
	sonar-scanner
