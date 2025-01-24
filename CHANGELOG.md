# Changelog

## 0.1.0-rc1 (2025-01-24)

Full Changelog: [v0.1.0-alpha.9...v0.1.0-rc1](https://github.com/dackerman/terraform-provider-demostore/compare/v0.1.0-alpha.9...v0.1.0-rc1)

### Features

* **api:** update via SDK Studio ([#67](https://github.com/dackerman/terraform-provider-demostore/issues/67)) ([f9fd240](https://github.com/dackerman/terraform-provider-demostore/commit/f9fd240014f0b07987d95ce49517a39cb5e9ec17))


### Bug Fixes

* **release:** don't try to announce to discord ([54f2313](https://github.com/dackerman/terraform-provider-demostore/commit/54f23139070abe7a074788cea0696a62d03d1f47))

## 0.1.0-alpha.9 (2025-01-24)

Full Changelog: [v0.1.0-alpha.8...v0.1.0-alpha.9](https://github.com/dackerman/terraform-provider-demostore/compare/v0.1.0-alpha.8...v0.1.0-alpha.9)

### Features

* fix(build): add perms and concurrency to release action ([cf46265](https://github.com/dackerman/terraform-provider-demostore/commit/cf462650e2c88a2d3ae8c06a72628cde7d58bbda))

## 0.1.0-alpha.8 (2025-01-23)

Full Changelog: [v0.1.0-alpha.7...v0.1.0-alpha.8](https://github.com/dackerman/terraform-provider-demostore/compare/v0.1.0-alpha.7...v0.1.0-alpha.8)

### Features

* fix(build): run go mod tidy ([cd6c426](https://github.com/dackerman/terraform-provider-demostore/commit/cd6c42659d401094f77fbadf84942e9dd75955fb))

## 0.1.0-alpha.7 (2025-01-23)

Full Changelog: [v0.1.0-alpha.6...v0.1.0-alpha.7](https://github.com/dackerman/terraform-provider-demostore/compare/v0.1.0-alpha.6...v0.1.0-alpha.7)

### Features

* chore(release): use correct gpg secret ([7e6dee3](https://github.com/dackerman/terraform-provider-demostore/commit/7e6dee334248117e688da4bc717bed039f8bb8d9))

## 0.1.0-alpha.6 (2025-01-23)

Full Changelog: [v0.1.0-alpha.5...v0.1.0-alpha.6](https://github.com/dackerman/terraform-provider-demostore/compare/v0.1.0-alpha.5...v0.1.0-alpha.6)

### Features

* **api:** update via SDK Studio ([#59](https://github.com/dackerman/terraform-provider-demostore/issues/59)) ([55295fe](https://github.com/dackerman/terraform-provider-demostore/commit/55295fefd76e3ca79449675de6eeff373e187c1f))
* remove resource state upon 404 during refresh operation ([#24](https://github.com/dackerman/terraform-provider-demostore/issues/24)) ([9761ade](https://github.com/dackerman/terraform-provider-demostore/commit/9761ade6cf00c4d2814ab36b35f2b3b7e49f22f7))
* **terraform:** improved warning when encountering 404 on resource get ([#58](https://github.com/dackerman/terraform-provider-demostore/issues/58)) ([8269f07](https://github.com/dackerman/terraform-provider-demostore/commit/8269f079e4f161470a6f4ffabf7eff02bda39145))
* tweak release workflow ([8a7d211](https://github.com/dackerman/terraform-provider-demostore/commit/8a7d21122a1c312adde546d0501f184969b23ed0))


### Bug Fixes

* **api:** encode objects as `application/json` in multipart encoding ([#45](https://github.com/dackerman/terraform-provider-demostore/issues/45)) ([2a2b4e2](https://github.com/dackerman/terraform-provider-demostore/commit/2a2b4e2ba022496a50799a11737fd75b934c271d))
* **docs:** reference provider type name in example code block ([#26](https://github.com/dackerman/terraform-provider-demostore/issues/26)) ([f603598](https://github.com/dackerman/terraform-provider-demostore/commit/f603598cceb06aaf6f1baf5f923ff7be80adaf62))


### Chores

* **internal:** bumps internal dependencies ([#27](https://github.com/dackerman/terraform-provider-demostore/issues/27)) ([074ed48](https://github.com/dackerman/terraform-provider-demostore/commit/074ed489baf13b6acf35055d19d51b9eef5870ae))
* **internal:** codegen related update ([#28](https://github.com/dackerman/terraform-provider-demostore/issues/28)) ([ce658fe](https://github.com/dackerman/terraform-provider-demostore/commit/ce658fe0114553360050c8ff91aa4da3d830b91f))
* **internal:** codegen related update ([#29](https://github.com/dackerman/terraform-provider-demostore/issues/29)) ([c872f26](https://github.com/dackerman/terraform-provider-demostore/commit/c872f260f4eac9776b23052f224c0d3feb4d4a0e))
* **internal:** codegen related update ([#30](https://github.com/dackerman/terraform-provider-demostore/issues/30)) ([b9ef715](https://github.com/dackerman/terraform-provider-demostore/commit/b9ef715b52189da06462690da63b8cd81978eb61))
* **internal:** codegen related update ([#31](https://github.com/dackerman/terraform-provider-demostore/issues/31)) ([6adb379](https://github.com/dackerman/terraform-provider-demostore/commit/6adb379a9e839754ebba0724e500fd3a82c3988b))
* **internal:** codegen related update ([#32](https://github.com/dackerman/terraform-provider-demostore/issues/32)) ([b7a0723](https://github.com/dackerman/terraform-provider-demostore/commit/b7a07237b0f6563a648f8f0975869d9c37196aa4))
* **internal:** codegen related update ([#33](https://github.com/dackerman/terraform-provider-demostore/issues/33)) ([ee91670](https://github.com/dackerman/terraform-provider-demostore/commit/ee916702e7efa6b333d24d2d82b01a610c6250f1))
* **internal:** codegen related update ([#34](https://github.com/dackerman/terraform-provider-demostore/issues/34)) ([3be1939](https://github.com/dackerman/terraform-provider-demostore/commit/3be19399255470eebb4cb0a13c15c49713f4fc13))
* **internal:** codegen related update ([#35](https://github.com/dackerman/terraform-provider-demostore/issues/35)) ([2d7c0d0](https://github.com/dackerman/terraform-provider-demostore/commit/2d7c0d0a8eb4995fd44565a2233ce043066f87c2))
* **internal:** codegen related update ([#36](https://github.com/dackerman/terraform-provider-demostore/issues/36)) ([a62b492](https://github.com/dackerman/terraform-provider-demostore/commit/a62b492739bd2c2081606238876b13697bcbb9ff))
* **internal:** codegen related update ([#37](https://github.com/dackerman/terraform-provider-demostore/issues/37)) ([fe9c95c](https://github.com/dackerman/terraform-provider-demostore/commit/fe9c95c98dabd354c74d00ca483a86d7e8efa7df))
* **internal:** codegen related update ([#38](https://github.com/dackerman/terraform-provider-demostore/issues/38)) ([9f26cb3](https://github.com/dackerman/terraform-provider-demostore/commit/9f26cb358832ff9e085be4148931072ae6bd576d))
* **internal:** codegen related update ([#39](https://github.com/dackerman/terraform-provider-demostore/issues/39)) ([8ff96db](https://github.com/dackerman/terraform-provider-demostore/commit/8ff96db560777240be8957ac56ed472412731483))
* **internal:** codegen related update ([#40](https://github.com/dackerman/terraform-provider-demostore/issues/40)) ([eaff974](https://github.com/dackerman/terraform-provider-demostore/commit/eaff9741dec4bef2d9017b71fd664444c631eee5))
* **internal:** codegen related update ([#41](https://github.com/dackerman/terraform-provider-demostore/issues/41)) ([4bb7d5e](https://github.com/dackerman/terraform-provider-demostore/commit/4bb7d5e1f0760cf9f8f74a2818ee0de508f39e60))
* **internal:** codegen related update ([#42](https://github.com/dackerman/terraform-provider-demostore/issues/42)) ([c86b138](https://github.com/dackerman/terraform-provider-demostore/commit/c86b13854e8e39fc568b9493c41322da9038d9fa))
* **internal:** codegen related update ([#43](https://github.com/dackerman/terraform-provider-demostore/issues/43)) ([9745a21](https://github.com/dackerman/terraform-provider-demostore/commit/9745a2140065fe02ca0649f2ddd52dc8e58efa1c))
* **internal:** codegen related update ([#44](https://github.com/dackerman/terraform-provider-demostore/issues/44)) ([939846f](https://github.com/dackerman/terraform-provider-demostore/commit/939846f0164a1051817925b929a50d3d9aacc195))
* **internal:** codegen related update ([#46](https://github.com/dackerman/terraform-provider-demostore/issues/46)) ([3ba0268](https://github.com/dackerman/terraform-provider-demostore/commit/3ba0268456c3d78bfc63b30b28fc8c53c5b089f9))
* **internal:** codegen related update ([#47](https://github.com/dackerman/terraform-provider-demostore/issues/47)) ([bd01085](https://github.com/dackerman/terraform-provider-demostore/commit/bd010859170dfa6e2d5a6bf86095070f618ef4cb))
* **internal:** codegen related update ([#48](https://github.com/dackerman/terraform-provider-demostore/issues/48)) ([85ed9db](https://github.com/dackerman/terraform-provider-demostore/commit/85ed9dbc6c6df4ed002edffd9de7c9eb2cf11f03))
* **internal:** codegen related update ([#49](https://github.com/dackerman/terraform-provider-demostore/issues/49)) ([3b603da](https://github.com/dackerman/terraform-provider-demostore/commit/3b603da9c122dad4425d4e7a7f1c360543129e94))
* **internal:** codegen related update ([#50](https://github.com/dackerman/terraform-provider-demostore/issues/50)) ([4da9aa7](https://github.com/dackerman/terraform-provider-demostore/commit/4da9aa72d713a84d7207aaafb83389355dbbcaa5))
* **internal:** codegen related update ([#51](https://github.com/dackerman/terraform-provider-demostore/issues/51)) ([4d35144](https://github.com/dackerman/terraform-provider-demostore/commit/4d3514405eb3dba3eb6efef777c3e659632757b1))
* **internal:** codegen related update ([#52](https://github.com/dackerman/terraform-provider-demostore/issues/52)) ([1ecd0d1](https://github.com/dackerman/terraform-provider-demostore/commit/1ecd0d1498c9e9c30f0f1a54e57a09997733b2ad))
* **internal:** codegen related update ([#53](https://github.com/dackerman/terraform-provider-demostore/issues/53)) ([9c65603](https://github.com/dackerman/terraform-provider-demostore/commit/9c65603e3068411e9bdbf67e8180300d57ed3fc8))
* **internal:** codegen related update ([#54](https://github.com/dackerman/terraform-provider-demostore/issues/54)) ([cba8f7f](https://github.com/dackerman/terraform-provider-demostore/commit/cba8f7f8f566d4fa8e9d535ab1e254f3c73c9a62))
* **internal:** codegen related update ([#55](https://github.com/dackerman/terraform-provider-demostore/issues/55)) ([e9eaa73](https://github.com/dackerman/terraform-provider-demostore/commit/e9eaa73a2e46be79eff36efc54403f79eaaae4e9))
* **internal:** codegen related update ([#56](https://github.com/dackerman/terraform-provider-demostore/issues/56)) ([265ebaa](https://github.com/dackerman/terraform-provider-demostore/commit/265ebaaf6446ba98c34e8c507a00c5fcc5570b4d))
* **internal:** codegen related update ([#57](https://github.com/dackerman/terraform-provider-demostore/issues/57)) ([d99d524](https://github.com/dackerman/terraform-provider-demostore/commit/d99d524d73bd8b54b50d760c3f26f5f5c77e469f))
* **internal:** codegen related update ([#60](https://github.com/dackerman/terraform-provider-demostore/issues/60)) ([8757cc8](https://github.com/dackerman/terraform-provider-demostore/commit/8757cc893f414808c0671bed2d4487d0febfb124))
* **internal:** codegen related update ([#61](https://github.com/dackerman/terraform-provider-demostore/issues/61)) ([9900c41](https://github.com/dackerman/terraform-provider-demostore/commit/9900c410b34413666e6052d0ed59e562323b1571))

## 0.1.0-alpha.5 (2024-12-06)

Full Changelog: [v0.1.0-alpha.4...v0.1.0-alpha.5](https://github.com/dackerman/terraform-provider-demostore/compare/v0.1.0-alpha.4...v0.1.0-alpha.5)

### Features

* add workflow ([24797d7](https://github.com/dackerman/terraform-provider-demostore/commit/24797d761b202cbc673158543eb2dffb126530b7))
* **api:** update via SDK Studio ([#18](https://github.com/dackerman/terraform-provider-demostore/issues/18)) ([e0b5c07](https://github.com/dackerman/terraform-provider-demostore/commit/e0b5c078a51620cebc34abb8376e304a79dd02be))
* better output ([ee54dd8](https://github.com/dackerman/terraform-provider-demostore/commit/ee54dd8f9644a8cebbbfebe8432cd27618bcc72e))
* ci ([4694abe](https://github.com/dackerman/terraform-provider-demostore/commit/4694abecd6148e0f26ed9ff77c00782ba2a3b06b))
* debug ([5c6f01d](https://github.com/dackerman/terraform-provider-demostore/commit/5c6f01d5be5f58fa32311630570f1d89b667d4c1))
* debug ([7b2b38c](https://github.com/dackerman/terraform-provider-demostore/commit/7b2b38c06254cdb139999752b41ceb058bf1b2b8))
* debug ([07c2a17](https://github.com/dackerman/terraform-provider-demostore/commit/07c2a17182e0dc656e8d0a1dc87db110796e25b6))
* i give up ([8f8c5df](https://github.com/dackerman/terraform-provider-demostore/commit/8f8c5df10f9d13c27dfe7d29d25ea23bedb48807))
* jq ([29a2fe8](https://github.com/dackerman/terraform-provider-demostore/commit/29a2fe8b4d3a0ac8961f88ab89a1d4e4dc1f126e))
* script ([bf7efa8](https://github.com/dackerman/terraform-provider-demostore/commit/bf7efa88cfe691eb52aa58a7b924e2039dcf098d))
* something else; ([efdba94](https://github.com/dackerman/terraform-provider-demostore/commit/efdba94ad33b353457795a0ca94cbf534f5d8f37))
* v1 script ([50fc627](https://github.com/dackerman/terraform-provider-demostore/commit/50fc62740e04043003d6d46c048bbdc3450c3cc7))
* wip ([1af4afb](https://github.com/dackerman/terraform-provider-demostore/commit/1af4afbee0ecfc6a263e23c6ab29c45c1167aa29))


### Chores

* **internal:** version bump ([#16](https://github.com/dackerman/terraform-provider-demostore/issues/16)) ([9048179](https://github.com/dackerman/terraform-provider-demostore/commit/9048179bcf8b3320a5cca0251e6abebe4f58bd97))

## 0.1.0-alpha.4 (2024-12-06)

Full Changelog: [v0.1.0-alpha.3...v0.1.0-alpha.4](https://github.com/dackerman/terraform-provider-demostore/compare/v0.1.0-alpha.3...v0.1.0-alpha.4)

### Features

* **api:** update via SDK Studio ([#13](https://github.com/dackerman/terraform-provider-demostore/issues/13)) ([aec80a0](https://github.com/dackerman/terraform-provider-demostore/commit/aec80a0cabff6494f2fb7471c135b2bb46fc46c2))
* **api:** update via SDK Studio ([#14](https://github.com/dackerman/terraform-provider-demostore/issues/14)) ([8b189a7](https://github.com/dackerman/terraform-provider-demostore/commit/8b189a7f405a601ea8c23b71465606d00c0a1e2d))
* **api:** update via SDK Studio ([#15](https://github.com/dackerman/terraform-provider-demostore/issues/15)) ([7ade39c](https://github.com/dackerman/terraform-provider-demostore/commit/7ade39ccc610bf90293c1845e8c465ec6484fe0c))


### Chores

* **internal:** version bump ([#11](https://github.com/dackerman/terraform-provider-demostore/issues/11)) ([23ec14e](https://github.com/dackerman/terraform-provider-demostore/commit/23ec14e48580e05c263c1a19d17601bf1dc16267))

## 0.1.0-alpha.3 (2024-12-05)

Full Changelog: [v0.1.0-alpha.2...v0.1.0-alpha.3](https://github.com/dackerman/terraform-provider-demostore/compare/v0.1.0-alpha.2...v0.1.0-alpha.3)

### Features

* **api:** update via SDK Studio ([#10](https://github.com/dackerman/terraform-provider-demostore/issues/10)) ([922badb](https://github.com/dackerman/terraform-provider-demostore/commit/922badb4294ec008a8108384cb83143f38503643))


### Chores

* **internal:** version bump ([#8](https://github.com/dackerman/terraform-provider-demostore/issues/8)) ([f16e2ca](https://github.com/dackerman/terraform-provider-demostore/commit/f16e2ca21a192d949e4f72bda2502b8657be1db6))

## 0.1.0-alpha.2 (2024-12-05)

Full Changelog: [v0.1.0-alpha.1...v0.1.0-alpha.2](https://github.com/dackerman/terraform-provider-demostore/compare/v0.1.0-alpha.1...v0.1.0-alpha.2)

### Chores

* **internal:** version bump ([#6](https://github.com/dackerman/terraform-provider-demostore/issues/6)) ([0bc85fb](https://github.com/dackerman/terraform-provider-demostore/commit/0bc85fb3be6d03a1b6cb7e0622f62bca6f0dcecd))

## 0.1.0-alpha.1 (2024-12-05)

Full Changelog: [v0.0.1-alpha.0...v0.1.0-alpha.1](https://github.com/dackerman/terraform-provider-demostore/compare/v0.0.1-alpha.0...v0.1.0-alpha.1)

### Features

* **api:** update via SDK Studio ([#3](https://github.com/dackerman/terraform-provider-demostore/issues/3)) ([49eefc3](https://github.com/dackerman/terraform-provider-demostore/commit/49eefc3861f3b1fe5ef835b11c4c3fef158dd102))
* **api:** update via SDK Studio ([#4](https://github.com/dackerman/terraform-provider-demostore/issues/4)) ([ba9b5a9](https://github.com/dackerman/terraform-provider-demostore/commit/ba9b5a9bebbc525beabd145284558ee5a3d9aea0))


### Chores

* go live ([61407b1](https://github.com/dackerman/terraform-provider-demostore/commit/61407b15e9169a58ed8d14e28d36afb753e2a55a))
* update SDK settings ([#1](https://github.com/dackerman/terraform-provider-demostore/issues/1)) ([4e5cfc6](https://github.com/dackerman/terraform-provider-demostore/commit/4e5cfc650e59a2679c10c2c5fdeff86beb749918))
* update SDK settings ([#2](https://github.com/dackerman/terraform-provider-demostore/issues/2)) ([32c83bc](https://github.com/dackerman/terraform-provider-demostore/commit/32c83bcd90fd265b1cb1a2c64a3595ba2ef3e1e8))
