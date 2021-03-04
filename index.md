|   REVIEWER    |  O(C)   | 1 | 2 | 3 | 4 | 5 | 6 | 7 | T  |
|---------------|---------|---|---|---|---|---|---|---|----|
| Reminiscent   |  3 ( 0) | 5 | 1 | 3 | 0 | 0 | 1 | 0 | 10 |
| SunRunAway    | 19 ( 1) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| XuHuaiyu      | 53 ( 3) | 1 | 3 | 0 | 0 | 0 | 2 | 1 |  7 |
| dyzsr         |  0 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| eurekaka      | 13 ( 1) | 0 | 1 | 0 | 0 | 0 | 0 | 0 |  1 |
| fzhedu        |  3 ( 0) | 0 | 1 | 2 | 0 | 0 | 0 | 0 |  3 |
| ichn-hu       |  1 ( 0) | 1 | 2 | 0 | 0 | 0 | 2 | 0 |  5 |
| lzmhhh123     | 35 ( 0) | 2 | 0 | 4 | 0 | 0 | 1 | 1 |  8 |
| nullnotnil    |  0 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| qw4990        | 53 ( 1) | 5 | 2 | 5 | 0 | 0 | 0 | 0 | 12 |
| rebelice      |  2 ( 0) | 2 | 2 | 0 | 0 | 0 | 3 | 1 |  8 |
| time-and-fate |  4 ( 0) | 3 | 0 | 0 | 0 | 0 | 0 | 0 |  3 |
| winoros       | 24 ( 2) | 4 | 8 | 1 | 0 | 0 | 0 | 1 | 14 |
| wshwsh12      | 15 ( 1) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |


<details> 
  <summary>Reminiscent's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                   PR                                                                   | C | LASTED |
|------------|----------------------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/21896 | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                             |   | 72d19h |
| tidb/22354 | [planner: do not cache prepared plan if optimization depends on mutable constant (#22349)](https://github.com/pingcap/tidb/pull/22354) |   | 50d23h |
| tidb/23074 | [planner: fix range partition prune bug for IN expr (#22894) (#22938)](https://github.com/pingcap/tidb/pull/23074)                     |   | 17h    |


## Reviewed in Last 7 Days

|    REPO     |                                                                        PR                                                                        | C | D |   R    |
|-------------|--------------------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/23075  | [statistics: fix merge idx hist with poped topn](https://github.com/pingcap/tidb/pull/23075)                                                     |   | 1 | 0h     |
| tidb/22918  | [sessionctx: add optimization-time and wait-TS-time into the slow log (#17869)](https://github.com/pingcap/tidb/pull/22918)                      |   | 1 | 6d22h  |
| tidb/23061  | [statistics: introduce a new kind of syntax to drop global-stats](https://github.com/pingcap/tidb/pull/23061)                                    |   | 1 | 0h     |
| parser/1181 | [statistics: introduce a new kind of syntax to drop global-stats](https://github.com/pingcap/parser/pull/1181)                                   |   | 1 | 0h     |
| tidb/23050  | [statistics: forbid global-stats in stats-ver1 and make analyze options take effect in global-stats](https://github.com/pingcap/tidb/pull/23050) |   | 1 | 19h    |
| tidb/23023  | [statistics: support dropping partition/global level statistics](https://github.com/pingcap/tidb/pull/23023)                                     |   | 2 | 1d0h   |
| parser/1177 | [statistics: add a new syntax to support dropping statistics of partitions](https://github.com/pingcap/parser/pull/1177)                         |   | 3 | 0h     |
| tidb/22625  | [planner, statistics: allow (auto) analyze single partition in dynamic-only mode](https://github.com/pingcap/tidb/pull/22625)                    |   | 3 | 30d20h |
| tidb/23016  | [statistics: add static partition prune mode checks](https://github.com/pingcap/tidb/pull/23016)                                                 |   | 3 | 2h     |
| tidb/22948  | [statistics: add more tests to check accurateness of global-stats](https://github.com/pingcap/tidb/pull/22948)                                   |   | 6 | 17h    |


</details> 


<details> 
  <summary>SunRunAway's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                  PR                                                                   | C | LASTED  |
|------------|---------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19178 | [executor: Refactor probe channel](https://github.com/pingcap/tidb/pull/19178)                                                        |   | 202d16h |
| tidb/19347 | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347)  |   | 194d23h |
| tidb/19807 | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                               |   | 180d10h |
| tidb/19900 | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                        | Y | 175d18h |
| tidb/20140 | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                    |   | 162d22h |
| tidb/20360 | [planner: refine explain info for batch cop](https://github.com/pingcap/tidb/pull/20360)                                              |   | 145d22h |
| tidb/20765 | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                     |   | 121d17h |
| tidb/21207 | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                       |   | 100d19h |
| tidb/21277 | [executor: fix split table with large integers](https://github.com/pingcap/tidb/pull/21277)                                           |   | 98d19h  |
| tidb/21834 | [planner: enhanced index range calculation plan](https://github.com/pingcap/tidb/pull/21834)                                          |   | 77d18h  |
| tidb/21876 | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)           |   | 75d19h  |
| tidb/21878 | [planner: do not push down lock to pointGet/bacthPointGet when selection exists](https://github.com/pingcap/tidb/pull/21878)          |   | 75d18h  |
| tidb/21956 | [planner/preprocessor: disallow into-outfile clause in some place](https://github.com/pingcap/tidb/pull/21956)                        |   | 70d23h  |
| tidb/22026 | [expression: separated arithmeticPlusIntSig](https://github.com/pingcap/tidb/pull/22026)                                              |   | 68d20h  |
| tidb/22114 | [test: fix globalkilltest (#21987)](https://github.com/pingcap/tidb/pull/22114)                                                       |   | 63d12h  |
| tidb/22181 | [planner, expression: fix error when using IN combined with subquery (#22080)](https://github.com/pingcap/tidb/pull/22181)            |   | 57d17h  |
| tidb/22217 | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                         |   | 56d17h  |
| tidb/22365 | [planner: check index valid while forUpdateRead (#22152)](https://github.com/pingcap/tidb/pull/22365)                                 |   | 50d19h  |
| tidb/22379 | [[experiment] executor: allow aggregation to spill disk when running out of memory quota](https://github.com/pingcap/tidb/pull/22379) |   | 49d19h  |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>XuHuaiyu's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                              PR                                                                              | C | LASTED  |
|--------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| docs-cn/5619 | [Update data-type-date-and-time.md](https://github.com/pingcap/docs-cn/pull/5619)                                                                            |   | 5d15h   |
| tidb/19900   | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                                               | Y | 175d18h |
| docs-cn/5620 | [Add details for Hexadecimal Literals](https://github.com/pingcap/docs-cn/pull/5620)                                                                         |   | 5d15h   |
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                                       | Y | 173d13h |
| tidb/20040   | [planner, expression: take NullFlag into consideration when optimize the `int non-const` <cmp > `non-int const`](https://github.com/pingcap/tidb/pull/20040) | Y | 168d13h |
| tidb/20140   | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                                           |   | 162d22h |
| tidb/20311   | [expression: fix overflow error when convert bit to int64 (#20266)](https://github.com/pingcap/tidb/pull/20311)                                              |   | 154d21h |
| tidb/20790   | [collation: add pinyin collation for chinese charset support](https://github.com/pingcap/tidb/pull/20790)                                                    |   | 120d20h |
| tidb/20905   | [planner: fix statement-optimize not work in `TryFastPlan`](https://github.com/pingcap/tidb/pull/20905)                                                      |   | 117d17h |
| tidb/20972   | [expression: POC implementation of Vitess hashing algorithm.](https://github.com/pingcap/tidb/pull/20972)                                                    |   | 113d0h  |
| tidb/21064   | [planner, executor: fix cast not check error](https://github.com/pingcap/tidb/pull/21064)                                                                    |   | 108d8h  |
| tidb/21149   | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                                |   | 104d14h |
| tidb/21228   | [executor: return the result immediately when combining LIMIT row_count with DISTINCT](https://github.com/pingcap/tidb/pull/21228)                           |   | 100d13h |
| tidb/21304   | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                                     |   | 98d12h  |
| tidb/21334   | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                                                |   | 97d14h  |
| tidb/21340   | [executor: initialize expensive query handler on domain creation](https://github.com/pingcap/tidb/pull/21340)                                                |   | 96d23h  |
| tidb/21476   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                        |   | 90d15h  |
| tidb/21536   | [executor: add slow-log file meta cache to avoid repeat read file meta information](https://github.com/pingcap/tidb/pull/21536)                              |   | 86d14h  |
| tidb/21564   | [ddl: fix Incorrect behavior of NO_ZERO_DATE when altering table](https://github.com/pingcap/tidb/pull/21564)                                                |   | 85d15h  |
| tidb/21853   | [expression: fix compatibility behaviors in time_format with MySQL (#21559)](https://github.com/pingcap/tidb/pull/21853)                                     |   | 76d18h  |
| tidb/21896   | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                                                   |   | 72d19h  |
| tidb/22131   | [privilege: remove leading and trailing space when create user and role](https://github.com/pingcap/tidb/pull/22131)                                         |   | 62d19h  |
| tidb/22149   | [session: set process info before building plan (#22101)](https://github.com/pingcap/tidb/pull/22149)                                                        |   | 58d19h  |
| tidb/22163   | [expression: separated arithmeticMinusIntSig](https://github.com/pingcap/tidb/pull/22163)                                                                    |   | 58d13h  |
| tidb/22186   | [executor: fix select into outfile with year type column has no data (#22175)](https://github.com/pingcap/tidb/pull/22186)                                   |   | 57d16h  |
| tidb/22294   | [planner, table: optimize the list partition pruner for range query](https://github.com/pingcap/tidb/pull/22294)                                             |   | 54d20h  |
| tidb/22307   | [ddl: fix update can see columns not public](https://github.com/pingcap/tidb/pull/22307)                                                                     |   | 54d16h  |
| tidb/22381   | [planner: check schema stale for plan cache when forUpdateRead](https://github.com/pingcap/tidb/pull/22381)                                                  |   | 49d14h  |
| tidb/22616   | [expression: from_unixtime accept 64-bit integers](https://github.com/pingcap/tidb/pull/22616)                                                               |   | 33d23h  |
| tidb/22617   | [metrics: fix wrong bucket name of coprocessor cache (#22454)](https://github.com/pingcap/tidb/pull/22617)                                                   |   | 33d23h  |
| tidb/22624   | [ planner: not pruning column used by union scan condition (#21640)](https://github.com/pingcap/tidb/pull/22624)                                             |   | 33d17h  |
| tidb/22631   | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                                                                              |   | 31d23h  |
| tidb/22640   | [*: refactor ExecuteInternal to return single resultset (#22546)](https://github.com/pingcap/tidb/pull/22640)                                                |   | 30d20h  |
| tidb/22696   | [expression: enable arithmetic Mod push down](https://github.com/pingcap/tidb/pull/22696)                                                                    |   | 28d17h  |
| tidb/22711   | [executor: Fix inline schema name](https://github.com/pingcap/tidb/pull/22711)                                                                               |   | 28d11h  |
| tidb/22722   | [planner, errno: make error code of ErrMixOfGroupFuncAndFields consistent with MySQL](https://github.com/pingcap/tidb/pull/22722)                            |   | 27d20h  |
| tidb/22736   | [executor: fix load data losing connection when batch_dml_size is set (#22724)](https://github.com/pingcap/tidb/pull/22736)                                  |   | 26d23h  |
| tidb/22786   | [config: deprecate some configs of `tikv-client.copr-cache`](https://github.com/pingcap/tidb/pull/22786)                                                     |   | 13d18h  |
| tidb/22814   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22814)                                          |   | 12d19h  |
| tidb/22815   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22815)                                          |   | 12d19h  |
| tidb/22832   | [expression: push down EXTRACT to TiFlash](https://github.com/pingcap/tidb/pull/22832)                                                                       |   | 12d1h   |
| tidb/22844   | [expression: do not adjust int when it is null and compared year (#22821)](https://github.com/pingcap/tidb/pull/22844)                                       |   | 11d19h  |
| tidb/22886   | [*: rename tidb_enable_tiflash_fallback_tikv to tidb_enable_engine_fallback](https://github.com/pingcap/tidb/pull/22886)                                     |   | 8d19h   |
| tidb/22914   | [partition: fix hash partition with not between condition get wrong result](https://github.com/pingcap/tidb/pull/22914)                                      |   | 7d17h   |
| tidb/22976   | [kv/union_store:remove tableinfo from union_store](https://github.com/pingcap/tidb/pull/22976)                                                               |   | 5d17h   |
| tidb/23012   | [executor: fix affected rows of ddls and complete uint tests](https://github.com/pingcap/tidb/pull/23012)                                                    |   | 3d16h   |
| tidb/23024   | [executor: make the memory tracker of Jsonobjectagg more accurate](https://github.com/pingcap/tidb/pull/23024)                                               |   | 2d18h   |
| tidb/23026   | [ddl: fix the create view clause will be restored with charset prefix](https://github.com/pingcap/tidb/pull/23026)                                           |   | 2d17h   |
| tidb/23056   | [MPP: Kill mpp queries](https://github.com/pingcap/tidb/pull/23056)                                                                                          |   | 1d12h   |
| tidb/23066   | [statistics: forbid getting global-stats through fast analyze and incremental analyze](https://github.com/pingcap/tidb/pull/23066)                           |   | 19h     |
| tidb/23072   | [executor: track memory usage of map in agg partial result.](https://github.com/pingcap/tidb/pull/23072)                                                     |   | 18h     |
| tidb/23087   | [executor: fix correlated column range in table reader for the clustered index table](https://github.com/pingcap/tidb/pull/23087)                            |   | 14h     |
| tidb/23092   | [*: fix a bug that collation is not handle for text type (#23045)](https://github.com/pingcap/tidb/pull/23092)                                               |   | 12h     |


## Reviewed in Last 7 Days

|    REPO    |                                                       PR                                                        | C | D |   R    |
|------------|-----------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/23034 | [executor: make the memory tracker of groupConcat more accurate.](https://github.com/pingcap/tidb/pull/23034)   |   | 1 | 1d23h  |
| tidb/23024 | [executor: make the memory tracker of Jsonobjectagg more accurate](https://github.com/pingcap/tidb/pull/23024)  |   | 2 | 1d15h  |
| tidb/22962 | [executor: track partialResultMap in unparalleled aggreagte.](https://github.com/pingcap/tidb/pull/22962)       |   | 2 | 4d3h   |
| tidb/22903 | [docs/design: update template](https://github.com/pingcap/tidb/pull/22903)                                      |   | 2 | 6d15h  |
| tidb/22507 | [types: fix the bug about the wrong query result for decimal type ](https://github.com/pingcap/tidb/pull/22507) |   | 6 | 32d3h  |
| tidb/22426 | [expression: fix bugs in builtinfunction ArithmeticMinusInt logic](https://github.com/pingcap/tidb/pull/22426)  |   | 6 | 38d17h |
| tidb/22912 | [executor: reduce useless memory allocation in min/max aggregate](https://github.com/pingcap/tidb/pull/22912)   |   | 7 | 20h    |


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
| tidb/19347 | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347)   |   | 194d23h |
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                              |   | 118d16h |
| tidb/21444 | [planner: ignore anonymous index while tiflash replica is available](https://github.com/pingcap/tidb/pull/21444)                       |   | 91d12h  |
| tidb/21994 | [range: fix overflow value access index ](https://github.com/pingcap/tidb/pull/21994)                                                  |   | 69d22h  |
| tidb/22342 | [session: fix two cases when updating bind info (#22338)](https://github.com/pingcap/tidb/pull/22342)                                  |   | 51d18h  |
| tidb/22354 | [planner: do not cache prepared plan if optimization depends on mutable constant (#22349)](https://github.com/pingcap/tidb/pull/22354) |   | 50d23h  |
| tidb/22369 | [session: fix the duplicate binding case when updating bind info (#22367)](https://github.com/pingcap/tidb/pull/22369)                 |   | 50d17h  |
| tidb/22416 | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                                      | Y | 46d11h  |
| tidb/22559 | [planner: split test data from test cases in cbo_test.go](https://github.com/pingcap/tidb/pull/22559)                                  |   | 35d19h  |
| tidb/22733 | [bindinfo: use new sql apis (#22653)](https://github.com/pingcap/tidb/pull/22733)                                                      |   | 27d15h  |
| tidb/22778 | [*: add support for dynamic privileges](https://github.com/pingcap/tidb/pull/22778)                                                    |   | 15d7h   |
| tidb/22953 | [planner: fix query range partition table got wrong result and TiDB panic](https://github.com/pingcap/tidb/pull/22953)                 |   | 6d14h   |
| tidb/23041 | [util: remove unused code](https://github.com/pingcap/tidb/pull/23041)                                                                 |   | 2d9h    |


## Reviewed in Last 7 Days

|    REPO    |                                                             PR                                                             | C | D | R  |
|------------|----------------------------------------------------------------------------------------------------------------------------|---|---|----|
| tidb/23047 | [planner, expression: fix error when using IN combined with subquery (#22080)](https://github.com/pingcap/tidb/pull/23047) |   | 2 | 0h |


</details> 


<details> 
  <summary>fzhedu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                         PR                                                          | C | LASTED |
|------------|---------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/22662 | [planner/core: let mpp support partition tables](https://github.com/pingcap/tidb/pull/22662)                        |   | 29d18h |
| tidb/22845 | [planner: fix bug of mpp wrongly set schema of exchanger](https://github.com/pingcap/tidb/pull/22845)               |   | 11d17h |
| tidb/22853 | [planner: fix LogicalPlans that contain Window Function are ambiguous ](https://github.com/pingcap/tidb/pull/22853) |   | 11d12h |


## Reviewed in Last 7 Days

|    REPO    |                                                                  PR                                                                   | C | D |   R   |
|------------|---------------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/22803 | [store/mockstore/unistore: refine and add more mpp tests](https://github.com/pingcap/tidb/pull/22803)                                 |   | 2 | 11d9h |
| tics/1342  | [Add background task to detect && cancel hanging MPP query](https://github.com/pingcap/tics/pull/1342)                                |   | 3 | 48d2h |
| tidb/23020 | [expression: Add warning info for exprs that can not be pushed to storage layer (#22713)](https://github.com/pingcap/tidb/pull/23020) |   | 3 | 0h    |


</details> 


<details> 
  <summary>ichn-hu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                            PR                                                            | C | LASTED |
|------------|--------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/21853 | [expression: fix compatibility behaviors in time_format with MySQL (#21559)](https://github.com/pingcap/tidb/pull/21853) |   | 76d18h |


## Reviewed in Last 7 Days

|    REPO    |                                                          PR                                                          | C | D |  R   |
|------------|----------------------------------------------------------------------------------------------------------------------|---|---|------|
| tidb/23072 | [executor: track memory usage of map in agg partial result.](https://github.com/pingcap/tidb/pull/23072)             |   | 1 | 0h   |
| tidb/23034 | [executor: make the memory tracker of groupConcat more accurate.](https://github.com/pingcap/tidb/pull/23034)        |   | 2 | 23h  |
| tidb/23024 | [executor: make the memory tracker of Jsonobjectagg more accurate](https://github.com/pingcap/tidb/pull/23024)       |   | 2 | 1d1h |
| tidb/22965 | [executor: introduce setWithMemoryUsage to track set memory in AggExec.](https://github.com/pingcap/tidb/pull/22965) |   | 6 | 5h   |
| tidb/22962 | [executor: track partialResultMap in unparalleled aggreagte.](https://github.com/pingcap/tidb/pull/22962)            |   | 6 | 5h   |


</details> 


<details> 
  <summary>lzmhhh123's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                             PR                                                                              | C | LASTED  |
|------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19347 | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347)                        |   | 194d23h |
| tidb/20444 | [expression: add json_merge_patch](https://github.com/pingcap/tidb/pull/20444)                                                                              |   | 140d21h |
| tidb/20465 | [expression: add uuidShortFunction](https://github.com/pingcap/tidb/pull/20465)                                                                             |   | 139d19h |
| tidb/20642 | [executor: modify admin executors to support partitioned table with global index](https://github.com/pingcap/tidb/pull/20642)                               |   | 128d15h |
| tidb/20903 | [planner: fix confused and unnecessary double-projection in plans.](https://github.com/pingcap/tidb/pull/20903)                                             |   | 117d17h |
| tidb/21018 | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                              |   | 111d16h |
| tidb/21195 | [brie: integrate lightning to suport IMPORT statement](https://github.com/pingcap/tidb/pull/21195)                                                          |   | 100d22h |
| tidb/21334 | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                                               |   | 97d14h  |
| tidb/21347 | [session: make rollback work on global variables](https://github.com/pingcap/tidb/pull/21347)                                                               |   | 96d19h  |
| tidb/21401 | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                                          |   | 93d11h  |
| tidb/21444 | [planner: ignore anonymous index while tiflash replica is available](https://github.com/pingcap/tidb/pull/21444)                                            |   | 91d12h  |
| tidb/21487 | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                                               |   | 90d4h   |
| tidb/21641 | [executor: Fix pessimistic lock doesn't work on the partition table for subquery/joins](https://github.com/pingcap/tidb/pull/21641)                         |   | 83d18h  |
| tidb/21651 | [planner: allow filter condition pushing down to IndexScan for prefix index](https://github.com/pingcap/tidb/pull/21651)                                    |   | 83d13h  |
| tidb/22126 | [*: add `sys` schema, `sys.SCHEMA_UNUSED_INDEXES` view and `sys.SCHEMA_INDEX_USAGE` view](https://github.com/pingcap/tidb/pull/22126)                       |   | 62d19h  |
| tidb/22149 | [session: set process info before building plan (#22101)](https://github.com/pingcap/tidb/pull/22149)                                                       |   | 58d19h  |
| tidb/22188 | [planner: do not use indexMerge when the path only use a single index (#22168)](https://github.com/pingcap/tidb/pull/22188)                                 |   | 57d13h  |
| tidb/22361 | [table: fix insert into _tidb_rowid panic and rebase it if needed (#22062)](https://github.com/pingcap/tidb/pull/22361)                                     |   | 50d20h  |
| tidb/22372 | [executor: fix SelectForUpdate in decorrelated subquery under pessimistic mode](https://github.com/pingcap/tidb/pull/22372)                                 |   | 50d9h   |
| tidb/22478 | [planner, executor: fix query partition table with global unique index get wrong result](https://github.com/pingcap/tidb/pull/22478)                        |   | 41d13h  |
| tidb/22631 | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                                                                             |   | 31d23h  |
| tidb/22662 | [planner/core: let mpp support partition tables](https://github.com/pingcap/tidb/pull/22662)                                                                |   | 29d18h  |
| tidb/22699 | [brie: add error info column and history backup/restore info in sql](https://github.com/pingcap/tidb/pull/22699)                                            |   | 28d16h  |
| tidb/22714 | [executor: add close recordSet in executor](https://github.com/pingcap/tidb/pull/22714)                                                                     |   | 27d23h  |
| tidb/22812 | [ executor: add new format specifier(%# %@ %.) for str_to_date expression (#22790)](https://github.com/pingcap/tidb/pull/22812)                             |   | 12d19h  |
| tidb/22834 | [executor, server: load_data.go is changed and add unit test](https://github.com/pingcap/tidb/pull/22834)                                                   |   | 12d0h   |
| tidb/22857 | [store: move mockstore/mocktikv to tikv/mockstore/mocktikv](https://github.com/pingcap/tidb/pull/22857)                                                     |   | 10d21h  |
| tidb/22910 | [util: optimize the performance of restore with db](https://github.com/pingcap/tidb/pull/22910)                                                             |   | 7d19h   |
| tidb/22974 | [parser: update parser to fix sql restore bug used in create view](https://github.com/pingcap/tidb/pull/22974)                                              |   | 5d17h   |
| tidb/23001 | [statistics: fix err check](https://github.com/pingcap/tidb/pull/23001)                                                                                     |   | 4d0h    |
| tidb/23022 | [executor: create PipelinedWindowExec based on current implementation and modify the windowProcessor interface](https://github.com/pingcap/tidb/pull/23022) |   | 2d18h   |
| tidb/23049 | [statistics: add more test cases for global-level stats](https://github.com/pingcap/tidb/pull/23049)                                                        |   | 1d18h   |
| tidb/23070 | [*: make linter enable `unused` #23013](https://github.com/pingcap/tidb/pull/23070)                                                                         |   | 18h     |
| tidb/23091 | [*: fix wrong replace or insert-on-dup behavior on prefixed clustered index](https://github.com/pingcap/tidb/pull/23091)                                    |   | 12h     |
| tidb/23094 | [planner: fix indexJoin(also hash, merge) on prefixed clustered index](https://github.com/pingcap/tidb/pull/23094)                                          |   | 9h      |


## Reviewed in Last 7 Days

|     REPO     |                                                                PR                                                                | C | D |   R   |
|--------------|----------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/22965   | [executor: introduce setWithMemoryUsage to track set memory in AggExec.](https://github.com/pingcap/tidb/pull/22965)             |   | 1 | 4d23h |
| tikv/9383    | [copr: support decoding the new index format](https://github.com/tikv/tikv/pull/9383)                                            |   | 1 | 64d0h |
| tidb/22980   | [planner: choose non-prefix column when both index key and handle have the same one](https://github.com/pingcap/tidb/pull/22980) |   | 3 | 2d18h |
| docs-cn/5561 | [Add sql optimization-related docs to toc](https://github.com/pingcap/docs-cn/pull/5561)                                         |   | 3 | 6d20h |
| tidb/22991   | [executor: fix err check](https://github.com/pingcap/tidb/pull/22991)                                                            |   | 3 | 1d5h  |
| tidb/22940   | [planner: enable column pruning for common handle](https://github.com/pingcap/tidb/pull/22940)                                   |   | 3 | 3d22h |
| tidb/22959   | [planner, executor: reset NotNullFlag when merge schema for join (#22955)](https://github.com/pingcap/tidb/pull/22959)           |   | 6 | 11h   |
| tidb/22955   | [planner, executor: reset NotNullFlag when merge schema for join](https://github.com/pingcap/tidb/pull/22955)                    |   | 7 | 0h    |


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
| docs-cn/5484 | [system variable: add tidb_enable_engine_fallback](https://github.com/pingcap/docs-cn/pull/5484)                                                            |   | 28d17h  |
| tidb/19029   | [types: fix unexpected NOT_NULL flags](https://github.com/pingcap/tidb/pull/19029)                                                                          |   | 209d22h |
| docs-cn/5561 | [Add sql optimization-related docs to toc](https://github.com/pingcap/docs-cn/pull/5561)                                                                    |   | 9d15h   |
| tidb/20708   | [*: separate auto_increment ID allocator from _tidb_rowid allocator](https://github.com/pingcap/tidb/pull/20708)                                            |   | 125d20h |
| tidb/20969   | [executor: Improve the performance of appending not fixed columns](https://github.com/pingcap/tidb/pull/20969)                                              |   | 113d9h  |
| tidb/20972   | [expression: POC implementation of Vitess hashing algorithm.](https://github.com/pingcap/tidb/pull/20972)                                                   |   | 113d0h  |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                              |   | 111d16h |
| tidb/21149   | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                               |   | 104d14h |
| tidb/21304   | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                                    |   | 98d12h  |
| tidb/21318   | [planner, expression: use the range of column types to simplify expressions](https://github.com/pingcap/tidb/pull/21318)                                    |   | 97d18h  |
| tidb/21401   | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                                          |   | 93d11h  |
| tidb/21476   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                       |   | 90d15h  |
| tidb/21508   | [execution: fix dayofweek('0000-00-00') behavior](https://github.com/pingcap/tidb/pull/21508)                                                               |   | 89d9h   |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                                 |   | 75d19h  |
| tidb/21887   | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                                             |   | 74d11h  |
| tidb/21930   | [planner: propagate NDV of column groups across plan nodes (#17854)](https://github.com/pingcap/tidb/pull/21930)                                            |   | 71d18h  |
| tidb/21954   | [planner/cascades: add rule `PushSelDownApply`](https://github.com/pingcap/tidb/pull/21954)                                                                 |   | 70d23h  |
| tidb/22090   | [planner: push aggregation operators down to projection and union by default](https://github.com/pingcap/tidb/pull/22090)                                   |   | 63d22h  |
| tidb/22146   | [executor: forbid SFU on view](https://github.com/pingcap/tidb/pull/22146)                                                                                  |   | 58d21h  |
| tidb/22217   | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                                               |   | 56d17h  |
| tidb/22234   | [executor, planner: ON DUPLICATE UPDATE can refer to un-project col (#14412)](https://github.com/pingcap/tidb/pull/22234)                                   |   | 56d15h  |
| tidb/22261   | [time: fix parse datetime won't truncate the reluctant string (#22232)](https://github.com/pingcap/tidb/pull/22261)                                         |   | 55d19h  |
| tidb/22307   | [ddl: fix update can see columns not public](https://github.com/pingcap/tidb/pull/22307)                                                                    |   | 54d16h  |
| tidb/22342   | [session: fix two cases when updating bind info (#22338)](https://github.com/pingcap/tidb/pull/22342)                                                       |   | 51d18h  |
| tidb/22369   | [session: fix the duplicate binding case when updating bind info (#22367)](https://github.com/pingcap/tidb/pull/22369)                                      |   | 50d17h  |
| tidb/22374   | [expression: separated arithmeticIntDivideSig](https://github.com/pingcap/tidb/pull/22374)                                                                  |   | 50d0h   |
| tidb/22415   | [ddl: refactor placement package](https://github.com/pingcap/tidb/pull/22415)                                                                               |   | 46d17h  |
| tidb/22471   | [ddl, executor: fix creating unique index without partition column error when enable-global-index is true](https://github.com/pingcap/tidb/pull/22471)      |   | 41d17h  |
| tidb/22507   | [types: fix the bug about the wrong query result for decimal type ](https://github.com/pingcap/tidb/pull/22507)                                             |   | 37d23h  |
| tidb/22541   | [expression: Support builtin function SOUNDEX](https://github.com/pingcap/tidb/pull/22541)                                                                  |   | 36d9h   |
| tidb/22559   | [planner: split test data from test cases in cbo_test.go](https://github.com/pingcap/tidb/pull/22559)                                                       |   | 35d19h  |
| tidb/22565   | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)                                          | Y | 35d17h  |
| tidb/22662   | [planner/core: let mpp support partition tables](https://github.com/pingcap/tidb/pull/22662)                                                                |   | 29d18h  |
| tidb/22733   | [bindinfo: use new sql apis (#22653)](https://github.com/pingcap/tidb/pull/22733)                                                                           |   | 27d15h  |
| tidb/22760   | [expression: fix wrong error info](https://github.com/pingcap/tidb/pull/22760)                                                                              |   | 24d0h   |
| tidb/22778   | [*: add support for dynamic privileges](https://github.com/pingcap/tidb/pull/22778)                                                                         |   | 15d7h   |
| tidb/22814   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22814)                                         |   | 12d19h  |
| tidb/22815   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22815)                                         |   | 12d19h  |
| tidb/22862   | [brie: fix the problem that ddl restored by BR via SQL is not replicated to downstream](https://github.com/pingcap/tidb/pull/22862)                         |   | 9d22h   |
| tidb/22869   | [executor: fix cast function will ignore tht error for point-get key construction](https://github.com/pingcap/tidb/pull/22869)                              |   | 9d16h   |
| tidb/22886   | [*: rename tidb_enable_tiflash_fallback_tikv to tidb_enable_engine_fallback](https://github.com/pingcap/tidb/pull/22886)                                    |   | 8d19h   |
| tidb/22915   | [planner: build correct MaxOneRow info from multi-column conditions](https://github.com/pingcap/tidb/pull/22915)                                            |   | 7d17h   |
| tidb/22923   | [expression: correct constant propagation for collation (#22666)](https://github.com/pingcap/tidb/pull/22923)                                               |   | 7d14h   |
| tidb/22924   | [planner: fix wrong index merge selection (#22825)](https://github.com/pingcap/tidb/pull/22924)                                                             |   | 7d13h   |
| tidb/22926   | [expression: add overflow check in multiplyInt](https://github.com/pingcap/tidb/pull/22926)                                                                 |   | 7d13h   |
| tidb/22984   | [executor: fix logging format of prepared statements (#16062)](https://github.com/pingcap/tidb/pull/22984)                                                  |   | 4d10h   |
| tidb/23008   | [types: fix err check](https://github.com/pingcap/tidb/pull/23008)                                                                                          |   | 4d0h    |
| tidb/23022   | [executor: create PipelinedWindowExec based on current implementation and modify the windowProcessor interface](https://github.com/pingcap/tidb/pull/23022) |   | 2d18h   |
| tidb/23062   | [*: fix structcheck lint warnings](https://github.com/pingcap/tidb/pull/23062)                                                                              |   | 19h     |
| tidb/23074   | [planner: fix range partition prune bug for IN expr (#22894) (#22938)](https://github.com/pingcap/tidb/pull/23074)                                          |   | 17h     |
| tidb/23086   | [statistics: make exponential backoff estimation more safe](https://github.com/pingcap/tidb/pull/23086)                                                     |   | 15h     |
| tidb/23088   | [statistics: delete extended stats cache item in current tidb synchronously](https://github.com/pingcap/tidb/pull/23088)                                    |   | 14h     |
| tidb/23089   | [statistics: report error for extended statistics register on partitioned tables](https://github.com/pingcap/tidb/pull/23089)                               |   | 14h     |


## Reviewed in Last 7 Days

|      REPO      |                                                              PR                                                               | C | D |   R    |
|----------------|-------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/23075     | [statistics: fix merge idx hist with poped topn](https://github.com/pingcap/tidb/pull/23075)                                  |   | 1 | 0h     |
| tidb/22803     | [store/mockstore/unistore: refine and add more mpp tests](https://github.com/pingcap/tidb/pull/22803)                         |   | 1 | 12d1h  |
| tidb/23057     | [statistics: fix the correlation estimation for version 2](https://github.com/pingcap/tidb/pull/23057)                        |   | 1 | 11h    |
| tidb/22910     | [util: optimize the performance of restore with db](https://github.com/pingcap/tidb/pull/22910)                               |   | 1 | 6d20h  |
| tidb/23047     | [planner, expression: fix error when using IN combined with subquery (#22080)](https://github.com/pingcap/tidb/pull/23047)    |   | 1 | 19h    |
| tidb/22662     | [planner/core: let mpp support partition tables](https://github.com/pingcap/tidb/pull/22662)                                  |   | 2 | 28d3h  |
| tidb/23049     | [statistics: add more test cases for global-level stats](https://github.com/pingcap/tidb/pull/23049)                          |   | 2 | 3h     |
| tidb/22625     | [planner, statistics: allow (auto) analyze single partition in dynamic-only mode](https://github.com/pingcap/tidb/pull/22625) |   | 3 | 30d20h |
| tidb/22931     | [statistics: enables global-level stats to be generated in fast analyze mode](https://github.com/pingcap/tidb/pull/22931)     |   | 3 | 4d5h   |
| tidb/22489     | [infoschema: support query partition_id from infoschema.partitions (#22240)](https://github.com/pingcap/tidb/pull/22489)      |   | 3 | 38d0h  |
| tidb/23016     | [statistics: add static partition prune mode checks](https://github.com/pingcap/tidb/pull/23016)                              |   | 3 | 2h     |
| tidb-test/1163 | [*: rename partition_prune_mode](https://github.com/pingcap/tidb-test/pull/1163)                                              |   | 3 | 6d0h   |


</details> 


<details> 
  <summary>rebelice's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                               PR                                                               | C | LASTED |
|------------|--------------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/22869 | [executor: fix cast function will ignore tht error for point-get key construction](https://github.com/pingcap/tidb/pull/22869) |   | 9d16h  |
| tidb/23074 | [planner: fix range partition prune bug for IN expr (#22894) (#22938)](https://github.com/pingcap/tidb/pull/23074)             |   | 17h    |


## Reviewed in Last 7 Days

|    REPO     |                                                                        PR                                                                        | C | D |  R   |
|-------------|--------------------------------------------------------------------------------------------------------------------------------------------------|---|---|------|
| tidb/23061  | [statistics: introduce a new kind of syntax to drop global-stats](https://github.com/pingcap/tidb/pull/23061)                                    |   | 1 | 0h   |
| parser/1181 | [statistics: introduce a new kind of syntax to drop global-stats](https://github.com/pingcap/parser/pull/1181)                                   |   | 1 | 0h   |
| tidb/23050  | [statistics: forbid global-stats in stats-ver1 and make analyze options take effect in global-stats](https://github.com/pingcap/tidb/pull/23050) |   | 2 | 4h   |
| tidb/23023  | [statistics: support dropping partition/global level statistics](https://github.com/pingcap/tidb/pull/23023)                                     |   | 2 | 21h  |
| tidb/22931  | [statistics: enables global-level stats to be generated in fast analyze mode](https://github.com/pingcap/tidb/pull/22931)                        |   | 6 | 1d4h |
| tidb/22948  | [statistics: add more tests to check accurateness of global-stats](https://github.com/pingcap/tidb/pull/22948)                                   |   | 6 | 20h  |
| tidb/22959  | [planner, executor: reset NotNullFlag when merge schema for join (#22955)](https://github.com/pingcap/tidb/pull/22959)                           |   | 6 | 11h  |
| tidb/22955  | [planner, executor: reset NotNullFlag when merge schema for join](https://github.com/pingcap/tidb/pull/22955)                                    |   | 7 | 1h   |


</details> 


<details> 
  <summary>time-and-fate's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                              PR                                                               | C | LASTED  |
|------------|-------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                     |   | 118d16h |
| tidb/22853 | [planner: fix LogicalPlans that contain Window Function are ambiguous ](https://github.com/pingcap/tidb/pull/22853)           |   | 11d12h  |
| tidb/23088 | [statistics: delete extended stats cache item in current tidb synchronously](https://github.com/pingcap/tidb/pull/23088)      |   | 14h     |
| tidb/23089 | [statistics: report error for extended statistics register on partitioned tables](https://github.com/pingcap/tidb/pull/23089) |   | 14h     |


## Reviewed in Last 7 Days

|    REPO    |                                                       PR                                                       | C | D |  R  |
|------------|----------------------------------------------------------------------------------------------------------------|---|---|-----|
| tidb/23086 | [statistics: make exponential backoff estimation more safe](https://github.com/pingcap/tidb/pull/23086)        |   | 1 | 1h  |
| tidb/23052 | [*: support `show stats_extended` to inspect extended stats cache](https://github.com/pingcap/tidb/pull/23052) |   | 1 | 20h |
| tidb/23057 | [statistics: fix the correlation estimation for version 2](https://github.com/pingcap/tidb/pull/23057)         |   | 1 | 11h |


</details> 


<details> 
  <summary>winoros's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                             PR                                                              | C | LASTED  |
|--------------|-----------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                      | Y | 173d13h |
| docs-cn/5484 | [system variable: add tidb_enable_engine_fallback](https://github.com/pingcap/docs-cn/pull/5484)                            |   | 28d17h  |
| tidb/20311   | [expression: fix overflow error when convert bit to int64 (#20266)](https://github.com/pingcap/tidb/pull/20311)             |   | 154d21h |
| tidb/20765   | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                           |   | 121d17h |
| tidb/20877   | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                   |   | 118d16h |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)              |   | 111d16h |
| tidb/21207   | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)             |   | 100d19h |
| tidb/21476   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                       |   | 90d15h  |
| tidb/21487   | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                               |   | 90d4h   |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876) |   | 75d19h  |
| tidb/21930   | [planner: propagate NDV of column groups across plan nodes (#17854)](https://github.com/pingcap/tidb/pull/21930)            |   | 71d18h  |
| tidb/21954   | [planner/cascades: add rule `PushSelDownApply`](https://github.com/pingcap/tidb/pull/21954)                                 |   | 70d23h  |
| tidb/22365   | [planner: check index valid while forUpdateRead (#22152)](https://github.com/pingcap/tidb/pull/22365)                       |   | 50d19h  |
| tidb/22504   | [*:Fix the fetchHotRegion bug that the count always zero](https://github.com/pingcap/tidb/pull/22504)                       |   | 38d19h  |
| tidb/22565   | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)          | Y | 35d17h  |
| tidb/22624   | [ planner: not pruning column used by union scan condition (#21640)](https://github.com/pingcap/tidb/pull/22624)            |   | 33d17h  |
| tidb/22804   | [session, util: update session to use new APIs (#22652)](https://github.com/pingcap/tidb/pull/22804)                        |   | 12d21h  |
| tidb/22830   | [planner: fix incorrect duration between compare](https://github.com/pingcap/tidb/pull/22830)                               |   | 12d9h   |
| tidb/22845   | [planner: fix bug of mpp wrongly set schema of exchanger](https://github.com/pingcap/tidb/pull/22845)                       |   | 11d17h  |
| tidb/22867   | [WIP: allow pushdown count distinct when enumerate physical plans](https://github.com/pingcap/tidb/pull/22867)              |   | 9d17h   |
| tidb/22923   | [expression: correct constant propagation for collation (#22666)](https://github.com/pingcap/tidb/pull/22923)               |   | 7d14h   |
| tidb/22971   | [Privileges: fix delete privilege check wrongly](https://github.com/pingcap/tidb/pull/22971)                                |   | 5d19h   |
| tidb/23088   | [statistics: delete extended stats cache item in current tidb synchronously](https://github.com/pingcap/tidb/pull/23088)    |   | 14h     |
| tidb/23092   | [*: fix a bug that collation is not handle for text type (#23045)](https://github.com/pingcap/tidb/pull/23092)              |   | 12h     |


## Reviewed in Last 7 Days

|     REPO     |                                                                PR                                                                 | C | D |   R    |
|--------------|-----------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/23087   | [executor: fix correlated column range in table reader for the clustered index table](https://github.com/pingcap/tidb/pull/23087) |   | 1 | 0h     |
| tidb/23078   | [store/copr: polish the tiflash-tikv fallback function.](https://github.com/pingcap/tidb/pull/23078)                              |   | 1 | 2h     |
| tidb/23045   | [*: fix a bug that collation is not handle for text type](https://github.com/pingcap/tidb/pull/23045)                             |   | 1 | 1d1h   |
| tidb/22940   | [planner: enable column pruning for common handle](https://github.com/pingcap/tidb/pull/22940)                                    |   | 1 | 5d22h  |
| tidb/22915   | [planner: build correct MaxOneRow info from multi-column conditions](https://github.com/pingcap/tidb/pull/22915)                  |   | 2 | 6d10h  |
| tidb/22914   | [partition: fix hash partition with not between condition get wrong result](https://github.com/pingcap/tidb/pull/22914)           |   | 2 | 6d10h  |
| tidb/22924   | [planner: fix wrong index merge selection (#22825)](https://github.com/pingcap/tidb/pull/22924)                                   |   | 2 | 6d6h   |
| tidb/23029   | [*: check `tidb_enable_extended_stats` in analyze and row count estimation](https://github.com/pingcap/tidb/pull/23029)           |   | 2 | 1d9h   |
| tidb/23041   | [util: remove unused code](https://github.com/pingcap/tidb/pull/23041)                                                            |   | 2 | 1d1h   |
| tidb/23052   | [*: support `show stats_extended` to inspect extended stats cache](https://github.com/pingcap/tidb/pull/23052)                    |   | 2 | 6h     |
| tidb/22090   | [planner: push aggregation operators down to projection and union by default](https://github.com/pingcap/tidb/pull/22090)         |   | 2 | 62d14h |
| tidb/22853   | [planner: fix LogicalPlans that contain Window Function are ambiguous ](https://github.com/pingcap/tidb/pull/22853)               |   | 2 | 9d20h  |
| tidb/22886   | [*: rename tidb_enable_tiflash_fallback_tikv to tidb_enable_engine_fallback](https://github.com/pingcap/tidb/pull/22886)          |   | 3 | 6d1h   |
| docs-cn/4669 | [sql-optimization: extended statistics documentation](https://github.com/pingcap/docs-cn/pull/4669)                               |   | 7 | 135d0h |


</details> 


<details> 
  <summary>wshwsh12's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                             PR                                                             | C | LASTED  |
|------------|----------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19557 | [*: Integrate timeline tracing with TiKV](https://github.com/pingcap/tidb/pull/19557)                                      |   | 187d23h |
| tidb/19807 | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                    |   | 180d10h |
| tidb/19957 | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                     | Y | 173d13h |
| tidb/21010 | [session: add a test case for issue 19127](https://github.com/pingcap/tidb/pull/21010)                                     |   | 111d19h |
| tidb/21487 | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                              |   | 90d4h   |
| tidb/21887 | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                            |   | 74d11h  |
| tidb/22176 | [store, executor: support stale read for tikv RPCContext](https://github.com/pingcap/tidb/pull/22176)                      |   | 57d19h  |
| tidb/22269 | [executor: check storage.block-cache.capacity value](https://github.com/pingcap/tidb/pull/22269)                           |   | 55d17h  |
| tidb/22378 | [executor: vectorize hash aggregate](https://github.com/pingcap/tidb/pull/22378)                                           |   | 49d19h  |
| tidb/22382 | [*: add infoschema client errors](https://github.com/pingcap/tidb/pull/22382)                                              |   | 49d5h   |
| tidb/22628 | [executor: Improve max/min window function with deque-based sliding window](https://github.com/pingcap/tidb/pull/22628)    |   | 32d23h  |
| tidb/22815 | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22815)        |   | 12d19h  |
| tidb/23041 | [util: remove unused code](https://github.com/pingcap/tidb/pull/23041)                                                     |   | 2d9h    |
| tidb/23064 | [sessionctx, executor: make the upper limit of tidb_ddl_reorg_worker_cnt work](https://github.com/pingcap/tidb/pull/23064) |   | 19h     |
| tidb/23079 | [executor: fix wrong key range of index scan when filter is comparing …](https://github.com/pingcap/tidb/pull/23079)       |   | 16h     |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 

