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
Time taken for tests:   0.524 seconds
Complete requests:      10000
Failed requests:        9991
   (Connect: 0, Receive: 0, Length: 9991, Exceptions: 0)
Total transferred:      1478890 bytes
HTML transferred:       308890 bytes
Requests per second:    19075.41 [#/sec] (mean)
Time per request:       5.242 [ms] (mean)
Time per request:       0.052 [ms] (mean, across all concurrent requests)
Transfer rate:          2754.93 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    3   0.7      2       8
Processing:     1    3   0.9      3      10
Waiting:        0    3   0.7      2       9
Total:          3    5   1.3      5      13
WARNING: The median and mean for the initial connection time are not within a normal deviation
        These results are probably not that reliable.
WARNING: The median and mean for the waiting time are not within a normal deviation
        These results are probably not that reliable.

Percentage of the requests served within a certain time (ms)
  50%      5
  66%      5
  75%      5
  80%      6
  90%      6
  95%      7
  98%     11
  99%     12
 100%     13 (longest request)
 ```

 -> New request
 ```bash
 curl http://localhost:3000
You are the visitor number 9979%
```
== Adding extra 300ms delay
```bash
$ ab -n 10000 -c 100 http://localhost:3000/
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
Time taken for tests:   30.790 seconds
Complete requests:      10000
Failed requests:        9999
   (Connect: 0, Receive: 0, Length: 9999, Exceptions: 0)
Total transferred:      1479093 bytes
HTML transferred:       309093 bytes
Requests per second:    324.78 [#/sec] (mean)
Time per request:       307.896 [ms] (mean)
Time per request:       3.079 [ms] (mean, across all concurrent requests)
Transfer rate:          46.91 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    2   1.8      2      10
Processing:   300  303   2.3    302     318
Waiting:      300  302   2.1    302     318
Total:        300  305   3.5    304     321

Percentage of the requests served within a certain time (ms)
  50%    304
  66%    305
  75%    306
  80%    307
  90%    310
  95%    312
  98%    314
  99%    315
 100%    321 (longest request)
 ```
 -> New request
 ```bash
 curl http://localhost:3000
You are the visitor number 9967%
```

**Conclusion:**
- It must be 10000 requests but it's only 9967 requests. Concurrency issue.