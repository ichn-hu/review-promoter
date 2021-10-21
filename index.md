|   REVIEWER    |  O(C)   | 1 | 2 | 3 | 4 | 5 | 6 | 7 | T |
|---------------|---------|---|---|---|---|---|---|---|---|
| Reminiscent   | 18 ( 0) | 1 | 1 | 2 | 0 | 0 | 0 | 0 | 4 |
| SunRunAway    |  4 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| XuHuaiyu      | 15 ( 0) | 0 | 0 | 0 | 1 | 0 | 1 | 2 | 4 |
| dyzsr         |  0 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| eurekaka      |  3 ( 1) | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| fzhedu        |  7 ( 0) | 2 | 2 | 0 | 0 | 0 | 0 | 0 | 4 |
| ichn-hu       |  9 ( 0) | 0 | 0 | 0 | 0 | 0 | 2 | 0 | 2 |
| lzmhhh123     | 16 ( 0) | 0 | 0 | 0 | 0 | 0 | 2 | 0 | 2 |
| nullnotnil    |  0 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 |
| qw4990        | 25 ( 0) | 1 | 0 | 2 | 0 | 0 | 2 | 2 | 7 |
| rebelice      | 21 ( 0) | 1 | 1 | 0 | 0 | 0 | 1 | 4 | 7 |
| time-and-fate | 18 ( 1) | 0 | 0 | 2 | 0 | 0 | 1 | 1 | 4 |
| winoros       | 29 ( 1) | 2 | 1 | 2 | 0 | 0 | 0 | 0 | 5 |
| wshwsh12      |  6 ( 0) | 0 | 1 | 0 | 0 | 0 | 0 | 0 | 1 |


<details> 
  <summary>Reminiscent's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                                               PR                                                                                               | C | LASTED  |
|------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/23590 | [planner, table: optimize the list partition pruner for range query](https://github.com/pingcap/tidb/pull/23590)                                                                               |   | 208d16h |
| tidb/26474 | [planner: fix the unstable unit test TestTableFromMeta (#26463)](https://github.com/pingcap/tidb/pull/26474)                                                                                   |   | 90d16h  |
| tidb/26475 | [planner: fix the unstable unit test TestTableFromMeta (#26463)](https://github.com/pingcap/tidb/pull/26475)                                                                                   |   | 90d16h  |
| tidb/26491 | [planner: fix the unstable test TestOrderedResultModeOnOtherOperators (#26481)](https://github.com/pingcap/tidb/pull/26491)                                                                    |   | 89d23h  |
| tidb/26492 | [planner: fix the unstable test TestOrderedResultModeOnOtherOperators (#26481)](https://github.com/pingcap/tidb/pull/26492)                                                                    |   | 89d23h  |
| tidb/26498 | [planner: fix the unstable unit test `TestAnalyzeIncremental` (#26460)](https://github.com/pingcap/tidb/pull/26498)                                                                            |   | 89d20h  |
| tidb/26499 | [planner: fix the unstable unit test `TestAnalyzeIncremental` (#26460)](https://github.com/pingcap/tidb/pull/26499)                                                                            |   | 89d20h  |
| tidb/26503 | [planner: fix goroutine leak problem in some unit tests (#26500)](https://github.com/pingcap/tidb/pull/26503)                                                                                  |   | 89d19h  |
| tidb/27636 | [planner, expression: avoid exprs with side effects in column pruning and agg pushdown (#27370)](https://github.com/pingcap/tidb/pull/27636)                                                   |   | 54d17h  |
| tidb/27773 | [statistics: remove redundant assignment for statistics.Column.Count](https://github.com/pingcap/tidb/pull/27773)                                                                              |   | 48d16h  |
| tidb/27837 | [planner: fix wrong plan caused by shallow copy schema columns (#27798)](https://github.com/pingcap/tidb/pull/27837)                                                                           |   | 44d16h  |
| tidb/28836 | [planner: fix the issue that plan-cache cannot be aware of changes of unsigned flags (#28827)](https://github.com/pingcap/tidb/pull/28836)                                                     |   | 5d23h   |
| tidb/28878 | [docs: add design doc for analyze predicate columns](https://github.com/pingcap/tidb/pull/28878)                                                                                               |   | 5d16h   |
| tidb/28927 | [planner: fix the issue that cached IndexJoin plans may return wrong results after changing parameters (#28915)](https://github.com/pingcap/tidb/pull/28927)                                   |   | 2d9h    |
| tidb/28928 | [planner: fix the issue that cached IndexJoin plans may return wrong results after changing parameters (#28915)](https://github.com/pingcap/tidb/pull/28928)                                   |   | 2d9h    |
| tidb/28992 | [planner: add an extra safe-guard selection upon DataSource to prevent wrong results caused by wrong rebuilt range when using plan-cache (#28976)](https://github.com/pingcap/tidb/pull/28992) |   | 6h      |
| tidb/28993 | [planner: add an extra safe-guard selection upon DataSource to prevent wrong results caused by wrong rebuilt range when using plan-cache (#28976)](https://github.com/pingcap/tidb/pull/28993) |   | 6h      |
| tidb/28994 | [planner: add an extra safe-guard selection upon DataSource to prevent wrong results caused by wrong rebuilt range when using plan-cache (#28976)](https://github.com/pingcap/tidb/pull/28994) |   | 6h      |


## Reviewed in Last 7 Days

|    REPO    |                                                                                          PR                                                                                           | C | D |  R   |
|------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---|------|
| tidb/28976 | [planner: add an extra safe-guard selection upon DataSource to prevent wrong results caused by wrong rebuilt range when using plan-cache](https://github.com/pingcap/tidb/pull/28976) |   | 1 | 2h   |
| tidb/28938 | [planner: fix the issue that plan-cache may return wrong results when the IndexLookup depends on a generated column](https://github.com/pingcap/tidb/pull/28938)                      |   | 2 | 2h   |
| tidb/28915 | [planner: fix the issue that cached IndexJoin plans may return wrong results after changing parameters](https://github.com/pingcap/tidb/pull/28915)                                   |   | 3 | 0h   |
| tidb/28837 | [planner: fix the issue that plan-cache cannot be aware of changes of unsigned flags (#28827)](https://github.com/pingcap/tidb/pull/28837)                                            |   | 3 | 3d0h |


</details> 


<details> 
  <summary>SunRunAway's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                        PR                                                                        | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/21834 | [planner: enhanced index range calculation plan](https://github.com/pingcap/tidb/pull/21834)                                                     |   | 308d18h |
| tidb/21956 | [planner/preprocessor: disallow into-outfile clause in some place](https://github.com/pingcap/tidb/pull/21956)                                   |   | 301d23h |
| tidb/25385 | [executor: global kill 32bits (local connID part)](https://github.com/pingcap/tidb/pull/25385)                                                   |   | 129d10h |
| tidb/27832 | [executor: fix a bug that can not insert null into a not null column in the empty SQL mode (#21237)](https://github.com/pingcap/tidb/pull/27832) |   | 44d16h  |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>XuHuaiyu's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                               PR                                                                | C | LASTED  |
|--------------|---------------------------------------------------------------------------------------------------------------------------------|---|---------|
| docs-cn/5561 | [Add sql optimization-related docs to toc](https://github.com/pingcap/docs-cn/pull/5561)                                        |   | 240d15h |
| tidb/21401   | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                              |   | 324d11h |
| docs-cn/6716 | [sysvar: add doc for tidb-restricted-read-only](https://github.com/pingcap/docs-cn/pull/6716)                                   |   | 90d18h  |
| tidb/26098   | [executor, planner: add support for SQL_CALC_FOUND_ROWS](https://github.com/pingcap/tidb/pull/26098)                            |   | 102d23h |
| tidb/26440   | [executor: a HashJoin demo in exchange parallel framework](https://github.com/pingcap/tidb/pull/26440)                          |   | 91d17h  |
| tidb/27315   | [go.mod: update parser to fix the parse error for subquery (#25647)](https://github.com/pingcap/tidb/pull/27315)                |   | 64d13h  |
| tidb/27378   | [distsql: fix goroutine/memory leak for streaming when query is cancelled (#27354)](https://github.com/pingcap/tidb/pull/27378) |   | 62d18h  |
| tidb/27396   | [*: set consistent assertion for DML](https://github.com/pingcap/tidb/pull/27396)                                               |   | 62d13h  |
| tidb/27403   | [expression: round function for int should use round half up rule](https://github.com/pingcap/tidb/pull/27403)                  |   | 62d11h  |
| tidb/27992   | [planner: add sub plan info of shuffleReceiver when query explain analyze](https://github.com/pingcap/tidb/pull/27992)          |   | 37d16h  |
| tidb/28870   | [expression: Fix wrong result of hour function in vectorized expression (#28857)](https://github.com/pingcap/tidb/pull/28870)   |   | 5d17h   |
| tidb/28871   | [expression: Fix wrong result of hour function in vectorized expression (#28857)](https://github.com/pingcap/tidb/pull/28871)   |   | 5d17h   |
| tidb/28872   | [expression: Fix wrong result of hour function in vectorized expression (#28857)](https://github.com/pingcap/tidb/pull/28872)   |   | 5d17h   |
| tidb/28953   | [sessionctx: Fix SET GLOBAL tidb_skip_isolation_level_check=1 (#27898)](https://github.com/pingcap/tidb/pull/28953)             |   | 1d17h   |
| tidb/28957   | [sessionctx: Fix SET GLOBAL tidb_skip_isolation_level_check=1 (#27898)](https://github.com/pingcap/tidb/pull/28957)             |   | 1d16h   |


## Reviewed in Last 7 Days

|     REPO      |                                                              PR                                                               | C | D |   R   |
|---------------|-------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/28874    | [expression: Fix wrong result of hour function in vectorized expression (#28857)](https://github.com/pingcap/tidb/pull/28874) |   | 4 | 2d17h |
| tidb/28857    | [expression: Fix wrong result of hour function in vectorized expression](https://github.com/pingcap/tidb/pull/28857)          |   | 6 | 0h    |
| tidb/28654    | [expression: not push invalid cast to tiflash (#28458)](https://github.com/pingcap/tidb/pull/28654)                           |   | 7 | 5d19h |
| community/583 | [fix: fix typo in tidb membership.json](https://github.com/pingcap/community/pull/583)                                        |   | 7 | 18h   |


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

|    REPO    |                                                PR                                                 | C | LASTED  |
|------------|---------------------------------------------------------------------------------------------------|---|---------|
| tidb/22416 | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416) | Y | 277d12h |
| tidb/23316 | [planner: Fix rebuild range for prepared plan](https://github.com/pingcap/tidb/pull/23316)        |   | 219d17h |
| tidb/27099 | [planner: support expression index for view](https://github.com/pingcap/tidb/pull/27099)          |   | 70d19h  |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>fzhedu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                         PR                                                                         | C | LASTED |
|------------|----------------------------------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/28147 | [planner: fix can not found column bug (#28067)](https://github.com/pingcap/tidb/pull/28147)                                                       |   | 33d19h |
| tidb/28262 | [distsql: avoid false positive error log about `invalid cop task execution summaries length` (#28188)](https://github.com/pingcap/tidb/pull/28262) |   | 28d16h |
| tidb/28263 | [distsql: avoid false positive error log about `invalid cop task execution summaries length` (#28188)](https://github.com/pingcap/tidb/pull/28263) |   | 28d16h |
| tidb/28287 | [copr: Fix bug that mpp node availability detect does not work in some corner cases (#28201)](https://github.com/pingcap/tidb/pull/28287)          |   | 27d21h |
| tidb/28288 | [copr: Fix bug that mpp node availability detect does not work in some corner cases (#28201)](https://github.com/pingcap/tidb/pull/28288)          |   | 27d21h |
| tidb/28651 | [expression: not push invalid cast to tiflash (#28458)](https://github.com/pingcap/tidb/pull/28651)                                                |   | 12d18h |
| tidb/28652 | [expression: not push invalid cast to tiflash (#28458)](https://github.com/pingcap/tidb/pull/28652)                                                |   | 12d18h |


## Reviewed in Last 7 Days

|    REPO    |                                                           PR                                                           | C | D |  R  |
|------------|------------------------------------------------------------------------------------------------------------------------|---|---|-----|
| tidb/28981 | [copr: (hotfix)MPP balance regions between TiFlash nodes with continuity.](https://github.com/pingcap/tidb/pull/28981) |   | 1 | 2h  |
| tidb/28972 | [copr: MPP balance regions between TiFlash nodes with continuity.](https://github.com/pingcap/tidb/pull/28972)         |   | 1 | 1h  |
| tics/3260  | [MPP: Restore concurrency before add ExchangeSenders](https://github.com/pingcap/tics/pull/3260)                       |   | 2 | 21h |
| tidb/28906 | [copr: MPP balance regions between TiFlash nodes with continuity. ](https://github.com/pingcap/tidb/pull/28906)        |   | 2 | 23h |


</details> 


<details> 
  <summary>ichn-hu's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                         PR                                                                         | C | LASTED  |
|--------------|----------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20903   | [planner: fix confused and unnecessary double-projection in plans.](https://github.com/pingcap/tidb/pull/20903)                                    |   | 348d17h |
| docs-cn/7238 | [system-variables: correct the description of tidb_allow_fallback_to_tikv](https://github.com/pingcap/docs-cn/pull/7238)                           |   | 20d19h  |
| tidb/22631   | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                                                                    |   | 262d23h |
| tidb/27119   | [executor: fix json_objectagg() on varbinary type](https://github.com/pingcap/tidb/pull/27119)                                                     |   | 70d16h  |
| tidb/27403   | [expression: round function for int should use round half up rule](https://github.com/pingcap/tidb/pull/27403)                                     |   | 62d11h  |
| tidb/27451   | [expression: fix wrong result for date add sub (#27244)](https://github.com/pingcap/tidb/pull/27451)                                               |   | 61d16h  |
| tidb/28262   | [distsql: avoid false positive error log about `invalid cop task execution summaries length` (#28188)](https://github.com/pingcap/tidb/pull/28262) |   | 28d16h  |
| tidb/28499   | [expression: align null flag of union columns and constants](https://github.com/pingcap/tidb/pull/28499)                                           |   | 20d15h  |
| tidb/28666   | [executor: fill extra partition ID column in UnionScan executor](https://github.com/pingcap/tidb/pull/28666)                                       |   | 12d11h  |


## Reviewed in Last 7 Days

|    REPO    |                                            PR                                            | C | D |  R  |
|------------|------------------------------------------------------------------------------------------|---|---|-----|
| tidb/28838 | [*: support show tiflash config](https://github.com/pingcap/tidb/pull/28838)             |   | 6 | 3h  |
| tidb/28828 | [planner: remove surplus FindColumnInfoByID](https://github.com/pingcap/tidb/pull/28828) |   | 6 | 18h |


</details> 


<details> 
  <summary>lzmhhh123's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                        PR                                                                        | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/22631 | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                                                                  |   | 262d23h |
| tidb/26005 | [expression: fix cast string like '.1a1' to decimal has no warnings information](https://github.com/pingcap/tidb/pull/26005)                     |   | 106d13h |
| tidb/26152 | [types: year function can't handle some date string](https://github.com/pingcap/tidb/pull/26152)                                                 |   | 100d14h |
| tidb/27212 | [planner: fix wrong charset about union result of date type and int](https://github.com/pingcap/tidb/pull/27212)                                 |   | 68d14h  |
| tidb/27611 | [types: fix incorrect return type about if function when argument type contains bit](https://github.com/pingcap/tidb/pull/27611)                 |   | 55d14h  |
| tidb/27832 | [executor: fix a bug that can not insert null into a not null column in the empty SQL mode (#21237)](https://github.com/pingcap/tidb/pull/27832) |   | 44d16h  |
| tikv/10616 | [copr: fix Max/Min bug when comparing signed and unsigned int64 (#10167)](https://github.com/tikv/tikv/pull/10616)                               |   | 89d21h  |
| tidb/27954 | [planner: Fix Empty string has different meanings in SELECT and UPDATE](https://github.com/pingcap/tidb/pull/27954)                              |   | 40d16h  |
| tikv/10617 | [copr: fix Max/Min bug when comparing signed and unsigned int64 (#10167)](https://github.com/tikv/tikv/pull/10617)                               |   | 89d21h  |
| tidb/28499 | [expression: align null flag of union columns and constants](https://github.com/pingcap/tidb/pull/28499)                                         |   | 20d15h  |
| tidb/28651 | [expression: not push invalid cast to tiflash (#28458)](https://github.com/pingcap/tidb/pull/28651)                                              |   | 12d18h  |
| tidb/28652 | [expression: not push invalid cast to tiflash (#28458)](https://github.com/pingcap/tidb/pull/28652)                                              |   | 12d18h  |
| tidb/28656 | [distsql: fix copr cache events metric](https://github.com/pingcap/tidb/pull/28656)                                                              |   | 12d17h  |
| tidb/28813 | [expression: simplify canFuncBePushed logic ](https://github.com/pingcap/tidb/pull/28813)                                                        |   | 6d19h   |
| tidb/28888 | [telemetry: nitpick for identical branches](https://github.com/pingcap/tidb/pull/28888)                                                          |   | 4d20h   |
| tidb/28924 | [planner: reject non-top sort for MPP](https://github.com/pingcap/tidb/pull/28924)                                                               |   | 2d14h   |


## Reviewed in Last 7 Days

|    REPO    |                                                PR                                                 | C | D |   R   |
|------------|---------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/28838 | [*: support show tiflash config](https://github.com/pingcap/tidb/pull/28838)                      |   | 6 | 2h    |
| tidb/28712 | [expression: add extra enum info for push down check](https://github.com/pingcap/tidb/pull/28712) |   | 6 | 3d18h |


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

|     REPO     |                                                                  PR                                                                  | C | LASTED  |
|--------------|--------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| docs-cn/5561 | [Add sql optimization-related docs to toc](https://github.com/pingcap/docs-cn/pull/5561)                                             |   | 240d15h |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                       |   | 342d17h |
| docs-cn/7237 | [Add restriction information for index merge to avoid misuse](https://github.com/pingcap/docs-cn/pull/7237)                          |   | 20d19h  |
| tidb/23590   | [planner, table: optimize the list partition pruner for range query](https://github.com/pingcap/tidb/pull/23590)                     |   | 208d16h |
| tidb/26323   | [planner: use multi-layer projections for subquery selection (#8190)](https://github.com/pingcap/tidb/pull/26323)                    |   | 94d6h   |
| tidb/26440   | [executor: a HashJoin demo in exchange parallel framework](https://github.com/pingcap/tidb/pull/26440)                               |   | 91d17h  |
| tidb/26499   | [planner: fix the unstable unit test `TestAnalyzeIncremental` (#26460)](https://github.com/pingcap/tidb/pull/26499)                  |   | 89d20h  |
| tidb/27315   | [go.mod: update parser to fix the parse error for subquery (#25647)](https://github.com/pingcap/tidb/pull/27315)                     |   | 64d13h  |
| tidb/27396   | [*: set consistent assertion for DML](https://github.com/pingcap/tidb/pull/27396)                                                    |   | 62d13h  |
| tidb/28038   | [*: collect column stats usage and periodically dump to the system table](https://github.com/pingcap/tidb/pull/28038)                |   | 36d16h  |
| tidb/28295   | [planner: keep the original join schema in predicate pushdown (#24862)](https://github.com/pingcap/tidb/pull/28295)                  |   | 27d16h  |
| tidb/28333   | [executor: fix detaching from GlobalTracker before executing select query](https://github.com/pingcap/tidb/pull/28333)               |   | 25d15h  |
| tidb/28428   | [*: support show column_stats_usage](https://github.com/pingcap/tidb/pull/28428)                                                     |   | 23d16h  |
| tidb/28666   | [executor: fill extra partition ID column in UnionScan executor](https://github.com/pingcap/tidb/pull/28666)                         |   | 12d11h  |
| tidb/28719   | [statistics: fix auto analyze triggered out of specified time (#28703)](https://github.com/pingcap/tidb/pull/28719)                  |   | 9d16h   |
| tidb/28721   | [statistics: fix auto analyze triggered out of specified time (#28703)](https://github.com/pingcap/tidb/pull/28721)                  |   | 9d16h   |
| tidb/28723   | [statistics: fix auto analyze triggered out of specified time (#28703)](https://github.com/pingcap/tidb/pull/28723)                  |   | 9d16h   |
| tidb/28744   | [planner: clone possible properties before saving them is unnecessary](https://github.com/pingcap/tidb/pull/28744)                   |   | 8d23h   |
| tidb/28819   | [planner: fix the wrong partition pruning when some conditions is out of range (#28296)](https://github.com/pingcap/tidb/pull/28819) |   | 6d18h   |
| tidb/28840   | [planner: add more test cases for plan cache](https://github.com/pingcap/tidb/pull/28840)                                            |   | 5d22h   |
| tidb/28878   | [docs: add design doc for analyze predicate columns](https://github.com/pingcap/tidb/pull/28878)                                     |   | 5d16h   |
| tidb/28944   | [planner: allow refineArgs for plan cache in some situations](https://github.com/pingcap/tidb/pull/28944)                            |   | 1d20h   |
| tidb/28953   | [sessionctx: Fix SET GLOBAL tidb_skip_isolation_level_check=1 (#27898)](https://github.com/pingcap/tidb/pull/28953)                  |   | 1d17h   |
| tidb/28957   | [sessionctx: Fix SET GLOBAL tidb_skip_isolation_level_check=1 (#27898)](https://github.com/pingcap/tidb/pull/28957)                  |   | 1d16h   |
| tidb/28984   | [config, planner: Change prepared-plan-cache.capacity default size from 100 to 1000](https://github.com/pingcap/tidb/pull/28984)     |   | 15h     |


## Reviewed in Last 7 Days

|    REPO    |                                                                  PR                                                                  | C | D |   R    |
|------------|--------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/28956 | [executor: refactor plan replayer](https://github.com/pingcap/tidb/pull/28956)                                                       |   | 1 | 1d0h   |
| tidb/27849 | [session: add system table mysql.column_stats_usage](https://github.com/pingcap/tidb/pull/27849)                                     |   | 3 | 41d5h  |
| tidb/28820 | [planner: fix the wrong partition pruning when some conditions is out of range (#28296)](https://github.com/pingcap/tidb/pull/28820) |   | 3 | 3d22h  |
| tidb/28826 | [*: rename RECREATOR to REPLAYER](https://github.com/pingcap/tidb/pull/28826)                                                        |   | 6 | 21h    |
| tidb/28774 | [planner: support rebuild the range of indexMerge when reuse the plan](https://github.com/pingcap/tidb/pull/28774)                   |   | 6 | 1d23h  |
| tidb/28799 | [domain, session: add plan replayer gc](https://github.com/pingcap/tidb/pull/28799)                                                  |   | 7 | 9h     |
| tidb/28296 | [planner: fix the wrong partition pruning when some conditions is out of range](https://github.com/pingcap/tidb/pull/28296)          |   | 7 | 20d20h |


</details> 


<details> 
  <summary>rebelice's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                                               PR                                                                                               | C | LASTED  |
|--------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| docs/5185    | [sql-statements, information-schema: add `END_TIME` field for table `ANALYZE_STATUS`](https://github.com/pingcap/docs/pull/5185)                                                               |   | 202d18h |
| docs-cn/5916 | [sql-statements, information-schema: add `END_TIME` field for table `ANALYZE_STATUS`](https://github.com/pingcap/docs-cn/pull/5916)                                                            |   | 202d17h |
| tidb/24033   | [statistics: fix some unstable tests in global stats (#23502)](https://github.com/pingcap/tidb/pull/24033)                                                                                     |   | 189d9h  |
| tidb/24669   | [planner: fix "order by + num " can use a column not in select fields](https://github.com/pingcap/tidb/pull/24669)                                                                             |   | 159d16h |
| tidb/26474   | [planner: fix the unstable unit test TestTableFromMeta (#26463)](https://github.com/pingcap/tidb/pull/26474)                                                                                   |   | 90d16h  |
| tidb/26475   | [planner: fix the unstable unit test TestTableFromMeta (#26463)](https://github.com/pingcap/tidb/pull/26475)                                                                                   |   | 90d16h  |
| tidb/26491   | [planner: fix the unstable test TestOrderedResultModeOnOtherOperators (#26481)](https://github.com/pingcap/tidb/pull/26491)                                                                    |   | 89d23h  |
| tidb/26492   | [planner: fix the unstable test TestOrderedResultModeOnOtherOperators (#26481)](https://github.com/pingcap/tidb/pull/26492)                                                                    |   | 89d23h  |
| tidb/26498   | [planner: fix the unstable unit test `TestAnalyzeIncremental` (#26460)](https://github.com/pingcap/tidb/pull/26498)                                                                            |   | 89d20h  |
| tidb/26499   | [planner: fix the unstable unit test `TestAnalyzeIncremental` (#26460)](https://github.com/pingcap/tidb/pull/26499)                                                                            |   | 89d20h  |
| tidb/26505   | [planner: fix goroutine leak problem in some unit tests (#26500)](https://github.com/pingcap/tidb/pull/26505)                                                                                  |   | 89d19h  |
| tidb/28038   | [*: collect column stats usage and periodically dump to the system table](https://github.com/pingcap/tidb/pull/28038)                                                                          |   | 36d16h  |
| tidb/28317   | [planner: remove duplicate predicates in the Selection operator](https://github.com/pingcap/tidb/pull/28317)                                                                                   |   | 26d8h   |
| tidb/28819   | [planner: fix the wrong partition pruning when some conditions is out of range (#28296)](https://github.com/pingcap/tidb/pull/28819)                                                           |   | 6d18h   |
| tidb/28836   | [planner: fix the issue that plan-cache cannot be aware of changes of unsigned flags (#28827)](https://github.com/pingcap/tidb/pull/28836)                                                     |   | 5d23h   |
| tidb/28840   | [planner: add more test cases for plan cache](https://github.com/pingcap/tidb/pull/28840)                                                                                                      |   | 5d22h   |
| tidb/28878   | [docs: add design doc for analyze predicate columns](https://github.com/pingcap/tidb/pull/28878)                                                                                               |   | 5d16h   |
| tidb/28944   | [planner: allow refineArgs for plan cache in some situations](https://github.com/pingcap/tidb/pull/28944)                                                                                      |   | 1d20h   |
| tidb/28992   | [planner: add an extra safe-guard selection upon DataSource to prevent wrong results caused by wrong rebuilt range when using plan-cache (#28976)](https://github.com/pingcap/tidb/pull/28992) |   | 6h      |
| tidb/28993   | [planner: add an extra safe-guard selection upon DataSource to prevent wrong results caused by wrong rebuilt range when using plan-cache (#28976)](https://github.com/pingcap/tidb/pull/28993) |   | 6h      |
| tidb/28994   | [planner: add an extra safe-guard selection upon DataSource to prevent wrong results caused by wrong rebuilt range when using plan-cache (#28976)](https://github.com/pingcap/tidb/pull/28994) |   | 6h      |


## Reviewed in Last 7 Days

|    REPO    |                                                                                          PR                                                                                           | C | D |   R    |
|------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/28976 | [planner: add an extra safe-guard selection upon DataSource to prevent wrong results caused by wrong rebuilt range when using plan-cache](https://github.com/pingcap/tidb/pull/28976) |   | 1 | 1h     |
| tidb/28938 | [planner: fix the issue that plan-cache may return wrong results when the IndexLookup depends on a generated column](https://github.com/pingcap/tidb/pull/28938)                      |   | 2 | 0h     |
| tidb/28774 | [planner: support rebuild the range of indexMerge when reuse the plan](https://github.com/pingcap/tidb/pull/28774)                                                                    |   | 6 | 2d3h   |
| tidb/28827 | [planner: fix the issue that plan-cache cannot be aware of changes of unsigned flags](https://github.com/pingcap/tidb/pull/28827)                                                     |   | 7 | 0h     |
| tidb/28821 | [*: fix typo and address unfinished comment for plan recreator](https://github.com/pingcap/tidb/pull/28821)                                                                           |   | 7 | 0h     |
| tidb/28790 | [planner: forbid constant fold when plan cache enable](https://github.com/pingcap/tidb/pull/28790)                                                                                    |   | 7 | 22h    |
| tidb/28296 | [planner: fix the wrong partition pruning when some conditions is out of range](https://github.com/pingcap/tidb/pull/28296)                                                           |   | 7 | 20d20h |


</details> 


<details> 
  <summary>time-and-fate's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                          PR                                                                           | C | LASTED  |
|------------|-------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/22416 | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                                                     | Y | 277d12h |
| tidb/25390 | [planner/core: fix `isTableAliasDuplicate`, use `schema.name` as key when table has a alias name](https://github.com/pingcap/tidb/pull/25390)         |   | 128d19h |
| tidb/26474 | [planner: fix the unstable unit test TestTableFromMeta (#26463)](https://github.com/pingcap/tidb/pull/26474)                                          |   | 90d16h  |
| tidb/26475 | [planner: fix the unstable unit test TestTableFromMeta (#26463)](https://github.com/pingcap/tidb/pull/26475)                                          |   | 90d16h  |
| tidb/26498 | [planner: fix the unstable unit test `TestAnalyzeIncremental` (#26460)](https://github.com/pingcap/tidb/pull/26498)                                   |   | 89d20h  |
| tidb/26499 | [planner: fix the unstable unit test `TestAnalyzeIncremental` (#26460)](https://github.com/pingcap/tidb/pull/26499)                                   |   | 89d20h  |
| tidb/26713 | [planner: use the converted datum based on the target column to point get](https://github.com/pingcap/tidb/pull/26713)                                |   | 84d11h  |
| tidb/27773 | [statistics: remove redundant assignment for statistics.Column.Count](https://github.com/pingcap/tidb/pull/27773)                                     |   | 48d16h  |
| tidb/28295 | [planner: keep the original join schema in predicate pushdown (#24862)](https://github.com/pingcap/tidb/pull/28295)                                   |   | 27d16h  |
| tidb/28428 | [*: support show column_stats_usage](https://github.com/pingcap/tidb/pull/28428)                                                                      |   | 23d16h  |
| tidb/28444 | [planner: fix the issue that planner may cache invalid plans for joins in some cases (#28432)](https://github.com/pingcap/tidb/pull/28444)            |   | 22d23h  |
| tidb/28445 | [planner: fix the issue that planner may cache invalid plans for joins in some cases (#28432)](https://github.com/pingcap/tidb/pull/28445)            |   | 22d23h  |
| tidb/28446 | [planner: fix the issue that planner may cache invalid plans for joins in some cases (#28432)](https://github.com/pingcap/tidb/pull/28446)            |   | 22d23h  |
| tidb/28491 | [util/ranger: check boundary condition when taking intersection of two columnValues](https://github.com/pingcap/tidb/pull/28491)                      |   | 20d20h  |
| tidb/28554 | [planner, statistics, sessionctx: add variable to enable/disable the outdated statistics to pseudo logic](https://github.com/pingcap/tidb/pull/28554) |   | 13d20h  |
| tidb/28819 | [planner: fix the wrong partition pruning when some conditions is out of range (#28296)](https://github.com/pingcap/tidb/pull/28819)                  |   | 6d18h   |
| tidb/28878 | [docs: add design doc for analyze predicate columns](https://github.com/pingcap/tidb/pull/28878)                                                      |   | 5d16h   |
| tidb/28924 | [planner: reject non-top sort for MPP](https://github.com/pingcap/tidb/pull/28924)                                                                    |   | 2d14h   |


## Reviewed in Last 7 Days

|    REPO    |                                                                     PR                                                                     | C | D |   R   |
|------------|--------------------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/28820 | [planner: fix the wrong partition pruning when some conditions is out of range (#28296)](https://github.com/pingcap/tidb/pull/28820)       |   | 3 | 3d22h |
| tidb/28837 | [planner: fix the issue that plan-cache cannot be aware of changes of unsigned flags (#28827)](https://github.com/pingcap/tidb/pull/28837) |   | 3 | 3d0h  |
| tidb/28799 | [domain, session: add plan replayer gc](https://github.com/pingcap/tidb/pull/28799)                                                        |   | 6 | 1d1h  |
| tidb/28821 | [*: fix typo and address unfinished comment for plan recreator](https://github.com/pingcap/tidb/pull/28821)                                |   | 7 | 1h    |


</details> 


<details> 
  <summary>winoros's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                              PR                                                                              | C | LASTED  |
|--------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| docs-cn/5916 | [sql-statements, information-schema: add `END_TIME` field for table `ANALYZE_STATUS`](https://github.com/pingcap/docs-cn/pull/5916)                          |   | 202d17h |
| tidb/20903   | [planner: fix confused and unnecessary double-projection in plans.](https://github.com/pingcap/tidb/pull/20903)                                              |   | 348d17h |
| docs/5783    | [migration: Add information about Vitess to TiDB migration](https://github.com/pingcap/docs/pull/5783)                                                       |   | 128d5h  |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                               |   | 342d17h |
| tidb/22416   | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                                                            | Y | 277d12h |
| tidb/22478   | [planner, executor: fix query partition table with global unique index get wrong result](https://github.com/pingcap/tidb/pull/22478)                         |   | 272d13h |
| tidb/24138   | [planner: Add Equivalence Rules to Transform BinaryOptSubquery to ExistsSubquery](https://github.com/pingcap/tidb/pull/24138)                                |   | 184d12h |
| tidb/26323   | [planner: use multi-layer projections for subquery selection (#8190)](https://github.com/pingcap/tidb/pull/26323)                                            |   | 94d6h   |
| tidb/26474   | [planner: fix the unstable unit test TestTableFromMeta (#26463)](https://github.com/pingcap/tidb/pull/26474)                                                 |   | 90d16h  |
| tidb/26475   | [planner: fix the unstable unit test TestTableFromMeta (#26463)](https://github.com/pingcap/tidb/pull/26475)                                                 |   | 90d16h  |
| tidb/26492   | [planner: fix the unstable test TestOrderedResultModeOnOtherOperators (#26481)](https://github.com/pingcap/tidb/pull/26492)                                  |   | 89d23h  |
| tidb/26503   | [planner: fix goroutine leak problem in some unit tests (#26500)](https://github.com/pingcap/tidb/pull/26503)                                                |   | 89d19h  |
| tidb/26505   | [planner: fix goroutine leak problem in some unit tests (#26500)](https://github.com/pingcap/tidb/pull/26505)                                                |   | 89d19h  |
| tidb/27636   | [planner, expression: avoid exprs with side effects in column pruning and agg pushdown (#27370)](https://github.com/pingcap/tidb/pull/27636)                 |   | 54d17h  |
| tidb/28295   | [planner: keep the original join schema in predicate pushdown (#24862)](https://github.com/pingcap/tidb/pull/28295)                                          |   | 27d16h  |
| tidb/28491   | [util/ranger: check boundary condition when taking intersection of two columnValues](https://github.com/pingcap/tidb/pull/28491)                             |   | 20d20h  |
| tidb/28554   | [planner, statistics, sessionctx: add variable to enable/disable the outdated statistics to pseudo logic](https://github.com/pingcap/tidb/pull/28554)        |   | 13d20h  |
| tidb/28719   | [statistics: fix auto analyze triggered out of specified time (#28703)](https://github.com/pingcap/tidb/pull/28719)                                          |   | 9d16h   |
| tidb/28721   | [statistics: fix auto analyze triggered out of specified time (#28703)](https://github.com/pingcap/tidb/pull/28721)                                          |   | 9d16h   |
| tidb/28723   | [statistics: fix auto analyze triggered out of specified time (#28703)](https://github.com/pingcap/tidb/pull/28723)                                          |   | 9d16h   |
| tidb/28744   | [planner: clone possible properties before saving them is unnecessary](https://github.com/pingcap/tidb/pull/28744)                                           |   | 8d23h   |
| tidb/28748   | [planner/cascades: fill group stats](https://github.com/pingcap/tidb/pull/28748)                                                                             |   | 8d22h   |
| tidb/28759   | [planner: make constant propagation more stringently](https://github.com/pingcap/tidb/pull/28759)                                                            |   | 8d17h   |
| tidb/28796   | [statistics: migrate test-infra to testify for selectivity_test.go](https://github.com/pingcap/tidb/pull/28796)                                              |   | 7d10h   |
| tidb/28878   | [docs: add design doc for analyze predicate columns](https://github.com/pingcap/tidb/pull/28878)                                                             |   | 5d16h   |
| tidb/28879   | [planner: make the error of access path miss more user friendly](https://github.com/pingcap/tidb/pull/28879)                                                 |   | 5d15h   |
| tidb/28924   | [planner: reject non-top sort for MPP](https://github.com/pingcap/tidb/pull/28924)                                                                           |   | 2d14h   |
| tidb/28927   | [planner: fix the issue that cached IndexJoin plans may return wrong results after changing parameters (#28915)](https://github.com/pingcap/tidb/pull/28927) |   | 2d9h    |
| tidb/28928   | [planner: fix the issue that cached IndexJoin plans may return wrong results after changing parameters (#28915)](https://github.com/pingcap/tidb/pull/28928) |   | 2d9h    |


## Reviewed in Last 7 Days

|    REPO    |                                                                         PR                                                                          | C | D |   R    |
|------------|-----------------------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/28987 | [table-level analyze options for both manual and auto analyze](https://github.com/pingcap/tidb/pull/28987)                                          |   | 1 | 4h     |
| tidb/28428 | [*: support show column_stats_usage](https://github.com/pingcap/tidb/pull/28428)                                                                    |   | 1 | 23d1h  |
| tidb/28038 | [*: collect column stats usage and periodically dump to the system table](https://github.com/pingcap/tidb/pull/28038)                               |   | 2 | 34d23h |
| tidb/28915 | [planner: fix the issue that cached IndexJoin plans may return wrong results after changing parameters](https://github.com/pingcap/tidb/pull/28915) |   | 3 | 0h     |
| tidb/28846 | [parser: analyze predicate/user-specified columns and show column_stats_usage](https://github.com/pingcap/tidb/pull/28846)                          |   | 3 | 3d1h   |


</details> 


<details> 
  <summary>wshwsh12's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                           PR                                                           | C | LASTED  |
|------------|------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/21401 | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                     |   | 324d11h |
| tidb/21887 | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                        |   | 305d11h |
| tidb/27562 | [executor: Support partition spilling for HashAgg](https://github.com/pingcap/tidb/pull/27562)                         |   | 56d23h  |
| tidb/27837 | [planner: fix wrong plan caused by shallow copy schema columns (#27798)](https://github.com/pingcap/tidb/pull/27837)   |   | 44d16h  |
| tidb/27992 | [planner: add sub plan info of shuffleReceiver when query explain analyze](https://github.com/pingcap/tidb/pull/27992) |   | 37d16h  |
| tidb/28333 | [executor: fix detaching from GlobalTracker before executing select query](https://github.com/pingcap/tidb/pull/28333) |   | 25d15h  |


## Reviewed in Last 7 Days

|    REPO    |                                              PR                                               | C | D | R  |
|------------|-----------------------------------------------------------------------------------------------|---|---|----|
| tidb/28936 | [expression, parser: fix the result of trim3Args](https://github.com/pingcap/tidb/pull/28936) |   | 2 | 6h |


</details> 

