import puppeteer from 'puppeteer';
import { load } from 'cheerio';

const browser = await puppeteer.launch();
const page = await browser.newPage();

const bc8Url = "https://www.asx.com.au/markets/company/VAL";

export class OfficeProfile {
  company_name;
  address;
  telephone;
  web_link;
  coordinates;

  constructor(company_name, address, telephone, web_link, coordinates) {
    this.company_name = company_name;
    this.address = address;
    this.telephone = telephone;
    this.web_link = web_link;
    this.coordinates = coordinates;
  }
}

export async function scrapeData(url = bc8Url) {
  // Navigate to the ASX company page and wait for network requests to settle
  await page.goto(url, {
    waitUntil: 'networkidle2',
  });

  // Wait specifically for the dynamic contact table to appear
  await page.waitForSelector('table.table-contact');

  const renderedHtml = await page.content();
  await browser.close();

  const $ = load(renderedHtml);
  const table = $("table.table-contact").filter(
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

  const location = await fetch(
    `https://nominatim.openstreetmap.org/search?format=jsonv2&limit=1&q=${encodeURIComponent(address)}`,
  );

  const [result] = await location.json();

  const coordinates = [
    Number(result.lon),
    Number(result.lat),
  ];

  return new OfficeProfile(company_name, address, telephone, webLink, coordinates);
}
