/** @type {import('next').NextConfig} */
const nextConfig = {
  i18n: {
    defaultLocale: "pt-BR",
    locales: ["pt-BR", "en-US"],
  },
};

const withNextItl = require("next-intl/plugin")("./i18n.ts");

module.exports = withNextItl(nextConfig);
