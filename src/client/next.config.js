// @ts-check
/** @type {import('next').NextConfig} */
const nextConfig = {
  i18n: {
    defaultLocale: "pt-BR",
    locales: ["pt-BR", "en-US"],
  },
};

const createNextIntlPlugin = require('next-intl/plugin');
const withNextIntl = createNextIntlPlugin();



module.exports = withNextIntl(nextConfig);
