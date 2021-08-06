# golang-fun
Set of boring tasks in golang


### How to run
```shell
$ go run GoFun
For given slice :       [1 2 3 4 5 6 7 8 9 10]
        Sum of evens:   30
        Sum of odds:    25

Fibonacci for 35:       9227465

1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, FizzBuzz, 16, 17, Fizz, 19, Buzz

Palindrome for 'rota':  'rotator'

array [1 2 2] has duplicates

array [1 2 3] has not duplicates

======================Contact List======================================

=============Adding contact johh===============

created object John: {johh Doe}
=============Adding contact Mister=============

created object Twister: {Mister Twister}
=============Updating contact johh to John=============

updated object John: {Jonh Doe}
=============Printing all=============

All objects: map[10:{Mister Twister} 74:{Jonh Doe}]
=============Deleting John=============

John deleted: 74
=============Printing all again=============

All objects: map[10:{Mister Twister}]
===================== Same CRUD for Tasks List======================================

======================Task List======================================

=============Adding contact johh===============

created task to Sleep well: {Sleep well High Not started}
=============Adding contact Mister=============

created task to Code: {Code more Very High Ongoing}
=============Updating task Sleep Well to just sleep=============

updated task Sleep: {Sleep Middle Not started}
=============Printing all=============

All tasks: map[21:{Sleep Middle Not started} 25:{Code more Very High Ongoing}]
=============Deleting task to Sleep =============

Forget to sleep: 21
=============Printing all again=============

All tasks: map[25:{Code more Very High Ongoing}]
```

### How to test
```shell
$ go test
PASS
ok      GoFun   0.516s

```
### How to check code coverage
```shell
$ go test -cover
PASS
coverage: 80.3% of statements
ok      GoFun   0.550s

```



### How to have a fun
Don't do this...
```shell
$ shutdown -h 0
```