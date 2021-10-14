|   REVIEWER    |  O(C)   | 1 | 2 | 3 | 4 | 5 | 6 | 7 | T  |
|---------------|---------|---|---|---|---|---|---|---|----|
| Reminiscent   | 12 ( 0) | 2 | 0 | 1 | 0 | 1 | 0 | 0 |  4 |
| SunRunAway    |  4 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| XuHuaiyu      | 12 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| dyzsr         |  0 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| eurekaka      |  4 ( 1) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| fzhedu        |  7 ( 0) | 5 | 0 | 3 | 0 | 0 | 0 | 0 |  8 |
| ichn-hu       |  9 ( 0) | 1 | 3 | 2 | 0 | 0 | 0 | 0 |  6 |
| lzmhhh123     | 14 ( 0) | 0 | 1 | 0 | 0 | 0 | 3 | 0 |  4 |
| nullnotnil    |  0 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |
| qw4990        | 21 ( 0) | 2 | 5 | 1 | 0 | 2 | 2 | 0 | 12 |
| rebelice      | 15 ( 0) | 0 | 1 | 0 | 0 | 0 | 1 | 1 |  3 |
| time-and-fate | 16 ( 1) | 0 | 1 | 0 | 0 | 0 | 0 | 0 |  1 |
| winoros       | 26 ( 1) | 0 | 5 | 3 | 0 | 3 | 1 | 0 | 12 |
| wshwsh12      |  5 ( 0) | 0 | 0 | 0 | 0 | 0 | 0 | 0 |  0 |


<details> 
  <summary>Reminiscent's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                      PR                                                                      | C | LASTED  |
|------------|----------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/23590 | [planner, table: optimize the list partition pruner for range query](https://github.com/pingcap/tidb/pull/23590)                             |   | 201d16h |
| tidb/26474 | [planner: fix the unstable unit test TestTableFromMeta (#26463)](https://github.com/pingcap/tidb/pull/26474)                                 |   | 83d16h  |
| tidb/26475 | [planner: fix the unstable unit test TestTableFromMeta (#26463)](https://github.com/pingcap/tidb/pull/26475)                                 |   | 83d16h  |
| tidb/26491 | [planner: fix the unstable test TestOrderedResultModeOnOtherOperators (#26481)](https://github.com/pingcap/tidb/pull/26491)                  |   | 82d23h  |
| tidb/26492 | [planner: fix the unstable test TestOrderedResultModeOnOtherOperators (#26481)](https://github.com/pingcap/tidb/pull/26492)                  |   | 82d23h  |
| tidb/26498 | [planner: fix the unstable unit test `TestAnalyzeIncremental` (#26460)](https://github.com/pingcap/tidb/pull/26498)                          |   | 82d20h  |
| tidb/26499 | [planner: fix the unstable unit test `TestAnalyzeIncremental` (#26460)](https://github.com/pingcap/tidb/pull/26499)                          |   | 82d20h  |
| tidb/26503 | [planner: fix goroutine leak problem in some unit tests (#26500)](https://github.com/pingcap/tidb/pull/26503)                                |   | 82d19h  |
| tidb/27636 | [planner, expression: avoid exprs with side effects in column pruning and agg pushdown (#27370)](https://github.com/pingcap/tidb/pull/27636) |   | 47d17h  |
| tidb/27773 | [statistics: remove redundant assignment for statistics.Column.Count](https://github.com/pingcap/tidb/pull/27773)                            |   | 41d16h  |
| tidb/27837 | [planner: fix wrong plan caused by shallow copy schema columns (#27798)](https://github.com/pingcap/tidb/pull/27837)                         |   | 37d16h  |
| tidb/27849 | [session: add system table mysql.column_stats_usage](https://github.com/pingcap/tidb/pull/27849)                                             |   | 36d23h  |


## Reviewed in Last 7 Days

|     REPO     |                                                                     PR                                                                     | C | D |   R    |
|--------------|--------------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| docs-cn/7261 | [Fix and improve some descriptions about optimizer](https://github.com/pingcap/docs-cn/pull/7261)                                          |   | 1 | 1h     |
| tidb/28543   | [planner/core: migrate test-infra to testify for `planner/core/integration_partition_test.go`](https://github.com/pingcap/tidb/pull/28543) |   | 1 | 9d16h  |
| docs-cn/6948 | [spm: add description for baseline capture filter](https://github.com/pingcap/docs-cn/pull/6948)                                           |   | 3 | 46d22h |
| tidb/28478   | [planner: fix the issue that some PointGet plans generated in physical-stage cannot be cached](https://github.com/pingcap/tidb/pull/28478) |   | 5 | 9d21h  |


</details> 


<details> 
  <summary>SunRunAway's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                        PR                                                                        | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/21834 | [planner: enhanced index range calculation plan](https://github.com/pingcap/tidb/pull/21834)                                                     |   | 301d18h |
| tidb/21956 | [planner/preprocessor: disallow into-outfile clause in some place](https://github.com/pingcap/tidb/pull/21956)                                   |   | 294d23h |
| tidb/25385 | [executor: global kill 32bits (local connID part)](https://github.com/pingcap/tidb/pull/25385)                                                   |   | 122d10h |
| tidb/27832 | [executor: fix a bug that can not insert null into a not null column in the empty SQL mode (#21237)](https://github.com/pingcap/tidb/pull/27832) |   | 37d16h  |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>XuHuaiyu's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                               PR                                                                | C | LASTED  |
|--------------|---------------------------------------------------------------------------------------------------------------------------------|---|---------|
| docs-cn/5561 | [Add sql optimization-related docs to toc](https://github.com/pingcap/docs-cn/pull/5561)                                        |   | 233d15h |
| tidb/21401   | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                              |   | 317d11h |
| docs-cn/6716 | [sysvar: add doc for tidb-restricted-read-only](https://github.com/pingcap/docs-cn/pull/6716)                                   |   | 83d18h  |
| tidb/26098   | [executor, planner: add support for SQL_CALC_FOUND_ROWS](https://github.com/pingcap/tidb/pull/26098)                            |   | 95d23h  |
| tidb/26440   | [executor: a HashJoin demo in exchange parallel framework](https://github.com/pingcap/tidb/pull/26440)                          |   | 84d16h  |
| tidb/27315   | [go.mod: update parser to fix the parse error for subquery (#25647)](https://github.com/pingcap/tidb/pull/27315)                |   | 57d13h  |
| tidb/27378   | [distsql: fix goroutine/memory leak for streaming when query is cancelled (#27354)](https://github.com/pingcap/tidb/pull/27378) |   | 55d18h  |
| tidb/27396   | [*: set consistent assertion for DML](https://github.com/pingcap/tidb/pull/27396)                                               |   | 55d13h  |
| tidb/27403   | [expression: round function for int should use round half up rule](https://github.com/pingcap/tidb/pull/27403)                  |   | 55d11h  |
| tidb/27561   | [server, privilege: Socket authentication](https://github.com/pingcap/tidb/pull/27561)                                          |   | 50d4h   |
| tidb/27992   | [planner: add sub plan info of shuffleReceiver when query explain analyze](https://github.com/pingcap/tidb/pull/27992)          |   | 30d15h  |
| tidb/28649   | [expression: limit valid decimal length (#28466)](https://github.com/pingcap/tidb/pull/28649)                                   |   | 5d19h   |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


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
| tidb/22416 | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416) | Y | 270d11h |
| tidb/23316 | [planner: Fix rebuild range for prepared plan](https://github.com/pingcap/tidb/pull/23316)        |   | 212d17h |
| tidb/27099 | [planner: support expression index for view](https://github.com/pingcap/tidb/pull/27099)          |   | 63d19h  |
| tidb/27849 | [session: add system table mysql.column_stats_usage](https://github.com/pingcap/tidb/pull/27849)  |   | 36d23h  |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 


<details> 
  <summary>fzhedu's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                         PR                                                                         | C | LASTED |
|------------|----------------------------------------------------------------------------------------------------------------------------------------------------|---|--------|
| tidb/28147 | [planner: fix can not found column bug (#28067)](https://github.com/pingcap/tidb/pull/28147)                                                       |   | 26d18h |
| tidb/28262 | [distsql: avoid false positive error log about `invalid cop task execution summaries length` (#28188)](https://github.com/pingcap/tidb/pull/28262) |   | 21d16h |
| tidb/28263 | [distsql: avoid false positive error log about `invalid cop task execution summaries length` (#28188)](https://github.com/pingcap/tidb/pull/28263) |   | 21d16h |
| tidb/28287 | [copr: Fix bug that mpp node availability detect does not work in some corner cases (#28201)](https://github.com/pingcap/tidb/pull/28287)          |   | 20d20h |
| tidb/28288 | [copr: Fix bug that mpp node availability detect does not work in some corner cases (#28201)](https://github.com/pingcap/tidb/pull/28288)          |   | 20d20h |
| tidb/28651 | [expression: not push invalid cast to tiflash (#28458)](https://github.com/pingcap/tidb/pull/28651)                                                |   | 5d18h  |
| tidb/28652 | [expression: not push invalid cast to tiflash (#28458)](https://github.com/pingcap/tidb/pull/28652)                                                |   | 5d18h  |


## Reviewed in Last 7 Days

|     REPO     |                                                                         PR                                                                         | C | D |   R    |
|--------------|----------------------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/28654   | [expression: not push invalid cast to tiflash (#28458)](https://github.com/pingcap/tidb/pull/28654)                                                |   | 1 | 4d22h  |
| tidb/28140   | [copr: avoid NPE when store is not available when balance batch cop task (#28110)](https://github.com/pingcap/tidb/pull/28140)                     |   | 1 | 26d0h  |
| tidb/28149   | [planner: fix can not found column bug (#28067)](https://github.com/pingcap/tidb/pull/28149)                                                       |   | 1 | 25d22h |
| tidb/28264   | [distsql: avoid false positive error log about `invalid cop task execution summaries length` (#28188)](https://github.com/pingcap/tidb/pull/28264) |   | 1 | 20d20h |
| tidb/28289   | [copr: Fix bug that mpp node availability detect does not work in some corner cases (#28201)](https://github.com/pingcap/tidb/pull/28289)          |   | 1 | 20d0h  |
| tidb/28742   | [sessionctx: fix data-race bug when alloc task id (#28022)](https://github.com/pingcap/tidb/pull/28742)                                            |   | 3 | 0h     |
| tidb/28740   | [distsql: avoid false positive error log about `invalid cop task execution summaries length](https://github.com/pingcap/tidb/pull/28740)           |   | 3 | 0h     |
| tiflash/1787 | [disable some tests](https://github.com/pingcap/tiflash/pull/1787)                                                                                 |   | 3 | 0h     |


</details> 


<details> 
  <summary>ichn-hu's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                         PR                                                                         | C | LASTED  |
|--------------|----------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20903   | [planner: fix confused and unnecessary double-projection in plans.](https://github.com/pingcap/tidb/pull/20903)                                    |   | 341d17h |
| docs-cn/7238 | [system-variables: correct the description of tidb_allow_fallback_to_tikv](https://github.com/pingcap/docs-cn/pull/7238)                           |   | 13d19h  |
| tidb/22631   | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                                                                    |   | 255d23h |
| tidb/27119   | [executor: fix json_objectagg() on varbinary type](https://github.com/pingcap/tidb/pull/27119)                                                     |   | 63d16h  |
| tidb/27403   | [expression: round function for int should use round half up rule](https://github.com/pingcap/tidb/pull/27403)                                     |   | 55d11h  |
| tidb/27451   | [expression: fix wrong result for date add sub (#27244)](https://github.com/pingcap/tidb/pull/27451)                                               |   | 54d16h  |
| tidb/28262   | [distsql: avoid false positive error log about `invalid cop task execution summaries length` (#28188)](https://github.com/pingcap/tidb/pull/28262) |   | 21d16h  |
| tidb/28666   | [executor: fill extra partition ID column in UnionScan executor](https://github.com/pingcap/tidb/pull/28666)                                       |   | 5d11h   |
| tidb/28712   | [expression: add extra enum info for push down check](https://github.com/pingcap/tidb/pull/28712)                                                  |   | 2d18h   |


## Reviewed in Last 7 Days

|    REPO    |                                                                         PR                                                                         | C | D |   R    |
|------------|----------------------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/28784 | [*: run go fmt in go1.17](https://github.com/pingcap/tidb/pull/28784)                                                                              |   | 1 | 5h     |
| tidb/28648 | [expression: limit valid decimal length (#28466)](https://github.com/pingcap/tidb/pull/28648)                                                      |   | 2 | 4d4h   |
| tidb/28649 | [expression: limit valid decimal length (#28466)](https://github.com/pingcap/tidb/pull/28649)                                                      |   | 2 | 4d4h   |
| tidb/28647 | [expression: limit valid decimal length (#28466)](https://github.com/pingcap/tidb/pull/28647)                                                      |   | 2 | 4d4h   |
| tidb/28263 | [distsql: avoid false positive error log about `invalid cop task execution summaries length` (#28188)](https://github.com/pingcap/tidb/pull/28263) |   | 3 | 19d16h |
| tidb/28264 | [distsql: avoid false positive error log about `invalid cop task execution summaries length` (#28188)](https://github.com/pingcap/tidb/pull/28264) |   | 3 | 19d16h |


</details> 


<details> 
  <summary>lzmhhh123's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                        PR                                                                        | C | LASTED  |
|------------|--------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/22631 | [executor: refine window processor](https://github.com/pingcap/tidb/pull/22631)                                                                  |   | 255d23h |
| tidb/26005 | [expression: fix cast string like '.1a1' to decimal has no warnings information](https://github.com/pingcap/tidb/pull/26005)                     |   | 99d13h  |
| tidb/26152 | [types: year function can't handle some date string](https://github.com/pingcap/tidb/pull/26152)                                                 |   | 93d14h  |
| tidb/27212 | [planner: fix wrong charset about union result of date type and int](https://github.com/pingcap/tidb/pull/27212)                                 |   | 61d14h  |
| tidb/27611 | [types: fix incorrect return type about if function when argument type contains bit](https://github.com/pingcap/tidb/pull/27611)                 |   | 48d14h  |
| tidb/27832 | [executor: fix a bug that can not insert null into a not null column in the empty SQL mode (#21237)](https://github.com/pingcap/tidb/pull/27832) |   | 37d16h  |
| tikv/10616 | [copr: fix Max/Min bug when comparing signed and unsigned int64 (#10167)](https://github.com/tikv/tikv/pull/10616)                               |   | 82d21h  |
| tidb/27954 | [planner: Fix Empty string has different meanings in SELECT and UPDATE](https://github.com/pingcap/tidb/pull/27954)                              |   | 33d16h  |
| tikv/10617 | [copr: fix Max/Min bug when comparing signed and unsigned int64 (#10167)](https://github.com/tikv/tikv/pull/10617)                               |   | 82d21h  |
| tidb/28649 | [expression: limit valid decimal length (#28466)](https://github.com/pingcap/tidb/pull/28649)                                                    |   | 5d19h   |
| tidb/28651 | [expression: not push invalid cast to tiflash (#28458)](https://github.com/pingcap/tidb/pull/28651)                                              |   | 5d18h   |
| tidb/28652 | [expression: not push invalid cast to tiflash (#28458)](https://github.com/pingcap/tidb/pull/28652)                                              |   | 5d18h   |
| tidb/28654 | [expression: not push invalid cast to tiflash (#28458)](https://github.com/pingcap/tidb/pull/28654)                                              |   | 5d18h   |
| tidb/28656 | [distsql: fix copr cache events metric](https://github.com/pingcap/tidb/pull/28656)                                                              |   | 5d17h   |


## Reviewed in Last 7 Days

|    REPO    |                                                    PR                                                    | C | D |   R   |
|------------|----------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/28712 | [expression: add extra enum info for push down check](https://github.com/pingcap/tidb/pull/28712)        |   | 2 | 19h   |
| tidb/28458 | [expression: not push invalid cast to tiflash](https://github.com/pingcap/tidb/pull/28458)               |   | 6 | 9d21h |
| tidb/28499 | [expression: align null flag of union columns and constants](https://github.com/pingcap/tidb/pull/28499) |   | 6 | 7d19h |
| tidb/28466 | [expression: limit valid decimal length](https://github.com/pingcap/tidb/pull/28466)                     |   | 6 | 9d15h |


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

|     REPO     |                                                                         PR                                                                          | C | LASTED  |
|--------------|-----------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| docs-cn/5561 | [Add sql optimization-related docs to toc](https://github.com/pingcap/docs-cn/pull/5561)                                                            |   | 233d15h |
| parser/1329  | [parser: support ANALYZE TABLE t PREDICATE COLUMNS / COLUMN c1 [, c2] ... and SHOW COLUMN_STATS_USAGE](https://github.com/pingcap/parser/pull/1329) |   | 40d15h  |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                      |   | 335d17h |
| docs-cn/7237 | [Add restriction information for index merge to avoid misuse](https://github.com/pingcap/docs-cn/pull/7237)                                         |   | 13d19h  |
| tidb/23590   | [planner, table: optimize the list partition pruner for range query](https://github.com/pingcap/tidb/pull/23590)                                    |   | 201d16h |
| tidb/26323   | [planner: use multi-layer projections for subquery selection (#8190)](https://github.com/pingcap/tidb/pull/26323)                                   |   | 87d6h   |
| tidb/26440   | [executor: a HashJoin demo in exchange parallel framework](https://github.com/pingcap/tidb/pull/26440)                                              |   | 84d16h  |
| tidb/26499   | [planner: fix the unstable unit test `TestAnalyzeIncremental` (#26460)](https://github.com/pingcap/tidb/pull/26499)                                 |   | 82d20h  |
| tidb/27315   | [go.mod: update parser to fix the parse error for subquery (#25647)](https://github.com/pingcap/tidb/pull/27315)                                    |   | 57d13h  |
| tidb/27396   | [*: set consistent assertion for DML](https://github.com/pingcap/tidb/pull/27396)                                                                   |   | 55d13h  |
| tidb/27849   | [session: add system table mysql.column_stats_usage](https://github.com/pingcap/tidb/pull/27849)                                                    |   | 36d23h  |
| tidb/28295   | [planner: keep the original join schema in predicate pushdown (#24862)](https://github.com/pingcap/tidb/pull/28295)                                 |   | 20d16h  |
| tidb/28296   | [planner: fix the wrong partition pruning when some conditions is out of range](https://github.com/pingcap/tidb/pull/28296)                         |   | 20d16h  |
| tidb/28333   | [executor: fix detaching from GlobalTracker before executing select query](https://github.com/pingcap/tidb/pull/28333)                              |   | 18d15h  |
| tidb/28666   | [executor: fill extra partition ID column in UnionScan executor](https://github.com/pingcap/tidb/pull/28666)                                        |   | 5d11h   |
| tidb/28669   | [executor,binlog: fix binlog column mismatch for pessmistic transaction on partition table](https://github.com/pingcap/tidb/pull/28669)             |   | 5d2h    |
| tidb/28719   | [statistics: fix auto analyze triggered out of specified time (#28703)](https://github.com/pingcap/tidb/pull/28719)                                 |   | 2d16h   |
| tidb/28721   | [statistics: fix auto analyze triggered out of specified time (#28703)](https://github.com/pingcap/tidb/pull/28721)                                 |   | 2d16h   |
| tidb/28723   | [statistics: fix auto analyze triggered out of specified time (#28703)](https://github.com/pingcap/tidb/pull/28723)                                 |   | 2d16h   |
| tidb/28744   | [planner: clone possible properties before saving them is unnecessary](https://github.com/pingcap/tidb/pull/28744)                                  |   | 1d23h   |
| tidb/28799   | [domain, session: add plan replayer gc](https://github.com/pingcap/tidb/pull/28799)                                                                 |   | 1h      |


## Reviewed in Last 7 Days

|    REPO    |                                                                         PR                                                                         | C | D |   R    |
|------------|----------------------------------------------------------------------------------------------------------------------------------------------------|---|---|--------|
| tidb/28790 | [planner: forbid constant fold when plan cache enable](https://github.com/pingcap/tidb/pull/28790)                                                 |   | 1 | 0h     |
| tidb/28774 | [planner: support rebuild the range of indexMerge when reuse the plan](https://github.com/pingcap/tidb/pull/28774)                                 |   | 1 | 4h     |
| tidb/28725 | [statistics: fix auto analyze triggered out of specified time (#28703)](https://github.com/pingcap/tidb/pull/28725)                                |   | 2 | 23h    |
| tidb/27302 | [statistics: fix "data too long" error when dumping stats from table with new collation data (#27033)](https://github.com/pingcap/tidb/pull/27302) |   | 2 | 56d0h  |
| tidb/28757 | [statistics: fix "data too long" error when dumping stats from table with new collation data (#27033)](https://github.com/pingcap/tidb/pull/28757) |   | 2 | 0h     |
| tidb/28758 | [statistics: fix auto analyze triggered out of specified time (#28703)](https://github.com/pingcap/tidb/pull/28758)                                |   | 2 | 0h     |
| tidb/28722 | [planner: generate the additional filter in table filter when enable plan cache](https://github.com/pingcap/tidb/pull/28722)                       |   | 2 | 22h    |
| tidb/28703 | [statistics: fix auto analyze triggered out of specified time](https://github.com/pingcap/tidb/pull/28703)                                         |   | 3 | 2h     |
| tidb/28678 | [statistics: avoid mutate global variable](https://github.com/pingcap/tidb/pull/28678)                                                             |   | 5 | 4h     |
| tidb/28655 | [planner: check whether the plan is valid in skylinePruning first](https://github.com/pingcap/tidb/pull/28655)                                     |   | 5 | 22h    |
| tidb/28013 | [executor: add auto id allocator execution runtime stats](https://github.com/pingcap/tidb/pull/28013)                                              |   | 6 | 24d18h |
| tidb/28475 | [planner: disable tiflash plan caching](https://github.com/pingcap/tidb/pull/28475)                                                                |   | 6 | 8d17h  |


</details> 


<details> 
  <summary>rebelice's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                 PR                                                                  | C | LASTED  |
|--------------|-------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| docs/5185    | [sql-statements, information-schema: add `END_TIME` field for table `ANALYZE_STATUS`](https://github.com/pingcap/docs/pull/5185)    |   | 195d17h |
| docs-cn/5916 | [sql-statements, information-schema: add `END_TIME` field for table `ANALYZE_STATUS`](https://github.com/pingcap/docs-cn/pull/5916) |   | 195d17h |
| tidb/24033   | [statistics: fix some unstable tests in global stats (#23502)](https://github.com/pingcap/tidb/pull/24033)                          |   | 182d9h  |
| tidb/24669   | [planner: fix "order by + num " can use a column not in select fields](https://github.com/pingcap/tidb/pull/24669)                  |   | 152d16h |
| tidb/26474   | [planner: fix the unstable unit test TestTableFromMeta (#26463)](https://github.com/pingcap/tidb/pull/26474)                        |   | 83d16h  |
| tidb/26475   | [planner: fix the unstable unit test TestTableFromMeta (#26463)](https://github.com/pingcap/tidb/pull/26475)                        |   | 83d16h  |
| tidb/26491   | [planner: fix the unstable test TestOrderedResultModeOnOtherOperators (#26481)](https://github.com/pingcap/tidb/pull/26491)         |   | 82d23h  |
| tidb/26492   | [planner: fix the unstable test TestOrderedResultModeOnOtherOperators (#26481)](https://github.com/pingcap/tidb/pull/26492)         |   | 82d23h  |
| tidb/26498   | [planner: fix the unstable unit test `TestAnalyzeIncremental` (#26460)](https://github.com/pingcap/tidb/pull/26498)                 |   | 82d20h  |
| tidb/26499   | [planner: fix the unstable unit test `TestAnalyzeIncremental` (#26460)](https://github.com/pingcap/tidb/pull/26499)                 |   | 82d20h  |
| tidb/26505   | [planner: fix goroutine leak problem in some unit tests (#26500)](https://github.com/pingcap/tidb/pull/26505)                       |   | 82d19h  |
| tidb/27849   | [session: add system table mysql.column_stats_usage](https://github.com/pingcap/tidb/pull/27849)                                    |   | 36d23h  |
| tidb/28317   | [planner: remove duplicate predicates in the Selection operator](https://github.com/pingcap/tidb/pull/28317)                        |   | 19d8h   |
| tidb/28774   | [planner: support rebuild the range of indexMerge when reuse the plan](https://github.com/pingcap/tidb/pull/28774)                  |   | 22h     |
| tidb/28790   | [planner: forbid constant fold when plan cache enable](https://github.com/pingcap/tidb/pull/28790)                                  |   | 16h     |


## Reviewed in Last 7 Days

|    REPO    |                                                              PR                                                               | C | D |   R   |
|------------|-------------------------------------------------------------------------------------------------------------------------------|---|---|-------|
| tidb/28722 | [planner: generate the additional filter in table filter when enable plan cache](https://github.com/pingcap/tidb/pull/28722)  |   | 2 | 22h   |
| tidb/28475 | [planner: disable tiflash plan caching](https://github.com/pingcap/tidb/pull/28475)                                           |   | 6 | 8d20h |
| tidb/28452 | [planner/core: migrate test-infra to testify for planner/core/plan_to_pb_test.go](https://github.com/pingcap/tidb/pull/28452) |   | 7 | 9d19h |


</details> 


<details> 
  <summary>time-and-fate's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                                          PR                                                                           | C | LASTED  |
|------------|-------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/22416 | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                                                     | Y | 270d11h |
| tidb/25390 | [planner/core: fix `isTableAliasDuplicate`, use `schema.name` as key when table has a alias name](https://github.com/pingcap/tidb/pull/25390)         |   | 121d19h |
| tidb/26474 | [planner: fix the unstable unit test TestTableFromMeta (#26463)](https://github.com/pingcap/tidb/pull/26474)                                          |   | 83d16h  |
| tidb/26475 | [planner: fix the unstable unit test TestTableFromMeta (#26463)](https://github.com/pingcap/tidb/pull/26475)                                          |   | 83d16h  |
| tidb/26498 | [planner: fix the unstable unit test `TestAnalyzeIncremental` (#26460)](https://github.com/pingcap/tidb/pull/26498)                                   |   | 82d20h  |
| tidb/26499 | [planner: fix the unstable unit test `TestAnalyzeIncremental` (#26460)](https://github.com/pingcap/tidb/pull/26499)                                   |   | 82d20h  |
| tidb/26713 | [planner: use the converted datum based on the target column to point get](https://github.com/pingcap/tidb/pull/26713)                                |   | 77d11h  |
| tidb/27773 | [statistics: remove redundant assignment for statistics.Column.Count](https://github.com/pingcap/tidb/pull/27773)                                     |   | 41d16h  |
| tidb/28295 | [planner: keep the original join schema in predicate pushdown (#24862)](https://github.com/pingcap/tidb/pull/28295)                                   |   | 20d16h  |
| tidb/28444 | [planner: fix the issue that planner may cache invalid plans for joins in some cases (#28432)](https://github.com/pingcap/tidb/pull/28444)            |   | 15d23h  |
| tidb/28445 | [planner: fix the issue that planner may cache invalid plans for joins in some cases (#28432)](https://github.com/pingcap/tidb/pull/28445)            |   | 15d23h  |
| tidb/28446 | [planner: fix the issue that planner may cache invalid plans for joins in some cases (#28432)](https://github.com/pingcap/tidb/pull/28446)            |   | 15d23h  |
| tidb/28491 | [util/ranger: check boundary condition when taking intersection of two columnValues](https://github.com/pingcap/tidb/pull/28491)                      |   | 13d20h  |
| tidb/28554 | [planner, statistics, sessionctx: add variable to enable/disable the outdated statistics to pseudo logic](https://github.com/pingcap/tidb/pull/28554) |   | 6d20h   |
| tidb/28790 | [planner: forbid constant fold when plan cache enable](https://github.com/pingcap/tidb/pull/28790)                                                    |   | 16h     |
| tidb/28799 | [domain, session: add plan replayer gc](https://github.com/pingcap/tidb/pull/28799)                                                                   |   | 1h      |


## Reviewed in Last 7 Days

|    REPO    |                                                PR                                                 | C | D |  R   |
|------------|---------------------------------------------------------------------------------------------------|---|---|------|
| tidb/28729 | [executor,distsql: fix analyze version 2 memory leak](https://github.com/pingcap/tidb/pull/28729) |   | 2 | 1d0h |


</details> 


<details> 
  <summary>winoros's detailed review report</summary> 

## To Be Reviewed

|     REPO     |                                                                          PR                                                                           | C | LASTED  |
|--------------|-------------------------------------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/20903   | [planner: fix confused and unnecessary double-projection in plans.](https://github.com/pingcap/tidb/pull/20903)                                       |   | 341d17h |
| docs-cn/5916 | [sql-statements, information-schema: add `END_TIME` field for table `ANALYZE_STATUS`](https://github.com/pingcap/docs-cn/pull/5916)                   |   | 195d17h |
| docs/5783    | [migration: Add information about Vitess to TiDB migration](https://github.com/pingcap/docs/pull/5783)                                                |   | 121d5h  |
| tidb/21018   | [planner: don't push down null sensitive join conditions (#19620)](https://github.com/pingcap/tidb/pull/21018)                                        |   | 335d17h |
| tidb/22416   | [core: fix subQuery at projection in only_full_group](https://github.com/pingcap/tidb/pull/22416)                                                     | Y | 270d11h |
| tidb/22478   | [planner, executor: fix query partition table with global unique index get wrong result](https://github.com/pingcap/tidb/pull/22478)                  |   | 265d13h |
| tidb/24138   | [planner: Add Equivalence Rules to Transform BinaryOptSubquery to ExistsSubquery](https://github.com/pingcap/tidb/pull/24138)                         |   | 177d12h |
| tidb/26323   | [planner: use multi-layer projections for subquery selection (#8190)](https://github.com/pingcap/tidb/pull/26323)                                     |   | 87d6h   |
| tidb/26474   | [planner: fix the unstable unit test TestTableFromMeta (#26463)](https://github.com/pingcap/tidb/pull/26474)                                          |   | 83d16h  |
| tidb/26475   | [planner: fix the unstable unit test TestTableFromMeta (#26463)](https://github.com/pingcap/tidb/pull/26475)                                          |   | 83d16h  |
| tidb/26492   | [planner: fix the unstable test TestOrderedResultModeOnOtherOperators (#26481)](https://github.com/pingcap/tidb/pull/26492)                           |   | 82d23h  |
| tidb/26503   | [planner: fix goroutine leak problem in some unit tests (#26500)](https://github.com/pingcap/tidb/pull/26503)                                         |   | 82d19h  |
| tidb/26505   | [planner: fix goroutine leak problem in some unit tests (#26500)](https://github.com/pingcap/tidb/pull/26505)                                         |   | 82d19h  |
| tidb/27636   | [planner, expression: avoid exprs with side effects in column pruning and agg pushdown (#27370)](https://github.com/pingcap/tidb/pull/27636)          |   | 47d17h  |
| tidb/27639   | [planner, expression: avoid exprs with side effects in column pruning and agg pushdown (#27370)](https://github.com/pingcap/tidb/pull/27639)          |   | 47d17h  |
| tidb/28295   | [planner: keep the original join schema in predicate pushdown (#24862)](https://github.com/pingcap/tidb/pull/28295)                                   |   | 20d16h  |
| tidb/28428   | [*: support show column_stats_usage](https://github.com/pingcap/tidb/pull/28428)                                                                      |   | 16d16h  |
| tidb/28491   | [util/ranger: check boundary condition when taking intersection of two columnValues](https://github.com/pingcap/tidb/pull/28491)                      |   | 13d20h  |
| tidb/28554   | [planner, statistics, sessionctx: add variable to enable/disable the outdated statistics to pseudo logic](https://github.com/pingcap/tidb/pull/28554) |   | 6d20h   |
| tidb/28558   | [statistics: migrate test-infra to testify for statistics_test.go](https://github.com/pingcap/tidb/pull/28558)                                        |   | 6d10h   |
| tidb/28719   | [statistics: fix auto analyze triggered out of specified time (#28703)](https://github.com/pingcap/tidb/pull/28719)                                   |   | 2d16h   |
| tidb/28721   | [statistics: fix auto analyze triggered out of specified time (#28703)](https://github.com/pingcap/tidb/pull/28721)                                   |   | 2d16h   |
| tidb/28723   | [statistics: fix auto analyze triggered out of specified time (#28703)](https://github.com/pingcap/tidb/pull/28723)                                   |   | 2d16h   |
| tidb/28744   | [planner: clone possible properties before saving them is unnecessary](https://github.com/pingcap/tidb/pull/28744)                                    |   | 1d23h   |
| tidb/28748   | [planner/cascades: fill group stats](https://github.com/pingcap/tidb/pull/28748)                                                                      |   | 1d22h   |
| tidb/28759   | [planner: make constant propagation more stringently](https://github.com/pingcap/tidb/pull/28759)                                                     |   | 1d17h   |


## Reviewed in Last 7 Days

|     REPO      |                                                                         PR                                                                         | C | D |    R    |
|---------------|----------------------------------------------------------------------------------------------------------------------------------------------------|---|---|---------|
| tidb/28725    | [statistics: fix auto analyze triggered out of specified time (#28703)](https://github.com/pingcap/tidb/pull/28725)                                |   | 2 | 1d0h    |
| tidb/27302    | [statistics: fix "data too long" error when dumping stats from table with new collation data (#27033)](https://github.com/pingcap/tidb/pull/27302) |   | 2 | 56d1h   |
| tidb/28758    | [statistics: fix auto analyze triggered out of specified time (#28703)](https://github.com/pingcap/tidb/pull/28758)                                |   | 2 | 1h      |
| tidb/28757    | [statistics: fix "data too long" error when dumping stats from table with new collation data (#27033)](https://github.com/pingcap/tidb/pull/28757) |   | 2 | 1h      |
| tidb/28550    | [statistics: migrate test-infra to testify for histogram_test.go](https://github.com/pingcap/tidb/pull/28550)                                      |   | 2 | 6d14h   |
| tidb/28729    | [executor,distsql: fix analyze version 2 memory leak](https://github.com/pingcap/tidb/pull/28729)                                                  |   | 3 | 0h      |
| tidb/28703    | [statistics: fix auto analyze triggered out of specified time](https://github.com/pingcap/tidb/pull/28703)                                         |   | 3 | 3h      |
| docs-cn/5561  | [Add sql optimization-related docs to toc](https://github.com/pingcap/docs-cn/pull/5561)                                                           |   | 3 | 230d20h |
| community/570 | [votes: add windtalker as tidb committer](https://github.com/pingcap/community/pull/570)                                                           |   | 5 | 10d4h   |
| community/571 | [votes: add LittleFall as tidb reviewer](https://github.com/pingcap/community/pull/571)                                                            |   | 5 | 10d3h   |
| tidb/28678    | [statistics: avoid mutate global variable](https://github.com/pingcap/tidb/pull/28678)                                                             |   | 5 | 1h      |
| tidb/28061    | [types: distinguish between string and non-string in execute statement arguments](https://github.com/pingcap/tidb/pull/28061)                      |   | 6 | 22d20h  |


</details> 


<details> 
  <summary>wshwsh12's detailed review report</summary> 

## To Be Reviewed

|    REPO    |                                                           PR                                                           | C | LASTED  |
|------------|------------------------------------------------------------------------------------------------------------------------|---|---------|
| tidb/21401 | [expression: incompatibility with MySQL for ADDTIME()](https://github.com/pingcap/tidb/pull/21401)                     |   | 317d11h |
| tidb/21887 | [types: support %X %V %W formats for STR_TO_DATE()](https://github.com/pingcap/tidb/pull/21887)                        |   | 298d11h |
| tidb/27837 | [planner: fix wrong plan caused by shallow copy schema columns (#27798)](https://github.com/pingcap/tidb/pull/27837)   |   | 37d16h  |
| tidb/27992 | [planner: add sub plan info of shuffleReceiver when query explain analyze](https://github.com/pingcap/tidb/pull/27992) |   | 30d15h  |
| tidb/28333 | [executor: fix detaching from GlobalTracker before executing select query](https://github.com/pingcap/tidb/pull/28333) |   | 18d15h  |


## Reviewed in Last 7 Days

| REPO | PR | C | D | R |
|------|----|---|---|---|


</details> 

