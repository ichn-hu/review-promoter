|   REVIEWER    |  O(C)   | 1 | 2 | 3 | 4 | 5 | 6 | 7 | T  |
|---------------|---------|---|---|---|---|---|---|---|----|
| Reminiscent   |  7 ( 0) | 2 | 3 | 0 | 0 | 0 | 0 | 0 |  5 |
| SunRunAway    | 54 ( 6) | 0 | 0 | 0 | 1 | 0 | 0 | 0 |  1 |
| XuHuaiyu      | 82 ( 4) | 1 | 0 | 0 | 1 | 0 | 0 | 0 |  2 |
| dyzsr         |  0 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| eurekaka      | 20 ( 5) | 3 | 3 | 2 | 2 | 0 | 0 | 0 | 10 |
| fzhedu        |  2 ( 2) | 1 | 0 | 0 | 0 | 0 | 0 | 0 |  1 |
| ichn-hu       |  3 ( 0) | 3 | 0 | 2 | 0 | 0 | 0 | 0 |  5 |
| lzmhhh123     | 48 ( 2) | 1 | 0 | 2 | 2 | 0 | 0 | 0 |  5 |
| nullnotnil    |  0 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| qw4990        | 63 ( 6) | 2 | 1 | 6 | 6 | 0 | 0 | 0 | 15 |
| rebelice      |  0 ( 0) | 1 | 0 | 0 | 0 | 0 | 0 | 0 |  1 |
| time-and-fate |  3 ( 0) | 1 | 0 | 0 | 0 | 0 | 0 | 0 |  1 |
| winoros       | 41 ( 4) | 0 | 0 | 2 | 2 | 0 | 0 | 0 |  4 |
| wshwsh12      | 16 ( 4) | 4 | 2 | 3 | 0 | 0 | 0 | 0 |  9 |


<details> 
  <summary>Reminiscent's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                               PR                                                                | C | LASTED |
|------------|---------------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/21137 | [executor: specially handle empty input for apply's outer child aggregate (#20544)](https://github.com/pingcap/tidb/pull/21137) |   | 49d20h |
| tidb/21550 | [planner : fix unsigned_decimal_col=-int_cnst access index (#21198)](https://github.com/pingcap/tidb/pull/21550)                |   | 30d20h |
| tidb/21614 | [planner: do not propagate column eq with different column types (#21495)](https://github.com/pingcap/tidb/pull/21614)          |   | 29d14h |
| tidb/21896 | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                      |   | 17d19h |
| tidb/21936 | [expression: fix wrong type inferring for ceiling function. (#21920)](https://github.com/pingcap/tidb/pull/21936)               |   | 16d17h |
| tidb/21957 | [planner: fix unknown columns in join using below agg (#21922)](https://github.com/pingcap/tidb/pull/21957)                     |   | 15d23h |
| tidb/21964 | [planner: add plancodec id for all type TableScan/IndexScan. (#21935)](https://github.com/pingcap/tidb/pull/21964)              |   | 15d18h |


## Reviewed in Last 7 Days

|    REPO    |                                                                  PR                                                                   | C | D |   R    |
|------------|---------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/22182 | [bindinfo: avoid duplicate bindings caused by concurrent baseline capture](https://github.com/pingcap/tidb/pull/22182)                |   | 1 | 1d19h  |
| tidb/22219 | [executor: store correct plan hint in statements_summary when log level is 'debug'](https://github.com/pingcap/tidb/pull/22219)       |   | 1 | 18h    |
| tidb/22222 | [planner: check error when correlatedAggregateResolver leaves ast.Node](https://github.com/pingcap/tidb/pull/22222)                   |   | 2 | 16h    |
| tidb/22126 | [*: add `sys` schema, `sys.SCHEMA_UNUSED_INDEXES` view and `sys.SCHEMA_INDEX_USAGE` view](https://github.com/pingcap/tidb/pull/22126) |   | 2 | 6d2h   |
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                             |   | 2 | 61d18h |


</details> 


<details> 
  <summary>SunRunAway's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                      PR                                                                       | C | LASTED  |
|--------------|-----------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| docs-cn/4913 | [explain: add indexes](https://github.com/pingcap/docs-cn/pull/4913)                                                                          |   | 52d18h  |
| tidb/15370   | [planner,executor: Refactor Shuffle and implement parallel Sort](https://github.com/pingcap/tidb/pull/15370)                                  | Y | 299d19h |
| docs-cn/4933 | [explain: add joins](https://github.com/pingcap/docs-cn/pull/4933)                                                                            |   | 48d20h  |
| tidb/15462   | [executor: implement `graceHashJoin`](https://github.com/pingcap/tidb/pull/15462)                                                             | Y | 295d18h |
| tidb/16967   | [executor: Refactor Shuffle and implement parallel sort (executor part)](https://github.com/pingcap/tidb/pull/16967)                          | Y | 250d10h |
| tidb/17238   | [*: refactor table.Allocator to improve readability](https://github.com/pingcap/tidb/pull/17238)                                              |   | 237d18h |
| tidb/19120   | [executor: Concurrently fetch chunks and insert them to a concurrent hash table in hash build](https://github.com/pingcap/tidb/pull/19120)    |   | 149d22h |
| tidb/19178   | [executor: Refactor probe channel](https://github.com/pingcap/tidb/pull/19178)                                                                |   | 147d17h |
| tidb/19347   | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347)          |   | 139d23h |
| tidb/19807   | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                                       |   | 125d11h |
| tidb/19900   | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                                | Y | 120d18h |
| tidb/20140   | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                            |   | 107d22h |
| tidb/20220   | [*: new secondary index value format](https://github.com/pingcap/tidb/pull/20220)                                                             |   | 104d17h |
| tidb/20316   | [docs/design: add design doc for index usage information](https://github.com/pingcap/tidb/pull/20316)                                         |   | 99d17h  |
| tidb/20335   | [planner, executor: enable inline projection for Selection](https://github.com/pingcap/tidb/pull/20335)                                       | Y | 96d18h  |
| tidb/20360   | [planner: refine explain info for batch cop](https://github.com/pingcap/tidb/pull/20360)                                                      |   | 90d22h  |
| tidb/20397   | [parser: replace ast.SelectLockInShareMode with ast.SelectLockForShare](https://github.com/pingcap/tidb/pull/20397)                           |   | 88d18h  |
| tidb/20615   | [utils: Avoid panic when getting memory](https://github.com/pingcap/tidb/pull/20615)                                                          |   | 76d2h   |
| tidb/20689   | [expression: make TIME function compatible with MySQL (#19158)](https://github.com/pingcap/tidb/pull/20689)                                   |   | 71d21h  |
| tidb/20752   | [*: trace statsCache and preparePlanCache by Global memory tracker.](https://github.com/pingcap/tidb/pull/20752)                              |   | 66d23h  |
| tidb/20765   | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                             |   | 66d17h  |
| tidb/21137   | [executor: specially handle empty input for apply's outer child aggregate (#20544)](https://github.com/pingcap/tidb/pull/21137)               |   | 49d20h  |
| tidb/21207   | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                               |   | 45d19h  |
| tidb/21277   | [executor: fix split table with large integers](https://github.com/pingcap/tidb/pull/21277)                                                   |   | 43d20h  |
| tidb/21310   | [types: convert string to MySQL BIT correctly](https://github.com/pingcap/tidb/pull/21310)                                                    |   | 42d22h  |
| tidb/21364   | [expression: Add test cases to cover the cases when invalid int value is casted as TIME (#18653)](https://github.com/pingcap/tidb/pull/21364) |   | 39d1h   |
| tidb/21381   | [*: optimize analyze cluster index table](https://github.com/pingcap/tidb/pull/21381)                                                         |   | 38d18h  |
| tidb/21386   | [expression: Disable cast decimal as string push down to TiFlash](https://github.com/pingcap/tidb/pull/21386)                                 |   | 38d16h  |
| tidb/21443   | [*: Let binary literal can be convert to enum and set (#20789)](https://github.com/pingcap/tidb/pull/21443)                                   |   | 36d14h  |
| tidb/21504   | [planner: fix invalid convert type in between...and... (#19820)](https://github.com/pingcap/tidb/pull/21504)                                  | Y | 34d15h  |
| tidb/21546   | [planner: do not push down the aggregation function with correlated column (#21453)](https://github.com/pingcap/tidb/pull/21546)              |   | 30d23h  |
| tidb/21573   | [expression: fix incorrect result of IsTrue function for time types (#21534)](https://github.com/pingcap/tidb/pull/21573)                     |   | 30d13h  |
| tidb/21810   | [expression: handle hybrid field types for where clause (#21724)](https://github.com/pingcap/tidb/pull/21810)                                 |   | 23d18h  |
| tidb/21813   | [expression: handle tp.flen overflow in to_base64 function (#20947)](https://github.com/pingcap/tidb/pull/21813)                              |   | 23d18h  |
| tidb/21834   | [planner: enhanced index range calculation plan](https://github.com/pingcap/tidb/pull/21834)                                                  |   | 22d19h  |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                   |   | 20d20h  |
| tidb/21877   | [planner: fix correlated aggregates which should be evaluated in outer query (#21431)](https://github.com/pingcap/tidb/pull/21877)            |   | 20d19h  |
| tidb/21878   | [planner: do not push down lock to pointGet/bacthPointGet when selection exists](https://github.com/pingcap/tidb/pull/21878)                  |   | 20d18h  |
| tidb/21890   | [*: redact some error code, part(3/3) (#21866)](https://github.com/pingcap/tidb/pull/21890)                                                   |   | 18d15h  |
| tidb/21936   | [expression: fix wrong type inferring for ceiling function. (#21920)](https://github.com/pingcap/tidb/pull/21936)                             |   | 16d17h  |
| tidb/21956   | [planner/preprocessor: disallow into-outfile clause in some place](https://github.com/pingcap/tidb/pull/21956)                                |   | 15d23h  |
| tidb/21960   | [types: Regard `TypeNewDecimal` as not a `hasVariantFieldLength` type. (#21849)](https://github.com/pingcap/tidb/pull/21960)                  |   | 15d21h  |
| tidb/22026   | [expression: separated arithmeticPlusIntSig](https://github.com/pingcap/tidb/pull/22026)                                                      |   | 13d21h  |
| tidb/22043   | [planner, executor: enhance the limit pushdown rule.](https://github.com/pingcap/tidb/pull/22043)                                             |   | 11d11h  |
| tidb/22089   | [executor: fix signed cluster index behavior (#22085)](https://github.com/pingcap/tidb/pull/22089)                                            |   | 8d23h   |
| tidb/22104   | [executor: fix incompatible escape behaviors in `select into outfile` (#22100)](https://github.com/pingcap/tidb/pull/22104)                   |   | 8d17h   |
| tidb/22106   | [executor: avoid log duplicate index name in slow-log (#22057)](https://github.com/pingcap/tidb/pull/22106)                                   |   | 8d14h   |
| tidb/22107   | [executor: avoid log duplicate index name in slow-log (#22057)](https://github.com/pingcap/tidb/pull/22107)                                   |   | 8d14h   |
| tidb/22114   | [test: fix globalkilltest (#21987)](https://github.com/pingcap/tidb/pull/22114)                                                               |   | 8d12h   |
| tidb/22120   | [executor: fix `update ignore` into not exists partition (#21984)](https://github.com/pingcap/tidb/pull/22120)                                |   | 7d23h   |
| tidb/22136   | [executor: improve the runtime stats of index lookup reader (#21982)](https://github.com/pingcap/tidb/pull/22136)                             |   | 7d17h   |
| tidb/22152   | [planner: check index valid while forUpdateRead](https://github.com/pingcap/tidb/pull/22152)                                                  |   | 3d19h   |
| tidb/22181   | [planner, expression: fix error when using IN combined with subquery (#22080)](https://github.com/pingcap/tidb/pull/22181)                    |   | 2d18h   |
| tidb/22217   | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                                 |   | 1d18h   |


## Reviewed in Last 7 Days

|   REPO    |                                     PR                                     | C | D |   R   |
|-----------|----------------------------------------------------------------------------|---|---|-------|
| docs/4565 | [tidb: add doc for global kill](https://github.com/pingcap/docs/pull/4565) |   | 4 | 2d14h |


</details> 


<details> 
  <summary>XuHuaiyu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                              PR                                                                              | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19292 | [planner: suppport left join in join reorder](https://github.com/pingcap/tidb/pull/19292)                                                                    |   | 141d17h |
| docs/4565  | [tidb: add doc for global kill](https://github.com/pingcap/docs/pull/4565)                                                                                   |   | 6d9h    |
| tidb/19900 | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                                               | Y | 120d18h |
| tidb/20040 | [planner, expression: take NullFlag into consideration when optimize the `int non-const` <cmp > `non-int const`](https://github.com/pingcap/tidb/pull/20040) | Y | 113d14h |
| tidb/20140 | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                                           |   | 107d22h |
| tidb/20311 | [expression: fix overflow error when convert bit to int64 (#20266)](https://github.com/pingcap/tidb/pull/20311)                                              |   | 99d21h  |
| tidb/20350 | [executor: support read global indexes in IndexMergeReader and index join](https://github.com/pingcap/tidb/pull/20350)                                       | Y | 93d14h  |
| tidb/20505 | [*: Add metrics for oom-action and sql memory usage.](https://github.com/pingcap/tidb/pull/20505)                                                            |   | 80d19h  |
| tidb/20576 | [*: fix stats feedback after tableReader handle multiple ranges](https://github.com/pingcap/tidb/pull/20576)                                                 |   | 78d13h  |
| tidb/20613 | [executor: fix issue of hash join fetch time inaccurate](https://github.com/pingcap/tidb/pull/20613)                                                         |   | 76d13h  |
| tidb/20752 | [*: trace statsCache and preparePlanCache by Global memory tracker.](https://github.com/pingcap/tidb/pull/20752)                                             |   | 66d23h  |
| tidb/20790 | [collation: add pinyin collation for chinese charset support](https://github.com/pingcap/tidb/pull/20790)                                                    |   | 65d21h  |
| tidb/20793 | [planner, executor: enable inline projection for Apply](https://github.com/pingcap/tidb/pull/20793)                                                          |   | 65d21h  |
| tidb/20905 | [planner: fix statement-optimize not work in `TryFastPlan`](https://github.com/pingcap/tidb/pull/20905)                                                      |   | 62d17h  |
| tidb/20972 | [expression: POC implementation of Vitess hashing algorithm.](https://github.com/pingcap/tidb/pull/20972)                                                    |   | 58d1h   |
| tidb/21064 | [planner, executor: fix cast not check error](https://github.com/pingcap/tidb/pull/21064)                                                                    |   | 53d9h   |
| tidb/21149 | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                                |   | 49d15h  |
| tidb/21228 | [executor: return the result immediately when combining LIMIT row_count with DISTINCT](https://github.com/pingcap/tidb/pull/21228)                           |   | 45d14h  |
| tidb/21304 | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                                     |   | 43d12h  |
| tidb/21334 | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                                                |   | 42d14h  |
| tidb/21340 | [executor: initialize expensive query handler on domain creation](https://github.com/pingcap/tidb/pull/21340)                                                |   | 42d0h   |
| tidb/21425 | [planner: natural join not consider rowid and null eq not propagate (#21328)](https://github.com/pingcap/tidb/pull/21425)                                    |   | 36d22h  |
| tidb/21473 | [ddl: check the generated column offset when modifies column (#21458)](https://github.com/pingcap/tidb/pull/21473)                                           |   | 35d17h  |
| tidb/21476 | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                        |   | 35d16h  |
| tidb/21477 | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21477)                                                        |   | 35d16h  |
| tidb/21483 | [executor, store/tikv: locks exist keys for point_get & batch_point_get (#21229)](https://github.com/pingcap/tidb/pull/21483)                                |   | 35d13h  |
| tidb/21488 | [planner: fix ambiguous field when resolve having expr  (#21165)](https://github.com/pingcap/tidb/pull/21488)                                                |   | 34d23h  |
| tidb/21504 | [planner: fix invalid convert type in between...and... (#19820)](https://github.com/pingcap/tidb/pull/21504)                                                 | Y | 34d15h  |
| tidb/21532 | [expression: set IsBooleanFlag for boolean scalar functions (#20706)](https://github.com/pingcap/tidb/pull/21532)                                            |   | 31d17h  |
| tidb/21536 | [executor: add slow-log file meta cache to avoid repeat read file meta information](https://github.com/pingcap/tidb/pull/21536)                              |   | 31d15h  |
| tidb/21550 | [planner : fix unsigned_decimal_col=-int_cnst access index (#21198)](https://github.com/pingcap/tidb/pull/21550)                                             |   | 30d20h  |
| tidb/21564 | [ddl: fix Incorrect behavior of NO_ZERO_DATE when altering table](https://github.com/pingcap/tidb/pull/21564)                                                |   | 30d16h  |
| tidb/21573 | [expression: fix incorrect result of IsTrue function for time types (#21534)](https://github.com/pingcap/tidb/pull/21573)                                    |   | 30d13h  |
| tidb/21590 | [expression: fix compatibility behaviors in sec_to_time with MySQL  (#21555)](https://github.com/pingcap/tidb/pull/21590)                                    |   | 29d21h  |
| tidb/21593 | [expression: fix convert number base for hybrid type (#21554)](https://github.com/pingcap/tidb/pull/21593)                                                   |   | 29d20h  |
| tidb/21602 | [expression: not evaluate time addition for timestamp with 2 args if 1st arg's year is zero (#21572)](https://github.com/pingcap/tidb/pull/21602)            |   | 29d18h  |
| tidb/21608 | [expression: fix error "invalid time format: '{0 0 0 0 0 0 0}'" for timestampAdd (#21591)](https://github.com/pingcap/tidb/pull/21608)                       |   | 29d16h  |
| tidb/21610 | [*: remove needless InInsertStmt (#19787)](https://github.com/pingcap/tidb/pull/21610)                                                                       |   | 29d15h  |
| tidb/21614 | [planner: do not propagate column eq with different column types (#21495)](https://github.com/pingcap/tidb/pull/21614)                                       |   | 29d14h  |
| tidb/21626 | [test: convert test to benchmard test to make ci stable (#21616)](https://github.com/pingcap/tidb/pull/21626)                                                |   | 28d23h  |
| tidb/21635 | [expression: handle invalid argument for addtime and subtime function  (#21600)](https://github.com/pingcap/tidb/pull/21635)                                 |   | 28d20h  |
| tidb/21673 | [expression, types: fix unexpected result from TIME() when fsp digits > 6 (#21652)](https://github.com/pingcap/tidb/pull/21673)                              |   | 27d17h  |
| tidb/21676 | [expression: fix compatibility of extract day_time unit functions (#21601)](https://github.com/pingcap/tidb/pull/21676)                                      |   | 27d17h  |
| tidb/21680 | [planner: report error when ORDER BY conflicts with DISTINCT (#21286)](https://github.com/pingcap/tidb/pull/21680)                                           |   | 27d16h  |
| tidb/21697 | [planner: check for only_full_group_by in ORDER BY and HAVING (#21216)](https://github.com/pingcap/tidb/pull/21697)                                          |   | 24d20h  |
| tidb/21711 | [expression: Fix unexpected panic when using IF function. (#21132)](https://github.com/pingcap/tidb/pull/21711)                                              |   | 24d17h  |
| tidb/21714 | [planner: fix the coercibility of the cast function (#21705)](https://github.com/pingcap/tidb/pull/21714)                                                    |   | 24d17h  |
| tidb/21718 | [types: fix compare object json type (#21703)](https://github.com/pingcap/tidb/pull/21718)                                                                   |   | 24d16h  |
| tidb/21785 | [types: fix compare float64 with float64 in json (#21709)](https://github.com/pingcap/tidb/pull/21785)                                                       |   | 23d22h  |
| tidb/21808 | [planner: fix the fail when we compare multi fields in the subquery (#21699)](https://github.com/pingcap/tidb/pull/21808)                                    |   | 23d18h  |
| tidb/21810 | [expression: handle hybrid field types for where clause (#21724)](https://github.com/pingcap/tidb/pull/21810)                                                |   | 23d18h  |
| tidb/21813 | [expression: handle tp.flen overflow in to_base64 function (#20947)](https://github.com/pingcap/tidb/pull/21813)                                             |   | 23d18h  |
| tidb/21839 | [planner/core: add 'split table using statistics' statement](https://github.com/pingcap/tidb/pull/21839)                                                     |   | 22d15h  |
| tidb/21842 | [planner: Shuffle hash agg](https://github.com/pingcap/tidb/pull/21842)                                                                                      |   | 22d11h  |
| tidb/21853 | [expression: fix compatibility behaviors in time_format with MySQL (#21559)](https://github.com/pingcap/tidb/pull/21853)                                     |   | 21d19h  |
| tidb/21870 | [types: report error for json object with key length >= 65536 (#21779)](https://github.com/pingcap/tidb/pull/21870)                                          |   | 20d23h  |
| tidb/21874 | [expression:truncate decimal value instead of return error (#21691)](https://github.com/pingcap/tidb/pull/21874)                                             |   | 20d21h  |
| tidb/21877 | [planner: fix correlated aggregates which should be evaluated in outer query (#21431)](https://github.com/pingcap/tidb/pull/21877)                           |   | 20d19h  |
| tidb/21896 | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                                                   |   | 17d19h  |
| tidb/21916 | [server: double type column from table should ignore its decimal (#21788)](https://github.com/pingcap/tidb/pull/21916)                                       |   | 16d23h  |
| tidb/21924 | [expression: fix type infer for tidb's builtin compare(least and greatest) (#21150)](https://github.com/pingcap/tidb/pull/21924)                             |   | 16d19h  |
| tidb/21936 | [expression: fix wrong type inferring for ceiling function. (#21920)](https://github.com/pingcap/tidb/pull/21936)                                            |   | 16d17h  |
| tidb/21957 | [planner: fix unknown columns in join using below agg (#21922)](https://github.com/pingcap/tidb/pull/21957)                                                  |   | 15d23h  |
| tidb/21958 | [expression: fix comparing json with string (#21903)](https://github.com/pingcap/tidb/pull/21958)                                                            |   | 15d22h  |
| tidb/21964 | [planner: add plancodec id for all type TableScan/IndexScan. (#21935)](https://github.com/pingcap/tidb/pull/21964)                                           |   | 15d18h  |
| tidb/21972 | [executor: throw error when prepared statement is execute, deallocate or prepare (#21962)](https://github.com/pingcap/tidb/pull/21972)                       |   | 15d16h  |
| tidb/22013 | [executor: fix unstable test Issue16696 (#22009)](https://github.com/pingcap/tidb/pull/22013)                                                                |   | 14d17h  |
| tidb/22014 | [executor: fix unstable test Issue16696 (#22009)](https://github.com/pingcap/tidb/pull/22014)                                                                |   | 14d17h  |
| tidb/22107 | [executor: avoid log duplicate index name in slow-log (#22057)](https://github.com/pingcap/tidb/pull/22107)                                                  |   | 8d14h   |
| tidb/22118 | [planner: check if columns count matches for batch point get in TryFastPlan (#22044)](https://github.com/pingcap/tidb/pull/22118)                            |   | 7d23h   |
| tidb/22119 | [executor: fix `update ignore` into not exists partition (#21984)](https://github.com/pingcap/tidb/pull/22119)                                               |   | 7d23h   |
| tidb/22120 | [executor: fix `update ignore` into not exists partition (#21984)](https://github.com/pingcap/tidb/pull/22120)                                               |   | 7d23h   |
| tidb/22131 | [privilege: remove leading and trailing space when create user and role](https://github.com/pingcap/tidb/pull/22131)                                         |   | 7d19h   |
| tidb/22136 | [executor: improve the runtime stats of index lookup reader (#21982)](https://github.com/pingcap/tidb/pull/22136)                                            |   | 7d17h   |
| tidb/22141 | [store: trace `loadRegion` to see the PD region cache loading (#22092)](https://github.com/pingcap/tidb/pull/22141)                                          |   | 4d0h    |
| tidb/22142 | [store: trace `loadRegion` to see the PD region cache loading (#22092)](https://github.com/pingcap/tidb/pull/22142)                                          |   | 4d0h    |
| tidb/22143 | [expression: return correct results for user variables of datetime type (#22078)](https://github.com/pingcap/tidb/pull/22143)                                |   | 3d23h   |
| tidb/22148 | [session: set process info before building plan (#22101)](https://github.com/pingcap/tidb/pull/22148)                                                        |   | 3d20h   |
| tidb/22149 | [session: set process info before building plan (#22101)](https://github.com/pingcap/tidb/pull/22149)                                                        |   | 3d20h   |
| tidb/22153 | [executor: refine bigint unsigned primary key duplicate error](https://github.com/pingcap/tidb/pull/22153)                                                   |   | 3d19h   |
| tidb/22186 | [executor: fix select into outfile with year type column has no data (#22175)](https://github.com/pingcap/tidb/pull/22186)                                   |   | 2d16h   |
| tidb/22221 | [executor: fix memTableReader decoding for old row format key-values](https://github.com/pingcap/tidb/pull/22221)                                            |   | 1d17h   |


## Reviewed in Last 7 Days

|    REPO    |                                              PR                                              | C | D |   R    |
|------------|----------------------------------------------------------------------------------------------|---|---|--------|
| tidb/21459 | [planner: push down projection for tiflash](https://github.com/pingcap/tidb/pull/21459)      |   | 1 | 34d23h |
| tidb/22101 | [session: set process info before building plan](https://github.com/pingcap/tidb/pull/22101) |   | 4 | 4d18h  |


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

|    REPO    |                                                                        PR                                                                        | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/14729 | [planner: fix constant propagation for PredicatePushDown](https://github.com/pingcap/tidb/pull/14729)                                            | Y | 331d18h |
| tidb/14831 | [planner/cascades: add implementationRule for IndexLookUpJoin](https://github.com/pingcap/tidb/pull/14831)                                       |   | 324d18h |
| tidb/15090 | [planner/cascades: refine the row count estimation of TiKV layer Selection](https://github.com/pingcap/tidb/pull/15090)                          |   | 310d18h |
| tidb/15157 | [planner/cascades: implement `HashCode` method for all the LogicalPlans](https://github.com/pingcap/tidb/pull/15157)                             | Y | 308d15h |
| tidb/15335 | [planner/cascades: add transformation rule PullAggregationUpApply & EliminateMaxOneRow](https://github.com/pingcap/tidb/pull/15335)              |   | 301d18h |
| tidb/15370 | [planner,executor: Refactor Shuffle and implement parallel Sort](https://github.com/pingcap/tidb/pull/15370)                                     | Y | 299d19h |
| tidb/17276 | [planner/cascades: add rule InjectProjectionBelowSort](https://github.com/pingcap/tidb/pull/17276)                                               | Y | 234d9h  |
| tidb/18882 | [planner, executor: add explain for `MetricSummaryTableExtractor`](https://github.com/pingcap/tidb/pull/18882)                                   | Y | 161d18h |
| tidb/19347 | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347)             |   | 139d23h |
| tidb/20580 | [statistics: add bucket ndv for index histogram](https://github.com/pingcap/tidb/pull/20580)                                                     |   | 77d21h  |
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                                        |   | 63d17h  |
| tidb/21444 | [planner: ignore anonymous index while tiflash replica is available](https://github.com/pingcap/tidb/pull/21444)                                 |   | 36d13h  |
| tidb/21488 | [planner: fix ambiguous field when resolve having expr  (#21165)](https://github.com/pingcap/tidb/pull/21488)                                    |   | 34d23h  |
| tidb/21573 | [expression: fix incorrect result of IsTrue function for time types (#21534)](https://github.com/pingcap/tidb/pull/21573)                        |   | 30d13h  |
| tidb/21680 | [planner: report error when ORDER BY conflicts with DISTINCT (#21286)](https://github.com/pingcap/tidb/pull/21680)                               |   | 27d16h  |
| tidb/21697 | [planner: check for only_full_group_by in ORDER BY and HAVING (#21216)](https://github.com/pingcap/tidb/pull/21697)                              |   | 24d20h  |
| tidb/21994 | [range: fix overflow value access index ](https://github.com/pingcap/tidb/pull/21994)                                                            |   | 14d23h  |
| tidb/22173 | [expression: fix unexpected panic when doing isNullRejected check](https://github.com/pingcap/tidb/pull/22173)                                   |   | 2d21h   |
| tidb/22222 | [planner: check error when correlatedAggregateResolver leaves ast.Node](https://github.com/pingcap/tidb/pull/22222)                              |   | 1d16h   |
| tidb/22255 | [planner: check convert outer join to inner join if prepare stmt meet prepared plan cache is enable](https://github.com/pingcap/tidb/pull/22255) |   | 21h     |


## Reviewed in Last 7 Days

|    REPO     |                                                             PR                                                              | C | D |   R    |
|-------------|-----------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/22202  | [executor: return error when region cache is invalid in fast analyze](https://github.com/pingcap/tidb/pull/22202)           |   | 1 | 1d3h   |
| tidb/22205  | [*: remove tidb-lightning pkg to fix unexpected config change (#22164)](https://github.com/pingcap/tidb/pull/22205)         |   | 1 | 1d0h   |
| tidb/22216  | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22216)               |   | 1 | 23h    |
| tidb/22218  | [*: bump tidb's parser to the latest version](https://github.com/pingcap/tidb/pull/22218)                                   |   | 2 | 0h     |
| tidb/21275  | [*: rewrite origin SQL with default DB for SQL bindings](https://github.com/pingcap/tidb/pull/21275)                        |   | 2 | 42d4h  |
| tidb/22169  | [statistics: fix stack overflow when use DNF expr immediately after add column](https://github.com/pingcap/tidb/pull/22169) |   | 2 | 23h    |
| tidb/14412  | [executor, planner: ON DUPLICATE UPDATE can refer to un-project col](https://github.com/pingcap/tidb/pull/14412)            |   | 3 | 362d0h |
| tidb/22080  | [planner, expression: fix error when using IN combined with subquery](https://github.com/pingcap/tidb/pull/22080)           |   | 3 | 6d17h  |
| tidb/22086  | [planner/core: fix a bug of adding enforcer.](https://github.com/pingcap/tidb/pull/22086)                                   |   | 4 | 5d17h  |
| parser/1146 | [yy_parser: enable supportWindowFunc by default (#986)](https://github.com/pingcap/parser/pull/1146)                        |   | 4 | 0h     |


</details> 


<details> 
  <summary>fzhedu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                   PR                                                   | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19845 | [expression:fix FORMAT compatibility issue #11206](https://github.com/pingcap/tidb/pull/19845)         | Y | 122d16h |
| tidb/20117 | [optimizer: fix issue on incorrect result of natural join](https://github.com/pingcap/tidb/pull/20117) | Y | 108d21h |


## Reviewed in Last 7 Days

|    REPO    |                                                 PR                                                 | C | D |   R   |
|------------|----------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/22053 | [execution: support explain analyze in mpp execution.](https://github.com/pingcap/tidb/pull/22053) |   | 1 | 9d22h |


</details> 


<details> 
  <summary>ichn-hu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                            PR                                                            | C | LASTED |
|------------|--------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/21676 | [expression: fix compatibility of extract day_time unit functions (#21601)](https://github.com/pingcap/tidb/pull/21676)  |   | 27d17h |
| tidb/21850 | [expression: add implicit eval int and real for function dayname (#21806)](https://github.com/pingcap/tidb/pull/21850)   |   | 21d19h |
| tidb/21853 | [expression: fix compatibility behaviors in time_format with MySQL (#21559)](https://github.com/pingcap/tidb/pull/21853) |   | 21d19h |


## Reviewed in Last 7 Days

|    REPO    |                                                          PR                                                           | C | D |   R   |
|------------|-----------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/22222 | [planner: check error when correlatedAggregateResolver leaves ast.Node](https://github.com/pingcap/tidb/pull/22222)   |   | 1 | 1d0h  |
| tidb/22202 | [executor: return error when region cache is invalid in fast analyze](https://github.com/pingcap/tidb/pull/22202)     |   | 1 | 1d3h  |
| tidb/22194 | [executor: fix `subquery returns more than 1 row` if multi-rows are same](https://github.com/pingcap/tidb/pull/22194) |   | 1 | 1d6h  |
| tidb/21310 | [types: convert string to MySQL BIT correctly](https://github.com/pingcap/tidb/pull/21310)                            |   | 3 | 40d2h |
| tidb/22175 | [executor: fix select into outfile with year type column has no data](https://github.com/pingcap/tidb/pull/22175)     |   | 3 | 0h    |


</details> 


<details> 
  <summary>lzmhhh123's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                           PR                                                                           | C | LASTED  |
|--------------|--------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/14729   | [planner: fix constant propagation for PredicatePushDown](https://github.com/pingcap/tidb/pull/14729)                                                  | Y | 331d18h |
| docs-cn/4913 | [explain: add indexes](https://github.com/pingcap/docs-cn/pull/4913)                                                                                   |   | 52d18h  |
| tidb/17414   | [add curCost based join reorder algorithm](https://github.com/pingcap/tidb/pull/17414)                                                                 |   | 226d18h |
| tidb/19347   | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347)                   |   | 139d23h |
| tidb/19698   | [*: update test cases to support new collation enabled by default](https://github.com/pingcap/tidb/pull/19698)                                         |   | 127d23h |
| tidb/20044   | [expression: Add column nullability checking before "refine args"](https://github.com/pingcap/tidb/pull/20044)                                         | Y | 113d7h  |
| tidb/20444   | [expression: add json_merge_patch](https://github.com/pingcap/tidb/pull/20444)                                                                         |   | 85d21h  |
| tidb/20465   | [expression: add uuidShortFunction](https://github.com/pingcap/tidb/pull/20465)                                                                        |   | 84d20h  |
| tidb/20505   | [*: Add metrics for oom-action and sql memory usage.](https://github.com/pingcap/tidb/pull/20505)                                                      |   | 80d19h  |
| tidb/20618   | [planner: fix update generated columns error](https://github.com/pingcap/tidb/pull/20618)                                                              |   | 75d20h  |
| tidb/20642   | [executor: modify admin executors to support partitioned table with global index](https://github.com/pingcap/tidb/pull/20642)                          |   | 73d16h  |
| tidb/20825   | [executor: add diagnosis rule to check Transparent Huge Pages(THP) enabled (#20611)](https://github.com/pingcap/tidb/pull/20825)                       |   | 64d19h  |
| tidb/20903   | [planner: fix confused and unnecessary double-projection in plans.](https://github.com/pingcap/tidb/pull/20903)                                        |   | 62d17h  |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                         |   | 56d17h  |
| tidb/21051   | [executor: change read slow-log file module to concurrent](https://github.com/pingcap/tidb/pull/21051)                                                 |   | 55d14h  |
| tidb/21137   | [executor: specially handle empty input for apply's outer child aggregate (#20544)](https://github.com/pingcap/tidb/pull/21137)                        |   | 49d20h  |
| tidb/21195   | [brie: integrate lightning to suport IMPORT statement](https://github.com/pingcap/tidb/pull/21195)                                                     |   | 45d23h  |
| tidb/21334   | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                                          |   | 42d14h  |
| tidb/21347   | [session: make rollback work on global variables](https://github.com/pingcap/tidb/pull/21347)                                                          |   | 41d20h  |
| tidb/21401   | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                                     |   | 38d12h  |
| tidb/21404   | [planner: fix unexpected bad plan when IndexJoin inner side estRow is 0. (#21084)](https://github.com/pingcap/tidb/pull/21404)                         |   | 37d23h  |
| tidb/21444   | [planner: ignore anonymous index while tiflash replica is available](https://github.com/pingcap/tidb/pull/21444)                                       |   | 36d13h  |
| tidb/21487   | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                                          |   | 35d5h   |
| tidb/21641   | [executor: Fix pessimistic lock doesn't work on the partition table for subquery/joins](https://github.com/pingcap/tidb/pull/21641)                    |   | 28d18h  |
| tidb/21651   | [planner: allow filter condition pushing down to IndexScan for prefix index](https://github.com/pingcap/tidb/pull/21651)                               |   | 28d14h  |
| tidb/21680   | [planner: report error when ORDER BY conflicts with DISTINCT (#21286)](https://github.com/pingcap/tidb/pull/21680)                                     |   | 27d16h  |
| tidb/21711   | [expression: Fix unexpected panic when using IF function. (#21132)](https://github.com/pingcap/tidb/pull/21711)                                        |   | 24d17h  |
| tidb/21808   | [planner: fix the fail when we compare multi fields in the subquery (#21699)](https://github.com/pingcap/tidb/pull/21808)                              |   | 23d18h  |
| tidb/21850   | [expression: add implicit eval int and real for function dayname (#21806)](https://github.com/pingcap/tidb/pull/21850)                                 |   | 21d19h  |
| tidb/21853   | [expression: fix compatibility behaviors in time_format with MySQL (#21559)](https://github.com/pingcap/tidb/pull/21853)                               |   | 21d19h  |
| tidb/21870   | [types: report error for json object with key length >= 65536 (#21779)](https://github.com/pingcap/tidb/pull/21870)                                    |   | 20d23h  |
| tidb/21877   | [planner: fix correlated aggregates which should be evaluated in outer query (#21431)](https://github.com/pingcap/tidb/pull/21877)                     |   | 20d19h  |
| tidb/21924   | [expression: fix type infer for tidb's builtin compare(least and greatest) (#21150)](https://github.com/pingcap/tidb/pull/21924)                       |   | 16d19h  |
| tidb/21954   | [planner/cascades: add rule `PushSelDownApply`](https://github.com/pingcap/tidb/pull/21954)                                                            |   | 15d23h  |
| tidb/21972   | [executor: throw error when prepared statement is execute, deallocate or prepare (#21962)](https://github.com/pingcap/tidb/pull/21972)                 |   | 15d16h  |
| tidb/22022   | [planner/codec: fix issue of decode plan error cause by without escape special char (#21937)](https://github.com/pingcap/tidb/pull/22022)              |   | 14d0h   |
| tidb/22089   | [executor: fix signed cluster index behavior (#22085)](https://github.com/pingcap/tidb/pull/22089)                                                     |   | 8d23h   |
| tidb/22124   | [planner: avoid using index_merge when there are multiple table filters (#22122)](https://github.com/pingcap/tidb/pull/22124)                          |   | 7d21h   |
| tidb/22126   | [*: add `sys` schema, `sys.SCHEMA_UNUSED_INDEXES` view and `sys.SCHEMA_INDEX_USAGE` view](https://github.com/pingcap/tidb/pull/22126)                  |   | 7d20h   |
| tidb/22130   | [planner: join reorder should not change the order of output columns (#16852)](https://github.com/pingcap/tidb/pull/22130)                             |   | 7d19h   |
| tidb/22137   | [expression: separated arithmeticModIntSig](https://github.com/pingcap/tidb/pull/22137)                                                                |   | 7d12h   |
| tidb/22148   | [session: set process info before building plan (#22101)](https://github.com/pingcap/tidb/pull/22148)                                                  |   | 3d20h   |
| tidb/22149   | [session: set process info before building plan (#22101)](https://github.com/pingcap/tidb/pull/22149)                                                  |   | 3d20h   |
| tidb/22174   | [expression, ddl: check the argument count for the generated column (#22154)](https://github.com/pingcap/tidb/pull/22174)                              |   | 2d21h   |
| tidb/22188   | [planner: do not use indexMerge when the path only use a single index (#22168)](https://github.com/pingcap/tidb/pull/22188)                            |   | 2d13h   |
| tidb/22191   | [expression: speed up Column.VecEvalReal by using MergeNulls](https://github.com/pingcap/tidb/pull/22191)                                              |   | 2d12h   |
| tidb/22251   | [expression, executor: fix runtime panic in WEIGHT_STRING function when the length of binary is too large](https://github.com/pingcap/tidb/pull/22251) |   | 21h     |
| tidb/22271   | [expression: handle duration type infer in least and greatest](https://github.com/pingcap/tidb/pull/22271)                                             |   | 17h     |


## Reviewed in Last 7 Days

|    REPO    |                                                        PR                                                        | C | D |   R   |
|------------|------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/22197 | [util: add test TestCodec](https://github.com/pingcap/tidb/pull/22197)                                           |   | 1 | 23h   |
| tidb/21310 | [types: convert string to MySQL BIT correctly](https://github.com/pingcap/tidb/pull/21310)                       |   | 3 | 40d2h |
| tidb/22154 | [expression, ddl: check the argument count for the generated column](https://github.com/pingcap/tidb/pull/22154) |   | 3 | 21h   |
| tidb/22073 | [executor: always decode the value first and then the handle](https://github.com/pingcap/tidb/pull/22073)        |   | 4 | 5d23h |
| tidb/22147 | [types: Add a limitation about float data type (#20929)](https://github.com/pingcap/tidb/pull/22147)             |   | 4 | 1h    |


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

|     REPO     |                                                                          PR                                                                          | C | LASTED  |
|--------------|------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/16305   | [expression: separate signatures for `ModInt`](https://github.com/pingcap/tidb/pull/16305)                                                           | Y | 270d0h  |
| docs-cn/4669 | [sql-optimization: extended statistics documentation](https://github.com/pingcap/docs-cn/pull/4669)                                                  |   | 86d17h  |
| tidb/16967   | [executor: Refactor Shuffle and implement parallel sort (executor part)](https://github.com/pingcap/tidb/pull/16967)                                 | Y | 250d10h |
| tidb/17396   | [types: improve StrToDate performance](https://github.com/pingcap/tidb/pull/17396)                                                                   | Y | 227d10h |
| tidb/18882   | [planner, executor: add explain for `MetricSummaryTableExtractor`](https://github.com/pingcap/tidb/pull/18882)                                       | Y | 161d18h |
| tidb/19029   | [types: fix unexpected NOT_NULL flags](https://github.com/pingcap/tidb/pull/19029)                                                                   |   | 154d22h |
| tidb/19120   | [executor: Concurrently fetch chunks and insert them to a concurrent hash table in hash build](https://github.com/pingcap/tidb/pull/19120)           |   | 149d22h |
| tidb/19292   | [planner: suppport left join in join reorder](https://github.com/pingcap/tidb/pull/19292)                                                            |   | 141d17h |
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                               | Y | 118d14h |
| tidb/20011   | [statistics: fix incorrect total count used in index selectivity computation](https://github.com/pingcap/tidb/pull/20011)                            |   | 114d15h |
| tidb/20316   | [docs/design: add design doc for index usage information](https://github.com/pingcap/tidb/pull/20316)                                                |   | 99d17h  |
| tidb/20354   | [planner: rename relational operators (#14575)](https://github.com/pingcap/tidb/pull/20354)                                                          | Y | 92d6h   |
| tidb/20399   | [*: make 'tidb_enable_change_column_type' available as a session variable](https://github.com/pingcap/tidb/pull/20399)                               |   | 88d16h  |
| tidb/20689   | [expression: make TIME function compatible with MySQL (#19158)](https://github.com/pingcap/tidb/pull/20689)                                          |   | 71d21h  |
| tidb/20708   | [*: separate auto_increment ID allocator from _tidb_rowid allocator](https://github.com/pingcap/tidb/pull/20708)                                     |   | 70d21h  |
| tidb/20972   | [expression: POC implementation of Vitess hashing algorithm.](https://github.com/pingcap/tidb/pull/20972)                                            |   | 58d1h   |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                       |   | 56d17h  |
| tidb/21137   | [executor: specially handle empty input for apply's outer child aggregate (#20544)](https://github.com/pingcap/tidb/pull/21137)                      |   | 49d20h  |
| tidb/21149   | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                        |   | 49d15h  |
| tidb/21304   | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                             |   | 43d12h  |
| tidb/21318   | [planner, expression: use the range of column types to simplify expressions](https://github.com/pingcap/tidb/pull/21318)                             |   | 42d19h  |
| tidb/21359   | [*: add runtime stats for split region statement](https://github.com/pingcap/tidb/pull/21359)                                                        |   | 41d13h  |
| tidb/21401   | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                                   |   | 38d12h  |
| tidb/21408   | [statistics: fix a bug which causes panic when using the clustered index and the new collation (#21379)](https://github.com/pingcap/tidb/pull/21408) |   | 37d20h  |
| tidb/21424   | [sessionctx: move set variable to sysvar struct](https://github.com/pingcap/tidb/pull/21424)                                                         |   | 37d5h   |
| tidb/21464   | [server: return results of ongoing queries when graceful shutdown (#19669)](https://github.com/pingcap/tidb/pull/21464)                              |   | 35d20h  |
| tidb/21471   | [session: fix ineffective EXPLAIN FOR CONNECTION statement (#21044)](https://github.com/pingcap/tidb/pull/21471)                                     |   | 35d17h  |
| tidb/21476   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                |   | 35d16h  |
| tidb/21477   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21477)                                                |   | 35d16h  |
| tidb/21508   | [execution: fix dayofweek('0000-00-00') behavior](https://github.com/pingcap/tidb/pull/21508)                                                        |   | 34d10h  |
| tidb/21525   | [expression: fix compatibility behaviors in zero datetime with MySQL (#21220)](https://github.com/pingcap/tidb/pull/21525)                           |   | 31d20h  |
| tidb/21610   | [*: remove needless InInsertStmt (#19787)](https://github.com/pingcap/tidb/pull/21610)                                                               |   | 29d15h  |
| tidb/21665   | [executor: fix LEAD and LAG's default value can not adapt to field type (#20747)](https://github.com/pingcap/tidb/pull/21665)                        |   | 27d19h  |
| tidb/21680   | [planner: report error when ORDER BY conflicts with DISTINCT (#21286)](https://github.com/pingcap/tidb/pull/21680)                                   |   | 27d16h  |
| tidb/21711   | [expression: Fix unexpected panic when using IF function. (#21132)](https://github.com/pingcap/tidb/pull/21711)                                      |   | 24d17h  |
| tidb/21842   | [planner: Shuffle hash agg](https://github.com/pingcap/tidb/pull/21842)                                                                              |   | 22d11h  |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                          |   | 20d20h  |
| tidb/21887   | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                                      |   | 19d11h  |
| tidb/21895   | [executor: fix load data in file get wrong result #20854](https://github.com/pingcap/tidb/pull/21895)                                                |   | 17d20h  |
| tidb/21924   | [expression: fix type infer for tidb's builtin compare(least and greatest) (#21150)](https://github.com/pingcap/tidb/pull/21924)                     |   | 16d19h  |
| tidb/21930   | [planner: propagate NDV of column groups across plan nodes (#17854)](https://github.com/pingcap/tidb/pull/21930)                                     |   | 16d18h  |
| tidb/21969   | [types:  Add a limitation about float data type (#20929)](https://github.com/pingcap/tidb/pull/21969)                                                |   | 15d17h  |
| tidb/21971   | [executor: fix `insert ignore` into not exists partition (#21904)](https://github.com/pingcap/tidb/pull/21971)                                       |   | 15d17h  |
| tidb/21977   | [expression: log functions that can not be pushed to cop](https://github.com/pingcap/tidb/pull/21977)                                                |   | 15d16h  |
| tidb/22021   | [distsql: fix cop stats string display when there is only 1 rpc (#21901) (#21999)](https://github.com/pingcap/tidb/pull/22021)                       |   | 14d0h   |
| tidb/22048   | [session: support validating read-only statement during read-only staleness transaction](https://github.com/pingcap/tidb/pull/22048)                 |   | 10d21h  |
| tidb/22090   | [planner: push aggregation operators down to projection and union by default](https://github.com/pingcap/tidb/pull/22090)                            |   | 8d22h   |
| tidb/22104   | [executor: fix incompatible escape behaviors in `select into outfile` (#22100)](https://github.com/pingcap/tidb/pull/22104)                          |   | 8d17h   |
| tidb/22106   | [executor: avoid log duplicate index name in slow-log (#22057)](https://github.com/pingcap/tidb/pull/22106)                                          |   | 8d14h   |
| tidb/22107   | [executor: avoid log duplicate index name in slow-log (#22057)](https://github.com/pingcap/tidb/pull/22107)                                          |   | 8d14h   |
| tidb/22110   | [config, session: promise the compatibility of oom-action when upgrading (#22102)](https://github.com/pingcap/tidb/pull/22110)                       |   | 8d13h   |
| tidb/22118   | [planner: check if columns count matches for batch point get in TryFastPlan (#22044)](https://github.com/pingcap/tidb/pull/22118)                    |   | 7d23h   |
| tidb/22127   | [*: support ALTER TABLE ADD / DROP TIDB_STATS](https://github.com/pingcap/tidb/pull/22127)                                                           |   | 7d20h   |
| tidb/22136   | [executor: improve the runtime stats of index lookup reader (#21982)](https://github.com/pingcap/tidb/pull/22136)                                    |   | 7d17h   |
| tidb/22143   | [expression: return correct results for user variables of datetime type (#22078)](https://github.com/pingcap/tidb/pull/22143)                        |   | 3d23h   |
| tidb/22146   | [executor: forbid SFU on view](https://github.com/pingcap/tidb/pull/22146)                                                                           |   | 3d22h   |
| tidb/22182   | [bindinfo: avoid duplicate bindings caused by concurrent baseline capture](https://github.com/pingcap/tidb/pull/22182)                               |   | 2d18h   |
| tidb/22217   | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                                        |   | 1d18h   |
| tidb/22219   | [executor: store correct plan hint in statements_summary when log level is 'debug'](https://github.com/pingcap/tidb/pull/22219)                      |   | 1d17h   |
| tidb/22234   | [executor, planner: ON DUPLICATE UPDATE can refer to un-project col (#14412)](https://github.com/pingcap/tidb/pull/22234)                            |   | 1d15h   |
| tidb/22240   | [infoschema: support query partition_id from infoschema.partitions](https://github.com/pingcap/tidb/pull/22240)                                      |   | 1d13h   |
| tidb/22261   | [time: fix parse datetime won't truncate the reluctant string (#22232)](https://github.com/pingcap/tidb/pull/22261)                                  |   | 19h     |
| tidb/22279   | [oracle: fix GetStaleTimestamp get stuck error](https://github.com/pingcap/tidb/pull/22279)                                                          |   | 13h     |


## Reviewed in Last 7 Days

|    REPO    |                                                                  PR                                                                   | C | D |   R    |
|------------|---------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/22216 | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22216)                         |   | 1 | 23h    |
| tidb/22173 | [expression: fix unexpected panic when doing isNullRejected check](https://github.com/pingcap/tidb/pull/22173)                        |   | 1 | 2d2h   |
| tidb/14412 | [executor, planner: ON DUPLICATE UPDATE can refer to un-project col](https://github.com/pingcap/tidb/pull/14412)                      |   | 2 | 363d3h |
| tidb/22168 | [planner: do not use indexMerge when the path only use a single index](https://github.com/pingcap/tidb/pull/22168)                    |   | 3 | 8h     |
| tidb/22126 | [*: add `sys` schema, `sys.SCHEMA_UNUSED_INDEXES` view and `sys.SCHEMA_INDEX_USAGE` view](https://github.com/pingcap/tidb/pull/22126) |   | 3 | 5d5h   |
| tidb/22169 | [statistics: fix stack overflow when use DNF expr immediately after add column](https://github.com/pingcap/tidb/pull/22169)           |   | 3 | 7h     |
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                             |   | 3 | 61d2h  |
| tidb/22080 | [planner, expression: fix error when using IN combined with subquery](https://github.com/pingcap/tidb/pull/22080)                     |   | 3 | 6d21h  |
| tidb/20580 | [statistics: add bucket ndv for index histogram](https://github.com/pingcap/tidb/pull/20580)                                          |   | 3 | 75d1h  |
| tidb/22156 | [mod: update parser to latest 4.0 version](https://github.com/pingcap/tidb/pull/22156)                                                |   | 4 | 4h     |
| tidb/21275 | [*: rewrite origin SQL with default DB for SQL bindings](https://github.com/pingcap/tidb/pull/21275)                                  |   | 4 | 40d4h  |
| tidb/22033 | [statistics: redesign the schema for `mysql.stats_extended`](https://github.com/pingcap/tidb/pull/22033)                              |   | 4 | 9d23h  |
| tidb/21712 | [statistics: no more counting feedback if it is invalid](https://github.com/pingcap/tidb/pull/21712)                                  |   | 4 | 20d23h |
| tidb/20750 | [executor, infoschema, planner: optimize query cluster_slow_query](https://github.com/pingcap/tidb/pull/20750)                        |   | 4 | 63d0h  |
| tidb/22078 | [expression: return correct results for user variables of datetime type](https://github.com/pingcap/tidb/pull/22078)                  |   | 4 | 5d17h  |


</details> 


<details> 
  <summary>rebelice's detailed review report</summary> 

## To Be Reviewed

| REPO | PR | C | LASTED |
|------|----|---|--------|


## Reviewed in Last 7 Days

|    REPO    |                                           PR                                            | C | D |   R    |
|------------|-----------------------------------------------------------------------------------------|---|---|--------|
| tidb/21459 | [planner: push down projection for tiflash](https://github.com/pingcap/tidb/pull/21459) |   | 1 | 34d22h |


</details> 


<details> 
  <summary>time-and-fate's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                       PR                                                       | C | LASTED |
|------------|----------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                      |   | 63d17h |
| tidb/22006 | [config: disable statistics feedback by default (#21923)](https://github.com/pingcap/tidb/pull/22006)          |   | 14d19h |
| tidb/22173 | [expression: fix unexpected panic when doing isNullRejected check](https://github.com/pingcap/tidb/pull/22173) |   | 2d21h  |


## Reviewed in Last 7 Days

|    REPO    |                                              PR                                              | C | D |   R   |
|------------|----------------------------------------------------------------------------------------------|---|---|-------|
| tidb/20580 | [statistics: add bucket ndv for index histogram](https://github.com/pingcap/tidb/pull/20580) |   | 1 | 77d0h |


</details> 


<details> 
  <summary>winoros's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                          PR                                                                          | C | LASTED  |
|--------------|------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/14424   | [expression: add nullable() method to check whether an expression can return null](https://github.com/pingcap/tidb/pull/14424)                       |   | 364d18h |
| docs-cn/4669 | [sql-optimization: extended statistics documentation](https://github.com/pingcap/docs-cn/pull/4669)                                                  |   | 86d17h  |
| tidb/14831   | [planner/cascades: add implementationRule for IndexLookUpJoin](https://github.com/pingcap/tidb/pull/14831)                                           |   | 324d18h |
| tidb/15090   | [planner/cascades: refine the row count estimation of TiKV layer Selection](https://github.com/pingcap/tidb/pull/15090)                              |   | 310d18h |
| tidb/15157   | [planner/cascades: implement `HashCode` method for all the LogicalPlans](https://github.com/pingcap/tidb/pull/15157)                                 | Y | 308d15h |
| tidb/15426   | [planner/cascades: add transformation rule PushSelDownApply & refactor PushSelDownJoin](https://github.com/pingcap/tidb/pull/15426)                  |   | 296d17h |
| tidb/16967   | [executor: Refactor Shuffle and implement parallel sort (executor part)](https://github.com/pingcap/tidb/pull/16967)                                 | Y | 250d10h |
| tidb/17414   | [add curCost based join reorder algorithm](https://github.com/pingcap/tidb/pull/17414)                                                               |   | 226d18h |
| tidb/17996   | [planner: push avg & distinct functions across join](https://github.com/pingcap/tidb/pull/17996)                                                     | Y | 208d11h |
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                               | Y | 118d14h |
| tidb/20011   | [statistics: fix incorrect total count used in index selectivity computation](https://github.com/pingcap/tidb/pull/20011)                            |   | 114d15h |
| tidb/20311   | [expression: fix overflow error when convert bit to int64 (#20266)](https://github.com/pingcap/tidb/pull/20311)                                      |   | 99d21h  |
| tidb/20765   | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                                    |   | 66d17h  |
| tidb/20877   | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                                            |   | 63d17h  |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                       |   | 56d17h  |
| tidb/21137   | [executor: specially handle empty input for apply's outer child aggregate (#20544)](https://github.com/pingcap/tidb/pull/21137)                      |   | 49d20h  |
| tidb/21207   | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                                      |   | 45d19h  |
| tidb/21230   | [planner, executor: fix haven't track the memory usage of PointGet/BatchPointGet](https://github.com/pingcap/tidb/pull/21230)                        |   | 45d11h  |
| tidb/21408   | [statistics: fix a bug which causes panic when using the clustered index and the new collation (#21379)](https://github.com/pingcap/tidb/pull/21408) |   | 37d20h  |
| tidb/21425   | [planner: natural join not consider rowid and null eq not propagate (#21328)](https://github.com/pingcap/tidb/pull/21425)                            |   | 36d22h  |
| tidb/21476   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                |   | 35d16h  |
| tidb/21477   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21477)                                                |   | 35d16h  |
| tidb/21487   | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                                        |   | 35d5h   |
| tidb/21614   | [planner: do not propagate column eq with different column types (#21495)](https://github.com/pingcap/tidb/pull/21614)                               |   | 29d14h  |
| tidb/21714   | [planner: fix the coercibility of the cast function (#21705)](https://github.com/pingcap/tidb/pull/21714)                                            |   | 24d17h  |
| tidb/21808   | [planner: fix the fail when we compare multi fields in the subquery (#21699)](https://github.com/pingcap/tidb/pull/21808)                            |   | 23d18h  |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                          |   | 20d20h  |
| tidb/21877   | [planner: fix correlated aggregates which should be evaluated in outer query (#21431)](https://github.com/pingcap/tidb/pull/21877)                   |   | 20d19h  |
| tidb/21930   | [planner: propagate NDV of column groups across plan nodes (#17854)](https://github.com/pingcap/tidb/pull/21930)                                     |   | 16d18h  |
| tidb/21957   | [planner: fix unknown columns in join using below agg (#21922)](https://github.com/pingcap/tidb/pull/21957)                                          |   | 15d23h  |
| tidb/21964   | [planner: add plancodec id for all type TableScan/IndexScan. (#21935)](https://github.com/pingcap/tidb/pull/21964)                                   |   | 15d18h  |
| tidb/21976   | [planner: report error for invalid window specs which are not used (#21083)](https://github.com/pingcap/tidb/pull/21976)                             |   | 15d16h  |
| tidb/22022   | [planner/codec: fix issue of decode plan error cause by without escape special char (#21937)](https://github.com/pingcap/tidb/pull/22022)            |   | 14d0h   |
| tidb/22090   | [planner: push aggregation operators down to projection and union by default](https://github.com/pingcap/tidb/pull/22090)                            |   | 8d22h   |
| tidb/22118   | [planner: check if columns count matches for batch point get in TryFastPlan (#22044)](https://github.com/pingcap/tidb/pull/22118)                    |   | 7d23h   |
| tidb/22127   | [*: support ALTER TABLE ADD / DROP TIDB_STATS](https://github.com/pingcap/tidb/pull/22127)                                                           |   | 7d20h   |
| tidb/22143   | [expression: return correct results for user variables of datetime type (#22078)](https://github.com/pingcap/tidb/pull/22143)                        |   | 3d23h   |
| tidb/22173   | [expression: fix unexpected panic when doing isNullRejected check](https://github.com/pingcap/tidb/pull/22173)                                       |   | 2d21h   |
| tidb/22182   | [bindinfo: avoid duplicate bindings caused by concurrent baseline capture](https://github.com/pingcap/tidb/pull/22182)                               |   | 2d18h   |
| tidb/22219   | [executor: store correct plan hint in statements_summary when log level is 'debug'](https://github.com/pingcap/tidb/pull/22219)                      |   | 1d17h   |
| tidb/22255   | [planner: check convert outer join to inner join if prepare stmt meet prepared plan cache is enable](https://github.com/pingcap/tidb/pull/22255)     |   | 21h     |


## Reviewed in Last 7 Days

|    REPO     |                                              PR                                               | C | D |   R    |
|-------------|-----------------------------------------------------------------------------------------------|---|---|--------|
| parser/1138 | [*: create / drop extended stats by ALTER TABLE](https://github.com/pingcap/parser/pull/1138) |   | 3 | 11d20h |
| tidb/22156  | [mod: update parser to latest 4.0 version](https://github.com/pingcap/tidb/pull/22156)        |   | 3 | 22h    |
| tidb/22086  | [planner/core: fix a bug of adding enforcer.](https://github.com/pingcap/tidb/pull/22086)     |   | 4 | 6d3h   |
| tidb/21014  | [statistics: GC index usage information](https://github.com/pingcap/tidb/pull/21014)          |   | 4 | 53d0h  |


</details> 


<details> 
  <summary>wshwsh12's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                           PR                                                                           | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/15462 | [executor: implement `graceHashJoin`](https://github.com/pingcap/tidb/pull/15462)                                                                      | Y | 295d18h |
| tidb/17996 | [planner: push avg & distinct functions across join](https://github.com/pingcap/tidb/pull/17996)                                                       | Y | 208d11h |
| tidb/19807 | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                                                |   | 125d11h |
| tidb/19957 | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                                 | Y | 118d14h |
| tidb/20044 | [expression: Add column nullability checking before "refine args"](https://github.com/pingcap/tidb/pull/20044)                                         | Y | 113d7h  |
| tidb/21381 | [*: optimize analyze cluster index table](https://github.com/pingcap/tidb/pull/21381)                                                                  |   | 38d18h  |
| tidb/21487 | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                                          |   | 35d5h   |
| tidb/21541 | [executor: Nested prepare stmt should not be prepared](https://github.com/pingcap/tidb/pull/21541)                                                     |   | 31d13h  |
| tidb/21839 | [planner/core: add 'split table using statistics' statement](https://github.com/pingcap/tidb/pull/21839)                                               |   | 22d15h  |
| tidb/21887 | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                                        |   | 19d11h  |
| tidb/21945 | [distsql: fix cop stats string display when there is only 1 rpc (#21901)](https://github.com/pingcap/tidb/pull/21945)                                  |   | 16d14h  |
| tidb/21957 | [planner: fix unknown columns in join using below agg (#21922)](https://github.com/pingcap/tidb/pull/21957)                                            |   | 15d23h  |
| tidb/22110 | [config, session: promise the compatibility of oom-action when upgrading (#22102)](https://github.com/pingcap/tidb/pull/22110)                         |   | 8d13h   |
| tidb/22194 | [executor: fix `subquery returns more than 1 row` if multi-rows are same](https://github.com/pingcap/tidb/pull/22194)                                  |   | 1d23h   |
| tidb/22251 | [expression, executor: fix runtime panic in WEIGHT_STRING function when the length of binary is too large](https://github.com/pingcap/tidb/pull/22251) |   | 21h     |
| tidb/22269 | [executor: check storage.block-cache.capacity value](https://github.com/pingcap/tidb/pull/22269)                                                       |   | 17h     |


## Reviewed in Last 7 Days

|      REPO      |                                                                           PR                                                                           | C | D |   R    |
|----------------|--------------------------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/22197     | [util: add test TestCodec](https://github.com/pingcap/tidb/pull/22197)                                                                                 |   | 1 | 1d2h   |
| tidb/22251     | [expression, executor: fix runtime panic in WEIGHT_STRING function when the length of binary is too large](https://github.com/pingcap/tidb/pull/22251) |   | 1 | 1h     |
| tidb/22232     | [time: fix parse datetime won't truncate the reluctant string](https://github.com/pingcap/tidb/pull/22232)                                             |   | 1 | 19h    |
| tidb-test/1150 | [mysql-test: fix parse timestamp](https://github.com/pingcap/tidb-test/pull/1150)                                                                      |   | 1 | 1h     |
| tidb/21155     | [util/chunk: fix slice out of bound panic](https://github.com/pingcap/tidb/pull/21155)                                                                 |   | 2 | 47d15h |
| tidb/22183     | [executor: fix data too long in union statement with select null and select varchar](https://github.com/pingcap/tidb/pull/22183)                       |   | 2 | 20h    |
| tidb/21916     | [server: double type column from table should ignore its decimal (#21788)](https://github.com/pingcap/tidb/pull/21916)                                 |   | 3 | 14d23h |
| tidb/21969     | [types:  Add a limitation about float data type (#20929)](https://github.com/pingcap/tidb/pull/21969)                                                  |   | 3 | 13d17h |
| tidb/22084     | [expression: set collation function flen](https://github.com/pingcap/tidb/pull/22084)                                                                  |   | 3 | 6d21h  |


</details> 

