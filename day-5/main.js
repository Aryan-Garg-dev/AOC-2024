import { promises as fs } from "fs";

async function parseInput(fileName){
  const parsedData = {
    rules: [
      { first: 0, second: 0 }
    ],
    updates: [[]]
  }
  try {
    const data = await fs.readFile(fileName, "utf-8");
    const [rulesInput, updatesInput] = data.split("\r\n\r\n")
    const rulesArr = rulesInput.split("\r\n");
    const updatesArr = updatesInput.split("\r\n");
    parsedData.rules.splice(0, 1);
    for (let rule of rulesArr){
      rule = rule.split("|");
      parsedData.rules.push({
        first: Number(rule[0]), 
        second: Number(rule[1])
      });
    }
    parsedData.updates.splice(0, 1);
    for (let update of updatesArr){
      update = update.split(",").map(val=>Number(val))
      parsedData.updates.push(update)
    }
    return parsedData;
  } catch(err){
    console.log(err);
    throw err;
  }
}

const input = await parseInput("input.txt");
const { rules, updates } = input;

const validateUpdate = (update)=>{
  for (const rule of rules){
    const i = update.indexOf(rule.first);
    const j = update.indexOf(rule.second);
    if ((i != -1 && j != -1) && i > j) return false;
  }
  return true;
}

const rectifyUpdate = (update)=>{
  for (const rule of rules){
    const i = update.indexOf(rule.first);
    const j = update.indexOf(rule.second);
    if ((i != -1 && j != -1) && i > j){
      const temp = update[i];
      update[i] = update[j];
      update[j] = temp;
    };
  }
}

let sum = 0;
let correctedSum = 0;
for (const update of updates){
  if (validateUpdate(update)){
    sum += (update[Math.floor(update.length / 2)]);
  } else {
    while (!validateUpdate(update)) rectifyUpdate(update);
    correctedSum += (update[Math.floor(update.length / 2)]);
  }
}
console.log(sum);
console.log(correctedSum);
