def permutation_dance():
    
    n, s = map(int, input().strip().split())
    
    
    birds = input().strip().split()
    
    
    for _ in range(s):
        action = input().strip().split()
        if action[0] == "insert":
            bird_name = action[1]
            position = int(action[2])
            birds.insert(position, bird_name)
        elif action[0] == "depart":
            bird_name = action[1]
            birds.remove(bird_name)
        elif action[0] == "relocate":
            bird_name = action[1]
            displacement = int(action[2])
            idx = birds.index(bird_name)
            birds.pop(idx)
            new_position = idx + displacement
            birds.insert(new_position, bird_name)
    
    
    print(" ".join(birds))


permutation_dance()
