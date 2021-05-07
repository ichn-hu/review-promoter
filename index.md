|   REVIEWER    |  O(C)   | 1 | 2 | 3 | 4 | 5 | 6 | 7 | T |
|---------------|---------|---|---|---|---|---|---|---|---|
| Reminiscent   |  6 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| SunRunAway    | 12 ( 1) | 1 | 0 | 0 | 0 | 0 | 0 | 1 | 2 |
| XuHuaiyu      | 80 ( 2) | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| dyzsr         |  1 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| eurekaka      | 20 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| fzhedu        |  2 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| ichn-hu       |  2 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| lzmhhh123     | 54 ( 0) | 4 | 0 | 0 | 0 | 0 | 0 | 1 | 5 |
| nullnotnil    |  0 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| qw4990        | 67 ( 2) | 2 | 0 | 0 | 0 | 0 | 0 | 0 | 2 |
| rebelice      |  6 ( 0) | 2 | 0 | 0 | 0 | 0 | 0 | 0 | 2 |
| time-and-fate |  8 ( 1) | 2 | 0 | 0 | 0 | 0 | 0 | 0 | 2 |
| winoros       | 39 ( 3) | 3 | 0 | 0 | 0 | 0 | 0 | 0 | 3 |
| wshwsh12      | 38 ( 1) | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |


<details> 
  <summary>Reminiscent's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                     PR                                                                     | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/21896 | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                                 |   | 136d19h |
| tidb/23474 | [planner: fix inappropriate null flag of null constants (#23457)](https://github.com/pingcap/tidb/pull/23474)                              |   | 44d18h  |
| tidb/23575 | [executor: fix update panic on join having statement (#23554)](https://github.com/pingcap/tidb/pull/23575)                                 |   | 41d21h  |
| tidb/23917 | [planner: fix wrong TableDual plans caused by comparing Binary and Bytes incorrectly (#23860)](https://github.com/pingcap/tidb/pull/23917) |   | 28d23h  |
| tidb/24016 | [planner: fix index-out-of-range error when checking only_full_group_by (#23844)](https://github.com/pingcap/tidb/pull/24016)              |   | 22d19h  |
| tidb/24357 | [statistics: fix a statistics GC problem that can cause duplicated fm-sketch records (#23830)](https://github.com/pingcap/tidb/pull/24357) |   | 8d14h   |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>SunRunAway's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                  PR                                                                   | C | LASTED  |
|------------|---------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19178 | [executor: Refactor probe channel](https://github.com/pingcap/tidb/pull/19178)                                                        |   | 266d17h |
| tidb/19807 | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                               |   | 244d11h |
| tidb/19900 | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                        | Y | 239d18h |
| tidb/20140 | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                    |   | 226d22h |
| tidb/20765 | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                     |   | 185d17h |
| tidb/21207 | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                       |   | 164d19h |
| tidb/21834 | [planner: enhanced index range calculation plan](https://github.com/pingcap/tidb/pull/21834)                                          |   | 141d19h |
| tidb/21876 | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)           |   | 139d19h |
| tidb/21878 | [planner: do not push down lock to pointGet/bacthPointGet when selection exists](https://github.com/pingcap/tidb/pull/21878)          |   | 139d18h |
| tidb/21956 | [planner/preprocessor: disallow into-outfile clause in some place](https://github.com/pingcap/tidb/pull/21956)                        |   | 134d23h |
| tidb/22217 | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                         |   | 120d18h |
| tidb/22379 | [[experiment] executor: allow aggregation to spill disk when running out of memory quota](https://github.com/pingcap/tidb/pull/22379) |   | 113d19h |


## Reviewed in Last 7 Days

|     REPO     |                                          PR                                           | C | D |   R    |
|--------------|---------------------------------------------------------------------------------------|---|---|--------|
| tidb/20749   | [executor: support global kill (32 bits)](https://github.com/pingcap/tidb/pull/20749) |   | 1 | 185d8h |
| docs-cn/6194 | [br: add restore to systables](https://github.com/pingcap/docs-cn/pull/6194)          |   | 7 | 2h     |


</details> 


<details> 
  <summary>XuHuaiyu's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                              PR                                                                               | C | LASTED  |
|--------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| docs-cn/5561 | [Add sql optimization-related docs to toc](https://github.com/pingcap/docs-cn/pull/5561)                                                                      |   | 73d15h  |
| docs-cn/5619 | [Update data-type-date-and-time.md](https://github.com/pingcap/docs-cn/pull/5619)                                                                             |   | 69d16h  |
| tidb/19900   | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                                                | Y | 239d18h |
| docs-cn/5671 | [tidb: Add time format description](https://github.com/pingcap/docs-cn/pull/5671)                                                                             |   | 63d11h  |
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                                        | Y | 237d14h |
| tidb/20140   | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                                            |   | 226d22h |
| tidb/20790   | [collation: add pinyin collation for chinese charset support](https://github.com/pingcap/tidb/pull/20790)                                                     |   | 184d21h |
| tidb/20969   | [executor: Improve the performance of appending not fixed columns](https://github.com/pingcap/tidb/pull/20969)                                                |   | 177d10h |
| tidb/21064   | [planner, executor: fix cast not check error](https://github.com/pingcap/tidb/pull/21064)                                                                     |   | 172d9h  |
| tidb/21149   | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                                 |   | 168d14h |
| tidb/21304   | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                                      |   | 162d12h |
| tidb/21334   | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                                                 |   | 161d14h |
| tidb/21401   | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                                            |   | 157d11h |
| tidb/21536   | [executor: add slow-log file meta cache to avoid repeat read file meta information](https://github.com/pingcap/tidb/pull/21536)                               |   | 150d15h |
| tidb/21564   | [ddl: fix Incorrect behavior of NO_ZERO_DATE when altering table](https://github.com/pingcap/tidb/pull/21564)                                                 |   | 149d16h |
| tidb/21896   | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                                                    |   | 136d19h |
| tidb/22131   | [privilege: remove leading and trailing space when create user and role](https://github.com/pingcap/tidb/pull/22131)                                          |   | 126d19h |
| tidb/22163   | [expression: separated arithmeticMinusIntSig](https://github.com/pingcap/tidb/pull/22163)                                                                     |   | 122d13h |
| tidb/22186   | [executor: fix select into outfile with year type column has no data (#22175)](https://github.com/pingcap/tidb/pull/22186)                                    |   | 121d16h |
| tidb/22616   | [expression: from_unixtime accept 64-bit integers](https://github.com/pingcap/tidb/pull/22616)                                                                |   | 97d23h  |
| tidb/22617   | [metrics: fix wrong bucket name of coprocessor cache (#22454)](https://github.com/pingcap/tidb/pull/22617)                                                    |   | 97d23h  |
| tidb/22631   | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                                                                               |   | 95d23h  |
| tidb/22696   | [expression: enable arithmetic Mod push down](https://github.com/pingcap/tidb/pull/22696)                                                                     |   | 92d17h  |
| tidb/22711   | [executor: Fix inline schema name](https://github.com/pingcap/tidb/pull/22711)                                                                                |   | 92d12h  |
| tidb/22722   | [planner, errno: make error code of ErrMixOfGroupFuncAndFields consistent with MySQL](https://github.com/pingcap/tidb/pull/22722)                             |   | 91d21h  |
| tidb/22814   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22814)                                           |   | 76d19h  |
| tidb/22908   | [txn: Add txn state's view](https://github.com/pingcap/tidb/pull/22908)                                                                                       |   | 71d21h  |
| tidb/23012   | [executor: fix affected rows of ddls and complete uint tests](https://github.com/pingcap/tidb/pull/23012)                                                     |   | 67d17h  |
| tidb/23152   | [expression: fix wrong error info (#22760)](https://github.com/pingcap/tidb/pull/23152)                                                                       |   | 60d14h  |
| tidb/23196   | [types: fix the bug about the wrong query result for decimal type  (#22507)](https://github.com/pingcap/tidb/pull/23196)                                      |   | 58d18h  |
| tidb/23220   | [Release 4.0](https://github.com/pingcap/tidb/pull/23220)                                                                                                     |   | 58d11h  |
| tidb/23233   | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23233)                                                        |   | 57d18h  |
| tidb/23257   | [executor: group_concat aggr panic when session.group_concat_max_len is small (#23131)](https://github.com/pingcap/tidb/pull/23257)                           |   | 56d18h  |
| tidb/23295   | [util, types: don't let SPM be affected by charset (#23161)](https://github.com/pingcap/tidb/pull/23295)                                                      |   | 55d11h  |
| tidb/23335   | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23335)                                      |   | 51d19h  |
| tidb/23336   | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23336)                                      |   | 51d19h  |
| tidb/23347   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                                          |   | 51d18h  |
| tidb/23348   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                                          |   | 51d18h  |
| tidb/23350   | [util/stringutil, util/ranger, planner: use hierarchical separators to simplify the parsing for info of EXPLAIN ](https://github.com/pingcap/tidb/pull/23350) |   | 51d17h  |
| tidb/23368   | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23368)                                         |   | 50d20h  |
| tidb/23397   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23397)                                                                |   | 49d18h  |
| tidb/23398   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23398)                                                                |   | 49d17h  |
| tidb/23405   | [domain: remove the exit chan, use context](https://github.com/pingcap/tidb/pull/23405)                                                                       |   | 49d17h  |
| tidb/23433   | [WIP: speed up for slow query logs retrieving ](https://github.com/pingcap/tidb/pull/23433)                                                                   |   | 48d17h  |
| tidb/23474   | [planner: fix inappropriate null flag of null constants (#23457)](https://github.com/pingcap/tidb/pull/23474)                                                 |   | 44d18h  |
| tidb/23497   | [expression: Let TiDB use Hyperscan to support multi-pattern-match](https://github.com/pingcap/tidb/pull/23497)                                               |   | 43d22h  |
| tidb/23517   | [*: Add the metric about the SQL with TiFlash Success  (#23426)](https://github.com/pingcap/tidb/pull/23517)                                                  |   | 43d12h  |
| tidb/23562   | [execution: reuse iterator in hash join](https://github.com/pingcap/tidb/pull/23562)                                                                          |   | 42d13h  |
| tidb/23640   | [*: fix the bug about YEAR(0.9) returns NULL instead of 0 in NO_ZERO_DATE mode](https://github.com/pingcap/tidb/pull/23640)                                   |   | 38d13h  |
| tidb/23661   | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23661)       |   | 37d20h  |
| tidb/23682   | [executor: fix a panic when batch point get is used for partition table (#23652)](https://github.com/pingcap/tidb/pull/23682)                                 |   | 37d17h  |
| tidb/23691   | [executor: fix index join on prefix column index (#23678)](https://github.com/pingcap/tidb/pull/23691)                                                        |   | 37d16h  |
| tidb/23705   | [executor: refineArgs() bug fix when compare int with very small decimal (#23694)](https://github.com/pingcap/tidb/pull/23705)                                |   | 37d13h  |
| tidb/23756   | [planner: fix set not null flag for outer join (#23727)](https://github.com/pingcap/tidb/pull/23756)                                                          |   | 36d15h  |
| tidb/23812   | [executor, planner: fix collation for hash join building (#23770)](https://github.com/pingcap/tidb/pull/23812)                                                |   | 35d12h  |
| tidb/23884   | [Metric: Collect TiKV Read Metric for SLI/SLO](https://github.com/pingcap/tidb/pull/23884)                                                                    |   | 29d20h  |
| tidb/23964   | [executor: GROUP_CONCAT(float) is not compatible with mysql](https://github.com/pingcap/tidb/pull/23964)                                                      |   | 24d17h  |
| tidb/24007   | [ddl: refactor rule [4/6]](https://github.com/pingcap/tidb/pull/24007)                                                                                        |   | 22d20h  |
| tidb/24016   | [planner: fix index-out-of-range error when checking only_full_group_by (#23844)](https://github.com/pingcap/tidb/pull/24016)                                 |   | 22d19h  |
| tidb/24026   | [types: fix type merge about bit type (#23857)](https://github.com/pingcap/tidb/pull/24026)                                                                   |   | 22d14h  |
| tidb/24033   | [statistics: fix some unstable tests in global stats (#23502)](https://github.com/pingcap/tidb/pull/24033)                                                    |   | 22d9h   |
| tidb/24053   | [executor: fix wrong convert from bit to string when do projection (#23960)](https://github.com/pingcap/tidb/pull/24053)                                      |   | 21d16h  |
| tidb/24060   | [statistics: fix some potential panic in statistics (#23988)](https://github.com/pingcap/tidb/pull/24060)                                                     |   | 21d13h  |
| tidb/24061   | [statistics: fix some potential panic in statistics (#23988)](https://github.com/pingcap/tidb/pull/24061)                                                     |   | 21d13h  |
| tidb/24079   | [planner: change descScanFactor to scanFactor when ExpectedCount is small. (#23972)](https://github.com/pingcap/tidb/pull/24079)                              |   | 20d19h  |
| tidb/24155   | [planner, executor: fix index merge partial table scan schema (#23936)](https://github.com/pingcap/tidb/pull/24155)                                           |   | 16d20h  |
| tidb/24179   | [expression: fix float64 overflow check in plus/minus real function](https://github.com/pingcap/tidb/pull/24179)                                              |   | 15d23h  |
| tidb/24228   | [executor: skip TestPrepareStmtAfterIsolationReadChange when race enable (#24200)](https://github.com/pingcap/tidb/pull/24228)                                |   | 13d22h  |
| tidb/24229   | [executor: speed up race test TestInsertReorgDelete (#24208)](https://github.com/pingcap/tidb/pull/24229)                                                     |   | 13d21h  |
| tidb/24234   | [executor: skip TestMppExecution when race is enabled (#24222)](https://github.com/pingcap/tidb/pull/24234)                                                   |   | 13d18h  |
| tidb/24241   | [planner/core: remove random test to reduce CI time (#24207)](https://github.com/pingcap/tidb/pull/24241)                                                     |   | 13d15h  |
| tidb/24266   | [expression: fix wrong flen infer for bit constant (#23867)](https://github.com/pingcap/tidb/pull/24266)                                                      |   | 11d18h  |
| tidb/24267   | [expression: fix wrong flen infer for bit constant (#23867)](https://github.com/pingcap/tidb/pull/24267)                                                      |   | 11d18h  |
| tidb/24287   | [planner/core: support union all for mpp.](https://github.com/pingcap/tidb/pull/24287)                                                                        |   | 10d19h  |
| tidb/24340   | [executor: fix projection executor panic and add failpoint test (#24231)](https://github.com/pingcap/tidb/pull/24340)                                         |   | 8d20h   |
| tidb/24341   | [executor: fix projection executor panic and add failpoint test (#24231)](https://github.com/pingcap/tidb/pull/24341)                                         |   | 8d20h   |
| tidb/24345   | [executor: fix data race of parallel apply operator (#24257)](https://github.com/pingcap/tidb/pull/24345)                                                     |   | 8d19h   |
| tidb/24354   | [expression: fix wrong type infer for agg function when type is null (#24290)](https://github.com/pingcap/tidb/pull/24354)                                    |   | 8d16h   |
| tidb/24371   | [*: avoid create new parser object in prepared exec](https://github.com/pingcap/tidb/pull/24371)                                                              |   | 7d20h   |
| tidb/24425   | [test: fix unstable TestIssue20658](https://github.com/pingcap/tidb/pull/24425)                                                                               |   | 17h     |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>dyzsr's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                 PR                                                                  | C | LASTED |
|------------|-------------------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/24018 | [ranger: fix the range construction behavior when the column's type is `YEAR` (#23559)](https://github.com/pingcap/tidb/pull/24018) |   | 22d18h |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>eurekaka's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                         PR                                                                         | C | LASTED  |
|------------|----------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                                          |   | 182d17h |
| tidb/23002 | [store/*: fix err check](https://github.com/pingcap/tidb/pull/23002)                                                                               |   | 68d0h   |
| tidb/23283 | [util: optimize the performance of restore with db (#22910)](https://github.com/pingcap/tidb/pull/23283)                                           |   | 55d17h  |
| tidb/23316 | [planner: Fix rebuild range for prepared plan](https://github.com/pingcap/tidb/pull/23316)                                                         |   | 52d17h  |
| tidb/23373 | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                  |   | 50d19h  |
| tidb/23543 | [statistics: fix auto analyze log information incomplete (#23522)](https://github.com/pingcap/tidb/pull/23543)                                     |   | 42d18h  |
| tidb/23689 | [planner: fix the panic when we calculate the partition range (#23651)](https://github.com/pingcap/tidb/pull/23689)                                |   | 37d16h  |
| tidb/23705 | [executor: refineArgs() bug fix when compare int with very small decimal (#23694)](https://github.com/pingcap/tidb/pull/23705)                     |   | 37d13h  |
| tidb/23756 | [planner: fix set not null flag for outer join (#23727)](https://github.com/pingcap/tidb/pull/23756)                                               |   | 36d15h  |
| tidb/23760 | [collation: fix tidb panic when compare string with collation](https://github.com/pingcap/tidb/pull/23760)                                         |   | 36d13h  |
| tidb/23938 | [planner,privilege: requires extra privileges for REPLACE and INSERT ON DUPLICATE statements (#23911)](https://github.com/pingcap/tidb/pull/23938) |   | 27d10h  |
| tidb/24033 | [statistics: fix some unstable tests in global stats (#23502)](https://github.com/pingcap/tidb/pull/24033)                                         |   | 22d9h   |
| tidb/24061 | [statistics: fix some potential panic in statistics (#23988)](https://github.com/pingcap/tidb/pull/24061)                                          |   | 21d13h  |
| tidb/24079 | [planner: change descScanFactor to scanFactor when ExpectedCount is small. (#23972)](https://github.com/pingcap/tidb/pull/24079)                   |   | 20d19h  |
| tidb/24147 | [docs/design: add proposal for common table expression](https://github.com/pingcap/tidb/pull/24147)                                                |   | 16d23h  |
| tidb/24155 | [planner, executor: fix index merge partial table scan schema (#23936)](https://github.com/pingcap/tidb/pull/24155)                                |   | 16d20h  |
| tidb/24214 | [plan: merge continuous selections and delete surely true expressions](https://github.com/pingcap/tidb/pull/24214)                                 |   | 14d12h  |
| tidb/24236 | [*: remove SchemaVersion in TransactionContext](https://github.com/pingcap/tidb/pull/24236)                                                        |   | 13d17h  |
| tidb/24258 | [Revert "planner: donot prune all columns for Projection (#24024)" (#24180)](https://github.com/pingcap/tidb/pull/24258)                           |   | 11d22h  |
| tidb/24317 | [statistics: skip reading mysql.stats_histograms if cached stats is up-to-date (#24175)](https://github.com/pingcap/tidb/pull/24317)               |   | 9d17h   |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>fzhedu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                          PR                                                           | C | LASTED |
|------------|-----------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/24340 | [executor: fix projection executor panic and add failpoint test (#24231)](https://github.com/pingcap/tidb/pull/24340) |   | 8d20h  |
| tidb/24341 | [executor: fix projection executor panic and add failpoint test (#24231)](https://github.com/pingcap/tidb/pull/24341) |   | 8d20h  |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>ichn-hu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                             PR                                                             | C | LASTED |
|------------|----------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/24354 | [expression: fix wrong type infer for agg function when type is null (#24290)](https://github.com/pingcap/tidb/pull/24354) |   | 8d16h  |
| tidb/24379 | [executor: enhancement for ListInDisk(support writing after reading)](https://github.com/pingcap/tidb/pull/24379)          |   | 7d17h  |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>lzmhhh123's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                           PR                                                                            | C | LASTED  |
|------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20444 | [expression: add json_merge_patch](https://github.com/pingcap/tidb/pull/20444)                                                                          |   | 204d21h |
| tidb/20465 | [expression: add uuidShortFunction](https://github.com/pingcap/tidb/pull/20465)                                                                         |   | 203d19h |
| tidb/20642 | [executor: modify admin executors to support partitioned table with global index](https://github.com/pingcap/tidb/pull/20642)                           |   | 192d15h |
| tidb/20903 | [planner: fix confused and unnecessary double-projection in plans.](https://github.com/pingcap/tidb/pull/20903)                                         |   | 181d17h |
| tidb/21018 | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                          |   | 175d17h |
| tidb/21195 | [brie: integrate lightning to suport IMPORT statement](https://github.com/pingcap/tidb/pull/21195)                                                      |   | 164d23h |
| tidb/21334 | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                                           |   | 161d14h |
| tidb/21347 | [session: make rollback work on global variables](https://github.com/pingcap/tidb/pull/21347)                                                           |   | 160d20h |
| tidb/21487 | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                                           |   | 154d5h  |
| tidb/21651 | [planner: allow filter condition pushing down to IndexScan for prefix index](https://github.com/pingcap/tidb/pull/21651)                                |   | 147d14h |
| tidb/22126 | [*: add `sys` schema, `sys.SCHEMA_UNUSED_INDEXES` view and `sys.SCHEMA_INDEX_USAGE` view](https://github.com/pingcap/tidb/pull/22126)                   |   | 126d20h |
| tidb/22361 | [table: fix insert into _tidb_rowid panic and rebase it if needed (#22062)](https://github.com/pingcap/tidb/pull/22361)                                 |   | 114d20h |
| tidb/22372 | [executor: fix SelectForUpdate in decorrelated subquery under pessimistic mode](https://github.com/pingcap/tidb/pull/22372)                             |   | 114d9h  |
| tidb/22478 | [planner, executor: fix query partition table with global unique index get wrong result](https://github.com/pingcap/tidb/pull/22478)                    |   | 105d13h |
| tidb/22631 | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                                                                         |   | 95d23h  |
| tidb/22699 | [brie: add error info column and history backup/restore info in sql](https://github.com/pingcap/tidb/pull/22699)                                        |   | 92d16h  |
| tidb/23022 | [executor: create PipelinedWindowExec](https://github.com/pingcap/tidb/pull/23022)                                                                      |   | 66d18h  |
| tidb/23149 | [core: support left join and right join for join reorder](https://github.com/pingcap/tidb/pull/23149)                                                   |   | 61d12h  |
| tidb/23257 | [executor: group_concat aggr panic when session.group_concat_max_len is small (#23131)](https://github.com/pingcap/tidb/pull/23257)                     |   | 56d18h  |
| tidb/23283 | [util: optimize the performance of restore with db (#22910)](https://github.com/pingcap/tidb/pull/23283)                                                |   | 55d17h  |
| tidb/23347 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                                    |   | 51d18h  |
| tidb/23348 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                                    |   | 51d18h  |
| tidb/23368 | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23368)                                   |   | 50d20h  |
| tidb/23373 | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                       |   | 50d19h  |
| tidb/23655 | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral (#23523)](https://github.com/pingcap/tidb/pull/23655)                |   | 37d22h  |
| tidb/23660 | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23660) |   | 37d20h  |
| tidb/23661 | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23661) |   | 37d20h  |
| tidb/23703 | [expression: fix approx_percent panic on bit column (#23687)](https://github.com/pingcap/tidb/pull/23703)                                               |   | 37d14h  |
| tidb/23705 | [executor: refineArgs() bug fix when compare int with very small decimal (#23694)](https://github.com/pingcap/tidb/pull/23705)                          |   | 37d13h  |
| tidb/23756 | [planner: fix set not null flag for outer join (#23727)](https://github.com/pingcap/tidb/pull/23756)                                                    |   | 36d15h  |
| tidb/23760 | [collation: fix tidb panic when compare string with collation](https://github.com/pingcap/tidb/pull/23760)                                              |   | 36d13h  |
| tidb/23812 | [executor, planner: fix collation for hash join building (#23770)](https://github.com/pingcap/tidb/pull/23812)                                          |   | 35d12h  |
| tidb/23940 | [config, ddl: allow auto inc columns in generated columns and expression indexes](https://github.com/pingcap/tidb/pull/23940)                           |   | 26d18h  |
| tidb/23968 | [statistics: fix unstable TestDropPartitionStats test](https://github.com/pingcap/tidb/pull/23968)                                                      |   | 24d15h  |
| tidb/23987 | [executor: Implements json_arrayagg function](https://github.com/pingcap/tidb/pull/23987)                                                               |   | 23d18h  |
| tidb/24016 | [planner: fix index-out-of-range error when checking only_full_group_by (#23844)](https://github.com/pingcap/tidb/pull/24016)                           |   | 22d19h  |
| tidb/24018 | [ranger: fix the range construction behavior when the column's type is `YEAR` (#23559)](https://github.com/pingcap/tidb/pull/24018)                     |   | 22d18h  |
| tidb/24054 | [executor: fix wrong convert from bit to string when do projection (#23960)](https://github.com/pingcap/tidb/pull/24054)                                |   | 21d16h  |
| tidb/24151 | [ddl: admin show ddl jobs output confusing with multiple jobs](https://github.com/pingcap/tidb/pull/24151)                                              |   | 16d21h  |
| tidb/24155 | [planner, executor: fix index merge partial table scan schema (#23936)](https://github.com/pingcap/tidb/pull/24155)                                     |   | 16d20h  |
| tidb/24185 | [executor: make column default value being aware of NO_ZERO_IN_DATE (#24174)](https://github.com/pingcap/tidb/pull/24185)                               |   | 15d19h  |
| tidb/24186 | [executor: make column default value being aware of NO_ZERO_IN_DATE (#24174)](https://github.com/pingcap/tidb/pull/24186)                               |   | 15d19h  |
| tidb/24211 | [*: support txn retry when auto id meets duplicate entry](https://github.com/pingcap/tidb/pull/24211)                                                   |   | 14d13h  |
| tidb/24234 | [executor: skip TestMppExecution when race is enabled (#24222)](https://github.com/pingcap/tidb/pull/24234)                                             |   | 13d18h  |
| tidb/24250 | [planner: rewrite `LIKE` as range for expression index](https://github.com/pingcap/tidb/pull/24250)                                                     |   | 12d21h  |
| tidb/24258 | [Revert "planner: donot prune all columns for Projection (#24024)" (#24180)](https://github.com/pingcap/tidb/pull/24258)                                |   | 11d22h  |
| tidb/24268 | [expression: fix cast real, decimal to time (#24120)](https://github.com/pingcap/tidb/pull/24268)                                                       |   | 11d17h  |
| tidb/24285 | [*: compatibility with staleread](https://github.com/pingcap/tidb/pull/24285)                                                                           |   | 10d19h  |
| tidb/24340 | [executor: fix projection executor panic and add failpoint test (#24231)](https://github.com/pingcap/tidb/pull/24340)                                   |   | 8d20h   |
| tidb/24341 | [executor: fix projection executor panic and add failpoint test (#24231)](https://github.com/pingcap/tidb/pull/24341)                                   |   | 8d20h   |
| tidb/24357 | [statistics: fix a statistics GC problem that can cause duplicated fm-sketch records (#23830)](https://github.com/pingcap/tidb/pull/24357)              |   | 8d14h   |
| tidb/24389 | [executor, meta: Allocate auto id for global temporary tables](https://github.com/pingcap/tidb/pull/24389)                                              |   | 7d0h    |
| tidb/24416 | [planner, privilege: Add security enhanced mode part 4](https://github.com/pingcap/tidb/pull/24416)                                                     |   | 1d4h    |
| tidb/24423 | [executor, statistics: support prefix column index case for full sampling analyze](https://github.com/pingcap/tidb/pull/24423)                          |   | 18h     |


## Reviewed in Last 7 Days

|    REPO    |                                                     PR                                                      | C | D |   R   |
|------------|-------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/24437 | [planner: fix column pruning bug for Apply and Join (#24369)](https://github.com/pingcap/tidb/pull/24437)   |   | 1 | 12h   |
| tidb/24431 | [util: fix enum index range for in/not in clause.](https://github.com/pingcap/tidb/pull/24431)              |   | 1 | 0h    |
| tidb/24369 | [planner: fix column pruning bug for Apply and Join](https://github.com/pingcap/tidb/pull/24369)            |   | 1 | 7d5h  |
| tikv/10101 | [copr: support group by enum column](https://github.com/tikv/tikv/pull/10101)                               | Y | 1 | 6d21h |
| tidb/24342 | [planner: create new column slice in PreparePossibleProperties](https://github.com/pingcap/tidb/pull/24342) |   | 7 | 1d22h |


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
| docs-cn/5561 | [Add sql optimization-related docs to toc](https://github.com/pingcap/docs-cn/pull/5561)                                                                |   | 73d15h  |
| tidb/19029   | [types: fix unexpected NOT_NULL flags](https://github.com/pingcap/tidb/pull/19029)                                                                      |   | 273d22h |
| docs/5498    | [partitioning: Corrected partition management](https://github.com/pingcap/docs/pull/5498)                                                               |   | 10d19h  |
| tidb/20708   | [*: separate auto_increment ID allocator from _tidb_rowid allocator](https://github.com/pingcap/tidb/pull/20708)                                        |   | 189d20h |
| tidb/20969   | [executor: Improve the performance of appending not fixed columns](https://github.com/pingcap/tidb/pull/20969)                                          |   | 177d10h |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                          |   | 175d17h |
| tidb/21149   | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                           |   | 168d14h |
| tidb/21304   | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                                |   | 162d12h |
| tidb/21318   | [planner, expression: use the range of column types to simplify expressions](https://github.com/pingcap/tidb/pull/21318)                                |   | 161d19h |
| tidb/21401   | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                                      |   | 157d11h |
| tidb/21508   | [execution: fix dayofweek('0000-00-00') behavior](https://github.com/pingcap/tidb/pull/21508)                                                           |   | 153d10h |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                             |   | 139d19h |
| tidb/21887   | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                                         |   | 138d11h |
| tidb/21954   | [planner/cascades: add rule `PushSelDownApply`](https://github.com/pingcap/tidb/pull/21954)                                                             |   | 134d23h |
| tidb/22146   | [executor: forbid SFU on view](https://github.com/pingcap/tidb/pull/22146)                                                                              |   | 122d22h |
| tidb/22217   | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                                           |   | 120d18h |
| tidb/22234   | [executor, planner: ON DUPLICATE UPDATE can refer to un-project col (#14412)](https://github.com/pingcap/tidb/pull/22234)                               |   | 120d15h |
| tidb/22261   | [time: fix parse datetime won't truncate the reluctant string (#22232)](https://github.com/pingcap/tidb/pull/22261)                                     |   | 119d19h |
| tidb/22374   | [expression: separated arithmeticIntDivideSig](https://github.com/pingcap/tidb/pull/22374)                                                              |   | 114d0h  |
| tidb/22415   | [ddl: refactor bundle[2/2] [6/6]](https://github.com/pingcap/tidb/pull/22415)                                                                           |   | 110d17h |
| tidb/22416   | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                                                       | Y | 110d12h |
| tidb/22541   | [expression: Support builtin function SOUNDEX](https://github.com/pingcap/tidb/pull/22541)                                                              |   | 100d9h  |
| tidb/22565   | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)                                      | Y | 99d17h  |
| tidb/22814   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22814)                                     |   | 76d19h  |
| tidb/22862   | [brie: fix the problem that ddl restored by BR via SQL is not replicated to downstream](https://github.com/pingcap/tidb/pull/22862)                     |   | 73d23h  |
| tidb/22923   | [expression: correct constant propagation for collation (#22666)](https://github.com/pingcap/tidb/pull/22923)                                           |   | 71d15h  |
| tidb/22924   | [planner: fix wrong index merge selection (#22825)](https://github.com/pingcap/tidb/pull/22924)                                                         |   | 71d14h  |
| tidb/23002   | [store/*: fix err check](https://github.com/pingcap/tidb/pull/23002)                                                                                    |   | 68d0h   |
| tidb/23022   | [executor: create PipelinedWindowExec](https://github.com/pingcap/tidb/pull/23022)                                                                      |   | 66d18h  |
| tidb/23152   | [expression: fix wrong error info (#22760)](https://github.com/pingcap/tidb/pull/23152)                                                                 |   | 60d14h  |
| tidb/23196   | [types: fix the bug about the wrong query result for decimal type  (#22507)](https://github.com/pingcap/tidb/pull/23196)                                |   | 58d18h  |
| tidb/23295   | [util, types: don't let SPM be affected by charset (#23161)](https://github.com/pingcap/tidb/pull/23295)                                                |   | 55d11h  |
| tidb/23316   | [planner: Fix rebuild range for prepared plan](https://github.com/pingcap/tidb/pull/23316)                                                              |   | 52d17h  |
| tidb/23373   | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                       |   | 50d19h  |
| tidb/23397   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23397)                                                          |   | 49d18h  |
| tidb/23398   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23398)                                                          |   | 49d17h  |
| tidb/23590   | [planner, table: optimize the list partition pruner for range query](https://github.com/pingcap/tidb/pull/23590)                                        |   | 41d16h  |
| tidb/23598   | [types: fix collation for binary literal (#23591)](https://github.com/pingcap/tidb/pull/23598)                                                          |   | 41d13h  |
| tidb/23655   | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral (#23523)](https://github.com/pingcap/tidb/pull/23655)                |   | 37d22h  |
| tidb/23658   | [*: collect transaction write duration/throughput metrics for SLI/SLO (#23462)](https://github.com/pingcap/tidb/pull/23658)                             |   | 37d22h  |
| tidb/23660   | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23660) |   | 37d20h  |
| tidb/23661   | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23661) |   | 37d20h  |
| tidb/23682   | [executor: fix a panic when batch point get is used for partition table (#23652)](https://github.com/pingcap/tidb/pull/23682)                           |   | 37d17h  |
| tidb/23689   | [planner: fix the panic when we calculate the partition range (#23651)](https://github.com/pingcap/tidb/pull/23689)                                     |   | 37d16h  |
| tidb/23730   | [distsql/*: typo fix for `dispatches`](https://github.com/pingcap/tidb/pull/23730)                                                                      |   | 36d18h  |
| tidb/23796   | [tests: make TestIndexLookupMergeJoinHang and TestIssue18068 stable (#23741)](https://github.com/pingcap/tidb/pull/23796)                               |   | 35d20h  |
| tidb/23812   | [executor, planner: fix collation for hash join building (#23770)](https://github.com/pingcap/tidb/pull/23812)                                          |   | 35d12h  |
| tidb/23963   | [executor: checking chunk is full precedes filtering](https://github.com/pingcap/tidb/pull/23963)                                                       |   | 24d17h  |
| tidb/23987   | [executor: Implements json_arrayagg function](https://github.com/pingcap/tidb/pull/23987)                                                               |   | 23d18h  |
| tidb/24018   | [ranger: fix the range construction behavior when the column's type is `YEAR` (#23559)](https://github.com/pingcap/tidb/pull/24018)                     |   | 22d18h  |
| tidb/24193   | [executor: implement CTEStorage](https://github.com/pingcap/tidb/pull/24193)                                                                            |   | 15d10h  |
| tidb/24214   | [plan: merge continuous selections and delete surely true expressions](https://github.com/pingcap/tidb/pull/24214)                                      |   | 14d12h  |
| tidb/24229   | [executor: speed up race test TestInsertReorgDelete (#24208)](https://github.com/pingcap/tidb/pull/24229)                                               |   | 13d21h  |
| tidb/24235   | [expression: try to fix TestExprPushDownToFlash tests](https://github.com/pingcap/tidb/pull/24235)                                                      |   | 13d17h  |
| tidb/24241   | [planner/core: remove random test to reduce CI time (#24207)](https://github.com/pingcap/tidb/pull/24241)                                               |   | 13d15h  |
| tidb/24266   | [expression: fix wrong flen infer for bit constant (#23867)](https://github.com/pingcap/tidb/pull/24266)                                                |   | 11d18h  |
| tidb/24267   | [expression: fix wrong flen infer for bit constant (#23867)](https://github.com/pingcap/tidb/pull/24267)                                                |   | 11d18h  |
| tidb/24328   | [*: implement tidb_bound_staleness built-in function](https://github.com/pingcap/tidb/pull/24328)                                                       |   | 9d10h   |
| tidb/24354   | [expression: fix wrong type infer for agg function when type is null (#24290)](https://github.com/pingcap/tidb/pull/24354)                              |   | 8d16h   |
| tidb/24359   | [domain, session: Add new sysvarcache to replace global values cache](https://github.com/pingcap/tidb/pull/24359)                                       |   | 8d7h    |
| tidb/24373   | [planner: filter conflict read_from_storage hints (#24313)](https://github.com/pingcap/tidb/pull/24373)                                                 |   | 7d19h   |
| tidb/24374   | [planner: filter conflict read_from_storage hints (#24313)](https://github.com/pingcap/tidb/pull/24374)                                                 |   | 7d19h   |
| tidb/24379   | [executor: enhancement for ListInDisk(support writing after reading)](https://github.com/pingcap/tidb/pull/24379)                                       |   | 7d17h   |
| tidb/24382   | [statistics: trigger auto-analyze based on histogram row count](https://github.com/pingcap/tidb/pull/24382)                                             |   | 7d16h   |
| tidb/24425   | [test: fix unstable TestIssue20658](https://github.com/pingcap/tidb/pull/24425)                                                                         |   | 17h     |
| tidb/24432   | [store/copr: invalidate stale regions for Mpp query. (#24410)](https://github.com/pingcap/tidb/pull/24432)                                              |   | 16h     |
| tidb/24437   | [planner: fix column pruning bug for Apply and Join (#24369)](https://github.com/pingcap/tidb/pull/24437)                                               |   | 13h     |


## Reviewed in Last 7 Days

|    REPO    |                                                PR                                                 | C | D |   R   |
|------------|---------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/24376 | [planner: prune partitions that will never be used](https://github.com/pingcap/tidb/pull/24376)   |   | 1 | 6d23h |
| tidb/24410 | [store/copr: invalidate stale regions for Mpp query.](https://github.com/pingcap/tidb/pull/24410) |   | 1 | 2d22h |


</details> 


<details> 
  <summary>rebelice's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                     PR                                                                     | C | LASTED |
|------------|--------------------------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/23836 | [parser, core: Implement force_index hint in parser and TiDB](https://github.com/pingcap/tidb/pull/23836)                                  |   | 34d18h |
| tidb/24033 | [statistics: fix some unstable tests in global stats (#23502)](https://github.com/pingcap/tidb/pull/24033)                                 |   | 22d9h  |
| tidb/24306 | [util/ranger: fix func name typo](https://github.com/pingcap/tidb/pull/24306)                                                              |   | 9d22h  |
| tidb/24339 | [server,session: do not create stats collector in HTTP API to avoid memory leak](https://github.com/pingcap/tidb/pull/24339)               |   | 8d22h  |
| tidb/24357 | [statistics: fix a statistics GC problem that can cause duplicated fm-sketch records (#23830)](https://github.com/pingcap/tidb/pull/24357) |   | 8d14h  |
| tidb/24374 | [planner: filter conflict read_from_storage hints (#24313)](https://github.com/pingcap/tidb/pull/24374)                                    |   | 7d19h  |


## Reviewed in Last 7 Days

|    REPO    |                                                                       PR                                                                        | C | D |  R  |
|------------|-------------------------------------------------------------------------------------------------------------------------------------------------|---|---|-----|
| tidb/24422 | [planner: add some test cases about partition table dynamic-mode and plan-cache](https://github.com/pingcap/tidb/pull/24422)                    |   | 1 | 18h |
| tidb/24430 | [planner: add some test cases about partition-table dynamic mode with global-stats and SQL binding](https://github.com/pingcap/tidb/pull/24430) |   | 1 | 16h |


</details> 


<details> 
  <summary>time-and-fate's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                       PR                                                                        | C | LASTED  |
|------------|-------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                                       |   | 182d17h |
| tidb/22416 | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                                               | Y | 110d12h |
| tidb/24155 | [planner, executor: fix index merge partial table scan schema (#23936)](https://github.com/pingcap/tidb/pull/24155)                             |   | 16d20h  |
| tidb/24230 | [*: consitent get infoschema](https://github.com/pingcap/tidb/pull/24230)                                                                       |   | 13d21h  |
| tidb/24373 | [planner: filter conflict read_from_storage hints (#24313)](https://github.com/pingcap/tidb/pull/24373)                                         |   | 7d19h   |
| tidb/24374 | [planner: filter conflict read_from_storage hints (#24313)](https://github.com/pingcap/tidb/pull/24374)                                         |   | 7d19h   |
| tidb/24422 | [planner: add some test cases about partition table dynamic-mode and plan-cache](https://github.com/pingcap/tidb/pull/24422)                    |   | 19h     |
| tidb/24430 | [planner: add some test cases about partition-table dynamic mode with global-stats and SQL binding](https://github.com/pingcap/tidb/pull/24430) |   | 16h     |


## Reviewed in Last 7 Days

|    REPO    |                                                     PR                                                      | C | D |  R   |
|------------|-------------------------------------------------------------------------------------------------------------|---|---|------|
| tidb/24382 | [statistics: trigger auto-analyze based on histogram row count](https://github.com/pingcap/tidb/pull/24382) |   | 1 | 7d3h |
| tidb/24376 | [planner: prune partitions that will never be used](https://github.com/pingcap/tidb/pull/24376)             |   | 1 | 7d1h |


</details> 


<details> 
  <summary>winoros's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                              PR                                                                               | C | LASTED  |
|--------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| docs-cn/5916 | [sql-statements, information-schema: add `END_TIME` field for table `ANALYZE_STATUS`](https://github.com/pingcap/docs-cn/pull/5916)                           |   | 35d17h  |
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                                        | Y | 237d14h |
| docs-cn/6113 | [config: update the default value of `feedback-probability`](https://github.com/pingcap/docs-cn/pull/6113)                                                    |   | 14d22h  |
| tidb/20765   | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                                             |   | 185d17h |
| tidb/20877   | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                                                     |   | 182d17h |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                                |   | 175d17h |
| tidb/21207   | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                                               |   | 164d19h |
| tidb/21487   | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                                                 |   | 154d5h  |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                                   |   | 139d19h |
| tidb/21954   | [planner/cascades: add rule `PushSelDownApply`](https://github.com/pingcap/tidb/pull/21954)                                                                   |   | 134d23h |
| tidb/22181   | [planner, expression: fix error when using IN combined with subquery (#22080)](https://github.com/pingcap/tidb/pull/22181)                                    |   | 121d18h |
| tidb/22416   | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                                                             | Y | 110d12h |
| tidb/22504   | [*:Fix the fetchHotRegion bug that the count always zero](https://github.com/pingcap/tidb/pull/22504)                                                         |   | 102d20h |
| tidb/22565   | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)                                            | Y | 99d17h  |
| tidb/22923   | [expression: correct constant propagation for collation (#22666)](https://github.com/pingcap/tidb/pull/22923)                                                 |   | 71d15h  |
| tidb/23233   | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23233)                                                        |   | 57d18h  |
| tidb/23347   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                                          |   | 51d18h  |
| tidb/23348   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                                          |   | 51d18h  |
| tidb/23350   | [util/stringutil, util/ranger, planner: use hierarchical separators to simplify the parsing for info of EXPLAIN ](https://github.com/pingcap/tidb/pull/23350) |   | 51d17h  |
| tidb/23373   | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                             |   | 50d19h  |
| tidb/23474   | [planner: fix inappropriate null flag of null constants (#23457)](https://github.com/pingcap/tidb/pull/23474)                                                 |   | 44d18h  |
| tidb/23543   | [statistics: fix auto analyze log information incomplete (#23522)](https://github.com/pingcap/tidb/pull/23543)                                                |   | 42d18h  |
| tidb/23598   | [types: fix collation for binary literal (#23591)](https://github.com/pingcap/tidb/pull/23598)                                                                |   | 41d13h  |
| tidb/23655   | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral (#23523)](https://github.com/pingcap/tidb/pull/23655)                      |   | 37d22h  |
| tidb/23689   | [planner: fix the panic when we calculate the partition range (#23651)](https://github.com/pingcap/tidb/pull/23689)                                           |   | 37d16h  |
| tidb/23772   | [tablecodec: fix text type decode for old row format (#23751)](https://github.com/pingcap/tidb/pull/23772)                                                    |   | 36d11h  |
| tidb/23849   | [ddl: tidb panic while query hash partition table with is null condition](https://github.com/pingcap/tidb/pull/23849)                                         |   | 31d13h  |
| tidb/23917   | [planner: fix wrong TableDual plans caused by comparing Binary and Bytes incorrectly (#23860)](https://github.com/pingcap/tidb/pull/23917)                    |   | 28d23h  |
| tidb/24018   | [ranger: fix the range construction behavior when the column's type is `YEAR` (#23559)](https://github.com/pingcap/tidb/pull/24018)                           |   | 22d18h  |
| tidb/24060   | [statistics: fix some potential panic in statistics (#23988)](https://github.com/pingcap/tidb/pull/24060)                                                     |   | 21d13h  |
| tidb/24061   | [statistics: fix some potential panic in statistics (#23988)](https://github.com/pingcap/tidb/pull/24061)                                                     |   | 21d13h  |
| tidb/24079   | [planner: change descScanFactor to scanFactor when ExpectedCount is small. (#23972)](https://github.com/pingcap/tidb/pull/24079)                              |   | 20d19h  |
| tidb/24138   | [planner: Add Equivalence Rules to Transform BinaryOptSubquery to ExistsSubquery](https://github.com/pingcap/tidb/pull/24138)                                 |   | 17d12h  |
| tidb/24241   | [planner/core: remove random test to reduce CI time (#24207)](https://github.com/pingcap/tidb/pull/24241)                                                     |   | 13d15h  |
| tidb/24258   | [Revert "planner: donot prune all columns for Projection (#24024)" (#24180)](https://github.com/pingcap/tidb/pull/24258)                                      |   | 11d22h  |
| tidb/24342   | [planner: create new column slice in PreparePossibleProperties](https://github.com/pingcap/tidb/pull/24342)                                                   |   | 8d20h   |
| tidb/24382   | [statistics: trigger auto-analyze based on histogram row count](https://github.com/pingcap/tidb/pull/24382)                                                   |   | 7d16h   |
| tidb/24418   | [*: fix some typos](https://github.com/pingcap/tidb/pull/24418)                                                                                               |   | 23h     |
| tidb/24437   | [planner: fix column pruning bug for Apply and Join (#24369)](https://github.com/pingcap/tidb/pull/24437)                                                     |   | 13h     |


## Reviewed in Last 7 Days

|    REPO    |                                                         PR                                                         | C | D |   R    |
|------------|--------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/24214 | [plan: merge continuous selections and delete surely true expressions](https://github.com/pingcap/tidb/pull/24214) |   | 1 | 13d20h |
| tidb/24369 | [planner: fix column pruning bug for Apply and Join](https://github.com/pingcap/tidb/pull/24369)                   |   | 1 | 7d5h   |
| tidb/24431 | [util: fix enum index range for in/not in clause.](https://github.com/pingcap/tidb/pull/24431)                     |   | 1 | 0h     |


</details> 


<details> 
  <summary>wshwsh12's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                       PR                                                                        | C | LASTED  |
|------------|-------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19807 | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                                         |   | 244d11h |
| tidb/19957 | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                          | Y | 237d14h |
| tidb/21487 | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                                   |   | 154d5h  |
| tidb/21887 | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                                 |   | 138d11h |
| tidb/22378 | [executor: vectorize hash aggregate](https://github.com/pingcap/tidb/pull/22378)                                                                |   | 113d19h |
| tidb/23336 | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23336)                        |   | 51d19h  |
| tidb/23347 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                            |   | 51d18h  |
| tidb/23348 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                            |   | 51d18h  |
| tidb/23368 | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23368)                           |   | 50d20h  |
| tidb/23397 | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23397)                                                  |   | 49d18h  |
| tidb/23398 | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23398)                                                  |   | 49d17h  |
| tidb/23519 | [executor: check privilege before adding](https://github.com/pingcap/tidb/pull/23519)                                                           |   | 43d0h   |
| tidb/23760 | [collation: fix tidb panic when compare string with collation](https://github.com/pingcap/tidb/pull/23760)                                      |   | 36d13h  |
| tidb/23968 | [statistics: fix unstable TestDropPartitionStats test](https://github.com/pingcap/tidb/pull/23968)                                              |   | 24d15h  |
| tidb/23979 | [executor, statistics: fix unstable `TestAnalyzeIndexExtractTopN`](https://github.com/pingcap/tidb/pull/23979)                                  |   | 23d23h  |
| tidb/24018 | [ranger: fix the range construction behavior when the column's type is `YEAR` (#23559)](https://github.com/pingcap/tidb/pull/24018)             |   | 22d18h  |
| tidb/24033 | [statistics: fix some unstable tests in global stats (#23502)](https://github.com/pingcap/tidb/pull/24033)                                      |   | 22d9h   |
| tidb/24050 | [expression: fix get var panic when types not match](https://github.com/pingcap/tidb/pull/24050)                                                |   | 21d17h  |
| tidb/24053 | [executor: fix wrong convert from bit to string when do projection (#23960)](https://github.com/pingcap/tidb/pull/24053)                        |   | 21d16h  |
| tidb/24054 | [executor: fix wrong convert from bit to string when do projection (#23960)](https://github.com/pingcap/tidb/pull/24054)                        |   | 21d16h  |
| tidb/24147 | [docs/design: add proposal for common table expression](https://github.com/pingcap/tidb/pull/24147)                                             |   | 16d23h  |
| tidb/24186 | [executor: make column default value being aware of NO_ZERO_IN_DATE (#24174)](https://github.com/pingcap/tidb/pull/24186)                       |   | 15d19h  |
| tidb/24228 | [executor: skip TestPrepareStmtAfterIsolationReadChange when race enable (#24200)](https://github.com/pingcap/tidb/pull/24228)                  |   | 13d22h  |
| tidb/24229 | [executor: speed up race test TestInsertReorgDelete (#24208)](https://github.com/pingcap/tidb/pull/24229)                                       |   | 13d21h  |
| tidb/24230 | [*: consitent get infoschema](https://github.com/pingcap/tidb/pull/24230)                                                                       |   | 13d21h  |
| tidb/24236 | [*: remove SchemaVersion in TransactionContext](https://github.com/pingcap/tidb/pull/24236)                                                     |   | 13d17h  |
| tidb/24266 | [expression: fix wrong flen infer for bit constant (#23867)](https://github.com/pingcap/tidb/pull/24266)                                        |   | 11d18h  |
| tidb/24267 | [expression: fix wrong flen infer for bit constant (#23867)](https://github.com/pingcap/tidb/pull/24267)                                        |   | 11d18h  |
| tidb/24268 | [expression: fix cast real, decimal to time (#24120)](https://github.com/pingcap/tidb/pull/24268)                                               |   | 11d17h  |
| tidb/24280 | [executor, session, variable: Move deprecation, synonyms, scope validation to sysvar struct](https://github.com/pingcap/tidb/pull/24280)        |   | 11d0h   |
| tidb/24341 | [executor: fix projection executor panic and add failpoint test (#24231)](https://github.com/pingcap/tidb/pull/24341)                           |   | 8d20h   |
| tidb/24345 | [executor: fix data race of parallel apply operator (#24257)](https://github.com/pingcap/tidb/pull/24345)                                       |   | 8d19h   |
| tidb/24354 | [expression: fix wrong type infer for agg function when type is null (#24290)](https://github.com/pingcap/tidb/pull/24354)                      |   | 8d16h   |
| tidb/24379 | [executor: enhancement for ListInDisk(support writing after reading)](https://github.com/pingcap/tidb/pull/24379)                               |   | 7d17h   |
| tidb/24380 | [ executor: Pass the SQL digest down to pessimistic lock request](https://github.com/pingcap/tidb/pull/24380)                                   |   | 7d17h   |
| tidb/24412 | [*: Add security enhanced mode part 3](https://github.com/pingcap/tidb/pull/24412)                                                              |   | 2d7h    |
| tidb/24422 | [planner: add some test cases about partition table dynamic-mode and plan-cache](https://github.com/pingcap/tidb/pull/24422)                    |   | 19h     |
| tidb/24430 | [planner: add some test cases about partition-table dynamic mode with global-stats and SQL binding](https://github.com/pingcap/tidb/pull/24430) |   | 16h     |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 

