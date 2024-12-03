import re

def get_muls_sum(str):
  first_pattern = r"mul\((\d{1,3}),(\d{1,3})\)"
  muls = re.findall(first_pattern, str)
  sum = 0
  for mul in muls:
    sum += int(mul[0])*int(mul[1])
  return sum

def get_uncorrupted_muls_sum(input):
  total_sum = 0
  i = 0
  enabled_at = 0
  mul_enabled = True
  
  while i < len(input):
    if input[i:i+4] == "do()": 
      if not mul_enabled:
        enabled_at = i + 4
        
      mul_enabled = True
      i += 4  
    elif input[i:i+7] == "don't()":
      if mul_enabled:
        enabled_section = input[enabled_at: i]
        total_sum += get_muls_sum(enabled_section)
        
      mul_enabled = False
      i = i + 7
    else:
      i = i + 1
      
  last_enabled_section = input[enabled_at : len(input)]
  total_sum += get_muls_sum(last_enabled_section)
      
  return total_sum


def main():
  with open("./input.txt") as inputFile:
    input = inputFile.readlines()
    input = "".join(input)
    print(get_muls_sum(input))
    print(get_uncorrupted_muls_sum(input))

      
    
    
    
main()
