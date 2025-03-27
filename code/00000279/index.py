from math import comb

a, x, n = map(int, input().split())
result = sum(comb(n, k) * x**k * a**(n-k) for k in range(n+1))

print(result)

