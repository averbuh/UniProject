#!/bin/bash

# Function to increment a specific part of the version (major, minor, patch)
function inc_version_part {
  local part="$1"
  local version="$2"
  IFS=. read -r major minor patch <<< "$version"

  temp_patch=$((patch++))
  # Handle overflow for patch version (increment minor if necessary)
  if [[ "$part" == "patch" && $temp_patch -gt 9 ]]; then
    patch=0
    minor=$((minor + 1))
  fi



  # Handle overflow for minor version (increment major if necessary)
  if [[ "$part" == "minor" && $(($minor++)) -gt 9 ]]; then
    minor=0
    major=$((major + 1))
  fi

  # Update the specific part and format the new version
  if [[ "$part" == "major" ]]; then
    echo "$major.$minor.$patch"
  elif [[ "$part" == "minor" ]]; then
    echo "$major.$minor.$patch"
  else
    echo "$major.$minor.$patch"
  fi
}

# Get the current version from the latest git tag (assuming tags follow "vX.Y.Z" format)
current_version=$(git describe --abbrev=0 --tags)

# Check if there are no tags (use initial version)
if [[ -z "$current_version" ]]; then
  current_version="0.0.0"
fi

# Remove the leading "v" from the version string
current_version=${current_version#"v"}

# Get the part to increment (major, minor, or patch)
part_to_increment="$1"  # This will be passed as an argument to the script

# Check for valid part argument
if [[ ! "$part_to_increment" =~ ^(major|minor|patch)$ ]]; then
  echo "Invalid part argument. Please use 'major', 'minor', or 'patch'."
  exit 1
fi

# Call the function to increment the version
new_version=$(inc_version_part "$part_to_increment" "$current_version")

# Create a new git tag with the incremented version prefixed with "v"
new_tag="v$new_version"

# Check if a tag with the new version already exists
if git describe --exact-match --tags "$new_tag" >/dev/null 2>&1; then
  echo "Error: Tag '$new_tag' already exists."
  exit 1
fi

# Confirm with user (optional)
read -p "Create tag '$new_tag'? (y/N) " confirmation
if [[ ! "$confirmation" =~ ^[Yy]$ ]]; then
  echo "Tag creation cancelled."
  exit 0
fi

# Add all changes to the staging area (optional, modify if needed)
git add .

# Create a commit message suggesting the version bump (optional, modify if needed)
git commit -m "Bump version to $new_version"

# Create the new git tag
git tag "$new_tag"

# Print success message
echo "Successfully created tag: $new_tag"

