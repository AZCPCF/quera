def do_intersect(a, b, c, d):
    
    if a > b:
        a, b = b, a
    if c > d:
        c, d = d, c
    
    return (a < c < b and (d > b or d < a)) or (c < a < d and (b > d or b < c))


n = int(input())
a, b = map(int, input().split())
c, d = map(int, input().split())


intersect = do_intersect(a, b, c, d)


if intersect:
    print(4)  
else:
    print(3)  
