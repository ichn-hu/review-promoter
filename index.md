|   REVIEWER    |  O(C)   | 1 | 2 | 3 | 4 | 5 | 6 | 7 | T |
|---------------|---------|---|---|---|---|---|---|---|---|
| Reminiscent   |  3 ( 0) | 0 | 1 | 1 | 0 | 0 | 0 | 0 | 2 |
| SunRunAway    | 24 ( 1) | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| XuHuaiyu      | 50 ( 3) | 0 | 0 | 5 | 0 | 0 | 0 | 0 | 5 |
| dyzsr         |  0 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| eurekaka      | 15 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| fzhedu        |  0 ( 0) | 0 | 1 | 1 | 0 | 0 | 0 | 0 | 2 |
| ichn-hu       |  2 ( 0) | 0 | 0 | 5 | 0 | 0 | 0 | 0 | 5 |
| lzmhhh123     | 40 ( 0) | 0 | 2 | 3 | 1 | 0 | 0 | 0 | 6 |
| nullnotnil    |  0 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| qw4990        | 48 ( 2) | 0 | 2 | 3 | 0 | 0 | 0 | 0 | 5 |
| rebelice      |  0 ( 0) | 0 | 0 | 1 | 0 | 0 | 0 | 0 | 1 |
| time-and-fate |  2 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| winoros       | 23 ( 2) | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| wshwsh12      | 19 ( 1) | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |


<details> 
  <summary>Reminiscent's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                   PR                                                                   | C | LASTED |
|------------|----------------------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/21896 | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                             |   | 62d19h |
| tidb/22354 | [planner: do not cache prepared plan if optimization depends on mutable constant (#22349)](https://github.com/pingcap/tidb/pull/22354) |   | 40d23h |
| tidb/22848 | [statistics: merge poped topn when generating the global histogram](https://github.com/pingcap/tidb/pull/22848)                        |   | 1d15h  |


## Reviewed in Last 7 Days

|    REPO    |                                                             PR                                                             | C | D |   R    |
|------------|----------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/22833 | [planner: fix where condition index out of range](https://github.com/pingcap/tidb/pull/22833)                              |   | 2 | 1h     |
| tidb/22603 | [statistics: merge the partition-level histograms to a global-level histogram](https://github.com/pingcap/tidb/pull/22603) |   | 3 | 21d16h |


</details> 


<details> 
  <summary>SunRunAway's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                  PR                                                                   | C | LASTED  |
|--------------|---------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19178   | [executor: Refactor probe channel](https://github.com/pingcap/tidb/pull/19178)                                                        |   | 192d16h |
| docs-cn/4669 | [sql-optimization: extended statistics documentation](https://github.com/pingcap/docs-cn/pull/4669)                                   |   | 131d17h |
| tidb/19347   | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347)  |   | 184d23h |
| tidb/19807   | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                               |   | 170d10h |
| tidb/19900   | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                        | Y | 165d18h |
| tidb/20140   | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                    |   | 152d22h |
| tidb/20220   | [*: new secondary index value format](https://github.com/pingcap/tidb/pull/20220)                                                     |   | 149d16h |
| tidb/20360   | [planner: refine explain info for batch cop](https://github.com/pingcap/tidb/pull/20360)                                              |   | 135d22h |
| tidb/20765   | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                     |   | 111d16h |
| tidb/21207   | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                       |   | 90d19h  |
| tidb/21277   | [executor: fix split table with large integers](https://github.com/pingcap/tidb/pull/21277)                                           |   | 88d19h  |
| tidb/21381   | [*: optimize analyze cluster index table](https://github.com/pingcap/tidb/pull/21381)                                                 |   | 83d17h  |
| tidb/21386   | [expression: Disable cast decimal as string push down to TiFlash](https://github.com/pingcap/tidb/pull/21386)                         |   | 83d16h  |
| tidb/21834   | [planner: enhanced index range calculation plan](https://github.com/pingcap/tidb/pull/21834)                                          |   | 67d18h  |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)           |   | 65d19h  |
| tidb/21878   | [planner: do not push down lock to pointGet/bacthPointGet when selection exists](https://github.com/pingcap/tidb/pull/21878)          |   | 65d18h  |
| tidb/21956   | [planner/preprocessor: disallow into-outfile clause in some place](https://github.com/pingcap/tidb/pull/21956)                        |   | 60d23h  |
| tidb/22026   | [expression: separated arithmeticPlusIntSig](https://github.com/pingcap/tidb/pull/22026)                                              |   | 58d20h  |
| tidb/22043   | [planner, executor: enhance the limit pushdown rule.](https://github.com/pingcap/tidb/pull/22043)                                     |   | 56d10h  |
| tidb/22114   | [test: fix globalkilltest (#21987)](https://github.com/pingcap/tidb/pull/22114)                                                       |   | 53d12h  |
| tidb/22181   | [planner, expression: fix error when using IN combined with subquery (#22080)](https://github.com/pingcap/tidb/pull/22181)            |   | 47d17h  |
| tidb/22217   | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                         |   | 46d17h  |
| tidb/22365   | [planner: check index valid while forUpdateRead (#22152)](https://github.com/pingcap/tidb/pull/22365)                                 |   | 40d19h  |
| tidb/22379   | [[experiment] executor: allow aggregation to spill disk when running out of memory quota](https://github.com/pingcap/tidb/pull/22379) |   | 39d19h  |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>XuHuaiyu's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                              PR                                                                              | C | LASTED  |
|--------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19900   | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                                               | Y | 165d18h |
| docs-cn/5323 | [Update parameter type description](https://github.com/pingcap/docs-cn/pull/5323)                                                                            |   | 34d18h  |
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                                       | Y | 163d13h |
| tidb/20040   | [planner, expression: take NullFlag into consideration when optimize the `int non-const` <cmp > `non-int const`](https://github.com/pingcap/tidb/pull/20040) | Y | 158d13h |
| tidb/20140   | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                                           |   | 152d22h |
| tidb/20311   | [expression: fix overflow error when convert bit to int64 (#20266)](https://github.com/pingcap/tidb/pull/20311)                                              |   | 144d21h |
| tidb/20576   | [*: fix stats feedback after tableReader handle multiple ranges](https://github.com/pingcap/tidb/pull/20576)                                                 |   | 123d13h |
| tidb/20790   | [collation: add pinyin collation for chinese charset support](https://github.com/pingcap/tidb/pull/20790)                                                    |   | 110d20h |
| tidb/20905   | [planner: fix statement-optimize not work in `TryFastPlan`](https://github.com/pingcap/tidb/pull/20905)                                                      |   | 107d17h |
| tidb/20972   | [expression: POC implementation of Vitess hashing algorithm.](https://github.com/pingcap/tidb/pull/20972)                                                    |   | 103d0h  |
| tidb/21064   | [planner, executor: fix cast not check error](https://github.com/pingcap/tidb/pull/21064)                                                                    |   | 98d8h   |
| tidb/21149   | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                                |   | 94d14h  |
| tidb/21228   | [executor: return the result immediately when combining LIMIT row_count with DISTINCT](https://github.com/pingcap/tidb/pull/21228)                           |   | 90d13h  |
| tidb/21304   | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                                     |   | 88d12h  |
| tidb/21334   | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                                                |   | 87d14h  |
| tidb/21340   | [executor: initialize expensive query handler on domain creation](https://github.com/pingcap/tidb/pull/21340)                                                |   | 86d23h  |
| tidb/21476   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                        |   | 80d15h  |
| tidb/21536   | [executor: add slow-log file meta cache to avoid repeat read file meta information](https://github.com/pingcap/tidb/pull/21536)                              |   | 76d14h  |
| tidb/21564   | [ddl: fix Incorrect behavior of NO_ZERO_DATE when altering table](https://github.com/pingcap/tidb/pull/21564)                                                |   | 75d15h  |
| tidb/21626   | [test: convert test to benchmard test to make ci stable (#21616)](https://github.com/pingcap/tidb/pull/21626)                                                |   | 73d22h  |
| tidb/21680   | [planner: report error when ORDER BY conflicts with DISTINCT (#21286)](https://github.com/pingcap/tidb/pull/21680)                                           |   | 72d16h  |
| tidb/21853   | [expression: fix compatibility behaviors in time_format with MySQL (#21559)](https://github.com/pingcap/tidb/pull/21853)                                     |   | 66d18h  |
| tidb/21896   | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                                                   |   | 62d19h  |
| tidb/22131   | [privilege: remove leading and trailing space when create user and role](https://github.com/pingcap/tidb/pull/22131)                                         |   | 52d19h  |
| tidb/22149   | [session: set process info before building plan (#22101)](https://github.com/pingcap/tidb/pull/22149)                                                        |   | 48d19h  |
| tidb/22163   | [expression: separated arithmeticMinusIntSig](https://github.com/pingcap/tidb/pull/22163)                                                                    |   | 48d13h  |
| tidb/22186   | [executor: fix select into outfile with year type column has no data (#22175)](https://github.com/pingcap/tidb/pull/22186)                                   |   | 47d16h  |
| tidb/22294   | [planner, table: optimize the list partition pruner for range query](https://github.com/pingcap/tidb/pull/22294)                                             |   | 44d19h  |
| tidb/22307   | [ddl: fix update can see columns not public](https://github.com/pingcap/tidb/pull/22307)                                                                     |   | 44d15h  |
| tidb/22381   | [planner: check schema stale for plan cache when forUpdateRead](https://github.com/pingcap/tidb/pull/22381)                                                  |   | 39d14h  |
| tidb/22418   | [expression: Optimize builtinArithmeticModRealSig and builtinGreatestDecimalSig using MergeNull method](https://github.com/pingcap/tidb/pull/22418)          |   | 35d0h   |
| tidb/22432   | [types,execute: fix errcode return like mysql when inserting incorrect int value ](https://github.com/pingcap/tidb/pull/22432)                               |   | 33d21h  |
| tidb/22507   | [types: fix the bug about the wrong query result for decimal type ](https://github.com/pingcap/tidb/pull/22507)                                              |   | 27d23h  |
| tidb/22616   | [expression: from_unixtime accept 64-bit integers](https://github.com/pingcap/tidb/pull/22616)                                                               |   | 23d23h  |
| tidb/22617   | [metrics: fix wrong bucket name of coprocessor cache (#22454)](https://github.com/pingcap/tidb/pull/22617)                                                   |   | 23d23h  |
| tidb/22624   | [ planner: not pruning column used by union scan condition (#21640)](https://github.com/pingcap/tidb/pull/22624)                                             |   | 23d16h  |
| tidb/22640   | [*: refactor ExecuteInternal to return single resultset (#22546)](https://github.com/pingcap/tidb/pull/22640)                                                |   | 20d20h  |
| tidb/22696   | [expression: enable arithmetic Mod push down](https://github.com/pingcap/tidb/pull/22696)                                                                    |   | 18d17h  |
| tidb/22711   | [executor: Fix inline schema name](https://github.com/pingcap/tidb/pull/22711)                                                                               |   | 18d11h  |
| tidb/22722   | [planner, errno: make error code of ErrMixOfGroupFuncAndFields consistent with MySQL](https://github.com/pingcap/tidb/pull/22722)                            |   | 17d20h  |
| tidb/22736   | [executor: fix load data losing connection when batch_dml_size is set (#22724)](https://github.com/pingcap/tidb/pull/22736)                                  |   | 16d23h  |
| tidb/22784   | [oracle: make @@txn_scope only support local and global](https://github.com/pingcap/tidb/pull/22784)                                                         |   | 3d19h   |
| tidb/22786   | [config: deprecate `tikv-client.copr-cache.enable`](https://github.com/pingcap/tidb/pull/22786)                                                              |   | 3d18h   |
| tidb/22814   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22814)                                          |   | 2d18h   |
| tidb/22815   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22815)                                          |   | 2d18h   |
| tidb/22824   | [*: apply golangci-lint to the new code](https://github.com/pingcap/tidb/pull/22824)                                                                         |   | 2d15h   |
| tidb/22832   | [expression: push down EXTRACT to TiFlash](https://github.com/pingcap/tidb/pull/22832)                                                                       |   | 2d1h    |
| tidb/22833   | [planner: fix where condition index out of range](https://github.com/pingcap/tidb/pull/22833)                                                                |   | 2d1h    |
| tidb/22844   | [expression: do not adjust int when it is null and compared year (#22821)](https://github.com/pingcap/tidb/pull/22844)                                       |   | 1d19h   |
| tidb/22856   | [mockstore: move mockstore/cluster to tikv/mockstore/cluster](https://github.com/pingcap/tidb/pull/22856)                                                    |   | 1d10h   |


## Reviewed in Last 7 Days

|    REPO    |                                                               PR                                                                | C | D |   R   |
|------------|---------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/22426 | [expression: fix bugs in builtinfunction ArithmeticMinusInt logic](https://github.com/pingcap/tidb/pull/22426)                  |   | 3 | 32d0h |
| tidb/22812 | [ executor: add new format specifier(%# %@ %.) for str_to_date expression (#22790)](https://github.com/pingcap/tidb/pull/22812) |   | 3 | 0h    |
| tidb/22790 | [ executor: add new format specifier(%# %@ %.) for str_to_date expression](https://github.com/pingcap/tidb/pull/22790)          |   | 3 | 21h   |
| tidb/22785 | [expression: fix enum and set type expression in where clause](https://github.com/pingcap/tidb/pull/22785)                      |   | 3 | 23h   |
| tidb/22737 | [executor: fix load data losing connection when batch_dml_size is set (#22724)](https://github.com/pingcap/tidb/pull/22737)     |   | 3 | 14d0h |


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
| tidb/19347 | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347)   |   | 184d23h |
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                              |   | 108d16h |
| tidb/21444 | [planner: ignore anonymous index while tiflash replica is available](https://github.com/pingcap/tidb/pull/21444)                       |   | 81d12h  |
| tidb/21680 | [planner: report error when ORDER BY conflicts with DISTINCT (#21286)](https://github.com/pingcap/tidb/pull/21680)                     |   | 72d16h  |
| tidb/21994 | [range: fix overflow value access index ](https://github.com/pingcap/tidb/pull/21994)                                                  |   | 59d22h  |
| tidb/22342 | [session: fix two cases when updating bind info (#22338)](https://github.com/pingcap/tidb/pull/22342)                                  |   | 41d18h  |
| tidb/22354 | [planner: do not cache prepared plan if optimization depends on mutable constant (#22349)](https://github.com/pingcap/tidb/pull/22354) |   | 40d23h  |
| tidb/22369 | [session: fix the duplicate binding case when updating bind info (#22367)](https://github.com/pingcap/tidb/pull/22369)                 |   | 40d17h  |
| tidb/22416 | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                                      |   | 36d11h  |
| tidb/22559 | [planner: split test data from test cases in cbo_test.go](https://github.com/pingcap/tidb/pull/22559)                                  |   | 25d19h  |
| tidb/22733 | [bindinfo: use new sql apis (#22653)](https://github.com/pingcap/tidb/pull/22733)                                                      |   | 17d15h  |
| tidb/22734 | [bindinfo: use new sql apis (#22653)](https://github.com/pingcap/tidb/pull/22734)                                                      |   | 17d15h  |
| tidb/22778 | [*: add support for dynamic privileges](https://github.com/pingcap/tidb/pull/22778)                                                    |   | 5d7h    |
| tidb/22805 | [session, util: update session to use new APIs (#22652)](https://github.com/pingcap/tidb/pull/22805)                                   |   | 2d21h   |
| tidb/22848 | [statistics: merge poped topn when generating the global histogram](https://github.com/pingcap/tidb/pull/22848)                        |   | 1d15h   |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>fzhedu's detailed review report</summary> 

## To Be Reviewed

| REPO | PR | C | LASTED |
|------|----|---|--------|


## Reviewed in Last 7 Days

|    REPO    |                                                              PR                                                              | C | D |   R    |
|------------|------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/22725 | [planner, distsql: fix the behaviour of building ranges for TiFlash](https://github.com/pingcap/tidb/pull/22725)             |   | 2 | 15d23h |
| tidb/22713 | [expression: Add warning info for exprs that can not be pushed to storage layer](https://github.com/pingcap/tidb/pull/22713) |   | 3 | 15d12h |


</details> 


<details> 
  <summary>ichn-hu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                            PR                                                            | C | LASTED |
|------------|--------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/21853 | [expression: fix compatibility behaviors in time_format with MySQL (#21559)](https://github.com/pingcap/tidb/pull/21853) |   | 66d18h |
| tidb/22844 | [expression: do not adjust int when it is null and compared year (#22821)](https://github.com/pingcap/tidb/pull/22844)   |   | 1d19h  |


## Reviewed in Last 7 Days

|    REPO    |                                                         PR                                                          | C | D |   R    |
|------------|---------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/22821 | [expression: do not adjust int when it is null and compared year](https://github.com/pingcap/tidb/pull/22821)       |   | 3 | 0h     |
| tidb/22815 | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22815) |   | 3 | 0h     |
| tidb/22814 | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22814) |   | 3 | 0h     |
| tidb/22785 | [expression: fix enum and set type expression in where clause](https://github.com/pingcap/tidb/pull/22785)          |   | 3 | 23h    |
| tidb/22701 | [expression: refine performance of EXTRACT function](https://github.com/pingcap/tidb/pull/22701)                    |   | 3 | 15d20h |


</details> 


<details> 
  <summary>lzmhhh123's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                  PR                                                                   | C | LASTED  |
|------------|---------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19347 | [executor: support new syntax `create/drop binding for digest` for tidb dashboard usage](https://github.com/pingcap/tidb/pull/19347)  |   | 184d23h |
| tidb/20444 | [expression: add json_merge_patch](https://github.com/pingcap/tidb/pull/20444)                                                        |   | 130d21h |
| tidb/20465 | [expression: add uuidShortFunction](https://github.com/pingcap/tidb/pull/20465)                                                       |   | 129d19h |
| tidb/20642 | [executor: modify admin executors to support partitioned table with global index](https://github.com/pingcap/tidb/pull/20642)         |   | 118d15h |
| tidb/20825 | [executor: add diagnosis rule to check Transparent Huge Pages(THP) enabled (#20611)](https://github.com/pingcap/tidb/pull/20825)      |   | 109d18h |
| tidb/20903 | [planner: fix confused and unnecessary double-projection in plans.](https://github.com/pingcap/tidb/pull/20903)                       |   | 107d17h |
| tidb/21018 | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                        |   | 101d16h |
| tidb/21195 | [brie: integrate lightning to suport IMPORT statement](https://github.com/pingcap/tidb/pull/21195)                                    |   | 90d22h  |
| tidb/21334 | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                         |   | 87d14h  |
| tidb/21347 | [session: make rollback work on global variables](https://github.com/pingcap/tidb/pull/21347)                                         |   | 86d19h  |
| tidb/21401 | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                    |   | 83d11h  |
| tidb/21444 | [planner: ignore anonymous index while tiflash replica is available](https://github.com/pingcap/tidb/pull/21444)                      |   | 81d12h  |
| tidb/21487 | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                         |   | 80d4h   |
| tidb/21641 | [executor: Fix pessimistic lock doesn't work on the partition table for subquery/joins](https://github.com/pingcap/tidb/pull/21641)   |   | 73d18h  |
| tidb/21651 | [planner: allow filter condition pushing down to IndexScan for prefix index](https://github.com/pingcap/tidb/pull/21651)              |   | 73d13h  |
| tidb/21680 | [planner: report error when ORDER BY conflicts with DISTINCT (#21286)](https://github.com/pingcap/tidb/pull/21680)                    |   | 72d16h  |
| tidb/21954 | [planner/cascades: add rule `PushSelDownApply`](https://github.com/pingcap/tidb/pull/21954)                                           |   | 60d23h  |
| tidb/22126 | [*: add `sys` schema, `sys.SCHEMA_UNUSED_INDEXES` view and `sys.SCHEMA_INDEX_USAGE` view](https://github.com/pingcap/tidb/pull/22126) |   | 52d19h  |
| tidb/22149 | [session: set process info before building plan (#22101)](https://github.com/pingcap/tidb/pull/22149)                                 |   | 48d19h  |
| tidb/22188 | [planner: do not use indexMerge when the path only use a single index (#22168)](https://github.com/pingcap/tidb/pull/22188)           |   | 47d13h  |
| tidb/22361 | [table: fix insert into _tidb_rowid panic and rebase it if needed (#22062)](https://github.com/pingcap/tidb/pull/22361)               |   | 40d20h  |
| tidb/22372 | [executor: fix SelectForUpdate in decorrelated subquery under pessimistic mode](https://github.com/pingcap/tidb/pull/22372)           |   | 40d9h   |
| tidb/22430 | [*: refactor table.Table interface, clean up unnecessay methods](https://github.com/pingcap/tidb/pull/22430)                          |   | 33d23h  |
| tidb/22478 | [planner, executor: fix query partition table with global unique index get wrong result](https://github.com/pingcap/tidb/pull/22478)  |   | 31d13h  |
| tidb/22625 | [planner, statistics: allow (auto) analyze single partition in dynamic-only mode](https://github.com/pingcap/tidb/pull/22625)         |   | 23d14h  |
| tidb/22631 | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                                                       |   | 21d22h  |
| tidb/22656 | [*: move new api out of session package (#22591)](https://github.com/pingcap/tidb/pull/22656)                                         |   | 20d0h   |
| tidb/22657 | [*: move new api out of session package (#22591)](https://github.com/pingcap/tidb/pull/22657)                                         |   | 20d0h   |
| tidb/22662 | [planner/core: let mpp support partition tables](https://github.com/pingcap/tidb/pull/22662)                                          |   | 19d18h  |
| tidb/22699 | [brie: add error info column and history backup/restore info in sql](https://github.com/pingcap/tidb/pull/22699)                      |   | 18d16h  |
| tidb/22714 | [executor: add close recordSet in executor](https://github.com/pingcap/tidb/pull/22714)                                               |   | 17d23h  |
| tidb/22728 | [store/tikv_driver:move MemManager from KVStore to tikvStore](https://github.com/pingcap/tidb/pull/22728)                             |   | 17d18h  |
| tidb/22734 | [bindinfo: use new sql apis (#22653)](https://github.com/pingcap/tidb/pull/22734)                                                     |   | 17d15h  |
| tidb/22752 | [server, sessionctx: Support for the status command of MySQL Shell](https://github.com/pingcap/tidb/pull/22752)                       |   | 16d1h   |
| tidb/22812 | [ executor: add new format specifier(%# %@ %.) for str_to_date expression (#22790)](https://github.com/pingcap/tidb/pull/22812)       |   | 2d19h   |
| tidb/22822 | [expression: fixed error msg in builtinArithmeticMinusIntSig](https://github.com/pingcap/tidb/pull/22822)                             |   | 2d16h   |
| tidb/22829 | [planner: let tikv know primary prefix columns info](https://github.com/pingcap/tidb/pull/22829)                                      |   | 2d13h   |
| tidb/22834 | [executor, server: load_data.go is changed and add unit test](https://github.com/pingcap/tidb/pull/22834)                             |   | 2d0h    |
| tidb/22857 | [store: move mockstore/mocktikv to tikv/mockstore/mocktikv](https://github.com/pingcap/tidb/pull/22857)                               |   | 20h     |
| tidb/22860 | [planner: include "CREATE/DROP TEMPORARY TABLE" to noop functions (##609)](https://github.com/pingcap/tidb/pull/22860)                |   | 8h      |


## Reviewed in Last 7 Days

|     REPO     |                                                           PR                                                           | C | D |   R    |
|--------------|------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| docs-cn/4933 | [explain: add joins](https://github.com/pingcap/docs-cn/pull/4933)                                                     |   | 2 | 91d21h |
| docs-cn/4913 | [explain: add indexes](https://github.com/pingcap/docs-cn/pull/4913)                                                   |   | 2 | 95d19h |
| tidb/22790   | [ executor: add new format specifier(%# %@ %.) for str_to_date expression](https://github.com/pingcap/tidb/pull/22790) |   | 3 | 21h    |
| tidb/22701   | [expression: refine performance of EXTRACT function](https://github.com/pingcap/tidb/pull/22701)                       |   | 3 | 15d19h |
| tidb/22426   | [expression: fix bugs in builtinfunction ArithmeticMinusInt logic](https://github.com/pingcap/tidb/pull/22426)         |   | 3 | 31d20h |
| tidb/22463   | [executor: make memory tracker for aggregate more accurate.](https://github.com/pingcap/tidb/pull/22463)               |   | 4 | 28d0h  |


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
| tidb/19029 | [types: fix unexpected NOT_NULL flags](https://github.com/pingcap/tidb/pull/19029)                                                                     |   | 199d22h |
| tidb/20354 | [planner: rename relational operators (#14575)](https://github.com/pingcap/tidb/pull/20354)                                                            | Y | 137d5h  |
| tidb/20708 | [*: separate auto_increment ID allocator from _tidb_rowid allocator](https://github.com/pingcap/tidb/pull/20708)                                       |   | 115d20h |
| tidb/20969 | [executor: Improve the performance of appending not fixed columns](https://github.com/pingcap/tidb/pull/20969)                                         |   | 103d9h  |
| tidb/20972 | [expression: POC implementation of Vitess hashing algorithm.](https://github.com/pingcap/tidb/pull/20972)                                              |   | 103d0h  |
| tidb/21018 | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                         |   | 101d16h |
| tidb/21149 | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                          |   | 94d14h  |
| tidb/21304 | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                               |   | 88d12h  |
| tidb/21318 | [planner, expression: use the range of column types to simplify expressions](https://github.com/pingcap/tidb/pull/21318)                               |   | 87d18h  |
| tidb/21401 | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                                     |   | 83d11h  |
| tidb/21476 | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                  |   | 80d15h  |
| tidb/21508 | [execution: fix dayofweek('0000-00-00') behavior](https://github.com/pingcap/tidb/pull/21508)                                                          |   | 79d9h   |
| tidb/21680 | [planner: report error when ORDER BY conflicts with DISTINCT (#21286)](https://github.com/pingcap/tidb/pull/21680)                                     |   | 72d16h  |
| tidb/21876 | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                            |   | 65d19h  |
| tidb/21887 | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                                        |   | 64d11h  |
| tidb/21930 | [planner: propagate NDV of column groups across plan nodes (#17854)](https://github.com/pingcap/tidb/pull/21930)                                       |   | 61d18h  |
| tidb/21977 | [expression: log functions that can not be pushed to cop](https://github.com/pingcap/tidb/pull/21977)                                                  |   | 60d15h  |
| tidb/22090 | [planner: push aggregation operators down to projection and union by default](https://github.com/pingcap/tidb/pull/22090)                              |   | 53d22h  |
| tidb/22146 | [executor: forbid SFU on view](https://github.com/pingcap/tidb/pull/22146)                                                                             |   | 48d21h  |
| tidb/22217 | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                                          |   | 46d17h  |
| tidb/22234 | [executor, planner: ON DUPLICATE UPDATE can refer to un-project col (#14412)](https://github.com/pingcap/tidb/pull/22234)                              |   | 46d15h  |
| tidb/22261 | [time: fix parse datetime won't truncate the reluctant string (#22232)](https://github.com/pingcap/tidb/pull/22261)                                    |   | 45d19h  |
| tidb/22307 | [ddl: fix update can see columns not public](https://github.com/pingcap/tidb/pull/22307)                                                               |   | 44d15h  |
| tidb/22342 | [session: fix two cases when updating bind info (#22338)](https://github.com/pingcap/tidb/pull/22342)                                                  |   | 41d18h  |
| tidb/22369 | [session: fix the duplicate binding case when updating bind info (#22367)](https://github.com/pingcap/tidb/pull/22369)                                 |   | 40d17h  |
| tidb/22374 | [expression: separated arithmeticIntDivideSig](https://github.com/pingcap/tidb/pull/22374)                                                             |   | 40d0h   |
| tidb/22415 | [ddl: refactor placement package](https://github.com/pingcap/tidb/pull/22415)                                                                          |   | 36d17h  |
| tidb/22456 | [distsql, executor: disable cache during staleness transaction](https://github.com/pingcap/tidb/pull/22456)                                            |   | 32d15h  |
| tidb/22471 | [ddl, executor: fix creating unique index without partition column error when enable-global-index is true](https://github.com/pingcap/tidb/pull/22471) |   | 31d17h  |
| tidb/22489 | [infoschema: support query partition_id from infoschema.partitions (#22240)](https://github.com/pingcap/tidb/pull/22489)                               |   | 30d19h  |
| tidb/22507 | [types: fix the bug about the wrong query result for decimal type ](https://github.com/pingcap/tidb/pull/22507)                                        |   | 27d23h  |
| tidb/22541 | [expression: Support builtin function SOUNDEX](https://github.com/pingcap/tidb/pull/22541)                                                             |   | 26d9h   |
| tidb/22565 | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)                                     | Y | 25d17h  |
| tidb/22641 | [*: do not report error for prepared stmt execution if tidb_snapshot is set (#22568)](https://github.com/pingcap/tidb/pull/22641)                      |   | 20d19h  |
| tidb/22649 | [planner: decorrelate LogicalApply with inner join as the inner child](https://github.com/pingcap/tidb/pull/22649)                                     |   | 20d16h  |
| tidb/22666 | [expression: correct constant propagation for collation](https://github.com/pingcap/tidb/pull/22666)                                                   |   | 19d18h  |
| tidb/22687 | [*: refactor the RestrictedSQLExecutor interface (#22579)](https://github.com/pingcap/tidb/pull/22687)                                                 |   | 18d21h  |
| tidb/22731 | [brie/: add GetVersion function for tidbGlueSession](https://github.com/pingcap/tidb/pull/22731)                                                       |   | 17d17h  |
| tidb/22733 | [bindinfo: use new sql apis (#22653)](https://github.com/pingcap/tidb/pull/22733)                                                                      |   | 17d15h  |
| tidb/22734 | [bindinfo: use new sql apis (#22653)](https://github.com/pingcap/tidb/pull/22734)                                                                      |   | 17d15h  |
| tidb/22760 | [expression: Fix wrong error info](https://github.com/pingcap/tidb/pull/22760)                                                                         |   | 14d0h   |
| tidb/22777 | [expression: Treat non null Enum as bool true](https://github.com/pingcap/tidb/pull/22777)                                                             |   | 5d15h   |
| tidb/22778 | [*: add support for dynamic privileges](https://github.com/pingcap/tidb/pull/22778)                                                                    |   | 5d7h    |
| tidb/22814 | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22814)                                    |   | 2d18h   |
| tidb/22815 | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22815)                                    |   | 2d18h   |
| tidb/22830 | [planner: fix incorrect duration between compare](https://github.com/pingcap/tidb/pull/22830)                                                          |   | 2d9h    |
| tidb/22846 | [stats, planner, sessionctx: handle compatibility between feedback and ver2 stats](https://github.com/pingcap/tidb/pull/22846)                         |   | 1d16h   |
| tidb/22848 | [statistics: merge poped topn when generating the global histogram](https://github.com/pingcap/tidb/pull/22848)                                        |   | 1d15h   |


## Reviewed in Last 7 Days

|    REPO    |                                                               PR                                                                | C | D |   R    |
|------------|---------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/22725 | [planner, distsql: fix the behaviour of building ranges for TiFlash](https://github.com/pingcap/tidb/pull/22725)                |   | 2 | 16d2h  |
| tidb/22682 | [statistics: introduce new estimation logic when index histogram fails to estimate](https://github.com/pingcap/tidb/pull/22682) |   | 2 | 17d15h |
| tidb/22433 | [statistics: merge partition-level TopN to global-level TopN](https://github.com/pingcap/tidb/pull/22433)                       |   | 3 | 31d4h  |
| tidb/22625 | [planner, statistics: allow (auto) analyze single partition in dynamic-only mode](https://github.com/pingcap/tidb/pull/22625)   |   | 3 | 20d22h |
| tidb/22603 | [statistics: merge the partition-level histograms to a global-level histogram](https://github.com/pingcap/tidb/pull/22603)      |   | 3 | 21d16h |


</details> 


<details> 
  <summary>rebelice's detailed review report</summary> 

## To Be Reviewed

| REPO | PR | C | LASTED |
|------|----|---|--------|


## Reviewed in Last 7 Days

|    REPO    |                                                    PR                                                     | C | D |   R   |
|------------|-----------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/22433 | [statistics: merge partition-level TopN to global-level TopN](https://github.com/pingcap/tidb/pull/22433) |   | 3 | 31d3h |


</details> 


<details> 
  <summary>time-and-fate's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                               PR                                                                | C | LASTED  |
|------------|---------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                       |   | 108d16h |
| tidb/22682 | [statistics: introduce new estimation logic when index histogram fails to estimate](https://github.com/pingcap/tidb/pull/22682) |   | 19d8h   |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>winoros's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                               PR                                                               | C | LASTED  |
|--------------|--------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                         | Y | 163d13h |
| docs-cn/4669 | [sql-optimization: extended statistics documentation](https://github.com/pingcap/docs-cn/pull/4669)                            |   | 131d17h |
| tidb/20311   | [expression: fix overflow error when convert bit to int64 (#20266)](https://github.com/pingcap/tidb/pull/20311)                |   | 144d21h |
| tidb/20765   | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                              |   | 111d16h |
| tidb/20877   | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                      |   | 108d16h |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                 |   | 101d16h |
| tidb/21207   | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                |   | 90d19h  |
| tidb/21476   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                          |   | 80d15h  |
| tidb/21487   | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                  |   | 80d4h   |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)    |   | 65d19h  |
| tidb/21930   | [planner: propagate NDV of column groups across plan nodes (#17854)](https://github.com/pingcap/tidb/pull/21930)               |   | 61d18h  |
| tidb/22090   | [planner: push aggregation operators down to projection and union by default](https://github.com/pingcap/tidb/pull/22090)      |   | 53d22h  |
| tidb/22365   | [planner: check index valid while forUpdateRead (#22152)](https://github.com/pingcap/tidb/pull/22365)                          |   | 40d19h  |
| tidb/22489   | [infoschema: support query partition_id from infoschema.partitions (#22240)](https://github.com/pingcap/tidb/pull/22489)       |   | 30d19h  |
| tidb/22504   | [*:Fix the fetchHotRegion bug that the count always zero](https://github.com/pingcap/tidb/pull/22504)                          |   | 28d19h  |
| tidb/22565   | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)             | Y | 25d17h  |
| tidb/22624   | [ planner: not pruning column used by union scan condition (#21640)](https://github.com/pingcap/tidb/pull/22624)               |   | 23d16h  |
| tidb/22804   | [session, util: update session to use new APIs (#22652)](https://github.com/pingcap/tidb/pull/22804)                           |   | 2d21h   |
| tidb/22805   | [session, util: update session to use new APIs (#22652)](https://github.com/pingcap/tidb/pull/22805)                           |   | 2d21h   |
| tidb/22825   | [planner: fix wrong index merge selection](https://github.com/pingcap/tidb/pull/22825)                                         |   | 2d15h   |
| tidb/22830   | [planner: fix incorrect duration between compare](https://github.com/pingcap/tidb/pull/22830)                                  |   | 2d9h    |
| tidb/22845   | [planner: fix bug of mpp wrongly set schema of exchanger](https://github.com/pingcap/tidb/pull/22845)                          |   | 1d17h   |
| tidb/22846   | [stats, planner, sessionctx: handle compatibility between feedback and ver2 stats](https://github.com/pingcap/tidb/pull/22846) |   | 1d16h   |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>wshwsh12's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                           PR                                                            | C | LASTED  |
|------------|-------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19557 | [*: Integrate timeline tracing with TiKV](https://github.com/pingcap/tidb/pull/19557)                                   |   | 177d23h |
| tidb/19807 | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                 |   | 170d10h |
| tidb/19957 | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                  | Y | 163d13h |
| tidb/21381 | [*: optimize analyze cluster index table](https://github.com/pingcap/tidb/pull/21381)                                   |   | 83d17h  |
| tidb/21487 | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                           |   | 80d4h   |
| tidb/21887 | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                         |   | 64d11h  |
| tidb/22269 | [executor: check storage.block-cache.capacity value](https://github.com/pingcap/tidb/pull/22269)                        |   | 45d16h  |
| tidb/22378 | [executor: vectorize hash aggregate](https://github.com/pingcap/tidb/pull/22378)                                        |   | 39d19h  |
| tidb/22382 | [*: add infoschema client errors](https://github.com/pingcap/tidb/pull/22382)                                           |   | 39d5h   |
| tidb/22426 | [expression: fix bugs in builtinfunction ArithmeticMinusInt logic](https://github.com/pingcap/tidb/pull/22426)          |   | 34d16h  |
| tidb/22475 | [*: fix parser error format when value overflow](https://github.com/pingcap/tidb/pull/22475)                            |   | 31d16h  |
| tidb/22628 | [executor: Improve max/min window function with deque-based sliding window](https://github.com/pingcap/tidb/pull/22628) |   | 22d23h  |
| tidb/22748 | [executor, privilege: fix failure on grant USAGE privilege operation](https://github.com/pingcap/tidb/pull/22748)       |   | 16d16h  |
| tidb/22772 | [Executor: Resolve bug where show index from perf_schema had no results](https://github.com/pingcap/tidb/pull/22772)    |   | 6d10h   |
| tidb/22803 | [store/mockstore/unistore: refine and add more mpp tests](https://github.com/pingcap/tidb/pull/22803)                   |   | 2d21h   |
| tidb/22815 | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22815)     |   | 2d18h   |
| tidb/22829 | [planner: let tikv know primary prefix columns info](https://github.com/pingcap/tidb/pull/22829)                        |   | 2d13h   |
| tidb/22836 | [expression: add integration test for partition pruning](https://github.com/pingcap/tidb/pull/22836)                    |   | 2d0h    |
| tidb/22858 | [execution: fix index's duplicate error message](https://github.com/pingcap/tidb/pull/22858)                            |   | 15h     |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 

