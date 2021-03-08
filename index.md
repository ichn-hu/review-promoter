|   REVIEWER    |  O(C)   | 1 | 2 | 3 | 4 | 5 | 6 | 7 | T  |
|---------------|---------|---|---|---|---|---|---|---|----|
| Reminiscent   |  4 ( 0) | 0 | 0 | 2 | 0 | 5 | 1 | 3 | 11 |
| SunRunAway    | 19 ( 1) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| XuHuaiyu      | 55 ( 3) | 0 | 0 | 1 | 5 | 1 | 2 | 0 |  9 |
| dyzsr         |  0 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| eurekaka      | 14 ( 1) | 0 | 0 | 0 | 0 | 0 | 1 | 0 |  1 |
| fzhedu        |  1 ( 0) | 0 | 0 | 1 | 0 | 0 | 1 | 2 |  4 |
| ichn-hu       |  1 ( 0) | 0 | 0 | 1 | 1 | 0 | 2 | 0 |  4 |
| lzmhhh123     | 35 ( 0) | 0 | 0 | 0 | 1 | 2 | 0 | 4 |  7 |
| nullnotnil    |  0 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| qw4990        | 55 ( 1) | 0 | 0 | 2 | 1 | 5 | 1 | 5 | 14 |
| rebelice      |  2 ( 0) | 0 | 0 | 4 | 1 | 2 | 2 | 0 |  9 |
| time-and-fate |  5 ( 0) | 0 | 0 | 1 | 1 | 3 | 0 | 0 |  5 |
| winoros       | 28 ( 2) | 0 | 0 | 2 | 1 | 3 | 8 | 1 | 15 |
| wshwsh12      | 17 ( 1) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |


<details> 
  <summary>Reminiscent's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                   PR                                                                   | C | LASTED |
|------------|----------------------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/21896 | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                             |   | 76d19h |
| tidb/22354 | [planner: do not cache prepared plan if optimization depends on mutable constant (#22349)](https://github.com/pingcap/tidb/pull/22354) |   | 54d23h |
| tidb/23074 | [planner: fix range partition prune bug for IN expr (#22894) (#22938)](https://github.com/pingcap/tidb/pull/23074)                     |   | 4d17h  |
| tidb/23138 | [statistics: add some test cases of global-stats to cover more column types](https://github.com/pingcap/tidb/pull/23138)               |   | 2d16h  |


## Reviewed in Last 7 Days

|    REPO     |                                                                        PR                                                                        | C | D |   R    |
|-------------|--------------------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/23113  | [statistics: add tests for `analyze` with dynamic partition prune mode](https://github.com/pingcap/tidb/pull/23113)                              |   | 3 | 22h    |
| tidb/23099  | [statistics: add more test cases about global-stats and fix some issues](https://github.com/pingcap/tidb/pull/23099)                             |   | 3 | 23h    |
| tidb/23075  | [statistics: fix merge idx hist with poped topn](https://github.com/pingcap/tidb/pull/23075)                                                     |   | 5 | 0h     |
| tidb/22918  | [sessionctx: add optimization-time and wait-TS-time into the slow log (#17869)](https://github.com/pingcap/tidb/pull/22918)                      |   | 5 | 6d22h  |
| tidb/23061  | [statistics: introduce a new kind of syntax to drop global-stats](https://github.com/pingcap/tidb/pull/23061)                                    |   | 5 | 0h     |
| parser/1181 | [statistics: introduce a new kind of syntax to drop global-stats](https://github.com/pingcap/parser/pull/1181)                                   |   | 5 | 0h     |
| tidb/23050  | [statistics: forbid global-stats in stats-ver1 and make analyze options take effect in global-stats](https://github.com/pingcap/tidb/pull/23050) |   | 5 | 19h    |
| tidb/23023  | [statistics: support dropping partition/global level statistics](https://github.com/pingcap/tidb/pull/23023)                                     |   | 6 | 1d0h   |
| parser/1177 | [statistics: add a new syntax to support dropping statistics of partitions](https://github.com/pingcap/parser/pull/1177)                         |   | 7 | 0h     |
| tidb/22625  | [planner, statistics: allow (auto) analyze single partition in dynamic-only mode](https://github.com/pingcap/tidb/pull/22625)                    |   | 7 | 30d20h |
| tidb/23016  | [statistics: add static partition prune mode checks](https://github.com/pingcap/tidb/pull/23016)                                                 |   | 7 | 2h     |


</details> 


<details> 
  <summary>SunRunAway's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                  PR                                                                   | C | LASTED  |
|------------|---------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19178 | [executor: Refactor probe channel](https://github.com/pingcap/tidb/pull/19178)                                                        |   | 206d16h |
| tidb/19347 | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347)  |   | 198d23h |
| tidb/19807 | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                               |   | 184d10h |
| tidb/19900 | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                        | Y | 179d18h |
| tidb/20140 | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                    |   | 166d22h |
| tidb/20360 | [planner: refine explain info for batch cop](https://github.com/pingcap/tidb/pull/20360)                                              |   | 149d22h |
| tidb/20765 | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                     |   | 125d17h |
| tidb/21207 | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                       |   | 104d19h |
| tidb/21277 | [executor: fix split table with large integers](https://github.com/pingcap/tidb/pull/21277)                                           |   | 102d20h |
| tidb/21834 | [planner: enhanced index range calculation plan](https://github.com/pingcap/tidb/pull/21834)                                          |   | 81d18h  |
| tidb/21876 | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)           |   | 79d19h  |
| tidb/21878 | [planner: do not push down lock to pointGet/bacthPointGet when selection exists](https://github.com/pingcap/tidb/pull/21878)          |   | 79d18h  |
| tidb/21956 | [planner/preprocessor: disallow into-outfile clause in some place](https://github.com/pingcap/tidb/pull/21956)                        |   | 74d23h  |
| tidb/22026 | [expression: separated arithmeticPlusIntSig](https://github.com/pingcap/tidb/pull/22026)                                              |   | 72d20h  |
| tidb/22114 | [test: fix globalkilltest (#21987)](https://github.com/pingcap/tidb/pull/22114)                                                       |   | 67d12h  |
| tidb/22181 | [planner, expression: fix error when using IN combined with subquery (#22080)](https://github.com/pingcap/tidb/pull/22181)            |   | 61d17h  |
| tidb/22217 | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                         |   | 60d17h  |
| tidb/22365 | [planner: check index valid while forUpdateRead (#22152)](https://github.com/pingcap/tidb/pull/22365)                                 |   | 54d19h  |
| tidb/22379 | [[experiment] executor: allow aggregation to spill disk when running out of memory quota](https://github.com/pingcap/tidb/pull/22379) |   | 53d19h  |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>XuHuaiyu's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                              PR                                                                              | C | LASTED  |
|--------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19900   | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                                               | Y | 179d18h |
| docs-cn/5671 | [tidb: Add time format description](https://github.com/pingcap/docs-cn/pull/5671)                                                                            |   | 3d11h   |
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                                       | Y | 177d14h |
| tidb/20040   | [planner, expression: take NullFlag into consideration when optimize the `int non-const` <cmp > `non-int const`](https://github.com/pingcap/tidb/pull/20040) | Y | 172d14h |
| tidb/20140   | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                                           |   | 166d22h |
| tidb/20311   | [expression: fix overflow error when convert bit to int64 (#20266)](https://github.com/pingcap/tidb/pull/20311)                                              |   | 158d21h |
| tidb/20790   | [collation: add pinyin collation for chinese charset support](https://github.com/pingcap/tidb/pull/20790)                                                    |   | 124d20h |
| tidb/20905   | [planner: fix statement-optimize not work in `TryFastPlan`](https://github.com/pingcap/tidb/pull/20905)                                                      |   | 121d17h |
| tidb/20972   | [expression: POC implementation of Vitess hashing algorithm.](https://github.com/pingcap/tidb/pull/20972)                                                    |   | 117d1h  |
| tidb/21064   | [planner, executor: fix cast not check error](https://github.com/pingcap/tidb/pull/21064)                                                                    |   | 112d8h  |
| tidb/21149   | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                                |   | 108d14h |
| tidb/21228   | [executor: return the result immediately when combining LIMIT row_count with DISTINCT](https://github.com/pingcap/tidb/pull/21228)                           |   | 104d13h |
| tidb/21304   | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                                     |   | 102d12h |
| tidb/21334   | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                                                |   | 101d14h |
| tidb/21340   | [executor: initialize expensive query handler on domain creation](https://github.com/pingcap/tidb/pull/21340)                                                |   | 100d23h |
| tidb/21476   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                        |   | 94d15h  |
| tidb/21536   | [executor: add slow-log file meta cache to avoid repeat read file meta information](https://github.com/pingcap/tidb/pull/21536)                              |   | 90d14h  |
| tidb/21564   | [ddl: fix Incorrect behavior of NO_ZERO_DATE when altering table](https://github.com/pingcap/tidb/pull/21564)                                                |   | 89d15h  |
| tidb/21853   | [expression: fix compatibility behaviors in time_format with MySQL (#21559)](https://github.com/pingcap/tidb/pull/21853)                                     |   | 80d19h  |
| tidb/21896   | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                                                   |   | 76d19h  |
| tidb/22131   | [privilege: remove leading and trailing space when create user and role](https://github.com/pingcap/tidb/pull/22131)                                         |   | 66d19h  |
| tidb/22149   | [session: set process info before building plan (#22101)](https://github.com/pingcap/tidb/pull/22149)                                                        |   | 62d19h  |
| tidb/22163   | [expression: separated arithmeticMinusIntSig](https://github.com/pingcap/tidb/pull/22163)                                                                    |   | 62d13h  |
| tidb/22186   | [executor: fix select into outfile with year type column has no data (#22175)](https://github.com/pingcap/tidb/pull/22186)                                   |   | 61d16h  |
| tidb/22294   | [planner, table: optimize the list partition pruner for range query](https://github.com/pingcap/tidb/pull/22294)                                             |   | 58d20h  |
| tidb/22307   | [ddl: fix update can see columns not public](https://github.com/pingcap/tidb/pull/22307)                                                                     |   | 58d16h  |
| tidb/22381   | [planner: check schema stale for plan cache when forUpdateRead](https://github.com/pingcap/tidb/pull/22381)                                                  |   | 53d14h  |
| tidb/22616   | [expression: from_unixtime accept 64-bit integers](https://github.com/pingcap/tidb/pull/22616)                                                               |   | 37d23h  |
| tidb/22617   | [metrics: fix wrong bucket name of coprocessor cache (#22454)](https://github.com/pingcap/tidb/pull/22617)                                                   |   | 37d23h  |
| tidb/22624   | [ planner: not pruning column used by union scan condition (#21640)](https://github.com/pingcap/tidb/pull/22624)                                             |   | 37d17h  |
| tidb/22631   | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                                                                              |   | 35d23h  |
| tidb/22696   | [expression: enable arithmetic Mod push down](https://github.com/pingcap/tidb/pull/22696)                                                                    |   | 32d17h  |
| tidb/22711   | [executor: Fix inline schema name](https://github.com/pingcap/tidb/pull/22711)                                                                               |   | 32d11h  |
| tidb/22722   | [planner, errno: make error code of ErrMixOfGroupFuncAndFields consistent with MySQL](https://github.com/pingcap/tidb/pull/22722)                            |   | 31d20h  |
| tidb/22736   | [executor: fix load data losing connection when batch_dml_size is set (#22724)](https://github.com/pingcap/tidb/pull/22736)                                  |   | 30d23h  |
| tidb/22786   | [config: deprecate some configs of `tikv-client.copr-cache`](https://github.com/pingcap/tidb/pull/22786)                                                     |   | 17d18h  |
| tidb/22814   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22814)                                          |   | 16d19h  |
| tidb/22815   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22815)                                          |   | 16d19h  |
| tidb/22832   | [expression: push down EXTRACT to TiFlash](https://github.com/pingcap/tidb/pull/22832)                                                                       |   | 16d1h   |
| tidb/22844   | [expression: do not adjust int when it is null and compared year (#22821)](https://github.com/pingcap/tidb/pull/22844)                                       |   | 15d19h  |
| tidb/22886   | [*: rename tidb_enable_tiflash_fallback_tikv to tidb_enable_engine_fallback](https://github.com/pingcap/tidb/pull/22886)                                     |   | 12d20h  |
| tidb/22914   | [partition: fix hash partition with not between condition get wrong result](https://github.com/pingcap/tidb/pull/22914)                                      |   | 11d18h  |
| tidb/23012   | [executor: fix affected rows of ddls and complete uint tests](https://github.com/pingcap/tidb/pull/23012)                                                    |   | 7d16h   |
| tidb/23056   | [MPP: Kill mpp queries](https://github.com/pingcap/tidb/pull/23056)                                                                                          |   | 5d12h   |
| tidb/23072   | [executor: track memory usage of map in agg partial result.](https://github.com/pingcap/tidb/pull/23072)                                                     |   | 4d18h   |
| tidb/23092   | [*: fix a bug that collation is not handle for text type (#23045)](https://github.com/pingcap/tidb/pull/23092)                                               |   | 4d12h   |
| tidb/23104   | [executor: fix wrong key range of index scan when filter is comparing year column with NULL (#23079)](https://github.com/pingcap/tidb/pull/23104)            |   | 3d18h   |
| tidb/23105   | [executor: fix wrong key range of index scan when filter is comparing year column with NULL (#23079)](https://github.com/pingcap/tidb/pull/23105)            |   | 3d18h   |
| tidb/23111   | [executor: fix linter --enable=deadcode check error in executor(#22979)](https://github.com/pingcap/tidb/pull/23111)                                         |   | 3d17h   |
| tidb/23123   | [planner: show cast type in EXPLAIN in coptask](https://github.com/pingcap/tidb/pull/23123)                                                                  |   | 3d13h   |
| tidb/23128   | [statistics: refactor the statistics package use the RestrictedSQLExecutor API (#22636)](https://github.com/pingcap/tidb/pull/23128)                         |   | 2d22h   |
| tidb/23131   | [executor: group_concat aggr panic when session.group_concat_max_len is small.](https://github.com/pingcap/tidb/pull/23131)                                  |   | 2d19h   |
| tidb/23135   | [executor: fix unexpected NotNullFlag in case when expr ret type (#23102)](https://github.com/pingcap/tidb/pull/23135)                                       |   | 2d17h   |
| tidb/23139   | [executor: inject random panic to AggExec](https://github.com/pingcap/tidb/pull/23139)                                                                       |   | 2d16h   |
| tidb/23152   | [expression: fix wrong error info (#22760)](https://github.com/pingcap/tidb/pull/23152)                                                                      |   | 14h     |


## Reviewed in Last 7 Days

|     REPO     |                                                                    PR                                                                    | C | D |   R    |
|--------------|------------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/22869   | [executor: fix cast function will ignore tht error for point-get key construction](https://github.com/pingcap/tidb/pull/22869)           |   | 3 | 10d17h |
| tidb/23102   | [executor: fix unexpected NotNullFlag in case when expr ret type](https://github.com/pingcap/tidb/pull/23102)                            |   | 4 | 18h    |
| docs-cn/5619 | [Update data-type-date-and-time.md](https://github.com/pingcap/docs-cn/pull/5619)                                                        |   | 4 | 5d21h  |
| docs-cn/5620 | [Add details for Hexadecimal Literals](https://github.com/pingcap/docs-cn/pull/5620)                                                     |   | 4 | 5d21h  |
| tidb/23079   | [executor: fix wrong key range of index scan when filter is comparing year column with NULL](https://github.com/pingcap/tidb/pull/23079) |   | 4 | 20h    |
| tidb/23024   | [executor: make the memory tracker of Jsonobjectagg more accurate](https://github.com/pingcap/tidb/pull/23024)                           |   | 4 | 2d19h  |
| tidb/23034   | [executor: make the memory tracker of groupConcat more accurate.](https://github.com/pingcap/tidb/pull/23034)                            |   | 5 | 1d23h  |
| tidb/22962   | [executor: track partialResultMap in unparalleled aggreagte.](https://github.com/pingcap/tidb/pull/22962)                                |   | 6 | 4d3h   |
| tidb/22903   | [docs/design: update template](https://github.com/pingcap/tidb/pull/22903)                                                               |   | 6 | 6d15h  |


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
| tidb/19347 | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347)   |   | 198d23h |
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                              |   | 122d16h |
| tidb/21444 | [planner: ignore anonymous index while tiflash replica is available](https://github.com/pingcap/tidb/pull/21444)                       |   | 95d12h  |
| tidb/21994 | [range: fix overflow value access index ](https://github.com/pingcap/tidb/pull/21994)                                                  |   | 73d22h  |
| tidb/22342 | [session: fix two cases when updating bind info (#22338)](https://github.com/pingcap/tidb/pull/22342)                                  |   | 55d18h  |
| tidb/22354 | [planner: do not cache prepared plan if optimization depends on mutable constant (#22349)](https://github.com/pingcap/tidb/pull/22354) |   | 54d23h  |
| tidb/22369 | [session: fix the duplicate binding case when updating bind info (#22367)](https://github.com/pingcap/tidb/pull/22369)                 |   | 54d17h  |
| tidb/22416 | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                                      | Y | 50d11h  |
| tidb/22559 | [planner: split test data from test cases in cbo_test.go](https://github.com/pingcap/tidb/pull/22559)                                  |   | 39d19h  |
| tidb/22733 | [bindinfo: use new sql apis (#22653)](https://github.com/pingcap/tidb/pull/22733)                                                      |   | 31d15h  |
| tidb/22778 | [*: add support for dynamic privileges](https://github.com/pingcap/tidb/pull/22778)                                                    |   | 19d7h   |
| tidb/22853 | [planner: fix LogicalPlans that contain Window Function are ambiguous ](https://github.com/pingcap/tidb/pull/22853)                    |   | 15d12h  |
| tidb/22910 | [util: optimize the performance of restore with db](https://github.com/pingcap/tidb/pull/22910)                                        |   | 11d19h  |
| tidb/22953 | [planner: fix query range partition table got wrong result and TiDB panic](https://github.com/pingcap/tidb/pull/22953)                 |   | 10d14h  |


## Reviewed in Last 7 Days

|    REPO    |                                                             PR                                                             | C | D | R  |
|------------|----------------------------------------------------------------------------------------------------------------------------|---|---|----|
| tidb/23047 | [planner, expression: fix error when using IN combined with subquery (#22080)](https://github.com/pingcap/tidb/pull/23047) |   | 6 | 0h |


</details> 


<details> 
  <summary>fzhedu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                         PR                                                          | C | LASTED |
|------------|---------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/22853 | [planner: fix LogicalPlans that contain Window Function are ambiguous ](https://github.com/pingcap/tidb/pull/22853) |   | 15d12h |


## Reviewed in Last 7 Days

|    REPO    |                                                                  PR                                                                   | C | D |   R   |
|------------|---------------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tics/1506  | [check block schema in exchange operator](https://github.com/pingcap/tics/pull/1506)                                                  |   | 3 | 2d5h  |
| tidb/22803 | [store/mockstore/unistore: refine and add more mpp tests](https://github.com/pingcap/tidb/pull/22803)                                 |   | 6 | 11d9h |
| tics/1342  | [Add background task to detect && cancel hanging MPP query](https://github.com/pingcap/tics/pull/1342)                                |   | 7 | 48d2h |
| tidb/23020 | [expression: Add warning info for exprs that can not be pushed to storage layer (#22713)](https://github.com/pingcap/tidb/pull/23020) |   | 7 | 0h    |


</details> 


<details> 
  <summary>ichn-hu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                            PR                                                            | C | LASTED |
|------------|--------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/21853 | [expression: fix compatibility behaviors in time_format with MySQL (#21559)](https://github.com/pingcap/tidb/pull/21853) |   | 80d19h |


## Reviewed in Last 7 Days

|    REPO    |                                                       PR                                                       | C | D |   R   |
|------------|----------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/23056 | [MPP: Kill mpp queries](https://github.com/pingcap/tidb/pull/23056)                                            |   | 3 | 2d21h |
| tidb/23072 | [executor: track memory usage of map in agg partial result.](https://github.com/pingcap/tidb/pull/23072)       |   | 4 | 19h   |
| tidb/23034 | [executor: make the memory tracker of groupConcat more accurate.](https://github.com/pingcap/tidb/pull/23034)  |   | 6 | 23h   |
| tidb/23024 | [executor: make the memory tracker of Jsonobjectagg more accurate](https://github.com/pingcap/tidb/pull/23024) |   | 6 | 1d1h  |


</details> 


<details> 
  <summary>lzmhhh123's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                             PR                                                                              | C | LASTED  |
|------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19347 | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347)                        |   | 198d23h |
| tidb/20444 | [expression: add json_merge_patch](https://github.com/pingcap/tidb/pull/20444)                                                                              |   | 144d21h |
| tidb/20465 | [expression: add uuidShortFunction](https://github.com/pingcap/tidb/pull/20465)                                                                             |   | 143d19h |
| tidb/20642 | [executor: modify admin executors to support partitioned table with global index](https://github.com/pingcap/tidb/pull/20642)                               |   | 132d15h |
| tidb/20903 | [planner: fix confused and unnecessary double-projection in plans.](https://github.com/pingcap/tidb/pull/20903)                                             |   | 121d17h |
| tidb/21018 | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                              |   | 115d17h |
| tidb/21195 | [brie: integrate lightning to suport IMPORT statement](https://github.com/pingcap/tidb/pull/21195)                                                          |   | 104d22h |
| tidb/21334 | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                                               |   | 101d14h |
| tidb/21347 | [session: make rollback work on global variables](https://github.com/pingcap/tidb/pull/21347)                                                               |   | 100d19h |
| tidb/21401 | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                                          |   | 97d11h  |
| tidb/21444 | [planner: ignore anonymous index while tiflash replica is available](https://github.com/pingcap/tidb/pull/21444)                                            |   | 95d12h  |
| tidb/21487 | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                                               |   | 94d4h   |
| tidb/21641 | [executor: Fix pessimistic lock doesn't work on the partition table for subquery/joins](https://github.com/pingcap/tidb/pull/21641)                         |   | 87d18h  |
| tidb/21651 | [planner: allow filter condition pushing down to IndexScan for prefix index](https://github.com/pingcap/tidb/pull/21651)                                    |   | 87d13h  |
| tidb/22126 | [*: add `sys` schema, `sys.SCHEMA_UNUSED_INDEXES` view and `sys.SCHEMA_INDEX_USAGE` view](https://github.com/pingcap/tidb/pull/22126)                       |   | 66d19h  |
| tidb/22149 | [session: set process info before building plan (#22101)](https://github.com/pingcap/tidb/pull/22149)                                                       |   | 62d19h  |
| tidb/22188 | [planner: do not use indexMerge when the path only use a single index (#22168)](https://github.com/pingcap/tidb/pull/22188)                                 |   | 61d13h  |
| tidb/22361 | [table: fix insert into _tidb_rowid panic and rebase it if needed (#22062)](https://github.com/pingcap/tidb/pull/22361)                                     |   | 54d20h  |
| tidb/22372 | [executor: fix SelectForUpdate in decorrelated subquery under pessimistic mode](https://github.com/pingcap/tidb/pull/22372)                                 |   | 54d9h   |
| tidb/22478 | [planner, executor: fix query partition table with global unique index get wrong result](https://github.com/pingcap/tidb/pull/22478)                        |   | 45d13h  |
| tidb/22631 | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                                                                             |   | 35d23h  |
| tidb/22699 | [brie: add error info column and history backup/restore info in sql](https://github.com/pingcap/tidb/pull/22699)                                            |   | 32d16h  |
| tidb/22812 | [ executor: add new format specifier(%# %@ %.) for str_to_date expression (#22790)](https://github.com/pingcap/tidb/pull/22812)                             |   | 16d19h  |
| tidb/22834 | [executor, server: load_data.go is changed and add unit test](https://github.com/pingcap/tidb/pull/22834)                                                   |   | 16d0h   |
| tidb/22857 | [mocktikv: split rpcHandler to kvHandler and coprHandler](https://github.com/pingcap/tidb/pull/22857)                                                       |   | 14d21h  |
| tidb/22910 | [util: optimize the performance of restore with db](https://github.com/pingcap/tidb/pull/22910)                                                             |   | 11d19h  |
| tidb/22974 | [parser: update parser to fix sql restore bug used in create view](https://github.com/pingcap/tidb/pull/22974)                                              |   | 9d17h   |
| tidb/23001 | [statistics: fix err check](https://github.com/pingcap/tidb/pull/23001)                                                                                     |   | 8d0h    |
| tidb/23022 | [executor: create PipelinedWindowExec based on current implementation and modify the windowProcessor interface](https://github.com/pingcap/tidb/pull/23022) |   | 6d18h   |
| tidb/23040 | [ddl: add truncate partition all support](https://github.com/pingcap/tidb/pull/23040)                                                                       |   | 6d13h   |
| tidb/23091 | [*: fix wrong replace or insert-on-dup behavior on prefixed clustered index](https://github.com/pingcap/tidb/pull/23091)                                    |   | 4d12h   |
| tidb/23094 | [planner: fix indexJoin(also hash, merge) on prefixed clustered index](https://github.com/pingcap/tidb/pull/23094)                                          |   | 4d9h    |
| tidb/23134 | [store/tikv:remove set/delete option from kv.Transaction](https://github.com/pingcap/tidb/pull/23134)                                                       |   | 2d18h   |
| tidb/23135 | [executor: fix unexpected NotNullFlag in case when expr ret type (#23102)](https://github.com/pingcap/tidb/pull/23135)                                      |   | 2d17h   |
| tidb/23149 | [core: support left join and right join for join reorder](https://github.com/pingcap/tidb/pull/23149)                                                       |   | 1d12h   |


## Reviewed in Last 7 Days

|     REPO     |                                                                PR                                                                 | C | D |   R   |
|--------------|-----------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/23087   | [executor: fix correlated column range in table reader for the clustered index table](https://github.com/pingcap/tidb/pull/23087) |   | 4 | 17h   |
| tidb/22965   | [executor: introduce setWithMemoryUsage to track set memory in AggExec.](https://github.com/pingcap/tidb/pull/22965)              |   | 5 | 4d23h |
| tikv/9383    | [copr: support decoding the new index format](https://github.com/tikv/tikv/pull/9383)                                             |   | 5 | 64d0h |
| tidb/22980   | [planner: choose non-prefix column when both index key and handle have the same one](https://github.com/pingcap/tidb/pull/22980)  |   | 7 | 2d18h |
| docs-cn/5561 | [Add sql optimization-related docs to toc](https://github.com/pingcap/docs-cn/pull/5561)                                          |   | 7 | 6d20h |
| tidb/22991   | [executor: fix err check](https://github.com/pingcap/tidb/pull/22991)                                                             |   | 7 | 1d5h  |
| tidb/22940   | [planner: enable column pruning for common handle](https://github.com/pingcap/tidb/pull/22940)                                    |   | 7 | 3d22h |


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
| docs-cn/5484 | [system variable: add tidb_enable_engine_fallback](https://github.com/pingcap/docs-cn/pull/5484)                                                            |   | 32d17h  |
| tidb/19029   | [types: fix unexpected NOT_NULL flags](https://github.com/pingcap/tidb/pull/19029)                                                                          |   | 213d22h |
| docs-cn/5561 | [Add sql optimization-related docs to toc](https://github.com/pingcap/docs-cn/pull/5561)                                                                    |   | 13d15h  |
| tidb/20708   | [*: separate auto_increment ID allocator from _tidb_rowid allocator](https://github.com/pingcap/tidb/pull/20708)                                            |   | 129d20h |
| tidb/20969   | [executor: Improve the performance of appending not fixed columns](https://github.com/pingcap/tidb/pull/20969)                                              |   | 117d9h  |
| tidb/20972   | [expression: POC implementation of Vitess hashing algorithm.](https://github.com/pingcap/tidb/pull/20972)                                                   |   | 117d1h  |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                              |   | 115d17h |
| tidb/21149   | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                               |   | 108d14h |
| tidb/21304   | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                                    |   | 102d12h |
| tidb/21318   | [planner, expression: use the range of column types to simplify expressions](https://github.com/pingcap/tidb/pull/21318)                                    |   | 101d19h |
| tidb/21401   | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                                          |   | 97d11h  |
| tidb/21476   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                       |   | 94d15h  |
| tidb/21508   | [execution: fix dayofweek('0000-00-00') behavior](https://github.com/pingcap/tidb/pull/21508)                                                               |   | 93d10h  |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                                 |   | 79d19h  |
| tidb/21887   | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                                             |   | 78d11h  |
| tidb/21930   | [planner: propagate NDV of column groups across plan nodes (#17854)](https://github.com/pingcap/tidb/pull/21930)                                            |   | 75d18h  |
| tidb/21954   | [planner/cascades: add rule `PushSelDownApply`](https://github.com/pingcap/tidb/pull/21954)                                                                 |   | 74d23h  |
| tidb/22090   | [planner: push aggregation operators down to projection and union by default](https://github.com/pingcap/tidb/pull/22090)                                   |   | 67d22h  |
| tidb/22146   | [executor: forbid SFU on view](https://github.com/pingcap/tidb/pull/22146)                                                                                  |   | 62d21h  |
| tidb/22217   | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                                               |   | 60d17h  |
| tidb/22234   | [executor, planner: ON DUPLICATE UPDATE can refer to un-project col (#14412)](https://github.com/pingcap/tidb/pull/22234)                                   |   | 60d15h  |
| tidb/22261   | [time: fix parse datetime won't truncate the reluctant string (#22232)](https://github.com/pingcap/tidb/pull/22261)                                         |   | 59d19h  |
| tidb/22307   | [ddl: fix update can see columns not public](https://github.com/pingcap/tidb/pull/22307)                                                                    |   | 58d16h  |
| tidb/22342   | [session: fix two cases when updating bind info (#22338)](https://github.com/pingcap/tidb/pull/22342)                                                       |   | 55d18h  |
| tidb/22369   | [session: fix the duplicate binding case when updating bind info (#22367)](https://github.com/pingcap/tidb/pull/22369)                                      |   | 54d17h  |
| tidb/22374   | [expression: separated arithmeticIntDivideSig](https://github.com/pingcap/tidb/pull/22374)                                                                  |   | 54d0h   |
| tidb/22415   | [ddl: refactor placement package](https://github.com/pingcap/tidb/pull/22415)                                                                               |   | 50d17h  |
| tidb/22507   | [types: fix the bug about the wrong query result for decimal type ](https://github.com/pingcap/tidb/pull/22507)                                             |   | 41d23h  |
| tidb/22541   | [expression: Support builtin function SOUNDEX](https://github.com/pingcap/tidb/pull/22541)                                                                  |   | 40d9h   |
| tidb/22559   | [planner: split test data from test cases in cbo_test.go](https://github.com/pingcap/tidb/pull/22559)                                                       |   | 39d19h  |
| tidb/22565   | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)                                          | Y | 39d17h  |
| tidb/22733   | [bindinfo: use new sql apis (#22653)](https://github.com/pingcap/tidb/pull/22733)                                                                           |   | 31d15h  |
| tidb/22778   | [*: add support for dynamic privileges](https://github.com/pingcap/tidb/pull/22778)                                                                         |   | 19d7h   |
| tidb/22814   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22814)                                         |   | 16d19h  |
| tidb/22815   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22815)                                         |   | 16d19h  |
| tidb/22862   | [brie: fix the problem that ddl restored by BR via SQL is not replicated to downstream](https://github.com/pingcap/tidb/pull/22862)                         |   | 13d22h  |
| tidb/22867   | [allow pushdown count distinct when enumerate physical plans](https://github.com/pingcap/tidb/pull/22867)                                                   |   | 13d17h  |
| tidb/22869   | [executor: fix cast function will ignore tht error for point-get key construction](https://github.com/pingcap/tidb/pull/22869)                              |   | 13d16h  |
| tidb/22886   | [*: rename tidb_enable_tiflash_fallback_tikv to tidb_enable_engine_fallback](https://github.com/pingcap/tidb/pull/22886)                                    |   | 12d20h  |
| tidb/22915   | [planner: build correct MaxOneRow info from multi-column conditions](https://github.com/pingcap/tidb/pull/22915)                                            |   | 11d17h  |
| tidb/22923   | [expression: correct constant propagation for collation (#22666)](https://github.com/pingcap/tidb/pull/22923)                                               |   | 11d15h  |
| tidb/22924   | [planner: fix wrong index merge selection (#22825)](https://github.com/pingcap/tidb/pull/22924)                                                             |   | 11d14h  |
| tidb/22926   | [expression: add overflow check in multiplyInt](https://github.com/pingcap/tidb/pull/22926)                                                                 |   | 11d13h  |
| tidb/22984   | [executor: fix logging format of prepared statements (#16062)](https://github.com/pingcap/tidb/pull/22984)                                                  |   | 8d10h   |
| tidb/23022   | [executor: create PipelinedWindowExec based on current implementation and modify the windowProcessor interface](https://github.com/pingcap/tidb/pull/23022) |   | 6d18h   |
| tidb/23062   | [*: fix structcheck lint warnings](https://github.com/pingcap/tidb/pull/23062)                                                                              |   | 4d19h   |
| tidb/23074   | [planner: fix range partition prune bug for IN expr (#22894) (#22938)](https://github.com/pingcap/tidb/pull/23074)                                          |   | 4d17h   |
| tidb/23088   | [statistics: delete extended stats cache item in current tidb synchronously](https://github.com/pingcap/tidb/pull/23088)                                    |   | 4d14h   |
| tidb/23105   | [executor: fix wrong key range of index scan when filter is comparing year column with NULL (#23079)](https://github.com/pingcap/tidb/pull/23105)           |   | 3d18h   |
| tidb/23119   | [statistics: remove existing deleted extended stats when add a new one](https://github.com/pingcap/tidb/pull/23119)                                         |   | 3d14h   |
| tidb/23132   | [planner: fix wrong table filters for index merge plan](https://github.com/pingcap/tidb/pull/23132)                                                         |   | 2d19h   |
| tidb/23133   | [plan/core: support mpp group by expressions.](https://github.com/pingcap/tidb/pull/23133)                                                                  |   | 2d19h   |
| tidb/23142   | [brie/: add GetVersion function for tidbGlueSession (#22731)](https://github.com/pingcap/tidb/pull/23142)                                                   |   | 2d14h   |
| tidb/23143   | [brie/: add GetVersion function for tidbGlueSession (#22731)](https://github.com/pingcap/tidb/pull/23143)                                                   |   | 2d14h   |
| tidb/23152   | [expression: fix wrong error info (#22760)](https://github.com/pingcap/tidb/pull/23152)                                                                     |   | 14h     |


## Reviewed in Last 7 Days

|      REPO      |                                                              PR                                                               | C | D |   R    |
|----------------|-------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/22845     | [planner: fix bug of mpp wrongly set schema of exchanger](https://github.com/pingcap/tidb/pull/22845)                         |   | 3 | 12d22h |
| tidb/23113     | [statistics: add tests for `analyze` with dynamic partition prune mode](https://github.com/pingcap/tidb/pull/23113)           |   | 3 | 17h    |
| tidb/22662     | [planner/core: let mpp support partition tables](https://github.com/pingcap/tidb/pull/22662)                                  |   | 4 | 30d4h  |
| tidb/23075     | [statistics: fix merge idx hist with poped topn](https://github.com/pingcap/tidb/pull/23075)                                  |   | 5 | 0h     |
| tidb/22803     | [store/mockstore/unistore: refine and add more mpp tests](https://github.com/pingcap/tidb/pull/22803)                         |   | 5 | 12d1h  |
| tidb/23057     | [statistics: fix the correlation estimation for version 2](https://github.com/pingcap/tidb/pull/23057)                        |   | 5 | 11h    |
| tidb/22910     | [util: optimize the performance of restore with db](https://github.com/pingcap/tidb/pull/22910)                               |   | 5 | 6d20h  |
| tidb/23047     | [planner, expression: fix error when using IN combined with subquery (#22080)](https://github.com/pingcap/tidb/pull/23047)    |   | 5 | 19h    |
| tidb/23049     | [statistics: add more test cases for global-level stats](https://github.com/pingcap/tidb/pull/23049)                          |   | 6 | 3h     |
| tidb/22625     | [planner, statistics: allow (auto) analyze single partition in dynamic-only mode](https://github.com/pingcap/tidb/pull/22625) |   | 7 | 30d20h |
| tidb/22931     | [statistics: enables global-level stats to be generated in fast analyze mode](https://github.com/pingcap/tidb/pull/22931)     |   | 7 | 4d5h   |
| tidb/22489     | [infoschema: support query partition_id from infoschema.partitions (#22240)](https://github.com/pingcap/tidb/pull/22489)      |   | 7 | 38d0h  |
| tidb/23016     | [statistics: add static partition prune mode checks](https://github.com/pingcap/tidb/pull/23016)                              |   | 7 | 2h     |
| tidb-test/1163 | [*: rename partition_prune_mode](https://github.com/pingcap/tidb-test/pull/1163)                                              |   | 7 | 6d0h   |


</details> 


<details> 
  <summary>rebelice's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                               PR                                                               | C | LASTED |
|------------|--------------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/22869 | [executor: fix cast function will ignore tht error for point-get key construction](https://github.com/pingcap/tidb/pull/22869) |   | 13d16h |
| tidb/23074 | [planner: fix range partition prune bug for IN expr (#22894) (#22938)](https://github.com/pingcap/tidb/pull/23074)             |   | 4d17h  |


## Reviewed in Last 7 Days

|    REPO     |                                                                        PR                                                                        | C | D |   R    |
|-------------|--------------------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/20360  | [planner: refine explain info for batch cop](https://github.com/pingcap/tidb/pull/20360)                                                         |   | 3 | 147d7h |
| tidb/23138  | [statistics: add some test cases of global-stats to cover more column types](https://github.com/pingcap/tidb/pull/23138)                         |   | 3 | 0h     |
| tidb/23099  | [statistics: add more test cases about global-stats and fix some issues](https://github.com/pingcap/tidb/pull/23099)                             |   | 3 | 23h    |
| tidb/23066  | [statistics: forbid getting global-stats through fast analyze and incremental analyze](https://github.com/pingcap/tidb/pull/23066)               |   | 3 | 1d19h  |
| tidb/23049  | [statistics: add more test cases for global-level stats](https://github.com/pingcap/tidb/pull/23049)                                             |   | 4 | 1d23h  |
| tidb/23061  | [statistics: introduce a new kind of syntax to drop global-stats](https://github.com/pingcap/tidb/pull/23061)                                    |   | 5 | 0h     |
| parser/1181 | [statistics: introduce a new kind of syntax to drop global-stats](https://github.com/pingcap/parser/pull/1181)                                   |   | 5 | 0h     |
| tidb/23050  | [statistics: forbid global-stats in stats-ver1 and make analyze options take effect in global-stats](https://github.com/pingcap/tidb/pull/23050) |   | 6 | 4h     |
| tidb/23023  | [statistics: support dropping partition/global level statistics](https://github.com/pingcap/tidb/pull/23023)                                     |   | 6 | 21h    |


</details> 


<details> 
  <summary>time-and-fate's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                            PR                                                            | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                |   | 122d16h |
| tidb/22853 | [planner: fix LogicalPlans that contain Window Function are ambiguous ](https://github.com/pingcap/tidb/pull/22853)      |   | 15d12h  |
| tidb/23088 | [statistics: delete extended stats cache item in current tidb synchronously](https://github.com/pingcap/tidb/pull/23088) |   | 4d14h   |
| tidb/23119 | [statistics: remove existing deleted extended stats when add a new one](https://github.com/pingcap/tidb/pull/23119)      |   | 3d14h   |
| tidb/23141 | [planner: fix panic when building index merge plan](https://github.com/pingcap/tidb/pull/23141)                          |   | 2d16h   |


## Reviewed in Last 7 Days

|    REPO    |                                                              PR                                                               | C | D |  R  |
|------------|-------------------------------------------------------------------------------------------------------------------------------|---|---|-----|
| tidb/23132 | [planner: fix wrong table filters for index merge plan](https://github.com/pingcap/tidb/pull/23132)                           |   | 3 | 0h  |
| tidb/23089 | [statistics: report error for extended statistics register on partitioned tables](https://github.com/pingcap/tidb/pull/23089) |   | 4 | 19h |
| tidb/23086 | [statistics: make exponential backoff estimation more safe](https://github.com/pingcap/tidb/pull/23086)                       |   | 5 | 1h  |
| tidb/23052 | [*: support `show stats_extended` to inspect extended stats cache](https://github.com/pingcap/tidb/pull/23052)                |   | 5 | 20h |
| tidb/23057 | [statistics: fix the correlation estimation for version 2](https://github.com/pingcap/tidb/pull/23057)                        |   | 5 | 11h |


</details> 


<details> 
  <summary>winoros's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                             PR                                                              | C | LASTED  |
|--------------|-----------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                      | Y | 177d14h |
| docs-cn/5484 | [system variable: add tidb_enable_engine_fallback](https://github.com/pingcap/docs-cn/pull/5484)                            |   | 32d17h  |
| tidb/20311   | [expression: fix overflow error when convert bit to int64 (#20266)](https://github.com/pingcap/tidb/pull/20311)             |   | 158d21h |
| tidb/20765   | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                           |   | 125d17h |
| tidb/20877   | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                   |   | 122d16h |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)              |   | 115d17h |
| tidb/21207   | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)             |   | 104d19h |
| tidb/21476   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                       |   | 94d15h  |
| tidb/21487   | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                               |   | 94d4h   |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876) |   | 79d19h  |
| tidb/21930   | [planner: propagate NDV of column groups across plan nodes (#17854)](https://github.com/pingcap/tidb/pull/21930)            |   | 75d18h  |
| tidb/21954   | [planner/cascades: add rule `PushSelDownApply`](https://github.com/pingcap/tidb/pull/21954)                                 |   | 74d23h  |
| tidb/22365   | [planner: check index valid while forUpdateRead (#22152)](https://github.com/pingcap/tidb/pull/22365)                       |   | 54d19h  |
| tidb/22504   | [*:Fix the fetchHotRegion bug that the count always zero](https://github.com/pingcap/tidb/pull/22504)                       |   | 42d19h  |
| tidb/22565   | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)          | Y | 39d17h  |
| tidb/22624   | [ planner: not pruning column used by union scan condition (#21640)](https://github.com/pingcap/tidb/pull/22624)            |   | 37d17h  |
| tidb/22804   | [session, util: update session to use new APIs (#22652)](https://github.com/pingcap/tidb/pull/22804)                        |   | 16d21h  |
| tidb/22830   | [planner: fix incorrect duration between compare](https://github.com/pingcap/tidb/pull/22830)                               |   | 16d10h  |
| tidb/22867   | [allow pushdown count distinct when enumerate physical plans](https://github.com/pingcap/tidb/pull/22867)                   |   | 13d17h  |
| tidb/22923   | [expression: correct constant propagation for collation (#22666)](https://github.com/pingcap/tidb/pull/22923)               |   | 11d15h  |
| tidb/22971   | [Privileges: fix delete privilege check wrongly](https://github.com/pingcap/tidb/pull/22971)                                |   | 9d19h   |
| tidb/23088   | [statistics: delete extended stats cache item in current tidb synchronously](https://github.com/pingcap/tidb/pull/23088)    |   | 4d14h   |
| tidb/23092   | [*: fix a bug that collation is not handle for text type (#23045)](https://github.com/pingcap/tidb/pull/23092)              |   | 4d12h   |
| tidb/23094   | [planner: fix indexJoin(also hash, merge) on prefixed clustered index](https://github.com/pingcap/tidb/pull/23094)          |   | 4d9h    |
| tidb/23117   | [planner: fix linter --enable=deadcode check error](https://github.com/pingcap/tidb/pull/23117)                             |   | 3d15h   |
| tidb/23132   | [planner: fix wrong table filters for index merge plan](https://github.com/pingcap/tidb/pull/23132)                         |   | 2d19h   |
| tidb/23133   | [plan/core: support mpp group by expressions.](https://github.com/pingcap/tidb/pull/23133)                                  |   | 2d19h   |
| tidb/23141   | [planner: fix panic when building index merge plan](https://github.com/pingcap/tidb/pull/23141)                             |   | 2d16h   |


## Reviewed in Last 7 Days

|    REPO    |                                                                PR                                                                 | C | D |   R    |
|------------|-----------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/23123 | [planner: show cast type in EXPLAIN in coptask](https://github.com/pingcap/tidb/pull/23123)                                       |   | 3 | 17h    |
| tidb/23122 | [test: make test TestUpdateWithTableReadLockWillFail stable](https://github.com/pingcap/tidb/pull/23122)                          |   | 3 | 16h    |
| tidb/23078 | [store/copr: polish the tiflash-tikv fallback function.](https://github.com/pingcap/tidb/pull/23078)                              |   | 4 | 23h    |
| tidb/23087 | [executor: fix correlated column range in table reader for the clustered index table](https://github.com/pingcap/tidb/pull/23087) |   | 5 | 0h     |
| tidb/23045 | [*: fix a bug that collation is not handle for text type](https://github.com/pingcap/tidb/pull/23045)                             |   | 5 | 1d1h   |
| tidb/22940 | [planner: enable column pruning for common handle](https://github.com/pingcap/tidb/pull/22940)                                    |   | 5 | 5d22h  |
| tidb/22915 | [planner: build correct MaxOneRow info from multi-column conditions](https://github.com/pingcap/tidb/pull/22915)                  |   | 6 | 6d10h  |
| tidb/22914 | [partition: fix hash partition with not between condition get wrong result](https://github.com/pingcap/tidb/pull/22914)           |   | 6 | 6d10h  |
| tidb/22924 | [planner: fix wrong index merge selection (#22825)](https://github.com/pingcap/tidb/pull/22924)                                   |   | 6 | 6d6h   |
| tidb/23029 | [*: check `tidb_enable_extended_stats` in analyze and row count estimation](https://github.com/pingcap/tidb/pull/23029)           |   | 6 | 1d9h   |
| tidb/23041 | [util: remove unused code](https://github.com/pingcap/tidb/pull/23041)                                                            |   | 6 | 1d1h   |
| tidb/23052 | [*: support `show stats_extended` to inspect extended stats cache](https://github.com/pingcap/tidb/pull/23052)                    |   | 6 | 6h     |
| tidb/22090 | [planner: push aggregation operators down to projection and union by default](https://github.com/pingcap/tidb/pull/22090)         |   | 6 | 62d14h |
| tidb/22853 | [planner: fix LogicalPlans that contain Window Function are ambiguous ](https://github.com/pingcap/tidb/pull/22853)               |   | 6 | 9d20h  |
| tidb/22886 | [*: rename tidb_enable_tiflash_fallback_tikv to tidb_enable_engine_fallback](https://github.com/pingcap/tidb/pull/22886)          |   | 7 | 6d1h   |


</details> 


<details> 
  <summary>wshwsh12's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                        PR                                                                         | C | LASTED  |
|------------|---------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19557 | [*: Integrate timeline tracing with TiKV](https://github.com/pingcap/tidb/pull/19557)                                                             |   | 191d23h |
| tidb/19807 | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                                           |   | 184d10h |
| tidb/19957 | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                            | Y | 177d14h |
| tidb/21487 | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                                     |   | 94d4h   |
| tidb/21887 | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                                   |   | 78d11h  |
| tidb/22269 | [executor: check storage.block-cache.capacity value](https://github.com/pingcap/tidb/pull/22269)                                                  |   | 59d17h  |
| tidb/22378 | [executor: vectorize hash aggregate](https://github.com/pingcap/tidb/pull/22378)                                                                  |   | 53d19h  |
| tidb/22382 | [*: add infoschema client errors](https://github.com/pingcap/tidb/pull/22382)                                                                     |   | 53d5h   |
| tidb/22628 | [executor: Improve max/min window function with deque-based sliding window](https://github.com/pingcap/tidb/pull/22628)                           |   | 36d23h  |
| tidb/22815 | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22815)                               |   | 16d19h  |
| tidb/23097 | [util: remove unused code #23013](https://github.com/pingcap/tidb/pull/23097)                                                                     |   | 3d21h   |
| tidb/23104 | [executor: fix wrong key range of index scan when filter is comparing year column with NULL (#23079)](https://github.com/pingcap/tidb/pull/23104) |   | 3d18h   |
| tidb/23105 | [executor: fix wrong key range of index scan when filter is comparing year column with NULL (#23079)](https://github.com/pingcap/tidb/pull/23105) |   | 3d18h   |
| tidb/23109 | [store/tikv:move set/delete option into tikv.Snapshot](https://github.com/pingcap/tidb/pull/23109)                                                |   | 3d17h   |
| tidb/23123 | [planner: show cast type in EXPLAIN in coptask](https://github.com/pingcap/tidb/pull/23123)                                                       |   | 3d13h   |
| tidb/23128 | [statistics: refactor the statistics package use the RestrictedSQLExecutor API (#22636)](https://github.com/pingcap/tidb/pull/23128)              |   | 2d22h   |
| tidb/23142 | [brie/: add GetVersion function for tidbGlueSession (#22731)](https://github.com/pingcap/tidb/pull/23142)                                         |   | 2d14h   |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 

