|   REVIEWER    |  O(C)   | 1 | 2 | 3 | 4 | 5 | 6 | 7 | T  |
|---------------|---------|---|---|---|---|---|---|---|----|
| Reminiscent   |  6 ( 0) | 0 | 3 | 1 | 1 | 0 | 2 | 4 | 11 |
| SunRunAway    | 16 ( 1) | 0 | 0 | 2 | 0 | 0 | 0 | 0 |  2 |
| XuHuaiyu      | 71 ( 2) | 1 | 0 | 1 | 0 | 0 | 2 | 2 |  6 |
| dyzsr         |  0 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| eurekaka      | 24 ( 1) | 0 | 1 | 0 | 0 | 0 | 4 | 1 |  6 |
| fzhedu        |  1 ( 0) | 0 | 0 | 0 | 0 | 0 | 3 | 1 |  4 |
| ichn-hu       |  3 ( 0) | 0 | 0 | 0 | 0 | 0 | 1 | 1 |  2 |
| lzmhhh123     | 50 ( 0) | 1 | 3 | 1 | 0 | 0 | 8 | 1 | 14 |
| nullnotnil    |  0 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| qw4990        | 63 ( 1) | 4 | 2 | 0 | 0 | 0 | 4 | 1 | 11 |
| rebelice      |  2 ( 0) | 1 | 0 | 0 | 0 | 0 | 1 | 2 |  4 |
| time-and-fate |  2 ( 0) | 2 | 0 | 0 | 0 | 0 | 0 | 0 |  2 |
| winoros       | 31 ( 2) | 1 | 1 | 2 | 0 | 0 | 2 | 2 |  8 |
| wshwsh12      | 19 ( 1) | 0 | 1 | 1 | 0 | 0 | 0 | 0 |  2 |


<details> 
  <summary>Reminiscent's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                     PR                                                                      | C | LASTED |
|------------|---------------------------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/21896 | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                                  |   | 86d19h |
| tidb/22354 | [planner: do not cache prepared plan if optimization depends on mutable constant (#22349)](https://github.com/pingcap/tidb/pull/22354)      |   | 64d23h |
| tidb/23074 | [planner: fix range partition prune bug for IN expr (#22894) (#22938)](https://github.com/pingcap/tidb/pull/23074)                          |   | 14d17h |
| tidb/23283 | [util: optimize the performance of restore with db (#22910)](https://github.com/pingcap/tidb/pull/23283)                                    |   | 5d17h  |
| tidb/23293 | [planner: fix the bug that wrong collation is used when try fast path for enum or set (#23217)](https://github.com/pingcap/tidb/pull/23293) |   | 5d14h  |
| tidb/23349 | [*: hide `index-usage-sync-lease` config](https://github.com/pingcap/tidb/pull/23349)                                                       |   | 1d17h  |


## Reviewed in Last 7 Days

|    REPO    |                                                                      PR                                                                      | C | D |   R    |
|------------|----------------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/23333 | [statistics: hide the `tidb_analyze_version` before it's GA (#23327)](https://github.com/pingcap/tidb/pull/23333)                            |   | 2 | 1h     |
| tidb/23328 | [planner: hide the switch of dynamic-pruning and global-stats (#23320)](https://github.com/pingcap/tidb/pull/23328)                          |   | 2 | 3h     |
| tidb/23327 | [statistics: hide the `tidb_analyze_version` before it's GA](https://github.com/pingcap/tidb/pull/23327)                                     |   | 2 | 6h     |
| tidb/23320 | [planner: hide the switch of dynamic-pruning and global-stats](https://github.com/pingcap/tidb/pull/23320)                                   |   | 3 | 0h     |
| tidb/23292 | [planner: fix the bug that wrong collation is used when try fast path for enum or set (#23217)](https://github.com/pingcap/tidb/pull/23292)  |   | 4 | 2d14h  |
| tidb/23217 | [planner: fix the bug that wrong collation is used when try fast path for enum or set](https://github.com/pingcap/tidb/pull/23217)           |   | 6 | 2d20h  |
| tidb/23265 | [statistics: add multi-column index test cases for global-statas](https://github.com/pingcap/tidb/pull/23265)                                |   | 6 | 17h    |
| tidb/22910 | [util: optimize the performance of restore with db](https://github.com/pingcap/tidb/pull/22910)                                              |   | 7 | 15d0h  |
| tidb/23219 | [statistics: fix a case that auto-analyze is triggered outside its time range (#23214)](https://github.com/pingcap/tidb/pull/23219)          |   | 7 | 1d16h  |
| tidb/23231 | [statistics: add a test case which builds global-stats on different versions of partition-stats](https://github.com/pingcap/tidb/pull/23231) |   | 7 | 22h    |
| tidb/22914 | [partition: fix hash partition with not between condition get wrong result](https://github.com/pingcap/tidb/pull/22914)                      |   | 7 | 14d18h |


</details> 


<details> 
  <summary>SunRunAway's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                  PR                                                                   | C | LASTED  |
|------------|---------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19178 | [executor: Refactor probe channel](https://github.com/pingcap/tidb/pull/19178)                                                        |   | 216d16h |
| tidb/19347 | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347)  |   | 208d23h |
| tidb/19807 | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                               |   | 194d11h |
| tidb/19900 | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                        | Y | 189d18h |
| tidb/20140 | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                    |   | 176d22h |
| tidb/20765 | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                     |   | 135d17h |
| tidb/21207 | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                       |   | 114d19h |
| tidb/21834 | [planner: enhanced index range calculation plan](https://github.com/pingcap/tidb/pull/21834)                                          |   | 91d18h  |
| tidb/21876 | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)           |   | 89d19h  |
| tidb/21878 | [planner: do not push down lock to pointGet/bacthPointGet when selection exists](https://github.com/pingcap/tidb/pull/21878)          |   | 89d18h  |
| tidb/21956 | [planner/preprocessor: disallow into-outfile clause in some place](https://github.com/pingcap/tidb/pull/21956)                        |   | 84d23h  |
| tidb/22026 | [expression: separated arithmeticPlusIntSig](https://github.com/pingcap/tidb/pull/22026)                                              |   | 82d20h  |
| tidb/22114 | [test: fix globalkilltest (#21987)](https://github.com/pingcap/tidb/pull/22114)                                                       |   | 77d12h  |
| tidb/22217 | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                         |   | 70d17h  |
| tidb/22365 | [planner: check index valid while forUpdateRead (#22152)](https://github.com/pingcap/tidb/pull/22365)                                 |   | 64d19h  |
| tidb/22379 | [[experiment] executor: allow aggregation to spill disk when running out of memory quota](https://github.com/pingcap/tidb/pull/22379) |   | 63d19h  |


## Reviewed in Last 7 Days

|    REPO    |                                             PR                                              | C | D |   R   |
|------------|---------------------------------------------------------------------------------------------|---|---|-------|
| tidb/23224 | [docs: Add Proposal for dynamic privileges](https://github.com/pingcap/tidb/pull/23224)     |   | 3 | 5d15h |
| tidb/23223 | [docs: add proposal for Security Enhanced Mode](https://github.com/pingcap/tidb/pull/23223) |   | 3 | 5d7h  |


</details> 


<details> 
  <summary>XuHuaiyu's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                              PR                                                                               | C | LASTED  |
|--------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| docs-cn/5619 | [Update data-type-date-and-time.md](https://github.com/pingcap/docs-cn/pull/5619)                                                                             |   | 19d16h  |
| docs-cn/5671 | [tidb: Add time format description](https://github.com/pingcap/docs-cn/pull/5671)                                                                             |   | 13d11h  |
| tidb/19900   | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                                                | Y | 189d18h |
| docs-cn/5754 | [add document for explain/explain analyze of mpp query.](https://github.com/pingcap/docs-cn/pull/5754)                                                        |   | 21h     |
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                                        | Y | 187d14h |
| tidb/20140   | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                                            |   | 176d22h |
| tidb/20311   | [expression: fix overflow error when convert bit to int64 (#20266)](https://github.com/pingcap/tidb/pull/20311)                                               |   | 168d21h |
| tidb/20790   | [collation: add pinyin collation for chinese charset support](https://github.com/pingcap/tidb/pull/20790)                                                     |   | 134d20h |
| tidb/20905   | [planner: fix statement-optimize not work in `TryFastPlan`](https://github.com/pingcap/tidb/pull/20905)                                                       |   | 131d17h |
| tidb/20972   | [expression: POC implementation of Vitess hashing algorithm.](https://github.com/pingcap/tidb/pull/20972)                                                     |   | 127d1h  |
| tidb/21064   | [planner, executor: fix cast not check error](https://github.com/pingcap/tidb/pull/21064)                                                                     |   | 122d8h  |
| tidb/21149   | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                                 |   | 118d14h |
| tidb/21228   | [executor: return the result immediately when combining LIMIT row_count with DISTINCT](https://github.com/pingcap/tidb/pull/21228)                            |   | 114d13h |
| tidb/21304   | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                                      |   | 112d12h |
| tidb/21334   | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                                                 |   | 111d14h |
| tidb/21401   | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                                            |   | 107d11h |
| tidb/21476   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                         |   | 104d15h |
| tidb/21536   | [executor: add slow-log file meta cache to avoid repeat read file meta information](https://github.com/pingcap/tidb/pull/21536)                               |   | 100d15h |
| tidb/21564   | [ddl: fix Incorrect behavior of NO_ZERO_DATE when altering table](https://github.com/pingcap/tidb/pull/21564)                                                 |   | 99d15h  |
| tidb/21853   | [expression: fix compatibility behaviors in time_format with MySQL (#21559)](https://github.com/pingcap/tidb/pull/21853)                                      |   | 90d19h  |
| tidb/21896   | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                                                    |   | 86d19h  |
| tidb/22131   | [privilege: remove leading and trailing space when create user and role](https://github.com/pingcap/tidb/pull/22131)                                          |   | 76d19h  |
| tidb/22149   | [session: set process info before building plan (#22101)](https://github.com/pingcap/tidb/pull/22149)                                                         |   | 72d19h  |
| tidb/22163   | [expression: separated arithmeticMinusIntSig](https://github.com/pingcap/tidb/pull/22163)                                                                     |   | 72d13h  |
| tidb/22186   | [executor: fix select into outfile with year type column has no data (#22175)](https://github.com/pingcap/tidb/pull/22186)                                    |   | 71d16h  |
| tidb/22294   | [planner, table: optimize the list partition pruner for range query](https://github.com/pingcap/tidb/pull/22294)                                              |   | 68d20h  |
| tidb/22307   | [ddl: fix update can see columns not public](https://github.com/pingcap/tidb/pull/22307)                                                                      |   | 68d16h  |
| tidb/22381   | [planner: check schema stale for plan cache when forUpdateRead](https://github.com/pingcap/tidb/pull/22381)                                                   |   | 63d14h  |
| tidb/22616   | [expression: from_unixtime accept 64-bit integers](https://github.com/pingcap/tidb/pull/22616)                                                                |   | 47d23h  |
| tidb/22617   | [metrics: fix wrong bucket name of coprocessor cache (#22454)](https://github.com/pingcap/tidb/pull/22617)                                                    |   | 47d23h  |
| tidb/22624   | [ planner: not pruning column used by union scan condition (#21640)](https://github.com/pingcap/tidb/pull/22624)                                              |   | 47d17h  |
| tidb/22631   | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                                                                               |   | 45d23h  |
| tidb/22696   | [expression: enable arithmetic Mod push down](https://github.com/pingcap/tidb/pull/22696)                                                                     |   | 42d17h  |
| tidb/22711   | [executor: Fix inline schema name](https://github.com/pingcap/tidb/pull/22711)                                                                                |   | 42d11h  |
| tidb/22722   | [planner, errno: make error code of ErrMixOfGroupFuncAndFields consistent with MySQL](https://github.com/pingcap/tidb/pull/22722)                             |   | 41d20h  |
| tidb/22814   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22814)                                           |   | 26d19h  |
| tidb/22815   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22815)                                           |   | 26d19h  |
| tidb/22844   | [expression: do not adjust int when it is null and compared year (#22821)](https://github.com/pingcap/tidb/pull/22844)                                        |   | 25d19h  |
| tidb/22914   | [partition: fix hash partition with not between condition get wrong result](https://github.com/pingcap/tidb/pull/22914)                                       |   | 21d18h  |
| tidb/22926   | [expression: add overflow check in multiplyInt](https://github.com/pingcap/tidb/pull/22926)                                                                   |   | 21d13h  |
| tidb/23012   | [executor: fix affected rows of ddls and complete uint tests](https://github.com/pingcap/tidb/pull/23012)                                                     |   | 17d16h  |
| tidb/23104   | [executor: fix wrong key range of index scan when filter is comparing year column with NULL (#23079)](https://github.com/pingcap/tidb/pull/23104)             |   | 13d18h  |
| tidb/23105   | [executor: fix wrong key range of index scan when filter is comparing year column with NULL (#23079)](https://github.com/pingcap/tidb/pull/23105)             |   | 13d18h  |
| tidb/23111   | [executor: fix linter --enable=deadcode check error in executor(#22979)](https://github.com/pingcap/tidb/pull/23111)                                          |   | 13d17h  |
| tidb/23128   | [statistics: refactor the statistics package use the RestrictedSQLExecutor API (#22636)](https://github.com/pingcap/tidb/pull/23128)                          |   | 12d22h  |
| tidb/23152   | [expression: fix wrong error info (#22760)](https://github.com/pingcap/tidb/pull/23152)                                                                       |   | 10d14h  |
| tidb/23191   | [planner/core: convert decimal type for mpp join before shuffling.](https://github.com/pingcap/tidb/pull/23191)                                               |   | 8d19h   |
| tidb/23196   | [types: fix the bug about the wrong query result for decimal type  (#22507)](https://github.com/pingcap/tidb/pull/23196)                                      |   | 8d18h   |
| tidb/23210   | [planner: fixed a bug that prevented SPM from taking effect (#23197)](https://github.com/pingcap/tidb/pull/23210)                                             |   | 8d16h   |
| tidb/23220   | [Release 4.0](https://github.com/pingcap/tidb/pull/23220)                                                                                                     |   | 8d11h   |
| tidb/23227   | [executor: hash join out of index panic when enum column value is zero (#23162)](https://github.com/pingcap/tidb/pull/23227)                                  |   | 7d22h   |
| tidb/23233   | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23233)                                                        |   | 7d18h   |
| tidb/23234   | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23234)                                                        |   | 7d18h   |
| tidb/23245   | [*: Add security enhanced mode as experimental](https://github.com/pingcap/tidb/pull/23245)                                                                   |   | 7d6h    |
| tidb/23257   | [executor: group_concat aggr panic when session.group_concat_max_len is small (#23131)](https://github.com/pingcap/tidb/pull/23257)                           |   | 6d18h   |
| tidb/23268   | [*: add infoschema client errors (#22382)](https://github.com/pingcap/tidb/pull/23268)                                                                        |   | 6d15h   |
| tidb/23278   | [executor: wrong result of nullif expr when used with is null expr. (#23170)](https://github.com/pingcap/tidb/pull/23278)                                     |   | 5d18h   |
| tidb/23295   | [util, types: don't let SPM be affected by charset (#23161)](https://github.com/pingcap/tidb/pull/23295)                                                      |   | 5d11h   |
| tidb/23332   | [excutor: fix the date precision of `builtinCastDurationAsStringSig.vecEvalString` #23314 #23286](https://github.com/pingcap/tidb/pull/23332)                 |   | 1d20h   |
| tidb/23335   | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23335)                                      |   | 1d19h   |
| tidb/23336   | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23336)                                      |   | 1d19h   |
| tidb/23337   | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23337)                                      |   | 1d19h   |
| tidb/23340   | [executor: fix unexpected NotNullFlag in case when expr ret type (#23102)](https://github.com/pingcap/tidb/pull/23340)                                        |   | 1d19h   |
| tidb/23347   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                                          |   | 1d18h   |
| tidb/23348   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                                          |   | 1d18h   |
| tidb/23350   | [util/stringutil, util/ranger, planner: use hierarchical separators to simplify the parsing for info of EXPLAIN ](https://github.com/pingcap/tidb/pull/23350) |   | 1d17h   |
| tidb/23361   | [*: add clustered index info in `show index from` stmt (#23329)](https://github.com/pingcap/tidb/pull/23361)                                                  |   | 1d14h   |
| tidb/23368   | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23368)                                         |   | 20h     |
| tidb/23369   | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23369)                                         |   | 20h     |
| tidb/23370   | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23370)                                         |   | 20h     |
| tidb/23374   | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23374)                                             |   | 19h     |


## Reviewed in Last 7 Days

|     REPO     |                                                       PR                                                       | C | D |   R    |
|--------------|----------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/23339   | [expression: fix refine compare constant](https://github.com/pingcap/tidb/pull/23339)                          |   | 1 | 20h    |
| docs-cn/5620 | [Add details for Hexadecimal Literals](https://github.com/pingcap/docs-cn/pull/5620)                           |   | 3 | 16d22h |
| tidb/23161   | [util, types: don't let SPM be affected by charset](https://github.com/pingcap/tidb/pull/23161)                |   | 6 | 4d7h   |
| tidb/23184   | [*: hide the system variables `tidb_track_aggregate_memory_usage`](https://github.com/pingcap/tidb/pull/23184) |   | 6 | 3d3h   |
| docs-cn/5704 | [remove system variable `tidb_track_aggregate_memory_usage`](https://github.com/pingcap/docs-cn/pull/5704)     |   | 7 | 23h    |
| tidb/23260   | [test: fix global kill e2e test](https://github.com/pingcap/tidb/pull/23260)                                   |   | 7 | 0h     |


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
| tidb/19347 | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347)   |   | 208d23h |
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                              |   | 132d17h |
| tidb/21444 | [planner: ignore anonymous index while tiflash replica is available](https://github.com/pingcap/tidb/pull/21444)                       |   | 105d12h |
| tidb/21994 | [range: fix overflow value access index ](https://github.com/pingcap/tidb/pull/21994)                                                  |   | 83d23h  |
| tidb/22342 | [session: fix two cases when updating bind info (#22338)](https://github.com/pingcap/tidb/pull/22342)                                  |   | 65d18h  |
| tidb/22354 | [planner: do not cache prepared plan if optimization depends on mutable constant (#22349)](https://github.com/pingcap/tidb/pull/22354) |   | 64d23h  |
| tidb/22369 | [session: fix the duplicate binding case when updating bind info (#22367)](https://github.com/pingcap/tidb/pull/22369)                 |   | 64d17h  |
| tidb/22416 | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                                      | Y | 60d11h  |
| tidb/22559 | [planner: split test data from test cases in cbo_test.go](https://github.com/pingcap/tidb/pull/22559)                                  |   | 49d19h  |
| tidb/22778 | [*: add support for dynamic privileges](https://github.com/pingcap/tidb/pull/22778)                                                    |   | 29d7h   |
| tidb/23137 | [planner: fix index merge row count estimation logic](https://github.com/pingcap/tidb/pull/23137)                                      |   | 12d17h  |
| tidb/23208 | [statistics, util/ranger: improve selectivity calculation for DNF filters (#18741)](https://github.com/pingcap/tidb/pull/23208)        |   | 8d16h   |
| tidb/23210 | [planner: fixed a bug that prevented SPM from taking effect (#23197)](https://github.com/pingcap/tidb/pull/23210)                      |   | 8d16h   |
| tidb/23216 | [Privileges: fix delete privilege check wrongly (#22971)](https://github.com/pingcap/tidb/pull/23216)                                  |   | 8d14h   |
| tidb/23270 | [*: deprecate alter-primary-key configuration](https://github.com/pingcap/tidb/pull/23270)                                             |   | 6d12h   |
| tidb/23283 | [util: optimize the performance of restore with db (#22910)](https://github.com/pingcap/tidb/pull/23283)                               |   | 5d17h   |
| tidb/23295 | [util, types: don't let SPM be affected by charset (#23161)](https://github.com/pingcap/tidb/pull/23295)                               |   | 5d11h   |
| tidb/23316 | [planner: Fix rebuild range for prepared plan](https://github.com/pingcap/tidb/pull/23316)                                             |   | 2d17h   |
| tidb/23346 | [statistics: fix build extended stats panic when met NULL](https://github.com/pingcap/tidb/pull/23346)                                 |   | 1d18h   |
| tidb/23349 | [*: hide `index-usage-sync-lease` config](https://github.com/pingcap/tidb/pull/23349)                                                  |   | 1d17h   |
| tidb/23365 | [planner: fix a bug that point get plan returns wrong column name](https://github.com/pingcap/tidb/pull/23365)                         |   | 22h     |
| tidb/23372 | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23372)                      |   | 19h     |
| tidb/23373 | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                      |   | 19h     |
| tidb/23374 | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23374)                      |   | 19h     |


## Reviewed in Last 7 Days

|    REPO    |                                                         PR                                                          | C | D |   R    |
|------------|---------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/22853 | [planner: fix LogicalPlans that contain Window Function are ambiguous ](https://github.com/pingcap/tidb/pull/22853) |   | 2 | 23d17h |
| tidb/23175 | [planner: fix plan cache not working caused by type difference](https://github.com/pingcap/tidb/pull/23175)         |   | 6 | 3d21h  |
| tidb/23161 | [util, types: don't let SPM be affected by charset](https://github.com/pingcap/tidb/pull/23161)                     |   | 6 | 4d1h   |
| tidb/23241 | [executor: fix get var expr when session var is hex literal](https://github.com/pingcap/tidb/pull/23241)            |   | 6 | 1d21h  |
| tidb/22910 | [util: optimize the performance of restore with db](https://github.com/pingcap/tidb/pull/22910)                     |   | 6 | 16d0h  |
| tidb/23209 | [planner: fixed a bug that prevented SPM from taking effect (#23197)](https://github.com/pingcap/tidb/pull/23209)   |   | 7 | 2d1h   |


</details> 


<details> 
  <summary>fzhedu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                         PR                                                          | C | LASTED |
|------------|---------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/22853 | [planner: fix LogicalPlans that contain Window Function are ambiguous ](https://github.com/pingcap/tidb/pull/22853) |   | 25d12h |


## Reviewed in Last 7 Days

|    REPO    |                                                                     PR                                                                      | C | D |   R   |
|------------|---------------------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/23191 | [planner/core: convert decimal type for mpp join before shuffling.](https://github.com/pingcap/tidb/pull/23191)                             |   | 6 | 3d4h  |
| tidb/23259 | [planner: not push down mpp/join to TiFlash in some cases that TiFlash not supported the query](https://github.com/pingcap/tidb/pull/23259) |   | 6 | 1d1h  |
| tidb/23203 | [planner/core: pass the elems info to fieldtype only for exchanger](https://github.com/pingcap/tidb/pull/23203)                             |   | 6 | 3d0h  |
| tidb/22867 | [expression, planner: allow pushdown count distinct when enumerate physical plans](https://github.com/pingcap/tidb/pull/22867)              |   | 7 | 17d1h |


</details> 


<details> 
  <summary>ichn-hu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                            PR                                                             | C | LASTED |
|------------|---------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/21853 | [expression: fix compatibility behaviors in time_format with MySQL (#21559)](https://github.com/pingcap/tidb/pull/21853)  |   | 90d19h |
| tidb/23278 | [executor: wrong result of nullif expr when used with is null expr. (#23170)](https://github.com/pingcap/tidb/pull/23278) |   | 5d18h  |
| tidb/23279 | [executor: wrong result of nullif expr when used with is null expr. (#23170)](https://github.com/pingcap/tidb/pull/23279) |   | 5d18h  |


## Reviewed in Last 7 Days

|    REPO    |                                                       PR                                                       | C | D |   R   |
|------------|----------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/23184 | [*: hide the system variables `tidb_track_aggregate_memory_usage`](https://github.com/pingcap/tidb/pull/23184) |   | 6 | 3d3h  |
| tidb/23056 | [MPP: Kill mpp queries](https://github.com/pingcap/tidb/pull/23056)                                            |   | 7 | 8d23h |


</details> 


<details> 
  <summary>lzmhhh123's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                             PR                                                                              | C | LASTED  |
|------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19347 | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347)                        |   | 208d23h |
| tidb/20444 | [expression: add json_merge_patch](https://github.com/pingcap/tidb/pull/20444)                                                                              |   | 154d21h |
| tidb/20465 | [expression: add uuidShortFunction](https://github.com/pingcap/tidb/pull/20465)                                                                             |   | 153d19h |
| tidb/20642 | [executor: modify admin executors to support partitioned table with global index](https://github.com/pingcap/tidb/pull/20642)                               |   | 142d15h |
| tidb/20903 | [planner: fix confused and unnecessary double-projection in plans.](https://github.com/pingcap/tidb/pull/20903)                                             |   | 131d17h |
| tidb/21018 | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                              |   | 125d17h |
| tidb/21195 | [brie: integrate lightning to suport IMPORT statement](https://github.com/pingcap/tidb/pull/21195)                                                          |   | 114d22h |
| tidb/21334 | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                                               |   | 111d14h |
| tidb/21347 | [session: make rollback work on global variables](https://github.com/pingcap/tidb/pull/21347)                                                               |   | 110d19h |
| tidb/21444 | [planner: ignore anonymous index while tiflash replica is available](https://github.com/pingcap/tidb/pull/21444)                                            |   | 105d12h |
| tidb/21487 | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                                               |   | 104d4h  |
| tidb/21641 | [executor: Fix pessimistic lock doesn't work on the partition table for subquery/joins](https://github.com/pingcap/tidb/pull/21641)                         |   | 97d18h  |
| tidb/21651 | [planner: allow filter condition pushing down to IndexScan for prefix index](https://github.com/pingcap/tidb/pull/21651)                                    |   | 97d13h  |
| tidb/22126 | [*: add `sys` schema, `sys.SCHEMA_UNUSED_INDEXES` view and `sys.SCHEMA_INDEX_USAGE` view](https://github.com/pingcap/tidb/pull/22126)                       |   | 76d19h  |
| tidb/22149 | [session: set process info before building plan (#22101)](https://github.com/pingcap/tidb/pull/22149)                                                       |   | 72d19h  |
| tidb/22188 | [planner: do not use indexMerge when the path only use a single index (#22168)](https://github.com/pingcap/tidb/pull/22188)                                 |   | 71d13h  |
| tidb/22361 | [table: fix insert into _tidb_rowid panic and rebase it if needed (#22062)](https://github.com/pingcap/tidb/pull/22361)                                     |   | 64d20h  |
| tidb/22372 | [executor: fix SelectForUpdate in decorrelated subquery under pessimistic mode](https://github.com/pingcap/tidb/pull/22372)                                 |   | 64d9h   |
| tidb/22478 | [planner, executor: fix query partition table with global unique index get wrong result](https://github.com/pingcap/tidb/pull/22478)                        |   | 55d13h  |
| tidb/22631 | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                                                                             |   | 45d23h  |
| tidb/22699 | [brie: add error info column and history backup/restore info in sql](https://github.com/pingcap/tidb/pull/22699)                                            |   | 42d16h  |
| tidb/22857 | [mocktikv: split rpcHandler to kvHandler and coprHandler](https://github.com/pingcap/tidb/pull/22857)                                                       |   | 24d21h  |
| tidb/22926 | [expression: add overflow check in multiplyInt](https://github.com/pingcap/tidb/pull/22926)                                                                 |   | 21d13h  |
| tidb/23001 | [statistics: fix err check](https://github.com/pingcap/tidb/pull/23001)                                                                                     |   | 18d0h   |
| tidb/23022 | [executor: create PipelinedWindowExec based on current implementation and modify the windowProcessor interface](https://github.com/pingcap/tidb/pull/23022) |   | 16d18h  |
| tidb/23040 | [ddl: add truncate partition all support](https://github.com/pingcap/tidb/pull/23040)                                                                       |   | 16d13h  |
| tidb/23149 | [core: support left join and right join for join reorder](https://github.com/pingcap/tidb/pull/23149)                                                       |   | 11d12h  |
| tidb/23210 | [planner: fixed a bug that prevented SPM from taking effect (#23197)](https://github.com/pingcap/tidb/pull/23210)                                           |   | 8d16h   |
| tidb/23237 | [plan: setting not null flag for extrak pk](https://github.com/pingcap/tidb/pull/23237)                                                                     |   | 7d17h   |
| tidb/23257 | [executor: group_concat aggr panic when session.group_concat_max_len is small (#23131)](https://github.com/pingcap/tidb/pull/23257)                         |   | 6d18h   |
| tidb/23278 | [executor: wrong result of nullif expr when used with is null expr. (#23170)](https://github.com/pingcap/tidb/pull/23278)                                   |   | 5d18h   |
| tidb/23279 | [executor: wrong result of nullif expr when used with is null expr. (#23170)](https://github.com/pingcap/tidb/pull/23279)                                   |   | 5d18h   |
| tidb/23283 | [util: optimize the performance of restore with db (#22910)](https://github.com/pingcap/tidb/pull/23283)                                                    |   | 5d17h   |
| tidb/23292 | [planner: fix the bug that wrong collation is used when try fast path for enum or set (#23217)](https://github.com/pingcap/tidb/pull/23292)                 |   | 5d14h   |
| tidb/23293 | [planner: fix the bug that wrong collation is used when try fast path for enum or set (#23217)](https://github.com/pingcap/tidb/pull/23293)                 |   | 5d14h   |
| tidb/23296 | [sig/execution: fix the bug that Wrong result of comparison operation(type date / type string)](https://github.com/pingcap/tidb/pull/23296)                 |   | 5d7h    |
| tidb/23307 | [util/chunk: replace outdated link with correct one](https://github.com/pingcap/tidb/pull/23307)                                                            |   | 2d20h   |
| tidb/23324 | [test: add testleak after checking and testing](https://github.com/pingcap/tidb/pull/23324)                                                                 |   | 2d13h   |
| tidb/23334 | [metrics/grafana: Remove duplicate items "Owner Watcher OPS"](https://github.com/pingcap/tidb/pull/23334)                                                   |   | 1d20h   |
| tidb/23337 | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23337)                                    |   | 1d19h   |
| tidb/23340 | [executor: fix unexpected NotNullFlag in case when expr ret type (#23102)](https://github.com/pingcap/tidb/pull/23340)                                      |   | 1d19h   |
| tidb/23347 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                                        |   | 1d18h   |
| tidb/23348 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                                        |   | 1d18h   |
| tidb/23361 | [*: add clustered index info in `show index from` stmt (#23329)](https://github.com/pingcap/tidb/pull/23361)                                                |   | 1d14h   |
| tidb/23368 | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23368)                                       |   | 20h     |
| tidb/23369 | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23369)                                       |   | 20h     |
| tidb/23370 | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23370)                                       |   | 20h     |
| tidb/23372 | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23372)                                           |   | 19h     |
| tidb/23373 | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                           |   | 19h     |
| tidb/23374 | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23374)                                           |   | 19h     |


## Reviewed in Last 7 Days

|      REPO      |                                                                       PR                                                                       | C | D |   R    |
|----------------|------------------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/23241     | [executor: fix get var expr when session var is hex literal](https://github.com/pingcap/tidb/pull/23241)                                       |   | 1 | 6d17h  |
| tidb/23356     | [planner: fix {index,hash,merge} join on range suffix condition clustered index (#23294)](https://github.com/pingcap/tidb/pull/23356)          |   | 2 | 0h     |
| tidb-test/1168 | [Add cast target type in explain result.](https://github.com/pingcap/tidb-test/pull/1168)                                                      |   | 2 | 0h     |
| tidb/23285     | [executor, expression: fix the incorrect result of AVG function](https://github.com/pingcap/tidb/pull/23285)                                   |   | 2 | 3d20h  |
| tidb/23294     | [planner: fix {index,hash,merge} join on range suffix condition clustered index](https://github.com/pingcap/tidb/pull/23294)                   |   | 3 | 2d13h  |
| tidb/23288     | [test: work around goroutine leak](https://github.com/pingcap/tidb/pull/23288)                                                                 |   | 6 | 0h     |
| tidb/22832     | [expression: push down EXTRACT to TiFlash](https://github.com/pingcap/tidb/pull/22832)                                                         |   | 6 | 20d9h  |
| tidb/23203     | [planner/core: pass the elems info to fieldtype only for exchanger](https://github.com/pingcap/tidb/pull/23203)                                |   | 6 | 3d0h   |
| tidb/23284     | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified.](https://github.com/pingcap/tidb/pull/23284) |   | 6 | 0h     |
| tidb-test/1149 | [mysql_test: add test with timezone change](https://github.com/pingcap/tidb-test/pull/1149)                                                    |   | 6 | 66d19h |
| tidb/23237     | [plan: setting not null flag for extrak pk](https://github.com/pingcap/tidb/pull/23237)                                                        |   | 6 | 1d19h  |
| tidb/23172     | [planner: set right null flag for constant value](https://github.com/pingcap/tidb/pull/23172)                                                  |   | 6 | 3d19h  |
| tidb/23276     | [Revert "executor: open childExec during execution for UnionExec (#21561)](https://github.com/pingcap/tidb/pull/23276)                         |   | 6 | 0h     |
| tikv/9787      | [copr: fix a bug that get_index_version doesn't consider some cases.](https://github.com/tikv/tikv/pull/9787)                                  |   | 7 | 0h     |


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
| docs-cn/5484 | [system variable: add tidb_allow_fallback_to_tikv](https://github.com/pingcap/docs-cn/pull/5484)                                                            |   | 42d17h  |
| docs/4781    | [system variable: add tidb_allow_fallback_to_tikv](https://github.com/pingcap/docs/pull/4781)                                                               |   | 42d17h  |
| tidb/19029   | [types: fix unexpected NOT_NULL flags](https://github.com/pingcap/tidb/pull/19029)                                                                          |   | 223d22h |
| docs-cn/5561 | [Add sql optimization-related docs to toc](https://github.com/pingcap/docs-cn/pull/5561)                                                                    |   | 23d15h  |
| tidb/20708   | [*: separate auto_increment ID allocator from _tidb_rowid allocator](https://github.com/pingcap/tidb/pull/20708)                                            |   | 139d20h |
| tidb/20969   | [executor: Improve the performance of appending not fixed columns](https://github.com/pingcap/tidb/pull/20969)                                              |   | 127d9h  |
| tidb/20972   | [expression: POC implementation of Vitess hashing algorithm.](https://github.com/pingcap/tidb/pull/20972)                                                   |   | 127d1h  |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                              |   | 125d17h |
| tidb/21149   | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                               |   | 118d14h |
| tidb/21304   | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                                    |   | 112d12h |
| tidb/21318   | [planner, expression: use the range of column types to simplify expressions](https://github.com/pingcap/tidb/pull/21318)                                    |   | 111d19h |
| tidb/21401   | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                                          |   | 107d11h |
| tidb/21476   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                       |   | 104d15h |
| tidb/21508   | [execution: fix dayofweek('0000-00-00') behavior](https://github.com/pingcap/tidb/pull/21508)                                                               |   | 103d10h |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                                 |   | 89d19h  |
| tidb/21887   | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                                             |   | 88d11h  |
| tidb/21954   | [planner/cascades: add rule `PushSelDownApply`](https://github.com/pingcap/tidb/pull/21954)                                                                 |   | 84d23h  |
| tidb/22146   | [executor: forbid SFU on view](https://github.com/pingcap/tidb/pull/22146)                                                                                  |   | 72d21h  |
| tidb/22217   | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                                               |   | 70d17h  |
| tidb/22234   | [executor, planner: ON DUPLICATE UPDATE can refer to un-project col (#14412)](https://github.com/pingcap/tidb/pull/22234)                                   |   | 70d15h  |
| tidb/22261   | [time: fix parse datetime won't truncate the reluctant string (#22232)](https://github.com/pingcap/tidb/pull/22261)                                         |   | 69d19h  |
| tidb/22294   | [planner, table: optimize the list partition pruner for range query](https://github.com/pingcap/tidb/pull/22294)                                            |   | 68d20h  |
| tidb/22307   | [ddl: fix update can see columns not public](https://github.com/pingcap/tidb/pull/22307)                                                                    |   | 68d16h  |
| tidb/22342   | [session: fix two cases when updating bind info (#22338)](https://github.com/pingcap/tidb/pull/22342)                                                       |   | 65d18h  |
| tidb/22369   | [session: fix the duplicate binding case when updating bind info (#22367)](https://github.com/pingcap/tidb/pull/22369)                                      |   | 64d17h  |
| tidb/22374   | [expression: separated arithmeticIntDivideSig](https://github.com/pingcap/tidb/pull/22374)                                                                  |   | 64d0h   |
| tidb/22415   | [ddl: refactor placement package](https://github.com/pingcap/tidb/pull/22415)                                                                               |   | 60d17h  |
| tidb/22541   | [expression: Support builtin function SOUNDEX](https://github.com/pingcap/tidb/pull/22541)                                                                  |   | 50d9h   |
| tidb/22559   | [planner: split test data from test cases in cbo_test.go](https://github.com/pingcap/tidb/pull/22559)                                                       |   | 49d19h  |
| tidb/22565   | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)                                          | Y | 49d17h  |
| tidb/22778   | [*: add support for dynamic privileges](https://github.com/pingcap/tidb/pull/22778)                                                                         |   | 29d7h   |
| tidb/22814   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22814)                                         |   | 26d19h  |
| tidb/22815   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22815)                                         |   | 26d19h  |
| tidb/22862   | [brie: fix the problem that ddl restored by BR via SQL is not replicated to downstream](https://github.com/pingcap/tidb/pull/22862)                         |   | 23d22h  |
| tidb/22915   | [planner: build correct MaxOneRow info from multi-column conditions](https://github.com/pingcap/tidb/pull/22915)                                            |   | 21d18h  |
| tidb/22923   | [expression: correct constant propagation for collation (#22666)](https://github.com/pingcap/tidb/pull/22923)                                               |   | 21d15h  |
| tidb/22924   | [planner: fix wrong index merge selection (#22825)](https://github.com/pingcap/tidb/pull/22924)                                                             |   | 21d14h  |
| tidb/22926   | [expression: add overflow check in multiplyInt](https://github.com/pingcap/tidb/pull/22926)                                                                 |   | 21d13h  |
| tidb/22984   | [executor: fix logging format of prepared statements (#16062)](https://github.com/pingcap/tidb/pull/22984)                                                  |   | 18d10h  |
| tidb/23022   | [executor: create PipelinedWindowExec based on current implementation and modify the windowProcessor interface](https://github.com/pingcap/tidb/pull/23022) |   | 16d18h  |
| tidb/23062   | [*: fix structcheck lint warnings](https://github.com/pingcap/tidb/pull/23062)                                                                              |   | 14d19h  |
| tidb/23074   | [planner: fix range partition prune bug for IN expr (#22894) (#22938)](https://github.com/pingcap/tidb/pull/23074)                                          |   | 14d17h  |
| tidb/23105   | [executor: fix wrong key range of index scan when filter is comparing year column with NULL (#23079)](https://github.com/pingcap/tidb/pull/23105)           |   | 13d18h  |
| tidb/23119   | [statistics: remove existing deleted extended stats when add a new one](https://github.com/pingcap/tidb/pull/23119)                                         |   | 13d14h  |
| tidb/23137   | [planner: fix index merge row count estimation logic](https://github.com/pingcap/tidb/pull/23137)                                                           |   | 12d17h  |
| tidb/23152   | [expression: fix wrong error info (#22760)](https://github.com/pingcap/tidb/pull/23152)                                                                     |   | 10d14h  |
| tidb/23171   | [store/tikv:move option from kv to tikv, and make define as a normal int](https://github.com/pingcap/tidb/pull/23171)                                       |   | 9d17h   |
| tidb/23196   | [types: fix the bug about the wrong query result for decimal type  (#22507)](https://github.com/pingcap/tidb/pull/23196)                                    |   | 8d18h   |
| tidb/23201   | [executor, server: load_data.go is changed and add unit test](https://github.com/pingcap/tidb/pull/23201)                                                   |   | 8d17h   |
| tidb/23208   | [statistics, util/ranger: improve selectivity calculation for DNF filters (#18741)](https://github.com/pingcap/tidb/pull/23208)                             |   | 8d16h   |
| tidb/23210   | [planner: fixed a bug that prevented SPM from taking effect (#23197)](https://github.com/pingcap/tidb/pull/23210)                                           |   | 8d16h   |
| tidb/23234   | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23234)                                                      |   | 7d18h   |
| tidb/23283   | [util: optimize the performance of restore with db (#22910)](https://github.com/pingcap/tidb/pull/23283)                                                    |   | 5d17h   |
| tidb/23284   | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified.](https://github.com/pingcap/tidb/pull/23284)              |   | 5d17h   |
| tidb/23292   | [planner: fix the bug that wrong collation is used when try fast path for enum or set (#23217)](https://github.com/pingcap/tidb/pull/23292)                 |   | 5d14h   |
| tidb/23293   | [planner: fix the bug that wrong collation is used when try fast path for enum or set (#23217)](https://github.com/pingcap/tidb/pull/23293)                 |   | 5d14h   |
| tidb/23295   | [util, types: don't let SPM be affected by charset (#23161)](https://github.com/pingcap/tidb/pull/23295)                                                    |   | 5d11h   |
| tidb/23316   | [planner: Fix rebuild range for prepared plan](https://github.com/pingcap/tidb/pull/23316)                                                                  |   | 2d17h   |
| tidb/23339   | [expression: fix refine compare constant](https://github.com/pingcap/tidb/pull/23339)                                                                       |   | 1d19h   |
| tidb/23345   | [sessionctx: hide extended stats variable in SHOW VARIABLES temporarily](https://github.com/pingcap/tidb/pull/23345)                                        |   | 1d18h   |
| tidb/23349   | [*: hide `index-usage-sync-lease` config](https://github.com/pingcap/tidb/pull/23349)                                                                       |   | 1d17h   |
| tidb/23361   | [*: add clustered index info in `show index from` stmt (#23329)](https://github.com/pingcap/tidb/pull/23361)                                                |   | 1d14h   |
| tidb/23373   | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                           |   | 19h     |


## Reviewed in Last 7 Days

|     REPO     |                                                                 PR                                                                 | C | D |   R   |
|--------------|------------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/23343   | [statistics: handle drop/trunctate partition events for global-stats](https://github.com/pingcap/tidb/pull/23343)                  |   | 1 | 23h   |
| tidb/22090   | [planner: push aggregation operators down to projection by default](https://github.com/pingcap/tidb/pull/22090)                    |   | 1 | 77d3h |
| docs-cn/5740 | [update SPM documentation for DML SQL Bind and baseline capture](https://github.com/pingcap/docs-cn/pull/5740)                     |   | 1 | 1d20h |
| tidb/23238   | [planner: fix wrong PointGet / TableDual plan reused in plan cache](https://github.com/pingcap/tidb/pull/23238)                    |   | 1 | 6d22h |
| tidb/23333   | [statistics: hide the `tidb_analyze_version` before it's GA (#23327)](https://github.com/pingcap/tidb/pull/23333)                  |   | 2 | 1h    |
| tidb/23327   | [statistics: hide the `tidb_analyze_version` before it's GA](https://github.com/pingcap/tidb/pull/23327)                           |   | 2 | 8h    |
| tidb/23161   | [util, types: don't let SPM be affected by charset](https://github.com/pingcap/tidb/pull/23161)                                    |   | 6 | 4d1h  |
| tidb/23217   | [planner: fix the bug that wrong collation is used when try fast path for enum or set](https://github.com/pingcap/tidb/pull/23217) |   | 6 | 2d20h |
| tidb/23256   | [statistics: add tests for global-stats when add and delete a single partition](https://github.com/pingcap/tidb/pull/23256)        |   | 6 | 23h   |
| tidb/23088   | [statistics: delete extended stats cache item in current tidb synchronously](https://github.com/pingcap/tidb/pull/23088)           |   | 6 | 8d14h |
| tidb/23240   | [config: disable prepare plan cache by default](https://github.com/pingcap/tidb/pull/23240)                                        |   | 7 | 21h   |


</details> 


<details> 
  <summary>rebelice's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                         PR                                                         | C | LASTED |
|------------|--------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/23074 | [planner: fix range partition prune bug for IN expr (#22894) (#22938)](https://github.com/pingcap/tidb/pull/23074) |   | 14d17h |
| tidb/23343 | [statistics: handle drop/trunctate partition events for global-stats](https://github.com/pingcap/tidb/pull/23343)  |   | 1d18h  |


## Reviewed in Last 7 Days

|    REPO    |                                                                 PR                                                                  | C | D |  R  |
|------------|-------------------------------------------------------------------------------------------------------------------------------------|---|---|-----|
| tidb/23354 | [planner: remove some risky cache operations in the plan builder](https://github.com/pingcap/tidb/pull/23354)                       |   | 1 | 22h |
| tidb/23256 | [statistics: add tests for global-stats when add and delete a single partition](https://github.com/pingcap/tidb/pull/23256)         |   | 6 | 23h |
| tidb/23265 | [statistics: add multi-column index test cases for global-statas](https://github.com/pingcap/tidb/pull/23265)                       |   | 7 | 2h  |
| tidb/23242 | [planner: fix a panic caused by unmatched FieldNames and ColsInfo in partition pruning](https://github.com/pingcap/tidb/pull/23242) |   | 7 | 18h |


</details> 


<details> 
  <summary>time-and-fate's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                         PR                                                          | C | LASTED  |
|------------|---------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                           |   | 132d17h |
| tidb/22853 | [planner: fix LogicalPlans that contain Window Function are ambiguous ](https://github.com/pingcap/tidb/pull/22853) |   | 25d12h  |


## Reviewed in Last 7 Days

|    REPO    |                                                         PR                                                          | C | D |   R    |
|------------|---------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/23238 | [planner: fix wrong PointGet / TableDual plan reused in plan cache](https://github.com/pingcap/tidb/pull/23238)     |   | 1 | 6d22h  |
| tidb/23119 | [statistics: remove existing deleted extended stats when add a new one](https://github.com/pingcap/tidb/pull/23119) |   | 1 | 12d19h |


</details> 


<details> 
  <summary>winoros's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                              PR                                                                               | C | LASTED  |
|--------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                                        | Y | 187d14h |
| docs-cn/5484 | [system variable: add tidb_allow_fallback_to_tikv](https://github.com/pingcap/docs-cn/pull/5484)                                                              |   | 42d17h  |
| docs/4781    | [system variable: add tidb_allow_fallback_to_tikv](https://github.com/pingcap/docs/pull/4781)                                                                 |   | 42d17h  |
| tidb/20311   | [expression: fix overflow error when convert bit to int64 (#20266)](https://github.com/pingcap/tidb/pull/20311)                                               |   | 168d21h |
| tidb/20765   | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                                             |   | 135d17h |
| tidb/20877   | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                                                     |   | 132d17h |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                                |   | 125d17h |
| tidb/21207   | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                                               |   | 114d19h |
| tidb/21476   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                         |   | 104d15h |
| tidb/21487   | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                                                 |   | 104d4h  |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                                   |   | 89d19h  |
| tidb/21954   | [planner/cascades: add rule `PushSelDownApply`](https://github.com/pingcap/tidb/pull/21954)                                                                   |   | 84d23h  |
| tidb/22181   | [planner, expression: fix error when using IN combined with subquery (#22080)](https://github.com/pingcap/tidb/pull/22181)                                    |   | 71d17h  |
| tidb/22365   | [planner: check index valid while forUpdateRead (#22152)](https://github.com/pingcap/tidb/pull/22365)                                                         |   | 64d19h  |
| tidb/22504   | [*:Fix the fetchHotRegion bug that the count always zero](https://github.com/pingcap/tidb/pull/22504)                                                         |   | 52d19h  |
| tidb/22565   | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)                                            | Y | 49d17h  |
| tidb/22624   | [ planner: not pruning column used by union scan condition (#21640)](https://github.com/pingcap/tidb/pull/22624)                                              |   | 47d17h  |
| tidb/22923   | [expression: correct constant propagation for collation (#22666)](https://github.com/pingcap/tidb/pull/22923)                                                 |   | 21d15h  |
| tidb/23163   | [plugin: fix linter --enable=deadcode check error](https://github.com/pingcap/tidb/pull/23163)                                                                |   | 9d19h   |
| tidb/23208   | [statistics, util/ranger: improve selectivity calculation for DNF filters (#18741)](https://github.com/pingcap/tidb/pull/23208)                               |   | 8d16h   |
| tidb/23215   | [Privileges: fix delete privilege check wrongly (#22971)](https://github.com/pingcap/tidb/pull/23215)                                                         |   | 8d14h   |
| tidb/23216   | [Privileges: fix delete privilege check wrongly (#22971)](https://github.com/pingcap/tidb/pull/23216)                                                         |   | 8d14h   |
| tidb/23233   | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23233)                                                        |   | 7d18h   |
| tidb/23234   | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23234)                                                        |   | 7d18h   |
| tidb/23246   | [planner: fix the panic in joinReOrderSolver.optimizeRecursive](https://github.com/pingcap/tidb/pull/23246)                                                   |   | 6d23h   |
| tidb/23345   | [sessionctx: hide extended stats variable in SHOW VARIABLES temporarily](https://github.com/pingcap/tidb/pull/23345)                                          |   | 1d18h   |
| tidb/23346   | [statistics: fix build extended stats panic when met NULL](https://github.com/pingcap/tidb/pull/23346)                                                        |   | 1d18h   |
| tidb/23347   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                                          |   | 1d18h   |
| tidb/23348   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                                          |   | 1d18h   |
| tidb/23350   | [util/stringutil, util/ranger, planner: use hierarchical separators to simplify the parsing for info of EXPLAIN ](https://github.com/pingcap/tidb/pull/23350) |   | 1d17h   |
| tidb/23373   | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                             |   | 19h     |


## Reviewed in Last 7 Days

|    REPO    |                                                                 PR                                                                  | C | D |   R   |
|------------|-------------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/23354 | [planner: remove some risky cache operations in the plan builder](https://github.com/pingcap/tidb/pull/23354)                       |   | 1 | 22h   |
| tidb/23328 | [planner: hide the switch of dynamic-pruning and global-stats (#23320)](https://github.com/pingcap/tidb/pull/23328)                 |   | 2 | 3h    |
| tidb/23320 | [planner: hide the switch of dynamic-pruning and global-stats](https://github.com/pingcap/tidb/pull/23320)                          |   | 3 | 9h    |
| tidb/23316 | [planner: Fix rebuild range for prepared plan](https://github.com/pingcap/tidb/pull/23316)                                          |   | 3 | 7h    |
| tidb/23288 | [test: work around goroutine leak](https://github.com/pingcap/tidb/pull/23288)                                                      |   | 6 | 0h    |
| tidb/23092 | [*: fix a bug that collation is not handle for text type (#23045)](https://github.com/pingcap/tidb/pull/23092)                      |   | 6 | 8d17h |
| tidb/23248 | [server: add tiflash fallback testcase](https://github.com/pingcap/tidb/pull/23248)                                                 |   | 7 | 3h    |
| tidb/23219 | [statistics: fix a case that auto-analyze is triggered outside its time range (#23214)](https://github.com/pingcap/tidb/pull/23219) |   | 7 | 1d16h |


</details> 


<details> 
  <summary>wshwsh12's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                        PR                                                                         | C | LASTED  |
|------------|---------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19807 | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                                           |   | 194d11h |
| tidb/19957 | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                            | Y | 187d14h |
| tidb/21487 | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                                     |   | 104d4h  |
| tidb/21887 | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                                   |   | 88d11h  |
| tidb/22378 | [executor: vectorize hash aggregate](https://github.com/pingcap/tidb/pull/22378)                                                                  |   | 63d19h  |
| tidb/22628 | [executor: Improve max/min window function with deque-based sliding window](https://github.com/pingcap/tidb/pull/22628)                           |   | 46d23h  |
| tidb/22815 | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22815)                               |   | 26d19h  |
| tidb/23104 | [executor: fix wrong key range of index scan when filter is comparing year column with NULL (#23079)](https://github.com/pingcap/tidb/pull/23104) |   | 13d18h  |
| tidb/23105 | [executor: fix wrong key range of index scan when filter is comparing year column with NULL (#23079)](https://github.com/pingcap/tidb/pull/23105) |   | 13d18h  |
| tidb/23128 | [statistics: refactor the statistics package use the RestrictedSQLExecutor API (#22636)](https://github.com/pingcap/tidb/pull/23128)              |   | 12d22h  |
| tidb/23234 | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23234)                                            |   | 7d18h   |
| tidb/23323 | [wip:test](https://github.com/pingcap/tidb/pull/23323)                                                                                            |   | 2d14h   |
| tidb/23336 | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23336)                          |   | 1d19h   |
| tidb/23347 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                              |   | 1d18h   |
| tidb/23348 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                              |   | 1d18h   |
| tidb/23368 | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23368)                             |   | 20h     |
| tidb/23369 | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23369)                             |   | 20h     |
| tidb/23370 | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23370)                             |   | 20h     |
| tidb/23376 | [store/tikv:mv kv.ReplicaReadType to tikv](https://github.com/pingcap/tidb/pull/23376)                                                            |   | 17h     |


## Reviewed in Last 7 Days

|    REPO    |                                          PR                                          | C | D |   R   |
|------------|--------------------------------------------------------------------------------------|---|---|-------|
| tidb/23301 | [expression: fix unused code in util.go](https://github.com/pingcap/tidb/pull/23301) |   | 2 | 1d23h |
| tidb/23260 | [test: fix global kill e2e test](https://github.com/pingcap/tidb/pull/23260)         |   | 3 | 4d1h  |


</details> 

