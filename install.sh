#!/usr/bin/env bash
set -e

APP_NAME="gopulse"
BIN_DIR="bin"
INSTALL_PATH="/usr/local/bin/${APP_NAME}"

# --- Colors ---
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${GREEN}🚀 Installing ${APP_NAME}...${NC}"

# --- Check Go installation ---
if ! command -v go >/dev/null 2>&1; then
  echo -e "${RED}❌ Go is not installed. Please install Go first.${NC}"
  exit 1
fi

# --- Build the binary ---
echo -e "${YELLOW}🏗️  Building ${APP_NAME}...${NC}"
mkdir -p "${BIN_DIR}"
go build -o "${BIN_DIR}/${APP_NAME}" ./cmd/${APP_NAME}

# --- Install the binary ---
echo -e "${YELLOW}📦 Installing to ${INSTALL_PATH}...${NC}"
sudo install -m755 "${BIN_DIR}/${APP_NAME}" "${INSTALL_PATH}"

echo -e "${GREEN}✅ ${APP_NAME} installed successfully!${NC}"
echo -e "${GREEN}👉 Run '${APP_NAME}' to start using it.${NC}"
