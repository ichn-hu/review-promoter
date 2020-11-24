|   REVIEWER    |  O(C)   | 1  | 2 | 3 | 4 | 5 | 6 | 7  | T  |
|---------------|---------|----|---|---|---|---|---|----|----|
| Reminiscent   |  4 ( 0) |  2 | 0 | 0 | 1 | 0 | 1 |  0 |  4 |
| SunRunAway    | 49 ( 5) |  2 | 0 | 0 | 0 | 4 | 1 | 13 | 20 |
| XuHuaiyu      | 44 ( 5) |  0 | 0 | 0 | 2 | 3 | 0 |  0 |  5 |
| dyzsr         |  0 ( 0) |  0 | 0 | 0 | 0 | 0 | 0 |  0 |  0 |
| eurekaka      | 18 ( 5) |  6 | 0 | 0 | 3 | 1 | 0 |  0 | 10 |
| fzhedu        |  3 ( 3) |  1 | 0 | 0 | 0 | 0 | 0 |  0 |  1 |
| ichn-hu       |  9 ( 2) |  1 | 0 | 0 | 1 | 2 | 0 |  2 |  6 |
| lzmhhh123     | 41 ( 2) | 11 | 0 | 0 | 8 | 6 | 0 |  5 | 30 |
| nullnotnil    |  0 ( 0) |  0 | 0 | 0 | 0 | 0 | 0 |  0 |  0 |
| qw4990        | 40 ( 6) |  2 | 0 | 0 | 2 | 0 | 0 |  6 | 10 |
| rebelice      |  0 ( 0) |  0 | 0 | 0 | 0 | 0 | 0 |  0 |  0 |
| time-and-fate |  2 ( 0) |  0 | 0 | 0 | 0 | 0 | 0 |  0 |  0 |
| winoros       | 37 ( 5) |  4 | 0 | 0 | 1 | 1 | 0 |  1 |  7 |
| wshwsh12      | 19 ( 4) |  3 | 0 | 0 | 0 | 1 | 0 |  4 |  8 |


<details> 
  <summary>Reminiscent's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                               PR                                                                | C | LASTED |
|------------|---------------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/20972 | [expression: POC implementation of Vitess hashing algorithm.](https://github.com/pingcap/tidb/pull/20972)                       |   | 13d0h  |
| tidb/21095 | [planner, executor: parallelize stream aggregation with shuffle.](https://github.com/pingcap/tidb/pull/21095)                   |   | 7d11h  |
| tidb/21137 | [executor: specially handle empty input for apply's outer child aggregate (#20544)](https://github.com/pingcap/tidb/pull/21137) |   | 4d20h  |
| tidb/21182 | [table: fix insert value into hash partition table which not int](https://github.com/pingcap/tidb/pull/21182)                   |   | 2d21h  |


## Reviewed in Last 7 Days

|    REPO    |                                                           PR                                                           | C | D |   R    |
|------------|------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/21155 | [util/chunk: fix slice out of bound panic](https://github.com/pingcap/tidb/pull/21155)                                 |   | 1 | 3d12h  |
| tidb/21174 | [expression: enable coprocessor pushdown of function UUID](https://github.com/pingcap/tidb/pull/21174)                 |   | 1 | 2d18h  |
| tidb/21061 | [planner/core: use constant propagate before predicates push down](https://github.com/pingcap/tidb/pull/21061)         |   | 4 | 4d22h  |
| tidb/20544 | [executor: specially handle empty input for apply's outer child aggregate](https://github.com/pingcap/tidb/pull/20544) |   | 6 | 28d20h |


</details> 


<details> 
  <summary>SunRunAway's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                                        PR                                                                                         | C | LASTED  |
|--------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| docs-cn/4685 | [sql tuning: add explain walkthrough](https://github.com/pingcap/docs-cn/pull/4685)                                                                                               |   | 39d23h  |
| docs/4219    | [toc: add sql optimization-related docs](https://github.com/pingcap/docs/pull/4219)                                                                                               |   | 11d10h  |
| docs-cn/4913 | [explain: add indexes](https://github.com/pingcap/docs-cn/pull/4913)                                                                                                              |   | 7d17h   |
| tidb/15370   | [planner,executor: Refactor Shuffle and implement parallel Sort](https://github.com/pingcap/tidb/pull/15370)                                                                      | Y | 254d18h |
| docs-cn/4933 | [explain: add joins](https://github.com/pingcap/docs-cn/pull/4933)                                                                                                                |   | 3d19h   |
| tidb/15462   | [executor: implement `graceHashJoin`](https://github.com/pingcap/tidb/pull/15462)                                                                                                 | Y | 250d17h |
| tidb/16967   | [executor: Refactor Shuffle and implement parallel sort (executor part)](https://github.com/pingcap/tidb/pull/16967)                                                              | Y | 205d10h |
| tidb/17238   | [*: refactor table.Allocator to improve readability](https://github.com/pingcap/tidb/pull/17238)                                                                                  |   | 192d18h |
| tidb/19120   | [executor: Concurrently fetch chunks and insert them to a concurrent hash table in hash build](https://github.com/pingcap/tidb/pull/19120)                                        |   | 104d21h |
| tidb/19178   | [executor: Refactor probe channel](https://github.com/pingcap/tidb/pull/19178)                                                                                                    |   | 102d16h |
| tidb/19347   | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347)                                              |   | 94d23h  |
| tidb/19807   | [executor: parallel evaluation for aggregate functions with distinct in hashAgg](https://github.com/pingcap/tidb/pull/19807)                                                      |   | 80d10h  |
| tidb/19900   | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                                                                    | Y | 75d18h  |
| tidb/20140   | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                                                                |   | 62d22h  |
| tidb/20220   | [*: new secondary index value format](https://github.com/pingcap/tidb/pull/20220)                                                                                                 |   | 59d16h  |
| tidb/20316   | [docs/design: add design doc for index usage information](https://github.com/pingcap/tidb/pull/20316)                                                                             |   | 54d17h  |
| tidb/20335   | [planner, executor: enable inline projection for Selection](https://github.com/pingcap/tidb/pull/20335)                                                                           | Y | 51d17h  |
| tidb/20360   | [planner: refine explain info for batch cop](https://github.com/pingcap/tidb/pull/20360)                                                                                          |   | 45d22h  |
| tidb/20397   | [parser: replace ast.SelectLockInShareMode with ast.SelectLockForShare](https://github.com/pingcap/tidb/pull/20397)                                                               |   | 43d18h  |
| tidb/20591   | [server: redact some error code](https://github.com/pingcap/tidb/pull/20591)                                                                                                      |   | 32d16h  |
| tidb/20615   | [utils: Avoid panic when getting memory](https://github.com/pingcap/tidb/pull/20615)                                                                                              |   | 31d2h   |
| tidb/20652   | [ddl: Convert ddl types automatically according to length](https://github.com/pingcap/tidb/pull/20652)                                                                            |   | 27d23h  |
| tidb/20689   | [expression: make TIME function compatible with MySQL (#19158)](https://github.com/pingcap/tidb/pull/20689)                                                                       |   | 26d20h  |
| tidb/20750   | [executor, infoschema, planner: optimize query cluster_slow_query](https://github.com/pingcap/tidb/pull/20750)                                                                    |   | 21d23h  |
| tidb/20752   | [*: trace statsCache and preparePlanCache by Global memory tracker.](https://github.com/pingcap/tidb/pull/20752)                                                                  |   | 21d22h  |
| tidb/20765   | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                                                                 |   | 21d16h  |
| tidb/20789   | [*: Let binary literal can be convert to enum and set](https://github.com/pingcap/tidb/pull/20789)                                                                                |   | 20d21h  |
| tidb/20799   | [planner: bypass the DNF restriction if index merge hint is specified](https://github.com/pingcap/tidb/pull/20799)                                                                |   | 20d16h  |
| tidb/20844   | [executor: introduce new variables to control Apply's behaviors and add more tests for it](https://github.com/pingcap/tidb/pull/20844)                                            |   | 19d13h  |
| tidb/20868   | [execution : fix Compatibility between select and mysql](https://github.com/pingcap/tidb/pull/20868)                                                                              |   | 18d18h  |
| tidb/20894   | [planner, store/tikv, executor:Support shuffled hash join and refine codes](https://github.com/pingcap/tidb/pull/20894)                                                           |   | 17d18h  |
| tidb/20919   | [util: redact log for expensive sqls (#20852)](https://github.com/pingcap/tidb/pull/20919)                                                                                        |   | 14d23h  |
| tidb/20942   | [executor,planner/core,util/plancodec: extend executor.ShuffleExec and planner.core.PhysicalShuffle to support multiple data sources](https://github.com/pingcap/tidb/pull/20942) |   | 14d12h  |
| tidb/20947   | [expression: handle tp.flen overflow in to_base64 function](https://github.com/pingcap/tidb/pull/20947)                                                                           |   | 14d0h   |
| tidb/20984   | [expression, planner: fix decimal results for aggregate functions (#20017)](https://github.com/pingcap/tidb/pull/20984)                                                           |   | 12d19h  |
| tidb/21052   | [executor: fix cannot use explain for with the statement queried by explain analyze](https://github.com/pingcap/tidb/pull/21052)                                                  |   | 10d13h  |
| tidb/21061   | [planner/core: use constant propagate before predicates push down](https://github.com/pingcap/tidb/pull/21061)                                                                    |   | 8d15h   |
| tidb/21100   | [*: support read only lock](https://github.com/pingcap/tidb/pull/21100)                                                                                                           |   | 6d19h   |
| tidb/21101   | [*: support SQL bind for Update / Delete / Insert / Replace (#20686)](https://github.com/pingcap/tidb/pull/21101)                                                                 |   | 6d18h   |
| tidb/21137   | [executor: specially handle empty input for apply's outer child aggregate (#20544)](https://github.com/pingcap/tidb/pull/21137)                                                   |   | 4d20h   |
| tidb/21146   | [planner: ban (index) merge join heuristically when convert eq cond to other cond (#21138)](https://github.com/pingcap/tidb/pull/21146)                                           |   | 4d16h   |
| tidb/21148   | [planner,executor: fix 'select ...(join on partition table) for update' panic](https://github.com/pingcap/tidb/pull/21148)                                                        |   | 4d15h   |
| tidb/21163   | [hptc ycsb-worloade: Make use of AppendRows in all cases ](https://github.com/pingcap/tidb/pull/21163)                                                                            |   | 3d21h   |
| tidb/21199   | [executor: load data statement shoule not be prepared (#21188)](https://github.com/pingcap/tidb/pull/21199)                                                                       |   | 20h     |
| tidb/21207   | [[WIP]planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                                                              |   | 18h     |
| tidb/21209   | [session: add encoded plan cache to avoid repeated encode plan when recording slow-log and statement (#21139)](https://github.com/pingcap/tidb/pull/21209)                        |   | 18h     |
| tidb/21219   | [planner: make index-hash-join and index-merge-join consider collation (#21201)](https://github.com/pingcap/tidb/pull/21219)                                                      |   | 16h     |
| tidb/21222   | [executor: load data statement shoule not be prepared (#21188)](https://github.com/pingcap/tidb/pull/21222)                                                                       |   | 14h     |
| tidb/21223   | [expression: Improve compatibility of string literal comparisons](https://github.com/pingcap/tidb/pull/21223)                                                                     |   | 14h     |


## Reviewed in Last 7 Days

|    REPO    |                                                                        PR                                                                         | C | D |   R    |
|------------|---------------------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/21139 | [session: add encoded plan cache to avoid repeated encode plan when recording slow-log and statement](https://github.com/pingcap/tidb/pull/21139) |   | 1 | 4d0h   |
| tidb/21117 | [table/tables: make the test case more stable](https://github.com/pingcap/tidb/pull/21117)                                                        |   | 1 | 5d18h  |
| tidb/21099 | [expression: adjust int constant when compare with year type](https://github.com/pingcap/tidb/pull/21099)                                         |   | 5 | 2d4h   |
| tidb/21133 | [distsql: fix issue of table reader runtime stats display wrong result. (#21072)](https://github.com/pingcap/tidb/pull/21133)                     |   | 5 | 0h     |
| tidb/21067 | [executor: fix analyze update panic cause by duplicate call analyze executor Close method (#20390)](https://github.com/pingcap/tidb/pull/21067)   |   | 5 | 3d0h   |
| tidb/20886 | [*: optimize for encoding huge plan (#20811)](https://github.com/pingcap/tidb/pull/20886)                                                         |   | 5 | 13d1h  |
| tidb/20544 | [executor: specially handle empty input for apply's outer child aggregate](https://github.com/pingcap/tidb/pull/20544)                            |   | 6 | 29d5h  |
| tidb/20898 | [executor: modify the error message of insert time value (#20847)](https://github.com/pingcap/tidb/pull/20898)                                    |   | 7 | 10d23h |
| tidb/20944 | [executor: fix issue of runtime stats of index merge join is wrong (#20892)](https://github.com/pingcap/tidb/pull/20944)                          |   | 7 | 7d16h  |
| tidb/20889 | [ddl: forbid invalid usage of window function in the generated column (#20855)](https://github.com/pingcap/tidb/pull/20889)                       |   | 7 | 11d3h  |
| tidb/21001 | [planner: check view recursion when building source from view (#20398)](https://github.com/pingcap/tidb/pull/21001)                               |   | 7 | 5d4h   |
| tidb/21002 | [planner: add missing table lock check for fast plan (#20948)](https://github.com/pingcap/tidb/pull/21002)                                        |   | 7 | 5d3h   |
| tidb/21050 | [planner, expression: fix a bug causes schema change after DML (#21027)](https://github.com/pingcap/tidb/pull/21050)                              |   | 7 | 3d19h  |
| tidb/21045 | [executor: fix show global variables return session variables also (#19341)](https://github.com/pingcap/tidb/pull/21045)                          |   | 7 | 3d22h  |
| tidb/21074 | [executor: fix The JSON Data can not import to TiDB correctly by `load data` (#21043)](https://github.com/pingcap/tidb/pull/21074)                |   | 7 | 1d3h   |
| tidb/20413 | [execute: fill arguments when returning `ErrUnknownSystemVar` in `SetSessionSystemVar` (#20102)](https://github.com/pingcap/tidb/pull/20413)      |   | 7 | 35d19h |
| tidb/21019 | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21019)                                    |   | 7 | 4d21h  |
| tidb/20609 | [planner: refine the explain message of expression (#16276) (#20410)](https://github.com/pingcap/tidb/pull/20609)                                 |   | 7 | 24d21h |
| tidb/21072 | [distsql: fix issue of table reader runtime stats display wrong result.](https://github.com/pingcap/tidb/pull/21072)                              |   | 7 | 1d1h   |
| tidb/21071 | [executor: fix unstable test of TestIntegrationCopCache](https://github.com/pingcap/tidb/pull/21071)                                              |   | 7 | 1d0h   |


</details> 


<details> 
  <summary>XuHuaiyu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                              PR                                                                              | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/17997 | [expression: make greatest/least type comparison compatible with MySQL](https://github.com/pingcap/tidb/pull/17997)                                          | Y | 162d19h |
| tidb/18742 | [*: Add memory information of executors if OOM action fires for debugging](https://github.com/pingcap/tidb/pull/18742)                                       |   | 123d16h |
| tidb/19292 | [planner: suppport left join in join reorder](https://github.com/pingcap/tidb/pull/19292)                                                                    |   | 96d16h  |
| tidb/19900 | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                                               | Y | 75d18h  |
| tidb/20040 | [planner, expression: take NullFlag into consideration when optimize the `int non-const` <cmp > `non-int const`](https://github.com/pingcap/tidb/pull/20040) | Y | 68d13h  |
| tidb/20140 | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                                           |   | 62d22h  |
| tidb/20233 | [expression, types: fix datetime and year comparison error](https://github.com/pingcap/tidb/pull/20233)                                                      | Y | 58d7h   |
| tidb/20311 | [expression: fix overflow error when convert bit to int64 (#20266)](https://github.com/pingcap/tidb/pull/20311)                                              |   | 54d21h  |
| tidb/20350 | [executor: support read global indexes in IndexMergeReader and index join](https://github.com/pingcap/tidb/pull/20350)                                       | Y | 48d13h  |
| tidb/20435 | [planner, expression: make non-lookup condition prune ahead and cache partition-by expression](https://github.com/pingcap/tidb/pull/20435)                   |   | 41d12h  |
| tidb/20505 | [*: Add metrics for oom-action and sql memory usage.](https://github.com/pingcap/tidb/pull/20505)                                                            |   | 35d18h  |
| tidb/20576 | [*: fix stats feedback after tableReader handle multiple ranges](https://github.com/pingcap/tidb/pull/20576)                                                 |   | 33d12h  |
| tidb/20577 | [executor: Add the HashAggExec runtime information](https://github.com/pingcap/tidb/pull/20577)                                                              |   | 32d23h  |
| tidb/20613 | [executor: fix issue of hash join fetch time inaccurate](https://github.com/pingcap/tidb/pull/20613)                                                         |   | 31d13h  |
| tidb/20706 | [expression: set IsBooleanFlag for boolean scalar functions](https://github.com/pingcap/tidb/pull/20706)                                                     |   | 25d20h  |
| tidb/20752 | [*: trace statsCache and preparePlanCache by Global memory tracker.](https://github.com/pingcap/tidb/pull/20752)                                             |   | 21d22h  |
| tidb/20790 | [collation: add pinyin collation for chinese charset support](https://github.com/pingcap/tidb/pull/20790)                                                    |   | 20d20h  |
| tidb/20793 | [planner, executor: enable inline projection for Apply](https://github.com/pingcap/tidb/pull/20793)                                                          |   | 20d20h  |
| tidb/20844 | [executor: introduce new variables to control Apply's behaviors and add more tests for it](https://github.com/pingcap/tidb/pull/20844)                       |   | 19d13h  |
| tidb/20868 | [execution : fix Compatibility between select and mysql](https://github.com/pingcap/tidb/pull/20868)                                                         |   | 18d18h  |
| tidb/20905 | [planner: fix statement-optimize not work in `TryFastPlan`](https://github.com/pingcap/tidb/pull/20905)                                                      |   | 17d16h  |
| tidb/20938 | [planner: fix update statement not blocked by primary (#20842)](https://github.com/pingcap/tidb/pull/20938)                                                  |   | 14d16h  |
| tidb/20972 | [expression: POC implementation of Vitess hashing algorithm.](https://github.com/pingcap/tidb/pull/20972)                                                    |   | 13d0h   |
| tidb/20987 | [expression, executor: allow insert strings with overflowed trailing spaces](https://github.com/pingcap/tidb/pull/20987)                                     |   | 12d17h  |
| tidb/21000 | [planner: check view recursion when building source from view (#20398)](https://github.com/pingcap/tidb/pull/21000)                                          |   | 11d23h  |
| tidb/21058 | [executor: fix sum(double) result if value is +Inf or -Inf](https://github.com/pingcap/tidb/pull/21058)                                                      |   | 8d20h   |
| tidb/21064 | [planner, executor: fix cast not check error](https://github.com/pingcap/tidb/pull/21064)                                                                    |   | 8d8h    |
| tidb/21078 | [planner/cascades: add rule `TransformJoinCondToSel` (#20460)](https://github.com/pingcap/tidb/pull/21078)                                                   |   | 7d20h   |
| tidb/21104 | [util/hint: fix ineffective index hint for delete / update in SQL Bind](https://github.com/pingcap/tidb/pull/21104)                                          |   | 6d18h   |
| tidb/21132 | [expresssion: Fix unexpected panic when using IF function.](https://github.com/pingcap/tidb/pull/21132)                                                      |   | 4d22h   |
| tidb/21146 | [planner: ban (index) merge join heuristically when convert eq cond to other cond (#21138)](https://github.com/pingcap/tidb/pull/21146)                      |   | 4d16h   |
| tidb/21149 | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                                |   | 4d14h   |
| tidb/21150 | [expression: fix type infer for tidb's builtin compare(least and greatest)](https://github.com/pingcap/tidb/pull/21150)                                      |   | 4d13h   |
| tidb/21155 | [util/chunk: fix slice out of bound panic](https://github.com/pingcap/tidb/pull/21155)                                                                       |   | 4d11h   |
| tidb/21166 | [mocktikv: select count result differs between tikv and mocktikv](https://github.com/pingcap/tidb/pull/21166)                                                |   | 3d19h   |
| tidb/21170 | [*: implement priority control for OOM Action](https://github.com/pingcap/tidb/pull/21170)                                                                   |   | 3d18h   |
| tidb/21176 | [expression: keep the original data type when doing date arithmetic operations  (#20940)](https://github.com/pingcap/tidb/pull/21176)                        |   | 3d16h   |
| tidb/21196 | [*: fix unable to configure memory-usage-alarm by config file.](https://github.com/pingcap/tidb/pull/21196)                                                  |   | 22h     |
| tidb/21208 | [executor: Fix IndexMergeReader works incorrectly under RC isolation level](https://github.com/pingcap/tidb/pull/21208)                                      |   | 18h     |
| tidb/21209 | [session: add encoded plan cache to avoid repeated encode plan when recording slow-log and statement (#21139)](https://github.com/pingcap/tidb/pull/21209)   |   | 18h     |
| tidb/21211 | [executor: fix bug when parsing the slow logs if the logNum is small (#20959)](https://github.com/pingcap/tidb/pull/21211)                                   |   | 18h     |
| tidb/21214 | [go.mod, bindinfo: update parser to fix binding doesn't work for prepared stmt with LIMIT](https://github.com/pingcap/tidb/pull/21214)                       |   | 17h     |
| tidb/21229 | [txn: locks exist keys for point_get & batch_point_get](https://github.com/pingcap/tidb/pull/21229)                                                          |   | 11h     |
| tidb/21233 | [executor: set the inner join keys' field length to unspecified](https://github.com/pingcap/tidb/pull/21233)                                                 |   | 7h      |


## Reviewed in Last 7 Days

|    REPO    |                                                               PR                                                               | C | D |   R   |
|------------|--------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/21160 | [util: fix wrong calculation to get memory usage. (#21158)](https://github.com/pingcap/tidb/pull/21160)                        |   | 4 | 3h    |
| tidb/21161 | [util: fix wrong calculation to get memory usage. (#21158)](https://github.com/pingcap/tidb/pull/21161)                        |   | 4 | 3h    |
| tidb/21130 | [*: inject projection for the plan pushed to TiDB (#21090)](https://github.com/pingcap/tidb/pull/21130)                        |   | 5 | 3h    |
| tidb/21138 | [planner: ban (index) merge join heuristically when convert eq cond to other cond](https://github.com/pingcap/tidb/pull/21138) |   | 5 | 0h    |
| tidb/20927 | [*: support to modify config server-memory-quota-ratio dynamically](https://github.com/pingcap/tidb/pull/20927)                |   | 5 | 9d23h |


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

|    REPO    |                                                                  PR                                                                  | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/14729 | [planner: fix constant propagation for PredicatePushDown](https://github.com/pingcap/tidb/pull/14729)                                | Y | 286d17h |
| tidb/14831 | [planner/cascades: add implementationRule for IndexLookUpJoin](https://github.com/pingcap/tidb/pull/14831)                           |   | 279d17h |
| tidb/15090 | [planner/cascades: refine the row count estimation of TiKV layer Selection](https://github.com/pingcap/tidb/pull/15090)              |   | 265d17h |
| tidb/15157 | [planner/cascades: implement `HashCode` method for all the LogicalPlans](https://github.com/pingcap/tidb/pull/15157)                 | Y | 263d14h |
| tidb/15335 | [planner/cascades: add transformation rule PullAggregationUpApply & EliminateMaxOneRow](https://github.com/pingcap/tidb/pull/15335)  |   | 256d17h |
| tidb/15370 | [planner,executor: Refactor Shuffle and implement parallel Sort](https://github.com/pingcap/tidb/pull/15370)                         | Y | 254d18h |
| tidb/17276 | [planner/cascades: add rule InjectProjectionBelowSort](https://github.com/pingcap/tidb/pull/17276)                                   | Y | 189d8h  |
| tidb/18882 | [planner, executor: add explain for `MetricSummaryTableExtractor`](https://github.com/pingcap/tidb/pull/18882)                       | Y | 116d17h |
| tidb/19347 | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347) |   | 94d23h  |
| tidb/20478 | [planner: add projection if expand virtual generated column](https://github.com/pingcap/tidb/pull/20478)                             |   | 39d13h  |
| tidb/20580 | [statistics: add bucket ndv for index histogram](https://github.com/pingcap/tidb/pull/20580)                                         |   | 32d20h  |
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                            |   | 18d16h  |
| tidb/21003 | [planner: fix index merge doesn't take effect when using prefix key (#20425)](https://github.com/pingcap/tidb/pull/21003)            |   | 11d21h  |
| tidb/21007 | [*: extract topn out of histogram correctly](https://github.com/pingcap/tidb/pull/21007)                                             |   | 11d20h  |
| tidb/21080 | [planner: detect unknown column in window clause](https://github.com/pingcap/tidb/pull/21080)                                        |   | 7d19h   |
| tidb/21110 | [planner: Add table engine name check](https://github.com/pingcap/tidb/pull/21110)                                                   |   | 6d16h   |
| tidb/21165 | [planner: fix ambiguous field when resolve having expr ](https://github.com/pingcap/tidb/pull/21165)                                 |   | 3d19h   |
| tidb/21216 | [planner: check for only_full_group_by in ORDER BY and HAVING](https://github.com/pingcap/tidb/pull/21216)                           |   | 17h     |


## Reviewed in Last 7 Days

|      REPO      |                                                                                    PR                                                                                     | C | D |   R    |
|----------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/21214     | [go.mod, bindinfo: update parser to fix binding doesn't work for prepared stmt with LIMIT](https://github.com/pingcap/tidb/pull/21214)                                    |   | 1 | 0h     |
| parser/1103    | [digester: consider paramMarker as literal when normalizing SQL (#1040)](https://github.com/pingcap/parser/pull/1103)                                                     |   | 1 | 2d21h  |
| tidb/21205     | [planner: `DELETE` cannot delete data in some cases when the database name is capitalized (#21202)](https://github.com/pingcap/tidb/pull/21205)                           |   | 1 | 0h     |
| tidb/21206     | [planner: `DELETE` cannot delete data in some cases when the database name is capitalized (#21202)](https://github.com/pingcap/tidb/pull/21206)                           |   | 1 | 0h     |
| tidb/21148     | [planner,executor: fix 'select ...(join on partition table) for update' panic](https://github.com/pingcap/tidb/pull/21148)                                                |   | 1 | 3d20h  |
| tidb/21202     | [planner: `DELETE` cannot delete data in some cases when the database name is capitalized](https://github.com/pingcap/tidb/pull/21202)                                    |   | 1 | 0h     |
| parser/1040    | [digester: consider paramMarker as literal when normalizing SQL](https://github.com/pingcap/parser/pull/1040)                                                             |   | 4 | 54d0h  |
| tidb/18017     | [planner, sessionctx : Add 'last_plan_from_binding' to help know whether sql's plan is matched with the hints in the binding](https://github.com/pingcap/tidb/pull/18017) | Y | 4 | 158d1h |
| tidb-test/1102 | [fix results for #19620 (#1101)](https://github.com/pingcap/tidb-test/pull/1102)                                                                                          |   | 4 | 7d23h  |
| tidb/21084     | [planner: fix unexpected bad plan when IndexJoin inner side estRow is 0.](https://github.com/pingcap/tidb/pull/21084)                                                     |   | 5 | 2d23h  |


</details> 


<details> 
  <summary>fzhedu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                            PR                                                             | C | LASTED |
|------------|---------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/19310 | [expression: make tidb_decode_key return json type and support escape string](https://github.com/pingcap/tidb/pull/19310) | Y | 96d0h  |
| tidb/19845 | [expression:fix FORMAT compatibility issue #11206](https://github.com/pingcap/tidb/pull/19845)                            | Y | 77d15h |
| tidb/20117 | [optimizer: fix issue on incorrect result of natural join](https://github.com/pingcap/tidb/pull/20117)                    | Y | 63d20h |


## Reviewed in Last 7 Days

|    REPO    |                                                           PR                                                            | C | D |   R    |
|------------|-------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/20894 | [planner, store/tikv, executor:Support shuffled hash join and refine codes](https://github.com/pingcap/tidb/pull/20894) |   | 1 | 16d19h |


</details> 


<details> 
  <summary>ichn-hu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                   PR                                                                    | C | LASTED  |
|------------|-----------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/17997 | [expression: make greatest/least type comparison compatible with MySQL](https://github.com/pingcap/tidb/pull/17997)                     | Y | 162d19h |
| tidb/20015 | [expression: handle zero date in `date_add()`](https://github.com/pingcap/tidb/pull/20015)                                              | Y | 69d13h  |
| tidb/20961 | [expression: convert a date to number if the date is used in numeric context](https://github.com/pingcap/tidb/pull/20961)               |   | 13d17h  |
| tidb/20981 | [expression: fix wrong inferred type for sum and avg (#20926)](https://github.com/pingcap/tidb/pull/20981)                              |   | 12d19h  |
| tidb/21074 | [executor: fix The JSON Data can not import to TiDB correctly by `load data` (#21043)](https://github.com/pingcap/tidb/pull/21074)      |   | 7d21h   |
| tidb/21144 | [expression: fix builtin IF truncation of type len (#20743)](https://github.com/pingcap/tidb/pull/21144)                                |   | 4d17h   |
| tidb/21146 | [planner: ban (index) merge join heuristically when convert eq cond to other cond (#21138)](https://github.com/pingcap/tidb/pull/21146) |   | 4d16h   |
| tidb/21176 | [expression: keep the original data type when doing date arithmetic operations  (#20940)](https://github.com/pingcap/tidb/pull/21176)   |   | 3d16h   |
| tidb/21220 | [expression: fix compatibility behaviors in zero datetime with MySQL](https://github.com/pingcap/tidb/pull/21220)                       |   | 16h     |


## Reviewed in Last 7 Days

|    REPO    |                                                               PR                                                               | C | D |   R    |
|------------|--------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/20747 | [executor: fix LEAD and LAG's default value can not adapt to field type](https://github.com/pingcap/tidb/pull/20747)           |   | 1 | 21d21h |
| tidb/20940 | [expression: keep the original data type when doing date arithmetic operations ](https://github.com/pingcap/tidb/pull/20940)   |   | 4 | 10d19h |
| tidb/21062 | [planner, type:  fix AggFieldType error when encouter unsigned and sign type](https://github.com/pingcap/tidb/pull/21062)      |   | 5 | 3d20h  |
| tidb/21138 | [planner: ban (index) merge join heuristically when convert eq cond to other cond](https://github.com/pingcap/tidb/pull/21138) |   | 5 | 0h     |
| tidb/20898 | [executor: modify the error message of insert time value (#20847)](https://github.com/pingcap/tidb/pull/20898)                 |   | 7 | 11d3h  |
| tidb/20206 | [table: fix zero date in different sqlmode](https://github.com/pingcap/tidb/pull/20206)                                        | Y | 7 | 53d18h |


</details> 


<details> 
  <summary>lzmhhh123's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                  PR                                                                   | C | LASTED  |
|--------------|---------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| docs-cn/4912 | [explain: add subqueries](https://github.com/pingcap/docs-cn/pull/4912)                                                               |   | 7d18h   |
| tidb/14729   | [planner: fix constant propagation for PredicatePushDown](https://github.com/pingcap/tidb/pull/14729)                                 | Y | 286d17h |
| docs-cn/4913 | [explain: add indexes](https://github.com/pingcap/docs-cn/pull/4913)                                                                  |   | 7d17h   |
| tidb/17414   | [add curCost based join reorder algorithm](https://github.com/pingcap/tidb/pull/17414)                                                |   | 181d17h |
| tidb/19347   | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347)  |   | 94d23h  |
| tidb/19698   | [*: update test cases to support new collation enabled by default](https://github.com/pingcap/tidb/pull/19698)                        |   | 82d22h  |
| tidb/19767   | [planner: rename needFrame to ignoreFrame](https://github.com/pingcap/tidb/pull/19767)                                                |   | 81d16h  |
| tidb/20044   | [expression: Add column nullability checking before "refine args"](https://github.com/pingcap/tidb/pull/20044)                        | Y | 68d7h   |
| tidb/20444   | [expression: add json_merge_patch](https://github.com/pingcap/tidb/pull/20444)                                                        |   | 40d21h  |
| tidb/20465   | [expression: add uuidShortFunction](https://github.com/pingcap/tidb/pull/20465)                                                       |   | 39d19h  |
| tidb/20505   | [*: Add metrics for oom-action and sql memory usage.](https://github.com/pingcap/tidb/pull/20505)                                     |   | 35d18h  |
| tidb/20543   | [planner: refine the error message when split region by decimal column](https://github.com/pingcap/tidb/pull/20543)                   |   | 34d14h  |
| tidb/20609   | [planner: refine the explain message of expression (#16276) (#20410)](https://github.com/pingcap/tidb/pull/20609)                     |   | 31d15h  |
| tidb/20618   | [planner: fix update generated columns error](https://github.com/pingcap/tidb/pull/20618)                                             |   | 30d20h  |
| tidb/20642   | [executor: modify admin executors to support partitioned table with global index](https://github.com/pingcap/tidb/pull/20642)         |   | 28d15h  |
| tidb/20785   | [errno, infoschema, executor, server: add client error infoschema tables](https://github.com/pingcap/tidb/pull/20785)                 |   | 20d23h  |
| tidb/20825   | [executor: add diagnosis rule to check Transparent Huge Pages(THP) enabled (#20611)](https://github.com/pingcap/tidb/pull/20825)      |   | 19d18h  |
| tidb/20865   | [executor:Add runtime information for UnionScanExec](https://github.com/pingcap/tidb/pull/20865)                                      |   | 18d18h  |
| tidb/20898   | [executor: modify the error message of insert time value (#20847)](https://github.com/pingcap/tidb/pull/20898)                        |   | 17d17h  |
| tidb/20903   | [planner: fix confused and unnecessary double-projection in plans.](https://github.com/pingcap/tidb/pull/20903)                       |   | 17d17h  |
| tidb/20929   | [types:  Add a limitation about float data type](https://github.com/pingcap/tidb/pull/20929)                                          |   | 14d18h  |
| tidb/20938   | [planner: fix update statement not blocked by primary (#20842)](https://github.com/pingcap/tidb/pull/20938)                           |   | 14d16h  |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                        |   | 11d16h  |
| tidb/21020   | [planner: generate proper hint for IndexHashJoin / IndexMergeJoin (#20992)](https://github.com/pingcap/tidb/pull/21020)               |   | 11d16h  |
| tidb/21051   | [executor: change read slow-log file module to concurrent](https://github.com/pingcap/tidb/pull/21051)                                |   | 10d14h  |
| tidb/21060   | [planner: fix distinct push across projection when read partition table](https://github.com/pingcap/tidb/pull/21060)                  |   | 8d16h   |
| tidb/21062   | [planner, type:  fix AggFieldType error when encouter unsigned and sign type](https://github.com/pingcap/tidb/pull/21062)             |   | 8d14h   |
| tidb/21078   | [planner/cascades: add rule `TransformJoinCondToSel` (#20460)](https://github.com/pingcap/tidb/pull/21078)                            |   | 7d20h   |
| tidb/21083   | [planner: reject invalid named window specs](https://github.com/pingcap/tidb/pull/21083)                                              |   | 7d18h   |
| tidb/21084   | [planner: fix unexpected bad plan when IndexJoin inner side estRow is 0.](https://github.com/pingcap/tidb/pull/21084)                 |   | 7d18h   |
| tidb/21107   | [*: differentiate types for user variables (#18973)](https://github.com/pingcap/tidb/pull/21107)                                      |   | 6d17h   |
| tidb/21120   | [planner: error by default for GROUP BY expr ASC|DESC](https://github.com/pingcap/tidb/pull/21120)                                    |   | 6d1h    |
| tidb/21137   | [executor: specially handle empty input for apply's outer child aggregate (#20544)](https://github.com/pingcap/tidb/pull/21137)       |   | 4d20h   |
| tidb/21143   | [bindinfo: dbname check for bindings should be case insensitive](https://github.com/pingcap/tidb/pull/21143)                          |   | 4d17h   |
| tidb/21173   | [planner: fix partition pruning when condition exceeds the range of column type](https://github.com/pingcap/tidb/pull/21173)          |   | 3d18h   |
| tidb/21176   | [expression: keep the original data type when doing date arithmetic operations  (#20940)](https://github.com/pingcap/tidb/pull/21176) |   | 3d16h   |
| tidb/21195   | [brie: integrate lightning to suport IMPORT statement](https://github.com/pingcap/tidb/pull/21195)                                    |   | 22h     |
| tidb/21199   | [executor: load data statement shoule not be prepared (#21188)](https://github.com/pingcap/tidb/pull/21199)                           |   | 20h     |
| tidb/21219   | [planner: make index-hash-join and index-merge-join consider collation (#21201)](https://github.com/pingcap/tidb/pull/21219)          |   | 16h     |
| tidb/21222   | [executor: load data statement shoule not be prepared (#21188)](https://github.com/pingcap/tidb/pull/21222)                           |   | 14h     |
| tidb/21233   | [executor: set the inner join keys' field length to unspecified](https://github.com/pingcap/tidb/pull/21233)                          |   | 7h      |


## Reviewed in Last 7 Days

|    REPO     |                                                                        PR                                                                         | C | D |   R    |
|-------------|---------------------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/21221  | [planner: make index-hash-join and index-merge-join consider the collation in the right way](https://github.com/pingcap/tidb/pull/21221)          |   | 1 | 1h     |
| tidb/21201  | [planner: make index-hash-join and index-merge-join consider collation](https://github.com/pingcap/tidb/pull/21201)                               |   | 1 | 2h     |
| tidb/21213  | [planner: set dbName for hinted query block table alias](https://github.com/pingcap/tidb/pull/21213)                                              |   | 1 | 0h     |
| tidb/20961  | [expression: convert a date to number if the date is used in numeric context](https://github.com/pingcap/tidb/pull/20961)                         |   | 1 | 12d23h |
| tidb/20965  | [planner: fold the GetVar into a constant if the query contains no SetVar for the same user variable](https://github.com/pingcap/tidb/pull/20965) |   | 1 | 12d19h |
| tidb/21205  | [planner: `DELETE` cannot delete data in some cases when the database name is capitalized (#21202)](https://github.com/pingcap/tidb/pull/21205)   |   | 1 | 0h     |
| tidb/21206  | [planner: `DELETE` cannot delete data in some cases when the database name is capitalized (#21202)](https://github.com/pingcap/tidb/pull/21206)   |   | 1 | 0h     |
| tidb/21202  | [planner: `DELETE` cannot delete data in some cases when the database name is capitalized](https://github.com/pingcap/tidb/pull/21202)            |   | 1 | 0h     |
| tidb/21188  | [executor: load data statement shoule not be prepared](https://github.com/pingcap/tidb/pull/21188)                                                |   | 1 | 16h    |
| tidb/21196  | [*: fix unable to configure memory-usage-alarm by config file.](https://github.com/pingcap/tidb/pull/21196)                                       |   | 1 | 0h     |
| tidb/21182  | [table: fix insert value into hash partition table which not int](https://github.com/pingcap/tidb/pull/21182)                                     |   | 1 | 1d23h  |
| tidb/20940  | [expression: keep the original data type when doing date arithmetic operations ](https://github.com/pingcap/tidb/pull/20940)                      |   | 4 | 10d20h |
| tidb/21061  | [planner/core: use constant propagate before predicates push down](https://github.com/pingcap/tidb/pull/21061)                                    |   | 4 | 4d22h  |
| tidb/21078  | [planner/cascades: add rule `TransformJoinCondToSel` (#20460)](https://github.com/pingcap/tidb/pull/21078)                                        |   | 4 | 4d3h   |
| tidb/21144  | [expression: fix builtin IF truncation of type len (#20743)](https://github.com/pingcap/tidb/pull/21144)                                          |   | 4 | 20h    |
| tidb/21161  | [util: fix wrong calculation to get memory usage. (#21158)](https://github.com/pingcap/tidb/pull/21161)                                           |   | 4 | 2h     |
| tidb/21160  | [util: fix wrong calculation to get memory usage. (#21158)](https://github.com/pingcap/tidb/pull/21160)                                           |   | 4 | 2h     |
| tidb/21124  | [planner: fix should not use point get plan](https://github.com/pingcap/tidb/pull/21124)                                                          |   | 4 | 1d13h  |
| tidb/21158  | [util: fix wrong calculation to get memory usage.](https://github.com/pingcap/tidb/pull/21158)                                                    |   | 4 | 0h     |
| tidb/21130  | [*: inject projection for the plan pushed to TiDB (#21090)](https://github.com/pingcap/tidb/pull/21130)                                           |   | 5 | 3h     |
| tidb/20743  | [expression: fix builtin IF truncation of type len](https://github.com/pingcap/tidb/pull/20743)                                                   |   | 5 | 19d8h  |
| parser/1101 | [parser, ast: track if order is implicit ASC/DESC](https://github.com/pingcap/parser/pull/1101)                                                   |   | 5 | 1d2h   |
| tidb/21058  | [executor: fix sum(double) result if value is +Inf or -Inf](https://github.com/pingcap/tidb/pull/21058)                                           |   | 5 | 3d21h  |
| tidb/21113  | [planner: disallow multi-updates on primary key (#20603)](https://github.com/pingcap/tidb/pull/21113)                                             |   | 5 | 1d16h  |
| tidb/21090  | [*: inject projection for the plan pushed to TiDB](https://github.com/pingcap/tidb/pull/21090)                                                    |   | 5 | 2d17h  |
| tidb/21099  | [expression: adjust int constant when compare with year type](https://github.com/pingcap/tidb/pull/21099)                                         |   | 7 | 3h     |
| tidb/21103  | [executor, planner: do not coalesce column for update/delete natural_using_join (#21085)](https://github.com/pingcap/tidb/pull/21103)             |   | 7 | 1h     |
| tidb/20206  | [table: fix zero date in different sqlmode](https://github.com/pingcap/tidb/pull/20206)                                                           | Y | 7 | 53d17h |
| tidb/21050  | [planner, expression: fix a bug causes schema change after DML (#21027)](https://github.com/pingcap/tidb/pull/21050)                              |   | 7 | 3d19h  |
| tidb/21085  | [executor, planner: do not coalesce column for update/delete natural_using_join](https://github.com/pingcap/tidb/pull/21085)                      |   | 7 | 23h    |


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

|    REPO    |                                                                     PR                                                                     | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/16305 | [expression: separate signatures for `ModInt`](https://github.com/pingcap/tidb/pull/16305)                                                 | Y | 224d23h |
| tidb/16967 | [executor: Refactor Shuffle and implement parallel sort (executor part)](https://github.com/pingcap/tidb/pull/16967)                       | Y | 205d10h |
| tidb/17396 | [types: improve StrToDate performance](https://github.com/pingcap/tidb/pull/17396)                                                         | Y | 182d9h  |
| tidb/18882 | [planner, executor: add explain for `MetricSummaryTableExtractor`](https://github.com/pingcap/tidb/pull/18882)                             | Y | 116d17h |
| tidb/19029 | [types: fix unexpected NOT_NULL flags](https://github.com/pingcap/tidb/pull/19029)                                                         |   | 109d22h |
| tidb/19120 | [executor: Concurrently fetch chunks and insert them to a concurrent hash table in hash build](https://github.com/pingcap/tidb/pull/19120) |   | 104d21h |
| tidb/19292 | [planner: suppport left join in join reorder](https://github.com/pingcap/tidb/pull/19292)                                                  |   | 96d16h  |
| tidb/19957 | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                     | Y | 73d13h  |
| tidb/20011 | [statistics: fix incorrect total count used in index selectivity computation](https://github.com/pingcap/tidb/pull/20011)                  |   | 69d15h  |
| tidb/20316 | [docs/design: add design doc for index usage information](https://github.com/pingcap/tidb/pull/20316)                                      |   | 54d17h  |
| tidb/20354 | [planner: rename relational operators (#14575)](https://github.com/pingcap/tidb/pull/20354)                                                | Y | 47d5h   |
| tidb/20399 | [*: make 'tidb_enable_change_column_type' available as a session variable](https://github.com/pingcap/tidb/pull/20399)                     |   | 43d15h  |
| tidb/20675 | [planner: fix explain-hint panic for joins generated by subquery](https://github.com/pingcap/tidb/pull/20675)                              |   | 27d17h  |
| tidb/20689 | [expression: make TIME function compatible with MySQL (#19158)](https://github.com/pingcap/tidb/pull/20689)                                |   | 26d20h  |
| tidb/20708 | [*: separate auto_increment ID allocator from _tidb_rowid allocator](https://github.com/pingcap/tidb/pull/20708)                           |   | 25d20h  |
| tidb/20747 | [executor: fix LEAD and LAG's default value can not adapt to field type](https://github.com/pingcap/tidb/pull/20747)                       |   | 22d19h  |
| tidb/20750 | [executor, infoschema, planner: optimize query cluster_slow_query](https://github.com/pingcap/tidb/pull/20750)                             |   | 21d23h  |
| tidb/20799 | [planner: bypass the DNF restriction if index merge hint is specified](https://github.com/pingcap/tidb/pull/20799)                         |   | 20d16h  |
| tidb/20836 | [planner: check for decimal format in cast expr](https://github.com/pingcap/tidb/pull/20836)                                               |   | 19d16h  |
| tidb/20883 | [*: support select from tablesample](https://github.com/pingcap/tidb/pull/20883)                                                           |   | 18d9h   |
| tidb/20929 | [types:  Add a limitation about float data type](https://github.com/pingcap/tidb/pull/20929)                                               |   | 14d18h  |
| tidb/20972 | [expression: POC implementation of Vitess hashing algorithm.](https://github.com/pingcap/tidb/pull/20972)                                  |   | 13d0h   |
| tidb/21018 | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                             |   | 11d16h  |
| tidb/21022 | [*: refine runtime stats display and tiny bug fix for metrics](https://github.com/pingcap/tidb/pull/21022)                                 |   | 11d15h  |
| tidb/21044 | [session: fix ineffective EXPLAIN FOR CONNECTION statement](https://github.com/pingcap/tidb/pull/21044)                                    |   | 10d17h  |
| tidb/21054 | [config: hide & deprecate enable-streaming (#20760)](https://github.com/pingcap/tidb/pull/21054)                                           |   | 10d8h   |
| tidb/21095 | [planner, executor: parallelize stream aggregation with shuffle.](https://github.com/pingcap/tidb/pull/21095)                              |   | 7d11h   |
| tidb/21101 | [*: support SQL bind for Update / Delete / Insert / Replace (#20686)](https://github.com/pingcap/tidb/pull/21101)                          |   | 6d18h   |
| tidb/21104 | [util/hint: fix ineffective index hint for delete / update in SQL Bind](https://github.com/pingcap/tidb/pull/21104)                        |   | 6d18h   |
| tidb/21105 | [executor: fix auto-id allocation during statements retry (#20659)](https://github.com/pingcap/tidb/pull/21105)                            |   | 6d18h   |
| tidb/21107 | [*: differentiate types for user variables (#18973)](https://github.com/pingcap/tidb/pull/21107)                                           |   | 6d17h   |
| tidb/21113 | [planner: disallow multi-updates on primary key (#20603)](https://github.com/pingcap/tidb/pull/21113)                                      |   | 6d15h   |
| tidb/21132 | [expresssion: Fix unexpected panic when using IF function.](https://github.com/pingcap/tidb/pull/21132)                                    |   | 4d22h   |
| tidb/21137 | [executor: specially handle empty input for apply's outer child aggregate (#20544)](https://github.com/pingcap/tidb/pull/21137)            |   | 4d20h   |
| tidb/21144 | [expression: fix builtin IF truncation of type len (#20743)](https://github.com/pingcap/tidb/pull/21144)                                   |   | 4d17h   |
| tidb/21149 | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                              |   | 4d14h   |
| tidb/21150 | [expression: fix type infer for tidb's builtin compare(least and greatest)](https://github.com/pingcap/tidb/pull/21150)                    |   | 4d13h   |
| tidb/21189 | [executor: modify lookupTableTask to return merged rows, and improve AppendRows](https://github.com/pingcap/tidb/pull/21189)               |   | 1d12h   |
| tidb/21220 | [expression: fix compatibility behaviors in zero datetime with MySQL](https://github.com/pingcap/tidb/pull/21220)                          |   | 16h     |
| tidb/21228 | [executor: return the result immediately when combining LIMIT row_count with DISTINCT](https://github.com/pingcap/tidb/pull/21228)         |   | 13h     |


## Reviewed in Last 7 Days

|    REPO    |                                                                                        PR                                                                                         | C | D |    R    |
|------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---|---------|
| tidb/20942 | [executor,planner/core,util/plancodec: extend executor.ShuffleExec and planner.core.PhysicalShuffle to support multiple data sources](https://github.com/pingcap/tidb/pull/20942) |   | 1 | 13d14h  |
| tidb/21095 | [planner, executor: parallelize stream aggregation with shuffle.](https://github.com/pingcap/tidb/pull/21095)                                                                     |   | 1 | 6d12h   |
| docs/4035  | [add a new document about how to analyze slow queries](https://github.com/pingcap/docs/pull/4035)                                                                                 |   | 4 | 38d19h  |
| tidb/20903 | [planner: fix confused and unnecessary double-projection in plans.](https://github.com/pingcap/tidb/pull/20903)                                                                   |   | 4 | 13d21h  |
| tidb/20653 | [executor:Add runtime stat for IndexMergeReaderExecutor](https://github.com/pingcap/tidb/pull/20653)                                                                              |   | 7 | 21d5h   |
| tidb/18973 | [*: differentiate types for user variables](https://github.com/pingcap/tidb/pull/18973)                                                                                           |   | 7 | 104d23h |
| tidb/21045 | [executor: fix show global variables return session variables also (#19341)](https://github.com/pingcap/tidb/pull/21045)                                                          |   | 7 | 3d22h   |
| tidb/21072 | [distsql: fix issue of table reader runtime stats display wrong result.](https://github.com/pingcap/tidb/pull/21072)                                                              |   | 7 | 1d3h    |
| tidb/21055 | [config: hide & deprecate enable-streaming (#20760)](https://github.com/pingcap/tidb/pull/21055)                                                                                  |   | 7 | 3d13h   |
| tidb/21073 | [plannr: build empty range for overflow predicate (#21042)](https://github.com/pingcap/tidb/pull/21073)                                                                           |   | 7 | 1d3h    |


</details> 


<details> 
  <summary>rebelice's detailed review report</summary> 

## To Be Reviewed

| REPO | PR | C | LASTED |
|------|----|---|--------|


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>time-and-fate's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                            PR                                                             | C | LASTED |
|------------|---------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/21003 | [planner: fix index merge doesn't take effect when using prefix key (#20425)](https://github.com/pingcap/tidb/pull/21003) |   | 11d21h |
| tidb/21007 | [*: extract topn out of histogram correctly](https://github.com/pingcap/tidb/pull/21007)                                  |   | 11d20h |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>winoros's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                           PR                                                                           | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/14424 | [expression: add nullable() method to check whether an expression can return null](https://github.com/pingcap/tidb/pull/14424)                         |   | 319d17h |
| tidb/14831 | [planner/cascades: add implementationRule for IndexLookUpJoin](https://github.com/pingcap/tidb/pull/14831)                                             |   | 279d17h |
| tidb/15090 | [planner/cascades: refine the row count estimation of TiKV layer Selection](https://github.com/pingcap/tidb/pull/15090)                                |   | 265d17h |
| tidb/15157 | [planner/cascades: implement `HashCode` method for all the LogicalPlans](https://github.com/pingcap/tidb/pull/15157)                                   | Y | 263d14h |
| tidb/15426 | [planner/cascades: add transformation rule PushSelDownApply & refactor PushSelDownJoin](https://github.com/pingcap/tidb/pull/15426)                    |   | 251d16h |
| tidb/16967 | [executor: Refactor Shuffle and implement parallel sort (executor part)](https://github.com/pingcap/tidb/pull/16967)                                   | Y | 205d10h |
| tidb/17414 | [add curCost based join reorder algorithm](https://github.com/pingcap/tidb/pull/17414)                                                                 |   | 181d17h |
| tidb/17996 | [planner: push avg & distinct functions across join](https://github.com/pingcap/tidb/pull/17996)                                                       | Y | 163d10h |
| tidb/19957 | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                                 | Y | 73d13h  |
| tidb/20011 | [statistics: fix incorrect total count used in index selectivity computation](https://github.com/pingcap/tidb/pull/20011)                              |   | 69d15h  |
| tidb/20091 | [statistics: change statistic cache to ristretto statscache](https://github.com/pingcap/tidb/pull/20091)                                               | Y | 66d22h  |
| tidb/20311 | [expression: fix overflow error when convert bit to int64 (#20266)](https://github.com/pingcap/tidb/pull/20311)                                        |   | 54d21h  |
| tidb/20482 | [planner: add EXPLAIN FORMAT=JSON](https://github.com/pingcap/tidb/pull/20482)                                                                         |   | 39d0h   |
| tidb/20664 | [executor, unistore: optimize analyze index by move it to analyze columns](https://github.com/pingcap/tidb/pull/20664)                                 |   | 27d18h  |
| tidb/20675 | [planner: fix explain-hint panic for joins generated by subquery](https://github.com/pingcap/tidb/pull/20675)                                          |   | 27d17h  |
| tidb/20765 | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                                      |   | 21d16h  |
| tidb/20799 | [planner: bypass the DNF restriction if index merge hint is specified](https://github.com/pingcap/tidb/pull/20799)                                     |   | 20d16h  |
| tidb/20836 | [planner: check for decimal format in cast expr](https://github.com/pingcap/tidb/pull/20836)                                                           |   | 19d16h  |
| tidb/20883 | [*: support select from tablesample](https://github.com/pingcap/tidb/pull/20883)                                                                       |   | 18d9h   |
| tidb/20889 | [ddl: forbid invalid usage of window function in the generated column (#20855)](https://github.com/pingcap/tidb/pull/20889)                            |   | 17d21h  |
| tidb/20965 | [planner: fold the GetVar into a constant if the query contains no SetVar for the same user variable](https://github.com/pingcap/tidb/pull/20965)      |   | 13d14h  |
| tidb/21000 | [planner: check view recursion when building source from view (#20398)](https://github.com/pingcap/tidb/pull/21000)                                    |   | 11d23h  |
| tidb/21014 | [statistics: GC index usage information](https://github.com/pingcap/tidb/pull/21014)                                                                   |   | 11d18h  |
| tidb/21018 | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                         |   | 11d16h  |
| tidb/21020 | [planner: generate proper hint for IndexHashJoin / IndexMergeJoin (#20992)](https://github.com/pingcap/tidb/pull/21020)                                |   | 11d16h  |
| tidb/21083 | [planner: reject invalid named window specs](https://github.com/pingcap/tidb/pull/21083)                                                               |   | 7d18h   |
| tidb/21086 | [*: seperate hash keys from join keys in IndexJoin (#20761)](https://github.com/pingcap/tidb/pull/21086)                                               |   | 7d17h   |
| tidb/21101 | [*: support SQL bind for Update / Delete / Insert / Replace (#20686)](https://github.com/pingcap/tidb/pull/21101)                                      |   | 6d18h   |
| tidb/21103 | [executor, planner: do not coalesce column for update/delete natural_using_join (#21085)](https://github.com/pingcap/tidb/pull/21103)                  |   | 6d18h   |
| tidb/21104 | [util/hint: fix ineffective index hint for delete / update in SQL Bind](https://github.com/pingcap/tidb/pull/21104)                                    |   | 6d18h   |
| tidb/21107 | [*: differentiate types for user variables (#18973)](https://github.com/pingcap/tidb/pull/21107)                                                       |   | 6d17h   |
| tidb/21124 | [planner: fix should not use point get plan](https://github.com/pingcap/tidb/pull/21124)                                                               |   | 5d10h   |
| tidb/21143 | [bindinfo: dbname check for bindings should be case insensitive](https://github.com/pingcap/tidb/pull/21143)                                           |   | 4d17h   |
| tidb/21146 | [planner: ban (index) merge join heuristically when convert eq cond to other cond (#21138)](https://github.com/pingcap/tidb/pull/21146)                |   | 4d16h   |
| tidb/21173 | [planner: fix partition pruning when condition exceeds the range of column type](https://github.com/pingcap/tidb/pull/21173)                           |   | 3d18h   |
| tidb/21213 | [planner: set dbName for hinted query block table alias](https://github.com/pingcap/tidb/pull/21213)                                                   |   | 18h     |
| tidb/21230 | [planner, executor: fix statement-level optimize hint invalid and memory tracker when `tryFastPlan` works](https://github.com/pingcap/tidb/pull/21230) |   | 10h     |


## Reviewed in Last 7 Days

|      REPO      |                                                                       PR                                                                        | C | D |   R   |
|----------------|-------------------------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/21221     | [planner: make index-hash-join and index-merge-join consider the collation in the right way](https://github.com/pingcap/tidb/pull/21221)        |   | 1 | 0h    |
| tidb/21219     | [planner: make index-hash-join and index-merge-join consider collation (#21201)](https://github.com/pingcap/tidb/pull/21219)                    |   | 1 | 1h    |
| tidb/21201     | [planner: make index-hash-join and index-merge-join consider collation](https://github.com/pingcap/tidb/pull/21201)                             |   | 1 | 3h    |
| tidb/21205     | [planner: `DELETE` cannot delete data in some cases when the database name is capitalized (#21202)](https://github.com/pingcap/tidb/pull/21205) |   | 1 | 0h    |
| tidb-test/1102 | [fix results for #19620 (#1101)](https://github.com/pingcap/tidb-test/pull/1102)                                                                |   | 4 | 8d0h  |
| tidb/21138     | [planner: ban (index) merge join heuristically when convert eq cond to other cond](https://github.com/pingcap/tidb/pull/21138)                  |   | 5 | 2h    |
| tidb/20686     | [*: support SQL bind for Update / Delete / Insert / Replace](https://github.com/pingcap/tidb/pull/20686)                                        |   | 7 | 20d3h |


</details> 


<details> 
  <summary>wshwsh12's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                           PR                                                                           | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/15462 | [executor: implement `graceHashJoin`](https://github.com/pingcap/tidb/pull/15462)                                                                      | Y | 250d17h |
| tidb/17052 | [[DNM] *: a prototype of readonly table](https://github.com/pingcap/tidb/pull/17052)                                                                   |   | 198d19h |
| tidb/17996 | [planner: push avg & distinct functions across join](https://github.com/pingcap/tidb/pull/17996)                                                       | Y | 163d10h |
| tidb/18742 | [*: Add memory information of executors if OOM action fires for debugging](https://github.com/pingcap/tidb/pull/18742)                                 |   | 123d16h |
| tidb/19807 | [executor: parallel evaluation for aggregate functions with distinct in hashAgg](https://github.com/pingcap/tidb/pull/19807)                           |   | 80d10h  |
| tidb/19957 | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                                 | Y | 73d13h  |
| tidb/20044 | [expression: Add column nullability checking before "refine args"](https://github.com/pingcap/tidb/pull/20044)                                         | Y | 68d7h   |
| tidb/20478 | [planner: add projection if expand virtual generated column](https://github.com/pingcap/tidb/pull/20478)                                               |   | 39d13h  |
| tidb/20503 | [expression: compatible with mysql's NO_ZERO_DATE](https://github.com/pingcap/tidb/pull/20503)                                                         |   | 36d11h  |
| tidb/20580 | [statistics: add bucket ndv for index histogram](https://github.com/pingcap/tidb/pull/20580)                                                           |   | 32d20h  |
| tidb/20664 | [executor, unistore: optimize analyze index by move it to analyze columns](https://github.com/pingcap/tidb/pull/20664)                                 |   | 27d18h  |
| tidb/20844 | [executor: introduce new variables to control Apply's behaviors and add more tests for it](https://github.com/pingcap/tidb/pull/20844)                 |   | 19d13h  |
| tidb/20861 | [executor:add runtime information for StreamAggExec](https://github.com/pingcap/tidb/pull/20861)                                                       |   | 18d19h  |
| tidb/20883 | [*: support select from tablesample](https://github.com/pingcap/tidb/pull/20883)                                                                       |   | 18d9h   |
| tidb/21100 | [*: support read only lock](https://github.com/pingcap/tidb/pull/21100)                                                                                |   | 6d19h   |
| tidb/21174 | [expression: enable coprocessor pushdown of function UUID](https://github.com/pingcap/tidb/pull/21174)                                                 |   | 3d18h   |
| tidb/21220 | [expression: fix compatibility behaviors in zero datetime with MySQL](https://github.com/pingcap/tidb/pull/21220)                                      |   | 16h     |
| tidb/21224 | [executor: fix the display of large unsigned handle when show table regions (#21026)](https://github.com/pingcap/tidb/pull/21224)                      |   | 14h     |
| tidb/21230 | [planner, executor: fix statement-level optimize hint invalid and memory tracker when `tryFastPlan` works](https://github.com/pingcap/tidb/pull/21230) |   | 10h     |


## Reviewed in Last 7 Days

|      REPO      |                                                                      PR                                                                      | C | D |   R    |
|----------------|----------------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/21169     | [expression: fix inaccurate error info for year column out of range (#20169)](https://github.com/pingcap/tidb/pull/21169)                    |   | 1 | 3d6h   |
| tidb/21057     | [hptc ycsb-worloade: Implement AppendRows for Chunk ](https://github.com/pingcap/tidb/pull/21057)                                            |   | 1 | 9d6h   |
| tidb/21168     | [store, executor: refactor ratelimitAction](https://github.com/pingcap/tidb/pull/21168)                                                      |   | 1 | 2d23h  |
| tidb/20886     | [*: optimize for encoding huge plan (#20811)](https://github.com/pingcap/tidb/pull/20886)                                                    |   | 5 | 13d1h  |
| tidb-test/1105 | [mysql_test: update test after @@tidb_constraint_check_in_place](https://github.com/pingcap/tidb-test/pull/1105)                             |   | 7 | 9h     |
| tidb/20939     | [executor: Make tidb_constraint_check_in_place session variable work for unique index](https://github.com/pingcap/tidb/pull/20939)           |   | 7 | 7d22h  |
| tidb/20413     | [execute: fill arguments when returning `ErrUnknownSystemVar` in `SetSessionSystemVar` (#20102)](https://github.com/pingcap/tidb/pull/20413) |   | 7 | 35d19h |
| tidb/20987     | [expression, executor: allow insert strings with overflowed trailing spaces](https://github.com/pingcap/tidb/pull/20987)                     |   | 7 | 5d21h  |


</details> 

