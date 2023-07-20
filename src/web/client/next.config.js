/** @type {import('next').NextConfig} */
const nextConfig = {
  experimental: {
    serverActions: true,
  },

};

const withNextItl = require("next-intl/plugin")("./i18n.ts");

module.exports = withNextItl(nextConfig);
