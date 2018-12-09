with open('input.txt') as f:
    input_ = f.read()

jumps = [int(i) for i in input_.split('\n') if i]
endpoint = len(jumps)

index = 0
iterations = 0
while 0 <= index < endpoint:
    jump = jumps[index]
    if jump > 2:
        jumps[index] -= 1
    else:
        jumps[index] += 1
    index += jump
    iterations += 1

print(iterations)
