|   REVIEWER    |  O(C)   | 1 | 2 | 3 | 4 | 5 | 6 | 7 | T  |
|---------------|---------|---|---|---|---|---|---|---|----|
| Reminiscent   |  1 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| SunRunAway    |  7 ( 1) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| XuHuaiyu      | 10 ( 0) | 0 | 0 | 1 | 0 | 3 | 0 | 0 |  4 |
| dyzsr         |  0 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| eurekaka      | 12 ( 0) | 0 | 0 | 1 | 0 | 2 | 3 | 0 |  6 |
| fzhedu        |  0 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| ichn-hu       |  3 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 1 |  1 |
| lzmhhh123     |  4 ( 0) | 2 | 0 | 2 | 0 | 4 | 0 | 7 | 15 |
| nullnotnil    |  0 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| qw4990        | 16 ( 0) | 0 | 0 | 1 | 1 | 1 | 0 | 1 |  4 |
| rebelice      | 12 ( 0) | 0 | 0 | 1 | 2 | 0 | 1 | 0 |  4 |
| time-and-fate | 14 ( 1) | 0 | 0 | 1 | 0 | 0 | 0 | 1 |  2 |
| winoros       | 14 ( 1) | 0 | 0 | 0 | 1 | 1 | 8 | 1 | 11 |
| wshwsh12      |  4 ( 0) | 0 | 0 | 1 | 1 | 2 | 2 | 0 |  6 |


<details> 
  <summary>Reminiscent's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                          PR                                          | C | LASTED |
|------------|--------------------------------------------------------------------------------------|---|--------|
| tidb/25583 | [bindinfo: fix SPM doesn't work for CTE](https://github.com/pingcap/tidb/pull/25583) |   | 14d15h |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>SunRunAway's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                              PR                                                              | C | LASTED  |
|------------|------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/19178 | [executor: Refactor probe channel](https://github.com/pingcap/tidb/pull/19178)                                               |   | 325d16h |
| tidb/19807 | [executor: parallel evaluation for hash aggregate distinct](https://github.com/pingcap/tidb/pull/19807)                      |   | 303d10h |
| tidb/19900 | [executor: enable inline projection for sort&topN](https://github.com/pingcap/tidb/pull/19900)                               | Y | 298d18h |
| tidb/21834 | [planner: enhanced index range calculation plan](https://github.com/pingcap/tidb/pull/21834)                                 |   | 200d18h |
| tidb/21878 | [planner: do not push down lock to pointGet/bacthPointGet when selection exists](https://github.com/pingcap/tidb/pull/21878) |   | 198d18h |
| tidb/21956 | [planner/preprocessor: disallow into-outfile clause in some place](https://github.com/pingcap/tidb/pull/21956)               |   | 193d23h |
| tidb/25385 | [executor: global kill 32bits (local connID part)](https://github.com/pingcap/tidb/pull/25385)                               |   | 21d10h  |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>XuHuaiyu's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                         PR                                                          | C | LASTED  |
|--------------|---------------------------------------------------------------------------------------------------------------------|---|---------|
| docs-cn/5561 | [Add sql optimization-related docs to toc](https://github.com/pingcap/docs-cn/pull/5561)                            |   | 132d15h |
| tidb/21401   | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                  |   | 216d11h |
| docs-cn/6409 | [Change tidb_memory_usage_alarm_ratio scope to instance ](https://github.com/pingcap/docs-cn/pull/6409)             |   | 30d15h  |
| tidb/25340   | [privilege: add restricted read only](https://github.com/pingcap/tidb/pull/25340)                                   |   | 24d15h  |
| tidb/25714   | [executor: support spill intermediate data for unparalleled hash agg](https://github.com/pingcap/tidb/pull/25714)   |   | 11d16h  |
| tidb/25792   | [docs/design: Support Spilling Unparalleled HashAgg](https://github.com/pingcap/tidb/pull/25792)                    |   | 6d20h   |
| tidb/25878   | [types: Fix issue about ineffectual assignment](https://github.com/pingcap/tidb/pull/25878)                         |   | 3d5h    |
| tidb/25901   | [executor: fix query empty table with IN clause reports 'invalid year'](https://github.com/pingcap/tidb/pull/25901) |   | 2d19h   |
| tidb/25906   | [config, sessionctx: deprecate streaming config](https://github.com/pingcap/tidb/pull/25906)                        |   | 2d17h   |
| tidb/25915   | [executor: fix hash join between datetime and timestamp](https://github.com/pingcap/tidb/pull/25915)                |   | 16h     |


## Reviewed in Last 7 Days

|    REPO    |                                                                 PR                                                                  | C | D |    R    |
|------------|-------------------------------------------------------------------------------------------------------------------------------------|---|---|---------|
| tidb/25845 | [planner,executor: fix 'select ...(join on partition table) for update' panic (#21148)](https://github.com/pingcap/tidb/pull/25845) |   | 3 | 2d1h    |
| tidb/20140 | [expressions: Support `bin-to-uuid` and `uuid-to-bin`](https://github.com/pingcap/tidb/pull/20140)                                  |   | 5 | 280d23h |
| tidb/22541 | [expression: Support builtin function SOUNDEX](https://github.com/pingcap/tidb/pull/22541)                                          |   | 5 | 154d10h |
| tipb/207   | [Support builtin function SOUNDEX](https://github.com/pingcap/tipb/pull/207)                                                        |   | 5 | 154d10h |


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

|    REPO    |                                                                               PR                                                                               | C | LASTED  |
|------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/23316 | [planner: Fix rebuild range for prepared plan](https://github.com/pingcap/tidb/pull/23316)                                                                     |   | 111d17h |
| tidb/23373 | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                              |   | 109d19h |
| tidb/23760 | [collation: fix tidb panic when compare string with collation](https://github.com/pingcap/tidb/pull/23760)                                                     |   | 95d13h  |
| tidb/24061 | [statistics: fix some potential panic in statistics (#23988)](https://github.com/pingcap/tidb/pull/24061)                                                      |   | 80d13h  |
| tidb/24556 | [planner: add MergeAdjacentWindow rule for cascades](https://github.com/pingcap/tidb/pull/24556)                                                               |   | 54d10h  |
| tidb/24649 | [server: close the temporary session in HTTP API to avoid memory leak (#24339)](https://github.com/pingcap/tidb/pull/24649)                                    |   | 52d0h   |
| tidb/24650 | [server: close the temporary session in HTTP API to avoid memory leak (#24339)](https://github.com/pingcap/tidb/pull/24650)                                    |   | 52d0h   |
| tidb/24921 | [planner: update IsCompleteModeAgg and transform function of RuleInjectProjectionBelowAgg to fix distinct agg bug](https://github.com/pingcap/tidb/pull/24921) |   | 38d19h  |
| tidb/25501 | [planner,executor: fix 'select ...(join on partition table) for update' panic (#21148)](https://github.com/pingcap/tidb/pull/25501)                            |   | 18d11h  |
| tidb/25737 | [planner: Log warnings when agg function can not be pushdown in explain statement (#25553)](https://github.com/pingcap/tidb/pull/25737)                        |   | 10d18h  |
| tidb/25845 | [planner,executor: fix 'select ...(join on partition table) for update' panic (#21148)](https://github.com/pingcap/tidb/pull/25845)                            |   | 4d19h   |
| tidb/25861 | [planner/core: thoroughly push down count-distinct agg in the MPP mode. (#25662)](https://github.com/pingcap/tidb/pull/25861)                                  |   | 3d19h   |


## Reviewed in Last 7 Days

|    REPO    |                                                                      PR                                                                       | C | D |   R    |
|------------|-----------------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/25583 | [bindinfo: fix SPM doesn't work for CTE](https://github.com/pingcap/tidb/pull/25583)                                                          |   | 3 | 11d20h |
| tidb/25471 | [planner: fix wrong aggregate pruning for some cases (#25289)](https://github.com/pingcap/tidb/pull/25471)                                    |   | 5 | 13d21h |
| tidb/24079 | [planner: change descScanFactor to scanFactor when ExpectedCount is small. (#23972)](https://github.com/pingcap/tidb/pull/24079)              |   | 5 | 74d20h |
| tidb/25818 | [planner: handle other-conditions from subqueries correctly when constructing IndexJoin (#25817)](https://github.com/pingcap/tidb/pull/25818) |   | 6 | 0h     |
| tidb/25819 | [planner: handle other-conditions from subqueries correctly when constructing IndexJoin (#25817)](https://github.com/pingcap/tidb/pull/25819) |   | 6 | 0h     |
| tidb/25817 | [planner: handle other-conditions from subqueries correctly when constructing IndexJoin](https://github.com/pingcap/tidb/pull/25817)          |   | 6 | 0h     |


</details> 


<details> 
  <summary>fzhedu's detailed review report</summary> 

## To Be Reviewed

| REPO | PR | C | LASTED |
|------|----|---|--------|


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>ichn-hu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                  PR                                                                  | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20903 | [planner: fix confused and unnecessary double-projection in plans.](https://github.com/pingcap/tidb/pull/20903)                      |   | 240d17h |
| tidb/22631 | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                                                      |   | 154d23h |
| tidb/25766 | [expression: Fix greatest and least function lost decimal precision compared with MySQL](https://github.com/pingcap/tidb/pull/25766) |   | 9d15h   |


## Reviewed in Last 7 Days

|    REPO    |                                                PR                                                 | C | D |   R   |
|------------|---------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/25523 | [expression: support datetime type for user variable](https://github.com/pingcap/tidb/pull/25523) |   | 7 | 11d1h |


</details> 


<details> 
  <summary>lzmhhh123's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                    PR                                                    | C | LASTED  |
|------------|----------------------------------------------------------------------------------------------------------|---|---------|
| tidb/22631 | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                          |   | 154d23h |
| tidb/25730 | [txn: change lock into put record for unique index key lock](https://github.com/pingcap/tidb/pull/25730) |   | 10d20h  |
| tidb/25822 | [executor: avoid unnecessary string EqualFold() comparisons](https://github.com/pingcap/tidb/pull/25822) |   | 5d16h   |
| tidb/25906 | [config, sessionctx: deprecate streaming config](https://github.com/pingcap/tidb/pull/25906)             |   | 2d17h   |


## Reviewed in Last 7 Days

|    REPO    |                                                                  PR                                                                  | C | D |    R    |
|------------|--------------------------------------------------------------------------------------------------------------------------------------|---|---|---------|
| tidb/24806 | [config: ignore tiflash when show config (#24770)](https://github.com/pingcap/tidb/pull/24806)                                       |   | 1 | 45d11h  |
| tidb/25562 | [expression: push down abs() to TiFlash (#24841)](https://github.com/pingcap/tidb/pull/25562)                                        |   | 1 | 16d18h  |
| tidb/25611 | [expression:  error information is inconsistent with MySQL about date or time literal](https://github.com/pingcap/tidb/pull/25611)   |   | 3 | 10d20h  |
| tidb/25854 | [executor: query memory table TIKV_REGION_PEERS and TIKV_REGION_STATUS get error](https://github.com/pingcap/tidb/pull/25854)        |   | 3 | 1d14h   |
| tidb/25766 | [expression: Fix greatest and least function lost decimal precision compared with MySQL](https://github.com/pingcap/tidb/pull/25766) |   | 5 | 4d20h   |
| tidb/24913 | [planner: fix incorrect usage of UNION and INTO](https://github.com/pingcap/tidb/pull/24913)                                         |   | 5 | 34d2h   |
| tidb/21018 | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                       |   | 5 | 229d19h |
| tidb/20444 | [expression: add json_merge_patch](https://github.com/pingcap/tidb/pull/20444)                                                       |   | 5 | 258d23h |
| tikv/10435 | [copr: make CM Sketch built with the same encoding as what TiDB assumes (#10418)](https://github.com/tikv/tikv/pull/10435)           | Y | 7 | 4d22h   |
| tikv/10433 | [copr: make CM Sketch built with the same encoding as what TiDB assumes (#10418)](https://github.com/tikv/tikv/pull/10433)           | Y | 7 | 4d22h   |
| tidb/25596 | [expression: Support mathematical functions pushdown to tiflash](https://github.com/pingcap/tidb/pull/25596)                         |   | 7 | 7d2h    |
| tidb/25665 | [expression: fix IN expr critical bug (#25653)](https://github.com/pingcap/tidb/pull/25665)                                          |   | 7 | 5d18h   |
| tidb/25740 | [planner: enforce projection push down (#25450)](https://github.com/pingcap/tidb/pull/25740)                                         |   | 7 | 3d19h   |
| tidb/25741 | [planner: enforce projection push down (#25450)](https://github.com/pingcap/tidb/pull/25741)                                         |   | 7 | 3d19h   |
| tidb/25767 | [expression: Improve the compatibility of `str_to_date` (#25386)](https://github.com/pingcap/tidb/pull/25767)                        |   | 7 | 2d15h   |


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

|     REPO     |                                                                                       PR                                                                                        | C | LASTED  |
|--------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                                                  |   | 234d17h |
| docs-cn/5561 | [Add sql optimization-related docs to toc](https://github.com/pingcap/docs-cn/pull/5561)                                                                                        |   | 132d15h |
| docs/5498    | [partitioning: Corrected partition management](https://github.com/pingcap/docs/pull/5498)                                                                                       |   | 69d19h  |
| tidb/21318   | [planner, expression: use the range of column types to simplify expressions](https://github.com/pingcap/tidb/pull/21318)                                                        |   | 220d18h |
| tidb/23295   | [util, types: don't let SPM be affected by charset (#23161)](https://github.com/pingcap/tidb/pull/23295)                                                                        |   | 114d11h |
| tidb/23590   | [planner, table: optimize the list partition pruner for range query](https://github.com/pingcap/tidb/pull/23590)                                                                |   | 100d16h |
| tidb/24663   | [planner: include schema name when checking duplicate table aliases](https://github.com/pingcap/tidb/pull/24663)                                                                |   | 51d16h  |
| tidb/24994   | [planner: don't extract hash keys from index join's OtherConds if inl_merge_join hint exists](https://github.com/pingcap/tidb/pull/24994)                                       |   | 34d17h  |
| tidb/25693   | [planner: fix index-out-of-range error when checking only_full_group_by and make sure limit outputs no more columns than its child](https://github.com/pingcap/tidb/pull/25693) |   | 11d22h  |
| tidb/25715   | [planner: fix row count estimation for partially pushed down selections](https://github.com/pingcap/tidb/pull/25715)                                                            |   | 11d16h  |
| tidb/25763   | [executor: reject setting read ts to a future time (#25732)](https://github.com/pingcap/tidb/pull/25763)                                                                        |   | 9d16h   |
| tidb/25769   | [planner: add some comment for checkOnlyFullGroupBy](https://github.com/pingcap/tidb/pull/25769)                                                                                |   | 9d12h   |
| tidb/25806   | [planner: check filter condition in func convertToPartialTableScan (#25294)](https://github.com/pingcap/tidb/pull/25806)                                                        |   | 6d15h   |
| tidb/25845   | [planner,executor: fix 'select ...(join on partition table) for update' panic (#21148)](https://github.com/pingcap/tidb/pull/25845)                                             |   | 4d19h   |
| tidb/25861   | [planner/core: thoroughly push down count-distinct agg in the MPP mode. (#25662)](https://github.com/pingcap/tidb/pull/25861)                                                   |   | 3d19h   |
| tidb/25870   | [executor: skip all test cases related to TiFlash+Partition since they are too slow (#25866)](https://github.com/pingcap/tidb/pull/25870)                                       |   | 3d15h   |


## Reviewed in Last 7 Days

|    REPO    |                                                          PR                                                          | C | D |   R    |
|------------|----------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/25226 | [planner: fix bug when unfolding wildcard in view definiton](https://github.com/pingcap/tidb/pull/25226)             |   | 3 | 24d15h |
| dm/1782    | [syncer: support filter row by SQL expression](https://github.com/pingcap/dm/pull/1782)                              |   | 4 | 14d10h |
| tidb/25743 | [sessionctx: add tidb_enable_list_partition global system variable](https://github.com/pingcap/tidb/pull/25743)      |   | 5 | 5d21h  |
| tidb/25662 | [planner/core: thoroughly push down count-distinct agg in the MPP mode.](https://github.com/pingcap/tidb/pull/25662) |   | 7 | 5d21h  |


</details> 


<details> 
  <summary>rebelice's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                    PR                                                                     | C | LASTED |
|--------------|-------------------------------------------------------------------------------------------------------------------------------------------|---|--------|
| docs/5185    | [sql-statements, information-schema: add `END_TIME` field for table `ANALYZE_STATUS`](https://github.com/pingcap/docs/pull/5185)          |   | 94d17h |
| docs-cn/5916 | [sql-statements, information-schema: add `END_TIME` field for table `ANALYZE_STATUS`](https://github.com/pingcap/docs-cn/pull/5916)       |   | 94d17h |
| tidb/24033   | [statistics: fix some unstable tests in global stats (#23502)](https://github.com/pingcap/tidb/pull/24033)                                |   | 81d9h  |
| docs-cn/6542 | [update docs related to partition table dynamic mode](https://github.com/pingcap/docs-cn/pull/6542)                                       |   | 6d19h  |
| tidb/24306   | [util/ranger: fix func name typo](https://github.com/pingcap/tidb/pull/24306)                                                             |   | 68d22h |
| tidb/24374   | [planner: filter conflict read_from_storage hints (#24313)](https://github.com/pingcap/tidb/pull/24374)                                   |   | 66d19h |
| tidb/24649   | [server: close the temporary session in HTTP API to avoid memory leak (#24339)](https://github.com/pingcap/tidb/pull/24649)               |   | 52d0h  |
| tidb/24650   | [server: close the temporary session in HTTP API to avoid memory leak (#24339)](https://github.com/pingcap/tidb/pull/24650)               |   | 52d0h  |
| tidb/24669   | [planner: fix "order by + num " can use a column not in select fields](https://github.com/pingcap/tidb/pull/24669)                        |   | 51d16h |
| tidb/25214   | [planner: don't push down topn to nil table plan side](https://github.com/pingcap/tidb/pull/25214)                                        |   | 27d16h |
| tidb/25861   | [planner/core: thoroughly push down count-distinct agg in the MPP mode. (#25662)](https://github.com/pingcap/tidb/pull/25861)             |   | 3d19h  |
| tidb/25870   | [executor: skip all test cases related to TiFlash+Partition since they are too slow (#25866)](https://github.com/pingcap/tidb/pull/25870) |   | 3d15h  |


## Reviewed in Last 7 Days

|    REPO    |                                                                PR                                                                | C | D |   R   |
|------------|----------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/25806 | [planner: check filter condition in func convertToPartialTableScan (#25294)](https://github.com/pingcap/tidb/pull/25806)         |   | 3 | 3d21h |
| tidb/25866 | [executor: skip all test cases related to TiFlash+Partition since they are too slow](https://github.com/pingcap/tidb/pull/25866) |   | 4 | 0h    |
| tidb/23836 | [parser, core: Implement force_index hint in parser and TiDB](https://github.com/pingcap/tidb/pull/23836)                        |   | 4 | 90d0h |
| tidb/25662 | [planner/core: thoroughly push down count-distinct agg in the MPP mode.](https://github.com/pingcap/tidb/pull/25662)             |   | 6 | 7d0h  |


</details> 


<details> 
  <summary>time-and-fate's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                      PR                                                                       | C | LASTED  |
|------------|-----------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/22416 | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                                             | Y | 169d11h |
| tidb/24374 | [planner: filter conflict read_from_storage hints (#24313)](https://github.com/pingcap/tidb/pull/24374)                                       |   | 66d19h  |
| tidb/24382 | [statistics: trigger auto-analyze based on histogram row count](https://github.com/pingcap/tidb/pull/24382)                                   |   | 66d15h  |
| tidb/24539 | [statistics: dump FMSketch to KV only for partition table with dynamic prune mode (#24453)](https://github.com/pingcap/tidb/pull/24539)       |   | 54d21h  |
| tidb/24556 | [planner: add MergeAdjacentWindow rule for cascades](https://github.com/pingcap/tidb/pull/24556)                                              |   | 54d10h  |
| tidb/24994 | [planner: don't extract hash keys from index join's OtherConds if inl_merge_join hint exists](https://github.com/pingcap/tidb/pull/24994)     |   | 34d17h  |
| tidb/25094 | [*: resolve select fields properly for coalesced columns of natural join](https://github.com/pingcap/tidb/pull/25094)                         |   | 31d18h  |
| tidb/25390 | [planner/core: fix `isTableAliasDuplicate`, use `schema.name` as key when table has a alias name](https://github.com/pingcap/tidb/pull/25390) |   | 20d19h  |
| tidb/25696 | [planner: generate wrong plan when update has subquery (#25660)](https://github.com/pingcap/tidb/pull/25696)                                  |   | 11d22h  |
| tidb/25698 | [planner: generate wrong plan when update has subquery (#25660)](https://github.com/pingcap/tidb/pull/25698)                                  |   | 11d22h  |
| tidb/25715 | [planner: fix row count estimation for partially pushed down selections](https://github.com/pingcap/tidb/pull/25715)                          |   | 11d16h  |
| tidb/25736 | [planner: Log warnings when agg function can not be pushdown in explain statement (#25553)](https://github.com/pingcap/tidb/pull/25736)       |   | 10d18h  |
| tidb/25737 | [planner: Log warnings when agg function can not be pushdown in explain statement (#25553)](https://github.com/pingcap/tidb/pull/25737)       |   | 10d18h  |
| tidb/25819 | [planner: handle other-conditions from subqueries correctly when constructing IndexJoin (#25817)](https://github.com/pingcap/tidb/pull/25819) |   | 5d18h   |


## Reviewed in Last 7 Days

|    REPO    |                                                          PR                                                          | C | D |   R   |
|------------|----------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/24720 | [*: update count / modify_count of mysql.stats_meta incrementally](https://github.com/pingcap/tidb/pull/24720)       |   | 3 | 45d1h |
| tidb/25789 | [statistics: skip dumping nil histograms for virtual columns in analyze](https://github.com/pingcap/tidb/pull/25789) |   | 7 | 1h    |


</details> 


<details> 
  <summary>winoros's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                                     PR                                                                                     | C | LASTED  |
|--------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20903   | [planner: fix confused and unnecessary double-projection in plans.](https://github.com/pingcap/tidb/pull/20903)                                                            |   | 240d17h |
| docs-cn/5916 | [sql-statements, information-schema: add `END_TIME` field for table `ANALYZE_STATUS`](https://github.com/pingcap/docs-cn/pull/5916)                                        |   | 94d17h  |
| docs/5783    | [migration: Add information about Vitess to TiDB migration](https://github.com/pingcap/docs/pull/5783)                                                                     |   | 20d5h   |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                                             |   | 234d17h |
| tidb/22416   | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                                                                          | Y | 169d11h |
| tidb/22504   | [*:Fix the fetchHotRegion bug that the count always zero](https://github.com/pingcap/tidb/pull/22504)                                                                      |   | 161d19h |
| tidb/23373   | [executor: fix get var expr when session var is hex literal (#23241)](https://github.com/pingcap/tidb/pull/23373)                                                          |   | 109d19h |
| tidb/24138   | [planner: Add Equivalence Rules to Transform BinaryOptSubquery to ExistsSubquery](https://github.com/pingcap/tidb/pull/24138)                                              |   | 76d12h  |
| tidb/24663   | [planner: include schema name when checking duplicate table aliases](https://github.com/pingcap/tidb/pull/24663)                                                           |   | 51d16h  |
| tidb/24921   | [planner: update IsCompleteModeAgg and transform function of RuleInjectProjectionBelowAgg to fix distinct agg bug](https://github.com/pingcap/tidb/pull/24921)             |   | 38d19h  |
| tidb/25686   | [*: always convert sysvar values when out of range](https://github.com/pingcap/tidb/pull/25686)                                                                            |   | 12d0h   |
| tidb/25828   | [executor: support forbid tiflash for stale read](https://github.com/pingcap/tidb/pull/25828)                                                                              |   | 5d14h   |
| tidb/25870   | [executor: skip all test cases related to TiFlash+Partition since they are too slow (#25866)](https://github.com/pingcap/tidb/pull/25870)                                  |   | 3d15h   |
| tidb/25874   | [util/stmtsummary: discard the plan if it is too long and enlarge the tidb_stmt_summary_max_stmt_count value to 3000 (#25843)](https://github.com/pingcap/tidb/pull/25874) |   | 3d15h   |


## Reviewed in Last 7 Days

|    REPO    |                                                                      PR                                                                       | C | D |   R    |
|------------|-----------------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/25866 | [executor: skip all test cases related to TiFlash+Partition since they are too slow](https://github.com/pingcap/tidb/pull/25866)              |   | 4 | 0h     |
| tidb/25583 | [bindinfo: fix SPM doesn't work for CTE](https://github.com/pingcap/tidb/pull/25583)                                                          |   | 5 | 10d5h  |
| tidb/25806 | [planner: check filter condition in func convertToPartialTableScan (#25294)](https://github.com/pingcap/tidb/pull/25806)                      |   | 6 | 22h    |
| tidb/25094 | [*: resolve select fields properly for coalesced columns of natural join](https://github.com/pingcap/tidb/pull/25094)                         |   | 6 | 26d0h  |
| tidb/24720 | [*: update count / modify_count of mysql.stats_meta incrementally](https://github.com/pingcap/tidb/pull/24720)                                |   | 6 | 42d0h  |
| tidb/25818 | [planner: handle other-conditions from subqueries correctly when constructing IndexJoin (#25817)](https://github.com/pingcap/tidb/pull/25818) |   | 6 | 0h     |
| tidb/25819 | [planner: handle other-conditions from subqueries correctly when constructing IndexJoin (#25817)](https://github.com/pingcap/tidb/pull/25819) |   | 6 | 0h     |
| tidb/25715 | [planner: fix row count estimation for partially pushed down selections](https://github.com/pingcap/tidb/pull/25715)                          |   | 6 | 5d21h  |
| tidb/25817 | [planner: handle other-conditions from subqueries correctly when constructing IndexJoin](https://github.com/pingcap/tidb/pull/25817)          |   | 6 | 0h     |
| tidb/24994 | [planner: don't extract hash keys from index join's OtherConds if inl_merge_join hint exists](https://github.com/pingcap/tidb/pull/24994)     |   | 6 | 28d22h |
| tidb/25789 | [statistics: skip dumping nil histograms for virtual columns in analyze](https://github.com/pingcap/tidb/pull/25789)                          |   | 7 | 0h     |


</details> 


<details> 
  <summary>wshwsh12's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                 PR                                                  | C | LASTED  |
|------------|-----------------------------------------------------------------------------------------------------|---|---------|
| tidb/21401 | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)  |   | 216d11h |
| tidb/21887 | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)     |   | 197d11h |
| tidb/22541 | [expression: Support builtin function SOUNDEX](https://github.com/pingcap/tidb/pull/22541)          |   | 159d9h  |
| tidb/24711 | [expression: add builtin function ``json_merge_patch``](https://github.com/pingcap/tidb/pull/24711) |   | 47d19h  |


## Reviewed in Last 7 Days

|    REPO    |                                                                 PR                                                                 | C | D |   R   |
|------------|------------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/25854 | [executor: query memory table TIKV_REGION_PEERS and TIKV_REGION_STATUS get error](https://github.com/pingcap/tidb/pull/25854)      |   | 3 | 1d14h |
| tidb/25389 | [expression: Improve the performance of `str_to_date`](https://github.com/pingcap/tidb/pull/25389)                                 |   | 4 | 17d9h |
| tidb/25611 | [expression:  error information is inconsistent with MySQL about date or time literal](https://github.com/pingcap/tidb/pull/25611) |   | 5 | 9d3h  |
| tidb/25810 | [executor: temporarily skip 2 unstable tests](https://github.com/pingcap/tidb/pull/25810)                                          |   | 5 | 1d5h  |
| tidb/25797 | [*: update tikv/client-go to improve failpoint performance issue](https://github.com/pingcap/tidb/pull/25797)                      |   | 6 | 22h   |
| tidb/25768 | [expression: Improve the compatibility of `str_to_date` (#25386)](https://github.com/pingcap/tidb/pull/25768)                      |   | 6 | 3d15h |


</details> 

