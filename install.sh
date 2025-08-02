#!/bin/bash

# Pulsarship Installation Script (Quiet Completions)

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

# Ensure Make or Go is installed
if ! command_exists make && ! command_exists go; then
    print_error "Neither Make nor Go is installed. Please install one to proceed."
    exit 1
fi

# Clone Pulsarship repo into temp dir
TMP_DIR=$(mktemp -d -t pulsarship-XXXXXX)
print_info "🚀 Cloning the Pulsarship repository into $TMP_DIR..."
git clone https://github.com/xeyossr/pulsarship "$TMP_DIR"
cd "$TMP_DIR"

# Build Pulsarship
if command_exists make; then
    make install
else
    go build -ldflags "\
  -X 'main.version=$(git describe --tags --abbrev=0)' \
  -X 'main.tag=$(git describe --tags --abbrev=0)' \
  -X 'main.commit=$(git rev-parse --short HEAD)' \
  -X 'main.buildTime=$(date -u +%Y-%m-%dT%H:%M:%SZ)' \
  -X 'main.buildEnv=$(go version)'" \
        -o pulsarship .
    sudo install -Dm755 pulsarship "/usr/bin/pulsarship"
fi

# Install completions silently to system-wide paths
COMPLETION_DIRS=(
    "/usr/share/bash-completion/completions"
    "/usr/share/fish/vendor_completions.d"
    "/usr/share/zsh/site-functions"
)

for dir in "${COMPLETION_DIRS[@]}"; do
    if [ ! -d "$dir" ]; then
        sudo install -dm755 "$dir"
    fi
done

# Install completions for Bash, Fish, and Zsh silently
sudo bash -c 'pulsarship completion bash > /usr/share/bash-completion/completions/pulsarship'
sudo bash -c 'pulsarship completion fish > /usr/share/fish/vendor_completions.d/pulsarship.fish'
sudo bash -c 'pulsarship completion zsh > /usr/share/zsh/site-functions/_pulsarship'

print_success "🎉 Pulsarship installation completed!"