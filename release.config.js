
module.exports = {
    "branches": ['main',
      {
        name: 'release/*-alpha', // Alpha branches
        prerelease: 'alpha' // Use 'alpha' as the prerelease identifier
      },
      {
        name: 'release/*-beta', // Beta branches like release/v1.2.2-beta
        prerelease: 'beta' // Use 'beta' as the prerelease identifier
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
            { "type": "BREAKING CHANGE", "release": "major" },
            { "type": "rc"}
          ]
        }
      ],
      "@semantic-release/release-notes-generator",
      "@semantic-release/changelog",
      "@semantic-release/github",
      "@semantic-release/git"
    ]
}