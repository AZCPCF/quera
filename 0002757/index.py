k = int(input())

ans = 1
p = 2

while p * p <= k:
    if k % p == 0:
        ans *= p
        while k % p == 0:
            k //= p
    p += 1

if k > 1:
    ans *= k

print(ans)