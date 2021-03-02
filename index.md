|   REVIEWER    |  O(C)   | 1 | 2 | 3 | 4 | 5 | 6 | 7 | T  |
|---------------|---------|---|---|---|---|---|---|---|----|
| Reminiscent   |  5 ( 0) | 3 | 0 | 0 | 1 | 0 | 1 | 0 |  5 |
| SunRunAway    | 19 ( 1) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| XuHuaiyu      | 57 ( 3) | 0 | 0 | 0 | 2 | 1 | 0 | 1 |  4 |
| dyzsr         |  0 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| eurekaka      | 14 ( 1) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| fzhedu        |  2 ( 0) | 2 | 0 | 0 | 0 | 0 | 1 | 0 |  3 |
| ichn-hu       |  3 ( 0) | 0 | 0 | 0 | 2 | 0 | 1 | 0 |  3 |
| lzmhhh123     | 34 ( 0) | 4 | 0 | 0 | 1 | 1 | 0 | 1 |  7 |
| nullnotnil    |  0 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| qw4990        | 50 ( 1) | 5 | 0 | 0 | 0 | 1 | 5 | 3 | 14 |
| rebelice      |  3 ( 0) | 0 | 0 | 0 | 3 | 1 | 1 | 1 |  6 |
| time-and-fate |  3 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| winoros       | 27 ( 2) | 2 | 0 | 0 | 0 | 1 | 1 | 0 |  4 |
| wshwsh12      | 19 ( 1) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |


<details> 
  <summary>Reminiscent's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                   PR                                                                   | C | LASTED |
|------------|----------------------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/21896 | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                             |   | 70d18h |
| tidb/22354 | [planner: do not cache prepared plan if optimization depends on mutable constant (#22349)](https://github.com/pingcap/tidb/pull/22354) |   | 48d22h |
| tidb/22918 | [sessionctx: add optimization-time and wait-TS-time into the slow log (#17869)](https://github.com/pingcap/tidb/pull/22918)            |   | 5d16h  |
| tidb/22938 | [planner: fix range partition prune bug for IN expr (#22894)](https://github.com/pingcap/tidb/pull/22938)                              |   | 4d18h  |
| tidb/23023 | [statistics: support dropping partition/global level statistics](https://github.com/pingcap/tidb/pull/23023)                           |   | 17h    |


## Reviewed in Last 7 Days

|    REPO     |                                                              PR                                                               | C | D |   R    |
|-------------|-------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| parser/1177 | [statistics: add a new syntax to support dropping statistics of partitions](https://github.com/pingcap/parser/pull/1177)      |   | 1 | 0h     |
| tidb/22625  | [planner, statistics: allow (auto) analyze single partition in dynamic-only mode](https://github.com/pingcap/tidb/pull/22625) |   | 1 | 30d20h |
| tidb/23016  | [statistics: add static partition prune mode checks](https://github.com/pingcap/tidb/pull/23016)                              |   | 1 | 2h     |
| tidb/22948  | [statistics: add more tests to check accurateness of global-stats](https://github.com/pingcap/tidb/pull/22948)                |   | 4 | 17h    |
| tidb/22911  | [statistics: make dump/load/show-stats be compatible with global-stats](https://github.com/pingcap/tidb/pull/22911)           |   | 6 | 1h     |


</details> 


<details> 
  <summary>SunRunAway's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                  PR                                                                   | C | LASTED  |
|------------|---------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19178 | [executor: Refactor probe channel](https://github.com/pingcap/tidb/pull/19178)                                                        |   | 200d16h |
| tidb/19347 | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347)  |   | 192d23h |
| tidb/19807 | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                               |   | 178d10h |
| tidb/19900 | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                        | Y | 173d18h |
| tidb/20140 | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                    |   | 160d22h |
| tidb/20220 | [*: new secondary index value format](https://github.com/pingcap/tidb/pull/20220)                                                     |   | 157d16h |
| tidb/20765 | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                     |   | 119d16h |
| tidb/21207 | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                       |   | 98d18h  |
| tidb/21277 | [executor: fix split table with large integers](https://github.com/pingcap/tidb/pull/21277)                                           |   | 96d19h  |
| tidb/21834 | [planner: enhanced index range calculation plan](https://github.com/pingcap/tidb/pull/21834)                                          |   | 75d18h  |
| tidb/21876 | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)           |   | 73d19h  |
| tidb/21878 | [planner: do not push down lock to pointGet/bacthPointGet when selection exists](https://github.com/pingcap/tidb/pull/21878)          |   | 73d17h  |
| tidb/21956 | [planner/preprocessor: disallow into-outfile clause in some place](https://github.com/pingcap/tidb/pull/21956)                        |   | 68d22h  |
| tidb/22026 | [expression: separated arithmeticPlusIntSig](https://github.com/pingcap/tidb/pull/22026)                                              |   | 66d20h  |
| tidb/22114 | [test: fix globalkilltest (#21987)](https://github.com/pingcap/tidb/pull/22114)                                                       |   | 61d12h  |
| tidb/22181 | [planner, expression: fix error when using IN combined with subquery (#22080)](https://github.com/pingcap/tidb/pull/22181)            |   | 55d17h  |
| tidb/22217 | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                         |   | 54d17h  |
| tidb/22365 | [planner: check index valid while forUpdateRead (#22152)](https://github.com/pingcap/tidb/pull/22365)                                 |   | 48d18h  |
| tidb/22379 | [[experiment] executor: allow aggregation to spill disk when running out of memory quota](https://github.com/pingcap/tidb/pull/22379) |   | 47d19h  |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>XuHuaiyu's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                              PR                                                                              | C | LASTED  |
|--------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| docs-cn/5619 | [Update data-type-date-and-time.md](https://github.com/pingcap/docs-cn/pull/5619)                                                                            |   | 3d15h   |
| tidb/19900   | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                                               | Y | 173d18h |
| docs-cn/5620 | [Add details for Hexadecimal Literals](https://github.com/pingcap/docs-cn/pull/5620)                                                                         |   | 3d15h   |
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                                       | Y | 171d13h |
| tidb/20040   | [planner, expression: take NullFlag into consideration when optimize the `int non-const` <cmp > `non-int const`](https://github.com/pingcap/tidb/pull/20040) | Y | 166d13h |
| tidb/20140   | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                                           |   | 160d22h |
| tidb/20311   | [expression: fix overflow error when convert bit to int64 (#20266)](https://github.com/pingcap/tidb/pull/20311)                                              |   | 152d21h |
| tidb/20790   | [collation: add pinyin collation for chinese charset support](https://github.com/pingcap/tidb/pull/20790)                                                    |   | 118d20h |
| tidb/20905   | [planner: fix statement-optimize not work in `TryFastPlan`](https://github.com/pingcap/tidb/pull/20905)                                                      |   | 115d16h |
| tidb/20972   | [expression: POC implementation of Vitess hashing algorithm.](https://github.com/pingcap/tidb/pull/20972)                                                    |   | 111d0h  |
| tidb/21064   | [planner, executor: fix cast not check error](https://github.com/pingcap/tidb/pull/21064)                                                                    |   | 106d8h  |
| tidb/21149   | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                                |   | 102d14h |
| tidb/21228   | [executor: return the result immediately when combining LIMIT row_count with DISTINCT](https://github.com/pingcap/tidb/pull/21228)                           |   | 98d13h  |
| tidb/21304   | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                                     |   | 96d12h  |
| tidb/21334   | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                                                |   | 95d13h  |
| tidb/21340   | [executor: initialize expensive query handler on domain creation](https://github.com/pingcap/tidb/pull/21340)                                                |   | 94d23h  |
| tidb/21476   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                        |   | 88d15h  |
| tidb/21536   | [executor: add slow-log file meta cache to avoid repeat read file meta information](https://github.com/pingcap/tidb/pull/21536)                              |   | 84d14h  |
| tidb/21564   | [ddl: fix Incorrect behavior of NO_ZERO_DATE when altering table](https://github.com/pingcap/tidb/pull/21564)                                                |   | 83d15h  |
| tidb/21853   | [expression: fix compatibility behaviors in time_format with MySQL (#21559)](https://github.com/pingcap/tidb/pull/21853)                                     |   | 74d18h  |
| tidb/21896   | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                                                   |   | 70d18h  |
| tidb/22131   | [privilege: remove leading and trailing space when create user and role](https://github.com/pingcap/tidb/pull/22131)                                         |   | 60d19h  |
| tidb/22149   | [session: set process info before building plan (#22101)](https://github.com/pingcap/tidb/pull/22149)                                                        |   | 56d19h  |
| tidb/22163   | [expression: separated arithmeticMinusIntSig](https://github.com/pingcap/tidb/pull/22163)                                                                    |   | 56d13h  |
| tidb/22186   | [executor: fix select into outfile with year type column has no data (#22175)](https://github.com/pingcap/tidb/pull/22186)                                   |   | 55d16h  |
| tidb/22294   | [planner, table: optimize the list partition pruner for range query](https://github.com/pingcap/tidb/pull/22294)                                             |   | 52d19h  |
| tidb/22307   | [ddl: fix update can see columns not public](https://github.com/pingcap/tidb/pull/22307)                                                                     |   | 52d15h  |
| tidb/22381   | [planner: check schema stale for plan cache when forUpdateRead](https://github.com/pingcap/tidb/pull/22381)                                                  |   | 47d14h  |
| tidb/22616   | [expression: from_unixtime accept 64-bit integers](https://github.com/pingcap/tidb/pull/22616)                                                               |   | 31d22h  |
| tidb/22617   | [metrics: fix wrong bucket name of coprocessor cache (#22454)](https://github.com/pingcap/tidb/pull/22617)                                                   |   | 31d22h  |
| tidb/22624   | [ planner: not pruning column used by union scan condition (#21640)](https://github.com/pingcap/tidb/pull/22624)                                             |   | 31d16h  |
| tidb/22631   | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                                                                              |   | 29d22h  |
| tidb/22640   | [*: refactor ExecuteInternal to return single resultset (#22546)](https://github.com/pingcap/tidb/pull/22640)                                                |   | 28d19h  |
| tidb/22696   | [expression: enable arithmetic Mod push down](https://github.com/pingcap/tidb/pull/22696)                                                                    |   | 26d16h  |
| tidb/22711   | [executor: Fix inline schema name](https://github.com/pingcap/tidb/pull/22711)                                                                               |   | 26d11h  |
| tidb/22722   | [planner, errno: make error code of ErrMixOfGroupFuncAndFields consistent with MySQL](https://github.com/pingcap/tidb/pull/22722)                            |   | 25d20h  |
| tidb/22736   | [executor: fix load data losing connection when batch_dml_size is set (#22724)](https://github.com/pingcap/tidb/pull/22736)                                  |   | 24d23h  |
| tidb/22786   | [config: deprecate `tikv-client.copr-cache.enable`](https://github.com/pingcap/tidb/pull/22786)                                                              |   | 11d18h  |
| tidb/22814   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22814)                                          |   | 10d18h  |
| tidb/22815   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22815)                                          |   | 10d18h  |
| tidb/22832   | [expression: push down EXTRACT to TiFlash](https://github.com/pingcap/tidb/pull/22832)                                                                       |   | 10d1h   |
| tidb/22844   | [expression: do not adjust int when it is null and compared year (#22821)](https://github.com/pingcap/tidb/pull/22844)                                       |   | 9d19h   |
| tidb/22886   | [*: rename tidb_enable_tiflash_fallback_tikv to tidb_enable_engine_fallback](https://github.com/pingcap/tidb/pull/22886)                                     |   | 6d19h   |
| tidb/22903   | [docs/design: update template](https://github.com/pingcap/tidb/pull/22903)                                                                                   |   | 6d12h   |
| tidb/22914   | [partition: fix hash partition with not between condition get wrong result](https://github.com/pingcap/tidb/pull/22914)                                      |   | 5d17h   |
| tidb/22961   | [statistics: refactor the statistics package use the RestrictedSQLExecutor API (#22636)](https://github.com/pingcap/tidb/pull/22961)                         |   | 3d23h   |
| tidb/22962   | [executor: track partialResultMap in unparalleled aggreagte.](https://github.com/pingcap/tidb/pull/22962)                                                    |   | 3d23h   |
| tidb/22965   | [executor: introduce setWithMemoryUsage to track set memory in AggExec.](https://github.com/pingcap/tidb/pull/22965)                                         |   | 3d22h   |
| tidb/22971   | [Privileges: fix delete privilege check wrongly](https://github.com/pingcap/tidb/pull/22971)                                                                 |   | 3d19h   |
| tidb/22976   | [kv/union_store:remove tableinfo from union_store](https://github.com/pingcap/tidb/pull/22976)                                                               |   | 3d16h   |
| tidb/22983   | [*: fix some structcheck lint warnings](https://github.com/pingcap/tidb/pull/22983)                                                                          |   | 2d12h   |
| tidb/22989   | [distsql: fix err check](https://github.com/pingcap/tidb/pull/22989)                                                                                         |   | 2d0h    |
| tidb/22992   | [expression: fix err check](https://github.com/pingcap/tidb/pull/22992)                                                                                      |   | 2d0h    |
| tidb/23012   | [executor: fix affected rows of creating database and dropping database](https://github.com/pingcap/tidb/pull/23012)                                         |   | 1d16h   |
| tidb/23024   | [executor: make the memory tracker of Jsonobjectagg more accurate](https://github.com/pingcap/tidb/pull/23024)                                               |   | 17h     |
| tidb/23026   | [ddl: fix the create view clause will be restored with charset prefix](https://github.com/pingcap/tidb/pull/23026)                                           |   | 17h     |
| tidb/23034   | [executor: make the memory tracker of groupConcat more accurate.](https://github.com/pingcap/tidb/pull/23034)                                                |   | 16h     |


## Reviewed in Last 7 Days

|    REPO    |                                                       PR                                                        | C | D |   R    |
|------------|-----------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/22507 | [types: fix the bug about the wrong query result for decimal type ](https://github.com/pingcap/tidb/pull/22507) |   | 4 | 32d3h  |
| tidb/22426 | [expression: fix bugs in builtinfunction ArithmeticMinusInt logic](https://github.com/pingcap/tidb/pull/22426)  |   | 4 | 38d17h |
| tidb/22912 | [executor: reduce useless memory allocation in min/max aggregate](https://github.com/pingcap/tidb/pull/22912)   |   | 5 | 20h    |
| tidb/22883 | [executor: fix regression by Tracker.Consume in aggregate.](https://github.com/pingcap/tidb/pull/22883)         |   | 7 | 0h     |


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
| tidb/19347 | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347)   |   | 192d23h |
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                              |   | 116d16h |
| tidb/21444 | [planner: ignore anonymous index while tiflash replica is available](https://github.com/pingcap/tidb/pull/21444)                       |   | 89d12h  |
| tidb/21994 | [range: fix overflow value access index ](https://github.com/pingcap/tidb/pull/21994)                                                  |   | 67d22h  |
| tidb/22342 | [session: fix two cases when updating bind info (#22338)](https://github.com/pingcap/tidb/pull/22342)                                  |   | 49d18h  |
| tidb/22354 | [planner: do not cache prepared plan if optimization depends on mutable constant (#22349)](https://github.com/pingcap/tidb/pull/22354) |   | 48d22h  |
| tidb/22369 | [session: fix the duplicate binding case when updating bind info (#22367)](https://github.com/pingcap/tidb/pull/22369)                 |   | 48d17h  |
| tidb/22416 | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                                      | Y | 44d11h  |
| tidb/22559 | [planner: split test data from test cases in cbo_test.go](https://github.com/pingcap/tidb/pull/22559)                                  |   | 33d19h  |
| tidb/22733 | [bindinfo: use new sql apis (#22653)](https://github.com/pingcap/tidb/pull/22733)                                                      |   | 25d15h  |
| tidb/22778 | [*: add support for dynamic privileges](https://github.com/pingcap/tidb/pull/22778)                                                    |   | 13d7h   |
| tidb/22953 | [planner: fix query range partition table got wrong result and TiDB panic](https://github.com/pingcap/tidb/pull/22953)                 |   | 4d14h   |
| tidb/22985 | [bindinfo: fix error check](https://github.com/pingcap/tidb/pull/22985)                                                                |   | 2d0h    |
| tidb/23041 | [util: remove unused code](https://github.com/pingcap/tidb/pull/23041)                                                                 |   | 9h      |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>fzhedu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                  PR                                                   | C | LASTED |
|------------|-------------------------------------------------------------------------------------------------------|---|--------|
| tidb/22662 | [planner/core: let mpp support partition tables](https://github.com/pingcap/tidb/pull/22662)          |   | 27d18h |
| tidb/22845 | [planner: fix bug of mpp wrongly set schema of exchanger](https://github.com/pingcap/tidb/pull/22845) |   | 9d17h  |


## Reviewed in Last 7 Days

|    REPO    |                                                                  PR                                                                   | C | D |   R   |
|------------|---------------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tics/1342  | [Add background task to detect && cancel hanging MPP query](https://github.com/pingcap/tics/pull/1342)                                |   | 1 | 48d2h |
| tidb/23020 | [expression: Add warning info for exprs that can not be pushed to storage layer (#22713)](https://github.com/pingcap/tidb/pull/23020) |   | 1 | 0h    |
| tidb/22803 | [store/mockstore/unistore: refine and add more mpp tests](https://github.com/pingcap/tidb/pull/22803)                                 |   | 6 | 5d9h  |


</details> 


<details> 
  <summary>ichn-hu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                            PR                                                            | C | LASTED |
|------------|--------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/21853 | [expression: fix compatibility behaviors in time_format with MySQL (#21559)](https://github.com/pingcap/tidb/pull/21853) |   | 74d18h |
| tidb/23024 | [executor: make the memory tracker of Jsonobjectagg more accurate](https://github.com/pingcap/tidb/pull/23024)           |   | 17h    |
| tidb/23034 | [executor: make the memory tracker of groupConcat more accurate.](https://github.com/pingcap/tidb/pull/23034)            |   | 16h    |


## Reviewed in Last 7 Days

|    REPO    |                                                          PR                                                          | C | D | R  |
|------------|----------------------------------------------------------------------------------------------------------------------|---|---|----|
| tidb/22965 | [executor: introduce setWithMemoryUsage to track set memory in AggExec.](https://github.com/pingcap/tidb/pull/22965) |   | 4 | 5h |
| tidb/22962 | [executor: track partialResultMap in unparalleled aggreagte.](https://github.com/pingcap/tidb/pull/22962)            |   | 4 | 5h |
| tidb/22912 | [executor: reduce useless memory allocation in min/max aggregate](https://github.com/pingcap/tidb/pull/22912)        |   | 6 | 3h |


</details> 


<details> 
  <summary>lzmhhh123's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                             PR                                                                              | C | LASTED  |
|------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19347 | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347)                        |   | 192d23h |
| tidb/20444 | [expression: add json_merge_patch](https://github.com/pingcap/tidb/pull/20444)                                                                              |   | 138d21h |
| tidb/20465 | [expression: add uuidShortFunction](https://github.com/pingcap/tidb/pull/20465)                                                                             |   | 137d19h |
| tidb/20642 | [executor: modify admin executors to support partitioned table with global index](https://github.com/pingcap/tidb/pull/20642)                               |   | 126d15h |
| tidb/20903 | [planner: fix confused and unnecessary double-projection in plans.](https://github.com/pingcap/tidb/pull/20903)                                             |   | 115d17h |
| tidb/21018 | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                              |   | 109d16h |
| tidb/21195 | [brie: integrate lightning to suport IMPORT statement](https://github.com/pingcap/tidb/pull/21195)                                                          |   | 98d22h  |
| tidb/21334 | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                                               |   | 95d13h  |
| tidb/21347 | [session: make rollback work on global variables](https://github.com/pingcap/tidb/pull/21347)                                                               |   | 94d19h  |
| tidb/21401 | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                                          |   | 91d11h  |
| tidb/21444 | [planner: ignore anonymous index while tiflash replica is available](https://github.com/pingcap/tidb/pull/21444)                                            |   | 89d12h  |
| tidb/21487 | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                                               |   | 88d4h   |
| tidb/21641 | [executor: Fix pessimistic lock doesn't work on the partition table for subquery/joins](https://github.com/pingcap/tidb/pull/21641)                         |   | 81d17h  |
| tidb/21651 | [planner: allow filter condition pushing down to IndexScan for prefix index](https://github.com/pingcap/tidb/pull/21651)                                    |   | 81d13h  |
| tidb/22126 | [*: add `sys` schema, `sys.SCHEMA_UNUSED_INDEXES` view and `sys.SCHEMA_INDEX_USAGE` view](https://github.com/pingcap/tidb/pull/22126)                       |   | 60d19h  |
| tidb/22149 | [session: set process info before building plan (#22101)](https://github.com/pingcap/tidb/pull/22149)                                                       |   | 56d19h  |
| tidb/22188 | [planner: do not use indexMerge when the path only use a single index (#22168)](https://github.com/pingcap/tidb/pull/22188)                                 |   | 55d13h  |
| tidb/22361 | [table: fix insert into _tidb_rowid panic and rebase it if needed (#22062)](https://github.com/pingcap/tidb/pull/22361)                                     |   | 48d19h  |
| tidb/22372 | [executor: fix SelectForUpdate in decorrelated subquery under pessimistic mode](https://github.com/pingcap/tidb/pull/22372)                                 |   | 48d9h   |
| tidb/22430 | [*: refactor table.Table interface, clean up unnecessay methods](https://github.com/pingcap/tidb/pull/22430)                                                |   | 41d23h  |
| tidb/22478 | [planner, executor: fix query partition table with global unique index get wrong result](https://github.com/pingcap/tidb/pull/22478)                        |   | 39d13h  |
| tidb/22631 | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                                                                             |   | 29d22h  |
| tidb/22656 | [*: move new api out of session package (#22591)](https://github.com/pingcap/tidb/pull/22656)                                                               |   | 28d0h   |
| tidb/22662 | [planner/core: let mpp support partition tables](https://github.com/pingcap/tidb/pull/22662)                                                                |   | 27d18h  |
| tidb/22699 | [brie: add error info column and history backup/restore info in sql](https://github.com/pingcap/tidb/pull/22699)                                            |   | 26d16h  |
| tidb/22714 | [executor: add close recordSet in executor](https://github.com/pingcap/tidb/pull/22714)                                                                     |   | 25d23h  |
| tidb/22812 | [ executor: add new format specifier(%# %@ %.) for str_to_date expression (#22790)](https://github.com/pingcap/tidb/pull/22812)                             |   | 10d19h  |
| tidb/22834 | [executor, server: load_data.go is changed and add unit test](https://github.com/pingcap/tidb/pull/22834)                                                   |   | 10d0h   |
| tidb/22857 | [store: move mockstore/mocktikv to tikv/mockstore/mocktikv](https://github.com/pingcap/tidb/pull/22857)                                                     |   | 8d20h   |
| tidb/22910 | [util: optimize the performance of restore with db](https://github.com/pingcap/tidb/pull/22910)                                                             |   | 5d19h   |
| tidb/22961 | [statistics: refactor the statistics package use the RestrictedSQLExecutor API (#22636)](https://github.com/pingcap/tidb/pull/22961)                        |   | 3d23h   |
| tidb/22974 | [parser: update parser to fix sql restore bug used in create view](https://github.com/pingcap/tidb/pull/22974)                                              |   | 3d17h   |
| tidb/23001 | [statistics: fix err check](https://github.com/pingcap/tidb/pull/23001)                                                                                     |   | 2d0h    |
| tidb/23022 | [executor: create PipelinedWindowExec based on current implementation and modify the windowProcessor interface](https://github.com/pingcap/tidb/pull/23022) |   | 18h     |


## Reviewed in Last 7 Days

|     REPO     |                                                                PR                                                                | C | D |   R   |
|--------------|----------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/22980   | [planner: choose non-prefix column when both index key and handle have the same one](https://github.com/pingcap/tidb/pull/22980) |   | 1 | 2d18h |
| docs-cn/5561 | [Add sql optimization-related docs to toc](https://github.com/pingcap/docs-cn/pull/5561)                                         |   | 1 | 6d20h |
| tidb/22991   | [executor: fix err check](https://github.com/pingcap/tidb/pull/22991)                                                            |   | 1 | 1d5h  |
| tidb/22940   | [planner: enable column pruning for common handle](https://github.com/pingcap/tidb/pull/22940)                                   |   | 1 | 3d22h |
| tidb/22959   | [planner, executor: reset NotNullFlag when merge schema for join (#22955)](https://github.com/pingcap/tidb/pull/22959)           |   | 4 | 11h   |
| tidb/22955   | [planner, executor: reset NotNullFlag when merge schema for join](https://github.com/pingcap/tidb/pull/22955)                    |   | 5 | 0h    |
| tidb/22883   | [executor: fix regression by Tracker.Consume in aggregate.](https://github.com/pingcap/tidb/pull/22883)                          |   | 7 | 3h    |


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
| docs-cn/5484 | [system variable: add tidb_enable_engine_fallback](https://github.com/pingcap/docs-cn/pull/5484)                                                            |   | 26d17h  |
| parser/1178  | [parser: support `show stats_extended` syntax](https://github.com/pingcap/parser/pull/1178)                                                                 |   | 15h     |
| tidb/19029   | [types: fix unexpected NOT_NULL flags](https://github.com/pingcap/tidb/pull/19029)                                                                          |   | 207d22h |
| docs-cn/5561 | [Add sql optimization-related docs to toc](https://github.com/pingcap/docs-cn/pull/5561)                                                                    |   | 7d15h   |
| tidb/20708   | [*: separate auto_increment ID allocator from _tidb_rowid allocator](https://github.com/pingcap/tidb/pull/20708)                                            |   | 123d20h |
| tidb/20969   | [executor: Improve the performance of appending not fixed columns](https://github.com/pingcap/tidb/pull/20969)                                              |   | 111d9h  |
| tidb/20972   | [expression: POC implementation of Vitess hashing algorithm.](https://github.com/pingcap/tidb/pull/20972)                                                   |   | 111d0h  |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                              |   | 109d16h |
| tidb/21149   | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                               |   | 102d14h |
| tidb/21304   | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                                    |   | 96d12h  |
| tidb/21318   | [planner, expression: use the range of column types to simplify expressions](https://github.com/pingcap/tidb/pull/21318)                                    |   | 95d18h  |
| tidb/21401   | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                                          |   | 91d11h  |
| tidb/21476   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                       |   | 88d15h  |
| tidb/21508   | [execution: fix dayofweek('0000-00-00') behavior](https://github.com/pingcap/tidb/pull/21508)                                                               |   | 87d9h   |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                                 |   | 73d19h  |
| tidb/21887   | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                                             |   | 72d11h  |
| tidb/21930   | [planner: propagate NDV of column groups across plan nodes (#17854)](https://github.com/pingcap/tidb/pull/21930)                                            |   | 69d18h  |
| tidb/21954   | [planner/cascades: add rule `PushSelDownApply`](https://github.com/pingcap/tidb/pull/21954)                                                                 |   | 68d23h  |
| tidb/22090   | [planner: push aggregation operators down to projection and union by default](https://github.com/pingcap/tidb/pull/22090)                                   |   | 61d22h  |
| tidb/22146   | [executor: forbid SFU on view](https://github.com/pingcap/tidb/pull/22146)                                                                                  |   | 56d21h  |
| tidb/22217   | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                                               |   | 54d17h  |
| tidb/22234   | [executor, planner: ON DUPLICATE UPDATE can refer to un-project col (#14412)](https://github.com/pingcap/tidb/pull/22234)                                   |   | 54d14h  |
| tidb/22261   | [time: fix parse datetime won't truncate the reluctant string (#22232)](https://github.com/pingcap/tidb/pull/22261)                                         |   | 53d19h  |
| tidb/22307   | [ddl: fix update can see columns not public](https://github.com/pingcap/tidb/pull/22307)                                                                    |   | 52d15h  |
| tidb/22342   | [session: fix two cases when updating bind info (#22338)](https://github.com/pingcap/tidb/pull/22342)                                                       |   | 49d18h  |
| tidb/22369   | [session: fix the duplicate binding case when updating bind info (#22367)](https://github.com/pingcap/tidb/pull/22369)                                      |   | 48d17h  |
| tidb/22374   | [expression: separated arithmeticIntDivideSig](https://github.com/pingcap/tidb/pull/22374)                                                                  |   | 48d0h   |
| tidb/22415   | [ddl: refactor placement package](https://github.com/pingcap/tidb/pull/22415)                                                                               |   | 44d17h  |
| tidb/22471   | [ddl, executor: fix creating unique index without partition column error when enable-global-index is true](https://github.com/pingcap/tidb/pull/22471)      |   | 39d17h  |
| tidb/22507   | [types: fix the bug about the wrong query result for decimal type ](https://github.com/pingcap/tidb/pull/22507)                                             |   | 35d23h  |
| tidb/22541   | [expression: Support builtin function SOUNDEX](https://github.com/pingcap/tidb/pull/22541)                                                                  |   | 34d9h   |
| tidb/22559   | [planner: split test data from test cases in cbo_test.go](https://github.com/pingcap/tidb/pull/22559)                                                       |   | 33d19h  |
| tidb/22565   | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)                                          | Y | 33d17h  |
| tidb/22733   | [bindinfo: use new sql apis (#22653)](https://github.com/pingcap/tidb/pull/22733)                                                                           |   | 25d15h  |
| tidb/22760   | [expression: fix wrong error info](https://github.com/pingcap/tidb/pull/22760)                                                                              |   | 22d0h   |
| tidb/22778   | [*: add support for dynamic privileges](https://github.com/pingcap/tidb/pull/22778)                                                                         |   | 13d7h   |
| tidb/22814   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22814)                                         |   | 10d18h  |
| tidb/22815   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22815)                                         |   | 10d18h  |
| tidb/22862   | [brie: fix the problem that ddl restored by BR via SQL is not replicated to downstream](https://github.com/pingcap/tidb/pull/22862)                         |   | 7d22h   |
| tidb/22869   | [executor: fix cast function will ignore tht error for point-get key construction](https://github.com/pingcap/tidb/pull/22869)                              |   | 7d16h   |
| tidb/22886   | [*: rename tidb_enable_tiflash_fallback_tikv to tidb_enable_engine_fallback](https://github.com/pingcap/tidb/pull/22886)                                    |   | 6d19h   |
| tidb/22915   | [planner: build correct MaxOneRow info from multi-column conditions](https://github.com/pingcap/tidb/pull/22915)                                            |   | 5d17h   |
| tidb/22923   | [expression: correct constant propagation for collation (#22666)](https://github.com/pingcap/tidb/pull/22923)                                               |   | 5d14h   |
| tidb/22924   | [planner: fix wrong index merge selection (#22825)](https://github.com/pingcap/tidb/pull/22924)                                                             |   | 5d13h   |
| tidb/22926   | [expression: add overflow check in multiplyInt](https://github.com/pingcap/tidb/pull/22926)                                                                 |   | 5d13h   |
| tidb/22938   | [planner: fix range partition prune bug for IN expr (#22894)](https://github.com/pingcap/tidb/pull/22938)                                                   |   | 4d18h   |
| tidb/22984   | [executor: fix logging format of prepared statements (#16062)](https://github.com/pingcap/tidb/pull/22984)                                                  |   | 2d9h    |
| tidb/23008   | [types: fix err check](https://github.com/pingcap/tidb/pull/23008)                                                                                          |   | 1d23h   |
| tidb/23022   | [executor: create PipelinedWindowExec based on current implementation and modify the windowProcessor interface](https://github.com/pingcap/tidb/pull/23022) |   | 18h     |
| tidb/23029   | [*: check `tidb_enable_extended_stats` in analyze and row count estimation](https://github.com/pingcap/tidb/pull/23029)                                     |   | 17h     |


## Reviewed in Last 7 Days

|      REPO      |                                                                  PR                                                                  | C | D |   R    |
|----------------|--------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/22625     | [planner, statistics: allow (auto) analyze single partition in dynamic-only mode](https://github.com/pingcap/tidb/pull/22625)        |   | 1 | 30d20h |
| tidb/22931     | [statistics: enables global-level stats to be generated in fast analyze mode](https://github.com/pingcap/tidb/pull/22931)            |   | 1 | 4d5h   |
| tidb/22489     | [infoschema: support query partition_id from infoschema.partitions (#22240)](https://github.com/pingcap/tidb/pull/22489)             |   | 1 | 38d0h  |
| tidb/23016     | [statistics: add static partition prune mode checks](https://github.com/pingcap/tidb/pull/23016)                                     |   | 1 | 2h     |
| tidb-test/1163 | [*: rename partition_prune_mode](https://github.com/pingcap/tidb-test/pull/1163)                                                     |   | 1 | 6d0h   |
| tidb/22910     | [util: optimize the performance of restore with db](https://github.com/pingcap/tidb/pull/22910)                                      |   | 5 | 1d0h   |
| tidb/22825     | [planner: fix wrong index merge selection](https://github.com/pingcap/tidb/pull/22825)                                               |   | 6 | 5d0h   |
| tidb/22803     | [store/mockstore/unistore: refine and add more mpp tests](https://github.com/pingcap/tidb/pull/22803)                                |   | 6 | 5d6h   |
| tidb/22866     | [*: modify the switch to control global stats](https://github.com/pingcap/tidb/pull/22866)                                           |   | 6 | 2d1h   |
| tidb/22666     | [expression: correct constant propagation for collation](https://github.com/pingcap/tidb/pull/22666)                                 |   | 6 | 22d0h  |
| tidb/22649     | [planner: decorrelate LogicalApply with inner join as the inner child](https://github.com/pingcap/tidb/pull/22649)                   |   | 6 | 22d21h |
| tidb/22682     | [statistics: introduce new estimation logic when index histogram fails to estimate](https://github.com/pingcap/tidb/pull/22682)      |   | 7 | 20d19h |
| tidb/22878     | [statistics: merge partition-level FMSketch to global-level FMSketch and update the NDV](https://github.com/pingcap/tidb/pull/22878) |   | 7 | 7h     |
| tidb/22836     | [expression: add integration test for partition pruning](https://github.com/pingcap/tidb/pull/22836)                                 |   | 7 | 3d3h   |


</details> 


<details> 
  <summary>rebelice's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                               PR                                                               | C | LASTED |
|------------|--------------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/22869 | [executor: fix cast function will ignore tht error for point-get key construction](https://github.com/pingcap/tidb/pull/22869) |   | 7d16h  |
| tidb/22938 | [planner: fix range partition prune bug for IN expr (#22894)](https://github.com/pingcap/tidb/pull/22938)                      |   | 4d18h  |
| tidb/23023 | [statistics: support dropping partition/global level statistics](https://github.com/pingcap/tidb/pull/23023)                   |   | 17h    |


## Reviewed in Last 7 Days

|    REPO    |                                                                  PR                                                                  | C | D |  R   |
|------------|--------------------------------------------------------------------------------------------------------------------------------------|---|---|------|
| tidb/22931 | [statistics: enables global-level stats to be generated in fast analyze mode](https://github.com/pingcap/tidb/pull/22931)            |   | 4 | 1d4h |
| tidb/22948 | [statistics: add more tests to check accurateness of global-stats](https://github.com/pingcap/tidb/pull/22948)                       |   | 4 | 20h  |
| tidb/22959 | [planner, executor: reset NotNullFlag when merge schema for join (#22955)](https://github.com/pingcap/tidb/pull/22959)               |   | 4 | 11h  |
| tidb/22955 | [planner, executor: reset NotNullFlag when merge schema for join](https://github.com/pingcap/tidb/pull/22955)                        |   | 5 | 1h   |
| tidb/22911 | [statistics: make dump/load/show-stats be compatible with global-stats](https://github.com/pingcap/tidb/pull/22911)                  |   | 6 | 2h   |
| tidb/22878 | [statistics: merge partition-level FMSketch to global-level FMSketch and update the NDV](https://github.com/pingcap/tidb/pull/22878) |   | 7 | 7h   |


</details> 


<details> 
  <summary>time-and-fate's detailed review report</summary> 

## To Be Reviewed

|    REPO     |                                                           PR                                                            | C | LASTED  |
|-------------|-------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20877  | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                               |   | 116d16h |
| parser/1178 | [parser: support `show stats_extended` syntax](https://github.com/pingcap/parser/pull/1178)                             |   | 15h     |
| tidb/23029  | [*: check `tidb_enable_extended_stats` in analyze and row count estimation](https://github.com/pingcap/tidb/pull/23029) |   | 17h     |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>winoros's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                             PR                                                              | C | LASTED  |
|--------------|-----------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                      | Y | 171d13h |
| docs-cn/5484 | [system variable: add tidb_enable_engine_fallback](https://github.com/pingcap/docs-cn/pull/5484)                            |   | 26d17h  |
| tidb/20311   | [expression: fix overflow error when convert bit to int64 (#20266)](https://github.com/pingcap/tidb/pull/20311)             |   | 152d21h |
| tidb/20765   | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                           |   | 119d16h |
| tidb/20877   | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                   |   | 116d16h |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)              |   | 109d16h |
| tidb/21207   | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)             |   | 98d18h  |
| tidb/21476   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                       |   | 88d15h  |
| tidb/21487   | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                               |   | 88d4h   |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876) |   | 73d19h  |
| tidb/21930   | [planner: propagate NDV of column groups across plan nodes (#17854)](https://github.com/pingcap/tidb/pull/21930)            |   | 69d18h  |
| tidb/21954   | [planner/cascades: add rule `PushSelDownApply`](https://github.com/pingcap/tidb/pull/21954)                                 |   | 68d23h  |
| tidb/22090   | [planner: push aggregation operators down to projection and union by default](https://github.com/pingcap/tidb/pull/22090)   |   | 61d22h  |
| tidb/22365   | [planner: check index valid while forUpdateRead (#22152)](https://github.com/pingcap/tidb/pull/22365)                       |   | 48d18h  |
| tidb/22504   | [*:Fix the fetchHotRegion bug that the count always zero](https://github.com/pingcap/tidb/pull/22504)                       |   | 36d19h  |
| tidb/22565   | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)          | Y | 33d17h  |
| tidb/22624   | [ planner: not pruning column used by union scan condition (#21640)](https://github.com/pingcap/tidb/pull/22624)            |   | 31d16h  |
| tidb/22804   | [session, util: update session to use new APIs (#22652)](https://github.com/pingcap/tidb/pull/22804)                        |   | 10d20h  |
| tidb/22830   | [planner: fix incorrect duration between compare](https://github.com/pingcap/tidb/pull/22830)                               |   | 10d9h   |
| tidb/22845   | [planner: fix bug of mpp wrongly set schema of exchanger](https://github.com/pingcap/tidb/pull/22845)                       |   | 9d17h   |
| tidb/22867   | [WIP: allow pushdown count distinct when enumerate physical plans](https://github.com/pingcap/tidb/pull/22867)              |   | 7d17h   |
| tidb/22914   | [partition: fix hash partition with not between condition get wrong result](https://github.com/pingcap/tidb/pull/22914)     |   | 5d17h   |
| tidb/22915   | [planner: build correct MaxOneRow info from multi-column conditions](https://github.com/pingcap/tidb/pull/22915)            |   | 5d17h   |
| tidb/22923   | [expression: correct constant propagation for collation (#22666)](https://github.com/pingcap/tidb/pull/22923)               |   | 5d14h   |
| tidb/22924   | [planner: fix wrong index merge selection (#22825)](https://github.com/pingcap/tidb/pull/22924)                             |   | 5d13h   |
| tidb/22971   | [Privileges: fix delete privilege check wrongly](https://github.com/pingcap/tidb/pull/22971)                                |   | 3d19h   |
| tidb/23016   | [statistics: add static partition prune mode checks](https://github.com/pingcap/tidb/pull/23016)                            |   | 22h     |


## Reviewed in Last 7 Days

|     REPO     |                                                            PR                                                            | C | D |   R    |
|--------------|--------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/22940   | [planner: enable column pruning for common handle](https://github.com/pingcap/tidb/pull/22940)                           |   | 1 | 3d23h  |
| tidb/22886   | [*: rename tidb_enable_tiflash_fallback_tikv to tidb_enable_engine_fallback](https://github.com/pingcap/tidb/pull/22886) |   | 1 | 6d1h   |
| docs-cn/4669 | [sql-optimization: extended statistics documentation](https://github.com/pingcap/docs-cn/pull/4669)                      |   | 5 | 135d0h |
| tidb/22666   | [expression: correct constant propagation for collation](https://github.com/pingcap/tidb/pull/22666)                     |   | 6 | 22d0h  |


</details> 


<details> 
  <summary>wshwsh12's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                  PR                                                                  | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19557 | [*: Integrate timeline tracing with TiKV](https://github.com/pingcap/tidb/pull/19557)                                                |   | 185d23h |
| tidb/19807 | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                              |   | 178d10h |
| tidb/19957 | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                               | Y | 171d13h |
| tidb/21010 | [session: add a test case for issue 19127](https://github.com/pingcap/tidb/pull/21010)                                               |   | 109d19h |
| tidb/21487 | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                        |   | 88d4h   |
| tidb/21887 | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                      |   | 72d11h  |
| tidb/22176 | [store, executor: support stale read for tikv RPCContext](https://github.com/pingcap/tidb/pull/22176)                                |   | 55d18h  |
| tidb/22269 | [executor: check storage.block-cache.capacity value](https://github.com/pingcap/tidb/pull/22269)                                     |   | 53d16h  |
| tidb/22378 | [executor: vectorize hash aggregate](https://github.com/pingcap/tidb/pull/22378)                                                     |   | 47d19h  |
| tidb/22382 | [*: add infoschema client errors](https://github.com/pingcap/tidb/pull/22382)                                                        |   | 47d5h   |
| tidb/22628 | [executor: Improve max/min window function with deque-based sliding window](https://github.com/pingcap/tidb/pull/22628)              |   | 30d23h  |
| tidb/22748 | [executor, privilege: fix failure on grant USAGE privilege operation](https://github.com/pingcap/tidb/pull/22748)                    |   | 24d16h  |
| tidb/22803 | [store/mockstore/unistore: refine and add more mpp tests](https://github.com/pingcap/tidb/pull/22803)                                |   | 10d21h  |
| tidb/22815 | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22815)                  |   | 10d18h  |
| tidb/22858 | [execution: fix index's duplicate error message](https://github.com/pingcap/tidb/pull/22858)                                         |   | 8d15h   |
| tidb/22961 | [statistics: refactor the statistics package use the RestrictedSQLExecutor API (#22636)](https://github.com/pingcap/tidb/pull/22961) |   | 3d23h   |
| tidb/22989 | [distsql: fix err check](https://github.com/pingcap/tidb/pull/22989)                                                                 |   | 2d0h    |
| tidb/23009 | [util: fix err check](https://github.com/pingcap/tidb/pull/23009)                                                                    |   | 1d23h   |
| tidb/23041 | [util: remove unused code](https://github.com/pingcap/tidb/pull/23041)                                                               |   | 9h      |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 

