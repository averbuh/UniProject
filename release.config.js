module.exports = {
    "branches": ['main',
    {
      "name": 'release*-alpha',
      "prerelease": true
    },
    {
      "name": 'release/v.*-beta',
      "prerelease": true
    },
  ],
    "plugins": [
      [
        "@semantic-release/commit-analyzer",
        {
          "preset": "conventionalcommits",
          "releaseRules": [
            { "type": "feat", "release": "minor" },
            { "type": "fix", "release": "patch" },
            { "type": "BREAKING CHANGE", "release": "major" }
          ]
        }
      ],
      "@semantic-release/release-notes-generator",
      "@semantic-release/changelog",
      "@semantic-release/github",
      "@semantic-release/git"
    ]
}