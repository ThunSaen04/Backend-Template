#!/bin/bash
# Regenerate all Swagger documentation
# Usage: ./scripts/swagger-gen.sh

echo "==> Generating Global Swagger docs (all modules)..."
swag init --generalInfo main.go --parseDependency --parseInternal --output docs

echo ""
echo "==> Generating Auth module Swagger docs..."
swag init --generalInfo handler/swagger_info.go --dir ./internal/modules/auth --parseDependency --parseInternal --output internal/modules/auth/docs --instanceName auth

echo ""
echo "Done! Swagger docs regenerated."
echo "  Global:  http://localhost:8080/swagger/index.html"
echo "  Auth:    http://localhost:8080/swagger/auth/index.html"
