|   REVIEWER    |  O(C)   | 1 | 2 | 3 | 4 | 5 | 6 | 7 | T  |
|---------------|---------|---|---|---|---|---|---|---|----|
| Reminiscent   |  6 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| SunRunAway    | 12 ( 1) | 0 | 2 | 0 | 0 | 0 | 0 | 0 |  2 |
| XuHuaiyu      | 85 ( 2) | 1 | 1 | 1 | 2 | 1 | 0 | 1 |  7 |
| dyzsr         |  1 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 1 |  1 |
| eurekaka      | 21 ( 0) | 1 | 1 | 0 | 0 | 1 | 0 | 0 |  3 |
| fzhedu        |  2 ( 0) | 1 | 0 | 0 | 0 | 1 | 0 | 1 |  3 |
| ichn-hu       |  2 ( 0) | 0 | 1 | 5 | 1 | 2 | 0 | 2 | 11 |
| lzmhhh123     | 53 ( 0) | 2 | 4 | 0 | 1 | 2 | 0 | 2 | 11 |
| nullnotnil    |  0 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| qw4990        | 67 ( 2) | 3 | 1 | 2 | 2 | 0 | 0 | 0 |  8 |
| rebelice      |  7 ( 0) | 1 | 0 | 0 | 1 | 0 | 0 | 0 |  2 |
| time-and-fate |  8 ( 1) | 3 | 2 | 0 | 0 | 0 | 0 | 0 |  5 |
| winoros       | 42 ( 3) | 4 | 4 | 0 | 0 | 0 | 0 | 0 |  8 |
| wshwsh12      | 34 ( 1) | 1 | 0 | 0 | 3 | 0 | 0 | 4 |  8 |


<details> 
  <summary>Reminiscent's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                     PR                                                                     | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/21896 | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                                 |   | 129d19h |
| tidb/23474 | [planner: fix inappropriate null flag of null constants (#23457)](https://github.com/pingcap/tidb/pull/23474)                              |   | 37d18h  |
| tidb/23575 | [executor: fix update panic on join having statement (#23554)](https://github.com/pingcap/tidb/pull/23575)                                 |   | 34d21h  |
| tidb/23917 | [planner: fix wrong TableDual plans caused by comparing Binary and Bytes incorrectly (#23860)](https://github.com/pingcap/tidb/pull/23917) |   | 21d23h  |
| tidb/24016 | [planner: fix index-out-of-range error when checking only_full_group_by (#23844)](https://github.com/pingcap/tidb/pull/24016)              |   | 15d19h  |
| tidb/24357 | [statistics: fix a statistics GC problem that can cause duplicated fm-sketch records (#23830)](https://github.com/pingcap/tidb/pull/24357) |   | 1d14h   |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>SunRunAway's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                  PR                                                                   | C | LASTED  |
|------------|---------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19178 | [executor: Refactor probe channel](https://github.com/pingcap/tidb/pull/19178)                                                        |   | 259d17h |
| tidb/19807 | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                               |   | 237d11h |
| tidb/19900 | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                        | Y | 232d18h |
| tidb/20140 | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                    |   | 219d22h |
| tidb/20765 | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                     |   | 178d17h |
| tidb/21207 | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                       |   | 157d19h |
| tidb/21834 | [planner: enhanced index range calculation plan](https://github.com/pingcap/tidb/pull/21834)                                          |   | 134d19h |
| tidb/21876 | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)           |   | 132d19h |
| tidb/21878 | [planner: do not push down lock to pointGet/bacthPointGet when selection exists](https://github.com/pingcap/tidb/pull/21878)          |   | 132d18h |
| tidb/21956 | [planner/preprocessor: disallow into-outfile clause in some place](https://github.com/pingcap/tidb/pull/21956)                        |   | 127d23h |
| tidb/22217 | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                         |   | 113d18h |
| tidb/22379 | [[experiment] executor: allow aggregation to spill disk when running out of memory quota](https://github.com/pingcap/tidb/pull/22379) |   | 106d19h |


## Reviewed in Last 7 Days

|    REPO    |                                          PR                                           | C | D |   R    |
|------------|---------------------------------------------------------------------------------------|---|---|--------|
| tidb/20749 | [executor: support global kill (32 bits)](https://github.com/pingcap/tidb/pull/20749) |   | 2 | 178d1h |
| tidb/24279 | [*: add security enhanced mode part 2](https://github.com/pingcap/tidb/pull/24279)    |   | 2 | 2d17h  |


</details> 


<details> 
  <summary>XuHuaiyu's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                              PR                                                                               | C | LASTED  |
|--------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| docs-cn/5561 | [Add sql optimization-related docs to toc](https://github.com/pingcap/docs-cn/pull/5561)                                                                      |   | 66d15h  |
| docs-cn/5619 | [Update data-type-date-and-time.md](https://github.com/pingcap/docs-cn/pull/5619)                                                                             |   | 62d16h  |
| tidb/19900   | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                                                | Y | 232d18h |
| docs-cn/5671 | [tidb: Add time format description](https://github.com/pingcap/docs-cn/pull/5671)                                                                             |   | 56d11h  |
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                                        | Y | 230d14h |
| tidb/20140   | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                                            |   | 219d22h |
| tidb/20749   | [executor: support global kill (32 bits)](https://github.com/pingcap/tidb/pull/20749)                                                                         |   | 179d2h  |
| tidb/20790   | [collation: add pinyin collation for chinese charset support](https://github.com/pingcap/tidb/pull/20790)                                                     |   | 177d21h |
| tidb/20969   | [executor: Improve the performance of appending not fixed columns](https://github.com/pingcap/tidb/pull/20969)                                                |   | 170d10h |
| tidb/21064   | [planner, executor: fix cast not check error](https://github.com/pingcap/tidb/pull/21064)                                                                     |   | 165d9h  |
| tidb/21149   | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                                 |   | 161d14h |
| tidb/21304   | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                                      |   | 155d12h |
| tidb/21334   | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                                                 |   | 154d14h |
| tidb/21401   | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                                            |   | 150d11h |
| tidb/21536   | [executor: add slow-log file meta cache to avoid repeat read file meta information](https://github.com/pingcap/tidb/pull/21536)                               |   | 143d15h |
| tidb/21564   | [ddl: fix Incorrect behavior of NO_ZERO_DATE when altering table](https://github.com/pingcap/tidb/pull/21564)                                                 |   | 142d16h |
| tidb/21896   | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                                                    |   | 129d19h |
| tidb/22131   | [privilege: remove leading and trailing space when create user and role](https://github.com/pingcap/tidb/pull/22131)                                          |   | 119d19h |
| tidb/22163   | [expression: separated arithmeticMinusIntSig](https://github.com/pingcap/tidb/pull/22163)                                                                     |   | 115d13h |
| tidb/22186   | [executor: fix select into outfile with year type column has no data (#22175)](https://github.com/pingcap/tidb/pull/22186)                                    |   | 114d16h |
| tidb/22616   | [expression: from_unixtime accept 64-bit integers](https://github.com/pingcap/tidb/pull/22616)                                                                |   | 90d23h  |
| tidb/22617   | [metrics: fix wrong bucket name of coprocessor cache (#22454)](https://github.com/pingcap/tidb/pull/22617)                                                    |   | 90d23h  |
| tidb/22624   | [ planner: not pruning column used by union scan condition (#21640)](https://github.com/pingcap/tidb/pull/22624)                                              |   | 90d17h  |
| tidb/22631   | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                                                                               |   | 88d23h  |
| tidb/22696   | [expression: enable arithmetic Mod push down](https://github.com/pingcap/tidb/pull/22696)                                                                     |   | 85d17h  |
| tidb/22711   | [executor: Fix inline schema name](https://github.com/pingcap/tidb/pull/22711)                                                                                |   | 85d12h  |
| tidb/22722   | [planner, errno: make error code of ErrMixOfGroupFuncAndFields consistent with MySQL](https://github.com/pingcap/tidb/pull/22722)                             |   | 84d21h  |
| tidb/22814   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22814)                                           |   | 69d19h  |
| tidb/22908   | [txn: Add txn state's view](https://github.com/pingcap/tidb/pull/22908)                                                                                       |   | 64d21h  |
| tidb/23012   | [executor: fix affected rows of ddls and complete uint tests](https://github.com/pingcap/tidb/pull/23012)                                                     |   | 60d17h  |
| tidb/23152   | [expression: fix wrong error info (#22760)](https://github.com/pingcap/tidb/pull/23152)                                                                       |   | 53d14h  |
| tidb/23196   | [types: fix the bug about the wrong query result for decimal type  (#22507)](https://github.com/pingcap/tidb/pull/23196)                                      |   | 51d18h  |
| tidb/23220   | [Release 4.0](https://github.com/pingcap/tidb/pull/23220)                                                                                                     |   | 51d11h  |
| tidb/23233   | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23233)                                                        |   | 50d18h  |
| tidb/23257   | [executor: group_concat aggr panic when session.group_concat_max_len is small (#23131)](https://github.com/pingcap/tidb/pull/23257)                           |   | 49d18h  |
| tidb/23295   | [util, types: don't let SPM be affected by charset (#23161)](https://github.com/pingcap/tidb/pull/23295)                                                      |   | 48d11h  |
| tidb/23335   | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23335)                                      |   | 44d19h  |
| tidb/23336   | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23336)                                      |   | 44d19h  |
| tidb/23347   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                                          |   | 44d18h  |
| tidb/23348   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                                          |   | 44d18h  |
| tidb/23350   | [util/stringutil, util/ranger, planner: use hierarchical separators to simplify the parsing for info of EXPLAIN ](https://github.com/pingcap/tidb/pull/23350) |   | 44d17h  |
| tidb/23368   | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23368)                                         |   | 43d20h  |
| tidb/23397   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23397)                                                                |   | 42d17h  |
| tidb/23398   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23398)                                                                |   | 42d17h  |
| tidb/23405   | [domain: remove the exit chan, use context](https://github.com/pingcap/tidb/pull/23405)                                                                       |   | 42d17h  |
| tidb/23433   | [WIP: speed up for slow query logs retrieving ](https://github.com/pingcap/tidb/pull/23433)                                                                   |   | 41d17h  |
| tidb/23474   | [planner: fix inappropriate null flag of null constants (#23457)](https://github.com/pingcap/tidb/pull/23474)                                                 |   | 37d18h  |
| tidb/23497   | [expression: Let TiDB use Hyperscan to support multi-pattern-match](https://github.com/pingcap/tidb/pull/23497)                                               |   | 36d22h  |
| tidb/23517   | [*: Add the metric about the SQL with TiFlash Success  (#23426)](https://github.com/pingcap/tidb/pull/23517)                                                  |   | 36d12h  |
| tidb/23562   | [execution: reuse iterator in hash join](https://github.com/pingcap/tidb/pull/23562)                                                                          |   | 35d13h  |
| tidb/23640   | [*: fix the bug about YEAR(0.9) returns NULL instead of 0 in NO_ZERO_DATE mode](https://github.com/pingcap/tidb/pull/23640)                                   |   | 31d13h  |
| tidb/23661   | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23661)       |   | 30d20h  |
| tidb/23682   | [executor: fix a panic when batch point get is used for partition table (#23652)](https://github.com/pingcap/tidb/pull/23682)                                 |   | 30d17h  |
| tidb/23691   | [executor: fix index join on prefix column index (#23678)](https://github.com/pingcap/tidb/pull/23691)                                                        |   | 30d16h  |
| tidb/23705   | [executor: refineArgs() bug fix when compare int with very small decimal (#23694)](https://github.com/pingcap/tidb/pull/23705)                                |   | 30d13h  |
| tidb/23756   | [planner: fix set not null flag for outer join (#23727)](https://github.com/pingcap/tidb/pull/23756)                                                          |   | 29d15h  |
| tidb/23812   | [executor, planner: fix collation for hash join building (#23770)](https://github.com/pingcap/tidb/pull/23812)                                                |   | 28d12h  |
| tidb/23878   | [functions: fix some string function has wrong collation and flag (#23835)](https://github.com/pingcap/tidb/pull/23878)                                       |   | 22d21h  |
| tidb/23884   | [Metric: Collect TiKV Read Metric for SLI/SLO](https://github.com/pingcap/tidb/pull/23884)                                                                    |   | 22d20h  |
| tidb/23888   | [executor: fix resource leak of Shuffle Executor.](https://github.com/pingcap/tidb/pull/23888)                                                                |   | 22d19h  |
| tidb/23958   | [executor: fix `show table status` for the database with upper-cased name (#23896)](https://github.com/pingcap/tidb/pull/23958)                               |   | 17d18h  |
| tidb/23964   | [executor: GROUP_CONCAT(float) is not compatible with mysql](https://github.com/pingcap/tidb/pull/23964)                                                      |   | 17d17h  |
| tidb/24007   | [ddl: refactor rule [4/6]](https://github.com/pingcap/tidb/pull/24007)                                                                                        |   | 15d20h  |
| tidb/24016   | [planner: fix index-out-of-range error when checking only_full_group_by (#23844)](https://github.com/pingcap/tidb/pull/24016)                                 |   | 15d19h  |
| tidb/24026   | [types: fix type merge about bit type (#23857)](https://github.com/pingcap/tidb/pull/24026)                                                                   |   | 15d14h  |
| tidb/24033   | [statistics: fix some unstable tests in global stats (#23502)](https://github.com/pingcap/tidb/pull/24033)                                                    |   | 15d9h   |
| tidb/24053   | [executor: fix wrong convert from bit to string when do projection (#23960)](https://github.com/pingcap/tidb/pull/24053)                                      |   | 14d16h  |
| tidb/24060   | [statistics: fix some potential panic in statistics (#23988)](https://github.com/pingcap/tidb/pull/24060)                                                     |   | 14d13h  |
| tidb/24061   | [statistics: fix some potential panic in statistics (#23988)](https://github.com/pingcap/tidb/pull/24061)                                                     |   | 14d13h  |
| tidb/24078   | [planner: change descScanFactor to scanFactor when ExpectedCount is small. (#23972)](https://github.com/pingcap/tidb/pull/24078)                              |   | 13d19h  |
| tidb/24079   | [planner: change descScanFactor to scanFactor when ExpectedCount is small. (#23972)](https://github.com/pingcap/tidb/pull/24079)                              |   | 13d19h  |
| tidb/24155   | [planner, executor: fix index merge partial table scan schema (#23936)](https://github.com/pingcap/tidb/pull/24155)                                           |   | 9d20h   |
| tidb/24179   | [expression: fix float64 overflow check in plus/minus real function](https://github.com/pingcap/tidb/pull/24179)                                              |   | 8d23h   |
| tidb/24228   | [executor: skip TestPrepareStmtAfterIsolationReadChange when race enable (#24200)](https://github.com/pingcap/tidb/pull/24228)                                |   | 6d22h   |
| tidb/24229   | [executor: speed up race test TestInsertReorgDelete (#24208)](https://github.com/pingcap/tidb/pull/24229)                                                     |   | 6d21h   |
| tidb/24234   | [executor: skip TestMppExecution when race is enabled (#24222)](https://github.com/pingcap/tidb/pull/24234)                                                   |   | 6d18h   |
| tidb/24241   | [planner/core: remove random test to reduce CI time (#24207)](https://github.com/pingcap/tidb/pull/24241)                                                     |   | 6d15h   |
| tidb/24266   | [expression: fix wrong flen infer for bit constant (#23867)](https://github.com/pingcap/tidb/pull/24266)                                                      |   | 4d18h   |
| tidb/24267   | [expression: fix wrong flen infer for bit constant (#23867)](https://github.com/pingcap/tidb/pull/24267)                                                      |   | 4d18h   |
| tidb/24287   | [planner/core: support union all for mpp.](https://github.com/pingcap/tidb/pull/24287)                                                                        |   | 3d19h   |
| tidb/24340   | [executor: fix projection executor panic and add failpoint test (#24231)](https://github.com/pingcap/tidb/pull/24340)                                         |   | 1d20h   |
| tidb/24341   | [executor: fix projection executor panic and add failpoint test (#24231)](https://github.com/pingcap/tidb/pull/24341)                                         |   | 1d20h   |
| tidb/24345   | [executor: fix data race of parallel apply operator (#24257)](https://github.com/pingcap/tidb/pull/24345)                                                     |   | 1d19h   |
| tidb/24354   | [expression: fix wrong type infer for agg function when type is null (#24290)](https://github.com/pingcap/tidb/pull/24354)                                    |   | 1d16h   |
| tidb/24371   | [*: avoid create new parser object in prepared exec](https://github.com/pingcap/tidb/pull/24371)                                                              |   | 19h     |


## Reviewed in Last 7 Days

|    REPO    |                                                                   PR                                                                   | C | D |    R    |
|------------|----------------------------------------------------------------------------------------------------------------------------------------|---|---|---------|
| tidb/21228 | [executor: return the result immediately when combining LIMIT row_count with DISTINCT](https://github.com/pingcap/tidb/pull/21228)     |   | 1 | 156d15h |
| tidb/24278 | [executor: accelerate TestVectorizedMergeJoin and TestVectorizedShuffleMergeJoin (#24177)](https://github.com/pingcap/tidb/pull/24278) |   | 2 | 2d10h   |
| tidb/24308 | [docs/design: update the proposal of `column type change`](https://github.com/pingcap/tidb/pull/24308)                                 |   | 3 | 0h      |
| tidb/24179 | [expression: fix float64 overflow check in plus/minus real function](https://github.com/pingcap/tidb/pull/24179)                       |   | 4 | 5d22h   |
| tidb/24290 | [expression: fix wrong type infer for agg function when type is null](https://github.com/pingcap/tidb/pull/24290)                      |   | 4 | 18h     |
| tidb/24257 | [executor: fix data race of parallel apply operator](https://github.com/pingcap/tidb/pull/24257)                                       |   | 5 | 0h      |
| docs/5445  | [tidb-config: add value range and type for token-limit](https://github.com/pingcap/docs/pull/5445)                                     |   | 7 | 0h      |


</details> 


<details> 
  <summary>dyzsr's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                 PR                                                                  | C | LASTED |
|------------|-------------------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/24018 | [ranger: fix the range construction behavior when the column's type is `YEAR` (#23559)](https://github.com/pingcap/tidb/pull/24018) |   | 15d18h |


## Reviewed in Last 7 Days

|    REPO    |                                      PR                                       | C | D | R  |
|------------|-------------------------------------------------------------------------------|---|---|----|
| tikv/10074 | [copr: fix unsound unsafe transmute](https://github.com/tikv/tikv/pull/10074) |   | 7 | 2h |


</details> 


<details> 
  <summary>eurekaka's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                         PR                                                                         | C | LASTED  |
|------------|----------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                                          |   | 175d17h |
| tidb/23002 | [store/*: fix err check](https://github.com/pingcap/tidb/pull/23002)                                                                               |   | 61d0h   |
| tidb/23283 | [util: optimize the performance of restore with db (#22910)](https://github.com/pingcap/tidb/pull/23283)                                           |   | 48d17h  |
| tidb/23316 | [planner: Fix rebuild range for prepared plan](https://github.com/pingcap/tidb/pull/23316)                                                         |   | 45d17h  |
| tidb/23373 | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                  |   | 43d19h  |
| tidb/23543 | [statistics: fix auto analyze log information incomplete (#23522)](https://github.com/pingcap/tidb/pull/23543)                                     |   | 35d18h  |
| tidb/23689 | [planner: fix the panic when we calculate the partition range (#23651)](https://github.com/pingcap/tidb/pull/23689)                                |   | 30d16h  |
| tidb/23705 | [executor: refineArgs() bug fix when compare int with very small decimal (#23694)](https://github.com/pingcap/tidb/pull/23705)                     |   | 30d13h  |
| tidb/23756 | [planner: fix set not null flag for outer join (#23727)](https://github.com/pingcap/tidb/pull/23756)                                               |   | 29d15h  |
| tidb/23760 | [collation: fix tidb panic when compare string with collation](https://github.com/pingcap/tidb/pull/23760)                                         |   | 29d13h  |
| tidb/23938 | [planner,privilege: requires extra privileges for REPLACE and INSERT ON DUPLICATE statements (#23911)](https://github.com/pingcap/tidb/pull/23938) |   | 20d10h  |
| tidb/23974 | [planner: do not push down to TiFlash if the table scan require to scan data in desc order (#23948)](https://github.com/pingcap/tidb/pull/23974)   |   | 17d12h  |
| tidb/24033 | [statistics: fix some unstable tests in global stats (#23502)](https://github.com/pingcap/tidb/pull/24033)                                         |   | 15d9h   |
| tidb/24061 | [statistics: fix some potential panic in statistics (#23988)](https://github.com/pingcap/tidb/pull/24061)                                          |   | 14d13h  |
| tidb/24079 | [planner: change descScanFactor to scanFactor when ExpectedCount is small. (#23972)](https://github.com/pingcap/tidb/pull/24079)                   |   | 13d19h  |
| tidb/24147 | [docs/design: add proposal for common table expression](https://github.com/pingcap/tidb/pull/24147)                                                |   | 9d23h   |
| tidb/24155 | [planner, executor: fix index merge partial table scan schema (#23936)](https://github.com/pingcap/tidb/pull/24155)                                |   | 9d20h   |
| tidb/24214 | [plan: merge continuous selections and delete surely true expressions](https://github.com/pingcap/tidb/pull/24214)                                 |   | 7d12h   |
| tidb/24236 | [*: remove SchemaVersion in TransactionContext](https://github.com/pingcap/tidb/pull/24236)                                                        |   | 6d17h   |
| tidb/24258 | [Revert "planner: donot prune all columns for Projection (#24024)" (#24180)](https://github.com/pingcap/tidb/pull/24258)                           |   | 4d22h   |
| tidb/24317 | [statistics: skip reading mysql.stats_histograms if cached stats is up-to-date (#24175)](https://github.com/pingcap/tidb/pull/24317)               |   | 2d17h   |


## Reviewed in Last 7 Days

|    REPO    |                                                              PR                                                              | C | D |   R    |
|------------|------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/23997 | [stats, executor: use a correct sampling to collect stats](https://github.com/pingcap/tidb/pull/23997)                       |   | 1 | 15d16h |
| tidb/24339 | [server,session: do not create stats collector in HTTP API to avoid memory leak](https://github.com/pingcap/tidb/pull/24339) |   | 2 | 4h     |
| tidb/24180 | [Revert "planner: donot prune all columns for Projection (#24024)"](https://github.com/pingcap/tidb/pull/24180)              |   | 5 | 3d23h  |


</details> 


<details> 
  <summary>fzhedu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                          PR                                                           | C | LASTED |
|------------|-----------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/24340 | [executor: fix projection executor panic and add failpoint test (#24231)](https://github.com/pingcap/tidb/pull/24340) |   | 1d20h  |
| tidb/24341 | [executor: fix projection executor panic and add failpoint test (#24231)](https://github.com/pingcap/tidb/pull/24341) |   | 1d20h  |


## Reviewed in Last 7 Days

|    REPO     |                                                           PR                                                           | C | D |   R    |
|-------------|------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| kvproto/751 | [mpp: support returning regions that need retry](https://github.com/pingcap/kvproto/pull/751)                          |   | 1 | 16d22h |
| tidb/24231  | [executor: fix projection executor panic and add failpoint test](https://github.com/pingcap/tidb/pull/24231)           |   | 5 | 2d1h   |
| tidb/24157  | [planner: let CopTiFlashConcurrencyFactor inflence the cost of whole plan](https://github.com/pingcap/tidb/pull/24157) |   | 7 | 3d3h   |


</details> 


<details> 
  <summary>ichn-hu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                             PR                                                             | C | LASTED |
|------------|----------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/24354 | [expression: fix wrong type infer for agg function when type is null (#24290)](https://github.com/pingcap/tidb/pull/24354) |   | 1d16h  |
| tidb/24379 | [executor: enhancement for ListInDisk(support writing after reading)](https://github.com/pingcap/tidb/pull/24379)          |   | 17h    |


## Reviewed in Last 7 Days

|    REPO    |                                                                   PR                                                                   | C | D |   R   |
|------------|----------------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/24290 | [expression: fix wrong type infer for agg function when type is null](https://github.com/pingcap/tidb/pull/24290)                      |   | 2 | 2d0h  |
| tidb/24266 | [expression: fix wrong flen infer for bit constant (#23867)](https://github.com/pingcap/tidb/pull/24266)                               |   | 3 | 2d2h  |
| tidb/24267 | [expression: fix wrong flen infer for bit constant (#23867)](https://github.com/pingcap/tidb/pull/24267)                               |   | 3 | 2d2h  |
| tidb/24300 | [expression: fix wrong collation for `concat` function (#24297)](https://github.com/pingcap/tidb/pull/24300)                           |   | 3 | 20h   |
| tidb/24304 | [executor, variable: move hidden variables to struct](https://github.com/pingcap/tidb/pull/24304)                                      |   | 3 | 12h   |
| tidb/24248 | [executor, variable: small cleanup](https://github.com/pingcap/tidb/pull/24248)                                                        |   | 3 | 3d14h |
| tidb/24278 | [executor: accelerate TestVectorizedMergeJoin and TestVectorizedShuffleMergeJoin (#24177)](https://github.com/pingcap/tidb/pull/24278) |   | 4 | 14h   |
| tidb/24177 | [executor: accelerate TestVectorizedMergeJoin and TestVectorizedShuffleMergeJoin](https://github.com/pingcap/tidb/pull/24177)          |   | 5 | 4d15h |
| tidb/24234 | [executor: skip TestMppExecution when race is enabled (#24222)](https://github.com/pingcap/tidb/pull/24234)                            |   | 5 | 1d19h |
| tidb/24235 | [expression: try to fix TestExprPushDownToFlash tests](https://github.com/pingcap/tidb/pull/24235)                                     |   | 7 | 0h    |
| tidb/24026 | [types: fix type merge about bit type (#23857)](https://github.com/pingcap/tidb/pull/24026)                                            |   | 7 | 8d19h |


</details> 


<details> 
  <summary>lzmhhh123's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                           PR                                                                            | C | LASTED  |
|------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20444 | [expression: add json_merge_patch](https://github.com/pingcap/tidb/pull/20444)                                                                          |   | 197d21h |
| tidb/20465 | [expression: add uuidShortFunction](https://github.com/pingcap/tidb/pull/20465)                                                                         |   | 196d19h |
| tidb/20642 | [executor: modify admin executors to support partitioned table with global index](https://github.com/pingcap/tidb/pull/20642)                           |   | 185d15h |
| tidb/20903 | [planner: fix confused and unnecessary double-projection in plans.](https://github.com/pingcap/tidb/pull/20903)                                         |   | 174d17h |
| tidb/21018 | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                          |   | 168d17h |
| tidb/21195 | [brie: integrate lightning to suport IMPORT statement](https://github.com/pingcap/tidb/pull/21195)                                                      |   | 157d23h |
| tidb/21334 | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                                           |   | 154d14h |
| tidb/21347 | [session: make rollback work on global variables](https://github.com/pingcap/tidb/pull/21347)                                                           |   | 153d20h |
| tidb/21487 | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                                           |   | 147d5h  |
| tidb/21651 | [planner: allow filter condition pushing down to IndexScan for prefix index](https://github.com/pingcap/tidb/pull/21651)                                |   | 140d14h |
| tidb/22126 | [*: add `sys` schema, `sys.SCHEMA_UNUSED_INDEXES` view and `sys.SCHEMA_INDEX_USAGE` view](https://github.com/pingcap/tidb/pull/22126)                   |   | 119d20h |
| tidb/22361 | [table: fix insert into _tidb_rowid panic and rebase it if needed (#22062)](https://github.com/pingcap/tidb/pull/22361)                                 |   | 107d20h |
| tidb/22372 | [executor: fix SelectForUpdate in decorrelated subquery under pessimistic mode](https://github.com/pingcap/tidb/pull/22372)                             |   | 107d9h  |
| tidb/22478 | [planner, executor: fix query partition table with global unique index get wrong result](https://github.com/pingcap/tidb/pull/22478)                    |   | 98d13h  |
| tidb/22631 | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                                                                         |   | 88d23h  |
| tidb/22699 | [brie: add error info column and history backup/restore info in sql](https://github.com/pingcap/tidb/pull/22699)                                        |   | 85d16h  |
| tidb/23022 | [executor: create PipelinedWindowExec](https://github.com/pingcap/tidb/pull/23022)                                                                      |   | 59d18h  |
| tidb/23149 | [core: support left join and right join for join reorder](https://github.com/pingcap/tidb/pull/23149)                                                   |   | 54d12h  |
| tidb/23257 | [executor: group_concat aggr panic when session.group_concat_max_len is small (#23131)](https://github.com/pingcap/tidb/pull/23257)                     |   | 49d18h  |
| tidb/23283 | [util: optimize the performance of restore with db (#22910)](https://github.com/pingcap/tidb/pull/23283)                                                |   | 48d17h  |
| tidb/23347 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                                    |   | 44d18h  |
| tidb/23348 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                                    |   | 44d18h  |
| tidb/23368 | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23368)                                   |   | 43d20h  |
| tidb/23373 | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                       |   | 43d19h  |
| tidb/23655 | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral (#23523)](https://github.com/pingcap/tidb/pull/23655)                |   | 30d22h  |
| tidb/23660 | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23660) |   | 30d20h  |
| tidb/23661 | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23661) |   | 30d20h  |
| tidb/23703 | [expression: fix approx_percent panic on bit column (#23687)](https://github.com/pingcap/tidb/pull/23703)                                               |   | 30d14h  |
| tidb/23705 | [executor: refineArgs() bug fix when compare int with very small decimal (#23694)](https://github.com/pingcap/tidb/pull/23705)                          |   | 30d13h  |
| tidb/23756 | [planner: fix set not null flag for outer join (#23727)](https://github.com/pingcap/tidb/pull/23756)                                                    |   | 29d15h  |
| tidb/23760 | [collation: fix tidb panic when compare string with collation](https://github.com/pingcap/tidb/pull/23760)                                              |   | 29d13h  |
| tidb/23812 | [executor, planner: fix collation for hash join building (#23770)](https://github.com/pingcap/tidb/pull/23812)                                          |   | 28d12h  |
| tidb/23940 | [config, ddl: allow auto inc columns in generated columns and expression indexes](https://github.com/pingcap/tidb/pull/23940)                           |   | 19d18h  |
| tidb/23968 | [statistics: fix unstable TestDropPartitionStats test](https://github.com/pingcap/tidb/pull/23968)                                                      |   | 17d15h  |
| tidb/23987 | [executor: Implements json_arrayagg function](https://github.com/pingcap/tidb/pull/23987)                                                               |   | 16d18h  |
| tidb/24016 | [planner: fix index-out-of-range error when checking only_full_group_by (#23844)](https://github.com/pingcap/tidb/pull/24016)                           |   | 15d19h  |
| tidb/24018 | [ranger: fix the range construction behavior when the column's type is `YEAR` (#23559)](https://github.com/pingcap/tidb/pull/24018)                     |   | 15d18h  |
| tidb/24054 | [executor: fix wrong convert from bit to string when do projection (#23960)](https://github.com/pingcap/tidb/pull/24054)                                |   | 14d16h  |
| tidb/24151 | [ddl: admin show ddl jobs output confusing with multiple jobs](https://github.com/pingcap/tidb/pull/24151)                                              |   | 9d21h   |
| tidb/24155 | [planner, executor: fix index merge partial table scan schema (#23936)](https://github.com/pingcap/tidb/pull/24155)                                     |   | 9d20h   |
| tidb/24185 | [executor: make column default value being aware of NO_ZERO_IN_DATE (#24174)](https://github.com/pingcap/tidb/pull/24185)                               |   | 8d19h   |
| tidb/24186 | [executor: make column default value being aware of NO_ZERO_IN_DATE (#24174)](https://github.com/pingcap/tidb/pull/24186)                               |   | 8d19h   |
| tidb/24211 | [*: support txn retry when auto id meets duplicate entry](https://github.com/pingcap/tidb/pull/24211)                                                   |   | 7d13h   |
| tidb/24234 | [executor: skip TestMppExecution when race is enabled (#24222)](https://github.com/pingcap/tidb/pull/24234)                                             |   | 6d18h   |
| tidb/24250 | [planner: rewrite `LIKE` as range for expression index](https://github.com/pingcap/tidb/pull/24250)                                                     |   | 5d21h   |
| tidb/24258 | [Revert "planner: donot prune all columns for Projection (#24024)" (#24180)](https://github.com/pingcap/tidb/pull/24258)                                |   | 4d22h   |
| tidb/24268 | [expression: fix cast real, decimal to time (#24120)](https://github.com/pingcap/tidb/pull/24268)                                                       |   | 4d17h   |
| tidb/24285 | [*: compatibility with staleread](https://github.com/pingcap/tidb/pull/24285)                                                                           |   | 3d19h   |
| tidb/24340 | [executor: fix projection executor panic and add failpoint test (#24231)](https://github.com/pingcap/tidb/pull/24340)                                   |   | 1d20h   |
| tidb/24341 | [executor: fix projection executor panic and add failpoint test (#24231)](https://github.com/pingcap/tidb/pull/24341)                                   |   | 1d20h   |
| tidb/24357 | [statistics: fix a statistics GC problem that can cause duplicated fm-sketch records (#23830)](https://github.com/pingcap/tidb/pull/24357)              |   | 1d14h   |
| tidb/24365 | [store/tikv: remove use of TxnScope transaction option in store/tikv](https://github.com/pingcap/tidb/pull/24365)                                       |   | 1d5h    |
| tidb/24389 | [executor, meta: Allocate auto id for global temporary tables](https://github.com/pingcap/tidb/pull/24389)                                              |   | 0h      |


## Reviewed in Last 7 Days

|    REPO    |                                                                   PR                                                                   | C | D |   R    |
|------------|----------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/23963 | [executor: checking chunk is full precedes filtering](https://github.com/pingcap/tidb/pull/23963)                                      |   | 1 | 17d5h  |
| tidb/24369 | [planner: fix column pruning bug for Apply and Join](https://github.com/pingcap/tidb/pull/24369)                                       |   | 1 | 1h     |
| tidb/24345 | [executor: fix data race of parallel apply operator (#24257)](https://github.com/pingcap/tidb/pull/24345)                              |   | 2 | 2h     |
| tidb/22691 | [planner, expression: support enum index scan](https://github.com/pingcap/tidb/pull/22691)                                             |   | 2 | 83d23h |
| tidb/24257 | [executor: fix data race of parallel apply operator](https://github.com/pingcap/tidb/pull/24257)                                       |   | 2 | 3d0h   |
| tikv/10057 | [statistics: introduce full sampling collect tech](https://github.com/tikv/tikv/pull/10057)                                            | Y | 2 | 6d14h  |
| tidb/24278 | [executor: accelerate TestVectorizedMergeJoin and TestVectorizedShuffleMergeJoin (#24177)](https://github.com/pingcap/tidb/pull/24278) |   | 4 | 15h    |
| tidb/24179 | [expression: fix float64 overflow check in plus/minus real function](https://github.com/pingcap/tidb/pull/24179)                       |   | 5 | 4d3h   |
| tipb/221   | [add more regexp functions](https://github.com/pingcap/tipb/pull/221)                                                                  |   | 5 | 1d21h  |
| tidb/24180 | [Revert "planner: donot prune all columns for Projection (#24024)"](https://github.com/pingcap/tidb/pull/24180)                        |   | 7 | 2d4h   |
| tidb/24212 | [*: turn on unused linter](https://github.com/pingcap/tidb/pull/24212)                                                                 |   | 7 | 13h    |


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
| docs-cn/5561 | [Add sql optimization-related docs to toc](https://github.com/pingcap/docs-cn/pull/5561)                                                                |   | 66d15h  |
| tidb/19029   | [types: fix unexpected NOT_NULL flags](https://github.com/pingcap/tidb/pull/19029)                                                                      |   | 266d22h |
| docs/5498    | [partitioning: Corrected partition management](https://github.com/pingcap/docs/pull/5498)                                                               |   | 3d19h   |
| tidb/20708   | [*: separate auto_increment ID allocator from _tidb_rowid allocator](https://github.com/pingcap/tidb/pull/20708)                                        |   | 182d20h |
| tidb/20969   | [executor: Improve the performance of appending not fixed columns](https://github.com/pingcap/tidb/pull/20969)                                          |   | 170d10h |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                          |   | 168d17h |
| tidb/21149   | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                           |   | 161d14h |
| tidb/21304   | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                                |   | 155d12h |
| tidb/21318   | [planner, expression: use the range of column types to simplify expressions](https://github.com/pingcap/tidb/pull/21318)                                |   | 154d19h |
| tidb/21401   | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                                      |   | 150d11h |
| tidb/21508   | [execution: fix dayofweek('0000-00-00') behavior](https://github.com/pingcap/tidb/pull/21508)                                                           |   | 146d10h |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                             |   | 132d19h |
| tidb/21887   | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                                         |   | 131d11h |
| tidb/21954   | [planner/cascades: add rule `PushSelDownApply`](https://github.com/pingcap/tidb/pull/21954)                                                             |   | 127d23h |
| tidb/22146   | [executor: forbid SFU on view](https://github.com/pingcap/tidb/pull/22146)                                                                              |   | 115d22h |
| tidb/22217   | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                                           |   | 113d18h |
| tidb/22234   | [executor, planner: ON DUPLICATE UPDATE can refer to un-project col (#14412)](https://github.com/pingcap/tidb/pull/22234)                               |   | 113d15h |
| tidb/22261   | [time: fix parse datetime won't truncate the reluctant string (#22232)](https://github.com/pingcap/tidb/pull/22261)                                     |   | 112d19h |
| tidb/22374   | [expression: separated arithmeticIntDivideSig](https://github.com/pingcap/tidb/pull/22374)                                                              |   | 107d0h  |
| tidb/22415   | [ddl: refactor bundle[2/2] [6/6]](https://github.com/pingcap/tidb/pull/22415)                                                                           |   | 103d17h |
| tidb/22416   | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                                                       | Y | 103d12h |
| tidb/22541   | [expression: Support builtin function SOUNDEX](https://github.com/pingcap/tidb/pull/22541)                                                              |   | 93d9h   |
| tidb/22565   | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)                                      | Y | 92d17h  |
| tidb/22814   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22814)                                     |   | 69d19h  |
| tidb/22862   | [brie: fix the problem that ddl restored by BR via SQL is not replicated to downstream](https://github.com/pingcap/tidb/pull/22862)                     |   | 66d23h  |
| tidb/22923   | [expression: correct constant propagation for collation (#22666)](https://github.com/pingcap/tidb/pull/22923)                                           |   | 64d15h  |
| tidb/22924   | [planner: fix wrong index merge selection (#22825)](https://github.com/pingcap/tidb/pull/22924)                                                         |   | 64d14h  |
| tidb/23002   | [store/*: fix err check](https://github.com/pingcap/tidb/pull/23002)                                                                                    |   | 61d0h   |
| tidb/23022   | [executor: create PipelinedWindowExec](https://github.com/pingcap/tidb/pull/23022)                                                                      |   | 59d18h  |
| tidb/23152   | [expression: fix wrong error info (#22760)](https://github.com/pingcap/tidb/pull/23152)                                                                 |   | 53d14h  |
| tidb/23196   | [types: fix the bug about the wrong query result for decimal type  (#22507)](https://github.com/pingcap/tidb/pull/23196)                                |   | 51d18h  |
| tidb/23295   | [util, types: don't let SPM be affected by charset (#23161)](https://github.com/pingcap/tidb/pull/23295)                                                |   | 48d11h  |
| tidb/23316   | [planner: Fix rebuild range for prepared plan](https://github.com/pingcap/tidb/pull/23316)                                                              |   | 45d17h  |
| tidb/23373   | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                       |   | 43d19h  |
| tidb/23397   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23397)                                                          |   | 42d17h  |
| tidb/23398   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23398)                                                          |   | 42d17h  |
| tidb/23590   | [planner, table: optimize the list partition pruner for range query](https://github.com/pingcap/tidb/pull/23590)                                        |   | 34d16h  |
| tidb/23598   | [types: fix collation for binary literal (#23591)](https://github.com/pingcap/tidb/pull/23598)                                                          |   | 34d13h  |
| tidb/23655   | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral (#23523)](https://github.com/pingcap/tidb/pull/23655)                |   | 30d22h  |
| tidb/23658   | [*: collect transaction write duration/throughput metrics for SLI/SLO (#23462)](https://github.com/pingcap/tidb/pull/23658)                             |   | 30d22h  |
| tidb/23660   | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23660) |   | 30d20h  |
| tidb/23661   | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23661) |   | 30d20h  |
| tidb/23682   | [executor: fix a panic when batch point get is used for partition table (#23652)](https://github.com/pingcap/tidb/pull/23682)                           |   | 30d17h  |
| tidb/23689   | [planner: fix the panic when we calculate the partition range (#23651)](https://github.com/pingcap/tidb/pull/23689)                                     |   | 30d16h  |
| tidb/23730   | [distsql/*: typo fix for `dispatches`](https://github.com/pingcap/tidb/pull/23730)                                                                      |   | 29d18h  |
| tidb/23796   | [tests: make TestIndexLookupMergeJoinHang and TestIssue18068 stable (#23741)](https://github.com/pingcap/tidb/pull/23796)                               |   | 28d20h  |
| tidb/23812   | [executor, planner: fix collation for hash join building (#23770)](https://github.com/pingcap/tidb/pull/23812)                                          |   | 28d12h  |
| tidb/23878   | [functions: fix some string function has wrong collation and flag (#23835)](https://github.com/pingcap/tidb/pull/23878)                                 |   | 22d21h  |
| tidb/23963   | [executor: checking chunk is full precedes filtering](https://github.com/pingcap/tidb/pull/23963)                                                       |   | 17d17h  |
| tidb/23987   | [executor: Implements json_arrayagg function](https://github.com/pingcap/tidb/pull/23987)                                                               |   | 16d18h  |
| tidb/23997   | [stats, executor: use a correct sampling to collect stats](https://github.com/pingcap/tidb/pull/23997)                                                  |   | 16d10h  |
| tidb/24018   | [ranger: fix the range construction behavior when the column's type is `YEAR` (#23559)](https://github.com/pingcap/tidb/pull/24018)                     |   | 15d18h  |
| tidb/24193   | [executor: implement CTEStorage](https://github.com/pingcap/tidb/pull/24193)                                                                            |   | 8d10h   |
| tidb/24214   | [plan: merge continuous selections and delete surely true expressions](https://github.com/pingcap/tidb/pull/24214)                                      |   | 7d12h   |
| tidb/24229   | [executor: speed up race test TestInsertReorgDelete (#24208)](https://github.com/pingcap/tidb/pull/24229)                                               |   | 6d21h   |
| tidb/24235   | [expression: try to fix TestExprPushDownToFlash tests](https://github.com/pingcap/tidb/pull/24235)                                                      |   | 6d17h   |
| tidb/24241   | [planner/core: remove random test to reduce CI time (#24207)](https://github.com/pingcap/tidb/pull/24241)                                               |   | 6d15h   |
| tidb/24266   | [expression: fix wrong flen infer for bit constant (#23867)](https://github.com/pingcap/tidb/pull/24266)                                                |   | 4d18h   |
| tidb/24267   | [expression: fix wrong flen infer for bit constant (#23867)](https://github.com/pingcap/tidb/pull/24267)                                                |   | 4d18h   |
| tidb/24328   | [*: implement tidb_bound_staleness built-in function](https://github.com/pingcap/tidb/pull/24328)                                                       |   | 2d10h   |
| tidb/24354   | [expression: fix wrong type infer for agg function when type is null (#24290)](https://github.com/pingcap/tidb/pull/24354)                              |   | 1d16h   |
| tidb/24359   | [domain, session: Add new sysvarcache to replace global values cache](https://github.com/pingcap/tidb/pull/24359)                                       |   | 1d7h    |
| tidb/24373   | [planner: filter conflict read_from_storage hints (#24313)](https://github.com/pingcap/tidb/pull/24373)                                                 |   | 19h     |
| tidb/24374   | [planner: filter conflict read_from_storage hints (#24313)](https://github.com/pingcap/tidb/pull/24374)                                                 |   | 19h     |
| tidb/24376   | [planner: prune partitions that will never be used](https://github.com/pingcap/tidb/pull/24376)                                                         |   | 18h     |
| tidb/24379   | [executor: enhancement for ListInDisk(support writing after reading)](https://github.com/pingcap/tidb/pull/24379)                                       |   | 17h     |
| tidb/24382   | [statistics: trigger auto-analyze based on histogram row count](https://github.com/pingcap/tidb/pull/24382)                                             |   | 16h     |


## Reviewed in Last 7 Days

|      REPO      |                                                              PR                                                              | C | D |   R   |
|----------------|------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/24370     | [statistics: fix three bugs about MergePartTopN2GlobalTopN](https://github.com/pingcap/tidb/pull/24370)                      |   | 1 | 1h    |
| tidb/24282     | [planner: remove useless predicates after partition pruning](https://github.com/pingcap/tidb/pull/24282)                     |   | 1 | 3d0h  |
| tidb/24339     | [server,session: do not create stats collector in HTTP API to avoid memory leak](https://github.com/pingcap/tidb/pull/24339) |   | 1 | 22h   |
| tidb/24157     | [planner: let CopTiFlashConcurrencyFactor inflence the cost of whole plan](https://github.com/pingcap/tidb/pull/24157)       |   | 2 | 8d1h  |
| tidb/24313     | [planner: filter conflict read_from_storage hints](https://github.com/pingcap/tidb/pull/24313)                               |   | 3 | 2h    |
| tidb-test/1185 | [randgen-test: fix test for tidb issue:16845](https://github.com/pingcap/tidb-test/pull/1185)                                |   | 3 | 2h    |
| tidb/24290     | [expression: fix wrong type infer for agg function when type is null](https://github.com/pingcap/tidb/pull/24290)            |   | 4 | 1h    |
| tidb/24245     | [planner: between .. and on int column can be used to prune hash partition](https://github.com/pingcap/tidb/pull/24245)      |   | 4 | 2d13h |


</details> 


<details> 
  <summary>rebelice's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                     PR                                                                     | C | LASTED |
|------------|--------------------------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/23836 | [parser, core: Implement force_index hint in parser and TiDB](https://github.com/pingcap/tidb/pull/23836)                                  |   | 27d18h |
| tidb/24033 | [statistics: fix some unstable tests in global stats (#23502)](https://github.com/pingcap/tidb/pull/24033)                                 |   | 15d9h  |
| tidb/24282 | [planner: remove useless predicates after partition pruning](https://github.com/pingcap/tidb/pull/24282)                                   |   | 3d20h  |
| tidb/24306 | [util/ranger: fix func name typo](https://github.com/pingcap/tidb/pull/24306)                                                              |   | 2d22h  |
| tidb/24339 | [server,session: do not create stats collector in HTTP API to avoid memory leak](https://github.com/pingcap/tidb/pull/24339)               |   | 1d22h  |
| tidb/24357 | [statistics: fix a statistics GC problem that can cause duplicated fm-sketch records (#23830)](https://github.com/pingcap/tidb/pull/24357) |   | 1d14h  |
| tidb/24374 | [planner: filter conflict read_from_storage hints (#24313)](https://github.com/pingcap/tidb/pull/24374)                                    |   | 19h    |


## Reviewed in Last 7 Days

|    REPO    |                                                           PR                                                            | C | D |   R   |
|------------|-------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/24376 | [planner: prune partitions that will never be used](https://github.com/pingcap/tidb/pull/24376)                         |   | 1 | 0h    |
| tidb/24245 | [planner: between .. and on int column can be used to prune hash partition](https://github.com/pingcap/tidb/pull/24245) |   | 4 | 2d15h |


</details> 


<details> 
  <summary>time-and-fate's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                         PR                                                          | C | LASTED  |
|------------|---------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                           |   | 175d17h |
| tidb/22416 | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                   | Y | 103d12h |
| tidb/24155 | [planner, executor: fix index merge partial table scan schema (#23936)](https://github.com/pingcap/tidb/pull/24155) |   | 9d20h   |
| tidb/24230 | [*: consitent get infoschema](https://github.com/pingcap/tidb/pull/24230)                                           |   | 6d21h   |
| tidb/24373 | [planner: filter conflict read_from_storage hints (#24313)](https://github.com/pingcap/tidb/pull/24373)             |   | 19h     |
| tidb/24374 | [planner: filter conflict read_from_storage hints (#24313)](https://github.com/pingcap/tidb/pull/24374)             |   | 19h     |
| tidb/24376 | [planner: prune partitions that will never be used](https://github.com/pingcap/tidb/pull/24376)                     |   | 18h     |
| tidb/24382 | [statistics: trigger auto-analyze based on histogram row count](https://github.com/pingcap/tidb/pull/24382)         |   | 16h     |


## Reviewed in Last 7 Days

|    REPO    |                                                                  PR                                                                  | C | D |   R    |
|------------|--------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/23997 | [stats, executor: use a correct sampling to collect stats](https://github.com/pingcap/tidb/pull/23997)                               |   | 1 | 15d19h |
| tidb/24313 | [planner: filter conflict read_from_storage hints](https://github.com/pingcap/tidb/pull/24313)                                       |   | 1 | 1d22h  |
| tidb/24352 | [statistics: skip reading mysql.stats_histograms if cached stats is up-to-date (#24175)](https://github.com/pingcap/tidb/pull/24352) |   | 1 | 17h    |
| tidb/24317 | [statistics: skip reading mysql.stats_histograms if cached stats is up-to-date (#24175)](https://github.com/pingcap/tidb/pull/24317) |   | 2 | 23h    |
| tidb/24204 | [planner: clone possible properties before saving them](https://github.com/pingcap/tidb/pull/24204)                                  |   | 2 | 5d18h  |


</details> 


<details> 
  <summary>winoros's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                              PR                                                                               | C | LASTED  |
|--------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| docs-cn/5916 | [sql-statements, information-schema: add `END_TIME` field for table `ANALYZE_STATUS`](https://github.com/pingcap/docs-cn/pull/5916)                           |   | 28d17h  |
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                                        | Y | 230d14h |
| docs-cn/6113 | [config: update the default value of `feedback-probability`](https://github.com/pingcap/docs-cn/pull/6113)                                                    |   | 7d22h   |
| tidb/20765   | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                                             |   | 178d17h |
| tidb/20877   | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                                                     |   | 175d17h |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                                |   | 168d17h |
| tidb/21207   | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                                               |   | 157d19h |
| tidb/21487   | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                                                 |   | 147d5h  |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                                   |   | 132d19h |
| tidb/21954   | [planner/cascades: add rule `PushSelDownApply`](https://github.com/pingcap/tidb/pull/21954)                                                                   |   | 127d23h |
| tidb/22181   | [planner, expression: fix error when using IN combined with subquery (#22080)](https://github.com/pingcap/tidb/pull/22181)                                    |   | 114d18h |
| tidb/22416   | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                                                             | Y | 103d12h |
| tidb/22504   | [*:Fix the fetchHotRegion bug that the count always zero](https://github.com/pingcap/tidb/pull/22504)                                                         |   | 95d19h  |
| tidb/22565   | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)                                            | Y | 92d17h  |
| tidb/22624   | [ planner: not pruning column used by union scan condition (#21640)](https://github.com/pingcap/tidb/pull/22624)                                              |   | 90d17h  |
| tidb/22923   | [expression: correct constant propagation for collation (#22666)](https://github.com/pingcap/tidb/pull/22923)                                                 |   | 64d15h  |
| tidb/23233   | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23233)                                                        |   | 50d18h  |
| tidb/23347   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                                          |   | 44d18h  |
| tidb/23348   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                                          |   | 44d18h  |
| tidb/23350   | [util/stringutil, util/ranger, planner: use hierarchical separators to simplify the parsing for info of EXPLAIN ](https://github.com/pingcap/tidb/pull/23350) |   | 44d17h  |
| tidb/23373   | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                             |   | 43d19h  |
| tidb/23474   | [planner: fix inappropriate null flag of null constants (#23457)](https://github.com/pingcap/tidb/pull/23474)                                                 |   | 37d18h  |
| tidb/23543   | [statistics: fix auto analyze log information incomplete (#23522)](https://github.com/pingcap/tidb/pull/23543)                                                |   | 35d18h  |
| tidb/23598   | [types: fix collation for binary literal (#23591)](https://github.com/pingcap/tidb/pull/23598)                                                                |   | 34d13h  |
| tidb/23655   | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral (#23523)](https://github.com/pingcap/tidb/pull/23655)                      |   | 30d22h  |
| tidb/23689   | [planner: fix the panic when we calculate the partition range (#23651)](https://github.com/pingcap/tidb/pull/23689)                                           |   | 30d16h  |
| tidb/23772   | [tablecodec: fix text type decode for old row format (#23751)](https://github.com/pingcap/tidb/pull/23772)                                                    |   | 29d11h  |
| tidb/23849   | [ddl: tidb panic while query hash partition table with is null condition](https://github.com/pingcap/tidb/pull/23849)                                         |   | 24d13h  |
| tidb/23917   | [planner: fix wrong TableDual plans caused by comparing Binary and Bytes incorrectly (#23860)](https://github.com/pingcap/tidb/pull/23917)                    |   | 21d23h  |
| tidb/23970   | [planner: fix a bug that point get plan returns wrong column name (#23365)](https://github.com/pingcap/tidb/pull/23970)                                       |   | 17d14h  |
| tidb/24018   | [ranger: fix the range construction behavior when the column's type is `YEAR` (#23559)](https://github.com/pingcap/tidb/pull/24018)                           |   | 15d18h  |
| tidb/24060   | [statistics: fix some potential panic in statistics (#23988)](https://github.com/pingcap/tidb/pull/24060)                                                     |   | 14d13h  |
| tidb/24061   | [statistics: fix some potential panic in statistics (#23988)](https://github.com/pingcap/tidb/pull/24061)                                                     |   | 14d13h  |
| tidb/24079   | [planner: change descScanFactor to scanFactor when ExpectedCount is small. (#23972)](https://github.com/pingcap/tidb/pull/24079)                              |   | 13d19h  |
| tidb/24097   | [planner: Remove redundant call to expression.ColumnSubstitute](https://github.com/pingcap/tidb/pull/24097)                                                   |   | 13d3h   |
| tidb/24138   | [planner: Add Equivalence Rules to Transform BinaryOptSubquery to ExistsSubquery](https://github.com/pingcap/tidb/pull/24138)                                 |   | 10d12h  |
| tidb/24241   | [planner/core: remove random test to reduce CI time (#24207)](https://github.com/pingcap/tidb/pull/24241)                                                     |   | 6d15h   |
| tidb/24258   | [Revert "planner: donot prune all columns for Projection (#24024)" (#24180)](https://github.com/pingcap/tidb/pull/24258)                                      |   | 4d22h   |
| tidb/24342   | [planner: create new column slice in PreparePossibleProperties](https://github.com/pingcap/tidb/pull/24342)                                                   |   | 1d20h   |
| tidb/24369   | [planner: fix column pruning bug for Apply and Join](https://github.com/pingcap/tidb/pull/24369)                                                              |   | 22h     |
| tidb/24376   | [planner: prune partitions that will never be used](https://github.com/pingcap/tidb/pull/24376)                                                               |   | 18h     |
| tidb/24382   | [statistics: trigger auto-analyze based on histogram row count](https://github.com/pingcap/tidb/pull/24382)                                                   |   | 16h     |


## Reviewed in Last 7 Days

|    REPO    |                                                                     PR                                                                     | C | D |   R   |
|------------|--------------------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/24370 | [statistics: fix three bugs about MergePartTopN2GlobalTopN](https://github.com/pingcap/tidb/pull/24370)                                    |   | 1 | 2h    |
| tidb/24282 | [planner: remove useless predicates after partition pruning](https://github.com/pingcap/tidb/pull/24282)                                   |   | 1 | 3d0h  |
| tidb/24352 | [statistics: skip reading mysql.stats_histograms if cached stats is up-to-date (#24175)](https://github.com/pingcap/tidb/pull/24352)       |   | 1 | 19h   |
| tidb/24357 | [statistics: fix a statistics GC problem that can cause duplicated fm-sketch records (#23830)](https://github.com/pingcap/tidb/pull/24357) |   | 1 | 17h   |
| tidb/24317 | [statistics: skip reading mysql.stats_histograms if cached stats is up-to-date (#24175)](https://github.com/pingcap/tidb/pull/24317)       |   | 2 | 1d0h  |
| tidb/22691 | [planner, expression: support enum index scan](https://github.com/pingcap/tidb/pull/22691)                                                 |   | 2 | 84d0h |
| tidb/24214 | [plan: merge continuous selections and delete surely true expressions](https://github.com/pingcap/tidb/pull/24214)                         |   | 2 | 5d17h |
| tidb/24204 | [planner: clone possible properties before saving them](https://github.com/pingcap/tidb/pull/24204)                                        |   | 2 | 5d22h |


</details> 


<details> 
  <summary>wshwsh12's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                 PR                                                                  | C | LASTED  |
|------------|-------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19807 | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                             |   | 237d11h |
| tidb/19957 | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                              | Y | 230d14h |
| tidb/21487 | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                       |   | 147d5h  |
| tidb/21887 | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                     |   | 131d11h |
| tidb/22378 | [executor: vectorize hash aggregate](https://github.com/pingcap/tidb/pull/22378)                                                    |   | 106d19h |
| tidb/23336 | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23336)            |   | 44d19h  |
| tidb/23347 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                |   | 44d18h  |
| tidb/23348 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                |   | 44d18h  |
| tidb/23368 | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23368)               |   | 43d20h  |
| tidb/23397 | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23397)                                      |   | 42d17h  |
| tidb/23398 | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23398)                                      |   | 42d17h  |
| tidb/23519 | [executor: check privilege before adding](https://github.com/pingcap/tidb/pull/23519)                                               |   | 36d0h   |
| tidb/23760 | [collation: fix tidb panic when compare string with collation](https://github.com/pingcap/tidb/pull/23760)                          |   | 29d13h  |
| tidb/23968 | [statistics: fix unstable TestDropPartitionStats test](https://github.com/pingcap/tidb/pull/23968)                                  |   | 17d15h  |
| tidb/23979 | [executor, statistics: fix unstable `TestAnalyzeIndexExtractTopN`](https://github.com/pingcap/tidb/pull/23979)                      |   | 16d23h  |
| tidb/24018 | [ranger: fix the range construction behavior when the column's type is `YEAR` (#23559)](https://github.com/pingcap/tidb/pull/24018) |   | 15d18h  |
| tidb/24033 | [statistics: fix some unstable tests in global stats (#23502)](https://github.com/pingcap/tidb/pull/24033)                          |   | 15d9h   |
| tidb/24050 | [expression: fix get var panic when types not match](https://github.com/pingcap/tidb/pull/24050)                                    |   | 14d17h  |
| tidb/24053 | [executor: fix wrong convert from bit to string when do projection (#23960)](https://github.com/pingcap/tidb/pull/24053)            |   | 14d16h  |
| tidb/24054 | [executor: fix wrong convert from bit to string when do projection (#23960)](https://github.com/pingcap/tidb/pull/24054)            |   | 14d16h  |
| tidb/24147 | [docs/design: add proposal for common table expression](https://github.com/pingcap/tidb/pull/24147)                                 |   | 9d23h   |
| tidb/24186 | [executor: make column default value being aware of NO_ZERO_IN_DATE (#24174)](https://github.com/pingcap/tidb/pull/24186)           |   | 8d19h   |
| tidb/24228 | [executor: skip TestPrepareStmtAfterIsolationReadChange when race enable (#24200)](https://github.com/pingcap/tidb/pull/24228)      |   | 6d22h   |
| tidb/24229 | [executor: speed up race test TestInsertReorgDelete (#24208)](https://github.com/pingcap/tidb/pull/24229)                           |   | 6d21h   |
| tidb/24230 | [*: consitent get infoschema](https://github.com/pingcap/tidb/pull/24230)                                                           |   | 6d21h   |
| tidb/24236 | [*: remove SchemaVersion in TransactionContext](https://github.com/pingcap/tidb/pull/24236)                                         |   | 6d17h   |
| tidb/24266 | [expression: fix wrong flen infer for bit constant (#23867)](https://github.com/pingcap/tidb/pull/24266)                            |   | 4d18h   |
| tidb/24267 | [expression: fix wrong flen infer for bit constant (#23867)](https://github.com/pingcap/tidb/pull/24267)                            |   | 4d18h   |
| tidb/24268 | [expression: fix cast real, decimal to time (#24120)](https://github.com/pingcap/tidb/pull/24268)                                   |   | 4d17h   |
| tidb/24280 | [executor, session, variable: Move deprecation and synonyms to sysvar struct](https://github.com/pingcap/tidb/pull/24280)           |   | 4d0h    |
| tidb/24341 | [executor: fix projection executor panic and add failpoint test (#24231)](https://github.com/pingcap/tidb/pull/24341)               |   | 1d20h   |
| tidb/24345 | [executor: fix data race of parallel apply operator (#24257)](https://github.com/pingcap/tidb/pull/24345)                           |   | 1d19h   |
| tidb/24354 | [expression: fix wrong type infer for agg function when type is null (#24290)](https://github.com/pingcap/tidb/pull/24354)          |   | 1d16h   |
| tidb/24380 | [ executor: Pass the SQL digest down to pessimistic lock request](https://github.com/pingcap/tidb/pull/24380)                       |   | 17h     |


## Reviewed in Last 7 Days

|    REPO    |                                                                             PR                                                                              | C | D |   R    |
|------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/24338 | [executor: fix a concurrent-access problem caused by accessing a single parser object in session concurrently ](https://github.com/pingcap/tidb/pull/24338) |   | 1 | 1d4h   |
| tidb/24234 | [executor: skip TestMppExecution when race is enabled (#24222)](https://github.com/pingcap/tidb/pull/24234)                                                 |   | 4 | 3d6h   |
| tidb/24260 | [executor: make IndexReaders in the inner side of IndexJoins support direct reading](https://github.com/pingcap/tidb/pull/24260)                            |   | 4 | 1d8h   |
| tidb/22686 | [expression: support enum pushdown](https://github.com/pingcap/tidb/pull/22686)                                                                             |   | 4 | 82d2h  |
| tidb/23867 | [expression: fix wrong flen infer for bit constant](https://github.com/pingcap/tidb/pull/23867)                                                             |   | 7 | 16d20h |
| tidb/24222 | [executor: skip TestMppExecution when race is enabled](https://github.com/pingcap/tidb/pull/24222)                                                          |   | 7 | 0h     |
| tidb/24208 | [executor: speed up race test TestInsertReorgDelete](https://github.com/pingcap/tidb/pull/24208)                                                            |   | 7 | 14h    |
| tidb/24200 | [executor: skip TestPrepareStmtAfterIsolationReadChange when race enable](https://github.com/pingcap/tidb/pull/24200)                                       |   | 7 | 19h    |


</details> 

