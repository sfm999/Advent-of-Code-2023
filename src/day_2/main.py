import re


def update_threshold(color, grab_list, max_values):
    if color in grab_list:
        index = grab_list.index(color)
        value = int(grab_list[index - 1])
        if max_values.get(color) < value:
            max_values[color] = value
    return max_values


def check_threshold(max_values):
    threshold = {"red": 12, "green": 13, "blue": 14}
    for key in threshold.keys():
        if threshold.get(key) < max_values.get(key):
            return False
    return True


def get_max_values(line):
    max_values = {"red": 0, "green": 0, "blue": 0}
    game = line.split(":")[1].split(";")
    for grab in game:
        split_grab = re.split(r'[ ,\n]', grab)
        max_values = update_threshold("blue", split_grab, max_values)
        max_values = update_threshold("red", split_grab, max_values)
        max_values = update_threshold("green", split_grab, max_values)
    return max_values


# Part 1
def part_a_solution(file):
    total = 0
    for line in file:
        game_id = int(line.split(":")[0].split(" ")[1])
        max_values = get_max_values(line)
        check = check_threshold(max_values)
        if check:
            total += game_id
    return total


# Part 2
def part_b_solution(file):
    total = 0
    for line in file:
        line_power = 1
        max_values = get_max_values(line)
        for item in max_values.values():
            line_power *= item
        total += line_power
    return total


if __name__ == "__main__":
    input_text = "data.txt"
    with open("data.txt", "r") as f:
        # print(part_a_solution(f))
        print(part_b_solution(f))