# Changelog

## [0.2.11](https://github.com/tuihub/librarian/compare/v0.2.10...v0.2.11) (2024-09-25)


### Features

* impl porter context control ([400890b](https://github.com/tuihub/librarian/commit/400890bf060e2c2993201aa5bafcca1579d6ab9d))
* refactor libcache ([b186f43](https://github.com/tuihub/librarian/commit/b186f435386ccfcc85a52f450884e6c7a7119dd1))
* refactor supervisor ([afbc8c3](https://github.com/tuihub/librarian/commit/afbc8c336befa55d71572905ba6e6f0d029e7064))
* support use porter without consul ([30efdaf](https://github.com/tuihub/librarian/commit/30efdaf1de09b8333dcdb77b36a26f10b1a10902))
* update kratos to v2.8.0 ([8d8d22b](https://github.com/tuihub/librarian/commit/8d8d22bcec3abb973e3489e3b5f9c7c958b698ab))
* update proto to v0.4.22 ([93eafb0](https://github.com/tuihub/librarian/commit/93eafb0e57dedf75ada05a449cc2425947f00b34))
* update proto to v0.4.23 ([52a939b](https://github.com/tuihub/librarian/commit/52a939b3a1404436d8102a33639f86a9d64d7da8))
* upgrade golangci version ([314720d](https://github.com/tuihub/librarian/commit/314720d63d5fd502b0af15f7c60faff12b4c798b))
* upgrade golangci-lint to v1.60.3 ([b74baf7](https://github.com/tuihub/librarian/commit/b74baf753ce89c88abc3559e9b5c1b61bb6f3e77))
* upgrade goverter to v1.5.0 ([9d3835b](https://github.com/tuihub/librarian/commit/9d3835b38988de654e4d51f1851d26e926c7dd4c))
* upgrade proto to v0.4.16 ([670e50f](https://github.com/tuihub/librarian/commit/670e50f10d34a57a1238ef7012a4576731664535))
* upgrade proto to v0.4.17 ([180cd7e](https://github.com/tuihub/librarian/commit/180cd7e8a131f7a9524ae03cb0b129baa44e1ead))
* upgrade proto to v0.4.18 ([326f56d](https://github.com/tuihub/librarian/commit/326f56dd6809f8323a80b54ee4c5ca69a5a049de))
* upgrade proto to v0.4.20 ([7f5ba6d](https://github.com/tuihub/librarian/commit/7f5ba6df657a5ad69cc21ba6e30b483aa569a8a1))
* upgrade proto to v0.4.21 ([fb1e216](https://github.com/tuihub/librarian/commit/fb1e2164774dccd22e5d55133c28a80c67fba3ce))


### Bug Fixes

* lint error ([7f01d1d](https://github.com/tuihub/librarian/commit/7f01d1d9359d67080bf6b45573ec0c6bec0b93e6))
* test new logic ([d047ccd](https://github.com/tuihub/librarian/commit/d047ccdc5c8a465ec9b29ef130fc25c07e807975))
* test notify ([52aa741](https://github.com/tuihub/librarian/commit/52aa7411e7b379dd399a4c7d009833248e2ef373))
* test porter context control ([8b3f8ca](https://github.com/tuihub/librarian/commit/8b3f8ca4841b219830acb2a852de67bc1bde48fa))
* update sync_map.go ([2ee9464](https://github.com/tuihub/librarian/commit/2ee946480b86cc404013ab2a9cdf14ff5d60842d))
* **yesod:** fix CreateFeedConfig ([08d22e6](https://github.com/tuihub/librarian/commit/08d22e6f8a3c4cf64a5ce9b55c3d5534c5cd0d38))

## [0.2.10](https://github.com/tuihub/librarian/compare/v0.2.9...v0.2.10) (2024-07-10)


### Features

* add built-in metrics ([291db25](https://github.com/tuihub/librarian/commit/291db255a35fcbff1213e4a267ea1b8addc33104))
* add independent otel support ([b5850e5](https://github.com/tuihub/librarian/commit/b5850e5beb066b2391926d59b5ad35a9fd1b30ed))
* add stdout logger ([b0427ef](https://github.com/tuihub/librarian/commit/b0427ef41b8199b78005dc7936cfe7b2641df3b6))
* impl `CreateFeedActionSet` `UpdateFeedActionSet` `ListFeedActionSets` ([f6f0de8](https://github.com/tuihub/librarian/commit/f6f0de8f1b690a87dd71db5b8affe6ad9703d416))
* impl builtin actions ([d2c8949](https://github.com/tuihub/librarian/commit/d2c8949b5e9277a0c6bc26bde27b56b5f621c9b5))
* impl feed pull state ([8592902](https://github.com/tuihub/librarian/commit/8592902d73434983b8150c795f0d01309a19f763))
* impl porter states notification ([c66f20c](https://github.com/tuihub/librarian/commit/c66f20c1c101b25da4d3d19aca6672e7fac56f5d))
* impl system notification ([d527588](https://github.com/tuihub/librarian/commit/d527588eca7df66ac103e2115e3ae29bb4a178b7))
* upgrade goverter to 1.4.0 ([0ad9821](https://github.com/tuihub/librarian/commit/0ad982116ec4e5d893ddd9671aecfedbb8e84319))
* upgrade proto to v0.4.12 ([6945d99](https://github.com/tuihub/librarian/commit/6945d997454b064f919f040b52921f124f14b6ba))
* upgrade proto to v0.4.13 ([1c2f081](https://github.com/tuihub/librarian/commit/1c2f0815aa3ceccad505417ecf32e66ae4125320))
* upgrade proto to v0.4.14 ([3da81e0](https://github.com/tuihub/librarian/commit/3da81e0607c421609f82370c44b76623072b4bbf))
* upgrade proto to v0.4.8 ([c3cefe9](https://github.com/tuihub/librarian/commit/c3cefe99a860e7430eb427acfa3b16228a1d1700))
* **yesod:** impl feed item collection api ([6c8e4eb](https://github.com/tuihub/librarian/commit/6c8e4ebdb0ffda0c8a5ade4bfd4a9676acacfc5e))


### Bug Fixes

* adjust magic numbers ([63c1689](https://github.com/tuihub/librarian/commit/63c1689b57089b04c46d609db16ff2a5f37ca3fe))
* improve feed states notification ([e9e2d34](https://github.com/tuihub/librarian/commit/e9e2d34b2540bf63f3f50afa120698df0dce605a))
* system notification tested ([f90ca07](https://github.com/tuihub/librarian/commit/f90ca07d32ee2b8bedac68ac552d5a8d1f3f8501))
* upgrade proto ([ac502f4](https://github.com/tuihub/librarian/commit/ac502f4afb914be64730a7816819dac461d07e46))
* **yesod:** add simple keyword filter ([0576ac4](https://github.com/tuihub/librarian/commit/0576ac4859892b2b9159122f55d29a2c290ae6f4))
* **yesod:** tested pull feed ([e3cb3c2](https://github.com/tuihub/librarian/commit/e3cb3c20bc97b80d6cdf05ea2030a4b7506b7037))

## [0.2.9](https://github.com/tuihub/librarian/compare/v0.2.8...v0.2.9) (2024-03-15)


### Bug Fixes

* **tiphereth:** fix user capacity ([4af5634](https://github.com/tuihub/librarian/commit/4af56348536dfaa6afe06b64c4c98bdc7332de40))

## [0.2.8](https://github.com/tuihub/librarian/compare/v0.2.7...v0.2.8) (2024-03-14)


### Features

* **tiphereth:** impl RegisterUser ([fb70689](https://github.com/tuihub/librarian/commit/fb7068900a0433f031c7004c200626cefc5904c7))


### Bug Fixes

* fix captcha delete ([ca5b399](https://github.com/tuihub/librarian/commit/ca5b399975a1011278ffd15a9be30001cc20e588))
* **gebura:** tested app inst run time ([fed04fa](https://github.com/tuihub/librarian/commit/fed04faa012db77ac6936990b8acec327aaa20e7))
* **gebura:** update SyncAppInfo ([4f87179](https://github.com/tuihub/librarian/commit/4f87179ba577d9415d4809928f1d84f492015ba9))
* **tiphereth:** sync pull account on linking ([699263b](https://github.com/tuihub/librarian/commit/699263bd1d5c1a21586f705fbd1bded6e25260a7))

## [0.2.7](https://github.com/tuihub/librarian/compare/v0.2.6...v0.2.7) (2024-03-11)


### Features

* add redis driver for mq ([02dc32b](https://github.com/tuihub/librarian/commit/02dc32b14e9d6e21fae7f47eda8fbac0274828cc))
* impl SearchNewAppInfos ([7c28bbe](https://github.com/tuihub/librarian/commit/7c28bbe898be8ed01db1fbd2e3b952d3eacd7973))
* upgrade proto to v0.4.4 ([92b5aef](https://github.com/tuihub/librarian/commit/92b5aefd445f1e5553f12108b426aab41286ce6f))


### Bug Fixes

* add mutex ([bb884c9](https://github.com/tuihub/librarian/commit/bb884c955c61dbe333858300eb972f265d5536e7))
* ignore rate limit in syncAppInfo ([de889c0](https://github.com/tuihub/librarian/commit/de889c07e50d6cd9335ca6183b4ba5567cef37de))

## [0.2.6](https://github.com/tuihub/librarian/compare/v0.2.5...v0.2.6) (2024-03-01)


### Bug Fixes

* **tiphereth:** error on creating new user session ([3151404](https://github.com/tuihub/librarian/commit/3151404ae32b3680542f3067cd97256bb23a36b7))

## [0.2.5](https://github.com/tuihub/librarian/compare/v0.2.4...v0.2.5) (2024-02-23)


### Bug Fixes

* check pointer before use ([a0b6ecb](https://github.com/tuihub/librarian/commit/a0b6ecb13192c2e111b3c90d5c4a32923941b5e8))

## [0.2.4](https://github.com/tuihub/librarian/compare/v0.2.3...v0.2.4) (2024-02-23)


### Features

* add sentry support ([23ea8ab](https://github.com/tuihub/librarian/commit/23ea8abb8d9cb8ee3d7e6aa7b549b48402a30c8f))
* upgrade gocron to v2 & support sentry cron ([019993b](https://github.com/tuihub/librarian/commit/019993badfd9205e54cdefcc7e5ffc4cff7e4075))


### Bug Fixes

* supervisor save instance on activation failed ([9e88651](https://github.com/tuihub/librarian/commit/9e8865176e32cf866118525c8b74d089b14b6190))

## [0.2.3](https://github.com/tuihub/librarian/compare/v0.2.2...v0.2.3) (2024-02-20)


### Features

* add toml support ([54e8daa](https://github.com/tuihub/librarian/commit/54e8daa731700dc007e04d674355fc62ccc2bbf2))
* remove mapper ([4cd88ab](https://github.com/tuihub/librarian/commit/4cd88ab4f33df74fdc622600b5452974a66ed49a))
* support zero config run ([34c2bc7](https://github.com/tuihub/librarian/commit/34c2bc76110e0aa49bcf0356a41c34cf664d513f))
* upgrade go version to 1.21 ([0a1b1cb](https://github.com/tuihub/librarian/commit/0a1b1cb140ac679e1ad80a59385ffbacab558271))
* upgrade goverter to v1.3.2 ([062d24a](https://github.com/tuihub/librarian/commit/062d24afb150dd10357bc4611d0f16d288554843))
* upgrade proto to v0.4.0 ([03cbc88](https://github.com/tuihub/librarian/commit/03cbc8835796522054b7f5bce5c9faf6877ee14e))
* upgrade proto to v0.4.1 ([664b7d0](https://github.com/tuihub/librarian/commit/664b7d0184f68ba34a7119539a369aa1cc3484c8))
* upgrade proto to v0.4.2 ([87228dd](https://github.com/tuihub/librarian/commit/87228dd9dfdd474979d38c924ced40fdf4195aad))
* upgrade protos to v0.3.9 ([f745b82](https://github.com/tuihub/librarian/commit/f745b82c0d60e821e50889d86315b607b9dbfa99))


### Bug Fixes

* allow run without consul ([caf6f3a](https://github.com/tuihub/librarian/commit/caf6f3adb1a40f9d853452972a0aa6f219e54d6b))
* impl porter connection status & instance info ([833921e](https://github.com/tuihub/librarian/commit/833921ea7bac778efb5bd24fbf32756a7f0ce2ec))
* several logic fixes ([1f9259f](https://github.com/tuihub/librarian/commit/1f9259f22263de98f2b5f0766795bdb2bb5e67e4))
* upgrade proto to v0.3.10 ([c05d45c](https://github.com/tuihub/librarian/commit/c05d45cac5057156576f944336206c58d7fcbf8d))

## [0.2.2](https://github.com/tuihub/librarian/compare/v0.2.1...v0.2.2) (2024-01-23)


### Features

* add consul config ([4570a11](https://github.com/tuihub/librarian/commit/4570a1139194628bb7af9dca7206094e923edb53))
* add demo mode ([6860c50](https://github.com/tuihub/librarian/commit/6860c507aa59d9ff5068cca9630e8a06f1dcaf37))

## [0.2.1](https://github.com/tuihub/librarian/compare/v0.2.0...v0.2.1) (2024-01-18)


### Features

* add stricter limit for anonymous requests ([b51b660](https://github.com/tuihub/librarian/commit/b51b66089f15ed7d6cfc4cfe56f998ea2231a8d6))
* **gebura:** bump proto version to v0.3.8 & fixes ([56913ac](https://github.com/tuihub/librarian/commit/56913acc33d6df706c6f5da6a83eeb7ccf8d9ad0))
* **gebura:** impl SyncApps ([8e6d184](https://github.com/tuihub/librarian/commit/8e6d184383bce7c642d918cf0b71a70e3a1c4659))
* **gebura:** upgrade proto to v0.3.6 ([00c3360](https://github.com/tuihub/librarian/commit/00c3360287ac4c78e6bc7e38f7af70badd3a1b09))
* **tiphereth:** impl session manage ([2555d7f](https://github.com/tuihub/librarian/commit/2555d7ff80467bfaa6bbf4d91e4e5d5cf95da76d))
* **tiphereth:** upgrade proto to v0.3.3 ([861acc6](https://github.com/tuihub/librarian/commit/861acc69433a126c60043bffbb595d4c3903045f))


### Bug Fixes

* move porters to individual repos ([1335c91](https://github.com/tuihub/librarian/commit/1335c91608810fc501e4ba98e61c96073bdd1289))
* test v0.2.0 features ([0170e80](https://github.com/tuihub/librarian/commit/0170e808d7be6d8cd2391a321b908539ffb94f3c))

## [0.2.0](https://github.com/tuihub/librarian/compare/v0.1.16...v0.2.0) (2024-01-10)


### ⚠ BREAKING CHANGES

* porter has compeletly redesigned, see documents for details

### Features

* add porter-sdk ([a157826](https://github.com/tuihub/librarian/commit/a157826488f3c2f5731437874e79fc57bbcbeefa))
* impl Porter manage and activate ([8eeca27](https://github.com/tuihub/librarian/commit/8eeca274d726a5c8ac0a1dd617e4b57acabaaf8e))
* move rss logic to porter-rss ([1913a82](https://github.com/tuihub/librarian/commit/1913a825a1a6af395dce80c4e1841da3720dfe7a))
* move s3 logic to sephirah ([5998be1](https://github.com/tuihub/librarian/commit/5998be15062cf8d99c8988178ca31acb30364ff8))
* move steam logic to porter-steam ([e5803ab](https://github.com/tuihub/librarian/commit/e5803abf70e02ab492d9f884c790dfe620756430))
* move telegram logic to porter-telegram ([442b108](https://github.com/tuihub/librarian/commit/442b1084d54d2bde816a6a49702f49c29128df40))
* porter has compeletly redesigned, see documents for details ([db830e6](https://github.com/tuihub/librarian/commit/db830e61b6f70729cb25e7509e2de79ff4ff658c))
* WIP ([6a4f4b0](https://github.com/tuihub/librarian/commit/6a4f4b0b068eff3c3c9dff0b6b5e168b1203086a))
* WIP ([9d3b9e0](https://github.com/tuihub/librarian/commit/9d3b9e0a5e661d30f60d9fad0711ce6991c8e25b))
* WIP ([b63eac4](https://github.com/tuihub/librarian/commit/b63eac46467c13491aace4887a1136527316ecbd))


### Bug Fixes

* test porter connection with consul ([a4e62f7](https://github.com/tuihub/librarian/commit/a4e62f702bce2a97d52c76f780d7c5e34a27621a))
* tested porter-rss ([9b718ad](https://github.com/tuihub/librarian/commit/9b718ad37f42bc35fd20f1526ba77e419596c2b1))
* tested with waiter ([9116b6b](https://github.com/tuihub/librarian/commit/9116b6b5560540331c0b2504a81a5bc7d7d289b6))

## [0.1.16](https://github.com/tuihub/librarian/compare/v0.1.15...v0.1.16) (2024-01-03)


### Features

* remove unauthorised error from log ([ddfcd4a](https://github.com/tuihub/librarian/commit/ddfcd4a5530024b836be90fdf9c9d89973744abc))
* upgrade proto version ([ec8f436](https://github.com/tuihub/librarian/commit/ec8f43644485dae408a09a21a4dfabec8a4b5688))

## [0.1.15](https://github.com/tuihub/librarian/compare/v0.1.14...v0.1.15) (2023-12-21)


### Features

* bump protos version to v0.2.46 ([fce6e58](https://github.com/tuihub/librarian/commit/fce6e58364b874f669527e1fa319a306058d78a5))
* **yesod:** impl ListFeedPlatforms & ReadFeedItem ([508dd87](https://github.com/tuihub/librarian/commit/508dd87b50157ccc480ccf97afebb15f792f4cf3))


### Bug Fixes

* feed icon & SumAppPackageRunTime ([91ac63d](https://github.com/tuihub/librarian/commit/91ac63d558c143c7dda022f5b9ad9d6c897f66c4))
* **netzach:** tested with client ([0303138](https://github.com/tuihub/librarian/commit/0303138358201337c5d78bdfb7d55419e85cd417))
* **netzach:** update sql structure ([edb2fa6](https://github.com/tuihub/librarian/commit/edb2fa6c0fc9a39107168cbd5e5a7a8ef7a6f995))

## [0.1.14](https://github.com/tuihub/librarian/compare/v0.1.13...v0.1.14) (2023-11-30)


### Bug Fixes

* **gebura:** re-bind account error & only search app name ([bb53b96](https://github.com/tuihub/librarian/commit/bb53b965db9e1600f39e81d742cc63d8ed82cb44))

## [0.1.13](https://github.com/tuihub/librarian/compare/v0.1.12...v0.1.13) (2023-10-22)


### Bug Fixes

* revert "fix: reduce disk writes" ([fb2084f](https://github.com/tuihub/librarian/commit/fb2084f4eb9498f789f7b291394bb6679a8ec98d))

## [0.1.12](https://github.com/tuihub/librarian/compare/v0.1.11...v0.1.12) (2023-10-14)


### Bug Fixes

* add go-favicon ([94cb00f](https://github.com/tuihub/librarian/commit/94cb00f31b60bab771a245fbaa77c04040bd5300))
* reduce disk writes ([2c28a36](https://github.com/tuihub/librarian/commit/2c28a361f30c5152dc56db2f67440196d5c13c71))
* reduce disk writes ([1552716](https://github.com/tuihub/librarian/commit/15527166bfce6e0705757f0ce32a9e7c9663d60d))
* several fixes ([ee4c0cc](https://github.com/tuihub/librarian/commit/ee4c0cc4f03a9f091728bfda653fbfa4a243f935))

## [0.1.11](https://github.com/TuiHub/Librarian/compare/v0.1.10...v0.1.11) (2023-08-11)


### Bug Fixes

* try fix release ci ([eacf514](https://github.com/TuiHub/Librarian/commit/eacf514d54e148114f21390a5f0fcbb780b9c30b))

## [0.1.10](https://github.com/TuiHub/Librarian/compare/v0.1.9...v0.1.10) (2023-08-11)


### Bug Fixes

* fix lint ([6d33128](https://github.com/TuiHub/Librarian/commit/6d33128e662a5cd014f0e2f6db86738d6e02a0d5))

## [0.1.9](https://github.com/TuiHub/Librarian/compare/v0.1.8...v0.1.9) (2023-08-11)


### Features

* add validator middleware ([5f0a1ab](https://github.com/TuiHub/Librarian/commit/5f0a1ab7f5e3400917297b12b4e809236fc11faf))
* Update go to 1.20 ([ab0e1fe](https://github.com/TuiHub/Librarian/commit/ab0e1fe8432ae785a8c80641bac47e3b009f0b9f))


### Bug Fixes

* **gebura:** flatten app info in GetPurchasedApps ([8e7396a](https://github.com/TuiHub/Librarian/commit/8e7396adba4bc2251bbca36dfa267ed32d1c977b))
* **gebura:** update appPackage logic ([426686e](https://github.com/TuiHub/Librarian/commit/426686ef47a74c83ae6d057f962154e8e25b75e6))
* several fix ([2756e3b](https://github.com/TuiHub/Librarian/commit/2756e3b957098e6a0f6a92ad85279df9cdb91d8f))

## [0.1.8](https://github.com/TuiHub/Librarian/compare/v0.1.7...v0.1.8) (2023-07-15)


### Features

* add zh locale support ([0e74203](https://github.com/TuiHub/Librarian/commit/0e74203649174bf1b2041a116fae7bbaf54aed9e))


### Bug Fixes

* **gebura:** correct search & flatten logic ([840b234](https://github.com/TuiHub/Librarian/commit/840b234b8dcae0291994ff311605013709619d5e))

## [0.1.7](https://github.com/TuiHub/Librarian/compare/v0.1.6...v0.1.7) (2023-07-14)


### Features

* **searcher:** add index support ([9cb9c35](https://github.com/TuiHub/Librarian/commit/9cb9c35808cfd34eab1ea60a53e6244187b190e0))
* use searcher to search app ([704913a](https://github.com/TuiHub/Librarian/commit/704913a6c0a9327b973a272824dff078c0477db5))


### Bug Fixes

* avoid duplicated data when re-link account ([8c3193a](https://github.com/TuiHub/Librarian/commit/8c3193a1780e57f288c64f5b3af6baae8c9b8b85))
* closure in libzap ([9f8de5e](https://github.com/TuiHub/Librarian/commit/9f8de5edbff552ab082e0647e1fa3e78ab75bb38))
* correct re-link account ([637d151](https://github.com/TuiHub/Librarian/commit/637d151665e67b1b6eabab2d30a3a67f653591ff))
* **tiphereth:** check bound before bind new account ([4f9a9f0](https://github.com/TuiHub/Librarian/commit/4f9a9f04e568f83b1973d5a13547aece820ec754))
* update purchase logic ([dddc3f4](https://github.com/TuiHub/Librarian/commit/dddc3f48b9590c30a9753bce453eb417c261b1b2))
* upgrade proto to v0.2.37 ([65c30db](https://github.com/TuiHub/Librarian/commit/65c30db0df97fa9c1347b8f8f727cf33007a7213))

## [0.1.6](https://github.com/TuiHub/Librarian/compare/v0.1.5...v0.1.6) (2023-07-04)


### Features

* upgrade proto & add GetApp ListFeedConfigCategories ([41379b1](https://github.com/TuiHub/Librarian/commit/41379b1bc881d059a8c5f41cc1814c9cd684aa5b))

## [0.1.5](https://github.com/TuiHub/Librarian/compare/v0.1.4...v0.1.5) (2023-07-01)


### Features

* support parse digest of feedItem ([ab1572b](https://github.com/TuiHub/Librarian/commit/ab1572b1c8debf44428df65b9d792bdb48aac1a2))


### Bug Fixes

* upgrade protos ([bce9c31](https://github.com/TuiHub/Librarian/commit/bce9c31c49f3a7f139f84096d70ac93cf895d081))
* **yesod:** impl added field ([79abe07](https://github.com/TuiHub/Librarian/commit/79abe0748e87d81fc450dd01077c69e2adfcccd8))

## [0.1.4](https://github.com/TuiHub/Librarian/compare/v0.1.3...v0.1.4) (2023-05-26)


### Features

* add chesed module ([742df82](https://github.com/TuiHub/Librarian/commit/742df82fc39036e024c0f9ea9bd0792f593c788f))
* **chesed:** impl ListImages DownloadImage ([f6e508f](https://github.com/TuiHub/Librarian/commit/f6e508fe3229ce1d42c3dab9282db005e9f8a1a8))
* **chesed:** impl SearchImages PresignedDownloadFile ([4b80aa9](https://github.com/TuiHub/Librarian/commit/4b80aa99cfe08dfa2acc76b641848a268f9f316a))
* impl PresignedUploadFile PresignedUploadFileStatus ([ba65b50](https://github.com/TuiHub/Librarian/commit/ba65b5068eb7ac0c0d56c9129380afd2aed6db72))


### Bug Fixes

* method name ([4159003](https://github.com/TuiHub/Librarian/commit/415900383a29dcdba22ec0739da02d6b226ac440))
* tested UploadImage ([17338cb](https://github.com/TuiHub/Librarian/commit/17338cb2fab65b6c44c21e01067b5cf91932d6ac))
* upgrade proto ([4db56db](https://github.com/TuiHub/Librarian/commit/4db56dbc00b1b116b0a1bb1fdd13588533ecebe8))

## [0.1.3](https://github.com/TuiHub/Librarian/compare/v0.1.2...v0.1.3) (2023-04-22)


### Features

* **chesed:** impl UploadImage ([519de90](https://github.com/TuiHub/Librarian/commit/519de905866c83db27af3f4566ae1653bba3fee2))


### Bug Fixes

* **chesed:** tested UploadImage ([ee5f186](https://github.com/TuiHub/Librarian/commit/ee5f18635a15bd578474c05d39c1356fd98d8a82))
* **tiphereth:** change error reason of GetToken ([ae5ec60](https://github.com/TuiHub/Librarian/commit/ae5ec609f51e82e24444e1164d7e5a341afe4cca))
* **yesod:** follow proto changes ([db288c4](https://github.com/TuiHub/Librarian/commit/db288c4c37ff549fcbce7e26e4571dde6e8af899))

## [0.1.2](https://github.com/TuiHub/Librarian/compare/v0.1.1...v0.1.2) (2023-04-18)


### Bug Fixes

* **angela:** check publishParsed before sort ([3e32080](https://github.com/TuiHub/Librarian/commit/3e32080b857b8a5dc327927979b7dc44147db6ea))

## [0.1.1](https://github.com/TuiHub/Librarian/compare/v0.1.0...v0.1.1) (2023-04-17)


### Features

* impl GetServerInformation ([771aa2c](https://github.com/TuiHub/Librarian/commit/771aa2c15e9f9dbe9691a13a03137f135f074b1b))


### Bug Fixes

* add unique check to feed_config ([4f96349](https://github.com/TuiHub/Librarian/commit/4f96349f396a35af3901931662dd33987a2fa5df))

## [0.1.0](https://github.com/TuiHub/Librarian/compare/v0.0.11...v0.1.0) (2023-04-05)


### ⚠ BREAKING CHANGES

* add consul support

### Features

* add consul support ([dd9a85f](https://github.com/TuiHub/Librarian/commit/dd9a85f0b59c435fcba535a21aa716ea83db0ad2))
* add Netzach module ([92e6863](https://github.com/TuiHub/Librarian/commit/92e686386e2f9fcfde41d5c6bfa3d8f667548d46))
* add ristretto & redis support ([fa38035](https://github.com/TuiHub/Librarian/commit/fa38035cec5d748765d5ead5dabd53b51fae0b0d))
* **porter:** support feature flag ([dded94c](https://github.com/TuiHub/Librarian/commit/dded94cadd5d407a113d9c8df24a6bcb80b170f4))


### Bug Fixes

* allow run in single mode without consul ([bf93c24](https://github.com/TuiHub/Librarian/commit/bf93c2450a5165e9147e53ace3c0e24334eb5791))
* avoid repeat queue ([dd2c9ca](https://github.com/TuiHub/Librarian/commit/dd2c9ca161140ac89a184a0152f0f1295db78537))
* tested push to telegram ([a2886d4](https://github.com/TuiHub/Librarian/commit/a2886d41bfc31883d3e7f08e57475964c4545233))
* tested ReportAppPackages ([92ff407](https://github.com/TuiHub/Librarian/commit/92ff40738e39d8d15077692a3370e95d01d792ab))
* use correct conversion ([f5f25e9](https://github.com/TuiHub/Librarian/commit/f5f25e9c38e43402371f6ec5f948548cfb01cc47))

## [0.0.11](https://github.com/TuiHub/Librarian/compare/v0.0.10...v0.0.11) (2023-03-29)


### Bug Fixes

* ci config ([71131ff](https://github.com/TuiHub/Librarian/commit/71131fffb8db838aae195ae4ef49f23c7e3fd5be))

## [0.0.10](https://github.com/TuiHub/Librarian/compare/v0.0.9...v0.0.10) (2023-03-29)


### Features

* add two build mode ([2ea35ef](https://github.com/TuiHub/Librarian/commit/2ea35efa3d64b661f5239ceba8aabe0cc566f64c))
* **gebura:** impl MergeApps SearchApps PurchaseApp GetAppLibrary ([98a8dda](https://github.com/TuiHub/Librarian/commit/98a8dda3d4dd653fc5e2287f67ab76524a9ff849))
* **gebura:** impl UnAssignAppPackage & review ([91bc8ab](https://github.com/TuiHub/Librarian/commit/91bc8ab5289a31487a0c8f4f90b1ab6b67005d58))


### Bug Fixes

* improve server config & startup logic ([061e3ac](https://github.com/TuiHub/Librarian/commit/061e3ac654518536f51735d85122c3a0030692fc))
* PullSteamApp ([537ae7d](https://github.com/TuiHub/Librarian/commit/537ae7dd493d8da248d4fd94ea66cf2dd955096a))

## [0.0.9](https://github.com/TuiHub/Librarian/compare/v0.0.8...v0.0.9) (2023-03-20)


### Features

* **searcher:** add bleve support & impl NewBatchIDs DescribeID SearchID ([9b0eb88](https://github.com/TuiHub/Librarian/commit/9b0eb8824cd1b8d42c760da1c0a862ede161c887))


### Bug Fixes

* NewTopic easier to use ([dcff5f4](https://github.com/TuiHub/Librarian/commit/dcff5f4476ce2592554c6d20568b9d68ca307bd0))
* update proto to 0.2.14 & impl GroupFeedItems ([af25691](https://github.com/TuiHub/Librarian/commit/af25691ae502f404a0029cb1889e0db7fe723506))
* upgrade protos ([01a8731](https://github.com/TuiHub/Librarian/commit/01a87311320f873b5168a1f5a7cb36f00e638a59))

## [0.0.8](https://github.com/TuiHub/Librarian/compare/v0.0.7...v0.0.8) (2023-03-04)


### Features

* impl GetUser ListFeeds ListFeedItems GetFeedItem GetBatchFeedItems ([d5cbb6c](https://github.com/TuiHub/Librarian/commit/d5cbb6c1a4311460a267581b5879d4d32f73f027))
* **mq:** add postgres driver ([88966e8](https://github.com/TuiHub/Librarian/commit/88966e8a8600cc221b0a164a47bd0fcfeace2620))
* tested Yesod interfaces ([f5b6e51](https://github.com/TuiHub/Librarian/commit/f5b6e515bd4681c4349a7c439029c12b470da224))


### Bug Fixes

* update protos to 0.2.9 & fix angela logic ([73cbc06](https://github.com/TuiHub/Librarian/commit/73cbc06e85cc2787da51ec70ef3f806a0386fc29))

## [0.0.7](https://github.com/TuiHub/Librarian/compare/v0.0.6-4...v0.0.7) (2023-02-28)


### Bug Fixes

* add data path config ([49681c7](https://github.com/TuiHub/Librarian/commit/49681c756a8b89e4e3eb0dc79ae18f528c49d056))

## [0.0.6-4](https://github.com/TuiHub/Librarian/compare/v0.0.6-3...v0.0.6-4) (2023-02-28)


### Continuous Integration

* fix config ([c1e3c2c](https://github.com/TuiHub/Librarian/commit/c1e3c2c8f03643800666f04325f882811cf9a6f5))

## [0.0.6-3](https://github.com/TuiHub/Librarian/compare/v0.0.6-2...v0.0.6-3) (2023-02-28)


### Continuous Integration

* fix config ([e32ce6a](https://github.com/TuiHub/Librarian/commit/e32ce6a18fdf092d7213d58315f49161103cf6f6))

## [0.0.6-2](https://github.com/TuiHub/Librarian/compare/v0.0.6-1...v0.0.6-2) (2023-02-28)


### Continuous Integration

* fix login step ([d6bb345](https://github.com/TuiHub/Librarian/commit/d6bb3457b17523b4faf700e5a88df245c5cdf2c1))

## [0.0.6-1](https://github.com/TuiHub/Librarian/compare/v0.0.6...v0.0.6-1) (2023-02-28)


### Continuous Integration

* add login step ([0e2f3d8](https://github.com/TuiHub/Librarian/commit/0e2f3d8ed212c6b4198012124645f88a25e955e7))

## [0.0.6](https://github.com/TuiHub/Librarian/compare/v0.0.5...v0.0.6) (2023-02-28)


### Features

* tiphereth is ready for fully test ([ca0697e](https://github.com/TuiHub/Librarian/commit/ca0697eda765e21421ecbf3f0f02f57e85ed7c5c))
* tiphereth is ready for fully test ([6020c18](https://github.com/TuiHub/Librarian/commit/6020c182b376d0f086a026859cae7e0d39763634))
* update sql structure & review tiphereth ([ea069e5](https://github.com/TuiHub/Librarian/commit/ea069e53edc8aacdd974c81b5885bc7bd06a2751))


### Bug Fixes

* add logger filter ([1be064a](https://github.com/TuiHub/Librarian/commit/1be064a9ebf283802b84831524c401b8e22f2d6f))
* tested Token and User interfaces ([cd5f646](https://github.com/TuiHub/Librarian/commit/cd5f646ad00893f38e378c79629f5d39377543af))

## [0.0.5](https://github.com/TuiHub/Librarian/compare/v0.0.4-4...v0.0.5) (2023-02-25)


### Features

* add gocron ([560f4a6](https://github.com/TuiHub/Librarian/commit/560f4a6a08bb9a6f069ac1cffc6ef19590ce53af))
* add gofeed ([ec1ce3d](https://github.com/TuiHub/Librarian/commit/ec1ce3d1155ddaccb82798e5645f3964ab8b7184))
* add new ent schema ([97600cc](https://github.com/TuiHub/Librarian/commit/97600cc06bac766eddab4dcc1873ab628a9f0134))
* impl CreateFeedConfig UpdateFeedConfig ([5dec242](https://github.com/TuiHub/Librarian/commit/5dec24202385d38af2a5dd787eb9e3a61b998866))
* use goverter to convert struct ([1b055b3](https://github.com/TuiHub/Librarian/commit/1b055b33a624484191e12d367879ed0c7db3b76b))


### Bug Fixes

* better log format ([5291f48](https://github.com/TuiHub/Librarian/commit/5291f4891e9228f52d8e3954fcc12db0741abffb))
* tested PullFeed ([26ca088](https://github.com/TuiHub/Librarian/commit/26ca0884358e71c20fe1e4b08b2b3b76032cc5c9))
* tested PullFeed cron ([543389d](https://github.com/TuiHub/Librarian/commit/543389df1050447db6c4fb0544c72327339c9832))


### Miscellaneous Chores

* release 0.0.5 ([4700a19](https://github.com/TuiHub/Librarian/commit/4700a19c8788e6c82d4b7799e2aa399de981dd07))

## [0.0.4-4](https://github.com/TuiHub/Librarian/compare/v0.0.4-3...v0.0.4-4) (2023-02-23)


### Miscellaneous Chores

* **ci:** fix config ([785549b](https://github.com/TuiHub/Librarian/commit/785549b09de5ef136acb02ce3db00d70331bb24a))

## [0.0.4-3](https://github.com/TuiHub/Librarian/compare/v0.0.4-2...v0.0.4-3) (2023-02-23)


### Miscellaneous Chores

* **ci:** use goreleaser to build binaries ([#29](https://github.com/TuiHub/Librarian/issues/29)) ([d766868](https://github.com/TuiHub/Librarian/commit/d7668681452754d239fe1cba2b7413a32e747a17))

## [0.0.4-2](https://github.com/TuiHub/Librarian/compare/v0.0.4-1...v0.0.4-2) (2023-02-22)


### Miscellaneous Chores

* **ci:** fix config ([0e685e6](https://github.com/TuiHub/Librarian/commit/0e685e6f4d3822a42e6253f71a48eb8be1726cfd))

## [0.0.4-1](https://github.com/TuiHub/Librarian/compare/v0.0.4...v0.0.4-1) (2023-02-22)


### Miscellaneous Chores

* **ci:** fix config ([d151369](https://github.com/TuiHub/Librarian/commit/d1513690b56efab773e5a9609d22994d1302b131))

## [0.0.4](https://github.com/TuiHub/Librarian/compare/v0.0.3...v0.0.4) (2023-02-22)


### Features

* **gebura:** impl CreateAppPackage UpdateAppPackage ListAppPackage BindAppPackage ReportAppPackage ([4258ee5](https://github.com/TuiHub/Librarian/commit/4258ee5a76f8301f62798e92733febd9395693f0))
* **porter:** add s3 support ([3300e1e](https://github.com/TuiHub/Librarian/commit/3300e1efc426e4a53d7333069f7d3929683d5d51))
* upgrade to go 1.19 && golangci-lint 1.50.1 ([#25](https://github.com/TuiHub/Librarian/issues/25)) ([59c28de](https://github.com/TuiHub/Librarian/commit/59c28de3aedea456d20417a302ef3479d012c354))


### Bug Fixes

* **sephirah:** set right vertex type ([27a834f](https://github.com/TuiHub/Librarian/commit/27a834fac35e035ed86aa28c2c49fc507e18803e))

## [0.0.3](https://github.com/TuiHub/Librarian/compare/v0.0.2...v0.0.3) (2022-11-19)


### Features

* add grpc_web support ([0832d3a](https://github.com/TuiHub/Librarian/commit/0832d3a6c515b6abe305b9caa3d692636ca58a6c))
* add grpc_web support ([643dd26](https://github.com/TuiHub/Librarian/commit/643dd26bdbb49cb43a665ed7ac7407252988417f))
* Mapper: impl edge rules ([d144d30](https://github.com/TuiHub/Librarian/commit/d144d30d7105c9c51633ff7f14167acd597ebe02))
* Sephirah: complete mapper function call for existing port ([0e2d07b](https://github.com/TuiHub/Librarian/commit/0e2d07b63e4f4c73494dfd9c60368631f20d18e2))


### Bug Fixes

* Tested link steam account with Mapper working ([7b29091](https://github.com/TuiHub/Librarian/commit/7b29091c639dfa603bb8b26d262b4477ca4cd76a))

## [0.0.2](https://github.com/TuiHub/Librarian/compare/v0.0.1...v0.0.2) (2022-10-03)


### Features

* Angela: add mq support ([b7fc3a7](https://github.com/TuiHub/Librarian/commit/b7fc3a779b3ebb971a3c330ec1206df4c1165b3f))
* Gebura: impl BindApp ([0ba7a6a](https://github.com/TuiHub/Librarian/commit/0ba7a6a5f06639bb7c6bc39af3d65fd6d9b7774e))
* link steam account generally available ([545091a](https://github.com/TuiHub/Librarian/commit/545091a890c04dd4a4018b504797686555f4ce36))
* Mapper: impl InsertVertex InsertEdge FetchEqualVertex ([6851088](https://github.com/TuiHub/Librarian/commit/6851088c027e0d9dd7d9ba396cfdedc8a735381b))
* Porter: add steam api support ([8deaec2](https://github.com/TuiHub/Librarian/commit/8deaec2c7b66d7e6baf8985033035d84abe8bc6d))
* Porter: impl PullAccount PullApp PullAccountAppRelation ([237a8b9](https://github.com/TuiHub/Librarian/commit/237a8b955f167cfbedaa319ad2546be6f5f05c00))
* Sephirah: impl LinkAccount ([54fa5b3](https://github.com/TuiHub/Librarian/commit/54fa5b3d613287485cdbcd56324a85303c27e7f7))


### Bug Fixes

* wire config ([88c6227](https://github.com/TuiHub/Librarian/commit/88c622746dd7a27eb7b7cba509358d732c1255ce))

## 0.0.1 (2022-09-30)


### Miscellaneous Chores

* release 0.0.1 ([dfba918](https://github.com/TuiHub/Librarian/commit/dfba9187eb248d9473113e64a412235663c85e0a))
