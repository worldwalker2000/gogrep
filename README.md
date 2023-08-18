# gogrep

This is a very simple grep command implemented in go. I did this as a project
to learn go but it turned out to be way to easy.

It uses goroutines to speed up grepping.

```shell
grep -rn "include" > /dev/null  0.34s user 0.03s system 99% cpu 0.373 total
~/code/gogrep/gogrep -rn "include" > /dev/null  0.10s user 0.13s system 356% cpu 0.067 total
```

It is actually faster (.34s vs .10).

This is not even close to a full grep as only the folling kinds of commands work

```shell
gogrep "pattern" file
gogrep -n "pattern" file
gogrep -rn "pattern"
```
