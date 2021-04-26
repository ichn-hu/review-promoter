|   REVIEWER    |  O(C)   | 1 | 2 | 3 | 4 | 5 | 6 | 7 | T  |
|---------------|---------|---|---|---|---|---|---|---|----|
| Reminiscent   |  6 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| SunRunAway    | 12 ( 1) | 0 | 0 | 1 | 0 | 0 | 0 | 0 |  1 |
| XuHuaiyu      | 86 ( 2) | 1 | 0 | 2 | 0 | 0 | 0 | 1 |  4 |
| dyzsr         |  1 ( 0) | 0 | 0 | 1 | 0 | 0 | 0 | 0 |  1 |
| eurekaka      | 20 ( 0) | 1 | 0 | 0 | 1 | 4 | 0 | 1 |  7 |
| fzhedu        |  0 ( 0) | 1 | 0 | 1 | 0 | 0 | 0 | 0 |  2 |
| ichn-hu       |  3 ( 0) | 3 | 0 | 2 | 0 | 0 | 0 | 0 |  5 |
| lzmhhh123     | 52 ( 0) | 2 | 0 | 2 | 6 | 2 | 1 | 3 | 16 |
| nullnotnil    |  0 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| qw4990        | 63 ( 2) | 0 | 0 | 0 | 0 | 2 | 0 | 0 |  2 |
| rebelice      |  3 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| time-and-fate |  5 ( 1) | 1 | 0 | 1 | 2 | 2 | 0 | 0 |  6 |
| winoros       | 41 ( 3) | 0 | 0 | 0 | 2 | 0 | 2 | 0 |  4 |
| wshwsh12      | 35 ( 1) | 0 | 0 | 4 | 0 | 0 | 2 | 1 |  7 |


<details> 
  <summary>Reminiscent's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                     PR                                                                     | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/21896 | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                                 |   | 125d19h |
| tidb/23474 | [planner: fix inappropriate null flag of null constants (#23457)](https://github.com/pingcap/tidb/pull/23474)                              |   | 33d18h  |
| tidb/23575 | [executor: fix update panic on join having statement (#23554)](https://github.com/pingcap/tidb/pull/23575)                                 |   | 30d21h  |
| tidb/23917 | [planner: fix wrong TableDual plans caused by comparing Binary and Bytes incorrectly (#23860)](https://github.com/pingcap/tidb/pull/23917) |   | 17d23h  |
| tidb/24016 | [planner: fix index-out-of-range error when checking only_full_group_by (#23844)](https://github.com/pingcap/tidb/pull/24016)              |   | 11d18h  |
| tidb/24188 | [util: fix bad number error with DISTINCT when dividing long decimals (#21783)](https://github.com/pingcap/tidb/pull/24188)                |   | 4d16h   |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>SunRunAway's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                  PR                                                                   | C | LASTED  |
|------------|---------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19178 | [executor: Refactor probe channel](https://github.com/pingcap/tidb/pull/19178)                                                        |   | 255d16h |
| tidb/19807 | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                               |   | 233d11h |
| tidb/19900 | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                        | Y | 228d18h |
| tidb/20140 | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                    |   | 215d22h |
| tidb/20765 | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                     |   | 174d17h |
| tidb/21207 | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                       |   | 153d19h |
| tidb/21834 | [planner: enhanced index range calculation plan](https://github.com/pingcap/tidb/pull/21834)                                          |   | 130d18h |
| tidb/21876 | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)           |   | 128d19h |
| tidb/21878 | [planner: do not push down lock to pointGet/bacthPointGet when selection exists](https://github.com/pingcap/tidb/pull/21878)          |   | 128d18h |
| tidb/21956 | [planner/preprocessor: disallow into-outfile clause in some place](https://github.com/pingcap/tidb/pull/21956)                        |   | 123d23h |
| tidb/22217 | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                         |   | 109d17h |
| tidb/22379 | [[experiment] executor: allow aggregation to spill disk when running out of memory quota](https://github.com/pingcap/tidb/pull/22379) |   | 102d19h |


## Reviewed in Last 7 Days

|    REPO    |                                          PR                                           | C | D |   R    |
|------------|---------------------------------------------------------------------------------------|---|---|--------|
| tidb/20749 | [executor: support global kill (32 bits)](https://github.com/pingcap/tidb/pull/20749) |   | 3 | 172d5h |


</details> 


<details> 
  <summary>XuHuaiyu's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                              PR                                                                               | C | LASTED  |
|--------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| docs-cn/5619 | [Update data-type-date-and-time.md](https://github.com/pingcap/docs-cn/pull/5619)                                                                             |   | 58d16h  |
| docs/5409    | [Add details for Hexadecimal Literals](https://github.com/pingcap/docs/pull/5409)                                                                             |   | 4d23h   |
| tidb/19900   | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                                                | Y | 228d18h |
| docs-cn/5671 | [tidb: Add time format description](https://github.com/pingcap/docs-cn/pull/5671)                                                                             |   | 52d11h  |
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                                        | Y | 226d14h |
| tidb/20140   | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                                            |   | 215d22h |
| tidb/20749   | [executor: support global kill (32 bits)](https://github.com/pingcap/tidb/pull/20749)                                                                         |   | 175d2h  |
| tidb/20790   | [collation: add pinyin collation for chinese charset support](https://github.com/pingcap/tidb/pull/20790)                                                     |   | 173d21h |
| tidb/20969   | [executor: Improve the performance of appending not fixed columns](https://github.com/pingcap/tidb/pull/20969)                                                |   | 166d9h  |
| tidb/21064   | [planner, executor: fix cast not check error](https://github.com/pingcap/tidb/pull/21064)                                                                     |   | 161d8h  |
| tidb/21149   | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                                 |   | 157d14h |
| tidb/21228   | [executor: return the result immediately when combining LIMIT row_count with DISTINCT](https://github.com/pingcap/tidb/pull/21228)                            |   | 153d13h |
| tidb/21304   | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                                      |   | 151d12h |
| tidb/21334   | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                                                 |   | 150d14h |
| tidb/21401   | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                                            |   | 146d11h |
| tidb/21536   | [executor: add slow-log file meta cache to avoid repeat read file meta information](https://github.com/pingcap/tidb/pull/21536)                               |   | 139d15h |
| tidb/21564   | [ddl: fix Incorrect behavior of NO_ZERO_DATE when altering table](https://github.com/pingcap/tidb/pull/21564)                                                 |   | 138d15h |
| tidb/21896   | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                                                    |   | 125d19h |
| tidb/22131   | [privilege: remove leading and trailing space when create user and role](https://github.com/pingcap/tidb/pull/22131)                                          |   | 115d19h |
| tidb/22163   | [expression: separated arithmeticMinusIntSig](https://github.com/pingcap/tidb/pull/22163)                                                                     |   | 111d13h |
| tidb/22186   | [executor: fix select into outfile with year type column has no data (#22175)](https://github.com/pingcap/tidb/pull/22186)                                    |   | 110d16h |
| tidb/22616   | [expression: from_unixtime accept 64-bit integers](https://github.com/pingcap/tidb/pull/22616)                                                                |   | 86d23h  |
| tidb/22617   | [metrics: fix wrong bucket name of coprocessor cache (#22454)](https://github.com/pingcap/tidb/pull/22617)                                                    |   | 86d23h  |
| tidb/22624   | [ planner: not pruning column used by union scan condition (#21640)](https://github.com/pingcap/tidb/pull/22624)                                              |   | 86d17h  |
| tidb/22631   | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                                                                               |   | 84d23h  |
| tidb/22696   | [expression: enable arithmetic Mod push down](https://github.com/pingcap/tidb/pull/22696)                                                                     |   | 81d17h  |
| tidb/22711   | [executor: Fix inline schema name](https://github.com/pingcap/tidb/pull/22711)                                                                                |   | 81d11h  |
| tidb/22722   | [planner, errno: make error code of ErrMixOfGroupFuncAndFields consistent with MySQL](https://github.com/pingcap/tidb/pull/22722)                             |   | 80d20h  |
| tidb/22814   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22814)                                           |   | 65d19h  |
| tidb/22908   | [txn: Add txn state's view](https://github.com/pingcap/tidb/pull/22908)                                                                                       |   | 60d20h  |
| tidb/23012   | [executor: fix affected rows of ddls and complete uint tests](https://github.com/pingcap/tidb/pull/23012)                                                     |   | 56d17h  |
| tidb/23152   | [expression: fix wrong error info (#22760)](https://github.com/pingcap/tidb/pull/23152)                                                                       |   | 49d14h  |
| tidb/23196   | [types: fix the bug about the wrong query result for decimal type  (#22507)](https://github.com/pingcap/tidb/pull/23196)                                      |   | 47d18h  |
| tidb/23220   | [Release 4.0](https://github.com/pingcap/tidb/pull/23220)                                                                                                     |   | 47d11h  |
| tidb/23233   | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23233)                                                        |   | 46d18h  |
| tidb/23257   | [executor: group_concat aggr panic when session.group_concat_max_len is small (#23131)](https://github.com/pingcap/tidb/pull/23257)                           |   | 45d18h  |
| tidb/23295   | [util, types: don't let SPM be affected by charset (#23161)](https://github.com/pingcap/tidb/pull/23295)                                                      |   | 44d11h  |
| tidb/23335   | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23335)                                      |   | 40d19h  |
| tidb/23336   | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23336)                                      |   | 40d19h  |
| tidb/23347   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                                          |   | 40d18h  |
| tidb/23348   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                                          |   | 40d18h  |
| tidb/23350   | [util/stringutil, util/ranger, planner: use hierarchical separators to simplify the parsing for info of EXPLAIN ](https://github.com/pingcap/tidb/pull/23350) |   | 40d17h  |
| tidb/23368   | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23368)                                         |   | 39d20h  |
| tidb/23397   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23397)                                                                |   | 38d17h  |
| tidb/23398   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23398)                                                                |   | 38d17h  |
| tidb/23405   | [domain: remove the exit chan, use context](https://github.com/pingcap/tidb/pull/23405)                                                                       |   | 38d17h  |
| tidb/23433   | [WIP: speed up for slow query logs retrieving ](https://github.com/pingcap/tidb/pull/23433)                                                                   |   | 37d17h  |
| tidb/23474   | [planner: fix inappropriate null flag of null constants (#23457)](https://github.com/pingcap/tidb/pull/23474)                                                 |   | 33d18h  |
| tidb/23497   | [expression: Let TiDB use Hyperscan to support multi-pattern-match](https://github.com/pingcap/tidb/pull/23497)                                               |   | 32d22h  |
| tidb/23517   | [*: Add the metric about the SQL with TiFlash Success  (#23426)](https://github.com/pingcap/tidb/pull/23517)                                                  |   | 32d12h  |
| tidb/23562   | [execution: reuse iterator in hash join](https://github.com/pingcap/tidb/pull/23562)                                                                          |   | 31d13h  |
| tidb/23640   | [*: fix the bug about YEAR(0.9) returns NULL instead of 0 in NO_ZERO_DATE mode](https://github.com/pingcap/tidb/pull/23640)                                   |   | 27d13h  |
| tidb/23661   | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23661)       |   | 26d20h  |
| tidb/23682   | [executor: fix a panic when batch point get is used for partition table (#23652)](https://github.com/pingcap/tidb/pull/23682)                                 |   | 26d16h  |
| tidb/23691   | [executor: fix index join on prefix column index (#23678)](https://github.com/pingcap/tidb/pull/23691)                                                        |   | 26d15h  |
| tidb/23705   | [executor: refineArgs() bug fix when compare int with very small decimal (#23694)](https://github.com/pingcap/tidb/pull/23705)                                |   | 26d13h  |
| tidb/23756   | [planner: fix set not null flag for outer join (#23727)](https://github.com/pingcap/tidb/pull/23756)                                                          |   | 25d14h  |
| tidb/23812   | [executor, planner: fix collation for hash join building (#23770)](https://github.com/pingcap/tidb/pull/23812)                                                |   | 24d12h  |
| tidb/23878   | [functions: fix some string function has wrong collation and flag (#23835)](https://github.com/pingcap/tidb/pull/23878)                                       |   | 18d21h  |
| tidb/23884   | [Metric: Collect TiKV Read Metric for SLI/SLO](https://github.com/pingcap/tidb/pull/23884)                                                                    |   | 18d20h  |
| tidb/23888   | [executor: fix resource leak of Shuffle Executor.](https://github.com/pingcap/tidb/pull/23888)                                                                |   | 18d18h  |
| tidb/23958   | [executor: fix `show table status` for the database with upper-cased name (#23896)](https://github.com/pingcap/tidb/pull/23958)                               |   | 13d18h  |
| tidb/23964   | [executor: GROUP_CONCAT(float) is not compatible with mysql](https://github.com/pingcap/tidb/pull/23964)                                                      |   | 13d16h  |
| tidb/24007   | [ddl: refactor rule [4/6]](https://github.com/pingcap/tidb/pull/24007)                                                                                        |   | 11d20h  |
| tidb/24016   | [planner: fix index-out-of-range error when checking only_full_group_by (#23844)](https://github.com/pingcap/tidb/pull/24016)                                 |   | 11d18h  |
| tidb/24026   | [types: fix type merge about bit type (#23857)](https://github.com/pingcap/tidb/pull/24026)                                                                   |   | 11d14h  |
| tidb/24033   | [statistics: fix some unstable tests in global stats (#23502)](https://github.com/pingcap/tidb/pull/24033)                                                    |   | 11d9h   |
| tidb/24053   | [executor: fix wrong convert from bit to string when do projection (#23960)](https://github.com/pingcap/tidb/pull/24053)                                      |   | 10d16h  |
| tidb/24060   | [statistics: fix some potential panic in statistics (#23988)](https://github.com/pingcap/tidb/pull/24060)                                                     |   | 10d13h  |
| tidb/24061   | [statistics: fix some potential panic in statistics (#23988)](https://github.com/pingcap/tidb/pull/24061)                                                     |   | 10d13h  |
| tidb/24078   | [planner: change descScanFactor to scanFactor when ExpectedCount is small. (#23972)](https://github.com/pingcap/tidb/pull/24078)                              |   | 9d19h   |
| tidb/24079   | [planner: change descScanFactor to scanFactor when ExpectedCount is small. (#23972)](https://github.com/pingcap/tidb/pull/24079)                              |   | 9d19h   |
| tidb/24155   | [planner, executor: fix index merge partial table scan schema (#23936)](https://github.com/pingcap/tidb/pull/24155)                                           |   | 5d20h   |
| tidb/24157   | [planner: let CopTiFlashConcurrencyFactor inflence the cost of whole plan](https://github.com/pingcap/tidb/pull/24157)                                        |   | 5d19h   |
| tidb/24179   | [expression: fix float64 overflow check in plus/minus real function](https://github.com/pingcap/tidb/pull/24179)                                              |   | 4d22h   |
| tidb/24196   | [*: support read and write operations for the global temporary table](https://github.com/pingcap/tidb/pull/24196)                                             |   | 3d23h   |
| tidb/24216   | [store/tikv: move kv.TxnInfo to store/tikv](https://github.com/pingcap/tidb/pull/24216)                                                                       |   | 3d10h   |
| tidb/24228   | [executor: skip TestPrepareStmtAfterIsolationReadChange when race enable (#24200)](https://github.com/pingcap/tidb/pull/24228)                                |   | 2d22h   |
| tidb/24229   | [executor: speed up race test TestInsertReorgDelete (#24208)](https://github.com/pingcap/tidb/pull/24229)                                                     |   | 2d21h   |
| tidb/24234   | [executor: skip TestMppExecution when race is enabled (#24222)](https://github.com/pingcap/tidb/pull/24234)                                                   |   | 2d18h   |
| tidb/24241   | [planner/core: remove random test to reduce CI time (#24207)](https://github.com/pingcap/tidb/pull/24241)                                                     |   | 2d15h   |
| tidb/24247   | [variable: change SetSessionSystemVar to accept string](https://github.com/pingcap/tidb/pull/24247)                                                           |   | 2d9h    |
| tidb/24261   | [executor: make IndexLookUps in the inner side of IndexJoins support direct reading](https://github.com/pingcap/tidb/pull/24261)                              |   | 19h     |
| tidb/24266   | [expression: fix wrong flen infer for bit constant (#23867)](https://github.com/pingcap/tidb/pull/24266)                                                      |   | 18h     |
| tidb/24267   | [expression: fix wrong flen infer for bit constant (#23867)](https://github.com/pingcap/tidb/pull/24267)                                                      |   | 17h     |
| tidb/24278   | [executor: accelerate TestVectorizedMergeJoin and TestVectorizedShuffleMergeJoin (#24177)](https://github.com/pingcap/tidb/pull/24278)                        |   | 10h     |


## Reviewed in Last 7 Days

|    REPO    |                                                                          PR                                                                          | C | D |   R   |
|------------|------------------------------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/24257 | [executor: fix data race of parallel apply operator](https://github.com/pingcap/tidb/pull/24257)                                                     |   | 1 | 0h    |
| docs/5445  | [tidb-config: add value range and type for token-limit](https://github.com/pingcap/docs/pull/5445)                                                   |   | 3 | 0h    |
| tidb/24177 | [executor: accelerate TestVectorizedMergeJoin and TestVectorizedShuffleMergeJoin](https://github.com/pingcap/tidb/pull/24177)                        |   | 3 | 2d14h |
| tidb/24132 | [executor: accelerate TestUpdateScanningHandles and TestIssue20658 and TestParallelStreamAggGroupConcat](https://github.com/pingcap/tidb/pull/24132) |   | 7 | 3h    |


</details> 


<details> 
  <summary>dyzsr's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                 PR                                                                  | C | LASTED |
|------------|-------------------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/24018 | [ranger: fix the range construction behavior when the column's type is `YEAR` (#23559)](https://github.com/pingcap/tidb/pull/24018) |   | 11d18h |


## Reviewed in Last 7 Days

|    REPO    |                                      PR                                       | C | D | R  |
|------------|-------------------------------------------------------------------------------|---|---|----|
| tikv/10074 | [copr: fix unsound unsafe transmute](https://github.com/tikv/tikv/pull/10074) |   | 3 | 2h |


</details> 


<details> 
  <summary>eurekaka's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                         PR                                                                         | C | LASTED  |
|------------|----------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                                          |   | 171d17h |
| tidb/23002 | [store/*: fix err check](https://github.com/pingcap/tidb/pull/23002)                                                                               |   | 57d0h   |
| tidb/23283 | [util: optimize the performance of restore with db (#22910)](https://github.com/pingcap/tidb/pull/23283)                                           |   | 44d17h  |
| tidb/23316 | [planner: Fix rebuild range for prepared plan](https://github.com/pingcap/tidb/pull/23316)                                                         |   | 41d17h  |
| tidb/23373 | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                  |   | 39d19h  |
| tidb/23543 | [statistics: fix auto analyze log information incomplete (#23522)](https://github.com/pingcap/tidb/pull/23543)                                     |   | 31d18h  |
| tidb/23689 | [planner: fix the panic when we calculate the partition range (#23651)](https://github.com/pingcap/tidb/pull/23689)                                |   | 26d16h  |
| tidb/23705 | [executor: refineArgs() bug fix when compare int with very small decimal (#23694)](https://github.com/pingcap/tidb/pull/23705)                     |   | 26d13h  |
| tidb/23756 | [planner: fix set not null flag for outer join (#23727)](https://github.com/pingcap/tidb/pull/23756)                                               |   | 25d14h  |
| tidb/23760 | [collation: fix tidb panic when compare string with collation](https://github.com/pingcap/tidb/pull/23760)                                         |   | 25d13h  |
| tidb/23938 | [planner,privilege: requires extra privileges for REPLACE and INSERT ON DUPLICATE statements (#23911)](https://github.com/pingcap/tidb/pull/23938) |   | 16d10h  |
| tidb/23974 | [planner: do not push down to TiFlash if the table scan require to scan data in desc order (#23948)](https://github.com/pingcap/tidb/pull/23974)   |   | 13d12h  |
| tidb/24033 | [statistics: fix some unstable tests in global stats (#23502)](https://github.com/pingcap/tidb/pull/24033)                                         |   | 11d9h   |
| tidb/24061 | [statistics: fix some potential panic in statistics (#23988)](https://github.com/pingcap/tidb/pull/24061)                                          |   | 10d13h  |
| tidb/24079 | [planner: change descScanFactor to scanFactor when ExpectedCount is small. (#23972)](https://github.com/pingcap/tidb/pull/24079)                   |   | 9d19h   |
| tidb/24147 | [docs/design: add proposal for common table expression](https://github.com/pingcap/tidb/pull/24147)                                                |   | 5d23h   |
| tidb/24155 | [planner, executor: fix index merge partial table scan schema (#23936)](https://github.com/pingcap/tidb/pull/24155)                                |   | 5d20h   |
| tidb/24214 | [plan: merge continuous selections and delete surely true expressions](https://github.com/pingcap/tidb/pull/24214)                                 |   | 3d12h   |
| tidb/24236 | [*: remove SchemaVersion in TransactionContext](https://github.com/pingcap/tidb/pull/24236)                                                        |   | 2d17h   |
| tidb/24258 | [Revert "planner: donot prune all columns for Projection (#24024)" (#24180)](https://github.com/pingcap/tidb/pull/24258)                           |   | 22h     |


## Reviewed in Last 7 Days

|    REPO     |                                                       PR                                                        | C | D |   R   |
|-------------|-----------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/24180  | [Revert "planner: donot prune all columns for Projection (#24024)"](https://github.com/pingcap/tidb/pull/24180) |   | 1 | 3d23h |
| docs/5392   | [releases: add tidb 5.0.1 release notes](https://github.com/pingcap/docs/pull/5392)                             |   | 4 | 2d23h |
| blog-cn/566 | [fix: remove the out-date description](https://github.com/pingcap/blog-cn/pull/566)                             |   | 5 | 1d4h  |
| tidb/24102  | [planner: Fix Join reorder occurs "index out of range" error](https://github.com/pingcap/tidb/pull/24102)       |   | 5 | 2d6h  |
| tipb/220    | [analyze: add proto fields for row based sampling](https://github.com/pingcap/tipb/pull/220)                    |   | 5 | 5d23h |
| tidb/24089  | [statistics: introduce the weighted reservoir sampling](https://github.com/pingcap/tidb/pull/24089)             |   | 5 | 4d23h |
| tidb/23936  | [planner, executor: fix index merge partial table scan schema](https://github.com/pingcap/tidb/pull/23936)      |   | 7 | 9d18h |


</details> 


<details> 
  <summary>fzhedu's detailed review report</summary> 

## To Be Reviewed

| REPO | PR | C | LASTED |
|------|----|---|--------|


## Reviewed in Last 7 Days

|    REPO    |                                                           PR                                                           | C | D |  R   |
|------------|------------------------------------------------------------------------------------------------------------------------|---|---|------|
| tidb/24231 | [executor: fix projection executor panic and add failpoint test](https://github.com/pingcap/tidb/pull/24231)           |   | 1 | 2d1h |
| tidb/24157 | [planner: let CopTiFlashConcurrencyFactor inflence the cost of whole plan](https://github.com/pingcap/tidb/pull/24157) |   | 3 | 3d3h |


</details> 


<details> 
  <summary>ichn-hu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                   PR                                                                   | C | LASTED |
|------------|----------------------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/24266 | [expression: fix wrong flen infer for bit constant (#23867)](https://github.com/pingcap/tidb/pull/24266)                               |   | 18h    |
| tidb/24267 | [expression: fix wrong flen infer for bit constant (#23867)](https://github.com/pingcap/tidb/pull/24267)                               |   | 17h    |
| tidb/24278 | [executor: accelerate TestVectorizedMergeJoin and TestVectorizedShuffleMergeJoin (#24177)](https://github.com/pingcap/tidb/pull/24278) |   | 10h    |


## Reviewed in Last 7 Days

|    REPO    |                                                              PR                                                               | C | D |   R   |
|------------|-------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/24177 | [executor: accelerate TestVectorizedMergeJoin and TestVectorizedShuffleMergeJoin](https://github.com/pingcap/tidb/pull/24177) |   | 1 | 4d15h |
| tidb/24234 | [executor: skip TestMppExecution when race is enabled (#24222)](https://github.com/pingcap/tidb/pull/24234)                   |   | 1 | 1d19h |
| tidb/24248 | [executor, variable: small cleanup](https://github.com/pingcap/tidb/pull/24248)                                               |   | 1 | 1d8h  |
| tidb/24235 | [expression: try to fix TestExprPushDownToFlash tests](https://github.com/pingcap/tidb/pull/24235)                            |   | 3 | 0h    |
| tidb/24026 | [types: fix type merge about bit type (#23857)](https://github.com/pingcap/tidb/pull/24026)                                   |   | 3 | 8d19h |


</details> 


<details> 
  <summary>lzmhhh123's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                           PR                                                                            | C | LASTED  |
|------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20444 | [expression: add json_merge_patch](https://github.com/pingcap/tidb/pull/20444)                                                                          |   | 193d21h |
| tidb/20465 | [expression: add uuidShortFunction](https://github.com/pingcap/tidb/pull/20465)                                                                         |   | 192d19h |
| tidb/20642 | [executor: modify admin executors to support partitioned table with global index](https://github.com/pingcap/tidb/pull/20642)                           |   | 181d15h |
| tidb/20903 | [planner: fix confused and unnecessary double-projection in plans.](https://github.com/pingcap/tidb/pull/20903)                                         |   | 170d17h |
| tidb/21018 | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                          |   | 164d17h |
| tidb/21195 | [brie: integrate lightning to suport IMPORT statement](https://github.com/pingcap/tidb/pull/21195)                                                      |   | 153d22h |
| tidb/21334 | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                                           |   | 150d14h |
| tidb/21347 | [session: make rollback work on global variables](https://github.com/pingcap/tidb/pull/21347)                                                           |   | 149d19h |
| tidb/21487 | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                                           |   | 143d4h  |
| tidb/21641 | [executor: Fix pessimistic lock doesn't work on the partition table for subquery/joins](https://github.com/pingcap/tidb/pull/21641)                     |   | 136d18h |
| tidb/21651 | [planner: allow filter condition pushing down to IndexScan for prefix index](https://github.com/pingcap/tidb/pull/21651)                                |   | 136d13h |
| tidb/22126 | [*: add `sys` schema, `sys.SCHEMA_UNUSED_INDEXES` view and `sys.SCHEMA_INDEX_USAGE` view](https://github.com/pingcap/tidb/pull/22126)                   |   | 115d20h |
| tidb/22361 | [table: fix insert into _tidb_rowid panic and rebase it if needed (#22062)](https://github.com/pingcap/tidb/pull/22361)                                 |   | 103d20h |
| tidb/22372 | [executor: fix SelectForUpdate in decorrelated subquery under pessimistic mode](https://github.com/pingcap/tidb/pull/22372)                             |   | 103d9h  |
| tidb/22478 | [planner, executor: fix query partition table with global unique index get wrong result](https://github.com/pingcap/tidb/pull/22478)                    |   | 94d13h  |
| tidb/22631 | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                                                                         |   | 84d23h  |
| tidb/22699 | [brie: add error info column and history backup/restore info in sql](https://github.com/pingcap/tidb/pull/22699)                                        |   | 81d16h  |
| tidb/23022 | [executor: create PipelinedWindowExec](https://github.com/pingcap/tidb/pull/23022)                                                                      |   | 55d18h  |
| tidb/23149 | [core: support left join and right join for join reorder](https://github.com/pingcap/tidb/pull/23149)                                                   |   | 50d12h  |
| tidb/23257 | [executor: group_concat aggr panic when session.group_concat_max_len is small (#23131)](https://github.com/pingcap/tidb/pull/23257)                     |   | 45d18h  |
| tidb/23283 | [util: optimize the performance of restore with db (#22910)](https://github.com/pingcap/tidb/pull/23283)                                                |   | 44d17h  |
| tidb/23347 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                                    |   | 40d18h  |
| tidb/23348 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                                    |   | 40d18h  |
| tidb/23368 | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23368)                                   |   | 39d20h  |
| tidb/23373 | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                       |   | 39d19h  |
| tidb/23655 | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral (#23523)](https://github.com/pingcap/tidb/pull/23655)                |   | 26d22h  |
| tidb/23660 | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23660) |   | 26d20h  |
| tidb/23661 | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23661) |   | 26d20h  |
| tidb/23703 | [expression: fix approx_percent panic on bit column (#23687)](https://github.com/pingcap/tidb/pull/23703)                                               |   | 26d14h  |
| tidb/23705 | [executor: refineArgs() bug fix when compare int with very small decimal (#23694)](https://github.com/pingcap/tidb/pull/23705)                          |   | 26d13h  |
| tidb/23756 | [planner: fix set not null flag for outer join (#23727)](https://github.com/pingcap/tidb/pull/23756)                                                    |   | 25d14h  |
| tidb/23760 | [collation: fix tidb panic when compare string with collation](https://github.com/pingcap/tidb/pull/23760)                                              |   | 25d13h  |
| tidb/23812 | [executor, planner: fix collation for hash join building (#23770)](https://github.com/pingcap/tidb/pull/23812)                                          |   | 24d12h  |
| tidb/23940 | [config, ddl: allow auto inc columns in generated columns and expression indexes](https://github.com/pingcap/tidb/pull/23940)                           |   | 15d18h  |
| tidb/23968 | [statistics: fix unstable TestDropPartitionStats test](https://github.com/pingcap/tidb/pull/23968)                                                      |   | 13d14h  |
| tidb/23987 | [executor: Implements json_arrayagg function](https://github.com/pingcap/tidb/pull/23987)                                                               |   | 12d18h  |
| tidb/24016 | [planner: fix index-out-of-range error when checking only_full_group_by (#23844)](https://github.com/pingcap/tidb/pull/24016)                           |   | 11d18h  |
| tidb/24018 | [ranger: fix the range construction behavior when the column's type is `YEAR` (#23559)](https://github.com/pingcap/tidb/pull/24018)                     |   | 11d18h  |
| tidb/24054 | [executor: fix wrong convert from bit to string when do projection (#23960)](https://github.com/pingcap/tidb/pull/24054)                                |   | 10d16h  |
| tidb/24155 | [planner, executor: fix index merge partial table scan schema (#23936)](https://github.com/pingcap/tidb/pull/24155)                                     |   | 5d20h   |
| tidb/24185 | [executor: make column default value being aware of NO_ZERO_IN_DATE (#24174)](https://github.com/pingcap/tidb/pull/24185)                               |   | 4d19h   |
| tidb/24186 | [executor: make column default value being aware of NO_ZERO_IN_DATE (#24174)](https://github.com/pingcap/tidb/pull/24186)                               |   | 4d19h   |
| tidb/24188 | [util: fix bad number error with DISTINCT when dividing long decimals (#21783)](https://github.com/pingcap/tidb/pull/24188)                             |   | 4d16h   |
| tidb/24211 | [*: support txn retry when auto id meets duplicate entry](https://github.com/pingcap/tidb/pull/24211)                                                   |   | 3d13h   |
| tidb/24231 | [executor: fix projection executor panic and add failpoint test](https://github.com/pingcap/tidb/pull/24231)                                            |   | 2d20h   |
| tidb/24234 | [executor: skip TestMppExecution when race is enabled (#24222)](https://github.com/pingcap/tidb/pull/24234)                                             |   | 2d18h   |
| tidb/24239 | [executor: make IndexMergeReader support reading partition table directly](https://github.com/pingcap/tidb/pull/24239)                                  |   | 2d16h   |
| tidb/24250 | [planner: rewritten LIKE as range for expression index](https://github.com/pingcap/tidb/pull/24250)                                                     |   | 1d21h   |
| tidb/24258 | [Revert "planner: donot prune all columns for Projection (#24024)" (#24180)](https://github.com/pingcap/tidb/pull/24258)                                |   | 22h     |
| tidb/24260 | [executor: make IndexReaders in the inner side of IndexJoins support direct reading](https://github.com/pingcap/tidb/pull/24260)                        |   | 20h     |
| tidb/24268 | [expression: fix cast real, decimal to time (#24120)](https://github.com/pingcap/tidb/pull/24268)                                                       |   | 17h     |
| tidb/24278 | [executor: accelerate TestVectorizedMergeJoin and TestVectorizedShuffleMergeJoin (#24177)](https://github.com/pingcap/tidb/pull/24278)                  |   | 10h     |


## Reviewed in Last 7 Days

|    REPO     |                                                             PR                                                              | C | D |   R   |
|-------------|-----------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/24179  | [expression: fix float64 overflow check in plus/minus real function](https://github.com/pingcap/tidb/pull/24179)            |   | 1 | 4d3h  |
| tipb/221    | [add more regexp functions](https://github.com/pingcap/tipb/pull/221)                                                       |   | 1 | 1d21h |
| tidb/24180  | [Revert "planner: donot prune all columns for Projection (#24024)"](https://github.com/pingcap/tidb/pull/24180)             |   | 3 | 2d4h  |
| tidb/24212  | [*: turn on unused linter](https://github.com/pingcap/tidb/pull/24212)                                                      |   | 3 | 13h   |
| tidb/22686  | [expression: support enum pushdown](https://github.com/pingcap/tidb/pull/22686)                                             |   | 4 | 78d5h |
| tidb/24202  | [Revert "planner: donot prune all columns for Projection (#24024) (#24093)"](https://github.com/pingcap/tidb/pull/24202)    |   | 4 | 1h    |
| tidb/24120  | [expression: fix cast real, decimal to time](https://github.com/pingcap/tidb/pull/24120)                                    |   | 4 | 3d1h  |
| tidb/24053  | [executor: fix wrong convert from bit to string when do projection (#23960)](https://github.com/pingcap/tidb/pull/24053)    |   | 4 | 6d21h |
| tidb/24191  | [executor: accelerate TestSortInDisk](https://github.com/pingcap/tidb/pull/24191)                                           |   | 4 | 16h   |
| tikv/10048  | [copr: fix cast real, decimal to time bug on overflow](https://github.com/tikv/tikv/pull/10048)                             | Y | 4 | 2d0h  |
| parser/1165 | [Add EnumSetAsIntFlag flag to control enum behavior.](https://github.com/pingcap/parser/pull/1165)                          |   | 5 | 76d3h |
| tidb/24178  | [planner/core: point get only work on TiKV](https://github.com/pingcap/tidb/pull/24178)                                     |   | 5 | 11h   |
| tidb/24125  | [telemetry: log when sending telemetry](https://github.com/pingcap/tidb/pull/24125)                                         |   | 6 | 22h   |
| tidb/24022  | [expression: don't propagateColumnEQ joinCondition when nullSensitive (#23989)](https://github.com/pingcap/tidb/pull/24022) |   | 7 | 4d21h |
| tidb/24098  | [executor: fix null div 0 for partition key expression is incorrect ](https://github.com/pingcap/tidb/pull/24098)           |   | 7 | 2d2h  |
| tidb/23936  | [planner, executor: fix index merge partial table scan schema](https://github.com/pingcap/tidb/pull/23936)                  |   | 7 | 9d17h |


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

|     REPO     |                                                                           PR                                                                            | C | LASTED  |
|--------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19029   | [types: fix unexpected NOT_NULL flags](https://github.com/pingcap/tidb/pull/19029)                                                                      |   | 262d22h |
| docs-cn/5561 | [Add sql optimization-related docs to toc](https://github.com/pingcap/docs-cn/pull/5561)                                                                |   | 62d15h  |
| tidb/20708   | [*: separate auto_increment ID allocator from _tidb_rowid allocator](https://github.com/pingcap/tidb/pull/20708)                                        |   | 178d20h |
| tidb/20969   | [executor: Improve the performance of appending not fixed columns](https://github.com/pingcap/tidb/pull/20969)                                          |   | 166d9h  |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                          |   | 164d17h |
| tidb/21149   | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                           |   | 157d14h |
| tidb/21304   | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                                |   | 151d12h |
| tidb/21318   | [planner, expression: use the range of column types to simplify expressions](https://github.com/pingcap/tidb/pull/21318)                                |   | 150d19h |
| tidb/21401   | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                                      |   | 146d11h |
| tidb/21508   | [execution: fix dayofweek('0000-00-00') behavior](https://github.com/pingcap/tidb/pull/21508)                                                           |   | 142d10h |
| tidb/21641   | [executor: Fix pessimistic lock doesn't work on the partition table for subquery/joins](https://github.com/pingcap/tidb/pull/21641)                     |   | 136d18h |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                             |   | 128d19h |
| tidb/21887   | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                                         |   | 127d11h |
| tidb/21954   | [planner/cascades: add rule `PushSelDownApply`](https://github.com/pingcap/tidb/pull/21954)                                                             |   | 123d23h |
| tidb/22146   | [executor: forbid SFU on view](https://github.com/pingcap/tidb/pull/22146)                                                                              |   | 111d21h |
| tidb/22217   | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                                           |   | 109d17h |
| tidb/22234   | [executor, planner: ON DUPLICATE UPDATE can refer to un-project col (#14412)](https://github.com/pingcap/tidb/pull/22234)                               |   | 109d15h |
| tidb/22261   | [time: fix parse datetime won't truncate the reluctant string (#22232)](https://github.com/pingcap/tidb/pull/22261)                                     |   | 108d19h |
| tidb/22374   | [expression: separated arithmeticIntDivideSig](https://github.com/pingcap/tidb/pull/22374)                                                              |   | 103d0h  |
| tidb/22415   | [ddl: refactor bundle[2/2] [6/6]](https://github.com/pingcap/tidb/pull/22415)                                                                           |   | 99d17h  |
| tidb/22416   | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                                                       | Y | 99d11h  |
| tidb/22541   | [expression: Support builtin function SOUNDEX](https://github.com/pingcap/tidb/pull/22541)                                                              |   | 89d9h   |
| tidb/22565   | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)                                      | Y | 88d17h  |
| tidb/22814   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22814)                                     |   | 65d19h  |
| tidb/22862   | [brie: fix the problem that ddl restored by BR via SQL is not replicated to downstream](https://github.com/pingcap/tidb/pull/22862)                     |   | 62d22h  |
| tidb/22923   | [expression: correct constant propagation for collation (#22666)](https://github.com/pingcap/tidb/pull/22923)                                           |   | 60d15h  |
| tidb/22924   | [planner: fix wrong index merge selection (#22825)](https://github.com/pingcap/tidb/pull/22924)                                                         |   | 60d14h  |
| tidb/23002   | [store/*: fix err check](https://github.com/pingcap/tidb/pull/23002)                                                                                    |   | 57d0h   |
| tidb/23022   | [executor: create PipelinedWindowExec](https://github.com/pingcap/tidb/pull/23022)                                                                      |   | 55d18h  |
| tidb/23152   | [expression: fix wrong error info (#22760)](https://github.com/pingcap/tidb/pull/23152)                                                                 |   | 49d14h  |
| tidb/23196   | [types: fix the bug about the wrong query result for decimal type  (#22507)](https://github.com/pingcap/tidb/pull/23196)                                |   | 47d18h  |
| tidb/23295   | [util, types: don't let SPM be affected by charset (#23161)](https://github.com/pingcap/tidb/pull/23295)                                                |   | 44d11h  |
| tidb/23316   | [planner: Fix rebuild range for prepared plan](https://github.com/pingcap/tidb/pull/23316)                                                              |   | 41d17h  |
| tidb/23373   | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                       |   | 39d19h  |
| tidb/23397   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23397)                                                          |   | 38d17h  |
| tidb/23398   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23398)                                                          |   | 38d17h  |
| tidb/23590   | [planner, table: optimize the list partition pruner for range query](https://github.com/pingcap/tidb/pull/23590)                                        |   | 30d16h  |
| tidb/23598   | [types: fix collation for binary literal (#23591)](https://github.com/pingcap/tidb/pull/23598)                                                          |   | 30d13h  |
| tidb/23655   | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral (#23523)](https://github.com/pingcap/tidb/pull/23655)                |   | 26d22h  |
| tidb/23658   | [*: collect transaction write duration/throughput metrics for SLI/SLO (#23462)](https://github.com/pingcap/tidb/pull/23658)                             |   | 26d22h  |
| tidb/23660   | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23660) |   | 26d20h  |
| tidb/23661   | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23661) |   | 26d20h  |
| tidb/23682   | [executor: fix a panic when batch point get is used for partition table (#23652)](https://github.com/pingcap/tidb/pull/23682)                           |   | 26d16h  |
| tidb/23689   | [planner: fix the panic when we calculate the partition range (#23651)](https://github.com/pingcap/tidb/pull/23689)                                     |   | 26d16h  |
| tidb/23730   | [distsql/*: typo fix for `dispatches`](https://github.com/pingcap/tidb/pull/23730)                                                                      |   | 25d18h  |
| tidb/23796   | [tests: make TestIndexLookupMergeJoinHang and TestIssue18068 stable (#23741)](https://github.com/pingcap/tidb/pull/23796)                               |   | 24d19h  |
| tidb/23812   | [executor, planner: fix collation for hash join building (#23770)](https://github.com/pingcap/tidb/pull/23812)                                          |   | 24d12h  |
| tidb/23878   | [functions: fix some string function has wrong collation and flag (#23835)](https://github.com/pingcap/tidb/pull/23878)                                 |   | 18d21h  |
| tidb/23963   | [executor: checking chunk is full precedes filtering](https://github.com/pingcap/tidb/pull/23963)                                                       |   | 13d17h  |
| tidb/23987   | [executor: Implements json_arrayagg function](https://github.com/pingcap/tidb/pull/23987)                                                               |   | 12d18h  |
| tidb/23997   | [stats, executor: use a correct sampling to collect stats](https://github.com/pingcap/tidb/pull/23997)                                                  |   | 12d9h   |
| tidb/24018   | [ranger: fix the range construction behavior when the column's type is `YEAR` (#23559)](https://github.com/pingcap/tidb/pull/24018)                     |   | 11d18h  |
| tidb/24188   | [util: fix bad number error with DISTINCT when dividing long decimals (#21783)](https://github.com/pingcap/tidb/pull/24188)                             |   | 4d16h   |
| tidb/24193   | [executor: implement CTEStorage](https://github.com/pingcap/tidb/pull/24193)                                                                            |   | 4d10h   |
| tidb/24214   | [plan: merge continuous selections and delete surely true expressions](https://github.com/pingcap/tidb/pull/24214)                                      |   | 3d12h   |
| tidb/24229   | [executor: speed up race test TestInsertReorgDelete (#24208)](https://github.com/pingcap/tidb/pull/24229)                                               |   | 2d21h   |
| tidb/24235   | [expression: try to fix TestExprPushDownToFlash tests](https://github.com/pingcap/tidb/pull/24235)                                                      |   | 2d17h   |
| tidb/24241   | [planner/core: remove random test to reduce CI time (#24207)](https://github.com/pingcap/tidb/pull/24241)                                               |   | 2d15h   |
| tidb/24245   | [planner: between .. and on int column can be used to prune hash partition](https://github.com/pingcap/tidb/pull/24245)                                 |   | 2d11h   |
| tidb/24248   | [executor, variable: small cleanup](https://github.com/pingcap/tidb/pull/24248)                                                                         |   | 2d7h    |
| tidb/24266   | [expression: fix wrong flen infer for bit constant (#23867)](https://github.com/pingcap/tidb/pull/24266)                                                |   | 18h     |
| tidb/24267   | [expression: fix wrong flen infer for bit constant (#23867)](https://github.com/pingcap/tidb/pull/24267)                                                |   | 17h     |
| tidb/24279   | [*: add security enhanced mode part 2](https://github.com/pingcap/tidb/pull/24279)                                                                      |   | 3h      |


## Reviewed in Last 7 Days

|     REPO      |                                                                 PR                                                                  | C | D | R  |
|---------------|-------------------------------------------------------------------------------------------------------------------------------------|---|---|----|
| community/439 | [planner: promote member and update links](https://github.com/pingcap/community/pull/439)                                           |   | 5 | 0h |
| tidb/24183    | [executor, statistics: remove NULL value from column histogram created by fast analyze](https://github.com/pingcap/tidb/pull/24183) |   | 5 | 0h |


</details> 


<details> 
  <summary>rebelice's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                           PR                                                           | C | LASTED |
|------------|------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/23537 | [planner: remove some risky cache operations in the plan builder (#23354)](https://github.com/pingcap/tidb/pull/23537) |   | 31d19h |
| tidb/23836 | [parser, core: Implement force_index hint in parser and TiDB](https://github.com/pingcap/tidb/pull/23836)              |   | 23d17h |
| tidb/24033 | [statistics: fix some unstable tests in global stats (#23502)](https://github.com/pingcap/tidb/pull/24033)             |   | 11d9h  |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>time-and-fate's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                         PR                                                          | C | LASTED  |
|------------|---------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                           |   | 171d17h |
| tidb/22416 | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                   | Y | 99d11h  |
| tidb/24155 | [planner, executor: fix index merge partial table scan schema (#23936)](https://github.com/pingcap/tidb/pull/24155) |   | 5d20h   |
| tidb/24204 | [planner: clone possible properties before saving them](https://github.com/pingcap/tidb/pull/24204)                 |   | 3d17h   |
| tidb/24230 | [*: consitent get infoschema](https://github.com/pingcap/tidb/pull/24230)                                           |   | 2d21h   |


## Reviewed in Last 7 Days

|    REPO    |                                                                 PR                                                                  | C | D |   R    |
|------------|-------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/23997 | [stats, executor: use a correct sampling to collect stats](https://github.com/pingcap/tidb/pull/23997)                              |   | 1 | 11d20h |
| tidb/24204 | [planner: clone possible properties before saving them](https://github.com/pingcap/tidb/pull/24204)                                 |   | 3 | 1d1h   |
| tidb/24089 | [statistics: introduce the weighted reservoir sampling](https://github.com/pingcap/tidb/pull/24089)                                 |   | 4 | 5d23h  |
| tipb/220   | [analyze: add proto fields for row based sampling](https://github.com/pingcap/tipb/pull/220)                                        |   | 4 | 6d22h  |
| tidb/24175 | [statistics: skip reading mysql.stats_histograms if cached stats is up-to-date](https://github.com/pingcap/tidb/pull/24175)         |   | 5 | 1d0h   |
| tidb/24183 | [executor, statistics: remove NULL value from column histogram created by fast analyze](https://github.com/pingcap/tidb/pull/24183) |   | 5 | 1h     |


</details> 


<details> 
  <summary>winoros's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                              PR                                                                               | C | LASTED  |
|--------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| docs-cn/5916 | [sql-statements, information-schema: add `END_TIME` field for table `ANALYZE_STATUS`](https://github.com/pingcap/docs-cn/pull/5916)                           |   | 24d17h  |
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                                        | Y | 226d14h |
| docs-cn/6113 | [config: update the default value of `feedback-probability`](https://github.com/pingcap/docs-cn/pull/6113)                                                    |   | 3d22h   |
| tidb/20765   | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                                             |   | 174d17h |
| tidb/20877   | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                                                     |   | 171d17h |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                                |   | 164d17h |
| tidb/21207   | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                                               |   | 153d19h |
| tidb/21487   | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                                                 |   | 143d4h  |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                                   |   | 128d19h |
| tidb/21954   | [planner/cascades: add rule `PushSelDownApply`](https://github.com/pingcap/tidb/pull/21954)                                                                   |   | 123d23h |
| tidb/22181   | [planner, expression: fix error when using IN combined with subquery (#22080)](https://github.com/pingcap/tidb/pull/22181)                                    |   | 110d17h |
| tidb/22416   | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                                                             | Y | 99d11h  |
| tidb/22504   | [*:Fix the fetchHotRegion bug that the count always zero](https://github.com/pingcap/tidb/pull/22504)                                                         |   | 91d19h  |
| tidb/22565   | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)                                            | Y | 88d17h  |
| tidb/22624   | [ planner: not pruning column used by union scan condition (#21640)](https://github.com/pingcap/tidb/pull/22624)                                              |   | 86d17h  |
| tidb/22923   | [expression: correct constant propagation for collation (#22666)](https://github.com/pingcap/tidb/pull/22923)                                                 |   | 60d15h  |
| tidb/23233   | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23233)                                                        |   | 46d18h  |
| tidb/23347   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                                          |   | 40d18h  |
| tidb/23348   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                                          |   | 40d18h  |
| tidb/23350   | [util/stringutil, util/ranger, planner: use hierarchical separators to simplify the parsing for info of EXPLAIN ](https://github.com/pingcap/tidb/pull/23350) |   | 40d17h  |
| tidb/23373   | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                             |   | 39d19h  |
| tidb/23474   | [planner: fix inappropriate null flag of null constants (#23457)](https://github.com/pingcap/tidb/pull/23474)                                                 |   | 33d18h  |
| tidb/23537   | [planner: remove some risky cache operations in the plan builder (#23354)](https://github.com/pingcap/tidb/pull/23537)                                        |   | 31d19h  |
| tidb/23543   | [statistics: fix auto analyze log information incomplete (#23522)](https://github.com/pingcap/tidb/pull/23543)                                                |   | 31d18h  |
| tidb/23598   | [types: fix collation for binary literal (#23591)](https://github.com/pingcap/tidb/pull/23598)                                                                |   | 30d13h  |
| tidb/23655   | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral (#23523)](https://github.com/pingcap/tidb/pull/23655)                      |   | 26d22h  |
| tidb/23689   | [planner: fix the panic when we calculate the partition range (#23651)](https://github.com/pingcap/tidb/pull/23689)                                           |   | 26d16h  |
| tidb/23772   | [tablecodec: fix text type decode for old row format (#23751)](https://github.com/pingcap/tidb/pull/23772)                                                    |   | 25d11h  |
| tidb/23849   | [ddl: tidb panic while query hash partition table with is null condition](https://github.com/pingcap/tidb/pull/23849)                                         |   | 20d13h  |
| tidb/23917   | [planner: fix wrong TableDual plans caused by comparing Binary and Bytes incorrectly (#23860)](https://github.com/pingcap/tidb/pull/23917)                    |   | 17d23h  |
| tidb/23946   | [planner: fix visit info for grant/revoke](https://github.com/pingcap/tidb/pull/23946)                                                                        |   | 14d6h   |
| tidb/23970   | [planner: fix a bug that point get plan returns wrong column name (#23365)](https://github.com/pingcap/tidb/pull/23970)                                       |   | 13d14h  |
| tidb/24018   | [ranger: fix the range construction behavior when the column's type is `YEAR` (#23559)](https://github.com/pingcap/tidb/pull/24018)                           |   | 11d18h  |
| tidb/24060   | [statistics: fix some potential panic in statistics (#23988)](https://github.com/pingcap/tidb/pull/24060)                                                     |   | 10d13h  |
| tidb/24061   | [statistics: fix some potential panic in statistics (#23988)](https://github.com/pingcap/tidb/pull/24061)                                                     |   | 10d13h  |
| tidb/24079   | [planner: change descScanFactor to scanFactor when ExpectedCount is small. (#23972)](https://github.com/pingcap/tidb/pull/24079)                              |   | 9d19h   |
| tidb/24097   | [planner: Remove redundant call to expression.ColumnSubstitute](https://github.com/pingcap/tidb/pull/24097)                                                   |   | 9d3h    |
| tidb/24138   | [planner: Add Equivalence Rules to Transform BinaryOptSubquery to ExistsSubquery](https://github.com/pingcap/tidb/pull/24138)                                 |   | 6d12h   |
| tidb/24204   | [planner: clone possible properties before saving them](https://github.com/pingcap/tidb/pull/24204)                                                           |   | 3d17h   |
| tidb/24241   | [planner/core: remove random test to reduce CI time (#24207)](https://github.com/pingcap/tidb/pull/24241)                                                     |   | 2d15h   |
| tidb/24258   | [Revert "planner: donot prune all columns for Projection (#24024)" (#24180)](https://github.com/pingcap/tidb/pull/24258)                                      |   | 22h     |


## Reviewed in Last 7 Days

|     REPO     |                                                             PR                                                              | C | D |  R   |
|--------------|-----------------------------------------------------------------------------------------------------------------------------|---|---|------|
| docs-cn/6061 | [releases: add tidb 5.0.1 release notes](https://github.com/pingcap/docs-cn/pull/6061)                                      |   | 4 | 3d2h |
| tidb/24202   | [Revert "planner: donot prune all columns for Projection (#24024) (#24093)"](https://github.com/pingcap/tidb/pull/24202)    |   | 4 | 0h   |
| tidb/24175   | [statistics: skip reading mysql.stats_histograms if cached stats is up-to-date](https://github.com/pingcap/tidb/pull/24175) |   | 6 | 2h   |
| docs/5392    | [releases: add tidb 5.0.1 release notes](https://github.com/pingcap/docs/pull/5392)                                         |   | 6 | 21h  |


</details> 


<details> 
  <summary>wshwsh12's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                   PR                                                                   | C | LASTED  |
|------------|----------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19807 | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                                |   | 233d11h |
| tidb/19957 | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                 | Y | 226d14h |
| tidb/21487 | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                          |   | 143d4h  |
| tidb/21887 | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                        |   | 127d11h |
| tidb/22378 | [executor: vectorize hash aggregate](https://github.com/pingcap/tidb/pull/22378)                                                       |   | 102d19h |
| tidb/22686 | [expression: support enum pushdown](https://github.com/pingcap/tidb/pull/22686)                                                        |   | 81d22h  |
| tidb/23336 | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23336)               |   | 40d19h  |
| tidb/23347 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                   |   | 40d18h  |
| tidb/23348 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                   |   | 40d18h  |
| tidb/23368 | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23368)                  |   | 39d20h  |
| tidb/23397 | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23397)                                         |   | 38d17h  |
| tidb/23398 | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23398)                                         |   | 38d17h  |
| tidb/23519 | [executor: check privilege before adding](https://github.com/pingcap/tidb/pull/23519)                                                  |   | 32d0h   |
| tidb/23760 | [collation: fix tidb panic when compare string with collation](https://github.com/pingcap/tidb/pull/23760)                             |   | 25d13h  |
| tidb/23866 | [executor,kv: support timebounded staleness transaction](https://github.com/pingcap/tidb/pull/23866)                                   |   | 19d16h  |
| tidb/23968 | [statistics: fix unstable TestDropPartitionStats test](https://github.com/pingcap/tidb/pull/23968)                                     |   | 13d14h  |
| tidb/23979 | [executor, statistics: fix unstable `TestAnalyzeIndexExtractTopN`](https://github.com/pingcap/tidb/pull/23979)                         |   | 12d23h  |
| tidb/24018 | [ranger: fix the range construction behavior when the column's type is `YEAR` (#23559)](https://github.com/pingcap/tidb/pull/24018)    |   | 11d18h  |
| tidb/24033 | [statistics: fix some unstable tests in global stats (#23502)](https://github.com/pingcap/tidb/pull/24033)                             |   | 11d9h   |
| tidb/24050 | [expression: fix get var panic when types not match](https://github.com/pingcap/tidb/pull/24050)                                       |   | 10d17h  |
| tidb/24053 | [executor: fix wrong convert from bit to string when do projection (#23960)](https://github.com/pingcap/tidb/pull/24053)               |   | 10d16h  |
| tidb/24054 | [executor: fix wrong convert from bit to string when do projection (#23960)](https://github.com/pingcap/tidb/pull/24054)               |   | 10d16h  |
| tidb/24147 | [docs/design: add proposal for common table expression](https://github.com/pingcap/tidb/pull/24147)                                    |   | 5d23h   |
| tidb/24186 | [executor: make column default value being aware of NO_ZERO_IN_DATE (#24174)](https://github.com/pingcap/tidb/pull/24186)              |   | 4d19h   |
| tidb/24228 | [executor: skip TestPrepareStmtAfterIsolationReadChange when race enable (#24200)](https://github.com/pingcap/tidb/pull/24228)         |   | 2d22h   |
| tidb/24229 | [executor: speed up race test TestInsertReorgDelete (#24208)](https://github.com/pingcap/tidb/pull/24229)                              |   | 2d21h   |
| tidb/24230 | [*: consitent get infoschema](https://github.com/pingcap/tidb/pull/24230)                                                              |   | 2d21h   |
| tidb/24234 | [executor: skip TestMppExecution when race is enabled (#24222)](https://github.com/pingcap/tidb/pull/24234)                            |   | 2d18h   |
| tidb/24236 | [*: remove SchemaVersion in TransactionContext](https://github.com/pingcap/tidb/pull/24236)                                            |   | 2d17h   |
| tidb/24257 | [executor: fix data race of parallel apply operator](https://github.com/pingcap/tidb/pull/24257)                                       |   | 23h     |
| tidb/24266 | [expression: fix wrong flen infer for bit constant (#23867)](https://github.com/pingcap/tidb/pull/24266)                               |   | 18h     |
| tidb/24267 | [expression: fix wrong flen infer for bit constant (#23867)](https://github.com/pingcap/tidb/pull/24267)                               |   | 17h     |
| tidb/24268 | [expression: fix cast real, decimal to time (#24120)](https://github.com/pingcap/tidb/pull/24268)                                      |   | 17h     |
| tidb/24278 | [executor: accelerate TestVectorizedMergeJoin and TestVectorizedShuffleMergeJoin (#24177)](https://github.com/pingcap/tidb/pull/24278) |   | 10h     |
| tidb/24280 | [executor, session, variable: Move deprecation and synonyms to sysvar struct](https://github.com/pingcap/tidb/pull/24280)              |   | 0h      |


## Reviewed in Last 7 Days

|    REPO    |                                                                          PR                                                                          | C | D |   R    |
|------------|------------------------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/23867 | [expression: fix wrong flen infer for bit constant](https://github.com/pingcap/tidb/pull/23867)                                                      |   | 3 | 16d20h |
| tidb/24222 | [executor: skip TestMppExecution when race is enabled](https://github.com/pingcap/tidb/pull/24222)                                                   |   | 3 | 0h     |
| tidb/24208 | [executor: speed up race test TestInsertReorgDelete](https://github.com/pingcap/tidb/pull/24208)                                                     |   | 3 | 14h    |
| tidb/24200 | [executor: skip TestPrepareStmtAfterIsolationReadChange when race enable](https://github.com/pingcap/tidb/pull/24200)                                |   | 3 | 19h    |
| tidb/23876 | [executor: fix scope ambiguity of joinResult](https://github.com/pingcap/tidb/pull/23876)                                                            |   | 6 | 14d1h  |
| tidb/24139 | [executor: accelerate TestShowVar (#24131)](https://github.com/pingcap/tidb/pull/24139)                                                              |   | 6 | 1d0h   |
| tidb/24132 | [executor: accelerate TestUpdateScanningHandles and TestIssue20658 and TestParallelStreamAggGroupConcat](https://github.com/pingcap/tidb/pull/24132) |   | 7 | 14h    |


</details> 

