# algorithm101

|nos|name|topic|state|desc|
|:-:|:-:|:-:|:-:|:-|
|122|two sum|dpointer|ac|-|
|88|merge sorted array|dpointer|ac|-|
|142|linked list cycle ii|dpointer|nac|1. `p1,p2=p1.Next,p2.Next.Next`<br/>`p2=head`<br/>`for p1!=p2 {`<br/>`p1,p2=p1.Next,p2.Next`<br/>`}`<br/>`return p1`<br/>2. use `for-break` to replace `do-while` in golang|
|76|minimum window substring|dpointer|nac|1. `map: need, have; set: needchars`<br/>2. `var ret, origlen = "", len(s)`<br/>3. 注意 ret 和 origlen 分离, `len(needchs) == 0 && len(s[l:r]) < origlen // 这里不是 len(ret) 否则 ret := "" 和 ret := s 例外`|
|633|sum of square number|dpointer|nac|二分|
|680|valid palindrome ii|dpointer|ac|1层递归|
|524|longest word in dictionary through deleting|dpointer|ac|-|

## math

gcd & lcm, 最大公约数和最小公倍数

leetcode, 1979 easy

steps:
```go
// gcd(a,b)
for {
    a, b = b%a, a
}

// gcd(a, b%a)
// if b%a == 0; return a;
for a != 0 {
    a, b = b%a, a
}
// a == 0 -> return b
return b
```

