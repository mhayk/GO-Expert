= Apache Bench
```bash
ab -n 10000 -c 100 http://localhost:3000/
```
output:
```
This is ApacheBench, Version 2.3 <$Revision: 1903618 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            3000

Document Path:          /
Document Length:        28 bytes

Concurrency Level:      100
Time taken for tests:   30.774 seconds
Complete requests:      10000
Failed requests:        9999
   (Connect: 0, Receive: 0, Length: 9999, Exceptions: 0)
Total transferred:      1479196 bytes
HTML transferred:       309196 bytes
Requests per second:    324.95 [#/sec] (mean)
Time per request:       307.742 [ms] (mean)
Time per request:       3.077 [ms] (mean, across all concurrent requests)
Transfer rate:          46.94 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    2   1.7      1       9
Processing:   300  302   1.9    302     312
Waiting:      300  302   1.7    302     312
Total:        300  304   3.1    303     319

Percentage of the requests served within a certain time (ms)
  50%    303
  66%    305
  75%    306
  80%    306
  90%    308
  95%    311
  98%    313
  99%    314
 100%    319 (longest request)
 ```
 -> New request
 ```bash
 curl http://localhost:3000
You are the visitor number 10001%
```

**Conclusion:**
- It must be 10001 requests. It's 10001 requests. No concurrency issue.
