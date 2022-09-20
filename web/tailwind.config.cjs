const fontFallback = [
  "ui-sans-serif",
  "system-ui",
  "-apple-system",
  "BlinkMacSystemFont",
  "Segoe UI",
  "Roboto",
  "Helvetica Neue",
  "Arial",
  "Noto Sans",
  "sans-serif",
  "Apple Color Emoji",
  "Segoe UI Emoji",
  "Segoe UI Symbol",
  "Noto Color Emoji",
];

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{ts,tsx}"],
  theme: {
    extend: {
      colors: {
        gray2: "#4f4f4f",
        gray3: "#828282",
        gray4: "#BDBDBD",
        "soft-blue": "#F6F8FB",
        "accent-blue": "#97BEF4",
        "primary-blue": "#2F80ED",
        light: "#fafafb",
      },
      fontFamily: {
        sans: ["Poppins", ...fontFallback],
        montserrat: ["Montserrat", ...fontFallback],
      },
    },
  },
  plugins: [],
};
