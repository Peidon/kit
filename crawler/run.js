import { closeStorage, saveProfile, allStockCodes} from "./storage.js";
import { readFile } from "node:fs/promises";

async function main() {
  const asx_code = await readFile("./mining_code.out", "utf8");
  const code_list = asx_code.split(",").filter((code) => code.trim() !== "");
  const existing_codes = await allStockCodes();

  // Process each code in sequence, waiting for each to complete before starting the next
  for (const code of code_list) {
    if (existing_codes.includes(code.trim().toUpperCase())) {
      console.log(`Profile for ${code.trim()} already exists`);
      continue;
    }
    try {
      const stockCode = code.trim().toUpperCase();
      const result = await saveProfile(stockCode);
      console.log(
        `Saved profile for ${result.company_name} (${result.stock_code})`,
      );
    } catch (error) {
      console.error(`Error saving profile for ${code.trim()}:`, error.message);
    }
  }
  closeStorage();
}

main();
