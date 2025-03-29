def find_winner(N):
    
    dp = [0] * (N + 1)
    
    
    dp[0] = 0  
    
    
    for i in range(1, N + 1):
        
        dp[i] = 0  
        for move in range(1, 7):  
            if i >= move and dp[i - move] == 0:
                dp[i] = 1  
                break
    
    
    return "Pat" if dp[N] == 1 else "Mat"


N = int(input())
print(find_winner(N))
