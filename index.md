|   REVIEWER    |  O(C)   | 1 | 2 | 3 | 4 | 5 | 6 | 7 | T  |
|---------------|---------|---|---|---|---|---|---|---|----|
| Reminiscent   |  6 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| SunRunAway    | 13 ( 1) | 0 | 0 | 0 | 0 | 1 | 0 | 0 |  1 |
| XuHuaiyu      | 83 ( 2) | 1 | 2 | 1 | 0 | 2 | 0 | 0 |  6 |
| dyzsr         |  1 ( 0) | 0 | 0 | 0 | 0 | 1 | 0 | 0 |  1 |
| eurekaka      | 21 ( 0) | 0 | 0 | 1 | 0 | 0 | 1 | 4 |  6 |
| fzhedu        |  0 ( 0) | 0 | 0 | 1 | 0 | 1 | 0 | 0 |  2 |
| ichn-hu       |  0 ( 0) | 5 | 1 | 2 | 0 | 2 | 0 | 0 | 10 |
| lzmhhh123     | 53 ( 0) | 0 | 1 | 2 | 0 | 2 | 6 | 2 | 13 |
| nullnotnil    |  0 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| qw4990        | 64 ( 2) | 2 | 2 | 0 | 0 | 0 | 0 | 2 |  6 |
| rebelice      |  4 ( 0) | 0 | 1 | 0 | 0 | 0 | 0 | 0 |  1 |
| time-and-fate |  7 ( 1) | 1 | 0 | 0 | 0 | 1 | 2 | 2 |  6 |
| winoros       | 40 ( 3) | 0 | 0 | 0 | 0 | 0 | 2 | 0 |  2 |
| wshwsh12      | 32 ( 1) | 0 | 3 | 0 | 0 | 4 | 0 | 0 |  7 |


<details> 
  <summary>Reminiscent's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                     PR                                                                     | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/21896 | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                                 |   | 127d19h |
| tidb/23474 | [planner: fix inappropriate null flag of null constants (#23457)](https://github.com/pingcap/tidb/pull/23474)                              |   | 35d18h  |
| tidb/23575 | [executor: fix update panic on join having statement (#23554)](https://github.com/pingcap/tidb/pull/23575)                                 |   | 32d21h  |
| tidb/23917 | [planner: fix wrong TableDual plans caused by comparing Binary and Bytes incorrectly (#23860)](https://github.com/pingcap/tidb/pull/23917) |   | 19d23h  |
| tidb/24016 | [planner: fix index-out-of-range error when checking only_full_group_by (#23844)](https://github.com/pingcap/tidb/pull/24016)              |   | 13d18h  |
| tidb/24188 | [util: fix bad number error with DISTINCT when dividing long decimals (#21783)](https://github.com/pingcap/tidb/pull/24188)                |   | 6d16h   |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>SunRunAway's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                  PR                                                                   | C | LASTED  |
|------------|---------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19178 | [executor: Refactor probe channel](https://github.com/pingcap/tidb/pull/19178)                                                        |   | 257d16h |
| tidb/19807 | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                               |   | 235d11h |
| tidb/19900 | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                        | Y | 230d18h |
| tidb/20140 | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                    |   | 217d22h |
| tidb/20765 | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                     |   | 176d17h |
| tidb/21207 | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                       |   | 155d19h |
| tidb/21834 | [planner: enhanced index range calculation plan](https://github.com/pingcap/tidb/pull/21834)                                          |   | 132d18h |
| tidb/21876 | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)           |   | 130d19h |
| tidb/21878 | [planner: do not push down lock to pointGet/bacthPointGet when selection exists](https://github.com/pingcap/tidb/pull/21878)          |   | 130d18h |
| tidb/21956 | [planner/preprocessor: disallow into-outfile clause in some place](https://github.com/pingcap/tidb/pull/21956)                        |   | 125d23h |
| tidb/22217 | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                         |   | 111d17h |
| tidb/22379 | [[experiment] executor: allow aggregation to spill disk when running out of memory quota](https://github.com/pingcap/tidb/pull/22379) |   | 104d19h |
| tidb/24279 | [*: add security enhanced mode part 2](https://github.com/pingcap/tidb/pull/24279)                                                    |   | 2d3h    |


## Reviewed in Last 7 Days

|    REPO    |                                          PR                                           | C | D |   R    |
|------------|---------------------------------------------------------------------------------------|---|---|--------|
| tidb/20749 | [executor: support global kill (32 bits)](https://github.com/pingcap/tidb/pull/20749) |   | 5 | 172d5h |


</details> 


<details> 
  <summary>XuHuaiyu's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                              PR                                                                               | C | LASTED  |
|--------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| docs-cn/5619 | [Update data-type-date-and-time.md](https://github.com/pingcap/docs-cn/pull/5619)                                                                             |   | 60d16h  |
| tidb/19900   | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                                                | Y | 230d18h |
| docs-cn/5671 | [tidb: Add time format description](https://github.com/pingcap/docs-cn/pull/5671)                                                                             |   | 54d11h  |
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                                        | Y | 228d14h |
| tidb/20140   | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                                            |   | 217d22h |
| tidb/20749   | [executor: support global kill (32 bits)](https://github.com/pingcap/tidb/pull/20749)                                                                         |   | 177d2h  |
| tidb/20790   | [collation: add pinyin collation for chinese charset support](https://github.com/pingcap/tidb/pull/20790)                                                     |   | 175d20h |
| tidb/20969   | [executor: Improve the performance of appending not fixed columns](https://github.com/pingcap/tidb/pull/20969)                                                |   | 168d9h  |
| tidb/21064   | [planner, executor: fix cast not check error](https://github.com/pingcap/tidb/pull/21064)                                                                     |   | 163d8h  |
| tidb/21149   | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                                 |   | 159d14h |
| tidb/21228   | [executor: return the result immediately when combining LIMIT row_count with DISTINCT](https://github.com/pingcap/tidb/pull/21228)                            |   | 155d13h |
| tidb/21304   | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                                      |   | 153d12h |
| tidb/21334   | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                                                 |   | 152d14h |
| tidb/21401   | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                                            |   | 148d11h |
| tidb/21536   | [executor: add slow-log file meta cache to avoid repeat read file meta information](https://github.com/pingcap/tidb/pull/21536)                               |   | 141d14h |
| tidb/21564   | [ddl: fix Incorrect behavior of NO_ZERO_DATE when altering table](https://github.com/pingcap/tidb/pull/21564)                                                 |   | 140d15h |
| tidb/21896   | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                                                    |   | 127d19h |
| tidb/22131   | [privilege: remove leading and trailing space when create user and role](https://github.com/pingcap/tidb/pull/22131)                                          |   | 117d19h |
| tidb/22163   | [expression: separated arithmeticMinusIntSig](https://github.com/pingcap/tidb/pull/22163)                                                                     |   | 113d13h |
| tidb/22186   | [executor: fix select into outfile with year type column has no data (#22175)](https://github.com/pingcap/tidb/pull/22186)                                    |   | 112d16h |
| tidb/22616   | [expression: from_unixtime accept 64-bit integers](https://github.com/pingcap/tidb/pull/22616)                                                                |   | 88d23h  |
| tidb/22617   | [metrics: fix wrong bucket name of coprocessor cache (#22454)](https://github.com/pingcap/tidb/pull/22617)                                                    |   | 88d23h  |
| tidb/22624   | [ planner: not pruning column used by union scan condition (#21640)](https://github.com/pingcap/tidb/pull/22624)                                              |   | 88d17h  |
| tidb/22631   | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                                                                               |   | 86d23h  |
| tidb/22696   | [expression: enable arithmetic Mod push down](https://github.com/pingcap/tidb/pull/22696)                                                                     |   | 83d17h  |
| tidb/22711   | [executor: Fix inline schema name](https://github.com/pingcap/tidb/pull/22711)                                                                                |   | 83d11h  |
| tidb/22722   | [planner, errno: make error code of ErrMixOfGroupFuncAndFields consistent with MySQL](https://github.com/pingcap/tidb/pull/22722)                             |   | 82d20h  |
| tidb/22814   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22814)                                           |   | 67d19h  |
| tidb/22908   | [txn: Add txn state's view](https://github.com/pingcap/tidb/pull/22908)                                                                                       |   | 62d20h  |
| tidb/23012   | [executor: fix affected rows of ddls and complete uint tests](https://github.com/pingcap/tidb/pull/23012)                                                     |   | 58d16h  |
| tidb/23152   | [expression: fix wrong error info (#22760)](https://github.com/pingcap/tidb/pull/23152)                                                                       |   | 51d14h  |
| tidb/23196   | [types: fix the bug about the wrong query result for decimal type  (#22507)](https://github.com/pingcap/tidb/pull/23196)                                      |   | 49d18h  |
| tidb/23220   | [Release 4.0](https://github.com/pingcap/tidb/pull/23220)                                                                                                     |   | 49d11h  |
| tidb/23233   | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23233)                                                        |   | 48d18h  |
| tidb/23257   | [executor: group_concat aggr panic when session.group_concat_max_len is small (#23131)](https://github.com/pingcap/tidb/pull/23257)                           |   | 47d18h  |
| tidb/23295   | [util, types: don't let SPM be affected by charset (#23161)](https://github.com/pingcap/tidb/pull/23295)                                                      |   | 46d11h  |
| tidb/23335   | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23335)                                      |   | 42d19h  |
| tidb/23336   | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23336)                                      |   | 42d19h  |
| tidb/23347   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                                          |   | 42d18h  |
| tidb/23348   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                                          |   | 42d18h  |
| tidb/23350   | [util/stringutil, util/ranger, planner: use hierarchical separators to simplify the parsing for info of EXPLAIN ](https://github.com/pingcap/tidb/pull/23350) |   | 42d17h  |
| tidb/23368   | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23368)                                         |   | 41d20h  |
| tidb/23397   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23397)                                                                |   | 40d17h  |
| tidb/23398   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23398)                                                                |   | 40d17h  |
| tidb/23405   | [domain: remove the exit chan, use context](https://github.com/pingcap/tidb/pull/23405)                                                                       |   | 40d17h  |
| tidb/23433   | [WIP: speed up for slow query logs retrieving ](https://github.com/pingcap/tidb/pull/23433)                                                                   |   | 39d17h  |
| tidb/23474   | [planner: fix inappropriate null flag of null constants (#23457)](https://github.com/pingcap/tidb/pull/23474)                                                 |   | 35d18h  |
| tidb/23497   | [expression: Let TiDB use Hyperscan to support multi-pattern-match](https://github.com/pingcap/tidb/pull/23497)                                               |   | 34d22h  |
| tidb/23517   | [*: Add the metric about the SQL with TiFlash Success  (#23426)](https://github.com/pingcap/tidb/pull/23517)                                                  |   | 34d12h  |
| tidb/23562   | [execution: reuse iterator in hash join](https://github.com/pingcap/tidb/pull/23562)                                                                          |   | 33d13h  |
| tidb/23640   | [*: fix the bug about YEAR(0.9) returns NULL instead of 0 in NO_ZERO_DATE mode](https://github.com/pingcap/tidb/pull/23640)                                   |   | 29d13h  |
| tidb/23661   | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23661)       |   | 28d20h  |
| tidb/23682   | [executor: fix a panic when batch point get is used for partition table (#23652)](https://github.com/pingcap/tidb/pull/23682)                                 |   | 28d16h  |
| tidb/23691   | [executor: fix index join on prefix column index (#23678)](https://github.com/pingcap/tidb/pull/23691)                                                        |   | 28d15h  |
| tidb/23705   | [executor: refineArgs() bug fix when compare int with very small decimal (#23694)](https://github.com/pingcap/tidb/pull/23705)                                |   | 28d13h  |
| tidb/23756   | [planner: fix set not null flag for outer join (#23727)](https://github.com/pingcap/tidb/pull/23756)                                                          |   | 27d14h  |
| tidb/23812   | [executor, planner: fix collation for hash join building (#23770)](https://github.com/pingcap/tidb/pull/23812)                                                |   | 26d12h  |
| tidb/23878   | [functions: fix some string function has wrong collation and flag (#23835)](https://github.com/pingcap/tidb/pull/23878)                                       |   | 20d21h  |
| tidb/23884   | [Metric: Collect TiKV Read Metric for SLI/SLO](https://github.com/pingcap/tidb/pull/23884)                                                                    |   | 20d19h  |
| tidb/23888   | [executor: fix resource leak of Shuffle Executor.](https://github.com/pingcap/tidb/pull/23888)                                                                |   | 20d18h  |
| tidb/23958   | [executor: fix `show table status` for the database with upper-cased name (#23896)](https://github.com/pingcap/tidb/pull/23958)                               |   | 15d18h  |
| tidb/23964   | [executor: GROUP_CONCAT(float) is not compatible with mysql](https://github.com/pingcap/tidb/pull/23964)                                                      |   | 15d16h  |
| tidb/24007   | [ddl: refactor rule [4/6]](https://github.com/pingcap/tidb/pull/24007)                                                                                        |   | 13d20h  |
| tidb/24016   | [planner: fix index-out-of-range error when checking only_full_group_by (#23844)](https://github.com/pingcap/tidb/pull/24016)                                 |   | 13d18h  |
| tidb/24026   | [types: fix type merge about bit type (#23857)](https://github.com/pingcap/tidb/pull/24026)                                                                   |   | 13d14h  |
| tidb/24033   | [statistics: fix some unstable tests in global stats (#23502)](https://github.com/pingcap/tidb/pull/24033)                                                    |   | 13d9h   |
| tidb/24053   | [executor: fix wrong convert from bit to string when do projection (#23960)](https://github.com/pingcap/tidb/pull/24053)                                      |   | 12d16h  |
| tidb/24060   | [statistics: fix some potential panic in statistics (#23988)](https://github.com/pingcap/tidb/pull/24060)                                                     |   | 12d13h  |
| tidb/24061   | [statistics: fix some potential panic in statistics (#23988)](https://github.com/pingcap/tidb/pull/24061)                                                     |   | 12d13h  |
| tidb/24078   | [planner: change descScanFactor to scanFactor when ExpectedCount is small. (#23972)](https://github.com/pingcap/tidb/pull/24078)                              |   | 11d19h  |
| tidb/24079   | [planner: change descScanFactor to scanFactor when ExpectedCount is small. (#23972)](https://github.com/pingcap/tidb/pull/24079)                              |   | 11d19h  |
| tidb/24136   | [store/tikv:use tikv.error.ErrNotExist instead of kv.ErrNotExist](https://github.com/pingcap/tidb/pull/24136)                                                 |   | 8d13h   |
| tidb/24155   | [planner, executor: fix index merge partial table scan schema (#23936)](https://github.com/pingcap/tidb/pull/24155)                                           |   | 7d20h   |
| tidb/24157   | [planner: let CopTiFlashConcurrencyFactor inflence the cost of whole plan](https://github.com/pingcap/tidb/pull/24157)                                        |   | 7d19h   |
| tidb/24196   | [*: support read and write operations for the global temporary table](https://github.com/pingcap/tidb/pull/24196)                                             |   | 5d23h   |
| tidb/24228   | [executor: skip TestPrepareStmtAfterIsolationReadChange when race enable (#24200)](https://github.com/pingcap/tidb/pull/24228)                                |   | 4d22h   |
| tidb/24229   | [executor: speed up race test TestInsertReorgDelete (#24208)](https://github.com/pingcap/tidb/pull/24229)                                                     |   | 4d21h   |
| tidb/24234   | [executor: skip TestMppExecution when race is enabled (#24222)](https://github.com/pingcap/tidb/pull/24234)                                                   |   | 4d18h   |
| tidb/24241   | [planner/core: remove random test to reduce CI time (#24207)](https://github.com/pingcap/tidb/pull/24241)                                                     |   | 4d15h   |
| tidb/24266   | [expression: fix wrong flen infer for bit constant (#23867)](https://github.com/pingcap/tidb/pull/24266)                                                      |   | 2d17h   |
| tidb/24267   | [expression: fix wrong flen infer for bit constant (#23867)](https://github.com/pingcap/tidb/pull/24267)                                                      |   | 2d17h   |
| tidb/24278   | [executor: accelerate TestVectorizedMergeJoin and TestVectorizedShuffleMergeJoin (#24177)](https://github.com/pingcap/tidb/pull/24278)                        |   | 2d10h   |
| tidb/24287   | [planner/core: support union all for mpp.](https://github.com/pingcap/tidb/pull/24287)                                                                        |   | 1d19h   |


## Reviewed in Last 7 Days

|    REPO    |                                                              PR                                                               | C | D |   R   |
|------------|-------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/24308 | [docs/design: update the proposal of `column type change`](https://github.com/pingcap/tidb/pull/24308)                        |   | 1 | 0h    |
| tidb/24179 | [expression: fix float64 overflow check in plus/minus real function](https://github.com/pingcap/tidb/pull/24179)              |   | 2 | 5d22h |
| tidb/24290 | [expression: fix wrong type infer for agg function when type is null](https://github.com/pingcap/tidb/pull/24290)             |   | 2 | 18h   |
| tidb/24257 | [executor: fix data race of parallel apply operator](https://github.com/pingcap/tidb/pull/24257)                              |   | 3 | 0h    |
| docs/5445  | [tidb-config: add value range and type for token-limit](https://github.com/pingcap/docs/pull/5445)                            |   | 5 | 0h    |
| tidb/24177 | [executor: accelerate TestVectorizedMergeJoin and TestVectorizedShuffleMergeJoin](https://github.com/pingcap/tidb/pull/24177) |   | 5 | 2d14h |


</details> 


<details> 
  <summary>dyzsr's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                 PR                                                                  | C | LASTED |
|------------|-------------------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/24018 | [ranger: fix the range construction behavior when the column's type is `YEAR` (#23559)](https://github.com/pingcap/tidb/pull/24018) |   | 13d18h |


## Reviewed in Last 7 Days

|    REPO    |                                      PR                                       | C | D | R  |
|------------|-------------------------------------------------------------------------------|---|---|----|
| tikv/10074 | [copr: fix unsound unsafe transmute](https://github.com/tikv/tikv/pull/10074) |   | 5 | 2h |


</details> 


<details> 
  <summary>eurekaka's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                         PR                                                                         | C | LASTED  |
|------------|----------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                                          |   | 173d17h |
| tidb/23002 | [store/*: fix err check](https://github.com/pingcap/tidb/pull/23002)                                                                               |   | 59d0h   |
| tidb/23283 | [util: optimize the performance of restore with db (#22910)](https://github.com/pingcap/tidb/pull/23283)                                           |   | 46d17h  |
| tidb/23316 | [planner: Fix rebuild range for prepared plan](https://github.com/pingcap/tidb/pull/23316)                                                         |   | 43d17h  |
| tidb/23373 | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                  |   | 41d19h  |
| tidb/23543 | [statistics: fix auto analyze log information incomplete (#23522)](https://github.com/pingcap/tidb/pull/23543)                                     |   | 33d18h  |
| tidb/23689 | [planner: fix the panic when we calculate the partition range (#23651)](https://github.com/pingcap/tidb/pull/23689)                                |   | 28d16h  |
| tidb/23705 | [executor: refineArgs() bug fix when compare int with very small decimal (#23694)](https://github.com/pingcap/tidb/pull/23705)                     |   | 28d13h  |
| tidb/23756 | [planner: fix set not null flag for outer join (#23727)](https://github.com/pingcap/tidb/pull/23756)                                               |   | 27d14h  |
| tidb/23760 | [collation: fix tidb panic when compare string with collation](https://github.com/pingcap/tidb/pull/23760)                                         |   | 27d13h  |
| tidb/23938 | [planner,privilege: requires extra privileges for REPLACE and INSERT ON DUPLICATE statements (#23911)](https://github.com/pingcap/tidb/pull/23938) |   | 18d10h  |
| tidb/23974 | [planner: do not push down to TiFlash if the table scan require to scan data in desc order (#23948)](https://github.com/pingcap/tidb/pull/23974)   |   | 15d12h  |
| tidb/24033 | [statistics: fix some unstable tests in global stats (#23502)](https://github.com/pingcap/tidb/pull/24033)                                         |   | 13d9h   |
| tidb/24061 | [statistics: fix some potential panic in statistics (#23988)](https://github.com/pingcap/tidb/pull/24061)                                          |   | 12d13h  |
| tidb/24079 | [planner: change descScanFactor to scanFactor when ExpectedCount is small. (#23972)](https://github.com/pingcap/tidb/pull/24079)                   |   | 11d19h  |
| tidb/24147 | [docs/design: add proposal for common table expression](https://github.com/pingcap/tidb/pull/24147)                                                |   | 7d23h   |
| tidb/24155 | [planner, executor: fix index merge partial table scan schema (#23936)](https://github.com/pingcap/tidb/pull/24155)                                |   | 7d20h   |
| tidb/24214 | [plan: merge continuous selections and delete surely true expressions](https://github.com/pingcap/tidb/pull/24214)                                 |   | 5d12h   |
| tidb/24236 | [*: remove SchemaVersion in TransactionContext](https://github.com/pingcap/tidb/pull/24236)                                                        |   | 4d17h   |
| tidb/24258 | [Revert "planner: donot prune all columns for Projection (#24024)" (#24180)](https://github.com/pingcap/tidb/pull/24258)                           |   | 2d22h   |
| tidb/24317 | [statistics: skip reading mysql.stats_histograms if cached stats is up-to-date (#24175)](https://github.com/pingcap/tidb/pull/24317)               |   | 17h     |


## Reviewed in Last 7 Days

|    REPO     |                                                       PR                                                        | C | D |   R   |
|-------------|-----------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/24180  | [Revert "planner: donot prune all columns for Projection (#24024)"](https://github.com/pingcap/tidb/pull/24180) |   | 3 | 3d23h |
| docs/5392   | [releases: add tidb 5.0.1 release notes](https://github.com/pingcap/docs/pull/5392)                             |   | 6 | 2d23h |
| blog-cn/566 | [fix: remove the out-date description](https://github.com/pingcap/blog-cn/pull/566)                             |   | 7 | 1d4h  |
| tidb/24102  | [planner: Fix Join reorder occurs "index out of range" error](https://github.com/pingcap/tidb/pull/24102)       |   | 7 | 2d6h  |
| tipb/220    | [analyze: add proto fields for row based sampling](https://github.com/pingcap/tipb/pull/220)                    |   | 7 | 5d23h |
| tidb/24089  | [statistics: introduce the weighted reservoir sampling](https://github.com/pingcap/tidb/pull/24089)             |   | 7 | 4d23h |


</details> 


<details> 
  <summary>fzhedu's detailed review report</summary> 

## To Be Reviewed

| REPO | PR | C | LASTED |
|------|----|---|--------|


## Reviewed in Last 7 Days

|    REPO    |                                                           PR                                                           | C | D |  R   |
|------------|------------------------------------------------------------------------------------------------------------------------|---|---|------|
| tidb/24231 | [executor: fix projection executor panic and add failpoint test](https://github.com/pingcap/tidb/pull/24231)           |   | 3 | 2d1h |
| tidb/24157 | [planner: let CopTiFlashConcurrencyFactor inflence the cost of whole plan](https://github.com/pingcap/tidb/pull/24157) |   | 5 | 3d3h |


</details> 


<details> 
  <summary>ichn-hu's detailed review report</summary> 

## To Be Reviewed

| REPO | PR | C | LASTED |
|------|----|---|--------|


## Reviewed in Last 7 Days

|    REPO    |                                                                   PR                                                                   | C | D |   R   |
|------------|----------------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/24266 | [expression: fix wrong flen infer for bit constant (#23867)](https://github.com/pingcap/tidb/pull/24266)                               |   | 1 | 2d2h  |
| tidb/24267 | [expression: fix wrong flen infer for bit constant (#23867)](https://github.com/pingcap/tidb/pull/24267)                               |   | 1 | 2d2h  |
| tidb/24300 | [expression: fix wrong collation for `concat` function (#24297)](https://github.com/pingcap/tidb/pull/24300)                           |   | 1 | 20h   |
| tidb/24304 | [executor, variable: move hidden variables to struct](https://github.com/pingcap/tidb/pull/24304)                                      |   | 1 | 12h   |
| tidb/24248 | [executor, variable: small cleanup](https://github.com/pingcap/tidb/pull/24248)                                                        |   | 1 | 3d14h |
| tidb/24278 | [executor: accelerate TestVectorizedMergeJoin and TestVectorizedShuffleMergeJoin (#24177)](https://github.com/pingcap/tidb/pull/24278) |   | 2 | 14h   |
| tidb/24177 | [executor: accelerate TestVectorizedMergeJoin and TestVectorizedShuffleMergeJoin](https://github.com/pingcap/tidb/pull/24177)          |   | 3 | 4d15h |
| tidb/24234 | [executor: skip TestMppExecution when race is enabled (#24222)](https://github.com/pingcap/tidb/pull/24234)                            |   | 3 | 1d19h |
| tidb/24235 | [expression: try to fix TestExprPushDownToFlash tests](https://github.com/pingcap/tidb/pull/24235)                                     |   | 5 | 0h    |
| tidb/24026 | [types: fix type merge about bit type (#23857)](https://github.com/pingcap/tidb/pull/24026)                                            |   | 5 | 8d19h |


</details> 


<details> 
  <summary>lzmhhh123's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                           PR                                                                            | C | LASTED  |
|------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20444 | [expression: add json_merge_patch](https://github.com/pingcap/tidb/pull/20444)                                                                          |   | 195d21h |
| tidb/20465 | [expression: add uuidShortFunction](https://github.com/pingcap/tidb/pull/20465)                                                                         |   | 194d19h |
| tidb/20642 | [executor: modify admin executors to support partitioned table with global index](https://github.com/pingcap/tidb/pull/20642)                           |   | 183d15h |
| tidb/20903 | [planner: fix confused and unnecessary double-projection in plans.](https://github.com/pingcap/tidb/pull/20903)                                         |   | 172d17h |
| tidb/21018 | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                          |   | 166d17h |
| tidb/21195 | [brie: integrate lightning to suport IMPORT statement](https://github.com/pingcap/tidb/pull/21195)                                                      |   | 155d22h |
| tidb/21334 | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                                           |   | 152d14h |
| tidb/21347 | [session: make rollback work on global variables](https://github.com/pingcap/tidb/pull/21347)                                                           |   | 151d19h |
| tidb/21487 | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                                           |   | 145d4h  |
| tidb/21641 | [executor: Fix pessimistic lock doesn't work on the partition table for subquery/joins](https://github.com/pingcap/tidb/pull/21641)                     |   | 138d18h |
| tidb/21651 | [planner: allow filter condition pushing down to IndexScan for prefix index](https://github.com/pingcap/tidb/pull/21651)                                |   | 138d13h |
| tidb/22126 | [*: add `sys` schema, `sys.SCHEMA_UNUSED_INDEXES` view and `sys.SCHEMA_INDEX_USAGE` view](https://github.com/pingcap/tidb/pull/22126)                   |   | 117d19h |
| tidb/22361 | [table: fix insert into _tidb_rowid panic and rebase it if needed (#22062)](https://github.com/pingcap/tidb/pull/22361)                                 |   | 105d20h |
| tidb/22372 | [executor: fix SelectForUpdate in decorrelated subquery under pessimistic mode](https://github.com/pingcap/tidb/pull/22372)                             |   | 105d9h  |
| tidb/22478 | [planner, executor: fix query partition table with global unique index get wrong result](https://github.com/pingcap/tidb/pull/22478)                    |   | 96d13h  |
| tidb/22631 | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                                                                         |   | 86d23h  |
| tidb/22699 | [brie: add error info column and history backup/restore info in sql](https://github.com/pingcap/tidb/pull/22699)                                        |   | 83d16h  |
| tidb/23022 | [executor: create PipelinedWindowExec](https://github.com/pingcap/tidb/pull/23022)                                                                      |   | 57d18h  |
| tidb/23149 | [core: support left join and right join for join reorder](https://github.com/pingcap/tidb/pull/23149)                                                   |   | 52d12h  |
| tidb/23257 | [executor: group_concat aggr panic when session.group_concat_max_len is small (#23131)](https://github.com/pingcap/tidb/pull/23257)                     |   | 47d18h  |
| tidb/23283 | [util: optimize the performance of restore with db (#22910)](https://github.com/pingcap/tidb/pull/23283)                                                |   | 46d17h  |
| tidb/23347 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                                    |   | 42d18h  |
| tidb/23348 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                                    |   | 42d18h  |
| tidb/23368 | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23368)                                   |   | 41d20h  |
| tidb/23373 | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                       |   | 41d19h  |
| tidb/23655 | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral (#23523)](https://github.com/pingcap/tidb/pull/23655)                |   | 28d22h  |
| tidb/23660 | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23660) |   | 28d20h  |
| tidb/23661 | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23661) |   | 28d20h  |
| tidb/23703 | [expression: fix approx_percent panic on bit column (#23687)](https://github.com/pingcap/tidb/pull/23703)                                               |   | 28d14h  |
| tidb/23705 | [executor: refineArgs() bug fix when compare int with very small decimal (#23694)](https://github.com/pingcap/tidb/pull/23705)                          |   | 28d13h  |
| tidb/23756 | [planner: fix set not null flag for outer join (#23727)](https://github.com/pingcap/tidb/pull/23756)                                                    |   | 27d14h  |
| tidb/23760 | [collation: fix tidb panic when compare string with collation](https://github.com/pingcap/tidb/pull/23760)                                              |   | 27d13h  |
| tidb/23812 | [executor, planner: fix collation for hash join building (#23770)](https://github.com/pingcap/tidb/pull/23812)                                          |   | 26d12h  |
| tidb/23940 | [config, ddl: allow auto inc columns in generated columns and expression indexes](https://github.com/pingcap/tidb/pull/23940)                           |   | 17d18h  |
| tidb/23968 | [statistics: fix unstable TestDropPartitionStats test](https://github.com/pingcap/tidb/pull/23968)                                                      |   | 15d14h  |
| tidb/23987 | [executor: Implements json_arrayagg function](https://github.com/pingcap/tidb/pull/23987)                                                               |   | 14d18h  |
| tidb/24016 | [planner: fix index-out-of-range error when checking only_full_group_by (#23844)](https://github.com/pingcap/tidb/pull/24016)                           |   | 13d18h  |
| tidb/24018 | [ranger: fix the range construction behavior when the column's type is `YEAR` (#23559)](https://github.com/pingcap/tidb/pull/24018)                     |   | 13d18h  |
| tidb/24054 | [executor: fix wrong convert from bit to string when do projection (#23960)](https://github.com/pingcap/tidb/pull/24054)                                |   | 12d16h  |
| tidb/24151 | [ddl: admin show ddl jobs output confusing with multiple jobs](https://github.com/pingcap/tidb/pull/24151)                                              |   | 7d21h   |
| tidb/24155 | [planner, executor: fix index merge partial table scan schema (#23936)](https://github.com/pingcap/tidb/pull/24155)                                     |   | 7d20h   |
| tidb/24185 | [executor: make column default value being aware of NO_ZERO_IN_DATE (#24174)](https://github.com/pingcap/tidb/pull/24185)                               |   | 6d19h   |
| tidb/24186 | [executor: make column default value being aware of NO_ZERO_IN_DATE (#24174)](https://github.com/pingcap/tidb/pull/24186)                               |   | 6d19h   |
| tidb/24188 | [util: fix bad number error with DISTINCT when dividing long decimals (#21783)](https://github.com/pingcap/tidb/pull/24188)                             |   | 6d16h   |
| tidb/24211 | [*: support txn retry when auto id meets duplicate entry](https://github.com/pingcap/tidb/pull/24211)                                                   |   | 5d13h   |
| tidb/24231 | [executor: fix projection executor panic and add failpoint test](https://github.com/pingcap/tidb/pull/24231)                                            |   | 4d20h   |
| tidb/24234 | [executor: skip TestMppExecution when race is enabled (#24222)](https://github.com/pingcap/tidb/pull/24234)                                             |   | 4d18h   |
| tidb/24250 | [planner: rewritten LIKE as range for expression index](https://github.com/pingcap/tidb/pull/24250)                                                     |   | 3d21h   |
| tidb/24258 | [Revert "planner: donot prune all columns for Projection (#24024)" (#24180)](https://github.com/pingcap/tidb/pull/24258)                                |   | 2d22h   |
| tidb/24268 | [expression: fix cast real, decimal to time (#24120)](https://github.com/pingcap/tidb/pull/24268)                                                       |   | 2d17h   |
| tidb/24285 | [*: compatibility with staleread](https://github.com/pingcap/tidb/pull/24285)                                                                           |   | 1d19h   |
| tidb/24304 | [executor, variable: move hidden variables to struct](https://github.com/pingcap/tidb/pull/24304)                                                       |   | 1d3h    |
| tidb/24335 | [store/tikv: remove IsolationLevel option](https://github.com/pingcap/tidb/pull/24335)                                                                  |   | 2h      |


## Reviewed in Last 7 Days

|    REPO     |                                                                   PR                                                                   | C | D |   R   |
|-------------|----------------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/24278  | [executor: accelerate TestVectorizedMergeJoin and TestVectorizedShuffleMergeJoin (#24177)](https://github.com/pingcap/tidb/pull/24278) |   | 2 | 15h   |
| tidb/24179  | [expression: fix float64 overflow check in plus/minus real function](https://github.com/pingcap/tidb/pull/24179)                       |   | 3 | 4d3h  |
| tipb/221    | [add more regexp functions](https://github.com/pingcap/tipb/pull/221)                                                                  |   | 3 | 1d21h |
| tidb/24180  | [Revert "planner: donot prune all columns for Projection (#24024)"](https://github.com/pingcap/tidb/pull/24180)                        |   | 5 | 2d4h  |
| tidb/24212  | [*: turn on unused linter](https://github.com/pingcap/tidb/pull/24212)                                                                 |   | 5 | 13h   |
| tidb/22686  | [expression: support enum pushdown](https://github.com/pingcap/tidb/pull/22686)                                                        |   | 6 | 78d5h |
| tidb/24202  | [Revert "planner: donot prune all columns for Projection (#24024) (#24093)"](https://github.com/pingcap/tidb/pull/24202)               |   | 6 | 1h    |
| tidb/24120  | [expression: fix cast real, decimal to time](https://github.com/pingcap/tidb/pull/24120)                                               |   | 6 | 3d1h  |
| tidb/24053  | [executor: fix wrong convert from bit to string when do projection (#23960)](https://github.com/pingcap/tidb/pull/24053)               |   | 6 | 6d21h |
| tidb/24191  | [executor: accelerate TestSortInDisk](https://github.com/pingcap/tidb/pull/24191)                                                      |   | 6 | 16h   |
| tikv/10048  | [copr: fix cast real, decimal to time bug on overflow](https://github.com/tikv/tikv/pull/10048)                                        | Y | 6 | 2d0h  |
| parser/1165 | [Add EnumSetAsIntFlag flag to control enum behavior.](https://github.com/pingcap/parser/pull/1165)                                     |   | 7 | 76d3h |
| tidb/24178  | [planner/core: point get only work on TiKV](https://github.com/pingcap/tidb/pull/24178)                                                |   | 7 | 11h   |


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
| tidb/19029   | [types: fix unexpected NOT_NULL flags](https://github.com/pingcap/tidb/pull/19029)                                                                      |   | 264d22h |
| docs-cn/5561 | [Add sql optimization-related docs to toc](https://github.com/pingcap/docs-cn/pull/5561)                                                                |   | 64d15h  |
| docs/5498    | [partitioning: Corrected partition management](https://github.com/pingcap/docs/pull/5498)                                                               |   | 1d19h   |
| tidb/20708   | [*: separate auto_increment ID allocator from _tidb_rowid allocator](https://github.com/pingcap/tidb/pull/20708)                                        |   | 180d20h |
| tidb/20969   | [executor: Improve the performance of appending not fixed columns](https://github.com/pingcap/tidb/pull/20969)                                          |   | 168d9h  |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                          |   | 166d17h |
| tidb/21149   | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                           |   | 159d14h |
| tidb/21304   | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                                |   | 153d12h |
| tidb/21318   | [planner, expression: use the range of column types to simplify expressions](https://github.com/pingcap/tidb/pull/21318)                                |   | 152d19h |
| tidb/21401   | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                                      |   | 148d11h |
| tidb/21508   | [execution: fix dayofweek('0000-00-00') behavior](https://github.com/pingcap/tidb/pull/21508)                                                           |   | 144d10h |
| tidb/21641   | [executor: Fix pessimistic lock doesn't work on the partition table for subquery/joins](https://github.com/pingcap/tidb/pull/21641)                     |   | 138d18h |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                             |   | 130d19h |
| tidb/21887   | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                                         |   | 129d11h |
| tidb/21954   | [planner/cascades: add rule `PushSelDownApply`](https://github.com/pingcap/tidb/pull/21954)                                                             |   | 125d23h |
| tidb/22146   | [executor: forbid SFU on view](https://github.com/pingcap/tidb/pull/22146)                                                                              |   | 113d21h |
| tidb/22217   | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                                           |   | 111d17h |
| tidb/22234   | [executor, planner: ON DUPLICATE UPDATE can refer to un-project col (#14412)](https://github.com/pingcap/tidb/pull/22234)                               |   | 111d15h |
| tidb/22261   | [time: fix parse datetime won't truncate the reluctant string (#22232)](https://github.com/pingcap/tidb/pull/22261)                                     |   | 110d19h |
| tidb/22374   | [expression: separated arithmeticIntDivideSig](https://github.com/pingcap/tidb/pull/22374)                                                              |   | 105d0h  |
| tidb/22415   | [ddl: refactor bundle[2/2] [6/6]](https://github.com/pingcap/tidb/pull/22415)                                                                           |   | 101d17h |
| tidb/22416   | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                                                       | Y | 101d11h |
| tidb/22541   | [expression: Support builtin function SOUNDEX](https://github.com/pingcap/tidb/pull/22541)                                                              |   | 91d9h   |
| tidb/22565   | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)                                      | Y | 90d17h  |
| tidb/22814   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22814)                                     |   | 67d19h  |
| tidb/22862   | [brie: fix the problem that ddl restored by BR via SQL is not replicated to downstream](https://github.com/pingcap/tidb/pull/22862)                     |   | 64d22h  |
| tidb/22923   | [expression: correct constant propagation for collation (#22666)](https://github.com/pingcap/tidb/pull/22923)                                           |   | 62d15h  |
| tidb/22924   | [planner: fix wrong index merge selection (#22825)](https://github.com/pingcap/tidb/pull/22924)                                                         |   | 62d14h  |
| tidb/23002   | [store/*: fix err check](https://github.com/pingcap/tidb/pull/23002)                                                                                    |   | 59d0h   |
| tidb/23022   | [executor: create PipelinedWindowExec](https://github.com/pingcap/tidb/pull/23022)                                                                      |   | 57d18h  |
| tidb/23152   | [expression: fix wrong error info (#22760)](https://github.com/pingcap/tidb/pull/23152)                                                                 |   | 51d14h  |
| tidb/23196   | [types: fix the bug about the wrong query result for decimal type  (#22507)](https://github.com/pingcap/tidb/pull/23196)                                |   | 49d18h  |
| tidb/23295   | [util, types: don't let SPM be affected by charset (#23161)](https://github.com/pingcap/tidb/pull/23295)                                                |   | 46d11h  |
| tidb/23316   | [planner: Fix rebuild range for prepared plan](https://github.com/pingcap/tidb/pull/23316)                                                              |   | 43d17h  |
| tidb/23373   | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                       |   | 41d19h  |
| tidb/23397   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23397)                                                          |   | 40d17h  |
| tidb/23398   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23398)                                                          |   | 40d17h  |
| tidb/23590   | [planner, table: optimize the list partition pruner for range query](https://github.com/pingcap/tidb/pull/23590)                                        |   | 32d16h  |
| tidb/23598   | [types: fix collation for binary literal (#23591)](https://github.com/pingcap/tidb/pull/23598)                                                          |   | 32d13h  |
| tidb/23655   | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral (#23523)](https://github.com/pingcap/tidb/pull/23655)                |   | 28d22h  |
| tidb/23658   | [*: collect transaction write duration/throughput metrics for SLI/SLO (#23462)](https://github.com/pingcap/tidb/pull/23658)                             |   | 28d22h  |
| tidb/23660   | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23660) |   | 28d20h  |
| tidb/23661   | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23661) |   | 28d20h  |
| tidb/23682   | [executor: fix a panic when batch point get is used for partition table (#23652)](https://github.com/pingcap/tidb/pull/23682)                           |   | 28d16h  |
| tidb/23689   | [planner: fix the panic when we calculate the partition range (#23651)](https://github.com/pingcap/tidb/pull/23689)                                     |   | 28d16h  |
| tidb/23730   | [distsql/*: typo fix for `dispatches`](https://github.com/pingcap/tidb/pull/23730)                                                                      |   | 27d18h  |
| tidb/23796   | [tests: make TestIndexLookupMergeJoinHang and TestIssue18068 stable (#23741)](https://github.com/pingcap/tidb/pull/23796)                               |   | 26d19h  |
| tidb/23812   | [executor, planner: fix collation for hash join building (#23770)](https://github.com/pingcap/tidb/pull/23812)                                          |   | 26d12h  |
| tidb/23878   | [functions: fix some string function has wrong collation and flag (#23835)](https://github.com/pingcap/tidb/pull/23878)                                 |   | 20d21h  |
| tidb/23963   | [executor: checking chunk is full precedes filtering](https://github.com/pingcap/tidb/pull/23963)                                                       |   | 15d17h  |
| tidb/23987   | [executor: Implements json_arrayagg function](https://github.com/pingcap/tidb/pull/23987)                                                               |   | 14d18h  |
| tidb/23997   | [stats, executor: use a correct sampling to collect stats](https://github.com/pingcap/tidb/pull/23997)                                                  |   | 14d9h   |
| tidb/24018   | [ranger: fix the range construction behavior when the column's type is `YEAR` (#23559)](https://github.com/pingcap/tidb/pull/24018)                     |   | 13d18h  |
| tidb/24188   | [util: fix bad number error with DISTINCT when dividing long decimals (#21783)](https://github.com/pingcap/tidb/pull/24188)                             |   | 6d16h   |
| tidb/24193   | [executor: implement CTEStorage](https://github.com/pingcap/tidb/pull/24193)                                                                            |   | 6d10h   |
| tidb/24214   | [plan: merge continuous selections and delete surely true expressions](https://github.com/pingcap/tidb/pull/24214)                                      |   | 5d12h   |
| tidb/24229   | [executor: speed up race test TestInsertReorgDelete (#24208)](https://github.com/pingcap/tidb/pull/24229)                                               |   | 4d21h   |
| tidb/24235   | [expression: try to fix TestExprPushDownToFlash tests](https://github.com/pingcap/tidb/pull/24235)                                                      |   | 4d17h   |
| tidb/24241   | [planner/core: remove random test to reduce CI time (#24207)](https://github.com/pingcap/tidb/pull/24241)                                               |   | 4d15h   |
| tidb/24266   | [expression: fix wrong flen infer for bit constant (#23867)](https://github.com/pingcap/tidb/pull/24266)                                                |   | 2d17h   |
| tidb/24267   | [expression: fix wrong flen infer for bit constant (#23867)](https://github.com/pingcap/tidb/pull/24267)                                                |   | 2d17h   |
| tidb/24279   | [*: add security enhanced mode part 2](https://github.com/pingcap/tidb/pull/24279)                                                                      |   | 2d3h    |
| tidb/24282   | [planner: remove useless predicates after partition pruning](https://github.com/pingcap/tidb/pull/24282)                                                |   | 1d20h   |
| tidb/24328   | [*: implement read_ts_in built-in function](https://github.com/pingcap/tidb/pull/24328)                                                                 |   | 10h     |


## Reviewed in Last 7 Days

|      REPO      |                                                                 PR                                                                  | C | D |   R   |
|----------------|-------------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/24313     | [planner: filter conflict read_from_storage hints](https://github.com/pingcap/tidb/pull/24313)                                      |   | 1 | 2h    |
| tidb-test/1185 | [randgen-test: fix test for tidb issue:16845](https://github.com/pingcap/tidb-test/pull/1185)                                       |   | 1 | 2h    |
| tidb/24290     | [expression: fix wrong type infer for agg function when type is null](https://github.com/pingcap/tidb/pull/24290)                   |   | 2 | 1h    |
| tidb/24245     | [planner: between .. and on int column can be used to prune hash partition](https://github.com/pingcap/tidb/pull/24245)             |   | 2 | 2d13h |
| community/439  | [planner: promote member and update links](https://github.com/pingcap/community/pull/439)                                           |   | 7 | 0h    |
| tidb/24183     | [executor, statistics: remove NULL value from column histogram created by fast analyze](https://github.com/pingcap/tidb/pull/24183) |   | 7 | 0h    |


</details> 


<details> 
  <summary>rebelice's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                     PR                                                     | C | LASTED |
|------------|------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/23836 | [parser, core: Implement force_index hint in parser and TiDB](https://github.com/pingcap/tidb/pull/23836)  |   | 25d17h |
| tidb/24033 | [statistics: fix some unstable tests in global stats (#23502)](https://github.com/pingcap/tidb/pull/24033) |   | 13d9h  |
| tidb/24282 | [planner: remove useless predicates after partition pruning](https://github.com/pingcap/tidb/pull/24282)   |   | 1d20h  |
| tidb/24306 | [util/ranger: fix func name typo](https://github.com/pingcap/tidb/pull/24306)                              |   | 22h    |


## Reviewed in Last 7 Days

|    REPO    |                                                           PR                                                            | C | D |   R   |
|------------|-------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/24245 | [planner: between .. and on int column can be used to prune hash partition](https://github.com/pingcap/tidb/pull/24245) |   | 2 | 2d15h |


</details> 


<details> 
  <summary>time-and-fate's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                  PR                                                                  | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                            |   | 173d17h |
| tidb/22416 | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                                    | Y | 101d11h |
| tidb/24155 | [planner, executor: fix index merge partial table scan schema (#23936)](https://github.com/pingcap/tidb/pull/24155)                  |   | 7d20h   |
| tidb/24204 | [planner: clone possible properties before saving them](https://github.com/pingcap/tidb/pull/24204)                                  |   | 5d17h   |
| tidb/24230 | [*: consitent get infoschema](https://github.com/pingcap/tidb/pull/24230)                                                            |   | 4d21h   |
| tidb/24313 | [planner: filter conflict read_from_storage hints](https://github.com/pingcap/tidb/pull/24313)                                       |   | 17h     |
| tidb/24317 | [statistics: skip reading mysql.stats_histograms if cached stats is up-to-date (#24175)](https://github.com/pingcap/tidb/pull/24317) |   | 17h     |


## Reviewed in Last 7 Days

|    REPO    |                                                                 PR                                                                  | C | D |   R    |
|------------|-------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/23997 | [stats, executor: use a correct sampling to collect stats](https://github.com/pingcap/tidb/pull/23997)                              |   | 1 | 13d18h |
| tidb/24204 | [planner: clone possible properties before saving them](https://github.com/pingcap/tidb/pull/24204)                                 |   | 5 | 1d1h   |
| tidb/24089 | [statistics: introduce the weighted reservoir sampling](https://github.com/pingcap/tidb/pull/24089)                                 |   | 6 | 5d23h  |
| tipb/220   | [analyze: add proto fields for row based sampling](https://github.com/pingcap/tipb/pull/220)                                        |   | 6 | 6d22h  |
| tidb/24175 | [statistics: skip reading mysql.stats_histograms if cached stats is up-to-date](https://github.com/pingcap/tidb/pull/24175)         |   | 7 | 1d0h   |
| tidb/24183 | [executor, statistics: remove NULL value from column histogram created by fast analyze](https://github.com/pingcap/tidb/pull/24183) |   | 7 | 1h     |


</details> 


<details> 
  <summary>winoros's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                              PR                                                                               | C | LASTED  |
|--------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| docs-cn/5916 | [sql-statements, information-schema: add `END_TIME` field for table `ANALYZE_STATUS`](https://github.com/pingcap/docs-cn/pull/5916)                           |   | 26d17h  |
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                                        | Y | 228d14h |
| docs-cn/6113 | [config: update the default value of `feedback-probability`](https://github.com/pingcap/docs-cn/pull/6113)                                                    |   | 5d22h   |
| tidb/20765   | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                                             |   | 176d17h |
| tidb/20877   | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                                                     |   | 173d17h |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                                |   | 166d17h |
| tidb/21207   | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                                               |   | 155d19h |
| tidb/21487   | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                                                 |   | 145d4h  |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                                   |   | 130d19h |
| tidb/21954   | [planner/cascades: add rule `PushSelDownApply`](https://github.com/pingcap/tidb/pull/21954)                                                                   |   | 125d23h |
| tidb/22181   | [planner, expression: fix error when using IN combined with subquery (#22080)](https://github.com/pingcap/tidb/pull/22181)                                    |   | 112d17h |
| tidb/22416   | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                                                             | Y | 101d11h |
| tidb/22504   | [*:Fix the fetchHotRegion bug that the count always zero](https://github.com/pingcap/tidb/pull/22504)                                                         |   | 93d19h  |
| tidb/22565   | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)                                            | Y | 90d17h  |
| tidb/22624   | [ planner: not pruning column used by union scan condition (#21640)](https://github.com/pingcap/tidb/pull/22624)                                              |   | 88d17h  |
| tidb/22923   | [expression: correct constant propagation for collation (#22666)](https://github.com/pingcap/tidb/pull/22923)                                                 |   | 62d15h  |
| tidb/23233   | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23233)                                                        |   | 48d18h  |
| tidb/23347   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                                          |   | 42d18h  |
| tidb/23348   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                                          |   | 42d18h  |
| tidb/23350   | [util/stringutil, util/ranger, planner: use hierarchical separators to simplify the parsing for info of EXPLAIN ](https://github.com/pingcap/tidb/pull/23350) |   | 42d17h  |
| tidb/23373   | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                             |   | 41d19h  |
| tidb/23474   | [planner: fix inappropriate null flag of null constants (#23457)](https://github.com/pingcap/tidb/pull/23474)                                                 |   | 35d18h  |
| tidb/23543   | [statistics: fix auto analyze log information incomplete (#23522)](https://github.com/pingcap/tidb/pull/23543)                                                |   | 33d18h  |
| tidb/23598   | [types: fix collation for binary literal (#23591)](https://github.com/pingcap/tidb/pull/23598)                                                                |   | 32d13h  |
| tidb/23655   | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral (#23523)](https://github.com/pingcap/tidb/pull/23655)                      |   | 28d22h  |
| tidb/23689   | [planner: fix the panic when we calculate the partition range (#23651)](https://github.com/pingcap/tidb/pull/23689)                                           |   | 28d16h  |
| tidb/23772   | [tablecodec: fix text type decode for old row format (#23751)](https://github.com/pingcap/tidb/pull/23772)                                                    |   | 27d11h  |
| tidb/23849   | [ddl: tidb panic while query hash partition table with is null condition](https://github.com/pingcap/tidb/pull/23849)                                         |   | 22d13h  |
| tidb/23917   | [planner: fix wrong TableDual plans caused by comparing Binary and Bytes incorrectly (#23860)](https://github.com/pingcap/tidb/pull/23917)                    |   | 19d23h  |
| tidb/23970   | [planner: fix a bug that point get plan returns wrong column name (#23365)](https://github.com/pingcap/tidb/pull/23970)                                       |   | 15d14h  |
| tidb/24018   | [ranger: fix the range construction behavior when the column's type is `YEAR` (#23559)](https://github.com/pingcap/tidb/pull/24018)                           |   | 13d18h  |
| tidb/24060   | [statistics: fix some potential panic in statistics (#23988)](https://github.com/pingcap/tidb/pull/24060)                                                     |   | 12d13h  |
| tidb/24061   | [statistics: fix some potential panic in statistics (#23988)](https://github.com/pingcap/tidb/pull/24061)                                                     |   | 12d13h  |
| tidb/24079   | [planner: change descScanFactor to scanFactor when ExpectedCount is small. (#23972)](https://github.com/pingcap/tidb/pull/24079)                              |   | 11d19h  |
| tidb/24097   | [planner: Remove redundant call to expression.ColumnSubstitute](https://github.com/pingcap/tidb/pull/24097)                                                   |   | 11d2h   |
| tidb/24138   | [planner: Add Equivalence Rules to Transform BinaryOptSubquery to ExistsSubquery](https://github.com/pingcap/tidb/pull/24138)                                 |   | 8d12h   |
| tidb/24204   | [planner: clone possible properties before saving them](https://github.com/pingcap/tidb/pull/24204)                                                           |   | 5d17h   |
| tidb/24241   | [planner/core: remove random test to reduce CI time (#24207)](https://github.com/pingcap/tidb/pull/24241)                                                     |   | 4d15h   |
| tidb/24258   | [Revert "planner: donot prune all columns for Projection (#24024)" (#24180)](https://github.com/pingcap/tidb/pull/24258)                                      |   | 2d22h   |
| tidb/24317   | [statistics: skip reading mysql.stats_histograms if cached stats is up-to-date (#24175)](https://github.com/pingcap/tidb/pull/24317)                          |   | 17h     |


## Reviewed in Last 7 Days

|     REPO     |                                                            PR                                                            | C | D |  R   |
|--------------|--------------------------------------------------------------------------------------------------------------------------|---|---|------|
| docs-cn/6061 | [releases: add tidb 5.0.1 release notes](https://github.com/pingcap/docs-cn/pull/6061)                                   |   | 6 | 3d2h |
| tidb/24202   | [Revert "planner: donot prune all columns for Projection (#24024) (#24093)"](https://github.com/pingcap/tidb/pull/24202) |   | 6 | 0h   |


</details> 


<details> 
  <summary>wshwsh12's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                 PR                                                                  | C | LASTED  |
|------------|-------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19807 | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                             |   | 235d11h |
| tidb/19957 | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                              | Y | 228d14h |
| tidb/21487 | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                       |   | 145d4h  |
| tidb/21887 | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                     |   | 129d11h |
| tidb/22378 | [executor: vectorize hash aggregate](https://github.com/pingcap/tidb/pull/22378)                                                    |   | 104d19h |
| tidb/23336 | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23336)            |   | 42d19h  |
| tidb/23347 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                |   | 42d18h  |
| tidb/23348 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                |   | 42d18h  |
| tidb/23368 | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23368)               |   | 41d20h  |
| tidb/23397 | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23397)                                      |   | 40d17h  |
| tidb/23398 | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23398)                                      |   | 40d17h  |
| tidb/23519 | [executor: check privilege before adding](https://github.com/pingcap/tidb/pull/23519)                                               |   | 34d0h   |
| tidb/23760 | [collation: fix tidb panic when compare string with collation](https://github.com/pingcap/tidb/pull/23760)                          |   | 27d13h  |
| tidb/23968 | [statistics: fix unstable TestDropPartitionStats test](https://github.com/pingcap/tidb/pull/23968)                                  |   | 15d14h  |
| tidb/23979 | [executor, statistics: fix unstable `TestAnalyzeIndexExtractTopN`](https://github.com/pingcap/tidb/pull/23979)                      |   | 14d23h  |
| tidb/24018 | [ranger: fix the range construction behavior when the column's type is `YEAR` (#23559)](https://github.com/pingcap/tidb/pull/24018) |   | 13d18h  |
| tidb/24033 | [statistics: fix some unstable tests in global stats (#23502)](https://github.com/pingcap/tidb/pull/24033)                          |   | 13d9h   |
| tidb/24050 | [expression: fix get var panic when types not match](https://github.com/pingcap/tidb/pull/24050)                                    |   | 12d17h  |
| tidb/24053 | [executor: fix wrong convert from bit to string when do projection (#23960)](https://github.com/pingcap/tidb/pull/24053)            |   | 12d16h  |
| tidb/24054 | [executor: fix wrong convert from bit to string when do projection (#23960)](https://github.com/pingcap/tidb/pull/24054)            |   | 12d16h  |
| tidb/24147 | [docs/design: add proposal for common table expression](https://github.com/pingcap/tidb/pull/24147)                                 |   | 7d23h   |
| tidb/24186 | [executor: make column default value being aware of NO_ZERO_IN_DATE (#24174)](https://github.com/pingcap/tidb/pull/24186)           |   | 6d19h   |
| tidb/24228 | [executor: skip TestPrepareStmtAfterIsolationReadChange when race enable (#24200)](https://github.com/pingcap/tidb/pull/24228)      |   | 4d22h   |
| tidb/24229 | [executor: speed up race test TestInsertReorgDelete (#24208)](https://github.com/pingcap/tidb/pull/24229)                           |   | 4d21h   |
| tidb/24230 | [*: consitent get infoschema](https://github.com/pingcap/tidb/pull/24230)                                                           |   | 4d21h   |
| tidb/24236 | [*: remove SchemaVersion in TransactionContext](https://github.com/pingcap/tidb/pull/24236)                                         |   | 4d17h   |
| tidb/24257 | [executor: fix data race of parallel apply operator](https://github.com/pingcap/tidb/pull/24257)                                    |   | 2d23h   |
| tidb/24266 | [expression: fix wrong flen infer for bit constant (#23867)](https://github.com/pingcap/tidb/pull/24266)                            |   | 2d17h   |
| tidb/24267 | [expression: fix wrong flen infer for bit constant (#23867)](https://github.com/pingcap/tidb/pull/24267)                            |   | 2d17h   |
| tidb/24268 | [expression: fix cast real, decimal to time (#24120)](https://github.com/pingcap/tidb/pull/24268)                                   |   | 2d17h   |
| tidb/24280 | [executor, session, variable: Move deprecation and synonyms to sysvar struct](https://github.com/pingcap/tidb/pull/24280)           |   | 2d0h    |
| tidb/24290 | [expression: fix wrong type infer for agg function when type is null](https://github.com/pingcap/tidb/pull/24290)                   |   | 1d18h   |


## Reviewed in Last 7 Days

|    REPO    |                                                                PR                                                                | C | D |   R    |
|------------|----------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/24234 | [executor: skip TestMppExecution when race is enabled (#24222)](https://github.com/pingcap/tidb/pull/24234)                      |   | 2 | 3d6h   |
| tidb/24260 | [executor: make IndexReaders in the inner side of IndexJoins support direct reading](https://github.com/pingcap/tidb/pull/24260) |   | 2 | 1d8h   |
| tidb/22686 | [expression: support enum pushdown](https://github.com/pingcap/tidb/pull/22686)                                                  |   | 2 | 82d2h  |
| tidb/23867 | [expression: fix wrong flen infer for bit constant](https://github.com/pingcap/tidb/pull/23867)                                  |   | 5 | 16d20h |
| tidb/24222 | [executor: skip TestMppExecution when race is enabled](https://github.com/pingcap/tidb/pull/24222)                               |   | 5 | 0h     |
| tidb/24208 | [executor: speed up race test TestInsertReorgDelete](https://github.com/pingcap/tidb/pull/24208)                                 |   | 5 | 14h    |
| tidb/24200 | [executor: skip TestPrepareStmtAfterIsolationReadChange when race enable](https://github.com/pingcap/tidb/pull/24200)            |   | 5 | 19h    |


</details> 

