|   REVIEWER    |  O(C)   | 1  | 2  | 3  | 4 | 5 | 6 | 7 | T  |
|---------------|---------|----|----|----|---|---|---|---|----|
| Reminiscent   |  4 ( 0) |  1 |  1 |  0 | 0 | 0 | 0 | 0 |  2 |
| SunRunAway    | 46 ( 5) |  0 |  0 |  0 | 0 | 0 | 0 | 0 |  0 |
| XuHuaiyu      | 57 ( 6) |  5 |  0 |  0 | 0 | 0 | 0 | 0 |  5 |
| dyzsr         |  0 ( 0) |  0 |  0 |  0 | 0 | 0 | 0 | 0 |  0 |
| eurekaka      | 20 ( 5) |  3 |  0 |  2 | 0 | 0 | 1 | 0 |  6 |
| fzhedu        |  2 ( 2) |  0 |  0 |  1 | 0 | 0 | 0 | 0 |  1 |
| ichn-hu       |  2 ( 0) |  0 |  0 |  1 | 0 | 0 | 0 | 0 |  1 |
| lzmhhh123     | 42 ( 2) |  4 |  4 |  1 | 0 | 0 | 0 | 0 |  9 |
| nullnotnil    |  0 ( 0) |  0 |  0 |  0 | 0 | 0 | 0 | 0 |  0 |
| qw4990        | 55 ( 8) | 15 | 37 | 10 | 0 | 0 | 3 | 0 | 65 |
| rebelice      |  0 ( 0) |  1 |  0 |  0 | 0 | 0 | 0 | 0 |  1 |
| time-and-fate |  1 ( 0) |  0 |  1 |  1 | 0 | 0 | 0 | 0 |  2 |
| winoros       | 31 ( 5) |  4 |  1 |  3 | 0 | 0 | 1 | 0 |  9 |
| wshwsh12      | 24 ( 5) |  0 |  0 |  0 | 0 | 0 | 0 | 0 |  0 |


<details> 
  <summary>Reminiscent's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                   PR                                                                   | C | LASTED |
|------------|----------------------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/21137 | [executor: specially handle empty input for apply's outer child aggregate (#20544)](https://github.com/pingcap/tidb/pull/21137)        |   | 69d20h |
| tidb/21896 | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                             |   | 37d19h |
| tidb/22330 | [planner: check error when correlatedAggregateResolver leaves ast.Node (#22222)](https://github.com/pingcap/tidb/pull/22330)           |   | 16d23h |
| tidb/22354 | [planner: do not cache prepared plan if optimization depends on mutable constant (#22349)](https://github.com/pingcap/tidb/pull/22354) |   | 15d23h |


## Reviewed in Last 7 Days

|    REPO    |                                                             PR                                                              | C | D |   R    |
|------------|-----------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/21964 | [planner: add plancodec id for all type TableScan/IndexScan. (#21935)](https://github.com/pingcap/tidb/pull/21964)          |   | 1 | 34d23h |
| tidb/22531 | [expression: do not rewrite `like` to `=` if new collation is enabled (#21893)](https://github.com/pingcap/tidb/pull/22531) |   | 2 | 18h    |


</details> 


<details> 
  <summary>SunRunAway's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                      PR                                                                       | C | LASTED  |
|--------------|-----------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| docs-cn/4913 | [explain: add indexes](https://github.com/pingcap/docs-cn/pull/4913)                                                                          |   | 72d17h  |
| tidb/15370   | [planner,executor: Refactor Shuffle and implement parallel Sort](https://github.com/pingcap/tidb/pull/15370)                                  | Y | 319d19h |
| docs-cn/4933 | [explain: add joins](https://github.com/pingcap/docs-cn/pull/4933)                                                                            |   | 68d20h  |
| tidb/15462   | [executor: implement `graceHashJoin`](https://github.com/pingcap/tidb/pull/15462)                                                             | Y | 315d17h |
| tidb/16967   | [executor: Refactor Shuffle and implement parallel sort (executor part)](https://github.com/pingcap/tidb/pull/16967)                          | Y | 270d10h |
| tidb/17238   | [*: refactor table.Allocator to improve readability](https://github.com/pingcap/tidb/pull/17238)                                              |   | 257d18h |
| tidb/19120   | [executor: Concurrently fetch chunks and insert them to a concurrent hash table in hash build](https://github.com/pingcap/tidb/pull/19120)    |   | 169d21h |
| tidb/19178   | [executor: Refactor probe channel](https://github.com/pingcap/tidb/pull/19178)                                                                |   | 167d17h |
| tidb/19347   | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347)          |   | 159d23h |
| tidb/19807   | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                                       |   | 145d11h |
| tidb/19900   | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                                | Y | 140d18h |
| tidb/20140   | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                            |   | 127d22h |
| tidb/20220   | [*: new secondary index value format](https://github.com/pingcap/tidb/pull/20220)                                                             |   | 124d16h |
| tidb/20316   | [docs/design: add design doc for index usage information](https://github.com/pingcap/tidb/pull/20316)                                         |   | 119d17h |
| tidb/20335   | [planner, executor: enable inline projection for Selection](https://github.com/pingcap/tidb/pull/20335)                                       | Y | 116d18h |
| tidb/20360   | [planner: refine explain info for batch cop](https://github.com/pingcap/tidb/pull/20360)                                                      |   | 110d22h |
| tidb/20397   | [parser: replace ast.SelectLockInShareMode with ast.SelectLockForShare](https://github.com/pingcap/tidb/pull/20397)                           |   | 108d18h |
| tidb/20615   | [utils: Avoid panic when getting memory](https://github.com/pingcap/tidb/pull/20615)                                                          |   | 96d2h   |
| tidb/20689   | [expression: make TIME function compatible with MySQL (#19158)](https://github.com/pingcap/tidb/pull/20689)                                   |   | 91d20h  |
| tidb/20752   | [*: trace statsCache and preparePlanCache by Global memory tracker.](https://github.com/pingcap/tidb/pull/20752)                              |   | 86d22h  |
| tidb/20765   | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                             |   | 86d17h  |
| tidb/21137   | [executor: specially handle empty input for apply's outer child aggregate (#20544)](https://github.com/pingcap/tidb/pull/21137)               |   | 69d20h  |
| tidb/21207   | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                               |   | 65d19h  |
| tidb/21277   | [executor: fix split table with large integers](https://github.com/pingcap/tidb/pull/21277)                                                   |   | 63d20h  |
| tidb/21364   | [expression: Add test cases to cover the cases when invalid int value is casted as TIME (#18653)](https://github.com/pingcap/tidb/pull/21364) |   | 59d1h   |
| tidb/21381   | [*: optimize analyze cluster index table](https://github.com/pingcap/tidb/pull/21381)                                                         |   | 58d17h  |
| tidb/21386   | [expression: Disable cast decimal as string push down to TiFlash](https://github.com/pingcap/tidb/pull/21386)                                 |   | 58d16h  |
| tidb/21834   | [planner: enhanced index range calculation plan](https://github.com/pingcap/tidb/pull/21834)                                                  |   | 42d18h  |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                   |   | 40d19h  |
| tidb/21877   | [planner: fix correlated aggregates which should be evaluated in outer query (#21431)](https://github.com/pingcap/tidb/pull/21877)            |   | 40d19h  |
| tidb/21878   | [planner: do not push down lock to pointGet/bacthPointGet when selection exists](https://github.com/pingcap/tidb/pull/21878)                  |   | 40d18h  |
| tidb/21890   | [*: redact some error code, part(3/3) (#21866)](https://github.com/pingcap/tidb/pull/21890)                                                   |   | 38d15h  |
| tidb/21956   | [planner/preprocessor: disallow into-outfile clause in some place](https://github.com/pingcap/tidb/pull/21956)                                |   | 35d23h  |
| tidb/22026   | [expression: separated arithmeticPlusIntSig](https://github.com/pingcap/tidb/pull/22026)                                                      |   | 33d20h  |
| tidb/22043   | [planner, executor: enhance the limit pushdown rule.](https://github.com/pingcap/tidb/pull/22043)                                             |   | 31d10h  |
| tidb/22089   | [executor: fix signed cluster index behavior (#22085)](https://github.com/pingcap/tidb/pull/22089)                                            |   | 28d22h  |
| tidb/22104   | [executor: fix incompatible escape behaviors in `select into outfile` (#22100)](https://github.com/pingcap/tidb/pull/22104)                   |   | 28d16h  |
| tidb/22107   | [executor: avoid log duplicate index name in slow-log (#22057)](https://github.com/pingcap/tidb/pull/22107)                                   |   | 28d14h  |
| tidb/22114   | [test: fix globalkilltest (#21987)](https://github.com/pingcap/tidb/pull/22114)                                                               |   | 28d12h  |
| tidb/22120   | [executor: fix `update ignore` into not exists partition (#21984)](https://github.com/pingcap/tidb/pull/22120)                                |   | 27d22h  |
| tidb/22181   | [planner, expression: fix error when using IN combined with subquery (#22080)](https://github.com/pingcap/tidb/pull/22181)                    |   | 22d17h  |
| tidb/22217   | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                                 |   | 21d17h  |
| tidb/22330   | [planner: check error when correlatedAggregateResolver leaves ast.Node (#22222)](https://github.com/pingcap/tidb/pull/22330)                  |   | 16d23h  |
| tidb/22365   | [planner: check index valid while forUpdateRead (#22152)](https://github.com/pingcap/tidb/pull/22365)                                         |   | 15d19h  |
| tidb/22379   | [[experiment] executor: allow aggregation to spill disk when running out of memory quota](https://github.com/pingcap/tidb/pull/22379)         |   | 14d19h  |
| tidb/22485   | [docs: Add design doc for security enhanced mode](https://github.com/pingcap/tidb/pull/22485)                                                 |   | 6d5h    |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>XuHuaiyu's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                              PR                                                                              | C | LASTED  |
|--------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19292   | [planner: suppport left join in join reorder](https://github.com/pingcap/tidb/pull/19292)                                                                    |   | 161d17h |
| docs-cn/5323 | [Update parameter type description](https://github.com/pingcap/docs-cn/pull/5323)                                                                            |   | 9d19h   |
| tidb/19900   | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                                               | Y | 140d18h |
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                                       | Y | 138d14h |
| tidb/20040   | [planner, expression: take NullFlag into consideration when optimize the `int non-const` <cmp > `non-int const`](https://github.com/pingcap/tidb/pull/20040) | Y | 133d14h |
| tidb/20140   | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                                           |   | 127d22h |
| tidb/20311   | [expression: fix overflow error when convert bit to int64 (#20266)](https://github.com/pingcap/tidb/pull/20311)                                              |   | 119d21h |
| tidb/20350   | [executor: support read global indexes in IndexMergeReader and index join](https://github.com/pingcap/tidb/pull/20350)                                       | Y | 113d14h |
| tidb/20505   | [*: Add metrics for oom-action and sql memory usage.](https://github.com/pingcap/tidb/pull/20505)                                                            |   | 100d19h |
| tidb/20576   | [*: fix stats feedback after tableReader handle multiple ranges](https://github.com/pingcap/tidb/pull/20576)                                                 |   | 98d13h  |
| tidb/20613   | [executor: fix issue of hash join fetch time inaccurate](https://github.com/pingcap/tidb/pull/20613)                                                         |   | 96d13h  |
| tidb/20752   | [*: trace statsCache and preparePlanCache by Global memory tracker.](https://github.com/pingcap/tidb/pull/20752)                                             |   | 86d22h  |
| tidb/20790   | [collation: add pinyin collation for chinese charset support](https://github.com/pingcap/tidb/pull/20790)                                                    |   | 85d21h  |
| tidb/20793   | [planner, executor: enable inline projection for Apply](https://github.com/pingcap/tidb/pull/20793)                                                          |   | 85d20h  |
| tidb/20905   | [planner: fix statement-optimize not work in `TryFastPlan`](https://github.com/pingcap/tidb/pull/20905)                                                      |   | 82d17h  |
| tidb/20972   | [expression: POC implementation of Vitess hashing algorithm.](https://github.com/pingcap/tidb/pull/20972)                                                    |   | 78d1h   |
| tidb/21064   | [planner, executor: fix cast not check error](https://github.com/pingcap/tidb/pull/21064)                                                                    |   | 73d8h   |
| tidb/21149   | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                                |   | 69d14h  |
| tidb/21228   | [executor: return the result immediately when combining LIMIT row_count with DISTINCT](https://github.com/pingcap/tidb/pull/21228)                           |   | 65d13h  |
| tidb/21304   | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                                     |   | 63d12h  |
| tidb/21334   | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                                                |   | 62d14h  |
| tidb/21340   | [executor: initialize expensive query handler on domain creation](https://github.com/pingcap/tidb/pull/21340)                                                |   | 61d23h  |
| tidb/21476   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                        |   | 55d15h  |
| tidb/21536   | [executor: add slow-log file meta cache to avoid repeat read file meta information](https://github.com/pingcap/tidb/pull/21536)                              |   | 51d15h  |
| tidb/21564   | [ddl: fix Incorrect behavior of NO_ZERO_DATE when altering table](https://github.com/pingcap/tidb/pull/21564)                                                |   | 50d15h  |
| tidb/21626   | [test: convert test to benchmard test to make ci stable (#21616)](https://github.com/pingcap/tidb/pull/21626)                                                |   | 48d23h  |
| tidb/21680   | [planner: report error when ORDER BY conflicts with DISTINCT (#21286)](https://github.com/pingcap/tidb/pull/21680)                                           |   | 47d16h  |
| tidb/21839   | [planner/core: add 'split table using statistics' statement](https://github.com/pingcap/tidb/pull/21839)                                                     |   | 42d15h  |
| tidb/21853   | [expression: fix compatibility behaviors in time_format with MySQL (#21559)](https://github.com/pingcap/tidb/pull/21853)                                     |   | 41d19h  |
| tidb/21877   | [planner: fix correlated aggregates which should be evaluated in outer query (#21431)](https://github.com/pingcap/tidb/pull/21877)                           |   | 40d19h  |
| tidb/21896   | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                                                   |   | 37d19h  |
| tidb/22014   | [executor: fix unstable test Issue16696 (#22009)](https://github.com/pingcap/tidb/pull/22014)                                                                |   | 34d17h  |
| tidb/22107   | [executor: avoid log duplicate index name in slow-log (#22057)](https://github.com/pingcap/tidb/pull/22107)                                                  |   | 28d14h  |
| tidb/22120   | [executor: fix `update ignore` into not exists partition (#21984)](https://github.com/pingcap/tidb/pull/22120)                                               |   | 27d22h  |
| tidb/22131   | [privilege: remove leading and trailing space when create user and role](https://github.com/pingcap/tidb/pull/22131)                                         |   | 27d19h  |
| tidb/22142   | [store: trace `loadRegion` to see the PD region cache loading (#22092)](https://github.com/pingcap/tidb/pull/22142)                                          |   | 24d0h   |
| tidb/22149   | [session: set process info before building plan (#22101)](https://github.com/pingcap/tidb/pull/22149)                                                        |   | 23d19h  |
| tidb/22163   | [expression: separated arithmeticMinusIntSig](https://github.com/pingcap/tidb/pull/22163)                                                                    |   | 23d13h  |
| tidb/22186   | [executor: fix select into outfile with year type column has no data (#22175)](https://github.com/pingcap/tidb/pull/22186)                                   |   | 22d16h  |
| tidb/22294   | [planner, table: optimize the list partition pruner for range query](https://github.com/pingcap/tidb/pull/22294)                                             |   | 19d20h  |
| tidb/22307   | [ddl: fix update can see columns not public](https://github.com/pingcap/tidb/pull/22307)                                                                     |   | 19d16h  |
| tidb/22381   | [planner: check schema stale for plan cache when forUpdateRead](https://github.com/pingcap/tidb/pull/22381)                                                  |   | 14d14h  |
| tidb/22403   | [planner: reorder inner joins simplified from outer joins (#22392)](https://github.com/pingcap/tidb/pull/22403)                                              |   | 12d22h  |
| tidb/22407   | [types: fix return err when decimal from string value](https://github.com/pingcap/tidb/pull/22407)                                                           |   | 12d19h  |
| tidb/22418   | [expression: Optimize builtinArithmeticModRealSig and builtinGreatestDecimalSig using MergeNull method](https://github.com/pingcap/tidb/pull/22418)          |   | 10d0h   |
| tidb/22426   | [expression: fix bugs in builtinfunction ArithmeticMinusInt logic](https://github.com/pingcap/tidb/pull/22426)                                               |   | 9d16h   |
| tidb/22432   | [types,execute: fix errcode return like mysql when inserting incorrect int value ](https://github.com/pingcap/tidb/pull/22432)                               |   | 8d21h   |
| tidb/22454   | [metrics: fix wrong bucket name of coprocessor cache](https://github.com/pingcap/tidb/pull/22454)                                                            |   | 7d16h   |
| tidb/22463   | [executor: make memory tracker for aggregate more accurate.](https://github.com/pingcap/tidb/pull/22463)                                                     |   | 6d23h   |
| tidb/22472   | [planner, statistics: build the global statistics for the partition table](https://github.com/pingcap/tidb/pull/22472)                                       |   | 6d17h   |
| tidb/22491   | [executor: skip null data in common handle during point-get (#22483)](https://github.com/pingcap/tidb/pull/22491)                                            |   | 5d19h   |
| tidb/22507   | [types: fix the bug about the wrong query result for decimal type ](https://github.com/pingcap/tidb/pull/22507)                                              |   | 2d23h   |
| tidb/22528   | [brie: update br pkg to latest master](https://github.com/pingcap/tidb/pull/22528)                                                                           |   | 1d22h   |
| tidb/22529   | [*: refactor `CompilePattern` and `DoMatch` used by `like` (#20610)](https://github.com/pingcap/tidb/pull/22529)                                             |   | 1d20h   |
| tidb/22546   | [*: refactor Execute/ExecuteInternal usage for single-statement](https://github.com/pingcap/tidb/pull/22546)                                                 |   | 23h     |
| tidb/22558   | [charset: implement utf8_unicode_ci and utf8mb4_unicode_ci collation (#18776)](https://github.com/pingcap/tidb/pull/22558)                                   | Y | 20h     |
| tidb/22582   | [expression: support handle two collation cannot substituted to each other (#19036)](https://github.com/pingcap/tidb/pull/22582)                             | Y | 11h     |


## Reviewed in Last 7 Days

|     REPO     |                                                               PR                                                                | C | D |   R    |
|--------------|---------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/22577   | [expression, types: fix unexpected result from TIME() when fsp digits > 6 (#21652)](https://github.com/pingcap/tidb/pull/22577) |   | 1 | 0h     |
| docs/4729    | [tidb: Refine the doc for global kill](https://github.com/pingcap/docs/pull/4729)                                               |   | 1 | 1h     |
| tidb/21676   | [expression: fix compatibility of extract day_time unit functions (#21601)](https://github.com/pingcap/tidb/pull/21676)         |   | 1 | 46d18h |
| tidb/21700   | [expression: fix incompatible result of `JSON_SEARCH()` (#20164)](https://github.com/pingcap/tidb/pull/21700)                   | Y | 1 | 43d19h |
| docs-cn/5386 | [tidb: add doc for global kill](https://github.com/pingcap/docs-cn/pull/5386)                                                   |   | 1 | 1d20h  |


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
| tidb/14729 | [planner: fix constant propagation for PredicatePushDown](https://github.com/pingcap/tidb/pull/14729)                                  | Y | 351d18h |
| tidb/14831 | [planner/cascades: add implementationRule for IndexLookUpJoin](https://github.com/pingcap/tidb/pull/14831)                             |   | 344d17h |
| tidb/15090 | [planner/cascades: refine the row count estimation of TiKV layer Selection](https://github.com/pingcap/tidb/pull/15090)                |   | 330d18h |
| tidb/15157 | [planner/cascades: implement `HashCode` method for all the LogicalPlans](https://github.com/pingcap/tidb/pull/15157)                   | Y | 328d14h |
| tidb/15335 | [planner/cascades: add transformation rule PullAggregationUpApply & EliminateMaxOneRow](https://github.com/pingcap/tidb/pull/15335)    |   | 321d17h |
| tidb/15370 | [planner,executor: Refactor Shuffle and implement parallel Sort](https://github.com/pingcap/tidb/pull/15370)                           | Y | 319d19h |
| tidb/17276 | [planner/cascades: add rule InjectProjectionBelowSort](https://github.com/pingcap/tidb/pull/17276)                                     | Y | 254d9h  |
| tidb/18882 | [planner, executor: add explain for `MetricSummaryTableExtractor`](https://github.com/pingcap/tidb/pull/18882)                         | Y | 181d17h |
| tidb/19347 | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347)   |   | 159d23h |
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                              |   | 83d17h  |
| tidb/21444 | [planner: ignore anonymous index while tiflash replica is available](https://github.com/pingcap/tidb/pull/21444)                       |   | 56d12h  |
| tidb/21680 | [planner: report error when ORDER BY conflicts with DISTINCT (#21286)](https://github.com/pingcap/tidb/pull/21680)                     |   | 47d16h  |
| tidb/21994 | [range: fix overflow value access index ](https://github.com/pingcap/tidb/pull/21994)                                                  |   | 34d23h  |
| tidb/22330 | [planner: check error when correlatedAggregateResolver leaves ast.Node (#22222)](https://github.com/pingcap/tidb/pull/22330)           |   | 16d23h  |
| tidb/22342 | [session: fix two cases when updating bind info (#22338)](https://github.com/pingcap/tidb/pull/22342)                                  |   | 16d18h  |
| tidb/22354 | [planner: do not cache prepared plan if optimization depends on mutable constant (#22349)](https://github.com/pingcap/tidb/pull/22354) |   | 15d23h  |
| tidb/22369 | [session: fix the duplicate binding case when updating bind info (#22367)](https://github.com/pingcap/tidb/pull/22369)                 |   | 15d17h  |
| tidb/22416 | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                                      |   | 11d11h  |
| tidb/22504 | [*:Fix the fetchHotRegion bug that the count always zero](https://github.com/pingcap/tidb/pull/22504)                                  |   | 3d19h   |
| tidb/22559 | [planner: split test data from test cases in cbo_test.go](https://github.com/pingcap/tidb/pull/22559)                                  |   | 19h     |


## Reviewed in Last 7 Days

|    REPO     |                                                         PR                                                         | C | D |  R   |
|-------------|--------------------------------------------------------------------------------------------------------------------|---|---|------|
| tidb/22443  | [planner: fix panic while get part of partition key values](https://github.com/pingcap/tidb/pull/22443)            |   | 1 | 7d3h |
| tidb/22452  | [planner: fix panic while get part of partition key values](https://github.com/pingcap/tidb/pull/22452)            |   | 1 | 7d0h |
| tidb/22565  | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565) | Y | 1 | 0h   |
| tidb/22492  | [go.mod, statistics, planner: replace TIDB_STATS with STATS_EXTENDED](https://github.com/pingcap/tidb/pull/22492)  |   | 3 | 3d4h |
| parser/1159 | [parser, ast: replace TIDB_STATS with STATS_EXTENDED](https://github.com/pingcap/parser/pull/1159)                 |   | 3 | 3d0h |
| tidb/22465  | [statistics: fix panic occurs when stats cache inconsistency](https://github.com/pingcap/tidb/pull/22465)          | Y | 6 | 22h  |


</details> 


<details> 
  <summary>fzhedu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                   PR                                                   | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19845 | [expression:fix FORMAT compatibility issue #11206](https://github.com/pingcap/tidb/pull/19845)         | Y | 142d16h |
| tidb/20117 | [optimizer: fix issue on incorrect result of natural join](https://github.com/pingcap/tidb/pull/20117) | Y | 128d20h |


## Reviewed in Last 7 Days

|    REPO    |                                                         PR                                                         | C | D |  R   |
|------------|--------------------------------------------------------------------------------------------------------------------|---|---|------|
| tidb/22428 | [unistore/cophandler: change the logic of mpp execution in unit test.](https://github.com/pingcap/tidb/pull/22428) |   | 3 | 7d4h |


</details> 


<details> 
  <summary>ichn-hu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                            PR                                                            | C | LASTED |
|------------|--------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/21853 | [expression: fix compatibility behaviors in time_format with MySQL (#21559)](https://github.com/pingcap/tidb/pull/21853) |   | 41d19h |
| tidb/22411 | [util/chunk: trigger disk spill for sort properly](https://github.com/pingcap/tidb/pull/22411)                           |   | 12d16h |


## Reviewed in Last 7 Days

|     REPO     |                                                             PR                                                              | C | D |   R   |
|--------------|-----------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| docs-cn/5378 | [Add incompatibility caused by deprecated features in mysql-compatibility.md](https://github.com/pingcap/docs-cn/pull/5378) |   | 3 | 2d21h |


</details> 


<details> 
  <summary>lzmhhh123's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                  PR                                                                   | C | LASTED  |
|--------------|---------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/14729   | [planner: fix constant propagation for PredicatePushDown](https://github.com/pingcap/tidb/pull/14729)                                 | Y | 351d18h |
| docs-cn/4913 | [explain: add indexes](https://github.com/pingcap/docs-cn/pull/4913)                                                                  |   | 72d17h  |
| tidb/17414   | [add curCost based join reorder algorithm](https://github.com/pingcap/tidb/pull/17414)                                                |   | 246d18h |
| tidb/19347   | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347)  |   | 159d23h |
| tidb/19698   | [*: update test cases to support new collation enabled by default](https://github.com/pingcap/tidb/pull/19698)                        |   | 147d22h |
| tidb/20044   | [expression: Add column nullability checking before "refine args"](https://github.com/pingcap/tidb/pull/20044)                        | Y | 133d7h  |
| tidb/20444   | [expression: add json_merge_patch](https://github.com/pingcap/tidb/pull/20444)                                                        |   | 105d21h |
| tidb/20465   | [expression: add uuidShortFunction](https://github.com/pingcap/tidb/pull/20465)                                                       |   | 104d19h |
| tidb/20505   | [*: Add metrics for oom-action and sql memory usage.](https://github.com/pingcap/tidb/pull/20505)                                     |   | 100d19h |
| tidb/20618   | [planner: fix update generated columns error](https://github.com/pingcap/tidb/pull/20618)                                             |   | 95d20h  |
| tidb/20642   | [executor: modify admin executors to support partitioned table with global index](https://github.com/pingcap/tidb/pull/20642)         |   | 93d15h  |
| tidb/20825   | [executor: add diagnosis rule to check Transparent Huge Pages(THP) enabled (#20611)](https://github.com/pingcap/tidb/pull/20825)      |   | 84d18h  |
| tidb/20903   | [planner: fix confused and unnecessary double-projection in plans.](https://github.com/pingcap/tidb/pull/20903)                       |   | 82d17h  |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                        |   | 76d17h  |
| tidb/21051   | [executor: change read slow-log file module to concurrent](https://github.com/pingcap/tidb/pull/21051)                                |   | 75d14h  |
| tidb/21137   | [executor: specially handle empty input for apply's outer child aggregate (#20544)](https://github.com/pingcap/tidb/pull/21137)       |   | 69d20h  |
| tidb/21195   | [brie: integrate lightning to suport IMPORT statement](https://github.com/pingcap/tidb/pull/21195)                                    |   | 65d22h  |
| tidb/21334   | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                         |   | 62d14h  |
| tidb/21347   | [session: make rollback work on global variables](https://github.com/pingcap/tidb/pull/21347)                                         |   | 61d19h  |
| tidb/21401   | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                    |   | 58d11h  |
| tidb/21444   | [planner: ignore anonymous index while tiflash replica is available](https://github.com/pingcap/tidb/pull/21444)                      |   | 56d12h  |
| tidb/21487   | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                         |   | 55d4h   |
| tidb/21641   | [executor: Fix pessimistic lock doesn't work on the partition table for subquery/joins](https://github.com/pingcap/tidb/pull/21641)   |   | 48d18h  |
| tidb/21651   | [planner: allow filter condition pushing down to IndexScan for prefix index](https://github.com/pingcap/tidb/pull/21651)              |   | 48d13h  |
| tidb/21680   | [planner: report error when ORDER BY conflicts with DISTINCT (#21286)](https://github.com/pingcap/tidb/pull/21680)                    |   | 47d16h  |
| tidb/21877   | [planner: fix correlated aggregates which should be evaluated in outer query (#21431)](https://github.com/pingcap/tidb/pull/21877)    |   | 40d19h  |
| tidb/21954   | [planner/cascades: add rule `PushSelDownApply`](https://github.com/pingcap/tidb/pull/21954)                                           |   | 35d23h  |
| tidb/22089   | [executor: fix signed cluster index behavior (#22085)](https://github.com/pingcap/tidb/pull/22089)                                    |   | 28d22h  |
| tidb/22126   | [*: add `sys` schema, `sys.SCHEMA_UNUSED_INDEXES` view and `sys.SCHEMA_INDEX_USAGE` view](https://github.com/pingcap/tidb/pull/22126) |   | 27d20h  |
| tidb/22130   | [planner: join reorder should not change the order of output columns (#16852)](https://github.com/pingcap/tidb/pull/22130)            |   | 27d19h  |
| tidb/22149   | [session: set process info before building plan (#22101)](https://github.com/pingcap/tidb/pull/22149)                                 |   | 23d19h  |
| tidb/22188   | [planner: do not use indexMerge when the path only use a single index (#22168)](https://github.com/pingcap/tidb/pull/22188)           |   | 22d13h  |
| tidb/22361   | [table: fix insert into _tidb_rowid panic and rebase it if needed (#22062)](https://github.com/pingcap/tidb/pull/22361)               |   | 15d20h  |
| tidb/22372   | [executor: fix SelectForUpdate in decorrelated subquery under pessimistic mode](https://github.com/pingcap/tidb/pull/22372)           |   | 15d9h   |
| tidb/22426   | [expression: fix bugs in builtinfunction ArithmeticMinusInt logic](https://github.com/pingcap/tidb/pull/22426)                        |   | 9d16h   |
| tidb/22428   | [unistore/cophandler: change the logic of mpp execution in unit test.](https://github.com/pingcap/tidb/pull/22428)                    |   | 9d14h   |
| tidb/22430   | [*: refactor table.Table interface, clean up unnecessay methods](https://github.com/pingcap/tidb/pull/22430)                          |   | 8d23h   |
| tidb/22433   | [statistics: merge partition-level TopN to global-level TopN](https://github.com/pingcap/tidb/pull/22433)                             |   | 8d19h   |
| tidb/22463   | [executor: make memory tracker for aggregate more accurate.](https://github.com/pingcap/tidb/pull/22463)                              |   | 6d23h   |
| tidb/22478   | [planner, executor: fix query partition table with global unique index get wrong result](https://github.com/pingcap/tidb/pull/22478)  |   | 6d13h   |
| tidb/22554   | [planner: change the content of AnalyzeTableID to build global-stats](https://github.com/pingcap/tidb/pull/22554)                     |   | 22h     |
| tidb/22568   | [*: do not report error for prepared stmt execution if tidb_snapshot is set](https://github.com/pingcap/tidb/pull/22568)              |   | 16h     |


## Reviewed in Last 7 Days

|    REPO    |                                                               PR                                                                | C | D |   R   |
|------------|---------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/22580 | [expression: handle duration type infer in least and greatest (#22271)](https://github.com/pingcap/tidb/pull/22580)             |   | 1 | 0h    |
| tidb/22562 | [expression: fix type infer for tidb's builtin compare(least and great…](https://github.com/pingcap/tidb/pull/22562)            |   | 1 | 5h    |
| tidb/22577 | [expression, types: fix unexpected result from TIME() when fsp digits > 6 (#21652)](https://github.com/pingcap/tidb/pull/22577) |   | 1 | 0h    |
| tidb/21850 | [expression: add implicit eval int and real for function dayname (#21806)](https://github.com/pingcap/tidb/pull/21850)          |   | 1 | 41d1h |
| tidb/21881 | [expression, types: fix datetime and year comparison error (#20233)](https://github.com/pingcap/tidb/pull/21881)                | Y | 2 | 39d2h |
| tidb/21853 | [expression: fix compatibility behaviors in time_format with MySQL (#21559)](https://github.com/pingcap/tidb/pull/21853)        |   | 2 | 40d4h |
| tidb/21870 | [types: report error for json object with key length >= 65536 (#21779)](https://github.com/pingcap/tidb/pull/21870)             |   | 2 | 39d4h |
| tidb/21808 | [planner: fix the fail when we compare multi fields in the subquery (#21699)](https://github.com/pingcap/tidb/pull/21808)       |   | 2 | 42d0h |
| tidb/22491 | [executor: skip null data in common handle during point-get (#22483)](https://github.com/pingcap/tidb/pull/22491)               |   | 3 | 2d21h |


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
| tidb/16305 | [expression: separate signatures for `ModInt`](https://github.com/pingcap/tidb/pull/16305)                                                             | Y | 289d23h |
| tidb/16967 | [executor: Refactor Shuffle and implement parallel sort (executor part)](https://github.com/pingcap/tidb/pull/16967)                                   | Y | 270d10h |
| tidb/17396 | [types: improve StrToDate performance](https://github.com/pingcap/tidb/pull/17396)                                                                     | Y | 247d10h |
| tidb/18882 | [planner, executor: add explain for `MetricSummaryTableExtractor`](https://github.com/pingcap/tidb/pull/18882)                                         | Y | 181d17h |
| tidb/19029 | [types: fix unexpected NOT_NULL flags](https://github.com/pingcap/tidb/pull/19029)                                                                     |   | 174d22h |
| tidb/19120 | [executor: Concurrently fetch chunks and insert them to a concurrent hash table in hash build](https://github.com/pingcap/tidb/pull/19120)             |   | 169d21h |
| tidb/19292 | [planner: suppport left join in join reorder](https://github.com/pingcap/tidb/pull/19292)                                                              |   | 161d17h |
| tidb/20011 | [statistics: fix incorrect total count used in index selectivity computation](https://github.com/pingcap/tidb/pull/20011)                              |   | 134d15h |
| tidb/20316 | [docs/design: add design doc for index usage information](https://github.com/pingcap/tidb/pull/20316)                                                  |   | 119d17h |
| tidb/20354 | [planner: rename relational operators (#14575)](https://github.com/pingcap/tidb/pull/20354)                                                            | Y | 112d5h  |
| tidb/20689 | [expression: make TIME function compatible with MySQL (#19158)](https://github.com/pingcap/tidb/pull/20689)                                            |   | 91d20h  |
| tidb/20708 | [*: separate auto_increment ID allocator from _tidb_rowid allocator](https://github.com/pingcap/tidb/pull/20708)                                       |   | 90d20h  |
| tidb/20972 | [expression: POC implementation of Vitess hashing algorithm.](https://github.com/pingcap/tidb/pull/20972)                                              |   | 78d1h   |
| tidb/21018 | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                         |   | 76d17h  |
| tidb/21137 | [executor: specially handle empty input for apply's outer child aggregate (#20544)](https://github.com/pingcap/tidb/pull/21137)                        |   | 69d20h  |
| tidb/21149 | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                          |   | 69d14h  |
| tidb/21304 | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                               |   | 63d12h  |
| tidb/21318 | [planner, expression: use the range of column types to simplify expressions](https://github.com/pingcap/tidb/pull/21318)                               |   | 62d19h  |
| tidb/21359 | [*: add runtime stats for split region statement](https://github.com/pingcap/tidb/pull/21359)                                                          |   | 61d13h  |
| tidb/21401 | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                                     |   | 58d11h  |
| tidb/21408 | [statistics: fix a bug which causes panic when using the clustered index and the new collation (#21379)](https://github.com/pingcap/tidb/pull/21408)   |   | 57d20h  |
| tidb/21424 | [sessionctx: move set variable to sysvar struct](https://github.com/pingcap/tidb/pull/21424)                                                           |   | 57d5h   |
| tidb/21476 | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                  |   | 55d15h  |
| tidb/21508 | [execution: fix dayofweek('0000-00-00') behavior](https://github.com/pingcap/tidb/pull/21508)                                                          |   | 54d10h  |
| tidb/21680 | [planner: report error when ORDER BY conflicts with DISTINCT (#21286)](https://github.com/pingcap/tidb/pull/21680)                                     |   | 47d16h  |
| tidb/21876 | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                            |   | 40d19h  |
| tidb/21887 | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                                        |   | 39d11h  |
| tidb/21930 | [planner: propagate NDV of column groups across plan nodes (#17854)](https://github.com/pingcap/tidb/pull/21930)                                       |   | 36d18h  |
| tidb/21977 | [expression: log functions that can not be pushed to cop](https://github.com/pingcap/tidb/pull/21977)                                                  |   | 35d16h  |
| tidb/22090 | [planner: push aggregation operators down to projection and union by default](https://github.com/pingcap/tidb/pull/22090)                              |   | 28d22h  |
| tidb/22104 | [executor: fix incompatible escape behaviors in `select into outfile` (#22100)](https://github.com/pingcap/tidb/pull/22104)                            |   | 28d16h  |
| tidb/22107 | [executor: avoid log duplicate index name in slow-log (#22057)](https://github.com/pingcap/tidb/pull/22107)                                            |   | 28d14h  |
| tidb/22146 | [executor: forbid SFU on view](https://github.com/pingcap/tidb/pull/22146)                                                                             |   | 23d21h  |
| tidb/22217 | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                                          |   | 21d17h  |
| tidb/22234 | [executor, planner: ON DUPLICATE UPDATE can refer to un-project col (#14412)](https://github.com/pingcap/tidb/pull/22234)                              |   | 21d15h  |
| tidb/22261 | [time: fix parse datetime won't truncate the reluctant string (#22232)](https://github.com/pingcap/tidb/pull/22261)                                    |   | 20d19h  |
| tidb/22307 | [ddl: fix update can see columns not public](https://github.com/pingcap/tidb/pull/22307)                                                               |   | 19d16h  |
| tidb/22342 | [session: fix two cases when updating bind info (#22338)](https://github.com/pingcap/tidb/pull/22342)                                                  |   | 16d18h  |
| tidb/22369 | [session: fix the duplicate binding case when updating bind info (#22367)](https://github.com/pingcap/tidb/pull/22369)                                 |   | 15d17h  |
| tidb/22374 | [expression: separated arithmeticIntDivideSig](https://github.com/pingcap/tidb/pull/22374)                                                             |   | 15d0h   |
| tidb/22406 | [executor: metrics slow query is divided into internal and general (#22350)](https://github.com/pingcap/tidb/pull/22406)                               |   | 12d19h  |
| tidb/22409 | [*: use CLUSTERED and NONCLUSTERED to control primary key type](https://github.com/pingcap/tidb/pull/22409)                                            |   | 12d18h  |
| tidb/22415 | [ddl: refactor placement package](https://github.com/pingcap/tidb/pull/22415)                                                                          |   | 11d17h  |
| tidb/22426 | [expression: fix bugs in builtinfunction ArithmeticMinusInt logic](https://github.com/pingcap/tidb/pull/22426)                                         |   | 9d16h   |
| tidb/22433 | [statistics: merge partition-level TopN to global-level TopN](https://github.com/pingcap/tidb/pull/22433)                                              |   | 8d19h   |
| tidb/22456 | [distsql, executor: disable cache during staleness transaction](https://github.com/pingcap/tidb/pull/22456)                                            |   | 7d15h   |
| tidb/22471 | [ddl, executor: fix creating unique index without partition column error when enable-global-index is true](https://github.com/pingcap/tidb/pull/22471) |   | 6d17h   |
| tidb/22489 | [infoschema: support query partition_id from infoschema.partitions (#22240)](https://github.com/pingcap/tidb/pull/22489)                               |   | 5d20h   |
| tidb/22490 | [infoschema: support query partition_id from infoschema.partitions (#22240)](https://github.com/pingcap/tidb/pull/22490)                               |   | 5d20h   |
| tidb/22507 | [types: fix the bug about the wrong query result for decimal type ](https://github.com/pingcap/tidb/pull/22507)                                        |   | 2d23h   |
| tidb/22541 | [expression: Support builtin function SOUNDEX](https://github.com/pingcap/tidb/pull/22541)                                                             |   | 1d9h    |
| tidb/22565 | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)                                     | Y | 17h     |
| tidb/22568 | [*: do not report error for prepared stmt execution if tidb_snapshot is set](https://github.com/pingcap/tidb/pull/22568)                               |   | 16h     |
| tidb/22581 | [expression: unicode_ci support when infer collation and charset information (#19142)](https://github.com/pingcap/tidb/pull/22581)                     | Y | 11h     |
| tidb/22582 | [expression: support handle two collation cannot substituted to each other (#19036)](https://github.com/pingcap/tidb/pull/22582)                       | Y | 11h     |


## Reviewed in Last 7 Days

|    REPO    |                                                                               PR                                                                                | C | D |   R    |
|------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/21464 | [server: return results of ongoing queries when graceful shutdown (#19669)](https://github.com/pingcap/tidb/pull/21464)                                         |   | 1 | 55d7h  |
| tidb/22575 | [*: fix LIKE expressions with _ following % (#17418)](https://github.com/pingcap/tidb/pull/22575)                                                               | Y | 1 | 2h     |
| tidb/21964 | [planner: add plancodec id for all type TableScan/IndexScan. (#21935)](https://github.com/pingcap/tidb/pull/21964)                                              |   | 1 | 35d5h  |
| tidb/22554 | [planner: change the content of AnalyzeTableID to build global-stats](https://github.com/pingcap/tidb/pull/22554)                                               |   | 1 | 7h     |
| tidb/22459 | [server: retry executing sql without tiflash after tiflash is down](https://github.com/pingcap/tidb/pull/22459)                                                 |   | 1 | 6d22h  |
| tidb/22564 | [statistics: add tests for version 2 stats and fix bug](https://github.com/pingcap/tidb/pull/22564)                                                             |   | 1 | 2h     |
| tidb/22561 | [metrics: add server info metric (#22556)](https://github.com/pingcap/tidb/pull/22561)                                                                          |   | 1 | 0h     |
| tidb/22560 | [metrics: add server info metric (#22556)](https://github.com/pingcap/tidb/pull/22560)                                                                          |   | 1 | 0h     |
| tidb/21425 | [planner: natural join not consider rowid and null eq not propagate (#21328)](https://github.com/pingcap/tidb/pull/21425)                                       |   | 1 | 56d2h  |
| tidb/21850 | [expression: add implicit eval int and real for function dayname (#21806)](https://github.com/pingcap/tidb/pull/21850)                                          |   | 1 | 40d23h |
| tidb/21553 | [table: fix zero date in different sqlmode (#20206)](https://github.com/pingcap/tidb/pull/21553)                                                                | Y | 1 | 49d22h |
| tidb/22174 | [expression, ddl: check the argument count for the generated column (#22154)](https://github.com/pingcap/tidb/pull/22174)                                       |   | 1 | 22d0h  |
| tidb/21976 | [planner: report error for invalid window specs which are not used (#21083)](https://github.com/pingcap/tidb/pull/21976)                                        |   | 1 | 34d18h |
| tidb/21676 | [expression: fix compatibility of extract day_time unit functions (#21601)](https://github.com/pingcap/tidb/pull/21676)                                         |   | 1 | 46d18h |
| tidb/22118 | [planner: check if columns count matches for batch point get in TryFastPlan (#22044)](https://github.com/pingcap/tidb/pull/22118)                               |   | 1 | 26d23h |
| tidb/21958 | [expression: fix comparing json with string (#21903)](https://github.com/pingcap/tidb/pull/21958)                                                               |   | 2 | 34d8h  |
| tidb/22537 | [planner: fix in-compatibility issues between TiDB agg and TiFlash agg in MPP mode](https://github.com/pingcap/tidb/pull/22537)                                 |   | 2 | 3h     |
| tidb/21665 | [executor: fix LEAD and LAG's default value can not adapt to field type (#20747)](https://github.com/pingcap/tidb/pull/21665)                                   |   | 2 | 46d5h  |
| tidb/21614 | [planner: do not propagate column eq with different column types (#21495)](https://github.com/pingcap/tidb/pull/21614)                                          |   | 2 | 48d0h  |
| tidb/21972 | [executor: throw error when prepared statement is execute, deallocate or prepare (#21962)](https://github.com/pingcap/tidb/pull/21972)                          |   | 2 | 34d2h  |
| tidb/21700 | [expression: fix incompatible result of `JSON_SEARCH()` (#20164)](https://github.com/pingcap/tidb/pull/21700)                                                   | Y | 2 | 43d4h  |
| tidb/21590 | [expression: fix compatibility behaviors in sec_to_time with MySQL  (#21555)](https://github.com/pingcap/tidb/pull/21590)                                       |   | 2 | 48d6h  |
| tidb/21957 | [planner: fix unknown columns in join using below agg (#21922)](https://github.com/pingcap/tidb/pull/21957)                                                     |   | 2 | 34d8h  |
| tidb/21714 | [planner: fix the coercibility of the cast function (#21705)](https://github.com/pingcap/tidb/pull/21714)                                                       |   | 2 | 43d2h  |
| tidb/21881 | [expression, types: fix datetime and year comparison error (#20233)](https://github.com/pingcap/tidb/pull/21881)                                                | Y | 2 | 39d2h  |
| tidb/21853 | [expression: fix compatibility behaviors in time_format with MySQL (#21559)](https://github.com/pingcap/tidb/pull/21853)                                        |   | 2 | 40d4h  |
| tidb/22428 | [unistore/cophandler: change the logic of mpp execution in unit test.](https://github.com/pingcap/tidb/pull/22428)                                              |   | 2 | 7d23h  |
| tidb/22119 | [executor: fix `update ignore` into not exists partition (#21984)](https://github.com/pingcap/tidb/pull/22119)                                                  |   | 2 | 26d7h  |
| tidb/22472 | [planner, statistics: build the global statistics for the partition table](https://github.com/pingcap/tidb/pull/22472)                                          |   | 2 | 5d2h   |
| tidb/22136 | [executor: improve the runtime stats of index lookup reader (#21982)](https://github.com/pingcap/tidb/pull/22136)                                               |   | 2 | 25d23h |
| tidb/21945 | [distsql: fix cop stats string display when there is only 1 rpc (#21901)](https://github.com/pingcap/tidb/pull/21945)                                           |   | 2 | 34d20h |
| tidb/22106 | [executor: avoid log duplicate index name in slow-log (#22057)](https://github.com/pingcap/tidb/pull/22106)                                                     |   | 2 | 26d20h |
| tidb/22148 | [session: set process info before building plan (#22101)](https://github.com/pingcap/tidb/pull/22148)                                                           |   | 2 | 22d1h  |
| tidb/22481 | [executor: fix load data in file get wrong result #20854 (#21895)](https://github.com/pingcap/tidb/pull/22481)                                                  |   | 2 | 4d18h  |
| tidb/22461 | [planner, executor, statistics: add tests for version 2 and fix bugs](https://github.com/pingcap/tidb/pull/22461)                                               |   | 2 | 5d14h  |
| tidb/21870 | [types: report error for json object with key length >= 65536 (#21779)](https://github.com/pingcap/tidb/pull/21870)                                             |   | 2 | 39d4h  |
| tidb/21810 | [expression: handle hybrid field types for where clause (#21724)](https://github.com/pingcap/tidb/pull/21810)                                                   |   | 2 | 42d0h  |
| tidb/21808 | [planner: fix the fail when we compare multi fields in the subquery (#21699)](https://github.com/pingcap/tidb/pull/21808)                                       |   | 2 | 42d0h  |
| tidb/21711 | [expression: Fix unexpected panic when using IF function. (#21132)](https://github.com/pingcap/tidb/pull/21711)                                                 |   | 2 | 42d23h |
| tidb/21697 | [planner: check for only_full_group_by in ORDER BY and HAVING (#21216)](https://github.com/pingcap/tidb/pull/21697)                                             |   | 2 | 43d1h  |
| tidb/21638 | [server: check LOAD DATA is into a base table (#20924)](https://github.com/pingcap/tidb/pull/21638)                                                             |   | 2 | 47d0h  |
| tidb/21604 | [expression, json: fix converting from string to decimal (#21592)](https://github.com/pingcap/tidb/pull/21604)                                                  |   | 2 | 47d22h |
| tidb/21550 | [planner : fix unsigned_decimal_col=-int_cnst access index (#21198)](https://github.com/pingcap/tidb/pull/21550)                                                |   | 2 | 49d1h  |
| tidb/21532 | [expression: set IsBooleanFlag for boolean scalar functions (#20706)](https://github.com/pingcap/tidb/pull/21532)                                               |   | 2 | 49d22h |
| tidb/21525 | [expression: fix compatibility behaviors in zero datetime with MySQL (#21220)](https://github.com/pingcap/tidb/pull/21525)                                      |   | 2 | 50d1h  |
| tidb/21504 | [planner: fix invalid convert type in between...and... (#19820)](https://github.com/pingcap/tidb/pull/21504)                                                    | Y | 2 | 52d21h |
| tidb/21471 | [session: fix ineffective EXPLAIN FOR CONNECTION statement (#21044)](https://github.com/pingcap/tidb/pull/21471)                                                |   | 2 | 53d23h |
| tidb/21874 | [expression:truncate decimal value instead of return error (#21691)](https://github.com/pingcap/tidb/pull/21874)                                                |   | 2 | 39d2h  |
| tidb/21936 | [expression: fix wrong type inferring for ceiling function. (#21920)](https://github.com/pingcap/tidb/pull/21936)                                               |   | 2 | 34d22h |
| tidb/21628 | [expression: change the round rule for approximate value to `round to nearest even`  (#21324)](https://github.com/pingcap/tidb/pull/21628)                      |   | 2 | 47d4h  |
| tidb/22527 | [session,executor: fix point get under @@tidb_snapshot (#22460)](https://github.com/pingcap/tidb/pull/22527)                                                    |   | 2 | 3h     |
| tidb/21971 | [executor: fix `insert ignore` into not exists partition (#21904)](https://github.com/pingcap/tidb/pull/21971)                                                  |   | 2 | 33d20h |
| tidb/19957 | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                                          | Y | 3 | 136d0h |
| tidb/22492 | [go.mod, statistics, planner: replace TIDB_STATS with STATS_EXTENDED](https://github.com/pingcap/tidb/pull/22492)                                               |   | 3 | 3d5h   |
| tidb/22110 | [config, session: promise the compatibility of oom-action when upgrading (#22102)](https://github.com/pingcap/tidb/pull/22110)                                  |   | 3 | 25d23h |
| tidb/22420 | [types: convert string to MySQL BIT correctly (#21310)](https://github.com/pingcap/tidb/pull/22420)                                                             |   | 3 | 7d7h   |
| tidb/22332 | [expression, executor: fix runtime panic in WEIGHT_STRING function when the length of binary is too large (#22251)](https://github.com/pingcap/tidb/pull/22332) |   | 3 | 14d8h  |
| tidb/22021 | [distsql: fix cop stats string display when there is only 1 rpc (#21901) (#21999)](https://github.com/pingcap/tidb/pull/22021)                                  |   | 3 | 31d9h  |
| tidb/22353 | [planner: do not cache prepared plan if optimization depends on mutable constant (#22349)](https://github.com/pingcap/tidb/pull/22353)                          |   | 3 | 13d8h  |
| tidb/22421 | [expression: handle duration type infer in least and greatest (#22271)](https://github.com/pingcap/tidb/pull/22421)                                             |   | 3 | 7d7h   |
| tidb/22518 | [planner: avoid potential panic when generating hints from joins (#22515)](https://github.com/pingcap/tidb/pull/22518)                                          |   | 3 | 1h     |
| tidb/22515 | [planner: avoid potential panic when generating hints from joins](https://github.com/pingcap/tidb/pull/22515)                                                   |   | 3 | 0h     |
| tidb/22479 | [statistics: fix inaccurate row count estimation caused by cross validation using invalid column stats](https://github.com/pingcap/tidb/pull/22479)             |   | 6 | 19h    |
| tidb/22240 | [infoschema: support query partition_id from infoschema.partitions](https://github.com/pingcap/tidb/pull/22240)                                                 |   | 6 | 15d16h |
| tidb/21842 | [planner: Shuffle hash agg](https://github.com/pingcap/tidb/pull/21842)                                                                                         |   | 6 | 36d11h |


</details> 


<details> 
  <summary>rebelice's detailed review report</summary> 

## To Be Reviewed

| REPO | PR | C | LASTED |
|------|----|---|--------|


## Reviewed in Last 7 Days

|    REPO    |                                                        PR                                                         | C | D | R  |
|------------|-------------------------------------------------------------------------------------------------------------------|---|---|----|
| tidb/22554 | [planner: change the content of AnalyzeTableID to build global-stats](https://github.com/pingcap/tidb/pull/22554) |   | 1 | 5h |


</details> 


<details> 
  <summary>time-and-fate's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                            PR                                             | C | LASTED |
|------------|-------------------------------------------------------------------------------------------|---|--------|
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877) |   | 83d17h |


## Reviewed in Last 7 Days

|    REPO    |                                                        PR                                                         | C | D |   R   |
|------------|-------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/22461 | [planner, executor, statistics: add tests for version 2 and fix bugs](https://github.com/pingcap/tidb/pull/22461) |   | 2 | 5d17h |
| tidb/22457 | [statistics: add more tests about ver2-stats](https://github.com/pingcap/tidb/pull/22457)                         |   | 3 | 4d21h |


</details> 


<details> 
  <summary>winoros's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                          PR                                                                          | C | LASTED  |
|--------------|------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/14424   | [expression: add nullable() method to check whether an expression can return null](https://github.com/pingcap/tidb/pull/14424)                       |   | 384d17h |
| docs-cn/4669 | [sql-optimization: extended statistics documentation](https://github.com/pingcap/docs-cn/pull/4669)                                                  |   | 106d17h |
| tidb/14831   | [planner/cascades: add implementationRule for IndexLookUpJoin](https://github.com/pingcap/tidb/pull/14831)                                           |   | 344d17h |
| tidb/15090   | [planner/cascades: refine the row count estimation of TiKV layer Selection](https://github.com/pingcap/tidb/pull/15090)                              |   | 330d18h |
| tidb/15157   | [planner/cascades: implement `HashCode` method for all the LogicalPlans](https://github.com/pingcap/tidb/pull/15157)                                 | Y | 328d14h |
| tidb/15426   | [planner/cascades: add transformation rule PushSelDownApply & refactor PushSelDownJoin](https://github.com/pingcap/tidb/pull/15426)                  |   | 316d16h |
| tidb/16967   | [executor: Refactor Shuffle and implement parallel sort (executor part)](https://github.com/pingcap/tidb/pull/16967)                                 | Y | 270d10h |
| tidb/17414   | [add curCost based join reorder algorithm](https://github.com/pingcap/tidb/pull/17414)                                                               |   | 246d18h |
| tidb/17996   | [planner: push avg & distinct functions across join](https://github.com/pingcap/tidb/pull/17996)                                                     | Y | 228d11h |
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                               | Y | 138d14h |
| tidb/20011   | [statistics: fix incorrect total count used in index selectivity computation](https://github.com/pingcap/tidb/pull/20011)                            |   | 134d15h |
| tidb/20311   | [expression: fix overflow error when convert bit to int64 (#20266)](https://github.com/pingcap/tidb/pull/20311)                                      |   | 119d21h |
| tidb/20765   | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                                    |   | 86d17h  |
| tidb/20877   | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                                            |   | 83d17h  |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                       |   | 76d17h  |
| tidb/21137   | [executor: specially handle empty input for apply's outer child aggregate (#20544)](https://github.com/pingcap/tidb/pull/21137)                      |   | 69d20h  |
| tidb/21207   | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                                      |   | 65d19h  |
| tidb/21408   | [statistics: fix a bug which causes panic when using the clustered index and the new collation (#21379)](https://github.com/pingcap/tidb/pull/21408) |   | 57d20h  |
| tidb/21476   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                |   | 55d15h  |
| tidb/21487   | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                                        |   | 55d4h   |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                          |   | 40d19h  |
| tidb/21877   | [planner: fix correlated aggregates which should be evaluated in outer query (#21431)](https://github.com/pingcap/tidb/pull/21877)                   |   | 40d19h  |
| tidb/21930   | [planner: propagate NDV of column groups across plan nodes (#17854)](https://github.com/pingcap/tidb/pull/21930)                                     |   | 36d18h  |
| tidb/22090   | [planner: push aggregation operators down to projection and union by default](https://github.com/pingcap/tidb/pull/22090)                            |   | 28d22h  |
| tidb/22365   | [planner: check index valid while forUpdateRead (#22152)](https://github.com/pingcap/tidb/pull/22365)                                                |   | 15d19h  |
| tidb/22409   | [*: use CLUSTERED and NONCLUSTERED to control primary key type](https://github.com/pingcap/tidb/pull/22409)                                          |   | 12d18h  |
| tidb/22459   | [server: retry executing sql without tiflash after tiflash is down](https://github.com/pingcap/tidb/pull/22459)                                      |   | 7d12h   |
| tidb/22489   | [infoschema: support query partition_id from infoschema.partitions (#22240)](https://github.com/pingcap/tidb/pull/22489)                             |   | 5d20h   |
| tidb/22490   | [infoschema: support query partition_id from infoschema.partitions (#22240)](https://github.com/pingcap/tidb/pull/22490)                             |   | 5d20h   |
| tidb/22504   | [*:Fix the fetchHotRegion bug that the count always zero](https://github.com/pingcap/tidb/pull/22504)                                                |   | 3d19h   |
| tidb/22565   | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)                                   | Y | 17h     |


## Reviewed in Last 7 Days

|    REPO    |                                                                         PR                                                                          | C | D |   R   |
|------------|-----------------------------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/22580 | [expression: handle duration type infer in least and greatest (#22271)](https://github.com/pingcap/tidb/pull/22580)                                 |   | 1 | 0h    |
| tidb/22577 | [expression, types: fix unexpected result from TIME() when fsp digits > 6 (#21652)](https://github.com/pingcap/tidb/pull/22577)                     |   | 1 | 0h    |
| tidb/22562 | [expression: fix type infer for tidb's builtin compare(least and great…](https://github.com/pingcap/tidb/pull/22562)                                |   | 1 | 5h    |
| tidb/22564 | [statistics: add tests for version 2 stats and fix bug](https://github.com/pingcap/tidb/pull/22564)                                                 |   | 1 | 1h    |
| tidb/22465 | [statistics: fix panic occurs when stats cache inconsistency](https://github.com/pingcap/tidb/pull/22465)                                           | Y | 2 | 5d4h  |
| tidb/22518 | [planner: avoid potential panic when generating hints from joins (#22515)](https://github.com/pingcap/tidb/pull/22518)                              |   | 3 | 1h    |
| tidb/22515 | [planner: avoid potential panic when generating hints from joins](https://github.com/pingcap/tidb/pull/22515)                                       |   | 3 | 1h    |
| tidb/22457 | [statistics: add more tests about ver2-stats](https://github.com/pingcap/tidb/pull/22457)                                                           |   | 3 | 4d18h |
| tidb/22479 | [statistics: fix inaccurate row count estimation caused by cross validation using invalid column stats](https://github.com/pingcap/tidb/pull/22479) |   | 6 | 19h   |


</details> 


<details> 
  <summary>wshwsh12's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                 PR                                                                 | C | LASTED  |
|------------|------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/15462 | [executor: implement `graceHashJoin`](https://github.com/pingcap/tidb/pull/15462)                                                  | Y | 315d17h |
| tidb/17996 | [planner: push avg & distinct functions across join](https://github.com/pingcap/tidb/pull/17996)                                   | Y | 228d11h |
| tidb/19557 | [*: Integrate timeline tracing with TiKV](https://github.com/pingcap/tidb/pull/19557)                                              |   | 152d23h |
| tidb/19807 | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                            |   | 145d11h |
| tidb/19957 | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                             | Y | 138d14h |
| tidb/20044 | [expression: Add column nullability checking before "refine args"](https://github.com/pingcap/tidb/pull/20044)                     | Y | 133d7h  |
| tidb/21381 | [*: optimize analyze cluster index table](https://github.com/pingcap/tidb/pull/21381)                                              |   | 58d17h  |
| tidb/21487 | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                      |   | 55d4h   |
| tidb/21839 | [planner/core: add 'split table using statistics' statement](https://github.com/pingcap/tidb/pull/21839)                           |   | 42d15h  |
| tidb/21887 | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                    |   | 39d11h  |
| tidb/22269 | [executor: check storage.block-cache.capacity value](https://github.com/pingcap/tidb/pull/22269)                                   |   | 20d17h  |
| tidb/22378 | [executor: vectorize hash aggregate](https://github.com/pingcap/tidb/pull/22378)                                                   |   | 14d19h  |
| tidb/22382 | [*: add infoschema client errors](https://github.com/pingcap/tidb/pull/22382)                                                      |   | 14d5h   |
| tidb/22406 | [executor: metrics slow query is divided into internal and general (#22350)](https://github.com/pingcap/tidb/pull/22406)           |   | 12d19h  |
| tidb/22411 | [util/chunk: trigger disk spill for sort properly](https://github.com/pingcap/tidb/pull/22411)                                     |   | 12d16h  |
| tidb/22426 | [expression: fix bugs in builtinfunction ArithmeticMinusInt logic](https://github.com/pingcap/tidb/pull/22426)                     |   | 9d16h   |
| tidb/22454 | [metrics: fix wrong bucket name of coprocessor cache](https://github.com/pingcap/tidb/pull/22454)                                  |   | 7d16h   |
| tidb/22475 | [*: fix parser error format when value overflow](https://github.com/pingcap/tidb/pull/22475)                                       |   | 6d16h   |
| tidb/22547 | [*: introduce new API ParseWithParams (#22499)](https://github.com/pingcap/tidb/pull/22547)                                        |   | 23h     |
| tidb/22548 | [*: introduce new API ParseWithParams (#22499)](https://github.com/pingcap/tidb/pull/22548)                                        |   | 23h     |
| tidb/22549 | [*: introduce new API ParseWithParams (#22499)](https://github.com/pingcap/tidb/pull/22549)                                        |   | 23h     |
| tidb/22570 | [server: add token usage gauge for tidb (#22511)](https://github.com/pingcap/tidb/pull/22570)                                      |   | 16h     |
| tidb/22581 | [expression: unicode_ci support when infer collation and charset information (#19142)](https://github.com/pingcap/tidb/pull/22581) | Y | 11h     |
| tidb/22585 | [executor: fix load-data result when field term be the prefix of line term](https://github.com/pingcap/tidb/pull/22585)            |   | 0h      |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 

