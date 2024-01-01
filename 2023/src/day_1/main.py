import re
from utils.functions import get_lines

# pylint: disable=missing-function-docstring


def part_one(data):
    total = 0
    for line in data:
        numbers = re.findall(r"\d", line)
        total += int(numbers[0] + numbers[-1])
    return total


def part_two(filename: str):
    numbers_as_words = {
        "one": "o1e",
        "two": "t2o",
        "three": "thr3e",
        "four": "fo4r",
        "five": "fi5e",
        "six": "s6x",
        "seven": "sev7n",
        "eight": "eig8t",
        "nine": "ni9e",
    }
    lines = get_lines(filename)

    pattern = "(" + "|".join(numbers_as_words.keys()) + "|"
    pattern += "|".join([str(x) for x in range(1, 10)]) + ")"

    total = 0
    for line in lines:
        for k, v in numbers_as_words.items():
            line = line.replace(k, v)
        nums = re.findall(pattern, line)
        total += int(nums[0] + nums[-1])
    return total


test_data = get_lines("test_data1.txt")
test_res = part_one(test_data)
print("Test result:", test_res)

real_data = get_lines("data.txt")
real_res = part_one(real_data)
print("Real result:", real_res)

print(part_two("data.txt"))
