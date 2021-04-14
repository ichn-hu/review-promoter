|   REVIEWER    |  O(C)   | 1 | 2 | 3 | 4 | 5 | 6 | 7 | T  |
|---------------|---------|---|---|---|---|---|---|---|----|
| Reminiscent   |  7 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 1 |  1 |
| SunRunAway    | 12 ( 1) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| XuHuaiyu      | 79 ( 3) | 0 | 0 | 0 | 0 | 0 | 0 | 1 |  1 |
| dyzsr         |  1 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| eurekaka      | 19 ( 0) | 1 | 0 | 0 | 0 | 0 | 0 | 1 |  2 |
| fzhedu        |  1 ( 0) | 0 | 0 | 0 | 0 | 2 | 0 | 1 |  3 |
| ichn-hu       |  2 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| lzmhhh123     | 51 ( 0) | 4 | 1 | 0 | 0 | 1 | 0 | 2 |  8 |
| nullnotnil    |  0 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| qw4990        | 60 ( 2) | 0 | 6 | 0 | 0 | 6 | 0 | 2 | 14 |
| rebelice      |  6 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 2 |  2 |
| time-and-fate |  5 ( 1) | 0 | 0 | 0 | 0 | 0 | 0 | 1 |  1 |
| winoros       | 41 ( 3) | 3 | 1 | 0 | 0 | 0 | 0 | 0 |  4 |
| wshwsh12      | 24 ( 1) | 1 | 1 | 0 | 0 | 0 | 0 | 4 |  6 |


<details> 
  <summary>Reminiscent's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                     PR                                                                     | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/21896 | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                                 |   | 113d19h |
| tidb/23474 | [planner: fix inappropriate null flag of null constants (#23457)](https://github.com/pingcap/tidb/pull/23474)                              |   | 21d18h  |
| tidb/23575 | [executor: fix update panic on join having statement (#23554)](https://github.com/pingcap/tidb/pull/23575)                                 |   | 18d21h  |
| tidb/23685 | [planner: fix the issue that planner hints don't work in some batch/point-get plans (#23666)](https://github.com/pingcap/tidb/pull/23685)  |   | 14d16h  |
| tidb/23917 | [planner: fix wrong TableDual plans caused by comparing Binary and Bytes incorrectly (#23860)](https://github.com/pingcap/tidb/pull/23917) |   | 5d23h   |
| tidb/23918 | [planner: fix wrong TableDual plans caused by comparing Binary and Bytes incorrectly (#23860)](https://github.com/pingcap/tidb/pull/23918) |   | 5d23h   |
| tidb/23962 | [statistics: use index fm-sketches instead of bucket NDV to calculate global NDV for indexes](https://github.com/pingcap/tidb/pull/23962)  |   | 1d17h   |


## Reviewed in Last 7 Days

|    REPO    |                                                       PR                                                       | C | D |  R   |
|------------|----------------------------------------------------------------------------------------------------------------|---|---|------|
| tidb/23838 | [statistics: collect FMSketch when processing index analyze tasks](https://github.com/pingcap/tidb/pull/23838) |   | 7 | 5d1h |


</details> 


<details> 
  <summary>SunRunAway's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                  PR                                                                   | C | LASTED  |
|------------|---------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19178 | [executor: Refactor probe channel](https://github.com/pingcap/tidb/pull/19178)                                                        |   | 243d16h |
| tidb/19807 | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                               |   | 221d10h |
| tidb/19900 | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                        | Y | 216d18h |
| tidb/20140 | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                    |   | 203d22h |
| tidb/20765 | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                     |   | 162d17h |
| tidb/21207 | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                       |   | 141d19h |
| tidb/21834 | [planner: enhanced index range calculation plan](https://github.com/pingcap/tidb/pull/21834)                                          |   | 118d18h |
| tidb/21876 | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)           |   | 116d19h |
| tidb/21878 | [planner: do not push down lock to pointGet/bacthPointGet when selection exists](https://github.com/pingcap/tidb/pull/21878)          |   | 116d18h |
| tidb/21956 | [planner/preprocessor: disallow into-outfile clause in some place](https://github.com/pingcap/tidb/pull/21956)                        |   | 111d23h |
| tidb/22217 | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                         |   | 97d17h  |
| tidb/22379 | [[experiment] executor: allow aggregation to spill disk when running out of memory quota](https://github.com/pingcap/tidb/pull/22379) |   | 90d19h  |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>XuHuaiyu's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                              PR                                                                               | C | LASTED  |
|--------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| docs-cn/5619 | [Update data-type-date-and-time.md](https://github.com/pingcap/docs-cn/pull/5619)                                                                             |   | 46d15h  |
| tidb/19900   | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                                                | Y | 216d18h |
| docs-cn/5671 | [tidb: Add time format description](https://github.com/pingcap/docs-cn/pull/5671)                                                                             |   | 40d11h  |
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                                        | Y | 214d13h |
| tidb/20140   | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                                            |   | 203d22h |
| tidb/20790   | [collation: add pinyin collation for chinese charset support](https://github.com/pingcap/tidb/pull/20790)                                                     |   | 161d20h |
| tidb/21064   | [planner, executor: fix cast not check error](https://github.com/pingcap/tidb/pull/21064)                                                                     |   | 149d8h  |
| tidb/21149   | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                                 |   | 145d14h |
| tidb/21228   | [executor: return the result immediately when combining LIMIT row_count with DISTINCT](https://github.com/pingcap/tidb/pull/21228)                            |   | 141d13h |
| tidb/21304   | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                                      |   | 139d12h |
| tidb/21334   | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                                                 |   | 138d14h |
| tidb/21401   | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                                            |   | 134d11h |
| tidb/21536   | [executor: add slow-log file meta cache to avoid repeat read file meta information](https://github.com/pingcap/tidb/pull/21536)                               |   | 127d14h |
| tidb/21564   | [ddl: fix Incorrect behavior of NO_ZERO_DATE when altering table](https://github.com/pingcap/tidb/pull/21564)                                                 |   | 126d15h |
| tidb/21896   | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                                                    |   | 113d19h |
| tidb/22131   | [privilege: remove leading and trailing space when create user and role](https://github.com/pingcap/tidb/pull/22131)                                          |   | 103d19h |
| tidb/22163   | [expression: separated arithmeticMinusIntSig](https://github.com/pingcap/tidb/pull/22163)                                                                     |   | 99d13h  |
| tidb/22186   | [executor: fix select into outfile with year type column has no data (#22175)](https://github.com/pingcap/tidb/pull/22186)                                    |   | 98d16h  |
| tidb/22616   | [expression: from_unixtime accept 64-bit integers](https://github.com/pingcap/tidb/pull/22616)                                                                |   | 74d23h  |
| tidb/22617   | [metrics: fix wrong bucket name of coprocessor cache (#22454)](https://github.com/pingcap/tidb/pull/22617)                                                    |   | 74d23h  |
| tidb/22624   | [ planner: not pruning column used by union scan condition (#21640)](https://github.com/pingcap/tidb/pull/22624)                                              |   | 74d17h  |
| tidb/22631   | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                                                                               |   | 72d23h  |
| tidb/22696   | [expression: enable arithmetic Mod push down](https://github.com/pingcap/tidb/pull/22696)                                                                     |   | 69d17h  |
| tidb/22711   | [executor: Fix inline schema name](https://github.com/pingcap/tidb/pull/22711)                                                                                |   | 69d11h  |
| tidb/22722   | [planner, errno: make error code of ErrMixOfGroupFuncAndFields consistent with MySQL](https://github.com/pingcap/tidb/pull/22722)                             |   | 68d20h  |
| tidb/22814   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22814)                                           |   | 53d19h  |
| tidb/22908   | [txn: Add txn state's view](https://github.com/pingcap/tidb/pull/22908)                                                                                       |   | 48d20h  |
| tidb/23012   | [executor: fix affected rows of ddls and complete uint tests](https://github.com/pingcap/tidb/pull/23012)                                                     |   | 44d16h  |
| tidb/23152   | [expression: fix wrong error info (#22760)](https://github.com/pingcap/tidb/pull/23152)                                                                       |   | 37d14h  |
| tidb/23196   | [types: fix the bug about the wrong query result for decimal type  (#22507)](https://github.com/pingcap/tidb/pull/23196)                                      |   | 35d18h  |
| tidb/23220   | [Release 4.0](https://github.com/pingcap/tidb/pull/23220)                                                                                                     |   | 35d11h  |
| tidb/23233   | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23233)                                                        |   | 34d18h  |
| tidb/23257   | [executor: group_concat aggr panic when session.group_concat_max_len is small (#23131)](https://github.com/pingcap/tidb/pull/23257)                           |   | 33d18h  |
| tidb/23295   | [util, types: don't let SPM be affected by charset (#23161)](https://github.com/pingcap/tidb/pull/23295)                                                      |   | 32d11h  |
| tidb/23335   | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23335)                                      |   | 28d19h  |
| tidb/23336   | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23336)                                      |   | 28d19h  |
| tidb/23347   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                                          |   | 28d18h  |
| tidb/23348   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                                          |   | 28d18h  |
| tidb/23350   | [util/stringutil, util/ranger, planner: use hierarchical separators to simplify the parsing for info of EXPLAIN ](https://github.com/pingcap/tidb/pull/23350) |   | 28d17h  |
| tidb/23368   | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23368)                                         |   | 27d20h  |
| tidb/23369   | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23369)                                         |   | 27d20h  |
| tidb/23397   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23397)                                                                |   | 26d17h  |
| tidb/23398   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23398)                                                                |   | 26d17h  |
| tidb/23405   | [domain: remove the exit chan, use context](https://github.com/pingcap/tidb/pull/23405)                                                                       |   | 26d17h  |
| tidb/23433   | [WIP: speed up for slow query logs retrieving ](https://github.com/pingcap/tidb/pull/23433)                                                                   |   | 25d17h  |
| tidb/23474   | [planner: fix inappropriate null flag of null constants (#23457)](https://github.com/pingcap/tidb/pull/23474)                                                 |   | 21d18h  |
| tidb/23487   | [planner: optimize count(distinct a) to count(a) if there is an unique key on a](https://github.com/pingcap/tidb/pull/23487)                                  | Y | 21d14h  |
| tidb/23497   | [expression: Let TiDB use Hyperscan to support multi-pattern-match](https://github.com/pingcap/tidb/pull/23497)                                               |   | 20d22h  |
| tidb/23517   | [*: Add the metric about the SQL with TiFlash Success  (#23426)](https://github.com/pingcap/tidb/pull/23517)                                                  |   | 20d12h  |
| tidb/23562   | [execution: reuse iterator in hash join](https://github.com/pingcap/tidb/pull/23562)                                                                          |   | 19d13h  |
| tidb/23640   | [*: fix the bug about YEAR(0.9) returns NULL instead of 0 in NO_ZERO_DATE mode](https://github.com/pingcap/tidb/pull/23640)                                   |   | 15d13h  |
| tidb/23661   | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23661)       |   | 14d20h  |
| tidb/23682   | [executor: fix a panic when batch point get is used for partition table (#23652)](https://github.com/pingcap/tidb/pull/23682)                                 |   | 14d16h  |
| tidb/23683   | [executor: fix a panic when batch point get is used for partition table (#23652)](https://github.com/pingcap/tidb/pull/23683)                                 |   | 14d16h  |
| tidb/23691   | [executor: fix index join on prefix column index (#23678)](https://github.com/pingcap/tidb/pull/23691)                                                        |   | 14d15h  |
| tidb/23705   | [executor: refineArgs() bug fix when compare int with very small decimal (#23694)](https://github.com/pingcap/tidb/pull/23705)                                |   | 14d13h  |
| tidb/23756   | [planner: fix set not null flag for outer join (#23727)](https://github.com/pingcap/tidb/pull/23756)                                                          |   | 13d14h  |
| tidb/23812   | [executor, planner: fix collation for hash join building (#23770)](https://github.com/pingcap/tidb/pull/23812)                                                |   | 12d12h  |
| tidb/23826   | [executor: fix 2nd index dup check after insert ignore on dup update primary (#23814)](https://github.com/pingcap/tidb/pull/23826)                            |   | 11d20h  |
| tidb/23857   | [types: fix type merge about bit type](https://github.com/pingcap/tidb/pull/23857)                                                                            |   | 7d19h   |
| tidb/23868   | [store, kv: Add information about async commit/1pc to tidb_last_txn_info (#23833)](https://github.com/pingcap/tidb/pull/23868)                                |   | 7d16h   |
| tidb/23878   | [functions: fix some string function has wrong collation and flag (#23835)](https://github.com/pingcap/tidb/pull/23878)                                       |   | 6d21h   |
| tidb/23879   | [functions: fix some string function has wrong collation and flag (#23835)](https://github.com/pingcap/tidb/pull/23879)                                       |   | 6d21h   |
| tidb/23884   | [Metric: Collect TiKV Read Metric for SLI/SLO](https://github.com/pingcap/tidb/pull/23884)                                                                    |   | 6d19h   |
| tidb/23888   | [executor: fix resource leak of Shuffle Executor.](https://github.com/pingcap/tidb/pull/23888)                                                                |   | 6d18h   |
| tidb/23902   | [expression: fix comparing year with datetime for equality](https://github.com/pingcap/tidb/pull/23902)                                                       |   | 6d15h   |
| tidb/23918   | [planner: fix wrong TableDual plans caused by comparing Binary and Bytes incorrectly (#23860)](https://github.com/pingcap/tidb/pull/23918)                    |   | 5d23h   |
| tidb/23931   | [*: don't allocate SessionIndexUsageCollector when indexUsageLease equals 0](https://github.com/pingcap/tidb/pull/23931)                                      |   | 4d18h   |
| tidb/23935   | [planner, sessionvar: avoid sending same task id to TiFlash (#23747)](https://github.com/pingcap/tidb/pull/23935)                                             |   | 4d13h   |
| tidb/23944   | [go.mod: update `gopsutil` to v3.21.3](https://github.com/pingcap/tidb/pull/23944)                                                                            |   | 2d14h   |
| tidb/23950   | [planner: fix comment mistake of `basePhysicalPlan. ExplainNormalizedInfo`](https://github.com/pingcap/tidb/pull/23950)                                       |   | 1d20h   |
| tidb/23956   | [*: Add the metric about the SQL with TiFlash Success  (#23426)](https://github.com/pingcap/tidb/pull/23956)                                                  |   | 1d19h   |
| tidb/23958   | [executor: fix `show table status` for the database with upper-cased name (#23896)](https://github.com/pingcap/tidb/pull/23958)                               |   | 1d18h   |
| tidb/23959   | [executor: fix `show table status` for the database with upper-cased name (#23896)](https://github.com/pingcap/tidb/pull/23959)                               |   | 1d18h   |
| tidb/23964   | [executor: GROUP_CONCAT(float) is not compatible with mysql](https://github.com/pingcap/tidb/pull/23964)                                                      |   | 1d16h   |
| tidb/23972   | [planner: change descScanFactor to scanFactor when rowCount is small.](https://github.com/pingcap/tidb/pull/23972)                                            |   | 1d13h   |
| tidb/23978   | [*: add security enhanced mode part 1](https://github.com/pingcap/tidb/pull/23978)                                                                            |   | 1d2h    |
| tidb/23981   | [planner: remove useless cast function in AggToProj](https://github.com/pingcap/tidb/pull/23981)                                                              |   | 20h     |
| tidb/23988   | [statistics: fix some potential panic in statistics](https://github.com/pingcap/tidb/pull/23988)                                                              |   | 18h     |


## Reviewed in Last 7 Days

|    REPO    |                                               PR                                                | C | D |  R  |
|------------|-------------------------------------------------------------------------------------------------|---|---|-----|
| tidb/23867 | [expression: fix wrong flen infer for bit constant](https://github.com/pingcap/tidb/pull/23867) |   | 7 | 16h |


</details> 


<details> 
  <summary>dyzsr's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                             PR                                                             | C | LASTED |
|------------|----------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/23559 | [ranger: fix the range construction behavior when the column's type is `YEAR`](https://github.com/pingcap/tidb/pull/23559) |   | 19d14h |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>eurekaka's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                         PR                                                                         | C | LASTED  |
|------------|----------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                                          |   | 159d16h |
| tidb/23002 | [store/*: fix err check](https://github.com/pingcap/tidb/pull/23002)                                                                               |   | 45d0h   |
| tidb/23283 | [util: optimize the performance of restore with db (#22910)](https://github.com/pingcap/tidb/pull/23283)                                           |   | 32d17h  |
| tidb/23316 | [planner: Fix rebuild range for prepared plan](https://github.com/pingcap/tidb/pull/23316)                                                         |   | 29d17h  |
| tidb/23373 | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                  |   | 27d19h  |
| tidb/23543 | [statistics: fix auto analyze log information incomplete (#23522)](https://github.com/pingcap/tidb/pull/23543)                                     |   | 19d18h  |
| tidb/23685 | [planner: fix the issue that planner hints don't work in some batch/point-get plans (#23666)](https://github.com/pingcap/tidb/pull/23685)          |   | 14d16h  |
| tidb/23689 | [planner: fix the panic when we calculate the partition range (#23651)](https://github.com/pingcap/tidb/pull/23689)                                |   | 14d16h  |
| tidb/23705 | [executor: refineArgs() bug fix when compare int with very small decimal (#23694)](https://github.com/pingcap/tidb/pull/23705)                     |   | 14d13h  |
| tidb/23756 | [planner: fix set not null flag for outer join (#23727)](https://github.com/pingcap/tidb/pull/23756)                                               |   | 13d14h  |
| tidb/23760 | [collation: fix tidb panic when compare string with collation](https://github.com/pingcap/tidb/pull/23760)                                         |   | 13d13h  |
| tidb/23883 | [*: don't allocate SessionIndexUsageCollector when indexUsageLease equals 0 (#23861)](https://github.com/pingcap/tidb/pull/23883)                  |   | 6d19h   |
| tidb/23926 | [planner: support explain verbose mode](https://github.com/pingcap/tidb/pull/23926)                                                                |   | 4d21h   |
| tidb/23938 | [planner,privilege: requires extra privileges for REPLACE and INSERT ON DUPLICATE statements (#23911)](https://github.com/pingcap/tidb/pull/23938) |   | 4d9h    |
| tidb/23939 | [planner,privilege: requires extra privileges for REPLACE and INSERT ON DUPLICATE statements (#23911)](https://github.com/pingcap/tidb/pull/23939) |   | 4d9h    |
| tidb/23969 | [planner: fix a bug that point get plan returns wrong column name (#23365)](https://github.com/pingcap/tidb/pull/23969)                            |   | 1d14h   |
| tidb/23970 | [planner: fix a bug that point get plan returns wrong column name (#23365)](https://github.com/pingcap/tidb/pull/23970)                            |   | 1d14h   |
| tidb/23974 | [planner: do not push down to TiFlash if the table scan require to scan data in desc order (#23948)](https://github.com/pingcap/tidb/pull/23974)   |   | 1d12h   |
| tidb/23988 | [statistics: fix some potential panic in statistics](https://github.com/pingcap/tidb/pull/23988)                                                   |   | 18h     |


## Reviewed in Last 7 Days

|    REPO    |                                                 PR                                                 | C | D |  R   |
|------------|----------------------------------------------------------------------------------------------------|---|---|------|
| tidb/23968 | [statistics: fix unstable TestDropPartitionStats test](https://github.com/pingcap/tidb/pull/23968) |   | 1 | 18h  |
| tidb/23718 | [*: add TableSample ID for PhysicalIDToTypeString()](https://github.com/pingcap/tidb/pull/23718)   |   | 7 | 7d5h |


</details> 


<details> 
  <summary>fzhedu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                        PR                                                         | C | LASTED |
|------------|-------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/23935 | [planner, sessionvar: avoid sending same task id to TiFlash (#23747)](https://github.com/pingcap/tidb/pull/23935) |   | 4d13h  |


## Reviewed in Last 7 Days

|    REPO    |                                                          PR                                                          | C | D |   R   |
|------------|----------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/23930 | [planner: make sure that join key type are always the same for mpp join](https://github.com/pingcap/tidb/pull/23930) |   | 5 | 5h    |
| tidb/23747 | [planner, sessionvar: avoid sending same task id to TiFlash](https://github.com/pingcap/tidb/pull/23747)             |   | 5 | 9d1h  |
| tics/1715  | [Do not allow columns with same name in TiFlash's Block](https://github.com/pingcap/tics/pull/1715)                  |   | 7 | 5d22h |


</details> 


<details> 
  <summary>ichn-hu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                               PR                                                | C | LASTED |
|------------|-------------------------------------------------------------------------------------------------|---|--------|
| tidb/23857 | [types: fix type merge about bit type](https://github.com/pingcap/tidb/pull/23857)              |   | 7d19h  |
| tidb/23867 | [expression: fix wrong flen infer for bit constant](https://github.com/pingcap/tidb/pull/23867) |   | 7d16h  |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>lzmhhh123's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                             PR                                                                              | C | LASTED  |
|------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20444 | [expression: add json_merge_patch](https://github.com/pingcap/tidb/pull/20444)                                                                              |   | 181d21h |
| tidb/20465 | [expression: add uuidShortFunction](https://github.com/pingcap/tidb/pull/20465)                                                                             |   | 180d19h |
| tidb/20642 | [executor: modify admin executors to support partitioned table with global index](https://github.com/pingcap/tidb/pull/20642)                               |   | 169d15h |
| tidb/20903 | [planner: fix confused and unnecessary double-projection in plans.](https://github.com/pingcap/tidb/pull/20903)                                             |   | 158d17h |
| tidb/21018 | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                              |   | 152d17h |
| tidb/21195 | [brie: integrate lightning to suport IMPORT statement](https://github.com/pingcap/tidb/pull/21195)                                                          |   | 141d22h |
| tidb/21334 | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                                               |   | 138d14h |
| tidb/21347 | [session: make rollback work on global variables](https://github.com/pingcap/tidb/pull/21347)                                                               |   | 137d19h |
| tidb/21487 | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                                               |   | 131d4h  |
| tidb/21641 | [executor: Fix pessimistic lock doesn't work on the partition table for subquery/joins](https://github.com/pingcap/tidb/pull/21641)                         |   | 124d18h |
| tidb/21651 | [planner: allow filter condition pushing down to IndexScan for prefix index](https://github.com/pingcap/tidb/pull/21651)                                    |   | 124d13h |
| tidb/22126 | [*: add `sys` schema, `sys.SCHEMA_UNUSED_INDEXES` view and `sys.SCHEMA_INDEX_USAGE` view](https://github.com/pingcap/tidb/pull/22126)                       |   | 103d19h |
| tidb/22361 | [table: fix insert into _tidb_rowid panic and rebase it if needed (#22062)](https://github.com/pingcap/tidb/pull/22361)                                     |   | 91d20h  |
| tidb/22372 | [executor: fix SelectForUpdate in decorrelated subquery under pessimistic mode](https://github.com/pingcap/tidb/pull/22372)                                 |   | 91d9h   |
| tidb/22478 | [planner, executor: fix query partition table with global unique index get wrong result](https://github.com/pingcap/tidb/pull/22478)                        |   | 82d13h  |
| tidb/22631 | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                                                                             |   | 72d23h  |
| tidb/22686 | [expression: support enum pushdown](https://github.com/pingcap/tidb/pull/22686)                                                                             |   | 69d22h  |
| tidb/22699 | [brie: add error info column and history backup/restore info in sql](https://github.com/pingcap/tidb/pull/22699)                                            |   | 69d16h  |
| tidb/23001 | [statistics: fix err check](https://github.com/pingcap/tidb/pull/23001)                                                                                     |   | 45d0h   |
| tidb/23022 | [executor: create PipelinedWindowExec based on current implementation and modify the windowProcessor interface](https://github.com/pingcap/tidb/pull/23022) |   | 43d18h  |
| tidb/23149 | [core: support left join and right join for join reorder](https://github.com/pingcap/tidb/pull/23149)                                                       |   | 38d12h  |
| tidb/23257 | [executor: group_concat aggr panic when session.group_concat_max_len is small (#23131)](https://github.com/pingcap/tidb/pull/23257)                         |   | 33d18h  |
| tidb/23283 | [util: optimize the performance of restore with db (#22910)](https://github.com/pingcap/tidb/pull/23283)                                                    |   | 32d17h  |
| tidb/23347 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                                        |   | 28d18h  |
| tidb/23348 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                                        |   | 28d18h  |
| tidb/23368 | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23368)                                       |   | 27d20h  |
| tidb/23369 | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23369)                                       |   | 27d20h  |
| tidb/23373 | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                           |   | 27d19h  |
| tidb/23559 | [ranger: fix the range construction behavior when the column's type is `YEAR`](https://github.com/pingcap/tidb/pull/23559)                                  |   | 19d14h  |
| tidb/23655 | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral (#23523)](https://github.com/pingcap/tidb/pull/23655)                    |   | 14d22h  |
| tidb/23656 | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral (#23523)](https://github.com/pingcap/tidb/pull/23656)                    |   | 14d22h  |
| tidb/23660 | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23660)     |   | 14d20h  |
| tidb/23661 | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23661)     |   | 14d20h  |
| tidb/23703 | [expression: fix approx_percent panic on bit column (#23687)](https://github.com/pingcap/tidb/pull/23703)                                                   |   | 14d13h  |
| tidb/23705 | [executor: refineArgs() bug fix when compare int with very small decimal (#23694)](https://github.com/pingcap/tidb/pull/23705)                              |   | 14d13h  |
| tidb/23714 | [*:Support record statment_history table evicted info](https://github.com/pingcap/tidb/pull/23714)                                                          |   | 14d1h   |
| tidb/23756 | [planner: fix set not null flag for outer join (#23727)](https://github.com/pingcap/tidb/pull/23756)                                                        |   | 13d14h  |
| tidb/23760 | [collation: fix tidb panic when compare string with collation](https://github.com/pingcap/tidb/pull/23760)                                                  |   | 13d13h  |
| tidb/23812 | [executor, planner: fix collation for hash join building (#23770)](https://github.com/pingcap/tidb/pull/23812)                                              |   | 12d12h  |
| tidb/23818 | [*: protect read only noop via tidb_enable_noop_functions](https://github.com/pingcap/tidb/pull/23818)                                                      |   | 12d5h   |
| tidb/23822 | [statistics: feedback not panic when no ndv collected (#23808)](https://github.com/pingcap/tidb/pull/23822)                                                 |   | 11d22h  |
| tidb/23902 | [expression: fix comparing year with datetime for equality](https://github.com/pingcap/tidb/pull/23902)                                                     |   | 6d15h   |
| tidb/23926 | [planner: support explain verbose mode](https://github.com/pingcap/tidb/pull/23926)                                                                         |   | 4d21h   |
| tidb/23935 | [planner, sessionvar: avoid sending same task id to TiFlash (#23747)](https://github.com/pingcap/tidb/pull/23935)                                           |   | 4d13h   |
| tidb/23936 | [planner: prevent appending `_tidb_rowid` to `DataSource` schema in index merge](https://github.com/pingcap/tidb/pull/23936)                                |   | 4d12h   |
| tidb/23939 | [planner,privilege: requires extra privileges for REPLACE and INSERT ON DUPLICATE statements (#23911)](https://github.com/pingcap/tidb/pull/23939)          |   | 4d9h    |
| tidb/23940 | [config, ddl: allow auto inc columns in generated columns and expression indexes](https://github.com/pingcap/tidb/pull/23940)                               |   | 3d17h   |
| tidb/23959 | [executor: fix `show table status` for the database with upper-cased name (#23896)](https://github.com/pingcap/tidb/pull/23959)                             |   | 1d18h   |
| tidb/23968 | [statistics: fix unstable TestDropPartitionStats test](https://github.com/pingcap/tidb/pull/23968)                                                          |   | 1d14h   |
| tidb/23987 | [executor: Implements json_arrayagg function](https://github.com/pingcap/tidb/pull/23987)                                                                   |   | 18h     |
| tidb/23989 | [expression: don't propagateColumnEQ joinCondition when nullSensitive](https://github.com/pingcap/tidb/pull/23989)                                          |   | 17h     |


## Reviewed in Last 7 Days

|    REPO    |                                                                        PR                                                                        | C | D |  R  |
|------------|--------------------------------------------------------------------------------------------------------------------------------------------------|---|---|-----|
| tidb/23982 | [expression: Fix incorrect pushdown function list.](https://github.com/pingcap/tidb/pull/23982)                                                  |   | 1 | 3h  |
| tidb/23981 | [planner: remove useless cast function in AggToProj](https://github.com/pingcap/tidb/pull/23981)                                                 |   | 1 | 1h  |
| tidb/23975 | [planner: do not push down to TiFlash if the table scan require to scan data in desc order (#23948)](https://github.com/pingcap/tidb/pull/23975) |   | 1 | 15h |
| tidb/23974 | [planner: do not push down to TiFlash if the table scan require to scan data in desc order (#23948)](https://github.com/pingcap/tidb/pull/23974) |   | 1 | 15h |
| tidb/23948 | [planner: do not push down to TiFlash if the table scan require to scan data in desc order](https://github.com/pingcap/tidb/pull/23948)          |   | 2 | 1h  |
| tidb/23922 | [planner: fix point-get response's original name](https://github.com/pingcap/tidb/pull/23922)                                                    |   | 5 | 18h |
| tidb/23888 | [executor: fix resource leak of Shuffle Executor.](https://github.com/pingcap/tidb/pull/23888)                                                   |   | 7 | 0h  |
| tidb/23885 | [executor,distsql: clean up useless interface](https://github.com/pingcap/tidb/pull/23885)                                                       |   | 7 | 0h  |


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
| tidb/19029   | [types: fix unexpected NOT_NULL flags](https://github.com/pingcap/tidb/pull/19029)                                                                          |   | 250d22h |
| docs-cn/5561 | [Add sql optimization-related docs to toc](https://github.com/pingcap/docs-cn/pull/5561)                                                                    |   | 50d15h  |
| tidb/20708   | [*: separate auto_increment ID allocator from _tidb_rowid allocator](https://github.com/pingcap/tidb/pull/20708)                                            |   | 166d20h |
| tidb/20969   | [executor: Improve the performance of appending not fixed columns](https://github.com/pingcap/tidb/pull/20969)                                              |   | 154d9h  |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                              |   | 152d17h |
| tidb/21149   | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                               |   | 145d14h |
| tidb/21304   | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                                    |   | 139d12h |
| tidb/21318   | [planner, expression: use the range of column types to simplify expressions](https://github.com/pingcap/tidb/pull/21318)                                    |   | 138d18h |
| tidb/21401   | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                                          |   | 134d11h |
| tidb/21508   | [execution: fix dayofweek('0000-00-00') behavior](https://github.com/pingcap/tidb/pull/21508)                                                               |   | 130d10h |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                                 |   | 116d19h |
| tidb/21887   | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                                             |   | 115d11h |
| tidb/21954   | [planner/cascades: add rule `PushSelDownApply`](https://github.com/pingcap/tidb/pull/21954)                                                                 |   | 111d23h |
| tidb/22146   | [executor: forbid SFU on view](https://github.com/pingcap/tidb/pull/22146)                                                                                  |   | 99d21h  |
| tidb/22217   | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                                               |   | 97d17h  |
| tidb/22234   | [executor, planner: ON DUPLICATE UPDATE can refer to un-project col (#14412)](https://github.com/pingcap/tidb/pull/22234)                                   |   | 97d15h  |
| tidb/22261   | [time: fix parse datetime won't truncate the reluctant string (#22232)](https://github.com/pingcap/tidb/pull/22261)                                         |   | 96d19h  |
| tidb/22374   | [expression: separated arithmeticIntDivideSig](https://github.com/pingcap/tidb/pull/22374)                                                                  |   | 91d0h   |
| tidb/22415   | [ddl: refactor placement package](https://github.com/pingcap/tidb/pull/22415)                                                                               |   | 87d17h  |
| tidb/22416   | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                                                           | Y | 87d11h  |
| tidb/22541   | [expression: Support builtin function SOUNDEX](https://github.com/pingcap/tidb/pull/22541)                                                                  |   | 77d9h   |
| tidb/22565   | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)                                          | Y | 76d17h  |
| tidb/22814   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22814)                                         |   | 53d19h  |
| tidb/22862   | [brie: fix the problem that ddl restored by BR via SQL is not replicated to downstream](https://github.com/pingcap/tidb/pull/22862)                         |   | 50d22h  |
| tidb/22923   | [expression: correct constant propagation for collation (#22666)](https://github.com/pingcap/tidb/pull/22923)                                               |   | 48d14h  |
| tidb/22924   | [planner: fix wrong index merge selection (#22825)](https://github.com/pingcap/tidb/pull/22924)                                                             |   | 48d13h  |
| tidb/22984   | [executor: fix logging format of prepared statements (#16062)](https://github.com/pingcap/tidb/pull/22984)                                                  |   | 45d10h  |
| tidb/23002   | [store/*: fix err check](https://github.com/pingcap/tidb/pull/23002)                                                                                        |   | 45d0h   |
| tidb/23022   | [executor: create PipelinedWindowExec based on current implementation and modify the windowProcessor interface](https://github.com/pingcap/tidb/pull/23022) |   | 43d18h  |
| tidb/23152   | [expression: fix wrong error info (#22760)](https://github.com/pingcap/tidb/pull/23152)                                                                     |   | 37d14h  |
| tidb/23196   | [types: fix the bug about the wrong query result for decimal type  (#22507)](https://github.com/pingcap/tidb/pull/23196)                                    |   | 35d18h  |
| tidb/23208   | [statistics, util/ranger: improve selectivity calculation for DNF filters (#18741)](https://github.com/pingcap/tidb/pull/23208)                             |   | 35d16h  |
| tidb/23295   | [util, types: don't let SPM be affected by charset (#23161)](https://github.com/pingcap/tidb/pull/23295)                                                    |   | 32d11h  |
| tidb/23316   | [planner: Fix rebuild range for prepared plan](https://github.com/pingcap/tidb/pull/23316)                                                                  |   | 29d17h  |
| tidb/23373   | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                           |   | 27d19h  |
| tidb/23397   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23397)                                                              |   | 26d17h  |
| tidb/23398   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23398)                                                              |   | 26d17h  |
| tidb/23448   | [wip :execution: parallel build hash table](https://github.com/pingcap/tidb/pull/23448)                                                                     |   | 23d12h  |
| tidb/23590   | [planner, table: optimize the list partition pruner for range query](https://github.com/pingcap/tidb/pull/23590)                                            |   | 18d16h  |
| tidb/23598   | [types: fix collation for binary literal (#23591)](https://github.com/pingcap/tidb/pull/23598)                                                              |   | 18d13h  |
| tidb/23655   | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral (#23523)](https://github.com/pingcap/tidb/pull/23655)                    |   | 14d22h  |
| tidb/23656   | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral (#23523)](https://github.com/pingcap/tidb/pull/23656)                    |   | 14d22h  |
| tidb/23658   | [*: collect transaction write duration/throughput metrics for SLI/SLO (#23462)](https://github.com/pingcap/tidb/pull/23658)                                 |   | 14d21h  |
| tidb/23660   | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23660)     |   | 14d20h  |
| tidb/23661   | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23661)     |   | 14d20h  |
| tidb/23682   | [executor: fix a panic when batch point get is used for partition table (#23652)](https://github.com/pingcap/tidb/pull/23682)                               |   | 14d16h  |
| tidb/23689   | [planner: fix the panic when we calculate the partition range (#23651)](https://github.com/pingcap/tidb/pull/23689)                                         |   | 14d16h  |
| tidb/23730   | [distsql/*: typo fix for `dispatches`](https://github.com/pingcap/tidb/pull/23730)                                                                          |   | 13d18h  |
| tidb/23796   | [tests: make TestIndexLookupMergeJoinHang and TestIssue18068 stable (#23741)](https://github.com/pingcap/tidb/pull/23796)                                   |   | 12d19h  |
| tidb/23812   | [executor, planner: fix collation for hash join building (#23770)](https://github.com/pingcap/tidb/pull/23812)                                              |   | 12d12h  |
| tidb/23867   | [expression: fix wrong flen infer for bit constant](https://github.com/pingcap/tidb/pull/23867)                                                             |   | 7d16h   |
| tidb/23878   | [functions: fix some string function has wrong collation and flag (#23835)](https://github.com/pingcap/tidb/pull/23878)                                     |   | 6d21h   |
| tidb/23879   | [functions: fix some string function has wrong collation and flag (#23835)](https://github.com/pingcap/tidb/pull/23879)                                     |   | 6d21h   |
| tidb/23918   | [planner: fix wrong TableDual plans caused by comparing Binary and Bytes incorrectly (#23860)](https://github.com/pingcap/tidb/pull/23918)                  |   | 5d23h   |
| tidb/23933   | [planner: append common handle columns into the schema of index merge table plan](https://github.com/pingcap/tidb/pull/23933)                               |   | 4d16h   |
| tidb/23954   | [*: collect transaction write duration/throughput metrics for SLI/SLO (#23462)](https://github.com/pingcap/tidb/pull/23954)                                 |   | 1d19h   |
| tidb/23963   | [executor: checking chunk is full precedes filtering](https://github.com/pingcap/tidb/pull/23963)                                                           |   | 1d17h   |
| tidb/23976   | [executor: skip ignore check for not update indexes (#23894)](https://github.com/pingcap/tidb/pull/23976)                                                   |   | 1d12h   |
| tidb/23980   | [planner: do not build MPP plan for scan with virtual columns](https://github.com/pingcap/tidb/pull/23980)                                                  |   | 22h     |
| tidb/23987   | [executor: Implements json_arrayagg function](https://github.com/pingcap/tidb/pull/23987)                                                                   |   | 18h     |


## Reviewed in Last 7 Days

|    REPO    |                                                                PR                                                                 | C | D |   R    |
|------------|-----------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/23733 | [planner: generate BatchPointGet for hash table partitions](https://github.com/pingcap/tidb/pull/23733)                           |   | 2 | 12d3h  |
| tidb/23931 | [*: don't allocate SessionIndexUsageCollector when indexUsageLease equals 0](https://github.com/pingcap/tidb/pull/23931)          |   | 2 | 3d4h   |
| tidb/23559 | [ranger: fix the range construction behavior when the column's type is `YEAR`](https://github.com/pingcap/tidb/pull/23559)        |   | 2 | 17d23h |
| tidb/23968 | [statistics: fix unstable TestDropPartitionStats test](https://github.com/pingcap/tidb/pull/23968)                                |   | 2 | 0h     |
| tidb/23926 | [planner: support explain verbose mode](https://github.com/pingcap/tidb/pull/23926)                                               |   | 2 | 3d4h   |
| tidb/23935 | [planner, sessionvar: avoid sending same task id to TiFlash (#23747)](https://github.com/pingcap/tidb/pull/23935)                 |   | 2 | 2d17h  |
| tidb/23922 | [planner: fix point-get response's original name](https://github.com/pingcap/tidb/pull/23922)                                     |   | 5 | 1d0h   |
| tidb/23674 | [*: add column `End_time` in show analyze status and add related log](https://github.com/pingcap/tidb/pull/23674)                 |   | 5 | 9d22h  |
| tidb/23890 | [statistics: check step overflow when converting a range to points for estimation](https://github.com/pingcap/tidb/pull/23890)    |   | 5 | 1d23h  |
| tidb/23895 | [statistics: add more tests for extended stats](https://github.com/pingcap/tidb/pull/23895)                                       |   | 5 | 1d22h  |
| tidb/23929 | [Revert "*: don't allocate SessionIndexUsageCollector when indexUsageLease equals 0"](https://github.com/pingcap/tidb/pull/23929) |   | 5 | 0h     |
| tidb/23862 | [statistics: make the global-level topN more accurate](https://github.com/pingcap/tidb/pull/23862)                                |   | 5 | 2d21h  |
| tidb/23883 | [*: don't allocate SessionIndexUsageCollector when indexUsageLease equals 0 (#23861)](https://github.com/pingcap/tidb/pull/23883) |   | 7 | 0h     |
| tidb/23861 | [*: don't allocate SessionIndexUsageCollector when indexUsageLease equals 0](https://github.com/pingcap/tidb/pull/23861)          |   | 7 | 18h    |


</details> 


<details> 
  <summary>rebelice's detailed review report</summary> 

## To Be Reviewed

|    REPO     |                                                                    PR                                                                     | C | LASTED |
|-------------|-------------------------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/23537  | [planner: remove some risky cache operations in the plan builder (#23354)](https://github.com/pingcap/tidb/pull/23537)                    |   | 19d19h |
| parser/1205 | [Implement force_index hint in parser and TiDB](https://github.com/pingcap/parser/pull/1205)                                              |   | 11d17h |
| tidb/23685  | [planner: fix the issue that planner hints don't work in some batch/point-get plans (#23666)](https://github.com/pingcap/tidb/pull/23685) |   | 14d16h |
| tidb/23836  | [parser, core: Implement force_index hint in parser and TiDB](https://github.com/pingcap/tidb/pull/23836)                                 |   | 11d17h |
| tidb/23926  | [planner: support explain verbose mode](https://github.com/pingcap/tidb/pull/23926)                                                       |   | 4d21h  |
| tidb/23962  | [statistics: use index fm-sketches instead of bucket NDV to calculate global NDV for indexes](https://github.com/pingcap/tidb/pull/23962) |   | 1d17h  |


## Reviewed in Last 7 Days

|    REPO    |                                                        PR                                                         | C | D |  R   |
|------------|-------------------------------------------------------------------------------------------------------------------|---|---|------|
| tidb/23674 | [*: add column `End_time` in show analyze status and add related log](https://github.com/pingcap/tidb/pull/23674) |   | 7 | 8d3h |
| tidb/23862 | [statistics: make the global-level topN more accurate](https://github.com/pingcap/tidb/pull/23862)                |   | 7 | 1d0h |


</details> 


<details> 
  <summary>time-and-fate's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                               PR                                                               | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                      |   | 159d16h |
| tidb/22416 | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                              | Y | 87d11h  |
| tidb/23890 | [statistics: check step overflow when converting a range to points for estimation](https://github.com/pingcap/tidb/pull/23890) |   | 6d18h   |
| tidb/23895 | [statistics: add more tests for extended stats](https://github.com/pingcap/tidb/pull/23895)                                    |   | 6d17h   |
| tidb/23933 | [planner: append common handle columns into the schema of index merge table plan](https://github.com/pingcap/tidb/pull/23933)  |   | 4d16h   |


## Reviewed in Last 7 Days

|    REPO    |                                                   PR                                                   | C | D |   R   |
|------------|--------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/23828 | [executor: do not show extended stats for dropped columns](https://github.com/pingcap/tidb/pull/23828) |   | 7 | 4d21h |


</details> 


<details> 
  <summary>winoros's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                              PR                                                                               | C | LASTED  |
|--------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                                        | Y | 214d13h |
| docs-cn/5916 | [sql-statements, information-schema: add `END_TIME` field for table `ANALYZE_STATUS`](https://github.com/pingcap/docs-cn/pull/5916)                           |   | 12d17h  |
| tidb/20765   | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                                             |   | 162d17h |
| tidb/20877   | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                                                     |   | 159d16h |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                                |   | 152d17h |
| tidb/21207   | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                                               |   | 141d19h |
| tidb/21487   | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                                                 |   | 131d4h  |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                                   |   | 116d19h |
| tidb/21954   | [planner/cascades: add rule `PushSelDownApply`](https://github.com/pingcap/tidb/pull/21954)                                                                   |   | 111d23h |
| tidb/22181   | [planner, expression: fix error when using IN combined with subquery (#22080)](https://github.com/pingcap/tidb/pull/22181)                                    |   | 98d17h  |
| tidb/22416   | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                                                             | Y | 87d11h  |
| tidb/22504   | [*:Fix the fetchHotRegion bug that the count always zero](https://github.com/pingcap/tidb/pull/22504)                                                         |   | 79d19h  |
| tidb/22565   | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)                                            | Y | 76d17h  |
| tidb/22624   | [ planner: not pruning column used by union scan condition (#21640)](https://github.com/pingcap/tidb/pull/22624)                                              |   | 74d17h  |
| tidb/22923   | [expression: correct constant propagation for collation (#22666)](https://github.com/pingcap/tidb/pull/22923)                                                 |   | 48d14h  |
| tidb/23208   | [statistics, util/ranger: improve selectivity calculation for DNF filters (#18741)](https://github.com/pingcap/tidb/pull/23208)                               |   | 35d16h  |
| tidb/23215   | [Privileges: fix delete privilege check wrongly (#22971)](https://github.com/pingcap/tidb/pull/23215)                                                         |   | 35d14h  |
| tidb/23233   | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23233)                                                        |   | 34d18h  |
| tidb/23347   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                                          |   | 28d18h  |
| tidb/23348   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                                          |   | 28d18h  |
| tidb/23350   | [util/stringutil, util/ranger, planner: use hierarchical separators to simplify the parsing for info of EXPLAIN ](https://github.com/pingcap/tidb/pull/23350) |   | 28d17h  |
| tidb/23373   | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                             |   | 27d19h  |
| tidb/23474   | [planner: fix inappropriate null flag of null constants (#23457)](https://github.com/pingcap/tidb/pull/23474)                                                 |   | 21d18h  |
| tidb/23537   | [planner: remove some risky cache operations in the plan builder (#23354)](https://github.com/pingcap/tidb/pull/23537)                                        |   | 19d19h  |
| tidb/23543   | [statistics: fix auto analyze log information incomplete (#23522)](https://github.com/pingcap/tidb/pull/23543)                                                |   | 19d18h  |
| tidb/23598   | [types: fix collation for binary literal (#23591)](https://github.com/pingcap/tidb/pull/23598)                                                                |   | 18d13h  |
| tidb/23655   | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral (#23523)](https://github.com/pingcap/tidb/pull/23655)                      |   | 14d22h  |
| tidb/23656   | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral (#23523)](https://github.com/pingcap/tidb/pull/23656)                      |   | 14d22h  |
| tidb/23683   | [executor: fix a panic when batch point get is used for partition table (#23652)](https://github.com/pingcap/tidb/pull/23683)                                 |   | 14d16h  |
| tidb/23689   | [planner: fix the panic when we calculate the partition range (#23651)](https://github.com/pingcap/tidb/pull/23689)                                           |   | 14d16h  |
| tidb/23772   | [tablecodec: fix text type decode for old row format (#23751)](https://github.com/pingcap/tidb/pull/23772)                                                    |   | 13d11h  |
| tidb/23849   | [ddl: tidb panic while query hash partition table with is null condition](https://github.com/pingcap/tidb/pull/23849)                                         |   | 8d13h   |
| tidb/23883   | [*: don't allocate SessionIndexUsageCollector when indexUsageLease equals 0 (#23861)](https://github.com/pingcap/tidb/pull/23883)                             |   | 6d19h   |
| tidb/23895   | [statistics: add more tests for extended stats](https://github.com/pingcap/tidb/pull/23895)                                                                   |   | 6d17h   |
| tidb/23917   | [planner: fix wrong TableDual plans caused by comparing Binary and Bytes incorrectly (#23860)](https://github.com/pingcap/tidb/pull/23917)                    |   | 5d23h   |
| tidb/23918   | [planner: fix wrong TableDual plans caused by comparing Binary and Bytes incorrectly (#23860)](https://github.com/pingcap/tidb/pull/23918)                    |   | 5d23h   |
| tidb/23946   | [planner: fix visit info for grant/revoke](https://github.com/pingcap/tidb/pull/23946)                                                                        |   | 2d6h    |
| tidb/23969   | [planner: fix a bug that point get plan returns wrong column name (#23365)](https://github.com/pingcap/tidb/pull/23969)                                       |   | 1d14h   |
| tidb/23970   | [planner: fix a bug that point get plan returns wrong column name (#23365)](https://github.com/pingcap/tidb/pull/23970)                                       |   | 1d14h   |
| tidb/23972   | [planner: change descScanFactor to scanFactor when rowCount is small.](https://github.com/pingcap/tidb/pull/23972)                                            |   | 1d13h   |
| tidb/23988   | [statistics: fix some potential panic in statistics](https://github.com/pingcap/tidb/pull/23988)                                                              |   | 18h     |


## Reviewed in Last 7 Days

|    REPO    |                                                                    PR                                                                     | C | D |   R   |
|------------|-------------------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/23991 | [distsql: refactor range splitting logic & comments on int64 boundary](https://github.com/pingcap/tidb/pull/23991)                        |   | 1 | 0h    |
| tidb/23981 | [planner: remove useless cast function in AggToProj](https://github.com/pingcap/tidb/pull/23981)                                          |   | 1 | 2h    |
| tidb/23962 | [statistics: use index fm-sketches instead of bucket NDV to calculate global NDV for indexes](https://github.com/pingcap/tidb/pull/23962) |   | 1 | 21h   |
| tidb/23933 | [planner: append common handle columns into the schema of index merge table plan](https://github.com/pingcap/tidb/pull/23933)             |   | 2 | 2d22h |


</details> 


<details> 
  <summary>wshwsh12's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                               PR                                                               | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19807 | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                        |   | 221d10h |
| tidb/19957 | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                         | Y | 214d13h |
| tidb/21487 | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                  |   | 131d4h  |
| tidb/21887 | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                |   | 115d11h |
| tidb/22378 | [executor: vectorize hash aggregate](https://github.com/pingcap/tidb/pull/22378)                                               |   | 90d19h  |
| tidb/22686 | [expression: support enum pushdown](https://github.com/pingcap/tidb/pull/22686)                                                |   | 69d22h  |
| tidb/23336 | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23336)       |   | 28d19h  |
| tidb/23347 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                           |   | 28d18h  |
| tidb/23348 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                           |   | 28d18h  |
| tidb/23368 | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23368)          |   | 27d20h  |
| tidb/23369 | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23369)          |   | 27d20h  |
| tidb/23397 | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23397)                                 |   | 26d17h  |
| tidb/23398 | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23398)                                 |   | 26d17h  |
| tidb/23519 | [executor: check privilege before adding](https://github.com/pingcap/tidb/pull/23519)                                          |   | 20d0h   |
| tidb/23705 | [executor: refineArgs() bug fix when compare int with very small decimal (#23694)](https://github.com/pingcap/tidb/pull/23705) |   | 14d13h  |
| tidb/23760 | [collation: fix tidb panic when compare string with collation](https://github.com/pingcap/tidb/pull/23760)                     |   | 13d13h  |
| tidb/23866 | [executor,kv: support timebounded staleness transaction](https://github.com/pingcap/tidb/pull/23866)                           |   | 7d16h   |
| tidb/23876 | [executor: fix scope ambiguity of joinResult](https://github.com/pingcap/tidb/pull/23876)                                      |   | 7d10h   |
| tidb/23954 | [*: collect transaction write duration/throughput metrics for SLI/SLO (#23462)](https://github.com/pingcap/tidb/pull/23954)    |   | 1d19h   |
| tidb/23960 | [executor: fix wrong convert from bit to string when do projection](https://github.com/pingcap/tidb/pull/23960)                |   | 1d17h   |
| tidb/23968 | [statistics: fix unstable TestDropPartitionStats test](https://github.com/pingcap/tidb/pull/23968)                             |   | 1d14h   |
| tidb/23979 | [executor, statistics: fix unstable `TestAnalyzeIndexExtractTopN`](https://github.com/pingcap/tidb/pull/23979)                 |   | 22h     |
| tidb/23985 | [execdetails: move execdetails to store/tikv](https://github.com/pingcap/tidb/pull/23985)                                      |   | 19h     |
| tidb/23991 | [distsql: refactor range splitting logic & comments on int64 boundary](https://github.com/pingcap/tidb/pull/23991)             |   | 16h     |


## Reviewed in Last 7 Days

|    REPO    |                                                           PR                                                            | C | D |   R   |
|------------|-------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/22628 | [executor: Improve max/min window function with deque-based sliding window](https://github.com/pingcap/tidb/pull/22628) |   | 1 | 73d6h |
| tidb/23260 | [test: fix global kill e2e test](https://github.com/pingcap/tidb/pull/23260)                                            |   | 2 | 32d0h |
| tidb/23915 | [expression: Implementation of Vitess hashing algorithm. (#23493)](https://github.com/pingcap/tidb/pull/23915)          |   | 7 | 0h    |
| tidb/23878 | [functions: fix some string function has wrong collation and flag (#23835)](https://github.com/pingcap/tidb/pull/23878) |   | 7 | 3h    |
| tidb/23879 | [functions: fix some string function has wrong collation and flag (#23835)](https://github.com/pingcap/tidb/pull/23879) |   | 7 | 3h    |
| tidb/23835 | [functions: fix some string function has wrong collation and flag](https://github.com/pingcap/tidb/pull/23835)          |   | 7 | 4d18h |


</details> 

