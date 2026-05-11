# MySQL 5.7 to 8.0 Migration Plan — Zero Downtime, 2 Weeks, Under $10k

## Overview

This plan outlines a two-week migration from MySQL 5.7 to 8.0 with zero downtime using a blue/green replication strategy. The approach keeps production on 5.7 until the new instance is fully validated, then cuts over with a brief read-only window.

---

## Week 1 — Preparation and Parallel Setup

### Days 1–2: Compatibility Audit
- Run the [MySQL Shell Upgrade Checker](https://dev.mysql.com/doc/mysql-shell/8.0/en/mysql-shell-utilities-upgrade.html) against the production instance: `mysqlcheck --all-databases` and `util.checkForServerUpgrade()`.
- Identify and fix:
  - Deprecated or removed SQL modes (e.g., `NO_ZERO_DATE`, `NO_ZERO_IN_DATE` enforcement changes).
  - `utf8` charset usage that behaves differently vs `utf8mb4` in 8.0.
  - `GROUP BY` queries relying on non-deterministic implicit ordering (removed in 8.0).
  - Stored procedures, triggers, or views using reserved words added in 8.0.
  - Any use of `SET @@session.sql_log_bin` (now requires `BINLOG_ADMIN`).
- Estimated effort: 1–2 engineers, 2 days.

### Days 3–4: Spin Up MySQL 8.0 Replica
- Provision a new MySQL 8.0 instance (same or slightly larger spec as current primary).
  - Option A: Cloud managed (RDS, Cloud SQL) — adds ~$300–600/month during transition (~$150 for 2-week overlap).
  - Option B: Self-managed VM — similar cost on major clouds.
- Configure MySQL 8.0 as a replica of the existing 5.7 primary using GTID-based replication.
  - Note: MySQL 8.0 can replicate from 5.7 as a downstream replica.
  - Enable binary logging on 5.7 if not already active.
  - Set `gtid_mode=ON` and `enforce_gtid_consistency=ON` on 5.7 (requires a staged enable — `OFF → WARN → ON_PERMISSIVE → ON`; plan a rolling restart).
- Seed the replica via `mysqldump` or `xtrabackup` snapshot.

### Day 5: Validate Replication Lag
- Monitor replication lag (`Seconds_Behind_Master`) and confirm it stays near zero under normal load.
- Run application read queries against the 8.0 replica in shadow/read-only mode and compare results.

### Days 6–7: Application Compatibility Testing
- Point a staging environment at the MySQL 8.0 instance.
- Run the full integration and regression test suite.
- Specifically test:
  - Authentication: 8.0 defaults to `caching_sha2_password`; older drivers may need `mysql_native_password` or driver upgrades.
  - JSON functions that changed behavior.
  - Window functions and CTEs (these are additive, but verify no naming conflicts).
- Fix any incompatibilities found.

---

## Week 2 — Validation, Cutover, and Stabilization

### Days 8–10: Load Testing and Final Checks
- Run load tests against the 8.0 staging instance at 1.5× production traffic levels.
- Confirm query performance — especially any that relied on implicit `GROUP BY` ordering or used `ONLY_FULL_GROUP_BY` workarounds.
- Verify backup/restore procedures work on 8.0.
- Confirm monitoring and alerting (slow query log, error log, replication lag alerts) are configured on the new instance.

### Day 11: Pre-Cutover Readiness Review
- Confirm replication lag is consistently < 1 second.
- Confirm all application compatibility issues are resolved and deployed to production code.
- Schedule the cutover window — ideally lowest-traffic period (e.g., 2–4 AM).
- Notify stakeholders and on-call team.

### Day 12: Cutover
The cutover sequence for zero (or near-zero) downtime:

1. **Enable maintenance mode / read-only** on the application layer (return HTTP 503 for writes, or use a feature flag). This window should be under 60 seconds.
2. Confirm replication lag on 8.0 replica reaches zero (`Seconds_Behind_Master = 0`).
3. **Stop replication** on the 8.0 instance and promote it to primary (`STOP REPLICA; RESET REPLICA ALL;`).
4. **Update application database connection strings** to point to the new MySQL 8.0 primary.
5. **Disable read-only / maintenance mode**.
6. Monitor error rates, query latency, and replication status for 30 minutes.

### Days 13–14: Post-Cutover Stabilization
- Keep the MySQL 5.7 instance running (stopped, not terminated) for the rollback window.
- Monitor slow query log for regressions.
- Confirm all background jobs and cron tasks ran successfully on 8.0.
- After 48 hours of clean operation, decommission 5.7 instance.

---

## Rollback Conditions

Initiate rollback if any of the following occur during or after cutover:

| Condition | Threshold |
|---|---|
| Error rate spike | > 2× baseline for > 5 minutes |
| p99 query latency | > 3× baseline sustained |
| Application exceptions related to DB | Any new DB-related exception class |
| Replication divergence detected | Any checksum mismatch |
| Critical business function failure | Payment, auth, or core data reads failing |

**Rollback procedure:**
1. Point application connection strings back to the 5.7 instance.
2. Re-enable 5.7 as primary (it should still be live and up-to-date up to the cutover point).
3. Any writes made to 8.0 after cutover will need to be evaluated — if the rollback window is short (< 10 minutes), assess whether re-syncing those writes is feasible or whether they must be replayed manually.
4. Page the team and conduct a post-mortem before re-attempting migration.

---

## Cost Estimate

| Item | Cost |
|---|---|
| Extra MySQL 8.0 instance (2 weeks) | ~$150–400 |
| Additional storage for parallel instance | ~$50–100 |
| Engineer time (2 engineers × ~5 days effective) | ~$4,000–7,000 (internal cost) |
| Tooling / mysqldump / xtrabackup (open source) | $0 |
| Load testing infrastructure (short burst) | ~$50–150 |
| Buffer for unexpected remediation work | ~$500–1,000 |
| **Total** | **~$5,000–9,000** |

This stays comfortably under the $10k cap assuming internal engineering rates around $100–150/hour and no major external consultant involvement.

---

## Key Risks and Mitigations

- **Authentication plugin mismatch**: Test driver compatibility early in Week 1. Update connectors before cutover.
- **`ONLY_FULL_GROUP_BY` strictness**: Run the upgrade checker and fix all offending queries before Week 2.
- **GTID enablement on 5.7**: This requires a staged process and a restart of 5.7. Schedule this carefully to minimize impact.
- **Timezone data**: MySQL 8.0 ships with updated timezone tables; verify `CONVERT_TZ()` and timestamp behavior.
- **Data type changes**: `YEAR(2)` is not supported; `FLOAT(M,D)` and `DOUBLE(M,D)` syntax is deprecated — audit schema.
