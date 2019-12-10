package main

import "flag"

import "time"

var Period = flag.Duration("period", 1*time.Second, "slepp period")
