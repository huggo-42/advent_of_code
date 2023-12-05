numbers = {
    'zero': 0,
    'one': 1,
    'two': 2,
    'three': 3,
    'four' : 4,
    'five' : 5,
    'six' : 6,
    'seven' : 7,
    'eight' : 8,
    'nine' : 9,
    '0': 0,
    '1': 1,
    '2': 2,
    '3': 3,
    '4' : 4,
    '5' : 5,
    '6' : 6,
    '7' : 7,
    '8' : 8,
    '9' : 9

} 

def get_first_and_last_number_as_int(string: str) -> int:
    ints = []

    for num in numbers: 
        rfind = string.rfind(num)
        if rfind != -1:
            ints.append({"index": rfind, "value": numbers[num]})
        find = string.find(num)
        if find != -1:
            ints.append({"index": find, "value": numbers[num]})

    highest_idx = lowest_idx = ints[0]["index"]
    first_int = last_int = ints[0]["value"] 

    for i in ints:
        if i["index"] > highest_idx:
            highest_idx = i["index"]
            last_int = i["value"]
        if i["index"] < lowest_idx:
            lowest_idx = i["index"]
            first_int = i["value"]

    return int(str(first_int) + str(last_int))


total = 0
for line in open("input.data", "r").readlines():
    total += get_first_and_last_number_as_int(line)
print('Sum of all of the calibration values: ', total)
