|   REVIEWER    |  O(C)   | 1  | 2 | 3 | 4 | 5 | 6 | 7 | T  |
|---------------|---------|----|---|---|---|---|---|---|----|
| Reminiscent   |  6 ( 0) |  2 | 2 | 0 | 0 | 0 | 0 | 4 |  8 |
| SunRunAway    | 12 ( 1) |  0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| XuHuaiyu      | 67 ( 3) |  0 | 3 | 3 | 0 | 0 | 0 | 2 |  8 |
| dyzsr         |  1 ( 0) |  0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| eurekaka      | 22 ( 1) |  0 | 0 | 1 | 1 | 0 | 0 | 2 |  4 |
| fzhedu        |  2 ( 0) |  0 | 2 | 0 | 0 | 0 | 0 | 0 |  2 |
| ichn-hu       |  2 ( 0) |  1 | 2 | 2 | 0 | 0 | 0 | 0 |  5 |
| lzmhhh123     | 54 ( 0) |  0 | 0 | 5 | 0 | 0 | 0 | 3 |  8 |
| nullnotnil    |  0 ( 0) |  0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| qw4990        | 58 ( 1) | 13 | 5 | 5 | 0 | 0 | 0 | 9 | 32 |
| rebelice      |  4 ( 0) |  0 | 0 | 2 | 0 | 0 | 0 | 0 |  2 |
| time-and-fate |  3 ( 0) |  0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| winoros       | 34 ( 2) |  3 | 4 | 4 | 3 | 0 | 0 | 4 | 18 |
| wshwsh12      | 23 ( 1) |  3 | 0 | 0 | 0 | 0 | 0 | 2 |  5 |


<details> 
  <summary>Reminiscent's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                    PR                                                                     | C | LASTED  |
|------------|-------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/21896 | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                                |   | 101d19h |
| tidb/23441 | [executor: Refactor probe channel & fix bug in chunks of join](https://github.com/pingcap/tidb/pull/23441)                                |   | 13d14h  |
| tidb/23474 | [planner: fix inappropriate null flag of null constants (#23457)](https://github.com/pingcap/tidb/pull/23474)                             |   | 9d18h   |
| tidb/23493 | [expression: Implementation of Vitess hashing algorithm.](https://github.com/pingcap/tidb/pull/23493)                                     |   | 9d6h    |
| tidb/23575 | [executor: fix update panic on join having statement (#23554)](https://github.com/pingcap/tidb/pull/23575)                                |   | 6d21h   |
| tidb/23685 | [planner: fix the issue that planner hints don't work in some batch/point-get plans (#23666)](https://github.com/pingcap/tidb/pull/23685) |   | 2d16h   |


## Reviewed in Last 7 Days

|    REPO    |                                                                PR                                                                 | C | D |   R   |
|------------|-----------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/23746 | [executor: add a test case for batch point queries on table partitions](https://github.com/pingcap/tidb/pull/23746)               |   | 1 | 16h   |
| tidb/23780 | [executor: fix 'index out of range' issue in index lookup join (#23755)](https://github.com/pingcap/tidb/pull/23780)              |   | 1 | 6h    |
| tidb/23754 | [config: remove index-usage-sync-lease from config.toml.example (#23740)](https://github.com/pingcap/tidb/pull/23754)             |   | 2 | 0h    |
| tidb/23740 | [config: remove index-usage-sync-lease from config.toml.example](https://github.com/pingcap/tidb/pull/23740)                      |   | 2 | 0h    |
| tidb/23572 | [planner, executor: IndexMerge supports reading extraHandleCol in partialTableReader](https://github.com/pingcap/tidb/pull/23572) |   | 7 | 2h    |
| tidb/23576 | [executor: fix update panic on join having statement (#23554)](https://github.com/pingcap/tidb/pull/23576)                        |   | 7 | 0h    |
| tidb/23554 | [executor: fix update panic on join having statement](https://github.com/pingcap/tidb/pull/23554)                                 |   | 7 | 16h   |
| tidb/23001 | [statistics: fix err check](https://github.com/pingcap/tidb/pull/23001)                                                           |   | 7 | 26d0h |


</details> 


<details> 
  <summary>SunRunAway's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                  PR                                                                   | C | LASTED  |
|------------|---------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19178 | [executor: Refactor probe channel](https://github.com/pingcap/tidb/pull/19178)                                                        |   | 231d17h |
| tidb/19807 | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                               |   | 209d11h |
| tidb/19900 | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                        | Y | 204d18h |
| tidb/20140 | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                    |   | 191d22h |
| tidb/20765 | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                     |   | 150d17h |
| tidb/21207 | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                       |   | 129d19h |
| tidb/21834 | [planner: enhanced index range calculation plan](https://github.com/pingcap/tidb/pull/21834)                                          |   | 106d18h |
| tidb/21876 | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)           |   | 104d19h |
| tidb/21878 | [planner: do not push down lock to pointGet/bacthPointGet when selection exists](https://github.com/pingcap/tidb/pull/21878)          |   | 104d18h |
| tidb/21956 | [planner/preprocessor: disallow into-outfile clause in some place](https://github.com/pingcap/tidb/pull/21956)                        |   | 99d23h  |
| tidb/22217 | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                         |   | 85d17h  |
| tidb/22379 | [[experiment] executor: allow aggregation to spill disk when running out of memory quota](https://github.com/pingcap/tidb/pull/22379) |   | 78d19h  |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>XuHuaiyu's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                              PR                                                                               | C | LASTED  |
|--------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| docs-cn/5619 | [Update data-type-date-and-time.md](https://github.com/pingcap/docs-cn/pull/5619)                                                                             |   | 34d16h  |
| docs-cn/5671 | [tidb: Add time format description](https://github.com/pingcap/docs-cn/pull/5671)                                                                             |   | 28d11h  |
| tidb/19900   | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                                                | Y | 204d18h |
| docs-cn/5915 | [update default oom action of TiDB](https://github.com/pingcap/docs-cn/pull/5915)                                                                             |   | 18h     |
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                                        | Y | 202d14h |
| tidb/20140   | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                                            |   | 191d22h |
| tidb/20311   | [expression: fix overflow error when convert bit to int64 (#20266)](https://github.com/pingcap/tidb/pull/20311)                                               |   | 183d21h |
| tidb/20790   | [collation: add pinyin collation for chinese charset support](https://github.com/pingcap/tidb/pull/20790)                                                     |   | 149d21h |
| tidb/21064   | [planner, executor: fix cast not check error](https://github.com/pingcap/tidb/pull/21064)                                                                     |   | 137d8h  |
| tidb/21149   | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                                 |   | 133d14h |
| tidb/21228   | [executor: return the result immediately when combining LIMIT row_count with DISTINCT](https://github.com/pingcap/tidb/pull/21228)                            |   | 129d13h |
| tidb/21304   | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                                      |   | 127d12h |
| tidb/21334   | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                                                 |   | 126d14h |
| tidb/21401   | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                                            |   | 122d11h |
| tidb/21476   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                         |   | 119d15h |
| tidb/21536   | [executor: add slow-log file meta cache to avoid repeat read file meta information](https://github.com/pingcap/tidb/pull/21536)                               |   | 115d15h |
| tidb/21564   | [ddl: fix Incorrect behavior of NO_ZERO_DATE when altering table](https://github.com/pingcap/tidb/pull/21564)                                                 |   | 114d15h |
| tidb/21896   | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                                                    |   | 101d19h |
| tidb/22131   | [privilege: remove leading and trailing space when create user and role](https://github.com/pingcap/tidb/pull/22131)                                          |   | 91d19h  |
| tidb/22163   | [expression: separated arithmeticMinusIntSig](https://github.com/pingcap/tidb/pull/22163)                                                                     |   | 87d13h  |
| tidb/22186   | [executor: fix select into outfile with year type column has no data (#22175)](https://github.com/pingcap/tidb/pull/22186)                                    |   | 86d16h  |
| tidb/22307   | [ddl: fix update can see columns not public](https://github.com/pingcap/tidb/pull/22307)                                                                      |   | 83d16h  |
| tidb/22616   | [expression: from_unixtime accept 64-bit integers](https://github.com/pingcap/tidb/pull/22616)                                                                |   | 62d23h  |
| tidb/22617   | [metrics: fix wrong bucket name of coprocessor cache (#22454)](https://github.com/pingcap/tidb/pull/22617)                                                    |   | 62d23h  |
| tidb/22624   | [ planner: not pruning column used by union scan condition (#21640)](https://github.com/pingcap/tidb/pull/22624)                                              |   | 62d17h  |
| tidb/22631   | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                                                                               |   | 60d23h  |
| tidb/22696   | [expression: enable arithmetic Mod push down](https://github.com/pingcap/tidb/pull/22696)                                                                     |   | 57d17h  |
| tidb/22711   | [executor: Fix inline schema name](https://github.com/pingcap/tidb/pull/22711)                                                                                |   | 57d11h  |
| tidb/22722   | [planner, errno: make error code of ErrMixOfGroupFuncAndFields consistent with MySQL](https://github.com/pingcap/tidb/pull/22722)                             |   | 56d20h  |
| tidb/22814   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22814)                                           |   | 41d19h  |
| tidb/22908   | [txn: Add txn state's view](https://github.com/pingcap/tidb/pull/22908)                                                                                       |   | 36d20h  |
| tidb/22926   | [expression: add overflow check in multiplyInt](https://github.com/pingcap/tidb/pull/22926)                                                                   |   | 36d13h  |
| tidb/23012   | [executor: fix affected rows of ddls and complete uint tests](https://github.com/pingcap/tidb/pull/23012)                                                     |   | 32d17h  |
| tidb/23152   | [expression: fix wrong error info (#22760)](https://github.com/pingcap/tidb/pull/23152)                                                                       |   | 25d14h  |
| tidb/23196   | [types: fix the bug about the wrong query result for decimal type  (#22507)](https://github.com/pingcap/tidb/pull/23196)                                      |   | 23d18h  |
| tidb/23220   | [Release 4.0](https://github.com/pingcap/tidb/pull/23220)                                                                                                     |   | 23d11h  |
| tidb/23227   | [executor: hash join out of index panic when enum column value is zero (#23162)](https://github.com/pingcap/tidb/pull/23227)                                  |   | 22d22h  |
| tidb/23233   | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23233)                                                        |   | 22d18h  |
| tidb/23257   | [executor: group_concat aggr panic when session.group_concat_max_len is small (#23131)](https://github.com/pingcap/tidb/pull/23257)                           |   | 21d18h  |
| tidb/23295   | [util, types: don't let SPM be affected by charset (#23161)](https://github.com/pingcap/tidb/pull/23295)                                                      |   | 20d11h  |
| tidb/23335   | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23335)                                      |   | 16d19h  |
| tidb/23336   | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23336)                                      |   | 16d19h  |
| tidb/23347   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                                          |   | 16d18h  |
| tidb/23348   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                                          |   | 16d18h  |
| tidb/23350   | [util/stringutil, util/ranger, planner: use hierarchical separators to simplify the parsing for info of EXPLAIN ](https://github.com/pingcap/tidb/pull/23350) |   | 16d17h  |
| tidb/23368   | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23368)                                         |   | 15d20h  |
| tidb/23369   | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23369)                                         |   | 15d20h  |
| tidb/23397   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23397)                                                                |   | 14d17h  |
| tidb/23398   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23398)                                                                |   | 14d17h  |
| tidb/23405   | [domain: remove the exit chan, use context](https://github.com/pingcap/tidb/pull/23405)                                                                       |   | 14d17h  |
| tidb/23433   | [WIP: speed up for slow query logs retrieving ](https://github.com/pingcap/tidb/pull/23433)                                                                   |   | 13d17h  |
| tidb/23441   | [executor: Refactor probe channel & fix bug in chunks of join](https://github.com/pingcap/tidb/pull/23441)                                                    |   | 13d14h  |
| tidb/23474   | [planner: fix inappropriate null flag of null constants (#23457)](https://github.com/pingcap/tidb/pull/23474)                                                 |   | 9d18h   |
| tidb/23487   | [planner: optimize count(distinct a) to count(a) if there is an unique key on a](https://github.com/pingcap/tidb/pull/23487)                                  | Y | 9d14h   |
| tidb/23493   | [expression: Implementation of Vitess hashing algorithm.](https://github.com/pingcap/tidb/pull/23493)                                                         |   | 9d6h    |
| tidb/23497   | [expression: Let TiDB use Hyperscan to support multi-pattern-match](https://github.com/pingcap/tidb/pull/23497)                                               |   | 8d22h   |
| tidb/23517   | [*: Add the metric about the SQL with TiFlash Success  (#23426)](https://github.com/pingcap/tidb/pull/23517)                                                  |   | 8d12h   |
| tidb/23562   | [execution: reuse iterator in hash join](https://github.com/pingcap/tidb/pull/23562)                                                                          |   | 7d13h   |
| tidb/23640   | [*: fix the bug about YEAR(0.9) returns NULL instead of 0 in NO_ZERO_DATE mode](https://github.com/pingcap/tidb/pull/23640)                                   |   | 3d13h   |
| tidb/23661   | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23661)       |   | 2d20h   |
| tidb/23682   | [executor: fix a panic when batch point get is used for partition table (#23652)](https://github.com/pingcap/tidb/pull/23682)                                 |   | 2d16h   |
| tidb/23683   | [executor: fix a panic when batch point get is used for partition table (#23652)](https://github.com/pingcap/tidb/pull/23683)                                 |   | 2d16h   |
| tidb/23691   | [executor: fix index join on prefix column index (#23678)](https://github.com/pingcap/tidb/pull/23691)                                                        |   | 2d15h   |
| tidb/23705   | [executor: refineArgs() bug fix when compare int with very small decimal (#23694)](https://github.com/pingcap/tidb/pull/23705)                                |   | 2d13h   |
| tidb/23756   | [planner: fix set not null flag for outer join (#23727)](https://github.com/pingcap/tidb/pull/23756)                                                          |   | 1d14h   |
| tidb/23757   | [planner: fix set not null flag for outer join (#23727)](https://github.com/pingcap/tidb/pull/23757)                                                          |   | 1d14h   |
| tidb/23812   | [executor, planner: fix collation for hash join building (#23770)](https://github.com/pingcap/tidb/pull/23812)                                                |   | 12h     |


## Reviewed in Last 7 Days

|     REPO     |                                                          PR                                                           | C | D |  R   |
|--------------|-----------------------------------------------------------------------------------------------------------------------|---|---|------|
| tidb/23770   | [executor, planner: fix collation for hash join building](https://github.com/pingcap/tidb/pull/23770)                 |   | 2 | 0h   |
| docs-cn/5892 | [add apply-cache](https://github.com/pingcap/docs-cn/pull/5892)                                                       |   | 2 | 15h  |
| tidb/23493   | [expression: Implementation of Vitess hashing algorithm.](https://github.com/pingcap/tidb/pull/23493)                 |   | 2 | 7d7h |
| tidb/23692   | [executor: fix index join on prefix column index (#23678)](https://github.com/pingcap/tidb/pull/23692)                |   | 3 | 0h   |
| tidb/23694   | [executor: refineArgs() bug fix when compare int with very small decimal](https://github.com/pingcap/tidb/pull/23694) |   | 3 | 0h   |
| parser/1202  | [mysql: modify TypeNewDecimal length in defaultLengthAndDecimalForCast](https://github.com/pingcap/parser/pull/1202)  |   | 3 | 14h  |
| tidb/23576   | [executor: fix update panic on join having statement (#23554)](https://github.com/pingcap/tidb/pull/23576)            |   | 7 | 0h   |
| tidb/23545   | [go.mod: update parser to fix panic on connection verification #23532](https://github.com/pingcap/tidb/pull/23545)    |   | 7 | 18h  |


</details> 


<details> 
  <summary>dyzsr's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                             PR                                                             | C | LASTED |
|------------|----------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/23559 | [ranger: fix the range construction behavior when the column's type is `YEAR`](https://github.com/pingcap/tidb/pull/23559) |   | 7d14h  |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>eurekaka's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                    PR                                                                     | C | LASTED  |
|------------|-------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                                 |   | 147d17h |
| docs/5150  | [SPM: update DML SQL Bind and baseline capture description (#5088)](https://github.com/pingcap/docs/pull/5150)                            |   | 3d15h   |
| tidb/21444 | [planner: ignore anonymous index while tiflash replica is available](https://github.com/pingcap/tidb/pull/21444)                          |   | 120d12h |
| tidb/22416 | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                                         | Y | 75d11h  |
| tidb/22559 | [planner: split test data from test cases in cbo_test.go](https://github.com/pingcap/tidb/pull/22559)                                     |   | 64d19h  |
| tidb/22853 | [planner: fix LogicalPlans that contain Window Function are ambiguous ](https://github.com/pingcap/tidb/pull/22853)                       |   | 40d12h  |
| tidb/23137 | [planner: fix index merge row count estimation logic](https://github.com/pingcap/tidb/pull/23137)                                         |   | 27d17h  |
| tidb/23208 | [statistics, util/ranger: improve selectivity calculation for DNF filters (#18741)](https://github.com/pingcap/tidb/pull/23208)           |   | 23d16h  |
| tidb/23283 | [util: optimize the performance of restore with db (#22910)](https://github.com/pingcap/tidb/pull/23283)                                  |   | 20d17h  |
| tidb/23295 | [util, types: don't let SPM be affected by charset (#23161)](https://github.com/pingcap/tidb/pull/23295)                                  |   | 20d11h  |
| tidb/23316 | [planner: Fix rebuild range for prepared plan](https://github.com/pingcap/tidb/pull/23316)                                                |   | 17d17h  |
| tidb/23365 | [planner: fix a bug that point get plan returns wrong column name](https://github.com/pingcap/tidb/pull/23365)                            |   | 15d22h  |
| tidb/23373 | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                         |   | 15d19h  |
| tidb/23543 | [statistics: fix auto analyze log information incomplete (#23522)](https://github.com/pingcap/tidb/pull/23543)                            |   | 7d18h   |
| tidb/23575 | [executor: fix update panic on join having statement (#23554)](https://github.com/pingcap/tidb/pull/23575)                                |   | 6d21h   |
| tidb/23685 | [planner: fix the issue that planner hints don't work in some batch/point-get plans (#23666)](https://github.com/pingcap/tidb/pull/23685) |   | 2d16h   |
| tidb/23689 | [planner: fix the panic when we calculate the partition range (#23651)](https://github.com/pingcap/tidb/pull/23689)                       |   | 2d16h   |
| tidb/23705 | [executor: refineArgs() bug fix when compare int with very small decimal (#23694)](https://github.com/pingcap/tidb/pull/23705)            |   | 2d13h   |
| tidb/23749 | [planner/core: fix a bug during add cast for decimal join key  (#23723)](https://github.com/pingcap/tidb/pull/23749)                      |   | 1d16h   |
| tidb/23756 | [planner: fix set not null flag for outer join (#23727)](https://github.com/pingcap/tidb/pull/23756)                                      |   | 1d14h   |
| tidb/23757 | [planner: fix set not null flag for outer join (#23727)](https://github.com/pingcap/tidb/pull/23757)                                      |   | 1d14h   |
| tidb/23760 | [collation: fix tidb panic when compare string with collation](https://github.com/pingcap/tidb/pull/23760)                                |   | 1d13h   |


## Reviewed in Last 7 Days

|     REPO     |                                                                PR                                                                 | C | D |   R   |
|--------------|-----------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| docs-cn/5835 | [Update 5.0 GA release notes and experimental features](https://github.com/pingcap/docs-cn/pull/5835)                             |   | 3 | 5d3h  |
| tidb/23628   | [planner, util/ranger:  apply PushDownNot to condition before pruning partition](https://github.com/pingcap/tidb/pull/23628)      |   | 4 | 2h    |
| tidb/23417   | [plan: reset not null flag](https://github.com/pingcap/tidb/pull/23417)                                                           |   | 7 | 7d17h |
| tidb/23572   | [planner, executor: IndexMerge supports reading extraHandleCol in partialTableReader](https://github.com/pingcap/tidb/pull/23572) |   | 7 | 0h    |


</details> 


<details> 
  <summary>fzhedu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                          PR                                                          | C | LASTED |
|------------|----------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/22853 | [planner: fix LogicalPlans that contain Window Function are ambiguous ](https://github.com/pingcap/tidb/pull/22853)  |   | 40d12h |
| tidb/23749 | [planner/core: fix a bug during add cast for decimal join key  (#23723)](https://github.com/pingcap/tidb/pull/23749) |   | 1d16h  |


## Reviewed in Last 7 Days

|    REPO    |                                                      PR                                                       | C | D |   R    |
|------------|---------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/23723 | [planner/core: fix a bug during add cast for decimal join key ](https://github.com/pingcap/tidb/pull/23723)   |   | 2 | 2h     |
| tidb/22955 | [planner, executor: reset NotNullFlag when merge schema for join](https://github.com/pingcap/tidb/pull/22955) |   | 2 | 33d17h |


</details> 


<details> 
  <summary>ichn-hu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                     PR                                                     | C | LASTED |
|------------|------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/23441 | [executor: Refactor probe channel & fix bug in chunks of join](https://github.com/pingcap/tidb/pull/23441) |   | 13d14h |
| tidb/23691 | [executor: fix index join on prefix column index (#23678)](https://github.com/pingcap/tidb/pull/23691)     |   | 2d15h  |


## Reviewed in Last 7 Days

|    REPO    |                                                    PR                                                     | C | D |   R   |
|------------|-----------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/23680 | [*: add test for modifying default length of cast as decimal](https://github.com/pingcap/tidb/pull/23680) |   | 1 | 1d23h |
| tidb/23702 | [expression: fix approx_percent panic on bit column (#23687)](https://github.com/pingcap/tidb/pull/23702) |   | 2 | 21h   |
| tidb/23703 | [expression: fix approx_percent panic on bit column (#23687)](https://github.com/pingcap/tidb/pull/23703) |   | 2 | 21h   |
| tidb/23687 | [expression: fix approx_percent panic on bit column](https://github.com/pingcap/tidb/pull/23687)          |   | 3 | 0h    |
| tidb/23678 | [executor: fix index join on prefix column index](https://github.com/pingcap/tidb/pull/23678)             |   | 3 | 0h    |


</details> 


<details> 
  <summary>lzmhhh123's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                             PR                                                                              | C | LASTED  |
|------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20444 | [expression: add json_merge_patch](https://github.com/pingcap/tidb/pull/20444)                                                                              |   | 169d21h |
| tidb/20465 | [expression: add uuidShortFunction](https://github.com/pingcap/tidb/pull/20465)                                                                             |   | 168d19h |
| tidb/20642 | [executor: modify admin executors to support partitioned table with global index](https://github.com/pingcap/tidb/pull/20642)                               |   | 157d15h |
| tidb/20903 | [planner: fix confused and unnecessary double-projection in plans.](https://github.com/pingcap/tidb/pull/20903)                                             |   | 146d17h |
| tidb/21018 | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                              |   | 140d17h |
| tidb/21195 | [brie: integrate lightning to suport IMPORT statement](https://github.com/pingcap/tidb/pull/21195)                                                          |   | 129d22h |
| tidb/21334 | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                                               |   | 126d14h |
| tidb/21347 | [session: make rollback work on global variables](https://github.com/pingcap/tidb/pull/21347)                                                               |   | 125d19h |
| tidb/21444 | [planner: ignore anonymous index while tiflash replica is available](https://github.com/pingcap/tidb/pull/21444)                                            |   | 120d12h |
| tidb/21487 | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                                               |   | 119d4h  |
| tidb/21641 | [executor: Fix pessimistic lock doesn't work on the partition table for subquery/joins](https://github.com/pingcap/tidb/pull/21641)                         |   | 112d18h |
| tidb/21651 | [planner: allow filter condition pushing down to IndexScan for prefix index](https://github.com/pingcap/tidb/pull/21651)                                    |   | 112d13h |
| tidb/22126 | [*: add `sys` schema, `sys.SCHEMA_UNUSED_INDEXES` view and `sys.SCHEMA_INDEX_USAGE` view](https://github.com/pingcap/tidb/pull/22126)                       |   | 91d20h  |
| tidb/22361 | [table: fix insert into _tidb_rowid panic and rebase it if needed (#22062)](https://github.com/pingcap/tidb/pull/22361)                                     |   | 79d20h  |
| tidb/22372 | [executor: fix SelectForUpdate in decorrelated subquery under pessimistic mode](https://github.com/pingcap/tidb/pull/22372)                                 |   | 79d9h   |
| tidb/22478 | [planner, executor: fix query partition table with global unique index get wrong result](https://github.com/pingcap/tidb/pull/22478)                        |   | 70d13h  |
| tidb/22631 | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                                                                             |   | 60d23h  |
| tidb/22686 | [expression: support enum pushdown](https://github.com/pingcap/tidb/pull/22686)                                                                             |   | 57d22h  |
| tidb/22699 | [brie: add error info column and history backup/restore info in sql](https://github.com/pingcap/tidb/pull/22699)                                            |   | 57d16h  |
| tidb/22926 | [expression: add overflow check in multiplyInt](https://github.com/pingcap/tidb/pull/22926)                                                                 |   | 36d13h  |
| tidb/23001 | [statistics: fix err check](https://github.com/pingcap/tidb/pull/23001)                                                                                     |   | 33d0h   |
| tidb/23022 | [executor: create PipelinedWindowExec based on current implementation and modify the windowProcessor interface](https://github.com/pingcap/tidb/pull/23022) |   | 31d18h  |
| tidb/23149 | [core: support left join and right join for join reorder](https://github.com/pingcap/tidb/pull/23149)                                                       |   | 26d12h  |
| tidb/23257 | [executor: group_concat aggr panic when session.group_concat_max_len is small (#23131)](https://github.com/pingcap/tidb/pull/23257)                         |   | 21d18h  |
| tidb/23283 | [util: optimize the performance of restore with db (#22910)](https://github.com/pingcap/tidb/pull/23283)                                                    |   | 20d17h  |
| tidb/23296 | [sig/execution: fix the bug that Wrong result of comparison operation(type date / type string)](https://github.com/pingcap/tidb/pull/23296)                 |   | 20d7h   |
| tidb/23347 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                                        |   | 16d18h  |
| tidb/23348 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                                        |   | 16d18h  |
| tidb/23368 | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23368)                                       |   | 15d20h  |
| tidb/23369 | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23369)                                       |   | 15d20h  |
| tidb/23373 | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                           |   | 15d19h  |
| tidb/23422 | [sig/execution: fix the bug that The result of 'varbinary + 1' is incorrect](https://github.com/pingcap/tidb/pull/23422)                                    |   | 14d6h   |
| tidb/23441 | [executor: Refactor probe channel & fix bug in chunks of join](https://github.com/pingcap/tidb/pull/23441)                                                  |   | 13d14h  |
| tidb/23493 | [expression: Implementation of Vitess hashing algorithm.](https://github.com/pingcap/tidb/pull/23493)                                                       |   | 9d6h    |
| tidb/23559 | [ranger: fix the range construction behavior when the column's type is `YEAR`](https://github.com/pingcap/tidb/pull/23559)                                  |   | 7d14h   |
| tidb/23655 | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral (#23523)](https://github.com/pingcap/tidb/pull/23655)                    |   | 2d22h   |
| tidb/23656 | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral (#23523)](https://github.com/pingcap/tidb/pull/23656)                    |   | 2d22h   |
| tidb/23660 | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23660)     |   | 2d20h   |
| tidb/23661 | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23661)     |   | 2d20h   |
| tidb/23680 | [*: add test for modifying default length of cast as decimal](https://github.com/pingcap/tidb/pull/23680)                                                   |   | 2d17h   |
| tidb/23700 | [tikv: distinguish server timeout for TiKV and TiFlash](https://github.com/pingcap/tidb/pull/23700)                                                         |   | 2d14h   |
| tidb/23703 | [expression: fix approx_percent panic on bit column (#23687)](https://github.com/pingcap/tidb/pull/23703)                                                   |   | 2d14h   |
| tidb/23705 | [executor: refineArgs() bug fix when compare int with very small decimal (#23694)](https://github.com/pingcap/tidb/pull/23705)                              |   | 2d13h   |
| tidb/23714 | [*:Support record statment_history table evicted info](https://github.com/pingcap/tidb/pull/23714)                                                          |   | 2d1h    |
| tidb/23738 | [executor: remove duplicate entry in show privileges](https://github.com/pingcap/tidb/pull/23738)                                                           |   | 1d17h   |
| tidb/23747 | [planner, sessionvar: avoid sending same task id to TiFlash](https://github.com/pingcap/tidb/pull/23747)                                                    |   | 1d16h   |
| tidb/23749 | [planner/core: fix a bug during add cast for decimal join key  (#23723)](https://github.com/pingcap/tidb/pull/23749)                                        |   | 1d16h   |
| tidb/23756 | [planner: fix set not null flag for outer join (#23727)](https://github.com/pingcap/tidb/pull/23756)                                                        |   | 1d14h   |
| tidb/23757 | [planner: fix set not null flag for outer join (#23727)](https://github.com/pingcap/tidb/pull/23757)                                                        |   | 1d14h   |
| tidb/23760 | [collation: fix tidb panic when compare string with collation](https://github.com/pingcap/tidb/pull/23760)                                                  |   | 1d13h   |
| tidb/23764 | [planner: fix partition selection of point plan for the update statement](https://github.com/pingcap/tidb/pull/23764)                                       |   | 1d13h   |
| tidb/23796 | [tests: make TestIndexLookupMergeJoinHang and TestIssue18068 stable (#23741)](https://github.com/pingcap/tidb/pull/23796)                                   |   | 19h     |
| tidb/23812 | [executor, planner: fix collation for hash join building (#23770)](https://github.com/pingcap/tidb/pull/23812)                                              |   | 12h     |
| tidb/23818 | [*: protect read only noop via tidb_enable_noop_functions](https://github.com/pingcap/tidb/pull/23818)                                                      |   | 5h      |


## Reviewed in Last 7 Days

|    REPO    |                                                                       PR                                                                       | C | D |   R    |
|------------|------------------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/23706 | [executor: refineArgs() bug fix when compare int with very small decimal (#23694)](https://github.com/pingcap/tidb/pull/23706)                 |   | 3 | 0h     |
| tidb/23694 | [executor: refineArgs() bug fix when compare int with very small decimal](https://github.com/pingcap/tidb/pull/23694)                          |   | 3 | 1h     |
| tidb/23697 | [planner: set schema column be the same with new expr one (#23679)](https://github.com/pingcap/tidb/pull/23697)                                |   | 3 | 0h     |
| tidb/23679 | [planner: set schema column be the same with new expr one](https://github.com/pingcap/tidb/pull/23679)                                         |   | 3 | 0h     |
| tidb/23284 | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified.](https://github.com/pingcap/tidb/pull/23284) |   | 3 | 17d20h |
| tidb/23597 | [planner: fix correlated columns in filter or access in MPP apply (#23509)](https://github.com/pingcap/tidb/pull/23597)                        |   | 7 | 0h     |
| tidb/23592 | [MPP: fix 2-phase agg chose wrong partition column during planning (#23557)](https://github.com/pingcap/tidb/pull/23592)                       |   | 7 | 2h     |
| tikv/9870  | [copr: cast invalid utf8 string to real bug (#9860)](https://github.com/tikv/tikv/pull/9870)                                                   | Y | 7 | 2d18h  |


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
| docs-cn/5561 | [Add sql optimization-related docs to toc](https://github.com/pingcap/docs-cn/pull/5561)                                                                    |   | 38d15h  |
| tidb/19029   | [types: fix unexpected NOT_NULL flags](https://github.com/pingcap/tidb/pull/19029)                                                                          |   | 238d22h |
| docs-cn/5785 | [update SPM documentation for DML SQL Bind and baseline capture (#5740)](https://github.com/pingcap/docs-cn/pull/5785)                                      |   | 10d18h  |
| tidb/20708   | [*: separate auto_increment ID allocator from _tidb_rowid allocator](https://github.com/pingcap/tidb/pull/20708)                                            |   | 154d20h |
| tidb/20969   | [executor: Improve the performance of appending not fixed columns](https://github.com/pingcap/tidb/pull/20969)                                              |   | 142d9h  |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                              |   | 140d17h |
| tidb/21149   | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                               |   | 133d14h |
| tidb/21304   | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                                    |   | 127d12h |
| tidb/21318   | [planner, expression: use the range of column types to simplify expressions](https://github.com/pingcap/tidb/pull/21318)                                    |   | 126d19h |
| tidb/21401   | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                                          |   | 122d11h |
| tidb/21476   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                       |   | 119d15h |
| tidb/21508   | [execution: fix dayofweek('0000-00-00') behavior](https://github.com/pingcap/tidb/pull/21508)                                                               |   | 118d10h |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                                 |   | 104d19h |
| tidb/21887   | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                                             |   | 103d11h |
| tidb/21954   | [planner/cascades: add rule `PushSelDownApply`](https://github.com/pingcap/tidb/pull/21954)                                                                 |   | 99d23h  |
| tidb/22146   | [executor: forbid SFU on view](https://github.com/pingcap/tidb/pull/22146)                                                                                  |   | 87d21h  |
| tidb/22217   | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                                               |   | 85d17h  |
| tidb/22234   | [executor, planner: ON DUPLICATE UPDATE can refer to un-project col (#14412)](https://github.com/pingcap/tidb/pull/22234)                                   |   | 85d15h  |
| tidb/22261   | [time: fix parse datetime won't truncate the reluctant string (#22232)](https://github.com/pingcap/tidb/pull/22261)                                         |   | 84d19h  |
| tidb/22307   | [ddl: fix update can see columns not public](https://github.com/pingcap/tidb/pull/22307)                                                                    |   | 83d16h  |
| tidb/22374   | [expression: separated arithmeticIntDivideSig](https://github.com/pingcap/tidb/pull/22374)                                                                  |   | 79d0h   |
| tidb/22415   | [ddl: refactor placement package](https://github.com/pingcap/tidb/pull/22415)                                                                               |   | 75d17h  |
| tidb/22541   | [expression: Support builtin function SOUNDEX](https://github.com/pingcap/tidb/pull/22541)                                                                  |   | 65d9h   |
| tidb/22565   | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)                                          | Y | 64d17h  |
| tidb/22814   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22814)                                         |   | 41d19h  |
| tidb/22862   | [brie: fix the problem that ddl restored by BR via SQL is not replicated to downstream](https://github.com/pingcap/tidb/pull/22862)                         |   | 38d22h  |
| tidb/22923   | [expression: correct constant propagation for collation (#22666)](https://github.com/pingcap/tidb/pull/22923)                                               |   | 36d15h  |
| tidb/22924   | [planner: fix wrong index merge selection (#22825)](https://github.com/pingcap/tidb/pull/22924)                                                             |   | 36d14h  |
| tidb/22926   | [expression: add overflow check in multiplyInt](https://github.com/pingcap/tidb/pull/22926)                                                                 |   | 36d13h  |
| tidb/22984   | [executor: fix logging format of prepared statements (#16062)](https://github.com/pingcap/tidb/pull/22984)                                                  |   | 33d10h  |
| tidb/23022   | [executor: create PipelinedWindowExec based on current implementation and modify the windowProcessor interface](https://github.com/pingcap/tidb/pull/23022) |   | 31d18h  |
| tidb/23152   | [expression: fix wrong error info (#22760)](https://github.com/pingcap/tidb/pull/23152)                                                                     |   | 25d14h  |
| tidb/23196   | [types: fix the bug about the wrong query result for decimal type  (#22507)](https://github.com/pingcap/tidb/pull/23196)                                    |   | 23d18h  |
| tidb/23208   | [statistics, util/ranger: improve selectivity calculation for DNF filters (#18741)](https://github.com/pingcap/tidb/pull/23208)                             |   | 23d16h  |
| tidb/23295   | [util, types: don't let SPM be affected by charset (#23161)](https://github.com/pingcap/tidb/pull/23295)                                                    |   | 20d11h  |
| tidb/23316   | [planner: Fix rebuild range for prepared plan](https://github.com/pingcap/tidb/pull/23316)                                                                  |   | 17d17h  |
| tidb/23373   | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                           |   | 15d19h  |
| tidb/23397   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23397)                                                              |   | 14d17h  |
| tidb/23398   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23398)                                                              |   | 14d17h  |
| tidb/23448   | [wip :execution: parallel build hash table](https://github.com/pingcap/tidb/pull/23448)                                                                     |   | 11d12h  |
| tidb/23493   | [expression: Implementation of Vitess hashing algorithm.](https://github.com/pingcap/tidb/pull/23493)                                                       |   | 9d6h    |
| tidb/23543   | [statistics: fix auto analyze log information incomplete (#23522)](https://github.com/pingcap/tidb/pull/23543)                                              |   | 7d18h   |
| tidb/23590   | [planner, table: optimize the list partition pruner for range query](https://github.com/pingcap/tidb/pull/23590)                                            |   | 6d16h   |
| tidb/23598   | [types: fix collation for binary literal (#23591)](https://github.com/pingcap/tidb/pull/23598)                                                              |   | 6d13h   |
| tidb/23655   | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral (#23523)](https://github.com/pingcap/tidb/pull/23655)                    |   | 2d22h   |
| tidb/23656   | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral (#23523)](https://github.com/pingcap/tidb/pull/23656)                    |   | 2d22h   |
| tidb/23658   | [*: collect transaction write duration/throughput metrics for SLI/SLO (#23462)](https://github.com/pingcap/tidb/pull/23658)                                 |   | 2d22h   |
| tidb/23660   | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23660)     |   | 2d20h   |
| tidb/23661   | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23661)     |   | 2d20h   |
| tidb/23674   | [*: add column `End_time` in show analyze status and add related log](https://github.com/pingcap/tidb/pull/23674)                                           |   | 2d17h   |
| tidb/23682   | [executor: fix a panic when batch point get is used for partition table (#23652)](https://github.com/pingcap/tidb/pull/23682)                               |   | 2d16h   |
| tidb/23689   | [planner: fix the panic when we calculate the partition range (#23651)](https://github.com/pingcap/tidb/pull/23689)                                         |   | 2d16h   |
| tidb/23718   | [*: add TableSample ID for PhysicalIDToTypeString()](https://github.com/pingcap/tidb/pull/23718)                                                            |   | 1d23h   |
| tidb/23728   | [planner: skip storage engine check for CRAETE VIEW statement](https://github.com/pingcap/tidb/pull/23728)                                                  |   | 1d19h   |
| tidb/23730   | [distsql/*: typo fix for `dispatches`](https://github.com/pingcap/tidb/pull/23730)                                                                          |   | 1d18h   |
| tidb/23764   | [planner: fix partition selection of point plan for the update statement](https://github.com/pingcap/tidb/pull/23764)                                       |   | 1d13h   |
| tidb/23796   | [tests: make TestIndexLookupMergeJoinHang and TestIssue18068 stable (#23741)](https://github.com/pingcap/tidb/pull/23796)                                   |   | 19h     |
| tidb/23812   | [executor, planner: fix collation for hash join building (#23770)](https://github.com/pingcap/tidb/pull/23812)                                              |   | 12h     |


## Reviewed in Last 7 Days

|     REPO     |                                                               PR                                                                | C | D |   R    |
|--------------|---------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/23808   | [statistics: feedback not panic when no ndv collected](https://github.com/pingcap/tidb/pull/23808)                              |   | 1 | 0h     |
| tidb/23770   | [executor, planner: fix collation for hash join building](https://github.com/pingcap/tidb/pull/23770)                           |   | 1 | 20h    |
| tidb/22559   | [planner: split test data from test cases in cbo_test.go](https://github.com/pingcap/tidb/pull/22559)                           |   | 1 | 64d0h  |
| tidb/23137   | [planner: fix index merge row count estimation logic](https://github.com/pingcap/tidb/pull/23137)                               |   | 1 | 26d22h |
| tidb/22915   | [planner: build correct MaxOneRow info from multi-column conditions](https://github.com/pingcap/tidb/pull/22915)                |   | 1 | 35d22h |
| tidb/23683   | [executor: fix a panic when batch point get is used for partition table (#23652)](https://github.com/pingcap/tidb/pull/23683)   |   | 1 | 1d18h  |
| tidb/23702   | [expression: fix approx_percent panic on bit column (#23687)](https://github.com/pingcap/tidb/pull/23702)                       |   | 1 | 1d15h  |
| tidb/23703   | [expression: fix approx_percent panic on bit column (#23687)](https://github.com/pingcap/tidb/pull/23703)                       |   | 1 | 1d15h  |
| tidb/23746   | [executor: add a test case for batch point queries on table partitions](https://github.com/pingcap/tidb/pull/23746)             |   | 1 | 17h    |
| tidb/23749   | [planner/core: fix a bug during add cast for decimal join key  (#23723)](https://github.com/pingcap/tidb/pull/23749)            |   | 1 | 17h    |
| tidb/23766   | [config: add a test for `config.toml.example`](https://github.com/pingcap/tidb/pull/23766)                                      |   | 1 | 14h    |
| tidb/23783   | [*: fix a bug that wrong index data on prefixed clustered index  (#23742)](https://github.com/pingcap/tidb/pull/23783)          |   | 1 | 0h     |
| tidb/23747   | [planner, sessionvar: avoid sending same task id to TiFlash](https://github.com/pingcap/tidb/pull/23747)                        |   | 1 | 17h    |
| tidb/23755   | [executor: fix 'index out of range' issue in index lookup join](https://github.com/pingcap/tidb/pull/23755)                     |   | 2 | 3h     |
| tidb/23769   | [planner: fix wrong IndexScan plan reused in plan cache (#23758)](https://github.com/pingcap/tidb/pull/23769)                   |   | 2 | 0h     |
| tidb/23754   | [config: remove index-usage-sync-lease from config.toml.example (#23740)](https://github.com/pingcap/tidb/pull/23754)           |   | 2 | 0h     |
| tidb/23740   | [config: remove index-usage-sync-lease from config.toml.example](https://github.com/pingcap/tidb/pull/23740)                    |   | 2 | 0h     |
| tidb/23723   | [planner/core: fix a bug during add cast for decimal join key ](https://github.com/pingcap/tidb/pull/23723)                     |   | 2 | 0h     |
| tidb/23699   | [util: fix range building for binary literal](https://github.com/pingcap/tidb/pull/23699)                                       |   | 3 | 0h     |
| tidb/23690   | [planner: fix the panic when we calculate the partition range (#23651)](https://github.com/pingcap/tidb/pull/23690)             |   | 3 | 1h     |
| tidb/23652   | [executor: fix a panic when batch point get is used for partition table](https://github.com/pingcap/tidb/pull/23652)            |   | 3 | 6h     |
| tidb/23651   | [planner: fix the panic when we calculate the partition range](https://github.com/pingcap/tidb/pull/23651)                      |   | 3 | 0h     |
| tidb/23523   | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral](https://github.com/pingcap/tidb/pull/23523) |   | 3 | 4d23h  |
| tidb/23597   | [planner: fix correlated columns in filter or access in MPP apply (#23509)](https://github.com/pingcap/tidb/pull/23597)         |   | 7 | 0h     |
| docs/5090    | [update tidb v4.0.12 release notes](https://github.com/pingcap/docs/pull/5090)                                                  |   | 7 | 2d21h  |
| docs-cn/5806 | [releases: add tidb release notes 4.0.12](https://github.com/pingcap/docs-cn/pull/5806)                                         |   | 7 | 2d20h  |
| tidb/23557   | [MPP: fix 2-phase agg chose wrong partition column during planning](https://github.com/pingcap/tidb/pull/23557)                 |   | 7 | 20h    |
| tidb/23574   | [planner/core: inject project for tiflash agg (#23480)](https://github.com/pingcap/tidb/pull/23574)                             |   | 7 | 3h     |
| tidb/23565   | [planner/core: convert decimal type for mpp join before shuffling. (#23191)](https://github.com/pingcap/tidb/pull/23565)        |   | 7 | 15h    |
| tipb/198     | [Adding vitess_hash function code to tipb](https://github.com/pingcap/tipb/pull/198)                                            |   | 7 | 135d5h |
| tidb/23509   | [planner: fix correlated columns in filter or access in MPP apply](https://github.com/pingcap/tidb/pull/23509)                  |   | 7 | 1d22h  |
| tidb/23480   | [planner/core: inject project for tiflash agg](https://github.com/pingcap/tidb/pull/23480)                                      |   | 7 | 2d18h  |


</details> 


<details> 
  <summary>rebelice's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                    PR                                                                     | C | LASTED |
|------------|-------------------------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/22853 | [planner: fix LogicalPlans that contain Window Function are ambiguous ](https://github.com/pingcap/tidb/pull/22853)                       |   | 40d12h |
| tidb/23537 | [planner: remove some risky cache operations in the plan builder (#23354)](https://github.com/pingcap/tidb/pull/23537)                    |   | 7d19h  |
| tidb/23674 | [*: add column `End_time` in show analyze status and add related log](https://github.com/pingcap/tidb/pull/23674)                         |   | 2d17h  |
| tidb/23685 | [planner: fix the issue that planner hints don't work in some batch/point-get plans (#23666)](https://github.com/pingcap/tidb/pull/23685) |   | 2d16h  |


## Reviewed in Last 7 Days

|    REPO    |                                                                PR                                                                | C | D | R  |
|------------|----------------------------------------------------------------------------------------------------------------------------------|---|---|----|
| tidb/23674 | [*: add column `End_time` in show analyze status and add related log](https://github.com/pingcap/tidb/pull/23674)                |   | 3 | 0h |
| tidb/23666 | [planner: fix the issue that planner hints don't work in some batch/point-get plans](https://github.com/pingcap/tidb/pull/23666) |   | 3 | 0h |


</details> 


<details> 
  <summary>time-and-fate's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                         PR                                                          | C | LASTED  |
|------------|---------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                           |   | 147d17h |
| tidb/22853 | [planner: fix LogicalPlans that contain Window Function are ambiguous ](https://github.com/pingcap/tidb/pull/22853) |   | 40d12h  |
| tidb/23728 | [planner: skip storage engine check for CRAETE VIEW statement](https://github.com/pingcap/tidb/pull/23728)          |   | 1d19h   |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>winoros's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                              PR                                                                               | C | LASTED  |
|--------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                                        | Y | 202d14h |
| docs-cn/5916 | [sql-statements, information-schema: add `END_TIME` field for table `ANALYZE_STATUS`](https://github.com/pingcap/docs-cn/pull/5916)                           |   | 17h     |
| tidb/20311   | [expression: fix overflow error when convert bit to int64 (#20266)](https://github.com/pingcap/tidb/pull/20311)                                               |   | 183d21h |
| tidb/20765   | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                                             |   | 150d17h |
| tidb/20877   | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                                                     |   | 147d17h |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                                |   | 140d17h |
| tidb/21207   | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                                               |   | 129d19h |
| tidb/21476   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                         |   | 119d15h |
| tidb/21487   | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                                                 |   | 119d4h  |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                                   |   | 104d19h |
| tidb/21954   | [planner/cascades: add rule `PushSelDownApply`](https://github.com/pingcap/tidb/pull/21954)                                                                   |   | 99d23h  |
| tidb/22181   | [planner, expression: fix error when using IN combined with subquery (#22080)](https://github.com/pingcap/tidb/pull/22181)                                    |   | 86d17h  |
| tidb/22504   | [*:Fix the fetchHotRegion bug that the count always zero](https://github.com/pingcap/tidb/pull/22504)                                                         |   | 67d19h  |
| tidb/22565   | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)                                            | Y | 64d17h  |
| tidb/22624   | [ planner: not pruning column used by union scan condition (#21640)](https://github.com/pingcap/tidb/pull/22624)                                              |   | 62d17h  |
| tidb/22923   | [expression: correct constant propagation for collation (#22666)](https://github.com/pingcap/tidb/pull/22923)                                                 |   | 36d15h  |
| tidb/23208   | [statistics, util/ranger: improve selectivity calculation for DNF filters (#18741)](https://github.com/pingcap/tidb/pull/23208)                               |   | 23d16h  |
| tidb/23215   | [Privileges: fix delete privilege check wrongly (#22971)](https://github.com/pingcap/tidb/pull/23215)                                                         |   | 23d14h  |
| tidb/23233   | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23233)                                                        |   | 22d18h  |
| tidb/23347   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                                          |   | 16d18h  |
| tidb/23348   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                                          |   | 16d18h  |
| tidb/23350   | [util/stringutil, util/ranger, planner: use hierarchical separators to simplify the parsing for info of EXPLAIN ](https://github.com/pingcap/tidb/pull/23350) |   | 16d17h  |
| tidb/23365   | [planner: fix a bug that point get plan returns wrong column name](https://github.com/pingcap/tidb/pull/23365)                                                |   | 15d22h  |
| tidb/23373   | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                             |   | 15d19h  |
| tidb/23474   | [planner: fix inappropriate null flag of null constants (#23457)](https://github.com/pingcap/tidb/pull/23474)                                                 |   | 9d18h   |
| tidb/23537   | [planner: remove some risky cache operations in the plan builder (#23354)](https://github.com/pingcap/tidb/pull/23537)                                        |   | 7d19h   |
| tidb/23543   | [statistics: fix auto analyze log information incomplete (#23522)](https://github.com/pingcap/tidb/pull/23543)                                                |   | 7d18h   |
| tidb/23598   | [types: fix collation for binary literal (#23591)](https://github.com/pingcap/tidb/pull/23598)                                                                |   | 6d13h   |
| tidb/23655   | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral (#23523)](https://github.com/pingcap/tidb/pull/23655)                      |   | 2d22h   |
| tidb/23656   | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral (#23523)](https://github.com/pingcap/tidb/pull/23656)                      |   | 2d22h   |
| tidb/23683   | [executor: fix a panic when batch point get is used for partition table (#23652)](https://github.com/pingcap/tidb/pull/23683)                                 |   | 2d16h   |
| tidb/23689   | [planner: fix the panic when we calculate the partition range (#23651)](https://github.com/pingcap/tidb/pull/23689)                                           |   | 2d16h   |
| tidb/23728   | [planner: skip storage engine check for CRAETE VIEW statement](https://github.com/pingcap/tidb/pull/23728)                                                    |   | 1d19h   |
| tidb/23772   | [tablecodec: fix text type decode for old row format (#23751)](https://github.com/pingcap/tidb/pull/23772)                                                    |   | 1d11h   |


## Reviewed in Last 7 Days

|    REPO    |                                                                               PR                                                                               | C | D |   R   |
|------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/23816 | [planner: fix "can't find column" when projection wrongly added above table reader after agg pushed down (#23804)](https://github.com/pingcap/tidb/pull/23816) |   | 1 | 0h    |
| tidb/23804 | [planner: fix "can't find column" when projection wrongly added above table reader after agg pushed down](https://github.com/pingcap/tidb/pull/23804)          |   | 1 | 6h    |
| tidb/23647 | [planner, privileges: do not require SELECT for unqualified DELETE](https://github.com/pingcap/tidb/pull/23647)                                                |   | 1 | 2d12h |
| tidb/23769 | [planner: fix wrong IndexScan plan reused in plan cache (#23758)](https://github.com/pingcap/tidb/pull/23769)                                                  |   | 2 | 0h    |
| tidb/23758 | [planner: fix wrong IndexScan plan reused in plan cache](https://github.com/pingcap/tidb/pull/23758)                                                           |   | 2 | 1h    |
| tidb/23751 | [tablecodec: fix text type decode for old row format](https://github.com/pingcap/tidb/pull/23751)                                                              |   | 2 | 0h    |
| tidb/23674 | [*: add column `End_time` in show analyze status and add related log](https://github.com/pingcap/tidb/pull/23674)                                              |   | 2 | 21h   |
| tidb/23699 | [util: fix range building for binary literal](https://github.com/pingcap/tidb/pull/23699)                                                                      |   | 3 | 0h    |
| tidb/23685 | [planner: fix the issue that planner hints don't work in some batch/point-get plans (#23666)](https://github.com/pingcap/tidb/pull/23685)                      |   | 3 | 0h    |
| tidb/23666 | [planner: fix the issue that planner hints don't work in some batch/point-get plans](https://github.com/pingcap/tidb/pull/23666)                               |   | 3 | 0h    |
| tidb/23559 | [ranger: fix the range construction behavior when the column's type is `YEAR`](https://github.com/pingcap/tidb/pull/23559)                                     |   | 3 | 4d18h |
| tidb/23641 | [planner, util/ranger:  apply PushDownNot to condition before pruning partition (#23628)](https://github.com/pingcap/tidb/pull/23641)                          |   | 4 | 0h    |
| tidb/23628 | [planner, util/ranger:  apply PushDownNot to condition before pruning partition](https://github.com/pingcap/tidb/pull/23628)                                   |   | 4 | 0h    |
| tidb/23523 | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral](https://github.com/pingcap/tidb/pull/23523)                                |   | 4 | 4d5h  |
| tidb/23589 | [store, plan: make mpp workable when some node is not available shortly.](https://github.com/pingcap/tidb/pull/23589)                                          |   | 7 | 3h    |
| tidb/23597 | [planner: fix correlated columns in filter or access in MPP apply (#23509)](https://github.com/pingcap/tidb/pull/23597)                                        |   | 7 | 0h    |
| tidb/23435 | [planner: check schema stale for plan cache when forUpdateRead (#22381)](https://github.com/pingcap/tidb/pull/23435)                                           |   | 7 | 7d3h  |
| tidb/23591 | [types: fix collation for binary literal](https://github.com/pingcap/tidb/pull/23591)                                                                          |   | 7 | 0h    |


</details> 


<details> 
  <summary>wshwsh12's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                               PR                                                               | C | LASTED  |
|--------------|--------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19807   | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                        |   | 209d11h |
| docs-cn/5915 | [update default oom action of TiDB](https://github.com/pingcap/docs-cn/pull/5915)                                              |   | 18h     |
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                         | Y | 202d14h |
| tidb/21487   | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                  |   | 119d4h  |
| tidb/21887   | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                |   | 103d11h |
| tidb/22378   | [executor: vectorize hash aggregate](https://github.com/pingcap/tidb/pull/22378)                                               |   | 78d19h  |
| tidb/22628   | [executor: Improve max/min window function with deque-based sliding window](https://github.com/pingcap/tidb/pull/22628)        |   | 61d23h  |
| tidb/23336   | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23336)       |   | 16d19h  |
| tidb/23347   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                           |   | 16d18h  |
| tidb/23348   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                           |   | 16d18h  |
| tidb/23368   | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23368)          |   | 15d20h  |
| tidb/23369   | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23369)          |   | 15d20h  |
| tidb/23397   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23397)                                 |   | 14d17h  |
| tidb/23398   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23398)                                 |   | 14d17h  |
| tidb/23519   | [executor: check privilege before adding](https://github.com/pingcap/tidb/pull/23519)                                          |   | 8d0h    |
| tidb/23683   | [executor: fix a panic when batch point get is used for partition table (#23652)](https://github.com/pingcap/tidb/pull/23683)  |   | 2d16h   |
| tidb/23705   | [executor: refineArgs() bug fix when compare int with very small decimal (#23694)](https://github.com/pingcap/tidb/pull/23705) |   | 2d13h   |
| tidb/23749   | [planner/core: fix a bug during add cast for decimal join key  (#23723)](https://github.com/pingcap/tidb/pull/23749)           |   | 1d16h   |
| tidb/23760   | [collation: fix tidb panic when compare string with collation](https://github.com/pingcap/tidb/pull/23760)                     |   | 1d13h   |
| tidb/23782   | [collation: fix some string function has wrong collation and flag](https://github.com/pingcap/tidb/pull/23782)                 |   | 1d0h    |
| tidb/23796   | [tests: make TestIndexLookupMergeJoinHang and TestIssue18068 stable (#23741)](https://github.com/pingcap/tidb/pull/23796)      |   | 19h     |
| tidb/23819   | [plugin: fix audit plugin will cause tidb panic (#23803)](https://github.com/pingcap/tidb/pull/23819)                          |   | 2h      |
| tidb/23820   | [plugin: fix audit plugin will cause tidb panic (#23803)](https://github.com/pingcap/tidb/pull/23820)                          |   | 2h      |


## Reviewed in Last 7 Days

|     REPO      |                                                          PR                                                          | C | D |   R   |
|---------------|----------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/23493    | [expression: Implementation of Vitess hashing algorithm.](https://github.com/pingcap/tidb/pull/23493)                |   | 1 | 8d10h |
| tidb/23780    | [executor: fix 'index out of range' issue in index lookup join (#23755)](https://github.com/pingcap/tidb/pull/23780) |   | 1 | 6h    |
| community/426 | [promote zimulala to sig/exec committer](https://github.com/pingcap/community/pull/426)                              |   | 1 | 0h    |
| tidb/23578    | [telemetry: fix incorrect value for copr cache (#23577)](https://github.com/pingcap/tidb/pull/23578)                 |   | 7 | 0h    |
| tidb/23577    | [telemetry: fix incorrect value for copr cache](https://github.com/pingcap/tidb/pull/23577)                          |   | 7 | 0h    |


</details> 

