import puppeteer from 'puppeteer';
import {load} from 'cheerio';

const browser = await puppeteer.launch();
const page = await browser.newPage();

const bc8Url = "https://www.asx.com.au/markets/company/BC8";

export class OfficeProfile {
  address;
  telephone;
  web_link;

  constructor(address, telephone, web_link) {
    this.address = address;
    this.telephone = telephone;
    this.web_link = web_link;
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

  // const response = await axios.get(url);
  const $ = load(renderedHtml);
  const tables = $("table");
  const table = tables.filter(
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

  return new OfficeProfile(address, telephone, webLink);
}
