def count_tilings(n):
    if n % 2 != 0:
        return 0  
    
    
    dp = [0] * (n + 1)
    dp[0] = 1
    dp[2] = 3
    
    
    for i in range(4, n + 1, 2):
        dp[i] = 4 * dp[i - 2] - dp[i - 4]
    
    return dp[n]

n = int(input())
result = count_tilings(n)
print(result * 2)  
