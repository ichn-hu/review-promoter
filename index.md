|   REVIEWER    |  O(C)   | 1 | 2 | 3 | 4 | 5 | 6 | 7 | T  |
|---------------|---------|---|---|---|---|---|---|---|----|
| Reminiscent   |  4 ( 0) | 4 | 4 | 3 | 3 | 0 | 0 | 2 | 16 |
| SunRunAway    | 19 ( 1) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| XuHuaiyu      | 68 ( 3) | 3 | 3 | 3 | 1 | 0 | 0 | 1 | 11 |
| dyzsr         |  0 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| eurekaka      | 20 ( 1) | 1 | 0 | 1 | 1 | 0 | 0 | 0 |  3 |
| fzhedu        |  1 ( 0) | 1 | 1 | 0 | 1 | 0 | 0 | 0 |  3 |
| ichn-hu       |  2 ( 0) | 2 | 0 | 2 | 0 | 0 | 0 | 0 |  4 |
| lzmhhh123     | 36 ( 0) | 1 | 2 | 2 | 5 | 0 | 0 | 0 | 10 |
| nullnotnil    |  0 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| qw4990        | 61 ( 1) | 2 | 4 | 0 | 6 | 0 | 0 | 2 | 14 |
| rebelice      |  2 ( 0) | 2 | 3 | 0 | 2 | 0 | 0 | 3 | 10 |
| time-and-fate |  3 ( 0) | 1 | 1 | 0 | 2 | 0 | 0 | 1 |  5 |
| winoros       | 30 ( 2) | 2 | 0 | 4 | 0 | 0 | 0 | 2 |  8 |
| wshwsh12      | 22 ( 1) | 0 | 1 | 1 | 0 | 0 | 0 | 0 |  2 |


<details> 
  <summary>Reminiscent's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                   PR                                                                   | C | LASTED |
|------------|----------------------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/21896 | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                             |   | 80d19h |
| tidb/22354 | [planner: do not cache prepared plan if optimization depends on mutable constant (#22349)](https://github.com/pingcap/tidb/pull/22354) |   | 58d23h |
| tidb/23074 | [planner: fix range partition prune bug for IN expr (#22894) (#22938)](https://github.com/pingcap/tidb/pull/23074)                     |   | 8d17h  |
| tidb/23265 | [statistics: add multi-column index test cases for global-statas](https://github.com/pingcap/tidb/pull/23265)                          |   | 16h    |


## Reviewed in Last 7 Days

|    REPO    |                                                                       PR                                                                        | C | D |   R    |
|------------|-------------------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/22910 | [util: optimize the performance of restore with db](https://github.com/pingcap/tidb/pull/22910)                                                 |   | 1 | 15d0h  |
| tidb/23219 | [statistics: fix a case that auto-analyze is triggered outside its time range (#23214)](https://github.com/pingcap/tidb/pull/23219)             |   | 1 | 1d16h  |
| tidb/23231 | [statistics: add a test case which builds global-stats on different versions of partition-stats](https://github.com/pingcap/tidb/pull/23231)    |   | 1 | 22h    |
| tidb/22914 | [partition: fix hash partition with not between condition get wrong result](https://github.com/pingcap/tidb/pull/22914)                         |   | 1 | 14d18h |
| tidb/23242 | [planner: fix a panic caused by unmatched FieldNames and ColsInfo in partition pruning](https://github.com/pingcap/tidb/pull/23242)             |   | 2 | 14h    |
| tidb/23240 | [config: disable prepare plan cache by default](https://github.com/pingcap/tidb/pull/23240)                                                     |   | 2 | 16h    |
| tidb/23226 | [server: remove unstable tiflash fallback testcase](https://github.com/pingcap/tidb/pull/23226)                                                 |   | 2 | 4h     |
| tidb/23228 | [statistics: update the count and modify variables of global-stats as well when dumping delta info](https://github.com/pingcap/tidb/pull/23228) |   | 2 | 1h     |
| tidb/23217 | [planner: fix the bug that wrong collation is used when try fast path for enum or set](https://github.com/pingcap/tidb/pull/23217)              |   | 3 | 14h    |
| tidb/23214 | [statistics: fix a case that auto-analyze is triggered outside its time range](https://github.com/pingcap/tidb/pull/23214)                      |   | 3 | 0h     |
| tidb/23176 | [statistics: test `tidb_partition_prune_mode` session variable](https://github.com/pingcap/tidb/pull/23176)                                     |   | 3 | 18h    |
| tidb/23169 | [statistics: add more test cases for global-stats with different analysis options](https://github.com/pingcap/tidb/pull/23169)                  |   | 4 | 2h     |
| tidb/23167 | [statistics: add test for analyze single partition](https://github.com/pingcap/tidb/pull/23167)                                                 |   | 4 | 2h     |
| tidb/23138 | [statistics: add some test cases of global-stats to cover more column types](https://github.com/pingcap/tidb/pull/23138)                        |   | 4 | 2d18h  |
| tidb/23113 | [statistics: add tests for `analyze` with dynamic partition prune mode](https://github.com/pingcap/tidb/pull/23113)                             |   | 7 | 22h    |
| tidb/23099 | [statistics: add more test cases about global-stats and fix some issues](https://github.com/pingcap/tidb/pull/23099)                            |   | 7 | 23h    |


</details> 


<details> 
  <summary>SunRunAway's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                  PR                                                                   | C | LASTED  |
|------------|---------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19178 | [executor: Refactor probe channel](https://github.com/pingcap/tidb/pull/19178)                                                        |   | 210d16h |
| tidb/19347 | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347)  |   | 202d23h |
| tidb/19807 | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                               |   | 188d11h |
| tidb/19900 | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                        | Y | 183d18h |
| tidb/20140 | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                    |   | 170d22h |
| tidb/20765 | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                     |   | 129d17h |
| tidb/21207 | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                       |   | 108d19h |
| tidb/21277 | [executor: fix split table with large integers](https://github.com/pingcap/tidb/pull/21277)                                           |   | 106d20h |
| tidb/21834 | [planner: enhanced index range calculation plan](https://github.com/pingcap/tidb/pull/21834)                                          |   | 85d18h  |
| tidb/21876 | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)           |   | 83d19h  |
| tidb/21878 | [planner: do not push down lock to pointGet/bacthPointGet when selection exists](https://github.com/pingcap/tidb/pull/21878)          |   | 83d18h  |
| tidb/21956 | [planner/preprocessor: disallow into-outfile clause in some place](https://github.com/pingcap/tidb/pull/21956)                        |   | 78d23h  |
| tidb/22026 | [expression: separated arithmeticPlusIntSig](https://github.com/pingcap/tidb/pull/22026)                                              |   | 76d20h  |
| tidb/22114 | [test: fix globalkilltest (#21987)](https://github.com/pingcap/tidb/pull/22114)                                                       |   | 71d12h  |
| tidb/22217 | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                         |   | 64d17h  |
| tidb/22365 | [planner: check index valid while forUpdateRead (#22152)](https://github.com/pingcap/tidb/pull/22365)                                 |   | 58d19h  |
| tidb/22379 | [[experiment] executor: allow aggregation to spill disk when running out of memory quota](https://github.com/pingcap/tidb/pull/22379) |   | 57d19h  |
| tidb/23223 | [docs: add proposal for Security Enhanced Mode](https://github.com/pingcap/tidb/pull/23223)                                           |   | 2d1h    |
| tidb/23224 | [docs: Add Proposal for dynamic privileges](https://github.com/pingcap/tidb/pull/23224)                                               |   | 2d0h    |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>XuHuaiyu's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                              PR                                                                              | C | LASTED  |
|--------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| docs-cn/5619 | [Update data-type-date-and-time.md](https://github.com/pingcap/docs-cn/pull/5619)                                                                            |   | 13d16h  |
| docs-cn/5620 | [Add details for Hexadecimal Literals](https://github.com/pingcap/docs-cn/pull/5620)                                                                         |   | 13d15h  |
| tidb/19900   | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                                               | Y | 183d18h |
| docs-cn/5671 | [tidb: Add time format description](https://github.com/pingcap/docs-cn/pull/5671)                                                                            |   | 7d11h   |
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                                       | Y | 181d14h |
| tidb/20040   | [planner, expression: take NullFlag into consideration when optimize the `int non-const` <cmp > `non-int const`](https://github.com/pingcap/tidb/pull/20040) | Y | 176d14h |
| tidb/20140   | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                                           |   | 170d22h |
| tidb/20311   | [expression: fix overflow error when convert bit to int64 (#20266)](https://github.com/pingcap/tidb/pull/20311)                                              |   | 162d21h |
| tidb/20790   | [collation: add pinyin collation for chinese charset support](https://github.com/pingcap/tidb/pull/20790)                                                    |   | 128d20h |
| tidb/20905   | [planner: fix statement-optimize not work in `TryFastPlan`](https://github.com/pingcap/tidb/pull/20905)                                                      |   | 125d17h |
| tidb/20972   | [expression: POC implementation of Vitess hashing algorithm.](https://github.com/pingcap/tidb/pull/20972)                                                    |   | 121d1h  |
| tidb/21064   | [planner, executor: fix cast not check error](https://github.com/pingcap/tidb/pull/21064)                                                                    |   | 116d8h  |
| tidb/21149   | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                                |   | 112d14h |
| tidb/21228   | [executor: return the result immediately when combining LIMIT row_count with DISTINCT](https://github.com/pingcap/tidb/pull/21228)                           |   | 108d13h |
| tidb/21304   | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                                     |   | 106d12h |
| tidb/21334   | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                                                |   | 105d14h |
| tidb/21476   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                        |   | 98d15h  |
| tidb/21536   | [executor: add slow-log file meta cache to avoid repeat read file meta information](https://github.com/pingcap/tidb/pull/21536)                              |   | 94d15h  |
| tidb/21564   | [ddl: fix Incorrect behavior of NO_ZERO_DATE when altering table](https://github.com/pingcap/tidb/pull/21564)                                                |   | 93d15h  |
| tidb/21853   | [expression: fix compatibility behaviors in time_format with MySQL (#21559)](https://github.com/pingcap/tidb/pull/21853)                                     |   | 84d19h  |
| tidb/21896   | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                                                   |   | 80d19h  |
| tidb/22131   | [privilege: remove leading and trailing space when create user and role](https://github.com/pingcap/tidb/pull/22131)                                         |   | 70d19h  |
| tidb/22149   | [session: set process info before building plan (#22101)](https://github.com/pingcap/tidb/pull/22149)                                                        |   | 66d19h  |
| tidb/22163   | [expression: separated arithmeticMinusIntSig](https://github.com/pingcap/tidb/pull/22163)                                                                    |   | 66d13h  |
| tidb/22186   | [executor: fix select into outfile with year type column has no data (#22175)](https://github.com/pingcap/tidb/pull/22186)                                   |   | 65d16h  |
| tidb/22294   | [planner, table: optimize the list partition pruner for range query](https://github.com/pingcap/tidb/pull/22294)                                             |   | 62d20h  |
| tidb/22307   | [ddl: fix update can see columns not public](https://github.com/pingcap/tidb/pull/22307)                                                                     |   | 62d16h  |
| tidb/22381   | [planner: check schema stale for plan cache when forUpdateRead](https://github.com/pingcap/tidb/pull/22381)                                                  |   | 57d14h  |
| tidb/22616   | [expression: from_unixtime accept 64-bit integers](https://github.com/pingcap/tidb/pull/22616)                                                               |   | 41d23h  |
| tidb/22617   | [metrics: fix wrong bucket name of coprocessor cache (#22454)](https://github.com/pingcap/tidb/pull/22617)                                                   |   | 41d23h  |
| tidb/22624   | [ planner: not pruning column used by union scan condition (#21640)](https://github.com/pingcap/tidb/pull/22624)                                             |   | 41d17h  |
| tidb/22631   | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                                                                              |   | 39d23h  |
| tidb/22696   | [expression: enable arithmetic Mod push down](https://github.com/pingcap/tidb/pull/22696)                                                                    |   | 36d17h  |
| tidb/22711   | [executor: Fix inline schema name](https://github.com/pingcap/tidb/pull/22711)                                                                               |   | 36d11h  |
| tidb/22722   | [planner, errno: make error code of ErrMixOfGroupFuncAndFields consistent with MySQL](https://github.com/pingcap/tidb/pull/22722)                            |   | 35d20h  |
| tidb/22814   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22814)                                          |   | 20d19h  |
| tidb/22815   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22815)                                          |   | 20d19h  |
| tidb/22832   | [expression: push down EXTRACT to TiFlash](https://github.com/pingcap/tidb/pull/22832)                                                                       |   | 20d1h   |
| tidb/22844   | [expression: do not adjust int when it is null and compared year (#22821)](https://github.com/pingcap/tidb/pull/22844)                                       |   | 19d19h  |
| tidb/22886   | [*: rename tiflash fallback switch](https://github.com/pingcap/tidb/pull/22886)                                                                              |   | 16d20h  |
| tidb/22914   | [partition: fix hash partition with not between condition get wrong result](https://github.com/pingcap/tidb/pull/22914)                                      |   | 15d18h  |
| tidb/23012   | [executor: fix affected rows of ddls and complete uint tests](https://github.com/pingcap/tidb/pull/23012)                                                    |   | 11d16h  |
| tidb/23056   | [MPP: Kill mpp queries](https://github.com/pingcap/tidb/pull/23056)                                                                                          |   | 9d12h   |
| tidb/23092   | [*: fix a bug that collation is not handle for text type (#23045)](https://github.com/pingcap/tidb/pull/23092)                                               |   | 8d12h   |
| tidb/23104   | [executor: fix wrong key range of index scan when filter is comparing year column with NULL (#23079)](https://github.com/pingcap/tidb/pull/23104)            |   | 7d18h   |
| tidb/23105   | [executor: fix wrong key range of index scan when filter is comparing year column with NULL (#23079)](https://github.com/pingcap/tidb/pull/23105)            |   | 7d18h   |
| tidb/23111   | [executor: fix linter --enable=deadcode check error in executor(#22979)](https://github.com/pingcap/tidb/pull/23111)                                         |   | 7d17h   |
| tidb/23123   | [planner: show cast type in EXPLAIN in coptask](https://github.com/pingcap/tidb/pull/23123)                                                                  |   | 7d13h   |
| tidb/23128   | [statistics: refactor the statistics package use the RestrictedSQLExecutor API (#22636)](https://github.com/pingcap/tidb/pull/23128)                         |   | 6d22h   |
| tidb/23135   | [executor: fix unexpected NotNullFlag in case when expr ret type (#23102)](https://github.com/pingcap/tidb/pull/23135)                                       |   | 6d17h   |
| tidb/23152   | [expression: fix wrong error info (#22760)](https://github.com/pingcap/tidb/pull/23152)                                                                      |   | 4d14h   |
| tidb/23161   | [util, types: don't let SPM be affected by charset](https://github.com/pingcap/tidb/pull/23161)                                                              |   | 3d19h   |
| tidb/23172   | [planner: set right null flag for constant value](https://github.com/pingcap/tidb/pull/23172)                                                                |   | 3d17h   |
| tidb/23184   | [*: hide the system variables `tidb_track_aggregate_memory_usage`](https://github.com/pingcap/tidb/pull/23184)                                               |   | 2d23h   |
| tidb/23191   | [planner/core: convert decimal type for mpp join before shuffling.](https://github.com/pingcap/tidb/pull/23191)                                              |   | 2d19h   |
| tidb/23196   | [types: fix the bug about the wrong query result for decimal type  (#22507)](https://github.com/pingcap/tidb/pull/23196)                                     |   | 2d18h   |
| tidb/23203   | [planner/core: pass the elems info to fieldtype only for exchanger](https://github.com/pingcap/tidb/pull/23203)                                              |   | 2d17h   |
| tidb/23209   | [planner: fixed a bug that prevented SPM from taking effect (#23197)](https://github.com/pingcap/tidb/pull/23209)                                            |   | 2d16h   |
| tidb/23210   | [planner: fixed a bug that prevented SPM from taking effect (#23197)](https://github.com/pingcap/tidb/pull/23210)                                            |   | 2d16h   |
| tidb/23211   | [executor: fix cast function will ignore tht error for point-get key construction (#22869)](https://github.com/pingcap/tidb/pull/23211)                      |   | 2d16h   |
| tidb/23220   | [Release 4.0](https://github.com/pingcap/tidb/pull/23220)                                                                                                    |   | 2d11h   |
| tidb/23227   | [executor: hash join out of index panic when enum column value is zero (#23162)](https://github.com/pingcap/tidb/pull/23227)                                 |   | 1d22h   |
| tidb/23233   | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23233)                                                       |   | 1d18h   |
| tidb/23234   | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23234)                                                       |   | 1d18h   |
| tidb/23245   | [*: Add security enhanced mode as experimental](https://github.com/pingcap/tidb/pull/23245)                                                                  |   | 1d6h    |
| tidb/23257   | [executor: group_concat aggr panic when session.group_concat_max_len is small (#23131)](https://github.com/pingcap/tidb/pull/23257)                          |   | 18h     |
| tidb/23259   | [planner: not push down mpp/join to TiFlash in some cases that TiFlash not supported the query](https://github.com/pingcap/tidb/pull/23259)                  |   | 18h     |
| tidb/23268   | [*: add infoschema client errors (#22382)](https://github.com/pingcap/tidb/pull/23268)                                                                       |   | 15h     |


## Reviewed in Last 7 Days

|     REPO     |                                                                  PR                                                                   | C | D |   R    |
|--------------|---------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| docs-cn/5704 | [remove system variable `tidb_track_aggregate_memory_usage`](https://github.com/pingcap/docs-cn/pull/5704)                            |   | 1 | 23h    |
| tidb/23260   | [test: fix global kill e2e test](https://github.com/pingcap/tidb/pull/23260)                                                          |   | 1 | 0h     |
| tidb/23184   | [*: hide the system variables `tidb_track_aggregate_memory_usage`](https://github.com/pingcap/tidb/pull/23184)                        |   | 1 | 2d5h   |
| tidb/22786   | [config: deprecate `tikv-client.copr-cache.enable` and invisible some copr-cache configs](https://github.com/pingcap/tidb/pull/22786) |   | 2 | 19d20h |
| docs-cn/5699 | [docs: remove some config fields for copr cache](https://github.com/pingcap/docs-cn/pull/5699)                                        |   | 2 | 0h     |
| tidb/22830   | [planner: fix incorrect duration between compare](https://github.com/pingcap/tidb/pull/22830)                                         |   | 2 | 18d10h |
| tidb/23197   | [planner: fixed a bug that prevented SPM from taking effect](https://github.com/pingcap/tidb/pull/23197)                              |   | 3 | 0h     |
| tidb/23139   | [executor: inject random panic to AggExec](https://github.com/pingcap/tidb/pull/23139)                                                |   | 3 | 3d21h  |
| tidb/23072   | [executor: track memory usage of map in agg partial result.](https://github.com/pingcap/tidb/pull/23072)                              |   | 3 | 5d19h  |
| tidb/23131   | [executor: group_concat aggr panic when session.group_concat_max_len is small](https://github.com/pingcap/tidb/pull/23131)            |   | 4 | 3d2h   |
| tidb/22869   | [executor: fix cast function will ignore tht error for point-get key construction](https://github.com/pingcap/tidb/pull/22869)        |   | 7 | 10d17h |


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
| tidb/19347 | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347)   |   | 202d23h |
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                              |   | 126d17h |
| tidb/21444 | [planner: ignore anonymous index while tiflash replica is available](https://github.com/pingcap/tidb/pull/21444)                       |   | 99d12h  |
| tidb/21994 | [range: fix overflow value access index ](https://github.com/pingcap/tidb/pull/21994)                                                  |   | 77d23h  |
| tidb/22342 | [session: fix two cases when updating bind info (#22338)](https://github.com/pingcap/tidb/pull/22342)                                  |   | 59d18h  |
| tidb/22354 | [planner: do not cache prepared plan if optimization depends on mutable constant (#22349)](https://github.com/pingcap/tidb/pull/22354) |   | 58d23h  |
| tidb/22369 | [session: fix the duplicate binding case when updating bind info (#22367)](https://github.com/pingcap/tidb/pull/22369)                 |   | 58d17h  |
| tidb/22416 | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                                      | Y | 54d11h  |
| tidb/22559 | [planner: split test data from test cases in cbo_test.go](https://github.com/pingcap/tidb/pull/22559)                                  |   | 43d19h  |
| tidb/22733 | [bindinfo: use new sql apis (#22653)](https://github.com/pingcap/tidb/pull/22733)                                                      |   | 35d15h  |
| tidb/22778 | [*: add support for dynamic privileges](https://github.com/pingcap/tidb/pull/22778)                                                    |   | 23d7h   |
| tidb/22853 | [planner: fix LogicalPlans that contain Window Function are ambiguous ](https://github.com/pingcap/tidb/pull/22853)                    |   | 19d12h  |
| tidb/22910 | [util: optimize the performance of restore with db](https://github.com/pingcap/tidb/pull/22910)                                        |   | 15d19h  |
| tidb/23161 | [util, types: don't let SPM be affected by charset](https://github.com/pingcap/tidb/pull/23161)                                        |   | 3d19h   |
| tidb/23175 | [planner: fix plan cache not working caused by type difference](https://github.com/pingcap/tidb/pull/23175)                            |   | 3d15h   |
| tidb/23208 | [statistics, util/ranger: improve selectivity calculation for DNF filters (#18741)](https://github.com/pingcap/tidb/pull/23208)        |   | 2d16h   |
| tidb/23210 | [planner: fixed a bug that prevented SPM from taking effect (#23197)](https://github.com/pingcap/tidb/pull/23210)                      |   | 2d16h   |
| tidb/23212 | [planner: fix query range partition table got wrong result and TiDB panic (#22953)](https://github.com/pingcap/tidb/pull/23212)        |   | 2d15h   |
| tidb/23216 | [Privileges: fix delete privilege check wrongly (#22971)](https://github.com/pingcap/tidb/pull/23216)                                  |   | 2d14h   |
| tidb/23265 | [statistics: add multi-column index test cases for global-statas](https://github.com/pingcap/tidb/pull/23265)                          |   | 16h     |


## Reviewed in Last 7 Days

|    REPO    |                                                        PR                                                         | C | D |   R    |
|------------|-------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/23209 | [planner: fixed a bug that prevented SPM from taking effect (#23197)](https://github.com/pingcap/tidb/pull/23209) |   | 1 | 2d1h   |
| tidb/23197 | [planner: fixed a bug that prevented SPM from taking effect](https://github.com/pingcap/tidb/pull/23197)          |   | 3 | 0h     |
| tidb/22924 | [planner: fix wrong index merge selection (#22825)](https://github.com/pingcap/tidb/pull/22924)                   |   | 4 | 11d17h |


</details> 


<details> 
  <summary>fzhedu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                         PR                                                          | C | LASTED |
|------------|---------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/22853 | [planner: fix LogicalPlans that contain Window Function are ambiguous ](https://github.com/pingcap/tidb/pull/22853) |   | 19d12h |


## Reviewed in Last 7 Days

|    REPO    |                                                               PR                                                               | C | D |   R   |
|------------|--------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/22867 | [expression, planner: allow pushdown count distinct when enumerate physical plans](https://github.com/pingcap/tidb/pull/22867) |   | 1 | 17d1h |
| tidb/23133 | [plan/core: support mpp group by expressions.](https://github.com/pingcap/tidb/pull/23133)                                     |   | 2 | 5d0h  |
| tics/1506  | [check block schema in exchange operator](https://github.com/pingcap/tics/pull/1506)                                           |   | 4 | 5d6h  |


</details> 


<details> 
  <summary>ichn-hu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                            PR                                                            | C | LASTED |
|------------|--------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/21853 | [expression: fix compatibility behaviors in time_format with MySQL (#21559)](https://github.com/pingcap/tidb/pull/21853) |   | 84d19h |
| tidb/23184 | [*: hide the system variables `tidb_track_aggregate_memory_usage`](https://github.com/pingcap/tidb/pull/23184)           |   | 2d23h  |


## Reviewed in Last 7 Days

|    REPO    |                                                        PR                                                        | C | D |   R   |
|------------|------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/23056 | [MPP: Kill mpp queries](https://github.com/pingcap/tidb/pull/23056)                                              |   | 1 | 8d23h |
| tidb/23184 | [*: hide the system variables `tidb_track_aggregate_memory_usage`](https://github.com/pingcap/tidb/pull/23184)   |   | 1 | 2d7h  |
| tidb/23170 | [executor: wrong result of nullif expr when used with is null expr.](https://github.com/pingcap/tidb/pull/23170) |   | 3 | 23h   |
| tidb/23139 | [executor: inject random panic to AggExec](https://github.com/pingcap/tidb/pull/23139)                           |   | 3 | 3d21h |


</details> 


<details> 
  <summary>lzmhhh123's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                             PR                                                                              | C | LASTED  |
|------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19347 | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347)                        |   | 202d23h |
| tidb/20444 | [expression: add json_merge_patch](https://github.com/pingcap/tidb/pull/20444)                                                                              |   | 148d21h |
| tidb/20465 | [expression: add uuidShortFunction](https://github.com/pingcap/tidb/pull/20465)                                                                             |   | 147d19h |
| tidb/20642 | [executor: modify admin executors to support partitioned table with global index](https://github.com/pingcap/tidb/pull/20642)                               |   | 136d15h |
| tidb/20903 | [planner: fix confused and unnecessary double-projection in plans.](https://github.com/pingcap/tidb/pull/20903)                                             |   | 125d17h |
| tidb/21018 | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                              |   | 119d17h |
| tidb/21195 | [brie: integrate lightning to suport IMPORT statement](https://github.com/pingcap/tidb/pull/21195)                                                          |   | 108d22h |
| tidb/21334 | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                                               |   | 105d14h |
| tidb/21347 | [session: make rollback work on global variables](https://github.com/pingcap/tidb/pull/21347)                                                               |   | 104d19h |
| tidb/21401 | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                                          |   | 101d11h |
| tidb/21444 | [planner: ignore anonymous index while tiflash replica is available](https://github.com/pingcap/tidb/pull/21444)                                            |   | 99d12h  |
| tidb/21487 | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                                               |   | 98d4h   |
| tidb/21641 | [executor: Fix pessimistic lock doesn't work on the partition table for subquery/joins](https://github.com/pingcap/tidb/pull/21641)                         |   | 91d18h  |
| tidb/21651 | [planner: allow filter condition pushing down to IndexScan for prefix index](https://github.com/pingcap/tidb/pull/21651)                                    |   | 91d13h  |
| tidb/22126 | [*: add `sys` schema, `sys.SCHEMA_UNUSED_INDEXES` view and `sys.SCHEMA_INDEX_USAGE` view](https://github.com/pingcap/tidb/pull/22126)                       |   | 70d19h  |
| tidb/22149 | [session: set process info before building plan (#22101)](https://github.com/pingcap/tidb/pull/22149)                                                       |   | 66d19h  |
| tidb/22188 | [planner: do not use indexMerge when the path only use a single index (#22168)](https://github.com/pingcap/tidb/pull/22188)                                 |   | 65d13h  |
| tidb/22361 | [table: fix insert into _tidb_rowid panic and rebase it if needed (#22062)](https://github.com/pingcap/tidb/pull/22361)                                     |   | 58d20h  |
| tidb/22372 | [executor: fix SelectForUpdate in decorrelated subquery under pessimistic mode](https://github.com/pingcap/tidb/pull/22372)                                 |   | 58d9h   |
| tidb/22478 | [planner, executor: fix query partition table with global unique index get wrong result](https://github.com/pingcap/tidb/pull/22478)                        |   | 49d13h  |
| tidb/22631 | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                                                                             |   | 39d23h  |
| tidb/22699 | [brie: add error info column and history backup/restore info in sql](https://github.com/pingcap/tidb/pull/22699)                                            |   | 36d16h  |
| tidb/22857 | [mocktikv: split rpcHandler to kvHandler and coprHandler](https://github.com/pingcap/tidb/pull/22857)                                                       |   | 18d21h  |
| tidb/22910 | [util: optimize the performance of restore with db](https://github.com/pingcap/tidb/pull/22910)                                                             |   | 15d19h  |
| tidb/23001 | [statistics: fix err check](https://github.com/pingcap/tidb/pull/23001)                                                                                     |   | 12d0h   |
| tidb/23022 | [executor: create PipelinedWindowExec based on current implementation and modify the windowProcessor interface](https://github.com/pingcap/tidb/pull/23022) |   | 10d18h  |
| tidb/23040 | [ddl: add truncate partition all support](https://github.com/pingcap/tidb/pull/23040)                                                                       |   | 10d13h  |
| tidb/23134 | [store/tikv:remove set/delete option from kv.Transaction](https://github.com/pingcap/tidb/pull/23134)                                                       |   | 6d18h   |
| tidb/23149 | [core: support left join and right join for join reorder](https://github.com/pingcap/tidb/pull/23149)                                                       |   | 5d12h   |
| tidb/23172 | [planner: set right null flag for constant value](https://github.com/pingcap/tidb/pull/23172)                                                               |   | 3d17h   |
| tidb/23210 | [planner: fixed a bug that prevented SPM from taking effect (#23197)](https://github.com/pingcap/tidb/pull/23210)                                           |   | 2d16h   |
| tidb/23217 | [planner: fix the bug that wrong collation is used when try fast path for enum or set](https://github.com/pingcap/tidb/pull/23217)                          |   | 2d14h   |
| tidb/23237 | [plan: setting not null flag for extrak pk](https://github.com/pingcap/tidb/pull/23237)                                                                     |   | 1d17h   |
| tidb/23241 | [executor: fix get var expr when session var is hex literal](https://github.com/pingcap/tidb/pull/23241)                                                    |   | 1d16h   |
| tidb/23256 | [statistics: add tests for global-stats when add and delete a single partition](https://github.com/pingcap/tidb/pull/23256)                                 |   | 19h     |
| tidb/23257 | [executor: group_concat aggr panic when session.group_concat_max_len is small (#23131)](https://github.com/pingcap/tidb/pull/23257)                         |   | 18h     |


## Reviewed in Last 7 Days

|    REPO    |                                                            PR                                                             | C | D |   R    |
|------------|---------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tikv/9787  | [copr: fix a bug that get_index_version doesn't consider some cases.](https://github.com/tikv/tikv/pull/9787)             |   | 1 | 0h     |
| tidb/23185 | [planner: fix prepared execute panic when sql without clustered pk condition](https://github.com/pingcap/tidb/pull/23185) |   | 2 | 1d4h   |
| tidb/23218 | [executor: truncate column values from index KV during admin check](https://github.com/pingcap/tidb/pull/23218)           |   | 2 | 17h    |
| tidb/23170 | [executor: wrong result of nullif expr when used with is null expr.](https://github.com/pingcap/tidb/pull/23170)          |   | 3 | 21h    |
| tidb/23164 | [planner: refine explain info for batch cop (#20360)](https://github.com/pingcap/tidb/pull/23164)                         |   | 3 | 22h    |
| tidb/23123 | [planner: show cast type in EXPLAIN in coptask](https://github.com/pingcap/tidb/pull/23123)                               |   | 4 | 3d17h  |
| tidb/20360 | [planner: refine explain info for batch cop](https://github.com/pingcap/tidb/pull/20360)                                  |   | 4 | 150d2h |
| tidb/23135 | [executor: fix unexpected NotNullFlag in case when expr ret type (#23102)](https://github.com/pingcap/tidb/pull/23135)    |   | 4 | 2d21h  |
| tidb/23094 | [planner: fix indexJoin(also hash, merge) on prefixed clustered index](https://github.com/pingcap/tidb/pull/23094)        |   | 4 | 4d13h  |
| tikv/9182  | [copr: fetch restore data while new collation is meet in `BatchTableScan`](https://github.com/tikv/tikv/pull/9182)        |   | 4 | 94d20h |


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
| docs-cn/5484 | [system variable: add tidb_enable_engine_fallback](https://github.com/pingcap/docs-cn/pull/5484)                                                            |   | 36d17h  |
| docs/4781    | [system variable: add tidb_enable_engine_fallback](https://github.com/pingcap/docs/pull/4781)                                                               |   | 36d17h  |
| tidb/19029   | [types: fix unexpected NOT_NULL flags](https://github.com/pingcap/tidb/pull/19029)                                                                          |   | 217d22h |
| docs-cn/5561 | [Add sql optimization-related docs to toc](https://github.com/pingcap/docs-cn/pull/5561)                                                                    |   | 17d15h  |
| tidb/20708   | [*: separate auto_increment ID allocator from _tidb_rowid allocator](https://github.com/pingcap/tidb/pull/20708)                                            |   | 133d20h |
| tidb/20969   | [executor: Improve the performance of appending not fixed columns](https://github.com/pingcap/tidb/pull/20969)                                              |   | 121d9h  |
| tidb/20972   | [expression: POC implementation of Vitess hashing algorithm.](https://github.com/pingcap/tidb/pull/20972)                                                   |   | 121d1h  |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                              |   | 119d17h |
| tidb/21149   | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                               |   | 112d14h |
| tidb/21304   | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                                    |   | 106d12h |
| tidb/21318   | [planner, expression: use the range of column types to simplify expressions](https://github.com/pingcap/tidb/pull/21318)                                    |   | 105d19h |
| tidb/21401   | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                                          |   | 101d11h |
| tidb/21476   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                       |   | 98d15h  |
| tidb/21508   | [execution: fix dayofweek('0000-00-00') behavior](https://github.com/pingcap/tidb/pull/21508)                                                               |   | 97d10h  |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                                 |   | 83d19h  |
| tidb/21887   | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                                             |   | 82d11h  |
| tidb/21930   | [planner: propagate NDV of column groups across plan nodes (#17854)](https://github.com/pingcap/tidb/pull/21930)                                            |   | 79d18h  |
| tidb/21954   | [planner/cascades: add rule `PushSelDownApply`](https://github.com/pingcap/tidb/pull/21954)                                                                 |   | 78d23h  |
| tidb/22090   | [planner: push aggregation operators down to projection and union by default](https://github.com/pingcap/tidb/pull/22090)                                   |   | 71d22h  |
| tidb/22146   | [executor: forbid SFU on view](https://github.com/pingcap/tidb/pull/22146)                                                                                  |   | 66d21h  |
| tidb/22217   | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                                               |   | 64d17h  |
| tidb/22234   | [executor, planner: ON DUPLICATE UPDATE can refer to un-project col (#14412)](https://github.com/pingcap/tidb/pull/22234)                                   |   | 64d15h  |
| tidb/22261   | [time: fix parse datetime won't truncate the reluctant string (#22232)](https://github.com/pingcap/tidb/pull/22261)                                         |   | 63d19h  |
| tidb/22294   | [planner, table: optimize the list partition pruner for range query](https://github.com/pingcap/tidb/pull/22294)                                            |   | 62d20h  |
| tidb/22307   | [ddl: fix update can see columns not public](https://github.com/pingcap/tidb/pull/22307)                                                                    |   | 62d16h  |
| tidb/22342   | [session: fix two cases when updating bind info (#22338)](https://github.com/pingcap/tidb/pull/22342)                                                       |   | 59d18h  |
| tidb/22369   | [session: fix the duplicate binding case when updating bind info (#22367)](https://github.com/pingcap/tidb/pull/22369)                                      |   | 58d17h  |
| tidb/22374   | [expression: separated arithmeticIntDivideSig](https://github.com/pingcap/tidb/pull/22374)                                                                  |   | 58d0h   |
| tidb/22415   | [ddl: refactor placement package](https://github.com/pingcap/tidb/pull/22415)                                                                               |   | 54d17h  |
| tidb/22541   | [expression: Support builtin function SOUNDEX](https://github.com/pingcap/tidb/pull/22541)                                                                  |   | 44d9h   |
| tidb/22559   | [planner: split test data from test cases in cbo_test.go](https://github.com/pingcap/tidb/pull/22559)                                                       |   | 43d19h  |
| tidb/22565   | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)                                          | Y | 43d17h  |
| tidb/22733   | [bindinfo: use new sql apis (#22653)](https://github.com/pingcap/tidb/pull/22733)                                                                           |   | 35d15h  |
| tidb/22778   | [*: add support for dynamic privileges](https://github.com/pingcap/tidb/pull/22778)                                                                         |   | 23d7h   |
| tidb/22814   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22814)                                         |   | 20d19h  |
| tidb/22815   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22815)                                         |   | 20d19h  |
| tidb/22862   | [brie: fix the problem that ddl restored by BR via SQL is not replicated to downstream](https://github.com/pingcap/tidb/pull/22862)                         |   | 17d22h  |
| tidb/22886   | [*: rename tiflash fallback switch](https://github.com/pingcap/tidb/pull/22886)                                                                             |   | 16d20h  |
| tidb/22915   | [planner: build correct MaxOneRow info from multi-column conditions](https://github.com/pingcap/tidb/pull/22915)                                            |   | 15d17h  |
| tidb/22923   | [expression: correct constant propagation for collation (#22666)](https://github.com/pingcap/tidb/pull/22923)                                               |   | 15d15h  |
| tidb/22924   | [planner: fix wrong index merge selection (#22825)](https://github.com/pingcap/tidb/pull/22924)                                                             |   | 15d14h  |
| tidb/22926   | [expression: add overflow check in multiplyInt](https://github.com/pingcap/tidb/pull/22926)                                                                 |   | 15d13h  |
| tidb/22984   | [executor: fix logging format of prepared statements (#16062)](https://github.com/pingcap/tidb/pull/22984)                                                  |   | 12d10h  |
| tidb/23022   | [executor: create PipelinedWindowExec based on current implementation and modify the windowProcessor interface](https://github.com/pingcap/tidb/pull/23022) |   | 10d18h  |
| tidb/23062   | [*: fix structcheck lint warnings](https://github.com/pingcap/tidb/pull/23062)                                                                              |   | 8d19h   |
| tidb/23074   | [planner: fix range partition prune bug for IN expr (#22894) (#22938)](https://github.com/pingcap/tidb/pull/23074)                                          |   | 8d17h   |
| tidb/23088   | [statistics: delete extended stats cache item in current tidb synchronously](https://github.com/pingcap/tidb/pull/23088)                                    |   | 8d14h   |
| tidb/23105   | [executor: fix wrong key range of index scan when filter is comparing year column with NULL (#23079)](https://github.com/pingcap/tidb/pull/23105)           |   | 7d18h   |
| tidb/23119   | [statistics: remove existing deleted extended stats when add a new one](https://github.com/pingcap/tidb/pull/23119)                                         |   | 7d14h   |
| tidb/23152   | [expression: fix wrong error info (#22760)](https://github.com/pingcap/tidb/pull/23152)                                                                     |   | 4d14h   |
| tidb/23161   | [util, types: don't let SPM be affected by charset](https://github.com/pingcap/tidb/pull/23161)                                                             |   | 3d19h   |
| tidb/23171   | [store/tikv:move option from kv to tikv, and make define as a normal int](https://github.com/pingcap/tidb/pull/23171)                                       |   | 3d17h   |
| tidb/23196   | [types: fix the bug about the wrong query result for decimal type  (#22507)](https://github.com/pingcap/tidb/pull/23196)                                    |   | 2d18h   |
| tidb/23201   | [executor, server: load_data.go is changed and add unit test](https://github.com/pingcap/tidb/pull/23201)                                                   |   | 2d17h   |
| tidb/23208   | [statistics, util/ranger: improve selectivity calculation for DNF filters (#18741)](https://github.com/pingcap/tidb/pull/23208)                             |   | 2d16h   |
| tidb/23209   | [planner: fixed a bug that prevented SPM from taking effect (#23197)](https://github.com/pingcap/tidb/pull/23209)                                           |   | 2d16h   |
| tidb/23210   | [planner: fixed a bug that prevented SPM from taking effect (#23197)](https://github.com/pingcap/tidb/pull/23210)                                           |   | 2d16h   |
| tidb/23211   | [executor: fix cast function will ignore tht error for point-get key construction (#22869)](https://github.com/pingcap/tidb/pull/23211)                     |   | 2d16h   |
| tidb/23234   | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23234)                                                      |   | 1d18h   |
| tidb/23238   | [planner: fix wrong PointGet / TableDual plan reused in plan cache](https://github.com/pingcap/tidb/pull/23238)                                             |   | 1d17h   |
| tidb/23248   | [server: add tiflash fallback testcase](https://github.com/pingcap/tidb/pull/23248)                                                                         |   | 21h     |


## Reviewed in Last 7 Days

|    REPO    |                                                                 PR                                                                 | C | D |   R    |
|------------|------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/23256 | [statistics: add tests for global-stats when add and delete a single partition](https://github.com/pingcap/tidb/pull/23256)        |   | 1 | 0h     |
| tidb/23240 | [config: disable prepare plan cache by default](https://github.com/pingcap/tidb/pull/23240)                                        |   | 1 | 21h    |
| tidb/23226 | [server: remove unstable tiflash fallback testcase](https://github.com/pingcap/tidb/pull/23226)                                    |   | 2 | 4h     |
| tidb/23181 | [statistics: test the auto analyze and feedback for the global-level stats](https://github.com/pingcap/tidb/pull/23181)            |   | 2 | 1d5h   |
| tidb/23217 | [planner: fix the bug that wrong collation is used when try fast path for enum or set](https://github.com/pingcap/tidb/pull/23217) |   | 2 | 20h    |
| tidb/23133 | [plan/core: support mpp group by expressions.](https://github.com/pingcap/tidb/pull/23133)                                         |   | 2 | 5d0h   |
| tidb/23176 | [statistics: test `tidb_partition_prune_mode` session variable](https://github.com/pingcap/tidb/pull/23176)                        |   | 4 | 0h     |
| tidb/23167 | [statistics: add test for analyze single partition](https://github.com/pingcap/tidb/pull/23167)                                    |   | 4 | 3h     |
| tidb/23166 | [planner: fix wrong table filters for index merge plan (#23132)](https://github.com/pingcap/tidb/pull/23166)                       |   | 4 | 3h     |
| tidb/23165 | [planner: fix wrong table filters for index merge plan (#23132)](https://github.com/pingcap/tidb/pull/23165)                       |   | 4 | 3h     |
| tidb/23088 | [statistics: delete extended stats cache item in current tidb synchronously](https://github.com/pingcap/tidb/pull/23088)           |   | 4 | 4d17h  |
| tidb/23141 | [planner: fix panic when building index merge plan](https://github.com/pingcap/tidb/pull/23141)                                    |   | 4 | 2d17h  |
| tidb/22845 | [planner: fix bug of mpp wrongly set schema of exchanger](https://github.com/pingcap/tidb/pull/22845)                              |   | 7 | 12d22h |
| tidb/23113 | [statistics: add tests for `analyze` with dynamic partition prune mode](https://github.com/pingcap/tidb/pull/23113)                |   | 7 | 17h    |


</details> 


<details> 
  <summary>rebelice's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                   PR                                                                    | C | LASTED |
|------------|-----------------------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/23074 | [planner: fix range partition prune bug for IN expr (#22894) (#22938)](https://github.com/pingcap/tidb/pull/23074)                      |   | 8d17h  |
| tidb/23211 | [executor: fix cast function will ignore tht error for point-get key construction (#22869)](https://github.com/pingcap/tidb/pull/23211) |   | 2d16h  |


## Reviewed in Last 7 Days

|    REPO    |                                                                       PR                                                                        | C | D |   R    |
|------------|-------------------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/23265 | [statistics: add multi-column index test cases for global-statas](https://github.com/pingcap/tidb/pull/23265)                                   |   | 1 | 2h     |
| tidb/23242 | [planner: fix a panic caused by unmatched FieldNames and ColsInfo in partition pruning](https://github.com/pingcap/tidb/pull/23242)             |   | 1 | 18h    |
| tidb/23181 | [statistics: test the auto analyze and feedback for the global-level stats](https://github.com/pingcap/tidb/pull/23181)                         |   | 2 | 1d5h   |
| tidb/23231 | [statistics: add a test case which builds global-stats on different versions of partition-stats](https://github.com/pingcap/tidb/pull/23231)    |   | 2 | 0h     |
| tidb/23228 | [statistics: update the count and modify variables of global-stats as well when dumping delta info](https://github.com/pingcap/tidb/pull/23228) |   | 2 | 0h     |
| tidb/23169 | [statistics: add more test cases for global-stats with different analysis options](https://github.com/pingcap/tidb/pull/23169)                  |   | 4 | 2h     |
| tidb/23138 | [statistics: add some test cases of global-stats to cover more column types](https://github.com/pingcap/tidb/pull/23138)                        |   | 4 | 2d19h  |
| tidb/20360 | [planner: refine explain info for batch cop](https://github.com/pingcap/tidb/pull/20360)                                                        |   | 7 | 147d7h |
| tidb/23099 | [statistics: add more test cases about global-stats and fix some issues](https://github.com/pingcap/tidb/pull/23099)                            |   | 7 | 23h    |
| tidb/23066 | [statistics: forbid getting global-stats through fast analyze and incremental analyze](https://github.com/pingcap/tidb/pull/23066)              |   | 7 | 1d19h  |


</details> 


<details> 
  <summary>time-and-fate's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                         PR                                                          | C | LASTED  |
|------------|---------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                           |   | 126d17h |
| tidb/22853 | [planner: fix LogicalPlans that contain Window Function are ambiguous ](https://github.com/pingcap/tidb/pull/22853) |   | 19d12h  |
| tidb/23119 | [statistics: remove existing deleted extended stats when add a new one](https://github.com/pingcap/tidb/pull/23119) |   | 7d14h   |


## Reviewed in Last 7 Days

|    REPO    |                                                            PR                                                            | C | D |   R   |
|------------|--------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/23238 | [planner: fix wrong PointGet / TableDual plan reused in plan cache](https://github.com/pingcap/tidb/pull/23238)          |   | 1 | 23h   |
| tidb/23088 | [statistics: delete extended stats cache item in current tidb synchronously](https://github.com/pingcap/tidb/pull/23088) |   | 2 | 6d15h |
| tidb/23165 | [planner: fix wrong table filters for index merge plan (#23132)](https://github.com/pingcap/tidb/pull/23165)             |   | 4 | 3h    |
| tidb/23166 | [planner: fix wrong table filters for index merge plan (#23132)](https://github.com/pingcap/tidb/pull/23166)             |   | 4 | 2h    |
| tidb/23132 | [planner: fix wrong table filters for index merge plan](https://github.com/pingcap/tidb/pull/23132)                      |   | 7 | 0h    |


</details> 


<details> 
  <summary>winoros's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                               PR                                                                | C | LASTED  |
|--------------|---------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                          | Y | 181d14h |
| docs-cn/5484 | [system variable: add tidb_enable_engine_fallback](https://github.com/pingcap/docs-cn/pull/5484)                                |   | 36d17h  |
| docs/4781    | [system variable: add tidb_enable_engine_fallback](https://github.com/pingcap/docs/pull/4781)                                   |   | 36d17h  |
| tidb/20311   | [expression: fix overflow error when convert bit to int64 (#20266)](https://github.com/pingcap/tidb/pull/20311)                 |   | 162d21h |
| tidb/20765   | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                               |   | 129d17h |
| tidb/20877   | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                       |   | 126d17h |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                  |   | 119d17h |
| tidb/21207   | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                 |   | 108d19h |
| tidb/21476   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                           |   | 98d15h  |
| tidb/21487   | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                   |   | 98d4h   |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)     |   | 83d19h  |
| tidb/21930   | [planner: propagate NDV of column groups across plan nodes (#17854)](https://github.com/pingcap/tidb/pull/21930)                |   | 79d18h  |
| tidb/21954   | [planner/cascades: add rule `PushSelDownApply`](https://github.com/pingcap/tidb/pull/21954)                                     |   | 78d23h  |
| tidb/22181   | [planner, expression: fix error when using IN combined with subquery (#22080)](https://github.com/pingcap/tidb/pull/22181)      |   | 65d17h  |
| tidb/22365   | [planner: check index valid while forUpdateRead (#22152)](https://github.com/pingcap/tidb/pull/22365)                           |   | 58d19h  |
| tidb/22504   | [*:Fix the fetchHotRegion bug that the count always zero](https://github.com/pingcap/tidb/pull/22504)                           |   | 46d19h  |
| tidb/22565   | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)              | Y | 43d17h  |
| tidb/22624   | [ planner: not pruning column used by union scan condition (#21640)](https://github.com/pingcap/tidb/pull/22624)                |   | 41d17h  |
| tidb/22886   | [*: rename tiflash fallback switch](https://github.com/pingcap/tidb/pull/22886)                                                 |   | 16d20h  |
| tidb/22923   | [expression: correct constant propagation for collation (#22666)](https://github.com/pingcap/tidb/pull/22923)                   |   | 15d15h  |
| tidb/23088   | [statistics: delete extended stats cache item in current tidb synchronously](https://github.com/pingcap/tidb/pull/23088)        |   | 8d14h   |
| tidb/23092   | [*: fix a bug that collation is not handle for text type (#23045)](https://github.com/pingcap/tidb/pull/23092)                  |   | 8d12h   |
| tidb/23163   | [plugin: fix linter --enable=deadcode check error](https://github.com/pingcap/tidb/pull/23163)                                  |   | 3d19h   |
| tidb/23208   | [statistics, util/ranger: improve selectivity calculation for DNF filters (#18741)](https://github.com/pingcap/tidb/pull/23208) |   | 2d16h   |
| tidb/23215   | [Privileges: fix delete privilege check wrongly (#22971)](https://github.com/pingcap/tidb/pull/23215)                           |   | 2d14h   |
| tidb/23216   | [Privileges: fix delete privilege check wrongly (#22971)](https://github.com/pingcap/tidb/pull/23216)                           |   | 2d14h   |
| tidb/23233   | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23233)                          |   | 1d18h   |
| tidb/23234   | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23234)                          |   | 1d18h   |
| tidb/23238   | [planner: fix wrong PointGet / TableDual plan reused in plan cache](https://github.com/pingcap/tidb/pull/23238)                 |   | 1d17h   |
| tidb/23246   | [planner: fix the panic in joinReOrderSolver.optimizeRecursive](https://github.com/pingcap/tidb/pull/23246)                     |   | 23h     |


## Reviewed in Last 7 Days

|    REPO    |                                                                 PR                                                                  | C | D |   R   |
|------------|-------------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/23248 | [server: add tiflash fallback testcase](https://github.com/pingcap/tidb/pull/23248)                                                 |   | 1 | 3h    |
| tidb/23219 | [statistics: fix a case that auto-analyze is triggered outside its time range (#23214)](https://github.com/pingcap/tidb/pull/23219) |   | 1 | 1d16h |
| tidb/23214 | [statistics: fix a case that auto-analyze is triggered outside its time range](https://github.com/pingcap/tidb/pull/23214)          |   | 3 | 0h    |
| tidb/22971 | [Privileges: fix delete privilege check wrongly](https://github.com/pingcap/tidb/pull/22971)                                        |   | 3 | 11d2h |
| tidb/23200 | [store: fix SetOption race error](https://github.com/pingcap/tidb/pull/23200)                                                       |   | 3 | 0h    |
| tidb/23103 | [server: refine tiflash fallback testcase](https://github.com/pingcap/tidb/pull/23103)                                              |   | 3 | 4d23h |
| tidb/23123 | [planner: show cast type in EXPLAIN in coptask](https://github.com/pingcap/tidb/pull/23123)                                         |   | 7 | 17h   |
| tidb/23122 | [test: make test TestUpdateWithTableReadLockWillFail stable](https://github.com/pingcap/tidb/pull/23122)                            |   | 7 | 16h   |


</details> 


<details> 
  <summary>wshwsh12's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                        PR                                                                         | C | LASTED  |
|------------|---------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19557 | [*: Integrate timeline tracing with TiKV](https://github.com/pingcap/tidb/pull/19557)                                                             |   | 195d23h |
| tidb/19807 | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                                           |   | 188d11h |
| tidb/19957 | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                            | Y | 181d14h |
| tidb/21487 | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                                     |   | 98d4h   |
| tidb/21887 | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                                   |   | 82d11h  |
| tidb/22269 | [executor: check storage.block-cache.capacity value](https://github.com/pingcap/tidb/pull/22269)                                                  |   | 63d17h  |
| tidb/22378 | [executor: vectorize hash aggregate](https://github.com/pingcap/tidb/pull/22378)                                                                  |   | 57d19h  |
| tidb/22628 | [executor: Improve max/min window function with deque-based sliding window](https://github.com/pingcap/tidb/pull/22628)                           |   | 40d23h  |
| tidb/22815 | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22815)                               |   | 20d19h  |
| tidb/23104 | [executor: fix wrong key range of index scan when filter is comparing year column with NULL (#23079)](https://github.com/pingcap/tidb/pull/23104) |   | 7d18h   |
| tidb/23105 | [executor: fix wrong key range of index scan when filter is comparing year column with NULL (#23079)](https://github.com/pingcap/tidb/pull/23105) |   | 7d18h   |
| tidb/23109 | [store/tikv:move set/delete option into tikv.Snapshot](https://github.com/pingcap/tidb/pull/23109)                                                |   | 7d17h   |
| tidb/23123 | [planner: show cast type in EXPLAIN in coptask](https://github.com/pingcap/tidb/pull/23123)                                                       |   | 7d13h   |
| tidb/23128 | [statistics: refactor the statistics package use the RestrictedSQLExecutor API (#22636)](https://github.com/pingcap/tidb/pull/23128)              |   | 6d22h   |
| tidb/23206 | [executor: fix the bug of step value overflow when split regions](https://github.com/pingcap/tidb/pull/23206)                                     |   | 2d16h   |
| tidb/23211 | [executor: fix cast function will ignore tht error for point-get key construction (#22869)](https://github.com/pingcap/tidb/pull/23211)           |   | 2d16h   |
| tidb/23227 | [executor: hash join out of index panic when enum column value is zero (#23162)](https://github.com/pingcap/tidb/pull/23227)                      |   | 1d22h   |
| tidb/23233 | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23233)                                            |   | 1d18h   |
| tidb/23234 | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23234)                                            |   | 1d18h   |
| tidb/23259 | [planner: not push down mpp/join to TiFlash in some cases that TiFlash not supported the query](https://github.com/pingcap/tidb/pull/23259)       |   | 18h     |
| tidb/23267 | [*: add infoschema client errors (#22382)](https://github.com/pingcap/tidb/pull/23267)                                                            |   | 15h     |
| tidb/23268 | [*: add infoschema client errors (#22382)](https://github.com/pingcap/tidb/pull/23268)                                                            |   | 15h     |


## Reviewed in Last 7 Days

|    REPO    |                                                         PR                                                          | C | D |   R    |
|------------|---------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/22830 | [planner: fix incorrect duration between compare](https://github.com/pingcap/tidb/pull/22830)                       |   | 2 | 18d14h |
| tidb/23162 | [executor: hash join out of index panic when enum column value is zero](https://github.com/pingcap/tidb/pull/23162) |   | 3 | 1d2h   |


</details> 

