// await import("./src/env.mjs");

/** @type {import("next").NextConfig} */
const config = {
  reactStrictMode: true,

  i18n: {
    locales: ["en", "pt_BR"],
    defaultLocale: "en",
  },
};
export default config;
