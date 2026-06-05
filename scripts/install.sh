#!/usr/bin/env bash
set -euo pipefail

# ============================================================================
# ciberbal-ai — Install Script
# One command to provision your pentesting AI stack on any OS.
#
# Usage:
#   curl -sL https://raw.githubusercontent.com/Balthael/ciberbal-ai/main/scripts/install.sh | bash
#
# Or download and run:
#   curl -sLO https://raw.githubusercontent.com/Balthael/ciberbal-ai/main/scripts/install.sh
#   chmod +x install.sh
#   ./install.sh
# ============================================================================

GITHUB_OWNER="Balthael"
GITHUB_REPO="ciberbal-ai"
BINARY_NAME="ciberbal-ai"
BREW_TAP="Balthael/homebrew-tap"

# ============================================================================
# Color support
# ============================================================================

setup_colors() {
    if [ -t 1 ] && [ "${TERM:-}" != "dumb" ]; then
        RED='\033[0;31m'
        GREEN='\033[0;32m'
        YELLOW='\033[1;33m'
        BLUE='\033[0;34m'
        CYAN='\033[0;36m'
        BOLD='\033[1m'
        DIM='\033[2m'
        NC='\033[0m'
    else
        RED='' GREEN='' YELLOW='' BLUE='' CYAN='' BOLD='' DIM='' NC=''
    fi
}

# ============================================================================
# Logging helpers
# ============================================================================

info()    { echo -e "${BLUE}[info]${NC}    $*"; }
success() { echo -e "${GREEN}[ok]${NC}      $*"; }
warn()    { echo -e "${YELLOW}[warn]${NC}    $*"; }
error()   { echo -e "${RED}[error]${NC}   $*" >&2; }
fatal()   { error "$@"; exit 1; }
step()    { echo -e "\n${CYAN}${BOLD}==>${NC} ${BOLD}$*${NC}"; }

# ============================================================================
# GitHub auth helpers
# ============================================================================

github_token() {
    if [ -n "${GH_TOKEN:-}" ]; then
        echo "$GH_TOKEN"
        return 0
    fi
    if [ -n "${GITHUB_TOKEN:-}" ]; then
        echo "$GITHUB_TOKEN"
        return 0
    fi
    return 1
}

github_curl() {
    local token
    if token="$(github_token 2>/dev/null)"; then
        curl -H "Authorization: Bearer ${token}" -H "Accept: application/vnd.github+json" "$@"
        return
    fi
    curl "$@"
}

# ============================================================================
# Help
# ============================================================================

show_help() {
    cat <<EOF
${BOLD}ciberbal-ai installer${NC}

Usage: install.sh [OPTIONS]

Options:
  --method METHOD   Force install method: brew, go, binary, source (default: auto-detect)
  --dir DIR         Custom install directory for binary method
  --insecure        Skip checksum verification (warning: reduces security)
  -h, --help        Show this help

Available install methods:
  1. source  — build from the local cloned repository
  2. brew    — Homebrew tap on macOS when available
  3. binary  — pre-built binary from GitHub Releases (default on Linux)
  4. go      — go install from module source (use --method go)

Notes:
  - The binary method only requires curl and tar. It does NOT require git.
  - This script installs only the ciberbal-ai binary and does not modify
    system packages. Agent dependencies (Node.js, etc.) are configured
    interactively when you first run 'ciberbal-ai'.
  - Checksum verification is fail-closed by default. Use --insecure to skip.

Examples:
  curl -sL https://raw.githubusercontent.com/${GITHUB_OWNER}/${GITHUB_REPO}/main/scripts/install.sh | bash
  ./install.sh --method binary
  ./install.sh --method binary --dir \$HOME/.local/bin
  ./install.sh --method binary --insecure

EOF
}

# ============================================================================
# Platform detection
# ============================================================================

detect_platform() {
    local uname_os uname_arch

    uname_os="$(uname -s)"
    uname_arch="$(uname -m)"

    case "$uname_os" in
        Darwin) OS="darwin"; OS_LABEL="macOS"; GORELEASER_OS="darwin" ;;
        Linux)  OS="linux";  OS_LABEL="Linux"; GORELEASER_OS="linux" ;;
        *)      fatal "Unsupported OS: $uname_os. Only macOS and Linux are supported." ;;
    esac

    case "$uname_arch" in
        x86_64|amd64)   ARCH="amd64" ;;
        arm64|aarch64)  ARCH="arm64" ;;
        *)              fatal "Unsupported architecture: $uname_arch. Only amd64 and arm64 are supported." ;;
    esac

    success "Platform: ${OS_LABEL} (${OS}/${ARCH})"
}

# ============================================================================
# GoReleaser archive naming
#
# From .goreleaser.yaml:
#   name_template: "{{ .ProjectName }}_{{ .Version }}_{{ .Os }}_{{ .Arch }}"
#
# GoReleaser v2 {{ .Os }} produces GOOS values (lowercase: darwin, linux)
# GoReleaser {{ .Arch }} produces GOARCH values (amd64, arm64)
# Examples:
#   ciberbal-ai_1.0.0_darwin_arm64.tar.gz
#   ciberbal-ai_1.0.0_linux_amd64.tar.gz
# ============================================================================

get_archive_name() {
    local version="$1"
    echo "${BINARY_NAME}_${version}_${GORELEASER_OS}_${ARCH}.tar.gz"
}

# ============================================================================
# Prerequisites
# ============================================================================

check_prerequisites() {
    step "Checking prerequisites"

    local missing=()

    case "${INSTALL_METHOD:-}" in
        binary)
            command -v curl &>/dev/null || missing+=("curl")
            command -v tar &>/dev/null || missing+=("tar")
            if [ "$INSECURE" != true ] && ! command -v sha256sum &>/dev/null && ! command -v shasum &>/dev/null; then
                missing+=("sha256sum or shasum")
            fi
            ;;
        brew)
            command -v brew &>/dev/null || missing+=("brew")
            ;;
        source|go)
            command -v go &>/dev/null || missing+=("go")
            ;;
        *)
            fatal "Install method was not resolved before prerequisite checks"
            ;;
    esac

    if [ ${#missing[@]} -gt 0 ]; then
        fatal "Missing required tools: ${missing[*]}. Please install them and try again."
    fi

    success "Prerequisites satisfied for method: ${INSTALL_METHOD}"
}

# ============================================================================
# Install method detection
# ============================================================================

detect_install_method() {
    if [ -n "${FORCE_METHOD:-}" ]; then
        case "$FORCE_METHOD" in
            brew|go|binary|source) INSTALL_METHOD="$FORCE_METHOD" ;;
            *) fatal "Unknown install method: $FORCE_METHOD. Use: brew, go, binary, or source" ;;
        esac
        info "Using forced method: $INSTALL_METHOD"
        return
    fi

    step "Detecting best install method"

    # Priority: source > macOS brew > binary > go
    # Source build is ideal when running from a cloned repository because it
    # avoids waiting for GitHub Releases or module-path migration.
    # Brew handles upgrades natively on macOS, but Linux defaults to the
    # release binary so a random Linux instance does not depend on Homebrew taps.
    # Binary download from GitHub Releases is always up-to-date.
    # go install is last resort because the Go module proxy can lag
    # behind new tags for up to 30 minutes, causing @latest to install
    # a stale version.
    local repo_root_detected=false
    if find_repo_root >/dev/null 2>&1; then
        repo_root_detected=true
    fi

    if [ "$repo_root_detected" = true ]; then
        if command -v go &>/dev/null; then
            INSTALL_METHOD="source"
            success "Local repository checkout detected — will build from source"
        else
            fatal "Local ciberbal-ai checkout detected, but Go is not available in PATH. Install Go first (for example: sudo apt install -y golang), then run ./scripts/install.sh again. To intentionally install the older published release instead, pass --method binary explicitly."
        fi
    elif [ "${OS:-}" = "darwin" ] && command -v brew &>/dev/null; then
        INSTALL_METHOD="brew"
        success "Homebrew found on macOS — will install via brew tap"
    else
        INSTALL_METHOD="binary"
        info "Will download pre-built binary from GitHub Releases"
    fi
}

# ============================================================================
# Install via Homebrew
# ============================================================================

install_brew() {
    step "Installing via Homebrew"

    # Always refresh the tap to pick up new releases
    info "Refreshing ${BREW_TAP}..."
    brew untap "$BREW_TAP" 2>/dev/null || true
    if ! brew tap "$BREW_TAP"; then
        fatal "Failed to tap $BREW_TAP"
    fi

    if brew list "$BINARY_NAME" &>/dev/null; then
        info "Already installed, upgrading ${BINARY_NAME}..."
        if brew upgrade "$BINARY_NAME" 2>/dev/null; then
            success "Upgraded ${BINARY_NAME} via Homebrew"
        else
            # "already up-to-date" also exits non-zero on some brew versions
            success "${BINARY_NAME} is already at the latest version"
        fi
    else
        info "Installing ${BINARY_NAME}..."
        if brew install "$BINARY_NAME"; then
            success "Installed ${BINARY_NAME} via Homebrew"
        else
            fatal "Failed to install ${BINARY_NAME} via Homebrew"
        fi
    fi
}

# ============================================================================
# Install via go install
# ============================================================================

install_go() {
    step "Installing via go install"

    local go_package="github.com/${GITHUB_OWNER,,}/${GITHUB_REPO}/cmd/${BINARY_NAME}@latest"

    info "Running: go install ${go_package}"
    if ! go install "$go_package"; then
        fatal "Failed to install via go install. Make sure Go is properly configured."
    fi

    # Verify GOBIN / GOPATH/bin is in PATH
    local gobin
    gobin="$(go env GOBIN)"
    if [ -z "$gobin" ]; then
        gobin="$(go env GOPATH)/bin"
    fi

    if [[ ":$PATH:" != *":$gobin:"* ]]; then
        warn "${gobin} is not in your PATH"
        warn "Add this to your shell profile: export PATH=\"\$PATH:${gobin}\""
    fi

    success "Installed ${BINARY_NAME} via go install"
}

# ============================================================================
# Install via local source build
# ============================================================================

find_repo_root() {
    local script_source script_dir candidate
    script_source="${BASH_SOURCE[0]:-}"
    if [ -z "$script_source" ] || [ ! -f "$script_source" ]; then
        return 1
    fi

    script_dir="$(cd -- "$(dirname -- "$script_source")" && pwd)"
    candidate="$(cd -- "${script_dir}/.." && pwd)"

    if [ -f "${candidate}/go.mod" ] && [ -f "${candidate}/cmd/${BINARY_NAME}/main.go" ]; then
        echo "$candidate"
        return 0
    fi

    return 1
}

install_source() {
    step "Installing from local source"

    if ! command -v go &>/dev/null; then
        fatal "Go is required for --method source, but 'go' was not found in PATH. Install Go or use a published release."
    fi

    local repo_root
    if ! repo_root="$(find_repo_root)"; then
        fatal "Could not detect a local ciberbal-ai repository checkout. Run this from a cloned repo or use --method binary once releases exist."
    fi

    info "Building from local repository: ${repo_root}"

    local tmpdir
    tmpdir="$(mktemp -d)"
    trap '[ -n "${tmpdir:-}" ] && rm -rf "$tmpdir"' EXIT

    if ! go build -o "${tmpdir}/${BINARY_NAME}" "${repo_root}/cmd/${BINARY_NAME}"; then
        fatal "Failed to build ${BINARY_NAME} from local source"
    fi

    local install_dir="${INSTALL_DIR:-}"
    if [ -z "$install_dir" ]; then
        if [ -d "/usr/local/bin" ] && [ -w "/usr/local/bin" ]; then
            install_dir="/usr/local/bin"
        elif [ "$(id -u)" = "0" ]; then
            install_dir="/usr/local/bin"
        else
            install_dir="${HOME}/.local/bin"
        fi
    fi

    mkdir -p "$install_dir"
    info "Installing to ${install_dir}/${BINARY_NAME}..."
    if cp "${tmpdir}/${BINARY_NAME}" "${install_dir}/${BINARY_NAME}" 2>/dev/null; then
        chmod +x "${install_dir}/${BINARY_NAME}"
    elif command -v sudo &>/dev/null; then
        warn "Permission denied. Trying with sudo..."
        sudo cp "${tmpdir}/${BINARY_NAME}" "${install_dir}/${BINARY_NAME}"
        sudo chmod +x "${install_dir}/${BINARY_NAME}"
    else
        fatal "Cannot write to ${install_dir}. Run with sudo or use --dir to specify a writable directory."
    fi

    success "Installed ${BINARY_NAME} from local source to ${install_dir}/${BINARY_NAME}"

    if [[ ":$PATH:" != *":${install_dir}:"* ]]; then
        warn "${install_dir} is not in your PATH"
        warn "Add this to your shell profile: export PATH=\"\$PATH:${install_dir}\""
    fi
}

# ============================================================================
# Install via binary download
# ============================================================================

get_latest_version() {
    local url="https://api.github.com/repos/${GITHUB_OWNER}/${GITHUB_REPO}/releases/latest"

    info "Fetching latest release from GitHub..."

    local response
    response="$(github_curl -sL -w "\n%{http_code}" "$url")" || fatal "Failed to fetch latest release"

    local http_code body
    http_code="$(echo "$response" | tail -n1)"
    body="$(echo "$response" | sed '$d')"

    if [ "$http_code" != "200" ]; then
        if [ "$http_code" = "403" ]; then
            warn "GitHub API returned HTTP 403. Trying non-API latest-release redirect..."
            if LATEST_VERSION="$(get_latest_version_from_redirect)"; then
                VERSION_NUMBER="${LATEST_VERSION#v}"
                success "Latest version: ${LATEST_VERSION}"
                return 0
            fi
            if find_repo_root >/dev/null 2>&1 && command -v go &>/dev/null; then
                warn "GitHub API is rate limited, but a local repository checkout is available. Falling back to local source build."
                INSTALL_METHOD="source"
                return 1
            fi
            fatal "GitHub API returned HTTP 403 and the latest release could not be resolved without the API. Try again later, export GH_TOKEN/GITHUB_TOKEN, or use --method source/go"
        fi
        if [ "$http_code" = "404" ]; then
            if find_repo_root >/dev/null 2>&1 && command -v go &>/dev/null; then
                warn "No GitHub release exists yet for ${GITHUB_OWNER}/${GITHUB_REPO}. Falling back to local source build from this cloned repository."
                INSTALL_METHOD="source"
                return 1
            fi
            if github_token >/dev/null 2>&1; then
                fatal "GitHub Releases returned HTTP 404 even with a token. Verify the repository, release visibility, and token permissions."
            fi
            fatal "GitHub Releases returned HTTP 404. If this repository is private, export GH_TOKEN or GITHUB_TOKEN before running the installer. Otherwise, no published release exists yet for ${GITHUB_OWNER}/${GITHUB_REPO}."
        fi
        fatal "GitHub API returned HTTP $http_code. Rate limited? Try again later or use --method source/go"
    fi

    # Extract tag_name — works without jq
    LATEST_VERSION="$(echo "$body" | sed -n 's/.*"tag_name"[[:space:]]*:[[:space:]]*"\([^"]*\)".*/\1/p' | head -1)"

    if [ -z "$LATEST_VERSION" ]; then
        fatal "Could not determine latest version from GitHub API response"
    fi

    # Strip leading 'v' for archive naming (goreleaser uses version without v prefix)
    VERSION_NUMBER="${LATEST_VERSION#v}"

    success "Latest version: ${LATEST_VERSION}"
}

get_latest_version_from_redirect() {
    local url="https://github.com/${GITHUB_OWNER}/${GITHUB_REPO}/releases/latest"
    local headers location

    headers="$(github_curl -sI "$url")" || return 1
    location="$(printf '%s\n' "$headers" | tr -d '\r' | sed -n 's/^[Ll]ocation:[[:space:]]*.*\/tag\/\(v[^[:space:]]*\).*/\1/p' | tail -n1)"

    if [ -z "$location" ]; then
        return 1
    fi

    printf '%s\n' "$location"
}

install_binary() {
    step "Installing pre-built binary"

    if ! get_latest_version; then
        if [ "${INSTALL_METHOD:-}" = "source" ]; then
            install_source
            return
        fi
        fatal "Could not determine latest release"
    fi

    local archive_name
    archive_name="$(get_archive_name "$VERSION_NUMBER")"
    local download_url="https://github.com/${GITHUB_OWNER}/${GITHUB_REPO}/releases/download/${LATEST_VERSION}/${archive_name}"
    local checksums_url="https://github.com/${GITHUB_OWNER}/${GITHUB_REPO}/releases/download/${LATEST_VERSION}/checksums.txt"

    # Create temp directory — clean up on exit
    local tmpdir
    tmpdir="$(mktemp -d)"
    trap '[ -n "${tmpdir:-}" ] && rm -rf "$tmpdir"' EXIT

    # Download archive
    info "Downloading ${archive_name}..."
    if ! github_curl -sfL -o "${tmpdir}/${archive_name}" "$download_url"; then
        fatal "Failed to download ${download_url}"
    fi

    # Verify file was actually downloaded (not a 404 HTML page)
    local file_size
    file_size="$(wc -c < "${tmpdir}/${archive_name}" | tr -d '[:space:]')"
    if [ "$file_size" -lt 1000 ]; then
        fatal "Downloaded file is suspiciously small (${file_size} bytes). Archive may not exist for this platform."
    fi

    success "Downloaded ${archive_name} (${file_size} bytes)"

    # Download and verify checksum (fail-closed by default)
    info "Verifying checksum..."
    if [ "$INSECURE" = true ]; then
        warn "--insecure: skipping checksum verification"
    else
        # Require a working checksum tool before downloading checksums.txt
        local checksum_tool=""
        if command -v sha256sum &>/dev/null; then
            checksum_tool="sha256sum"
        elif command -v shasum &>/dev/null; then
            checksum_tool="shasum"
        else
            fatal "No sha256sum or shasum found. Cannot verify download integrity.\nInstall one of these tools or use --insecure to skip (not recommended)."
        fi

        if ! github_curl -sL -o "${tmpdir}/checksums.txt" "$checksums_url"; then
            fatal "Could not download checksums.txt from ${checksums_url}.\nUse --insecure to skip checksum verification (not recommended)."
        fi

        local expected_checksum
        expected_checksum="$(grep -F "${archive_name}" "${tmpdir}/checksums.txt" 2>/dev/null | awk '{print $1}' || true)"

        if [ -z "$expected_checksum" ]; then
            fatal "Archive '${archive_name}' not found in checksums.txt.\nUse --insecure to skip checksum verification (not recommended)."
        fi

        local actual_checksum
        if [ "$checksum_tool" = "sha256sum" ]; then
            actual_checksum="$(sha256sum "${tmpdir}/${archive_name}" | awk '{print $1}')"
        else
            actual_checksum="$(shasum -a 256 "${tmpdir}/${archive_name}" | awk '{print $1}')"
        fi

        if [ "$actual_checksum" != "$expected_checksum" ]; then
            fatal "Checksum mismatch!\n  Expected: ${expected_checksum}\n  Got:      ${actual_checksum}"
        fi
        success "Checksum verified"
    fi

    # Extract binary
    info "Extracting ${BINARY_NAME}..."
    if ! tar -xzf "${tmpdir}/${archive_name}" -C "$tmpdir"; then
        fatal "Failed to extract archive"
    fi

    if [ ! -f "${tmpdir}/${BINARY_NAME}" ]; then
        fatal "Binary '${BINARY_NAME}' not found in archive"
    fi

    # Determine install directory
    local install_dir="${INSTALL_DIR:-}"

    if [ -z "$install_dir" ]; then
        if [ -d "/usr/local/bin" ] && [ -w "/usr/local/bin" ]; then
            install_dir="/usr/local/bin"
        elif [ "$(id -u)" = "0" ]; then
            install_dir="/usr/local/bin"
        else
            install_dir="${HOME}/.local/bin"
        fi
    fi

    # Create install dir if needed
    mkdir -p "$install_dir"

    # Install binary
    info "Installing to ${install_dir}/${BINARY_NAME}..."
    if cp "${tmpdir}/${BINARY_NAME}" "${install_dir}/${BINARY_NAME}" 2>/dev/null; then
        chmod +x "${install_dir}/${BINARY_NAME}"
    elif command -v sudo &>/dev/null; then
        warn "Permission denied. Trying with sudo..."
        sudo cp "${tmpdir}/${BINARY_NAME}" "${install_dir}/${BINARY_NAME}"
        sudo chmod +x "${install_dir}/${BINARY_NAME}"
    else
        fatal "Cannot write to ${install_dir}. Run with sudo or use --dir to specify a writable directory."
    fi

    success "Installed ${BINARY_NAME} to ${install_dir}/${BINARY_NAME}"

    # Check if install dir is in PATH
    if [[ ":$PATH:" != *":${install_dir}:"* ]]; then
        warn "${install_dir} is not in your PATH"
        echo ""
        warn "Add this to your shell profile (~/.bashrc, ~/.zshrc, etc.):"
        echo -e "  ${DIM}export PATH=\"\$PATH:${install_dir}\"${NC}"
        echo ""
    fi
}

# ============================================================================
# Verify installation
# ============================================================================

verify_installation() {
    step "Verifying installation"

    # Allow PATH changes to take effect
    hash -r 2>/dev/null || true

    if command -v "$BINARY_NAME" &>/dev/null; then
        local version_output
        version_output="$("$BINARY_NAME" version 2>&1 || true)"
        success "${BINARY_NAME} is installed: ${version_output}"
        return 0
    fi

    # Check explicit/custom and common locations even if not in PATH.
    local locations=()
    if [ -n "${INSTALL_DIR:-}" ]; then
        locations+=("${INSTALL_DIR}/${BINARY_NAME}")
    fi
    locations+=(
        "/usr/local/bin/${BINARY_NAME}"
        "${HOME}/.local/bin/${BINARY_NAME}"
    )

    if command -v go &>/dev/null; then
        local gopath
        gopath="$(go env GOPATH 2>/dev/null || true)"
        if [ -n "$gopath" ]; then
            locations+=("${gopath}/bin/${BINARY_NAME}")
        fi
    fi

    for loc in "${locations[@]}"; do
        if [ -n "$loc" ] && [ -x "$loc" ]; then
            local version_output
            version_output="$("$loc" version 2>&1 || true)"
            success "Found ${BINARY_NAME} at ${loc}: ${version_output}"
            warn "Binary location is not in your PATH. Add it to use '${BINARY_NAME}' directly."
            return 0
        fi
    done

    warn "Could not verify installation. You may need to restart your shell."
    return 0
}

# ============================================================================
# Print next steps
# ============================================================================

print_banner() {
    echo ""
    echo -e "${CYAN}${BOLD}"
    echo "   ==============================================="
    echo "                   Ciberbal-ai"
    echo "   ==============================================="
    echo -e "${NC}"
    echo -e "  ${DIM}One command to provision your pentesting AI stack on any OS${NC}"
    echo ""
}

print_next_steps() {
    echo ""
    echo -e "${GREEN}${BOLD}Binary installed!${NC}"
    echo ""
    echo -e "${DIM}This script installed only the ${BINARY_NAME} binary.${NC}"
    echo -e "${DIM}No system packages were modified. Agent dependencies${NC}"
    echo -e "${DIM}(Node.js, npm, etc.) are configured when you run the tool.${NC}"
    echo ""
    echo -e "${BOLD}Next steps:${NC}"
    echo -e "  ${CYAN}1.${NC} Run ${BOLD}${BINARY_NAME}${NC} to start the TUI installer"
    echo -e "  ${CYAN}2.${NC} Select your AI agent(s) and tools to configure"
    echo -e "  ${CYAN}3.${NC} Follow the interactive prompts — dependencies are"
    echo -e "     installed only for the agents you choose"
    echo ""
    echo -e "${DIM}For help: ${BINARY_NAME} --help${NC}"
    echo -e "${DIM}Docs:     https://github.com/${GITHUB_OWNER}/${GITHUB_REPO}${NC}"
    echo ""
}

# ============================================================================
# Main
# ============================================================================

main() {
    setup_colors

    # Parse arguments
    FORCE_METHOD=""
    INSTALL_DIR=""
    INSECURE=false

    while [ $# -gt 0 ]; do
        case "$1" in
            --method)
                [ $# -lt 2 ] && fatal "--method requires an argument"
                FORCE_METHOD="$2"; shift 2
                ;;
            --dir)
                [ $# -lt 2 ] && fatal "--dir requires an argument"
                INSTALL_DIR="$2"; shift 2
                ;;
            --insecure)
                INSECURE=true; shift
                ;;
            -h|--help)
                setup_colors
                show_help
                exit 0
                ;;
            *)
                fatal "Unknown option: $1. Use --help for usage."
                ;;
        esac
    done

    print_banner

    step "Detecting platform"
    detect_platform

    detect_install_method
    check_prerequisites

    case "$INSTALL_METHOD" in
        brew)   install_brew ;;
        go)     install_go ;;
        binary) install_binary ;;
        source) install_source ;;
    esac

    verify_installation
    print_next_steps
}

main "$@"
