# Shell script to create a pre-release branch
git checkout main
git pull origin main
latest_tag=$(npx semantic-release --dry-run | grep "The next release version is" | awk '{print $6}')
pre_release_branch="prerelease/v${latest_tag}-rc.1"
git checkout -b $pre_release_branch
git push origin $pre_release_branch


git tag -l 'v?.?.?' --sort=creatordate | tail -n 1