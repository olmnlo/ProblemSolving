def kangaroo():
    kangaros = input()
    values = kangaros.split(" ")
    x1 = int(values[0])
    v1 = int(values[1])
    x2 = int(values[2])
    v2 = int(values[3])

    if v1 == v2:
        return "YES" if x1 == x2 else "NO"
    elif (x2 - x1) * (v1 - v2) < 0:
        return "NO"
    elif (x2 - x1) % (v1 - v2) == 0:
        return "YES"
    else:
        return "NO"

print(kangaroo())