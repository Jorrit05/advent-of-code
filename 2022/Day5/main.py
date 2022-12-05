from collections import deque

def init_deques(file_name : str):
    with open(file_name) as f:
        for line in f.readlines():
            print(line.split('['))
            print(len(line.split('[')))
            if line == "\n":
                break

if __name__ == "__main__":
    init_deques("test.txt")
    # queue = deque([])
    # print(queue)