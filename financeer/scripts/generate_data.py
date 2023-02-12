#!/usr/bin/env python3
import csv
import random
import sys

CATEGORY = ["Dining", "Living", "Leisure"]

num_rows = int(sys.argv[1])
file_path = '../test_data/generated_data.csv'
print("Writing", num_rows, " rows to", file_path)

f = open(file_path, 'w')
writer = csv.writer(f)

for i in range(num_rows):
    random_category = random.choice(CATEGORY)
    random_price = random.randrange(1, 500001)
    expense_record = ["sample_data", random_category, random_price]
    writer.writerow(expense_record)
    pass
