import { saveProfile } from "./storage.js";
import { readFile } from "node:fs/promises";

async function main() {
  const asx_code = await readFile("./mining_code.out", "utf8");
  const code_list = asx_code.split(",").filter((code) => code.trim() !== "");
  code_list.forEach(async (code) => {
    try {
      const result = await saveProfile(code.trim());
      console.log(
        `Saved profile for ${result.company_name} (${result.stock_code})`,
      );
    } catch (error) {
      console.error(`Error saving profile for ${code.trim()}:`, error.message);
    }
  });
}

main();
