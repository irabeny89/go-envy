## Unreleased (e35ad22..987289d)

#### Miscellaneous Tasks

- update changelog and readme formatting - (987289d) - Irabeny
- align commit types and advise dev to use cocogitto - (e35ad22) - Irabeny

- - -

## v0.4.0 - 2026-04-22

#### Features

- add git hook setup script for conventional commits - (fc86ec7) - Irabeny

#### Documentation

- improve formatting and readability of CI/CD pipeline documentation - (4865f96) - Irabeny
- add CI/CD documentation - (3f65382) - Irabeny

#### Miscellaneous Tasks

- (**version**) v0.3.1 - (4bc36a9) - Irabeny

- - -

## v0.3.0 - 2026-04-22

#### Features

- (**ci**) configure tag prefix and improve release workflow - (689d0b7) - Irabeny
- add GOENVY_DEFAULT_PATH support to allow configuring the default .env file path - (fda2c0e) - Irabeny

#### Bug Fixes

- (**ci**) remove redundant v prefix from release workflow - (091a271) - Irabeny

#### Performance Improvements

- add optional path param to LoadEnv - (bf3607c) - Irabeny

#### Documentation

- (**test**) add package comment to tests - (cc77609) - Irabeny
- rewrite README to include installation, usage examples, and configuration documentation - (b5ff098) - Irabeny
- show how to use LoadEnv with path param - (ee975a8) - Irabeny

#### Refactorings

- add automated .env loading, improve multiline parsing, and update API to return errors - (dd4eeed) - Irabeny

- - -

## v0.2.0 - 2026-04-22

#### Features

- add function to load environment variables using a path argument - (9d6cb75) - Irabeny

#### Documentation

- update readme with new feature LoadEnvPath usage - (90ca88c) - Irabeny
- add optional to step 2 of goenvy package docs - (4549f3f) - Irabeny
- improve goenvy package documentation - (6121d0a) - Irabeny

#### Tests

- test LoadEnvPath function and update env files for this purpose - (e20dc09) - Irabeny

- - -

## v0.1.0 - 2026-04-22

#### Features

- initial commit to load env - (d757b33) - Irabeny

#### Bug Fixes

- rewrite LoadEnv function to cover edge cases like single line quoted texts - (8e32ecc) - Irabeny

#### Documentation

- add environment variables sample to readme file - (940bfb3) - Irabeny
- add readme file - (6a550f8) - Irabeny
- add code example to package - (db3f0e6) - Irabeny

#### Tests

- add code test coverage report - (faf2f60) - Irabeny
- test the fix on single line quoted strings and other edge cases - (d350cd1) - Irabeny

#### Continuous Integration

- remove the v letter attached the version that was causing a bug in workflow - (d151e5e) - Irabeny
- update workflow name from Test to CI/CD - (f17e330) - Irabeny
- update ci/cd workflow - (bf9d06f) - Irabeny
- add github actions workflow to test and bump versions - (230e922) - Irabeny

#### Refactorings

- update github actions workflow - (8fa3649) - Irabeny
- remove unnecessary steps from workflow - (85741d7) - Irabeny
- rewrite workflow - (9e935d6) - Irabeny

#### Miscellaneous Tasks

- add cocogitto semantic version config - (7d84356) - Irabeny
- add test github actions workflow - (01bc1a7) - Irabeny
- update variables in .env file - (b353b29) - Irabeny
- update go module directive value - (e490eef) - Irabeny

#### Styles

- (**workflow**) prefix semantic version with letter v - (8fcec01) - Irabeny
