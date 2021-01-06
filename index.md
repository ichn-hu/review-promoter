|   REVIEWER    |  O(C)   | 1 | 2 | 3 | 4 | 5 | 6 | 7 | T  |
|---------------|---------|---|---|---|---|---|---|---|----|
| Reminiscent   |  8 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 3 |  3 |
| SunRunAway    | 53 ( 6) | 0 | 1 | 0 | 0 | 0 | 0 | 0 |  1 |
| XuHuaiyu      | 83 ( 4) | 0 | 1 | 0 | 0 | 0 | 0 | 9 | 10 |
| dyzsr         |  0 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| eurekaka      | 18 ( 5) | 4 | 2 | 0 | 0 | 0 | 3 | 0 |  9 |
| fzhedu        |  2 ( 2) | 1 | 0 | 0 | 0 | 0 | 0 | 0 |  1 |
| ichn-hu       |  4 ( 0) | 2 | 0 | 0 | 0 | 0 | 0 | 0 |  2 |
| lzmhhh123     | 48 ( 2) | 2 | 2 | 0 | 0 | 0 | 3 | 2 |  9 |
| nullnotnil    |  0 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| qw4990        | 59 ( 6) | 7 | 6 | 0 | 0 | 0 | 4 | 5 | 22 |
| rebelice      |  0 ( 0) | 0 | 1 | 0 | 0 | 0 | 0 | 0 |  1 |
| time-and-fate |  2 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 2 |  2 |
| winoros       | 39 ( 4) | 2 | 2 | 0 | 0 | 0 | 1 | 2 |  7 |
| wshwsh12      | 14 ( 4) | 3 | 0 | 0 | 0 | 0 | 0 | 6 |  9 |


<details> 
  <summary>Reminiscent's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                               PR                                                                | C | LASTED |
|------------|---------------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/21137 | [executor: specially handle empty input for apply's outer child aggregate (#20544)](https://github.com/pingcap/tidb/pull/21137) |   | 47d20h |
| tidb/21550 | [planner : fix unsigned_decimal_col=-int_cnst access index (#21198)](https://github.com/pingcap/tidb/pull/21550)                |   | 28d20h |
| tidb/21614 | [planner: do not propagate column eq with different column types (#21495)](https://github.com/pingcap/tidb/pull/21614)          |   | 27d14h |
| tidb/21896 | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                      |   | 15d19h |
| tidb/21936 | [expression: fix wrong type inferring for ceiling function. (#21920)](https://github.com/pingcap/tidb/pull/21936)               |   | 14d17h |
| tidb/21957 | [planner: fix unknown columns in join using below agg (#21922)](https://github.com/pingcap/tidb/pull/21957)                     |   | 13d23h |
| tidb/21964 | [planner: add plancodec id for all type TableScan/IndexScan. (#21935)](https://github.com/pingcap/tidb/pull/21964)              |   | 13d18h |
| tidb/22182 | [bindinfo: avoid duplicate bindings caused by concurrent baseline capture](https://github.com/pingcap/tidb/pull/22182)          |   | 18h    |


## Reviewed in Last 7 Days

|    REPO    |                                                             PR                                                              | C | D | R  |
|------------|-----------------------------------------------------------------------------------------------------------------------------|---|---|----|
| tidb/22104 | [executor: fix incompatible escape behaviors in `select into outfile` (#22100)](https://github.com/pingcap/tidb/pull/22104) |   | 7 | 0h |
| tidb/22103 | [executor: fix incompatible escape behaviors in `select into outfile` (#22100)](https://github.com/pingcap/tidb/pull/22103) |   | 7 | 0h |
| tidb/22100 | [executor: fix incompatible escape behaviors in `select into outfile`](https://github.com/pingcap/tidb/pull/22100)          |   | 7 | 0h |


</details> 


<details> 
  <summary>SunRunAway's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                      PR                                                                       | C | LASTED  |
|--------------|-----------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| docs-cn/4913 | [explain: add indexes](https://github.com/pingcap/docs-cn/pull/4913)                                                                          |   | 50d18h  |
| tidb/15370   | [planner,executor: Refactor Shuffle and implement parallel Sort](https://github.com/pingcap/tidb/pull/15370)                                  | Y | 297d19h |
| docs-cn/4933 | [explain: add joins](https://github.com/pingcap/docs-cn/pull/4933)                                                                            |   | 46d20h  |
| tidb/15462   | [executor: implement `graceHashJoin`](https://github.com/pingcap/tidb/pull/15462)                                                             | Y | 293d18h |
| tidb/16967   | [executor: Refactor Shuffle and implement parallel sort (executor part)](https://github.com/pingcap/tidb/pull/16967)                          | Y | 248d10h |
| tidb/17238   | [*: refactor table.Allocator to improve readability](https://github.com/pingcap/tidb/pull/17238)                                              |   | 235d18h |
| tidb/19120   | [executor: Concurrently fetch chunks and insert them to a concurrent hash table in hash build](https://github.com/pingcap/tidb/pull/19120)    |   | 147d21h |
| tidb/19178   | [executor: Refactor probe channel](https://github.com/pingcap/tidb/pull/19178)                                                                |   | 145d17h |
| tidb/19347   | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347)          |   | 137d23h |
| tidb/19807   | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                                       |   | 123d11h |
| tidb/19900   | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                                | Y | 118d18h |
| tidb/20140   | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                            |   | 105d22h |
| tidb/20220   | [*: new secondary index value format](https://github.com/pingcap/tidb/pull/20220)                                                             |   | 102d17h |
| tidb/20316   | [docs/design: add design doc for index usage information](https://github.com/pingcap/tidb/pull/20316)                                         |   | 97d17h  |
| tidb/20335   | [planner, executor: enable inline projection for Selection](https://github.com/pingcap/tidb/pull/20335)                                       | Y | 94d18h  |
| tidb/20360   | [planner: refine explain info for batch cop](https://github.com/pingcap/tidb/pull/20360)                                                      |   | 88d22h  |
| tidb/20397   | [parser: replace ast.SelectLockInShareMode with ast.SelectLockForShare](https://github.com/pingcap/tidb/pull/20397)                           |   | 86d18h  |
| tidb/20615   | [utils: Avoid panic when getting memory](https://github.com/pingcap/tidb/pull/20615)                                                          |   | 74d2h   |
| tidb/20689   | [expression: make TIME function compatible with MySQL (#19158)](https://github.com/pingcap/tidb/pull/20689)                                   |   | 69d21h  |
| tidb/20752   | [*: trace statsCache and preparePlanCache by Global memory tracker.](https://github.com/pingcap/tidb/pull/20752)                              |   | 64d23h  |
| tidb/20765   | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                             |   | 64d17h  |
| tidb/21137   | [executor: specially handle empty input for apply's outer child aggregate (#20544)](https://github.com/pingcap/tidb/pull/21137)               |   | 47d20h  |
| tidb/21207   | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                               |   | 43d19h  |
| tidb/21277   | [executor: fix split table with large integers](https://github.com/pingcap/tidb/pull/21277)                                                   |   | 41d20h  |
| tidb/21310   | [types: convert string to MySQL BIT correctly](https://github.com/pingcap/tidb/pull/21310)                                                    |   | 40d22h  |
| tidb/21364   | [expression: Add test cases to cover the cases when invalid int value is casted as TIME (#18653)](https://github.com/pingcap/tidb/pull/21364) |   | 37d1h   |
| tidb/21381   | [*: optimize analyze cluster index table](https://github.com/pingcap/tidb/pull/21381)                                                         |   | 36d17h  |
| tidb/21386   | [expression: Disable cast decimal as string push down to TiFlash](https://github.com/pingcap/tidb/pull/21386)                                 |   | 36d16h  |
| tidb/21443   | [*: Let binary literal can be convert to enum and set (#20789)](https://github.com/pingcap/tidb/pull/21443)                                   |   | 34d14h  |
| tidb/21504   | [planner: fix invalid convert type in between...and... (#19820)](https://github.com/pingcap/tidb/pull/21504)                                  | Y | 32d15h  |
| tidb/21546   | [planner: do not push down the aggregation function with correlated column (#21453)](https://github.com/pingcap/tidb/pull/21546)              |   | 28d23h  |
| tidb/21573   | [expression: fix incorrect result of IsTrue function for time types (#21534)](https://github.com/pingcap/tidb/pull/21573)                     |   | 28d13h  |
| tidb/21810   | [expression: handle hybrid field types for where clause (#21724)](https://github.com/pingcap/tidb/pull/21810)                                 |   | 21d18h  |
| tidb/21813   | [expression: handle tp.flen overflow in to_base64 function (#20947)](https://github.com/pingcap/tidb/pull/21813)                              |   | 21d17h  |
| tidb/21834   | [planner: enhanced index range calculation plan](https://github.com/pingcap/tidb/pull/21834)                                                  |   | 20d19h  |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                   |   | 18d20h  |
| tidb/21877   | [planner: fix correlated aggregates which should be evaluated in outer query (#21431)](https://github.com/pingcap/tidb/pull/21877)            |   | 18d19h  |
| tidb/21878   | [planner: do not push down lock to pointGet/bacthPointGet when selection exists](https://github.com/pingcap/tidb/pull/21878)                  |   | 18d18h  |
| tidb/21890   | [*: redact some error code, part(3/3) (#21866)](https://github.com/pingcap/tidb/pull/21890)                                                   |   | 16d15h  |
| tidb/21936   | [expression: fix wrong type inferring for ceiling function. (#21920)](https://github.com/pingcap/tidb/pull/21936)                             |   | 14d17h  |
| tidb/21956   | [planner/preprocessor: disallow into-outfile clause in some place](https://github.com/pingcap/tidb/pull/21956)                                |   | 13d23h  |
| tidb/21960   | [types: Regard `TypeNewDecimal` as not a `hasVariantFieldLength` type. (#21849)](https://github.com/pingcap/tidb/pull/21960)                  |   | 13d21h  |
| tidb/22026   | [expression: separated arithmeticPlusIntSig](https://github.com/pingcap/tidb/pull/22026)                                                      |   | 11d20h  |
| tidb/22043   | [planner, executor: enhance the limit pushdown rule.](https://github.com/pingcap/tidb/pull/22043)                                             |   | 9d11h   |
| tidb/22089   | [executor: fix signed cluster index behavior (#22085)](https://github.com/pingcap/tidb/pull/22089)                                            |   | 6d23h   |
| tidb/22104   | [executor: fix incompatible escape behaviors in `select into outfile` (#22100)](https://github.com/pingcap/tidb/pull/22104)                   |   | 6d17h   |
| tidb/22106   | [executor: avoid log duplicate index name in slow-log (#22057)](https://github.com/pingcap/tidb/pull/22106)                                   |   | 6d14h   |
| tidb/22107   | [executor: avoid log duplicate index name in slow-log (#22057)](https://github.com/pingcap/tidb/pull/22107)                                   |   | 6d14h   |
| tidb/22114   | [test: fix globalkilltest (#21987)](https://github.com/pingcap/tidb/pull/22114)                                                               |   | 6d12h   |
| tidb/22120   | [executor: fix `update ignore` into not exists partition (#21984)](https://github.com/pingcap/tidb/pull/22120)                                |   | 5d23h   |
| tidb/22136   | [executor: improve the runtime stats of index lookup reader (#21982)](https://github.com/pingcap/tidb/pull/22136)                             |   | 5d17h   |
| tidb/22152   | [planner: check index valid while forUpdateRead](https://github.com/pingcap/tidb/pull/22152)                                                  |   | 1d18h   |
| tidb/22181   | [planner, expression: fix error when using IN combined with subquery (#22080)](https://github.com/pingcap/tidb/pull/22181)                    |   | 18h     |


## Reviewed in Last 7 Days

|   REPO    |                                     PR                                     | C | D |   R   |
|-----------|----------------------------------------------------------------------------|---|---|-------|
| docs/4565 | [tidb: add doc for global kill](https://github.com/pingcap/docs/pull/4565) |   | 2 | 2d14h |


</details> 


<details> 
  <summary>XuHuaiyu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                              PR                                                                              | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19292 | [planner: suppport left join in join reorder](https://github.com/pingcap/tidb/pull/19292)                                                                    |   | 139d17h |
| docs/4565  | [tidb: add doc for global kill](https://github.com/pingcap/docs/pull/4565)                                                                                   |   | 4d9h    |
| tidb/19900 | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                                               | Y | 118d18h |
| tidb/20040 | [planner, expression: take NullFlag into consideration when optimize the `int non-const` <cmp > `non-int const`](https://github.com/pingcap/tidb/pull/20040) | Y | 111d14h |
| tidb/20140 | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                                           |   | 105d22h |
| tidb/20311 | [expression: fix overflow error when convert bit to int64 (#20266)](https://github.com/pingcap/tidb/pull/20311)                                              |   | 97d21h  |
| tidb/20350 | [executor: support read global indexes in IndexMergeReader and index join](https://github.com/pingcap/tidb/pull/20350)                                       | Y | 91d14h  |
| tidb/20505 | [*: Add metrics for oom-action and sql memory usage.](https://github.com/pingcap/tidb/pull/20505)                                                            |   | 78d19h  |
| tidb/20576 | [*: fix stats feedback after tableReader handle multiple ranges](https://github.com/pingcap/tidb/pull/20576)                                                 |   | 76d13h  |
| tidb/20613 | [executor: fix issue of hash join fetch time inaccurate](https://github.com/pingcap/tidb/pull/20613)                                                         |   | 74d13h  |
| tidb/20752 | [*: trace statsCache and preparePlanCache by Global memory tracker.](https://github.com/pingcap/tidb/pull/20752)                                             |   | 64d23h  |
| tidb/20790 | [collation: add pinyin collation for chinese charset support](https://github.com/pingcap/tidb/pull/20790)                                                    |   | 63d21h  |
| tidb/20793 | [planner, executor: enable inline projection for Apply](https://github.com/pingcap/tidb/pull/20793)                                                          |   | 63d21h  |
| tidb/20905 | [planner: fix statement-optimize not work in `TryFastPlan`](https://github.com/pingcap/tidb/pull/20905)                                                      |   | 60d17h  |
| tidb/20972 | [expression: POC implementation of Vitess hashing algorithm.](https://github.com/pingcap/tidb/pull/20972)                                                    |   | 56d1h   |
| tidb/21064 | [planner, executor: fix cast not check error](https://github.com/pingcap/tidb/pull/21064)                                                                    |   | 51d9h   |
| tidb/21149 | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                                |   | 47d15h  |
| tidb/21155 | [util/chunk: fix slice out of bound panic](https://github.com/pingcap/tidb/pull/21155)                                                                       |   | 47d12h  |
| tidb/21228 | [executor: return the result immediately when combining LIMIT row_count with DISTINCT](https://github.com/pingcap/tidb/pull/21228)                           |   | 43d14h  |
| tidb/21304 | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                                     |   | 41d12h  |
| tidb/21334 | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                                                |   | 40d14h  |
| tidb/21340 | [executor: initialize expensive query handler on domain creation](https://github.com/pingcap/tidb/pull/21340)                                                |   | 40d0h   |
| tidb/21425 | [planner: natural join not consider rowid and null eq not propagate (#21328)](https://github.com/pingcap/tidb/pull/21425)                                    |   | 34d22h  |
| tidb/21459 | [planner: push down projection for tiflash](https://github.com/pingcap/tidb/pull/21459)                                                                      |   | 33d22h  |
| tidb/21473 | [ddl: check the generated column offset when modifies column (#21458)](https://github.com/pingcap/tidb/pull/21473)                                           |   | 33d17h  |
| tidb/21476 | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                        |   | 33d16h  |
| tidb/21477 | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21477)                                                        |   | 33d16h  |
| tidb/21483 | [executor, store/tikv: locks exist keys for point_get & batch_point_get (#21229)](https://github.com/pingcap/tidb/pull/21483)                                |   | 33d13h  |
| tidb/21488 | [planner: fix ambiguous field when resolve having expr  (#21165)](https://github.com/pingcap/tidb/pull/21488)                                                |   | 32d23h  |
| tidb/21504 | [planner: fix invalid convert type in between...and... (#19820)](https://github.com/pingcap/tidb/pull/21504)                                                 | Y | 32d15h  |
| tidb/21532 | [expression: set IsBooleanFlag for boolean scalar functions (#20706)](https://github.com/pingcap/tidb/pull/21532)                                            |   | 29d17h  |
| tidb/21536 | [executor: add slow-log file meta cache to avoid repeat read file meta information](https://github.com/pingcap/tidb/pull/21536)                              |   | 29d15h  |
| tidb/21550 | [planner : fix unsigned_decimal_col=-int_cnst access index (#21198)](https://github.com/pingcap/tidb/pull/21550)                                             |   | 28d20h  |
| tidb/21564 | [ddl: fix Incorrect behavior of NO_ZERO_DATE when altering table](https://github.com/pingcap/tidb/pull/21564)                                                |   | 28d16h  |
| tidb/21573 | [expression: fix incorrect result of IsTrue function for time types (#21534)](https://github.com/pingcap/tidb/pull/21573)                                    |   | 28d13h  |
| tidb/21590 | [expression: fix compatibility behaviors in sec_to_time with MySQL  (#21555)](https://github.com/pingcap/tidb/pull/21590)                                    |   | 27d21h  |
| tidb/21593 | [expression: fix convert number base for hybrid type (#21554)](https://github.com/pingcap/tidb/pull/21593)                                                   |   | 27d20h  |
| tidb/21602 | [expression: not evaluate time addition for timestamp with 2 args if 1st arg's year is zero (#21572)](https://github.com/pingcap/tidb/pull/21602)            |   | 27d17h  |
| tidb/21608 | [expression: fix error "invalid time format: '{0 0 0 0 0 0 0}'" for timestampAdd (#21591)](https://github.com/pingcap/tidb/pull/21608)                       |   | 27d16h  |
| tidb/21610 | [*: remove needless InInsertStmt (#19787)](https://github.com/pingcap/tidb/pull/21610)                                                                       |   | 27d15h  |
| tidb/21614 | [planner: do not propagate column eq with different column types (#21495)](https://github.com/pingcap/tidb/pull/21614)                                       |   | 27d14h  |
| tidb/21626 | [test: convert test to benchmard test to make ci stable (#21616)](https://github.com/pingcap/tidb/pull/21626)                                                |   | 26d23h  |
| tidb/21635 | [expression: handle invalid argument for addtime and subtime function  (#21600)](https://github.com/pingcap/tidb/pull/21635)                                 |   | 26d20h  |
| tidb/21673 | [expression, types: fix unexpected result from TIME() when fsp digits > 6 (#21652)](https://github.com/pingcap/tidb/pull/21673)                              |   | 25d17h  |
| tidb/21676 | [expression: fix compatibility of extract day_time unit functions (#21601)](https://github.com/pingcap/tidb/pull/21676)                                      |   | 25d17h  |
| tidb/21680 | [planner: report error when ORDER BY conflicts with DISTINCT (#21286)](https://github.com/pingcap/tidb/pull/21680)                                           |   | 25d16h  |
| tidb/21697 | [planner: check for only_full_group_by in ORDER BY and HAVING (#21216)](https://github.com/pingcap/tidb/pull/21697)                                          |   | 22d20h  |
| tidb/21711 | [expression: Fix unexpected panic when using IF function. (#21132)](https://github.com/pingcap/tidb/pull/21711)                                              |   | 22d17h  |
| tidb/21714 | [planner: fix the coercibility of the cast function (#21705)](https://github.com/pingcap/tidb/pull/21714)                                                    |   | 22d17h  |
| tidb/21718 | [types: fix compare object json type (#21703)](https://github.com/pingcap/tidb/pull/21718)                                                                   |   | 22d16h  |
| tidb/21785 | [types: fix compare float64 with float64 in json (#21709)](https://github.com/pingcap/tidb/pull/21785)                                                       |   | 21d22h  |
| tidb/21808 | [planner: fix the fail when we compare multi fields in the subquery (#21699)](https://github.com/pingcap/tidb/pull/21808)                                    |   | 21d18h  |
| tidb/21810 | [expression: handle hybrid field types for where clause (#21724)](https://github.com/pingcap/tidb/pull/21810)                                                |   | 21d18h  |
| tidb/21813 | [expression: handle tp.flen overflow in to_base64 function (#20947)](https://github.com/pingcap/tidb/pull/21813)                                             |   | 21d17h  |
| tidb/21839 | [planner/core: add 'split table using statistics' statement](https://github.com/pingcap/tidb/pull/21839)                                                     |   | 20d15h  |
| tidb/21853 | [expression: fix compatibility behaviors in time_format with MySQL (#21559)](https://github.com/pingcap/tidb/pull/21853)                                     |   | 19d19h  |
| tidb/21870 | [types: report error for json object with key length >= 65536 (#21779)](https://github.com/pingcap/tidb/pull/21870)                                          |   | 18d23h  |
| tidb/21874 | [expression:truncate decimal value instead of return error (#21691)](https://github.com/pingcap/tidb/pull/21874)                                             |   | 18d20h  |
| tidb/21877 | [planner: fix correlated aggregates which should be evaluated in outer query (#21431)](https://github.com/pingcap/tidb/pull/21877)                           |   | 18d19h  |
| tidb/21896 | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                                                   |   | 15d19h  |
| tidb/21916 | [server: double type column from table should ignore its decimal (#21788)](https://github.com/pingcap/tidb/pull/21916)                                       |   | 14d23h  |
| tidb/21924 | [expression: fix type infer for tidb's builtin compare(least and greatest) (#21150)](https://github.com/pingcap/tidb/pull/21924)                             |   | 14d19h  |
| tidb/21936 | [expression: fix wrong type inferring for ceiling function. (#21920)](https://github.com/pingcap/tidb/pull/21936)                                            |   | 14d17h  |
| tidb/21957 | [planner: fix unknown columns in join using below agg (#21922)](https://github.com/pingcap/tidb/pull/21957)                                                  |   | 13d23h  |
| tidb/21958 | [expression: fix comparing json with string (#21903)](https://github.com/pingcap/tidb/pull/21958)                                                            |   | 13d22h  |
| tidb/21964 | [planner: add plancodec id for all type TableScan/IndexScan. (#21935)](https://github.com/pingcap/tidb/pull/21964)                                           |   | 13d18h  |
| tidb/21972 | [executor: throw error when prepared statement is execute, deallocate or prepare (#21962)](https://github.com/pingcap/tidb/pull/21972)                       |   | 13d16h  |
| tidb/22013 | [executor: fix unstable test Issue16696 (#22009)](https://github.com/pingcap/tidb/pull/22013)                                                                |   | 12d17h  |
| tidb/22014 | [executor: fix unstable test Issue16696 (#22009)](https://github.com/pingcap/tidb/pull/22014)                                                                |   | 12d17h  |
| tidb/22107 | [executor: avoid log duplicate index name in slow-log (#22057)](https://github.com/pingcap/tidb/pull/22107)                                                  |   | 6d14h   |
| tidb/22118 | [planner: check if columns count matches for batch point get in TryFastPlan (#22044)](https://github.com/pingcap/tidb/pull/22118)                            |   | 5d23h   |
| tidb/22119 | [executor: fix `update ignore` into not exists partition (#21984)](https://github.com/pingcap/tidb/pull/22119)                                               |   | 5d23h   |
| tidb/22120 | [executor: fix `update ignore` into not exists partition (#21984)](https://github.com/pingcap/tidb/pull/22120)                                               |   | 5d23h   |
| tidb/22131 | [privilege: remove leading and trailing space when create user and role](https://github.com/pingcap/tidb/pull/22131)                                         |   | 5d19h   |
| tidb/22136 | [executor: improve the runtime stats of index lookup reader (#21982)](https://github.com/pingcap/tidb/pull/22136)                                            |   | 5d17h   |
| tidb/22141 | [store: trace `loadRegion` to see the PD region cache loading (#22092)](https://github.com/pingcap/tidb/pull/22141)                                          |   | 2d0h    |
| tidb/22142 | [store: trace `loadRegion` to see the PD region cache loading (#22092)](https://github.com/pingcap/tidb/pull/22142)                                          |   | 2d0h    |
| tidb/22143 | [expression: return correct results for user variables of datetime type (#22078)](https://github.com/pingcap/tidb/pull/22143)                                |   | 1d23h   |
| tidb/22148 | [session: set process info before building plan (#22101)](https://github.com/pingcap/tidb/pull/22148)                                                        |   | 1d19h   |
| tidb/22149 | [session: set process info before building plan (#22101)](https://github.com/pingcap/tidb/pull/22149)                                                        |   | 1d19h   |
| tidb/22153 | [executor: refine bigint unsigned primary key duplicate error](https://github.com/pingcap/tidb/pull/22153)                                                   |   | 1d18h   |
| tidb/22173 | [expression: fix unexpected panic when doing isNullRejected check](https://github.com/pingcap/tidb/pull/22173)                                               |   | 21h     |
| tidb/22186 | [executor: fix select into outfile with year type column has no data (#22175)](https://github.com/pingcap/tidb/pull/22186)                                   |   | 16h     |


## Reviewed in Last 7 Days

|    REPO    |                                                               PR                                                               | C | D |   R   |
|------------|--------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/22101 | [session: set process info before building plan](https://github.com/pingcap/tidb/pull/22101)                                   |   | 2 | 4d18h |
| tidb/22111 | [config, session: promise the compatibility of oom-action when upgrading (#22102)](https://github.com/pingcap/tidb/pull/22111) |   | 7 | 1h    |
| tidb/22115 | [util: add cache for mem info in container (#22109)](https://github.com/pingcap/tidb/pull/22115)                               |   | 7 | 0h    |
| tidb/22112 | [bootstrap: add mysql.user columns for parser#1121 (#21856)](https://github.com/pingcap/tidb/pull/22112)                       |   | 7 | 0h    |
| tidb/22104 | [executor: fix incompatible escape behaviors in `select into outfile` (#22100)](https://github.com/pingcap/tidb/pull/22104)    |   | 7 | 2h    |
| tidb/22103 | [executor: fix incompatible escape behaviors in `select into outfile` (#22100)](https://github.com/pingcap/tidb/pull/22103)    |   | 7 | 2h    |
| tidb/22100 | [executor: fix incompatible escape behaviors in `select into outfile`](https://github.com/pingcap/tidb/pull/22100)             |   | 7 | 0h    |
| tidb/22095 | [select into outfile `untime error: index out of range` (#22038)](https://github.com/pingcap/tidb/pull/22095)                  |   | 7 | 1h    |
| tidb/22094 | [select into outfile `untime error: index out of range` (#22038)](https://github.com/pingcap/tidb/pull/22094)                  |   | 7 | 1h    |
| tidb/22038 | [select into outfile `untime error: index out of range`](https://github.com/pingcap/tidb/pull/22038)                           |   | 7 | 4d17h |


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
| tidb/14729 | [planner: fix constant propagation for PredicatePushDown](https://github.com/pingcap/tidb/pull/14729)                                | Y | 329d18h |
| tidb/14831 | [planner/cascades: add implementationRule for IndexLookUpJoin](https://github.com/pingcap/tidb/pull/14831)                           |   | 322d18h |
| tidb/15090 | [planner/cascades: refine the row count estimation of TiKV layer Selection](https://github.com/pingcap/tidb/pull/15090)              |   | 308d18h |
| tidb/15157 | [planner/cascades: implement `HashCode` method for all the LogicalPlans](https://github.com/pingcap/tidb/pull/15157)                 | Y | 306d15h |
| tidb/15335 | [planner/cascades: add transformation rule PullAggregationUpApply & EliminateMaxOneRow](https://github.com/pingcap/tidb/pull/15335)  |   | 299d18h |
| tidb/15370 | [planner,executor: Refactor Shuffle and implement parallel Sort](https://github.com/pingcap/tidb/pull/15370)                         | Y | 297d19h |
| tidb/17276 | [planner/cascades: add rule InjectProjectionBelowSort](https://github.com/pingcap/tidb/pull/17276)                                   | Y | 232d9h  |
| tidb/18882 | [planner, executor: add explain for `MetricSummaryTableExtractor`](https://github.com/pingcap/tidb/pull/18882)                       | Y | 159d18h |
| tidb/19347 | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347) |   | 137d23h |
| tidb/20580 | [statistics: add bucket ndv for index histogram](https://github.com/pingcap/tidb/pull/20580)                                         |   | 75d21h  |
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                            |   | 61d17h  |
| tidb/21444 | [planner: ignore anonymous index while tiflash replica is available](https://github.com/pingcap/tidb/pull/21444)                     |   | 34d12h  |
| tidb/21459 | [planner: push down projection for tiflash](https://github.com/pingcap/tidb/pull/21459)                                              |   | 33d22h  |
| tidb/21488 | [planner: fix ambiguous field when resolve having expr  (#21165)](https://github.com/pingcap/tidb/pull/21488)                        |   | 32d23h  |
| tidb/21573 | [expression: fix incorrect result of IsTrue function for time types (#21534)](https://github.com/pingcap/tidb/pull/21573)            |   | 28d13h  |
| tidb/21680 | [planner: report error when ORDER BY conflicts with DISTINCT (#21286)](https://github.com/pingcap/tidb/pull/21680)                   |   | 25d16h  |
| tidb/21697 | [planner: check for only_full_group_by in ORDER BY and HAVING (#21216)](https://github.com/pingcap/tidb/pull/21697)                  |   | 22d20h  |
| tidb/21994 | [range: fix overflow value access index ](https://github.com/pingcap/tidb/pull/21994)                                                |   | 12d23h  |


## Reviewed in Last 7 Days

|    REPO     |                                                              PR                                                               | C | D |   R    |
|-------------|-------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/22169  | [statistics: fix stack overflow when use DNF expr immediately after add column](https://github.com/pingcap/tidb/pull/22169)   |   | 1 | 4h     |
| tidb/21275  | [*: rewrite origin SQL with default DB for SQL bindings](https://github.com/pingcap/tidb/pull/21275)                          |   | 1 | 41d3h  |
| tidb/14412  | [executor, planner: ON DUPLICATE UPDATE can refer to un-project col](https://github.com/pingcap/tidb/pull/14412)              |   | 1 | 362d0h |
| tidb/22080  | [planner, expression: fix error when using IN combined with subquery](https://github.com/pingcap/tidb/pull/22080)             |   | 1 | 6d17h  |
| tidb/22086  | [planner/core: fix a bug of adding enforcer.](https://github.com/pingcap/tidb/pull/22086)                                     |   | 2 | 5d17h  |
| parser/1146 | [yy_parser: enable supportWindowFunc by default (#986)](https://github.com/pingcap/parser/pull/1146)                          |   | 2 | 0h     |
| tidb/22124  | [planner: avoid using index_merge when there are multiple table filters (#22122)](https://github.com/pingcap/tidb/pull/22124) |   | 6 | 0h     |
| tidb/22125  | [planner: avoid using index_merge when there are multiple table filters (#22122)](https://github.com/pingcap/tidb/pull/22125) |   | 6 | 0h     |
| tidb/22122  | [planner: avoid using index_merge when there are multiple table filters](https://github.com/pingcap/tidb/pull/22122)          |   | 6 | 1h     |


</details> 


<details> 
  <summary>fzhedu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                   PR                                                   | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19845 | [expression:fix FORMAT compatibility issue #11206](https://github.com/pingcap/tidb/pull/19845)         | Y | 120d16h |
| tidb/20117 | [optimizer: fix issue on incorrect result of natural join](https://github.com/pingcap/tidb/pull/20117) | Y | 106d21h |


## Reviewed in Last 7 Days

|    REPO    |                                                 PR                                                 | C | D |   R   |
|------------|----------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/22053 | [execution: support explain analyze in mpp execution.](https://github.com/pingcap/tidb/pull/22053) |   | 1 | 7d23h |


</details> 


<details> 
  <summary>ichn-hu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                            PR                                                            | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/18312 | [expression: fix compatible problem with mysql when parse datetime](https://github.com/pingcap/tidb/pull/18312)          |   | 188d17h |
| tidb/21676 | [expression: fix compatibility of extract day_time unit functions (#21601)](https://github.com/pingcap/tidb/pull/21676)  |   | 25d17h  |
| tidb/21850 | [expression: add implicit eval int and real for function dayname (#21806)](https://github.com/pingcap/tidb/pull/21850)   |   | 19d19h  |
| tidb/21853 | [expression: fix compatibility behaviors in time_format with MySQL (#21559)](https://github.com/pingcap/tidb/pull/21853) |   | 19d19h  |


## Reviewed in Last 7 Days

|    REPO    |                                                        PR                                                         | C | D |   R   |
|------------|-------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/21310 | [types: convert string to MySQL BIT correctly](https://github.com/pingcap/tidb/pull/21310)                        |   | 1 | 40d2h |
| tidb/22175 | [executor: fix select into outfile with year type column has no data](https://github.com/pingcap/tidb/pull/22175) |   | 1 | 0h    |


</details> 


<details> 
  <summary>lzmhhh123's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                    PR                                                                     | C | LASTED  |
|--------------|-------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/14729   | [planner: fix constant propagation for PredicatePushDown](https://github.com/pingcap/tidb/pull/14729)                                     | Y | 329d18h |
| docs-cn/4913 | [explain: add indexes](https://github.com/pingcap/docs-cn/pull/4913)                                                                      |   | 50d18h  |
| tidb/17414   | [add curCost based join reorder algorithm](https://github.com/pingcap/tidb/pull/17414)                                                    |   | 224d18h |
| tidb/19347   | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347)      |   | 137d23h |
| tidb/19698   | [*: update test cases to support new collation enabled by default](https://github.com/pingcap/tidb/pull/19698)                            |   | 125d23h |
| tidb/20044   | [expression: Add column nullability checking before "refine args"](https://github.com/pingcap/tidb/pull/20044)                            | Y | 111d7h  |
| tidb/20444   | [expression: add json_merge_patch](https://github.com/pingcap/tidb/pull/20444)                                                            |   | 83d21h  |
| tidb/20465   | [expression: add uuidShortFunction](https://github.com/pingcap/tidb/pull/20465)                                                           |   | 82d20h  |
| tidb/20505   | [*: Add metrics for oom-action and sql memory usage.](https://github.com/pingcap/tidb/pull/20505)                                         |   | 78d19h  |
| tidb/20618   | [planner: fix update generated columns error](https://github.com/pingcap/tidb/pull/20618)                                                 |   | 73d20h  |
| tidb/20642   | [executor: modify admin executors to support partitioned table with global index](https://github.com/pingcap/tidb/pull/20642)             |   | 71d16h  |
| tidb/20825   | [executor: add diagnosis rule to check Transparent Huge Pages(THP) enabled (#20611)](https://github.com/pingcap/tidb/pull/20825)          |   | 62d19h  |
| tidb/20903   | [planner: fix confused and unnecessary double-projection in plans.](https://github.com/pingcap/tidb/pull/20903)                           |   | 60d17h  |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                            |   | 54d17h  |
| tidb/21051   | [executor: change read slow-log file module to concurrent](https://github.com/pingcap/tidb/pull/21051)                                    |   | 53d14h  |
| tidb/21137   | [executor: specially handle empty input for apply's outer child aggregate (#20544)](https://github.com/pingcap/tidb/pull/21137)           |   | 47d20h  |
| tidb/21195   | [brie: integrate lightning to suport IMPORT statement](https://github.com/pingcap/tidb/pull/21195)                                        |   | 43d23h  |
| tidb/21275   | [*: rewrite origin SQL with default DB for SQL bindings](https://github.com/pingcap/tidb/pull/21275)                                      |   | 41d22h  |
| tidb/21334   | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                             |   | 40d14h  |
| tidb/21347   | [session: make rollback work on global variables](https://github.com/pingcap/tidb/pull/21347)                                             |   | 39d20h  |
| tidb/21401   | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                        |   | 36d12h  |
| tidb/21404   | [planner: fix unexpected bad plan when IndexJoin inner side estRow is 0. (#21084)](https://github.com/pingcap/tidb/pull/21404)            |   | 35d23h  |
| tidb/21444   | [planner: ignore anonymous index while tiflash replica is available](https://github.com/pingcap/tidb/pull/21444)                          |   | 34d12h  |
| tidb/21487   | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                             |   | 33d5h   |
| tidb/21641   | [executor: Fix pessimistic lock doesn't work on the partition table for subquery/joins](https://github.com/pingcap/tidb/pull/21641)       |   | 26d18h  |
| tidb/21651   | [planner: allow filter condition pushing down to IndexScan for prefix index](https://github.com/pingcap/tidb/pull/21651)                  |   | 26d14h  |
| tidb/21680   | [planner: report error when ORDER BY conflicts with DISTINCT (#21286)](https://github.com/pingcap/tidb/pull/21680)                        |   | 25d16h  |
| tidb/21711   | [expression: Fix unexpected panic when using IF function. (#21132)](https://github.com/pingcap/tidb/pull/21711)                           |   | 22d17h  |
| tidb/21808   | [planner: fix the fail when we compare multi fields in the subquery (#21699)](https://github.com/pingcap/tidb/pull/21808)                 |   | 21d18h  |
| tidb/21842   | [planner: Shuffle hash agg](https://github.com/pingcap/tidb/pull/21842)                                                                   |   | 20d10h  |
| tidb/21850   | [expression: add implicit eval int and real for function dayname (#21806)](https://github.com/pingcap/tidb/pull/21850)                    |   | 19d19h  |
| tidb/21853   | [expression: fix compatibility behaviors in time_format with MySQL (#21559)](https://github.com/pingcap/tidb/pull/21853)                  |   | 19d19h  |
| tidb/21870   | [types: report error for json object with key length >= 65536 (#21779)](https://github.com/pingcap/tidb/pull/21870)                       |   | 18d23h  |
| tidb/21877   | [planner: fix correlated aggregates which should be evaluated in outer query (#21431)](https://github.com/pingcap/tidb/pull/21877)        |   | 18d19h  |
| tidb/21924   | [expression: fix type infer for tidb's builtin compare(least and greatest) (#21150)](https://github.com/pingcap/tidb/pull/21924)          |   | 14d19h  |
| tidb/21954   | [planner/cascades: add rule `PushSelDownApply`](https://github.com/pingcap/tidb/pull/21954)                                               |   | 13d23h  |
| tidb/21972   | [executor: throw error when prepared statement is execute, deallocate or prepare (#21962)](https://github.com/pingcap/tidb/pull/21972)    |   | 13d16h  |
| tidb/22022   | [planner/codec: fix issue of decode plan error cause by without escape special char (#21937)](https://github.com/pingcap/tidb/pull/22022) |   | 12d0h   |
| tidb/22053   | [execution: support explain analyze in mpp execution.](https://github.com/pingcap/tidb/pull/22053)                                        |   | 8d18h   |
| tidb/22089   | [executor: fix signed cluster index behavior (#22085)](https://github.com/pingcap/tidb/pull/22089)                                        |   | 6d23h   |
| tidb/22124   | [planner: avoid using index_merge when there are multiple table filters (#22122)](https://github.com/pingcap/tidb/pull/22124)             |   | 5d21h   |
| tidb/22126   | [*: add `sys` schema, `sys.SCHEMA_UNUSED_INDEXES` view and `sys.SCHEMA_INDEX_USAGE` view](https://github.com/pingcap/tidb/pull/22126)     |   | 5d20h   |
| tidb/22130   | [planner: join reorder should not change the order of output columns (#16852)](https://github.com/pingcap/tidb/pull/22130)                |   | 5d19h   |
| tidb/22137   | [expression: separated arithmeticModIntSig](https://github.com/pingcap/tidb/pull/22137)                                                   |   | 5d12h   |
| tidb/22148   | [session: set process info before building plan (#22101)](https://github.com/pingcap/tidb/pull/22148)                                     |   | 1d19h   |
| tidb/22149   | [session: set process info before building plan (#22101)](https://github.com/pingcap/tidb/pull/22149)                                     |   | 1d19h   |
| tidb/22174   | [expression, ddl: check the argument count for the generated column (#22154)](https://github.com/pingcap/tidb/pull/22174)                 |   | 20h     |
| tidb/22188   | [planner: do not use indexMerge when the path only use a single index (#22168)](https://github.com/pingcap/tidb/pull/22188)               |   | 13h     |


## Reviewed in Last 7 Days

|    REPO    |                                                               PR                                                               | C | D |   R   |
|------------|--------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/21310 | [types: convert string to MySQL BIT correctly](https://github.com/pingcap/tidb/pull/21310)                                     |   | 1 | 40d2h |
| tidb/22154 | [expression, ddl: check the argument count for the generated column](https://github.com/pingcap/tidb/pull/22154)               |   | 1 | 21h   |
| tidb/22073 | [executor: always decode the value first and then the handle](https://github.com/pingcap/tidb/pull/22073)                      |   | 2 | 5d23h |
| tidb/22147 | [types: Add a limitation about float data type (#20929)](https://github.com/pingcap/tidb/pull/22147)                           |   | 2 | 1h    |
| tidb/22101 | [session: set process info before building plan](https://github.com/pingcap/tidb/pull/22101)                                   |   | 6 | 1d0h  |
| tidb/21976 | [planner: report error for invalid window specs which are not used (#21083)](https://github.com/pingcap/tidb/pull/21976)       |   | 6 | 7d22h |
| tidb/21969 | [types:  Add a limitation about float data type (#20929)](https://github.com/pingcap/tidb/pull/21969)                          |   | 6 | 8d0h  |
| tidb/22111 | [config, session: promise the compatibility of oom-action when upgrading (#22102)](https://github.com/pingcap/tidb/pull/22111) |   | 7 | 1h    |
| tidb/22085 | [executor: fix signed cluster index behavior](https://github.com/pingcap/tidb/pull/22085)                                      |   | 7 | 13h   |


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
| tidb/16305   | [expression: separate signatures for `ModInt`](https://github.com/pingcap/tidb/pull/16305)                                                           | Y | 268d0h  |
| docs-cn/4669 | [sql-optimization: extended statistics documentation](https://github.com/pingcap/docs-cn/pull/4669)                                                  |   | 84d17h  |
| tidb/16967   | [executor: Refactor Shuffle and implement parallel sort (executor part)](https://github.com/pingcap/tidb/pull/16967)                                 | Y | 248d10h |
| tidb/17396   | [types: improve StrToDate performance](https://github.com/pingcap/tidb/pull/17396)                                                                   | Y | 225d10h |
| tidb/18882   | [planner, executor: add explain for `MetricSummaryTableExtractor`](https://github.com/pingcap/tidb/pull/18882)                                       | Y | 159d18h |
| tidb/19029   | [types: fix unexpected NOT_NULL flags](https://github.com/pingcap/tidb/pull/19029)                                                                   |   | 152d22h |
| tidb/19120   | [executor: Concurrently fetch chunks and insert them to a concurrent hash table in hash build](https://github.com/pingcap/tidb/pull/19120)           |   | 147d21h |
| tidb/19292   | [planner: suppport left join in join reorder](https://github.com/pingcap/tidb/pull/19292)                                                            |   | 139d17h |
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                               | Y | 116d14h |
| tidb/20011   | [statistics: fix incorrect total count used in index selectivity computation](https://github.com/pingcap/tidb/pull/20011)                            |   | 112d15h |
| tidb/20316   | [docs/design: add design doc for index usage information](https://github.com/pingcap/tidb/pull/20316)                                                |   | 97d17h  |
| tidb/20354   | [planner: rename relational operators (#14575)](https://github.com/pingcap/tidb/pull/20354)                                                          | Y | 90d6h   |
| tidb/20399   | [*: make 'tidb_enable_change_column_type' available as a session variable](https://github.com/pingcap/tidb/pull/20399)                               |   | 86d16h  |
| tidb/20689   | [expression: make TIME function compatible with MySQL (#19158)](https://github.com/pingcap/tidb/pull/20689)                                          |   | 69d21h  |
| tidb/20708   | [*: separate auto_increment ID allocator from _tidb_rowid allocator](https://github.com/pingcap/tidb/pull/20708)                                     |   | 68d20h  |
| tidb/20972   | [expression: POC implementation of Vitess hashing algorithm.](https://github.com/pingcap/tidb/pull/20972)                                            |   | 56d1h   |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                       |   | 54d17h  |
| tidb/21137   | [executor: specially handle empty input for apply's outer child aggregate (#20544)](https://github.com/pingcap/tidb/pull/21137)                      |   | 47d20h  |
| tidb/21149   | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                        |   | 47d15h  |
| tidb/21304   | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                             |   | 41d12h  |
| tidb/21318   | [planner, expression: use the range of column types to simplify expressions](https://github.com/pingcap/tidb/pull/21318)                             |   | 40d19h  |
| tidb/21359   | [*: add runtime stats for split region statement](https://github.com/pingcap/tidb/pull/21359)                                                        |   | 39d13h  |
| tidb/21401   | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                                   |   | 36d12h  |
| tidb/21408   | [statistics: fix a bug which causes panic when using the clustered index and the new collation (#21379)](https://github.com/pingcap/tidb/pull/21408) |   | 35d20h  |
| tidb/21424   | [sessionctx: move set variable to sysvar struct](https://github.com/pingcap/tidb/pull/21424)                                                         |   | 35d5h   |
| tidb/21464   | [server: return results of ongoing queries when graceful shutdown (#19669)](https://github.com/pingcap/tidb/pull/21464)                              |   | 33d20h  |
| tidb/21471   | [session: fix ineffective EXPLAIN FOR CONNECTION statement (#21044)](https://github.com/pingcap/tidb/pull/21471)                                     |   | 33d17h  |
| tidb/21476   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                |   | 33d16h  |
| tidb/21477   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21477)                                                |   | 33d16h  |
| tidb/21508   | [execution: fix dayofweek('0000-00-00') behavior](https://github.com/pingcap/tidb/pull/21508)                                                        |   | 32d10h  |
| tidb/21525   | [expression: fix compatibility behaviors in zero datetime with MySQL (#21220)](https://github.com/pingcap/tidb/pull/21525)                           |   | 29d20h  |
| tidb/21610   | [*: remove needless InInsertStmt (#19787)](https://github.com/pingcap/tidb/pull/21610)                                                               |   | 27d15h  |
| tidb/21665   | [executor: fix LEAD and LAG's default value can not adapt to field type (#20747)](https://github.com/pingcap/tidb/pull/21665)                        |   | 25d19h  |
| tidb/21680   | [planner: report error when ORDER BY conflicts with DISTINCT (#21286)](https://github.com/pingcap/tidb/pull/21680)                                   |   | 25d16h  |
| tidb/21711   | [expression: Fix unexpected panic when using IF function. (#21132)](https://github.com/pingcap/tidb/pull/21711)                                      |   | 22d17h  |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                          |   | 18d20h  |
| tidb/21887   | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                                      |   | 17d11h  |
| tidb/21895   | [executor: fix load data in file get wrong result #20854](https://github.com/pingcap/tidb/pull/21895)                                                |   | 15d20h  |
| tidb/21924   | [expression: fix type infer for tidb's builtin compare(least and greatest) (#21150)](https://github.com/pingcap/tidb/pull/21924)                     |   | 14d19h  |
| tidb/21930   | [planner: propagate NDV of column groups across plan nodes (#17854)](https://github.com/pingcap/tidb/pull/21930)                                     |   | 14d18h  |
| tidb/21969   | [types:  Add a limitation about float data type (#20929)](https://github.com/pingcap/tidb/pull/21969)                                                |   | 13d17h  |
| tidb/21971   | [executor: fix `insert ignore` into not exists partition (#21904)](https://github.com/pingcap/tidb/pull/21971)                                       |   | 13d17h  |
| tidb/21977   | [expression: log functions that can not be pushed to cop](https://github.com/pingcap/tidb/pull/21977)                                                |   | 13d16h  |
| tidb/22021   | [distsql: fix cop stats string display when there is only 1 rpc (#21901) (#21999)](https://github.com/pingcap/tidb/pull/22021)                       |   | 12d0h   |
| tidb/22048   | [session: support validating read-only statement during read-only staleness transaction](https://github.com/pingcap/tidb/pull/22048)                 |   | 8d21h   |
| tidb/22090   | [planner: push aggregation operators down to projection and union by default](https://github.com/pingcap/tidb/pull/22090)                            |   | 6d22h   |
| tidb/22104   | [executor: fix incompatible escape behaviors in `select into outfile` (#22100)](https://github.com/pingcap/tidb/pull/22104)                          |   | 6d17h   |
| tidb/22106   | [executor: avoid log duplicate index name in slow-log (#22057)](https://github.com/pingcap/tidb/pull/22106)                                          |   | 6d14h   |
| tidb/22107   | [executor: avoid log duplicate index name in slow-log (#22057)](https://github.com/pingcap/tidb/pull/22107)                                          |   | 6d14h   |
| tidb/22110   | [config, session: promise the compatibility of oom-action when upgrading (#22102)](https://github.com/pingcap/tidb/pull/22110)                       |   | 6d13h   |
| tidb/22118   | [planner: check if columns count matches for batch point get in TryFastPlan (#22044)](https://github.com/pingcap/tidb/pull/22118)                    |   | 5d23h   |
| tidb/22127   | [*: support ALTER TABLE ADD / DROP TIDB_STATS](https://github.com/pingcap/tidb/pull/22127)                                                           |   | 5d20h   |
| tidb/22136   | [executor: improve the runtime stats of index lookup reader (#21982)](https://github.com/pingcap/tidb/pull/22136)                                    |   | 5d17h   |
| tidb/22143   | [expression: return correct results for user variables of datetime type (#22078)](https://github.com/pingcap/tidb/pull/22143)                        |   | 1d23h   |
| tidb/22146   | [executor: forbid SFU on view](https://github.com/pingcap/tidb/pull/22146)                                                                           |   | 1d22h   |
| tidb/22164   | [*: remove tidb-lightning pkg to fix unexpected config change](https://github.com/pingcap/tidb/pull/22164)                                           |   | 1d13h   |
| tidb/22173   | [expression: fix unexpected panic when doing isNullRejected check](https://github.com/pingcap/tidb/pull/22173)                                       |   | 21h     |
| tidb/22182   | [bindinfo: avoid duplicate bindings caused by concurrent baseline capture](https://github.com/pingcap/tidb/pull/22182)                               |   | 18h     |
| tidb/22183   | [executor: fix data too long in union statement with select null and select varchar](https://github.com/pingcap/tidb/pull/22183)                     |   | 18h     |


## Reviewed in Last 7 Days

|    REPO    |                                                                  PR                                                                   | C | D |   R    |
|------------|---------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/22168 | [planner: do not use indexMerge when the path only use a single index](https://github.com/pingcap/tidb/pull/22168)                    |   | 1 | 8h     |
| tidb/22126 | [*: add `sys` schema, `sys.SCHEMA_UNUSED_INDEXES` view and `sys.SCHEMA_INDEX_USAGE` view](https://github.com/pingcap/tidb/pull/22126) |   | 1 | 5d5h   |
| tidb/22169 | [statistics: fix stack overflow when use DNF expr immediately after add column](https://github.com/pingcap/tidb/pull/22169)           |   | 1 | 7h     |
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                             |   | 1 | 61d2h  |
| tidb/14412 | [executor, planner: ON DUPLICATE UPDATE can refer to un-project col](https://github.com/pingcap/tidb/pull/14412)                      |   | 1 | 362d1h |
| tidb/22080 | [planner, expression: fix error when using IN combined with subquery](https://github.com/pingcap/tidb/pull/22080)                     |   | 1 | 6d21h  |
| tidb/20580 | [statistics: add bucket ndv for index histogram](https://github.com/pingcap/tidb/pull/20580)                                          |   | 1 | 75d1h  |
| tidb/22156 | [mod: update parser to latest 4.0 version](https://github.com/pingcap/tidb/pull/22156)                                                |   | 2 | 4h     |
| tidb/21275 | [*: rewrite origin SQL with default DB for SQL bindings](https://github.com/pingcap/tidb/pull/21275)                                  |   | 2 | 40d4h  |
| tidb/22033 | [statistics: redesign the schema for `mysql.stats_extended`](https://github.com/pingcap/tidb/pull/22033)                              |   | 2 | 9d23h  |
| tidb/21712 | [statistics: no more counting feedback if it is invalid](https://github.com/pingcap/tidb/pull/21712)                                  |   | 2 | 20d23h |
| tidb/20750 | [executor, infoschema, planner: optimize query cluster_slow_query](https://github.com/pingcap/tidb/pull/20750)                        |   | 2 | 63d0h  |
| tidb/22078 | [expression: return correct results for user variables of datetime type](https://github.com/pingcap/tidb/pull/22078)                  |   | 2 | 5d17h  |
| tidb/22125 | [planner: avoid using index_merge when there are multiple table filters (#22122)](https://github.com/pingcap/tidb/pull/22125)         |   | 6 | 1h     |
| tidb/16852 | [planner: join reorder should not change the order of output columns](https://github.com/pingcap/tidb/pull/16852)                     |   | 6 | 248d3h |
| tidb/22124 | [planner: avoid using index_merge when there are multiple table filters (#22122)](https://github.com/pingcap/tidb/pull/22124)         |   | 6 | 0h     |
| tidb/22122 | [planner: avoid using index_merge when there are multiple table filters](https://github.com/pingcap/tidb/pull/22122)                  |   | 6 | 1h     |
| tidb/22102 | [config, session: promise the compatibility of oom-action when upgrading](https://github.com/pingcap/tidb/pull/22102)                 |   | 7 | 0h     |
| tidb/22057 | [executor: avoid log duplicate index name in slow-log](https://github.com/pingcap/tidb/pull/22057)                                    |   | 7 | 1d19h  |
| tidb/22095 | [select into outfile `untime error: index out of range` (#22038)](https://github.com/pingcap/tidb/pull/22095)                         |   | 7 | 1h     |
| tidb/22094 | [select into outfile `untime error: index out of range` (#22038)](https://github.com/pingcap/tidb/pull/22094)                         |   | 7 | 1h     |
| tidb/22038 | [select into outfile `untime error: index out of range`](https://github.com/pingcap/tidb/pull/22038)                                  |   | 7 | 4d14h  |


</details> 


<details> 
  <summary>rebelice's detailed review report</summary> 

## To Be Reviewed

| REPO | PR | C | LASTED |
|------|----|---|--------|


## Reviewed in Last 7 Days

|    REPO    |                                           PR                                            | C | D |   R   |
|------------|-----------------------------------------------------------------------------------------|---|---|-------|
| tidb/21459 | [planner: push down projection for tiflash](https://github.com/pingcap/tidb/pull/21459) |   | 2 | 32d5h |


</details> 


<details> 
  <summary>time-and-fate's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                  PR                                                   | C | LASTED |
|------------|-------------------------------------------------------------------------------------------------------|---|--------|
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)             |   | 61d17h |
| tidb/22006 | [config: disable statistics feedback by default (#21923)](https://github.com/pingcap/tidb/pull/22006) |   | 12d19h |


## Reviewed in Last 7 Days

|     REPO     |                                              PR                                              | C | D |   R   |
|--------------|----------------------------------------------------------------------------------------------|---|---|-------|
| tidb/20580   | [statistics: add bucket ndv for index histogram](https://github.com/pingcap/tidb/pull/20580) |   | 7 | 69d3h |
| docs-cn/5164 | [statistics: add docs for version 2](https://github.com/pingcap/docs-cn/pull/5164)           |   | 7 | 5d22h |


</details> 


<details> 
  <summary>winoros's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                          PR                                                                          | C | LASTED  |
|--------------|------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/14424   | [expression: add nullable() method to check whether an expression can return null](https://github.com/pingcap/tidb/pull/14424)                       |   | 362d18h |
| docs-cn/4669 | [sql-optimization: extended statistics documentation](https://github.com/pingcap/docs-cn/pull/4669)                                                  |   | 84d17h  |
| tidb/14831   | [planner/cascades: add implementationRule for IndexLookUpJoin](https://github.com/pingcap/tidb/pull/14831)                                           |   | 322d18h |
| tidb/15090   | [planner/cascades: refine the row count estimation of TiKV layer Selection](https://github.com/pingcap/tidb/pull/15090)                              |   | 308d18h |
| tidb/15157   | [planner/cascades: implement `HashCode` method for all the LogicalPlans](https://github.com/pingcap/tidb/pull/15157)                                 | Y | 306d15h |
| tidb/15426   | [planner/cascades: add transformation rule PushSelDownApply & refactor PushSelDownJoin](https://github.com/pingcap/tidb/pull/15426)                  |   | 294d17h |
| tidb/16967   | [executor: Refactor Shuffle and implement parallel sort (executor part)](https://github.com/pingcap/tidb/pull/16967)                                 | Y | 248d10h |
| tidb/17414   | [add curCost based join reorder algorithm](https://github.com/pingcap/tidb/pull/17414)                                                               |   | 224d18h |
| tidb/17996   | [planner: push avg & distinct functions across join](https://github.com/pingcap/tidb/pull/17996)                                                     | Y | 206d11h |
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                               | Y | 116d14h |
| tidb/20011   | [statistics: fix incorrect total count used in index selectivity computation](https://github.com/pingcap/tidb/pull/20011)                            |   | 112d15h |
| tidb/20311   | [expression: fix overflow error when convert bit to int64 (#20266)](https://github.com/pingcap/tidb/pull/20311)                                      |   | 97d21h  |
| tidb/20765   | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                                    |   | 64d17h  |
| tidb/20877   | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                                            |   | 61d17h  |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                       |   | 54d17h  |
| tidb/21137   | [executor: specially handle empty input for apply's outer child aggregate (#20544)](https://github.com/pingcap/tidb/pull/21137)                      |   | 47d20h  |
| tidb/21207   | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                                      |   | 43d19h  |
| tidb/21230   | [planner, executor: fix haven't track the memory usage of PointGet/BatchPointGet](https://github.com/pingcap/tidb/pull/21230)                        |   | 43d10h  |
| tidb/21408   | [statistics: fix a bug which causes panic when using the clustered index and the new collation (#21379)](https://github.com/pingcap/tidb/pull/21408) |   | 35d20h  |
| tidb/21425   | [planner: natural join not consider rowid and null eq not propagate (#21328)](https://github.com/pingcap/tidb/pull/21425)                            |   | 34d22h  |
| tidb/21476   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                |   | 33d16h  |
| tidb/21477   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21477)                                                |   | 33d16h  |
| tidb/21487   | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                                        |   | 33d5h   |
| tidb/21614   | [planner: do not propagate column eq with different column types (#21495)](https://github.com/pingcap/tidb/pull/21614)                               |   | 27d14h  |
| tidb/21714   | [planner: fix the coercibility of the cast function (#21705)](https://github.com/pingcap/tidb/pull/21714)                                            |   | 22d17h  |
| tidb/21808   | [planner: fix the fail when we compare multi fields in the subquery (#21699)](https://github.com/pingcap/tidb/pull/21808)                            |   | 21d18h  |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                          |   | 18d20h  |
| tidb/21877   | [planner: fix correlated aggregates which should be evaluated in outer query (#21431)](https://github.com/pingcap/tidb/pull/21877)                   |   | 18d19h  |
| tidb/21930   | [planner: propagate NDV of column groups across plan nodes (#17854)](https://github.com/pingcap/tidb/pull/21930)                                     |   | 14d18h  |
| tidb/21957   | [planner: fix unknown columns in join using below agg (#21922)](https://github.com/pingcap/tidb/pull/21957)                                          |   | 13d23h  |
| tidb/21964   | [planner: add plancodec id for all type TableScan/IndexScan. (#21935)](https://github.com/pingcap/tidb/pull/21964)                                   |   | 13d18h  |
| tidb/21976   | [planner: report error for invalid window specs which are not used (#21083)](https://github.com/pingcap/tidb/pull/21976)                             |   | 13d16h  |
| tidb/22022   | [planner/codec: fix issue of decode plan error cause by without escape special char (#21937)](https://github.com/pingcap/tidb/pull/22022)            |   | 12d0h   |
| tidb/22090   | [planner: push aggregation operators down to projection and union by default](https://github.com/pingcap/tidb/pull/22090)                            |   | 6d22h   |
| tidb/22118   | [planner: check if columns count matches for batch point get in TryFastPlan (#22044)](https://github.com/pingcap/tidb/pull/22118)                    |   | 5d23h   |
| tidb/22127   | [*: support ALTER TABLE ADD / DROP TIDB_STATS](https://github.com/pingcap/tidb/pull/22127)                                                           |   | 5d20h   |
| tidb/22143   | [expression: return correct results for user variables of datetime type (#22078)](https://github.com/pingcap/tidb/pull/22143)                        |   | 1d23h   |
| tidb/22173   | [expression: fix unexpected panic when doing isNullRejected check](https://github.com/pingcap/tidb/pull/22173)                                       |   | 21h     |
| tidb/22182   | [bindinfo: avoid duplicate bindings caused by concurrent baseline capture](https://github.com/pingcap/tidb/pull/22182)                               |   | 18h     |


## Reviewed in Last 7 Days

|    REPO     |                                                        PR                                                         | C | D |   R    |
|-------------|-------------------------------------------------------------------------------------------------------------------|---|---|--------|
| parser/1138 | [*: create / drop extended stats by ALTER TABLE](https://github.com/pingcap/parser/pull/1138)                     |   | 1 | 11d20h |
| tidb/22156  | [mod: update parser to latest 4.0 version](https://github.com/pingcap/tidb/pull/22156)                            |   | 1 | 22h    |
| tidb/22086  | [planner/core: fix a bug of adding enforcer.](https://github.com/pingcap/tidb/pull/22086)                         |   | 2 | 6d3h   |
| tidb/21014  | [statistics: GC index usage information](https://github.com/pingcap/tidb/pull/21014)                              |   | 2 | 53d0h  |
| tidb/22033  | [statistics: redesign the schema for `mysql.stats_extended`](https://github.com/pingcap/tidb/pull/22033)          |   | 6 | 5d22h  |
| docs/4546   | [Add version 2 info in statistics](https://github.com/pingcap/docs/pull/4546)                                     |   | 7 | 0h     |
| tidb/22080  | [planner, expression: fix error when using IN combined with subquery](https://github.com/pingcap/tidb/pull/22080) |   | 7 | 21h    |


</details> 


<details> 
  <summary>wshwsh12's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                               PR                                                               | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/15462 | [executor: implement `graceHashJoin`](https://github.com/pingcap/tidb/pull/15462)                                              | Y | 293d18h |
| tidb/17996 | [planner: push avg & distinct functions across join](https://github.com/pingcap/tidb/pull/17996)                               | Y | 206d11h |
| tidb/19807 | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                        |   | 123d11h |
| tidb/19957 | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                         | Y | 116d14h |
| tidb/20044 | [expression: Add column nullability checking before "refine args"](https://github.com/pingcap/tidb/pull/20044)                 | Y | 111d7h  |
| tidb/21381 | [*: optimize analyze cluster index table](https://github.com/pingcap/tidb/pull/21381)                                          |   | 36d17h  |
| tidb/21487 | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                  |   | 33d5h   |
| tidb/21541 | [executor: Nested prepare stmt should not be prepared](https://github.com/pingcap/tidb/pull/21541)                             |   | 29d13h  |
| tidb/21839 | [planner/core: add 'split table using statistics' statement](https://github.com/pingcap/tidb/pull/21839)                       |   | 20d15h  |
| tidb/21887 | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                |   | 17d11h  |
| tidb/21945 | [distsql: fix cop stats string display when there is only 1 rpc (#21901)](https://github.com/pingcap/tidb/pull/21945)          |   | 14d14h  |
| tidb/21957 | [planner: fix unknown columns in join using below agg (#21922)](https://github.com/pingcap/tidb/pull/21957)                    |   | 13d23h  |
| tidb/21988 | [*: Move tikv gc configuration to sysvars](https://github.com/pingcap/tidb/pull/21988)                                         |   | 13d2h   |
| tidb/22110 | [config, session: promise the compatibility of oom-action when upgrading (#22102)](https://github.com/pingcap/tidb/pull/22110) |   | 6d13h   |


## Reviewed in Last 7 Days

|    REPO    |                                                               PR                                                               | C | D |   R    |
|------------|--------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/21916 | [server: double type column from table should ignore its decimal (#21788)](https://github.com/pingcap/tidb/pull/21916)         |   | 1 | 14d23h |
| tidb/21969 | [types:  Add a limitation about float data type (#20929)](https://github.com/pingcap/tidb/pull/21969)                          |   | 1 | 13d17h |
| tidb/22084 | [expression: set collation function flen](https://github.com/pingcap/tidb/pull/22084)                                          |   | 1 | 6d21h  |
| tidb/22116 | [util: add cache for mem info in container (#22109)](https://github.com/pingcap/tidb/pull/22116)                               |   | 7 | 12h    |
| tidb/22111 | [config, session: promise the compatibility of oom-action when upgrading (#22102)](https://github.com/pingcap/tidb/pull/22111) |   | 7 | 1h     |
| tidb/22115 | [util: add cache for mem info in container (#22109)](https://github.com/pingcap/tidb/pull/22115)                               |   | 7 | 0h     |
| tidb/22112 | [bootstrap: add mysql.user columns for parser#1121 (#21856)](https://github.com/pingcap/tidb/pull/22112)                       |   | 7 | 0h     |
| tidb/22109 | [util: add cache for mem info in container](https://github.com/pingcap/tidb/pull/22109)                                        |   | 7 | 0h     |
| tidb/22102 | [config, session: promise the compatibility of oom-action when upgrading](https://github.com/pingcap/tidb/pull/22102)          |   | 7 | 4h     |


</details> 

