def get_lines(filename: str):
    with open(file=filename, mode="r", encoding="utf-8") as fr:
        data = [x.strip() for x in fr.readlines()]
        fr.close()
    return data
