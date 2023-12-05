def get_first_and_last_int_as_int(string: str):
    ints = []
    for char in string:
        if char.isdigit():
            ints.append(char)
    if len(ints) == 0:
        return 0;
    elif len(ints) == 1:
        return int(str(ints[0]) + str(ints[0]))
    else:
        return int(str(ints[0]) + str(ints[len(ints) - 1]))

total = 0
for line in open("input.data", "r").readlines():
    total += get_first_and_last_int_as_int(line)
print('sum of all of the calibration values: ', total)
