import pg from "pg";
import { scrapeData } from "./head_office.js";

const { Pool } = pg;
const pool = new Pool({
  connectionString: process.env.DATABASE_URL,
});

/** Scrape a company profile and insert it into company_profile. */
export async function saveProfile(stockCode) {
  const url = `https://www.asx.com.au/markets/company/${stockCode}`;
  const data = await scrapeData(url);

  if (
    !data.company_name ||
    !data.address ||
    !data.telephone ||
    !data.web_link
  ) {
    throw new Error("Scraped profile is missing a required company field");
  }
  if (!Number.isFinite(data.longitude) || !Number.isFinite(data.latitude)) {
    throw new Error("Scraped profile has invalid office coordinates");
  }

  const query = `
    INSERT INTO company_profile (
      company_name,
      stock_code,
      web_link,
      office_address,
      office_telephone,
      office_location
    ) VALUES ($1, $2, $3, $4, $5,
      ST_SetSRID(ST_MakePoint($6, $7), 4326)::geography
    )
    RETURNING id, company_name, stock_code
  `;
  const values = [
    data.company_name,
    stockCode.toUpperCase(),
    data.web_link,
    data.address,
    data.telephone,
    data.longitude,
    data.latitude,
  ];

  const result = await pool.query(query, values);
  return result.rows[0];
}

export async function closeStorage() {
  await pool.end();
}
