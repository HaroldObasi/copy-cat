package bootstrap

const TAILWIND_CONTENT string = `
/** @type {import('tailwindcss').Config} */
export default {
content: ["./*.html"],
theme: {
	extend: {},
},
plugins: [],
};
`

const POST_CSS_CONTENT string = `
export default {
	plugins: {
		tailwindcss: {},
		autoprefixer: {},
	},
};
`

const STYLE_CSS_CONTENT string = `
@tailwind base;
@tailwind components;
@tailwind utilities;

@layer components {
  .custom-chip {
    @apply bg-gray-200 text-sm inline-flex hover:bg-gray-300 transition-colors duration-200 text-black  p-2 rounded-lg items-center gap-1;
  }

  .custom-chip-sm {
    @apply text-xs py-1 rounded-lg;
  }
}
`

const MAIN_JS_CONTENT string = `
import './style.css'
import { createIcons, icons } from "lucide";

createIcons({ icons });
`
