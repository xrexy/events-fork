/** @type {import('tailwindcss').Config} */
// eslint-disable-next-line no-undef
module.exports = {
  content: ["./src/**/*.vue"],
  darkMode: "class",
  theme: {
    extend: {
      colors: {
        brand: {
          100: "#e5d3e6",
          200: "#caa6cd",
          300: "#b07ab4",
          400: "#954d9b",
          500: "#7b2182",
          600: "#621a68",
          700: "#4a144e",
          800: "#310d34",
          900: "#19071a",
        },
      },
    },
  },
  plugins: [],
};
