#!/bin/bash

set -e

# Detect changed folders
detect_changed_folders() {
  git fetch origin main
  echo "$(git diff --name-only origin/main | awk -F'/' '{print $1}' | sort -u)"
}

install_sonar_cloud() {
    curl -O https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-6.2.1.4610-linux-x64.zip
    unzip *.zip
    export PATH=$PATH:"${GITHUB_WORKSPACE}/sonar-scanner-6.2.1.4610-linux-x64/bin"
    echo "Done: Download and Unzip Sonar Cloud"
}

# Detect language based on folder contents
detect_language() {
  local folder=$1
  if [[ -f "$folder/pom.xml" ]]; then
    echo "java"
  elif [[ -f "$folder/package.json" ]]; then
    echo "js"
  elif ls "$folder"/*.go >/dev/null 2>&1; then
    echo "go"
  else
    echo "unknown"
  fi
}

# Main CI process
run_ci() {
  local folder=$1
  local language=$2

  cd "$folder"
  
  case $language in
    java)
    #   ./gradlew sonarqube -Dsonar.login=$SONAR_TOKEN
      mvn clean install
      ;;
    js | ts)
      npm install
      npm run lint
      ;;
    go)
      golangci-lint run
      go test -coverprofile coverage.out ./...
      go test -json ./... > test-report.out
      ;;
    *)
      echo "No CI steps for $language in $folder, skipping."
      ;;
  esac
  sonar-scanner 
    \ -Dsonar.token=$PSON_TOKEN
    \ -Dsonar.sources=. 
    \ -Dsonar.host.url=https://sonarcloud.io

  # Build and push Docker image
#   if [[ -f "Dockerfile" ]]; then
#     docker build -t $DOCKER_USERNAME/$folder:latest .
#     echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
#     docker push "$DOCKER_USERNAME/$folder:latest"
#   fi

  cd ..
}

# Run CI for each changed folder
CHANGED_FOLDERS=$(detect_changed_folders)

if [[ -z "$CHANGED_FOLDERS" ]]; then
  echo "No relevant changes detected."
  exit 0
fi

echo "Detected changed folders: $CHANGED_FOLDERS"

install_sonar_cloud

for folder in $CHANGED_FOLDERS; do
  if [[ ! -d "$folder" ]]; then
    echo "Folder $folder does not exist, skipping."
    continue
  fi

  if [[ "$folder" ==  ".github" ]]; then
    continue
  fi

  language=$(detect_language "$folder")
  echo "Detected language for $folder: $language"

  if [[ "$language" == "unknown" ]]; then
    echo "Unknown language in $folder, skipping."
    continue
  fi

  run_ci "$folder" "$language"
done
