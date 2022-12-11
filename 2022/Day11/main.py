#!/usr/bin/env python

import time
import os
import operator
import math

operator_map = {
                "*" : operator.mul,
                "-" : operator.sub,
                "+" : operator.add
            }

class Monkey():

    """
    Monkey business class
    """

    def __init__(self, name : int, start_items : list, operation_op: operator, operation_int , test_criteria : int, true_cond : int, false_cond : int):
        self.name = name
        self.item_stash = self.process_start_items(start_items)
        self.observed_worry_level = 0
        self.test_criteria = test_criteria
        self.test_condition = lambda x : self.observed_worry_level % x
        self.true_condition = true_cond
        self.false_condition = false_cond
        self.operation_int = operation_int
        self.operation = lambda x, y : operation_op(x, y)
        self.nr_of_inspections = 0

    def process_start_items(self, items):
        stash = []
        for item in items:
            stash.append(int(item.replace(',', "")))
        return stash

    def doTest(self):
        self.observed_worry_level = math.floor(self.observed_worry_level / 3)
        if self.test_condition(self.test_criteria) == 0:
            return self.true_condition
        return self.false_condition

    def doOperation(self):
        if self.operation_int == 'old':
            self.observed_worry_level = self.operation(self.observed_worry_level, self.observed_worry_level)
        else:
            self.observed_worry_level = self.operation(self.observed_worry_level, int(self.operation_int))

    def throw_items(self):
        move_list = []

        for item in self.item_stash:
            self.nr_of_inspections += 1
            self.observed_worry_level = item
            self.doOperation()
            target_monkey = self.doTest()
            move_list.append((target_monkey, self.observed_worry_level))

        self.item_stash = []
        return move_list



def create_monkey(current_monkey):
    return Monkey(int(current_monkey[0][1].replace(':', "")), current_monkey[1][2:],
                operator_map[current_monkey[2][4]],
                current_monkey[2][-1],
                int(current_monkey[3][-1]),
                int(current_monkey[4][-1]),
                int(current_monkey[5][-1]))

def monkey_generator(file_name : str):
    with open(f"{os.getcwd()}/Day11/{file_name}") as f:
        lines = f.read().splitlines()
        monkey_map = {}
        current_monkey = []

        for line in lines:
            if line == '':
                m = create_monkey(current_monkey)
                monkey_map[m.name] = m
                current_monkey = []
                continue
            current_monkey.append(line.split())

        m = create_monkey(current_monkey)
        monkey_map[m.name] = m

    return monkey_map

def process_moves(move_list, monkey_list):
    for (target, item) in move_list:
        monkey_list[target].item_stash.append(item)

def process_monkey(monkey_map, i):
    monkey = monkey_map[i]
    move_list = monkey.throw_items()
    process_moves(move_list, monkey_map)

if __name__ == "__main__":
    st = time.time()
    monkey_map = monkey_generator("test.txt")
    nr_monkeys = len(monkey_map)
    counter = 0
    for i in range(20):
        for j in range(nr_monkeys):
            process_monkey(monkey_map, j)
            counter += 1

    result_list = []
    for k,v in monkey_map.items():
        # print(monkey_map[k].item_stash)
        result_list.append(monkey_map[k].nr_of_inspections)

    result_list.sort()
    # 107822
    print(result_list[-1] * result_list[-2])
    print('Execution time:', {time.time() - st}, 'seconds')
    os.EX_OK
