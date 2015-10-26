#Default

```
$ httpd-2.4.17 ./support/ab -n 10000 -c 200 http://127.0.0.1:8080/
This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
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
Server Hostname:        127.0.0.1
Server Port:            8080

Document Path:          /
Document Length:        7 bytes

Concurrency Level:      200
Time taken for tests:   25.454 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      1230000 bytes
HTML transferred:       70000 bytes
Requests per second:    392.87 [#/sec] (mean)
Time per request:       509.071 [ms] (mean)
Time per request:       2.545 [ms] (mean, across all concurrent requests)
Transfer rate:          47.19 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    4   2.0      4      11
Processing:   500  505   2.3    505     517
Waiting:      500  504   2.0    504     515
Total:        501  509   3.6    509     522

Percentage of the requests served within a certain time (ms)
  50%    509
  66%    510
  75%    511
  80%    512
  90%    513
  95%    515
  98%    516
  99%    517
 100%    522 (longest request)
```

# Limited
```
$ httpd-2.4.17 ./support/ab -n 10000 -c 200 http://127.0.0.1:8080/
This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
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
Server Hostname:        127.0.0.1
Server Port:            8080

Document Path:          /
Document Length:        7 bytes

Concurrency Level:      200
Time taken for tests:   50.592 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      1230000 bytes
HTML transferred:       70000 bytes
Requests per second:    197.66 [#/sec] (mean)
Time per request:       1011.830 [ms] (mean)
Time per request:       5.059 [ms] (mean, across all concurrent requests)
Transfer rate:          23.74 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    3   1.2      3      10
Processing:   504 1004  50.2   1009    1030
Waiting:      503 1002  50.2   1007    1029
Total:        513 1007  49.6   1011    1031

Percentage of the requests served within a certain time (ms)
  50%   1011
  66%   1013
  75%   1015
  80%   1016
  90%   1019
  95%   1024
  98%   1027
  99%   1028
 100%   1031 (longest request)
```
