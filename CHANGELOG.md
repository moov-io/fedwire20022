## v0.3.5 (Released 2026-08-26)

BUG FIXES

- fix: accept ISO 20022 times without timezone or with fractional seconds ([#62](https://github.com/moov-io/fedwire20022/pull/62))
  - Fed TODI traffic sends `xs:time` values like `10:51:23.123`. Unmarshal only tried `Z` and `±hh:mm` layouts, so those messages failed to parse.

BUILD

- ci: run CodeQL only on master and PRs to master ([#61](https://github.com/moov-io/fedwire20022/pull/61))
- ci: use Renovate for automerge, drop Dependabot
- ci: skip CodeQL when GitHub cannot accept SARIF ([#60](https://github.com/moov-io/fedwire20022/pull/60))
- ci: skip code owners on go.mod/go.sum ([#59](https://github.com/moov-io/fedwire20022/pull/59))

## v0.3.4 (Released 2026-08-25)

BUG FIXES

- fix: keep underscores on `FedwireFundsAccountBalanceReport_Master` and `FedwireFundsAccountBalanceReport_Self` so live Fed `camt.052` account-balance reports unmarshal ([#58](https://github.com/moov-io/fedwire20022/pull/58))
  - Generated tags had dropped the XSD underscores, which made mq-relay nack TODI replies into a redelivery loop.

BUILD

- fix(deps): update module github.com/moov-io/base to v0.63.3
- fix(deps): update module github.com/stretchr/testify to v1.12.0
- chore(deps): update go toolchain directive to v1.26.6
- security: enable weekly scanners and fix alerts
- security: add Dependabot cooldown
- Add zizmor GitHub Actions security analysis ([#53](https://github.com/moov-io/fedwire20022/pull/53))

## v0.3.3 (Released 2026-08-07)

IMPROVEMENTS

- chore(deps): bump github.com/moov-io/base to v0.62.1 ([#48](https://github.com/moov-io/fedwire20022/pull/48))
  - Bound `pgxpool.Close` wait on `sql.DB` shutdown to avoid process-exit deadlocks (Postgres/AlloyDB)
- test(fuzz): add Fedwire Funds XML and type fuzzers
- ci: add scheduled Go fuzz testing workflow

BUILD

- chore(deps): update go toolchain directive to v1.26.5 ([#46](https://github.com/moov-io/fedwire20022/pull/46))
- fix(deps): update module github.com/moov-io/base to v0.61.2 ([#43](https://github.com/moov-io/fedwire20022/pull/43))

## v0.3.2 (Released 2026-05-07)

IMPROVEMENTS

- fedwire: restore Z on ISOTime marshal
- fix: remove Z from ISODate marshal

## v0.3.1 (Released 2026-05-01)

IMPROVEMENTS

- fix: remove hyphens from IMAD cycle date

## v0.3.0 (Released 2026-04-22)

IMPROVEMENTS

- fix: correctly set xmlns ([#41](https://github.com/moov-io/fedwire20022/pull/41))
- feat: add fedwiretest for flipping messages (in testing)
- templates: updates to parse incoming xml
- docs: explain incoming/outgoing naming

BUILD

- chore(deps): update dependency go to v1.26.2 ([#40](https://github.com/moov-io/fedwire20022/pull/40))

## v0.2.0 (Released 2026-03-16)

ADDITIONS

- feat: add Input Message Accountability Data (IMAD)
- feat: add BusinessMessageID helpers
- feat: add BusinessMessageIDHash for consistent message IDs

IMPROVEMENTS

- fix: switch to correct xsd

BUILD

- chore(deps): update dependency go to v1.26.1 ([#39](https://github.com/moov-io/fedwire20022/pull/39))
- fix(deps): update module github.com/moov-io/base to v0.61.1 ([#38](https://github.com/moov-io/fedwire20022/pull/38))

## v0.1.0 (Released 2024-11-22)

INITIAL RELEASE

This is the first release of moov-io/fedwire20022 so please try it out and give us some feedback! Let us know what works and doesn't for your usecase.
Join our slack channel (`#wire`) on the [moov-io slack](https://slack.moov.io/) to have an interactive discussion about the development of the project.
