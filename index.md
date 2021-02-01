|   REVIEWER    |  O(C)   | 1 | 2 | 3 | 4 | 5  | 6  | 7  | T  |
|---------------|---------|---|---|---|---|----|----|----|----|
| Reminiscent   |  3 ( 0) | 0 | 0 | 0 | 0 |  1 |  1 |  0 |  2 |
| SunRunAway    | 44 ( 5) | 0 | 0 | 0 | 0 |  0 |  0 |  0 |  0 |
| XuHuaiyu      | 56 ( 4) | 0 | 0 | 1 | 2 |  5 |  0 |  0 |  8 |
| dyzsr         |  0 ( 0) | 0 | 0 | 0 | 0 |  0 |  0 |  0 |  0 |
| eurekaka      | 19 ( 5) | 0 | 0 | 1 | 1 |  3 |  0 |  2 |  7 |
| fzhedu        |  2 ( 2) | 0 | 0 | 0 | 1 |  0 |  0 |  1 |  2 |
| ichn-hu       |  2 ( 0) | 0 | 0 | 0 | 0 |  0 |  0 |  1 |  1 |
| lzmhhh123     | 39 ( 2) | 0 | 0 | 2 | 1 |  4 |  4 |  1 | 12 |
| nullnotnil    |  0 ( 0) | 0 | 0 | 0 | 0 |  0 |  0 |  0 |  0 |
| qw4990        | 52 ( 6) | 0 | 0 | 4 | 9 | 13 | 37 | 10 | 73 |
| rebelice      |  0 ( 0) | 0 | 0 | 0 | 0 |  1 |  0 |  0 |  1 |
| time-and-fate |  1 ( 0) | 0 | 0 | 2 | 0 |  0 |  0 |  1 |  3 |
| winoros       | 30 ( 5) | 0 | 0 | 2 | 2 |  4 |  1 |  3 | 12 |
| wshwsh12      | 24 ( 4) | 0 | 0 | 2 | 0 |  0 |  0 |  0 |  2 |


<details> 
  <summary>Reminiscent's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                   PR                                                                   | C | LASTED |
|------------|----------------------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/21896 | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                             |   | 41d19h |
| tidb/22330 | [planner: check error when correlatedAggregateResolver leaves ast.Node (#22222)](https://github.com/pingcap/tidb/pull/22330)           |   | 20d23h |
| tidb/22354 | [planner: do not cache prepared plan if optimization depends on mutable constant (#22349)](https://github.com/pingcap/tidb/pull/22354) |   | 19d23h |


## Reviewed in Last 7 Days

|    REPO    |                                                             PR                                                              | C | D |   R    |
|------------|-----------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/21964 | [planner: add plancodec id for all type TableScan/IndexScan. (#21935)](https://github.com/pingcap/tidb/pull/21964)          |   | 5 | 34d23h |
| tidb/22531 | [expression: do not rewrite `like` to `=` if new collation is enabled (#21893)](https://github.com/pingcap/tidb/pull/22531) |   | 6 | 18h    |


</details> 


<details> 
  <summary>SunRunAway's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                      PR                                                                       | C | LASTED  |
|--------------|-----------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| docs-cn/4913 | [explain: add indexes](https://github.com/pingcap/docs-cn/pull/4913)                                                                          |   | 76d18h  |
| tidb/15370   | [planner,executor: Refactor Shuffle and implement parallel Sort](https://github.com/pingcap/tidb/pull/15370)                                  | Y | 323d19h |
| docs-cn/4933 | [explain: add joins](https://github.com/pingcap/docs-cn/pull/4933)                                                                            |   | 72d20h  |
| tidb/15462   | [executor: implement `graceHashJoin`](https://github.com/pingcap/tidb/pull/15462)                                                             | Y | 319d17h |
| tidb/16967   | [executor: Refactor Shuffle and implement parallel sort (executor part)](https://github.com/pingcap/tidb/pull/16967)                          | Y | 274d10h |
| tidb/17238   | [*: refactor table.Allocator to improve readability](https://github.com/pingcap/tidb/pull/17238)                                              |   | 261d18h |
| tidb/19120   | [executor: Concurrently fetch chunks and insert them to a concurrent hash table in hash build](https://github.com/pingcap/tidb/pull/19120)    |   | 173d21h |
| tidb/19178   | [executor: Refactor probe channel](https://github.com/pingcap/tidb/pull/19178)                                                                |   | 171d17h |
| tidb/19347   | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347)          |   | 163d23h |
| tidb/19807   | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                                       |   | 149d11h |
| tidb/19900   | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                                | Y | 144d18h |
| tidb/20140   | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                            |   | 131d22h |
| tidb/20220   | [*: new secondary index value format](https://github.com/pingcap/tidb/pull/20220)                                                             |   | 128d16h |
| tidb/20316   | [docs/design: add design doc for index usage information](https://github.com/pingcap/tidb/pull/20316)                                         |   | 123d17h |
| tidb/20335   | [planner, executor: enable inline projection for Selection](https://github.com/pingcap/tidb/pull/20335)                                       | Y | 120d18h |
| tidb/20360   | [planner: refine explain info for batch cop](https://github.com/pingcap/tidb/pull/20360)                                                      |   | 114d22h |
| tidb/20397   | [parser: replace ast.SelectLockInShareMode with ast.SelectLockForShare](https://github.com/pingcap/tidb/pull/20397)                           |   | 112d18h |
| tidb/20615   | [utils: Avoid panic when getting memory](https://github.com/pingcap/tidb/pull/20615)                                                          |   | 100d2h  |
| tidb/20689   | [expression: make TIME function compatible with MySQL (#19158)](https://github.com/pingcap/tidb/pull/20689)                                   |   | 95d20h  |
| tidb/20752   | [*: trace statsCache and preparePlanCache by Global memory tracker.](https://github.com/pingcap/tidb/pull/20752)                              |   | 90d22h  |
| tidb/20765   | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                             |   | 90d17h  |
| tidb/21207   | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                               |   | 69d19h  |
| tidb/21277   | [executor: fix split table with large integers](https://github.com/pingcap/tidb/pull/21277)                                                   |   | 67d20h  |
| tidb/21364   | [expression: Add test cases to cover the cases when invalid int value is casted as TIME (#18653)](https://github.com/pingcap/tidb/pull/21364) |   | 63d1h   |
| tidb/21381   | [*: optimize analyze cluster index table](https://github.com/pingcap/tidb/pull/21381)                                                         |   | 62d17h  |
| tidb/21386   | [expression: Disable cast decimal as string push down to TiFlash](https://github.com/pingcap/tidb/pull/21386)                                 |   | 62d16h  |
| tidb/21834   | [planner: enhanced index range calculation plan](https://github.com/pingcap/tidb/pull/21834)                                                  |   | 46d18h  |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                   |   | 44d19h  |
| tidb/21878   | [planner: do not push down lock to pointGet/bacthPointGet when selection exists](https://github.com/pingcap/tidb/pull/21878)                  |   | 44d18h  |
| tidb/21890   | [*: redact some error code, part(3/3) (#21866)](https://github.com/pingcap/tidb/pull/21890)                                                   |   | 42d15h  |
| tidb/21956   | [planner/preprocessor: disallow into-outfile clause in some place](https://github.com/pingcap/tidb/pull/21956)                                |   | 39d23h  |
| tidb/22026   | [expression: separated arithmeticPlusIntSig](https://github.com/pingcap/tidb/pull/22026)                                                      |   | 37d20h  |
| tidb/22043   | [planner, executor: enhance the limit pushdown rule.](https://github.com/pingcap/tidb/pull/22043)                                             |   | 35d10h  |
| tidb/22089   | [executor: fix signed cluster index behavior (#22085)](https://github.com/pingcap/tidb/pull/22089)                                            |   | 32d22h  |
| tidb/22104   | [executor: fix incompatible escape behaviors in `select into outfile` (#22100)](https://github.com/pingcap/tidb/pull/22104)                   |   | 32d16h  |
| tidb/22107   | [executor: avoid log duplicate index name in slow-log (#22057)](https://github.com/pingcap/tidb/pull/22107)                                   |   | 32d14h  |
| tidb/22114   | [test: fix globalkilltest (#21987)](https://github.com/pingcap/tidb/pull/22114)                                                               |   | 32d12h  |
| tidb/22120   | [executor: fix `update ignore` into not exists partition (#21984)](https://github.com/pingcap/tidb/pull/22120)                                |   | 31d22h  |
| tidb/22181   | [planner, expression: fix error when using IN combined with subquery (#22080)](https://github.com/pingcap/tidb/pull/22181)                    |   | 26d17h  |
| tidb/22217   | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                                 |   | 25d17h  |
| tidb/22330   | [planner: check error when correlatedAggregateResolver leaves ast.Node (#22222)](https://github.com/pingcap/tidb/pull/22330)                  |   | 20d23h  |
| tidb/22365   | [planner: check index valid while forUpdateRead (#22152)](https://github.com/pingcap/tidb/pull/22365)                                         |   | 19d19h  |
| tidb/22379   | [[experiment] executor: allow aggregation to spill disk when running out of memory quota](https://github.com/pingcap/tidb/pull/22379)         |   | 18d19h  |
| tidb/22485   | [docs: Add design doc for security enhanced mode](https://github.com/pingcap/tidb/pull/22485)                                                 |   | 10d5h   |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>XuHuaiyu's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                              PR                                                                              | C | LASTED  |
|--------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19292   | [planner: suppport left join in join reorder](https://github.com/pingcap/tidb/pull/19292)                                                                    |   | 165d17h |
| docs-cn/5323 | [Update parameter type description](https://github.com/pingcap/docs-cn/pull/5323)                                                                            |   | 13d19h  |
| tidb/19900   | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                                               | Y | 144d18h |
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                                       | Y | 142d14h |
| tidb/20040   | [planner, expression: take NullFlag into consideration when optimize the `int non-const` <cmp > `non-int const`](https://github.com/pingcap/tidb/pull/20040) | Y | 137d14h |
| tidb/20140   | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                                           |   | 131d22h |
| tidb/20311   | [expression: fix overflow error when convert bit to int64 (#20266)](https://github.com/pingcap/tidb/pull/20311)                                              |   | 123d21h |
| tidb/20350   | [executor: support read global indexes in IndexMergeReader and index join](https://github.com/pingcap/tidb/pull/20350)                                       | Y | 117d14h |
| tidb/20505   | [*: Add metrics for oom-action and sql memory usage.](https://github.com/pingcap/tidb/pull/20505)                                                            |   | 104d19h |
| tidb/20576   | [*: fix stats feedback after tableReader handle multiple ranges](https://github.com/pingcap/tidb/pull/20576)                                                 |   | 102d13h |
| tidb/20613   | [executor: fix issue of hash join fetch time inaccurate](https://github.com/pingcap/tidb/pull/20613)                                                         |   | 100d13h |
| tidb/20752   | [*: trace statsCache and preparePlanCache by Global memory tracker.](https://github.com/pingcap/tidb/pull/20752)                                             |   | 90d22h  |
| tidb/20790   | [collation: add pinyin collation for chinese charset support](https://github.com/pingcap/tidb/pull/20790)                                                    |   | 89d21h  |
| tidb/20793   | [planner, executor: enable inline projection for Apply](https://github.com/pingcap/tidb/pull/20793)                                                          |   | 89d20h  |
| tidb/20905   | [planner: fix statement-optimize not work in `TryFastPlan`](https://github.com/pingcap/tidb/pull/20905)                                                      |   | 86d17h  |
| tidb/20972   | [expression: POC implementation of Vitess hashing algorithm.](https://github.com/pingcap/tidb/pull/20972)                                                    |   | 82d1h   |
| tidb/21064   | [planner, executor: fix cast not check error](https://github.com/pingcap/tidb/pull/21064)                                                                    |   | 77d8h   |
| tidb/21149   | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                                |   | 73d14h  |
| tidb/21228   | [executor: return the result immediately when combining LIMIT row_count with DISTINCT](https://github.com/pingcap/tidb/pull/21228)                           |   | 69d13h  |
| tidb/21304   | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                                     |   | 67d12h  |
| tidb/21334   | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                                                |   | 66d14h  |
| tidb/21340   | [executor: initialize expensive query handler on domain creation](https://github.com/pingcap/tidb/pull/21340)                                                |   | 66d0h   |
| tidb/21476   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                        |   | 59d15h  |
| tidb/21536   | [executor: add slow-log file meta cache to avoid repeat read file meta information](https://github.com/pingcap/tidb/pull/21536)                              |   | 55d15h  |
| tidb/21564   | [ddl: fix Incorrect behavior of NO_ZERO_DATE when altering table](https://github.com/pingcap/tidb/pull/21564)                                                |   | 54d15h  |
| tidb/21626   | [test: convert test to benchmard test to make ci stable (#21616)](https://github.com/pingcap/tidb/pull/21626)                                                |   | 52d23h  |
| tidb/21680   | [planner: report error when ORDER BY conflicts with DISTINCT (#21286)](https://github.com/pingcap/tidb/pull/21680)                                           |   | 51d16h  |
| tidb/21839   | [planner/core: add 'split table using statistics' statement](https://github.com/pingcap/tidb/pull/21839)                                                     |   | 46d15h  |
| tidb/21853   | [expression: fix compatibility behaviors in time_format with MySQL (#21559)](https://github.com/pingcap/tidb/pull/21853)                                     |   | 45d19h  |
| tidb/21896   | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                                                   |   | 41d19h  |
| tidb/22014   | [executor: fix unstable test Issue16696 (#22009)](https://github.com/pingcap/tidb/pull/22014)                                                                |   | 38d17h  |
| tidb/22107   | [executor: avoid log duplicate index name in slow-log (#22057)](https://github.com/pingcap/tidb/pull/22107)                                                  |   | 32d14h  |
| tidb/22120   | [executor: fix `update ignore` into not exists partition (#21984)](https://github.com/pingcap/tidb/pull/22120)                                               |   | 31d22h  |
| tidb/22131   | [privilege: remove leading and trailing space when create user and role](https://github.com/pingcap/tidb/pull/22131)                                         |   | 31d19h  |
| tidb/22142   | [store: trace `loadRegion` to see the PD region cache loading (#22092)](https://github.com/pingcap/tidb/pull/22142)                                          |   | 28d0h   |
| tidb/22149   | [session: set process info before building plan (#22101)](https://github.com/pingcap/tidb/pull/22149)                                                        |   | 27d19h  |
| tidb/22163   | [expression: separated arithmeticMinusIntSig](https://github.com/pingcap/tidb/pull/22163)                                                                    |   | 27d13h  |
| tidb/22186   | [executor: fix select into outfile with year type column has no data (#22175)](https://github.com/pingcap/tidb/pull/22186)                                   |   | 26d16h  |
| tidb/22294   | [planner, table: optimize the list partition pruner for range query](https://github.com/pingcap/tidb/pull/22294)                                             |   | 23d20h  |
| tidb/22307   | [ddl: fix update can see columns not public](https://github.com/pingcap/tidb/pull/22307)                                                                     |   | 23d16h  |
| tidb/22381   | [planner: check schema stale for plan cache when forUpdateRead](https://github.com/pingcap/tidb/pull/22381)                                                  |   | 18d15h  |
| tidb/22403   | [planner: reorder inner joins simplified from outer joins (#22392)](https://github.com/pingcap/tidb/pull/22403)                                              |   | 16d22h  |
| tidb/22407   | [types: fix return err when decimal from string value](https://github.com/pingcap/tidb/pull/22407)                                                           |   | 16d19h  |
| tidb/22418   | [expression: Optimize builtinArithmeticModRealSig and builtinGreatestDecimalSig using MergeNull method](https://github.com/pingcap/tidb/pull/22418)          |   | 14d0h   |
| tidb/22426   | [expression: fix bugs in builtinfunction ArithmeticMinusInt logic](https://github.com/pingcap/tidb/pull/22426)                                               |   | 13d16h  |
| tidb/22432   | [types,execute: fix errcode return like mysql when inserting incorrect int value ](https://github.com/pingcap/tidb/pull/22432)                               |   | 12d21h  |
| tidb/22463   | [executor: make memory tracker for aggregate more accurate.](https://github.com/pingcap/tidb/pull/22463)                                                     |   | 10d23h  |
| tidb/22472   | [planner, statistics: build the global statistics for the partition table](https://github.com/pingcap/tidb/pull/22472)                                       |   | 10d17h  |
| tidb/22491   | [executor: skip null data in common handle during point-get (#22483)](https://github.com/pingcap/tidb/pull/22491)                                            |   | 9d20h   |
| tidb/22507   | [types: fix the bug about the wrong query result for decimal type ](https://github.com/pingcap/tidb/pull/22507)                                              |   | 6d23h   |
| tidb/22546   | [*: refactor ExecuteInternal to return single resultset](https://github.com/pingcap/tidb/pull/22546)                                                         |   | 4d23h   |
| tidb/22597   | [session, exectutor: Guarantee external consistency by default; Add an explicit begin statement to disable it](https://github.com/pingcap/tidb/pull/22597)   |   | 3d17h   |
| tidb/22616   | [expression: from_unixtime accept 64-bit integers](https://github.com/pingcap/tidb/pull/22616)                                                               |   | 2d23h   |
| tidb/22617   | [metrics: fix wrong bucket name of coprocessor cache (#22454)](https://github.com/pingcap/tidb/pull/22617)                                                   |   | 2d23h   |
| tidb/22618   | [metrics: fix wrong bucket name of coprocessor cache (#22454)](https://github.com/pingcap/tidb/pull/22618)                                                   |   | 2d23h   |
| tidb/22624   | [ planner: not pruning column used by union scan condition (#21640)](https://github.com/pingcap/tidb/pull/22624)                                             |   | 2d17h   |


## Reviewed in Last 7 Days

|     REPO     |                                                               PR                                                                | C | D |   R    |
|--------------|---------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/22454   | [metrics: fix wrong bucket name of coprocessor cache](https://github.com/pingcap/tidb/pull/22454)                               |   | 3 | 8d17h  |
| tidb/22593   | [server: remove TokenLimitGauge and use ConfigStatus instead. (#22590)](https://github.com/pingcap/tidb/pull/22593)             |   | 4 | 0h     |
| tidb/22570   | [server: add token usage gauge for tidb (#22511)](https://github.com/pingcap/tidb/pull/22570)                                   |   | 4 | 16h    |
| tidb/22577   | [expression, types: fix unexpected result from TIME() when fsp digits > 6 (#21652)](https://github.com/pingcap/tidb/pull/22577) |   | 5 | 0h     |
| docs/4729    | [tidb: Refine the doc for global kill](https://github.com/pingcap/docs/pull/4729)                                               |   | 5 | 1h     |
| tidb/21676   | [expression: fix compatibility of extract day_time unit functions (#21601)](https://github.com/pingcap/tidb/pull/21676)         |   | 5 | 46d18h |
| tidb/21700   | [expression: fix incompatible result of `JSON_SEARCH()` (#20164)](https://github.com/pingcap/tidb/pull/21700)                   | Y | 5 | 43d19h |
| docs-cn/5386 | [tidb: add doc for global kill](https://github.com/pingcap/docs-cn/pull/5386)                                                   |   | 5 | 1d20h  |


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
| tidb/14729 | [planner: fix constant propagation for PredicatePushDown](https://github.com/pingcap/tidb/pull/14729)                                  | Y | 355d18h |
| tidb/14831 | [planner/cascades: add implementationRule for IndexLookUpJoin](https://github.com/pingcap/tidb/pull/14831)                             |   | 348d17h |
| tidb/15090 | [planner/cascades: refine the row count estimation of TiKV layer Selection](https://github.com/pingcap/tidb/pull/15090)                |   | 334d18h |
| tidb/15157 | [planner/cascades: implement `HashCode` method for all the LogicalPlans](https://github.com/pingcap/tidb/pull/15157)                   | Y | 332d14h |
| tidb/15335 | [planner/cascades: add transformation rule PullAggregationUpApply & EliminateMaxOneRow](https://github.com/pingcap/tidb/pull/15335)    |   | 325d18h |
| tidb/15370 | [planner,executor: Refactor Shuffle and implement parallel Sort](https://github.com/pingcap/tidb/pull/15370)                           | Y | 323d19h |
| tidb/17276 | [planner/cascades: add rule InjectProjectionBelowSort](https://github.com/pingcap/tidb/pull/17276)                                     | Y | 258d9h  |
| tidb/18882 | [planner, executor: add explain for `MetricSummaryTableExtractor`](https://github.com/pingcap/tidb/pull/18882)                         | Y | 185d17h |
| tidb/19347 | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347)   |   | 163d23h |
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                              |   | 87d17h  |
| tidb/21444 | [planner: ignore anonymous index while tiflash replica is available](https://github.com/pingcap/tidb/pull/21444)                       |   | 60d12h  |
| tidb/21680 | [planner: report error when ORDER BY conflicts with DISTINCT (#21286)](https://github.com/pingcap/tidb/pull/21680)                     |   | 51d16h  |
| tidb/21994 | [range: fix overflow value access index ](https://github.com/pingcap/tidb/pull/21994)                                                  |   | 38d23h  |
| tidb/22330 | [planner: check error when correlatedAggregateResolver leaves ast.Node (#22222)](https://github.com/pingcap/tidb/pull/22330)           |   | 20d23h  |
| tidb/22342 | [session: fix two cases when updating bind info (#22338)](https://github.com/pingcap/tidb/pull/22342)                                  |   | 20d18h  |
| tidb/22354 | [planner: do not cache prepared plan if optimization depends on mutable constant (#22349)](https://github.com/pingcap/tidb/pull/22354) |   | 19d23h  |
| tidb/22369 | [session: fix the duplicate binding case when updating bind info (#22367)](https://github.com/pingcap/tidb/pull/22369)                 |   | 19d17h  |
| tidb/22416 | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                                      |   | 15d12h  |
| tidb/22559 | [planner: split test data from test cases in cbo_test.go](https://github.com/pingcap/tidb/pull/22559)                                  |   | 4d19h   |


## Reviewed in Last 7 Days

|    REPO     |                                                         PR                                                         | C | D |   R   |
|-------------|--------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/22624  | [ planner: not pruning column used by union scan condition (#21640)](https://github.com/pingcap/tidb/pull/22624)   |   | 3 | 0h    |
| tidb/22504  | [*:Fix the fetchHotRegion bug that the count always zero](https://github.com/pingcap/tidb/pull/22504)              |   | 4 | 3d21h |
| tidb/22443  | [planner: fix panic while get part of partition key values](https://github.com/pingcap/tidb/pull/22443)            |   | 5 | 7d3h  |
| tidb/22452  | [planner: fix panic while get part of partition key values](https://github.com/pingcap/tidb/pull/22452)            |   | 5 | 7d0h  |
| tidb/22565  | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565) | Y | 5 | 0h    |
| tidb/22492  | [go.mod, statistics, planner: replace TIDB_STATS with STATS_EXTENDED](https://github.com/pingcap/tidb/pull/22492)  |   | 7 | 3d4h  |
| parser/1159 | [parser, ast: replace TIDB_STATS with STATS_EXTENDED](https://github.com/pingcap/parser/pull/1159)                 |   | 7 | 3d0h  |


</details> 


<details> 
  <summary>fzhedu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                   PR                                                   | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19845 | [expression:fix FORMAT compatibility issue #11206](https://github.com/pingcap/tidb/pull/19845)         | Y | 146d16h |
| tidb/20117 | [optimizer: fix issue on incorrect result of natural join](https://github.com/pingcap/tidb/pull/20117) | Y | 132d20h |


## Reviewed in Last 7 Days

|    REPO    |                                                         PR                                                         | C | D |  R   |
|------------|--------------------------------------------------------------------------------------------------------------------|---|---|------|
| tics/1379  | [Fix ExchangeSender: remove duplicated write stream operation](https://github.com/pingcap/tics/pull/1379)          |   | 4 | 3d0h |
| tidb/22428 | [unistore/cophandler: change the logic of mpp execution in unit test.](https://github.com/pingcap/tidb/pull/22428) |   | 7 | 7d4h |


</details> 


<details> 
  <summary>ichn-hu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                            PR                                                            | C | LASTED |
|------------|--------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/21853 | [expression: fix compatibility behaviors in time_format with MySQL (#21559)](https://github.com/pingcap/tidb/pull/21853) |   | 45d19h |
| tidb/22411 | [util/chunk: trigger disk spill for sort properly](https://github.com/pingcap/tidb/pull/22411)                           |   | 16d16h |


## Reviewed in Last 7 Days

|     REPO     |                                                             PR                                                              | C | D |   R   |
|--------------|-----------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| docs-cn/5378 | [Add incompatibility caused by deprecated features in mysql-compatibility.md](https://github.com/pingcap/docs-cn/pull/5378) |   | 7 | 2d21h |


</details> 


<details> 
  <summary>lzmhhh123's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                  PR                                                                   | C | LASTED  |
|--------------|---------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/14729   | [planner: fix constant propagation for PredicatePushDown](https://github.com/pingcap/tidb/pull/14729)                                 | Y | 355d18h |
| docs-cn/4913 | [explain: add indexes](https://github.com/pingcap/docs-cn/pull/4913)                                                                  |   | 76d18h  |
| tidb/17414   | [add curCost based join reorder algorithm](https://github.com/pingcap/tidb/pull/17414)                                                |   | 250d18h |
| tidb/19347   | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347)  |   | 163d23h |
| tidb/19698   | [*: update test cases to support new collation enabled by default](https://github.com/pingcap/tidb/pull/19698)                        |   | 151d22h |
| tidb/20044   | [expression: Add column nullability checking before "refine args"](https://github.com/pingcap/tidb/pull/20044)                        | Y | 137d7h  |
| tidb/20444   | [expression: add json_merge_patch](https://github.com/pingcap/tidb/pull/20444)                                                        |   | 109d21h |
| tidb/20465   | [expression: add uuidShortFunction](https://github.com/pingcap/tidb/pull/20465)                                                       |   | 108d19h |
| tidb/20505   | [*: Add metrics for oom-action and sql memory usage.](https://github.com/pingcap/tidb/pull/20505)                                     |   | 104d19h |
| tidb/20618   | [planner: fix update generated columns error](https://github.com/pingcap/tidb/pull/20618)                                             |   | 99d20h  |
| tidb/20642   | [executor: modify admin executors to support partitioned table with global index](https://github.com/pingcap/tidb/pull/20642)         |   | 97d15h  |
| tidb/20825   | [executor: add diagnosis rule to check Transparent Huge Pages(THP) enabled (#20611)](https://github.com/pingcap/tidb/pull/20825)      |   | 88d18h  |
| tidb/20903   | [planner: fix confused and unnecessary double-projection in plans.](https://github.com/pingcap/tidb/pull/20903)                       |   | 86d17h  |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                        |   | 80d17h  |
| tidb/21051   | [executor: change read slow-log file module to concurrent](https://github.com/pingcap/tidb/pull/21051)                                |   | 79d14h  |
| tidb/21195   | [brie: integrate lightning to suport IMPORT statement](https://github.com/pingcap/tidb/pull/21195)                                    |   | 69d23h  |
| tidb/21334   | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                         |   | 66d14h  |
| tidb/21347   | [session: make rollback work on global variables](https://github.com/pingcap/tidb/pull/21347)                                         |   | 65d19h  |
| tidb/21401   | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                    |   | 62d11h  |
| tidb/21444   | [planner: ignore anonymous index while tiflash replica is available](https://github.com/pingcap/tidb/pull/21444)                      |   | 60d12h  |
| tidb/21487   | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                         |   | 59d4h   |
| tidb/21641   | [executor: Fix pessimistic lock doesn't work on the partition table for subquery/joins](https://github.com/pingcap/tidb/pull/21641)   |   | 52d18h  |
| tidb/21651   | [planner: allow filter condition pushing down to IndexScan for prefix index](https://github.com/pingcap/tidb/pull/21651)              |   | 52d13h  |
| tidb/21680   | [planner: report error when ORDER BY conflicts with DISTINCT (#21286)](https://github.com/pingcap/tidb/pull/21680)                    |   | 51d16h  |
| tidb/21954   | [planner/cascades: add rule `PushSelDownApply`](https://github.com/pingcap/tidb/pull/21954)                                           |   | 39d23h  |
| tidb/22089   | [executor: fix signed cluster index behavior (#22085)](https://github.com/pingcap/tidb/pull/22089)                                    |   | 32d22h  |
| tidb/22126   | [*: add `sys` schema, `sys.SCHEMA_UNUSED_INDEXES` view and `sys.SCHEMA_INDEX_USAGE` view](https://github.com/pingcap/tidb/pull/22126) |   | 31d20h  |
| tidb/22149   | [session: set process info before building plan (#22101)](https://github.com/pingcap/tidb/pull/22149)                                 |   | 27d19h  |
| tidb/22188   | [planner: do not use indexMerge when the path only use a single index (#22168)](https://github.com/pingcap/tidb/pull/22188)           |   | 26d13h  |
| tidb/22361   | [table: fix insert into _tidb_rowid panic and rebase it if needed (#22062)](https://github.com/pingcap/tidb/pull/22361)               |   | 19d20h  |
| tidb/22372   | [executor: fix SelectForUpdate in decorrelated subquery under pessimistic mode](https://github.com/pingcap/tidb/pull/22372)           |   | 19d9h   |
| tidb/22426   | [expression: fix bugs in builtinfunction ArithmeticMinusInt logic](https://github.com/pingcap/tidb/pull/22426)                        |   | 13d16h  |
| tidb/22428   | [unistore/cophandler: change the logic of mpp execution in unit test.](https://github.com/pingcap/tidb/pull/22428)                    |   | 13d14h  |
| tidb/22430   | [*: refactor table.Table interface, clean up unnecessay methods](https://github.com/pingcap/tidb/pull/22430)                          |   | 12d23h  |
| tidb/22433   | [statistics: merge partition-level TopN to global-level TopN](https://github.com/pingcap/tidb/pull/22433)                             |   | 12d19h  |
| tidb/22463   | [executor: make memory tracker for aggregate more accurate.](https://github.com/pingcap/tidb/pull/22463)                              |   | 10d23h  |
| tidb/22478   | [planner, executor: fix query partition table with global unique index get wrong result](https://github.com/pingcap/tidb/pull/22478)  |   | 10d13h  |
| tidb/22554   | [planner: change the content of AnalyzeTableID to build global-stats](https://github.com/pingcap/tidb/pull/22554)                     |   | 4d22h   |
| tidb/22607   | [store/tikv: expose failpoints that are used externally](https://github.com/pingcap/tidb/pull/22607)                                  |   | 3d13h   |


## Reviewed in Last 7 Days

|    REPO    |                                                                 PR                                                                 | C | D |   R    |
|------------|------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/21553 | [table: fix zero date in different sqlmode (#20206)](https://github.com/pingcap/tidb/pull/21553)                                   | Y | 3 | 52d1h  |
| tidb/22568 | [*: do not report error for prepared stmt execution if tidb_snapshot is set](https://github.com/pingcap/tidb/pull/22568)           |   | 3 | 1d23h  |
| tidb/21877 | [planner: fix correlated aggregates which should be evaluated in outer query (#21431)](https://github.com/pingcap/tidb/pull/21877) |   | 4 | 40d22h |
| tidb/22580 | [expression: handle duration type infer in least and greatest (#22271)](https://github.com/pingcap/tidb/pull/22580)                |   | 5 | 0h     |
| tidb/22562 | [expression: fix type infer for tidb's builtin compare(least and great…](https://github.com/pingcap/tidb/pull/22562)               |   | 5 | 5h     |
| tidb/22577 | [expression, types: fix unexpected result from TIME() when fsp digits > 6 (#21652)](https://github.com/pingcap/tidb/pull/22577)    |   | 5 | 0h     |
| tidb/21850 | [expression: add implicit eval int and real for function dayname (#21806)](https://github.com/pingcap/tidb/pull/21850)             |   | 5 | 41d1h  |
| tidb/21881 | [expression, types: fix datetime and year comparison error (#20233)](https://github.com/pingcap/tidb/pull/21881)                   | Y | 6 | 39d2h  |
| tidb/21853 | [expression: fix compatibility behaviors in time_format with MySQL (#21559)](https://github.com/pingcap/tidb/pull/21853)           |   | 6 | 40d4h  |
| tidb/21870 | [types: report error for json object with key length >= 65536 (#21779)](https://github.com/pingcap/tidb/pull/21870)                |   | 6 | 39d4h  |
| tidb/21808 | [planner: fix the fail when we compare multi fields in the subquery (#21699)](https://github.com/pingcap/tidb/pull/21808)          |   | 6 | 42d0h  |
| tidb/22491 | [executor: skip null data in common handle during point-get (#22483)](https://github.com/pingcap/tidb/pull/22491)                  |   | 7 | 2d21h  |


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

|    REPO    |                                                                           PR                                                                           | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/16305 | [expression: separate signatures for `ModInt`](https://github.com/pingcap/tidb/pull/16305)                                                             | Y | 294d0h  |
| tidb/16967 | [executor: Refactor Shuffle and implement parallel sort (executor part)](https://github.com/pingcap/tidb/pull/16967)                                   | Y | 274d10h |
| tidb/17396 | [types: improve StrToDate performance](https://github.com/pingcap/tidb/pull/17396)                                                                     | Y | 251d10h |
| tidb/18882 | [planner, executor: add explain for `MetricSummaryTableExtractor`](https://github.com/pingcap/tidb/pull/18882)                                         | Y | 185d17h |
| tidb/19029 | [types: fix unexpected NOT_NULL flags](https://github.com/pingcap/tidb/pull/19029)                                                                     |   | 178d22h |
| tidb/19120 | [executor: Concurrently fetch chunks and insert them to a concurrent hash table in hash build](https://github.com/pingcap/tidb/pull/19120)             |   | 173d21h |
| tidb/19292 | [planner: suppport left join in join reorder](https://github.com/pingcap/tidb/pull/19292)                                                              |   | 165d17h |
| tidb/20011 | [statistics: fix incorrect total count used in index selectivity computation](https://github.com/pingcap/tidb/pull/20011)                              |   | 138d15h |
| tidb/20316 | [docs/design: add design doc for index usage information](https://github.com/pingcap/tidb/pull/20316)                                                  |   | 123d17h |
| tidb/20354 | [planner: rename relational operators (#14575)](https://github.com/pingcap/tidb/pull/20354)                                                            | Y | 116d5h  |
| tidb/20689 | [expression: make TIME function compatible with MySQL (#19158)](https://github.com/pingcap/tidb/pull/20689)                                            |   | 95d20h  |
| tidb/20708 | [*: separate auto_increment ID allocator from _tidb_rowid allocator](https://github.com/pingcap/tidb/pull/20708)                                       |   | 94d20h  |
| tidb/20972 | [expression: POC implementation of Vitess hashing algorithm.](https://github.com/pingcap/tidb/pull/20972)                                              |   | 82d1h   |
| tidb/21018 | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                         |   | 80d17h  |
| tidb/21149 | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                          |   | 73d14h  |
| tidb/21304 | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                               |   | 67d12h  |
| tidb/21318 | [planner, expression: use the range of column types to simplify expressions](https://github.com/pingcap/tidb/pull/21318)                               |   | 66d19h  |
| tidb/21359 | [*: add runtime stats for split region statement](https://github.com/pingcap/tidb/pull/21359)                                                          |   | 65d13h  |
| tidb/21401 | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                                     |   | 62d11h  |
| tidb/21424 | [sessionctx: move set variable to sysvar struct](https://github.com/pingcap/tidb/pull/21424)                                                           |   | 61d5h   |
| tidb/21476 | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                  |   | 59d15h  |
| tidb/21508 | [execution: fix dayofweek('0000-00-00') behavior](https://github.com/pingcap/tidb/pull/21508)                                                          |   | 58d10h  |
| tidb/21680 | [planner: report error when ORDER BY conflicts with DISTINCT (#21286)](https://github.com/pingcap/tidb/pull/21680)                                     |   | 51d16h  |
| tidb/21876 | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                            |   | 44d19h  |
| tidb/21887 | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                                        |   | 43d11h  |
| tidb/21930 | [planner: propagate NDV of column groups across plan nodes (#17854)](https://github.com/pingcap/tidb/pull/21930)                                       |   | 40d18h  |
| tidb/21977 | [expression: log functions that can not be pushed to cop](https://github.com/pingcap/tidb/pull/21977)                                                  |   | 39d16h  |
| tidb/22090 | [planner: push aggregation operators down to projection and union by default](https://github.com/pingcap/tidb/pull/22090)                              |   | 32d22h  |
| tidb/22104 | [executor: fix incompatible escape behaviors in `select into outfile` (#22100)](https://github.com/pingcap/tidb/pull/22104)                            |   | 32d16h  |
| tidb/22107 | [executor: avoid log duplicate index name in slow-log (#22057)](https://github.com/pingcap/tidb/pull/22107)                                            |   | 32d14h  |
| tidb/22146 | [executor: forbid SFU on view](https://github.com/pingcap/tidb/pull/22146)                                                                             |   | 27d21h  |
| tidb/22217 | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                                          |   | 25d17h  |
| tidb/22234 | [executor, planner: ON DUPLICATE UPDATE can refer to un-project col (#14412)](https://github.com/pingcap/tidb/pull/22234)                              |   | 25d15h  |
| tidb/22261 | [time: fix parse datetime won't truncate the reluctant string (#22232)](https://github.com/pingcap/tidb/pull/22261)                                    |   | 24d19h  |
| tidb/22307 | [ddl: fix update can see columns not public](https://github.com/pingcap/tidb/pull/22307)                                                               |   | 23d16h  |
| tidb/22342 | [session: fix two cases when updating bind info (#22338)](https://github.com/pingcap/tidb/pull/22342)                                                  |   | 20d18h  |
| tidb/22369 | [session: fix the duplicate binding case when updating bind info (#22367)](https://github.com/pingcap/tidb/pull/22369)                                 |   | 19d17h  |
| tidb/22374 | [expression: separated arithmeticIntDivideSig](https://github.com/pingcap/tidb/pull/22374)                                                             |   | 19d0h   |
| tidb/22406 | [executor: metrics slow query is divided into internal and general (#22350)](https://github.com/pingcap/tidb/pull/22406)                               |   | 16d19h  |
| tidb/22409 | [*: use CLUSTERED and NONCLUSTERED to control primary key type](https://github.com/pingcap/tidb/pull/22409)                                            |   | 16d18h  |
| tidb/22415 | [ddl: refactor placement package](https://github.com/pingcap/tidb/pull/22415)                                                                          |   | 15d17h  |
| tidb/22426 | [expression: fix bugs in builtinfunction ArithmeticMinusInt logic](https://github.com/pingcap/tidb/pull/22426)                                         |   | 13d16h  |
| tidb/22433 | [statistics: merge partition-level TopN to global-level TopN](https://github.com/pingcap/tidb/pull/22433)                                              |   | 12d19h  |
| tidb/22456 | [distsql, executor: disable cache during staleness transaction](https://github.com/pingcap/tidb/pull/22456)                                            |   | 11d15h  |
| tidb/22471 | [ddl, executor: fix creating unique index without partition column error when enable-global-index is true](https://github.com/pingcap/tidb/pull/22471) |   | 10d17h  |
| tidb/22489 | [infoschema: support query partition_id from infoschema.partitions (#22240)](https://github.com/pingcap/tidb/pull/22489)                               |   | 9d20h   |
| tidb/22490 | [infoschema: support query partition_id from infoschema.partitions (#22240)](https://github.com/pingcap/tidb/pull/22490)                               |   | 9d20h   |
| tidb/22507 | [types: fix the bug about the wrong query result for decimal type ](https://github.com/pingcap/tidb/pull/22507)                                        |   | 6d23h   |
| tidb/22541 | [expression: Support builtin function SOUNDEX](https://github.com/pingcap/tidb/pull/22541)                                                             |   | 5d9h    |
| tidb/22565 | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)                                     | Y | 4d17h   |
| tidb/22568 | [*: do not report error for prepared stmt execution if tidb_snapshot is set](https://github.com/pingcap/tidb/pull/22568)                               |   | 4d16h   |
| tidb/22608 | [statistics: fix bug when bootstrap version 2 stats](https://github.com/pingcap/tidb/pull/22608)                                                       |   | 3d7h    |


## Reviewed in Last 7 Days

|    REPO    |                                                                               PR                                                                                | C | D |   R    |
|------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/21553 | [table: fix zero date in different sqlmode (#20206)](https://github.com/pingcap/tidb/pull/21553)                                                                | Y | 3 | 52d1h  |
| tidb/22614 | [expression: incorporate unicode_ci into constant propagation (#19555)](https://github.com/pingcap/tidb/pull/22614)                                             | Y | 3 | 4h     |
| tidb/22615 | [expression: allow function coercibility derive to DERIVATIO ... (#19462)](https://github.com/pingcap/tidb/pull/22615)                                          | Y | 3 | 4h     |
| tidb/21137 | [executor: specially handle empty input for apply's outer child aggregate (#20544)](https://github.com/pingcap/tidb/pull/21137)                                 |   | 3 | 70d23h |
| tidb/22599 | [expression: fix incorrect collation when cast non-string type arg to string type (#19186)](https://github.com/pingcap/tidb/pull/22599)                         | Y | 4 | 1h     |
| tidb/22459 | [server: retry executing sql without tiflash after tiflash is down](https://github.com/pingcap/tidb/pull/22459)                                                 |   | 4 | 7d21h  |
| tidb/22602 | [expression: fix wrong collation and coercibility (#19169)](https://github.com/pingcap/tidb/pull/22602)                                                         | Y | 4 | 0h     |
| tidb/22529 | [*: refactor `CompilePattern` and `DoMatch` used by `like` (#20610)](https://github.com/pingcap/tidb/pull/22529)                                                |   | 4 | 2d3h   |
| tidb/22582 | [expression: support handle two collation cannot substituted to each other (#19036)](https://github.com/pingcap/tidb/pull/22582)                                | Y | 4 | 16h    |
| tidb/21408 | [statistics: fix a bug which causes panic when using the clustered index and the new collation (#21379)](https://github.com/pingcap/tidb/pull/21408)            |   | 4 | 58d0h  |
| tidb/22588 | [config: add config metrics (#22586)](https://github.com/pingcap/tidb/pull/22588)                                                                               |   | 4 | 2h     |
| tidb/20981 | [expression: fix wrong inferred type for sum and avg (#20926)](https://github.com/pingcap/tidb/pull/20981)                                                      |   | 4 | 78d0h  |
| tidb/21877 | [planner: fix correlated aggregates which should be evaluated in outer query (#21431)](https://github.com/pingcap/tidb/pull/21877)                              |   | 4 | 40d22h |
| tidb/21464 | [server: return results of ongoing queries when graceful shutdown (#19669)](https://github.com/pingcap/tidb/pull/21464)                                         |   | 5 | 55d7h  |
| tidb/22575 | [*: fix LIKE expressions with _ following % (#17418)](https://github.com/pingcap/tidb/pull/22575)                                                               | Y | 5 | 2h     |
| tidb/21964 | [planner: add plancodec id for all type TableScan/IndexScan. (#21935)](https://github.com/pingcap/tidb/pull/21964)                                              |   | 5 | 35d5h  |
| tidb/22554 | [planner: change the content of AnalyzeTableID to build global-stats](https://github.com/pingcap/tidb/pull/22554)                                               |   | 5 | 7h     |
| tidb/22564 | [statistics: add tests for version 2 stats and fix bug](https://github.com/pingcap/tidb/pull/22564)                                                             |   | 5 | 2h     |
| tidb/22561 | [metrics: add server info metric (#22556)](https://github.com/pingcap/tidb/pull/22561)                                                                          |   | 5 | 0h     |
| tidb/22560 | [metrics: add server info metric (#22556)](https://github.com/pingcap/tidb/pull/22560)                                                                          |   | 5 | 0h     |
| tidb/21425 | [planner: natural join not consider rowid and null eq not propagate (#21328)](https://github.com/pingcap/tidb/pull/21425)                                       |   | 5 | 56d2h  |
| tidb/21850 | [expression: add implicit eval int and real for function dayname (#21806)](https://github.com/pingcap/tidb/pull/21850)                                          |   | 5 | 40d23h |
| tidb/22174 | [expression, ddl: check the argument count for the generated column (#22154)](https://github.com/pingcap/tidb/pull/22174)                                       |   | 5 | 22d0h  |
| tidb/21976 | [planner: report error for invalid window specs which are not used (#21083)](https://github.com/pingcap/tidb/pull/21976)                                        |   | 5 | 34d18h |
| tidb/21676 | [expression: fix compatibility of extract day_time unit functions (#21601)](https://github.com/pingcap/tidb/pull/21676)                                         |   | 5 | 46d18h |
| tidb/22118 | [planner: check if columns count matches for batch point get in TryFastPlan (#22044)](https://github.com/pingcap/tidb/pull/22118)                               |   | 5 | 26d23h |
| tidb/21958 | [expression: fix comparing json with string (#21903)](https://github.com/pingcap/tidb/pull/21958)                                                               |   | 6 | 34d8h  |
| tidb/22537 | [planner: fix in-compatibility issues between TiDB agg and TiFlash agg in MPP mode](https://github.com/pingcap/tidb/pull/22537)                                 |   | 6 | 3h     |
| tidb/21665 | [executor: fix LEAD and LAG's default value can not adapt to field type (#20747)](https://github.com/pingcap/tidb/pull/21665)                                   |   | 6 | 46d5h  |
| tidb/21614 | [planner: do not propagate column eq with different column types (#21495)](https://github.com/pingcap/tidb/pull/21614)                                          |   | 6 | 48d0h  |
| tidb/21972 | [executor: throw error when prepared statement is execute, deallocate or prepare (#21962)](https://github.com/pingcap/tidb/pull/21972)                          |   | 6 | 34d2h  |
| tidb/21700 | [expression: fix incompatible result of `JSON_SEARCH()` (#20164)](https://github.com/pingcap/tidb/pull/21700)                                                   | Y | 6 | 43d4h  |
| tidb/21590 | [expression: fix compatibility behaviors in sec_to_time with MySQL  (#21555)](https://github.com/pingcap/tidb/pull/21590)                                       |   | 6 | 48d6h  |
| tidb/21957 | [planner: fix unknown columns in join using below agg (#21922)](https://github.com/pingcap/tidb/pull/21957)                                                     |   | 6 | 34d8h  |
| tidb/21714 | [planner: fix the coercibility of the cast function (#21705)](https://github.com/pingcap/tidb/pull/21714)                                                       |   | 6 | 43d2h  |
| tidb/21881 | [expression, types: fix datetime and year comparison error (#20233)](https://github.com/pingcap/tidb/pull/21881)                                                | Y | 6 | 39d2h  |
| tidb/21853 | [expression: fix compatibility behaviors in time_format with MySQL (#21559)](https://github.com/pingcap/tidb/pull/21853)                                        |   | 6 | 40d4h  |
| tidb/22428 | [unistore/cophandler: change the logic of mpp execution in unit test.](https://github.com/pingcap/tidb/pull/22428)                                              |   | 6 | 7d23h  |
| tidb/22119 | [executor: fix `update ignore` into not exists partition (#21984)](https://github.com/pingcap/tidb/pull/22119)                                                  |   | 6 | 26d7h  |
| tidb/22472 | [planner, statistics: build the global statistics for the partition table](https://github.com/pingcap/tidb/pull/22472)                                          |   | 6 | 5d2h   |
| tidb/22136 | [executor: improve the runtime stats of index lookup reader (#21982)](https://github.com/pingcap/tidb/pull/22136)                                               |   | 6 | 25d23h |
| tidb/21945 | [distsql: fix cop stats string display when there is only 1 rpc (#21901)](https://github.com/pingcap/tidb/pull/21945)                                           |   | 6 | 34d20h |
| tidb/22106 | [executor: avoid log duplicate index name in slow-log (#22057)](https://github.com/pingcap/tidb/pull/22106)                                                     |   | 6 | 26d20h |
| tidb/22148 | [session: set process info before building plan (#22101)](https://github.com/pingcap/tidb/pull/22148)                                                           |   | 6 | 22d1h  |
| tidb/22481 | [executor: fix load data in file get wrong result #20854 (#21895)](https://github.com/pingcap/tidb/pull/22481)                                                  |   | 6 | 4d18h  |
| tidb/22461 | [planner, executor, statistics: add tests for version 2 and fix bugs](https://github.com/pingcap/tidb/pull/22461)                                               |   | 6 | 5d14h  |
| tidb/21870 | [types: report error for json object with key length >= 65536 (#21779)](https://github.com/pingcap/tidb/pull/21870)                                             |   | 6 | 39d4h  |
| tidb/21810 | [expression: handle hybrid field types for where clause (#21724)](https://github.com/pingcap/tidb/pull/21810)                                                   |   | 6 | 42d0h  |
| tidb/21808 | [planner: fix the fail when we compare multi fields in the subquery (#21699)](https://github.com/pingcap/tidb/pull/21808)                                       |   | 6 | 42d0h  |
| tidb/21711 | [expression: Fix unexpected panic when using IF function. (#21132)](https://github.com/pingcap/tidb/pull/21711)                                                 |   | 6 | 42d23h |
| tidb/21697 | [planner: check for only_full_group_by in ORDER BY and HAVING (#21216)](https://github.com/pingcap/tidb/pull/21697)                                             |   | 6 | 43d1h  |
| tidb/21638 | [server: check LOAD DATA is into a base table (#20924)](https://github.com/pingcap/tidb/pull/21638)                                                             |   | 6 | 47d0h  |
| tidb/21604 | [expression, json: fix converting from string to decimal (#21592)](https://github.com/pingcap/tidb/pull/21604)                                                  |   | 6 | 47d22h |
| tidb/21550 | [planner : fix unsigned_decimal_col=-int_cnst access index (#21198)](https://github.com/pingcap/tidb/pull/21550)                                                |   | 6 | 49d1h  |
| tidb/21532 | [expression: set IsBooleanFlag for boolean scalar functions (#20706)](https://github.com/pingcap/tidb/pull/21532)                                               |   | 6 | 49d22h |
| tidb/21525 | [expression: fix compatibility behaviors in zero datetime with MySQL (#21220)](https://github.com/pingcap/tidb/pull/21525)                                      |   | 6 | 50d1h  |
| tidb/21504 | [planner: fix invalid convert type in between...and... (#19820)](https://github.com/pingcap/tidb/pull/21504)                                                    | Y | 6 | 52d21h |
| tidb/21471 | [session: fix ineffective EXPLAIN FOR CONNECTION statement (#21044)](https://github.com/pingcap/tidb/pull/21471)                                                |   | 6 | 53d23h |
| tidb/21874 | [expression:truncate decimal value instead of return error (#21691)](https://github.com/pingcap/tidb/pull/21874)                                                |   | 6 | 39d2h  |
| tidb/21936 | [expression: fix wrong type inferring for ceiling function. (#21920)](https://github.com/pingcap/tidb/pull/21936)                                               |   | 6 | 34d22h |
| tidb/21628 | [expression: change the round rule for approximate value to `round to nearest even`  (#21324)](https://github.com/pingcap/tidb/pull/21628)                      |   | 6 | 47d4h  |
| tidb/22527 | [session,executor: fix point get under @@tidb_snapshot (#22460)](https://github.com/pingcap/tidb/pull/22527)                                                    |   | 6 | 3h     |
| tidb/21971 | [executor: fix `insert ignore` into not exists partition (#21904)](https://github.com/pingcap/tidb/pull/21971)                                                  |   | 6 | 33d20h |
| tidb/19957 | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                                          | Y | 7 | 136d0h |
| tidb/22492 | [go.mod, statistics, planner: replace TIDB_STATS with STATS_EXTENDED](https://github.com/pingcap/tidb/pull/22492)                                               |   | 7 | 3d5h   |
| tidb/22110 | [config, session: promise the compatibility of oom-action when upgrading (#22102)](https://github.com/pingcap/tidb/pull/22110)                                  |   | 7 | 25d23h |
| tidb/22420 | [types: convert string to MySQL BIT correctly (#21310)](https://github.com/pingcap/tidb/pull/22420)                                                             |   | 7 | 7d7h   |
| tidb/22332 | [expression, executor: fix runtime panic in WEIGHT_STRING function when the length of binary is too large (#22251)](https://github.com/pingcap/tidb/pull/22332) |   | 7 | 14d8h  |
| tidb/22021 | [distsql: fix cop stats string display when there is only 1 rpc (#21901) (#21999)](https://github.com/pingcap/tidb/pull/22021)                                  |   | 7 | 31d9h  |
| tidb/22353 | [planner: do not cache prepared plan if optimization depends on mutable constant (#22349)](https://github.com/pingcap/tidb/pull/22353)                          |   | 7 | 13d8h  |
| tidb/22421 | [expression: handle duration type infer in least and greatest (#22271)](https://github.com/pingcap/tidb/pull/22421)                                             |   | 7 | 7d7h   |
| tidb/22518 | [planner: avoid potential panic when generating hints from joins (#22515)](https://github.com/pingcap/tidb/pull/22518)                                          |   | 7 | 1h     |
| tidb/22515 | [planner: avoid potential panic when generating hints from joins](https://github.com/pingcap/tidb/pull/22515)                                                   |   | 7 | 0h     |


</details> 


<details> 
  <summary>rebelice's detailed review report</summary> 

## To Be Reviewed

| REPO | PR | C | LASTED |
|------|----|---|--------|


## Reviewed in Last 7 Days

|    REPO    |                                                        PR                                                         | C | D | R  |
|------------|-------------------------------------------------------------------------------------------------------------------|---|---|----|
| tidb/22554 | [planner: change the content of AnalyzeTableID to build global-stats](https://github.com/pingcap/tidb/pull/22554) |   | 5 | 5h |


</details> 


<details> 
  <summary>time-and-fate's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                            PR                                             | C | LASTED |
|------------|-------------------------------------------------------------------------------------------|---|--------|
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877) |   | 87d17h |


## Reviewed in Last 7 Days

|    REPO    |                                                        PR                                                         | C | D |   R   |
|------------|-------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/22608 | [statistics: fix bug when bootstrap version 2 stats](https://github.com/pingcap/tidb/pull/22608)                  |   | 3 | 14h   |
| tidb/22461 | [planner, executor, statistics: add tests for version 2 and fix bugs](https://github.com/pingcap/tidb/pull/22461) |   | 3 | 8d15h |
| tidb/22457 | [statistics: add more tests about ver2-stats](https://github.com/pingcap/tidb/pull/22457)                         |   | 7 | 4d21h |


</details> 


<details> 
  <summary>winoros's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                 PR                                                                  | C | LASTED  |
|--------------|-------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/14424   | [expression: add nullable() method to check whether an expression can return null](https://github.com/pingcap/tidb/pull/14424)      |   | 388d17h |
| docs-cn/4669 | [sql-optimization: extended statistics documentation](https://github.com/pingcap/docs-cn/pull/4669)                                 |   | 110d17h |
| tidb/14831   | [planner/cascades: add implementationRule for IndexLookUpJoin](https://github.com/pingcap/tidb/pull/14831)                          |   | 348d17h |
| tidb/15090   | [planner/cascades: refine the row count estimation of TiKV layer Selection](https://github.com/pingcap/tidb/pull/15090)             |   | 334d18h |
| tidb/15157   | [planner/cascades: implement `HashCode` method for all the LogicalPlans](https://github.com/pingcap/tidb/pull/15157)                | Y | 332d14h |
| tidb/15426   | [planner/cascades: add transformation rule PushSelDownApply & refactor PushSelDownJoin](https://github.com/pingcap/tidb/pull/15426) |   | 320d16h |
| tidb/16967   | [executor: Refactor Shuffle and implement parallel sort (executor part)](https://github.com/pingcap/tidb/pull/16967)                | Y | 274d10h |
| tidb/17414   | [add curCost based join reorder algorithm](https://github.com/pingcap/tidb/pull/17414)                                              |   | 250d18h |
| tidb/17996   | [planner: push avg & distinct functions across join](https://github.com/pingcap/tidb/pull/17996)                                    | Y | 232d11h |
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                              | Y | 142d14h |
| tidb/20011   | [statistics: fix incorrect total count used in index selectivity computation](https://github.com/pingcap/tidb/pull/20011)           |   | 138d15h |
| tidb/20311   | [expression: fix overflow error when convert bit to int64 (#20266)](https://github.com/pingcap/tidb/pull/20311)                     |   | 123d21h |
| tidb/20765   | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                   |   | 90d17h  |
| tidb/20877   | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                           |   | 87d17h  |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                      |   | 80d17h  |
| tidb/21207   | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                     |   | 69d19h  |
| tidb/21476   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                               |   | 59d15h  |
| tidb/21487   | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                       |   | 59d4h   |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)         |   | 44d19h  |
| tidb/21930   | [planner: propagate NDV of column groups across plan nodes (#17854)](https://github.com/pingcap/tidb/pull/21930)                    |   | 40d18h  |
| tidb/22090   | [planner: push aggregation operators down to projection and union by default](https://github.com/pingcap/tidb/pull/22090)           |   | 32d22h  |
| tidb/22365   | [planner: check index valid while forUpdateRead (#22152)](https://github.com/pingcap/tidb/pull/22365)                               |   | 19d19h  |
| tidb/22409   | [*: use CLUSTERED and NONCLUSTERED to control primary key type](https://github.com/pingcap/tidb/pull/22409)                         |   | 16d18h  |
| tidb/22459   | [server: retry executing sql without tiflash after tiflash is down](https://github.com/pingcap/tidb/pull/22459)                     |   | 11d12h  |
| tidb/22489   | [infoschema: support query partition_id from infoschema.partitions (#22240)](https://github.com/pingcap/tidb/pull/22489)            |   | 9d20h   |
| tidb/22490   | [infoschema: support query partition_id from infoschema.partitions (#22240)](https://github.com/pingcap/tidb/pull/22490)            |   | 9d20h   |
| tidb/22504   | [*:Fix the fetchHotRegion bug that the count always zero](https://github.com/pingcap/tidb/pull/22504)                               |   | 7d19h   |
| tidb/22565   | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)                  | Y | 4d17h   |
| tidb/22603   | [statistics: merge the partition-level histograms to a global-level histogram](https://github.com/pingcap/tidb/pull/22603)          |   | 3d16h   |
| tidb/22624   | [ planner: not pruning column used by union scan condition (#21640)](https://github.com/pingcap/tidb/pull/22624)                    |   | 2d17h   |


## Reviewed in Last 7 Days

|    REPO    |                                                                          PR                                                                          | C | D |   R    |
|------------|------------------------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/22583 | [expression: do not rewrite like to eq when using unicode ci](https://github.com/pingcap/tidb/pull/22583)                                            |   | 3 | 1d15h  |
| tidb/21137 | [executor: specially handle empty input for apply's outer child aggregate (#20544)](https://github.com/pingcap/tidb/pull/21137)                      |   | 3 | 70d23h |
| tidb/22594 | [Revert "executor: fix load data in file get wrong result #20854 (#21895)"](https://github.com/pingcap/tidb/pull/22594)                              |   | 4 | 0h     |
| tidb/21408 | [statistics: fix a bug which causes panic when using the clustered index and the new collation (#21379)](https://github.com/pingcap/tidb/pull/21408) |   | 4 | 58d0h  |
| tidb/22580 | [expression: handle duration type infer in least and greatest (#22271)](https://github.com/pingcap/tidb/pull/22580)                                  |   | 5 | 0h     |
| tidb/22577 | [expression, types: fix unexpected result from TIME() when fsp digits > 6 (#21652)](https://github.com/pingcap/tidb/pull/22577)                      |   | 5 | 0h     |
| tidb/22562 | [expression: fix type infer for tidb's builtin compare(least and great…](https://github.com/pingcap/tidb/pull/22562)                                 |   | 5 | 5h     |
| tidb/22564 | [statistics: add tests for version 2 stats and fix bug](https://github.com/pingcap/tidb/pull/22564)                                                  |   | 5 | 1h     |
| tidb/22465 | [statistics: fix panic occurs when stats cache inconsistency](https://github.com/pingcap/tidb/pull/22465)                                            | Y | 6 | 5d4h   |
| tidb/22518 | [planner: avoid potential panic when generating hints from joins (#22515)](https://github.com/pingcap/tidb/pull/22518)                               |   | 7 | 1h     |
| tidb/22515 | [planner: avoid potential panic when generating hints from joins](https://github.com/pingcap/tidb/pull/22515)                                        |   | 7 | 1h     |
| tidb/22457 | [statistics: add more tests about ver2-stats](https://github.com/pingcap/tidb/pull/22457)                                                            |   | 7 | 4d18h  |


</details> 


<details> 
  <summary>wshwsh12's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                            PR                                                            | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/15462 | [executor: implement `graceHashJoin`](https://github.com/pingcap/tidb/pull/15462)                                        | Y | 319d17h |
| tidb/17996 | [planner: push avg & distinct functions across join](https://github.com/pingcap/tidb/pull/17996)                         | Y | 232d11h |
| tidb/19557 | [*: Integrate timeline tracing with TiKV](https://github.com/pingcap/tidb/pull/19557)                                    |   | 156d23h |
| tidb/19807 | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                  |   | 149d11h |
| tidb/19957 | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                   | Y | 142d14h |
| tidb/20044 | [expression: Add column nullability checking before "refine args"](https://github.com/pingcap/tidb/pull/20044)           | Y | 137d7h  |
| tidb/21381 | [*: optimize analyze cluster index table](https://github.com/pingcap/tidb/pull/21381)                                    |   | 62d17h  |
| tidb/21487 | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                            |   | 59d4h   |
| tidb/21839 | [planner/core: add 'split table using statistics' statement](https://github.com/pingcap/tidb/pull/21839)                 |   | 46d15h  |
| tidb/21887 | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                          |   | 43d11h  |
| tidb/22261 | [time: fix parse datetime won't truncate the reluctant string (#22232)](https://github.com/pingcap/tidb/pull/22261)      |   | 24d19h  |
| tidb/22269 | [executor: check storage.block-cache.capacity value](https://github.com/pingcap/tidb/pull/22269)                         |   | 24d17h  |
| tidb/22378 | [executor: vectorize hash aggregate](https://github.com/pingcap/tidb/pull/22378)                                         |   | 18d19h  |
| tidb/22382 | [*: add infoschema client errors](https://github.com/pingcap/tidb/pull/22382)                                            |   | 18d5h   |
| tidb/22406 | [executor: metrics slow query is divided into internal and general (#22350)](https://github.com/pingcap/tidb/pull/22406) |   | 16d19h  |
| tidb/22411 | [util/chunk: trigger disk spill for sort properly](https://github.com/pingcap/tidb/pull/22411)                           |   | 16d16h  |
| tidb/22426 | [expression: fix bugs in builtinfunction ArithmeticMinusInt logic](https://github.com/pingcap/tidb/pull/22426)           |   | 13d16h  |
| tidb/22475 | [*: fix parser error format when value overflow](https://github.com/pingcap/tidb/pull/22475)                             |   | 10d16h  |
| tidb/22547 | [*: introduce new API ParseWithParams (#22499)](https://github.com/pingcap/tidb/pull/22547)                              |   | 4d23h   |
| tidb/22548 | [*: introduce new API ParseWithParams (#22499)](https://github.com/pingcap/tidb/pull/22548)                              |   | 4d23h   |
| tidb/22549 | [*: introduce new API ParseWithParams (#22499)](https://github.com/pingcap/tidb/pull/22549)                              |   | 4d23h   |
| tidb/22617 | [metrics: fix wrong bucket name of coprocessor cache (#22454)](https://github.com/pingcap/tidb/pull/22617)               |   | 2d23h   |
| tidb/22618 | [metrics: fix wrong bucket name of coprocessor cache (#22454)](https://github.com/pingcap/tidb/pull/22618)               |   | 2d23h   |
| tidb/22628 | [executor: Improve max/min window function with deque-based sliding window](https://github.com/pingcap/tidb/pull/22628)  |   | 1d23h   |


## Reviewed in Last 7 Days

|    REPO    |                                                         PR                                                          | C | D |   R   |
|------------|---------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/22260 | [time: fix parse datetime won't truncate the reluctant string (#22232)](https://github.com/pingcap/tidb/pull/22260) |   | 3 | 22d3h |
| tidb/22454 | [metrics: fix wrong bucket name of coprocessor cache](https://github.com/pingcap/tidb/pull/22454)                   |   | 3 | 8d17h |


</details> 

