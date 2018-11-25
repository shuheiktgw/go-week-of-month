go-week-of-month
====
[![GitHub release](http://img.shields.io/github/release/shuheiktgw/go-week-of-month.svg?style=flat-square)](https://github.com/shuheiktgw/go-week-of-month/releases/latest)
[![CircleCI](https://circleci.com/gh/shuheiktgw/go-week-of-month.svg?style=svg)](https://circleci.com/gh/shuheiktgw/go-week-of-month)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

`go-week-of-month` provides useful functions to get week of month.

## Usage

### `GetWeekOfMonth`
`GetWeekOfMonth` returns the number of the week of the month. Weeks start on Monday. For example, 2018/11/05 belongs to the second week of November so `GetWeekOfMonth` returns two.

``` go
package main

import (
	"os"
	"time"

	"github.com/shuheiktgw/go-week-of-month"
)

func main() {
    t := time.Now()
    
    # Returns the number of the week of the month
    goWeekOfMonth.GetWeekOfMonth(&t)
}
```

### `GetWeekOfMonthByDay`
`GetWeekOfMonthByDay` returns the number of the week of month by day. For example, 2018/11/05 is the first Monday in 2018/11, so `GetWeekOfMonthByDay` returns one.

``` go
package main

import (
	"os"
	"time"

	"github.com/shuheiktgw/go-week-of-month"
)

func main() {
    t := time.Now()
    
    # Returns the number of the week of the month by day
    goWeekOfMonth.GetWeekOfMonthByDay(&t)
}
```



## Install

``` bash
$ go get github.com/shuheiktgw/go-week-of-month
```

## Author
[Shuhei Kitagawa](https://github.com/shuheiktgw)