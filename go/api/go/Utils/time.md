```mermaid
graph TB
 A[查数据库 URL]-->B[给URL分类预先分类]
B-->C[Redis调度员]
C--到时间--> D{拿到id 查数据库}
D--更新 -->E{更新数据库时间和调度时间}
D--没更新 -->F{更新调度时间}
```