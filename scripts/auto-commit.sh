#!/usr/bin/env bash
# auto-commit.sh
#
# Script for automatic commits with detailed change descriptions
# Compatible with macOS, Linux and Windows (Git Bash/WSL)
# Generates commits following Conventional Commits based on detected changes

set -euo pipefail

# Detect operating system
case "$(uname -s)" in
    Darwin*)    OS="macOS" ;;
    Linux*)     OS="Linux" ;;
    CYGWIN*|MINGW*|MSYS*) OS="Windows" ;;
    *)          OS="Unknown" ;;
esac

# Colors for output
if [[ "$OS" == "Windows" ]]; then
    # Windows may have issues with colors, use simplified version
    RED='[ERROR]'
    GREEN='[SUCCESS]'
    YELLOW='[WARNING]'
    BLUE='[INFO]'
    NC=''
else
    RED='\033[0;31m'
    GREEN='\033[0;32m'
    YELLOW='\033[1;33m'
    BLUE='\033[0;34m'
    NC='\033[0m'
fi

echo -e "${BLUE}🚀 Auto-commit script - Detected: $OS${NC}"

# Verify that we are in a git repository
if ! git rev-parse --git-dir > /dev/null 2>&1; then
    echo -e "${RED}❌ Error: You are not in a Git repository${NC}"
    exit 1
fi

# Function to analyze file type and suggest scope
get_file_scope() {
    local file="$1"
    case "$file" in
        *.sh)           echo "scripts" ;;
        *.md)           echo "docs" ;;
        *.json|*.yaml|*.yml) echo "config" ;;
        package*.json)  echo "deps" ;;
        .gitignore)     echo "git" ;;
        Dockerfile*)    echo "docker" ;;
        *.go)           echo "go" ;;
        *.js|*.ts)      echo "node" ;;
        *.py)           echo "python" ;;
        *)              echo "misc" ;;
    esac
}

# Function to analyze change type
analyze_change_type() {
    local file="$1"
    local status="$2"
    
    case "$status" in
        "A")    # Added file
            if [[ "$file" == *.md ]]; then
                echo "docs"
            elif [[ "$file" == *.sh ]]; then
                echo "feat"
            else
                echo "feat"
            fi
            ;;
        "M")    # Modified file
            if [[ "$file" == *"test"* ]] || [[ "$file" == *"spec"* ]]; then
                echo "test"
            elif [[ "$file" == *.md ]]; then
                echo "docs"
            elif [[ "$file" == *"fix"* ]] || [[ "$file" == *"bug"* ]]; then
                echo "fix"
            else
                echo "feat"
            fi
            ;;
        "D")    # Deleted file
            echo "remove"
            ;;
        "R")    # Renamed file
            echo "refactor"
            ;;
        *)      echo "chore" ;;
    esac
}

# Get modified files
echo -e "\n${BLUE}📋 Analyzing changes...${NC}"

# Check if there are staged changes
if ! git diff --cached --quiet; then
    echo -e "${BLUE}📦 Changes in staging area:${NC}"
    STAGED_FILES=$(git diff --cached --name-status)
else
    echo -e "${YELLOW}⚠️  No changes in staging area. Adding modified files...${NC}"
    
    # Add modified files automatically
    if git diff --quiet; then
        echo -e "${YELLOW}ℹ️  No changes to commit${NC}"
        exit 0
    fi
    
    git add -A
    STAGED_FILES=$(git diff --cached --name-status)
fi

echo "$STAGED_FILES"

# Analyze changes and generate commit message
declare -A changes_by_type
declare -A changes_by_scope
declare -a detailed_changes

echo -e "\n${BLUE}🔍 Analyzing change types...${NC}"

while IFS=$'\t' read -r status file; do
    # Clean the status (may have extra characters)
    status=$(echo "$status" | tr -d ' ')
    
    change_type=$(analyze_change_type "$file" "$status")
    scope=$(get_file_scope "$file")
    
    # Count changes by type
    changes_by_type["$change_type"]=$((${changes_by_type["$change_type"]:-0} + 1))
    changes_by_scope["$scope"]=$((${changes_by_scope["$scope"]:-0} + 1))
    
    # Generate detailed description
    case "$status" in
        "A") detailed_changes+=("Added $file") ;;
        "M") detailed_changes+=("Modified $file") ;;
        "D") detailed_changes+=("Deleted $file") ;;
        "R"*) detailed_changes+=("Renamed $file") ;;
        *) detailed_changes+=("Changed $file") ;;
    esac
    
done <<< "$STAGED_FILES"

# Determine primary commit type
primary_type="feat"
max_count=0
for type in "${!changes_by_type[@]}"; do
    if [[ ${changes_by_type[$type]} -gt $max_count ]]; then
        max_count=${changes_by_type[$type]}
        primary_type="$type"
    fi
done

# Determine primary scope
primary_scope="misc"
max_count=0
for scope in "${!changes_by_scope[@]}"; do
    if [[ ${changes_by_scope[$scope]} -gt $max_count ]]; then
        max_count=${changes_by_scope[$scope]}
        primary_scope="$scope"
    fi
done

# Generate change summary
total_files=${#detailed_changes[@]}
change_summary=""

if [[ $total_files -eq 1 ]]; then
    change_summary="${detailed_changes[0]}"
else
    change_summary="Updated $total_files files"
    
    # Add details of most common types
    for type in "${!changes_by_type[@]}"; do
        count=${changes_by_type[$type]}
        if [[ $count -gt 1 ]]; then
            change_summary="$change_summary ($count $type changes)"
        fi
    done
fi

# Generate commit message
if [[ "$primary_scope" != "misc" ]]; then
    commit_message="$primary_type($primary_scope): $change_summary"
else
    commit_message="$primary_type: $change_summary"
fi

# Show summary to user
echo -e "\n${BLUE}📊 Change summary:${NC}"
echo "  Total files: $total_files"
echo "  Primary type: $primary_type"
echo "  Primary scope: $primary_scope"
echo ""

echo -e "${BLUE}📝 Change types:${NC}"
for type in "${!changes_by_type[@]}"; do
    echo "  $type: ${changes_by_type[$type]} files"
done

echo -e "\n${BLUE}📂 Affected scopes:${NC}"
for scope in "${!changes_by_scope[@]}"; do
    echo "  $scope: ${changes_by_scope[$scope]} files"
done

echo -e "\n${BLUE}📋 Modified files:${NC}"
for change in "${detailed_changes[@]}"; do
    echo "  - $change"
done

# Generate detailed message
detailed_message="$commit_message

Changes summary:
- Total files modified: $total_files
- Primary change type: $primary_type
- Primary scope: $primary_scope

File changes:"

for change in "${detailed_changes[@]}"; do
    detailed_message="$detailed_message
- $change"
done

# Add system information
detailed_message="$detailed_message

System info:
- OS: $OS
- Date: $(date '+%Y-%m-%d %H:%M:%S')
- Git user: $(git config user.name) <$(git config user.email)>"

echo -e "\n${YELLOW}💡 Proposed commit message:${NC}"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "$commit_message"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"

# Ask user if they want to continue
echo -e "\n${BLUE}Do you want to continue with this commit? [Y/n/e(dit)]${NC}"
read -r response

case "$response" in
    [eE])
        echo -e "${BLUE}✏️  Editing commit message...${NC}"
        
        # Create temporary file with the message
        if [[ "$OS" == "Windows" ]]; then
            temp_file="$(mktemp).txt"
        else
            temp_file=$(mktemp)
        fi
        echo "$detailed_message" > "$temp_file"
        
        # Open editor
        if [[ -n "${EDITOR:-}" ]]; then
            "$EDITOR" "$temp_file"
        elif command -v code >/dev/null 2>&1; then
            code --wait "$temp_file"
        elif command -v nano >/dev/null 2>&1; then
            nano "$temp_file"
        elif command -v vim >/dev/null 2>&1; then
            vim "$temp_file"
        else
            echo -e "${YELLOW}⚠️  No editor found. Using original message.${NC}"
        fi
        
        detailed_message=$(cat "$temp_file")
        rm -f "$temp_file"
        ;;
    [nN])
        echo -e "${YELLOW}❌ Commit cancelled${NC}"
        exit 0
        ;;
    *)
        echo -e "${GREEN}✅ Continuing with commit...${NC}"
        ;;
esac

# Perform the commit
echo -e "\n${BLUE}🚀 Performing commit...${NC}"

if git commit -m "$detailed_message"; then
    echo -e "\n${GREEN}✅ Commit completed successfully!${NC}"
    
    # Show commit information
    echo -e "\n${BLUE}📋 Commit information:${NC}"
    git log -1 --oneline
    
    # Ask if they want to push
    echo -e "\n${BLUE}Do you want to push to remote repository? [y/N]${NC}"
    read -r push_response
    
    case "$push_response" in
        [yY])
            echo -e "${BLUE}📤 Pushing...${NC}"
            current_branch=$(git branch --show-current)
            
            if git push origin "$current_branch"; then
                echo -e "${GREEN}✅ Push completed successfully!${NC}"
            else
                echo -e "${YELLOW}⚠️  Push error. You can do it manually with: git push origin $current_branch${NC}"
            fi
            ;;
        *)
            echo -e "${BLUE}ℹ️  Push skipped. You can do it manually later.${NC}"
            ;;
    esac
    
else
    echo -e "${RED}❌ Commit error${NC}"
    exit 1
fi

echo -e "\n${GREEN}🎉 Process completed!${NC}"
