
并发太高（CPU使用率高）的时候，在 RLock 里的 goroutine 可能会被随机抢占，导致调度时间长
https://xargin.com/a-rlock-story/
https://cloud.tencent.com/developer/article/1560331