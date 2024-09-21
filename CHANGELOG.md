## [1.2.2](https://github.com/averbuh/UniProject/compare/v1.2.1...v1.2.2) (2024-09-21)


### Bug Fixes

* Fix branch name configuration for alpha and beta releases ([49eb4e7](https://github.com/averbuh/UniProject/commit/49eb4e7fc568dad94ab94313afca2d39e8c8c418))
* Fix missing comma in release configuration file ([b0968a3](https://github.com/averbuh/UniProject/commit/b0968a3243d6fdc9a559ce0ca96b8bc2be457fd2))
* Fix setting of prerelease values to boolean type in .releaserc.json configuration file. This change corrects the format to ensure proper prerelease handling. ([e338ece](https://github.com/averbuh/UniProject/commit/e338ece7950cdde8260e32f64eff623abd87139b))
* Fix unnecessary commented-out code related to branch filtering in CircleCI configuration. ([a4ff8fd](https://github.com/averbuh/UniProject/commit/a4ff8fd1db2b744f0ce6f22994f2dba1bffe2b8b))
* issue related to unnecessary commented code and add clarifying comment in CircleCI configuration file. ([7813d8b](https://github.com/averbuh/UniProject/commit/7813d8b234bb171e48a4469bf7da50a8db19e784))
* Refine branch naming conventions for alpha and beta releases ([ebe39d6](https://github.com/averbuh/UniProject/commit/ebe39d6656e3cbd9ba2f5c433b77d4fc383a9884))
* some fix ([20ae929](https://github.com/averbuh/UniProject/commit/20ae92985cd24f543fc10399dc45c5a91dbd9a17))
* some fix ([3495246](https://github.com/averbuh/UniProject/commit/34952462cdb82446967dbeec2f633ab8d5e39722))

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
