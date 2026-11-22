import puppeteer from "puppeteer";
import { load } from "cheerio";

export class OfficeProfile {
  company_name;
  address;
  telephone;
  web_link;
  longitude;
  latitude;

  constructor(company_name, address, telephone, web_link, location) {
    this.company_name = company_name;
    this.address = address;
    this.telephone = telephone;
    this.web_link = web_link;
    this.longitude = Number(location.lon);
    this.latitude = Number(location.lat);
  }
}

export async function scrapeData(url, rough = false) {
  const browser = await puppeteer.launch();
  let renderedHtml;
  try {
    const page = await browser.newPage();
    await page.goto(url, {
      waitUntil: "networkidle2",
      timeout: 60_000,
    });
    await page.waitForSelector("table.table-contact", { timeout: 30_000 });
    renderedHtml = await page.content();
  } finally {
    await browser.close();
  }

  const $ = load(renderedHtml);
  const table = $("table.table-contact")
    .filter(
      (_index, element) =>
        $(element).find("caption").text().trim() ===
        "Head office contact information",
    )
    .first();

  if (table.length === 0) {
    throw new Error("Head office contact information table was not found");
  }

  const rows = table.find("tr");
  const address = $(rows[0]).find("td").text().trim();
  const telephone = $(rows[1]).find("td").text().trim();
  const webLink =
    $(rows[rows.length - 1])
      .find("td a")
      .attr("href")
      ?.trim() ?? "";

  const company_name = $("h1.panel-title")?.find("span")?.first().text().trim();
  if (company_name.length === 0) {
    throw new Error("Company Name was not found");
  }

  const location = await fetchCoordinates(address, rough);

  return new OfficeProfile(company_name, address, telephone, webLink, location);
}

export async function fetchCoordinates(address, rough = false) {
  const response = await fetch(
    `https://nominatim.openstreetmap.org/search?format=jsonv2&limit=1&q=${encode(address, rough)}`,
    {
      headers: {
        Accept: "application/json",
        "User-Agent": "xupedon@qq.com",
      },
    },
  );

  const body = await response.text();

  if (!response.ok) {
    throw new Error(
      `Nominatim request failed: ${response.status} ${response.statusText}. ${body}`,
    );
  }

  const results = JSON.parse(body);
  if (results.length === 0) {
    throw new Error(`No coordinates found for: ${address}`);
  }
  return results[0];
}

export function encode(address, rough = false) {
  const phrases = address.split(",");
  if (phrases.length > 5) {
    address = phrases
      .slice(phrases.length - 5, phrases.length)
      .join(",")
      .trim();
  }
  if (rough) {
    // remove the number of fist phrase, e.g. "123 Main St" becomes "Main St"
    address = phrases
      .map((phrase, index) => (index === 0 ? phrase.replace(/^\d+\s+/, "") : phrase))
      .join(",")
      .trim();
  }
  return encodeURIComponent(address);
}
