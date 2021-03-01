|   REVIEWER    |  O(C)   | 1 | 2 | 3 | 4 | 5 | 6 | 7 | T  |
|---------------|---------|---|---|---|---|---|---|---|----|
| Reminiscent   |  4 ( 0) | 0 | 0 | 1 | 0 | 1 | 1 | 2 |  5 |
| SunRunAway    | 21 ( 1) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| XuHuaiyu      | 52 ( 3) | 0 | 0 | 2 | 1 | 0 | 1 | 6 | 10 |
| dyzsr         |  0 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| eurekaka      | 13 ( 1) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| fzhedu        |  2 ( 0) | 0 | 0 | 0 | 0 | 1 | 0 | 0 |  1 |
| ichn-hu       |  1 ( 0) | 0 | 0 | 2 | 0 | 1 | 0 | 1 |  4 |
| lzmhhh123     | 43 ( 0) | 0 | 0 | 1 | 2 | 0 | 1 | 1 |  5 |
| nullnotnil    |  0 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| qw4990        | 50 ( 1) | 0 | 0 | 1 | 1 | 5 | 3 | 5 | 15 |
| rebelice      |  2 ( 0) | 0 | 0 | 3 | 1 | 1 | 1 | 1 |  7 |
| time-and-fate |  1 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 2 |  2 |
| winoros       | 30 ( 2) | 0 | 0 | 0 | 1 | 1 | 0 | 3 |  5 |
| wshwsh12      | 20 ( 1) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |


<details> 
  <summary>Reminiscent's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                   PR                                                                   | C | LASTED |
|------------|----------------------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/21896 | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                             |   | 69d19h |
| tidb/22354 | [planner: do not cache prepared plan if optimization depends on mutable constant (#22349)](https://github.com/pingcap/tidb/pull/22354) |   | 47d23h |
| tidb/22918 | [sessionctx: add optimization-time and wait-TS-time into the slow log (#17869)](https://github.com/pingcap/tidb/pull/22918)            |   | 4d17h  |
| tidb/22938 | [planner: fix range partition prune bug for IN expr (#22894)](https://github.com/pingcap/tidb/pull/22938)                              |   | 3d18h  |


## Reviewed in Last 7 Days

|    REPO    |                                                              PR                                                               | C | D |   R    |
|------------|-------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/22948 | [statistics: add more tests to check accurateness of global-stats](https://github.com/pingcap/tidb/pull/22948)                |   | 3 | 17h    |
| tidb/22911 | [statistics: make dump/load/show-stats be compatible with global-stats](https://github.com/pingcap/tidb/pull/22911)           |   | 5 | 1h     |
| tidb/22625 | [planner, statistics: allow (auto) analyze single partition in dynamic-only mode](https://github.com/pingcap/tidb/pull/22625) |   | 6 | 24d14h |
| tidb/22866 | [*: modify the switch to control global stats](https://github.com/pingcap/tidb/pull/22866)                                    |   | 7 | 0h     |
| tidb/22848 | [statistics: merge poped topn when generating the global histogram](https://github.com/pingcap/tidb/pull/22848)               |   | 7 | 1d19h  |


</details> 


<details> 
  <summary>SunRunAway's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                  PR                                                                   | C | LASTED  |
|--------------|---------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19178   | [executor: Refactor probe channel](https://github.com/pingcap/tidb/pull/19178)                                                        |   | 199d16h |
| docs-cn/4669 | [sql-optimization: extended statistics documentation](https://github.com/pingcap/docs-cn/pull/4669)                                   |   | 138d17h |
| tidb/19347   | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347)  |   | 191d23h |
| tidb/19807   | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                               |   | 177d10h |
| tidb/19900   | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                        | Y | 172d18h |
| tidb/20140   | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                    |   | 159d22h |
| tidb/20220   | [*: new secondary index value format](https://github.com/pingcap/tidb/pull/20220)                                                     |   | 156d16h |
| tidb/20765   | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                     |   | 118d17h |
| tidb/21207   | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                       |   | 97d19h  |
| tidb/21277   | [executor: fix split table with large integers](https://github.com/pingcap/tidb/pull/21277)                                           |   | 95d19h  |
| tidb/21381   | [*: optimize analyze cluster index table](https://github.com/pingcap/tidb/pull/21381)                                                 |   | 90d17h  |
| tidb/21834   | [planner: enhanced index range calculation plan](https://github.com/pingcap/tidb/pull/21834)                                          |   | 74d18h  |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)           |   | 72d19h  |
| tidb/21878   | [planner: do not push down lock to pointGet/bacthPointGet when selection exists](https://github.com/pingcap/tidb/pull/21878)          |   | 72d18h  |
| tidb/21956   | [planner/preprocessor: disallow into-outfile clause in some place](https://github.com/pingcap/tidb/pull/21956)                        |   | 67d23h  |
| tidb/22026   | [expression: separated arithmeticPlusIntSig](https://github.com/pingcap/tidb/pull/22026)                                              |   | 65d20h  |
| tidb/22114   | [test: fix globalkilltest (#21987)](https://github.com/pingcap/tidb/pull/22114)                                                       |   | 60d12h  |
| tidb/22181   | [planner, expression: fix error when using IN combined with subquery (#22080)](https://github.com/pingcap/tidb/pull/22181)            |   | 54d17h  |
| tidb/22217   | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                         |   | 53d17h  |
| tidb/22365   | [planner: check index valid while forUpdateRead (#22152)](https://github.com/pingcap/tidb/pull/22365)                                 |   | 47d19h  |
| tidb/22379   | [[experiment] executor: allow aggregation to spill disk when running out of memory quota](https://github.com/pingcap/tidb/pull/22379) |   | 46d19h  |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>XuHuaiyu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                              PR                                                                              | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19900 | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                                               | Y | 172d18h |
| tidb/19957 | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                                       | Y | 170d13h |
| tidb/20040 | [planner, expression: take NullFlag into consideration when optimize the `int non-const` <cmp > `non-int const`](https://github.com/pingcap/tidb/pull/20040) | Y | 165d13h |
| tidb/20140 | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                                           |   | 159d22h |
| tidb/20311 | [expression: fix overflow error when convert bit to int64 (#20266)](https://github.com/pingcap/tidb/pull/20311)                                              |   | 151d21h |
| tidb/20790 | [collation: add pinyin collation for chinese charset support](https://github.com/pingcap/tidb/pull/20790)                                                    |   | 117d20h |
| tidb/20905 | [planner: fix statement-optimize not work in `TryFastPlan`](https://github.com/pingcap/tidb/pull/20905)                                                      |   | 114d17h |
| tidb/20972 | [expression: POC implementation of Vitess hashing algorithm.](https://github.com/pingcap/tidb/pull/20972)                                                    |   | 110d1h  |
| tidb/21064 | [planner, executor: fix cast not check error](https://github.com/pingcap/tidb/pull/21064)                                                                    |   | 105d8h  |
| tidb/21149 | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                                |   | 101d14h |
| tidb/21228 | [executor: return the result immediately when combining LIMIT row_count with DISTINCT](https://github.com/pingcap/tidb/pull/21228)                           |   | 97d13h  |
| tidb/21304 | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                                     |   | 95d12h  |
| tidb/21334 | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                                                |   | 94d14h  |
| tidb/21340 | [executor: initialize expensive query handler on domain creation](https://github.com/pingcap/tidb/pull/21340)                                                |   | 93d23h  |
| tidb/21476 | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                        |   | 87d15h  |
| tidb/21536 | [executor: add slow-log file meta cache to avoid repeat read file meta information](https://github.com/pingcap/tidb/pull/21536)                              |   | 83d14h  |
| tidb/21564 | [ddl: fix Incorrect behavior of NO_ZERO_DATE when altering table](https://github.com/pingcap/tidb/pull/21564)                                                |   | 82d15h  |
| tidb/21853 | [expression: fix compatibility behaviors in time_format with MySQL (#21559)](https://github.com/pingcap/tidb/pull/21853)                                     |   | 73d18h  |
| tidb/21896 | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                                                   |   | 69d19h  |
| tidb/22131 | [privilege: remove leading and trailing space when create user and role](https://github.com/pingcap/tidb/pull/22131)                                         |   | 59d19h  |
| tidb/22149 | [session: set process info before building plan (#22101)](https://github.com/pingcap/tidb/pull/22149)                                                        |   | 55d19h  |
| tidb/22163 | [expression: separated arithmeticMinusIntSig](https://github.com/pingcap/tidb/pull/22163)                                                                    |   | 55d13h  |
| tidb/22186 | [executor: fix select into outfile with year type column has no data (#22175)](https://github.com/pingcap/tidb/pull/22186)                                   |   | 54d16h  |
| tidb/22294 | [planner, table: optimize the list partition pruner for range query](https://github.com/pingcap/tidb/pull/22294)                                             |   | 51d20h  |
| tidb/22307 | [ddl: fix update can see columns not public](https://github.com/pingcap/tidb/pull/22307)                                                                     |   | 51d16h  |
| tidb/22381 | [planner: check schema stale for plan cache when forUpdateRead](https://github.com/pingcap/tidb/pull/22381)                                                  |   | 46d14h  |
| tidb/22616 | [expression: from_unixtime accept 64-bit integers](https://github.com/pingcap/tidb/pull/22616)                                                               |   | 30d23h  |
| tidb/22617 | [metrics: fix wrong bucket name of coprocessor cache (#22454)](https://github.com/pingcap/tidb/pull/22617)                                                   |   | 30d23h  |
| tidb/22624 | [ planner: not pruning column used by union scan condition (#21640)](https://github.com/pingcap/tidb/pull/22624)                                             |   | 30d17h  |
| tidb/22640 | [*: refactor ExecuteInternal to return single resultset (#22546)](https://github.com/pingcap/tidb/pull/22640)                                                |   | 27d20h  |
| tidb/22696 | [expression: enable arithmetic Mod push down](https://github.com/pingcap/tidb/pull/22696)                                                                    |   | 25d17h  |
| tidb/22711 | [executor: Fix inline schema name](https://github.com/pingcap/tidb/pull/22711)                                                                               |   | 25d11h  |
| tidb/22722 | [planner, errno: make error code of ErrMixOfGroupFuncAndFields consistent with MySQL](https://github.com/pingcap/tidb/pull/22722)                            |   | 24d20h  |
| tidb/22736 | [executor: fix load data losing connection when batch_dml_size is set (#22724)](https://github.com/pingcap/tidb/pull/22736)                                  |   | 23d23h  |
| tidb/22786 | [config: deprecate `tikv-client.copr-cache.enable`](https://github.com/pingcap/tidb/pull/22786)                                                              |   | 10d18h  |
| tidb/22814 | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22814)                                          |   | 9d19h   |
| tidb/22815 | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22815)                                          |   | 9d19h   |
| tidb/22832 | [expression: push down EXTRACT to TiFlash](https://github.com/pingcap/tidb/pull/22832)                                                                       |   | 9d1h    |
| tidb/22844 | [expression: do not adjust int when it is null and compared year (#22821)](https://github.com/pingcap/tidb/pull/22844)                                       |   | 8d19h   |
| tidb/22886 | [*: rename tidb_enable_tiflash_fallback_tikv to tidb_enable_engine_fallback](https://github.com/pingcap/tidb/pull/22886)                                     |   | 5d19h   |
| tidb/22903 | [docs/design: update template](https://github.com/pingcap/tidb/pull/22903)                                                                                   |   | 5d13h   |
| tidb/22914 | [partition: fix hash partition with not between condition get wrong result](https://github.com/pingcap/tidb/pull/22914)                                      |   | 4d18h   |
| tidb/22922 | [store/tikv: move coprocessor out](https://github.com/pingcap/tidb/pull/22922)                                                                               |   | 4d15h   |
| tidb/22957 | [config: use tidb_enable_list_partition to enable list table partition feature (#22864)](https://github.com/pingcap/tidb/pull/22957)                         |   | 3d12h   |
| tidb/22961 | [statistics: refactor the statistics package use the RestrictedSQLExecutor API (#22636)](https://github.com/pingcap/tidb/pull/22961)                         |   | 2d23h   |
| tidb/22962 | [executor: track partialResultMap in unparalleled aggreagte.](https://github.com/pingcap/tidb/pull/22962)                                                    |   | 2d23h   |
| tidb/22965 | [executor: introduce setWithMemoryUsage to track set memory in AggExec.](https://github.com/pingcap/tidb/pull/22965)                                         |   | 2d22h   |
| tidb/22971 | [Privileges: fix delete privilege check wrongly](https://github.com/pingcap/tidb/pull/22971)                                                                 |   | 2d19h   |
| tidb/22976 | [kv/union_store:remove tableinfo from union_store](https://github.com/pingcap/tidb/pull/22976)                                                               |   | 2d17h   |
| tidb/22983 | [*: fix some structcheck lint warnings](https://github.com/pingcap/tidb/pull/22983)                                                                          |   | 1d12h   |
| tidb/22992 | [expression: fix err check](https://github.com/pingcap/tidb/pull/22992)                                                                                      |   | 1d0h    |
| tidb/23012 | [executor: fix affected rows of creating database and dropping database](https://github.com/pingcap/tidb/pull/23012)                                         |   | 16h     |


## Reviewed in Last 7 Days

|    REPO    |                                                                  PR                                                                   | C | D |   R    |
|------------|---------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/22507 | [types: fix the bug about the wrong query result for decimal type ](https://github.com/pingcap/tidb/pull/22507)                       |   | 3 | 32d3h  |
| tidb/22426 | [expression: fix bugs in builtinfunction ArithmeticMinusInt logic](https://github.com/pingcap/tidb/pull/22426)                        |   | 3 | 38d17h |
| tidb/22912 | [executor: reduce useless memory allocation in min/max aggregate](https://github.com/pingcap/tidb/pull/22912)                         |   | 4 | 20h    |
| tidb/22883 | [executor: fix regression by Tracker.Consume in aggregate.](https://github.com/pingcap/tidb/pull/22883)                               |   | 6 | 0h     |
| tidb/22822 | [expression: fixed error msg in builtinArithmeticMinusIntSig](https://github.com/pingcap/tidb/pull/22822)                             |   | 7 | 2d23h  |
| tidb/22559 | [planner: split test data from test cases in cbo_test.go](https://github.com/pingcap/tidb/pull/22559)                                 |   | 7 | 26d2h  |
| tidb/22714 | [executor: add close recordSet in executor](https://github.com/pingcap/tidb/pull/22714)                                               |   | 7 | 18d6h  |
| tidb/22760 | [expression: fix wrong error info](https://github.com/pingcap/tidb/pull/22760)                                                        |   | 7 | 14d7h  |
| tidb/22830 | [planner: fix incorrect duration between compare](https://github.com/pingcap/tidb/pull/22830)                                         |   | 7 | 2d17h  |
| tidb/22861 | [executor: track memory usage of aggregate unparelled logic and distinct partial result.](https://github.com/pingcap/tidb/pull/22861) |   | 7 | 5h     |


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
| tidb/19347 | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347)   |   | 191d23h |
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                              |   | 115d16h |
| tidb/21444 | [planner: ignore anonymous index while tiflash replica is available](https://github.com/pingcap/tidb/pull/21444)                       |   | 88d12h  |
| tidb/21994 | [range: fix overflow value access index ](https://github.com/pingcap/tidb/pull/21994)                                                  |   | 66d22h  |
| tidb/22342 | [session: fix two cases when updating bind info (#22338)](https://github.com/pingcap/tidb/pull/22342)                                  |   | 48d18h  |
| tidb/22354 | [planner: do not cache prepared plan if optimization depends on mutable constant (#22349)](https://github.com/pingcap/tidb/pull/22354) |   | 47d23h  |
| tidb/22369 | [session: fix the duplicate binding case when updating bind info (#22367)](https://github.com/pingcap/tidb/pull/22369)                 |   | 47d17h  |
| tidb/22416 | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                                      | Y | 43d11h  |
| tidb/22559 | [planner: split test data from test cases in cbo_test.go](https://github.com/pingcap/tidb/pull/22559)                                  |   | 32d19h  |
| tidb/22733 | [bindinfo: use new sql apis (#22653)](https://github.com/pingcap/tidb/pull/22733)                                                      |   | 24d15h  |
| tidb/22778 | [*: add support for dynamic privileges](https://github.com/pingcap/tidb/pull/22778)                                                    |   | 12d7h   |
| tidb/22953 | [planner: fix query range partition table got wrong result and TiDB panic](https://github.com/pingcap/tidb/pull/22953)                 |   | 3d14h   |
| tidb/22985 | [bindinfo: fix error check](https://github.com/pingcap/tidb/pull/22985)                                                                |   | 1d1h    |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>fzhedu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                  PR                                                   | C | LASTED |
|------------|-------------------------------------------------------------------------------------------------------|---|--------|
| tidb/22662 | [planner/core: let mpp support partition tables](https://github.com/pingcap/tidb/pull/22662)          |   | 26d18h |
| tidb/22845 | [planner: fix bug of mpp wrongly set schema of exchanger](https://github.com/pingcap/tidb/pull/22845) |   | 8d17h  |


## Reviewed in Last 7 Days

|    REPO    |                                                  PR                                                   | C | D |  R   |
|------------|-------------------------------------------------------------------------------------------------------|---|---|------|
| tidb/22803 | [store/mockstore/unistore: refine and add more mpp tests](https://github.com/pingcap/tidb/pull/22803) |   | 5 | 5d9h |


</details> 


<details> 
  <summary>ichn-hu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                            PR                                                            | C | LASTED |
|------------|--------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/21853 | [expression: fix compatibility behaviors in time_format with MySQL (#21559)](https://github.com/pingcap/tidb/pull/21853) |   | 73d18h |


## Reviewed in Last 7 Days

|    REPO    |                                                           PR                                                           | C | D |   R   |
|------------|------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/22965 | [executor: introduce setWithMemoryUsage to track set memory in AggExec.](https://github.com/pingcap/tidb/pull/22965)   |   | 3 | 5h    |
| tidb/22962 | [executor: track partialResultMap in unparalleled aggreagte.](https://github.com/pingcap/tidb/pull/22962)              |   | 3 | 5h    |
| tidb/22912 | [executor: reduce useless memory allocation in min/max aggregate](https://github.com/pingcap/tidb/pull/22912)          |   | 5 | 3h    |
| tidb/22844 | [expression: do not adjust int when it is null and compared year (#22821)](https://github.com/pingcap/tidb/pull/22844) |   | 7 | 1d20h |


</details> 


<details> 
  <summary>lzmhhh123's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                  PR                                                                   | C | LASTED  |
|--------------|---------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19347   | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347)  |   | 191d23h |
| docs-cn/5561 | [Add sql optimization-related docs to toc](https://github.com/pingcap/docs-cn/pull/5561)                                              |   | 6d15h   |
| tidb/20444   | [expression: add json_merge_patch](https://github.com/pingcap/tidb/pull/20444)                                                        |   | 137d21h |
| tidb/20465   | [expression: add uuidShortFunction](https://github.com/pingcap/tidb/pull/20465)                                                       |   | 136d19h |
| tidb/20642   | [executor: modify admin executors to support partitioned table with global index](https://github.com/pingcap/tidb/pull/20642)         |   | 125d15h |
| tidb/20903   | [planner: fix confused and unnecessary double-projection in plans.](https://github.com/pingcap/tidb/pull/20903)                       |   | 114d17h |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                        |   | 108d17h |
| tidb/21195   | [brie: integrate lightning to suport IMPORT statement](https://github.com/pingcap/tidb/pull/21195)                                    |   | 97d22h  |
| tidb/21334   | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                         |   | 94d14h  |
| tidb/21347   | [session: make rollback work on global variables](https://github.com/pingcap/tidb/pull/21347)                                         |   | 93d19h  |
| tidb/21401   | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                    |   | 90d11h  |
| tidb/21444   | [planner: ignore anonymous index while tiflash replica is available](https://github.com/pingcap/tidb/pull/21444)                      |   | 88d12h  |
| tidb/21487   | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                         |   | 87d4h   |
| tidb/21641   | [executor: Fix pessimistic lock doesn't work on the partition table for subquery/joins](https://github.com/pingcap/tidb/pull/21641)   |   | 80d18h  |
| tidb/21651   | [planner: allow filter condition pushing down to IndexScan for prefix index](https://github.com/pingcap/tidb/pull/21651)              |   | 80d13h  |
| tidb/22126   | [*: add `sys` schema, `sys.SCHEMA_UNUSED_INDEXES` view and `sys.SCHEMA_INDEX_USAGE` view](https://github.com/pingcap/tidb/pull/22126) |   | 59d19h  |
| tidb/22149   | [session: set process info before building plan (#22101)](https://github.com/pingcap/tidb/pull/22149)                                 |   | 55d19h  |
| tidb/22188   | [planner: do not use indexMerge when the path only use a single index (#22168)](https://github.com/pingcap/tidb/pull/22188)           |   | 54d13h  |
| tidb/22361   | [table: fix insert into _tidb_rowid panic and rebase it if needed (#22062)](https://github.com/pingcap/tidb/pull/22361)               |   | 47d20h  |
| tidb/22372   | [executor: fix SelectForUpdate in decorrelated subquery under pessimistic mode](https://github.com/pingcap/tidb/pull/22372)           |   | 47d9h   |
| tidb/22430   | [*: refactor table.Table interface, clean up unnecessay methods](https://github.com/pingcap/tidb/pull/22430)                          |   | 40d23h  |
| tidb/22478   | [planner, executor: fix query partition table with global unique index get wrong result](https://github.com/pingcap/tidb/pull/22478)  |   | 38d13h  |
| tidb/22625   | [planner, statistics: allow (auto) analyze single partition in dynamic-only mode](https://github.com/pingcap/tidb/pull/22625)         |   | 30d14h  |
| tidb/22631   | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                                                       |   | 28d23h  |
| tidb/22656   | [*: move new api out of session package (#22591)](https://github.com/pingcap/tidb/pull/22656)                                         |   | 27d0h   |
| tidb/22662   | [planner/core: let mpp support partition tables](https://github.com/pingcap/tidb/pull/22662)                                          |   | 26d18h  |
| tidb/22699   | [brie: add error info column and history backup/restore info in sql](https://github.com/pingcap/tidb/pull/22699)                      |   | 25d16h  |
| tidb/22714   | [executor: add close recordSet in executor](https://github.com/pingcap/tidb/pull/22714)                                               |   | 24d23h  |
| tidb/22812   | [ executor: add new format specifier(%# %@ %.) for str_to_date expression (#22790)](https://github.com/pingcap/tidb/pull/22812)       |   | 9d19h   |
| tidb/22829   | [planner: let tikv know primary prefix columns info](https://github.com/pingcap/tidb/pull/22829)                                      |   | 9d13h   |
| tidb/22834   | [executor, server: load_data.go is changed and add unit test](https://github.com/pingcap/tidb/pull/22834)                             |   | 9d0h    |
| tidb/22857   | [store: move mockstore/mocktikv to tikv/mockstore/mocktikv](https://github.com/pingcap/tidb/pull/22857)                               |   | 7d21h   |
| tidb/22866   | [*: modify the switch to control global stats](https://github.com/pingcap/tidb/pull/22866)                                            |   | 6d18h   |
| tidb/22904   | [server: support MVCC-get HTTP API for clustered index](https://github.com/pingcap/tidb/pull/22904)                                   |   | 5d12h   |
| tidb/22910   | [util: optimize the performance of restore with db](https://github.com/pingcap/tidb/pull/22910)                                       |   | 4d19h   |
| tidb/22931   | [statistics: enables global-level stats to be generated in fast analyze mode](https://github.com/pingcap/tidb/pull/22931)             |   | 3d23h   |
| tidb/22944   | [planner/core: support mpp shuffled by Expression](https://github.com/pingcap/tidb/pull/22944)                                        |   | 3d17h   |
| tidb/22961   | [statistics: refactor the statistics package use the RestrictedSQLExecutor API (#22636)](https://github.com/pingcap/tidb/pull/22961)  |   | 2d23h   |
| tidb/22974   | [parser: update parser to fix sql restore bug used in create view](https://github.com/pingcap/tidb/pull/22974)                        |   | 2d17h   |
| tidb/22980   | [planner: choose non-prefix column when both index key and handle have the same one](https://github.com/pingcap/tidb/pull/22980)      |   | 2d13h   |
| tidb/22991   | [executor: fix err check](https://github.com/pingcap/tidb/pull/22991)                                                                 |   | 1d0h    |
| tidb/23001   | [statistics: fix err check](https://github.com/pingcap/tidb/pull/23001)                                                               |   | 1d0h    |
| tidb/23011   | [*: fix some varcheck lint warnings](https://github.com/pingcap/tidb/pull/23011)                                                      |   | 19h     |


## Reviewed in Last 7 Days

|    REPO    |                                                           PR                                                           | C | D |   R   |
|------------|------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/22959 | [planner, executor: reset NotNullFlag when merge schema for join (#22955)](https://github.com/pingcap/tidb/pull/22959) |   | 3 | 11h   |
| tidb/22955 | [planner, executor: reset NotNullFlag when merge schema for join](https://github.com/pingcap/tidb/pull/22955)          |   | 4 | 0h    |
| tidb/22940 | [planner: enable column pruning for common handle](https://github.com/pingcap/tidb/pull/22940)                         |   | 4 | 0h    |
| tidb/22883 | [executor: fix regression by Tracker.Consume in aggregate.](https://github.com/pingcap/tidb/pull/22883)                |   | 6 | 3h    |
| tidb/22822 | [expression: fixed error msg in builtinArithmeticMinusIntSig](https://github.com/pingcap/tidb/pull/22822)              |   | 7 | 2d23h |


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

|     REPO     |                                                                           PR                                                                           | C | LASTED  |
|--------------|--------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| docs-cn/5484 | [system variable: add tidb_enable_engine_fallback](https://github.com/pingcap/docs-cn/pull/5484)                                                       |   | 25d17h  |
| tidb/19029   | [types: fix unexpected NOT_NULL flags](https://github.com/pingcap/tidb/pull/19029)                                                                     |   | 206d22h |
| docs-cn/5561 | [Add sql optimization-related docs to toc](https://github.com/pingcap/docs-cn/pull/5561)                                                               |   | 6d15h   |
| tidb/20708   | [*: separate auto_increment ID allocator from _tidb_rowid allocator](https://github.com/pingcap/tidb/pull/20708)                                       |   | 122d20h |
| tidb/20969   | [executor: Improve the performance of appending not fixed columns](https://github.com/pingcap/tidb/pull/20969)                                         |   | 110d9h  |
| tidb/20972   | [expression: POC implementation of Vitess hashing algorithm.](https://github.com/pingcap/tidb/pull/20972)                                              |   | 110d1h  |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                         |   | 108d17h |
| tidb/21149   | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                          |   | 101d14h |
| tidb/21304   | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                               |   | 95d12h  |
| tidb/21318   | [planner, expression: use the range of column types to simplify expressions](https://github.com/pingcap/tidb/pull/21318)                               |   | 94d18h  |
| tidb/21401   | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                                     |   | 90d11h  |
| tidb/21476   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                  |   | 87d15h  |
| tidb/21508   | [execution: fix dayofweek('0000-00-00') behavior](https://github.com/pingcap/tidb/pull/21508)                                                          |   | 86d10h  |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                            |   | 72d19h  |
| tidb/21887   | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                                        |   | 71d11h  |
| tidb/21930   | [planner: propagate NDV of column groups across plan nodes (#17854)](https://github.com/pingcap/tidb/pull/21930)                                       |   | 68d18h  |
| tidb/21954   | [planner/cascades: add rule `PushSelDownApply`](https://github.com/pingcap/tidb/pull/21954)                                                            |   | 67d23h  |
| tidb/22090   | [planner: push aggregation operators down to projection and union by default](https://github.com/pingcap/tidb/pull/22090)                              |   | 60d22h  |
| tidb/22146   | [executor: forbid SFU on view](https://github.com/pingcap/tidb/pull/22146)                                                                             |   | 55d21h  |
| tidb/22217   | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                                          |   | 53d17h  |
| tidb/22234   | [executor, planner: ON DUPLICATE UPDATE can refer to un-project col (#14412)](https://github.com/pingcap/tidb/pull/22234)                              |   | 53d15h  |
| tidb/22261   | [time: fix parse datetime won't truncate the reluctant string (#22232)](https://github.com/pingcap/tidb/pull/22261)                                    |   | 52d19h  |
| tidb/22307   | [ddl: fix update can see columns not public](https://github.com/pingcap/tidb/pull/22307)                                                               |   | 51d16h  |
| tidb/22342   | [session: fix two cases when updating bind info (#22338)](https://github.com/pingcap/tidb/pull/22342)                                                  |   | 48d18h  |
| tidb/22369   | [session: fix the duplicate binding case when updating bind info (#22367)](https://github.com/pingcap/tidb/pull/22369)                                 |   | 47d17h  |
| tidb/22374   | [expression: separated arithmeticIntDivideSig](https://github.com/pingcap/tidb/pull/22374)                                                             |   | 47d0h   |
| tidb/22415   | [ddl: refactor placement package](https://github.com/pingcap/tidb/pull/22415)                                                                          |   | 43d17h  |
| tidb/22471   | [ddl, executor: fix creating unique index without partition column error when enable-global-index is true](https://github.com/pingcap/tidb/pull/22471) |   | 38d17h  |
| tidb/22489   | [infoschema: support query partition_id from infoschema.partitions (#22240)](https://github.com/pingcap/tidb/pull/22489)                               |   | 37d19h  |
| tidb/22507   | [types: fix the bug about the wrong query result for decimal type ](https://github.com/pingcap/tidb/pull/22507)                                        |   | 34d23h  |
| tidb/22541   | [expression: Support builtin function SOUNDEX](https://github.com/pingcap/tidb/pull/22541)                                                             |   | 33d9h   |
| tidb/22559   | [planner: split test data from test cases in cbo_test.go](https://github.com/pingcap/tidb/pull/22559)                                                  |   | 32d19h  |
| tidb/22565   | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)                                     | Y | 32d17h  |
| tidb/22641   | [*: do not report error for prepared stmt execution if tidb_snapshot is set (#22568)](https://github.com/pingcap/tidb/pull/22641)                      |   | 27d19h  |
| tidb/22733   | [bindinfo: use new sql apis (#22653)](https://github.com/pingcap/tidb/pull/22733)                                                                      |   | 24d15h  |
| tidb/22760   | [expression: fix wrong error info](https://github.com/pingcap/tidb/pull/22760)                                                                         |   | 21d0h   |
| tidb/22778   | [*: add support for dynamic privileges](https://github.com/pingcap/tidb/pull/22778)                                                                    |   | 12d7h   |
| tidb/22814   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22814)                                    |   | 9d19h   |
| tidb/22815   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22815)                                    |   | 9d19h   |
| tidb/22862   | [brie: fix the problem that ddl restored by BR via SQL is not replicated to downstream](https://github.com/pingcap/tidb/pull/22862)                    |   | 6d22h   |
| tidb/22869   | [executor: fix cast function will ignore tht error for point-get key construction](https://github.com/pingcap/tidb/pull/22869)                         |   | 6d16h   |
| tidb/22886   | [*: rename tidb_enable_tiflash_fallback_tikv to tidb_enable_engine_fallback](https://github.com/pingcap/tidb/pull/22886)                               |   | 5d19h   |
| tidb/22915   | [planner: build correct MaxOneRow info from multi-column conditions](https://github.com/pingcap/tidb/pull/22915)                                       |   | 4d17h   |
| tidb/22923   | [expression: correct constant propagation for collation (#22666)](https://github.com/pingcap/tidb/pull/22923)                                          |   | 4d14h   |
| tidb/22924   | [planner: fix wrong index merge selection (#22825)](https://github.com/pingcap/tidb/pull/22924)                                                        |   | 4d14h   |
| tidb/22926   | [expression: add overflow check in multiplyInt](https://github.com/pingcap/tidb/pull/22926)                                                            |   | 4d13h   |
| tidb/22938   | [planner: fix range partition prune bug for IN expr (#22894)](https://github.com/pingcap/tidb/pull/22938)                                              |   | 3d18h   |
| tidb/22977   | [config: fix data race caused by config.Labels](https://github.com/pingcap/tidb/pull/22977)                                                            |   | 2d17h   |
| tidb/22984   | [executor: fix logging format of prepared statements (#16062)](https://github.com/pingcap/tidb/pull/22984)                                             |   | 1d10h   |
| tidb/23008   | [types: fix err check](https://github.com/pingcap/tidb/pull/23008)                                                                                     |   | 1d0h    |


## Reviewed in Last 7 Days

|    REPO    |                                                                  PR                                                                  | C | D |   R    |
|------------|--------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/22931 | [statistics: enables global-level stats to be generated in fast analyze mode](https://github.com/pingcap/tidb/pull/22931)            |   | 3 | 1d4h   |
| tidb/22910 | [util: optimize the performance of restore with db](https://github.com/pingcap/tidb/pull/22910)                                      |   | 4 | 1d0h   |
| tidb/22825 | [planner: fix wrong index merge selection](https://github.com/pingcap/tidb/pull/22825)                                               |   | 5 | 5d0h   |
| tidb/22803 | [store/mockstore/unistore: refine and add more mpp tests](https://github.com/pingcap/tidb/pull/22803)                                |   | 5 | 5d6h   |
| tidb/22866 | [*: modify the switch to control global stats](https://github.com/pingcap/tidb/pull/22866)                                           |   | 5 | 2d1h   |
| tidb/22666 | [expression: correct constant propagation for collation](https://github.com/pingcap/tidb/pull/22666)                                 |   | 5 | 22d0h  |
| tidb/22649 | [planner: decorrelate LogicalApply with inner join as the inner child](https://github.com/pingcap/tidb/pull/22649)                   |   | 5 | 22d21h |
| tidb/22682 | [statistics: introduce new estimation logic when index histogram fails to estimate](https://github.com/pingcap/tidb/pull/22682)      |   | 6 | 20d19h |
| tidb/22878 | [statistics: merge partition-level FMSketch to global-level FMSketch and update the NDV](https://github.com/pingcap/tidb/pull/22878) |   | 6 | 7h     |
| tidb/22836 | [expression: add integration test for partition pruning](https://github.com/pingcap/tidb/pull/22836)                                 |   | 6 | 3d3h   |
| tidb/22841 | [statistics: support to store FMSketch and add FMSketch to column stats](https://github.com/pingcap/tidb/pull/22841)                 |   | 7 | 2d6h   |
| tidb/22677 | [*: use `explain format = 'brief'` for tests](https://github.com/pingcap/tidb/pull/22677)                                            |   | 7 | 19d21h |
| tidb/22846 | [stats, planner, sessionctx: handle compatibility between feedback and ver2 stats](https://github.com/pingcap/tidb/pull/22846)       |   | 7 | 1d21h  |
| tidb/22848 | [statistics: merge poped topn when generating the global histogram](https://github.com/pingcap/tidb/pull/22848)                      |   | 7 | 1d17h  |
| tidb/22833 | [planner: fix where condition index out of range](https://github.com/pingcap/tidb/pull/22833)                                        |   | 7 | 2d2h   |


</details> 


<details> 
  <summary>rebelice's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                               PR                                                               | C | LASTED |
|------------|--------------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/22869 | [executor: fix cast function will ignore tht error for point-get key construction](https://github.com/pingcap/tidb/pull/22869) |   | 6d16h  |
| tidb/22938 | [planner: fix range partition prune bug for IN expr (#22894)](https://github.com/pingcap/tidb/pull/22938)                      |   | 3d18h  |


## Reviewed in Last 7 Days

|    REPO    |                                                                  PR                                                                  | C | D |  R   |
|------------|--------------------------------------------------------------------------------------------------------------------------------------|---|---|------|
| tidb/22931 | [statistics: enables global-level stats to be generated in fast analyze mode](https://github.com/pingcap/tidb/pull/22931)            |   | 3 | 1d4h |
| tidb/22948 | [statistics: add more tests to check accurateness of global-stats](https://github.com/pingcap/tidb/pull/22948)                       |   | 3 | 20h  |
| tidb/22959 | [planner, executor: reset NotNullFlag when merge schema for join (#22955)](https://github.com/pingcap/tidb/pull/22959)               |   | 3 | 11h  |
| tidb/22955 | [planner, executor: reset NotNullFlag when merge schema for join](https://github.com/pingcap/tidb/pull/22955)                        |   | 4 | 1h   |
| tidb/22911 | [statistics: make dump/load/show-stats be compatible with global-stats](https://github.com/pingcap/tidb/pull/22911)                  |   | 5 | 2h   |
| tidb/22878 | [statistics: merge partition-level FMSketch to global-level FMSketch and update the NDV](https://github.com/pingcap/tidb/pull/22878) |   | 6 | 7h   |
| tidb/22841 | [statistics: support to store FMSketch and add FMSketch to column stats](https://github.com/pingcap/tidb/pull/22841)                 |   | 7 | 2d6h |


</details> 


<details> 
  <summary>time-and-fate's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                            PR                                             | C | LASTED  |
|------------|-------------------------------------------------------------------------------------------|---|---------|
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877) |   | 115d16h |


## Reviewed in Last 7 Days

|    REPO    |                                                               PR                                                                | C | D |   R    |
|------------|---------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/22677 | [*: use `explain format = 'brief'` for tests](https://github.com/pingcap/tidb/pull/22677)                                       |   | 7 | 19d21h |
| tidb/22682 | [statistics: introduce new estimation logic when index histogram fails to estimate](https://github.com/pingcap/tidb/pull/22682) |   | 7 | 19d14h |


</details> 


<details> 
  <summary>winoros's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                PR                                                                | C | LASTED  |
|--------------|----------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                           | Y | 170d13h |
| docs-cn/5484 | [system variable: add tidb_enable_engine_fallback](https://github.com/pingcap/docs-cn/pull/5484)                                 |   | 25d17h  |
| tidb/20311   | [expression: fix overflow error when convert bit to int64 (#20266)](https://github.com/pingcap/tidb/pull/20311)                  |   | 151d21h |
| tidb/20765   | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                |   | 118d17h |
| tidb/20877   | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                        |   | 115d16h |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                   |   | 108d17h |
| tidb/21207   | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                  |   | 97d19h  |
| tidb/21476   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                            |   | 87d15h  |
| tidb/21487   | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                    |   | 87d4h   |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)      |   | 72d19h  |
| tidb/21930   | [planner: propagate NDV of column groups across plan nodes (#17854)](https://github.com/pingcap/tidb/pull/21930)                 |   | 68d18h  |
| tidb/21954   | [planner/cascades: add rule `PushSelDownApply`](https://github.com/pingcap/tidb/pull/21954)                                      |   | 67d23h  |
| tidb/22090   | [planner: push aggregation operators down to projection and union by default](https://github.com/pingcap/tidb/pull/22090)        |   | 60d22h  |
| tidb/22365   | [planner: check index valid while forUpdateRead (#22152)](https://github.com/pingcap/tidb/pull/22365)                            |   | 47d19h  |
| tidb/22489   | [infoschema: support query partition_id from infoschema.partitions (#22240)](https://github.com/pingcap/tidb/pull/22489)         |   | 37d19h  |
| tidb/22504   | [*:Fix the fetchHotRegion bug that the count always zero](https://github.com/pingcap/tidb/pull/22504)                            |   | 35d19h  |
| tidb/22565   | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)               | Y | 32d17h  |
| tidb/22624   | [ planner: not pruning column used by union scan condition (#21640)](https://github.com/pingcap/tidb/pull/22624)                 |   | 30d17h  |
| tidb/22804   | [session, util: update session to use new APIs (#22652)](https://github.com/pingcap/tidb/pull/22804)                             |   | 9d21h   |
| tidb/22830   | [planner: fix incorrect duration between compare](https://github.com/pingcap/tidb/pull/22830)                                    |   | 9d10h   |
| tidb/22845   | [planner: fix bug of mpp wrongly set schema of exchanger](https://github.com/pingcap/tidb/pull/22845)                            |   | 8d17h   |
| tidb/22867   | [WIP: allow pushdown count distinct when enumerate physical plans](https://github.com/pingcap/tidb/pull/22867)                   |   | 6d17h   |
| tidb/22886   | [*: rename tidb_enable_tiflash_fallback_tikv to tidb_enable_engine_fallback](https://github.com/pingcap/tidb/pull/22886)         |   | 5d19h   |
| tidb/22914   | [partition: fix hash partition with not between condition get wrong result](https://github.com/pingcap/tidb/pull/22914)          |   | 4d18h   |
| tidb/22915   | [planner: build correct MaxOneRow info from multi-column conditions](https://github.com/pingcap/tidb/pull/22915)                 |   | 4d17h   |
| tidb/22923   | [expression: correct constant propagation for collation (#22666)](https://github.com/pingcap/tidb/pull/22923)                    |   | 4d14h   |
| tidb/22924   | [planner: fix wrong index merge selection (#22825)](https://github.com/pingcap/tidb/pull/22924)                                  |   | 4d14h   |
| tidb/22940   | [planner: enable column pruning for common handle](https://github.com/pingcap/tidb/pull/22940)                                   |   | 3d17h   |
| tidb/22971   | [Privileges: fix delete privilege check wrongly](https://github.com/pingcap/tidb/pull/22971)                                     |   | 2d19h   |
| tidb/22980   | [planner: choose non-prefix column when both index key and handle have the same one](https://github.com/pingcap/tidb/pull/22980) |   | 2d13h   |


## Reviewed in Last 7 Days

|     REPO     |                                                               PR                                                               | C | D |   R    |
|--------------|--------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| docs-cn/4669 | [sql-optimization: extended statistics documentation](https://github.com/pingcap/docs-cn/pull/4669)                            |   | 4 | 135d0h |
| tidb/22666   | [expression: correct constant propagation for collation](https://github.com/pingcap/tidb/pull/22666)                           |   | 5 | 22d0h  |
| tidb/21651   | [planner: allow filter condition pushing down to IndexScan for prefix index](https://github.com/pingcap/tidb/pull/21651)       |   | 7 | 73d19h |
| tidb/22846   | [stats, planner, sessionctx: handle compatibility between feedback and ver2 stats](https://github.com/pingcap/tidb/pull/22846) |   | 7 | 1d22h  |
| tidb/22825   | [planner: fix wrong index merge selection](https://github.com/pingcap/tidb/pull/22825)                                         |   | 7 | 2d20h  |


</details> 


<details> 
  <summary>wshwsh12's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                  PR                                                                  | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19557 | [*: Integrate timeline tracing with TiKV](https://github.com/pingcap/tidb/pull/19557)                                                |   | 184d23h |
| tidb/19807 | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                              |   | 177d10h |
| tidb/19957 | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                               | Y | 170d13h |
| tidb/21010 | [session: add a test case for issue 19127](https://github.com/pingcap/tidb/pull/21010)                                               |   | 108d19h |
| tidb/21381 | [*: optimize analyze cluster index table](https://github.com/pingcap/tidb/pull/21381)                                                |   | 90d17h  |
| tidb/21487 | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                        |   | 87d4h   |
| tidb/21887 | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                      |   | 71d11h  |
| tidb/22176 | [store, executor: support stale read for tikv RPCContext](https://github.com/pingcap/tidb/pull/22176)                                |   | 54d19h  |
| tidb/22269 | [executor: check storage.block-cache.capacity value](https://github.com/pingcap/tidb/pull/22269)                                     |   | 52d17h  |
| tidb/22378 | [executor: vectorize hash aggregate](https://github.com/pingcap/tidb/pull/22378)                                                     |   | 46d19h  |
| tidb/22382 | [*: add infoschema client errors](https://github.com/pingcap/tidb/pull/22382)                                                        |   | 46d5h   |
| tidb/22628 | [executor: Improve max/min window function with deque-based sliding window](https://github.com/pingcap/tidb/pull/22628)              |   | 29d23h  |
| tidb/22748 | [executor, privilege: fix failure on grant USAGE privilege operation](https://github.com/pingcap/tidb/pull/22748)                    |   | 23d16h  |
| tidb/22803 | [store/mockstore/unistore: refine and add more mpp tests](https://github.com/pingcap/tidb/pull/22803)                                |   | 9d21h   |
| tidb/22815 | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22815)                  |   | 9d19h   |
| tidb/22829 | [planner: let tikv know primary prefix columns info](https://github.com/pingcap/tidb/pull/22829)                                     |   | 9d13h   |
| tidb/22858 | [execution: fix index's duplicate error message](https://github.com/pingcap/tidb/pull/22858)                                         |   | 7d15h   |
| tidb/22961 | [statistics: refactor the statistics package use the RestrictedSQLExecutor API (#22636)](https://github.com/pingcap/tidb/pull/22961) |   | 2d23h   |
| tidb/22989 | [distsql: fix err check](https://github.com/pingcap/tidb/pull/22989)                                                                 |   | 1d0h    |
| tidb/23009 | [util: fix err check](https://github.com/pingcap/tidb/pull/23009)                                                                    |   | 1d0h    |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 

