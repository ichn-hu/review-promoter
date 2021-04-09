|   REVIEWER    |  O(C)   | 1 | 2 | 3 | 4 | 5 | 6 | 7  | T  |
|---------------|---------|---|---|---|---|---|---|----|----|
| Reminiscent   |  6 ( 0) | 0 | 1 | 3 | 0 | 0 | 0 |  1 |  5 |
| SunRunAway    | 12 ( 1) | 0 | 0 | 0 | 0 | 0 | 0 |  0 |  0 |
| XuHuaiyu      | 73 ( 3) | 0 | 1 | 1 | 0 | 0 | 0 |  0 |  2 |
| dyzsr         |  1 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 |  0 |  0 |
| eurekaka      | 15 ( 0) | 0 | 1 | 0 | 0 | 0 | 0 | 10 | 11 |
| fzhedu        |  1 ( 0) | 0 | 2 | 0 | 0 | 0 | 0 |  0 |  2 |
| ichn-hu       |  2 ( 0) | 0 | 0 | 1 | 0 | 0 | 0 |  0 |  1 |
| lzmhhh123     | 49 ( 0) | 0 | 2 | 3 | 0 | 0 | 1 |  1 |  7 |
| nullnotnil    |  0 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 |  0 |  0 |
| qw4990        | 61 ( 2) | 0 | 3 | 7 | 0 | 0 | 0 |  1 | 11 |
| rebelice      |  5 ( 0) | 0 | 2 | 0 | 0 | 0 | 0 |  1 |  3 |
| time-and-fate |  5 ( 1) | 0 | 1 | 0 | 0 | 0 | 0 |  2 |  3 |
| winoros       | 40 ( 3) | 0 | 0 | 1 | 0 | 0 | 0 |  0 |  1 |
| wshwsh12      | 25 ( 1) | 0 | 4 | 2 | 0 | 0 | 0 |  0 |  6 |


<details> 
  <summary>Reminiscent's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                     PR                                                                     | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/21896 | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                                 |   | 108d19h |
| tidb/23474 | [planner: fix inappropriate null flag of null constants (#23457)](https://github.com/pingcap/tidb/pull/23474)                              |   | 16d18h  |
| tidb/23575 | [executor: fix update panic on join having statement (#23554)](https://github.com/pingcap/tidb/pull/23575)                                 |   | 13d21h  |
| tidb/23685 | [planner: fix the issue that planner hints don't work in some batch/point-get plans (#23666)](https://github.com/pingcap/tidb/pull/23685)  |   | 9d16h   |
| tidb/23917 | [planner: fix wrong TableDual plans caused by comparing Binary and Bytes incorrectly (#23860)](https://github.com/pingcap/tidb/pull/23917) |   | 23h     |
| tidb/23918 | [planner: fix wrong TableDual plans caused by comparing Binary and Bytes incorrectly (#23860)](https://github.com/pingcap/tidb/pull/23918) |   | 23h     |


## Reviewed in Last 7 Days

|    REPO    |                                                                PR                                                                 | C | D |   R   |
|------------|-----------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/23838 | [statistics: collect FMSketch when processing index analyze tasks](https://github.com/pingcap/tidb/pull/23838)                    |   | 2 | 5d1h  |
| tidb/23860 | [planner: fix wrong TableDual plans caused by comparing Binary and Bytes incorrectly](https://github.com/pingcap/tidb/pull/23860) |   | 3 | 4h    |
| tidb/23834 | [statistics: print log when tidb marks extended stats as deleted internally](https://github.com/pingcap/tidb/pull/23834)          |   | 3 | 3d18h |
| tidb/23844 | [planner: fix index-out-of-range error when checking only_full_group_by](https://github.com/pingcap/tidb/pull/23844)              |   | 3 | 2d18h |
| tidb/23830 | [statistics: fix a statistics GC problem that can cause duplicated fm-sketch records](https://github.com/pingcap/tidb/pull/23830) |   | 7 | 0h    |


</details> 


<details> 
  <summary>SunRunAway's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                  PR                                                                   | C | LASTED  |
|------------|---------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19178 | [executor: Refactor probe channel](https://github.com/pingcap/tidb/pull/19178)                                                        |   | 238d16h |
| tidb/19807 | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                               |   | 216d10h |
| tidb/19900 | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                        | Y | 211d18h |
| tidb/20140 | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                    |   | 198d22h |
| tidb/20765 | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                     |   | 157d17h |
| tidb/21207 | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                       |   | 136d19h |
| tidb/21834 | [planner: enhanced index range calculation plan](https://github.com/pingcap/tidb/pull/21834)                                          |   | 113d18h |
| tidb/21876 | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)           |   | 111d19h |
| tidb/21878 | [planner: do not push down lock to pointGet/bacthPointGet when selection exists](https://github.com/pingcap/tidb/pull/21878)          |   | 111d18h |
| tidb/21956 | [planner/preprocessor: disallow into-outfile clause in some place](https://github.com/pingcap/tidb/pull/21956)                        |   | 106d23h |
| tidb/22217 | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                         |   | 92d17h  |
| tidb/22379 | [[experiment] executor: allow aggregation to spill disk when running out of memory quota](https://github.com/pingcap/tidb/pull/22379) |   | 85d19h  |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>XuHuaiyu's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                              PR                                                                               | C | LASTED  |
|--------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| docs-cn/5619 | [Update data-type-date-and-time.md](https://github.com/pingcap/docs-cn/pull/5619)                                                                             |   | 41d15h  |
| docs-cn/5671 | [tidb: Add time format description](https://github.com/pingcap/docs-cn/pull/5671)                                                                             |   | 35d11h  |
| tidb/19900   | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                                                                | Y | 211d18h |
| docs-cn/5952 | [update system-variable for statements-summary](https://github.com/pingcap/docs-cn/pull/5952)                                                                 |   | 2d23h   |
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                                        | Y | 209d13h |
| tidb/20140   | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                                            |   | 198d22h |
| tidb/20311   | [expression: fix overflow error when convert bit to int64 (#20266)](https://github.com/pingcap/tidb/pull/20311)                                               |   | 190d21h |
| tidb/20790   | [collation: add pinyin collation for chinese charset support](https://github.com/pingcap/tidb/pull/20790)                                                     |   | 156d20h |
| tidb/21064   | [planner, executor: fix cast not check error](https://github.com/pingcap/tidb/pull/21064)                                                                     |   | 144d8h  |
| tidb/21149   | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                                 |   | 140d14h |
| tidb/21228   | [executor: return the result immediately when combining LIMIT row_count with DISTINCT](https://github.com/pingcap/tidb/pull/21228)                            |   | 136d13h |
| tidb/21304   | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                                      |   | 134d12h |
| tidb/21334   | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                                                 |   | 133d14h |
| tidb/21401   | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                                            |   | 129d11h |
| tidb/21476   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                         |   | 126d15h |
| tidb/21536   | [executor: add slow-log file meta cache to avoid repeat read file meta information](https://github.com/pingcap/tidb/pull/21536)                               |   | 122d14h |
| tidb/21564   | [ddl: fix Incorrect behavior of NO_ZERO_DATE when altering table](https://github.com/pingcap/tidb/pull/21564)                                                 |   | 121d15h |
| tidb/21896   | [planner: fix union doesn't handle collate correctly (#21854)](https://github.com/pingcap/tidb/pull/21896)                                                    |   | 108d19h |
| tidb/22131   | [privilege: remove leading and trailing space when create user and role](https://github.com/pingcap/tidb/pull/22131)                                          |   | 98d19h  |
| tidb/22163   | [expression: separated arithmeticMinusIntSig](https://github.com/pingcap/tidb/pull/22163)                                                                     |   | 94d13h  |
| tidb/22186   | [executor: fix select into outfile with year type column has no data (#22175)](https://github.com/pingcap/tidb/pull/22186)                                    |   | 93d16h  |
| tidb/22616   | [expression: from_unixtime accept 64-bit integers](https://github.com/pingcap/tidb/pull/22616)                                                                |   | 69d23h  |
| tidb/22617   | [metrics: fix wrong bucket name of coprocessor cache (#22454)](https://github.com/pingcap/tidb/pull/22617)                                                    |   | 69d23h  |
| tidb/22624   | [ planner: not pruning column used by union scan condition (#21640)](https://github.com/pingcap/tidb/pull/22624)                                              |   | 69d17h  |
| tidb/22631   | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                                                                               |   | 67d23h  |
| tidb/22696   | [expression: enable arithmetic Mod push down](https://github.com/pingcap/tidb/pull/22696)                                                                     |   | 64d17h  |
| tidb/22711   | [executor: Fix inline schema name](https://github.com/pingcap/tidb/pull/22711)                                                                                |   | 64d11h  |
| tidb/22722   | [planner, errno: make error code of ErrMixOfGroupFuncAndFields consistent with MySQL](https://github.com/pingcap/tidb/pull/22722)                             |   | 63d20h  |
| tidb/22814   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22814)                                           |   | 48d19h  |
| tidb/22908   | [txn: Add txn state's view](https://github.com/pingcap/tidb/pull/22908)                                                                                       |   | 43d20h  |
| tidb/23012   | [executor: fix affected rows of ddls and complete uint tests](https://github.com/pingcap/tidb/pull/23012)                                                     |   | 39d16h  |
| tidb/23152   | [expression: fix wrong error info (#22760)](https://github.com/pingcap/tidb/pull/23152)                                                                       |   | 32d14h  |
| tidb/23196   | [types: fix the bug about the wrong query result for decimal type  (#22507)](https://github.com/pingcap/tidb/pull/23196)                                      |   | 30d18h  |
| tidb/23220   | [Release 4.0](https://github.com/pingcap/tidb/pull/23220)                                                                                                     |   | 30d11h  |
| tidb/23227   | [executor: hash join out of index panic when enum column value is zero (#23162)](https://github.com/pingcap/tidb/pull/23227)                                  |   | 29d22h  |
| tidb/23233   | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23233)                                                        |   | 29d18h  |
| tidb/23257   | [executor: group_concat aggr panic when session.group_concat_max_len is small (#23131)](https://github.com/pingcap/tidb/pull/23257)                           |   | 28d18h  |
| tidb/23295   | [util, types: don't let SPM be affected by charset (#23161)](https://github.com/pingcap/tidb/pull/23295)                                                      |   | 27d11h  |
| tidb/23335   | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23335)                                      |   | 23d19h  |
| tidb/23336   | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23336)                                      |   | 23d19h  |
| tidb/23347   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                                          |   | 23d18h  |
| tidb/23348   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                                          |   | 23d18h  |
| tidb/23350   | [util/stringutil, util/ranger, planner: use hierarchical separators to simplify the parsing for info of EXPLAIN ](https://github.com/pingcap/tidb/pull/23350) |   | 23d17h  |
| tidb/23368   | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23368)                                         |   | 22d20h  |
| tidb/23369   | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23369)                                         |   | 22d20h  |
| tidb/23397   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23397)                                                                |   | 21d17h  |
| tidb/23398   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23398)                                                                |   | 21d17h  |
| tidb/23405   | [domain: remove the exit chan, use context](https://github.com/pingcap/tidb/pull/23405)                                                                       |   | 21d17h  |
| tidb/23433   | [WIP: speed up for slow query logs retrieving ](https://github.com/pingcap/tidb/pull/23433)                                                                   |   | 20d17h  |
| tidb/23474   | [planner: fix inappropriate null flag of null constants (#23457)](https://github.com/pingcap/tidb/pull/23474)                                                 |   | 16d18h  |
| tidb/23487   | [planner: optimize count(distinct a) to count(a) if there is an unique key on a](https://github.com/pingcap/tidb/pull/23487)                                  | Y | 16d14h  |
| tidb/23497   | [expression: Let TiDB use Hyperscan to support multi-pattern-match](https://github.com/pingcap/tidb/pull/23497)                                               |   | 15d22h  |
| tidb/23517   | [*: Add the metric about the SQL with TiFlash Success  (#23426)](https://github.com/pingcap/tidb/pull/23517)                                                  |   | 15d12h  |
| tidb/23562   | [execution: reuse iterator in hash join](https://github.com/pingcap/tidb/pull/23562)                                                                          |   | 14d13h  |
| tidb/23640   | [*: fix the bug about YEAR(0.9) returns NULL instead of 0 in NO_ZERO_DATE mode](https://github.com/pingcap/tidb/pull/23640)                                   |   | 10d13h  |
| tidb/23661   | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23661)       |   | 9d20h   |
| tidb/23682   | [executor: fix a panic when batch point get is used for partition table (#23652)](https://github.com/pingcap/tidb/pull/23682)                                 |   | 9d16h   |
| tidb/23683   | [executor: fix a panic when batch point get is used for partition table (#23652)](https://github.com/pingcap/tidb/pull/23683)                                 |   | 9d16h   |
| tidb/23691   | [executor: fix index join on prefix column index (#23678)](https://github.com/pingcap/tidb/pull/23691)                                                        |   | 9d15h   |
| tidb/23705   | [executor: refineArgs() bug fix when compare int with very small decimal (#23694)](https://github.com/pingcap/tidb/pull/23705)                                |   | 9d13h   |
| tidb/23756   | [planner: fix set not null flag for outer join (#23727)](https://github.com/pingcap/tidb/pull/23756)                                                          |   | 8d14h   |
| tidb/23757   | [planner: fix set not null flag for outer join (#23727)](https://github.com/pingcap/tidb/pull/23757)                                                          |   | 8d14h   |
| tidb/23812   | [executor, planner: fix collation for hash join building (#23770)](https://github.com/pingcap/tidb/pull/23812)                                                |   | 7d12h   |
| tidb/23826   | [executor: fix 2nd index dup check after insert ignore on dup update primary (#23814)](https://github.com/pingcap/tidb/pull/23826)                            |   | 6d20h   |
| tidb/23857   | [types: fix type merge about bit type](https://github.com/pingcap/tidb/pull/23857)                                                                            |   | 2d19h   |
| tidb/23862   | [statistics: make the global-level topN more accurate](https://github.com/pingcap/tidb/pull/23862)                                                            |   | 2d17h   |
| tidb/23868   | [store, kv: Add information about async commit/1pc to tidb_last_txn_info (#23833)](https://github.com/pingcap/tidb/pull/23868)                                |   | 2d16h   |
| tidb/23878   | [functions: fix some string function has wrong collation and flag (#23835)](https://github.com/pingcap/tidb/pull/23878)                                       |   | 1d21h   |
| tidb/23879   | [functions: fix some string function has wrong collation and flag (#23835)](https://github.com/pingcap/tidb/pull/23879)                                       |   | 1d21h   |
| tidb/23884   | [Metric: Collect TiKV Read Metric for SLI/SLO](https://github.com/pingcap/tidb/pull/23884)                                                                    |   | 1d19h   |
| tidb/23888   | [executor: fix resource leak of Shuffle Executor.](https://github.com/pingcap/tidb/pull/23888)                                                                |   | 1d18h   |
| tidb/23896   | [executor: fix `show table status` for the database with upper-cased name](https://github.com/pingcap/tidb/pull/23896)                                        |   | 1d17h   |
| tidb/23918   | [planner: fix wrong TableDual plans caused by comparing Binary and Bytes incorrectly (#23860)](https://github.com/pingcap/tidb/pull/23918)                    |   | 23h     |


## Reviewed in Last 7 Days

|     REPO     |                                                  PR                                                   | C | D |   R   |
|--------------|-------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/23867   | [expression: fix wrong flen infer for bit constant](https://github.com/pingcap/tidb/pull/23867)       |   | 2 | 16h   |
| docs-cn/5835 | [Update 5.0 GA release notes and experimental features](https://github.com/pingcap/docs-cn/pull/5835) |   | 3 | 12d4h |


</details> 


<details> 
  <summary>dyzsr's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                             PR                                                             | C | LASTED |
|------------|----------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/23559 | [ranger: fix the range construction behavior when the column's type is `YEAR`](https://github.com/pingcap/tidb/pull/23559) |   | 14d14h |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>eurekaka's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                    PR                                                                     | C | LASTED  |
|------------|-------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                                 |   | 154d16h |
| docs/5150  | [SPM: update DML SQL Bind and baseline capture description (#5088)](https://github.com/pingcap/docs/pull/5150)                            |   | 10d15h  |
| tidb/23283 | [util: optimize the performance of restore with db (#22910)](https://github.com/pingcap/tidb/pull/23283)                                  |   | 27d17h  |
| tidb/23316 | [planner: Fix rebuild range for prepared plan](https://github.com/pingcap/tidb/pull/23316)                                                |   | 24d17h  |
| tidb/23373 | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                         |   | 22d19h  |
| tidb/23543 | [statistics: fix auto analyze log information incomplete (#23522)](https://github.com/pingcap/tidb/pull/23543)                            |   | 14d18h  |
| tidb/23685 | [planner: fix the issue that planner hints don't work in some batch/point-get plans (#23666)](https://github.com/pingcap/tidb/pull/23685) |   | 9d16h   |
| tidb/23689 | [planner: fix the panic when we calculate the partition range (#23651)](https://github.com/pingcap/tidb/pull/23689)                       |   | 9d16h   |
| tidb/23705 | [executor: refineArgs() bug fix when compare int with very small decimal (#23694)](https://github.com/pingcap/tidb/pull/23705)            |   | 9d13h   |
| tidb/23749 | [planner/core: fix a bug during add cast for decimal join key  (#23723)](https://github.com/pingcap/tidb/pull/23749)                      |   | 8d16h   |
| tidb/23756 | [planner: fix set not null flag for outer join (#23727)](https://github.com/pingcap/tidb/pull/23756)                                      |   | 8d14h   |
| tidb/23757 | [planner: fix set not null flag for outer join (#23727)](https://github.com/pingcap/tidb/pull/23757)                                      |   | 8d14h   |
| tidb/23760 | [collation: fix tidb panic when compare string with collation](https://github.com/pingcap/tidb/pull/23760)                                |   | 8d13h   |
| tidb/23883 | [*: don't allocate SessionIndexUsageCollector when indexUsageLease equals 0 (#23861)](https://github.com/pingcap/tidb/pull/23883)         |   | 1d20h   |
| tidb/23911 | [planner,privilege: requires extra privileges for REPLACE and INSERT ON DUPLICATE statements](https://github.com/pingcap/tidb/pull/23911) |   | 1d7h    |


## Reviewed in Last 7 Days

|    REPO    |                                                               PR                                                                | C | D |    R    |
|------------|---------------------------------------------------------------------------------------------------------------------------------|---|---|---------|
| tidb/23718 | [*: add TableSample ID for PhysicalIDToTypeString()](https://github.com/pingcap/tidb/pull/23718)                                |   | 2 | 7d5h    |
| tidb/21444 | [planner: ignore anonymous index while tiflash replica is available](https://github.com/pingcap/tidb/pull/21444)                |   | 7 | 120d19h |
| tidb/22416 | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                               | Y | 7 | 75d19h  |
| tidb/22559 | [planner: split test data from test cases in cbo_test.go](https://github.com/pingcap/tidb/pull/22559)                           |   | 7 | 65d2h   |
| tidb/22853 | [planner: fix LogicalPlans that contain Window Function are ambiguous ](https://github.com/pingcap/tidb/pull/22853)             |   | 7 | 40d19h  |
| tidb/23137 | [planner: fix index merge row count estimation logic](https://github.com/pingcap/tidb/pull/23137)                               |   | 7 | 28d0h   |
| tidb/23208 | [statistics, util/ranger: improve selectivity calculation for DNF filters (#18741)](https://github.com/pingcap/tidb/pull/23208) |   | 7 | 23d23h  |
| tidb/23295 | [util, types: don't let SPM be affected by charset (#23161)](https://github.com/pingcap/tidb/pull/23295)                        |   | 7 | 20d18h  |
| tidb/23365 | [planner: fix a bug that point get plan returns wrong column name](https://github.com/pingcap/tidb/pull/23365)                  |   | 7 | 16d4h   |
| tidb/23575 | [executor: fix update panic on join having statement (#23554)](https://github.com/pingcap/tidb/pull/23575)                      |   | 7 | 7d4h    |
| tidb/23831 | [statistic: solve write write conflicts of mysql.stats_histograms](https://github.com/pingcap/tidb/pull/23831)                  |   | 7 | 1h      |


</details> 


<details> 
  <summary>fzhedu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                          PR                                                          | C | LASTED |
|------------|----------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/23749 | [planner/core: fix a bug during add cast for decimal join key  (#23723)](https://github.com/pingcap/tidb/pull/23749) |   | 8d16h  |


## Reviewed in Last 7 Days

|    REPO    |                                                    PR                                                    | C | D |   R   |
|------------|----------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/23747 | [planner, sessionvar: avoid sending same task id to TiFlash](https://github.com/pingcap/tidb/pull/23747) |   | 2 | 6d18h |
| tics/1715  | [Do not allow columns with same name in TiFlash's Block](https://github.com/pingcap/tics/pull/1715)      |   | 2 | 5d22h |


</details> 


<details> 
  <summary>ichn-hu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                               PR                                                | C | LASTED |
|------------|-------------------------------------------------------------------------------------------------|---|--------|
| tidb/23857 | [types: fix type merge about bit type](https://github.com/pingcap/tidb/pull/23857)              |   | 2d19h  |
| tidb/23867 | [expression: fix wrong flen infer for bit constant](https://github.com/pingcap/tidb/pull/23867) |   | 2d16h  |


## Reviewed in Last 7 Days

|    REPO    |                                                   PR                                                   | C | D |  R   |
|------------|--------------------------------------------------------------------------------------------------------|---|---|------|
| tidb/23691 | [executor: fix index join on prefix column index (#23678)](https://github.com/pingcap/tidb/pull/23691) |   | 3 | 7d0h |


</details> 


<details> 
  <summary>lzmhhh123's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                             PR                                                                              | C | LASTED  |
|------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20444 | [expression: add json_merge_patch](https://github.com/pingcap/tidb/pull/20444)                                                                              |   | 176d21h |
| tidb/20465 | [expression: add uuidShortFunction](https://github.com/pingcap/tidb/pull/20465)                                                                             |   | 175d19h |
| tidb/20642 | [executor: modify admin executors to support partitioned table with global index](https://github.com/pingcap/tidb/pull/20642)                               |   | 164d15h |
| tidb/20903 | [planner: fix confused and unnecessary double-projection in plans.](https://github.com/pingcap/tidb/pull/20903)                                             |   | 153d17h |
| tidb/21018 | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                              |   | 147d17h |
| tidb/21195 | [brie: integrate lightning to suport IMPORT statement](https://github.com/pingcap/tidb/pull/21195)                                                          |   | 136d22h |
| tidb/21334 | [*: make rollback work on user-defined variables](https://github.com/pingcap/tidb/pull/21334)                                                               |   | 133d14h |
| tidb/21347 | [session: make rollback work on global variables](https://github.com/pingcap/tidb/pull/21347)                                                               |   | 132d19h |
| tidb/21487 | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                                               |   | 126d4h  |
| tidb/21641 | [executor: Fix pessimistic lock doesn't work on the partition table for subquery/joins](https://github.com/pingcap/tidb/pull/21641)                         |   | 119d18h |
| tidb/21651 | [planner: allow filter condition pushing down to IndexScan for prefix index](https://github.com/pingcap/tidb/pull/21651)                                    |   | 119d13h |
| tidb/22126 | [*: add `sys` schema, `sys.SCHEMA_UNUSED_INDEXES` view and `sys.SCHEMA_INDEX_USAGE` view](https://github.com/pingcap/tidb/pull/22126)                       |   | 98d19h  |
| tidb/22361 | [table: fix insert into _tidb_rowid panic and rebase it if needed (#22062)](https://github.com/pingcap/tidb/pull/22361)                                     |   | 86d20h  |
| tidb/22372 | [executor: fix SelectForUpdate in decorrelated subquery under pessimistic mode](https://github.com/pingcap/tidb/pull/22372)                                 |   | 86d9h   |
| tidb/22478 | [planner, executor: fix query partition table with global unique index get wrong result](https://github.com/pingcap/tidb/pull/22478)                        |   | 77d13h  |
| tidb/22631 | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                                                                             |   | 67d23h  |
| tidb/22686 | [expression: support enum pushdown](https://github.com/pingcap/tidb/pull/22686)                                                                             |   | 64d22h  |
| tidb/22699 | [brie: add error info column and history backup/restore info in sql](https://github.com/pingcap/tidb/pull/22699)                                            |   | 64d16h  |
| tidb/23001 | [statistics: fix err check](https://github.com/pingcap/tidb/pull/23001)                                                                                     |   | 40d0h   |
| tidb/23022 | [executor: create PipelinedWindowExec based on current implementation and modify the windowProcessor interface](https://github.com/pingcap/tidb/pull/23022) |   | 38d18h  |
| tidb/23149 | [core: support left join and right join for join reorder](https://github.com/pingcap/tidb/pull/23149)                                                       |   | 33d12h  |
| tidb/23257 | [executor: group_concat aggr panic when session.group_concat_max_len is small (#23131)](https://github.com/pingcap/tidb/pull/23257)                         |   | 28d18h  |
| tidb/23283 | [util: optimize the performance of restore with db (#22910)](https://github.com/pingcap/tidb/pull/23283)                                                    |   | 27d17h  |
| tidb/23347 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                                        |   | 23d18h  |
| tidb/23348 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                                        |   | 23d18h  |
| tidb/23368 | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23368)                                       |   | 22d20h  |
| tidb/23369 | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23369)                                       |   | 22d20h  |
| tidb/23373 | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                           |   | 22d19h  |
| tidb/23559 | [ranger: fix the range construction behavior when the column's type is `YEAR`](https://github.com/pingcap/tidb/pull/23559)                                  |   | 14d14h  |
| tidb/23655 | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral (#23523)](https://github.com/pingcap/tidb/pull/23655)                    |   | 9d22h   |
| tidb/23656 | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral (#23523)](https://github.com/pingcap/tidb/pull/23656)                    |   | 9d22h   |
| tidb/23660 | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23660)     |   | 9d20h   |
| tidb/23661 | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23661)     |   | 9d20h   |
| tidb/23703 | [expression: fix approx_percent panic on bit column (#23687)](https://github.com/pingcap/tidb/pull/23703)                                                   |   | 9d13h   |
| tidb/23705 | [executor: refineArgs() bug fix when compare int with very small decimal (#23694)](https://github.com/pingcap/tidb/pull/23705)                              |   | 9d13h   |
| tidb/23714 | [*:Support record statment_history table evicted info](https://github.com/pingcap/tidb/pull/23714)                                                          |   | 9d1h    |
| tidb/23733 | [planner: generate BatchPointGet for hash table partitions](https://github.com/pingcap/tidb/pull/23733)                                                     |   | 8d18h   |
| tidb/23738 | [executor: remove duplicate entry in show privileges](https://github.com/pingcap/tidb/pull/23738)                                                           |   | 8d17h   |
| tidb/23747 | [planner, sessionvar: avoid sending same task id to TiFlash](https://github.com/pingcap/tidb/pull/23747)                                                    |   | 8d16h   |
| tidb/23749 | [planner/core: fix a bug during add cast for decimal join key  (#23723)](https://github.com/pingcap/tidb/pull/23749)                                        |   | 8d16h   |
| tidb/23756 | [planner: fix set not null flag for outer join (#23727)](https://github.com/pingcap/tidb/pull/23756)                                                        |   | 8d14h   |
| tidb/23757 | [planner: fix set not null flag for outer join (#23727)](https://github.com/pingcap/tidb/pull/23757)                                                        |   | 8d14h   |
| tidb/23760 | [collation: fix tidb panic when compare string with collation](https://github.com/pingcap/tidb/pull/23760)                                                  |   | 8d13h   |
| tidb/23812 | [executor, planner: fix collation for hash join building (#23770)](https://github.com/pingcap/tidb/pull/23812)                                              |   | 7d12h   |
| tidb/23818 | [*: protect read only noop via tidb_enable_noop_functions](https://github.com/pingcap/tidb/pull/23818)                                                      |   | 7d5h    |
| tidb/23822 | [statistics: feedback not panic when no ndv collected (#23808)](https://github.com/pingcap/tidb/pull/23822)                                                 |   | 6d22h   |
| tidb/23859 | [draft:store/tikv:remove tidb/kv from package tikv/unionstore](https://github.com/pingcap/tidb/pull/23859)                                                  |   | 2d18h   |
| tidb/23871 | [store/tikv: move BufferBatchGetter from tidb/kv to tikv/unionstore](https://github.com/pingcap/tidb/pull/23871)                                            |   | 2d15h   |
| tidb/23902 | [expression: fix comparing year with datetime for equality](https://github.com/pingcap/tidb/pull/23902)                                                     |   | 1d15h   |


## Reviewed in Last 7 Days

|    REPO    |                                                            PR                                                             | C | D |   R   |
|------------|---------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/23888 | [executor: fix resource leak of Shuffle Executor.](https://github.com/pingcap/tidb/pull/23888)                            |   | 2 | 0h    |
| tidb/23885 | [executor,distsql: clean up useless interface](https://github.com/pingcap/tidb/pull/23885)                                |   | 2 | 0h    |
| tidb/23796 | [tests: make TestIndexLookupMergeJoinHang and TestIssue18068 stable (#23741)](https://github.com/pingcap/tidb/pull/23796) |   | 3 | 5d3h  |
| tidb/23700 | [tikv: distinguish server timeout for TiKV and TiFlash](https://github.com/pingcap/tidb/pull/23700)                       |   | 3 | 6d18h |
| tidb/23844 | [planner: fix index-out-of-range error when checking only_full_group_by](https://github.com/pingcap/tidb/pull/23844)      |   | 3 | 2d19h |
| tidb/23845 | [types/json: replace binary.Read with binary.BigEndian.Uint16](https://github.com/pingcap/tidb/pull/23845)                |   | 6 | 12h   |
| tidb/23680 | [*: add test for modifying default length of cast as decimal](https://github.com/pingcap/tidb/pull/23680)                 |   | 7 | 3d0h  |


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
| docs-cn/5561 | [Add sql optimization-related docs to toc](https://github.com/pingcap/docs-cn/pull/5561)                                                                    |   | 45d15h  |
| tidb/19029   | [types: fix unexpected NOT_NULL flags](https://github.com/pingcap/tidb/pull/19029)                                                                          |   | 245d22h |
| docs-cn/5785 | [update SPM documentation for DML SQL Bind and baseline capture (#5740)](https://github.com/pingcap/docs-cn/pull/5785)                                      |   | 17d18h  |
| tidb/20708   | [*: separate auto_increment ID allocator from _tidb_rowid allocator](https://github.com/pingcap/tidb/pull/20708)                                            |   | 161d20h |
| tidb/20969   | [executor: Improve the performance of appending not fixed columns](https://github.com/pingcap/tidb/pull/20969)                                              |   | 149d9h  |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                              |   | 147d17h |
| tidb/21149   | [executor:Add runtime stat for IndexMergeReaderExecutor (#20653)](https://github.com/pingcap/tidb/pull/21149)                                               |   | 140d14h |
| tidb/21304   | [executor: Add the HashAggExec runtime information (#20577)](https://github.com/pingcap/tidb/pull/21304)                                                    |   | 134d12h |
| tidb/21318   | [planner, expression: use the range of column types to simplify expressions](https://github.com/pingcap/tidb/pull/21318)                                    |   | 133d19h |
| tidb/21401   | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                                                          |   | 129d11h |
| tidb/21476   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                       |   | 126d15h |
| tidb/21508   | [execution: fix dayofweek('0000-00-00') behavior](https://github.com/pingcap/tidb/pull/21508)                                                               |   | 125d10h |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                                 |   | 111d19h |
| tidb/21887   | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                                             |   | 110d11h |
| tidb/21954   | [planner/cascades: add rule `PushSelDownApply`](https://github.com/pingcap/tidb/pull/21954)                                                                 |   | 106d23h |
| tidb/22146   | [executor: forbid SFU on view](https://github.com/pingcap/tidb/pull/22146)                                                                                  |   | 94d21h  |
| tidb/22217   | [*: rewrite origin SQL with default DB for SQL bindings (#21275)](https://github.com/pingcap/tidb/pull/22217)                                               |   | 92d17h  |
| tidb/22234   | [executor, planner: ON DUPLICATE UPDATE can refer to un-project col (#14412)](https://github.com/pingcap/tidb/pull/22234)                                   |   | 92d15h  |
| tidb/22261   | [time: fix parse datetime won't truncate the reluctant string (#22232)](https://github.com/pingcap/tidb/pull/22261)                                         |   | 91d19h  |
| tidb/22374   | [expression: separated arithmeticIntDivideSig](https://github.com/pingcap/tidb/pull/22374)                                                                  |   | 86d0h   |
| tidb/22415   | [ddl: refactor placement package](https://github.com/pingcap/tidb/pull/22415)                                                                               |   | 82d17h  |
| tidb/22416   | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                                                           | Y | 82d11h  |
| tidb/22541   | [expression: Support builtin function SOUNDEX](https://github.com/pingcap/tidb/pull/22541)                                                                  |   | 72d9h   |
| tidb/22565   | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)                                          | Y | 71d17h  |
| tidb/22814   | [expression: fix enum and set type expression in where clause (#22785)](https://github.com/pingcap/tidb/pull/22814)                                         |   | 48d19h  |
| tidb/22862   | [brie: fix the problem that ddl restored by BR via SQL is not replicated to downstream](https://github.com/pingcap/tidb/pull/22862)                         |   | 45d22h  |
| tidb/22923   | [expression: correct constant propagation for collation (#22666)](https://github.com/pingcap/tidb/pull/22923)                                               |   | 43d15h  |
| tidb/22924   | [planner: fix wrong index merge selection (#22825)](https://github.com/pingcap/tidb/pull/22924)                                                             |   | 43d14h  |
| tidb/22984   | [executor: fix logging format of prepared statements (#16062)](https://github.com/pingcap/tidb/pull/22984)                                                  |   | 40d10h  |
| tidb/23022   | [executor: create PipelinedWindowExec based on current implementation and modify the windowProcessor interface](https://github.com/pingcap/tidb/pull/23022) |   | 38d18h  |
| tidb/23152   | [expression: fix wrong error info (#22760)](https://github.com/pingcap/tidb/pull/23152)                                                                     |   | 32d14h  |
| tidb/23196   | [types: fix the bug about the wrong query result for decimal type  (#22507)](https://github.com/pingcap/tidb/pull/23196)                                    |   | 30d18h  |
| tidb/23208   | [statistics, util/ranger: improve selectivity calculation for DNF filters (#18741)](https://github.com/pingcap/tidb/pull/23208)                             |   | 30d16h  |
| tidb/23295   | [util, types: don't let SPM be affected by charset (#23161)](https://github.com/pingcap/tidb/pull/23295)                                                    |   | 27d11h  |
| tidb/23316   | [planner: Fix rebuild range for prepared plan](https://github.com/pingcap/tidb/pull/23316)                                                                  |   | 24d17h  |
| tidb/23373   | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                           |   | 22d19h  |
| tidb/23397   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23397)                                                              |   | 21d17h  |
| tidb/23398   | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23398)                                                              |   | 21d17h  |
| tidb/23448   | [wip :execution: parallel build hash table](https://github.com/pingcap/tidb/pull/23448)                                                                     |   | 18d12h  |
| tidb/23590   | [planner, table: optimize the list partition pruner for range query](https://github.com/pingcap/tidb/pull/23590)                                            |   | 13d16h  |
| tidb/23598   | [types: fix collation for binary literal (#23591)](https://github.com/pingcap/tidb/pull/23598)                                                              |   | 13d13h  |
| tidb/23655   | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral (#23523)](https://github.com/pingcap/tidb/pull/23655)                    |   | 9d22h   |
| tidb/23656   | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral (#23523)](https://github.com/pingcap/tidb/pull/23656)                    |   | 9d22h   |
| tidb/23658   | [*: collect transaction write duration/throughput metrics for SLI/SLO (#23462)](https://github.com/pingcap/tidb/pull/23658)                                 |   | 9d22h   |
| tidb/23660   | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23660)     |   | 9d20h   |
| tidb/23661   | [expression: Maintain separate scalar function pushdown lists for each engine instead of unified. (#23284)](https://github.com/pingcap/tidb/pull/23661)     |   | 9d20h   |
| tidb/23674   | [*: add column `End_time` in show analyze status and add related log](https://github.com/pingcap/tidb/pull/23674)                                           |   | 9d17h   |
| tidb/23682   | [executor: fix a panic when batch point get is used for partition table (#23652)](https://github.com/pingcap/tidb/pull/23682)                               |   | 9d16h   |
| tidb/23689   | [planner: fix the panic when we calculate the partition range (#23651)](https://github.com/pingcap/tidb/pull/23689)                                         |   | 9d16h   |
| tidb/23730   | [distsql/*: typo fix for `dispatches`](https://github.com/pingcap/tidb/pull/23730)                                                                          |   | 8d18h   |
| tidb/23733   | [planner: generate BatchPointGet for hash table partitions](https://github.com/pingcap/tidb/pull/23733)                                                     |   | 8d18h   |
| tidb/23796   | [tests: make TestIndexLookupMergeJoinHang and TestIssue18068 stable (#23741)](https://github.com/pingcap/tidb/pull/23796)                                   |   | 7d19h   |
| tidb/23812   | [executor, planner: fix collation for hash join building (#23770)](https://github.com/pingcap/tidb/pull/23812)                                              |   | 7d12h   |
| tidb/23831   | [statistic: solve write write conflicts of mysql.stats_histograms](https://github.com/pingcap/tidb/pull/23831)                                              |   | 6d18h   |
| tidb/23854   | [store/tikv: move CommitDetails and LockKeysDetails to tikv/util](https://github.com/pingcap/tidb/pull/23854)                                               |   | 2d21h   |
| tidb/23867   | [expression: fix wrong flen infer for bit constant](https://github.com/pingcap/tidb/pull/23867)                                                             |   | 2d16h   |
| tidb/23878   | [functions: fix some string function has wrong collation and flag (#23835)](https://github.com/pingcap/tidb/pull/23878)                                     |   | 1d21h   |
| tidb/23879   | [functions: fix some string function has wrong collation and flag (#23835)](https://github.com/pingcap/tidb/pull/23879)                                     |   | 1d21h   |
| tidb/23890   | [statistics: check step overflow when converting a range to points for estimation](https://github.com/pingcap/tidb/pull/23890)                              |   | 1d18h   |
| tidb/23895   | [statistics: add more tests for extended stats](https://github.com/pingcap/tidb/pull/23895)                                                                 |   | 1d17h   |
| tidb/23918   | [planner: fix wrong TableDual plans caused by comparing Binary and Bytes incorrectly (#23860)](https://github.com/pingcap/tidb/pull/23918)                  |   | 23h     |


## Reviewed in Last 7 Days

|     REPO     |                                                                PR                                                                 | C | D |   R   |
|--------------|-----------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/23862   | [statistics: make the global-level topN more accurate](https://github.com/pingcap/tidb/pull/23862)                                |   | 2 | 1d0h  |
| tidb/23883   | [*: don't allocate SessionIndexUsageCollector when indexUsageLease equals 0 (#23861)](https://github.com/pingcap/tidb/pull/23883) |   | 2 | 0h    |
| tidb/23861   | [*: don't allocate SessionIndexUsageCollector when indexUsageLease equals 0](https://github.com/pingcap/tidb/pull/23861)          |   | 2 | 18h   |
| tidb/23543   | [statistics: fix auto analyze log information incomplete (#23522)](https://github.com/pingcap/tidb/pull/23543)                    |   | 3 | 12d0h |
| tidb/23718   | [*: add TableSample ID for PhysicalIDToTypeString()](https://github.com/pingcap/tidb/pull/23718)                                  |   | 3 | 6d3h  |
| tidb/23728   | [planner: skip storage engine check for CRAETE VIEW statement](https://github.com/pingcap/tidb/pull/23728)                        |   | 3 | 5d23h |
| tidb/23733   | [planner: generate BatchPointGet for hash table partitions](https://github.com/pingcap/tidb/pull/23733)                           |   | 3 | 5d22h |
| tidb/23828   | [executor: do not show extended stats for dropped columns](https://github.com/pingcap/tidb/pull/23828)                            |   | 3 | 3d23h |
| tidb/23822   | [statistics: feedback not panic when no ndv collected (#23808)](https://github.com/pingcap/tidb/pull/23822)                       |   | 3 | 4d2h  |
| tidb/23834   | [statistics: print log when tidb marks extended stats as deleted internally](https://github.com/pingcap/tidb/pull/23834)          |   | 3 | 3d18h |
| docs-cn/5806 | [releases: add tidb release notes 4.0.12](https://github.com/pingcap/docs-cn/pull/5806)                                           |   | 7 | 9d15h |


</details> 


<details> 
  <summary>rebelice's detailed review report</summary> 

## To Be Reviewed

|    REPO     |                                                                    PR                                                                     | C | LASTED |
|-------------|-------------------------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/23537  | [planner: remove some risky cache operations in the plan builder (#23354)](https://github.com/pingcap/tidb/pull/23537)                    |   | 14d19h |
| parser/1205 | [Implement force_index hint in parser and TiDB](https://github.com/pingcap/parser/pull/1205)                                              |   | 6d17h  |
| tidb/23685  | [planner: fix the issue that planner hints don't work in some batch/point-get plans (#23666)](https://github.com/pingcap/tidb/pull/23685) |   | 9d16h  |
| tidb/23836  | [parser, core: Implement force_index hint in parser and TiDB](https://github.com/pingcap/tidb/pull/23836)                                 |   | 6d17h  |
| tidb/23838  | [statistics: collect FMSketch when processing index analyze tasks](https://github.com/pingcap/tidb/pull/23838)                            |   | 6d16h  |


## Reviewed in Last 7 Days

|    REPO    |                                                                PR                                                                 | C | D |  R   |
|------------|-----------------------------------------------------------------------------------------------------------------------------------|---|---|------|
| tidb/23674 | [*: add column `End_time` in show analyze status and add related log](https://github.com/pingcap/tidb/pull/23674)                 |   | 2 | 8d3h |
| tidb/23862 | [statistics: make the global-level topN more accurate](https://github.com/pingcap/tidb/pull/23862)                                |   | 2 | 1d0h |
| tidb/23830 | [statistics: fix a statistics GC problem that can cause duplicated fm-sketch records](https://github.com/pingcap/tidb/pull/23830) |   | 7 | 0h   |


</details> 


<details> 
  <summary>time-and-fate's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                               PR                                                               | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20877 | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                      |   | 154d16h |
| tidb/22416 | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                              | Y | 82d11h  |
| tidb/23733 | [planner: generate BatchPointGet for hash table partitions](https://github.com/pingcap/tidb/pull/23733)                        |   | 8d18h   |
| tidb/23890 | [statistics: check step overflow when converting a range to points for estimation](https://github.com/pingcap/tidb/pull/23890) |   | 1d18h   |
| tidb/23895 | [statistics: add more tests for extended stats](https://github.com/pingcap/tidb/pull/23895)                                    |   | 1d17h   |


## Reviewed in Last 7 Days

|    REPO    |                                                     PR                                                      | C | D |   R   |
|------------|-------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/23828 | [executor: do not show extended stats for dropped columns](https://github.com/pingcap/tidb/pull/23828)      |   | 2 | 4d21h |
| tidb/23728 | [planner: skip storage engine check for CRAETE VIEW statement](https://github.com/pingcap/tidb/pull/23728)  |   | 7 | 2d2h  |
| tidb/23822 | [statistics: feedback not panic when no ndv collected (#23808)](https://github.com/pingcap/tidb/pull/23822) |   | 7 | 1h    |


</details> 


<details> 
  <summary>winoros's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                              PR                                                                               | C | LASTED  |
|--------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19957   | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                                        | Y | 209d13h |
| docs-cn/5916 | [sql-statements, information-schema: add `END_TIME` field for table `ANALYZE_STATUS`](https://github.com/pingcap/docs-cn/pull/5916)                           |   | 7d17h   |
| tidb/20311   | [expression: fix overflow error when convert bit to int64 (#20266)](https://github.com/pingcap/tidb/pull/20311)                                               |   | 190d21h |
| tidb/20765   | [planner: support stable result mode](https://github.com/pingcap/tidb/pull/20765)                                                                             |   | 157d17h |
| tidb/20877   | [statistics: collect index usage information](https://github.com/pingcap/tidb/pull/20877)                                                                     |   | 154d16h |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                                |   | 147d17h |
| tidb/21207   | [planner: fix the inappropriate out-of-range range estimation rule](https://github.com/pingcap/tidb/pull/21207)                                               |   | 136d19h |
| tidb/21476   | [planner: check for decimal format in cast expr (#20836)](https://github.com/pingcap/tidb/pull/21476)                                                         |   | 126d15h |
| tidb/21487   | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                                                 |   | 126d4h  |
| tidb/21876   | [planner: bypass the DNF restriction if index merge hint is specified (#20799)](https://github.com/pingcap/tidb/pull/21876)                                   |   | 111d19h |
| tidb/21954   | [planner/cascades: add rule `PushSelDownApply`](https://github.com/pingcap/tidb/pull/21954)                                                                   |   | 106d23h |
| tidb/22181   | [planner, expression: fix error when using IN combined with subquery (#22080)](https://github.com/pingcap/tidb/pull/22181)                                    |   | 93d17h  |
| tidb/22416   | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                                                             | Y | 82d11h  |
| tidb/22504   | [*:Fix the fetchHotRegion bug that the count always zero](https://github.com/pingcap/tidb/pull/22504)                                                         |   | 74d19h  |
| tidb/22565   | [statistics: fix panic occurs when stats cache inconsistency (#22465)](https://github.com/pingcap/tidb/pull/22565)                                            | Y | 71d17h  |
| tidb/22624   | [ planner: not pruning column used by union scan condition (#21640)](https://github.com/pingcap/tidb/pull/22624)                                              |   | 69d17h  |
| tidb/22923   | [expression: correct constant propagation for collation (#22666)](https://github.com/pingcap/tidb/pull/22923)                                                 |   | 43d15h  |
| tidb/23208   | [statistics, util/ranger: improve selectivity calculation for DNF filters (#18741)](https://github.com/pingcap/tidb/pull/23208)                               |   | 30d16h  |
| tidb/23215   | [Privileges: fix delete privilege check wrongly (#22971)](https://github.com/pingcap/tidb/pull/23215)                                                         |   | 30d14h  |
| tidb/23233   | [planner: fix incorrect duration between compare (#22830)](https://github.com/pingcap/tidb/pull/23233)                                                        |   | 29d18h  |
| tidb/23347   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                                          |   | 23d18h  |
| tidb/23348   | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                                          |   | 23d18h  |
| tidb/23350   | [util/stringutil, util/ranger, planner: use hierarchical separators to simplify the parsing for info of EXPLAIN ](https://github.com/pingcap/tidb/pull/23350) |   | 23d17h  |
| tidb/23365   | [planner: fix a bug that point get plan returns wrong column name](https://github.com/pingcap/tidb/pull/23365)                                                |   | 22d22h  |
| tidb/23373   | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                             |   | 22d19h  |
| tidb/23474   | [planner: fix inappropriate null flag of null constants (#23457)](https://github.com/pingcap/tidb/pull/23474)                                                 |   | 16d18h  |
| tidb/23537   | [planner: remove some risky cache operations in the plan builder (#23354)](https://github.com/pingcap/tidb/pull/23537)                                        |   | 14d19h  |
| tidb/23543   | [statistics: fix auto analyze log information incomplete (#23522)](https://github.com/pingcap/tidb/pull/23543)                                                |   | 14d18h  |
| tidb/23598   | [types: fix collation for binary literal (#23591)](https://github.com/pingcap/tidb/pull/23598)                                                                |   | 13d13h  |
| tidb/23655   | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral (#23523)](https://github.com/pingcap/tidb/pull/23655)                      |   | 9d22h   |
| tidb/23656   | [planner, type: remove the prefix 0 in the bit array when we get the BinaryLiteral (#23523)](https://github.com/pingcap/tidb/pull/23656)                      |   | 9d22h   |
| tidb/23683   | [executor: fix a panic when batch point get is used for partition table (#23652)](https://github.com/pingcap/tidb/pull/23683)                                 |   | 9d16h   |
| tidb/23689   | [planner: fix the panic when we calculate the partition range (#23651)](https://github.com/pingcap/tidb/pull/23689)                                           |   | 9d16h   |
| tidb/23772   | [tablecodec: fix text type decode for old row format (#23751)](https://github.com/pingcap/tidb/pull/23772)                                                    |   | 8d11h   |
| tidb/23831   | [statistic: solve write write conflicts of mysql.stats_histograms](https://github.com/pingcap/tidb/pull/23831)                                                |   | 6d18h   |
| tidb/23849   | [ddl: tidb panic while query hash partition table with is null condition](https://github.com/pingcap/tidb/pull/23849)                                         |   | 3d13h   |
| tidb/23883   | [*: don't allocate SessionIndexUsageCollector when indexUsageLease equals 0 (#23861)](https://github.com/pingcap/tidb/pull/23883)                             |   | 1d20h   |
| tidb/23895   | [statistics: add more tests for extended stats](https://github.com/pingcap/tidb/pull/23895)                                                                   |   | 1d17h   |
| tidb/23917   | [planner: fix wrong TableDual plans caused by comparing Binary and Bytes incorrectly (#23860)](https://github.com/pingcap/tidb/pull/23917)                    |   | 23h     |
| tidb/23918   | [planner: fix wrong TableDual plans caused by comparing Binary and Bytes incorrectly (#23860)](https://github.com/pingcap/tidb/pull/23918)                    |   | 23h     |


## Reviewed in Last 7 Days

|    REPO    |                                                                PR                                                                 | C | D | R  |
|------------|-----------------------------------------------------------------------------------------------------------------------------------|---|---|----|
| tidb/23860 | [planner: fix wrong TableDual plans caused by comparing Binary and Bytes incorrectly](https://github.com/pingcap/tidb/pull/23860) |   | 3 | 3h |


</details> 


<details> 
  <summary>wshwsh12's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                     PR                                                                     | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19807 | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                                    |   | 216d10h |
| tidb/19957 | [executor: add builtin aggregate function `json_arrayagg`](https://github.com/pingcap/tidb/pull/19957)                                     | Y | 209d13h |
| tidb/21487 | [*: ensure TABLE statement works](https://github.com/pingcap/tidb/pull/21487)                                                              |   | 126d4h  |
| tidb/21887 | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                                            |   | 110d11h |
| tidb/22378 | [executor: vectorize hash aggregate](https://github.com/pingcap/tidb/pull/22378)                                                           |   | 85d19h  |
| tidb/22628 | [executor: Improve max/min window function with deque-based sliding window](https://github.com/pingcap/tidb/pull/22628)                    |   | 68d23h  |
| tidb/23336 | [expression: fix unexpected constant fold when year compare string (#23281)](https://github.com/pingcap/tidb/pull/23336)                   |   | 23d19h  |
| tidb/23347 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23347)                                       |   | 23d18h  |
| tidb/23348 | [planner: show cast type in EXPLAIN in coptask (#23123)](https://github.com/pingcap/tidb/pull/23348)                                       |   | 23d18h  |
| tidb/23368 | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23368)                      |   | 22d20h  |
| tidb/23369 | [executor, expression: fix the incorrect result of AVG function (#23285)](https://github.com/pingcap/tidb/pull/23369)                      |   | 22d20h  |
| tidb/23397 | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23397)                                             |   | 21d17h  |
| tidb/23398 | [expression: fix refine compare constant (#23339)](https://github.com/pingcap/tidb/pull/23398)                                             |   | 21d17h  |
| tidb/23519 | [executor: check privilege before adding](https://github.com/pingcap/tidb/pull/23519)                                                      |   | 15d0h   |
| tidb/23683 | [executor: fix a panic when batch point get is used for partition table (#23652)](https://github.com/pingcap/tidb/pull/23683)              |   | 9d16h   |
| tidb/23705 | [executor: refineArgs() bug fix when compare int with very small decimal (#23694)](https://github.com/pingcap/tidb/pull/23705)             |   | 9d13h   |
| tidb/23749 | [planner/core: fix a bug during add cast for decimal join key  (#23723)](https://github.com/pingcap/tidb/pull/23749)                       |   | 8d16h   |
| tidb/23760 | [collation: fix tidb panic when compare string with collation](https://github.com/pingcap/tidb/pull/23760)                                 |   | 8d13h   |
| tidb/23796 | [tests: make TestIndexLookupMergeJoinHang and TestIssue18068 stable (#23741)](https://github.com/pingcap/tidb/pull/23796)                  |   | 7d19h   |
| tidb/23853 | [ddl: grant nonexistent user/role incompatible error](https://github.com/pingcap/tidb/pull/23853)                                          |   | 2d21h   |
| tidb/23866 | [executor,kv: support timebounded staleness transaction](https://github.com/pingcap/tidb/pull/23866)                                       |   | 2d16h   |
| tidb/23868 | [store, kv: Add information about async commit/1pc to tidb_last_txn_info (#23833)](https://github.com/pingcap/tidb/pull/23868)             |   | 2d16h   |
| tidb/23876 | [executor: fix scope ambiguity of joinResult](https://github.com/pingcap/tidb/pull/23876)                                                  |   | 2d10h   |
| tidb/23917 | [planner: fix wrong TableDual plans caused by comparing Binary and Bytes incorrectly (#23860)](https://github.com/pingcap/tidb/pull/23917) |   | 23h     |
| tidb/23918 | [planner: fix wrong TableDual plans caused by comparing Binary and Bytes incorrectly (#23860)](https://github.com/pingcap/tidb/pull/23918) |   | 23h     |


## Reviewed in Last 7 Days

|    REPO    |                                                           PR                                                            | C | D |   R   |
|------------|-------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/23915 | [expression: Implementation of Vitess hashing algorithm. (#23493)](https://github.com/pingcap/tidb/pull/23915)          |   | 2 | 0h    |
| tidb/23878 | [functions: fix some string function has wrong collation and flag (#23835)](https://github.com/pingcap/tidb/pull/23878) |   | 2 | 3h    |
| tidb/23879 | [functions: fix some string function has wrong collation and flag (#23835)](https://github.com/pingcap/tidb/pull/23879) |   | 2 | 3h    |
| tidb/23835 | [functions: fix some string function has wrong collation and flag](https://github.com/pingcap/tidb/pull/23835)          |   | 2 | 4d18h |
| tidb/23823 | [plugin: fix audit plugin will cause tidb panic (#23803)](https://github.com/pingcap/tidb/pull/23823)                   |   | 3 | 3d21h |
| tidb/23819 | [plugin: fix audit plugin will cause tidb panic (#23803)](https://github.com/pingcap/tidb/pull/23819)                   |   | 3 | 4d2h  |


</details> 

