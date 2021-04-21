|   REVIEWER    |  O(C)   | 1 | 2 | 3 | 4 | 5 | 6 | 7 | T  |
|---------------|---------|---|---|---|---|---|---|---|----|
| Reminiscent   |  5 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| SunRunAway    | 12 ( 1) | 0 | 0 | 0 | 0 | 0 | 0 | 1 |  1 |
| XuHuaiyu      | 74 ( 3) | 0 | 1 | 0 | 0 | 1 | 1 | 0 |  3 |
| dyzsr         |  1 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| eurekaka      | 18 ( 0) | 1 | 1 | 0 | 0 | 4 | 4 | 6 | 16 |
| fzhedu        |  0 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| ichn-hu       |  1 ( 0) | 0 | 0 | 0 | 0 | 0 | 1 | 1 |  2 |
| lzmhhh123     | 49 ( 0) | 1 | 3 | 0 | 0 | 2 | 2 | 2 | 10 |
| nullnotnil    |  0 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| qw4990        | 55 ( 2) | 0 | 0 | 0 | 0 | 1 | 5 | 1 |  7 |
| rebelice      |  3 ( 0) | 0 | 0 | 0 | 0 | 1 | 0 | 0 |  1 |
| time-and-fate |  4 ( 1) | 0 | 1 | 0 | 0 | 1 | 5 | 2 |  9 |
| winoros       | 38 ( 3) | 3 | 0 | 0 | 0 | 7 | 4 | 0 | 14 |
| wshwsh12      | 24 ( 1) | 3 | 0 | 0 | 0 | 1 | 1 | 2 |  7 |


<details> 
  <summary>Reminiscent's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                     PR                                                                     | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/21896 | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                                 |   | 120d19h |
| tidb/23474 | [planner: fix inappropriate null flag of null constants (#23457)](https://github.com/pingcap/tidb/pull/23474)                              |   | 28d18h  |
| tidb/23575 | [executor: fix update panic on join having statement (#23554)](https://github.com/pingcap/tidb/pull/23575)                                 |   | 25d21h  |
| tidb/23917 | [planner: fix wrong TableDual plans caused by comparing Binary and Bytes incorrectly (#23860)](https://github.com/pingcap/tidb/pull/23917) |   | 12d23h  |
| tidb/24016 | [planner: fix index-out-of-range error when checking only_full_group_by (#23844)](https://github.com/pingcap/tidb/pull/24016)              |   | 6d18h   |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>SunRunAway's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                  PR                                                                   | C | LASTED  |
|------------|---------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19178 | [executor: Refactor probe channel](https://github.com/pingcap/tidb/pull/19178)                                                        |   | 250d16h |
| tidb/19807 | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                               |   | 228d10h |
| tidb/19900 | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                        | Y | 223d18h |
| tidb/20140 | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                    |   | 210d22h |
| tidb/20765 | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                     |   | 169d17h |
| tidb/21207 | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                       |   | 148d19h |
| tidb/21834 | [planner: enhanced index range calculation plan](https://github.com/pingcap/tidb/pull/21834)                                          |   | 125d18h |
| tidb/21876 | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)           |   | 123d19h |
| tidb/21878 | [planner: do not push down lock to pointGet/bacthPointGet when selection exists](https://github.com/pingcap/tidb/pull/21878)          |   | 123d18h |
| tidb/21956 | [planner/preprocessor: disallow into-outfile clause in some place](https://github.com/pingcap/tidb/pull/21956)                        |   | 118d23h |
| tidb/22217 | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                         |   | 104d17h |
| tidb/22379 | [[experiment] executor: allow aggregation to spill disk when running out of memory quota](https://github.com/pingcap/tidb/pull/22379) |   | 97d19h  |


## Reviewed in Last 7 Days

|    REPO    |                                          PR                                           | C | D |   R    |
|------------|---------------------------------------------------------------------------------------|---|---|--------|
| tidb/20749 | [executor: support global kill (32 bits)](https://github.com/pingcap/tidb/pull/20749) |   | 7 | 163d5h |


</details> 


<details> 
  <summary>XuHuaiyu's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                              PR                                                                               | C | LASTED  |
|--------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| docs-cn/5619 | [Update data-type-date-and-time.md](https://github.com/pingcap/docs-cn/pull/5619)                                                                             |   | 53d15h  |
| docs-cn/5671 | [tidb: Add time format description](https://github.com/pingcap/docs-cn/pull/5671)                                                                             |   | 47d11h  |
| tidb/19900   | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                                                | Y | 223d18h |
| docs-cn/6062 | [add MySQL compatibility in user-defined-variables.md](https://github.com/pingcap/docs-cn/pull/6062)                                                          | Y | 1d16h   |
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                                        | Y | 221d13h |
| tidb/20140   | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                                            |   | 210d22h |
| tidb/20790   | [collation: add pinyin collation for chinese charset support](https://github.com/pingcap/tidb/pull/20790)                                                     |   | 168d20h |
| tidb/21064   | [planner, executor: fix cast not check error](https://github.com/pingcap/tidb/pull/21064)                                                                     |   | 156d8h  |
| tidb/21149   | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                                 |   | 152d14h |
| tidb/21228   | [executor: return the result immediately when combining LIMIT row_count with DISTINCT](https://github.com/pingcap/tidb/pull/21228)                            |   | 148d13h |
| tidb/21304   | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                                      |   | 146d12h |
| tidb/21334   | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                                                 |   | 145d14h |
| tidb/21401   | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                                            |   | 141d11h |
| tidb/21536   | [executor: add slow-log file meta cache to avoid repeat read file meta information](https://github.com/pingcap/tidb/pull/21536)                               |   | 134d14h |
| tidb/21564   | [ddl: fix Incorrect behavior of NO_ZERO_DATE when altering table](https://github.com/pingcap/tidb/pull/21564)                                                 |   | 133d15h |
| tidb/21896   | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                                                    |   | 120d19h |
| tidb/22131   | [privilege: remove leading and trailing space when create user and role](https://github.com/pingcap/tidb/pull/22131)                                          |   | 110d19h |
| tidb/22163   | [expression: separated arithmeticMinusIntSig](https://github.com/pingcap/tidb/pull/22163)                                                                     |   | 106d13h |
| tidb/22186   | [executor: fix select into outfile with year type column has no data (#22175)](https://github.com/pingcap/tidb/pull/22186)                                    |   | 105d16h |
| tidb/22616   | [expression: from_unixtime accept 64-bit integers](https://github.com/pingcap/tidb/pull/22616)                                                                |   | 81d23h  |
| tidb/22617   | [metrics: fix wrong bucket name of coprocessor cache (#22454)](https://github.com/pingcap/tidb/pull/22617)                                                    |   | 81d23h  |
| tidb/22624   | [ planner: not pruning column used by union scan condition (#21640)](https://github.com/pingcap/tidb/pull/22624)                                              |   | 81d17h  |
| tidb/22631   | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                                                                               |   | 79d23h  |
| tidb/22696   | [expression: enable arithmetic Mod push down](https://github.com/pingcap/tidb/pull/22696)                                                                     |   | 76d17h  |
| tidb/22711   | [executor: Fix inline schema name](https://github.com/pingcap/tidb/pull/22711)                                                                                |   | 76d11h  |
| tidb/22722   | [planner, errno: make error code of ErrMixOfGroupFuncAndFields consistent with MySQL](https://github.com/pingcap/tidb/pull/22722)                             |   | 75d20h  |
| tidb/22814   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22814)                                           |   | 60d19h  |
| tidb/22908   | [txn: Add txn state's view](https://github.com/pingcap/tidb/pull/22908)                                                                                       |   | 55d20h  |
| tidb/23012   | [executor: fix affected rows of ddls and complete uint tests](https://github.com/pingcap/tidb/pull/23012)                                                     |   | 51d16h  |
| tidb/23152   | [expression: fix wrong error info (#22760)](https://github.com/pingcap/tidb/pull/23152)                                                                       |   | 44d14h  |
| tidb/23196   | [types: fix the bug about the wrong query result for decimal type  (#22507)](https://github.com/pingcap/tidb/pull/23196)                                      |   | 42d18h  |
| tidb/23220   | [Release 4.0](https://github.com/pingcap/tidb/pull/23220)                                                                                                     |   | 42d11h  |
| tidb/23233   | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23233)                                                        |   | 41d18h  |
| tidb/23257   | [executor: group_concat aggr panic when session.group_concat_max_len is small (#23131)](https://github.com/pingcap/tidb/pull/23257)                           |   | 40d18h  |
| tidb/23295   | [util, types: don't let SPM be affected by charset (#23161)](https://github.com/pingcap/tidb/pull/23295)                                                      |   | 39d11h  |
| tidb/23335   | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23335)                                      |   | 35d19h  |
| tidb/23336   | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23336)                                      |   | 35d19h  |
| tidb/23347   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                                          |   | 35d18h  |
| tidb/23348   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                                          |   | 35d18h  |
| tidb/23350   | [util/stringutil, util/ranger, planner: use hierarchical separators to simplify the parsing for info of EXPLAIN ](https://github.com/pingcap/tidb/pull/23350) |   | 35d17h  |
| tidb/23368   | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23368)                                         |   | 34d20h  |
| tidb/23397   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23397)                                                                |   | 33d17h  |
| tidb/23398   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23398)                                                                |   | 33d17h  |
| tidb/23405   | [domain: remove the exit chan, use context](https://github.com/pingcap/tidb/pull/23405)                                                                       |   | 33d17h  |
| tidb/23433   | [WIP: speed up for slow query logs retrieving ](https://github.com/pingcap/tidb/pull/23433)                                                                   |   | 32d17h  |
| tidb/23474   | [planner: fix inappropriate null flag of null constants (#23457)](https://github.com/pingcap/tidb/pull/23474)                                                 |   | 28d18h  |
| tidb/23497   | [expression: Let TiDB use Hyperscan to support multi-pattern-match](https://github.com/pingcap/tidb/pull/23497)                                               |   | 27d22h  |
| tidb/23517   | [*: Add the metric about the SQL with TiFlash Success  (#23426)](https://github.com/pingcap/tidb/pull/23517)                                                  |   | 27d12h  |
| tidb/23562   | [execution: reuse iterator in hash join](https://github.com/pingcap/tidb/pull/23562)                                                                          |   | 26d13h  |
| tidb/23640   | [*: fix the bug about YEAR(0.9) returns NULL instead of 0 in NO_ZERO_DATE mode](https://github.com/pingcap/tidb/pull/23640)                                   |   | 22d13h  |
| tidb/23661   | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23661)       |   | 21d20h  |
| tidb/23682   | [executor: fix a panic when batch point get is used for partition table (#23652)](https://github.com/pingcap/tidb/pull/23682)                                 |   | 21d16h  |
| tidb/23691   | [executor: fix index join on prefix column index (#23678)](https://github.com/pingcap/tidb/pull/23691)                                                        |   | 21d15h  |
| tidb/23705   | [executor: refineArgs() bug fix when compare int with very small decimal (#23694)](https://github.com/pingcap/tidb/pull/23705)                                |   | 21d13h  |
| tidb/23756   | [planner: fix set not null flag for outer join (#23727)](https://github.com/pingcap/tidb/pull/23756)                                                          |   | 20d14h  |
| tidb/23812   | [executor, planner: fix collation for hash join building (#23770)](https://github.com/pingcap/tidb/pull/23812)                                                |   | 19d12h  |
| tidb/23867   | [expression: fix wrong flen infer for bit constant](https://github.com/pingcap/tidb/pull/23867)                                                               |   | 14d16h  |
| tidb/23878   | [functions: fix some string function has wrong collation and flag (#23835)](https://github.com/pingcap/tidb/pull/23878)                                       |   | 13d21h  |
| tidb/23884   | [Metric: Collect TiKV Read Metric for SLI/SLO](https://github.com/pingcap/tidb/pull/23884)                                                                    |   | 13d19h  |
| tidb/23888   | [executor: fix resource leak of Shuffle Executor.](https://github.com/pingcap/tidb/pull/23888)                                                                |   | 13d18h  |
| tidb/23958   | [executor: fix `show table status` for the database with upper-cased name (#23896)](https://github.com/pingcap/tidb/pull/23958)                               |   | 8d18h   |
| tidb/23964   | [executor: GROUP_CONCAT(float) is not compatible with mysql](https://github.com/pingcap/tidb/pull/23964)                                                      |   | 8d16h   |
| tidb/24007   | [ddl: refactor rule [4/6]](https://github.com/pingcap/tidb/pull/24007)                                                                                        |   | 6d20h   |
| tidb/24016   | [planner: fix index-out-of-range error when checking only_full_group_by (#23844)](https://github.com/pingcap/tidb/pull/24016)                                 |   | 6d18h   |
| tidb/24026   | [types: fix type merge about bit type (#23857)](https://github.com/pingcap/tidb/pull/24026)                                                                   |   | 6d14h   |
| tidb/24033   | [statistics: fix some unstable tests in global stats (#23502)](https://github.com/pingcap/tidb/pull/24033)                                                    |   | 6d9h    |
| tidb/24053   | [executor: fix wrong convert from bit to string when do projection (#23960)](https://github.com/pingcap/tidb/pull/24053)                                      |   | 5d16h   |
| tidb/24060   | [statistics: fix some potential panic in statistics (#23988)](https://github.com/pingcap/tidb/pull/24060)                                                     |   | 5d13h   |
| tidb/24061   | [statistics: fix some potential panic in statistics (#23988)](https://github.com/pingcap/tidb/pull/24061)                                                     |   | 5d13h   |
| tidb/24078   | [planner: change descScanFactor to scanFactor when ExpectedCount is small. (#23972)](https://github.com/pingcap/tidb/pull/24078)                              |   | 4d19h   |
| tidb/24079   | [planner: change descScanFactor to scanFactor when ExpectedCount is small. (#23972)](https://github.com/pingcap/tidb/pull/24079)                              |   | 4d19h   |
| tidb/24155   | [planner, executor: fix index merge partial table scan schema (#23936)](https://github.com/pingcap/tidb/pull/24155)                                           |   | 19h     |
| tidb/24157   | [planner: let CopTiFlashConcurrencyFactor inflence the cost of whole plan](https://github.com/pingcap/tidb/pull/24157)                                        |   | 19h     |
| tidb/24177   | [executor: accelerate TestVectorizedMergeJoin](https://github.com/pingcap/tidb/pull/24177)                                                                    |   | 14h     |


## Reviewed in Last 7 Days

|    REPO    |                                                                          PR                                                                          | C | D |  R  |
|------------|------------------------------------------------------------------------------------------------------------------------------------------------------|---|---|-----|
| tidb/24132 | [executor: accelerate TestUpdateScanningHandles and TestIssue20658 and TestParallelStreamAggGroupConcat](https://github.com/pingcap/tidb/pull/24132) |   | 2 | 3h  |
| tidb/24059 | [executor: bypass encoding invalid datetime for index join (#24051)](https://github.com/pingcap/tidb/pull/24059)                                     |   | 5 | 20h |
| tidb/24051 | [executor: bypass encoding invalid datetime for index join](https://github.com/pingcap/tidb/pull/24051)                                              |   | 6 | 0h  |


</details> 


<details> 
  <summary>dyzsr's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                 PR                                                                  | C | LASTED |
|------------|-------------------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/24018 | [ranger: fix the range construction behavior when the column's type is `YEAR` (#23559)](https://github.com/pingcap/tidb/pull/24018) |   | 6d18h  |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>eurekaka's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                         PR                                                                         | C | LASTED  |
|------------|----------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                                          |   | 166d16h |
| tidb/23002 | [store/*: fix err check](https://github.com/pingcap/tidb/pull/23002)                                                                               |   | 52d0h   |
| tidb/23283 | [util: optimize the performance of restore with db (#22910)](https://github.com/pingcap/tidb/pull/23283)                                           |   | 39d17h  |
| tidb/23316 | [planner: Fix rebuild range for prepared plan](https://github.com/pingcap/tidb/pull/23316)                                                         |   | 36d17h  |
| tidb/23373 | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                  |   | 34d19h  |
| tidb/23543 | [statistics: fix auto analyze log information incomplete (#23522)](https://github.com/pingcap/tidb/pull/23543)                                     |   | 26d18h  |
| tidb/23689 | [planner: fix the panic when we calculate the partition range (#23651)](https://github.com/pingcap/tidb/pull/23689)                                |   | 21d16h  |
| tidb/23705 | [executor: refineArgs() bug fix when compare int with very small decimal (#23694)](https://github.com/pingcap/tidb/pull/23705)                     |   | 21d13h  |
| tidb/23756 | [planner: fix set not null flag for outer join (#23727)](https://github.com/pingcap/tidb/pull/23756)                                               |   | 20d14h  |
| tidb/23760 | [collation: fix tidb panic when compare string with collation](https://github.com/pingcap/tidb/pull/23760)                                         |   | 20d13h  |
| tidb/23938 | [planner,privilege: requires extra privileges for REPLACE and INSERT ON DUPLICATE statements (#23911)](https://github.com/pingcap/tidb/pull/23938) |   | 11d10h  |
| tidb/23974 | [planner: do not push down to TiFlash if the table scan require to scan data in desc order (#23948)](https://github.com/pingcap/tidb/pull/23974)   |   | 8d12h   |
| tidb/24033 | [statistics: fix some unstable tests in global stats (#23502)](https://github.com/pingcap/tidb/pull/24033)                                         |   | 6d9h    |
| tidb/24061 | [statistics: fix some potential panic in statistics (#23988)](https://github.com/pingcap/tidb/pull/24061)                                          |   | 5d13h   |
| tidb/24079 | [planner: change descScanFactor to scanFactor when ExpectedCount is small. (#23972)](https://github.com/pingcap/tidb/pull/24079)                   |   | 4d19h   |
| tidb/24102 | [planner: Fix Join reorder occurs "index out of range" error](https://github.com/pingcap/tidb/pull/24102)                                          |   | 2d0h    |
| tidb/24147 | [docs/design: add proposal for common table expression](https://github.com/pingcap/tidb/pull/24147)                                                |   | 23h     |
| tidb/24155 | [planner, executor: fix index merge partial table scan schema (#23936)](https://github.com/pingcap/tidb/pull/24155)                                |   | 19h     |


## Reviewed in Last 7 Days

|    REPO    |                                                                     PR                                                                     | C | D |   R    |
|------------|--------------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/24089 | [statistics: introduce the weighted reservoir sampling](https://github.com/pingcap/tidb/pull/24089)                                        |   | 1 | 3d23h  |
| tidb/23936 | [planner, executor: fix index merge partial table scan schema](https://github.com/pingcap/tidb/pull/23936)                                 |   | 2 | 9d18h  |
| tidb/24022 | [expression: don't propagateColumnEQ joinCondition when nullSensitive (#23989)](https://github.com/pingcap/tidb/pull/24022)                |   | 5 | 1d23h  |
| tidb/24050 | [expression: fix get var panic when types not match](https://github.com/pingcap/tidb/pull/24050)                                           |   | 5 | 23h    |
| tidb/24081 | [planner/core: push down topn to mpp](https://github.com/pingcap/tidb/pull/24081)                                                          |   | 5 | 1h     |
| tidb/24060 | [statistics: fix some potential panic in statistics (#23988)](https://github.com/pingcap/tidb/pull/24060)                                  |   | 5 | 19h    |
| tidb/23970 | [planner: fix a bug that point get plan returns wrong column name (#23365)](https://github.com/pingcap/tidb/pull/23970)                    |   | 6 | 2d20h  |
| tidb/24025 | [session, executor: skip some frequent unstable test cases (#24003)](https://github.com/pingcap/tidb/pull/24025)                           |   | 6 | 18h    |
| tidb/23969 | [planner: fix a bug that point get plan returns wrong column name (#23365)](https://github.com/pingcap/tidb/pull/23969)                    |   | 6 | 2d18h  |
| tidb/24023 | [expression: don't propagateColumnEQ joinCondition when nullSensitive (#23989)](https://github.com/pingcap/tidb/pull/24023)                |   | 6 | 17h    |
| tidb/23474 | [planner: fix inappropriate null flag of null constants (#23457)](https://github.com/pingcap/tidb/pull/23474)                              |   | 7 | 22d2h  |
| tidb/23487 | [planner: optimize count(distinct a) to count(a) if there is an unique key on a](https://github.com/pingcap/tidb/pull/23487)               | Y | 7 | 21d22h |
| tidb/23917 | [planner: fix wrong TableDual plans caused by comparing Binary and Bytes incorrectly (#23860)](https://github.com/pingcap/tidb/pull/23917) |   | 7 | 6d7h   |
| tidb/23918 | [planner: fix wrong TableDual plans caused by comparing Binary and Bytes incorrectly (#23860)](https://github.com/pingcap/tidb/pull/23918) |   | 7 | 6d7h   |
| tidb/23988 | [statistics: fix some potential panic in statistics](https://github.com/pingcap/tidb/pull/23988)                                           |   | 7 | 1d2h   |
| tidb/23989 | [expression: don't propagateColumnEQ joinCondition when nullSensitive](https://github.com/pingcap/tidb/pull/23989)                         |   | 7 | 23h    |


</details> 


<details> 
  <summary>fzhedu's detailed review report</summary> 

## To Be Reviewed

| REPO | PR | C | LASTED |
|------|----|---|--------|


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>ichn-hu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                             PR                                              | C | LASTED |
|------------|---------------------------------------------------------------------------------------------|---|--------|
| tidb/24026 | [types: fix type merge about bit type (#23857)](https://github.com/pingcap/tidb/pull/24026) |   | 6d14h  |


## Reviewed in Last 7 Days

|    REPO    |                                               PR                                                | C | D |   R   |
|------------|-------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/23867 | [expression: fix wrong flen infer for bit constant](https://github.com/pingcap/tidb/pull/23867) |   | 6 | 8d17h |
| tidb/23857 | [types: fix type merge about bit type](https://github.com/pingcap/tidb/pull/23857)              |   | 7 | 8d2h  |


</details> 


<details> 
  <summary>lzmhhh123's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                             PR                                                                              | C | LASTED  |
|------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20444 | [expression: add json_merge_patch](https://github.com/pingcap/tidb/pull/20444)                                                                              |   | 188d21h |
| tidb/20465 | [expression: add uuidShortFunction](https://github.com/pingcap/tidb/pull/20465)                                                                             |   | 187d19h |
| tidb/20642 | [executor: modify admin executors to support partitioned table with global index](https://github.com/pingcap/tidb/pull/20642)                               |   | 176d15h |
| tidb/20903 | [planner: fix confused and unnecessary double-projection in plans.](https://github.com/pingcap/tidb/pull/20903)                                             |   | 165d17h |
| tidb/21018 | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                              |   | 159d17h |
| tidb/21195 | [brie: integrate lightning to suport IMPORT statement](https://github.com/pingcap/tidb/pull/21195)                                                          |   | 148d22h |
| tidb/21334 | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                                               |   | 145d14h |
| tidb/21347 | [session: make rollback work on global variables](https://github.com/pingcap/tidb/pull/21347)                                                               |   | 144d19h |
| tidb/21487 | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                                               |   | 138d4h  |
| tidb/21641 | [executor: Fix pessimistic lock doesn't work on the partition table for subquery/joins](https://github.com/pingcap/tidb/pull/21641)                         |   | 131d18h |
| tidb/21651 | [planner: allow filter condition pushing down to IndexScan for prefix index](https://github.com/pingcap/tidb/pull/21651)                                    |   | 131d13h |
| tidb/22126 | [*: add `sys` schema, `sys.SCHEMA_UNUSED_INDEXES` view and `sys.SCHEMA_INDEX_USAGE` view](https://github.com/pingcap/tidb/pull/22126)                       |   | 110d19h |
| tidb/22361 | [table: fix insert into _tidb_rowid panic and rebase it if needed (#22062)](https://github.com/pingcap/tidb/pull/22361)                                     |   | 98d20h  |
| tidb/22372 | [executor: fix SelectForUpdate in decorrelated subquery under pessimistic mode](https://github.com/pingcap/tidb/pull/22372)                                 |   | 98d9h   |
| tidb/22478 | [planner, executor: fix query partition table with global unique index get wrong result](https://github.com/pingcap/tidb/pull/22478)                        |   | 89d13h  |
| tidb/22631 | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                                                                             |   | 79d23h  |
| tidb/22686 | [expression: support enum pushdown](https://github.com/pingcap/tidb/pull/22686)                                                                             |   | 76d22h  |
| tidb/22699 | [brie: add error info column and history backup/restore info in sql](https://github.com/pingcap/tidb/pull/22699)                                            |   | 76d16h  |
| tidb/23001 | [statistics: fix err check](https://github.com/pingcap/tidb/pull/23001)                                                                                     |   | 52d0h   |
| tidb/23022 | [executor: create PipelinedWindowExec based on current implementation and modify the windowProcessor interface](https://github.com/pingcap/tidb/pull/23022) |   | 50d18h  |
| tidb/23149 | [core: support left join and right join for join reorder](https://github.com/pingcap/tidb/pull/23149)                                                       |   | 45d12h  |
| tidb/23257 | [executor: group_concat aggr panic when session.group_concat_max_len is small (#23131)](https://github.com/pingcap/tidb/pull/23257)                         |   | 40d18h  |
| tidb/23283 | [util: optimize the performance of restore with db (#22910)](https://github.com/pingcap/tidb/pull/23283)                                                    |   | 39d17h  |
| tidb/23347 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                                        |   | 35d18h  |
| tidb/23348 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                                        |   | 35d18h  |
| tidb/23368 | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23368)                                       |   | 34d20h  |
| tidb/23373 | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                           |   | 34d19h  |
| tidb/23655 | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral (#23523)](https://github.com/pingcap/tidb/pull/23655)                    |   | 21d22h  |
| tidb/23660 | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23660)     |   | 21d20h  |
| tidb/23661 | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23661)     |   | 21d20h  |
| tidb/23703 | [expression: fix approx_percent panic on bit column (#23687)](https://github.com/pingcap/tidb/pull/23703)                                                   |   | 21d13h  |
| tidb/23705 | [executor: refineArgs() bug fix when compare int with very small decimal (#23694)](https://github.com/pingcap/tidb/pull/23705)                              |   | 21d13h  |
| tidb/23714 | [*:Support record statment_history table evicted info](https://github.com/pingcap/tidb/pull/23714)                                                          |   | 21d1h   |
| tidb/23756 | [planner: fix set not null flag for outer join (#23727)](https://github.com/pingcap/tidb/pull/23756)                                                        |   | 20d14h  |
| tidb/23760 | [collation: fix tidb panic when compare string with collation](https://github.com/pingcap/tidb/pull/23760)                                                  |   | 20d13h  |
| tidb/23812 | [executor, planner: fix collation for hash join building (#23770)](https://github.com/pingcap/tidb/pull/23812)                                              |   | 19d12h  |
| tidb/23940 | [config, ddl: allow auto inc columns in generated columns and expression indexes](https://github.com/pingcap/tidb/pull/23940)                               |   | 10d17h  |
| tidb/23968 | [statistics: fix unstable TestDropPartitionStats test](https://github.com/pingcap/tidb/pull/23968)                                                          |   | 8d14h   |
| tidb/23987 | [executor: Implements json_arrayagg function](https://github.com/pingcap/tidb/pull/23987)                                                                   |   | 7d18h   |
| tidb/24016 | [planner: fix index-out-of-range error when checking only_full_group_by (#23844)](https://github.com/pingcap/tidb/pull/24016)                               |   | 6d18h   |
| tidb/24018 | [ranger: fix the range construction behavior when the column's type is `YEAR` (#23559)](https://github.com/pingcap/tidb/pull/24018)                         |   | 6d18h   |
| tidb/24053 | [executor: fix wrong convert from bit to string when do projection (#23960)](https://github.com/pingcap/tidb/pull/24053)                                    |   | 5d16h   |
| tidb/24054 | [executor: fix wrong convert from bit to string when do projection (#23960)](https://github.com/pingcap/tidb/pull/24054)                                    |   | 5d16h   |
| tidb/24092 | [planner: do not prune all columns for Projection (#24024)](https://github.com/pingcap/tidb/pull/24092)                                                     |   | 4d16h   |
| tidb/24120 | [expression: Fix cast real, decimal to time](https://github.com/pingcap/tidb/pull/24120)                                                                    |   | 1d19h   |
| tidb/24139 | [executor: accelerate TestShowVar (#24131)](https://github.com/pingcap/tidb/pull/24139)                                                                     |   | 1d10h   |
| tidb/24155 | [planner, executor: fix index merge partial table scan schema (#23936)](https://github.com/pingcap/tidb/pull/24155)                                         |   | 19h     |
| tidb/24174 | [executor: make column default value being aware of NO_ZERO_IN_DATE](https://github.com/pingcap/tidb/pull/24174)                                            |   | 16h     |
| tidb/24178 | [planner/core: point get only work on TiKV](https://github.com/pingcap/tidb/pull/24178)                                                                     |   | 10h     |


## Reviewed in Last 7 Days

|    REPO    |                                                             PR                                                              | C | D |   R   |
|------------|-----------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/24125 | [telemetry: log when sending telemetry](https://github.com/pingcap/tidb/pull/24125)                                         |   | 1 | 22h   |
| tidb/24022 | [expression: don't propagateColumnEQ joinCondition when nullSensitive (#23989)](https://github.com/pingcap/tidb/pull/24022) |   | 2 | 4d21h |
| tidb/24098 | [executor: fix null div 0 for partition key expression is incorrect ](https://github.com/pingcap/tidb/pull/24098)           |   | 2 | 2d2h  |
| tidb/23936 | [planner, executor: fix index merge partial table scan schema](https://github.com/pingcap/tidb/pull/23936)                  |   | 2 | 9d17h |
| tidb/24093 | [planner: donot prune all columns for Projection (#24024)](https://github.com/pingcap/tidb/pull/24093)                      |   | 5 | 0h    |
| tidb/24024 | [planner: donot prune all columns for Projection](https://github.com/pingcap/tidb/pull/24024)                               |   | 5 | 1d20h |
| tidb/24023 | [expression: don't propagateColumnEQ joinCondition when nullSensitive (#23989)](https://github.com/pingcap/tidb/pull/24023) |   | 6 | 17h   |
| tikv/10018 | [copr: fix IN expr didn't handle unsigned/signed int properly (#9823)](https://github.com/tikv/tikv/pull/10018)             |   | 6 | 19h   |
| tidb/23960 | [executor: fix wrong convert from bit to string when do projection](https://github.com/pingcap/tidb/pull/23960)             |   | 7 | 1d21h |
| tidb/23989 | [expression: don't propagateColumnEQ joinCondition when nullSensitive](https://github.com/pingcap/tidb/pull/23989)          |   | 7 | 21h   |


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
| tidb/19029   | [types: fix unexpected NOT_NULL flags](https://github.com/pingcap/tidb/pull/19029)                                                                          |   | 257d22h |
| docs-cn/5561 | [Add sql optimization-related docs to toc](https://github.com/pingcap/docs-cn/pull/5561)                                                                    |   | 57d15h  |
| tidb/20708   | [*: separate auto_increment ID allocator from _tidb_rowid allocator](https://github.com/pingcap/tidb/pull/20708)                                            |   | 173d20h |
| tidb/20969   | [executor: Improve the performance of appending not fixed columns](https://github.com/pingcap/tidb/pull/20969)                                              |   | 161d9h  |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                              |   | 159d17h |
| tidb/21149   | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                               |   | 152d14h |
| tidb/21304   | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                                    |   | 146d12h |
| tidb/21318   | [planner, expression: use the range of column types to simplify expressions](https://github.com/pingcap/tidb/pull/21318)                                    |   | 145d19h |
| tidb/21401   | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                                          |   | 141d11h |
| tidb/21508   | [execution: fix dayofweek('0000-00-00') behavior](https://github.com/pingcap/tidb/pull/21508)                                                               |   | 137d10h |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                                 |   | 123d19h |
| tidb/21887   | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                                             |   | 122d11h |
| tidb/21954   | [planner/cascades: add rule `PushSelDownApply`](https://github.com/pingcap/tidb/pull/21954)                                                                 |   | 118d23h |
| tidb/22146   | [executor: forbid SFU on view](https://github.com/pingcap/tidb/pull/22146)                                                                                  |   | 106d21h |
| tidb/22217   | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                                               |   | 104d17h |
| tidb/22234   | [executor, planner: ON DUPLICATE UPDATE can refer to un-project col (#14412)](https://github.com/pingcap/tidb/pull/22234)                                   |   | 104d15h |
| tidb/22261   | [time: fix parse datetime won't truncate the reluctant string (#22232)](https://github.com/pingcap/tidb/pull/22261)                                         |   | 103d19h |
| tidb/22374   | [expression: separated arithmeticIntDivideSig](https://github.com/pingcap/tidb/pull/22374)                                                                  |   | 98d0h   |
| tidb/22415   | [ddl: refactor bundle[2/2] [6/6]](https://github.com/pingcap/tidb/pull/22415)                                                                               |   | 94d17h  |
| tidb/22416   | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                                                           | Y | 94d11h  |
| tidb/22541   | [expression: Support builtin function SOUNDEX](https://github.com/pingcap/tidb/pull/22541)                                                                  |   | 84d9h   |
| tidb/22565   | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)                                          | Y | 83d17h  |
| tidb/22814   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22814)                                         |   | 60d19h  |
| tidb/22862   | [brie: fix the problem that ddl restored by BR via SQL is not replicated to downstream](https://github.com/pingcap/tidb/pull/22862)                         |   | 57d22h  |
| tidb/22923   | [expression: correct constant propagation for collation (#22666)](https://github.com/pingcap/tidb/pull/22923)                                               |   | 55d15h  |
| tidb/22924   | [planner: fix wrong index merge selection (#22825)](https://github.com/pingcap/tidb/pull/22924)                                                             |   | 55d14h  |
| tidb/23002   | [store/*: fix err check](https://github.com/pingcap/tidb/pull/23002)                                                                                        |   | 52d0h   |
| tidb/23022   | [executor: create PipelinedWindowExec based on current implementation and modify the windowProcessor interface](https://github.com/pingcap/tidb/pull/23022) |   | 50d18h  |
| tidb/23152   | [expression: fix wrong error info (#22760)](https://github.com/pingcap/tidb/pull/23152)                                                                     |   | 44d14h  |
| tidb/23196   | [types: fix the bug about the wrong query result for decimal type  (#22507)](https://github.com/pingcap/tidb/pull/23196)                                    |   | 42d18h  |
| tidb/23295   | [util, types: don't let SPM be affected by charset (#23161)](https://github.com/pingcap/tidb/pull/23295)                                                    |   | 39d11h  |
| tidb/23316   | [planner: Fix rebuild range for prepared plan](https://github.com/pingcap/tidb/pull/23316)                                                                  |   | 36d17h  |
| tidb/23373   | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                           |   | 34d19h  |
| tidb/23397   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23397)                                                              |   | 33d17h  |
| tidb/23398   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23398)                                                              |   | 33d17h  |
| tidb/23590   | [planner, table: optimize the list partition pruner for range query](https://github.com/pingcap/tidb/pull/23590)                                            |   | 25d16h  |
| tidb/23598   | [types: fix collation for binary literal (#23591)](https://github.com/pingcap/tidb/pull/23598)                                                              |   | 25d13h  |
| tidb/23655   | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral (#23523)](https://github.com/pingcap/tidb/pull/23655)                    |   | 21d22h  |
| tidb/23658   | [*: collect transaction write duration/throughput metrics for SLI/SLO (#23462)](https://github.com/pingcap/tidb/pull/23658)                                 |   | 21d22h  |
| tidb/23660   | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23660)     |   | 21d20h  |
| tidb/23661   | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23661)     |   | 21d20h  |
| tidb/23682   | [executor: fix a panic when batch point get is used for partition table (#23652)](https://github.com/pingcap/tidb/pull/23682)                               |   | 21d16h  |
| tidb/23689   | [planner: fix the panic when we calculate the partition range (#23651)](https://github.com/pingcap/tidb/pull/23689)                                         |   | 21d16h  |
| tidb/23730   | [distsql/*: typo fix for `dispatches`](https://github.com/pingcap/tidb/pull/23730)                                                                          |   | 20d18h  |
| tidb/23796   | [tests: make TestIndexLookupMergeJoinHang and TestIssue18068 stable (#23741)](https://github.com/pingcap/tidb/pull/23796)                                   |   | 19d19h  |
| tidb/23812   | [executor, planner: fix collation for hash join building (#23770)](https://github.com/pingcap/tidb/pull/23812)                                              |   | 19d12h  |
| tidb/23867   | [expression: fix wrong flen infer for bit constant](https://github.com/pingcap/tidb/pull/23867)                                                             |   | 14d16h  |
| tidb/23878   | [functions: fix some string function has wrong collation and flag (#23835)](https://github.com/pingcap/tidb/pull/23878)                                     |   | 13d21h  |
| tidb/23963   | [executor: checking chunk is full precedes filtering](https://github.com/pingcap/tidb/pull/23963)                                                           |   | 8d17h   |
| tidb/23987   | [executor: Implements json_arrayagg function](https://github.com/pingcap/tidb/pull/23987)                                                                   |   | 7d18h   |
| tidb/23997   | [stats, executor: use a correct sampling to collect stats](https://github.com/pingcap/tidb/pull/23997)                                                      |   | 7d9h    |
| tidb/24018   | [ranger: fix the range construction behavior when the column's type is `YEAR` (#23559)](https://github.com/pingcap/tidb/pull/24018)                         |   | 6d18h   |
| tidb/24035   | [store/tikv: new config TxnScope in tikv config](https://github.com/pingcap/tidb/pull/24035)                                                                |   | 5d22h   |
| tidb/24089   | [statistics: introduce the weighted reservoir sampling](https://github.com/pingcap/tidb/pull/24089)                                                         |   | 4d17h   |
| tidb/24100   | [variable: refactor session/global validation (part 8)](https://github.com/pingcap/tidb/pull/24100)                                                         |   | 2d8h    |


## Reviewed in Last 7 Days

|    REPO    |                                                                    PR                                                                    | C | D |   R   |
|------------|------------------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/24063 | [expression: check NotNullFlag of columns when simplifying binary operations (#24049)](https://github.com/pingcap/tidb/pull/24063)       |   | 5 | 12h   |
| tidb/24041 | [util: fix range building for binary literal (#23699)](https://github.com/pingcap/tidb/pull/24041)                                       |   | 6 | 6h    |
| tidb/24043 | [planner: fix wrong PointGet / TableDual plan reused in plan cache (#23238)](https://github.com/pingcap/tidb/pull/24043)                 |   | 6 | 1h    |
| tidb/24042 | [planner: append common handle columns into the schema of index merge table plan (#23933)](https://github.com/pingcap/tidb/pull/24042)   |   | 6 | 0h    |
| tidb/23656 | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral (#23523)](https://github.com/pingcap/tidb/pull/23656) |   | 6 | 16d2h |
| tidb/24033 | [statistics: fix some unstable tests in global stats (#23502)](https://github.com/pingcap/tidb/pull/24033)                               |   | 6 | 9h    |
| tidb/24003 | [session, executor: skip some frequent unstable test cases](https://github.com/pingcap/tidb/pull/24003)                                  |   | 7 | 0h    |


</details> 


<details> 
  <summary>rebelice's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                           PR                                                           | C | LASTED |
|------------|------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/23537 | [planner: remove some risky cache operations in the plan builder (#23354)](https://github.com/pingcap/tidb/pull/23537) |   | 26d19h |
| tidb/23836 | [parser, core: Implement force_index hint in parser and TiDB](https://github.com/pingcap/tidb/pull/23836)              |   | 18d17h |
| tidb/24033 | [statistics: fix some unstable tests in global stats (#23502)](https://github.com/pingcap/tidb/pull/24033)             |   | 6d9h   |


## Reviewed in Last 7 Days

|    REPO    |                                        PR                                         | C | D | R  |
|------------|-----------------------------------------------------------------------------------|---|---|----|
| tidb/24081 | [planner/core: push down topn to mpp](https://github.com/pingcap/tidb/pull/24081) |   | 5 | 0h |


</details> 


<details> 
  <summary>time-and-fate's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                             PR                                                              | C | LASTED  |
|------------|-----------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                   |   | 166d16h |
| tidb/22416 | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                           | Y | 94d11h  |
| tidb/24155 | [planner, executor: fix index merge partial table scan schema (#23936)](https://github.com/pingcap/tidb/pull/24155)         |   | 19h     |
| tidb/24175 | [statistics: skip reading mysql.stats_histograms if cached stats is up-to-date](https://github.com/pingcap/tidb/pull/24175) |   | 16h     |


## Reviewed in Last 7 Days

|    REPO    |                                                                     PR                                                                     | C | D |   R   |
|------------|--------------------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/24089 | [statistics: introduce the weighted reservoir sampling](https://github.com/pingcap/tidb/pull/24089)                                        |   | 2 | 3d3h  |
| tidb/24063 | [expression: check NotNullFlag of columns when simplifying binary operations (#24049)](https://github.com/pingcap/tidb/pull/24063)         |   | 5 | 12h   |
| tidb/24058 | [planner: do not build MPP plan for scan with virtual columns (#23980)](https://github.com/pingcap/tidb/pull/24058)                        |   | 6 | 0h    |
| tidb/24049 | [expression: check NotNullFlag of columns when simplifying binary operations](https://github.com/pingcap/tidb/pull/24049)                  |   | 6 | 1h    |
| tidb/24043 | [planner: fix wrong PointGet / TableDual plan reused in plan cache (#23238)](https://github.com/pingcap/tidb/pull/24043)                   |   | 6 | 0h    |
| tidb/23656 | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral (#23523)](https://github.com/pingcap/tidb/pull/23656)   |   | 6 | 16d2h |
| tidb/23918 | [planner: fix wrong TableDual plans caused by comparing Binary and Bytes incorrectly (#23860)](https://github.com/pingcap/tidb/pull/23918) |   | 6 | 7d3h  |
| tidb/23895 | [statistics: add more tests for extended stats](https://github.com/pingcap/tidb/pull/23895)                                                |   | 7 | 7d3h  |
| tidb/23890 | [statistics: check step overflow when converting a range to points for estimation](https://github.com/pingcap/tidb/pull/23890)             |   | 7 | 7d1h  |


</details> 


<details> 
  <summary>winoros's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                              PR                                                                               | C | LASTED  |
|--------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                                        | Y | 221d13h |
| docs-cn/5916 | [sql-statements, information-schema: add `END_TIME` field for table `ANALYZE_STATUS`](https://github.com/pingcap/docs-cn/pull/5916)                           |   | 19d17h  |
| tidb/20765   | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                                             |   | 169d17h |
| tidb/20877   | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                                                     |   | 166d16h |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                                |   | 159d17h |
| tidb/21207   | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                                               |   | 148d19h |
| tidb/21487   | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                                                 |   | 138d4h  |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                                   |   | 123d19h |
| tidb/21954   | [planner/cascades: add rule `PushSelDownApply`](https://github.com/pingcap/tidb/pull/21954)                                                                   |   | 118d23h |
| tidb/22181   | [planner, expression: fix error when using IN combined with subquery (#22080)](https://github.com/pingcap/tidb/pull/22181)                                    |   | 105d17h |
| tidb/22416   | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                                                             | Y | 94d11h  |
| tidb/22504   | [*:Fix the fetchHotRegion bug that the count always zero](https://github.com/pingcap/tidb/pull/22504)                                                         |   | 86d19h  |
| tidb/22565   | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)                                            | Y | 83d17h  |
| tidb/22624   | [ planner: not pruning column used by union scan condition (#21640)](https://github.com/pingcap/tidb/pull/22624)                                              |   | 81d17h  |
| tidb/22923   | [expression: correct constant propagation for collation (#22666)](https://github.com/pingcap/tidb/pull/22923)                                                 |   | 55d15h  |
| tidb/23233   | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23233)                                                        |   | 41d18h  |
| tidb/23347   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                                          |   | 35d18h  |
| tidb/23348   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                                          |   | 35d18h  |
| tidb/23350   | [util/stringutil, util/ranger, planner: use hierarchical separators to simplify the parsing for info of EXPLAIN ](https://github.com/pingcap/tidb/pull/23350) |   | 35d17h  |
| tidb/23373   | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                             |   | 34d19h  |
| tidb/23474   | [planner: fix inappropriate null flag of null constants (#23457)](https://github.com/pingcap/tidb/pull/23474)                                                 |   | 28d18h  |
| tidb/23537   | [planner: remove some risky cache operations in the plan builder (#23354)](https://github.com/pingcap/tidb/pull/23537)                                        |   | 26d19h  |
| tidb/23543   | [statistics: fix auto analyze log information incomplete (#23522)](https://github.com/pingcap/tidb/pull/23543)                                                |   | 26d18h  |
| tidb/23598   | [types: fix collation for binary literal (#23591)](https://github.com/pingcap/tidb/pull/23598)                                                                |   | 25d13h  |
| tidb/23655   | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral (#23523)](https://github.com/pingcap/tidb/pull/23655)                      |   | 21d22h  |
| tidb/23689   | [planner: fix the panic when we calculate the partition range (#23651)](https://github.com/pingcap/tidb/pull/23689)                                           |   | 21d16h  |
| tidb/23772   | [tablecodec: fix text type decode for old row format (#23751)](https://github.com/pingcap/tidb/pull/23772)                                                    |   | 20d11h  |
| tidb/23849   | [ddl: tidb panic while query hash partition table with is null condition](https://github.com/pingcap/tidb/pull/23849)                                         |   | 15d13h  |
| tidb/23917   | [planner: fix wrong TableDual plans caused by comparing Binary and Bytes incorrectly (#23860)](https://github.com/pingcap/tidb/pull/23917)                    |   | 12d23h  |
| tidb/23946   | [planner: fix visit info for grant/revoke](https://github.com/pingcap/tidb/pull/23946)                                                                        |   | 9d6h    |
| tidb/23970   | [planner: fix a bug that point get plan returns wrong column name (#23365)](https://github.com/pingcap/tidb/pull/23970)                                       |   | 8d14h   |
| tidb/24018   | [ranger: fix the range construction behavior when the column's type is `YEAR` (#23559)](https://github.com/pingcap/tidb/pull/24018)                           |   | 6d18h   |
| tidb/24060   | [statistics: fix some potential panic in statistics (#23988)](https://github.com/pingcap/tidb/pull/24060)                                                     |   | 5d13h   |
| tidb/24061   | [statistics: fix some potential panic in statistics (#23988)](https://github.com/pingcap/tidb/pull/24061)                                                     |   | 5d13h   |
| tidb/24079   | [planner: change descScanFactor to scanFactor when ExpectedCount is small. (#23972)](https://github.com/pingcap/tidb/pull/24079)                              |   | 4d19h   |
| tidb/24092   | [planner: do not prune all columns for Projection (#24024)](https://github.com/pingcap/tidb/pull/24092)                                                       |   | 4d16h   |
| tidb/24097   | [planner: Remove redundant call to expression.ColumnSubstitute](https://github.com/pingcap/tidb/pull/24097)                                                   |   | 4d2h    |
| tidb/24138   | [planner: Add Equivalence Rules to Transform BinaryOptSubquery to ExistsSubquery](https://github.com/pingcap/tidb/pull/24138)                                 |   | 1d12h   |


## Reviewed in Last 7 Days

|     REPO     |                                                                          PR                                                                          | C | D |    R    |
|--------------|------------------------------------------------------------------------------------------------------------------------------------------------------|---|---|---------|
| tidb/24175   | [statistics: skip reading mysql.stats_histograms if cached stats is up-to-date](https://github.com/pingcap/tidb/pull/24175)                          |   | 1 | 2h      |
| docs/5392    | [releases: add tidb 5.0.1 release notes](https://github.com/pingcap/docs/pull/5392)                                                                  |   | 1 | 21h     |
| docs-cn/6061 | [releases: add tidb 5.0.1 release notes](https://github.com/pingcap/docs-cn/pull/6061)                                                               |   | 1 | 20h     |
| tidb/24093   | [planner: donot prune all columns for Projection (#24024)](https://github.com/pingcap/tidb/pull/24093)                                               |   | 5 | 0h      |
| tidb/23215   | [Privileges: fix delete privilege check wrongly (#22971)](https://github.com/pingcap/tidb/pull/23215)                                                |   | 5 | 37d21h  |
| tidb/21299   | [statistics: fix the panic when analyze with collation enabled (#21262)](https://github.com/pingcap/tidb/pull/21299)                                 |   | 5 | 141d20h |
| tidb/24078   | [planner: change descScanFactor to scanFactor when ExpectedCount is small. (#23972)](https://github.com/pingcap/tidb/pull/24078)                     |   | 5 | 1h      |
| tidb/23972   | [planner: change descScanFactor to scanFactor when rowCount is small.](https://github.com/pingcap/tidb/pull/23972)                                   |   | 5 | 3d17h   |
| tidb/24024   | [planner: donot prune all columns for Projection](https://github.com/pingcap/tidb/pull/24024)                                                        |   | 5 | 1d20h   |
| tidb/21408   | [statistics: fix a bug which causes panic when using the clustered index and the new collation (#21379)](https://github.com/pingcap/tidb/pull/21408) |   | 5 | 135d23h |
| tidb/24058   | [planner: do not build MPP plan for scan with virtual columns (#23980)](https://github.com/pingcap/tidb/pull/24058)                                  |   | 6 | 1h      |
| tidb/24041   | [util: fix range building for binary literal (#23699)](https://github.com/pingcap/tidb/pull/24041)                                                   |   | 6 | 6h      |
| tidb/24042   | [planner: append common handle columns into the schema of index merge table plan (#23933)](https://github.com/pingcap/tidb/pull/24042)               |   | 6 | 0h      |
| tidb/24025   | [session, executor: skip some frequent unstable test cases (#24003)](https://github.com/pingcap/tidb/pull/24025)                                     |   | 6 | 19h     |


</details> 


<details> 
  <summary>wshwsh12's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                 PR                                                                  | C | LASTED  |
|------------|-------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19807 | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                             |   | 228d10h |
| tidb/19957 | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                              | Y | 221d13h |
| tidb/21487 | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                       |   | 138d4h  |
| tidb/21887 | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                     |   | 122d11h |
| tidb/22378 | [executor: vectorize hash aggregate](https://github.com/pingcap/tidb/pull/22378)                                                    |   | 97d19h  |
| tidb/22686 | [expression: support enum pushdown](https://github.com/pingcap/tidb/pull/22686)                                                     |   | 76d22h  |
| tidb/23336 | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23336)            |   | 35d19h  |
| tidb/23347 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                |   | 35d18h  |
| tidb/23348 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                |   | 35d18h  |
| tidb/23368 | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23368)               |   | 34d20h  |
| tidb/23397 | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23397)                                      |   | 33d17h  |
| tidb/23398 | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23398)                                      |   | 33d17h  |
| tidb/23519 | [executor: check privilege before adding](https://github.com/pingcap/tidb/pull/23519)                                               |   | 27d0h   |
| tidb/23760 | [collation: fix tidb panic when compare string with collation](https://github.com/pingcap/tidb/pull/23760)                          |   | 20d13h  |
| tidb/23866 | [executor,kv: support timebounded staleness transaction](https://github.com/pingcap/tidb/pull/23866)                                |   | 14d16h  |
| tidb/23968 | [statistics: fix unstable TestDropPartitionStats test](https://github.com/pingcap/tidb/pull/23968)                                  |   | 8d14h   |
| tidb/23979 | [executor, statistics: fix unstable `TestAnalyzeIndexExtractTopN`](https://github.com/pingcap/tidb/pull/23979)                      |   | 7d22h   |
| tidb/24018 | [ranger: fix the range construction behavior when the column's type is `YEAR` (#23559)](https://github.com/pingcap/tidb/pull/24018) |   | 6d18h   |
| tidb/24033 | [statistics: fix some unstable tests in global stats (#23502)](https://github.com/pingcap/tidb/pull/24033)                          |   | 6d9h    |
| tidb/24050 | [expression: fix get var panic when types not match](https://github.com/pingcap/tidb/pull/24050)                                    |   | 5d17h   |
| tidb/24053 | [executor: fix wrong convert from bit to string when do projection (#23960)](https://github.com/pingcap/tidb/pull/24053)            |   | 5d16h   |
| tidb/24054 | [executor: fix wrong convert from bit to string when do projection (#23960)](https://github.com/pingcap/tidb/pull/24054)            |   | 5d16h   |
| tidb/24147 | [docs/design: add proposal for common table expression](https://github.com/pingcap/tidb/pull/24147)                                 |   | 23h     |
| tidb/24177 | [executor: accelerate TestVectorizedMergeJoin](https://github.com/pingcap/tidb/pull/24177)                                          |   | 14h     |


## Reviewed in Last 7 Days

|    REPO    |                                                                          PR                                                                          | C | D |   R    |
|------------|------------------------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/23876 | [executor: fix scope ambiguity of joinResult](https://github.com/pingcap/tidb/pull/23876)                                                            |   | 1 | 14d1h  |
| tidb/24139 | [executor: accelerate TestShowVar (#24131)](https://github.com/pingcap/tidb/pull/24139)                                                              |   | 1 | 1d0h   |
| tidb/24132 | [executor: accelerate TestUpdateScanningHandles and TestIssue20658 and TestParallelStreamAggGroupConcat](https://github.com/pingcap/tidb/pull/24132) |   | 1 | 14h    |
| tidb/24059 | [executor: bypass encoding invalid datetime for index join (#24051)](https://github.com/pingcap/tidb/pull/24059)                                     |   | 5 | 20h    |
| tidb/24051 | [executor: bypass encoding invalid datetime for index join](https://github.com/pingcap/tidb/pull/24051)                                              |   | 6 | 0h     |
| tidb/24019 | [tests: fix the graceshutdown e2e that is failing](https://github.com/pingcap/tidb/pull/24019)                                                       |   | 7 | 0h     |
| tidb/22686 | [expression: support enum pushdown](https://github.com/pingcap/tidb/pull/22686)                                                                      |   | 7 | 69d22h |


</details> 

