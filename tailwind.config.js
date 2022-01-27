// tailwind.config.js
module.exports = {
  purge: ["./src/**/*.{js,jsx,ts,tsx}", "./public/index.html"],
  theme: {
    extend: {
      colors: {
        Body: "#15202b",
        Primary: "#1da1f2",
        Pink: "#f91880",
        Green: "#00ba7c",
        DarkGray: "#657786",
        LightGray: "#AAB8C2",
        ExtraLightGray: "#E1E8ED",
        ExtraExtraLightGray: "#F5F8FA",
      },
    },
    fontFamily: {
      Roboto: ["Roboto, sans-serif"],
    },
  },
  variants: {
    extend: {},
  },
  plugins: [],
};
