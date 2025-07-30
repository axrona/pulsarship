#!/bin/bash

# Pulsarship Installation Script

# Function to print error messages with red color
function print_error {
    echo -e "\033[1;31m❌ ERROR: $1\033[0m"
}

# Function to print success messages with green color
function print_success {
    echo -e "\033[1;32m✅ SUCCESS: $1\033[0m"
}

# Function to print info messages with blue color
function print_info {
    echo -e "\033[1;34mℹ️ INFO: $1\033[0m"
}

# Function to print warning messages with yellow color
function print_warning {
    echo -e "\033[1;33m⚠️ WARNING: $1\033[0m"
}

# Function to check if a command exists
function command_exists {
    command -v "$1" &>/dev/null
}

# Ensure Git is installed
if ! command_exists git; then
    print_error "Git is not installed. Please install Git first."
    exit 1
fi

# Ensure Make is installed (if not, fall back to Go build)
if ! command_exists make; then
    print_warning "Make is not installed. Falling back to 'go build' for compilation."
    if ! command_exists go; then
        print_error "Go is not installed either. Please install Go or Make to proceed."
        exit 1
    fi
fi

# Clear screen before starting the Pulsarship installation
clear
print_info "🚀 Cloning the Pulsarship repository..."

# Clone Pulsarship repository and install
git clone https://github.com/xeyossr/pulsarship
cd pulsarship

# If Make is not available, use Go build
if command_exists make; then
    print_info "⚡ Running 'make install'..."
    make install
else
    print_info "⚡ Running 'go build'..."
    go build -ldflags "\
  -X 'main.version=$(git describe --tags --abbrev=0)' \
  -X 'main.tag=$(git describe --tags)' \
  -X 'main.commit=$(git rev-parse --short HEAD)' \
  -X 'main.buildTime=$(date -u +%Y-%m-%dT%H:%M:%SZ)' \
  -X 'main.buildEnv=$(go version)'" \
        -o pulsarship .

    print_info "⚡ Installing to /usr/bin ..."
    sudo install -Dm755 pulsarship "/usr/bin/pulsarship"
fi

# Clear screen and display final success message
clear
print_success "🎉 Pulsarship installation completed!"
