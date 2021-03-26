|   REVIEWER    |  O(C)   | 1 | 2 | 3 | 4 | 5 | 6 | 7 | T  |
|---------------|---------|---|---|---|---|---|---|---|----|
| Reminiscent   |  7 ( 0) | 4 | 1 | 2 | 1 | 0 | 0 | 3 | 11 |
| SunRunAway    | 15 ( 1) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| XuHuaiyu      | 77 ( 3) | 0 | 1 | 0 | 0 | 0 | 0 | 0 |  1 |
| dyzsr         |  0 ( 0) | 1 | 0 | 0 | 0 | 0 | 0 | 0 |  1 |
| eurekaka      | 23 ( 1) | 4 | 1 | 0 | 1 | 0 | 0 | 0 |  6 |
| fzhedu        |  2 ( 0) | 0 | 2 | 0 | 0 | 0 | 0 | 1 |  3 |
| ichn-hu       |  3 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| lzmhhh123     | 56 ( 0) | 2 | 1 | 1 | 0 | 0 | 0 | 1 |  5 |
| nullnotnil    |  0 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| qw4990        | 65 ( 1) | 6 | 1 | 3 | 2 | 0 | 0 | 3 | 15 |
| rebelice      |  2 ( 0) | 1 | 2 | 0 | 1 | 0 | 0 | 0 |  4 |
| time-and-fate |  4 ( 0) | 3 | 0 | 0 | 0 | 0 | 0 | 0 |  3 |
| winoros       | 34 ( 2) | 8 | 6 | 3 | 3 | 0 | 0 | 2 | 22 |
| wshwsh12      | 25 ( 1) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |


<details> 
  <summary>Reminiscent's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                     PR                                                                      | C | LASTED |
|------------|---------------------------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/21896 | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                                  |   | 94d19h |
| tidb/22354 | [planner: do not cache prepared plan if optimization depends on mutable constant (#22349)](https://github.com/pingcap/tidb/pull/22354)      |   | 72d22h |
| tidb/23293 | [planner: fix the bug that wrong collation is used when try fast path for enum or set (#23217)](https://github.com/pingcap/tidb/pull/23293) |   | 13d14h |
| tidb/23441 | [executor: Refactor probe channel & fix bug in chunks of join](https://github.com/pingcap/tidb/pull/23441)                                  |   | 6d14h  |
| tidb/23469 | [*: hide `index-usage-sync-lease` config (#23349)](https://github.com/pingcap/tidb/pull/23469)                                              |   | 2d21h  |
| tidb/23474 | [planner: fix inappropriate null flag of null constants (#23457)](https://github.com/pingcap/tidb/pull/23474)                               |   | 2d18h  |
| tidb/23493 | [expression: Implementation of Vitess hashing algorithm.](https://github.com/pingcap/tidb/pull/23493)                                       |   | 2d6h   |


## Reviewed in Last 7 Days

|    REPO    |                                                                PR                                                                | C | D |   R    |
|------------|----------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/23536 | [planner: remove some risky cache operations in the plan builder (#23354)](https://github.com/pingcap/tidb/pull/23536)           |   | 1 | 0h     |
| tidb/23524 | [partition: fix hash partition with not between condition get wrong result (#22914)](https://github.com/pingcap/tidb/pull/23524) |   | 1 | 2h     |
| tidb/23525 | [partition: fix hash partition with not between condition get wrong result (#22914)](https://github.com/pingcap/tidb/pull/23525) |   | 1 | 2h     |
| tidb/23520 | [statistics: optimize the global histogram merging algorithm](https://github.com/pingcap/tidb/pull/23520)                        |   | 1 | 3h     |
| tidb/23481 | [planner: fix inappropriate null flag of null constants (#23457)](https://github.com/pingcap/tidb/pull/23481)                    |   | 2 | 22h    |
| tidb/23486 | [statistics: recalculate bucket repeats when merging global statistics](https://github.com/pingcap/tidb/pull/23486)              |   | 3 | 14h    |
| tidb/23283 | [util: optimize the performance of restore with db (#22910)](https://github.com/pingcap/tidb/pull/23283)                         |   | 3 | 10d23h |
| tidb/23457 | [planner: fix inappropriate null flag of null constants](https://github.com/pingcap/tidb/pull/23457)                             |   | 4 | 0h     |
| tidb/23346 | [statistics: fix build extended stats panic when met NULL](https://github.com/pingcap/tidb/pull/23346)                           |   | 7 | 2d18h  |
| tidb/23119 | [statistics: remove existing deleted extended stats when add a new one](https://github.com/pingcap/tidb/pull/23119)              |   | 7 | 14d14h |
| tidb/23413 | [statistics: remove the dependency of stats GC on LastUpdateVersion](https://github.com/pingcap/tidb/pull/23413)                 |   | 7 | 14h    |


</details> 


<details> 
  <summary>SunRunAway's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                  PR                                                                   | C | LASTED  |
|------------|---------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19178 | [executor: Refactor probe channel](https://github.com/pingcap/tidb/pull/19178)                                                        |   | 224d16h |
| tidb/19807 | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                               |   | 202d10h |
| tidb/19900 | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                        | Y | 197d18h |
| tidb/20140 | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                    |   | 184d22h |
| tidb/20765 | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                     |   | 143d16h |
| tidb/21207 | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                       |   | 122d19h |
| tidb/21834 | [planner: enhanced index range calculation plan](https://github.com/pingcap/tidb/pull/21834)                                          |   | 99d18h  |
| tidb/21876 | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)           |   | 97d19h  |
| tidb/21878 | [planner: do not push down lock to pointGet/bacthPointGet when selection exists](https://github.com/pingcap/tidb/pull/21878)          |   | 97d18h  |
| tidb/21956 | [planner/preprocessor: disallow into-outfile clause in some place](https://github.com/pingcap/tidb/pull/21956)                        |   | 92d23h  |
| tidb/22026 | [expression: separated arithmeticPlusIntSig](https://github.com/pingcap/tidb/pull/22026)                                              |   | 90d20h  |
| tidb/22114 | [test: fix globalkilltest (#21987)](https://github.com/pingcap/tidb/pull/22114)                                                       |   | 85d12h  |
| tidb/22217 | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                         |   | 78d17h  |
| tidb/22365 | [planner: check index valid while forUpdateRead (#22152)](https://github.com/pingcap/tidb/pull/22365)                                 |   | 72d19h  |
| tidb/22379 | [[experiment] executor: allow aggregation to spill disk when running out of memory quota](https://github.com/pingcap/tidb/pull/22379) |   | 71d19h  |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>XuHuaiyu's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                              PR                                                                               | C | LASTED  |
|--------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| docs-cn/5619 | [Update data-type-date-and-time.md](https://github.com/pingcap/docs-cn/pull/5619)                                                                             |   | 27d15h  |
| docs/5075    | [hide the config global-kill](https://github.com/pingcap/docs/pull/5075)                                                                                      |   | 3d17h   |
| tidb/19900   | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                                                | Y | 197d18h |
| docs-cn/5671 | [tidb: Add time format description](https://github.com/pingcap/docs-cn/pull/5671)                                                                             |   | 21d11h  |
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                                        | Y | 195d13h |
| tidb/20140   | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                                            |   | 184d22h |
| tidb/20311   | [expression: fix overflow error when convert bit to int64 (#20266)](https://github.com/pingcap/tidb/pull/20311)                                               |   | 176d21h |
| tidb/20790   | [collation: add pinyin collation for chinese charset support](https://github.com/pingcap/tidb/pull/20790)                                                     |   | 142d20h |
| tidb/21064   | [planner, executor: fix cast not check error](https://github.com/pingcap/tidb/pull/21064)                                                                     |   | 130d8h  |
| tidb/21149   | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                                 |   | 126d14h |
| tidb/21228   | [executor: return the result immediately when combining LIMIT row_count with DISTINCT](https://github.com/pingcap/tidb/pull/21228)                            |   | 122d13h |
| tidb/21304   | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                                      |   | 120d12h |
| tidb/21334   | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                                                 |   | 119d14h |
| tidb/21401   | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                                            |   | 115d11h |
| tidb/21476   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                         |   | 112d15h |
| tidb/21536   | [executor: add slow-log file meta cache to avoid repeat read file meta information](https://github.com/pingcap/tidb/pull/21536)                               |   | 108d14h |
| tidb/21564   | [ddl: fix Incorrect behavior of NO_ZERO_DATE when altering table](https://github.com/pingcap/tidb/pull/21564)                                                 |   | 107d15h |
| tidb/21853   | [expression: fix compatibility behaviors in time_format with MySQL (#21559)](https://github.com/pingcap/tidb/pull/21853)                                      |   | 98d18h  |
| tidb/21896   | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                                                    |   | 94d19h  |
| tidb/22131   | [privilege: remove leading and trailing space when create user and role](https://github.com/pingcap/tidb/pull/22131)                                          |   | 84d19h  |
| tidb/22149   | [session: set process info before building plan (#22101)](https://github.com/pingcap/tidb/pull/22149)                                                         |   | 80d19h  |
| tidb/22163   | [expression: separated arithmeticMinusIntSig](https://github.com/pingcap/tidb/pull/22163)                                                                     |   | 80d13h  |
| tidb/22186   | [executor: fix select into outfile with year type column has no data (#22175)](https://github.com/pingcap/tidb/pull/22186)                                    |   | 79d16h  |
| tidb/22294   | [planner, table: optimize the list column partition pruner for range query](https://github.com/pingcap/tidb/pull/22294)                                       |   | 76d19h  |
| tidb/22307   | [ddl: fix update can see columns not public](https://github.com/pingcap/tidb/pull/22307)                                                                      |   | 76d15h  |
| tidb/22616   | [expression: from_unixtime accept 64-bit integers](https://github.com/pingcap/tidb/pull/22616)                                                                |   | 55d23h  |
| tidb/22617   | [metrics: fix wrong bucket name of coprocessor cache (#22454)](https://github.com/pingcap/tidb/pull/22617)                                                    |   | 55d23h  |
| tidb/22624   | [ planner: not pruning column used by union scan condition (#21640)](https://github.com/pingcap/tidb/pull/22624)                                              |   | 55d16h  |
| tidb/22631   | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                                                                               |   | 53d22h  |
| tidb/22696   | [expression: enable arithmetic Mod push down](https://github.com/pingcap/tidb/pull/22696)                                                                     |   | 50d17h  |
| tidb/22711   | [executor: Fix inline schema name](https://github.com/pingcap/tidb/pull/22711)                                                                                |   | 50d11h  |
| tidb/22722   | [planner, errno: make error code of ErrMixOfGroupFuncAndFields consistent with MySQL](https://github.com/pingcap/tidb/pull/22722)                             |   | 49d20h  |
| tidb/22814   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22814)                                           |   | 34d18h  |
| tidb/22815   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22815)                                           |   | 34d18h  |
| tidb/22908   | [txn: Add txn state's view](https://github.com/pingcap/tidb/pull/22908)                                                                                       |   | 29d20h  |
| tidb/22926   | [expression: add overflow check in multiplyInt](https://github.com/pingcap/tidb/pull/22926)                                                                   |   | 29d13h  |
| tidb/23012   | [executor: fix affected rows of ddls and complete uint tests](https://github.com/pingcap/tidb/pull/23012)                                                     |   | 25d16h  |
| tidb/23105   | [executor: fix wrong key range of index scan when filter is comparing year column with NULL (#23079)](https://github.com/pingcap/tidb/pull/23105)             |   | 21d18h  |
| tidb/23128   | [statistics: refactor the statistics package use the RestrictedSQLExecutor API (#22636)](https://github.com/pingcap/tidb/pull/23128)                          |   | 20d22h  |
| tidb/23152   | [expression: fix wrong error info (#22760)](https://github.com/pingcap/tidb/pull/23152)                                                                       |   | 18d14h  |
| tidb/23196   | [types: fix the bug about the wrong query result for decimal type  (#22507)](https://github.com/pingcap/tidb/pull/23196)                                      |   | 16d18h  |
| tidb/23210   | [planner: fixed a bug that prevented SPM from taking effect (#23197)](https://github.com/pingcap/tidb/pull/23210)                                             |   | 16d16h  |
| tidb/23220   | [Release 4.0](https://github.com/pingcap/tidb/pull/23220)                                                                                                     |   | 16d11h  |
| tidb/23227   | [executor: hash join out of index panic when enum column value is zero (#23162)](https://github.com/pingcap/tidb/pull/23227)                                  |   | 15d21h  |
| tidb/23233   | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23233)                                                        |   | 15d18h  |
| tidb/23234   | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23234)                                                        |   | 15d18h  |
| tidb/23245   | [*: Add security enhanced mode as experimental](https://github.com/pingcap/tidb/pull/23245)                                                                   |   | 15d6h   |
| tidb/23257   | [executor: group_concat aggr panic when session.group_concat_max_len is small (#23131)](https://github.com/pingcap/tidb/pull/23257)                           |   | 14d18h  |
| tidb/23278   | [executor: wrong result of nullif expr when used with is null expr. (#23170)](https://github.com/pingcap/tidb/pull/23278)                                     |   | 13d18h  |
| tidb/23295   | [util, types: don't let SPM be affected by charset (#23161)](https://github.com/pingcap/tidb/pull/23295)                                                      |   | 13d11h  |
| tidb/23335   | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23335)                                      |   | 9d19h   |
| tidb/23336   | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23336)                                      |   | 9d19h   |
| tidb/23337   | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23337)                                      |   | 9d19h   |
| tidb/23340   | [executor: fix unexpected NotNullFlag in case when expr ret type (#23102)](https://github.com/pingcap/tidb/pull/23340)                                        |   | 9d19h   |
| tidb/23347   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                                          |   | 9d17h   |
| tidb/23348   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                                          |   | 9d17h   |
| tidb/23350   | [util/stringutil, util/ranger, planner: use hierarchical separators to simplify the parsing for info of EXPLAIN ](https://github.com/pingcap/tidb/pull/23350) |   | 9d17h   |
| tidb/23368   | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23368)                                         |   | 8d20h   |
| tidb/23369   | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23369)                                         |   | 8d20h   |
| tidb/23370   | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23370)                                         |   | 8d20h   |
| tidb/23374   | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23374)                                             |   | 8d19h   |
| tidb/23397   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23397)                                                                |   | 7d17h   |
| tidb/23398   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23398)                                                                |   | 7d17h   |
| tidb/23399   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23399)                                                                |   | 7d17h   |
| tidb/23405   | [domain: remove the exit chan, use context](https://github.com/pingcap/tidb/pull/23405)                                                                       |   | 7d17h   |
| tidb/23433   | [WIP: speed up for slow query logs retrieving ](https://github.com/pingcap/tidb/pull/23433)                                                                   |   | 6d17h   |
| tidb/23435   | [planner: check schema stale for plan cache when forUpdateRead (#22381)](https://github.com/pingcap/tidb/pull/23435)                                          |   | 6d16h   |
| tidb/23441   | [executor: Refactor probe channel & fix bug in chunks of join](https://github.com/pingcap/tidb/pull/23441)                                                    |   | 6d14h   |
| tidb/23474   | [planner: fix inappropriate null flag of null constants (#23457)](https://github.com/pingcap/tidb/pull/23474)                                                 |   | 2d18h   |
| tidb/23487   | [planner: optimize count(distinct a) to count(a) if there is an unique key on a](https://github.com/pingcap/tidb/pull/23487)                                  | Y | 2d13h   |
| tidb/23493   | [expression: Implementation of Vitess hashing algorithm.](https://github.com/pingcap/tidb/pull/23493)                                                         |   | 2d6h    |
| tidb/23497   | [expression: Let TiDB use Hyperscan to support multi-pattern-match](https://github.com/pingcap/tidb/pull/23497)                                               |   | 1d22h   |
| tidb/23517   | [*: Add the metric about the SQL with TiFlash Success  (#23426)](https://github.com/pingcap/tidb/pull/23517)                                                  |   | 1d11h   |
| tidb/23524   | [partition: fix hash partition with not between condition get wrong result (#22914)](https://github.com/pingcap/tidb/pull/23524)                              |   | 22h     |
| tidb/23533   | [executor: fix batchGet/PointGet overflow result in clustered index (#23516)](https://github.com/pingcap/tidb/pull/23533)                                     |   | 19h     |
| tidb/23562   | [execution: reuse iterator in hash join](https://github.com/pingcap/tidb/pull/23562)                                                                          |   | 13h     |
| tidb/23565   | [planner/core: convert decimal type for mpp join before shuffling. (#23191)](https://github.com/pingcap/tidb/pull/23565)                                      |   | 9h      |


## Reviewed in Last 7 Days

|    REPO    |                                          PR                                          | C | D |  R   |
|------------|--------------------------------------------------------------------------------------|---|---|------|
| tidb/23467 | [hot-fix: paginate indexLookUp executor](https://github.com/pingcap/tidb/pull/23467) |   | 2 | 1d0h |


</details> 


<details> 
  <summary>dyzsr's detailed review report</summary> 

## To Be Reviewed

| REPO | PR | C | LASTED |
|------|----|---|--------|


## Reviewed in Last 7 Days

|    REPO    |                                                     PR                                                     | C | D |   R   |
|------------|------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/23491 | [executor,planner: fix update join update unmatched outer row](https://github.com/pingcap/tidb/pull/23491) |   | 1 | 1d20h |


</details> 


<details> 
  <summary>eurekaka's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                   PR                                                                   | C | LASTED  |
|------------|----------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                              |   | 140d16h |
| tidb/21444 | [planner: ignore anonymous index while tiflash replica is available](https://github.com/pingcap/tidb/pull/21444)                       |   | 113d12h |
| tidb/22342 | [session: fix two cases when updating bind info (#22338)](https://github.com/pingcap/tidb/pull/22342)                                  |   | 73d18h  |
| tidb/22354 | [planner: do not cache prepared plan if optimization depends on mutable constant (#22349)](https://github.com/pingcap/tidb/pull/22354) |   | 72d22h  |
| tidb/22369 | [session: fix the duplicate binding case when updating bind info (#22367)](https://github.com/pingcap/tidb/pull/22369)                 |   | 72d17h  |
| tidb/22416 | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                                      | Y | 68d11h  |
| tidb/22559 | [planner: split test data from test cases in cbo_test.go](https://github.com/pingcap/tidb/pull/22559)                                  |   | 57d19h  |
| tidb/22778 | [*: add support for dynamic privileges](https://github.com/pingcap/tidb/pull/22778)                                                    |   | 37d7h   |
| tidb/23137 | [planner: fix index merge row count estimation logic](https://github.com/pingcap/tidb/pull/23137)                                      |   | 20d17h  |
| tidb/23208 | [statistics, util/ranger: improve selectivity calculation for DNF filters (#18741)](https://github.com/pingcap/tidb/pull/23208)        |   | 16d16h  |
| tidb/23210 | [planner: fixed a bug that prevented SPM from taking effect (#23197)](https://github.com/pingcap/tidb/pull/23210)                      |   | 16d16h  |
| tidb/23216 | [Privileges: fix delete privilege check wrongly (#22971)](https://github.com/pingcap/tidb/pull/23216)                                  |   | 16d14h  |
| tidb/23283 | [util: optimize the performance of restore with db (#22910)](https://github.com/pingcap/tidb/pull/23283)                               |   | 13d17h  |
| tidb/23295 | [util, types: don't let SPM be affected by charset (#23161)](https://github.com/pingcap/tidb/pull/23295)                               |   | 13d11h  |
| tidb/23316 | [planner: Fix rebuild range for prepared plan](https://github.com/pingcap/tidb/pull/23316)                                             |   | 10d17h  |
| tidb/23365 | [planner: fix a bug that point get plan returns wrong column name](https://github.com/pingcap/tidb/pull/23365)                         |   | 8d21h   |
| tidb/23373 | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                      |   | 8d19h   |
| tidb/23374 | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23374)                      |   | 8d19h   |
| tidb/23469 | [*: hide `index-usage-sync-lease` config (#23349)](https://github.com/pingcap/tidb/pull/23469)                                         |   | 2d21h   |
| tidb/23536 | [planner: remove some risky cache operations in the plan builder (#23354)](https://github.com/pingcap/tidb/pull/23536)                 |   | 19h     |
| tidb/23543 | [statistics: fix auto analyze log information incomplete (#23522)](https://github.com/pingcap/tidb/pull/23543)                         |   | 18h     |
| tidb/23554 | [executor: fix update panic on join having statement](https://github.com/pingcap/tidb/pull/23554)                                      |   | 16h     |
| tidb/23565 | [planner/core: convert decimal type for mpp join before shuffling. (#23191)](https://github.com/pingcap/tidb/pull/23565)               |   | 9h      |


## Reviewed in Last 7 Days

|     REPO     |                                                              PR                                                              | C | D |   R    |
|--------------|------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/23487   | [planner: optimize count(distinct a) to count(a) if there is an unique key on a](https://github.com/pingcap/tidb/pull/23487) | Y | 1 | 1d19h  |
| tidb/23480   | [planner/core: inject project for tiflash agg](https://github.com/pingcap/tidb/pull/23480)                                   |   | 1 | 1d22h  |
| docs-cn/5811 | [Revert "prepared plan cache: Enable by default"](https://github.com/pingcap/docs-cn/pull/5811)                              |   | 1 | 1d3h   |
| tidb/23522   | [statistics: fix auto analyze log information incomplete](https://github.com/pingcap/tidb/pull/23522)                        |   | 1 | 2h     |
| docs/5088    | [SPM: update DML SQL Bind and baseline capture description](https://github.com/pingcap/docs/pull/5088)                       |   | 2 | 22h    |
| tidb/23175   | [planner: fix plan cache not working caused by type difference](https://github.com/pingcap/tidb/pull/23175)                  |   | 4 | 13d19h |


</details> 


<details> 
  <summary>fzhedu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                            PR                                                            | C | LASTED |
|------------|--------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/22853 | [planner: fix LogicalPlans that contain Window Function are ambiguous ](https://github.com/pingcap/tidb/pull/22853)      |   | 33d12h |
| tidb/23565 | [planner/core: convert decimal type for mpp join before shuffling. (#23191)](https://github.com/pingcap/tidb/pull/23565) |   | 9h     |


## Reviewed in Last 7 Days

|    REPO    |                                                       PR                                                        | C | D |   R   |
|------------|-----------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/23191 | [planner/core: convert decimal type for mpp join before shuffling.](https://github.com/pingcap/tidb/pull/23191) |   | 2 | 15d8h |
| tidb/23441 | [executor: Refactor probe channel & fix bug in chunks of join](https://github.com/pingcap/tidb/pull/23441)      |   | 2 | 5d2h  |
| tidb/23401 | [planner, sessionctx: turn on the mpp by default](https://github.com/pingcap/tidb/pull/23401)                   |   | 7 | 1d3h  |


</details> 


<details> 
  <summary>ichn-hu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                            PR                                                             | C | LASTED |
|------------|---------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/21853 | [expression: fix compatibility behaviors in time_format with MySQL (#21559)](https://github.com/pingcap/tidb/pull/21853)  |   | 98d18h |
| tidb/23278 | [executor: wrong result of nullif expr when used with is null expr. (#23170)](https://github.com/pingcap/tidb/pull/23278) |   | 13d18h |
| tidb/23441 | [executor: Refactor probe channel & fix bug in chunks of join](https://github.com/pingcap/tidb/pull/23441)                |   | 6d14h  |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>lzmhhh123's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                             PR                                                                              | C | LASTED  |
|------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20444 | [expression: add json_merge_patch](https://github.com/pingcap/tidb/pull/20444)                                                                              |   | 162d21h |
| tidb/20465 | [expression: add uuidShortFunction](https://github.com/pingcap/tidb/pull/20465)                                                                             |   | 161d19h |
| tidb/20642 | [executor: modify admin executors to support partitioned table with global index](https://github.com/pingcap/tidb/pull/20642)                               |   | 150d15h |
| tidb/20903 | [planner: fix confused and unnecessary double-projection in plans.](https://github.com/pingcap/tidb/pull/20903)                                             |   | 139d17h |
| tidb/21018 | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                              |   | 133d16h |
| tidb/21195 | [brie: integrate lightning to suport IMPORT statement](https://github.com/pingcap/tidb/pull/21195)                                                          |   | 122d22h |
| tidb/21334 | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                                               |   | 119d14h |
| tidb/21347 | [session: make rollback work on global variables](https://github.com/pingcap/tidb/pull/21347)                                                               |   | 118d19h |
| tidb/21444 | [planner: ignore anonymous index while tiflash replica is available](https://github.com/pingcap/tidb/pull/21444)                                            |   | 113d12h |
| tidb/21487 | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                                               |   | 112d4h  |
| tidb/21641 | [executor: Fix pessimistic lock doesn't work on the partition table for subquery/joins](https://github.com/pingcap/tidb/pull/21641)                         |   | 105d17h |
| tidb/21651 | [planner: allow filter condition pushing down to IndexScan for prefix index](https://github.com/pingcap/tidb/pull/21651)                                    |   | 105d13h |
| tidb/22126 | [*: add `sys` schema, `sys.SCHEMA_UNUSED_INDEXES` view and `sys.SCHEMA_INDEX_USAGE` view](https://github.com/pingcap/tidb/pull/22126)                       |   | 84d19h  |
| tidb/22149 | [session: set process info before building plan (#22101)](https://github.com/pingcap/tidb/pull/22149)                                                       |   | 80d19h  |
| tidb/22188 | [planner: do not use indexMerge when the path only use a single index (#22168)](https://github.com/pingcap/tidb/pull/22188)                                 |   | 79d13h  |
| tidb/22361 | [table: fix insert into _tidb_rowid panic and rebase it if needed (#22062)](https://github.com/pingcap/tidb/pull/22361)                                     |   | 72d20h  |
| tidb/22372 | [executor: fix SelectForUpdate in decorrelated subquery under pessimistic mode](https://github.com/pingcap/tidb/pull/22372)                                 |   | 72d9h   |
| tidb/22478 | [planner, executor: fix query partition table with global unique index get wrong result](https://github.com/pingcap/tidb/pull/22478)                        |   | 63d13h  |
| tidb/22631 | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                                                                             |   | 53d22h  |
| tidb/22686 | [expression: support enum pushdown](https://github.com/pingcap/tidb/pull/22686)                                                                             |   | 50d22h  |
| tidb/22699 | [brie: add error info column and history backup/restore info in sql](https://github.com/pingcap/tidb/pull/22699)                                            |   | 50d16h  |
| tidb/22857 | [mocktikv: split rpcHandler to kvHandler and coprHandler](https://github.com/pingcap/tidb/pull/22857)                                                       |   | 32d20h  |
| tidb/22926 | [expression: add overflow check in multiplyInt](https://github.com/pingcap/tidb/pull/22926)                                                                 |   | 29d13h  |
| tidb/23001 | [statistics: fix err check](https://github.com/pingcap/tidb/pull/23001)                                                                                     |   | 26d0h   |
| tidb/23022 | [executor: create PipelinedWindowExec based on current implementation and modify the windowProcessor interface](https://github.com/pingcap/tidb/pull/23022) |   | 24d18h  |
| tidb/23149 | [core: support left join and right join for join reorder](https://github.com/pingcap/tidb/pull/23149)                                                       |   | 19d12h  |
| tidb/23210 | [planner: fixed a bug that prevented SPM from taking effect (#23197)](https://github.com/pingcap/tidb/pull/23210)                                           |   | 16d16h  |
| tidb/23257 | [executor: group_concat aggr panic when session.group_concat_max_len is small (#23131)](https://github.com/pingcap/tidb/pull/23257)                         |   | 14d18h  |
| tidb/23278 | [executor: wrong result of nullif expr when used with is null expr. (#23170)](https://github.com/pingcap/tidb/pull/23278)                                   |   | 13d18h  |
| tidb/23283 | [util: optimize the performance of restore with db (#22910)](https://github.com/pingcap/tidb/pull/23283)                                                    |   | 13d17h  |
| tidb/23293 | [planner: fix the bug that wrong collation is used when try fast path for enum or set (#23217)](https://github.com/pingcap/tidb/pull/23293)                 |   | 13d14h  |
| tidb/23296 | [sig/execution: fix the bug that Wrong result of comparison operation(type date / type string)](https://github.com/pingcap/tidb/pull/23296)                 |   | 13d6h   |
| tidb/23307 | [util/chunk: replace outdated link with correct one](https://github.com/pingcap/tidb/pull/23307)                                                            |   | 10d20h  |
| tidb/23334 | [metrics/grafana: Remove duplicate items "Owner Watcher OPS"](https://github.com/pingcap/tidb/pull/23334)                                                   |   | 9d19h   |
| tidb/23337 | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23337)                                    |   | 9d19h   |
| tidb/23340 | [executor: fix unexpected NotNullFlag in case when expr ret type (#23102)](https://github.com/pingcap/tidb/pull/23340)                                      |   | 9d19h   |
| tidb/23347 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                                        |   | 9d17h   |
| tidb/23348 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                                        |   | 9d17h   |
| tidb/23368 | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23368)                                       |   | 8d20h   |
| tidb/23369 | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23369)                                       |   | 8d20h   |
| tidb/23370 | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23370)                                       |   | 8d20h   |
| tidb/23373 | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                           |   | 8d19h   |
| tidb/23374 | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23374)                                           |   | 8d19h   |
| tidb/23399 | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23399)                                                              |   | 7d17h   |
| tidb/23417 | [plan: reset not null flag](https://github.com/pingcap/tidb/pull/23417)                                                                                     |   | 7d11h   |
| tidb/23422 | [sig/execution: fix the bug that The result of 'varbinary + 1' is incorrect](https://github.com/pingcap/tidb/pull/23422)                                    |   | 7d6h    |
| tidb/23441 | [executor: Refactor probe channel & fix bug in chunks of join](https://github.com/pingcap/tidb/pull/23441)                                                  |   | 6d14h   |
| tidb/23447 | [*: Support record stmtsummary evicted count](https://github.com/pingcap/tidb/pull/23447)                                                                   |   | 4d18h   |
| tidb/23461 | [execdetails: refine cop task execution stats display in plan](https://github.com/pingcap/tidb/pull/23461)                                                  |   | 3d12h   |
| tidb/23469 | [*: hide `index-usage-sync-lease` config (#23349)](https://github.com/pingcap/tidb/pull/23469)                                                              |   | 2d21h   |
| tidb/23480 | [planner/core: inject project for tiflash agg](https://github.com/pingcap/tidb/pull/23480)                                                                  |   | 2d16h   |
| tidb/23493 | [expression: Implementation of Vitess hashing algorithm.](https://github.com/pingcap/tidb/pull/23493)                                                       |   | 2d6h    |
| tidb/23523 | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral](https://github.com/pingcap/tidb/pull/23523)                             |   | 22h     |
| tidb/23536 | [planner: remove some risky cache operations in the plan builder (#23354)](https://github.com/pingcap/tidb/pull/23536)                                      |   | 19h     |
| tidb/23557 | [MPP: fix agg bug during planning](https://github.com/pingcap/tidb/pull/23557)                                                                              |   | 14h     |
| tidb/23559 | [util: fix wrong year type](https://github.com/pingcap/tidb/pull/23559)                                                                                     |   | 14h     |


## Reviewed in Last 7 Days

|    REPO    |                                                           PR                                                           | C | D |   R   |
|------------|------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/23489 | [telemetry: add copr-cache, tiflash, cluster index, async commit (#23454)](https://github.com/pingcap/tidb/pull/23489) |   | 1 | 1d20h |
| tidb/23518 | [telemetry: Fix window algorithm](https://github.com/pingcap/tidb/pull/23518)                                          |   | 1 | 8h    |
| tidb/23470 | [telemetry: add transaction usage info](https://github.com/pingcap/tidb/pull/23470)                                    |   | 2 | 21h   |
| tikv/9860  | [copr: cast invalid utf8 string to real bug](https://github.com/tikv/tikv/pull/9860)                                   | Y | 3 | 5h    |
| tikv/9850  | [copr: fix IN expr didn't handle unsigned/signed int properly (#9823)](https://github.com/tikv/tikv/pull/9850)         |   | 7 | 1h    |


</details> 


<details> 
  <summary>nullnotnil's detailed review report</summary> 

## To Be Reviewed

| REPO | PR | C | LASTED |
|------|----|---|--------|


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>qw4990's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                             PR                                                                              | C | LASTED  |
|--------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| docs-cn/5561 | [Add sql optimization-related docs to toc](https://github.com/pingcap/docs-cn/pull/5561)                                                                    |   | 31d15h  |
| docs/5090    | [update tidb v4.0.12 release notes](https://github.com/pingcap/docs/pull/5090)                                                                              |   | 2d14h   |
| docs-cn/5785 | [update SPM documentation for DML SQL Bind and baseline capture (#5740)](https://github.com/pingcap/docs-cn/pull/5785)                                      |   | 3d18h   |
| tidb/19029   | [types: fix unexpected NOT_NULL flags](https://github.com/pingcap/tidb/pull/19029)                                                                          |   | 231d22h |
| docs-cn/5806 | [releases: add tidb release notes 4.0.12](https://github.com/pingcap/docs-cn/pull/5806)                                                                     |   | 2d14h   |
| tidb/20708   | [*: separate auto_increment ID allocator from _tidb_rowid allocator](https://github.com/pingcap/tidb/pull/20708)                                            |   | 147d20h |
| tidb/20969   | [executor: Improve the performance of appending not fixed columns](https://github.com/pingcap/tidb/pull/20969)                                              |   | 135d9h  |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                              |   | 133d16h |
| tidb/21149   | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                               |   | 126d14h |
| tidb/21304   | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                                    |   | 120d12h |
| tidb/21318   | [planner, expression: use the range of column types to simplify expressions](https://github.com/pingcap/tidb/pull/21318)                                    |   | 119d18h |
| tidb/21401   | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                                          |   | 115d11h |
| tidb/21476   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                       |   | 112d15h |
| tidb/21508   | [execution: fix dayofweek('0000-00-00') behavior](https://github.com/pingcap/tidb/pull/21508)                                                               |   | 111d9h  |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                                 |   | 97d19h  |
| tidb/21887   | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                                             |   | 96d11h  |
| tidb/21954   | [planner/cascades: add rule `PushSelDownApply`](https://github.com/pingcap/tidb/pull/21954)                                                                 |   | 92d23h  |
| tidb/22146   | [executor: forbid SFU on view](https://github.com/pingcap/tidb/pull/22146)                                                                                  |   | 80d21h  |
| tidb/22217   | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                                               |   | 78d17h  |
| tidb/22234   | [executor, planner: ON DUPLICATE UPDATE can refer to un-project col (#14412)](https://github.com/pingcap/tidb/pull/22234)                                   |   | 78d15h  |
| tidb/22261   | [time: fix parse datetime won't truncate the reluctant string (#22232)](https://github.com/pingcap/tidb/pull/22261)                                         |   | 77d19h  |
| tidb/22294   | [planner, table: optimize the list column partition pruner for range query](https://github.com/pingcap/tidb/pull/22294)                                     |   | 76d19h  |
| tidb/22307   | [ddl: fix update can see columns not public](https://github.com/pingcap/tidb/pull/22307)                                                                    |   | 76d15h  |
| tidb/22342   | [session: fix two cases when updating bind info (#22338)](https://github.com/pingcap/tidb/pull/22342)                                                       |   | 73d18h  |
| tidb/22369   | [session: fix the duplicate binding case when updating bind info (#22367)](https://github.com/pingcap/tidb/pull/22369)                                      |   | 72d17h  |
| tidb/22374   | [expression: separated arithmeticIntDivideSig](https://github.com/pingcap/tidb/pull/22374)                                                                  |   | 72d0h   |
| tidb/22415   | [ddl: refactor placement package](https://github.com/pingcap/tidb/pull/22415)                                                                               |   | 68d17h  |
| tidb/22541   | [expression: Support builtin function SOUNDEX](https://github.com/pingcap/tidb/pull/22541)                                                                  |   | 58d9h   |
| tidb/22559   | [planner: split test data from test cases in cbo_test.go](https://github.com/pingcap/tidb/pull/22559)                                                       |   | 57d19h  |
| tidb/22565   | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)                                          | Y | 57d17h  |
| tidb/22814   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22814)                                         |   | 34d18h  |
| tidb/22815   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22815)                                         |   | 34d18h  |
| tidb/22862   | [brie: fix the problem that ddl restored by BR via SQL is not replicated to downstream](https://github.com/pingcap/tidb/pull/22862)                         |   | 31d22h  |
| tidb/22915   | [planner: build correct MaxOneRow info from multi-column conditions](https://github.com/pingcap/tidb/pull/22915)                                            |   | 29d17h  |
| tidb/22923   | [expression: correct constant propagation for collation (#22666)](https://github.com/pingcap/tidb/pull/22923)                                               |   | 29d14h  |
| tidb/22924   | [planner: fix wrong index merge selection (#22825)](https://github.com/pingcap/tidb/pull/22924)                                                             |   | 29d13h  |
| tidb/22926   | [expression: add overflow check in multiplyInt](https://github.com/pingcap/tidb/pull/22926)                                                                 |   | 29d13h  |
| tidb/22984   | [executor: fix logging format of prepared statements (#16062)](https://github.com/pingcap/tidb/pull/22984)                                                  |   | 26d10h  |
| tidb/23022   | [executor: create PipelinedWindowExec based on current implementation and modify the windowProcessor interface](https://github.com/pingcap/tidb/pull/23022) |   | 24d18h  |
| tidb/23105   | [executor: fix wrong key range of index scan when filter is comparing year column with NULL (#23079)](https://github.com/pingcap/tidb/pull/23105)           |   | 21d18h  |
| tidb/23137   | [planner: fix index merge row count estimation logic](https://github.com/pingcap/tidb/pull/23137)                                                           |   | 20d17h  |
| tidb/23152   | [expression: fix wrong error info (#22760)](https://github.com/pingcap/tidb/pull/23152)                                                                     |   | 18d14h  |
| tidb/23196   | [types: fix the bug about the wrong query result for decimal type  (#22507)](https://github.com/pingcap/tidb/pull/23196)                                    |   | 16d18h  |
| tidb/23201   | [executor, server: load_data.go is changed and add unit test](https://github.com/pingcap/tidb/pull/23201)                                                   |   | 16d17h  |
| tidb/23208   | [statistics, util/ranger: improve selectivity calculation for DNF filters (#18741)](https://github.com/pingcap/tidb/pull/23208)                             |   | 16d16h  |
| tidb/23210   | [planner: fixed a bug that prevented SPM from taking effect (#23197)](https://github.com/pingcap/tidb/pull/23210)                                           |   | 16d16h  |
| tidb/23234   | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23234)                                                      |   | 15d18h  |
| tidb/23284   | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified.](https://github.com/pingcap/tidb/pull/23284)              |   | 13d17h  |
| tidb/23293   | [planner: fix the bug that wrong collation is used when try fast path for enum or set (#23217)](https://github.com/pingcap/tidb/pull/23293)                 |   | 13d14h  |
| tidb/23295   | [util, types: don't let SPM be affected by charset (#23161)](https://github.com/pingcap/tidb/pull/23295)                                                    |   | 13d11h  |
| tidb/23316   | [planner: Fix rebuild range for prepared plan](https://github.com/pingcap/tidb/pull/23316)                                                                  |   | 10d17h  |
| tidb/23373   | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                           |   | 8d19h   |
| tidb/23397   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23397)                                                              |   | 7d17h   |
| tidb/23398   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23398)                                                              |   | 7d17h   |
| tidb/23399   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23399)                                                              |   | 7d17h   |
| tidb/23448   | [wip :execution: parallel build hash table](https://github.com/pingcap/tidb/pull/23448)                                                                     |   | 4d11h   |
| tidb/23462   | [*: collect transaction write duration/throughput metrics for SLI/SLO](https://github.com/pingcap/tidb/pull/23462)                                          |   | 3d12h   |
| tidb/23493   | [expression: Implementation of Vitess hashing algorithm.](https://github.com/pingcap/tidb/pull/23493)                                                       |   | 2d6h    |
| tidb/23509   | [mpp: fix correlated columns in filter or access in MPP apply](https://github.com/pingcap/tidb/pull/23509)                                                  |   | 1d17h   |
| tidb/23523   | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral](https://github.com/pingcap/tidb/pull/23523)                             |   | 22h     |
| tidb/23533   | [executor: fix batchGet/PointGet overflow result in clustered index (#23516)](https://github.com/pingcap/tidb/pull/23533)                                   |   | 19h     |
| tidb/23543   | [statistics: fix auto analyze log information incomplete (#23522)](https://github.com/pingcap/tidb/pull/23543)                                              |   | 18h     |
| tidb/23550   | [ranger: handle decimal overflow properly when building index ranges (#23535)](https://github.com/pingcap/tidb/pull/23550)                                  |   | 17h     |
| tidb/23557   | [MPP: fix agg bug during planning](https://github.com/pingcap/tidb/pull/23557)                                                                              |   | 14h     |
| tidb/23565   | [planner/core: convert decimal type for mpp join before shuffling. (#23191)](https://github.com/pingcap/tidb/pull/23565)                                    |   | 9h      |


## Reviewed in Last 7 Days

|     REPO     |                                                                                       PR                                                                                        | C | D |   R    |
|--------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/23492   | [planner, sessionctx: turn on the mpp by default (#23401)](https://github.com/pingcap/tidb/pull/23492)                                                                          |   | 1 | 2d1h   |
| tidb/23191   | [planner/core: convert decimal type for mpp join before shuffling.](https://github.com/pingcap/tidb/pull/23191)                                                                 |   | 1 | 16d1h  |
| tidb/23544   | [statistics: fix auto analyze log information incomplete (#23522)](https://github.com/pingcap/tidb/pull/23544)                                                                  |   | 1 | 0h     |
| docs/5092    | [SQL variables: add a new variable `tidb_enable_parallel_apply`](https://github.com/pingcap/docs/pull/5092)                                                                     |   | 1 | 1d1h   |
| tidb/23437   | [statistics: check duplicate columns and types for extended stats](https://github.com/pingcap/tidb/pull/23437)                                                                  |   | 1 | 5d21h  |
| tidb/23522   | [statistics: fix auto analyze log information incomplete](https://github.com/pingcap/tidb/pull/23522)                                                                           |   | 1 | 3h     |
| tidb/23502   | [statistics: fix some unstable tests in global stats](https://github.com/pingcap/tidb/pull/23502)                                                                               |   | 2 | 1h     |
| tidb/23283   | [util: optimize the performance of restore with db (#22910)](https://github.com/pingcap/tidb/pull/23283)                                                                        |   | 3 | 10d23h |
| tidb/23471   | [planner: fix only_full_group_by check not enough (#23404)](https://github.com/pingcap/tidb/pull/23471)                                                                         |   | 3 | 0h     |
| tidb/23404   | [planner: fix only_full_group_by check not enough](https://github.com/pingcap/tidb/pull/23404)                                                                                  |   | 3 | 4d18h  |
| tidb/23458   | [planner: incorrect results returned when executing a prepared point-get update meets nondeterministic function. (#22067) (#23455)](https://github.com/pingcap/tidb/pull/23458) |   | 4 | 0h     |
| tidb/23455   | [planner: incorrect results returned when executing a prepared point-get update meets nondeterministic function. (#22067)](https://github.com/pingcap/tidb/pull/23455)          |   | 4 | 1h     |
| tidb/23343   | [statistics: handle drop partition events for global-stats](https://github.com/pingcap/tidb/pull/23343)                                                                         |   | 7 | 3d1h   |
| docs-cn/5484 | [system variable: add tidb_allow_fallback_to_tikv](https://github.com/pingcap/docs-cn/pull/5484)                                                                                |   | 7 | 44d0h  |
| tidb/23401   | [planner, sessionctx: turn on the mpp by default](https://github.com/pingcap/tidb/pull/23401)                                                                                   |   | 7 | 23h    |


</details> 


<details> 
  <summary>rebelice's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                           PR                                                           | C | LASTED |
|------------|------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/23536 | [planner: remove some risky cache operations in the plan builder (#23354)](https://github.com/pingcap/tidb/pull/23536) |   | 19h    |
| tidb/23537 | [planner: remove some risky cache operations in the plan builder (#23354)](https://github.com/pingcap/tidb/pull/23537) |   | 19h    |


## Reviewed in Last 7 Days

|    REPO    |                                                         PR                                                          | C | D |   R   |
|------------|---------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/23520 | [statistics: optimize the global histogram merging algorithm](https://github.com/pingcap/tidb/pull/23520)           |   | 1 | 2h    |
| tidb/23502 | [statistics: fix some unstable tests in global stats](https://github.com/pingcap/tidb/pull/23502)                   |   | 2 | 0h    |
| tidb/23486 | [statistics: recalculate bucket repeats when merging global statistics](https://github.com/pingcap/tidb/pull/23486) |   | 2 | 15h   |
| tidb/23343 | [statistics: handle drop partition events for global-stats](https://github.com/pingcap/tidb/pull/23343)             |   | 4 | 5d22h |


</details> 


<details> 
  <summary>time-and-fate's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                             PR                                                             | C | LASTED  |
|------------|----------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                  |   | 140d16h |
| tidb/22853 | [planner: fix LogicalPlans that contain Window Function are ambiguous ](https://github.com/pingcap/tidb/pull/22853)        |   | 33d12h  |
| tidb/22915 | [planner: build correct MaxOneRow info from multi-column conditions](https://github.com/pingcap/tidb/pull/22915)           |   | 29d17h  |
| tidb/23550 | [ranger: handle decimal overflow properly when building index ranges (#23535)](https://github.com/pingcap/tidb/pull/23550) |   | 17h     |


## Reviewed in Last 7 Days

|    REPO    |                                                             PR                                                             | C | D |   R   |
|------------|----------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/23551 | [ranger: handle decimal overflow properly when building index ranges (#23535)](https://github.com/pingcap/tidb/pull/23551) |   | 1 | 0h    |
| tidb/23413 | [statistics: remove the dependency of stats GC on LastUpdateVersion](https://github.com/pingcap/tidb/pull/23413)           |   | 1 | 6d20h |
| tidb/23535 | [ranger: handle decimal overflow properly when building index ranges](https://github.com/pingcap/tidb/pull/23535)          |   | 1 | 1h    |


</details> 


<details> 
  <summary>winoros's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                              PR                                                                               | C | LASTED  |
|------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19957 | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                                        | Y | 195d13h |
| tidb/20311 | [expression: fix overflow error when convert bit to int64 (#20266)](https://github.com/pingcap/tidb/pull/20311)                                               |   | 176d21h |
| tidb/20765 | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                                             |   | 143d16h |
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                                                     |   | 140d16h |
| tidb/21018 | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                                |   | 133d16h |
| tidb/21207 | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                                               |   | 122d19h |
| tidb/21476 | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                         |   | 112d15h |
| tidb/21487 | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                                                 |   | 112d4h  |
| tidb/21876 | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                                   |   | 97d19h  |
| tidb/21954 | [planner/cascades: add rule `PushSelDownApply`](https://github.com/pingcap/tidb/pull/21954)                                                                   |   | 92d23h  |
| tidb/22181 | [planner, expression: fix error when using IN combined with subquery (#22080)](https://github.com/pingcap/tidb/pull/22181)                                    |   | 79d17h  |
| tidb/22365 | [planner: check index valid while forUpdateRead (#22152)](https://github.com/pingcap/tidb/pull/22365)                                                         |   | 72d19h  |
| tidb/22504 | [*:Fix the fetchHotRegion bug that the count always zero](https://github.com/pingcap/tidb/pull/22504)                                                         |   | 60d19h  |
| tidb/22565 | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)                                            | Y | 57d17h  |
| tidb/22624 | [ planner: not pruning column used by union scan condition (#21640)](https://github.com/pingcap/tidb/pull/22624)                                              |   | 55d16h  |
| tidb/22923 | [expression: correct constant propagation for collation (#22666)](https://github.com/pingcap/tidb/pull/22923)                                                 |   | 29d14h  |
| tidb/23163 | [plugin: fix linter --enable=deadcode check error](https://github.com/pingcap/tidb/pull/23163)                                                                |   | 17d18h  |
| tidb/23208 | [statistics, util/ranger: improve selectivity calculation for DNF filters (#18741)](https://github.com/pingcap/tidb/pull/23208)                               |   | 16d16h  |
| tidb/23215 | [Privileges: fix delete privilege check wrongly (#22971)](https://github.com/pingcap/tidb/pull/23215)                                                         |   | 16d14h  |
| tidb/23216 | [Privileges: fix delete privilege check wrongly (#22971)](https://github.com/pingcap/tidb/pull/23216)                                                         |   | 16d14h  |
| tidb/23233 | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23233)                                                        |   | 15d18h  |
| tidb/23234 | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23234)                                                        |   | 15d18h  |
| tidb/23347 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                                          |   | 9d17h   |
| tidb/23348 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                                          |   | 9d17h   |
| tidb/23350 | [util/stringutil, util/ranger, planner: use hierarchical separators to simplify the parsing for info of EXPLAIN ](https://github.com/pingcap/tidb/pull/23350) |   | 9d17h   |
| tidb/23365 | [planner: fix a bug that point get plan returns wrong column name](https://github.com/pingcap/tidb/pull/23365)                                                |   | 8d21h   |
| tidb/23373 | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                             |   | 8d19h   |
| tidb/23435 | [planner: check schema stale for plan cache when forUpdateRead (#22381)](https://github.com/pingcap/tidb/pull/23435)                                          |   | 6d16h   |
| tidb/23469 | [*: hide `index-usage-sync-lease` config (#23349)](https://github.com/pingcap/tidb/pull/23469)                                                                |   | 2d21h   |
| tidb/23474 | [planner: fix inappropriate null flag of null constants (#23457)](https://github.com/pingcap/tidb/pull/23474)                                                 |   | 2d18h   |
| tidb/23537 | [planner: remove some risky cache operations in the plan builder (#23354)](https://github.com/pingcap/tidb/pull/23537)                                        |   | 19h     |
| tidb/23543 | [statistics: fix auto analyze log information incomplete (#23522)](https://github.com/pingcap/tidb/pull/23543)                                                |   | 18h     |
| tidb/23550 | [ranger: handle decimal overflow properly when building index ranges (#23535)](https://github.com/pingcap/tidb/pull/23550)                                    |   | 17h     |
| tidb/23564 | [statistics: sort the test result, stable tests](https://github.com/pingcap/tidb/pull/23564)                                                                  |   | 12h     |


## Reviewed in Last 7 Days

|     REPO     |                                                                PR                                                                | C | D |    R    |
|--------------|----------------------------------------------------------------------------------------------------------------------------------|---|---|---------|
| tidb/23413   | [statistics: remove the dependency of stats GC on LastUpdateVersion](https://github.com/pingcap/tidb/pull/23413)                 |   | 1 | 6d22h   |
| tidb/23551   | [ranger: handle decimal overflow properly when building index ranges (#23535)](https://github.com/pingcap/tidb/pull/23551)       |   | 1 | 0h      |
| tidb/23535   | [ranger: handle decimal overflow properly when building index ranges](https://github.com/pingcap/tidb/pull/23535)                |   | 1 | 1h      |
| tidb/23523   | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral](https://github.com/pingcap/tidb/pull/23523)  |   | 1 | 4h      |
| tidb/23437   | [statistics: check duplicate columns and types for extended stats](https://github.com/pingcap/tidb/pull/23437)                   |   | 1 | 5d21h   |
| tidb/23536   | [planner: remove some risky cache operations in the plan builder (#23354)](https://github.com/pingcap/tidb/pull/23536)           |   | 1 | 0h      |
| tidb/23524   | [partition: fix hash partition with not between condition get wrong result (#22914)](https://github.com/pingcap/tidb/pull/23524) |   | 1 | 2h      |
| tidb/23525   | [partition: fix hash partition with not between condition get wrong result (#22914)](https://github.com/pingcap/tidb/pull/23525) |   | 1 | 2h      |
| tidb/23487   | [planner: optimize count(distinct a) to count(a) if there is an unique key on a](https://github.com/pingcap/tidb/pull/23487)     | Y | 2 | 23h     |
| tidb/23503   | [store/copr: return ErrTiFlashServerTimeout when MPP meets network error (#23434)](https://github.com/pingcap/tidb/pull/23503)   |   | 2 | 4h      |
| tidb/23481   | [planner: fix inappropriate null flag of null constants (#23457)](https://github.com/pingcap/tidb/pull/23481)                    |   | 2 | 1d1h    |
| tidb/23467   | [hot-fix: paginate indexLookUp executor](https://github.com/pingcap/tidb/pull/23467)                                             |   | 2 | 1d5h    |
| tidb/23246   | [planner: fix the panic in joinReOrderSolver.optimizeRecursive](https://github.com/pingcap/tidb/pull/23246)                      |   | 2 | 13d5h   |
| tidb/23434   | [store/copr: return ErrTiFlashServerTimeout when MPP meets network error](https://github.com/pingcap/tidb/pull/23434)            |   | 2 | 4d20h   |
| tidb/23471   | [planner: fix only_full_group_by check not enough (#23404)](https://github.com/pingcap/tidb/pull/23471)                          |   | 3 | 4h      |
| tidb/23457   | [planner: fix inappropriate null flag of null constants](https://github.com/pingcap/tidb/pull/23457)                             |   | 3 | 19h     |
| tidb/18247   | [docs/design: proposal for non-recursive common table expression (CTE)](https://github.com/pingcap/tidb/pull/18247)              | Y | 3 | 267d10h |
| tidb/23175   | [planner: fix plan cache not working caused by type difference](https://github.com/pingcap/tidb/pull/23175)                      |   | 4 | 14d0h   |
| tidb/23349   | [*: hide `index-usage-sync-lease` config](https://github.com/pingcap/tidb/pull/23349)                                            |   | 4 | 6d1h    |
| tidb/23404   | [planner: fix only_full_group_by check not enough](https://github.com/pingcap/tidb/pull/23404)                                   |   | 4 | 3d21h   |
| tidb/22381   | [planner: check schema stale for plan cache when forUpdateRead](https://github.com/pingcap/tidb/pull/22381)                      |   | 7 | 64d21h  |
| docs-cn/5484 | [system variable: add tidb_allow_fallback_to_tikv](https://github.com/pingcap/docs-cn/pull/5484)                                 |   | 7 | 43d23h  |


</details> 


<details> 
  <summary>wshwsh12's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                        PR                                                                         | C | LASTED  |
|------------|---------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19807 | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                                           |   | 202d10h |
| tidb/19957 | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                            | Y | 195d13h |
| tidb/21487 | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                                     |   | 112d4h  |
| tidb/21887 | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                                   |   | 96d11h  |
| tidb/22378 | [executor: vectorize hash aggregate](https://github.com/pingcap/tidb/pull/22378)                                                                  |   | 71d19h  |
| tidb/22628 | [executor: Improve max/min window function with deque-based sliding window](https://github.com/pingcap/tidb/pull/22628)                           |   | 54d23h  |
| tidb/22815 | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22815)                               |   | 34d18h  |
| tidb/23105 | [executor: fix wrong key range of index scan when filter is comparing year column with NULL (#23079)](https://github.com/pingcap/tidb/pull/23105) |   | 21d18h  |
| tidb/23128 | [statistics: refactor the statistics package use the RestrictedSQLExecutor API (#22636)](https://github.com/pingcap/tidb/pull/23128)              |   | 20d22h  |
| tidb/23234 | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23234)                                            |   | 15d18h  |
| tidb/23336 | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23336)                          |   | 9d19h   |
| tidb/23347 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                              |   | 9d17h   |
| tidb/23348 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                              |   | 9d17h   |
| tidb/23368 | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23368)                             |   | 8d20h   |
| tidb/23369 | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23369)                             |   | 8d20h   |
| tidb/23370 | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23370)                             |   | 8d20h   |
| tidb/23397 | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23397)                                                    |   | 7d17h   |
| tidb/23398 | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23398)                                                    |   | 7d17h   |
| tidb/23399 | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23399)                                                    |   | 7d17h   |
| tidb/23435 | [planner: check schema stale for plan cache when forUpdateRead (#22381)](https://github.com/pingcap/tidb/pull/23435)                              |   | 6d16h   |
| tidb/23445 | [chore: Refactor code quality issues](https://github.com/pingcap/tidb/pull/23445)                                                                 |   | 5d20h   |
| tidb/23469 | [*: hide `index-usage-sync-lease` config (#23349)](https://github.com/pingcap/tidb/pull/23469)                                                    |   | 2d21h   |
| tidb/23519 | [executor: check privilege before adding](https://github.com/pingcap/tidb/pull/23519)                                                             |   | 23h     |
| tidb/23554 | [executor: fix update panic on join having statement](https://github.com/pingcap/tidb/pull/23554)                                                 |   | 16h     |
| tidb/23565 | [planner/core: convert decimal type for mpp join before shuffling. (#23191)](https://github.com/pingcap/tidb/pull/23565)                          |   | 9h      |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 

