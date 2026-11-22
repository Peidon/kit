import pg from "pg";

const { Pool } = pg;
const pool = new Pool({
  connectionString: process.env.DATABASE_URL,
});

/** insert company profile into company_profile. */
export async function saveProfile(data) {
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

  try {
    const result = await pool.query(query, values);
    return result.rows[0];
  } catch (error) {
    if (error.code === "23505") {
      throw new Error(
        `Profile for stock code ${stockCode} already exists in the database`,
      );
    }
    throw error;
  }
}

export async function allStockCodes() {
  const query = "SELECT stock_code FROM company_profile";
  const result = await pool.query(query);
  return result.rows.map((row) => row.stock_code);
}

export async function closeStorage() {
  await pool.end();
}
