const path = require('path');

module.exports = {
  content: ["./src/**/*.{html,jsx}"],
  theme: {
    extend: {
      spacing: {
        '40': '5rem',
      },
      backgroundImage: {
        'custom-background': `url('${path.resolve(__dirname, './src/assets/Background.jpg')}')`,
        'cyber': `url('${path.resolve(__dirname, './src/assets/Cyber.jpg')}')`
      },
      colors: {
        'gradient-start': '#BAB0D4',
        'gradient-end': '#4C1A57',
      },
    },
  },
  plugins: [
    require('tailwindcss-gradients'),
  ],
}

