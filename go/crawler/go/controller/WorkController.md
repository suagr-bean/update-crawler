```mermaid
graph TD
A[workstart time]-->B[workcontroller]
B-->C{条件函数 time}
C--time满足-->D[执行 workcontroller]
C--time 不满足-->E[阻塞]

```