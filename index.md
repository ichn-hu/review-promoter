|   REVIEWER    |  O(C)   | 1 | 2 | 3 | 4 | 5  | 6 | 7 | T  |
|---------------|---------|---|---|---|---|----|---|---|----|
| Reminiscent   |  5 ( 0) | 1 | 0 | 0 | 0 |  4 | 0 | 3 |  8 |
| SunRunAway    | 16 ( 1) | 0 | 0 | 0 | 0 |  2 | 0 | 0 |  2 |
| XuHuaiyu      | 76 ( 2) | 0 | 0 | 0 | 0 |  2 | 1 | 0 |  3 |
| dyzsr         |  0 ( 0) | 0 | 0 | 0 | 0 |  0 | 0 | 0 |  0 |
| eurekaka      | 21 ( 1) | 1 | 0 | 0 | 0 |  0 | 0 | 1 |  2 |
| fzhedu        |  1 ( 0) | 0 | 0 | 0 | 2 |  0 | 0 | 0 |  2 |
| ichn-hu       |  2 ( 0) | 0 | 0 | 0 | 0 |  0 | 0 | 0 |  0 |
| lzmhhh123     | 51 ( 0) | 0 | 0 | 0 | 1 | 12 | 1 | 3 | 17 |
| nullnotnil    |  0 ( 0) | 0 | 0 | 0 | 0 |  0 | 0 | 0 |  0 |
| qw4990        | 58 ( 1) | 2 | 0 | 0 | 4 |  5 | 2 | 2 | 15 |
| rebelice      |  0 ( 0) | 1 | 0 | 0 | 0 |  0 | 1 | 0 |  2 |
| time-and-fate |  3 ( 0) | 0 | 0 | 0 | 0 |  0 | 2 | 0 |  2 |
| winoros       | 30 ( 2) | 3 | 0 | 0 | 3 |  4 | 1 | 1 | 12 |
| wshwsh12      | 23 ( 1) | 0 | 0 | 0 | 0 |  1 | 0 | 1 |  2 |


<details> 
  <summary>Reminiscent's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                     PR                                                                      | C | LASTED |
|------------|---------------------------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/21896 | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                                  |   | 91d19h |
| tidb/22354 | [planner: do not cache prepared plan if optimization depends on mutable constant (#22349)](https://github.com/pingcap/tidb/pull/22354)      |   | 69d23h |
| tidb/23283 | [util: optimize the performance of restore with db (#22910)](https://github.com/pingcap/tidb/pull/23283)                                    |   | 10d17h |
| tidb/23293 | [planner: fix the bug that wrong collation is used when try fast path for enum or set (#23217)](https://github.com/pingcap/tidb/pull/23293) |   | 10d14h |
| tidb/23413 | [statistics: remove the dependency of stats GC on LastUpdateVersion](https://github.com/pingcap/tidb/pull/23413)                            |   | 4d14h  |


## Reviewed in Last 7 Days

|    REPO    |                                                              PR                                                               | C | D |   R    |
|------------|-------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/23457 | [planner: fix inappropriate null flag of null constants](https://github.com/pingcap/tidb/pull/23457)                          |   | 1 | 0h     |
| tidb/23346 | [statistics: fix build extended stats panic when met NULL](https://github.com/pingcap/tidb/pull/23346)                        |   | 5 | 2d18h  |
| tidb/23119 | [statistics: remove existing deleted extended stats when add a new one](https://github.com/pingcap/tidb/pull/23119)           |   | 5 | 14d14h |
| tidb/23413 | [statistics: remove the dependency of stats GC on LastUpdateVersion](https://github.com/pingcap/tidb/pull/23413)              |   | 5 | 14h    |
| tidb/23392 | [sessionctx: hide extended stats variable in SHOW VARIABLES temporarily (#23345)](https://github.com/pingcap/tidb/pull/23392) |   | 5 | 0h     |
| tidb/23333 | [statistics: hide the `tidb_analyze_version` before it's GA (#23327)](https://github.com/pingcap/tidb/pull/23333)             |   | 7 | 1h     |
| tidb/23328 | [planner: hide the switch of dynamic-pruning and global-stats (#23320)](https://github.com/pingcap/tidb/pull/23328)           |   | 7 | 3h     |
| tidb/23327 | [statistics: hide the `tidb_analyze_version` before it's GA](https://github.com/pingcap/tidb/pull/23327)                      |   | 7 | 6h     |


</details> 


<details> 
  <summary>SunRunAway's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                  PR                                                                   | C | LASTED  |
|------------|---------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19178 | [executor: Refactor probe channel](https://github.com/pingcap/tidb/pull/19178)                                                        |   | 221d17h |
| tidb/19347 | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347)  |   | 213d23h |
| tidb/19807 | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                               |   | 199d11h |
| tidb/19900 | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                        | Y | 194d18h |
| tidb/20140 | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                    |   | 181d22h |
| tidb/20765 | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                     |   | 140d17h |
| tidb/21207 | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                       |   | 119d19h |
| tidb/21834 | [planner: enhanced index range calculation plan](https://github.com/pingcap/tidb/pull/21834)                                          |   | 96d18h  |
| tidb/21876 | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)           |   | 94d19h  |
| tidb/21878 | [planner: do not push down lock to pointGet/bacthPointGet when selection exists](https://github.com/pingcap/tidb/pull/21878)          |   | 94d18h  |
| tidb/21956 | [planner/preprocessor: disallow into-outfile clause in some place](https://github.com/pingcap/tidb/pull/21956)                        |   | 89d23h  |
| tidb/22026 | [expression: separated arithmeticPlusIntSig](https://github.com/pingcap/tidb/pull/22026)                                              |   | 87d20h  |
| tidb/22114 | [test: fix globalkilltest (#21987)](https://github.com/pingcap/tidb/pull/22114)                                                       |   | 82d12h  |
| tidb/22217 | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                         |   | 75d17h  |
| tidb/22365 | [planner: check index valid while forUpdateRead (#22152)](https://github.com/pingcap/tidb/pull/22365)                                 |   | 69d19h  |
| tidb/22379 | [[experiment] executor: allow aggregation to spill disk when running out of memory quota](https://github.com/pingcap/tidb/pull/22379) |   | 68d19h  |


## Reviewed in Last 7 Days

|    REPO    |                                             PR                                              | C | D |   R   |
|------------|---------------------------------------------------------------------------------------------|---|---|-------|
| tidb/23224 | [docs: Add Proposal for dynamic privileges](https://github.com/pingcap/tidb/pull/23224)     |   | 5 | 8d14h |
| tidb/23223 | [docs: add proposal for Security Enhanced Mode](https://github.com/pingcap/tidb/pull/23223) |   | 5 | 8d14h |


</details> 


<details> 
  <summary>XuHuaiyu's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                              PR                                                                               | C | LASTED  |
|--------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| docs-cn/5619 | [Update data-type-date-and-time.md](https://github.com/pingcap/docs-cn/pull/5619)                                                                             |   | 24d16h  |
| docs/5075    | [hide the config global-kill](https://github.com/pingcap/docs/pull/5075)                                                                                      |   | 17h     |
| docs-cn/5671 | [tidb: Add time format description](https://github.com/pingcap/docs-cn/pull/5671)                                                                             |   | 18d11h  |
| docs-cn/5754 | [add document for explain/explain analyze of mpp query.](https://github.com/pingcap/docs-cn/pull/5754)                                                        |   | 5d21h   |
| docs-cn/5778 | [docs: remove some config fields for copr cache (#5699)](https://github.com/pingcap/docs-cn/pull/5778)                                                        |   | 3d13h   |
| tidb/19900   | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                                                | Y | 194d18h |
| docs-cn/5779 | [remove system variable `tidb_track_aggregate_memory_usage` (#5704)](https://github.com/pingcap/docs-cn/pull/5779)                                            |   | 3d13h   |
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                                        | Y | 192d14h |
| tidb/20140   | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                                            |   | 181d22h |
| tidb/20311   | [expression: fix overflow error when convert bit to int64 (#20266)](https://github.com/pingcap/tidb/pull/20311)                                               |   | 173d21h |
| tidb/20790   | [collation: add pinyin collation for chinese charset support](https://github.com/pingcap/tidb/pull/20790)                                                     |   | 139d21h |
| tidb/21064   | [planner, executor: fix cast not check error](https://github.com/pingcap/tidb/pull/21064)                                                                     |   | 127d8h  |
| tidb/21149   | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                                 |   | 123d14h |
| tidb/21228   | [executor: return the result immediately when combining LIMIT row_count with DISTINCT](https://github.com/pingcap/tidb/pull/21228)                            |   | 119d13h |
| tidb/21304   | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                                      |   | 117d12h |
| tidb/21334   | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                                                 |   | 116d14h |
| tidb/21401   | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                                            |   | 112d11h |
| tidb/21476   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                         |   | 109d15h |
| tidb/21536   | [executor: add slow-log file meta cache to avoid repeat read file meta information](https://github.com/pingcap/tidb/pull/21536)                               |   | 105d15h |
| tidb/21564   | [ddl: fix Incorrect behavior of NO_ZERO_DATE when altering table](https://github.com/pingcap/tidb/pull/21564)                                                 |   | 104d15h |
| tidb/21853   | [expression: fix compatibility behaviors in time_format with MySQL (#21559)](https://github.com/pingcap/tidb/pull/21853)                                      |   | 95d19h  |
| tidb/21896   | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                                                    |   | 91d19h  |
| tidb/22131   | [privilege: remove leading and trailing space when create user and role](https://github.com/pingcap/tidb/pull/22131)                                          |   | 81d19h  |
| tidb/22149   | [session: set process info before building plan (#22101)](https://github.com/pingcap/tidb/pull/22149)                                                         |   | 77d19h  |
| tidb/22163   | [expression: separated arithmeticMinusIntSig](https://github.com/pingcap/tidb/pull/22163)                                                                     |   | 77d13h  |
| tidb/22186   | [executor: fix select into outfile with year type column has no data (#22175)](https://github.com/pingcap/tidb/pull/22186)                                    |   | 76d16h  |
| tidb/22294   | [planner, table: optimize the list partition pruner for range query](https://github.com/pingcap/tidb/pull/22294)                                              |   | 73d20h  |
| tidb/22307   | [ddl: fix update can see columns not public](https://github.com/pingcap/tidb/pull/22307)                                                                      |   | 73d16h  |
| tidb/22616   | [expression: from_unixtime accept 64-bit integers](https://github.com/pingcap/tidb/pull/22616)                                                                |   | 52d23h  |
| tidb/22617   | [metrics: fix wrong bucket name of coprocessor cache (#22454)](https://github.com/pingcap/tidb/pull/22617)                                                    |   | 52d23h  |
| tidb/22624   | [ planner: not pruning column used by union scan condition (#21640)](https://github.com/pingcap/tidb/pull/22624)                                              |   | 52d17h  |
| tidb/22631   | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                                                                               |   | 50d23h  |
| tidb/22696   | [expression: enable arithmetic Mod push down](https://github.com/pingcap/tidb/pull/22696)                                                                     |   | 47d17h  |
| tidb/22711   | [executor: Fix inline schema name](https://github.com/pingcap/tidb/pull/22711)                                                                                |   | 47d11h  |
| tidb/22722   | [planner, errno: make error code of ErrMixOfGroupFuncAndFields consistent with MySQL](https://github.com/pingcap/tidb/pull/22722)                             |   | 46d20h  |
| tidb/22814   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22814)                                           |   | 31d19h  |
| tidb/22815   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22815)                                           |   | 31d19h  |
| tidb/22908   | [txn: Add txn state's view](https://github.com/pingcap/tidb/pull/22908)                                                                                       |   | 26d20h  |
| tidb/22914   | [partition: fix hash partition with not between condition get wrong result](https://github.com/pingcap/tidb/pull/22914)                                       |   | 26d18h  |
| tidb/22926   | [expression: add overflow check in multiplyInt](https://github.com/pingcap/tidb/pull/22926)                                                                   |   | 26d13h  |
| tidb/23012   | [executor: fix affected rows of ddls and complete uint tests](https://github.com/pingcap/tidb/pull/23012)                                                     |   | 22d17h  |
| tidb/23105   | [executor: fix wrong key range of index scan when filter is comparing year column with NULL (#23079)](https://github.com/pingcap/tidb/pull/23105)             |   | 18d18h  |
| tidb/23128   | [statistics: refactor the statistics package use the RestrictedSQLExecutor API (#22636)](https://github.com/pingcap/tidb/pull/23128)                          |   | 17d22h  |
| tidb/23152   | [expression: fix wrong error info (#22760)](https://github.com/pingcap/tidb/pull/23152)                                                                       |   | 15d14h  |
| tidb/23191   | [planner/core: convert decimal type for mpp join before shuffling.](https://github.com/pingcap/tidb/pull/23191)                                               |   | 13d19h  |
| tidb/23196   | [types: fix the bug about the wrong query result for decimal type  (#22507)](https://github.com/pingcap/tidb/pull/23196)                                      |   | 13d18h  |
| tidb/23210   | [planner: fixed a bug that prevented SPM from taking effect (#23197)](https://github.com/pingcap/tidb/pull/23210)                                             |   | 13d16h  |
| tidb/23220   | [Release 4.0](https://github.com/pingcap/tidb/pull/23220)                                                                                                     |   | 13d11h  |
| tidb/23227   | [executor: hash join out of index panic when enum column value is zero (#23162)](https://github.com/pingcap/tidb/pull/23227)                                  |   | 12d22h  |
| tidb/23233   | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23233)                                                        |   | 12d18h  |
| tidb/23234   | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23234)                                                        |   | 12d18h  |
| tidb/23245   | [*: Add security enhanced mode as experimental](https://github.com/pingcap/tidb/pull/23245)                                                                   |   | 12d6h   |
| tidb/23257   | [executor: group_concat aggr panic when session.group_concat_max_len is small (#23131)](https://github.com/pingcap/tidb/pull/23257)                           |   | 11d18h  |
| tidb/23278   | [executor: wrong result of nullif expr when used with is null expr. (#23170)](https://github.com/pingcap/tidb/pull/23278)                                     |   | 10d18h  |
| tidb/23295   | [util, types: don't let SPM be affected by charset (#23161)](https://github.com/pingcap/tidb/pull/23295)                                                      |   | 10d11h  |
| tidb/23335   | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23335)                                      |   | 6d19h   |
| tidb/23336   | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23336)                                      |   | 6d19h   |
| tidb/23337   | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23337)                                      |   | 6d19h   |
| tidb/23340   | [executor: fix unexpected NotNullFlag in case when expr ret type (#23102)](https://github.com/pingcap/tidb/pull/23340)                                        |   | 6d19h   |
| tidb/23347   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                                          |   | 6d18h   |
| tidb/23348   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                                          |   | 6d18h   |
| tidb/23350   | [util/stringutil, util/ranger, planner: use hierarchical separators to simplify the parsing for info of EXPLAIN ](https://github.com/pingcap/tidb/pull/23350) |   | 6d17h   |
| tidb/23368   | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23368)                                         |   | 5d20h   |
| tidb/23369   | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23369)                                         |   | 5d20h   |
| tidb/23370   | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23370)                                         |   | 5d20h   |
| tidb/23374   | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23374)                                             |   | 5d19h   |
| tidb/23397   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23397)                                                                |   | 4d17h   |
| tidb/23398   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23398)                                                                |   | 4d17h   |
| tidb/23399   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23399)                                                                |   | 4d17h   |
| tidb/23401   | [planner, sessionctx: turn on the mpp by default](https://github.com/pingcap/tidb/pull/23401)                                                                 |   | 4d17h   |
| tidb/23405   | [domain: remove the exit chan, use context](https://github.com/pingcap/tidb/pull/23405)                                                                       |   | 4d17h   |
| tidb/23426   | [*: Add the metric about the SQL with TiFlash Success ](https://github.com/pingcap/tidb/pull/23426)                                                           |   | 3d22h   |
| tidb/23433   | [WIP: speed up for slow query logs retrieving ](https://github.com/pingcap/tidb/pull/23433)                                                                   |   | 3d17h   |
| tidb/23435   | [planner: check schema stale for plan cache when forUpdateRead (#22381)](https://github.com/pingcap/tidb/pull/23435)                                          |   | 3d16h   |
| tidb/23444   | [wip: execution: refine allocating iterator for slice](https://github.com/pingcap/tidb/pull/23444)                                                            |   | 3d10h   |
| tidb/23457   | [planner: fix inappropriate null flag of null constants](https://github.com/pingcap/tidb/pull/23457)                                                          |   | 14h     |


## Reviewed in Last 7 Days

|    REPO    |                                                                        PR                                                                         | C | D |   R   |
|------------|---------------------------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/23104 | [executor: fix wrong key range of index scan when filter is comparing year column with NULL (#23079)](https://github.com/pingcap/tidb/pull/23104) |   | 5 | 14d4h |
| tidb/23279 | [executor: wrong result of nullif expr when used with is null expr. (#23170)](https://github.com/pingcap/tidb/pull/23279)                         |   | 5 | 6d3h  |
| tidb/23339 | [expression: fix refine compare constant](https://github.com/pingcap/tidb/pull/23339)                                                             |   | 6 | 20h   |


</details> 


<details> 
  <summary>dyzsr's detailed review report</summary> 

## To Be Reviewed

| REPO | PR | C | LASTED |
|------|----|---|--------|


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>eurekaka's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                   PR                                                                   | C | LASTED  |
|------------|----------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19347 | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347)   |   | 213d23h |
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                              |   | 137d17h |
| tidb/21444 | [planner: ignore anonymous index while tiflash replica is available](https://github.com/pingcap/tidb/pull/21444)                       |   | 110d12h |
| tidb/21994 | [range: fix overflow value access index ](https://github.com/pingcap/tidb/pull/21994)                                                  |   | 88d23h  |
| tidb/22342 | [session: fix two cases when updating bind info (#22338)](https://github.com/pingcap/tidb/pull/22342)                                  |   | 70d18h  |
| tidb/22354 | [planner: do not cache prepared plan if optimization depends on mutable constant (#22349)](https://github.com/pingcap/tidb/pull/22354) |   | 69d23h  |
| tidb/22369 | [session: fix the duplicate binding case when updating bind info (#22367)](https://github.com/pingcap/tidb/pull/22369)                 |   | 69d17h  |
| tidb/22416 | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                                      | Y | 65d11h  |
| tidb/22559 | [planner: split test data from test cases in cbo_test.go](https://github.com/pingcap/tidb/pull/22559)                                  |   | 54d19h  |
| tidb/22778 | [*: add support for dynamic privileges](https://github.com/pingcap/tidb/pull/22778)                                                    |   | 34d7h   |
| tidb/23137 | [planner: fix index merge row count estimation logic](https://github.com/pingcap/tidb/pull/23137)                                      |   | 17d17h  |
| tidb/23208 | [statistics, util/ranger: improve selectivity calculation for DNF filters (#18741)](https://github.com/pingcap/tidb/pull/23208)        |   | 13d16h  |
| tidb/23210 | [planner: fixed a bug that prevented SPM from taking effect (#23197)](https://github.com/pingcap/tidb/pull/23210)                      |   | 13d16h  |
| tidb/23216 | [Privileges: fix delete privilege check wrongly (#22971)](https://github.com/pingcap/tidb/pull/23216)                                  |   | 13d14h  |
| tidb/23283 | [util: optimize the performance of restore with db (#22910)](https://github.com/pingcap/tidb/pull/23283)                               |   | 10d17h  |
| tidb/23295 | [util, types: don't let SPM be affected by charset (#23161)](https://github.com/pingcap/tidb/pull/23295)                               |   | 10d11h  |
| tidb/23316 | [planner: Fix rebuild range for prepared plan](https://github.com/pingcap/tidb/pull/23316)                                             |   | 7d17h   |
| tidb/23365 | [planner: fix a bug that point get plan returns wrong column name](https://github.com/pingcap/tidb/pull/23365)                         |   | 5d22h   |
| tidb/23373 | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                      |   | 5d19h   |
| tidb/23374 | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23374)                      |   | 5d19h   |
| tidb/23404 | [planner: fix only_full_group_by check not enough](https://github.com/pingcap/tidb/pull/23404)                                         |   | 4d17h   |


## Reviewed in Last 7 Days

|    REPO    |                                                         PR                                                          | C | D |   R    |
|------------|---------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/23175 | [planner: fix plan cache not working caused by type difference](https://github.com/pingcap/tidb/pull/23175)         |   | 1 | 13d19h |
| tidb/22853 | [planner: fix LogicalPlans that contain Window Function are ambiguous ](https://github.com/pingcap/tidb/pull/22853) |   | 7 | 23d17h |


</details> 


<details> 
  <summary>fzhedu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                         PR                                                          | C | LASTED |
|------------|---------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/22853 | [planner: fix LogicalPlans that contain Window Function are ambiguous ](https://github.com/pingcap/tidb/pull/22853) |   | 30d12h |


## Reviewed in Last 7 Days

|    REPO    |                                                       PR                                                        | C | D |   R   |
|------------|-----------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/23401 | [planner, sessionctx: turn on the mpp by default](https://github.com/pingcap/tidb/pull/23401)                   |   | 4 | 1d3h  |
| tidb/23191 | [planner/core: convert decimal type for mpp join before shuffling.](https://github.com/pingcap/tidb/pull/23191) |   | 4 | 9d21h |


</details> 


<details> 
  <summary>ichn-hu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                            PR                                                             | C | LASTED |
|------------|---------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/21853 | [expression: fix compatibility behaviors in time_format with MySQL (#21559)](https://github.com/pingcap/tidb/pull/21853)  |   | 95d19h |
| tidb/23278 | [executor: wrong result of nullif expr when used with is null expr. (#23170)](https://github.com/pingcap/tidb/pull/23278) |   | 10d18h |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>lzmhhh123's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                             PR                                                                              | C | LASTED  |
|--------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19347   | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347)                        |   | 213d23h |
| docs-cn/5781 | [hide the config global-kill (#5770)](https://github.com/pingcap/docs-cn/pull/5781)                                                                         |   | 3d13h   |
| tidb/20444   | [expression: add json_merge_patch](https://github.com/pingcap/tidb/pull/20444)                                                                              |   | 159d21h |
| tidb/20465   | [expression: add uuidShortFunction](https://github.com/pingcap/tidb/pull/20465)                                                                             |   | 158d19h |
| tidb/20642   | [executor: modify admin executors to support partitioned table with global index](https://github.com/pingcap/tidb/pull/20642)                               |   | 147d15h |
| tidb/20903   | [planner: fix confused and unnecessary double-projection in plans.](https://github.com/pingcap/tidb/pull/20903)                                             |   | 136d17h |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                              |   | 130d17h |
| tidb/21195   | [brie: integrate lightning to suport IMPORT statement](https://github.com/pingcap/tidb/pull/21195)                                                          |   | 119d22h |
| tidb/21334   | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                                               |   | 116d14h |
| tidb/21347   | [session: make rollback work on global variables](https://github.com/pingcap/tidb/pull/21347)                                                               |   | 115d19h |
| tidb/21444   | [planner: ignore anonymous index while tiflash replica is available](https://github.com/pingcap/tidb/pull/21444)                                            |   | 110d12h |
| tidb/21487   | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                                               |   | 109d4h  |
| tidb/21641   | [executor: Fix pessimistic lock doesn't work on the partition table for subquery/joins](https://github.com/pingcap/tidb/pull/21641)                         |   | 102d18h |
| tidb/21651   | [planner: allow filter condition pushing down to IndexScan for prefix index](https://github.com/pingcap/tidb/pull/21651)                                    |   | 102d13h |
| tidb/22126   | [*: add `sys` schema, `sys.SCHEMA_UNUSED_INDEXES` view and `sys.SCHEMA_INDEX_USAGE` view](https://github.com/pingcap/tidb/pull/22126)                       |   | 81d20h  |
| tidb/22149   | [session: set process info before building plan (#22101)](https://github.com/pingcap/tidb/pull/22149)                                                       |   | 77d19h  |
| tidb/22188   | [planner: do not use indexMerge when the path only use a single index (#22168)](https://github.com/pingcap/tidb/pull/22188)                                 |   | 76d13h  |
| tidb/22361   | [table: fix insert into _tidb_rowid panic and rebase it if needed (#22062)](https://github.com/pingcap/tidb/pull/22361)                                     |   | 69d20h  |
| tidb/22372   | [executor: fix SelectForUpdate in decorrelated subquery under pessimistic mode](https://github.com/pingcap/tidb/pull/22372)                                 |   | 69d9h   |
| tidb/22478   | [planner, executor: fix query partition table with global unique index get wrong result](https://github.com/pingcap/tidb/pull/22478)                        |   | 60d13h  |
| tidb/22631   | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                                                                             |   | 50d23h  |
| tidb/22686   | [expression: support enum pushdown](https://github.com/pingcap/tidb/pull/22686)                                                                             |   | 47d22h  |
| tidb/22699   | [brie: add error info column and history backup/restore info in sql](https://github.com/pingcap/tidb/pull/22699)                                            |   | 47d16h  |
| tidb/22857   | [mocktikv: split rpcHandler to kvHandler and coprHandler](https://github.com/pingcap/tidb/pull/22857)                                                       |   | 29d21h  |
| tidb/22926   | [expression: add overflow check in multiplyInt](https://github.com/pingcap/tidb/pull/22926)                                                                 |   | 26d13h  |
| tidb/23001   | [statistics: fix err check](https://github.com/pingcap/tidb/pull/23001)                                                                                     |   | 23d0h   |
| tidb/23022   | [executor: create PipelinedWindowExec based on current implementation and modify the windowProcessor interface](https://github.com/pingcap/tidb/pull/23022) |   | 21d18h  |
| tidb/23149   | [core: support left join and right join for join reorder](https://github.com/pingcap/tidb/pull/23149)                                                       |   | 16d12h  |
| tidb/23210   | [planner: fixed a bug that prevented SPM from taking effect (#23197)](https://github.com/pingcap/tidb/pull/23210)                                           |   | 13d16h  |
| tidb/23257   | [executor: group_concat aggr panic when session.group_concat_max_len is small (#23131)](https://github.com/pingcap/tidb/pull/23257)                         |   | 11d18h  |
| tidb/23278   | [executor: wrong result of nullif expr when used with is null expr. (#23170)](https://github.com/pingcap/tidb/pull/23278)                                   |   | 10d18h  |
| tidb/23283   | [util: optimize the performance of restore with db (#22910)](https://github.com/pingcap/tidb/pull/23283)                                                    |   | 10d17h  |
| tidb/23293   | [planner: fix the bug that wrong collation is used when try fast path for enum or set (#23217)](https://github.com/pingcap/tidb/pull/23293)                 |   | 10d14h  |
| tidb/23296   | [sig/execution: fix the bug that Wrong result of comparison operation(type date / type string)](https://github.com/pingcap/tidb/pull/23296)                 |   | 10d7h   |
| tidb/23307   | [util/chunk: replace outdated link with correct one](https://github.com/pingcap/tidb/pull/23307)                                                            |   | 7d20h   |
| tidb/23334   | [metrics/grafana: Remove duplicate items "Owner Watcher OPS"](https://github.com/pingcap/tidb/pull/23334)                                                   |   | 6d20h   |
| tidb/23337   | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23337)                                    |   | 6d19h   |
| tidb/23340   | [executor: fix unexpected NotNullFlag in case when expr ret type (#23102)](https://github.com/pingcap/tidb/pull/23340)                                      |   | 6d19h   |
| tidb/23347   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                                        |   | 6d18h   |
| tidb/23348   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                                        |   | 6d18h   |
| tidb/23368   | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23368)                                       |   | 5d20h   |
| tidb/23369   | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23369)                                       |   | 5d20h   |
| tidb/23370   | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23370)                                       |   | 5d20h   |
| tidb/23373   | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                           |   | 5d19h   |
| tidb/23374   | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23374)                                           |   | 5d19h   |
| tidb/23399   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23399)                                                              |   | 4d17h   |
| tidb/23417   | [plan: reset not null flag](https://github.com/pingcap/tidb/pull/23417)                                                                                     |   | 4d11h   |
| tidb/23422   | [sig/execution: fix the bug that The result of 'varbinary + 1' is incorrect](https://github.com/pingcap/tidb/pull/23422)                                    |   | 4d6h    |
| tidb/23441   | [executor: Refactor probe channel & fix bug in chunks of join](https://github.com/pingcap/tidb/pull/23441)                                                  |   | 3d14h   |
| tidb/23447   | [*: Support record stmtsummary evicted count](https://github.com/pingcap/tidb/pull/23447)                                                                   |   | 1d18h   |
| tidb/23461   | [execdetails: refine cop task execution stats display in plan](https://github.com/pingcap/tidb/pull/23461)                                                  |   | 12h     |


## Reviewed in Last 7 Days

|        REPO         |                                                                      PR                                                                       | C | D |   R   |
|---------------------|-----------------------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tikv/9850           | [copr: fix IN expr didn't handle unsigned/signed int properly (#9823)](https://github.com/tikv/tikv/pull/9850)                                |   | 4 | 1h    |
| tidb/23410          | [test: add testleak after checking and testing (#23324)](https://github.com/pingcap/tidb/pull/23410)                                          |   | 5 | 2h    |
| tidb/23409          | [plan: setting not null flag for extrak pk (#23237)](https://github.com/pingcap/tidb/pull/23409)                                              |   | 5 | 2h    |
| tidb/23407          | [*: hide the config `global-kill` and session var `tidb_enable_index_merge_join` (#23395)](https://github.com/pingcap/tidb/pull/23407)        |   | 5 | 0h    |
| automated-tests/606 | [fix test case after @@tidb_enable_cluster_index became global var](https://github.com/pingcap/automated-tests/pull/606)                      |   | 5 | 0h    |
| tidb/23332          | [excutor: fix the date precision of `builtinCastDurationAsStringSig.vecEvalString` #23314 #23286](https://github.com/pingcap/tidb/pull/23332) |   | 5 | 2d3h  |
| docs-cn/5770        | [hide the config global-kill](https://github.com/pingcap/docs-cn/pull/5770)                                                                   |   | 5 | 0h    |
| tidb/23237          | [plan: setting not null flag for extrak pk](https://github.com/pingcap/tidb/pull/23237)                                                       |   | 5 | 7d23h |
| tidb/23395          | [*: hide the config `global-kill` and session var `tidb_enable_index_merge_join`](https://github.com/pingcap/tidb/pull/23395)                 |   | 5 | 0h    |
| tidb/23361          | [*: add clustered index info in `show index from` stmt (#23329)](https://github.com/pingcap/tidb/pull/23361)                                  |   | 5 | 1d18h |
| tidb/23372          | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23372)                             |   | 5 | 23h   |
| tidb/23324          | [test: add testleak after checking and testing](https://github.com/pingcap/tidb/pull/23324)                                                   |   | 5 | 2d17h |
| tikv/9823           | [copr: fix IN expr didn't handle unsigned/signed int properly](https://github.com/tikv/tikv/pull/9823)                                        |   | 5 | 2d7h  |
| tidb/23241          | [executor: fix get var expr when session var is hex literal](https://github.com/pingcap/tidb/pull/23241)                                      |   | 6 | 6d17h |
| tidb/23356          | [planner: fix {index,hash,merge} join on range suffix condition clustered index (#23294)](https://github.com/pingcap/tidb/pull/23356)         |   | 7 | 0h    |
| tidb-test/1168      | [Add cast target type in explain result.](https://github.com/pingcap/tidb-test/pull/1168)                                                     |   | 7 | 0h    |
| tidb/23285          | [executor, expression: fix the incorrect result of AVG function](https://github.com/pingcap/tidb/pull/23285)                                  |   | 7 | 3d20h |


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
| docs/4781    | [system variable: add tidb_allow_fallback_to_tikv](https://github.com/pingcap/docs/pull/4781)                                                               |   | 47d17h  |
| docs-cn/5561 | [Add sql optimization-related docs to toc](https://github.com/pingcap/docs-cn/pull/5561)                                                                    |   | 28d15h  |
| docs-cn/5785 | [update SPM documentation for DML SQL Bind and baseline capture (#5740)](https://github.com/pingcap/docs-cn/pull/5785)                                      |   | 18h     |
| tidb/19029   | [types: fix unexpected NOT_NULL flags](https://github.com/pingcap/tidb/pull/19029)                                                                          |   | 228d22h |
| docs-cn/5786 | [update SPM documentation for DML SQL Bind and baseline capture (#5740)](https://github.com/pingcap/docs-cn/pull/5786)                                      |   | 18h     |
| tidb/20708   | [*: separate auto_increment ID allocator from _tidb_rowid allocator](https://github.com/pingcap/tidb/pull/20708)                                            |   | 144d20h |
| tidb/20969   | [executor: Improve the performance of appending not fixed columns](https://github.com/pingcap/tidb/pull/20969)                                              |   | 132d9h  |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                              |   | 130d17h |
| tidb/21149   | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                               |   | 123d14h |
| tidb/21304   | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                                    |   | 117d12h |
| tidb/21318   | [planner, expression: use the range of column types to simplify expressions](https://github.com/pingcap/tidb/pull/21318)                                    |   | 116d19h |
| tidb/21401   | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                                          |   | 112d11h |
| tidb/21476   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                       |   | 109d15h |
| tidb/21508   | [execution: fix dayofweek('0000-00-00') behavior](https://github.com/pingcap/tidb/pull/21508)                                                               |   | 108d10h |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                                 |   | 94d19h  |
| tidb/21887   | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                                             |   | 93d11h  |
| tidb/21954   | [planner/cascades: add rule `PushSelDownApply`](https://github.com/pingcap/tidb/pull/21954)                                                                 |   | 89d23h  |
| tidb/22146   | [executor: forbid SFU on view](https://github.com/pingcap/tidb/pull/22146)                                                                                  |   | 77d21h  |
| tidb/22217   | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                                               |   | 75d17h  |
| tidb/22234   | [executor, planner: ON DUPLICATE UPDATE can refer to un-project col (#14412)](https://github.com/pingcap/tidb/pull/22234)                                   |   | 75d15h  |
| tidb/22261   | [time: fix parse datetime won't truncate the reluctant string (#22232)](https://github.com/pingcap/tidb/pull/22261)                                         |   | 74d19h  |
| tidb/22294   | [planner, table: optimize the list partition pruner for range query](https://github.com/pingcap/tidb/pull/22294)                                            |   | 73d20h  |
| tidb/22307   | [ddl: fix update can see columns not public](https://github.com/pingcap/tidb/pull/22307)                                                                    |   | 73d16h  |
| tidb/22342   | [session: fix two cases when updating bind info (#22338)](https://github.com/pingcap/tidb/pull/22342)                                                       |   | 70d18h  |
| tidb/22369   | [session: fix the duplicate binding case when updating bind info (#22367)](https://github.com/pingcap/tidb/pull/22369)                                      |   | 69d17h  |
| tidb/22374   | [expression: separated arithmeticIntDivideSig](https://github.com/pingcap/tidb/pull/22374)                                                                  |   | 69d0h   |
| tidb/22415   | [ddl: refactor placement package](https://github.com/pingcap/tidb/pull/22415)                                                                               |   | 65d17h  |
| tidb/22541   | [expression: Support builtin function SOUNDEX](https://github.com/pingcap/tidb/pull/22541)                                                                  |   | 55d9h   |
| tidb/22559   | [planner: split test data from test cases in cbo_test.go](https://github.com/pingcap/tidb/pull/22559)                                                       |   | 54d19h  |
| tidb/22565   | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)                                          | Y | 54d17h  |
| tidb/22814   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22814)                                         |   | 31d19h  |
| tidb/22815   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22815)                                         |   | 31d19h  |
| tidb/22862   | [brie: fix the problem that ddl restored by BR via SQL is not replicated to downstream](https://github.com/pingcap/tidb/pull/22862)                         |   | 28d22h  |
| tidb/22915   | [planner: build correct MaxOneRow info from multi-column conditions](https://github.com/pingcap/tidb/pull/22915)                                            |   | 26d18h  |
| tidb/22923   | [expression: correct constant propagation for collation (#22666)](https://github.com/pingcap/tidb/pull/22923)                                               |   | 26d15h  |
| tidb/22924   | [planner: fix wrong index merge selection (#22825)](https://github.com/pingcap/tidb/pull/22924)                                                             |   | 26d14h  |
| tidb/22926   | [expression: add overflow check in multiplyInt](https://github.com/pingcap/tidb/pull/22926)                                                                 |   | 26d13h  |
| tidb/22984   | [executor: fix logging format of prepared statements (#16062)](https://github.com/pingcap/tidb/pull/22984)                                                  |   | 23d10h  |
| tidb/23022   | [executor: create PipelinedWindowExec based on current implementation and modify the windowProcessor interface](https://github.com/pingcap/tidb/pull/23022) |   | 21d18h  |
| tidb/23062   | [*: fix structcheck lint warnings](https://github.com/pingcap/tidb/pull/23062)                                                                              |   | 19d19h  |
| tidb/23105   | [executor: fix wrong key range of index scan when filter is comparing year column with NULL (#23079)](https://github.com/pingcap/tidb/pull/23105)           |   | 18d18h  |
| tidb/23137   | [planner: fix index merge row count estimation logic](https://github.com/pingcap/tidb/pull/23137)                                                           |   | 17d17h  |
| tidb/23152   | [expression: fix wrong error info (#22760)](https://github.com/pingcap/tidb/pull/23152)                                                                     |   | 15d14h  |
| tidb/23196   | [types: fix the bug about the wrong query result for decimal type  (#22507)](https://github.com/pingcap/tidb/pull/23196)                                    |   | 13d18h  |
| tidb/23201   | [executor, server: load_data.go is changed and add unit test](https://github.com/pingcap/tidb/pull/23201)                                                   |   | 13d17h  |
| tidb/23208   | [statistics, util/ranger: improve selectivity calculation for DNF filters (#18741)](https://github.com/pingcap/tidb/pull/23208)                             |   | 13d16h  |
| tidb/23210   | [planner: fixed a bug that prevented SPM from taking effect (#23197)](https://github.com/pingcap/tidb/pull/23210)                                           |   | 13d16h  |
| tidb/23234   | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23234)                                                      |   | 12d18h  |
| tidb/23283   | [util: optimize the performance of restore with db (#22910)](https://github.com/pingcap/tidb/pull/23283)                                                    |   | 10d17h  |
| tidb/23284   | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified.](https://github.com/pingcap/tidb/pull/23284)              |   | 10d17h  |
| tidb/23293   | [planner: fix the bug that wrong collation is used when try fast path for enum or set (#23217)](https://github.com/pingcap/tidb/pull/23293)                 |   | 10d14h  |
| tidb/23295   | [util, types: don't let SPM be affected by charset (#23161)](https://github.com/pingcap/tidb/pull/23295)                                                    |   | 10d11h  |
| tidb/23316   | [planner: Fix rebuild range for prepared plan](https://github.com/pingcap/tidb/pull/23316)                                                                  |   | 7d17h   |
| tidb/23373   | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                           |   | 5d19h   |
| tidb/23397   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23397)                                                              |   | 4d17h   |
| tidb/23398   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23398)                                                              |   | 4d17h   |
| tidb/23399   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23399)                                                              |   | 4d17h   |
| tidb/23448   | [wip :execution: parallel build hash table](https://github.com/pingcap/tidb/pull/23448)                                                                     |   | 1d12h   |


## Reviewed in Last 7 Days

|     REPO     |                                                                                       PR                                                                                        | C | D |    R    |
|--------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---|---------|
| tidb/23458   | [planner: incorrect results returned when executing a prepared point-get update meets nondeterministic function. (#22067) (#23455)](https://github.com/pingcap/tidb/pull/23458) |   | 1 | 0h      |
| tidb/23455   | [planner: incorrect results returned when executing a prepared point-get update meets nondeterministic function. (#22067)](https://github.com/pingcap/tidb/pull/23455)          |   | 1 | 1h      |
| tidb/23343   | [statistics: handle drop partition events for global-stats](https://github.com/pingcap/tidb/pull/23343)                                                                         |   | 4 | 3d1h    |
| docs-cn/5484 | [system variable: add tidb_allow_fallback_to_tikv](https://github.com/pingcap/docs-cn/pull/5484)                                                                                |   | 4 | 44d0h   |
| tidb/23401   | [planner, sessionctx: turn on the mpp by default](https://github.com/pingcap/tidb/pull/23401)                                                                                   |   | 4 | 23h     |
| tidb/23191   | [planner/core: convert decimal type for mpp join before shuffling.](https://github.com/pingcap/tidb/pull/23191)                                                                 |   | 4 | 10d0h   |
| docs-cn/5740 | [update SPM documentation for DML SQL Bind and baseline capture](https://github.com/pingcap/docs-cn/pull/5740)                                                                  |   | 5 | 3d0h    |
| tidb/23292   | [planner: fix the bug that wrong collation is used when try fast path for enum or set (#23217)](https://github.com/pingcap/tidb/pull/23292)                                     |   | 5 | 5d20h   |
| tidb/23349   | [*: hide `index-usage-sync-lease` config](https://github.com/pingcap/tidb/pull/23349)                                                                                           |   | 5 | 1d22h   |
| tidb/20905   | [planner: fix statement-optimize not work in `TryFastPlan`](https://github.com/pingcap/tidb/pull/20905)                                                                         |   | 5 | 131d20h |
| tidb/23345   | [sessionctx: hide extended stats variable in SHOW VARIABLES temporarily](https://github.com/pingcap/tidb/pull/23345)                                                            |   | 5 | 1d19h   |
| tidb/22090   | [planner: push aggregation operators down to projection by default](https://github.com/pingcap/tidb/pull/22090)                                                                 |   | 6 | 77d3h   |
| tidb/23238   | [planner: fix wrong PointGet / TableDual plan reused in plan cache](https://github.com/pingcap/tidb/pull/23238)                                                                 |   | 6 | 6d22h   |
| tidb/23333   | [statistics: hide the `tidb_analyze_version` before it's GA (#23327)](https://github.com/pingcap/tidb/pull/23333)                                                               |   | 7 | 1h      |
| tidb/23327   | [statistics: hide the `tidb_analyze_version` before it's GA](https://github.com/pingcap/tidb/pull/23327)                                                                        |   | 7 | 8h      |


</details> 


<details> 
  <summary>rebelice's detailed review report</summary> 

## To Be Reviewed

| REPO | PR | C | LASTED |
|------|----|---|--------|


## Reviewed in Last 7 Days

|    REPO    |                                                      PR                                                       | C | D |   R   |
|------------|---------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/23343 | [statistics: handle drop partition events for global-stats](https://github.com/pingcap/tidb/pull/23343)       |   | 1 | 5d22h |
| tidb/23354 | [planner: remove some risky cache operations in the plan builder](https://github.com/pingcap/tidb/pull/23354) |   | 6 | 22h   |


</details> 


<details> 
  <summary>time-and-fate's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                         PR                                                          | C | LASTED  |
|------------|---------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                           |   | 137d17h |
| tidb/22853 | [planner: fix LogicalPlans that contain Window Function are ambiguous ](https://github.com/pingcap/tidb/pull/22853) |   | 30d12h  |
| tidb/22915 | [planner: build correct MaxOneRow info from multi-column conditions](https://github.com/pingcap/tidb/pull/22915)    |   | 26d18h  |


## Reviewed in Last 7 Days

|    REPO    |                                                         PR                                                          | C | D |   R    |
|------------|---------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/23238 | [planner: fix wrong PointGet / TableDual plan reused in plan cache](https://github.com/pingcap/tidb/pull/23238)     |   | 6 | 6d22h  |
| tidb/23119 | [statistics: remove existing deleted extended stats when add a new one](https://github.com/pingcap/tidb/pull/23119) |   | 6 | 12d19h |


</details> 


<details> 
  <summary>winoros's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                              PR                                                                               | C | LASTED  |
|------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19957 | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                                        | Y | 192d14h |
| docs/4781  | [system variable: add tidb_allow_fallback_to_tikv](https://github.com/pingcap/docs/pull/4781)                                                                 |   | 47d17h  |
| tidb/20311 | [expression: fix overflow error when convert bit to int64 (#20266)](https://github.com/pingcap/tidb/pull/20311)                                               |   | 173d21h |
| tidb/20765 | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                                             |   | 140d17h |
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                                                     |   | 137d17h |
| tidb/21018 | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                                |   | 130d17h |
| tidb/21207 | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                                               |   | 119d19h |
| tidb/21476 | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                         |   | 109d15h |
| tidb/21487 | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                                                 |   | 109d4h  |
| tidb/21876 | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                                   |   | 94d19h  |
| tidb/21954 | [planner/cascades: add rule `PushSelDownApply`](https://github.com/pingcap/tidb/pull/21954)                                                                   |   | 89d23h  |
| tidb/22181 | [planner, expression: fix error when using IN combined with subquery (#22080)](https://github.com/pingcap/tidb/pull/22181)                                    |   | 76d17h  |
| tidb/22365 | [planner: check index valid while forUpdateRead (#22152)](https://github.com/pingcap/tidb/pull/22365)                                                         |   | 69d19h  |
| tidb/22504 | [*:Fix the fetchHotRegion bug that the count always zero](https://github.com/pingcap/tidb/pull/22504)                                                         |   | 57d19h  |
| tidb/22565 | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)                                            | Y | 54d17h  |
| tidb/22624 | [ planner: not pruning column used by union scan condition (#21640)](https://github.com/pingcap/tidb/pull/22624)                                              |   | 52d17h  |
| tidb/22923 | [expression: correct constant propagation for collation (#22666)](https://github.com/pingcap/tidb/pull/22923)                                                 |   | 26d15h  |
| tidb/23163 | [plugin: fix linter --enable=deadcode check error](https://github.com/pingcap/tidb/pull/23163)                                                                |   | 14d19h  |
| tidb/23208 | [statistics, util/ranger: improve selectivity calculation for DNF filters (#18741)](https://github.com/pingcap/tidb/pull/23208)                               |   | 13d16h  |
| tidb/23215 | [Privileges: fix delete privilege check wrongly (#22971)](https://github.com/pingcap/tidb/pull/23215)                                                         |   | 13d14h  |
| tidb/23216 | [Privileges: fix delete privilege check wrongly (#22971)](https://github.com/pingcap/tidb/pull/23216)                                                         |   | 13d14h  |
| tidb/23233 | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23233)                                                        |   | 12d18h  |
| tidb/23234 | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23234)                                                        |   | 12d18h  |
| tidb/23246 | [planner: fix the panic in joinReOrderSolver.optimizeRecursive](https://github.com/pingcap/tidb/pull/23246)                                                   |   | 11d23h  |
| tidb/23347 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                                          |   | 6d18h   |
| tidb/23348 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                                          |   | 6d18h   |
| tidb/23350 | [util/stringutil, util/ranger, planner: use hierarchical separators to simplify the parsing for info of EXPLAIN ](https://github.com/pingcap/tidb/pull/23350) |   | 6d17h   |
| tidb/23365 | [planner: fix a bug that point get plan returns wrong column name](https://github.com/pingcap/tidb/pull/23365)                                                |   | 5d22h   |
| tidb/23373 | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                             |   | 5d19h   |
| tidb/23435 | [planner: check schema stale for plan cache when forUpdateRead (#22381)](https://github.com/pingcap/tidb/pull/23435)                                          |   | 3d16h   |


## Reviewed in Last 7 Days

|     REPO     |                                                          PR                                                          | C | D |    R    |
|--------------|----------------------------------------------------------------------------------------------------------------------|---|---|---------|
| tidb/23175   | [planner: fix plan cache not working caused by type difference](https://github.com/pingcap/tidb/pull/23175)          |   | 1 | 14d0h   |
| tidb/23349   | [*: hide `index-usage-sync-lease` config](https://github.com/pingcap/tidb/pull/23349)                                |   | 1 | 6d1h    |
| tidb/23404   | [planner: fix only_full_group_by check not enough](https://github.com/pingcap/tidb/pull/23404)                       |   | 1 | 3d21h   |
| tidb/23437   | [statistics: check duplicate columns and types for extended stats](https://github.com/pingcap/tidb/pull/23437)       |   | 4 | 0h      |
| tidb/22381   | [planner: check schema stale for plan cache when forUpdateRead](https://github.com/pingcap/tidb/pull/22381)          |   | 4 | 64d21h  |
| docs-cn/5484 | [system variable: add tidb_allow_fallback_to_tikv](https://github.com/pingcap/docs-cn/pull/5484)                     |   | 4 | 43d23h  |
| docs-cn/5762 | [Update statistics.md](https://github.com/pingcap/docs-cn/pull/5762)                                                 |   | 5 | 5h      |
| tidb/20905   | [planner: fix statement-optimize not work in `TryFastPlan`](https://github.com/pingcap/tidb/pull/20905)              |   | 5 | 131d20h |
| tidb/23346   | [statistics: fix build extended stats panic when met NULL](https://github.com/pingcap/tidb/pull/23346)               |   | 5 | 1d21h   |
| tidb/23345   | [sessionctx: hide extended stats variable in SHOW VARIABLES temporarily](https://github.com/pingcap/tidb/pull/23345) |   | 5 | 1d21h   |
| tidb/23354   | [planner: remove some risky cache operations in the plan builder](https://github.com/pingcap/tidb/pull/23354)        |   | 6 | 22h     |
| tidb/23328   | [planner: hide the switch of dynamic-pruning and global-stats (#23320)](https://github.com/pingcap/tidb/pull/23328)  |   | 7 | 3h      |


</details> 


<details> 
  <summary>wshwsh12's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                                       PR                                                                                        | C | LASTED  |
|------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19807 | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                                                                         |   | 199d11h |
| tidb/19957 | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                                                          | Y | 192d14h |
| tidb/21487 | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                                                                   |   | 109d4h  |
| tidb/21887 | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                                                                 |   | 93d11h  |
| tidb/22378 | [executor: vectorize hash aggregate](https://github.com/pingcap/tidb/pull/22378)                                                                                                |   | 68d19h  |
| tidb/22628 | [executor: Improve max/min window function with deque-based sliding window](https://github.com/pingcap/tidb/pull/22628)                                                         |   | 51d23h  |
| tidb/22815 | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22815)                                                             |   | 31d19h  |
| tidb/23105 | [executor: fix wrong key range of index scan when filter is comparing year column with NULL (#23079)](https://github.com/pingcap/tidb/pull/23105)                               |   | 18d18h  |
| tidb/23128 | [statistics: refactor the statistics package use the RestrictedSQLExecutor API (#22636)](https://github.com/pingcap/tidb/pull/23128)                                            |   | 17d22h  |
| tidb/23234 | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23234)                                                                          |   | 12d18h  |
| tidb/23336 | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23336)                                                        |   | 6d19h   |
| tidb/23347 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                                                            |   | 6d18h   |
| tidb/23348 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                                                            |   | 6d18h   |
| tidb/23368 | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23368)                                                           |   | 5d20h   |
| tidb/23369 | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23369)                                                           |   | 5d20h   |
| tidb/23370 | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23370)                                                           |   | 5d20h   |
| tidb/23397 | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23397)                                                                                  |   | 4d17h   |
| tidb/23398 | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23398)                                                                                  |   | 4d17h   |
| tidb/23399 | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23399)                                                                                  |   | 4d17h   |
| tidb/23435 | [planner: check schema stale for plan cache when forUpdateRead (#22381)](https://github.com/pingcap/tidb/pull/23435)                                                            |   | 3d16h   |
| tidb/23445 | [chore: Refactor code quality issues](https://github.com/pingcap/tidb/pull/23445)                                                                                               |   | 2d20h   |
| tidb/23454 | [telemetry: add copr-cache, tiflash, cluster index, async commit](https://github.com/pingcap/tidb/pull/23454)                                                                   |   | 17h     |
| tidb/23458 | [planner: incorrect results returned when executing a prepared point-get update meets nondeterministic function. (#22067) (#23455)](https://github.com/pingcap/tidb/pull/23458) |   | 14h     |


## Reviewed in Last 7 Days

|    REPO    |                                          PR                                           | C | D |   R   |
|------------|---------------------------------------------------------------------------------------|---|---|-------|
| tidb/23339 | [expression: fix refine compare constant](https://github.com/pingcap/tidb/pull/23339) |   | 5 | 2d0h  |
| tidb/23301 | [expression: fix unused code in util.go](https://github.com/pingcap/tidb/pull/23301)  |   | 7 | 1d23h |


</details> 

