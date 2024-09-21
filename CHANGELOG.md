# [1.1.0-beta.1](https://github.com/averbuh/UniProject/compare/v1.0.2...v1.1.0-beta.1) (2024-09-21)


### Bug Fixes

* Enable CloudWatch monitoring for EKS workloads, adjust instance type to t3.medium for worker nodes, and update release branch settings. These changes align resource allocations and deployment configurations, enhancing system monitoring and scalability. ([5e0b4a1](https://github.com/averbuh/UniProject/commit/5e0b4a1df8017db070b78199c5a28fb6a54d9e77))
* Fix branch name configuration for alpha and beta releases ([0407b81](https://github.com/averbuh/UniProject/commit/0407b8155b4fb3c7b79d66aa3b074703193c7d4d))
* Fix missing comma in release configuration file ([9080b99](https://github.com/averbuh/UniProject/commit/9080b99888c0a03dea72981ff358bd5f95542f96))
* Fix setting of prerelease values to boolean type in .releaserc.json configuration file. This change corrects the format to ensure proper prerelease handling. ([00b280a](https://github.com/averbuh/UniProject/commit/00b280afe52ca7d4323b7dd9213fa63ec6d83fec))
* Fix unnecessary commented-out code related to branch filtering in CircleCI configuration. ([ad7e14a](https://github.com/averbuh/UniProject/commit/ad7e14ac86f0d77885703654826c03bc16e06ccb))
* issue related to unnecessary commented code and add clarifying comment in CircleCI configuration file. ([3120b4d](https://github.com/averbuh/UniProject/commit/3120b4d9c5c3eb902da3024575315443a624fb77))
* Refine branch naming conventions for alpha and beta releases ([f861d34](https://github.com/averbuh/UniProject/commit/f861d34e5ed300e048597215cb44275d2e77b0a4))
* some fix ([51babf6](https://github.com/averbuh/UniProject/commit/51babf645f11a7e759c38f9a7cf4de7913cdaa75))
* some fix ([45074ca](https://github.com/averbuh/UniProject/commit/45074ca229672e3a791ef834719c5482429d1e69))


### Features

* Init ([4078178](https://github.com/averbuh/UniProject/commit/407817857ac6bebe3e00706a9b3d55d0e6649324))
* Refactor branch naming conventions and executor configurations ([985675a](https://github.com/averbuh/UniProject/commit/985675a985c8b4b742f6d4015c7dbd0f284882f2)), closes [#ISSUE-123](https://github.com/averbuh/UniProject/issues/ISSUE-123)
* Some changes ([71eec89](https://github.com/averbuh/UniProject/commit/71eec8951f6b8169beeb2a41975e01cd2a485a74))
* test ([77b12c8](https://github.com/averbuh/UniProject/commit/77b12c862beb9facf980f106832cc617c4e9a929))
* test ([3ee8961](https://github.com/averbuh/UniProject/commit/3ee89611939875819a60c229df6463382a49d765))
* test ([866651b](https://github.com/averbuh/UniProject/commit/866651bdaa7d01a48a9350bbc7997c1a3f364929))
* test ([5e02093](https://github.com/averbuh/UniProject/commit/5e0209354576cc66c00afa2d9de5f95faa3457ec))
* test ([32c4745](https://github.com/averbuh/UniProject/commit/32c4745e595624efbe87054b4a5e03f621f25439))
* test ([2195b07](https://github.com/averbuh/UniProject/commit/2195b07dbdbe0d62da914133b5419a54a3df5c1a))
* test ([146036b](https://github.com/averbuh/UniProject/commit/146036b622f891eca2e08b9d54cdf5f16eec7dd2))
* test ([8ae86cb](https://github.com/averbuh/UniProject/commit/8ae86cbb0f79b524519abc186b332c467d263e6a))
* test ([6381ac5](https://github.com/averbuh/UniProject/commit/6381ac5f1cfc9718187ffaff1117bcea7085dd9f))
* test ([40f84df](https://github.com/averbuh/UniProject/commit/40f84df219ee4d6992c0e4cb0094692d3e252716))
* Update CircleCI config to use 'run-recipes-go' flag instead of 'run-pipeline' to align with actual workflow. This change removes unused 'run-frontend' parameter. (#issue123) ([f5097fc](https://github.com/averbuh/UniProject/commit/f5097fc173e41a738464a501bcb648a5a0cb5be2)), closes [#issue123](https://github.com/averbuh/UniProject/issues/issue123)

## [1.2.1](https://github.com/averbuh/UniProject/compare/v1.2.0...v1.2.1) (2024-09-21)


### Bug Fixes

* Enable CloudWatch monitoring for EKS workloads, adjust instance type to t3.medium for worker nodes, and update release branch settings. These changes align resource allocations and deployment configurations, enhancing system monitoring and scalability. ([0b4cbc6](https://github.com/averbuh/UniProject/commit/0b4cbc6a655645006fbd98dd2243d141ba5851a9))

# [1.2.0](https://github.com/averbuh/UniProject/compare/v1.1.0...v1.2.0) (2024-08-13)


### Features

* Init ([83ae3d0](https://github.com/averbuh/UniProject/commit/83ae3d0669b160d05c4bd399e1498897e7beb32a))
* Merge pull request [#42](https://github.com/averbuh/UniProject/issues/42) from averbuh/rc/v1.2.0 ([bcc1390](https://github.com/averbuh/UniProject/commit/bcc139070bb959770d70f73f1c77f690e14dd469))
* Merge pull request [#43](https://github.com/averbuh/UniProject/issues/43) from averbuh/beta ([8511f17](https://github.com/averbuh/UniProject/commit/8511f17e342fa937e7f608c279a5bad6be4f8410))
* Refactor branch naming conventions and executor configurations ([bd6f931](https://github.com/averbuh/UniProject/commit/bd6f931a3333bb3b195225d38de64d2359d78b50)), closes [#ISSUE-123](https://github.com/averbuh/UniProject/issues/ISSUE-123)
* test ([a34e0a3](https://github.com/averbuh/UniProject/commit/a34e0a3b21a8704939c1e22af55b841d52f5baa7))
* test ([8836509](https://github.com/averbuh/UniProject/commit/8836509427b4bc3d2936fae5e916ad5d6874b4e2))
* test ([68ba9b4](https://github.com/averbuh/UniProject/commit/68ba9b456b5d62f88638885e8ac8d76c280faab4))
* test ([b796e17](https://github.com/averbuh/UniProject/commit/b796e172febc129af8c3d2fb367495843d0124ec))
* test ([a29c61a](https://github.com/averbuh/UniProject/commit/a29c61a94c0ecf5ce4c6fb0a17e9d3dca4094ad1))
* test ([33d5e00](https://github.com/averbuh/UniProject/commit/33d5e00963ea90014f3067c5692945c6ab1f8842))
* test ([ffa5263](https://github.com/averbuh/UniProject/commit/ffa5263c73ce6012f115365b29205ce64466163e))
* test ([654915b](https://github.com/averbuh/UniProject/commit/654915b0775e6937fb1d6447830d558eea222671))
* test ([6465f29](https://github.com/averbuh/UniProject/commit/6465f2903cdd2dcf38870e626f1e2f6aecbd3d90))
* test ([e9c3965](https://github.com/averbuh/UniProject/commit/e9c3965779f105b29bc7b4bffde18bd8c7a6f2d3))
* Update CircleCI config to use 'run-recipes-go' flag instead of 'run-pipeline' to align with actual workflow. This change removes unused 'run-frontend' parameter. (#issue123) ([089c99e](https://github.com/averbuh/UniProject/commit/089c99ed4f044ef1a3cc49859277e29913bd53ab)), closes [#issue123](https://github.com/averbuh/UniProject/issues/issue123)

# [1.2.0-beta.1](https://github.com/averbuh/UniProject/compare/v1.1.0...v1.2.0-beta.1) (2024-08-13)


### Features

* Init ([83ae3d0](https://github.com/averbuh/UniProject/commit/83ae3d0669b160d05c4bd399e1498897e7beb32a))
* Merge pull request [#42](https://github.com/averbuh/UniProject/issues/42) from averbuh/rc/v1.2.0 ([bcc1390](https://github.com/averbuh/UniProject/commit/bcc139070bb959770d70f73f1c77f690e14dd469))
* Refactor branch naming conventions and executor configurations ([bd6f931](https://github.com/averbuh/UniProject/commit/bd6f931a3333bb3b195225d38de64d2359d78b50)), closes [#ISSUE-123](https://github.com/averbuh/UniProject/issues/ISSUE-123)
* test ([a34e0a3](https://github.com/averbuh/UniProject/commit/a34e0a3b21a8704939c1e22af55b841d52f5baa7))
* test ([8836509](https://github.com/averbuh/UniProject/commit/8836509427b4bc3d2936fae5e916ad5d6874b4e2))
* test ([68ba9b4](https://github.com/averbuh/UniProject/commit/68ba9b456b5d62f88638885e8ac8d76c280faab4))
* test ([b796e17](https://github.com/averbuh/UniProject/commit/b796e172febc129af8c3d2fb367495843d0124ec))
* test ([a29c61a](https://github.com/averbuh/UniProject/commit/a29c61a94c0ecf5ce4c6fb0a17e9d3dca4094ad1))
* test ([33d5e00](https://github.com/averbuh/UniProject/commit/33d5e00963ea90014f3067c5692945c6ab1f8842))
* test ([ffa5263](https://github.com/averbuh/UniProject/commit/ffa5263c73ce6012f115365b29205ce64466163e))
* test ([654915b](https://github.com/averbuh/UniProject/commit/654915b0775e6937fb1d6447830d558eea222671))
* test ([6465f29](https://github.com/averbuh/UniProject/commit/6465f2903cdd2dcf38870e626f1e2f6aecbd3d90))
* test ([e9c3965](https://github.com/averbuh/UniProject/commit/e9c3965779f105b29bc7b4bffde18bd8c7a6f2d3))
* Update CircleCI config to use 'run-recipes-go' flag instead of 'run-pipeline' to align with actual workflow. This change removes unused 'run-frontend' parameter. (#issue123) ([089c99e](https://github.com/averbuh/UniProject/commit/089c99ed4f044ef1a3cc49859277e29913bd53ab)), closes [#issue123](https://github.com/averbuh/UniProject/issues/issue123)

# [1.1.0](https://github.com/averbuh/UniProject/compare/v1.0.2...v1.1.0) (2024-08-13)


### Features

* Some changes ([55c78e9](https://github.com/averbuh/UniProject/commit/55c78e98a0cd8307f90d983a91bb294950e4e3eb))

## [1.0.2](https://github.com/averbuh/UniProject/compare/v1.0.1...v1.0.2) (2024-08-13)


### Bug Fixes

* .config ([6abf507](https://github.com/averbuh/UniProject/commit/6abf50704b011bff485847e8680d86992c08d22e))
* config ([497cc10](https://github.com/averbuh/UniProject/commit/497cc10c32df4953b16ff0496549d27e3915f2d3))
* Merge pull request [#40](https://github.com/averbuh/UniProject/issues/40) from averbuh/feature/some-test ([2a31bf5](https://github.com/averbuh/UniProject/commit/2a31bf5efd107616e8a029f8d9a8cd26d2de7ed3))

## [1.0.1](https://github.com/averbuh/UniProject/compare/v1.0.0...v1.0.1) (2024-08-09)


### Bug Fixes

* Refine commit-analyzer plugin preset and release rules, update ApplicationSet ([8f721fe](https://github.com/averbuh/UniProject/commit/8f721fe1458ec9fa181d01390a1da6f3d5e5f578))
* test versioning ([1240ae8](https://github.com/averbuh/UniProject/commit/1240ae84672cbcb3eeb60f25315dd183e11b33f8))

# 1.0.0 (2024-08-09)


### Bug Fixes

* Added semantic versioning ([35f0ff9](https://github.com/averbuh/UniProject/commit/35f0ff98cd6cdaadc62f720251288a9bc10d1b7c))
* Configure mainconfig.yaml ([c08e835](https://github.com/averbuh/UniProject/commit/c08e83559b08717458552fb856f884ee2796a0dd))
* fix semantic release ([ec5b2c8](https://github.com/averbuh/UniProject/commit/ec5b2c87d190050cf6a21f114cd2e7aa455ba3b1))
* Main config.yaml ([bc5d624](https://github.com/averbuh/UniProject/commit/bc5d6244dd1d1954e9c02d65e35025d53d3d7ef2))
* some changes ([fefefea](https://github.com/averbuh/UniProject/commit/fefefea5cec437a934b4c96e81ad7a94bb829eff))
* test ([5d3622c](https://github.com/averbuh/UniProject/commit/5d3622c7b5416ea488de6cc5c78985e04d25445a))
* test ([8eab357](https://github.com/averbuh/UniProject/commit/8eab357b49b4f2b0d0d8aa5c42a252ef227b950a))
* test pipeline ([5659fb3](https://github.com/averbuh/UniProject/commit/5659fb385d4e8378535ef79818968167a70f008b))
